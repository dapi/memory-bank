package lint

import (
	"fmt"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"
	"unicode/utf8"
)

var ignoredDirectories = map[string]bool{
	".git": true, ".hg": true, ".svn": true, ".venv": true,
	"node_modules": true, "tmp": true, "log": true, "vendor": true,
}

func NormalizeScopeRoot(scopeRoot string) (string, error) {
	normalized := path.Clean(strings.TrimSpace(scopeRoot))
	if normalized == "" || normalized == "." {
		return "", fmt.Errorf("--scope-root must point to a repository-relative directory")
	}
	return strings.TrimSuffix(normalized, "/"), nil
}

func Run(options Options) (Report, error) {
	scopeRoot, err := NormalizeScopeRoot(options.ScopeRoot)
	if err != nil {
		return Report{}, err
	}
	repoRoot, err := filepath.Abs(options.RepoRoot)
	if err != nil {
		return Report{}, err
	}
	documents, err := loadDocuments(repoRoot)
	if err != nil {
		return Report{}, err
	}
	entrypoints, missingEntrypoints := deriveEntrypoints(documents, scopeRoot, options.Entrypoints)
	return buildReport(repoRoot, scopeRoot, entrypoints, missingEntrypoints, options.MaxDepth, documents), nil
}

func loadDocuments(repoRoot string) (map[string]document, error) {
	documents := make(map[string]document)
	err := filepath.WalkDir(repoRoot, func(fullPath string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.IsDir() {
			if fullPath != repoRoot && ignoredDirectories[entry.Name()] {
				return filepath.SkipDir
			}
			return nil
		}
		if !strings.HasSuffix(entry.Name(), ".md") {
			return nil
		}
		contents, readErr := os.ReadFile(fullPath)
		if readErr != nil {
			return readErr
		}
		relativePath, relativeErr := filepath.Rel(repoRoot, fullPath)
		if relativeErr != nil {
			return relativeErr
		}
		text := string(contents)
		documents[filepath.ToSlash(relativePath)] = document{text: text, frontmatter: parseFrontmatter(text)}
		return nil
	})
	return documents, err
}

func isScopedMarkdown(documentPath, scopeRoot string) bool {
	return strings.HasPrefix(documentPath, scopeRoot+"/") && strings.HasSuffix(documentPath, ".md")
}

func isScopedReadme(documentPath, scopeRoot string) bool {
	return isScopedMarkdown(documentPath, scopeRoot) && path.Base(documentPath) == "README.md"
}

func resolveEntrypointPath(entrypoint, scopeRoot string, knownPaths map[string]bool) (string, string, bool) {
	primaryCandidate, primaryOK := normalizeCLIDocumentPath(entrypoint)
	scopedInput := path.Join(scopeRoot, strings.TrimLeft(entrypoint, "/"))
	scopedCandidate, scopedOK := normalizeCLIDocumentPath(scopedInput)
	normalizedInput := strings.Trim(strings.TrimSpace(entrypoint), "<>")
	explicitRepoRoot := strings.HasPrefix(normalizedInput, "/") || strings.HasPrefix(normalizedInput, "./")
	alreadyScoped := primaryOK && (primaryCandidate == scopeRoot || strings.HasPrefix(primaryCandidate, scopeRoot+"/"))

	if !explicitRepoRoot && !alreadyScoped && scopedOK && knownPaths[scopedCandidate] {
		return scopedCandidate, entrypoint, true
	}
	if primaryOK && knownPaths[primaryCandidate] {
		return primaryCandidate, entrypoint, true
	}
	if scopedOK && knownPaths[scopedCandidate] {
		return scopedCandidate, entrypoint, true
	}
	fallback := entrypoint
	if primaryOK {
		fallback = primaryCandidate
	} else if scopedOK {
		fallback = scopedCandidate
	}
	return "", fallback, false
}

func deriveEntrypoints(documents map[string]document, scopeRoot string, configuredEntrypoints []string) ([]string, []string) {
	knownPaths := make(map[string]bool, len(documents))
	for documentPath := range documents {
		knownPaths[documentPath] = true
	}
	resolved := []string{}
	missing := []string{}
	if len(configuredEntrypoints) > 0 {
		seen := make(map[string]bool)
		for _, entrypoint := range configuredEntrypoints {
			resolvedPath, missingHint, ok := resolveEntrypointPath(entrypoint, scopeRoot, knownPaths)
			if !ok {
				missing = append(missing, missingHint)
				continue
			}
			if !seen[resolvedPath] {
				resolved = append(resolved, resolvedPath)
				seen[resolvedPath] = true
			}
		}
		return resolved, missing
	}

	defaultEntrypoint := scopeRoot + "/README.md"
	if knownPaths[defaultEntrypoint] {
		return []string{defaultEntrypoint}, []string{}
	}
	return []string{}, []string{defaultEntrypoint}
}

func extractDerivedFromPaths(frontmatter map[string]any) []string {
	rawValue, ok := frontmatter["derived_from"]
	if !ok {
		return nil
	}
	values, isList := rawValue.([]any)
	if !isList {
		values = []any{rawValue}
	}
	paths := []string{}
	for _, value := range values {
		switch typedValue := value.(type) {
		case string:
			if typedValue != "" {
				paths = append(paths, typedValue)
			}
		case map[string]string:
			if dependencyPath := typedValue["path"]; dependencyPath != "" {
				paths = append(paths, dependencyPath)
			}
		}
	}
	return paths
}

func validateFrontmatterDependencies(documents map[string]document, scopeRoot string) []FrontmatterDependency {
	issues := []FrontmatterDependency{}
	for _, sourcePath := range sortedDocumentPaths(documents) {
		if !isScopedMarkdown(sourcePath, scopeRoot) {
			continue
		}
		for _, rawPath := range extractDerivedFromPaths(documents[sourcePath].frontmatter) {
			target, ok := normalizeInternalMarkdownTarget(sourcePath, rawPath)
			if ok {
				if _, exists := documents[target]; !exists {
					issues = append(issues, FrontmatterDependency{
						Source: sourcePath, Field: "derived_from", Value: rawPath, Target: target,
					})
				}
			}
		}
	}
	return issues
}

func buildLinkGraph(documents map[string]document, scopeRoot string) (map[string]map[string]bool, map[string]map[string]bool, map[string]map[string]bool) {
	outgoing := make(map[string]map[string]bool)
	incomingInScope := make(map[string]map[string]bool)
	brokenLinks := make(map[string]map[string]bool)
	for _, sourcePath := range sortedDocumentPaths(documents) {
		for _, target := range extractInternalMarkdownLinks(sourcePath, documents[sourcePath].text) {
			if _, exists := documents[target]; exists {
				addToSet(outgoing, sourcePath, target)
				if sourcePath != target && isScopedMarkdown(sourcePath, scopeRoot) && isScopedMarkdown(target, scopeRoot) {
					addToSet(incomingInScope, target, sourcePath)
				}
			} else if isScopedMarkdown(sourcePath, scopeRoot) {
				addToSet(brokenLinks, sourcePath, target)
			}
		}
	}
	return outgoing, incomingInScope, brokenLinks
}

func deriveIndexPaths(documents map[string]document, scopeRoot string) []string {
	indexPaths := []string{}
	for _, documentPath := range sortedDocumentPaths(documents) {
		if isScopedMarkdown(documentPath, scopeRoot) && stringFrontmatterValue(documents[documentPath].frontmatter, "doc_function") == "index" {
			indexPaths = append(indexPaths, documentPath)
		}
	}
	return indexPaths
}

func deriveExpectedReadmeIndices(documents map[string]document, scopeRoot string) []string {
	readmes := []string{}
	for _, documentPath := range sortedDocumentPaths(documents) {
		if !isScopedReadme(documentPath, scopeRoot) {
			continue
		}
		frontmatter := documents[documentPath].frontmatter
		if stringFrontmatterValue(frontmatter, "doc_function") == "template" {
			continue
		}
		if stringFrontmatterValue(frontmatter, "doc_kind") == "feature-support" && stringFrontmatterValue(frontmatter, "doc_function") == "reference" {
			continue
		}
		readmes = append(readmes, documentPath)
	}
	return readmes
}

type childAnnotation struct {
	target     string
	annotation string
}

func annotationTextForChildLinks(indexPath, text string) []childAnnotation {
	sectionPrefix := path.Dir(indexPath)
	if sectionPrefix == "." {
		sectionPrefix = ""
	}
	lines := strings.Split(stripFencedCodeBlocks(text), "\n")
	annotations := []childAnnotation{}
	for indexLine := 0; indexLine < len(lines); indexLine++ {
		line := lines[indexLine]
		rawDestination, found := firstBulletLinkDestination(line)
		if !found {
			continue
		}
		target, ok := normalizeInternalMarkdownTarget(indexPath, rawDestination)
		if !ok {
			continue
		}
		childPrefix := sectionPrefix + "/"
		if sectionPrefix != "" && !strings.HasPrefix(target, childPrefix) {
			continue
		}

		fragments := []string{}
		inlineAnnotation := strings.Trim(removeMarkdownLinks(line), " -\t:")
		if inlineAnnotation != "" {
			fragments = append(fragments, inlineAnnotation)
		}
		for continuationIndex := indexLine + 1; continuationIndex < len(lines); continuationIndex++ {
			continuation := lines[continuationIndex]
			if strings.TrimSpace(continuation) == "" {
				break
			}
			if strings.HasPrefix(continuation, "  ") || strings.HasPrefix(continuation, "\t") {
				fragments = append(fragments, strings.TrimSpace(continuation))
				continue
			}
			break
		}
		annotations = append(annotations, childAnnotation{target: target, annotation: strings.TrimSpace(strings.Join(fragments, " "))})
	}
	return annotations
}

func validateIndexDocument(indexPath string, documents map[string]document) []string {
	document, exists := documents[indexPath]
	if !exists {
		return []string{"missing expected index file"}
	}
	issues := []string{}
	if len(document.frontmatter) == 0 {
		issues = append(issues, "missing YAML frontmatter")
	}
	purpose, purposeOK := document.frontmatter["purpose"].(string)
	if !purposeOK || strings.TrimSpace(purpose) == "" {
		issues = append(issues, "missing 'purpose' in frontmatter")
	}
	if stringFrontmatterValue(document.frontmatter, "doc_function") != "index" {
		issues = append(issues, "expected `doc_function: index`")
	}
	for _, child := range annotationTextForChildLinks(indexPath, document.text) {
		normalizedAnnotation := strings.ToLower(strings.Trim(whitespaceRE.ReplaceAllString(child.annotation, " "), " -:\t"))
		basename := strings.ToLower(path.Base(child.target))
		basenameWithoutExtension := strings.TrimSuffix(basename, path.Ext(basename))
		switch {
		case normalizedAnnotation == "":
			issues = append(issues, "missing annotation for child link -> "+child.target)
		case normalizedAnnotation == basename || normalizedAnnotation == basenameWithoutExtension:
			issues = append(issues, "annotation repeats filename for child link -> "+child.target)
		case utf8.RuneCountInString(normalizedAnnotation) < 12:
			issues = append(issues, "annotation too short for child link -> "+child.target)
		}
	}
	return issues
}

func expectedParentIndex(documentPath string, indexPaths map[string]bool, scopeRoot string) *string {
	if !isScopedMarkdown(documentPath, scopeRoot) {
		return nil
	}
	currentDirectory := path.Dir(documentPath)
	if path.Base(documentPath) == "README.md" {
		currentDirectory = path.Dir(currentDirectory)
	}
	for currentDirectory != "" && currentDirectory != "." {
		candidate := path.Join(currentDirectory, "README.md")
		if indexPaths[candidate] && candidate != documentPath {
			return stringPointer(candidate)
		}
		parentDirectory := path.Dir(currentDirectory)
		if parentDirectory == currentDirectory {
			break
		}
		currentDirectory = parentDirectory
	}
	return nil
}

func buildNavigationReachability(outgoing map[string]map[string]bool, navigationNodes map[string]bool, entrypoints []string) map[string]reachability {
	reachable := make(map[string]reachability)
	navigationDepths := make(map[string]int)
	queue := []string{}
	for _, entrypoint := range entrypoints {
		reachable[entrypoint] = reachability{depth: 0, route: []string{entrypoint}}
		navigationDepths[entrypoint] = 0
		queue = append(queue, entrypoint)
	}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		currentInfo := reachable[current]
		currentDepth := navigationDepths[current]
		for _, target := range sortedSet(outgoing[current]) {
			candidateDepth := currentDepth + 1
			candidateRoute := append(append([]string{}, currentInfo.route...), target)
			best, hasBest := reachable[target]
			if !hasBest || candidateDepth < best.depth {
				reachable[target] = reachability{depth: candidateDepth, route: candidateRoute}
			}
			if navigationNodes[target] {
				bestDepth, hasBestDepth := navigationDepths[target]
				if !hasBestDepth || candidateDepth < bestDepth {
					navigationDepths[target] = candidateDepth
					queue = append(queue, target)
				}
			}
		}
	}
	return reachable
}

func flattenBrokenLinks(brokenLinks map[string]map[string]bool) []BrokenLink {
	items := []BrokenLink{}
	for _, sourcePath := range sortedOuterSet(brokenLinks) {
		for _, target := range sortedSet(brokenLinks[sourcePath]) {
			items = append(items, BrokenLink{Source: sourcePath, Target: target})
		}
	}
	return items
}

func buildReport(repoRoot, scopeRoot string, entrypoints, missingEntrypoints []string, maxDepth int, documents map[string]document) Report {
	scopedMarkdownPaths := []string{}
	for _, documentPath := range sortedDocumentPaths(documents) {
		if isScopedMarkdown(documentPath, scopeRoot) {
			scopedMarkdownPaths = append(scopedMarkdownPaths, documentPath)
		}
	}
	indexPaths := deriveIndexPaths(documents, scopeRoot)
	expectedReadmeIndices := deriveExpectedReadmeIndices(documents, scopeRoot)
	outgoing, incomingInScope, brokenLinks := buildLinkGraph(documents, scopeRoot)

	report := Report{
		FormatVersion:      1,
		RepoRoot:           repoRoot,
		ScopeRoot:          scopeRoot,
		Entrypoints:        append([]string{}, entrypoints...),
		MissingEntrypoints: append([]string{}, missingEntrypoints...),
		MaxDepth:           maxDepth,
		Stats:              Stats{MarkdownFilesInScope: len(scopedMarkdownPaths), IndexDocumentsInScope: len(indexPaths)},
		Errors: Errors{
			Config: []ConfigError{}, BrokenLinks: flattenBrokenLinks(brokenLinks),
			FrontmatterDependencies: validateFrontmatterDependencies(documents, scopeRoot),
			Orphans:                 []NavigationIssue{}, Unreachable: []NavigationIssue{}, IndexContract: []IndexContractIssue{},
		},
		Warnings: Warnings{DeepReachable: []DeepReachableWarning{}},
	}
	if len(missingEntrypoints) > 0 {
		report.Errors.Config = append(report.Errors.Config, ConfigError{Message: "Configured entrypoints are missing.", Paths: append([]string{}, missingEntrypoints...)})
	}
	if len(scopedMarkdownPaths) == 0 {
		report.Errors.Config = append(report.Errors.Config, ConfigError{Message: "Scope contains no markdown files.", Paths: []string{scopeRoot}})
	}
	if len(entrypoints) == 0 {
		paths := append([]string{}, missingEntrypoints...)
		if len(paths) == 0 {
			paths = []string{scopeRoot + "/README.md"}
		}
		report.Errors.Config = append(report.Errors.Config, ConfigError{Message: "No valid entrypoints were resolved.", Paths: paths})
	}

	indexPathSet := boolSet(indexPaths)
	entrypointSet := boolSet(entrypoints)
	if len(entrypoints) > 0 && len(scopedMarkdownPaths) > 0 {
		navigationNodes := boolSet(indexPaths)
		for entrypoint := range entrypointSet {
			navigationNodes[entrypoint] = true
		}
		reachable := buildNavigationReachability(outgoing, navigationNodes, entrypoints)
		for _, documentPath := range scopedMarkdownPaths {
			inboundSources := sortedSet(incomingInScope[documentPath])
			parentIndex := expectedParentIndex(documentPath, indexPathSet, scopeRoot)
			if !entrypointSet[documentPath] && len(inboundSources) == 0 {
				report.Errors.Orphans = append(report.Errors.Orphans, NavigationIssue{
					Path: documentPath, ExpectedParentIndex: parentIndex, InboundLinks: inboundSources,
				})
			}
			reachabilityInfo, isReachable := reachable[documentPath]
			if !isReachable {
				report.Errors.Unreachable = append(report.Errors.Unreachable, NavigationIssue{
					Path: documentPath, ExpectedParentIndex: parentIndex, InboundLinks: inboundSources,
				})
				continue
			}
			if reachabilityInfo.depth > maxDepth {
				report.Warnings.DeepReachable = append(report.Warnings.DeepReachable, DeepReachableWarning{
					Path: documentPath, Depth: reachabilityInfo.depth, MaxDepth: maxDepth,
					ExpectedParentIndex: parentIndex, Route: append([]string{}, reachabilityInfo.route...),
				})
			}
		}
	}

	for _, indexPath := range expectedReadmeIndices {
		issues := validateIndexDocument(indexPath, documents)
		if len(issues) > 0 {
			report.Errors.IndexContract = append(report.Errors.IndexContract, IndexContractIssue{
				Path: indexPath, Issues: issues, ExpectedParentIndex: expectedParentIndex(indexPath, indexPathSet, scopeRoot),
			})
		}
	}
	sort.Slice(report.Warnings.DeepReachable, func(left, right int) bool {
		if report.Warnings.DeepReachable[left].Depth == report.Warnings.DeepReachable[right].Depth {
			return report.Warnings.DeepReachable[left].Path < report.Warnings.DeepReachable[right].Path
		}
		return report.Warnings.DeepReachable[left].Depth < report.Warnings.DeepReachable[right].Depth
	})

	report.Stats.BrokenLinkCount = len(report.Errors.BrokenLinks)
	report.Stats.FrontmatterDependencyIssueCount = len(report.Errors.FrontmatterDependencies)
	report.Stats.OrphanCount = len(report.Errors.Orphans)
	report.Stats.UnreachableCount = len(report.Errors.Unreachable)
	report.Stats.IndexContractIssueCount = len(report.Errors.IndexContract)
	report.Stats.DeepReachableWarningCount = len(report.Warnings.DeepReachable)
	report.Stats.EntrypointCount = len(entrypoints)
	if len(report.Errors.Config) > 0 || len(report.Errors.BrokenLinks) > 0 || len(report.Errors.FrontmatterDependencies) > 0 ||
		len(report.Errors.Orphans) > 0 || len(report.Errors.Unreachable) > 0 || len(report.Errors.IndexContract) > 0 {
		report.ExitCode = 1
	}
	return report
}

func addToSet(sets map[string]map[string]bool, key, value string) {
	if sets[key] == nil {
		sets[key] = make(map[string]bool)
	}
	sets[key][value] = true
}

func boolSet(values []string) map[string]bool {
	set := make(map[string]bool, len(values))
	for _, value := range values {
		set[value] = true
	}
	return set
}

func sortedDocumentPaths(documents map[string]document) []string {
	paths := make([]string, 0, len(documents))
	for documentPath := range documents {
		paths = append(paths, documentPath)
	}
	sort.Strings(paths)
	return paths
}

func sortedOuterSet(sets map[string]map[string]bool) []string {
	keys := make([]string, 0, len(sets))
	for key := range sets {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}

func sortedSet(set map[string]bool) []string {
	values := make([]string, 0, len(set))
	for value := range set {
		values = append(values, value)
	}
	sort.Strings(values)
	return values
}

func stringFrontmatterValue(frontmatter map[string]any, key string) string {
	value, _ := frontmatter[key].(string)
	return value
}

func stringPointer(value string) *string {
	return &value
}

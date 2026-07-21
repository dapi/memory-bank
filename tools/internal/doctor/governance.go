package doctor

import (
	"bytes"
	"fmt"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"gopkg.in/yaml.v3"
)

type governedDocument struct {
	path        string
	frontmatter map[string]any
	content     string
}

var designNotRequired = regexp.MustCompile(`(?im)^\s*(?:[-*]\s+|\d+\.\s+|\|\s*)?design\s+required\s*:\s*no\b`)

func (report *Report) checkGovernance(scopeRoot string) {
	documents := map[string]governedDocument{}
	root := filepath.Join(report.RepoRoot, filepath.FromSlash(scopeRoot))
	err := filepath.WalkDir(root, func(fullPath string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.IsDir() || !strings.HasSuffix(strings.ToLower(entry.Name()), ".md") {
			return nil
		}
		relative, err := filepath.Rel(report.RepoRoot, fullPath)
		if err != nil {
			return err
		}
		documentPath := filepath.ToSlash(relative)
		if entry.Type()&os.ModeSymlink != 0 {
			report.add(Finding{Code: "governance.unsafe_symlink", Severity: Error, Group: "frontmatter_governance", Path: documentPath, Message: "Governed document is a symlink.", Remediation: "Replace it with a regular repository-owned Markdown file."})
			return nil
		}
		data, err := os.ReadFile(fullPath)
		if err != nil {
			return err
		}
		frontmatter, found, err := parseFrontmatter(data)
		if err != nil {
			report.add(Finding{Code: "governance.frontmatter_invalid", Severity: Error, Group: "frontmatter_governance", Path: documentPath, Message: err.Error(), Remediation: "Fix the YAML frontmatter according to memory-bank/dna/frontmatter.md."})
			return nil
		}
		if !found {
			report.add(Finding{Code: "governance.frontmatter_missing", Severity: Error, Group: "frontmatter_governance", Path: documentPath, Message: "Governed Markdown document has no YAML frontmatter.", Remediation: "Add frontmatter with at least a valid status field."})
			return nil
		}
		documents[documentPath] = governedDocument{path: documentPath, frontmatter: frontmatter, content: string(data)}
		validateGovernedDocument(report, documents[documentPath])
		return nil
	})
	if err != nil {
		report.add(Finding{Code: "governance.scope_unreadable", Severity: Error, Group: "frontmatter_governance", Path: scopeRoot, Message: err.Error(), Remediation: "Restore a readable memory-bank documentation tree."})
		return
	}
	report.checkDerivedFromCycles(documents)
	report.checkFeatureLifecycle(documents, scopeRoot)
}

func parseFrontmatter(data []byte) (map[string]any, bool, error) {
	// YAML permits CRLF line endings. Normalize them before recognizing the
	// Markdown delimiters so governed documents work consistently across
	// platforms.
	data = bytes.ReplaceAll(data, []byte("\r\n"), []byte("\n"))
	if !bytes.HasPrefix(data, []byte("---\n")) {
		return nil, false, nil
	}
	remainder := data[4:]
	end := -1
	for offset := 0; offset < len(remainder); {
		candidate := bytes.Index(remainder[offset:], []byte("\n---"))
		if candidate < 0 {
			break
		}
		candidate += offset
		// A delimiter must occupy its entire line. Without this check a value
		// such as "---not-a-delimiter" silently closes the frontmatter.
		afterDelimiter := candidate + len("\n---")
		if afterDelimiter == len(remainder) || remainder[afterDelimiter] == '\n' {
			end = candidate
			break
		}
		offset = afterDelimiter
	}
	if end < 0 {
		return nil, true, fmt.Errorf("unterminated YAML frontmatter")
	}
	frontmatter := map[string]any{}
	decoder := yaml.NewDecoder(bytes.NewReader(remainder[:end]))
	if err := decoder.Decode(&frontmatter); err != nil {
		return nil, true, fmt.Errorf("invalid YAML frontmatter: %w", err)
	}
	return frontmatter, true, nil
}

func validateGovernedDocument(report *Report, document governedDocument) {
	status, ok := document.frontmatter["status"].(string)
	if !ok || !oneOf(status, "draft", "active", "archived") {
		report.add(Finding{Code: "governance.status_invalid", Severity: Error, Group: "frontmatter_governance", Path: document.path, Subject: fmt.Sprint(document.frontmatter["status"]), Message: "status is missing or outside the governed enum.", Remediation: "Set status to draft, active, or archived."})
	}
	if delivery, exists := document.frontmatter["delivery_status"]; exists {
		value, valid := delivery.(string)
		if !valid || !oneOf(value, "planned", "in_progress", "done", "cancelled") {
			report.add(Finding{Code: "governance.delivery_status_invalid", Severity: Error, Group: "frontmatter_governance", Path: document.path, Subject: fmt.Sprint(delivery), Message: "delivery_status is outside the governed enum.", Remediation: "Use planned, in_progress, done, or cancelled."})
		}
		if path.Base(document.path) != "brief.md" {
			report.add(Finding{Code: "lifecycle.delivery_status_wrong_owner", Severity: Error, Group: "lifecycle_consistency", Path: document.path, Message: "delivery_status is owned only by a canonical brief.md.", Remediation: "Move lifecycle state to the package brief.md and remove the duplicate field."})
		}
	}
	if decision, exists := document.frontmatter["decision_status"]; exists {
		value, valid := decision.(string)
		if !valid || !oneOf(value, "proposed", "accepted", "superseded", "rejected") {
			report.add(Finding{Code: "governance.decision_status_invalid", Severity: Error, Group: "frontmatter_governance", Path: document.path, Subject: fmt.Sprint(decision), Message: "decision_status is outside the governed enum.", Remediation: "Use proposed, accepted, superseded, or rejected."})
		}
		docKind, _ := document.frontmatter["doc_kind"].(string)
		if docKind != "adr" {
			report.add(Finding{Code: "lifecycle.decision_status_wrong_owner", Severity: Error, Group: "lifecycle_consistency", Path: document.path, Message: "decision_status is owned only by ADR documents.", Remediation: "Move decision lifecycle state to an ADR and remove the field here."})
		}
	}
	if _, exists := document.frontmatter["derived_from"]; exists && len(derivedFromTargets(document)) == 0 {
		report.add(Finding{Code: "governance.derived_from_invalid", Severity: Error, Group: "frontmatter_governance", Path: document.path, Message: "derived_from is present but contains no usable path.", Remediation: "Use a path string or an object with a non-empty path field."})
	}
	if isADR(document.path) {
		if _, exists := document.frontmatter["decision_status"]; !exists {
			report.add(Finding{Code: "lifecycle.adr_decision_status_missing", Severity: Error, Group: "lifecycle_consistency", Path: document.path, Message: "Instantiated ADR does not declare decision_status.", Remediation: "Add the ADR decision lifecycle state."})
		}
	}
}

func isADR(documentPath string) bool {
	base := path.Base(documentPath)
	return strings.Contains(documentPath, "/adr/") && strings.HasPrefix(base, "ADR-") && base != "ADR-XXX.md"
}

func oneOf(value string, values ...string) bool {
	for _, candidate := range values {
		if value == candidate {
			return true
		}
	}
	return false
}

func derivedFromTargets(document governedDocument) []string {
	raw, exists := document.frontmatter["derived_from"]
	if !exists {
		return nil
	}
	values, ok := raw.([]any)
	if !ok {
		values = []any{raw}
	}
	targets := []string{}
	for _, item := range values {
		var rawPath string
		switch value := item.(type) {
		case string:
			rawPath = value
		case map[string]any:
			rawPath, _ = value["path"].(string)
		}
		if rawPath == "" {
			continue
		}
		if strings.HasPrefix(rawPath, "/") {
			targets = append(targets, path.Clean(strings.TrimPrefix(rawPath, "/")))
		} else {
			targets = append(targets, path.Clean(path.Join(path.Dir(document.path), rawPath)))
		}
	}
	return targets
}

func (report *Report) checkDerivedFromCycles(documents map[string]governedDocument) {
	state := map[string]int{}
	stack := []string{}
	seenCycles := map[string]bool{}
	var visit func(string)
	visit = func(current string) {
		state[current] = 1
		stack = append(stack, current)
		for _, target := range derivedFromTargets(documents[current]) {
			if _, exists := documents[target]; !exists {
				continue
			}
			if state[target] == 0 {
				visit(target)
			} else if state[target] == 1 {
				start := 0
				for stack[start] != target {
					start++
				}
				cycle := append(append([]string{}, stack[start:]...), target)
				key := strings.Join(cycle, " -> ")
				if !seenCycles[key] {
					seenCycles[key] = true
					report.add(Finding{Code: "governance.derived_from_cycle", Severity: Error, Group: "frontmatter_governance", Path: current, Subject: key, Message: "derived_from contains a dependency cycle.", Remediation: "Restore an acyclic upstream-to-downstream dependency graph."})
				}
			}
		}
		stack = stack[:len(stack)-1]
		state[current] = 2
	}
	paths := make([]string, 0, len(documents))
	for documentPath := range documents {
		paths = append(paths, documentPath)
	}
	sort.Strings(paths)
	for _, documentPath := range paths {
		if state[documentPath] == 0 {
			visit(documentPath)
		}
	}
}

func (report *Report) checkFeatureLifecycle(documents map[string]governedDocument, scopeRoot string) {
	prefix := strings.TrimSuffix(scopeRoot, "/") + "/features/"
	packages := map[string]map[string]governedDocument{}
	for documentPath, document := range documents {
		if !strings.HasPrefix(documentPath, prefix) {
			continue
		}
		relative := strings.TrimPrefix(documentPath, prefix)
		parts := strings.Split(relative, "/")
		if len(parts) < 2 || !strings.HasPrefix(parts[0], "FT-") {
			continue
		}
		if packages[parts[0]] == nil {
			packages[parts[0]] = map[string]governedDocument{}
		}
		packages[parts[0]][strings.Join(parts[1:], "/")] = document
	}
	for packageName, files := range packages {
		brief, hasBrief := files["brief.md"]
		_, hasDesign := files["design.md"]
		plan, hasPlan := files["implementation-plan.md"]
		packagePath := path.Join(prefix, packageName)
		if !hasBrief {
			report.add(Finding{Code: "lifecycle.feature_brief_missing", Severity: Error, Group: "lifecycle_consistency", Path: packagePath, Message: "Feature package contains stage artifacts without canonical brief.md.", Remediation: "Create brief.md from the governed feature template before later-stage artifacts."})
			continue
		}
		delivery, _ := brief.frontmatter["delivery_status"].(string)
		if delivery == "" {
			report.add(Finding{Code: "lifecycle.delivery_status_missing", Severity: Error, Group: "lifecycle_consistency", Path: brief.path, Message: "Canonical feature brief does not own delivery_status.", Remediation: "Add the package lifecycle state to brief.md."})
		}
		if hasPlan {
			briefStatus, _ := brief.frontmatter["status"].(string)
			if briefStatus != "active" {
				report.add(Finding{Code: "lifecycle.plan_brief_not_active", Severity: Error, Group: "lifecycle_consistency", Path: brief.path, Message: "An implementation plan requires an active feature brief.", Remediation: "Complete the Problem Ready gate and set brief.md status to active before creating the implementation plan."})
			}
		}
		if hasPlan && !hasDesign && !designNotRequired.MatchString(brief.content) {
			report.add(Finding{Code: "lifecycle.plan_without_design", Severity: Warning, Group: "lifecycle_consistency", Path: plan.path, Message: "Implementation plan exists without a design stage artifact.", Remediation: "Document why design is not required, or add design.md before execution planning."})
		}
		if hasPlan && hasDesign && !designNotRequired.MatchString(brief.content) {
			design := files["design.md"]
			designStatus, _ := design.frontmatter["status"].(string)
			if designStatus != "active" {
				report.add(Finding{Code: "lifecycle.plan_design_not_active", Severity: Error, Group: "lifecycle_consistency", Path: design.path, Message: "An implementation plan requires an active design when design is required.", Remediation: "Complete the Solution Ready gate and set design.md status to active before creating the implementation plan."})
			}
		}
		planStatus, _ := plan.frontmatter["status"].(string)
		if delivery == "in_progress" && (!hasPlan || planStatus != "active") {
			report.add(Finding{Code: "lifecycle.execution_plan_not_active", Severity: Error, Group: "lifecycle_consistency", Path: packagePath, Message: "An in-progress feature requires an active implementation-plan.md.", Remediation: "Create and activate the implementation plan before execution."})
		}
		if delivery == "done" && (!hasPlan || planStatus != "archived") {
			report.add(Finding{Code: "lifecycle.done_plan_not_archived", Severity: Error, Group: "lifecycle_consistency", Path: packagePath, Message: "A done feature requires an archived implementation-plan.md.", Remediation: "Complete the Done gate and archive the implementation plan."})
		}
		if delivery == "cancelled" && hasPlan && planStatus != "archived" {
			report.add(Finding{Code: "lifecycle.cancelled_plan_not_archived", Severity: Error, Group: "lifecycle_consistency", Path: packagePath, Message: "A cancelled feature requires implementation-plan.md to be absent or archived.", Remediation: "Archive the implementation plan, or remove it if it was never used."})
		}
	}
}

func (report *Report) addNavigationFindings() {
	for _, item := range report.Navigation.Errors.Config {
		report.add(Finding{Code: "navigation.config", Severity: Error, Group: "navigation", Subject: strings.Join(item.Paths, ","), Message: item.Message, Remediation: "Fix the lint scope or configured entrypoints."})
	}
	for _, item := range report.Navigation.Errors.BrokenLinks {
		report.add(Finding{Code: "navigation.broken_link", Severity: Error, Group: "navigation", Path: item.Source, Subject: item.Target, Message: "Internal Markdown link target does not exist.", Remediation: "Fix or remove the relative link."})
	}
	for _, item := range report.Navigation.Errors.FrontmatterDependencies {
		report.add(Finding{Code: "navigation.broken_derived_from", Severity: Error, Group: "navigation", Path: item.Source, Subject: item.Target, Message: "derived_from target does not exist.", Remediation: "Fix the upstream dependency path."})
	}
	for _, item := range report.Navigation.Errors.Orphans {
		report.add(Finding{Code: "navigation.orphan", Severity: Error, Group: "navigation", Path: item.Path, Message: "Document has no in-scope inbound link.", Remediation: "Add it to the appropriate README index or remove it."})
	}
	for _, item := range report.Navigation.Errors.Unreachable {
		report.add(Finding{Code: "navigation.unreachable", Severity: Error, Group: "navigation", Path: item.Path, Message: "Document is not reachable through index navigation.", Remediation: "Connect it to the entrypoint through governed README indices."})
	}
	for _, item := range report.Navigation.Errors.IndexContract {
		report.add(Finding{Code: "navigation.index_contract", Severity: Error, Group: "navigation", Path: item.Path, Subject: strings.Join(item.Issues, "; "), Message: "README index violates the navigation contract.", Remediation: "Fix its frontmatter and annotated child links."})
	}
	for _, item := range report.Navigation.Warnings.DeepReachable {
		report.add(Finding{Code: "navigation.deep_reachable", Severity: Warning, Group: "navigation", Path: item.Path, Subject: fmt.Sprintf("depth=%d", item.Depth), Message: "Document is reachable only beyond the configured depth.", Remediation: "Add a shorter index route or increase --max-depth deliberately."})
	}
}

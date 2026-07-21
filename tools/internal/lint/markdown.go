package lint

import (
	"path"
	"regexp"
	"strings"
)

var (
	fencedCodeBlockRE = regexp.MustCompile("(?s)```.*?```")
	markdownLinkRE    = regexp.MustCompile(`\[([^\]]+)\]\(([^)]+)\)`)
	frontmatterRE     = regexp.MustCompile(`(?s)^---\n(.*?)\n---(?:\n|$)`)
	bulletPrefixRE    = regexp.MustCompile(`^\s*-\s+`)
	whitespaceRE      = regexp.MustCompile(`\s+`)
)

func stripFencedCodeBlocks(text string) string {
	return fencedCodeBlockRE.ReplaceAllString(text, "")
}

func extractMarkdownLinkDestination(rawURL string) string {
	url := strings.TrimSpace(rawURL)
	if url == "" {
		return ""
	}
	if strings.HasPrefix(url, "<") {
		if closingIndex := strings.Index(url, ">"); closingIndex >= 0 {
			return strings.TrimSpace(url[1:closingIndex])
		}
	}
	if fields := strings.Fields(url); len(fields) > 0 {
		url = fields[0]
	}
	return strings.Trim(url, "<>")
}

func stripFrontmatterValue(value string) string {
	return strings.Trim(strings.TrimSpace(value), "\"'")
}

func parseFrontmatterListItem(item string) any {
	if strings.HasPrefix(item, "{") && strings.HasSuffix(item, "}") {
		fields := make(map[string]string)
		for _, part := range strings.Split(strings.Trim(item, "{}"), ",") {
			key, value, ok := strings.Cut(part, ":")
			if !ok {
				continue
			}
			fields[strings.TrimSpace(key)] = stripFrontmatterValue(value)
		}
		return fields
	}
	if strings.HasPrefix(item, "path:") {
		return map[string]string{"path": stripFrontmatterValue(strings.TrimPrefix(item, "path:"))}
	}
	return stripFrontmatterValue(item)
}

func parseFrontmatter(text string) map[string]any {
	match := frontmatterRE.FindStringSubmatch(text)
	if match == nil {
		return map[string]any{}
	}

	frontmatter := make(map[string]any)
	currentKey := ""
	for _, line := range strings.Split(match[1], "\n") {
		if line == "" {
			continue
		}
		strippedLine := strings.TrimSpace(line)
		indented := strings.HasPrefix(line, " ") || strings.HasPrefix(line, "\t")
		if indented && strings.HasPrefix(strippedLine, "- ") && currentKey != "" {
			currentValue, ok := frontmatter[currentKey].([]any)
			if !ok {
				currentValue = []any{}
			}
			currentValue = append(currentValue, parseFrontmatterListItem(strings.TrimSpace(strings.TrimPrefix(strippedLine, "- "))))
			frontmatter[currentKey] = currentValue
			continue
		}
		if indented || strings.HasPrefix(line, "-") || !strings.Contains(line, ":") {
			continue
		}

		key, value, _ := strings.Cut(line, ":")
		currentKey = strings.TrimSpace(key)
		strippedValue := strings.TrimSpace(value)
		if strippedValue == "" {
			frontmatter[currentKey] = ""
		} else {
			frontmatter[currentKey] = stripFrontmatterValue(strippedValue)
		}
	}
	return frontmatter
}

func normalizeInternalMarkdownTarget(sourcePath, rawURL string) (string, bool) {
	url := extractMarkdownLinkDestination(rawURL)
	if url == "" || strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://") ||
		strings.HasPrefix(url, "mailto:") || strings.HasPrefix(url, "#") {
		return "", false
	}
	if before, _, found := strings.Cut(url, "#"); found {
		url = before
	}
	if before, _, found := strings.Cut(url, "?"); found {
		url = before
	}
	url = strings.TrimSpace(url)
	if url == "" {
		return "", false
	}

	extension := strings.ToLower(path.Ext(url))
	if extension != "" && extension != ".md" {
		return "", false
	}

	var resolved string
	if strings.HasPrefix(url, "/") {
		resolved = path.Clean(strings.TrimLeft(url, "/"))
	} else {
		resolved = path.Clean(path.Join(path.Dir(sourcePath), url))
	}
	if extension == "" {
		resolved = path.Join(resolved, "README.md")
	}
	return resolved, true
}

func normalizeCLIDocumentPath(rawPath string) (string, bool) {
	documentPath := strings.Trim(strings.TrimSpace(rawPath), "<>")
	if documentPath == "" {
		return "", false
	}
	if before, _, found := strings.Cut(documentPath, "#"); found {
		documentPath = before
	}
	if before, _, found := strings.Cut(documentPath, "?"); found {
		documentPath = before
	}
	documentPath = strings.TrimSpace(documentPath)
	if documentPath == "" {
		return "", false
	}

	extension := strings.ToLower(path.Ext(documentPath))
	if extension != "" && extension != ".md" {
		return "", false
	}
	normalized := path.Clean(strings.TrimLeft(documentPath, "/"))
	if normalized == "" || normalized == "." {
		return "", false
	}
	if extension == "" {
		normalized = path.Join(normalized, "README.md")
	}
	return normalized, true
}

func extractInternalMarkdownLinks(sourcePath, text string) []string {
	strippedText := stripFencedCodeBlocks(text)
	matches := markdownLinkRE.FindAllStringSubmatchIndex(strippedText, -1)
	links := make([]string, 0, len(matches))
	for _, match := range matches {
		if isImageLink(strippedText, match[0]) {
			continue
		}
		target, ok := normalizeInternalMarkdownTarget(sourcePath, strippedText[match[4]:match[5]])
		if ok {
			links = append(links, target)
		}
	}
	return links
}

func removeMarkdownLinks(text string) string {
	matches := markdownLinkRE.FindAllStringSubmatchIndex(text, -1)
	var result strings.Builder
	lastEnd := 0
	for _, match := range matches {
		if isImageLink(text, match[0]) {
			continue
		}
		result.WriteString(text[lastEnd:match[0]])
		lastEnd = match[1]
	}
	result.WriteString(text[lastEnd:])
	return result.String()
}

func firstBulletLinkDestination(line string) (string, bool) {
	if !bulletPrefixRE.MatchString(line) {
		return "", false
	}
	for _, match := range markdownLinkRE.FindAllStringSubmatchIndex(line, -1) {
		if !isImageLink(line, match[0]) {
			return line[match[4]:match[5]], true
		}
	}
	return "", false
}

func isImageLink(text string, matchStart int) bool {
	return matchStart > 0 && text[matchStart-1] == '!'
}

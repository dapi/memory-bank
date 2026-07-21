// Package agentinstructions defines the managed Memory Bank routing block.
package agentinstructions

import (
	"bytes"
	"fmt"
	"strings"
)

const (
	DefaultTarget = "AGENTS.md"
	StartMarker   = "<!-- MEMORY BANK START -->"
	EndMarker     = "<!-- MEMORY BANK END -->"
	BlockVersion  = 1
)

var CurrentBlock = []byte(fmt.Sprintf(`%s
<!-- MEMORY BANK MANAGED BLOCK VERSION: %d -->
Before substantial delivery work, read memory-bank/README.md, memory-bank/dna/README.md, and memory-bank/flows/routing.md.
Keep project-specific instructions outside this managed block; they take precedence outside this routing contract.
%s
`, StartMarker, BlockVersion, EndMarker))

type Status string

const (
	Missing   Status = "missing"
	Current   Status = "current"
	Outdated  Status = "outdated"
	Ambiguous Status = "ambiguous"
)

type Plan struct {
	Status Status
	Data   []byte
	Diff   string
}

type markerLine struct {
	start int
	end   int
}

// BuildPlan returns a safe whole-file replacement while preserving every byte
// outside the exact managed markers. A newly appended block is separated from
// existing content by one blank line; no existing newline is rewritten.
func BuildPlan(original []byte) Plan {
	rawStarts := bytes.Count(original, []byte(StartMarker))
	rawEnds := bytes.Count(original, []byte(EndMarker))
	upper := bytes.ToUpper(original)
	markerLikeStarts := bytes.Count(upper, []byte("<!-- MEMORY BANK START"))
	markerLikeEnds := bytes.Count(upper, []byte("<!-- MEMORY BANK END"))
	starts, ends := standaloneMarkers(original)
	if markerLikeStarts != rawStarts || markerLikeEnds != rawEnds || rawStarts != len(starts) || rawEnds != len(ends) {
		return Plan{Status: Ambiguous}
	}
	if len(starts) == 0 && len(ends) == 0 {
		data := append([]byte(nil), original...)
		if len(data) > 0 {
			if data[len(data)-1] != '\n' {
				data = append(data, '\n')
			}
			data = append(data, '\n')
		}
		data = append(data, CurrentBlock...)
		return Plan{Status: Missing, Data: data, Diff: blockDiff(nil, CurrentBlock)}
	}
	if len(starts) != 1 || len(ends) != 1 {
		return Plan{Status: Ambiguous}
	}
	if ends[0].start < starts[0].start {
		return Plan{Status: Ambiguous}
	}
	start := starts[0].start
	end := ends[0].end
	existing := original[start:end]
	if bytes.Equal(existing, CurrentBlock) {
		return Plan{Status: Current, Data: append([]byte(nil), original...)}
	}
	data := make([]byte, 0, len(original)-len(existing)+len(CurrentBlock))
	data = append(data, original[:start]...)
	data = append(data, CurrentBlock...)
	data = append(data, original[end:]...)
	return Plan{Status: Outdated, Data: data, Diff: blockDiff(existing, CurrentBlock)}
}

// standaloneMarkers recognizes ownership boundaries only when the marker is
// the complete logical line. The returned end includes CRLF/LF when present so
// replacement cannot leave part of the managed line ending behind.
func standaloneMarkers(data []byte) (starts, ends []markerLine) {
	for lineStart := 0; lineStart < len(data); {
		newline := bytes.IndexByte(data[lineStart:], '\n')
		lineEnd := len(data)
		next := len(data)
		if newline >= 0 {
			lineEnd = lineStart + newline
			next = lineEnd + 1
		}
		logicalEnd := lineEnd
		if logicalEnd > lineStart && data[logicalEnd-1] == '\r' {
			logicalEnd--
		}
		line := data[lineStart:logicalEnd]
		switch {
		case bytes.Equal(line, []byte(StartMarker)):
			starts = append(starts, markerLine{start: lineStart, end: next})
		case bytes.Equal(line, []byte(EndMarker)):
			ends = append(ends, markerLine{start: lineStart, end: next})
		}
		lineStart = next
	}
	return starts, ends
}

func blockDiff(old, current []byte) string {
	var result strings.Builder
	result.WriteString("--- managed block (current)\n+++ managed block (planned)\n")
	for _, line := range lines(old) {
		result.WriteString("-")
		result.WriteString(line)
		result.WriteByte('\n')
	}
	for _, line := range lines(current) {
		result.WriteString("+")
		result.WriteString(line)
		result.WriteByte('\n')
	}
	return result.String()
}

func lines(data []byte) []string {
	trimmed := strings.TrimSuffix(string(data), "\n")
	if trimmed == "" {
		return nil
	}
	return strings.Split(trimmed, "\n")
}

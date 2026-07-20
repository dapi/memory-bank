package lint

import (
	"fmt"
	"io"
	"strings"
)

func PrintTextReport(writer io.Writer, report Report) {
	fmt.Fprintln(writer, "Memory-bank link audit")
	fmt.Fprintf(writer, "Repo root: %s\n", report.RepoRoot)
	fmt.Fprintf(writer, "Scope root: %s\n", report.ScopeRoot)
	entrypoints := strings.Join(report.Entrypoints, ", ")
	if entrypoints == "" {
		entrypoints = "(none)"
	}
	fmt.Fprintf(writer, "Entrypoints: %s\n", entrypoints)
	fmt.Fprintf(writer, "Navigation depth threshold: %d\n", report.MaxDepth)
	fmt.Fprintf(writer, "Markdown files in scope: %d\n", report.Stats.MarkdownFilesInScope)
	fmt.Fprintf(writer, "Index documents in scope: %d\n\n", report.Stats.IndexDocumentsInScope)

	if len(report.Errors.Config) > 0 {
		fmt.Fprintln(writer, "Configuration errors:")
		for _, item := range report.Errors.Config {
			fmt.Fprintf(writer, "  - %s\n", item.Message)
			for _, itemPath := range item.Paths {
				fmt.Fprintf(writer, "    * %s\n", itemPath)
			}
		}
		fmt.Fprintln(writer)
	}

	if len(report.Errors.BrokenLinks) > 0 {
		fmt.Fprintln(writer, "Broken internal markdown links:")
		for _, item := range report.Errors.BrokenLinks {
			fmt.Fprintf(writer, "  - %s -> %s\n", item.Source, item.Target)
		}
		fmt.Fprintln(writer)
	} else {
		fmt.Fprintln(writer, "OK: no broken internal markdown links in scope.")
		fmt.Fprintln(writer)
	}

	if len(report.Errors.FrontmatterDependencies) > 0 {
		fmt.Fprintln(writer, "Broken frontmatter markdown dependencies:")
		for _, item := range report.Errors.FrontmatterDependencies {
			fmt.Fprintf(writer, "  - %s %s: %s -> %s\n", item.Source, item.Field, item.Value, item.Target)
		}
		fmt.Fprintln(writer)
	} else {
		fmt.Fprintln(writer, "OK: no broken frontmatter markdown dependencies in scope.")
		fmt.Fprintln(writer)
	}

	if len(report.Errors.Orphans) > 0 {
		fmt.Fprintln(writer, "Orphan markdown files in scope:")
		for _, item := range report.Errors.Orphans {
			fmt.Fprintf(writer, "  - %s\n", item.Path)
			fmt.Fprintf(writer, "    expected_parent_index: %s\n", optionalPath(item.ExpectedParentIndex))
		}
		fmt.Fprintln(writer)
	} else {
		fmt.Fprintln(writer, "OK: no orphan markdown files in scope.")
		fmt.Fprintln(writer)
	}

	if len(report.Errors.Unreachable) > 0 {
		fmt.Fprintln(writer, "Markdown files missing from index navigation:")
		for _, item := range report.Errors.Unreachable {
			fmt.Fprintf(writer, "  - %s\n", item.Path)
			fmt.Fprintf(writer, "    expected_parent_index: %s\n", optionalPath(item.ExpectedParentIndex))
			if len(item.InboundLinks) > 0 {
				fmt.Fprintf(writer, "    inbound_links: %s\n", strings.Join(item.InboundLinks, ", "))
			}
		}
		fmt.Fprintln(writer)
	} else {
		fmt.Fprintln(writer, "OK: all scoped markdown files are reachable from the configured entrypoints via index navigation.")
		fmt.Fprintln(writer)
	}

	if len(report.Warnings.DeepReachable) > 0 {
		fmt.Fprintln(writer, "Warnings: documents reachable only deeper than the configured threshold:")
		for _, item := range report.Warnings.DeepReachable {
			fmt.Fprintf(writer, "  - %s\n", item.Path)
			fmt.Fprintf(writer, "    depth: %d\n", item.Depth)
			fmt.Fprintf(writer, "    expected_parent_index: %s\n", optionalPath(item.ExpectedParentIndex))
			fmt.Fprintf(writer, "    route: %s\n", strings.Join(item.Route, " -> "))
		}
		fmt.Fprintln(writer)
	} else {
		fmt.Fprintln(writer, "OK: no documents are reachable only deeper than the configured threshold.")
		fmt.Fprintln(writer)
	}

	fmt.Fprintln(writer, "Index compliance:")
	if len(report.Errors.IndexContract) > 0 {
		for _, item := range report.Errors.IndexContract {
			fmt.Fprintf(writer, "  - %s\n", item.Path)
			for _, issue := range item.Issues {
				fmt.Fprintf(writer, "    * %s\n", issue)
			}
		}
		fmt.Fprintln(writer)
	} else {
		fmt.Fprintln(writer, "  - OK")
		fmt.Fprintln(writer)
	}

	result := "OK"
	if report.ExitCode != 0 {
		result = "FAIL"
	}
	fmt.Fprintf(writer, "Result: %s\n", result)
	fmt.Fprintln(writer, "Machine-readable output: re-run with --json for a structured report suitable for auto-indexing.")
}

func optionalPath(value *string) string {
	if value == nil {
		return "(none)"
	}
	return *value
}

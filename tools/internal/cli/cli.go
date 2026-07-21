// Package cli defines the shared command contract for the memory-bank binaries.
package cli

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"path/filepath"
	"sort"

	"github.com/dapi/memory-bank/tools/internal/lint"
	"github.com/dapi/memory-bank/tools/internal/ownership"
	"github.com/dapi/memory-bank/tools/internal/repository"
)

const (
	defaultScopeRoot = "memory-bank"
	defaultMaxDepth  = 3
	exitSuccess      = 0
	exitFailure      = 1
	exitUsage        = 2
)

type entrypointFlags []string

func (values *entrypointFlags) String() string {
	return fmt.Sprint([]string(*values))
}

func (values *entrypointFlags) Set(value string) error {
	*values = append(*values, value)
	return nil
}

// Run executes the primary, subcommand-based memory-bank CLI.
func Run(arguments []string, version string, stdout, stderr io.Writer) int {
	if len(arguments) == 0 {
		printRootUsage(stderr)
		return exitUsage
	}

	switch arguments[0] {
	case "lint":
		return RunLint(arguments[1:], "memory-bank lint", version, stdout, stderr)
	case "init":
		return runOwnership(arguments[1:], "init", stdout, stderr)
	case "update":
		return runOwnership(arguments[1:], "update", stdout, stderr)
	case "--version", "-version":
		if len(arguments) != 1 {
			fmt.Fprintf(stderr, "memory-bank: unexpected arguments: %v\n", arguments[1:])
			return exitUsage
		}
		fmt.Fprintf(stdout, "memory-bank %s\n", version)
		return exitSuccess
	case "--help", "-h", "-help", "help":
		if len(arguments) != 1 {
			fmt.Fprintf(stderr, "memory-bank: unexpected arguments: %v\n", arguments[1:])
			return exitUsage
		}
		printRootUsage(stdout)
		return exitSuccess
	default:
		fmt.Fprintf(stderr, "memory-bank: unknown command %q\n\n", arguments[0])
		printRootUsage(stderr)
		return exitUsage
	}
}

func printRootUsage(writer io.Writer) {
	fmt.Fprintln(writer, "Memory Bank documentation tooling.")
	fmt.Fprintln(writer)
	fmt.Fprintln(writer, "Usage: memory-bank <command> [options]")
	fmt.Fprintln(writer)
	fmt.Fprintln(writer, "Commands:")
	fmt.Fprintln(writer, "  init    Adopt or install a template and create its ownership lock")
	fmt.Fprintln(writer, "  update  Safely update a template using its ownership lock")
	fmt.Fprintln(writer, "  lint    Audit markdown navigation integrity")
	fmt.Fprintln(writer)
	fmt.Fprintln(writer, "Options:")
	fmt.Fprintln(writer, "  --help       Show this help")
	fmt.Fprintln(writer, "  --version    Print the version and exit")
}

func runOwnership(arguments []string, command string, stdout, stderr io.Writer) int {
	flags := flag.NewFlagSet("memory-bank "+command, flag.ContinueOnError)
	flags.SetOutput(stderr)
	flags.Usage = func() {
		fmt.Fprintf(stderr, "Usage: memory-bank %s --source DIR --template-version VERSION --source-ref REF [options]\n", command)
		flags.PrintDefaults()
	}
	repoRootArgument := addRepoRootFlag(flags)
	sourceRootArgument := flags.String("source", "", "clean template Git checkout containing memory-bank/")
	templateVersion := flags.String("template-version", "", "human-readable template version")
	sourceRef := flags.String("source-ref", "", "full commit SHA matching the source checkout HEAD")
	dryRun := flags.Bool("dry-run", false, "print the complete mutation plan without applying it")
	jsonOutput := addJSONOutputFlag(flags)
	if err := flags.Parse(arguments); err != nil {
		if errors.Is(err, flag.ErrHelp) {
			return exitSuccess
		}
		return exitUsage
	}
	if flags.NArg() > 0 {
		fmt.Fprintf(stderr, "memory-bank %s: unexpected arguments: %v\n", command, flags.Args())
		return exitUsage
	}
	if *sourceRootArgument == "" || *templateVersion == "" || *sourceRef == "" {
		fmt.Fprintf(stderr, "memory-bank %s: --source, --template-version, and --source-ref are required\n", command)
		return exitUsage
	}
	repoRoot, err := repository.ResolveRoot(*repoRootArgument)
	if err != nil {
		fmt.Fprintln(stderr, err)
		return exitFailure
	}
	sourceRoot, err := filepath.Abs(*sourceRootArgument)
	if err != nil {
		fmt.Fprintln(stderr, err)
		return exitFailure
	}
	options := ownership.Options{
		RepoRoot: repoRoot, SourceRoot: sourceRoot, TemplateVersion: *templateVersion,
		SourceRef: *sourceRef, DryRun: *dryRun,
	}
	var report ownership.Report
	if command == "init" {
		report, err = ownership.Init(options)
	} else {
		report, err = ownership.Update(options)
	}
	if err != nil {
		if report.Applied {
			if outputErr := writeResult(stdout, *jsonOutput, report, func(writer io.Writer) {
				printOwnershipReport(writer, report)
			}); outputErr != nil {
				fmt.Fprintln(stderr, outputErr)
			}
		}
		fmt.Fprintln(stderr, err)
		return exitFailure
	}
	if err := writeResult(stdout, *jsonOutput, report, func(writer io.Writer) {
		printOwnershipReport(writer, report)
	}); err != nil {
		fmt.Fprintln(stderr, err)
		return exitFailure
	}
	if report.ConflictCount > 0 {
		return exitFailure
	}
	return exitSuccess
}

func printOwnershipReport(writer io.Writer, report ownership.Report) {
	decisions := append([]ownership.Decision(nil), report.Decisions...)
	sort.Slice(decisions, func(i, j int) bool { return decisions[i].Path < decisions[j].Path })
	for _, decision := range decisions {
		fmt.Fprintf(writer, "%s\t%s\t%s\t%s\n", decision.Action, decision.Ownership, decision.Path, decision.Reason)
	}
	switch {
	case report.ConflictCount > 0:
		fmt.Fprintf(writer, "update not applied: %d conflict(s)\n", report.ConflictCount)
	case report.DryRun:
		fmt.Fprintln(writer, "dry run: no files changed")
	case report.Applied:
		fmt.Fprintln(writer, "update applied atomically")
	}
}

// RunLint executes the lint command. commandName controls only human-readable
// usage and diagnostics, allowing the compatibility binary to preserve its CLI.
func RunLint(arguments []string, commandName, version string, stdout, stderr io.Writer) int {
	flags := flag.NewFlagSet(commandName, flag.ContinueOnError)
	flags.SetOutput(stderr)
	flags.Usage = func() {
		fmt.Fprintln(stderr, "Audit markdown navigation integrity for a memory-bank-like documentation tree.")
		fmt.Fprintln(stderr)
		fmt.Fprintf(stderr, "Usage: %s [options]\n", commandName)
		flags.PrintDefaults()
	}

	var configuredEntrypoints entrypointFlags
	repoRootArgument := addRepoRootFlag(flags)
	scopeRootArgument := flags.String("scope-root", defaultScopeRoot, "repository-relative directory to audit")
	maxDepth := flags.Int("max-depth", defaultMaxDepth, "maximum allowed navigation depth before a warning")
	jsonOutput := addJSONOutputFlag(flags)
	versionOutput := flags.Bool("version", false, "print the version and exit")
	flags.Var(&configuredEntrypoints, "entrypoint", "markdown navigation entrypoint; may be repeated")

	if err := flags.Parse(arguments); err != nil {
		if errors.Is(err, flag.ErrHelp) {
			return exitSuccess
		}
		return exitUsage
	}
	if flags.NArg() > 0 {
		fmt.Fprintf(stderr, "%s: unexpected arguments: %v\n", commandName, flags.Args())
		return exitUsage
	}
	if *versionOutput {
		fmt.Fprintf(stdout, "%s %s\n", commandName, version)
		return exitSuccess
	}
	if *maxDepth < 0 {
		fmt.Fprintf(stderr, "%s: --max-depth must be greater than or equal to 0\n", commandName)
		return exitUsage
	}
	scopeRoot, err := lint.NormalizeScopeRoot(*scopeRootArgument)
	if err != nil {
		fmt.Fprintln(stderr, err)
		return exitFailure
	}
	repoRoot, err := repository.ResolveRoot(*repoRootArgument)
	if err != nil {
		fmt.Fprintln(stderr, err)
		return exitFailure
	}
	report, err := lint.Run(lint.Options{
		RepoRoot: repoRoot, ScopeRoot: scopeRoot, Entrypoints: configuredEntrypoints, MaxDepth: *maxDepth,
	})
	if err != nil {
		fmt.Fprintln(stderr, err)
		return exitFailure
	}

	if err := writeResult(stdout, *jsonOutput, report, func(writer io.Writer) {
		lint.PrintTextReport(writer, report)
	}); err != nil {
		fmt.Fprintln(stderr, err)
		return exitFailure
	}
	return report.ExitCode
}

func addRepoRootFlag(flags *flag.FlagSet) *string {
	return flags.String("repo-root", "", "filesystem path to the repository root")
}

func addJSONOutputFlag(flags *flag.FlagSet) *bool {
	return flags.Bool("json", false, "emit a machine-readable JSON report")
}

func writeResult(writer io.Writer, jsonOutput bool, result any, writeText func(io.Writer)) error {
	if !jsonOutput {
		writeText(writer)
		return nil
	}
	encoder := json.NewEncoder(writer)
	encoder.SetEscapeHTML(false)
	encoder.SetIndent("", "  ")
	return encoder.Encode(result)
}

package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

var version = "dev"

const (
	defaultScopeRoot = "memory-bank"
	defaultMaxDepth  = 3
)

type entrypointFlags []string

func (values *entrypointFlags) String() string {
	return fmt.Sprint([]string(*values))
}

func (values *entrypointFlags) Set(value string) error {
	*values = append(*values, value)
	return nil
}

func main() {
	os.Exit(run(os.Args[1:], os.Stdout, os.Stderr))
}

func run(arguments []string, stdout, stderr io.Writer) int {
	flags := flag.NewFlagSet("memory-bank-lint", flag.ContinueOnError)
	flags.SetOutput(stderr)
	flags.Usage = func() {
		fmt.Fprintln(stderr, "Audit markdown navigation integrity for a memory-bank-like documentation tree.")
		fmt.Fprintln(stderr)
		fmt.Fprintln(stderr, "Usage: memory-bank-lint [options]")
		flags.PrintDefaults()
	}

	var configuredEntrypoints entrypointFlags
	repoRootArgument := flags.String("repo-root", "", "filesystem path to the repository root")
	scopeRootArgument := flags.String("scope-root", defaultScopeRoot, "repository-relative directory to audit")
	maxDepth := flags.Int("max-depth", defaultMaxDepth, "maximum allowed navigation depth before a warning")
	jsonOutput := flags.Bool("json", false, "emit a machine-readable JSON report")
	versionOutput := flags.Bool("version", false, "print the version and exit")
	flags.Var(&configuredEntrypoints, "entrypoint", "markdown navigation entrypoint; may be repeated")

	if err := flags.Parse(arguments); err != nil {
		if errors.Is(err, flag.ErrHelp) {
			return 0
		}
		return 2
	}
	if flags.NArg() > 0 {
		fmt.Fprintf(stderr, "memory-bank-lint: unexpected arguments: %v\n", flags.Args())
		return 2
	}
	if *versionOutput {
		fmt.Fprintf(stdout, "memory-bank-lint %s\n", version)
		return 0
	}
	if *maxDepth < 0 {
		fmt.Fprintln(stderr, "memory-bank-lint: --max-depth must be greater than or equal to 0")
		return 2
	}
	scopeRoot, err := NormalizeScopeRoot(*scopeRootArgument)
	if err != nil {
		fmt.Fprintln(stderr, err)
		return 1
	}
	repoRoot, err := ResolveRepoRoot(*repoRootArgument, scopeRoot)
	if err != nil {
		fmt.Fprintln(stderr, err)
		return 1
	}
	report, err := Run(Options{
		RepoRoot: repoRoot, ScopeRoot: scopeRoot, Entrypoints: configuredEntrypoints, MaxDepth: *maxDepth,
	})
	if err != nil {
		fmt.Fprintln(stderr, err)
		return 1
	}

	if *jsonOutput {
		encoder := json.NewEncoder(stdout)
		encoder.SetEscapeHTML(false)
		encoder.SetIndent("", "  ")
		if err := encoder.Encode(report); err != nil {
			fmt.Fprintln(stderr, err)
			return 1
		}
	} else {
		PrintTextReport(stdout, report)
	}
	return report.ExitCode
}

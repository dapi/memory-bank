// Package doctor diagnoses Memory Bank adoption and governance without mutation.
package doctor

import "github.com/dapi/memory-bank/tools/internal/lint"

// ReportFormatVersion identifies the aggregate doctor JSON schema. Version 2
// reflects the replacement of the former ownership-style drift/conflict fields
// with the aggregate summary and findings contract.
const ReportFormatVersion = 2

type Profile string

const (
	ProfileAuto       Profile = "auto"
	ProfileTemplate   Profile = "template"
	ProfileDownstream Profile = "downstream"
)

type Severity string

const (
	Error   Severity = "error"
	Warning Severity = "warning"
	Info    Severity = "info"
)

type Finding struct {
	Code        string   `json:"code"`
	Severity    Severity `json:"severity"`
	Group       string   `json:"group"`
	Path        string   `json:"path,omitempty"`
	Subject     string   `json:"subject,omitempty"`
	Message     string   `json:"message"`
	Remediation string   `json:"remediation"`
}

type TemplateIdentity struct {
	SchemaVersion int    `json:"schema_version"`
	Version       string `json:"version"`
	SourceRef     string `json:"source_ref"`
}

type Summary struct {
	Errors   int `json:"errors"`
	Warnings int `json:"warnings"`
	Info     int `json:"info"`
}

type Report struct {
	FormatVersion    int              `json:"format_version"`
	Profile          Profile          `json:"profile"`
	RepoRoot         string           `json:"repo_root"`
	TemplateIdentity TemplateIdentity `json:"template_identity"`
	Summary          Summary          `json:"summary"`
	Findings         []Finding        `json:"findings"`
	Navigation       lint.Report      `json:"navigation"`
}

type Options struct {
	RepoRoot  string
	ScopeRoot string
	AgentFile string
	Profile   Profile
	MaxDepth  int
}

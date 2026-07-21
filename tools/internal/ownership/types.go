// Package ownership implements the versioned Memory Bank ownership and update contract.
package ownership

import "time"

const (
	LockFileName         = "memory-bank/.lock"
	CurrentSchemaVersion = 1
	ReportFormatVersion  = 1
)

type Class string

const (
	Managed   Class = "managed"
	Adapted   Class = "adapted"
	UserOwned Class = "user-owned"
	Generated Class = "generated"
)

type Template struct {
	Version   string `json:"version"`
	SourceRef string `json:"source_ref"`
}

type UpdateRecord struct {
	Version string    `json:"version"`
	At      time.Time `json:"at"`
}

type File struct {
	Ownership     Class  `json:"ownership"`
	BaseDigest    string `json:"base_digest,omitempty"`
	PayloadDigest string `json:"payload_digest,omitempty"`
	BaseMode      string `json:"base_mode,omitempty"`
	PayloadMode   string `json:"payload_mode,omitempty"`
}

type Lock struct {
	SchemaVersion int             `json:"schema_version"`
	Template      Template        `json:"template"`
	LastUpdate    UpdateRecord    `json:"last_update"`
	Files         map[string]File `json:"files"`
}

type Action string

const (
	Create     Action = "create"
	UpdateFile Action = "update"
	Preserve   Action = "preserve"
	Conflict   Action = "conflict"
	Delete     Action = "delete"
)

type Decision struct {
	Path      string `json:"path"`
	Ownership Class  `json:"ownership"`
	Action    Action `json:"action"`
	Reason    string `json:"reason"`
}

type Report struct {
	FormatVersion int        `json:"format_version"`
	DryRun        bool       `json:"dry_run"`
	Applied       bool       `json:"applied"`
	Decisions     []Decision `json:"decisions"`
	ConflictCount int        `json:"conflict_count"`
}

type Options struct {
	RepoRoot        string
	SourceRoot      string
	TemplateVersion string
	SourceRef       string
	DryRun          bool
	Now             func() time.Time
	// verifySource is replaced by unit tests that use synthetic source trees.
	// CLI callers always use the Git-backed provenance verifier.
	verifySource func(string, string) error
	// BeforeMutation is used by tests after staging to simulate an interrupted update.
	BeforeMutation func(Decision) error
}

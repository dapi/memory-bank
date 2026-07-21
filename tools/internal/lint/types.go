package lint

type Options struct {
	RepoRoot    string
	ScopeRoot   string
	Entrypoints []string
	MaxDepth    int
}

type Report struct {
	FormatVersion      int      `json:"format_version"`
	RepoRoot           string   `json:"repo_root"`
	ScopeRoot          string   `json:"scope_root"`
	Entrypoints        []string `json:"entrypoints"`
	MissingEntrypoints []string `json:"missing_entrypoints"`
	MaxDepth           int      `json:"max_depth"`
	Stats              Stats    `json:"stats"`
	Errors             Errors   `json:"errors"`
	Warnings           Warnings `json:"warnings"`
	ExitCode           int      `json:"exit_code"`
}

type Stats struct {
	MarkdownFilesInScope            int `json:"markdown_files_in_scope"`
	IndexDocumentsInScope           int `json:"index_documents_in_scope"`
	BrokenLinkCount                 int `json:"broken_link_count"`
	FrontmatterDependencyIssueCount int `json:"frontmatter_dependency_issue_count"`
	OrphanCount                     int `json:"orphan_count"`
	UnreachableCount                int `json:"unreachable_count"`
	IndexContractIssueCount         int `json:"index_contract_issue_count"`
	DeepReachableWarningCount       int `json:"deep_reachable_warning_count"`
	EntrypointCount                 int `json:"entrypoint_count"`
}

type Errors struct {
	Config                  []ConfigError           `json:"config"`
	BrokenLinks             []BrokenLink            `json:"broken_links"`
	FrontmatterDependencies []FrontmatterDependency `json:"frontmatter_dependencies"`
	Orphans                 []NavigationIssue       `json:"orphans"`
	Unreachable             []NavigationIssue       `json:"unreachable"`
	IndexContract           []IndexContractIssue    `json:"index_contract"`
}

type Warnings struct {
	DeepReachable []DeepReachableWarning `json:"deep_reachable"`
}

type ConfigError struct {
	Message string   `json:"message"`
	Paths   []string `json:"paths"`
}

type BrokenLink struct {
	Source string `json:"source"`
	Target string `json:"target"`
}

type FrontmatterDependency struct {
	Source string `json:"source"`
	Field  string `json:"field"`
	Value  string `json:"value"`
	Target string `json:"target"`
}

type NavigationIssue struct {
	Path                string   `json:"path"`
	ExpectedParentIndex *string  `json:"expected_parent_index"`
	InboundLinks        []string `json:"inbound_links"`
}

type IndexContractIssue struct {
	Path                string   `json:"path"`
	Issues              []string `json:"issues"`
	ExpectedParentIndex *string  `json:"expected_parent_index"`
}

type DeepReachableWarning struct {
	Path                string   `json:"path"`
	Depth               int      `json:"depth"`
	MaxDepth            int      `json:"max_depth"`
	ExpectedParentIndex *string  `json:"expected_parent_index"`
	Route               []string `json:"route"`
}

type document struct {
	text        string
	frontmatter map[string]any
}

type reachability struct {
	depth int
	route []string
}

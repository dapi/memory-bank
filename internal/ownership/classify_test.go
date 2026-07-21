package ownership

import "testing"

func TestCurrentTemplateBoundary(t *testing.T) {
	tests := map[string]Class{
		"memory-bank/dna/governance.md":           Managed,
		"memory-bank/flows/templates/brief.md":    Managed,
		"memory-bank/prompts/PROMPT-001.md":       Managed,
		"memory-bank/domain/model.md":             Adapted,
		"memory-bank/engineering/architecture.md": Adapted,
		"memory-bank/README.md":                   Adapted,
		"memory-bank/features/README.md":          Managed,
		"memory-bank/features/FT-001/brief.md":    UserOwned,
		"memory-bank/features/FT-001/README.md":   UserOwned,
		"memory-bank/custom/project-note.md":      UserOwned,
		"memory-bank/.generated/index.json":       Generated,
	}
	for path, want := range tests {
		if got := Classify(path); got != want {
			t.Errorf("Classify(%q) = %q, want %q", path, got, want)
		}
	}
}

func TestOnlyTopLevelArtifactIndexesAreManaged(t *testing.T) {
	for _, directory := range []string{"prd", "epics", "use-cases", "features", "adr"} {
		t.Run(directory, func(t *testing.T) {
			if got := Classify("memory-bank/" + directory + "/README.md"); got != Managed {
				t.Fatalf("top-level index classified as %q, want %q", got, Managed)
			}
			if got := Classify("memory-bank/" + directory + "/instance/README.md"); got != UserOwned {
				t.Fatalf("nested artifact README classified as %q, want %q", got, UserOwned)
			}
		})
	}
}

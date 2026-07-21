package ownership

import "strings"

// Classify defines the template boundary. Paths must be repository-relative.
func Classify(path string) Class {
	if strings.HasPrefix(path, "memory-bank/.generated/") {
		return Generated
	}
	if path == "memory-bank/README.md" {
		return Adapted
	}
	for _, prefix := range []string{"memory-bank/dna/", "memory-bank/flows/", "memory-bank/prompts/"} {
		if strings.HasPrefix(path, prefix) {
			return Managed
		}
	}
	for _, prefix := range []string{"memory-bank/product/", "memory-bank/domain/", "memory-bank/engineering/", "memory-bank/ops/"} {
		if strings.HasPrefix(path, prefix) {
			return Adapted
		}
	}
	for _, prefix := range []string{"memory-bank/prd/", "memory-bank/epics/", "memory-bank/use-cases/", "memory-bank/features/", "memory-bank/adr/"} {
		if strings.HasPrefix(path, prefix) {
			if path == prefix+"README.md" {
				return Managed
			}
			return UserOwned
		}
	}
	// Unknown files below memory-bank are downstream-owned by default. This is
	// deliberately fail-safe for directories introduced by a project.
	return UserOwned
}

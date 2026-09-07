---
title: "FT-141: Implementation plan"
doc_kind: feature
doc_function: derived
purpose: Execute the template-owned component declarations and adoption documentation.
derived_from:
  - brief.md
  - design.md
status: active
audience: humans_and_agents
---

# FT-141: Implementation plan

## Preconditions and grounding

PRE-01: ADR-002 accepted, design active, and independent review of this plan clean before payload writes.
Template baseline: f1f04de843aef45a2425d4a7351d577bbf89e940. CLI implementation belongs to CLI #62;
this plan owns only the template payload, its installer entrypoint and producer integration checks.

| Grounding | Inspected owner and observed fact | Execution consequence |
| --- | --- | --- |
| GRND-01 | template/memory-bank/dna/README.md: Universal Governance Baseline imports flows/priming | STEP-01 replaces the inverse dependency with a DNA-only reading sequence |
| GRND-02 | template/memory-bank/dna/frontmatter.md and lifecycle.md: conditional type/flow fields are global | STEP-01 moves specialized meaning to Documents or process extensions |
| GRND-03 | template/memory-bank/flows/templates/feature/brief.md: complete process-bearing template | STEP-02 adds independent base templates and retains old paths as extension entrypoints |
| GRND-04 | template/memory-bank/engineering/testing-conventions.md and section indexes: mandatory flow imports | STEP-02 removes mandatory reverse dependencies from Documents |
| GRND-05 | tools/refresh-memory-bank-projection.rb: plan/apply! preserves real project files and creates generic symlinks | STEP-03 refreshes the projection after adding payload files |
| GRND-06 | tools/validate-priming-manifests.rb and repository AGENTS.md: manifest/link/doctor checks are existing delivery gates | STEP-03 adds component-specific semantic and binary checks alongside them |

## Implementation priming

Read in order before the corresponding step: template/memory-bank/dna/README.md and
frontmatter.md (GRND-01/02, STEP-01); template/memory-bank/flows/templates/feature/brief.md
and template/memory-bank/engineering/testing-conventions.md (GRND-03/04, STEP-02);
tools/refresh-memory-bank-projection.rb#plan and tools/validate-priming-manifests.rb
(GRND-05/06, STEP-03). These paths were inspected at the grounded revision. The accepted
CTR-01 wire owner defines behavior and serialization; implementation does not invent a second schema.

## Steps and verification

1. STEP-01 / SOL-01 / INV-01 / REQ-01/07: make all six DNA documents standalone;
   add DNA rules and the exhaustive components.json inventory. Declare core/docs/full/legacy
   and optional adapter dependencies. Add the root source envelope only together with the
   capability-gated entrypoint. CHK-01/07: dependency closure and semantic audit; EVID-01/07
   are automated producer checks and exact-commit CLI consumer integration.
2. STEP-02 / SOL-03 / REQ-02/03/05/07: add document-types and base templates for ADR,
   feature, PRD, use case, research and epic; put specialized fields under their owning type
   or flow. Retain flow template paths as thin extensions with links to base contracts;
   keep companion process templates where they belong. Install immutable contract bundles,
   frozen engine artifact and compatibility corpus supplied by the CLI owner. Rewrite
   Documents section indexes and hidden mandatory dependencies; preserve human catalog
   contents. CHK-02/03/05/10/11: base/flow differentiation, immutable bundles and migration
   consumer fixtures. EVID-02/03/05/10/11 reside in template checks and linked CLI #62 tests.
3. STEP-03 / SOL-04 / REQ-04/06/08/09: add tools/install-components.sh capability gate,
   component matrix integration script and required CI; update root README.md then its Russian
   adaptation, migration guide and project-local projection. CHK-04/06/08/09: actual bridge
   refuses new source, pre-bridge entrypoint stops before installer, current Linux/macOS CLI applies
   the entire matrix and preserves state on negative paths; unsupported hosts report component
   capabilities unavailable while keeping legacy support. EVID-04/06/08/09 are pinned
   binary/source identities and command/CI results linked in PRs.

No CLI implementation or runtime wrapper is copied into this repository. Bundle engine artifacts
must match the trusted CLI bytes; the template's declarative producer checks validate the format,
while the CLI owner tests parser/transaction behavior. Script environment uses existing Go, Ruby,
Bash and the task-built pinned CLI binaries; no host agent install or environment reconfiguration.

## Required checks and checkpoints

CP-01: standalone DNA/Documents have no flow/adapter dependency in Markdown, derived_from,
embedded metadata, priming paths or mandatory textual instructions. Unknown inventory and
contract combinations fail producer checks. CP-02: exact template commit passes the real CLI
preset/adapter, docs-to-full, legacy migration and negative integrity/path matrix. CP-03: template
lint/doctor/priming validation, projection lint, diff checks and required CI are green; separate
independent implementation and simplification reviews are clean on the delivered revision.

Run rg --files template; ruby tools/validate-priming-manifests.rb template/memory-bank;
memory-bank-cli lint --scope-root template/memory-bank --entrypoint template/memory-bank/README.md;
memory-bank-cli doctor --profile template; project-local lint; git diff --check. The component
binary matrix is required in addition to these existing checks. Test commands and actual commit
identities are recorded in the PR. SC-12 explicitly varies source, old lock, resolution,
observed file bytes/Git mode/actual permissions (0600 versus 0644), directory state, selection,
write intents and registry bytes; each stale input must reject before writes. Recovery fixtures
come from the CLI owner and distinguish complete rollback, failed rollback requiring complete
before-state restoration, and committed cleanup failure. No template-side transaction writer
is introduced. Identities and evidence are recorded in the PR; no claim of released capability precedes a real release.

## Failure and completion

STOP-01: producer/consumer mismatch returns to CTR-01 and its owners before affected code continues.
STOP-02: missing binary capability or unresolved migration conflict preserves the old installation.
OQ-01: release tags remain assigned by the release owner; this blocks publication only.
No other design question is delegated to this execution plan. User authorized worktrees,
implementation, review/fix and PRs; merge/release/live migration have no execution step here.

Completion requires every applicable SC/CHK/EVID row in the brief and required CI at the same
revision as the final independent review. The PR explicitly links the bridge-first release dependency.

Plan Ready: independent code-converge document review completed clean at 2026-09-07T01:16:56Z against b94560c plus this staged plan and gate promotions. Execution is authorized; delivery evidence remains pending.

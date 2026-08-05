---
title: "FT-109: Task Capsule Contract Implementation Plan"
doc_kind: feature
doc_function: derived
purpose: "Derived execution plan for publishing and validating the generic Task Capsule contract."
derived_from:
  - brief.md
  - design.md
  - ../../../template/memory-bank/flows/feature.md
  - ../../../template/memory-bank/engineering/testing-policy.md
  - ../../ops/development.md
status: active
audience: humans_and_agents
---

# FT-109: Task Capsule Contract Implementation Plan

## Цель текущего плана

Publish the generic Task Capsule contract and update the Feature/Epic/session
handoff documentation indexes and lifecycle boundaries without implementing a
runtime or validator.

## Grounding Evidence

- `GRND-01` at `8066bed81fa4245d6fd3e380c39312720622b296`: `template/memory-bank/flows/{routing,feature,epic}.md` define route, lifecycle owners and gates; they are the required integration points.
- `GRND-02` at `8066bed81fa4245d6fd3e380c39312720622b296`: `template/memory-bank/flows/templates/process/session-handoff.md` already defines continuation sections; the new capsule must extend, not replace, it.
- `GRND-03` at `8066bed81fa4245d6fd3e380c39312720622b296`: `template/memory-bank/dna/{governance,frontmatter}.md` require active governed docs and upstream ownership; the contract must preserve those rules.
- `GRND-04` at `8066bed81fa4245d6fd3e380c39312720622b296`: `memory-bank/features/README.md` confirms project-local FT packages are the instantiated delivery owners; no existing package conflicts with `FT-109`.

## Implementation Priming

| Order | Path | Section / symbol | Grounding | Purpose | Required step |
| --- | --- | --- | --- | --- | --- |
| 1 | `template/memory-bank/flows/task-capsule.md` | whole document | `GRND-01`, `GRND-02`, `GRND-03` | canonical schema, ownership and resume contract | `STEP-01` |
| 2 | `template/memory-bank/flows/templates/process/task-capsule.md` | whole document | `GRND-02`, `GRND-03` | reusable instantiation template | `STEP-01` |
| 3 | `template/memory-bank/flows/{README,feature,epic}.md` | indexes and lifecycle rules | `GRND-01`, `GRND-02` | discoverability and lifecycle mapping | `STEP-02` |
| 4 | `template/memory-bank/flows/templates/process/session-handoff.md` | wrapper notes | `GRND-02` | continuation integration | `STEP-02` |
| 5 | `memory-bank/features/FT-109/{brief,design,task-capsule}.md` | canonical owner package | `GRND-04` | issue traceability and resume example | `STEP-03` |

## Grounding / Support References

- Existing references: `template/memory-bank/flows/routing.md`, `feature.md`, `epic.md`, `templates/process/session-handoff.md`.
- Existing validation commands: `ruby tools/validate-priming-manifests.rb template/memory-bank`; `memory-bank-cli lint --scope-root template/memory-bank --entrypoint template/memory-bank/README.md`; `memory-bank-cli doctor --profile template`; `git diff --check`.
- No executable test surface exists; documentation profile requires targeted schema/frontmatter/link checks and semantic read-through.

## Current State / Reference Points

The repository is documentation-only, the worktree is clean at the grounded
SHA, and the generic template has no Task Capsule contract. `memory-bank/` is
the project-local canonical adaptation layer; `template/memory-bank/` is the
generic payload.

## Test Strategy

- Run required template structure, priming-manifest, link/reachability,
  governance/marker, whitespace and required-key checks.
- Perform semantic read-through of the capsule against Feature, Epic and
  session-handoff owners, then run `codex review --uncommitted` for the
  repository diff.
- No manual-only gap: the acceptance is documentation-only and all stated
  predicates have deterministic file or command evidence.

## Open Questions / Ambiguities

`OQ-01 none`: issue scope is sufficient; runtime and validator work are explicit
follow-ups and must not enter this delivery unit.

## Environment Contract

- Worktree: `/Users/danil/worktrees/feature/issue-109-define-task-capsule-contract-for-memory`
- Branch: `feature/issue-109-define-task-capsule-contract-for-memory`
- Grounded immutable revision: `8066bed81fa4245d6fd3e380c39312720622b296`
- Required external CI: GitHub checks for the eventual PR; no PR exists yet.

## Preconditions

- `PRE-01` FT-109 brief and design remain active and internally consistent.
- `PRE-02` No parent/linked issue changes the scope or canonical owner.
- `PRE-03` Generic changes remain under `template/memory-bank/`; no project-specific details leak into the template.

## Design Realization Mapping

| Realization target | Design refs | Requirement refs | Target step | Check/evidence |
| --- | --- | --- | --- | --- |
| Canonical schema and SSoT boundary | `SOL-01..SOL-04`, `CTR-01`, `INV-01..03` | `REQ-01`, `REQ-04` | `STEP-01` | `CHK-01` / `EVID-01` |
| Feature/Epic lifecycle mapping | `SD-01..03`, `CTR-02` | `REQ-02`, `REQ-03` | `STEP-02` | `CHK-02` / `EVID-02` |
| Reusable package example | `SOL-03` | `REQ-01`, `REQ-03` | `STEP-03` | `CHK-02` / `EVID-02` |

## Workstreams

- `WS-01` Publish canonical contract and process template.
- `WS-02` Integrate links and update lifecycle/session handoff mapping.
- `WS-03` Validate, review and freeze evidence.

## Approval Gates

None. No runtime, production, security, migration or external-effect action is
within scope.

## Порядок работ

| Step ID | Actor | Implements | Goal | Touchpoints | Artifact | Verifies | Evidence IDs | Check command / procedure | Blocked by | Needs approval | Escalate if |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| `STEP-01` | delivery-orchestrator | `REQ-01`, `REQ-04`, `SOL-*`, `CTR-01`, `INV-*` | Publish canonical schema, example and projection boundary | `template/memory-bank/flows/task-capsule.md`, process template | Generic contract | `SC-01`, `SC-04` | `EVID-01` | required-key scan and semantic read-through | `PRE-01` | none | owner conflict |
| `STEP-02` | delivery-orchestrator | `REQ-02`, `REQ-03`, `SD-*`, `CTR-02` | Map Feature/Epic/session handoff lifecycle | flow indexes and lifecycle docs | Updated flow docs | `SC-02`, `SC-03` | `EVID-02` | template lint/doctor and link audit | `STEP-01` | none | scope expands |
| `STEP-03` | delivery-orchestrator | `REQ-01..04` | Freeze FT-109 artifacts and validate candidate | FT-109 package and checks | Reviewable repository diff | `SC-01..04` | `EVID-01`, `EVID-02` | all required commands and review | `STEP-02` | none | failing required check |

## Parallelizable Work

- `PAR-01` No parallel writes: all docs share the same contract and one writer lease.

## Checkpoints

| Checkpoint ID | Refs | Condition | Evidence IDs |
| --- | --- | --- | --- |
| `CP-01` | `STEP-01` | Canonical schema and ownership are complete | `EVID-01` |
| `CP-02` | `STEP-02` | Feature/Epic/session handoff links and mappings are consistent | `EVID-02` |
| `CP-03` | `STEP-03` | Required checks and clean review complete on frozen candidate | `EVID-01`, `EVID-02` |

## Execution Risks

| Risk ID | Risk | Impact | Mitigation | Trigger |
| --- | --- | --- | --- | --- |
| `ER-01` | Capsule wording duplicates lifecycle owners | SSoT violation | Keep owner table and refs explicit; review against governance | Duplicate status/requirements found |
| `ER-02` | Flow mapping accidentally expands into runtime work | Scope breach | Keep runtime/validator as follow-up issues | New executable surface proposed |

## Stop Conditions / Fallback

| Stop ID | Related refs | Trigger | Immediate action | Safe fallback state |
| --- | --- | --- | --- | --- |
| `STOP-01` | `DEC-01`, `ER-01` | Conflicting canonical owner or unresolved storage trade-off | Pause and route to owner/human decision | Last clean docs revision |
| `STOP-02` | `ER-02` | Runtime/validator scope appears | Stop mutation and create separate Task Routing handoff | Generic contract only |

## Plan-local Evidence

| Evidence ID | Artifact | Producer | Path contract | Reused by checkpoints |
| --- | --- | --- | --- | --- |
| `EVID-09` | Frozen candidate revision and review record | reviewer | Run Ledger / PR review record | `CP-03` |

## Готово для приемки

- Generic contract, process template and lifecycle mappings are published.
- FT-109 owners trace every requirement to concrete checks and evidence.
- Required documentation validation, semantic read-through, review, commit,
  push, PR and required CI are complete.

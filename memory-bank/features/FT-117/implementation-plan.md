---
title: "FT-117: Implementation Plan"
doc_kind: feature
doc_function: derived
purpose: "Ретроспективный execution и verification plan для документарной реализации FT-117: canonical autonomy contract, flow handoffs, validation rules и deterministic checks."
derived_from:
  - brief.md
  - design.md
status: draft
audience: humans_and_agents
must_not_define:
  - ft_117_scope
  - ft_117_selected_design
  - ft_117_acceptance_criteria
  - ft_117_validation_profile
---

# FT-117: Implementation Plan

## Цель текущего плана

Зафиксировать и проверить уже доставленный в issue #117 governed documentation
contract без runtime changes: self-contained Structured Decision Protocol,
ownership, approval evidence, routing/flow priming и deterministic checks.

## Grounding Evidence

- Grounded repository revision: `d3639b1`
- Grounded at: `2026-08-14`

| Grounding ID | Inspected path / command | Observed current-state fact | Plan impact |
| --- | --- | --- | --- |
| `GRND-01` | `template/memory-bank/engineering/autonomy-boundaries.md` at `d3639b1` | Issue implementation already contains the self-contained protocol baseline; this package adds traceability and review evidence around that baseline | `STEP-01` verifies the canonical protocol owner and records any bounded correction |
| `GRND-02` | `template/memory-bank/flows/routing.md`, `template/memory-bank/flows/research.md` and `template/memory-bank/flows/priming/{routing,research}.yaml` | Routing had Structured Decision Protocol and read-only P0 direction; Research needed an explicit protocol dependency, bounded-probe return and bootstrap priming input | `STEP-02` aligns the Research handoff and verifies repeat-routing evidence |
| `GRND-03` | `template/memory-bank/flows/bug-fix.md`, `feature.md`, `engineering/validation-profiles.md` | Bug Fix/Feature and validation docs contained the relevant lifecycle gates and needed direct autonomy dependency/approval alignment | `STEP-03` verifies flow and profile consistency |
| `GRND-04` | `template/memory-bank/flows/priming/bug-fix.yaml`, `feature.yaml` | Candidate implementation includes the autonomy owner in the relevant priming source sets | `STEP-04` verifies exact source-set alignment |
| `GRND-05` | `ruby tools/validate-priming-manifests.rb template/memory-bank`; template-scope and project-scope `memory-bank-cli lint`; `memory-bank-cli doctor` | Documentation validation commands are available and cover both the generic payload and project-level FT-117 package | `STEP-05` runs all checks and records outputs |
| `GRND-06` | `template/`, `memory-bank/features/README.md`, `memory-bank/.lock` | `template/` is the generic downstream payload; repository-specific FT-117 history belongs in the project-level package and its index must be adapted rather than managed | `STEP-04` verifies package placement and ownership metadata |

## Implementation Priming

| Order | Exact path / stable source | Section / symbol | Grounding refs | Purpose | Required before |
| --- | --- | --- | --- | --- | --- |
| `1` | `template/memory-bank/engineering/autonomy-boundaries.md` | `# Structured Decision Protocol`, `# Human Gate` | `GRND-01` | Confirm canonical fields, outcomes and approval contract | `STEP-01` |
| `2` | `template/memory-bank/flows/routing.md` | `# Routing Predicates`, `# Human Routing` | `GRND-02` | Confirm P0/Research/Human Routing handoffs | `STEP-02` |
| `3` | `template/memory-bank/flows/research.md` | `# Terminal Dispositions and Handoff`, `# Boundary Rules` | `GRND-02` | Confirm bounded-probe return to the originating protocol and repeat routing | `STEP-02` |
| `4` | `template/memory-bank/flows/priming/research.yaml` | `stages.bootstrap` | `GRND-02` | Confirm Research receives the canonical protocol before bootstrap | `STEP-02` |
| `5` | `template/memory-bank/flows/bug-fix.md` | `# Entry Gate` | `GRND-03` | Confirm Bug Fix reroute behavior and autonomy dependency | `STEP-03` |
| `6` | `template/memory-bank/flows/feature.md` | `# Upstream Ready → Plan Ready` | `GRND-03` | Confirm review exhaustion replan/probe behavior and approval gate wording | `STEP-03` |
| `7` | `template/memory-bank/engineering/validation-profiles.md` | `# Escalation And Downgrade Rules` | `GRND-03` | Confirm approval is checked at repository and live execution gates | `STEP-03` |
| `8` | `template/memory-bank/flows/priming/bug-fix.yaml` | `stages.entry` | `GRND-04` | Confirm Bug Fix receives autonomy source before entry decision | `STEP-04` |
| `9` | `template/memory-bank/flows/priming/feature.yaml` | `stages.bootstrap_brief` | `GRND-04` | Confirm Feature receives autonomy source before design/plan routing | `STEP-04` |
| `10` | `memory-bank/features/README.md`, `memory-bank/.lock` | instantiated package index and ownership entry | `GRND-06` | Confirm FT-117 is project-level and does not leak into the generic payload | `STEP-04` |

The issue implementation is already present at `d3639b1`; any additional write
must be limited to reconciling this package or its canonical documentation
owners. If a new implementation revision is proposed, refresh `GRND-*` first.

## Grounding / Support References

| Document | Role in this plan | Facts reused | Conflict action |
| --- | --- | --- | --- |
| `brief.md` | Canonical problem, scope, profile and verify owner | `REQ-*`, `SC-*`, `CHK-*`, `EVID-*` | Update `brief.md` first |
| `design.md` | Selected solution and design-pack owner map | `SOL-*`, `SD-*`, `C4-00` | Update direct external canonical owner first |
| `../../../template/memory-bank/engineering/autonomy-boundaries.md` | External canonical protocol owner | Protocol, carrier and approval semantics | Update owner before dependent flow docs |
| `../../../template/memory-bank/flows/routing.md` | External routing owner | P0, Research and Human Routing | Update routing owner before this plan |
| `../../../template/memory-bank/flows/research.md` | External research owner | Probe lifecycle, evidence and handoff | Update research owner before this plan |
| `../../../template/memory-bank/engineering/validation-profiles.md` | External validation owner | Profile and exact approval gate | Update profile owner before plan |

## Current State / Reference Points

| Path / module | Grounding refs | Current role | Why relevant | Reuse / mirror |
| --- | --- | --- | --- | --- |
| `template/memory-bank/engineering/autonomy-boundaries.md` | `GRND-01` | Canonical autonomy and escalation rules | Owns protocol and approval evidence | Preserve one canonical owner |
| `template/memory-bank/flows/routing.md` | `GRND-02` | Route selection and Human Routing | Owns P0 and Research handoff | Link, do not duplicate protocol |
| `template/memory-bank/flows/bug-fix.md` | `GRND-03` | Bug Fix lifecycle | Owns expected-behavior and reroute gate | Link to protocol |
| `template/memory-bank/flows/feature.md` | `GRND-03` | Feature lifecycle | Owns design/plan and review convergence | Link to protocol |
| `template/memory-bank/engineering/validation-profiles.md` | `GRND-03` | Validation floor and approval timing | Owns profile decision rules | Keep approval at exact gate |
| `template/memory-bank/flows/priming/{bug-fix,feature}.yaml` | `GRND-04` | Stage source sets | Own required priming inputs | Include autonomy owner |
| `memory-bank/features/FT-117/`, `memory-bank/features/README.md`, `memory-bank/.lock` | `GRND-06` | Project-level delivery record, index and ownership metadata | Keeps repository history outside installed template payload | Preserve adapted ownership and zero FT-117 references under `template/` |

## Test Strategy

| Test surface | Canonical refs | Existing coverage | Planned automated coverage | Required local suites / commands | Required CI suites / jobs | Manual-only gap / justification | Manual-only approval ref |
| --- | --- | --- | --- | --- | --- | --- | --- |
| Governed Markdown/YAML structure | `REQ-01`–`REQ-08`, `CHK-01`–`CHK-04` | Repository has manifest, link and doctor checks | Run manifest validator, lint for both documentation roots, doctor and diff check | `ruby tools/validate-priming-manifests.rb template/memory-bank`; `memory-bank-cli lint --scope-root template/memory-bank --entrypoint template/memory-bank/README.md`; `memory-bank-cli lint --scope-root memory-bank --entrypoint memory-bank/README.md`; `memory-bank-cli doctor --profile template`; `git diff --check` | Required documentation CI jobs | `CHK-05` semantic review remains manual because no runtime test suite evaluates policy meaning | `none` |
| Cross-flow semantics | `SC-01`–`SC-09`, `NEG-01`–`NEG-03`, `CHK-05` | Text contracts exist in canonical owners | Deterministic targeted `rg` assertions plus semantic package review | `rg -n` for protocol fields, outcomes, tie-breakers and approval evidence | Documentation CI | Semantic review of canonical owner references | `none` |

## Open Questions / Ambiguities

`none` — issue intent, scope, validation profile and selected documentation
solution are sufficiently bounded. Product/value judgment remains a runtime
Human Gate rule, not an open question for this feature.

## Environment Contract

| Area | Contract | Used by | Failure symptom |
| --- | --- | --- | --- |
| setup | Run from repository root with `memory-bank-cli` installed | `STEP-05` | Command unavailable or wrong scope |
| test | Use the four commands in `CHK-01`–`CHK-04` | `STEP-05` | Evidence cannot establish documentation integrity |
| access / network / secrets | No secrets, external writes or network access required for local validation | All steps | Stop if a check unexpectedly requires external authorization |

## Preconditions

| Precondition ID | Canonical ref | Required state | Used by steps | Blocks start |
| --- | --- | --- | --- | --- |
| `PRE-01` | `brief.md` | `delivery_status: planned`, profile and verify contract are stable; promote to `in_progress` only after Plan Ready | `STEP-01`–`STEP-05` | yes |
| `PRE-02` | `design.md` | `Solution Ready`: `SOL-01`–`SOL-06`, `SD-01`–`SD-04`, `C4-00` recorded | `STEP-01`–`STEP-05` | yes |

## Design Realization Mapping

| Canonical solution refs | Owner | Realization target | Steps | Checks | Evidence |
| --- | --- | --- | --- | --- | --- |
| `SOL-01`, `SD-01`, `SD-02`, `SD-04` | `../../../template/memory-bank/engineering/autonomy-boundaries.md` | Structured Decision Protocol and approval evidence sections | `STEP-01` | `CHK-02`, `CHK-05` | `EVID-02` |
| `SOL-02`, `SD-03` | `../../../template/memory-bank/flows/routing.md` | P0 and Research/Human Routing rules | `STEP-02` | `CHK-02`, `CHK-05` | `EVID-03` |
| `SOL-05`, `SD-03` | `../../../template/memory-bank/flows/research.md` | Probe lifecycle and evidence handoff | `STEP-02` | `CHK-02`, `CHK-05` | `EVID-03` |
| `SOL-03` | `../../../template/memory-bank/flows/bug-fix.md`, `../../../template/memory-bank/flows/feature.md` | Reroute, replan and review convergence rules | `STEP-03` | `CHK-02`, `CHK-05` | `EVID-03` |
| `SOL-04`, `SOL-06`, `SD-04` | `../../../template/memory-bank/engineering/validation-profiles.md`, `../../../template/memory-bank/engineering/autonomy-boundaries.md` | Profile approval timing, edit-step autonomy and minimum contract | `STEP-03` | `CHK-02`, `CHK-05` | `EVID-04` |
| `C4-00` | `design.md` | No runtime architecture artifact | `STEP-01` | `CHK-05` | `EVID-02` |

## Workstreams

| Workstream | Implements | Result | Owner | Dependencies |
| --- | --- | --- | --- | --- |
| `WS-1` | `SOL-01`, `SD-01`, `SD-02`, `SD-04` | Canonical protocol, carrier and approval contract | agent | `PRE-01`, `PRE-02` |
| `WS-2` | `SOL-02`, `SOL-03`, `SOL-04`, `SOL-05`, `SOL-06` | Flow, validation, Research and priming alignment | agent | `WS-1` canonical terminology |
| `WS-3` | `CHK-01`–`CHK-05` | Validation outputs and semantic review evidence | agent/reviewer | `WS-1`, `WS-2` |

## Approval Gates

| Approval Gate ID | Trigger | Applies to | Why approval is required | Approver / evidence |
| --- | --- | --- | --- | --- |
| `AG-01` | This feature changes only repository documentation; no external execution is included | `none` | No risk-bearing execution step in this package | `none` |

## Порядок работ

| Step ID | Actor | Implements | Goal | Touchpoints | Artifact | Verifies | Evidence IDs | Check command / procedure | Blocked by | Needs approval | Escalate if |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| `STEP-01` | agent | `SOL-01`, `SD-01`, `SD-02`, `SD-04` | Проверить protocol schema, roles, tie-breakers и approval evidence против уже доставленного baseline | `engineering/autonomy-boundaries.md` | Canonical autonomy contract review | `CHK-02`, `CHK-05` | `EVID-02` | Review frontmatter and protocol sections | `PRE-01`, `PRE-02` | `none` | Canonical policy conflict cannot be bounded |
| `STEP-02` | agent | `SOL-02`, `SOL-05`, `SD-03` | Согласовать P0, Research bounded-probe handoff, возврат evidence в originating protocol и Human Routing contract | `flows/routing.md`, `flows/research.md` | Routing and Research contract alignment | `CHK-02`, `CHK-05` | `EVID-03` | Review route predicates, probe lifecycle, repeat protocol/routing and required evidence | `STEP-01` | `none` | Route requires missing value judgment |
| `STEP-03` | agent | `SOL-03`, `SOL-04`, `SOL-06` | Проверить Bug Fix, Feature, ordinary code/PR review и validation approval gates | `flows/bug-fix.md`, `flows/feature.md`, `engineering/validation-profiles.md`, `engineering/autonomy-boundaries.md` | Flow/profile contract review | `CHK-02`, `CHK-05` | `EVID-03`, `EVID-04` | Review dependencies, convergence, edit-step autonomy and exact gate | `STEP-01` | `none` | Risk or policy cannot be contained |
| `STEP-04` | agent | `SOL-02`, `SOL-03`, `GRND-06` | Проверить autonomy owner в relevant priming source sets и отсутствие repository-specific FT-117 package в generic payload | `flows/priming/bug-fix.yaml`, `flows/priming/feature.yaml`, `memory-bank/features/FT-117/`, `memory-bank/features/README.md`, `memory-bank/.lock` | Priming alignment and package-placement review | `CHK-01`, `CHK-02`, `CHK-03` | `EVID-03`, `EVID-05` | Run manifest validator, both link lints, doctor and targeted template leak search | `STEP-01` | `none` | Path cannot resolve or managed-content drift remains |
| `STEP-05` | agent/reviewer | `REQ-01`–`REQ-08` | Выполнить deterministic checks и semantic review | Repository scope | Check outputs and review result | `CHK-01`–`CHK-05` | `EVID-05` | Run all validation commands and review package traceability | `STEP-01`–`STEP-04` | `none` | Any required check fails |

## Parallelizable Work

- `PAR-01` После стабилизации canonical terminology можно параллельно review
  routing и validation references.
- `PAR-02` Priming manifest edits могут идти вместе с flow reference review, но
  validation commands выполняются после всех writes.

## Checkpoints

| Checkpoint ID | Refs | Condition | Evidence IDs |
| --- | --- | --- | --- |
| `CP-01` | `STEP-01`, `SOL-01`, `SD-01`–`SD-04` | Protocol and ownership contract complete | `EVID-02` |
| `CP-02` | `STEP-02`–`STEP-04`, `SOL-02`–`SOL-06` | Flow, Research, validation and priming contracts align with canonical owner | `EVID-03`, `EVID-04` |
| `CP-03` | `STEP-05`, `CHK-01`–`CHK-05` | All deterministic checks pass and semantic review is recorded | `EVID-05` |

## Execution Risks

| Risk ID | Risk | Impact | Mitigation | Trigger |
| --- | --- | --- | --- | --- |
| `ER-01` | Downstream flow invents a second protocol or approval contract | Governance drift | Keep ownership map and direct links; update canonical owner first | Duplicate rationale or permission state appears |
| `ER-02` | Documentation check passes while semantics contradict | False confidence | Require `CHK-05` semantic review against brief scenarios | Any scenario lacks owner/evidence |
| `ER-03` | New task scope introduces runtime or external effect | Wrong validation profile | Repeat Task Routing and raise profile before execution | Contract, deployment, live state or external write enters scope |

## Stop Conditions / Fallback

| Stop ID | Related refs | Trigger | Immediate action | Safe fallback state |
| --- | --- | --- | --- | --- |
| `STOP-01` | `CON-01`, `ER-03` | Change would override policy, require external write or add runtime boundary | Stop writes, update brief/design and reroute | Documentation-only scope unchanged |
| `STOP-02` | `CHK-01`–`CHK-04` | Any deterministic check fails | Fix canonical owner or stop; rerun from failed check | No `delivery_status: done` |
| `STOP-03` | `CHK-05` | Semantic review finds missing scenario coverage or duplicate owner | Update upstream canonical owner, then dependent package docs | Feature remains `planned`; plan remains `draft` until a clean Plan Ready review |

## Plan-local Evidence

| Evidence ID | Artifact | Producer | Path contract | Reused by checkpoints |
| --- | --- | --- | --- | --- |
| `EVID-05` | Validation command outputs and semantic review of this package | agent/reviewer | Issue/PR review record or repository-approved run ledger; do not append verdict to frozen artifact | `CP-03` |

## Готово для приемки

- `brief.md` and `design.md` are active and all canonical IDs have owners.
- `implementation-plan.md` has no unresolved implementation placeholders.
- `CHK-01`–`CHK-04` are green and `CHK-05` semantic review is recorded.
- No external execution, production mutation, merge, release or deployment is
  performed by this package.
- The package remains `planned` until the clean Plan Ready review; only then may
  execution promote `brief.md` to `delivery_status: in_progress`.

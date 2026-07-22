---
title: "FT-53: Implementation Plan"
doc_kind: feature
doc_function: derived
purpose: "Execution plan for the brownfield adaptation protocol without redefining FT-53 problem-space facts or validation profile."
derived_from:
  - brief.md
  - decision-log.md
status: active
audience: humans_and_agents
must_not_define:
  - ft_53_scope
  - ft_53_selected_design
  - ft_53_acceptance_criteria
  - ft_53_blocker_state
  - ft_53_validation_profile
---

# FT-53: Implementation Plan

## Goal

Deliver the documentation outcome defined by `REQ-53-01`–`REQ-53-06`, retaining
`brief.md` as the canonical owner of requirements and verification.

## Grounding / Support References

| Document | Role in this plan | Facts reused | Conflict action |
| --- | --- | --- | --- |
| [`brief.md`](brief.md) | Canonical problem, validation and verify owner | `REQ-53-*`, `SC-53-*`, `CHK-53-*`, `EVID-53-*` | Update brief first |
| [`decision-log.md`](decision-log.md) | Feature-local decision evidence | `DEC-53-01`, `DEC-53-02` | Update decision log and brief if problem-space facts change |
| [`../../../docs/adoption.md`](../../../docs/adoption.md) | Existing brownfield guidance to reconcile | Current inventory, order and readiness sections | Update canonical guidance rather than duplicate it |
| [`../../../docs/greenfield-integration-protocol.md`](../../../docs/greenfield-integration-protocol.md) | Adjacent lifecycle reference | Existing install/adaptation ordering | Escalate lifecycle conflict to brief before edits |

`design.md`, support artifacts, contracts, diagrams and ADRs: `none`, per
`DEC-53-01` in [`brief.md`](brief.md#design-requirement-decision).

## Current State / Reference Points

| Path / module | Current role | Why relevant | Reuse / mirror |
| --- | --- | --- | --- |
| `docs/adoption.md` | Entry guidance for brownfield and greenfield adoption | Must link to the dedicated protocol and no longer carry contradictory lifecycle detail | Keep it as concise routing/summary guidance |
| `docs/greenfield-integration-protocol.md` | Executable neighboring protocol | Defines the expected level of lifecycle detail and source-evidence discipline | Mirror structure where applicable, not greenfield-specific facts |
| `README.md` | Root documentation navigation | Must make the new protocol discoverable | Add one annotated link |

## Test Strategy

Validation profile is `documentation`, selected only in [`brief.md`](brief.md#validation-profile-decision).

| Test surface | Canonical refs | Planned coverage | Required local commands | Manual-only gap / justification |
| --- | --- | --- | --- | --- |
| Protocol completeness | `REQ-53-01`–`REQ-53-05`, `CHK-53-01` | Targeted semantic read-through | `git diff -- docs/adoption.md docs/brownfield-adaptation-protocol.md README.md` | Reviewer verifies prose meaning |
| Guidance reconciliation | `REQ-53-06`, `CHK-53-02` | Whitespace and lifecycle read-through | `git diff --check` | Reviewer compares lifecycle boundaries |
| Memory Bank navigation | `CHK-53-03` | Links and indexes | `memory-bank lint` | If unavailable, state exact gap and use brief fallback |

## Open Questions / Ambiguities

none. `DEC-53-01` and `DEC-53-02` are closed from issue evidence in
[`decision-log.md`](decision-log.md); no open question blocks planning.

## Environment Contract

| Area | Contract | Used by | Failure symptom |
| --- | --- | --- | --- |
| setup | Run from repository root with Git available | `STEP-53-01`–`STEP-53-04` | Paths or diff evidence cannot be resolved |
| test | `memory-bank lint` is the primary navigation check; `git diff --check` checks whitespace | `CHK-53-02`, `CHK-53-03` | Missing command produces a stated verification gap, not a pass |
| access / network / secrets | No secrets, external access or runtime environment is required | all steps | Any request to copy sensitive project data stops the work and triggers scope escalation |

## Preconditions

| Precondition ID | Canonical ref | Required state | Used by steps | Blocks start |
| --- | --- | --- | --- | --- |
| `PRE-53-01` | `DEC-53-01` | Feature remains documentation-only | `STEP-53-01`–`STEP-53-04` | yes |
| `PRE-53-02` | `DEC-53-02`, `DEC-53-03` | Intake PRD lifecycle remains evidence-backed, staged and starts at `./brownfield-intake-prd.md` | `STEP-53-01` | yes |

## Design Realization Mapping

Not applicable: `Design required: no` in [`brief.md`](brief.md#design-requirement-decision).

## Workstreams

| Workstream | Implements | Result | Owner | Dependencies |
| --- | --- | --- | --- | --- |
| `WS-53-01` | `REQ-53-01`–`REQ-53-05` | Dedicated generic brownfield protocol | agent | `PRE-53-01`, `PRE-53-02` |
| `WS-53-02` | `REQ-53-06` | Reconciled adoption/root navigation | agent | `WS-53-01` |
| `WS-53-03` | `CHK-53-01`–`CHK-53-03` | Verification evidence or explicit gap | agent / reviewer | `WS-53-01`, `WS-53-02` |

## Approval Gates

None. The `documentation` validation profile requires ordinary review only;
any scope expansion to runtime, secrets or project-specific facts must stop and
be escalated under `STOP-53-01`.

## Work Sequence

| Step ID | Actor | Implements | Goal | Touchpoints | Artifact | Verifies | Evidence IDs | Check command / procedure | Blocked by | Needs approval | Escalate if |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| `STEP-53-01` | agent | `REQ-53-01`–`REQ-53-05` | Write protocol with staged discovery, intake at `./brownfield-intake-prd.md`, adaptation, conversion, conditional inventory, safety and DoD | `docs/brownfield-adaptation-protocol.md` | New protocol | `CHK-53-01` | `EVID-53-01` | Compare every section with brief and issue | `PRE-53-01`, `PRE-53-02` | none | An unresolvable conflict with canonical guidance appears |
| `STEP-53-02` | agent | `REQ-53-06` | Replace detailed brownfield lifecycle in adoption docs with concise route; link root navigation | `docs/adoption.md`, `README.md` | Reachable guidance | `CHK-53-02` | `EVID-53-02` | Read all three lifecycle documents together | `STEP-53-01` | none | A required change affects unrelated guidance |
| `STEP-53-03` | agent | `CHK-53-01`, `CHK-53-02` | Run documentation integrity checks | changed docs | Review result | `CHK-53-01`, `CHK-53-02` | `EVID-53-01`, `EVID-53-02` | `git diff --check`; semantic read-through | `STEP-53-02` | none | A requirement cannot be evidenced from issue or docs |
| `STEP-53-04` | agent | `CHK-53-03` | Verify Memory Bank links/indexes and record result | `memory-bank/` | Lint result or gap | `CHK-53-03` | `EVID-53-03` | `memory-bank lint`, otherwise fallback in brief | `STEP-53-02` | none | Lint reports unresolved issue outside feature scope |

## Checkpoints

| Checkpoint ID | Refs | Condition | Evidence IDs |
| --- | --- | --- | --- |
| `CP-53-01` | `STEP-53-01`, `CHK-53-01` | All required protocol sections trace to `REQ-53-*` IDs | `EVID-53-01` |
| `CP-53-02` | `STEP-53-02`, `CHK-53-02` | Adoption, root and greenfield paths have compatible lifecycle boundaries | `EVID-53-02` |
| `CP-53-03` | `STEP-53-04`, `CHK-53-03` | Lint passes or exact verification gap is captured | `EVID-53-03` |

## Execution Risks

| Risk ID | Risk | Impact | Mitigation | Trigger |
| --- | --- | --- | --- | --- |
| `ER-53-01` | Protocol duplicates and drifts from adoption guidance | Conflicting instructions | Keep adoption as short routing and protocol as detailed lifecycle owner | Same lifecycle rule has different wording/order |
| `ER-53-02` | Generic guidance asserts downstream facts | Unsafe adaptation | Require source references, confidence and explicit unknown/conflict records | Draft names an unsupported product, endpoint, secret or owner |

## Stop Conditions / Fallback

| Stop ID | Related refs | Trigger | Immediate action | Safe fallback state |
| --- | --- | --- | --- | --- |
| `STOP-53-01` | `CON-53-01`, `NS-53-01`, `NS-53-02` | Edit needs project-specific facts, secrets, runtime behavior or unproven lifecycle choice | Stop and request human decision with evidence/options | Keep scoped documentation draft; do not claim completion |

## Plan-local Evidence

| Evidence ID | Artifact | Producer | Path contract | Reused by checkpoints |
| --- | --- | --- | --- | --- |
| `EVID-53-09` | Feature-pack review report | documentation quality agent | `memory-bank/features/FT-53/feature-review-report.md` | final handoff |

## Ready for Acceptance

- All workstreams and checkpoints are complete.
- `CHK-53-01`–`CHK-53-03` have evidence or the permitted verification gap is explicit.
- Final acceptance follows [`brief.md`](brief.md#verify), not this plan.

---
title: "FT-012: Design"
doc_kind: feature
doc_function: canonical
purpose: "Solution-space документ для FT-012. Фиксирует выбранный подход к compact task-flow family, task templates, routing indexes и promotion boundaries без переопределения problem space или execution contract."
derived_from:
  - brief.md
  - ../../flows/feature-flow.md
  - https://github.com/dapi/memory-bank/issues/12
status: active
audience: humans_and_agents
must_not_define:
  - ft_012_scope
  - ft_012_acceptance_criteria
  - ft_012_evidence_contract
  - implementation_sequence
---

# FT-012: Design

## Design Pack

| Artifact | Role | Owns |
| --- | --- | --- |
| `design.md` | Feature-local solution owner | `SOL-*`, `ALT-*`, `TRD-*`, `C4-*`, feature-local `CTR-*`, `INV-*`, `FM-*`, `RB-*` |
| `decision-log.md` | FPF-backed decision log | Local reasoning entries for decisions that do not need global ADR |

## Context

Issue #12 names a concrete source family from `alfagen/mercury`: `task-flow.md`, `bugfix-flow.md`, `refactor-flow.md`, `templates/task/*` and `tasks/README.md`. The target repository is a generic template, so the design problem is not whether to copy those files verbatim, but how to adapt their reusable process structure while keeping `feature-flow` as the owner of feature delivery and preventing compact task profiles from hiding higher-risk work.

## C4 Applicability

| C4 ID | Decision | Trigger / reason | Artifact |
| --- | --- | --- | --- |
| `C4-00` | `not required` | The feature changes documentation governance and templates only. It does not change runtime/deployable/container boundaries, external system interactions, APIs, schemas or storage topology. | `none` |

## Selected Solution

- `SOL-01` Add `task-flow.md` as the family-level non-feature lifecycle and carrier/package owner for `tracker-only`, `bugfix-compact`, `refactor-small` and `managed-task`, satisfying `REQ-01`, `REQ-06`.
- `SOL-02` Add `bugfix-flow.md` and `refactor-flow.md` as profile-specific process owners, so root cause/regression coverage and behavior invariant/checkpoint rules do not live in the family flow, satisfying `REQ-01`.
- `SOL-03` Add `memory-bank/flows/templates/task/package-README.md`, `bugfix.md` and `refactor.md` as governed wrapper templates for issue/PR compact notes and durable `TASK-XXX/` packages, satisfying `REQ-02`.
- `SOL-04` Add `memory-bank/tasks/README.md` as an optional destination index for managed non-feature packages, satisfying `REQ-03`.
- `SOL-05` Update `workflows.md` and navigation indexes so task profiles are discoverable before feature escalation, while preserving `feature-flow.md` and `epic-flow.md` as promotion targets, satisfying `REQ-04`, `REQ-06`.
- `SOL-06` Apply source sanitization: adapt generic routing/profile/promotion semantics from source docs, but drop source project examples, source issue ids and source-only operational docs, satisfying `REQ-05`.

## Alternatives Considered

| Alternative ID | Option | Why not selected |
| --- | --- | --- |
| `ALT-01` | Keep using `feature-flow` for all bugfix/refactor/chore tasks | Rejected by issue #12 problem statement and acceptance: compact task flow must explicitly differ from feature-flow. |
| `ALT-02` | Copy `alfagen/mercury` task docs verbatim | Rejected by `REQ-05`: source docs include downstream examples and references outside issue #12 scope. |
| `ALT-03` | Add only a short section to `workflows.md`, without standalone `task-flow`, `bugfix-flow`, `refactor-flow` | Rejected by issue #12 scope, which explicitly asks for the three generic flow documents and `TASK-XXX` templates. |
| `ALT-04` | Add `workflow-decision-log.md` and `workflow-metrics.md` because source `workflows.md` references them | Rejected for FT-012: those artifacts are not in issue #12 scope and would broaden the delivery unit. |

## Trade-offs

| Trade-off ID | Decision | Benefit | Cost / Risk |
| --- | --- | --- | --- |
| `TRD-01` | Use three flow docs rather than a single long task-flow | Clear ownership: family routing/package rules vs bugfix/refactor process details | More documents and more index links to keep coherent |
| `TRD-02` | Keep task packages without `design.md` or `implementation-plan.md` | Prevents compact non-feature work from becoming hidden feature-flow | Some medium work must promote earlier rather than accumulating local task docs |
| `TRD-03` | Update `workflows.md` selector as a related routing doc | Makes compact task profiles discoverable from the current task entrypoint | Requires careful wording to avoid implying feature-flow is deprecated |

## Accepted Local Decisions

- `SD-01` `workflows.md` is in implementation scope as a related canonical routing doc because compact profiles cannot be reliably discovered from indexes alone.
- `SD-02` Source-only `workflow-decision-log.md` and `workflow-metrics.md` are excluded from FT-012; if needed later, they require a separate issue or feature package.
- `SD-03` `TASK-XXX/` packages must not contain `design.md` or `implementation-plan.md`; the appearance of solution-space or execution-plan needs is a promotion signal.
- `SD-04` `memory-bank/tasks/README.md` will not include source example packages. It may describe naming and package rules generically.

## Contracts

| Contract ID | Input / Output | Producer / Consumer | Semantics / Constraints |
| --- | --- | --- | --- |
| `CTR-01` | Routing Signature fields | `workflows.md` -> task/feature/epic flows | Must include enough fields to route kind, size, risk, change surface, contract change, design need, evidence need, owner doc need and workflow profile. |
| `CTR-02` | Promotion triggers | `workflows.md`, `task-flow.md`, profile flows/templates -> agents/reviewers | Any capability, contract, high/critical risk, design reasoning or multi-delivery-unit trigger routes out of compact task profile. |
| `CTR-03` | Durable task package shape | `task-flow.md`, task templates, `tasks/README.md` -> `memory-bank/tasks/TASK-XXX/` | Durable package contains `README.md` plus exactly one primary owner doc: `bugfix.md` or `refactor.md`, unless explicitly used as routing index for follow-ups. |
| `CTR-04` | Source sanitization | Implementation -> template repository | Source provenance may be cited in FT-012 feature docs, but generated generic template docs must not embed source project terms or examples. |

## Invariants

- `INV-01` Compact task profiles never replace feature/epic/ADR ownership when promotion triggers are present.
- `INV-02` `feature-flow.md` remains canonical for feature packages and must not be redefined by task-flow docs.
- `INV-03` All new governed docs and templates include YAML frontmatter with `status`.
- `INV-04` Navigation from `memory-bank/README.md` to task flows/templates/tasks destination remains reachable by link audit.
- `INV-05` Source-specific facts from `alfagen/mercury` are not copied into generic target docs.

## Failure Modes

- `FM-01` Compact task docs become a loophole for feature/contract/high-risk work. Mitigation: duplicate promotion triggers in selector, task-flow, profile flows and templates.
- `FM-02` New task docs are created but not discoverable. Mitigation: update root, flows, templates and tasks indexes and run `CHK-01`.
- `FM-03` Source-specific examples leak into the generic template. Mitigation: run `CHK-03` against known source-specific terms and remove any matches outside FT-012 provenance.
- `FM-04` Workflows selector conflicts with `feature-flow`. Mitigation: wording must say task-flow is a minimal non-feature profile family and feature-flow remains required for capability/contract/design work.

## Rollout / Backout

| Stage ID | Stage | Entry condition | Backout |
| --- | --- | --- | --- |
| `RB-01` | Add feature package docs | Issue #12 read and feature-flow reviewed | Remove `memory-bank/features/FT-012/` and related index link before implementation begins |
| `RB-02` | Add task-flow family docs/templates/indexes | `brief.md` and `design.md` active | Revert new task docs/index links if link audit or review shows scope conflict |
| `RB-03` | Verify and handoff | Docs complete and checks available | Leave feature in `planned` with explicit blockers if checks fail |

## ADR / External Design Dependencies

| Artifact | Current status | Used for | Rule |
| --- | --- | --- | --- |
| GitHub issue [#12](https://github.com/dapi/memory-bank/issues/12) | open | Scope, source list and acceptance | Source of requirements for FT-012 |
| `alfagen/mercury` source docs | external source evidence | Reusable task-flow/profile/template semantics | Adapt only generic process structure; do not copy source-specific examples |
| ADR | none | No reusable architecture decision required | Create ADR only if implementation changes cross-feature architecture policy beyond issue #12 |

## Traceability

| Requirement ID | Solution refs | Contracts / invariants | Failure / rollout refs |
| --- | --- | --- | --- |
| `REQ-01` | `SOL-01`, `SOL-02`, `ALT-03`, `TRD-01` | `CTR-01`, `CTR-02`, `INV-01`, `INV-02`, `INV-03` | `FM-01`, `FM-04`, `RB-02` |
| `REQ-02` | `SOL-03`, `TRD-02` | `CTR-02`, `CTR-03`, `INV-01`, `INV-03` | `FM-01`, `RB-02` |
| `REQ-03` | `SOL-04`, `SD-04` | `CTR-03`, `CTR-04`, `INV-03`, `INV-05` | `FM-02`, `FM-03`, `RB-02` |
| `REQ-04` | `SOL-05`, `SD-01`, `TRD-03` | `CTR-01`, `CTR-02`, `INV-02`, `INV-04` | `FM-02`, `FM-04`, `RB-02` |
| `REQ-05` | `SOL-06`, `ALT-02`, `SD-02`, `SD-04` | `CTR-04`, `INV-05` | `FM-03`, `RB-02` |
| `REQ-06` | `SOL-01`, `SOL-05`, `SD-03` | `CTR-02`, `CTR-03`, `INV-01`, `INV-02` | `FM-01`, `FM-04`, `RB-02` |

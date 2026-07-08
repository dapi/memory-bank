---
title: "FT-013: Design"
doc_kind: feature
doc_function: canonical
purpose: "Solution-space документ для FT-013. Фиксирует выбранный подход к workflow decision log, workflow metrics и developer brief без переопределения problem space или execution contract."
derived_from:
  - brief.md
status: active
audience: humans_and_agents
must_not_define:
  - ft_013_scope
  - ft_013_acceptance_criteria
  - ft_013_evidence_contract
  - implementation_sequence
---
# FT-013: Design

## Design Pack

| Artifact | Role | Owns |
| --- | --- | --- |
| `design.md` | Feature-local solution owner | `SOL-*`, `ALT-*`, `TRD-*`, `C4-*`, feature-local `CTR-*`, `INV-*`, `FM-*`, `RB-*` |
| `decision-log.md` | Feature-local decision provenance | FPF-backed `DL-FT013-*` decisions used by this design |

## Context

Issue 13 asks to adapt a proven source pattern into the generic memory-bank template: workflow decision log, workflow metrics and developer brief for workflow selector-а. Current `flows/workflows.md` provides coarse task routing, while issue 12 owns the future compact task-flow, bugfix-flow and refactor-flow artifacts. FT-013 therefore needs a solution that adds optional, reachable process docs now without silently implementing issue 12 or adding broken links to absent local files.

The design problem is not runtime architecture. It is documentation ownership: where the new process docs live, how generic they must be, how safety-first metrics are expressed, and how the issue 12 dependency is represented.

## C4 Applicability

| C4 ID | Decision | Trigger / reason | Artifact |
| --- | --- | --- | --- |
| `C4-00` | `not required` | The feature changes markdown governance/process docs only. It does not change API, schema, runtime topology, deployable/container boundaries, security boundary, queue/storage or external integration. | `none` |

## Selected Solution

- `SOL-01` Add `memory-bank/flows/workflow-decision-log.md` as an optional governed process document. It owns workflow selector rationale, decision entries and review/update rules for selector/profile changes.
- `SOL-02` Add `memory-bank/flows/workflow-metrics.md` as an optional governed metrics template. It owns metric cards and decision rules for evaluating workflow profile changes, with safety/evidence gates before speed.
- `SOL-03` Add `memory-bank/flows/workflow-routing-developer-brief.md` as a short developer-facing guide/template. It explains workflow profile choice, promotion triggers and expected evidence at a level that stays usable before issue 12 lands.
- `SOL-04` Update `memory-bank/flows/README.md` so the optional docs are reachable from the flows index.
- `SOL-05` Update `memory-bank/flows/workflows.md` with a compact pointer from routing rules to the optional decision log, metrics and developer brief.
- `SOL-06` Keep issue 12 as an explicit adjacent dependency in FT-013 docs and implementation notes, but do not create `task-flow.md`, `bugfix-flow.md`, `refactor-flow.md`, task templates or `memory-bank/tasks/`.

## Alternatives Considered

| Alternative ID | Option | Why not selected |
| --- | --- | --- |
| `ALT-01` | Copy source docs into `memory-bank/flows/` verbatim | Rejected because issue 13 explicitly forbids project-specific dates, project names and operational details. |
| `ALT-02` | Put all three docs under `memory-bank/flows/templates/` only | Rejected because issue and source pattern point to reusable flow-level docs, and acceptance requires optional docs reachable from `flows/README.md`; putting them only under templates would hide them from routing use. |
| `ALT-03` | Wait for issue 12 before adding any workflow routing metrics/docs | Rejected because issue 13 can deliver generic optional evaluation/rationale docs now, as long as it avoids local links to absent task-flow artifacts. |
| `ALT-04` | Add task-flow, bugfix-flow and refactor-flow links speculatively | Rejected because those files do not exist in current repo and issue 12 owns their creation. This would fail link audit and expand scope. |

## Trade-offs

| Trade-off ID | Decision | Benefit | Cost / Risk |
| --- | --- | --- | --- |
| `TRD-01` | Keep docs in `flows/` rather than only `flows/templates/` | High discoverability from routing docs and direct reuse by downstream projects | Docs must be written as generic guidance with placeholders, not as instantiated project pilot records |
| `TRD-02` | Use configurable placeholders/examples instead of source pilot dates and fixed thresholds | Avoids leaking source operational context into the generic template | Downstream projects must set concrete dates/thresholds during adoption |
| `TRD-03` | Mention compact profiles generically before issue 12 lands | Lets developer brief explain the selector intent now | Requires careful wording to avoid promising local docs that are not present yet |

## Accepted Local Decisions

- `SD-01` FT-013 creates a feature-local `decision-log.md` for FPF-backed documentation decisions made during package creation and review-improve.
- `SD-02` `workflow-metrics.md` must treat speed/lead-time as a secondary observation; any safety misroute or missing critical evidence blocks a positive success verdict.
- `SD-03` Generic workflow docs must not link to absent local `task-flow.md`, `bugfix-flow.md` or `refactor-flow.md`. If issue 12 later adds them, a follow-up may add concrete links.
- `SD-04` `workflow-routing-developer-brief.md` is a reusable guide/template, not a project-specific announcement. It may describe profile families in generic terms and point to existing canonical `workflows.md`, `feature-flow.md` and `epic-flow.md`.

## Contracts

| Contract ID | Input / Output | Producer / Consumer | Semantics / Constraints |
| --- | --- | --- | --- |
| `CTR-01` | New `workflow-*.md` docs under `memory-bank/flows/` | Implementer / flows readers | Each doc has YAML frontmatter, `status`, relative `derived_from` only to existing local docs, and no broken internal links. |
| `CTR-02` | `flows/README.md` navigation | Implementer / link audit and readers | The index includes annotated links to the optional workflow docs. |
| `CTR-03` | `workflows.md` routing pointer | Implementer / agents choosing workflow | Routing rules point readers to decision log, metrics and developer brief as optional aids without making them mandatory lifecycle gates. |
| `CTR-04` | Generic adaptation boundary | Implementer / reviewers | Delivered workflow docs omit source project names, source pilot dates, source operational commands and downstream-specific domain details. |

## Invariants

- `INV-01` Optional workflow docs must not become required gates for every task unless a future governance change explicitly promotes them.
- `INV-02` Safety/evidence metrics outrank speed metrics in any selector success decision.
- `INV-03` Task-flow issue 12 remains the owner for compact task-flow artifacts.
- `INV-04` Feature-flow docs remain unchanged by FT-013 unless a future issue explicitly changes lifecycle rules.

## Failure Modes

- `FM-01` If new docs link to absent task-flow files, `check_memory_bank_index.py` fails. Mitigation: link only existing local docs and represent issue 12 dependency in feature docs, not reusable flow docs.
- `FM-02` If metrics reuse source pilot dates or commands as defaults, generic template quality fails. Mitigation: use placeholders, configurable windows and generic evidence sources.
- `FM-03` If `workflow-routing-developer-brief.md` repeats detailed issue 12 scope, FT-013 expands beyond its boundary. Mitigation: keep profile descriptions generic and leave task-flow artifacts to issue 12.

## Rollout / Backout

| Stage ID | Stage | Entry condition | Backout |
| --- | --- | --- | --- |
| `RB-01` | Add workflow docs and navigation | `brief.md` and `design.md` active | Remove new `workflow-*.md` docs and README/workflows links if link audit or generic leak scan fails and cannot be fixed locally |
| `RB-02` | Verify generic adaptation | New docs exist | Revert or rewrite project-specific fragments before handoff |

## ADR / External Design Dependencies

| Artifact | Current status | Used for | Rule |
| --- | --- | --- | --- |
| [GitHub issue 12](https://github.com/dapi/memory-bank/issues/12) | open | Adjacent owner for compact task-flow, bugfix-flow, refactor-flow and task templates | It is a dependency/context marker only; not a finalized local design artifact for FT-013. |
| [GitHub issue 13](https://github.com/dapi/memory-bank/issues/13) | open | Source task and acceptance criteria | Scope and acceptance in `brief.md` are derived from issue 13. |

## Traceability

| Requirement ID | Solution refs | Contracts / invariants | Failure / rollout refs |
| --- | --- | --- | --- |
| `REQ-01` | `SOL-01`, `TRD-01`, `C4-00`, `SD-01` | `CTR-01`, `INV-01` | `FM-02`, `RB-01`, `RB-02` |
| `REQ-02` | `SOL-02`, `TRD-02`, `SD-02` | `CTR-01`, `CTR-04`, `INV-02` | `FM-02`, `RB-02` |
| `REQ-03` | `SOL-03`, `TRD-03`, `SD-03`, `SD-04` | `CTR-01`, `CTR-04`, `INV-03` | `FM-01`, `FM-03`, `RB-01` |
| `REQ-04` | `SOL-04`, `SOL-05` | `CTR-02`, `CTR-03`, `INV-01` | `FM-01`, `RB-01` |
| `REQ-05` | `SOL-06`, `ALT-03`, `ALT-04`, `SD-03` | `INV-03`, `INV-04` | `FM-01`, `FM-03` |
| `REQ-06` | `ALT-01`, `TRD-02`, `SD-04` | `CTR-04` | `FM-02`, `RB-02` |

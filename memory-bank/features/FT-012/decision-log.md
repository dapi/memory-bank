---
title: "FT-012: Decision Log"
doc_kind: feature
doc_function: reference
purpose: "Feature-local decision log для FPF-backed решений по issue #12. Фиксирует существенные неоднозначности, факты, выбранные решения и последствия без замены brief/design/plan."
derived_from:
  - brief.md
  - design.md
status: active
audience: humans_and_agents
must_not_define:
  - ft_012_scope
  - ft_012_acceptance_criteria
  - implementation_sequence
---

# FT-012: Decision Log

## Принципы фиксации

Decision log фиксирует только решения, которые materially влияют на целостность feature package или scope issue #12 и не требуют global ADR. Обоснования используют FPF в plain-language форме:

- `Bounded Context`: source repository context and template repository context are separate; transfer requires explicit fit.
- `Strict Distinction`: selector, family flow, profile flow, task package and feature package have different owners and must not replace each other.
- `Evidence Graph`: claims cite concrete carriers such as issue #12, local docs and source files.
- `Reasoning Cycle`: propose -> analyze -> test; new decisions stay tied to checks in `brief.md`.

## Decisions

### `DL-001`: Exclude Source-Only Workflow Decision/Metrics Docs

| Field | Value |
| --- | --- |
| Status | accepted |
| Closed in cycle | cycle 1 |
| Question | Should FT-012 also add source-only `workflow-decision-log.md` and `workflow-metrics.md` because source `workflows.md` references them? |
| Available facts | Issue #12 source list names `task-flow.md`, `bugfix-flow.md`, `refactor-flow.md`, `templates/task/*` and `tasks/README.md`. Issue #12 scope names those documents plus related indexes. Current repository has `workflows.md` but no `workflow-decision-log.md` or `workflow-metrics.md`. Source `workflows.md` references those source-only docs. |
| FPF reasoning | Bounded Context: `alfagen/mercury` is a source context, while this repository is a generic template context; same filenames are not automatically in scope. Evidence Graph: the only requirement carrier for this feature is issue #12 plus existing local governance docs. Strict Distinction: metrics/decision-log governance for workflow evolution is different work from adding compact task carrier/profile flow. |
| Decision | Do not add `workflow-decision-log.md` or `workflow-metrics.md` in FT-012. Keep them out of scope unless a later issue explicitly asks for workflow governance metrics/decision history. |
| Consequences | `brief.md` records `NS-01`; `design.md` records `ALT-04` and `SD-02`. Implementation may update `workflows.md` selector, but must not import source-only governance docs. |

### `DL-002`: Include `workflows.md` as Related Routing Scope

| Field | Value |
| --- | --- |
| Status | accepted |
| Closed in cycle | cycle 1 |
| Question | Is updating `memory-bank/flows/workflows.md` in scope when issue #12 explicitly lists flow docs/templates/indexes but not this file by name? |
| Available facts | Current `workflows.md` is the repository's task routing document and currently routes small/medium feature, bugfix, refactor and incident work without compact document profiles. Issue #12 acceptance requires compact task flow to explicitly differ from `feature-flow` and promotion triggers to prevent hiding feature/contract/high-risk work in task packages. Source `task-flow.md` applies after selector from `workflows.md`. |
| FPF reasoning | Strict Distinction: `workflows.md` is the selector owner; `task-flow.md` is the family lifecycle owner; profile flows are specialized process owners. If selector is not updated, the new task family exists but has no canonical entrypoint. Deduction from acceptance: promotion triggers cannot be consistently applied if the canonical routing layer does not name compact profiles. |
| Decision | Treat `workflows.md` as a related routing doc in `REQ-04` and update it during implementation. |
| Consequences | `design.md` records `SD-01` and `TRD-03`; `implementation-plan.md` includes `memory-bank/flows/workflows.md` in `STEP-04`. |

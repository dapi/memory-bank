---
title: Workflow Routing Developer Brief
doc_kind: guide
doc_function: review_brief
purpose: "Краткое reusable объяснение для разработчиков и агентов: как выбирать workflow profile и когда повышать задачу до более строгого flow."
derived_from:
  - ../dna/governance.md
  - workflows.md
  - workflow-decision-log.md
  - workflow-metrics.md
status: active
audience: humans_and_agents
---
# Workflow Routing Developer Brief

Этот brief можно адаптировать для команды или вставлять в issue/PR instructions. Canonical routing rules остаются в [workflows.md](workflows.md); [workflow-decision-log.md](workflow-decision-log.md) объясняет причины changes to selector, а [workflow-metrics.md](workflow-metrics.md) помогает проверить, что profiles не ухудшили safety/evidence.

## Короткое правило

Выбирай минимальный workflow profile, который не теряет контроль над риском, evidence и traceability.

- Если работа меняет capability или user/operator/system behavior, используй feature-level flow.
- Если работа меняет contract, rollout, security boundary, data shape, integration или requires approval, повышай profile.
- Если работа маленькая, low-risk и не требует durable owner-docs, достаточно compact carrier в issue/PR.
- Если compact carrier становится тесным, повышай до managed package или feature package, а не прячь complexity в prose.

## Routing Signature

Перед стартом зафиксируй короткую routing signature.

| Field | Values / examples | Why it matters |
| --- | --- | --- |
| `kind` | feature / bugfix / chore / refactor / incident / epic / unknown | Determines the closest workflow family. |
| `size` | tiny / small / medium / large / unknown | Helps avoid heavyweight docs for small safe work. |
| `risk` | low / normal / high / critical / unknown | High risk requires stronger gates and evidence. |
| `change_surface` | single_component / multi_component / cross_boundary / external / unknown | Larger or external surfaces need stronger review. |
| `contract_change` | none / api / schema / event / security / financial / integration / rollout / unknown | Contract changes usually require promotion. |
| `design_need` | none / local_reasoning / required / unknown | Design uncertainty should not be solved inside execution steps. |
| `evidence_need` | focused_test / regression_test / manual_evidence / ci_evidence / external_e2e / unknown | Evidence need drives docs and checks. |
| `owner_doc_need` | none / task_note / managed_task_package / feature_package / epic_package / unknown | Shows where durable context should live. |

If `risk`, `contract_change`, `design_need` or `evidence_need` is `unknown`, do discovery or choose a stricter profile.

## Profile Guide

| Situation | Recommended profile | Carrier / owner |
| --- | --- | --- |
| Tiny or small low-risk change with clear evidence | `tracker-only` | Issue/PR body, comments or checklist |
| Local bugfix with reproducible symptom and clear root cause | `bugfix-compact` or `managed-task` | Issue/PR carrier or `memory-bank/tasks/TASK-XXX/` |
| Refactor/chore with no behavior change | `refactor-small` or `managed-task` | Issue/PR carrier or `memory-bank/tasks/TASK-XXX/` |
| Work needs checkpoints, durable context or several review passes, but is not a feature | `managed-task` | `memory-bank/tasks/TASK-XXX/` |
| New or materially changed capability, contract/risk trigger, design reasoning or acceptance traceability | `feature-package` | `memory-bank/features/FT-XXX/` following [feature-flow.md](feature-flow.md) |
| Multiple independent delivery units, roadmap, risk register or subissues | `epic-package` | `memory-bank/epics/EP-XXX/` following [epic-flow.md](epic-flow.md) |

For exact gates and owner boundaries, follow [workflows.md](workflows.md) first, then the concrete flow selected by the profile: [task-flow.md](task-flow.md), [bugfix-flow.md](bugfix-flow.md), [refactor-flow.md](refactor-flow.md), [feature-flow.md](feature-flow.md) or [epic-flow.md](epic-flow.md).

## Promotion Triggers

Promote to a stricter profile when any trigger is true:

- observable behavior changes beyond the original routing signature;
- API, schema, event, CLI, file format, config, security, financial, data or integration contract changes;
- rollout/backout, approval, migration or operational coordination is needed;
- implementation requires alternatives/trade-off reasoning;
- evidence cannot be captured in the selected carrier without losing traceability;
- review finds missing scope, missing evidence or repeated ambiguity;
- one task becomes several delivery units.

## What Reviewers Check

- Routing signature exists before implementation.
- Selected profile matches risk and contract triggers.
- Compact profile includes enough evidence for the changed behavior.
- Managed/feature/epic work links from issue/PR to owner docs.
- Feature work is not hidden inside a compact or task carrier.
- Speed improvement is not treated as success if safety/evidence worsened.

## Starter Comment Template

```markdown
## Workflow Routing

| Field | Value |
| --- | --- |
| kind |  |
| size |  |
| risk |  |
| change_surface |  |
| contract_change |  |
| design_need |  |
| evidence_need |  |
| owner_doc_need |  |
| workflow_profile |  |

Selected profile:

Rationale:

Promotion triggers checked:

Required carrier/docs:

Verification/evidence:
```

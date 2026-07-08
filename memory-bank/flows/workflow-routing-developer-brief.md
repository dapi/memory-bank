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
| `kind` | feature / bugfix / refactor / chore / incident / docs | Determines the closest workflow family. |
| `size` | tiny / small / medium / large | Helps avoid heavyweight docs for small safe work. |
| `risk` | low / medium / high / critical | High risk requires stronger gates and evidence. |
| `change_surface` | docs / tests / code / API / schema / infra / operations | Larger or external surfaces need stronger review. |
| `contract_change` | none / internal / external / unknown | Contract changes usually require promotion. |
| `design_need` | no / yes / unknown | Design uncertainty should not be solved inside execution steps. |
| `evidence_need` | standard / regression / rollout / approval / unknown | Evidence need drives docs and checks. |

If `risk`, `contract_change`, `design_need` or `evidence_need` is `unknown`, do discovery or choose a stricter profile.

## Profile Guide

| Situation | Recommended profile | Carrier / owner |
| --- | --- | --- |
| Tiny or small low-risk change with clear evidence | Compact issue/PR carrier | Issue/PR body, comments or checklist |
| Local bugfix with reproducible symptom and clear root cause | Bugfix-oriented compact or managed profile | Issue/PR carrier or project-defined task package |
| Refactor/chore with no behavior change | Refactor-oriented compact or managed profile | Issue/PR carrier or project-defined task package |
| Work needs checkpoints, durable context or several review passes, but is not a feature | Managed task profile | Project-defined task package |
| New or materially changed capability, contract/risk trigger, design reasoning or acceptance traceability | Feature package | `memory-bank/features/FT-XXX/` following [feature-flow.md](feature-flow.md) |
| Multiple independent delivery units, roadmap, risk register or subissues | Epic package | `memory-bank/epics/EP-XXX/` following [epic-flow.md](epic-flow.md) |

This generic template does not require a local task-flow document to exist. If a project adds one later, link the project-specific task profile docs from this brief.

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

Selected profile:

Rationale:

Promotion triggers checked:

Required carrier/docs:

Verification/evidence:
```

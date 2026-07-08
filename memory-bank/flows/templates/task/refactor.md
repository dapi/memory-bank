---
title: TASK-XXX Refactor Template
doc_kind: task
doc_function: template
purpose: "Compact template для refactor/chore profile: фиксирует intent, behavior invariants, change surface, checkpoints и verification без full feature package."
derived_from:
  - ../../workflows.md
  - ../../task-flow.md
  - ../../refactor-flow.md
  - ../../../dna/frontmatter.md
status: active
audience: humans_and_agents
template_for: refactor_task
template_target_path: ../../../tasks/TASK-XXX/refactor.md
---

# TASK-XXX: Refactor Name

Этот template можно использовать в двух режимах:

- как секцию GitHub issue/PR body для `refactor-small`;
- как `memory-bank/tasks/TASK-XXX/refactor.md`, если нужен durable managed-task package.

## Instantiated Frontmatter

```yaml
title: "TASK-XXX: Refactor Name"
doc_kind: task
doc_function: canonical
purpose: "Compact refactor/chore note: intent, invariants, change surface, checkpoints and verification."
derived_from:
  - ../../flows/workflows.md
  - ../../flows/task-flow.md
  - ../../flows/refactor-flow.md
status: active
delivery_status: planned
audience: humans_and_agents
workflow_profile: refactor-small
must_not_define:
  - feature_scope
  - new_runtime_behavior
```

## Instantiated Body

```markdown
# TASK-XXX: Refactor Name

## Routing

| Field | Value |
| --- | --- |
| `kind` | `refactor` / `chore` |
| `size` | `tiny` / `small` / `medium` |
| `risk` | `low` / `normal` / `high` |
| `change_surface` | `single_component` / `multi_component` / `cross_boundary` |
| `contract_change` | `none` |
| `workflow_profile` | `refactor-small` / `managed-task` |

Promotion check: если меняется observable behavior, contract, rollout, security/financial boundary или несколько bounded contexts, поднять profile по `workflows.md`.

## Intent

- `INT-01` Какую техническую проблему устраняем.

## Behavior Invariants

- `INV-01` Какое observable behavior должно остаться неизменным.
- `INV-02` Какие публичные contracts не должны измениться.

## Change Surface

- `SURF-01` Затронутые модули/paths.
- `NS-01` Что явно не входит в refactor.

## Checkpoints

| Checkpoint ID | Purpose | Evidence |
| --- | --- | --- |
| `CP-01` | Сохранить поведение до/после refactor | focused specs / CI / diff review |

## Verification

| Check ID | Covers | How to check | Expected result | Evidence |
| --- | --- | --- | --- | --- |
| `CHK-01` | `INV-01`, `INV-02` | focused specs / existing suite / static scan | Поведение и contracts сохранены | PR checks / CI / attached log |
```

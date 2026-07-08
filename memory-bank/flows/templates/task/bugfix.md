---
title: TASK-XXX Bugfix Template
doc_kind: task
doc_function: template
purpose: "Compact template для bugfix profile: фиксирует symptom, reproduction, root cause, fix boundary и regression evidence без full feature package."
derived_from:
  - ../../routing.md
  - ../../bug-fix.md
  - ../../../dna/frontmatter.md
status: active
audience: humans_and_agents
template_for: bugfix_task
template_target_path: ../../../tasks/TASK-XXX/bugfix.md
---

# TASK-XXX: Bugfix Name

Этот template можно использовать в двух режимах:

- как секцию GitHub issue/PR body для `bugfix-compact`;
- как `memory-bank/tasks/TASK-XXX/bugfix.md`, если нужен durable managed-task package.

## Instantiated Frontmatter

```yaml
title: "TASK-XXX: Bugfix Name"
doc_kind: task
doc_function: canonical
purpose: "Compact bugfix note: symptom, reproduction, root cause, fix boundary and regression evidence."
derived_from:
  - ../../flows/routing.md
  - ../../flows/bug-fix.md
status: active
delivery_status: planned
audience: humans_and_agents
workflow_profile: bugfix-compact
must_not_define:
  - feature_scope
  - architecture_decision
```

## Instantiated Body

```markdown
# TASK-XXX: Bugfix Name

## Routing

| Field | Value |
| --- | --- |
| `kind` | `bugfix` |
| `size` | `tiny` / `small` / `medium` |
| `risk` | `low` / `normal` / `high` |
| `change_surface` | `single_component` / `multi_component` / `cross_boundary` |
| `contract_change` | `none` |
| `workflow_profile` | `bugfix-compact` |

Promotion check: если `contract_change != none`, `risk=high/critical` или root cause требует design reasoning, остановиться и повторить routing по `flows/routing.md`.

## Symptom

- Что сломано:
- Где наблюдается:
- Кто/что затронуто:

## Reproduction

- `REPRO-01` Минимальный способ воспроизвести дефект или ссылка на evidence.

## Root Cause

- `RC-01` Проверяемая причина дефекта.

## Fix Boundary

- `FIX-01` Что меняем.
- `NS-01` Что не меняем.

## Regression Coverage

| Check ID | Covers | How to check | Expected result | Evidence |
| --- | --- | --- | --- | --- |
| `CHK-01` | `REPRO-01`, `RC-01`, `FIX-01` | focused spec / manual procedure | Дефект не воспроизводится, happy path сохранен | PR checks / CI / attached log |

Manual-only exception допустим только с причиной, процедурой и approval ref.
```

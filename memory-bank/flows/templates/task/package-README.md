---
title: TASK-XXX README Template
doc_kind: task
doc_function: template
purpose: "Шаблон routing/index layer для durable managed task package в memory-bank/tasks/TASK-XXX/."
derived_from:
  - ../../routing.md
  - ../../bug-fix.md
  - ../../refactoring.md
  - ../../../dna/frontmatter.md
status: active
audience: humans_and_agents
template_for: managed_task_readme
template_target_path: ../../../tasks/TASK-XXX/README.md
---

# TASK-XXX: Managed Task Package

Этот template используется только для durable `memory-bank/tasks/TASK-XXX/` package. Для issue/PR-only compact profiles `README.md` не нужен.

## Инстанцированный Frontmatter

```yaml
title: "TASK-XXX: Task Package"
doc_kind: task
doc_function: index
purpose: "Routing/index layer для managed non-feature task package."
derived_from:
  - ../../flows/routing.md
status: active
delivery_status: planned
audience: humans_and_agents
workflow_profile: managed-task
```

## Инстанцированное Тело

```markdown
# TASK-XXX: Task Package

## Маршрутизация

| Field | Value |
| --- | --- |
| `kind` | `bugfix` / `refactor` / `chore` |
| `size` | `small` / `medium` |
| `risk` | `low` / `normal` |
| `change_surface` | `single_component` / `multi_component` |
| `contract_change` | `none` |
| `workflow_profile` | `managed-task` |

Promotion check: если появляется capability, contract change, high/critical risk, design reasoning или несколько delivery units, остановиться и повторить routing по `flows/routing.md`.

## Аннотированный индекс

Оставь только route для primary owner этого package. В одном обычном `TASK-XXX/` должен быть `bugfix.md` или `refactor.md`, но не оба сразу.

- [`bugfix.md`](bugfix.md)
  Читать, когда package является managed bugfix.
  Отвечает на вопрос: какой symptom/root cause/fix boundary и regression coverage действуют.

- [`refactor.md`](refactor.md)
  Читать, когда package является managed refactor/chore.
  Отвечает на вопрос: какой intent, behavior invariants, change surface и verification действуют.

## Ссылки

- GitHub issue:
- PR:
- Related incident/ADR/runbook:
```

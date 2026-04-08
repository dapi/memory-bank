---
doc_kind: governance
doc_function: canonical
purpose: SSoT implementation и правила dependency tree. Отвечает на вопрос — кто владеет каким фактом.
derived_from:
  - principles.md
status: active
---
# Document Governance

`Governed document` — markdown-файл в `memory-bank/` с валидным YAML frontmatter. Принцип SSoT определён в [principles.md](principles.md). Этот документ описывает механизм его исполнения.

## SSoT Implementation

1. Authoritative только `active`-документы. `draft` не переопределяет `active`.
2. Среди допустимых по status побеждает upstream: сначала `canonical_for`, затем dependency tree.
3. Публикационный статус (`status`) отделён от lifecycle сущности (`delivery_status`, `decision_status`).

## Source Dependency Tree

1. Поле `derived_from` перечисляет прямые upstream-документы. Authority течёт upstream → downstream.
2. Корневой документ — `principles.md`, не имеет `derived_from`. Для каждого `active` non-root документа `derived_from` обязательно.
3. Циклические зависимости запрещены. Изменение upstream может потребовать обновления downstream.

## Governance-specific Frontmatter Fields

Governance-документы (DNA, flows) используют дополнительные поля, не входящие в общую schema (`frontmatter.md`):

| Поле | Значения | Назначение |
|-|-|-|
| `doc_kind` | `governance`, `project` | Тип документа. Governance — meta-правила, project — domain/ops |
| `doc_function` | `canonical`, `index`, `template` | Роль: canonical owner факта, навигационный индекс или шаблон |

Эти поля обязательны для governance-документов и не требуются в domain/ops/engineering документах.

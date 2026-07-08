---
title: Feature Packages Index
doc_kind: feature
doc_function: index
purpose: Навигация по instantiated feature packages. Читать, чтобы найти существующую delivery-единицу или понять, где создавать новую.
derived_from:
  - ../dna/governance.md
  - ../flows/feature-flow.md
status: active
audience: humans_and_agents
---

# Feature Packages Index

Каталог `memory-bank/features/` хранит instantiated feature packages вида `FT-XXX/`.

## Rules

- Каждый package создается по правилам из [`../flows/feature-flow.md`](../flows/feature-flow.md).
- Bootstrap package начинается с `README.md` и `brief.md`; после `Problem Ready` в него добавляется `design.md`, если `brief.md` фиксирует `Design required: yes`; `implementation-plan.md` появляется после готовности нужных upstream owners.
- Для bootstrap и downstream-документов используй шаблоны из [`../flows/templates/feature/`](../flows/templates/feature/).
- Если работа требует roadmap, risk register и нескольких delivery subissues, сначала создай или обнови epic package в [`../epics/README.md`](../epics/README.md).
- По умолчанию feature ссылается на общий product context из [`../product/context.md`](../product/context.md), а при изменении предметных правил также на соответствующие документы из [`../domain/README.md`](../domain/README.md).
- Если feature реализует или существенно меняет устойчивый сценарий проекта, она должна ссылаться на соответствующий `UC-*` из [`../use-cases/README.md`](../use-cases/README.md).
- В шаблонном репозитории этот каталог может быть пустым. Это нормально.

## Реестр

| Feature ID | Title | Status | Delivery status | Source |
| --- | --- | --- | --- | --- |
| [`FT-015`](FT-015-pattern-operational-agentic-use-cases/README.md) | Operational and Agentic Use Cases Pattern | `active` | `done` | [GitHub issue #15](https://github.com/dapi/memory-bank/issues/15) |

## Naming

- Базовый формат: `FT-XXX/`
- Вместо `XXX` используй идентификатор, принятый в проекте: issue id, ticket id или другой стабильный ключ
- Один package = одна delivery-единица

## Packages

- [`FT-012`](FT-012/)
  Читать, когда нужно: вести delivery по issue #12 о compact task-flow для bugfix/refactor/chore через `feature-flow`.
  Отвечает на вопрос: какие scope, design, plan и локальные решения управляют добавлением task-flow family.

- [FT-013](FT-013/)
  Читать, когда нужно: открыть feature package для issue 13 про generic workflow routing metrics, decision log и developer brief.

- [`FT-014/`](FT-014/) — delivery package for issue 14: add lightweight epic intake brief template and related epic-flow/template index updates.

- [`FT-015`](FT-015-pattern-operational-agentic-use-cases/README.md)
  Читать, когда нужно: открыть feature package для issue 15 про operational/agentic use case guidance.

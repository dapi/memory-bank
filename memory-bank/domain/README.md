---
title: Domain Documentation Index
doc_kind: domain
doc_function: index
purpose: Навигация по domain-level документации шаблона. Читать для фиксации предметной модели, ubiquitous language, бизнес-правил, состояний, событий и bounded contexts.
derived_from:
  - ../dna/governance.md
status: active
audience: humans_and_agents
---

# Domain Documentation Index

Каталог `memory-bank/domain/` хранит предметную модель проекта: язык домена, бизнес-сущности, правила, состояния, события и bounded contexts. Этот слой описывает то, что должно оставаться истинным независимо от текущей продуктовой инициативы или технической реализации.

Domain-документы не определяют market positioning, product metrics, UI design system, concurrency pattern, deployment config или implementation sequence.

## На Какие Вопросы Отвечает Domain

- Какие понятия существуют в предметной области и что они означают?
- Какие сущности, value objects, actors или aggregates важны для reasoning?
- Какие бизнес-правила и инварианты нельзя нарушать?
- Какие состояния и переходы допустимы?
- Какие domain events являются бизнес-значимыми фактами?
- Где проходят bounded contexts и language boundaries?

## Граница С `product/`

| Layer | Отвечает на вопросы | Не отвечает на вопросы |
| --- | --- | --- |
| `product/` | Зачем существует продукт, для кого он, какие outcomes и metrics важны | Какие domain entities, states, invariants и events существуют |
| `domain/` | Что истинно в предметной области и какие правила обязана соблюдать система | Почему именно эта аудитория приоритетна, как продукт позиционируется, какой roadmap выбран |

Пример:

- Product: "Уменьшить количество ручных операций для segment `SEG-01`".
- Domain: "`Invoice` не может быть marked paid без подтвержденного payment event".

## Граница С Engineering

- `domain/context-map.md` описывает business bounded contexts и language ownership.
- `engineering/architecture.md` описывает code/module boundaries, runtime patterns, concurrency, error handling и configuration ownership.
- Если документ отвечает на вопрос "какое бизнес-правило истинно?", он принадлежит `domain/`.
- Если документ отвечает на вопрос "как это безопасно реализовать в системе?", он принадлежит `engineering/`.

## Аннотированный Индекс

- [Glossary](glossary.md) — ubiquitous language, термины, запрещенные двусмысленности и canonical names.
- [Domain Model](model.md) — ключевые domain concepts, relationships, ownership и model notes.
- [Domain Rules](rules.md) — бизнес-правила, инварианты, policies и rule ownership.
- [States](states.md) — lifecycle states, allowed transitions и terminal states.
- [Events](events.md) — domain events как бизнес-значимые факты и их минимальный contract.
- [Context Map](context-map.md) — bounded contexts, upstream/downstream relations и language boundaries.

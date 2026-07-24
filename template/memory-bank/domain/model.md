---
title: Domain Model
doc_kind: domain
doc_function: canonical
purpose: Каноничное описание ключевых domain concepts, relationships, ownership и model boundaries.
derived_from:
  - ../dna/governance.md
  - glossary.md
status: active
audience: humans_and_agents
canonical_for:
  - domain_model
  - domain_concepts
---

# Domain Model

Этот документ описывает conceptual model предметной области. Он не должен подменять database schema, API contract или code module layout.

## Concepts

| Concept | Kind | Owns / Represents | Key relationships | Notes |
| --- | --- | --- | --- | --- |
| `ConceptName` | entity / value object / actor / aggregate / policy | Что означает | С чем связан | Важные ограничения |

## Relationship Map

Опиши связи на уровне бизнеса, а не таблиц базы данных.

Пример:

- `Order` belongs to `Customer`.
- `Payment` confirms or rejects an attempt to settle `Order`.
- `Refund` references a previously captured `Payment`.

## Concept Ownership

| Concept | Canonical owner | Allowed writers | Allowed readers | Notes |
| --- | --- | --- | --- | --- |
| `ConceptName` | Context или team | Кто может менять state | Кто может читать через public contract | Boundary notes |

## Model Boundaries

- `MB-01` Что сознательно не является domain concept, даже если существует в UI или storage.
- `MB-02` Какой legacy term сохраняется только для compatibility и не должен расширяться.

## Related Documents

- Бизнес-правила фиксируются в [`rules.md`](rules.md).
- Состояния и transitions фиксируются в [`states.md`](states.md).
- Bounded contexts фиксируются в [`context-map.md`](context-map.md).

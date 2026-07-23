---
title: Domain Events
doc_kind: domain
doc_function: canonical
purpose: Каноничное место для domain events как бизнес-значимых фактов, их meaning, producers, consumers и минимального payload contract.
derived_from:
  - ../dna/governance.md
  - model.md
  - rules.md
status: active
audience: humans_and_agents
canonical_for:
  - domain_events
  - business_events
---

# Domain Events

Этот документ описывает события, которые являются значимыми фактами предметной области. Technical logs, analytics events и infrastructure messages живут в engineering/ops/product docs, если у них нет domain meaning.

## Events

| Event ID | Event | Meaning | Producer | Consumers | Minimal facts |
| --- | --- | --- | --- | --- | --- |
| `DE-01` | `DomainEventName` | Что стало истинным | Context / component | Кто реагирует | Какие факты обязательны |

## Event Rules

- Событие называется в прошедшем времени или как факт, который уже произошел.
- Событие не должно означать command или request.
- Если event меняет allowed state transitions, обнови [`states.md`](states.md).
- Если event переносит responsibility между contexts, обнови [`context-map.md`](context-map.md).

## Delivery Semantics

Опиши только business expectations: например, whether duplicate event must be harmless or ordering matters for domain correctness. Technical retry, queue, lock и error handling rules фиксируй в [`../engineering/architecture.md`](../engineering/architecture.md).

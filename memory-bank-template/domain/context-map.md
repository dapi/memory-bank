---
title: Domain Context Map
doc_kind: domain
doc_function: canonical
purpose: Каноничное место для bounded contexts, upstream/downstream relations, language ownership и business integration boundaries.
derived_from:
  - ../dna/governance.md
  - glossary.md
  - model.md
status: active
audience: humans_and_agents
canonical_for:
  - bounded_contexts
  - domain_context_map
---

# Domain Context Map

Этот документ фиксирует business bounded contexts. Он не описывает runtime deployment, package layout или service topology, если они не совпадают с domain boundary.

## Bounded Contexts

| Context | Owns language / rules for | Upstream contexts | Downstream contexts | Must not know |
| --- | --- | --- | --- | --- |
| `context-name` | Какие concepts и rules принадлежат context | От кого зависит | Кто зависит от него | Какие details запрещены |

## Context Relationships

| Relationship ID | Upstream | Downstream | Contract | Notes |
| --- | --- | --- | --- | --- |
| `REL-01` | Context owner of source facts | Context consuming facts | API / event / manual process / policy | Важные ограничения |

## Shared Kernel / Published Language

- `SK-01` Какие terms, value objects или policies общие для нескольких contexts.
- `PL-01` Какой published language обязателен на boundary.

## Boundary Rules

- Context владеет своими domain facts и public contracts.
- Другой context не должен читать или менять internal state в обход published boundary.
- Если technical module boundary отличается от domain boundary, объясни это в [`../engineering/architecture.md`](../engineering/architecture.md).

## Open Boundary Questions

- `OQ-01` Где context ownership пока неясен.
- `OQ-02` Какое legacy coupling требует ADR, migration plan или explicit exception.

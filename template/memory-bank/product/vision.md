---
title: Product Vision
doc_kind: product
doc_function: canonical
purpose: Каноничное место для долгосрочного направления продукта, strategic bets, experience principles и product non-goals.
derived_from:
  - ../dna/governance.md
  - context.md
status: active
audience: humans_and_agents
canonical_for:
  - product_vision
  - product_strategy_principles
---

# Product Vision

Этот документ фиксирует устойчивое направление продукта. Он должен помогать принимать решения между competing features, но не заменяет PRD, roadmap или domain rules.

## Product Promise

Опиши в 1-3 абзацах, какой результат продукт обещает пользователю или customer segment.

## Strategic Bets

| Bet ID | Bet | Why now | Evidence | Review cadence |
| --- | --- | --- | --- | --- |
| `BET-01` | Что считаем важной ставкой | Почему это актуально | На чем основано | Когда пересматриваем |

## Experience Principles

- `XP-01` Какой принцип должен сохраняться во всех ключевых product surfaces.
- `XP-02` Какой trade-off продукт делает осознанно.

## Product Non-Goals

- `PNG-01` Что продукт сознательно не пытается решать.
- `PNG-02` Какую аудиторию, use case или business model не оптимизируем сейчас.

## Decision Rules

Опиши правила, которые помогают выбрать между двумя инициативами.

Пример:

- Если две инициативы дают сопоставимый impact, приоритет получает та, которая улучшает core workflow из [`context.md`](context.md).
- Если инициатива требует нового domain concept, сначала обнови [`../domain/model.md`](../domain/model.md) и [`../domain/rules.md`](../domain/rules.md).

## Source Documents

- Strategy memo, board deck, customer research, roadmap artifact или другая ссылка.
- Если источника нет, зафиксируй это явно.

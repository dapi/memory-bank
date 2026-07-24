---
title: Domain Rules
doc_kind: domain
doc_function: canonical
purpose: Каноничное место для бизнес-правил, инвариантов, policies и rule ownership.
derived_from:
  - ../dna/governance.md
  - model.md
status: active
audience: humans_and_agents
canonical_for:
  - domain_rules
  - domain_invariants
---

# Domain Rules

Этот документ фиксирует правила предметной области, которые обязана соблюдать любая реализация. Он не описывает UI behavior, test plan или technical exception handling, если они не являются частью business rule.

## Invariants

| Rule ID | Rule | Applies to | Why it exists | Source |
| --- | --- | --- | --- | --- |
| `DR-01` | Что всегда должно быть истинно | Concept / context | Почему правило важно | SME / policy / PRD / unknown |

## Policies

| Policy ID | Policy | Input | Output / Verdict | Owner |
| --- | --- | --- | --- | --- |
| `POL-01` | Как принимается бизнес-решение | Какие факты нужны | Какой verdict получается | Context / team |

## Cross-Context Rules

- `XDR-01` Правило, которое требует координации нескольких bounded contexts.
- `XDR-02` Правило, где один context обязан использовать public event/API другого context.

## Rule Change Policy

- Если feature меняет domain invariant, обнови этот документ до или вместе с feature `brief.md` / required `design.md`.
- Если правило локально только для одной delivery-единицы, держи его в feature package, пока оно не станет shared domain rule.
- Если правило является архитектурным решением, фиксируй его в ADR и ссылайся отсюда.

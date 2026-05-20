---
title: Product Roadmap
doc_kind: product
doc_function: canonical
purpose: Каноничное место для product themes, bets, horizons и dependencies без превращения roadmap в feature backlog.
derived_from:
  - ../dna/governance.md
  - context.md
  - vision.md
  - metrics.md
status: active
audience: humans_and_agents
canonical_for:
  - product_roadmap
  - product_themes
---

# Product Roadmap

Этот документ описывает направление и sequencing продуктовых тем. Он не должен становиться списком всех feature packages: delivery-единицы живут в [`../features/README.md`](../features/README.md), а инициативы — в [`../prd/README.md`](../prd/README.md).

## Horizons

| Horizon | Theme | Intended outcome | Candidate PRD / Feature | Dependency | Status |
| --- | --- | --- | --- | --- | --- |
| `now` | Что делаем ближайшим горизонтом | Какой outcome ожидаем | `PRD-XXX` / `FT-XXX` / `unknown` | Что должно быть готово | draft / active |
| `next` | Следующая ставка | Какой outcome ожидаем | `PRD-XXX` / `unknown` | Что блокирует | draft |
| `later` | Дальняя тема | Почему это важно | `unknown` | Что нужно узнать | idea |

## Roadmap Rules

- Roadmap theme описывает product intent, а не implementation plan.
- Если тема требует нескольких delivery slices, создай PRD и перечисли downstream features там.
- Если тема меняет предметную модель, сначала обнови [`../domain/model.md`](../domain/model.md), [`../domain/rules.md`](../domain/rules.md) или [`../domain/context-map.md`](../domain/context-map.md).

## Open Bets

- `BET-01` Какая ставка еще требует validation.
- `OQ-01` Какой вопрос нужно закрыть, прежде чем переводить тему в PRD или feature package.

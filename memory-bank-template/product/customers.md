---
title: Customers And Users
doc_kind: product
doc_function: canonical
purpose: Каноничное описание customer/user segments, jobs to be done, pains, evidence и assumptions.
derived_from:
  - ../dna/governance.md
  - context.md
status: active
audience: humans_and_agents
canonical_for:
  - product_customers
  - user_segments
  - jobs_to_be_done
---

# Customers And Users

Этот документ описывает людей, команды или организации, для которых создается продукт. Он не определяет domain entities: если customer segment совпадает по названию с domain concept, различай product-смысл и domain-смысл явно.

## Segments

| Segment ID | Segment | Job To Be Done | Current Pain | Success Signal | Evidence |
| --- | --- | --- | --- | --- | --- |
| `SEG-01` | Кто это | Какую работу пытается выполнить | Что мешает сейчас | Что покажет улучшение | Ссылка или `unknown` |

## Users And Actors

| Actor ID | Actor | Uses product how | Decision power | Notes |
| --- | --- | --- | --- | --- |
| `ACT-01` | Роль пользователя | Где и как взаимодействует с продуктом | Buyer / admin / operator / end user | Важные ограничения |

Если actor становится участником устойчивого сценария, use case фиксируй в [`../use-cases/README.md`](../use-cases/README.md).

## Research Inputs

- Customer interviews, support tickets, sales notes, analytics cohorts или usability studies.
- Если evidence пока нет, пометь assumption как `unvalidated`.

## Assumptions

- `ASM-01` Какое предположение о customer/user пока не подтверждено.
- `ASM-02` Какое предположение влияет на product priority или scope.

## Must Not Assume

- `NA-01` Какую потребность, сегмент или behavior нельзя молча додумывать без evidence.

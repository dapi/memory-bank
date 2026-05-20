---
title: Product Documentation Index
doc_kind: product
doc_function: index
purpose: Навигация по product-level документации шаблона. Читать, чтобы понять зачем существует продукт, для кого он создается и как измеряется успех.
derived_from:
  - ../dna/governance.md
status: active
audience: humans_and_agents
---

# Product Documentation Index

Каталог `memory-bank/product/` хранит устойчивый продуктовый контекст проекта: why, users, outcomes, metrics, positioning и roadmap. Этот слой помогает не повторять общий product background в PRD, use cases и feature packages.

Product-документы не определяют предметную модель, архитектуру реализации, feature acceptance criteria или execution sequence.

## На Какие Вопросы Отвечает Product

- Зачем существует продукт или платформа?
- Для кого он создается: customers, users, segments, actors?
- Какие customer jobs, pains и outcomes важны?
- Какие метрики показывают успех на уровне продукта?
- Как продукт позиционируется относительно альтернатив?
- Какие themes, bets или roadmap horizons направляют дальнейшую работу?

## Граница С `domain/`

| Layer | Отвечает на вопросы | Не отвечает на вопросы |
| --- | --- | --- |
| `product/` | Why, for whom, what outcome, how success is measured, how product is positioned | Какие domain entities существуют, какие инварианты обязательны, как устроена реализация |
| `domain/` | Какие понятия, правила, состояния, события и bounded contexts существуют в предметной области | Зачем бизнесу эта инициатива, какие market segments приоритетны, какие каналы продвижения выбраны |

Пример:

- Product: "Сократить время оператора на обработку заявки и увеличить долю self-service completion".
- Domain: "`Application` не может перейти в `approved`, пока обязательные checks не имеют финальный verdict".

## Граница С PRD

- `product/` — project-wide и long-lived knowledge base.
- `prd/PRD-XXX-short-name.md` — initiative-specific wrapper: какую продуктовую проблему берем в работу сейчас, для каких пользователей и с каким scope.
- Если документ только повторяет общий context, customers или metrics, обнови `product/`, а не заводи новый PRD.

## Аннотированный Индекс

- [Product Context](context.md) — общий продуктовый контекст, ключевые workflows, product constraints и source documents.
- [Vision](vision.md) — долгосрочное направление продукта, strategic bets, experience principles и non-goals.
- [Customers](customers.md) — customer/user segments, jobs to be done, pains, evidence и assumptions.
- [Metrics](metrics.md) — product metrics, baselines, targets, measurement ownership и instrumentation constraints.
- [Marketing](marketing.md) — positioning, messaging, channels, competitive alternatives и launch constraints.
- [Roadmap](roadmap.md) — product themes, bets, horizons и зависимости без превращения в feature backlog.

---
title: Product Context
doc_kind: product
doc_function: canonical
purpose: Каноничное project-wide описание продукта, проблемного пространства и top-level outcomes. Читать перед PRD, use cases и feature briefs, чтобы не повторять общий контекст в каждой delivery-единице.
derived_from:
  - ../dna/governance.md
status: active
audience: humans_and_agents
canonical_for:
  - project_product_context
  - product_problem_space
  - top_level_outcomes
must_not_define:
  - domain_model
  - domain_invariants
  - implementation_sequence
  - architecture_decision
---

# Product Context

Этот документ фиксирует общий продуктовый контекст проекта. Downstream-документы должны ссылаться на него, а не переписывать один и тот же background каждый раз.

PRD, если он нужен, уточняет отдельную инициативу относительно уже зафиксированного project-wide контекста.

## Boundary With PRD And Domain

- `product/context.md` — общий для всего проекта контекст: продукт, пользователи, ключевые product workflows, top-level outcomes и устойчивые product constraints.
- `prd/PRD-XXX-short-name.md` — инициативный слой: какая именно продуктовая проблема берется в работу сейчас, для каких пользователей и с каким scope.
- `domain/` — предметная модель: language, entities, states, invariants, events и bounded contexts, которые должны оставаться истинными независимо от текущей инициативы.
- Если новый документ просто повторяет общий фон проекта и не вводит initiative-specific scope, PRD создавать не нужно.

## Product Context

Опиши проект в 2-4 коротких абзацах:

- кто основные customers и users;
- какую задачу продукт помогает решать;
- почему существующее решение недостаточно;
- какие продуктовые границы у системы или платформы.

Пример:

> Команда поддерживает внутреннюю SaaS-платформу для операционной автоматизации. Пользователи ожидают предсказуемые workflows, прозрачные статусы и быстрый доступ к критичным действиям. Любая новая feature должна либо сокращать операционную нагрузку, либо уменьшать риск ошибок, либо ускорять путь пользователя к целевому результату.

## Core Product Workflows

- `WF-01` Ключевой пользовательский поток номер один.
- `WF-02` Ключевой пользовательский поток номер два.
- `WF-03` Внутренний или операционный поток, который важно не сломать.

Если workflow становится устойчивым canonical scenario с trigger, preconditions, main flow и postconditions, заведи отдельный `UC-*` в [`../use-cases/README.md`](../use-cases/README.md).

## Top-Level Outcomes

Подробные definitions и ownership метрик фиксируй в [`metrics.md`](metrics.md). Здесь оставь только краткий executive summary.

| Metric ID | Metric | Baseline | Target | Measurement method |
| --- | --- | --- | --- | --- |
| `MET-01` | Что считаем успехом на уровне продукта | Текущее состояние | Желаемый уровень | Как измеряем |

## Product Constraints

- `PCON-01` Ограничение продукта, рынка, customer promise или go-to-market, которое влияет на downstream-фичи.
- `PCON-02` Ограничение compliance, интеграций или customer operations, если оно задает продуктовую границу.

Domain-level invariants и state rules фиксируй в [`../domain/rules.md`](../domain/rules.md) и [`../domain/states.md`](../domain/states.md).

## Source Documents

- Добавь сюда ссылки на strategy docs, roadmap, customer research, analytics dashboards или другие upstream-артефакты, если они существуют.
- Если upstream-источников пока нет, так и напиши, не выдумывай их.

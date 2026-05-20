---
title: Frontend Engineering
doc_kind: engineering
doc_function: canonical
purpose: Шаблон описания UI-поверхностей, frontend stack, component boundaries, design system integration и i18n-слоя.
derived_from:
  - ../dna/governance.md
  - ../product/context.md
status: active
audience: humans_and_agents
---

# Frontend Engineering

Этот документ должен описывать реальные UI-поверхности downstream-проекта. Если в системе нет отдельного frontend-слоя, сократи документ до минимально полезного набора правил.

Product-level experience principles живут в [`../product/vision.md`](../product/vision.md). Domain language и rules живут в [`../domain/`](../domain/README.md). Здесь фиксируй engineering contract для UI.

## UI Surfaces

Опиши основные интерфейсы системы.

Пример:

- public web;
- internal backoffice;
- mobile app;
- embedded widgets;
- shared component library.

Для каждой поверхности полезно зафиксировать:

- где лежит код;
- какой stack используется;
- где проходит boundary с backend;
- что считается canonical owner для design decisions.

## Component And Styling Rules

Опиши проектные правила по UI-компонентам:

- используется ли единая design system;
- где живут shared components;
- можно ли создавать ad hoc UI без общего компонента;
- какой слой владеет токенами темы, spacing, typography и states.

Пример записи:

- новые UI-элементы сначала ищут место в `packages/ui`;
- локальный CSS допустим только внутри feature boundary;
- сложная интерактивность требует ADR или явного архитектурного решения.

## Interaction Patterns

Опиши canonical pattern для интерактивности: server-rendered UI, SPA, islands, HTMX/Turbo-like подход, native mobile и т.д.

Вместо project-specific выбора можно использовать шаблонную формулировку:

- для новых feature используй текущий основной interactive stack;
- не смешивай два конкурирующих паттерна без явного основания;
- если проект живет в переходном состоянии между стеками, зафиксируй migration rule и allowed exceptions.

## Localization

Документируй:

- откуда берутся переводы;
- как они попадают в UI;
- где кэшируются или versionируются;
- как добавлять новые ключи и кто владеет fallback behavior.

Если в проекте есть несколько источников переводов, зафиксируй приоритеты и merge order.

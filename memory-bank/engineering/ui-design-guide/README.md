---
title: UI Design Guide Index
doc_kind: engineering
doc_function: index
purpose: Project-level навигация по shared и surface-specific UI references. Читать, чтобы найти existing components, helper APIs, examples, screenshots и source paths для нужной UI surface.
derived_from:
  - ../../dna/governance.md
  - ../frontend.md
status: active
audience: humans_and_agents
canonical_for:
  - project_ui_design_guide_routing
must_not_define:
  - product_requirements
  - domain_rules
  - frontend_architecture_contract
  - feature_interface_requirements
  - implementation_source_of_truth
  - implementation_sequence
---

# UI Design Guide

Этот каталог маршрутизирует к project-level references по существующему UI kit. При адаптации заполни заготовки для реальных UI surfaces проекта, переведи их в `status: active` и удали неприменимые файлы вместе со ссылками из этого index.

## Ownership

- [`../frontend.md`](../frontend.md) владеет frontend stack, boundaries и обязательными engineering rules.
- Product и domain documents владеют product intent, business language и state semantics.
- `features/FT-XXX/ui-reference/README.md` описывает interface change конкретной feature.
- Код владеет фактическими component APIs, signatures и behavior. Этот guide владеет только curated discovery map; перед изменением проверяй source paths и examples по текущему checkout.

## Organization By UI Surface

Не смешивай в одном документе surfaces с разными component libraries, interaction patterns, release boundaries или owners. Каталог уже содержит draft-заготовки:

- `public-web.md` — public website и customer-facing flows;
- `admin.md` — operator/admin UI;
- `mobile.md` — native или mobile-specific UI;
- `shared-components.md` — только действительно shared components и tokens, которые не принадлежат одной surface.

Если в проекте одна компактная UI surface и весь material помещается в [`../frontend.md`](../frontend.md), удали все surface-заготовки и оставь в этом index короткую запись, что UI conventions покрыты `frontend.md`.

## Аннотированный Индекс

- [`public-web.md`](public-web.md) — draft reference для customer-facing web UI: routes, responsive behavior, public forms и accessibility patterns.
- [`admin.md`](admin.md) — draft reference для operator/admin UI: dense workflows, permissions, tables и bulk actions.
- [`mobile.md`](mobile.md) — draft reference для native или mobile-specific UI: navigation, lifecycle, offline и platform patterns.
- [`shared-components.md`](shared-components.md) — draft reference для UI assets, helpers и tokens, которые действительно используются несколькими surfaces.

При адаптации перепиши аннотации под реальный проект и удали ссылки на неприменимые заготовки.

## Surface Document Contract

Каждый surface document должен иметь governed frontmatter с `doc_kind: engineering`, `doc_function: reference`, `derived_from` на этот index и `../frontend.md`, а также только нужные секции из списка:

- component/pattern catalog с existing uses и source paths;
- forms, validation и error presentation;
- actions, navigation и interaction states;
- tables, collections, empty/loading/error states;
- visual labels со ссылками на semantic owners;
- helper APIs, representative examples и screenshots;
- agent entry points: что исследовать перед типовой UI задачей.

Не копируй в surface documents requirements, domain rules, frontend architecture contract, feature-specific interface design или implementation sequence.

## Maintenance

Обновляй соответствующий surface document, когда shared component, helper API, representative example или source path добавлен, удален или materially changed. Если запись не удается подтвердить по коду, исправь или удали ее до использования guide как implementation context.

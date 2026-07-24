---
title: Public Web UI Guide
doc_kind: engineering
doc_function: reference
purpose: Draft-заготовка reference по existing components, helpers, examples и source paths customer-facing web UI.
derived_from:
  - README.md
  - ../frontend.md
status: draft
audience: humans_and_agents
must_not_define:
  - product_requirements
  - domain_rules
  - frontend_architecture_contract
  - feature_interface_requirements
  - implementation_sequence
---

# Public Web UI Guide

Адаптируй этот reference, если проект имеет customer-facing web UI. Если surface отсутствует или полностью покрыта [`../frontend.md`](../frontend.md), удали файл и ссылку из [`README.md`](README.md).

## Surface And Entry Points

- Реальные routes, layout entry points и source roots.
- Responsive, accessibility, localization и browser-support context.

## Components And Patterns

| Component / pattern | Existing use | Source paths | Examples / screenshots | Owner rule |
| --- | --- | --- | --- | --- |
| Замени реальным component | Где и когда он используется | Реальные paths | Linkable examples | Ссылка на owner rule |

## Forms, Actions And Navigation

- Form helpers, validation и error presentation.
- Primary/secondary actions, loading/disabled/confirmation states.
- Navigation, redirects и return paths.

## Agent Entry Points

- Что исследовать перед изменением public web UI.
- Какие tests, examples и screenshots считаются representative.

## Maintenance

Код владеет implementation truth. Подтверждай paths и APIs по current checkout и обновляй этот reference при materially changed reusable patterns.

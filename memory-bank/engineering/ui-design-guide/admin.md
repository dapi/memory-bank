---
title: Admin UI Guide
doc_kind: engineering
doc_function: reference
purpose: Draft-заготовка reference по existing components, helpers, examples и source paths operator/admin UI.
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
  - authorization_policy
  - implementation_sequence
---

# Admin UI Guide

Адаптируй этот reference, если проект имеет operator/admin UI. Если surface отсутствует или полностью покрыта [`../frontend.md`](../frontend.md), удали файл и ссылку из [`README.md`](README.md).

## Surface And Entry Points

- Реальные admin routes, layout entry points и source roots.
- Roles и permissions только как presentation context со ссылками на canonical policy owners.

## Components And Dense Workflows

| Component / pattern | Existing use | Source paths | Examples / screenshots | Owner rule |
| --- | --- | --- | --- | --- |
| Замени реальным component | Где и когда он используется | Реальные paths | Linkable examples | Ссылка на owner rule |

## Tables, Filters And Bulk Actions

- Sorting, filtering, pagination и saved-view patterns.
- Empty/loading/error states, selection и bulk-action confirmation.
- Forms, validation и destructive-action safeguards.

## Agent Entry Points

- Что исследовать перед изменением admin UI.
- Какие roles, states, tests и screenshots нужны для representative review.

## Maintenance

Код владеет implementation truth. Подтверждай paths и APIs по current checkout и обновляй этот reference при materially changed reusable patterns.

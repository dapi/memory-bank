---
title: Shared UI Components Guide
doc_kind: engineering
doc_function: reference
purpose: Draft-заготовка reference по UI components, helpers и tokens, которые реально переиспользуются несколькими UI surfaces.
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

# Shared UI Components Guide

Адаптируй этот reference только для assets, которые реально используются несколькими UI surfaces. Surface-local components оставляй в `public-web.md`, `admin.md`, `mobile.md` или другом соответствующем reference. Если shared layer отсутствует, удали файл и ссылку из [`README.md`](README.md).

## Consumers And Boundaries

- Какие UI surfaces потребляют shared layer.
- Где лежат package/source roots и как проходит ownership boundary.

## Components, Helpers And Tokens

| Asset | Consumers | Source paths | Examples / screenshots | Usage constraints |
| --- | --- | --- | --- | --- |
| Замени реальным asset | Реальные surfaces | Реальные paths | Linkable examples | Ссылка на owner rule |

## Adoption And Deprecation

- Как подключать existing shared assets без копирования.
- Как распознать deprecated component/helper и какой active replacement использовать.

## Agent Entry Points

- Что исследовать перед созданием нового shared component или helper.
- Какие existing uses и tests показывают intended usage.

## Maintenance

Код владеет implementation truth. Подтверждай paths и APIs по current checkout и обновляй этот reference при materially changed reusable patterns.

---
title: Mobile UI Guide
doc_kind: engineering
doc_function: reference
purpose: Draft-заготовка reference по existing components, helpers, examples и source paths native или mobile-specific UI.
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

# Mobile UI Guide

Адаптируй этот reference, если проект имеет native или mobile-specific UI. Если surface отсутствует или полностью покрыта [`../frontend.md`](../frontend.md), удали файл и ссылку из [`README.md`](README.md).

## Platforms And Entry Points

- Реальные platforms, navigation roots, source roots и supported form factors.
- Platform-specific boundaries и shared cross-platform layer.

## Components And Navigation

| Component / pattern | Platform coverage | Source paths | Examples / screenshots | Owner rule |
| --- | --- | --- | --- | --- |
| Замени реальным component | iOS / Android / shared | Реальные paths | Linkable examples | Ссылка на owner rule |

## Lifecycle And Device States

- Loading, offline, background/foreground и interrupted-flow behavior.
- Permissions, keyboard, safe-area и accessibility presentation.
- Forms, actions и navigation transitions.

## Agent Entry Points

- Что исследовать перед изменением mobile UI.
- Какие devices, tests и screenshots считаются representative.

## Maintenance

Код владеет implementation truth. Подтверждай paths и APIs по current checkout и обновляй этот reference при materially changed reusable patterns.

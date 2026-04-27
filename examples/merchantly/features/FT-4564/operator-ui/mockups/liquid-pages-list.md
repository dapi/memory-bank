---
title: "FT-4564 Mockup: Liquid Pages List"
doc_kind: feature-support
doc_function: reference
purpose: "Low-fidelity markdown mockup каталога FT-4564 Liquid-шаблонов публичных страниц."
derived_from:
  - ../README.md
  - ../../feature.md
  - ../../spec.md
status: active
audience: humans_and_agents
---

# FT-4564 Mockup: Liquid Pages List

## Назначение

- показать discoverable entrypoint для public page templates;
- показать bounded catalog `Liquid pages`;
- сделать visible разницу между `configured state` и `effective verdict`;
- не дать оператору вручную создавать произвольные FT-4564 keys.

## Wireframe

```text
+--------------------------------------------------------------------------------------------------+
| Operator / Контент / Liquid-шаблоны публичных страниц                                           |
+--------------------------------------------------------------------------------------------------+
| Liquid pages                                                                                     |
+--------------------------------------------------------------------------------------------------+
| Liquid-шаблоны публичных страниц                                                                 |
| Pages are system-owned and bounded to FT-4564.                                                    |
+--------------------------------------------------------------------------------------------------+
| LIQUID PAGES                                                                                     |
+--------------------------------------------------------------------------------------------------+
| Page                  | Scenario                         | Configured | Template | Effective      |
|-----------------------+----------------------------------+------------+----------+----------------|
| Детали заказа         | public order page               | Liquid on  | Draft    | Liquid         |
| key: show             | order required                  |            | saved    |                |
| actions: [Disable] [Edit] [Preview]                                                           |
|--------------------------------------------------------------------------------------------------|
| Успешная оплата       | orders/paid + payment/show(ok)  | Liquid on  | Draft    | orders/paid:   |
| key: paid             | order optional on callback flow |            | saved    | Liquid         |
|                       |                                  |            |          | callback:      |
|                       |                                  |            |          | Liquid         |
| actions: [Disable] [Edit] [Preview]                                                           |
|--------------------------------------------------------------------------------------------------|
| Неудачная оплата      | payment/show(failure)           | Liquid on  | Missing  | Baseline       |
| key: failed           | order optional                  |            | body     | fallback       |
| actions: [Disable] [Edit] [Preview]                                                           |
|--------------------------------------------------------------------------------------------------|
| Заказ создан          | post-create order page          | Baseline   | Never    | Baseline       |
| key: created          | order required                  |            | edited   |                |
| actions: [Enable] [Edit] [Preview]                                                            |
|--------------------------------------------------------------------------------------------------|
| Заказ отменен         | canceled order state page       | Liquid on  | Draft    | Liquid         |
| key: canceled         | order required                  |            | saved    |                |
| actions: [Disable] [Edit] [Preview]                                                           |
+--------------------------------------------------------------------------------------------------+
| Callout: Liquid pages are system-owned FT-4564 templates.                                         |
+--------------------------------------------------------------------------------------------------+
```

## Что должно читаться с экрана

- `Liquid pages` всегда показывает fixed catalog из пяти keys.
- `paid` не схлопывается в synthetic aggregate status.
- `template body missing` виден отдельно от `Liquid disabled`.

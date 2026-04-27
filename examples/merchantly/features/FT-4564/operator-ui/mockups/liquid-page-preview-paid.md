---
title: "FT-4564 Mockup: Paid Preview Tab"
doc_kind: feature-support
doc_function: reference
purpose: "Low-fidelity markdown mockup preview tab для multi-surface page family paid."
derived_from:
  - ../README.md
  - ../../feature.md
  - ../../spec.md
status: active
audience: humans_and_agents
---

# FT-4564 Mockup: Paid Preview Tab

## Назначение

- показать preview для multi-surface key `paid`;
- не скрывать выбор surface за неявной магией;
- не auto-pick-ать случайный live order;
- честно показать, что callback success flow не гарантирует `order`.

## Wireframe

```text
+--------------------------------------------------------------------------------------------------+
| [Шаблон] [Предпросмотр]* [Переменные и хелперы]                                                  |
+--------------------------------------------------------------------------------------------------+
| Surface selector: (•) orders/paid   ( ) payment/show(state=success)                             |
| Context selector: [Choose preview context ▼]                                                     |
+--------------------------------------------------------------------------------------------------+
| Empty state until context selected:                                                              |
| Preview requires explicit context. For callback success flow, order is optional.                 |
| [Select context]                                                                                 |
+--------------------------------------------------------------------------------------------------+
| Direct public iframe src: /orders/:external_id?...&view=paid                                     |
| If payment/show selected: /payments/page_template_preview/success?...                           |
+--------------------------------------------------------------------------------------------------+
| Public render frame                                                                              |
| ------------------------------------------------------------------------------------------------ |
| Оплата прошла успешно                                                                            |
| Спасибо за покупку.                                                                              |
| [Вернуться в магазин]                                                                            |
| ------------------------------------------------------------------------------------------------ |
+--------------------------------------------------------------------------------------------------+
```

## Что должно читаться с экрана

- preview для `paid` начинается с explicit surface selection;
- до выбора context UI показывает empty state, а не произвольный runtime snapshot;
- для callback success flow UI отдельно напоминает, что `order` optional;
- payment/show preview не строится через route конкретной платёжки;
- preview frame показывает direct public render URL из `spec.md`, без `/operator/shop` и без вложенной операторской.

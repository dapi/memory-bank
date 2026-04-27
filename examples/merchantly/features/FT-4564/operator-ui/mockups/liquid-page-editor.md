---
title: "FT-4564 Mockup: Liquid Page Editor"
doc_kind: feature-support
doc_function: reference
purpose: "Low-fidelity markdown mockup экрана управления одной FT-4564 Liquid-страницей."
derived_from:
  - ../README.md
  - ../../feature.md
  - ../../spec.md
status: active
audience: humans_and_agents
---

# FT-4564 Mockup: Liquid Page Editor

## Назначение

- дать один editor surface для page key;
- объединить template body, enable/disable, preview и context reference;
- явно показать per-surface semantics для multi-surface key `paid`;
- не скрывать, что `order` на payment result pages optional.

## Wireframe

```text
+--------------------------------------------------------------------------------------------------+
| Operator / Контент / Liquid-шаблоны публичных страниц / paid                                      |
+--------------------------------------------------------------------------------------------------+
| [Back to list]                                                            [Disable] [Preview]   |
+--------------------------------------------------------------------------------------------------+
| Страница успешной оплаты                                                                  key: paid |
| Customer sees this after successful payment. Multi-surface page family.                         |
| Configured: Liquid enabled                                                                      |
+--------------------------------------------------------------------------------------------------+
| COVERED SURFACES                                                                                 |
| - orders/paid                    -> effective verdict: Liquid                                   |
| - payment/show(state=success)    -> effective verdict: Liquid                                   |
+--------------------------------------------------------------------------------------------------+
| [Шаблон]* [Предпросмотр] [Переменные и хелперы]                                                  |
+--------------------------------------------------------------------------------------------------+
| STATUS BAR                                                                                       |
| Toggle: [ON] Enable Liquid                                                                       |
| Badges: [Configured: Liquid enabled] [Template: Draft saved] [Runtime: mixed by surface]        |
+--------------------------------------------------------------------------------------------------+
| TEMPLATE BODY                                                                                    |
|--------------------------------------------------------------------------------------------------|
| {% if order %}                                                                                   |
|   <h1>Спасибо за оплату заказа #{{ order.id }}</h1>                                              |
| {% else %}                                                                                       |
|   <h1>Оплата прошла успешно</h1>                                                                 |
| {% endif %}                                                                                      |
|                                                                                                  |
| <a href="{{ vendor.home_url }}">Вернуться в магазин</a>                                          |
|--------------------------------------------------------------------------------------------------|
| [Save] [Preview saved version] [Disable Liquid]                                                  |
+--------------------------------------------------------------------------------------------------+
| CONTEXT REFERENCE                                                                                |
| Variable          | Availability | Meaning                                                       |
|-------------------+--------------+---------------------------------------------------------------|
| vendor            | required     | store identity and links                                      |
| order             | optional     | present on orders/paid, not guaranteed on callback surface   |
| outcome.state     | required     | success/failure metadata                                     |
| helpers           | required     | standard liquid helpers                                      |
+--------------------------------------------------------------------------------------------------+
```

## Что должно читаться с экрана

- editing, preview и context reference относятся к одному page key и не разнесены по разным разделам;
- per-surface verdict для `paid` показан в header/status area;
- optional nature of `order` для payment result family видна прямо в context reference.

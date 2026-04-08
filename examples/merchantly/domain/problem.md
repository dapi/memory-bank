---
title: Project Problem Statement
doc_kind: domain
doc_function: canonical
purpose: Каноничное описание продукта Merchantly, проблемного пространства и top-level outcomes. Читать перед PRD и feature-спеками.
derived_from:
  - ../dna/governance.md
status: active
audience: humans_and_agents
canonical_for:
  - project_problem_statement
  - product_context
  - top_level_outcomes
---

# Project Problem Statement

Этот документ фиксирует общий продуктовый контекст Merchantly. PRD и feature-документы должны ссылаться на него, а не заново пересказывать базовый background проекта.

## Product Context

Merchantly помогает продавцам запускать и поддерживать интернет-магазины на hosted-платформе. Для бизнеса критичны стабильная витрина, предсказуемый checkout, корректные платежи и возможность быстро вносить изменения в каталог, доставку, настройки и контент.

У продукта как минимум две устойчивые пользовательские поверхности. Покупатель взаимодействует с публичной витриной и оформлением заказа, а продавец или оператор работает в backoffice-панели с товарами, заказами, доставкой, оплатами и операционными настройками. Любое изменение должно учитывать обе стороны потока, даже если код меняется только в одном слое.

Merchantly развивается как многосоставная платформа: монолит, публичный React frontend, операторский UI на legacy-стеке, внешние платежные интеграции и tenant-specific настройки. Поэтому локально простое изменение может создать regressions в checkout, переводах, биллинге или vendor-specific поведении, если не держать в голове общие инварианты продукта.

## Core Workflows

- `WF-01` Продавец настраивает магазин, каталог, способы оплаты и доставки так, чтобы витрина оставалась корректной для его tenant.
- `WF-02` Покупатель просматривает витрину, добавляет товары в корзину, оформляет заказ и успешно проходит оплату.
- `WF-03` Оператор или support-команда обрабатывает заказы, диагностирует проблемы магазина и решает инциденты без нарушения tenant isolation и без зависимости operator-контекста от admin-контекста.

## Outcomes

| Metric ID | Metric | Baseline | Target | Measurement method |
| --- | --- | --- | --- | --- |
| `MET-01` | Критичный путь storefront -> checkout -> payment остается работоспособным после изменений | Стабильность контролируется через smoke и production monitoring, regressions дорогие | Изменения не деградируют оформление заказа и оплату в поддерживаемых сценариях | Release smoke, targeted regression coverage, production monitoring |
| `MET-02` | Операторские сценарии по заказам, каталогу и настройкам остаются предсказуемыми | Legacy UI и смешанный стек повышают риск случайных regressions | Новые изменения не ломают существующие operator workflows и не смешивают контексты | Acceptance scenarios, локальные проверки, support feedback |
| `MET-03` | Tenant-specific настройки, переводы и интеграции применяются предсказуемо | Конфигурация распределена между кодом, YAML, БД и внешними сервисами | Изменения явно учитывают vendor-specific behavior и не создают silent misconfiguration | Config review, targeted tests, staged verification |

## Constraints

- `PCON-01` Merchantly multi-tenant: поведение магазина, переводы, настройки оплаты и доставки зависят от конкретного vendor и не должны случайно протекать между tenants.
- `PCON-02` Контексты Admin и Operator изолированы: нельзя смешивать контроллеры, helper methods, views и implicit assumptions между ними.
- `PCON-03` Платежи, биллинг и job execution завязаны на внешние системы и concurrency constraints; для таких изменений нужны idempotency, корректные блокировки и аккуратная диагностика ошибок.

## Source Documents

- [Architecture Patterns](architecture.md)
- [Frontend](frontend.md)
- [CloudPayments](../ops/payments.md)
- [Production & Stage Access](../ops/production.md)

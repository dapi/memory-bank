---
title: Merchantly Glossary
doc_kind: domain
doc_function: glossary
purpose: Устойчивый словарь Merchantly-терминов. Читать, чтобы одинаково понимать продуктовые и технические сущности в feature-документах и обсуждениях.
derived_from:
  - problem.md
  - architecture.md
  - frontend.md
  - ../ops/payments.md
status: active
audience: humans_and_agents
---

# Merchantly Glossary

| Term | Meaning | Where it matters |
| --- | --- | --- |
| `Vendor` | Tenant платформы: отдельный продавец или магазин со своими настройками, переводами, платежами и доставкой. | Любые изменения в storefront, checkout, конфигурации и billing должны учитывать tenant isolation. |
| `Storefront` | Публичная витрина магазина для покупателя. | Изменения в каталоге, корзине и checkout влияют на критичный customer path. |
| `Checkout` | Путь покупателя от корзины до успешной оплаты. | Любые regressions здесь критичны для `MET-01` и release smoke. |
| `Operator` | Backoffice-контекст для продавца или support-команды. | Нельзя смешивать его с admin-контекстом в controllers, views и helper methods. |
| `Admin` | Отдельный внутренний административный контекст приложения. | Должен оставаться изолированным от operator surface. |
| `Member` | Пользователь operator-панели, связанный с vendor и ролями доступа. | Важен для owner/contact fallback и operator workflows. |
| `CloudPayments` | Платёжный шлюз Merchantly для charge-операций и связанной диагностики. | Все изменения в оплатах, 3DS и ручной диагностике платежей проходят через него. |
| `CloudKassir` | Кассовый провайдер, получающий inline receipt через CloudPayments. | Используется в системной фискализации по 54-ФЗ. |
| `OrangeData` | Legacy-провайдер фискализации, от которого Merchantly уходит в новых delivery units. | Часто фигурирует как legacy dependency и источник rollback/cleanup задач. |
| `Billing::VendorLock` | Обёртка над transaction-level advisory lock для критичных биллинговых секций. | Нужна, когда изменение затрагивает списания, инвойсы или recovery path. |
| `limits_concurrency` | Механизм SolidQueue для ограничения параллельного выполнения jobs. | Используется для job concurrency вместо session-level advisory locks. |
| `Bugsnag` | Система error tracking для runtime и background jobs. | Ручное `Bugsnag.notify` допустимо только когда нужен дополнительный domain context. |
| `Stimulus` | Предпочтительный JS-подход для новой интерактивности в operator UI. | Нужен при добавлении dropdowns, tabs, inline-edit и другого поведения в legacy frontend. |
| `Design Guide` | Локальный набор UI-правил и готовых паттернов operator-панели. | Снижает риск UI regressions и несогласованных решений в HAML/helpers слое. |

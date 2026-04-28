---
title: Architecture Patterns
doc_kind: domain
doc_function: canonical
purpose: Изоляция контекстов, блокировки, обработка ошибок в jobs. Читать при написании кода, затрагивающего Admin/Operator, биллинг, jobs.
derived_from:
  - ../../../memory-bank/dna/governance.md
status: active
audience: humans_and_agents
---

# Architecture Patterns

## Изоляция контекстов Admin и Operator

Контексты Admin (Administrate) и Operator — **разные изолированные области**. Они НЕ должны зависеть друг от друга.

| Контекст | Контроллер | Views | Helper methods |
|-|-|-|-|
| Admin | `Admin::ApplicationController` | `app/views/admin/` | `current_admin_user` |
| Operator | `Operator::BaseController` | `app/views/operator/` | `current_operator`, `current_member` |

**Правила:**
- Views из `app/views/operator/` **НЕ должны** использовать `current_admin_user` или другие admin-специфичные методы
- Views из `app/views/admin/` **НЕ должны** рендерить partials из `app/views/operator/`
- Общая view — в `app/views/shared/`, контекст через locals
- Не используй `defined?(current_admin_user)` — признак нарушения изоляции

## Блокировки (Locking)

### Billing locks (advisory)

Для биллинговых операций — `Billing::VendorLock` с transaction-level advisory locks (совместимо с PgBouncer transaction mode):

```ruby
Billing::VendorLock.with_lock(vendor) do
  # критическая секция — внутри DB-транзакции
end
```

**Важно:** `with_lock` оборачивает блок в DB-транзакцию. Все DB-операции атомарны. Внешние API-вызовы (CloudPayments) **НЕ** откатываются — нужен idempotent recovery.

### Job concurrency (SolidQueue)

Для предотвращения параллельного выполнения — `limits_concurrency` из SolidQueue:

```ruby
class MyJob < ApplicationJob
  limits_concurrency key: ->(vendor_id) { vendor_id }, to: 1

  def perform(vendor_id)
    # только один экземпляр с данным vendor_id одновременно
  end
end
```

**Важно:** НЕ используй `with_advisory_lock` напрямую:
- Биллинг → `Billing::VendorLock.with_lock`
- Конкурентность jobs → `limits_concurrency`
- Session-level advisory locks несовместимы с PgBouncer transaction mode

## Bugsnag и обработка ошибок в Jobs

### НЕ оборачивай Bugsnag.notify в begin/rescue
`Bugsnag.notify` — инфраструктурный вызов. Если Bugsnag упадёт — пусть ошибка поднимется наружу.

### НЕ добавляй ручной rescue+Bugsnag.notify в ActiveJob'ах
`ApplicationJob` уже обрабатывает ошибки через `retry_on StandardError` с Bugsnag-нотификацией при исчерпании ретраев. Ручной `rescue StandardError => e; Bugsnag.notify(e); raise` — **дублирование** (двойные нотификации).

**Когда ручной Bugsnag оправдан:** код добавляет **специфичный domain-контекст** (invoice_id, vendor_id), который нельзя получить из стандартного `arguments`.

## Configuration

Конфигурация через:
- `config/configs/application_config.rb` — класс `ApplicationConfig` (Anyway::Config)
- `config/application.yml` — YAML с defaults и per-environment секциями
- ENV переменные с префиксом `APP_` (автоматически подхватываются)

**При изменении конфигурации:**
1. Добавь/измени параметр в `ApplicationConfig`
2. Добавь/измени default в `config/application.yml`
3. **Обнови** `memory-bank/ops/config.md`

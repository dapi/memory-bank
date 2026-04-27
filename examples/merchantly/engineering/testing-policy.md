---
title: Testing Policy
doc_kind: engineering
doc_function: canonical
purpose: "Описывает testing policy репозитория: обязательность test case design, требования к automated regression coverage и допустимые manual-only gaps."
derived_from:
  - ../../../memory-bank/dna/governance.md
  - ../../../memory-bank/flows/feature-flow.md
status: active
canonical_for:
  - repository_testing_policy
  - feature_test_case_inventory_rules
  - automated_test_requirements
  - sufficient_test_coverage_definition
  - manual_only_verification_exceptions
  - simplify_review_discipline
  - verification_context_separation
must_not_define:
  - feature_acceptance_criteria
  - feature_scope
audience: humans_and_agents
---

# Testing Policy

## Stack

- **Framework:** RSpec (`spec/`)
- **Data:** Fixtures only (FactoryBot запрещён), `spec/fixtures` + `FixtureBuilders`
- **Assertions:** `I18n.t` для переводов, стандартные RSpec matchers
- **Оптимизация:** TestProf (`let_it_be`, `before_all`)
- **Запуск:** `make spec` / `RAILS_ENV=test bundle exec rspec spec`
- **CI:** GitHub Actions

## Core Rules

- Любое изменение поведения, которое можно проверить детерминированно, обязано получить automated regression coverage.
- Любой новый или измененный contract (API endpoint, schema, service interface) обязан получить request- или integration-level spec.
- Любой bugfix обязан добавить regression test на воспроизводимый сценарий.
- Required automated tests считаются закрывающими риск только если они проходят локально и в CI.
- Manual-only verify допустим только как явное исключение и не заменяет automated coverage там, где automation реалистична.

## Ownership Split

- Canonical test cases delivery-единицы задаются в `feature.md` через `SC-*`, feature-specific `NEG-*`, `CHK-*` и `EVID-*`.
- `solution.md` владеет selected design, to-be C4 architecture model, `CTR-*`, `FM-*` и локальными `RB-*`, но не подменяет canonical verify contract.
- `implementation-plan.md` владеет только стратегией исполнения: какие spec-файлы будут добавлены или обновлены, какие gaps временно остаются manual-only и почему.

## Feature Flow Expectations

Canonical lifecycle gates живут в [../../../memory-bank/flows/feature-flow.md](../../../memory-bank/flows/feature-flow.md):

- к `Problem Ready` `feature.md` уже фиксирует test case inventory (минимум один `SC-*`);
- к `Solution Ready` `solution.md` фиксирует selected design, to-be C4 architecture model, contracts и solution-level failure modes;
- к `Plan Ready` `implementation-plan.md` содержит `Test Strategy` с planned spec coverage и manual-only gaps;
- к `Done` required specs добавлены, `make spec` зелёный локально и в CI.

## Что Считается Sufficient Coverage

- Покрыт основной changed behavior и ближайший regression path.
- Покрыты новые или измененные API endpoints и service contracts.
- Покрыты критичные failure modes из `FM-*` в `solution.md`, bug history или acceptance risks.
- Покрыты feature-specific negative/edge scenarios, если они меняют verdict.
- Процент line coverage сам по себе недостаточен — нужен scenario- и contract-level coverage.

## Когда Manual-Only Допустим

- Сценарий зависит от live infra, платёжных систем (CloudPayments), внешних API или human оценки UI.
- Для каждого manual-only gap: причина, ручная процедура, owner follow-up.
- Если manual-only gap оставляет без regression protection критичный путь, feature не считается завершённой.

## Simplify Review

Отдельный проход верификации после функционального тестирования. Цель: убедиться, что реализация минимально сложна.

- Выполняется после прохождения specs, но до closure gate.
- Паттерны: premature abstractions, глубокая вложенность (> 3 уровней), дублирование логики, dead code, overengineering.
- Три похожие строки лучше premature abstraction. Абстракция оправдана при ≥ 3 повторениях.

## Verification Context Separation

Разные этапы верификации — отдельные проходы:

1. **Функциональная верификация** — specs проходят, acceptance scenarios покрыты
2. **Simplify review** — код минимально сложен
3. **Acceptance test** — end-to-end по `SC-*`

Для small features допустимо в одной сессии, но simplify review не пропускается.

## Spec Conventions (Merchantly-specific)

- Новые specs в `spec/`, подключая `require_relative` для `rails_helper`
- Fixtures only — **никаких фабрик**
- Database-affecting specs требуют openbill reset (`make test` делает автоматически)
- Текст проверять через `I18n.t`, данные брать из `spec/fixtures` и `FixtureBuilders`
- Используй `let_it_be` и `before_all` (TestProf) для оптимизации тяжёлых setup

## Fixture Library

Все тестовые данные — в `spec/fixtures/`. FactoryBot удалён из проекта.

### Базовые сущности
- `vendor_base`, `vendor_crossborder` (`vendors.yml`) — каноничные продавцы
- Workflow: `workflow_new`, `workflow_processing` (`workflow_states.yml`)
- Платежи/доставки: `vendor_payment_cash`, `vendor_payment_online`, `vendor_delivery_pickup`, `vendor_delivery_courier`

### Каталог и цены
- Категории: `category_root`, `category_sale` (`categories.yml`)
- Товары: `product_ring` + `product_item_ring_default`, `product_bracelet` + `product_item_bracelet_default`
- Цены: `price_kind_retail` + `product_ring_price` / `product_bracelet_price`

### Клиентский путь
- `client_anna`, `cart_default`, `operator_main`, `member_manager`, `role_manager`
- Заказ: `order_basic` + `order_item_ring`
- Корзина: `cart_item_ring`

### Интеграции и шаблоны
- AmoCRM: `vendor_amo_crm_575` (`vendor_amo_crms.yml`)
- Шаблоны: `vendor_template_default`, `vendor_template_source` (`vendor_templates.yml`)

### Правила использования
- Подключай через стандартные helpers: `vendors(:vendor_base)`, `orders(:order_basic)`
- Для мутабельных кейсов — `dup_record` (см. `spec/support/fixture_builders.rb`)
- Новые ключи: `snake_case + короткий суффикс сценария`
- **Не мутируй фикстуры напрямую** — клонируй через `dup_record` или `Model.create!(fixture.attributes.except('id'))`

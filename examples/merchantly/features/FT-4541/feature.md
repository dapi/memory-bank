---
title: "FT-4541: Переход фискализации с OrangeData на CloudKassir"
doc_kind: feature
doc_function: canonical
purpose: "Canonical feature-документ: замена нерабочего OrangeData на CloudKassir (inline receipt через CloudPayments) для фискализации платежей магазинов платформы по 54-ФЗ."
derived_from:
  - ../../../../memory-bank/flows/feature-flow.md
  - ../../engineering/testing-policy.md
status: active
delivery_status: planned
audience: humans_and_agents
must_not_define:
  - implementation_sequence
---

# FT-4541: Переход фискализации с OrangeData на CloudKassir

## What

### Problem

Платформа Kiiiosk обязана по 54-ФЗ отправлять фискальные чеки при приёме оплаты от магазинов (подписка, SMS, пополнение баланса). Сервис OrangeData перестал работать ~28.03.2026 из-за истечения фискального накопителя (ФН `7280440500087355`). Чеки не отправляются — нарушение 54-ФЗ с риском штрафов ФНС.

### Outcome

| Metric ID | Metric | Baseline | Target | Measurement method |
|-|-|-|-|-|
| `MET-01` | Доля платежей с фискальным чеком | 0% (OrangeData не работает) | 100% платежей через CloudPayments | Проверка в ЛК CloudKassir: все charge имеют receipt |
| `MET-02` | Чеки коррекции за период простоя | Не сформированы | Сформированы и отправлены | Наличие чека коррекции в ОФД + уведомление ФНС |

### Scope

- `REQ-01` При рекуррентном платеже (`Billing::CloudPaymentsRecurrentCharge#charge!`) в CloudPayments передаётся inline `receipt` с параметрами ФФД.
- `REQ-02` При прямой оплате картой (`System::PaymentsController#pay`) в CloudPayments передаётся inline `receipt`. Receipt передаётся в первоначальном `cards.charge`, не в `post3ds` (3DS — подтверждение, не новый charge).
- `REQ-03` Email/phone для receipt берётся по fallback-цепочке: `vendor.owners.first` → `vendor.email`/`vendor.phone` → Bugsnag alert.
- `REQ-04` Новые параметры конфигурации в `ApplicationConfig`: `system_receipt_inn`, `system_receipt_taxation_system`, `system_receipt_vat`, `send_system_receipts`.
- `REQ-05` Feature flag `send_system_receipts` управляет отправкой receipt (замена `send_orange_data_receipts`).
- `REQ-06` Сформировать чеки коррекции за период простоя OrangeData (~28.03.2026 — дата включения CloudKassir) через rake task.

### Non-Scope

- `NS-01` Фискализация покупок покупатель→магазин — ответственность магазинов, не платформы.
- `NS-02` Доработка gem `BrandyMint/cloud_payments` — если `json_receipt` параметр не поддерживается, доработка gem выносится в отдельную задачу (gem принимает произвольный hash, доработка маловероятна).
- `NS-03` Изменение логики биллинга, расчёта сумм, рекуррентных платежей — только добавление receipt.
- `NS-04` Миграции БД — фича не требует изменения schema.
- `NS-05` Чеки возврата при рефанде — отдельная задача. В коде нет вызовов `payments.refund`; рефанды выполняются вручную через ЛК CloudPayments. При настроенной кассе CloudKassir чеки возврата формируются CloudPayments автоматически.
- `NS-06` Удаление кода OrangeData (`lib/orange_data/`, `OrangeDataJob`, сертификаты `config/orange_data/`, конфигурация `orange_data.*`) — отдельная delivery unit после подтверждения работы CloudKassir в production.

### Constraints / Assumptions

- `ASM-01` CloudKassir будет зарегистрирована в ФНС и включена в рабочий режим до момента deploy на production. Бизнес-блокер: получение КЭП, регистрация в ЛК nalog.gov.ru, отправка карточки куратору.
- `ASM-02` Gem `cloud_payments` передаёт произвольные ключи в hash charge-запроса (проверено: `tokens.charge(attributes)` принимает hash).
- `ASM-03` CloudPayments при невалидном receipt отклоняет charge целиком (атомарность inline receipt).
- `CON-01` Параметры receipt фиксированы: `TaxationSystem: 1` (УСН Доход), `Vat: VatNone`, `Method: 4` (полный расчёт), `Object: 4` (услуга).
- `DEC-01` Решено: рефанды — out of scope (NS-05). В коде нет программных рефандов; при ручном рефанде через ЛК CloudPayments чек возврата формируется автоматически при настроенной кассе.
- `DEC-02` Решено: fallback при отсутствии контактов — платёж проходит без receipt + Bugsnag alert (крайне редкий edge case, не блокировать оплату).
- `DEC-03` Решено: удаление OrangeData — out of scope (NS-06). Код остаётся до подтверждения работы CloudKassir в production, обеспечивая backout path.
- `NT-01` Не модифицировать существующую логику advisory lock, pre-charge recovery и error handling в `CloudPaymentsRecurrentCharge`. Только добавление receipt в charge hash.

## How

### Solution

Заменить async post-payment OrangeData job на inline `receipt` объект в CloudPayments charge request. CloudPayments сам передаёт данные в CloudKassir при успешной оплате. Это устраняет рассинхронизацию (нет отдельного job) и обеспечивает атомарность: либо платёж + чек, либо ничего.

### Change Surface

| Surface | Type | Why it changes |
|-|-|-|
| `app/modules/billing/cloud_payments_recurrent_charge.rb` | code | Добавление `receipt` в `charge!` hash (строка 94) |
| `app/controllers/system/payments_controller.rb` | code | Добавление `receipt` в `#pay` charge hash (не `#post3ds` — это 3DS confirmation, не charge) |
| `app/services/system_receipt_builder.rb` (новый) | code | Сервис формирования receipt hash для CloudPayments |
| `config/configs/application_config.rb` | config | Новые параметры `system_receipt_*`, `send_system_receipts` |
| `config/application.yml` | config | Значения для новых параметров по средам |
| `ops/config.md` | doc | Документация новых параметров |
| `lib/tasks/fiscal_correction.rake` (новый) | code | Rake task для чеков коррекции |

### Flow

**Приход (charge):**
1. `CloudPaymentsRecurrentCharge#charge!` или `PaymentsController#pay` формирует charge-запрос.
2. Если `ApplicationConfig.send_system_receipts` → `SystemReceiptBuilder.new(invoice:).build` формирует receipt hash.
3. Receipt добавляется как `json_data: { CloudPayments: { CustomerReceipt: receipt } }` в charge hash.
4. CloudPayments при успешном charge передаёт receipt в CloudKassir → ОФД → ФНС.

**Fallback контактов:**
5. `SystemReceiptBuilder` ищет email/phone: `vendor.owners.first.email` → `vendor.owners.first.phone` → `vendor.email` → `vendor.phone` → nil.
6. Если nil — receipt не включается, Bugsnag alert отправляется.

### Contracts

| Contract ID | Input / Output | Producer / Consumer | Notes |
|-|-|-|-|
| `CTR-01` | `json_data.CloudPayments.CustomerReceipt` в charge request | `SystemReceiptBuilder` / CloudPayments API | Формат CloudPayments Online Receipt API v2 |
| `CTR-02` | `ApplicationConfig.send_system_receipts` boolean flag | `application.yml` / charge code | `false` → receipt не добавляется, charge работает как раньше |
| `CTR-03` | `ApplicationConfig.system_receipt_*` параметры | `application.yml` / `SystemReceiptBuilder` | ИНН, СНО, НДС для receipt |

### Failure Modes

- `FM-01` **Невалидный receipt блокирует платёж.** CloudPayments отклоняет charge при невалидном receipt (ASM-03). Митигация: валидация receipt перед отправкой + feature flag для быстрого отката.
- `FM-02` **Отсутствие email/phone у vendor.** Receipt без `customerContact` невалиден по 54-ФЗ. Митигация: fallback-цепочка (REQ-04), при исчерпании — пропуск receipt + Bugsnag (DEC-02).
- `FM-03` **CloudKassir не зарегистрирована / не активна.** Receipt отправляется, но CloudKassir не обрабатывает. Митигация: бизнес-блокер (ASM-01), проверка на stage перед production.

### Rollout / Backout

- `RB-01` **Rollout:** deploy с `send_system_receipts: false` → включить flag → проверить первые чеки в ЛК CloudKassir → оставить включённым.
- `RB-02` **Backout:** выключить `send_system_receipts` → платежи продолжают работать без receipt (как до фичи). Код OrangeData остаётся на месте (NS-06), обеспечивая возможность отката.

### ADR Dependencies

Нет ADR-зависимостей. Фича не вводит архитектурных решений, требующих отдельного ADR — используется существующий CloudPayments gem и ApplicationConfig.

## Verify

### Exit Criteria

- `EC-01` Рекуррентный платёж с `send_system_receipts: true` отправляет receipt с корректными ФФД-параметрами в CloudPayments.
- `EC-02` Прямая оплата картой отправляет receipt аналогично.
- `EC-03` При `send_system_receipts: false` платежи работают без receipt (backward compatibility).
- `EC-04` Feature flag и конфигурация добавлены в `ApplicationConfig` и `application.yml`.
- `EC-05` Rake task для чеков коррекции формирует корректный чек.

### Acceptance Scenarios

- `SC-01` **Happy path: рекуррентный платёж с receipt.** Магазин имеет автосписание. При charge в CloudPayments передаётся `json_data` с `CustomerReceipt` содержащим `Type: Income`, `TaxationSystem: 1`, `Items` с данными invoice, `Email` владельца магазина. Платёж проходит, receipt доставляется в ОФД.
- `SC-02` **Happy path: прямая оплата с receipt.** Оператор магазина оплачивает счёт через форму. Аналогично SC-01 receipt передаётся в charge. При 3DS — receipt передаётся в первоначальном charge (не в post3ds).
- `SC-03` **Fallback контактов.** У `vendor.owners.first` нет email. Receipt содержит `Phone` владельца. Если нет phone — берётся `vendor.email`. Если совсем нет — платёж проходит без receipt, Bugsnag alert.
- `SC-04` **Feature flag выключен.** `send_system_receipts: false`. Charge не содержит `json_data` с receipt. Платёж работает как до фичи.
- `SC-05` **Чек коррекции.** Rake task формирует чек коррекции за период простоя с корректными параметрами.

### Negative / Edge Cases

- `NEG-01` **Invoice без title.** `invoice.title` пустой или nil — receipt `Label` должен иметь fallback (например `invoice.description.truncate(128)`).
- `NEG-02` **Amount = 0.** Нулевой invoice не должен отправлять receipt (нет фискализации для нулевых сумм).

### Traceability matrix

| Requirement ID | Design refs | Acceptance refs | Checks | Evidence IDs |
|-|-|-|-|-|
| `REQ-01` | `ASM-02`, `CON-01`, `CTR-01`, `FM-01`, `NT-01` | `EC-01`, `SC-01` | `CHK-01` | `EVID-01` |
| `REQ-02` | `ASM-02`, `CON-01`, `CTR-01`, `FM-01` | `EC-02`, `SC-02` | `CHK-01` | `EVID-01` |
| `REQ-03` | `DEC-02`, `FM-02`, `CTR-01` | `SC-03` | `CHK-01` | `EVID-01` |
| `REQ-04` | `CTR-02`, `CTR-03` | `EC-04` | `CHK-02` | `EVID-02` |
| `REQ-05` | `CTR-02`, `RB-01`, `RB-02` | `EC-03`, `SC-04` | `CHK-02` | `EVID-02` |
| `REQ-06` | `ASM-01`, `FM-03` | `EC-05`, `SC-05` | `CHK-03` | `EVID-03` |

### Checks

| Check ID | Covers | How to check | Expected result | Evidence path |
|-|-|-|-|-|
| `CHK-01` | `EC-01`, `EC-02`, `EC-03`, `SC-01`–`SC-04`, `NEG-01`, `NEG-02` | `direnv exec . bundle exec rspec spec/modules/billing/cloud_payments_recurrent_charge_spec.rb spec/controllers/system/payments_controller_spec.rb spec/services/system_receipt_builder_spec.rb` | Все тесты зелёные | `artifacts/ft-4541/verify/chk-01/` |
| `CHK-02` | `EC-04`, `SC-04` | `direnv exec . bundle exec rspec spec/configs/application_config_spec.rb` | Конфигурация загружается с дефолтами | `artifacts/ft-4541/verify/chk-02/` |
| `CHK-03` | `EC-05`, `SC-05` | `direnv exec . bundle exec rspec spec/tasks/fiscal_correction_rake_spec.rb` | Rake task формирует корректный receipt | `artifacts/ft-4541/verify/chk-03/` |

### Test matrix

| Check ID | Evidence IDs | Evidence path |
|-|-|-|
| `CHK-01` | `EVID-01` | `artifacts/ft-4541/verify/chk-01/` |
| `CHK-02` | `EVID-02` | `artifacts/ft-4541/verify/chk-02/` |
| `CHK-03` | `EVID-03` | `artifacts/ft-4541/verify/chk-03/` |

### Evidence

- `EVID-01` Лог прохождения RSpec тестов для charge + receipt (рекуррент, прямая оплата, fallback, flag off, edge cases).
- `EVID-02` Лог прохождения RSpec тестов для ApplicationConfig с новыми параметрами.
- `EVID-03` Лог прохождения RSpec тестов для rake task чеков коррекции.

### Evidence contract

| Evidence ID | Artifact | Producer | Path contract | Reused by checks |
|-|-|-|-|-|
| `EVID-01` | RSpec output log | CI / local rspec run | `artifacts/ft-4541/verify/chk-01/` | `CHK-01` |
| `EVID-02` | RSpec output log | CI / local rspec run | `artifacts/ft-4541/verify/chk-02/` | `CHK-02` |
| `EVID-03` | RSpec output log | CI / local rspec run | `artifacts/ft-4541/verify/chk-03/` | `CHK-03` |

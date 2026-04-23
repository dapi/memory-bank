---
title: "FT-4541: Переход фискализации с OrangeData на CloudKassir"
doc_kind: feature
doc_function: canonical
purpose: "Canonical feature-документ problem space: замена нерабочей схемы OrangeData на working flow фискализации системных платежей платформы по 54-ФЗ."
derived_from:
  - ../../../../memory-bank/flows/feature-flow.md
  - ../../engineering/testing-policy.md
status: active
delivery_status: planned
audience: humans_and_agents
must_not_define:
  - selected_design
  - implementation_sequence
---

# FT-4541: Переход фискализации с OrangeData на CloudKassir

## What

### Problem

Платформа Kiiiosk обязана по 54-ФЗ отправлять фискальные чеки при приёме оплаты от магазинов (подписка, SMS, пополнение баланса). OrangeData перестал работать 28.03.2026 из-за истечения фискального накопителя (ФН `7280440500087355`). Чеки по системным платежам не отправляются, что создаёт нарушение 54-ФЗ и риск штрафов ФНС.

### Outcome

| Metric ID | Metric | Baseline | Target | Measurement method |
| --- | --- | --- | --- | --- |
| `MET-01` | Доля системных платежей с фискальным чеком | 0% после отказа OrangeData 28.03.2026 | 100% системных платежей получают чек через CloudKassir | Проверка в ЛК CloudKassir: у каждого charge есть receipt |
| `MET-02` | Чеки коррекции за период простоя | Не сформированы | Сформированы и отправлены за период 28.03.2026 -> дата включения CloudKassir | Наличие чеков коррекции в ОФД и подтверждение в CloudKassir |

### Scope

- `REQ-01` При рекуррентном платеже (`Billing::CloudPaymentsRecurrentCharge#charge!`) в CloudPayments передаётся inline receipt с параметрами ФФД.
- `REQ-02` При прямой оплате картой (`System::PaymentsController#pay`) в CloudPayments передаётся inline receipt для того же системного платежа.
- `REQ-03` Для receipt используется fallback по доступным контактам магазина; при полном отсутствии контактов платёж не должен блокироваться без дополнительного human approval.
- `REQ-04` В `ApplicationConfig` появляются параметры `system_receipt_inn`, `system_receipt_taxation_system`, `system_receipt_vat`, `send_system_receipts`.
- `REQ-05` Отправка receipt управляется feature flag `send_system_receipts`.
- `REQ-06` Появляется операционный путь для чеков коррекции за период простоя между отказом OrangeData и включением CloudKassir.

### Non-Scope

- `NS-01` Фискализация платежей покупатель -> магазин не входит в эту delivery-единицу.
- `NS-02` Изменение логики биллинга, расчёта сумм и существующего платежного lifecycle не входит в scope; фича добавляет только receipt-related behavior.
- `NS-03` Миграции БД не требуются и в scope не входят.
- `NS-04` Возвратные чеки для refund не входят в scope этой фичи.
- `NS-05` Полное удаление кода OrangeData и связанных артефактов не входит в scope этой delivery-единицы.

### Constraints / Assumptions

- `ASM-01` CloudKassir будет зарегистрирована в ФНС и включена в рабочий режим до production rollout.
- `CON-01` Параметры receipt для системных платежей зафиксированы: `TaxationSystem: 1` (УСН Доход), `Vat: VatNone`, `Method: 4` (полный расчёт), `Object: 4` (услуга).
- `INV-01` При `send_system_receipts: false` системные платежи должны оставаться backward-compatible с текущим production behavior.

### Blocking Decisions

На стадии `Problem Ready` unresolved blocking decisions отсутствуют. Все принятые локальные решения по выбранному подходу вынесены в [`solution.md`](solution.md).

## Verify

### Exit Criteria

- `EC-01` Рекуррентный платёж с `send_system_receipts: true` отправляет корректный receipt в CloudPayments / CloudKassir.
- `EC-02` Прямая оплата картой отправляет корректный receipt для того же системного платежа.
- `EC-03` При `send_system_receipts: false` платежи продолжают работать без regressions в текущем сценарии.
- `EC-04` Feature flag и новые параметры конфигурации доступны через `ApplicationConfig`.
- `EC-05` Существует исполнимый путь для формирования чеков коррекции за период простоя.

### Acceptance Scenarios

- `SC-01` **Happy path: рекуррентный платёж с receipt.** Системный платёж проходит, а receipt с корректными ФФД-параметрами передаётся в CloudPayments и появляется в CloudKassir.
- `SC-02` **Happy path: прямая оплата с receipt.** Оператор магазина оплачивает счёт картой, receipt передаётся вместе с системным платежом и не теряется при 3DS-подтверждении.
- `SC-03` **Fallback контактов.** Если приоритетный контакт недоступен, receipt использует следующий доступный контакт; если контактов нет совсем, платёж проходит без receipt и создаётся Bugsnag alert.
- `SC-04` **Feature flag выключен.** При `send_system_receipts: false` системный платёж проходит без receipt и без изменения текущего поведения.
- `SC-05` **Чек коррекции.** Операционная процедура формирует корректный чек коррекции за период простоя.

### Negative / Edge Cases

- `NEG-01` Если invoice не содержит пригодного заголовка, receipt использует безопасный fallback label.
- `NEG-02` Нулевой invoice не должен приводить к отправке receipt.

### Traceability matrix

| Requirement ID | Acceptance refs | Checks | Evidence IDs |
| --- | --- | --- | --- |
| `REQ-01` | `EC-01`, `SC-01` | `CHK-01` | `EVID-01` |
| `REQ-02` | `EC-02`, `SC-02` | `CHK-01` | `EVID-01` |
| `REQ-03` | `EC-03`, `SC-03`, `SC-04` | `CHK-01` | `EVID-01` |
| `REQ-04` | `EC-04` | `CHK-02` | `EVID-02` |
| `REQ-05` | `EC-03`, `SC-04` | `CHK-01`, `CHK-02` | `EVID-01`, `EVID-02` |
| `REQ-06` | `EC-05`, `SC-05` | `CHK-03` | `EVID-03` |

### Checks

| Check ID | Covers | How to check | Expected result | Evidence path |
| --- | --- | --- | --- | --- |
| `CHK-01` | `EC-01`, `EC-02`, `EC-03`, `SC-01`-`SC-04`, `NEG-01`, `NEG-02` | `direnv exec . bundle exec rspec spec/modules/billing/cloud_payments_recurrent_charge_spec.rb spec/controllers/system/payments_controller_spec.rb spec/services/system_receipt_builder_spec.rb` | Тесты системных receipt flows зелёные | `artifacts/ft-4541/verify/chk-01/` |
| `CHK-02` | `EC-04`, `SC-04` | `direnv exec . bundle exec rspec spec/configs/application_config_spec.rb` | Конфигурация с новыми параметрами загружается корректно | `artifacts/ft-4541/verify/chk-02/` |
| `CHK-03` | `EC-05`, `SC-05` | `direnv exec . bundle exec rspec spec/tasks/fiscal_correction_rake_spec.rb` | Операционный путь чеков коррекции формирует корректный receipt | `artifacts/ft-4541/verify/chk-03/` |

### Test matrix

| Check ID | Evidence IDs | Evidence path |
| --- | --- | --- |
| `CHK-01` | `EVID-01` | `artifacts/ft-4541/verify/chk-01/` |
| `CHK-02` | `EVID-02` | `artifacts/ft-4541/verify/chk-02/` |
| `CHK-03` | `EVID-03` | `artifacts/ft-4541/verify/chk-03/` |

### Evidence

- `EVID-01` Лог прохождения RSpec тестов для receipt в recurrent/direct payment flows, fallback контактов и flag-off path.
- `EVID-02` Лог прохождения RSpec тестов для `ApplicationConfig` с новыми параметрами.
- `EVID-03` Лог прохождения RSpec тестов для rake task чеков коррекции.

### Evidence contract

| Evidence ID | Artifact | Producer | Path contract | Reused by checks |
| --- | --- | --- | --- | --- |
| `EVID-01` | RSpec output log | CI / local rspec run | `artifacts/ft-4541/verify/chk-01/` | `CHK-01` |
| `EVID-02` | RSpec output log | CI / local rspec run | `artifacts/ft-4541/verify/chk-02/` | `CHK-02` |
| `EVID-03` | RSpec output log | CI / local rspec run | `artifacts/ft-4541/verify/chk-03/` | `CHK-03` |

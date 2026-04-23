---
title: "FT-4541: Solution"
doc_kind: feature
doc_function: canonical
purpose: "Canonical solution-документ для FT-4541. Фиксирует выбранный design для перехода с OrangeData на CloudKassir без переопределения problem space."
derived_from:
  - feature.md
status: active
audience: humans_and_agents
must_not_define:
  - ft_4541_scope
  - ft_4541_acceptance_criteria
  - ft_4541_delivery_status
  - implementation_sequence
---

# FT-4541: Solution

## Selected Design

- `SOL-01` Системные receipt отправляются inline в исходный CloudPayments charge request вместо отдельного async OrangeData job, чтобы receipt был атомарной частью платежа.
- `SOL-02` Общий `SystemReceiptBuilder` формирует payload receipt для recurrent и direct payment flows с едиными ФФД-параметрами и fallback по контактам.
- `SOL-03` Конфигурация и feature flag `send_system_receipts` разделяют code rollout и business enablement, обеспечивая быстрый backout без удаления нового кода.
- `SOL-04` Отдельный rake task закрывает correction path за период простоя между отказом OrangeData 28.03.2026 и датой включения CloudKassir.

## Requirement Mapping

| Requirement ID | Solution refs | Notes |
| --- | --- | --- |
| `REQ-01` | `SOL-01`, `SOL-02`, `SD-05`, `CTR-01` | Recurrent charge обогащается receipt payload без изменения существующей recovery/error logic |
| `REQ-02` | `SOL-01`, `SOL-02`, `SD-01`, `CTR-01` | Direct card payment получает receipt в исходном charge, а не в `post3ds` |
| `REQ-03` | `SOL-02`, `SD-02`, `FM-02` | Contact fallback и Bugsnag alert являются частью builder-поведения |
| `REQ-04` | `SOL-03`, `CTR-02`, `CTR-03` | Новая config surface даёт canonical source для enablement и receipt params |
| `REQ-05` | `SOL-03`, `CTR-02`, `RB-01`, `RB-02` | Feature flag управляет rollout/backout выбранного решения |
| `REQ-06` | `SOL-04`, `RB-01`, `FM-03` | Correction path отделён от online charge path, но опирается на ту же кассовую инфраструктуру |

## Accepted Local Decisions

- `SD-01` Для прямой оплаты receipt передаётся в исходный `cards.charge`, а не в `post3ds`: 3DS здесь подтверждает уже созданный charge, а не инициирует новый.
- `SD-02` Contact fallback идет по цепочке `vendor.owners.first.email -> vendor.owners.first.phone -> vendor.email -> vendor.phone`; если контактов нет, платёж не блокируется, а создаётся Bugsnag alert.
- `SD-03` Refund receipts остаются вне scope этой фичи: ручной refund path через кабинет CloudPayments не мигрируется в этот delivery slice.
- `SD-04` Код OrangeData и связанные артефакты сохраняются до production validation CloudKassir; backout делается через feature flag, а не через immediate code removal.
- `SD-05` В `CloudPaymentsRecurrentCharge` нельзя трогать advisory lock, pre-charge recovery и существующий error handling: receipt добавляется только в payload charge-запроса.

## Change Surface

| Ref | Surface | Type | Why it changes |
| --- | --- | --- | --- |
| `SOL-01` | `app/modules/billing/cloud_payments_recurrent_charge.rb` | code | Добавление receipt в recurrent charge payload |
| `SOL-01` | `app/controllers/system/payments_controller.rb` | code | Добавление receipt в direct charge payload |
| `SOL-02` | `app/services/system_receipt_builder.rb` | code | Единый builder для system receipt payload |
| `SOL-03` | `config/configs/application_config.rb` | config | Новые параметры `system_receipt_*` и `send_system_receipts` |
| `SOL-03` | `config/application.yml` | config | Значения по средам для новых параметров |
| `SOL-03` | `ops/config.md` | doc | Документация новой конфигурационной поверхности |
| `SOL-04` | `lib/tasks/fiscal_correction.rake` | code | Операционный correction path за период простоя |

## Internal Flow

1. `SOL-01` `CloudPaymentsRecurrentCharge#charge!` или `System::PaymentsController#pay` собирает charge request для системного платежа.
2. `SOL-02` Если `send_system_receipts` включён, `SystemReceiptBuilder` собирает `CustomerReceipt` с каноническими ФФД-параметрами и доступным контактом.
3. `SOL-01` Receipt добавляется в `json_data.CloudPayments.CustomerReceipt` исходного charge request.
4. `SOL-01` При успешном charge CloudPayments передаёт receipt в CloudKassir и далее в ОФД / ФНС.
5. `SD-02` Если пригодного контакта нет, charge остаётся допустимым, receipt пропускается, а система сигнализирует в Bugsnag.
6. `SOL-04` Для downtime period correction receipt формируется отдельной rake task, а не online charge flow.

## Contracts

| Contract ID | Related refs | Input / Output | Producer / Consumer | Notes |
| --- | --- | --- | --- | --- |
| `CTR-01` | `SOL-01`, `SOL-02`, `REQ-01`, `REQ-02` | `json_data.CloudPayments.CustomerReceipt` в charge request | `SystemReceiptBuilder` / CloudPayments API | Формат Online Receipt API для system charges |
| `CTR-02` | `SOL-03`, `REQ-05` | `ApplicationConfig.send_system_receipts` boolean flag | `application.yml` / charge code | `false` -> receipt не добавляется, платеж работает как раньше |
| `CTR-03` | `SOL-02`, `SOL-03`, `REQ-04` | `ApplicationConfig.system_receipt_*` параметры | `application.yml` / `SystemReceiptBuilder` | ИНН, СНО, НДС и другие фиксированные receipt params |

## Failure Modes

- `FM-01` Невалидный receipt блокирует весь charge, потому что inline receipt является частью исходного платежного запроса.
- `FM-02` Полное отсутствие контактных данных делает receipt непригодным для отправки и переводит кейс в Bugsnag alert + no-receipt path.
- `FM-03` Если CloudKassir не зарегистрирована или не активна к моменту rollout, correction path и online receipts не дадут целевой фискальный результат.

## Rollout / Backout

- `RB-01` Rollout: deploy с `send_system_receipts: false` -> включить feature flag -> проверить первые чеки в CloudKassir -> оставить включённым.
- `RB-02` Backout: выключить `send_system_receipts` -> системные платежи продолжают работать без receipt, пока новый кассовый path не будет исправлен.

## ADR Dependencies

Нет ADR-зависимостей. Решение локально для feature package и не вводит нового cross-feature архитектурного baseline.

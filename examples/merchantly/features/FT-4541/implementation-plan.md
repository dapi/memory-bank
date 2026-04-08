---
title: "FT-4541: Implementation Plan"
doc_kind: feature
doc_function: derived
purpose: "Execution-план реализации FT-4541. Фиксирует discovery context, шаги, риски и test strategy без переопределения canonical feature-фактов."
derived_from:
  - feature.md
status: active
audience: humans_and_agents
must_not_define:
  - ft_4541_scope
  - ft_4541_architecture
  - ft_4541_acceptance_criteria
  - ft_4541_blocker_state
---

# План имплементации

## Цель текущего плана

Заменить нерабочую post-payment фискализацию OrangeData на inline receipt через CloudPayments/CloudKassir. Результат: 100% платежей платформы формируют фискальный чек по 54-ФЗ. Rake task формирует чеки коррекции за период простоя.

## Current State / Reference Points

| Path / module | Current role | Why relevant | Reuse / mirror |
|-|-|-|-|
| `app/modules/billing/cloud_payments_recurrent_charge.rb:90-105` | Рекуррентный charge через `CloudPayments.client.payments.tokens.charge(hash)` | Точка вставки receipt в charge hash (`REQ-01`) | Паттерн передачи hash в `tokens.charge` — добавить `json_data` ключ |
| `app/controllers/system/payments_controller.rb:20-34` | Прямая оплата через `CloudPayments.client.payments.cards.charge(hash)` | Точка вставки receipt для прямой оплаты (`REQ-02`) | Аналогично — добавить `json_data` в hash метода `#pay` |
| `app/models/cloud_payment_receipt_data.rb` | Формирование receipt для **магазинных** платежей (order/invoice) | **Reference pattern** для CloudPayments receipt format. Использует `Items`, `taxationSystem`, `email`/`phone` | Зеркалить структуру receipt hash; НЕ reuse напрямую (другой контекст: system vs vendor) |
| `lib/orange_data/system_requestor.rb` | Текущий OrangeData receipt builder для системных платежей | **Reference pattern** для fallback контактов (`client_email_or_phone`) и receipt полей | Скопировать fallback-логику email/phone, адаптировать под CloudPayments формат |
| `config/configs/application_config.rb:123-124` | `orange_data: {}`, `cloud_payments_receipt_data: {}` в attr_config; строка 400: `send_orange_data_receipts: false` | Точка добавления `send_system_receipts` и `system_receipt_*` параметров (`REQ-04`, `REQ-05`) | Паттерн flat-параметров: `send_orange_data_receipts` → аналогично `send_system_receipts` |
| `config/application.yml:748-749` | `cloud_payments_receipt_data.taxation_system: 1` (используется `CloudPaymentReceiptData`) | Reference для значений конфигурации по средам | Добавить `system_receipt_inn`, `system_receipt_taxation_system`, `system_receipt_vat`, `send_system_receipts` |
| `app/jobs/billing/transaction_job.rb:49` | Триггер `OrangeDataJob` после cloudpayments income | **НЕ трогать** — inline receipt заменяет post-payment job. Строка 49 станет dead code после включения flag | `NT-01`: не модифицировать логику OrangeDataJob |
| `app/jobs/orange_data_job.rb` | OrangeData async job | Остаётся на месте (`NS-06`), не трогаем | — |
| `ops/config.md` | Документация конфигурации | Добавить описание новых параметров | Формат существующих записей |

## Test Strategy

| Test surface | Canonical refs | Existing coverage | Planned automated coverage | Required local suites / commands | Required CI suites / jobs | Manual-only gap / justification | Manual-only approval ref |
|-|-|-|-|-|-|-|-|
| `spec/services/system_receipt_builder_spec.rb` (новый) | `REQ-01`-`REQ-03`, `SC-01`-`SC-03`, `NEG-01`, `NEG-02`, `CHK-01` | Отсутствует | Unit-тесты: happy path receipt, fallback email→phone→vendor.email→vendor.phone→nil+Bugsnag, пустой title fallback, amount=0 skip | `direnv exec . bundle exec rspec spec/services/system_receipt_builder_spec.rb` | rspec job | — | `none` |
| `spec/modules/billing/cloud_payments_recurrent_charge_spec.rb` | `REQ-01`, `REQ-05`, `SC-01`, `SC-04`, `CHK-01` | 474 строки, покрыты charge, errors, recovery | Добавить: charge с receipt (flag on), charge без receipt (flag off), проверка `json_data` в hash | `direnv exec . bundle exec rspec spec/modules/billing/cloud_payments_recurrent_charge_spec.rb` | rspec job | — | `none` |
| `spec/controllers/system/payments_controller_spec.rb` | `REQ-02`, `REQ-05`, `SC-02`, `SC-04`, `CHK-01` | 168 строк, покрыты pay/post3ds/success | Добавить: pay с receipt (flag on), pay без receipt (flag off), проверка `json_data` не попадает в post3ds | `direnv exec . bundle exec rspec spec/controllers/system/payments_controller_spec.rb` | rspec job | — | `none` |
| `spec/configs/application_config_spec.rb` (новый, директория `spec/configs/` не существует — создать) | `REQ-04`, `REQ-05`, `SC-04`, `CHK-02` | Отсутствует, директория `spec/configs/` не существует | Проверка что новые параметры загружаются с дефолтами | `direnv exec . bundle exec rspec spec/configs/` | rspec job | — | `none` |
| `spec/tasks/fiscal_correction_rake_spec.rb` (новый, директория `spec/tasks/` не существует — создать) | `REQ-06`, `SC-05`, `CHK-03` | Отсутствует, директория `spec/tasks/` не существует | Rake task формирует чек коррекции с корректными параметрами | `direnv exec . bundle exec rspec spec/tasks/fiscal_correction_rake_spec.rb` | rspec job | — | `none` |
| Production receipt в CloudKassir ЛК | `MET-01`, `EC-01`, `EC-02` | — | — | — | — | Проверка реального чека в ЛК CloudKassir после первого production charge | `AG-02` |

## Open Questions / Ambiguities

| Open Question ID | Question | Why unresolved | Blocks | Default action / escalation owner |
|-|-|-|-|-|
| `OQ-01` | ~~Поддерживает ли gem `cloud_payments` ключ `json_data`?~~ **RESOLVED** | Grounding подтвердил: `tokens.rb:6` → `request(:charge, attributes)` → `Base#request` → HTTP body без фильтрации. Произвольные ключи проходят as-is | — | Закрыт. `PRE-01` выполнен |
| `OQ-02` | Точный формат CloudPayments Online Receipt API v2 для поля `CustomerReceipt` | Документация CloudPayments может отличаться от reference `CloudPaymentReceiptData` | `STEP-02` | Default: использовать формат из `CloudPaymentReceiptData#data` + добавить `Type: Income`. Verify в sandbox. Escalation: @danil |
| `OQ-03` | ~~Какой ИНН использовать для `system_receipt_inn`?~~ **RESOLVED** | ИНН зафиксирован в issue #4541: `212906363506` | — | Закрыт. Значение для `system_receipt_inn: "212906363506"` |
| `OQ-04` | Точные даты периода простоя для чеков коррекции | Начало ~28.03.2026, конец = дата включения CloudKassir. Точные даты нужны для rake task | `STEP-07` | Default: query `OpenbillTransaction` за период. Escalation: @danil |

## Environment Contract

| Area | Contract | Used by | Failure symptom |
|-|-|-|-|
| setup | `direnv exec . bundle install` — gem `cloud_payments` доступен | Все `STEP-*` | `LoadError` при запуске specs |
| test | `direnv exec . bundle exec rspec <path>` — эталонная verify команда | `CHK-01`, `CHK-02`, `CHK-03` | Красные specs = недостоверный verify |
| test | `direnv exec . make test` — полный suite перед PR | `CP-03` | CI fail |
| access | CloudPayments sandbox/test mode — для интеграционной проверки receipt | `STEP-06` | Отсутствие test credentials блокирует manual verify |
| access | CloudKassir ЛК — зарегистрирована и активна (`ASM-01`) | `AG-02` | Чеки не обрабатываются, `FM-03` |

## Preconditions

| Precondition ID | Canonical ref | Required state | Used by steps | Blocks start |
|-|-|-|-|-|
| `PRE-01` | `ASM-02` | Gem `cloud_payments` передаёт произвольные ключи hash в HTTP body (не фильтрует) — **VERIFIED** (`OQ-01` RESOLVED) | `STEP-01`, `STEP-02`, `STEP-03` | no |
| `PRE-02` | `ASM-01` | CloudKassir зарегистрирована в ФНС (бизнес-блокер для production, не для разработки) | `STEP-06`, `AG-02` | no (блокирует только production rollout) |
| `PRE-03` | `CON-01` | Параметры ФФД зафиксированы: TaxationSystem=1, Vat=VatNone, Method=4, Object=4 | `STEP-02` | yes |

## Workstreams

| Workstream | Implements | Result | Owner | Dependencies |
|-|-|-|-|-|
| `WS-1` | `REQ-01`, `REQ-02`, `REQ-03`, `REQ-04`, `REQ-05` | `SystemReceiptBuilder`, config, integration в charge code | agent | `PRE-01`, `PRE-03` |
| `WS-2` | `REQ-06` | Rake task для чеков коррекции | agent | `WS-1` (reuse `SystemReceiptBuilder`), `OQ-04` |

## Approval Gates

| Approval Gate ID | Trigger | Applies to | Why approval is required | Approver / evidence |
|-|-|-|-|-|
| `AG-01` | Перед включением `send_system_receipts: true` на production | `STEP-06` | Невалидный receipt блокирует платёж (`FM-01`). Бизнес-блокер: CloudKassir должна быть активна (`ASM-01`) | @danil / confirmation в PR или chat |
| `AG-02` | Проверка первых production чеков в ЛК CloudKassir | `STEP-06` | Подтверждение что чеки реально доходят до ОФД (`MET-01`) | @danil / screenshot ЛК CloudKassir |

## Порядок работ

| Step ID | Actor | Implements | Goal | Touchpoints | Artifact | Verifies | Evidence IDs | Check command / procedure | Blocked by | Needs approval | Escalate if |
|-|-|-|-|-|-|-|-|-|-|-|-|
| `STEP-01` | agent | `REQ-04`, `REQ-05`, `CTR-02`, `CTR-03` | Добавить параметры конфигурации `system_receipt_inn`, `system_receipt_taxation_system`, `system_receipt_vat`, `send_system_receipts` в `ApplicationConfig` и `application.yml` | `config/configs/application_config.rb`, `config/application.yml`, `ops/config.md` | Обновлённые файлы конфигурации | `CHK-02` | `EVID-02` | `direnv exec . bundle exec rspec spec/configs/` | — | `none` | Config не загружается |
| `STEP-02` | agent | `REQ-01`, `REQ-02`, `REQ-03`, `CTR-01` | Создать `SystemReceiptBuilder` — сервис формирования receipt hash для CloudPayments | `app/services/system_receipt_builder.rb` (новый), `spec/services/system_receipt_builder_spec.rb` (новый) | Сервис + unit tests | `CHK-01` | `EVID-01` | `direnv exec . bundle exec rspec spec/services/system_receipt_builder_spec.rb` | `STEP-01` | `none` | `OQ-02` не снят |
| `STEP-03` | agent | `REQ-01`, `REQ-05`, `NT-01` | Интегрировать receipt в `CloudPaymentsRecurrentCharge#charge!` — добавить `json_data` в hash при `send_system_receipts: true` | `app/modules/billing/cloud_payments_recurrent_charge.rb:90-105`, `spec/modules/billing/cloud_payments_recurrent_charge_spec.rb` | Обновлённый charge + тесты | `CHK-01` | `EVID-01` | `direnv exec . bundle exec rspec spec/modules/billing/cloud_payments_recurrent_charge_spec.rb` | `STEP-02` | `none` | — |
| `STEP-04` | agent | `REQ-02`, `REQ-05` | Интегрировать receipt в `PaymentsController#pay` — добавить `json_data` в hash при `send_system_receipts: true`. НЕ трогать `#post3ds` | `app/controllers/system/payments_controller.rb:20-34`, `spec/controllers/system/payments_controller_spec.rb` | Обновлённый controller + тесты | `CHK-01` | `EVID-01` | `direnv exec . bundle exec rspec spec/controllers/system/payments_controller_spec.rb` | `STEP-02` | `none` | — |
| `STEP-05` | agent | `REQ-06` | Создать rake task `fiscal_correction` для чеков коррекции за период простоя OrangeData | `lib/tasks/fiscal_correction.rake` (новый), `spec/tasks/fiscal_correction_rake_spec.rb` (новый) | Rake task + тесты | `CHK-03` | `EVID-03` | `direnv exec . bundle exec rspec spec/tasks/fiscal_correction_rake_spec.rb` | `STEP-02`, `OQ-04` | `none` | Неизвестны точные даты периода |
| `STEP-06` | human | `RB-01` | Production rollout: deploy с `send_system_receipts: false` → включить flag → проверить чеки в ЛК CloudKassir | `config/application.yml` (production), ЛК CloudKassir | Подтверждение чеков в ОФД | `CHK-01` | `EVID-01` | Визуальная проверка чеков в ЛК CloudKassir | `STEP-03`, `STEP-04`, `PRE-02`, `AG-01` | `AG-01`, `AG-02` | CloudKassir не активна (`FM-03`) |
| `STEP-07` | human | `REQ-06` | Запустить rake task чеков коррекции на production | production console | Чеки коррекции в ОФД | `CHK-03` | `EVID-03` | Проверка чеков коррекции в ЛК CloudKassir | `STEP-05`, `STEP-06`, `OQ-04` | `AG-01` | Период простоя неизвестен |

## Parallelizable work

- `PAR-01` `STEP-03` и `STEP-04` могут выполняться параллельно (разные файлы, оба зависят от `STEP-02`).
- `PAR-02` `STEP-05` может идти параллельно с `STEP-03`/`STEP-04` (отдельный rake task, зависит только от `STEP-02`).

## Checkpoints

| Checkpoint ID | Refs | Condition | Evidence IDs |
|-|-|-|-|
| `CP-01` | `STEP-01`, `CHK-02` | Config параметры загружаются, дефолты корректны | `EVID-02` |
| `CP-02` | `STEP-02`, `STEP-03`, `STEP-04`, `CHK-01` | Все unit/integration тесты для receipt зелёные | `EVID-01` |
| `CP-03` | `STEP-01`-`STEP-05` | `make test` зелёный, RuboCop clean | `EVID-01`, `EVID-02`, `EVID-03` |
| `CP-04` | `STEP-06`, `AG-02` | Первые production чеки видны в ЛК CloudKassir | `EVID-01` |

## Риски исполнения

| Risk ID | Risk | Impact | Mitigation | Trigger |
|-|-|-|-|-|
| `ER-01` | ~~Gem `cloud_payments` фильтрует неизвестные ключи в charge hash~~ **MITIGATED** | ~~Receipt не передаётся в CloudPayments~~ | Grounding подтвердил: gem передаёт произвольные ключи as-is (`OQ-01` RESOLVED) | — |
| `ER-02` | CloudKassir не зарегистрирована к моменту deploy | Чеки отправляются, но не обрабатываются (`FM-03`) | `AG-01` блокирует production rollout до подтверждения `ASM-01` | `PRE-02` не выполнена к дате deploy |
| `ER-03` | Невалидный receipt format блокирует платежи (`FM-01`) | 100% платежей отклоняются CloudPayments | Feature flag `send_system_receipts: false` — мгновенный backout (`RB-02`) | Первый charge после включения flag отклонён |
| `ER-04` | У vendor нет ни owner email/phone, ни vendor email/phone | Платёж проходит без чека, нарушение 54-ФЗ для этого конкретного vendor | Bugsnag alert + ручная коррекция. По `DEC-02` — не блокировать оплату | Bugsnag alert `SystemReceiptBuilder: no contact` |

## Stop conditions / fallback

| Stop ID | Related refs | Trigger | Immediate action | Safe fallback state |
|-|-|-|-|-|
| `STOP-01` | `FM-01`, `RB-02` | Production charge отклонён из-за невалидного receipt | Выключить `send_system_receipts` → платежи работают без receipt | `send_system_receipts: false`, OrangeData код на месте (`NS-06`) |
| `STOP-02` | `OQ-01`, `ER-01` | ~~Gem `cloud_payments` фильтрует `json_data`~~ **MITIGATED** | ~~Остановить разработку~~ Grounding подтвердил: gem передаёт ключи as-is (`OQ-01` RESOLVED) | — |
| `STOP-03` | `FM-03`, `ASM-01` | CloudKassir не активна после deploy | Не включать flag, держать `send_system_receipts: false` | Платежи без receipt (текущее состояние) |
| `STOP-04` | `FM-02`, `DEC-02`, `ER-04` | Систематические Bugsnag alerts "no contact" (>5% платежей без receipt) | Выключить flag, провести аудит vendors без контактных данных, заполнить email/phone | `send_system_receipts: false` до завершения аудита |

## Готово для приемки

- [ ] Все `STEP-01` — `STEP-05` выполнены (code complete)
- [ ] `CP-03` пройден: `make test` зелёный, RuboCop clean
- [ ] `AG-01` получен: подтверждение от @danil на production rollout
- [ ] `STEP-06` выполнен: `send_system_receipts: true` на production
- [ ] `AG-02` получен: чеки видны в ЛК CloudKassir
- [ ] `STEP-07` выполнен: чеки коррекции сформированы
- [ ] Simplify review пройден

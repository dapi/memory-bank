---
title: "FT-4564 Use Cases"
doc_kind: feature
doc_function: index
purpose: "Use cases для FT-4564: customer-facing сценарии Liquid-кастомизации публичных страниц `show`, `paid`, `failed`, `created`, `canceled` и operator authoring scenarios для `VendorPageTemplate`."
derived_from:
  - ../feature.md
  - ../spec.md
  - ../operator-ui/README.md
  - ../../../engineering/adr/0005-ft-4564-order-page-resolution-and-fallback-policy.md
  - ../../../engineering/adr/0006-missing-custom-liquid-template-falls-back-to-existing-baseline-page.md
status: active
audience: humans_and_agents
---

# FT-4564 Use Cases

Use cases ниже описывают только публичные страницы, которые FT-4564 переводит на page-specific Liquid-шаблоны через `VendorPageTemplate`. Документ не описывает другие механизмы content authoring и не делает их частью задачи.

Для `paid` и `failed` business label не означает единственный route/template: concrete render-surface ownership задаётся accepted solution docs и используется при интерпретации use cases ниже. Status-only redirect `/orders/:id?status=failed` остаётся boundary case: финальный `vendor/orders/show` принадлежит `show`, а не `failed`.

Adjacent surface `vendor/orders/payment` (`orders/payment`) intentionally не входит в этот набор use cases: это payment-entry page вне scope FT-4564, даже если она соседствует с post-create flow.

## Overview

Этот набор use cases предназначен для двух целей одновременно:

- как derived companion doc к `feature.md`, который упаковывает пользовательские сценарии в review-friendly форму;
- как вспомогательный mapping `UC -> REQ -> EC -> CHK` для оценки реализации.

Canonical acceptance / test inventory для FT-4564 остаётся в `feature.md` через `SC-*`, `NEG-*`, `CHK-*` и `EVID-*`. Этот документ не подменяет `feature.md`, а даёт удобную пользовательскую проекцию тех же контрактов и заготовку для последующей декомпозиции в test cases.

`UC-H01`-`UC-H05` покрывают независимый per-page opt-in без смены глобального engine. `UC-E01`-`UC-E02` фиксируют boundary conditions и контракт контекста. `UC-ER01`-`UC-ER02` разделяют два разных failure classes: `template body absent` и `runtime render error`. `UC-OP-*` задают operator-facing сценарии, а `TC-OP-*` ниже превращают их в derived test cases с preconditions, steps и expected result.

## Happy Path

| ID | Use Case | Описание |
|---|---|---|
| `UC-H01` | Liquid для публичной страницы заказа `show` | Вендор включает Liquid для `show`, пользователь открывает публичную страницу заказа в regular-state, и система рендерит Liquid-шаблон `show` с данными заказа. |
| `UC-H02` | Liquid для страницы неудачной оплаты `failed` | Вендор включает Liquid для `failed`, пользователь попадает на customer-facing render surface, принадлежащую key `failed`, и система отображает Liquid-кастомизацию `failed`; status-only order-page landing не засчитывается как `failed` surface. |
| `UC-H03` | Liquid для страницы успешной оплаты `paid` | Вендор включает Liquid для `paid`, пользователь попадает на customer-facing render surface, принадлежащую key `paid`, и система отображает Liquid-кастомизацию `paid`; use case должен покрывать и `vendor/orders/paid`, и `vendor/payment/show(state=success)`. |
| `UC-H04` | Liquid для страницы `created` | Вендор включает Liquid для `created`, пользователь проходит post-create flow, и система отображает Liquid-шаблон `created`. |
| `UC-H05` | Liquid для страницы `canceled` | Вендор включает Liquid для `canceled`, пользователь попадает в cancel-state, и система отображает Liquid-шаблон `canceled`. |

## Edge Cases

| ID | Use Case | Описание |
|---|---|---|
| `UC-E01` | Для страницы не включён page-specific Liquid | Если для конкретной страницы Liquid не включён, эта страница использует current baseline page implementation. |
| `UC-E02` | Liquid context contract для opted-in страниц | Если для страницы включён Liquid, шаблон получает `vendor`, `order` при наличии order context, page outcome/state metadata и стандартные helpers; для payment result pages `order` optional, для order-state pages required. |

## Error Cases

| ID | Use Case | Описание |
|---|---|---|
| `UC-ER01` | Liquid выбран, но template body отсутствует | Если для page key выбран Liquid, но template body отсутствует, система делает fallback на existing baseline page implementation для той же страницы. |
| `UC-ER02` | Runtime error во время Liquid render | Если Liquid выбран и во время render возникает ошибка, система следует fail-fast / error-reporting policy из ADR-0005 и не делает тихий fallback на baseline implementation. |

## Mobile Use Cases

| ID | Use Case | Описание |
|---|---|---|
| `—` | Отдельные mobile-specific scenarios не выделяются | FT-4564 меняет server-side resolution/rendering policy и не вводит новый mobile-specific UX contract. Mobile coverage выполняется через общие visual/regression checks, а не через отдельный acceptance scenario. |

## Operator Happy Path

| ID | Use Case | Описание |
|---|---|---|
| `UC-OP-01` | Operator открывает каталог public page templates | Оператор открывает entrypoint FT-4564 и видит bounded catalog из пяти system-owned keys: `show`, `paid`, `failed`, `created`, `canceled`. |
| `UC-OP-02` | Operator включает Liquid для одного page key | Оператор включает Liquid для выбранного key, и configured state меняется только у него без побочных изменений у соседних keys. |
| `UC-OP-03` | Operator выключает Liquid для одного page key | Оператор выключает Liquid для выбранного key, и он возвращается в `Baseline`, не затрагивая другие rows каталога. |
| `UC-OP-04` | Operator открывает editor выбранного key | Оператор открывает editor страницы и видит header со status semantics и tabs `Шаблон`, `Предпросмотр`, `Переменные и хелперы`. |
| `UC-OP-05` | Operator сохраняет валидный Liquid template body | Оператор редактирует и сохраняет Liquid body; после reload editor показывает сохранённый шаблон и актуальный template status. |
| `UC-OP-06` | Operator открывает preview из каталога или editor | Action `Preview` deep-link-ит в `Предпросмотр` tab того же editor-а, а preview iframe показывает direct customer-facing public render surface без `/operator/shop`; для payment-result surface используется provider-neutral `/payments/page_template_preview/:state`. |
| `UC-OP-07` | Operator видит context reference | Editor показывает variables/helpers и честные availability labels, включая `order: required`, `optional` или `unavailable`. |

## Operator Edge Cases

| ID | Use Case | Описание |
|---|---|---|
| `UC-OP-E01` | Operator не может создать ad-hoc Liquid key | Каталог FT-4564 остаётся system-owned: оператор не может вручную создать или удалить произвольный Liquid page key. |
| `UC-OP-E02` | `paid` показывает verdict per surface | Для key `paid` UI показывает отдельные effective verdicts для `vendor/orders/paid` и `vendor/payment/show(state=success)`, а не один synthetic aggregate badge. |
| `UC-OP-E03` | Preview order-aware page требует явный context | Для `show`, `created` и `canceled` preview не auto-pick-ает случайный live order, а требует явный выбор preview context. |
| `UC-OP-E04` | Preview `paid` требует явный выбор surface | Для `paid` оператор сначала выбирает variant surface (`orders/paid` или `payment/show(state=success)`), а затем соответствующий preview context. |
| `UC-OP-E05` | Catalog не имеет empty state | Даже если ни один Liquid template ещё не создан, catalog всё равно показывает все пять rows, меняются только их внутренние статусы. |
| `UC-OP-E06` | Preview payment-result не привязан к платёжке | Для `paid/payment.show(success)` и `failed/payment.show(failure)` preview строится через provider-neutral preview-only route, а не через `/payments/<provider>/success|failure`. |

## Operator Error Cases

| ID | Use Case | Описание |
|---|---|---|
| `UC-OP-ER01` | Liquid включён, но template body отсутствует | Для включённого page key UI явно показывает `template body missing`, а effective verdict остаётся baseline fallback, а не мнимый success. |
| `UC-OP-ER02` | Preview запускается при dirty editor | Если в editor есть несохранённые изменения, preview не использует случайный draft silently: UI требует сначала сохранить шаблон или явно сообщает, что preview строится по последней сохранённой версии. |
| `UC-OP-ER03` | Operator проверяет битый Liquid template | Оператор сохраняет заведомо битый Liquid template и через agreed verify path на opted-in customer-facing surface подтверждает наблюдаемую runtime error semantics, а не тихий fallback на baseline implementation; preview может использоваться как authoring helper, но не является отдельным owner-ом runtime policy. |

## Operator Test Cases

### `TC-OP-01` Fixed catalog public page templates

- Covers: `UC-OP-01`, `UC-OP-E01`, `UC-OP-E05`
- Primary refs: `REQ-12`, `EC-12`, `EC-16`, `SC-15`, `CHK-14`, `CHK-18`
- Preconditions:
  1. Оператор авторизован и имеет доступ к entrypoint FT-4564.
  2. Наличие или отсутствие сохранённых Liquid templates не влияет на старт кейса.
- Steps:
  1. Открыть catalog public page templates.
  2. Проверить, что catalog содержит ровно keys `show`, `paid`, `failed`, `created`, `canceled`.
  3. Проверить отсутствие UI для ручного создания или удаления FT-4564 Liquid keys.
  4. Проверить, что catalog не исчезает даже если templates ещё не создавались.
- Expected result:
  1. Entry point показывает фиксированный system-owned catalog из пяти rows.
  2. Для Liquid keys нет ad-hoc create/delete flow.
  3. Пустой state catalog не возникает; меняются только row statuses.
- Automation candidate: `automated`

### `TC-OP-02` Enable Liquid для одного page key

- Covers: `UC-OP-02`
- Primary refs: `REQ-13`, `EC-13`, `CHK-15`
- Preconditions:
  1. Выбран page key с configured mode `Baseline`.
  2. Остальные keys имеют известные стартовые configured states.
- Steps:
  1. Открыть catalog public page templates.
  2. Нажать `Enable` для одного выбранного key.
  3. Сохранить изменение, если UI требует submit.
  4. Перезагрузить страницу или повторно открыть catalog.
  5. Сравнить configured states всех rows с исходным snapshot.
- Expected result:
  1. Только выбранный key меняет configured mode на `Liquid enabled`.
  2. Остальные keys не меняют configured mode.
  3. Effective status считается отдельно от configured mode и не подменяет его.
- Automation candidate: `automated`

### `TC-OP-03` Disable Liquid для одного page key

- Covers: `UC-OP-03`
- Primary refs: `REQ-13`, `EC-13`, `CHK-15`
- Preconditions:
  1. Выбран page key с configured mode `Liquid enabled`.
  2. Остальные keys имеют известные стартовые configured states.
- Steps:
  1. Открыть catalog public page templates.
  2. Нажать `Disable` для выбранного key.
  3. Сохранить изменение, если UI требует submit.
  4. Перезагрузить страницу или повторно открыть catalog.
  5. Сравнить configured states всех rows с исходным snapshot.
- Expected result:
  1. Только выбранный key возвращается в `Baseline`.
  2. Соседние keys не меняются.
  3. Status semantics остаётся honest и не скрывает baseline verdict.
- Automation candidate: `automated`

### `TC-OP-04` Multi-surface verdict для `paid` показывается раздельно

- Covers: `UC-OP-E02`
- Primary refs: `REQ-13`, `EC-13`, `CHK-15`
- Preconditions:
  1. Для `paid` доступны обе covered surfaces: `vendor/orders/paid` и `vendor/payment/show(state=success)`.
  2. Подготовлено состояние, в котором verdict surfaces можно различить или как минимум показать отдельно.
- Steps:
  1. Открыть catalog public page templates.
  2. Найти row `paid`.
  3. Проверить effective status breakdown для обеих surfaces.
  4. Открыть editor `paid` и повторно проверить surface breakdown в header/status area.
- Expected result:
  1. UI показывает две отдельные surface verdicts для `paid`.
  2. UI не схлопывает surfaces в один synthetic aggregate badge.
  3. Если verdicts расходятся, divergence отображается явно.
- Automation candidate: `automated`

### `TC-OP-05` Editor shell и context reference для выбранного key

- Covers: `UC-OP-04`, `UC-OP-07`
- Primary refs: `REQ-14`, `REQ-15`, `EC-14`, `EC-15`, `CHK-17`
- Preconditions:
  1. Catalog public page templates доступен.
  2. Выбран page key для открытия editor-а.
- Steps:
  1. Открыть editor выбранного key через action `Edit`.
  2. Проверить header: title, stable key, configured/effective state.
  3. Проверить наличие tab-ов `Шаблон`, `Предпросмотр`, `Переменные и хелперы`.
  4. Открыть tab `Переменные и хелперы`.
  5. Проверить наличие variables/helpers и availability labels.
- Expected result:
  1. Editor открывается как dedicated surface для одного key.
  2. Header и tabs соответствуют contract из operator UI docs.
  3. Context reference показывает helpers и honest labels `required` / `optional` / `unavailable`.
- Automation candidate: `automated`

### `TC-OP-06` Save valid Liquid template body

- Covers: `UC-OP-05`
- Primary refs: `REQ-14`, `EC-14`, `CHK-16`, `CHK-19`
- Preconditions:
  1. Оператор находится в editor выбранного key.
  2. Подготовлен валидный Liquid body для сохранения.
- Steps:
  1. Открыть tab `Шаблон`.
  2. Вставить или изменить валидный Liquid body.
  3. Нажать `Save`.
  4. Перезагрузить editor.
  5. Проверить сохранённое содержимое и template status.
- Expected result:
  1. Сохранение проходит без потери выбранного key.
  2. После reload editor показывает последнюю сохранённую версию body.
  3. Template status больше не выглядит как `Never edited`; при непустом body не должен показываться `template body missing`.
- Automation candidate: `mixed`

### `TC-OP-07` Preview deep-link, path contract и явный context selection

- Covers: `UC-OP-06`, `UC-OP-E03`, `UC-OP-E04`, `UC-OP-E06`
- Primary refs: `REQ-14`, `REQ-15`, `EC-14`, `EC-15`, `CHK-16`, `CHK-17`, `CHK-19`
- Preconditions:
  1. Для выбранного key сохранён template body.
  2. Для order-aware preview доступен deterministic preview context.
- Steps:
  1. В catalog нажать `Preview` для `show`.
  2. Проверить, что открылся tab `Предпросмотр` того же editor-а.
  3. Проверить, что preview не auto-pick-ает случайный live order и требует явный context.
  4. Повторить сценарий для `paid`.
  5. Для `paid` сначала выбрать surface variant, затем соответствующий context.
  6. Для `failed` открыть payment-result preview.
- Expected result:
  1. `Preview` deep-link-ит в editor tab, а не на detached screen.
  2. Для order-aware pages preview требует явный context.
  3. Для `paid` preview требует сначала surface selection, потом context selection.
  4. Preview iframe показывает direct public render surface без operator sidebar, publish controls или `/operator/shop`.
  5. Replayable order-state preview строится как `/orders/:external_id` с preview query и `view=show|canceled|paid`; `created` preview использует тот же path с `view=created` только как preview-only forced render of post-create surface.
  6. Payment-result preview строится как `/payments/page_template_preview/success` или `/payments/page_template_preview/failure` с preview query, без `/payments/<provider>/success|failure`.
- Automation candidate: `mixed`

### `TC-OP-08` Visible missing-template state в catalog и editor

- Covers: `UC-OP-ER01`
- Primary refs: `REQ-09`, `REQ-13`, `EC-06`, `EC-13`, `CHK-07`, `CHK-15`
- Preconditions:
  1. Page key переведён в `Liquid enabled`.
  2. Для выбранного key отсутствует сохранённый template body.
- Steps:
  1. Открыть catalog public page templates.
  2. Проверить row выбранного key.
  3. Открыть editor того же key.
  4. Проверить warnings/status area.
- Expected result:
  1. Catalog и editor явно показывают `template body missing`.
  2. UI не изображает такое состояние как успешный Liquid render verdict.
  3. Effective status остаётся совместимым с baseline fallback policy.
- Automation candidate: `automated`

### `TC-OP-09` Dirty editor не preview-ится silently

- Covers: `UC-OP-ER02`
- Primary refs: `REQ-14`, `EC-14`, `CHK-16`
- Preconditions:
  1. В editor уже существует сохранённая версия template body.
  2. Оператор внёс несохранённые изменения.
- Steps:
  1. Открыть editor.
  2. Изменить template body без сохранения.
  3. Перейти в `Предпросмотр` или нажать `Preview`.
- Expected result:
  1. UI не preview-ит случайный unsaved draft silently.
  2. Либо показывается требование сначала сохранить изменения, либо явно сообщается, что preview строится по последней сохранённой версии.
  3. Поведение одинаково интерпретируемо из catalog и editor.
- Automation candidate: `manual`

### `TC-OP-10` Broken Liquid template остаётся наблюдаемым ошибочным состоянием

- Covers: `UC-OP-ER03`
- Primary refs: `REQ-09`, `REQ-14`, `EC-06`, `EC-14`, `CHK-07`, `CHK-16`
- Preconditions:
  1. Оператор имеет возможность сохранить заведомо битый Liquid template.
  2. Для key доступен preview или иной verify path.
- Steps:
  1. Открыть editor выбранного key.
  2. Сохранить заведомо битый Liquid template.
  3. Открыть preview или agreed verify path.
  4. Зафиксировать наблюдаемое поведение.
- Expected result:
  1. Broken template приводит к наблюдаемой error semantics.
  2. Система не делает тихий fallback на baseline page, который скрывает ошибку.
  3. Результат совместим с fail-fast policy и runtime-error verify contract.
- Automation candidate: `mixed`

## Traceability Matrix

| UC ID | Покрываемые REQ | Покрываемые EC | Check IDs | Примечание |
|---|---|---|---|---|
| `UC-H01` | `REQ-04` | `EC-01`, `EC-04` | `CHK-12`, `CHK-01` | Независимый opt-in для `show` |
| `UC-H02` | `REQ-02` | `EC-01`, `EC-02` | `CHK-13`, `CHK-04` | Независимый opt-in для `failed` |
| `UC-H03` | `REQ-03` | `EC-01`, `EC-03` | `CHK-12`, `CHK-13`, `CHK-09` | Независимый opt-in для `paid` на order-state и callback/result success surfaces |
| `UC-H04` | `REQ-07` | `EC-01`, `EC-04` | `CHK-12`, `CHK-05` | Независимый opt-in для `created` |
| `UC-H05` | `REQ-08` | `EC-01`, `EC-04` | `CHK-12`, `CHK-06` | Независимый opt-in для `canceled` |
| `UC-E01` | `REQ-06` | `EC-05` | `CHK-03` | Страница без opt-in сохраняет baseline behavior |
| `UC-E02` | `REQ-10`, `REQ-11` | `EC-07`, `EC-10` | `CHK-11` | Проверка Liquid context для order-aware и order-less payment-result flows |
| `UC-ER01` | `REQ-09` | `EC-06` | `CHK-07`, `CHK-10` | `template body absent` |
| `UC-ER02` | `REQ-09` | `EC-06` | `CHK-07`, `CHK-10` | `runtime render error` |
| `UC-OP-01` | `REQ-12` | `EC-12` | `CHK-14` | Fixed catalog из пяти keys |
| `UC-OP-02` | `REQ-13` | `EC-13` | `CHK-15` | Independent enable per key |
| `UC-OP-03` | `REQ-13` | `EC-13` | `CHK-15` | Independent disable per key |
| `UC-OP-04` | `REQ-14`, `REQ-15` | `EC-14`, `EC-15` | `CHK-16`, `CHK-17` | Editor shell + context reference |
| `UC-OP-05` | `REQ-14` | `EC-14` | `CHK-16`, `CHK-19` | Save/edit flow для template body; persistence/reload автоматизированы через `CHK-19`, visual UX остаётся в `CHK-16` |
| `UC-OP-06` | `REQ-14` | `EC-14` | `CHK-16`, `CHK-19` | Preview deep-link в editor tab |
| `UC-OP-07` | `REQ-15` | `EC-15` | `CHK-17` | Variables/helpers и availability labels |
| `UC-OP-E01` | `REQ-12` | `EC-12`, `EC-16` | `CHK-14`, `CHK-18` | No ad-hoc key creation/deletion; fixed catalog ownership |
| `UC-OP-E02` | `REQ-13` | `EC-13` | `CHK-15` | Per-surface verdict for `paid` |
| `UC-OP-E03` | `REQ-14`, `REQ-15` | `EC-14`, `EC-15` | `CHK-16`, `CHK-17`, `CHK-19` | Explicit order-aware preview context |
| `UC-OP-E04` | `REQ-14` | `EC-14` | `CHK-16`, `CHK-19` | Explicit preview surface selection for `paid` |
| `UC-OP-E05` | `REQ-12` | `EC-12` | `CHK-14` | No empty catalog state |
| `UC-OP-E06` | `REQ-14` | `EC-14` | `CHK-16`, `CHK-19` | Payment-result preview uses provider-neutral preview-only route, not provider callback route |
| `UC-OP-ER01` | `REQ-09`, `REQ-13` | `EC-06`, `EC-13` | `CHK-07`, `CHK-15` | Missing template body visible in UI and consistent with runtime policy |
| `UC-OP-ER02` | `REQ-14` | `EC-14` | `CHK-19`, `CHK-16` | Dirty editor / preview contract |
| `UC-OP-ER03` | `REQ-09` | `EC-06` | `CHK-07` | Broken template verifies runtime fail-fast semantics on opted-in customer-facing surface |

## Test Ownership

### Автоматизированные тесты

- `UC-H01`, `UC-H04`, `UC-H05` — request/integration specs для opted-in order-state pages (`CHK-12`).
- `UC-H02` — request/integration specs для payment result surfaces (`CHK-13`).
- `UC-H03` — request/integration specs для `paid` на обеих concrete surfaces: order-state `vendor/orders/paid` (`CHK-12`) и callback/result success surface `vendor/payment/show(state=success)` (`CHK-13`).
- `UC-E01` — automated coverage для baseline behavior без per-page opt-in (`CHK-03`).
- `UC-E02` — automated tests на Liquid context contract (`CHK-11`).
- `UC-ER01`, `UC-ER02` — automated tests для failure semantics + ADR traceability review (`CHK-07`, `CHK-10`).
- `UC-OP-01`, `UC-OP-E05` — automated operator specs на fixed catalog (`CHK-14`).
- `UC-OP-E01` — automated operator specs на отсутствие ad-hoc create/delete flow для system-owned public page keys (`CHK-18`).
- `UC-OP-02`, `UC-OP-03`, `UC-OP-E02`, `UC-OP-ER01` — automated operator specs на toggle/state semantics и honest effective verdict (`CHK-15`).
- `UC-OP-04`, `UC-OP-07`, `UC-OP-E03` — automated operator specs на editor shell и context reference (`CHK-17` с дополнительным UI/assertion coverage по editor state).
- `UC-OP-05`, `UC-OP-06`, `UC-OP-E03`, `UC-OP-E04`, `UC-OP-E06`, `UC-OP-ER02` — automated specs на template body save/reload и deterministic preview/editor routing: deep-link в preview tab, direct public iframe URL по contract из `spec.md`, provider-neutral payment-result preview route, preview по последней сохранённой версии, explicit surface selection и explicit context selection без hidden auto-pick (`CHK-19`).

### Ручное тестирование

- `UC-H01` — визуальная проверка публичной страницы заказа с включённым Liquid (`CHK-01`).
- `UC-H02` — клиентский failure flow, приводящий к render surface для `failed`, с включённым Liquid; status-only order-page redirect фиксируется как boundary, не как `failed` render (`CHK-04`).
- `UC-H03` — success flows с включённым Liquid для `paid`, покрывающие и `vendor/orders/paid`, и `vendor/payment/show(state=success)` (`CHK-09`).
- `UC-H04` — post-create flow с включённым Liquid для `created` (`CHK-05`).
- `UC-H05` — canceled-state flow с включённым Liquid для `canceled` (`CHK-06`).
- `UC-OP-04`, `UC-OP-05`, `UC-OP-06`, `UC-OP-E03`, `UC-OP-E04`, `UC-OP-E06`, `UC-OP-ER02` — ручная визуальная проверка editor/preview UX поверх deterministic contract coverage из `CHK-19` (`CHK-16`).
- `UC-OP-ER03` — ручная проверка operator-assisted verify path допустима как дополнительный end-to-end сценарий, но canonical runtime failure semantics остаётся в `CHK-07`, а не в preview UX checks.

## Notes

- `EC-09` и `EC-11` остаются общими regression / closure contracts feature и подтверждаются `CHK-08`, а не отдельным use case.
- `UC-H01`-`UC-H05` являются основным acceptance набором для vendor-visible behavior.
- `UC-H03` intentionally не схлопывает `paid` в одну concrete page: accepted solution для этого key включает и order-state success surface, и callback/result success surface.
- `UC-H02` intentionally не расширяет `failed` на `/orders/:id?status=failed`: если flow рендерит `vendor/orders/show`, customization принадлежит `show`.
- `UC-E02` нужен отдельно, потому что без явной проверки context contract FT-4564 остаётся функционально неполной даже при корректном визуальном рендере.
- `UC-OP-*` intentionally не заменяют detailed screen design из `operator-ui/README.md`; они фиксируют testable operator journeys и candidate decomposition для future specs/checklists.
- `TC-OP-*` остаются derived test cases внутри companion-doc и не создают новые canonical `CHK-*`; preview/editor coverage в `feature.md` разделена между manual visual `CHK-16`, context-reference `CHK-17` и deterministic routing `CHK-19`, а runtime fail-fast semantics для битого template остаётся owner-ом `CHK-07`.

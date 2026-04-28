---
title: "FT-4564: Per-page Liquid override и кастомные шаблоны для страниц заказа"
doc_kind: feature
doc_function: canonical
purpose: "Canonical feature scope для Liquid-кастомизации клиентских страниц заказа и результатов оплаты (`failed`, `paid`, `show`, `created`, `canceled`) и operator-facing authoring surface для их настройки. Архитектурные решения о resolution unit, fallback policy и предварительном refactor выносятся в отдельный ADR."
derived_from:
  - ../../../../AGENTS.md
  - ../../../../memory-bank/flows/feature-flow.md
  - ../../engineering/testing-policy.md
  - ../../engineering/adr/0003-per-action-engine-override.md
  - ../../engineering/adr/0005-ft-4564-order-page-resolution-and-fallback-policy.md
  - ../../engineering/adr/0006-missing-custom-liquid-template-falls-back-to-existing-baseline-page.md
  - GitHub Issue #4564
related_docs:
  - ./spec.md
  - ./operator-ui/README.md
  - ./use-cases/README.md
  - ../../engineering/adr/0005-ft-4564-order-page-resolution-and-fallback-policy.md
  - ../../engineering/adr/0006-missing-custom-liquid-template-falls-back-to-existing-baseline-page.md
status: active
delivery_status: planned
audience: humans_and_agents
must_not_define:
  - implementation_sequence
---

# FT-4564: Per-page Liquid override и кастомные шаблоны для страниц заказа

## What

### Problem

Заказчик (Marcelo Miracles, Telegram: @XY000) хочет возможность делать кастомные страницы через шаблоны Liquid для следующих клиентских страниц, сформулированных в бизнес-терминах:

1. **Страница неудачной оплаты** — куда пользователь попадает при выходе из любого виджета оплаты (CloudPayments, YooMoney, Tinkoff и др.). Требуется кастомизация текста, дизайна и CTA (например, «Попробовать снова», «Связаться с поддержкой»).
2. **Страница успешной оплаты** — подтверждение оплаты. Требуется кастомизация: благодарность, детали заказа, кнопки «Вернуться в магазин», аналитические события (реализуются заказчиком самостоятельно через вставку JS-кода в кастомный Liquid-шаблон; со стороны платформы дополнительных механизмов не требуется).
3. **Страница заказа клиента** — публичная страница по ссылке, которую видит клиент (статус заказа, трек-номер, состав заказа). Требуется полная кастомизация: вывод статуса, трек-номеров службы доставки, история статусов, кастомные блоки.

Для полного покрытия customer-facing order flow FT-4564 включает не только исходные три бизнес-страницы, но и state-specific страницы, без которых опыт кастомизации будет частичным.

| Бизнес-формулировка заказчика | Страницы / page keys в scope | Пояснение |
| --- | --- | --- |
| Неудачная оплата | `failed`, `canceled` | `failed` покрывает renderable outcome-level failure result surface; `canceled` покрывает canceled-state страницу заказа. Status-only redirect на публичную страницу заказа остаётся owned by `show` / `canceled` по accepted spec и не становится отдельной `failed` page. |
| Успешная оплата | `paid` | Прямое соответствие бизнес-сценарию успешной оплаты. |
| Страница заказа клиента | `show`, `created` | Помимо обычной публичной страницы заказа нужен created-state, иначе customer-facing order flow будет кастомизироваться неполностью. |

Таким образом, `created` и `canceled` фиксируются в feature не как новый коммерческий scope, а как необходимые state-specific страницы. Exact internal resolution unit и mapping к concrete render surfaces сознательно не фиксируются в этом документе и вынесены в accepted ADR-0005 и sibling `spec.md`.

Для `paid` и `failed` эти labels в `feature.md` остаются бизнесовыми именами capability, а не обещанием единственного route/template. Concrete render-surface ownership для них задаётся accepted ADR/spec. В частности, flow вида `/payments/failure -> /orders/:id?status=failed` считается transition-only bridge: если финальный render происходит на `vendor/orders/show`, он кастомизируется через `show`, а не через `failed`.

Кроме runtime-поведения FT-4564 требует operator-facing control plane: оператор должен понимать, какие страницы доступны для Liquid-кастомизации, где включить или выключить page-specific Liquid, где редактировать шаблон, как получить preview и где посмотреть доступный Liquid context для выбранной страницы. Детальный screen design этого operator UI выносится в `operator-ui/README.md`, а этот документ фиксирует только problem-space contract и verify.

Все пять страниц в scope должны:
- Поддерживать vendor-configurable Liquid customization без полного переключения глобального engine
- Быть доступны в operator UI через фиксированный catalog допустимых page keys, а не через ad-hoc создание произвольных FT-4564 страниц
- Использовать dedicated public page template storage (`VendorPageTemplate` / `vendor_page_templates`), а не переиспользовать unrelated content mechanisms
- Предоставлять Liquid context: `vendor`, `order` при наличии order context и стандартные helpers
- При отсутствии page-specific Liquid настройки сохранять текущее глобальное поведение
- Иметь явно принятую и задокументированную политику поведения для случая `Liquid template body absent` и для `runtime render error` до старта реализации

Дополнительно: FT-4564 требует безопасного постраничного включения Liquid без полного переключения глобального engine. Problem-space contract требует отделить opt-in state от наличия Liquid template body; concrete persistence shape принадлежит `spec.md`, а не этому документу.

### Outcome

| Metric ID | Metric | Baseline | Target | Measurement method |
| --- | --- | --- | --- | --- |
| `MET-01` | Вендоры с хотя бы одной Liquid-страницей (`show`, `paid`, `failed`, `created` или `canceled`) | 0 | ≥ 1 (пилот) | Metric readout procedure через `CHK-20` / `EVID-10`; target оценивается после пилотного включения и не подменяет pre-release closure gates |

### Scope

#### Бизнес-требования заказчика (из issue #4564)

- `REQ-02` Liquid-кастомизация для клиентской страницы неудачной оплаты (`failed`) на renderable failure result surfaces, которые accepted solution docs относят к page key `failed`. Status-only order-page redirect после failed payment не расширяет ownership `failed` и остаётся boundary `NS-08`. *Источник: issue #4564 п.1.*
- `REQ-03` Liquid-кастомизация для клиентской страницы успешной оплаты (`paid`). *Источник: issue #4564 п.2. Аналитические события реализуются заказчиком через JS-вставку в кастомный шаблон; со стороны платформы отдельных механизмов не требуется.*
- `REQ-04` Liquid-кастомизация для публичной страницы заказа клиента (`show`). *Источник: issue #4564 п.3.*

#### Производные функциональные требования, необходимые для полного покрытия customer scope в Merchantly

- `REQ-05` Resolution priority для in-scope public page surfaces: page-specific Liquid opt-in управляет FT-4564 surface; если Liquid для page key не включён, используется текущее глобальное поведение. *Нужно ввести dedicated public page template mechanism без расширения feature scope на unrelated content customization.*
- `REQ-06` Если для страницы Liquid не включён, используется текущее глобальное поведение. *Это отдельный контракт, не равный policy для `template body absent` / `runtime render error`.*
- `REQ-07` Liquid-кастомизация для страницы создания заказа (`created`) как производной state-specific страницы, необходимой для полного покрытия customer-facing order flow. *Это не новый бизнес-scope, а stateful декомпозиция исходного требования.*
- `REQ-08` Liquid-кастомизация для публичной страницы заказа в состоянии cancel (`canceled`) как производной state-specific страницы, необходимой для полного покрытия customer-facing order flow после неуспешной оплаты / отмены. *Это не новый бизнес-scope, а stateful декомпозиция исходного требования.*
- `REQ-09` Для страницы с включённой page-specific Liquid настройкой поведение при отсутствии Liquid template body и при runtime render error должно быть детерминированным, аргументированным, зафиксированным в ADR и покрытым проверками. *Нельзя оставлять эту семантику неявной.*
- `REQ-10` Liquid context для opted-in страниц предоставляет `vendor`, page outcome/state metadata и стандартные helpers; при наличии order context шаблон также получает `order`. *Нужно для реальной кастомизации результата оплаты и публичной страницы заказа без скрытых требований к единой runtime-surface форме.*
- `REQ-11` Payment result surfaces (`paid`, `failed`) должны оставаться корректно рендеримыми и не требовать `order` в тех customer-facing flows, где order context отсутствует; при наличии order context он передаётся в шаблон. *Нужно, чтобы callback/result surfaces не ломались из-за отсутствия universally available `order`.*
- `REQ-12` В операторской должен существовать discoverable authoring surface для FT-4564 с fixed catalog допустимых Liquid page keys (`show`, `paid`, `failed`, `created`, `canceled`); Liquid keys не создаются вручную как произвольные записи. Concrete entrypoint и IA фиксируются в `spec.md` / `operator-ui/README.md`, а не в problem-space contract. *Нужно, чтобы authoring surface оставался bounded и работал именно с public page templates.*
- `REQ-13` Для каждого допустимого page key оператор видит current mode (`Baseline` / `Liquid enabled`) и honest effective status (`Liquid`, `Baseline`, `template body missing`); если page key покрывает несколько concrete render surfaces, effective status раскрывается per surface без lossy aggregate verdict. Оператор может независимо включать/выключать page-specific Liquid без изменения других keys. *Нужно, чтобы operator UI не скрывал реальные resolution semantics FT-4564.*
- `REQ-14` Оператор может перейти из каталога в dedicated editor выбранного page key, изменить Liquid template body, сохранить его и запустить preview из operator UI; preview работает по последней сохранённой версии шаблона и не preview-ит unsaved draft silently; если page key покрывает несколько concrete surfaces, preview должен поддерживать каждую из них через явный surface selection. *Нужен единый authoring flow вместо распределённых действий по нескольким разделам и без скрытого схлопывания multi-surface keys в один preview mode.*
- `REQ-15` Editor выбранного page key показывает page-specific context reference: доступные переменные, helpers и пометку, где `order` required, optional или unavailable. *Без этого Liquid authoring остаётся непрозрачным и повышает риск невалидных шаблонов.*

### Non-Scope

- `NS-01` Перенос checkout (`orders/new`) — см. FT-4542 / #4454.
- `NS-02` Перенос каталога, корзины, личного кабинета — см. #4501.
- `NS-03` Провайдерские страницы оплаты и payment widgets (CloudPayments, YooMoney и др.) — FT-4564 покрывает только customer-facing result/order pages после выхода из платёжного flow.
- `NS-04` Полный отказ от React (глобальный engine switch).
- `NS-05` Явный per-page selector `React` для order/result pages.
- `NS-06` Redesign соседних content-разделов операторской вне page keys FT-4564.
- `NS-07` WYSIWYG / drag-and-drop page builder для Liquid-шаблонов.
- `NS-08` Отдельная Liquid-кастомизация status-only order-page landing вроде `/orders/:id?status=failed` как page key `failed`. Такие bridge flows либо заканчиваются на renderable surface `failed`, либо остаются owned by финальным order-state key (`show` / `canceled`).
- `NS-09` Standalone customer-facing страница инициации оплаты `vendor/orders/payment` (`orders/payment`) не входит в FT-4564. Эта adjacent surface остаётся текущим baseline payment-entry step и не считается ни order-result page, ни payment-result page в scope данной фичи.

### Constraints / Assumptions

- `ASM-01` Customer-facing order/payment flow уже сегодня stateful и не сводится к одной единственной публичной странице заказа.
- `CON-01` Нельзя сломать существующие React-страницы при частичном переходе на Liquid.
- `CON-02` FT-4564 хранит настройки public page templates в dedicated `VendorPageTemplate` / `vendor_page_templates` storage и не переиспользует unrelated content storage.
- `CON-03` FT-4564 должен оставаться совместимым с ADR-0003: accepted contract для explicit `engine:` override нельзя неявно сломать или переопределить.
- `CON-04` FT-4564 должна оставаться верифицируемой через automated regression checks для customer-facing order/result surfaces в scope; конкретный suite composition, baseline procedure и approval gates задаются downstream в `implementation-plan.md`, а не в этом документе.
- `CON-05` Operator UI должен показывать только FT-4564 public page template state; unrelated content state не должен входить в Liquid effective verdict.

### Behavioral Contracts

| Contract ID | Description | Notes |
| --- | --- | --- |
| `CTR-01` | Resolution priority | Для in-scope FT-4564 surfaces: `Liquid enabled + body` -> Liquid; `Liquid enabled + body absent` -> existing baseline page implementation; `Liquid disabled/not configured` -> текущее глобальное поведение |
| `CTR-02` | Liquid context для шаблонов | Feature contract требует `vendor`, page outcome/state metadata и стандартные helpers; `order` обязателен для order-aware surfaces и optional для payment result flows, где order context может отсутствовать. Дополнительные переменные допустимы, но не входят в обязательный контракт FT-4564. |
| `CTR-03` | Поведение без page-specific Liquid настройки | Если для страницы Liquid не включён, используется текущее глобальное поведение |
| `CTR-04` | `Template body absent` / `runtime render error` policy | Поведение определяется accepted ADR-0005 / ADR-0006; feature.md фиксирует только требование детерминированности и verify coverage |
| `CTR-05` | Operator authoring discoverability | Operator UI даёт discoverable authoring surface и показывает fixed catalog FT-4564 page keys; concrete entrypoint / IA owned by `spec.md` и `operator-ui/README.md` |
| `CTR-06` | Operator status semantics | Catalog и editor показывают configured state и FT-4564 effective render state отдельно; multi-surface page keys раскрывают effective verdict per surface, включая `template body missing` |
| `CTR-07` | Operator editor contract | Preview и context reference доступны из editor выбранного page key; preview использует последнюю сохранённую версию template body, явный context selection и не auto-pick-ает случайный live order; multi-surface page keys требуют явный surface selection и поддерживают preview каждой принадлежащей surface; labels availability для `order` и metadata согласованы с `CTR-02` |

### Failure Modes

- `FM-01` Для страницы включён Liquid, но Liquid template body отсутствует → система должна применить accepted missing-template policy, зафиксированную в ADR-0005 / ADR-0006, для соответствующей страницы.
- `FM-02` Public page templates моделируются как ad-hoc records вместо fixed catalog → оператор получает ложный scope и может создать неподдерживаемый page key.
- `FM-03` Runtime render error на Liquid-странице → система должна применить accepted runtime-error policy из ADR-0005 с явной наблюдаемостью ошибки.
- `FM-04` Operator UI смешивает configured mode и effective render verdict → включённый, но пустой template body выглядит как успешный Liquid render.

### ADR Dependencies

| ADR | Current `decision_status` | Used for | Execution rule |
| --- | --- | --- | --- |
| [ADR-0003: Per-Action Engine Override](../../engineering/adr/0003-per-action-engine-override.md) | `accepted` | Baseline contract для explicit `engine:` override | FT-4564 должен оставаться совместимым с accepted baseline ADR-0003 |
| [ADR-0005: FT-4564 Order Page Resolution and Fallback Policy](../../engineering/adr/0005-ft-4564-order-page-resolution-and-fallback-policy.md) | `accepted` | Resolution unit, runtime error policy и bounded refactor | Feature package должен быть синхронизирован с accepted baseline ADR-0005 |
| [ADR-0006: Missing Custom Liquid Template Falls Back to Existing Baseline Page](../../engineering/adr/0006-missing-custom-liquid-template-falls-back-to-existing-baseline-page.md) | `accepted` | Cross-feature fallback rule для `Liquid selected + template body absent` | FT-4564 и будущие custom Liquid features должны сохранять fallback на existing baseline page implementation |

## Verify

### Exit Criteria

- `EC-01` Вендор может независимо включить page-specific Liquid mode для любой из пяти страниц в scope (`show`, `paid`, `failed`, `created`, `canceled`) без смены глобального engine.
- `EC-02` При попадании на customer-facing render surface, принадлежащую page key `failed`, отображается Liquid-кастомизация для `failed`, если она включена; status-only order-page redirects не считаются `failed` surfaces.
- `EC-03` При попадании на любую customer-facing render surface, принадлежащую page key `paid`, отображается Liquid-кастомизация для `paid`, если она включена; для current accepted solution это минимум `vendor/orders/paid` и `vendor/payment/show(state=success)`.
- `EC-04` Публичные order states `show`, `created`, `canceled` поддерживают независимую Liquid-кастомизацию.
- `EC-05` Если для страницы Liquid не включён, она продолжает рендериться через текущее глобальное поведение.
- `EC-06` Поведение при отсутствии Liquid template body и при runtime render error соответствует принятому ADR-0005 / ADR-0006 policy.
- `EC-07` Liquid context для opted-in страниц предоставляет `vendor`, `order` при наличии order context, page outcome/state metadata и стандартные helpers.
- `EC-08` FT-4564 использует dedicated `VendorPageTemplate` storage и не создаёт ad-hoc public page keys вне fixed catalog.
- `EC-09` Существующие React-based order/result surfaces остаются неизменными при частичном переходе других страниц на Liquid.
- `EC-10` Payment result surfaces (`paid`, `failed`) остаются корректно рендеримыми и contract-valid в flows без `order`; в этом случае шаблон получает `vendor`, outcome/state metadata и стандартные helpers, а `order` передаётся только при наличии context.
- `EC-11` Required automated regression suites для change surface FT-4564 зелёные локально и в CI.
- `EC-12` В operator entrypoint для FT-4564 оператор видит fixed catalog допустимых page keys (`show`, `paid`, `failed`, `created`, `canceled`).
- `EC-13` Оператор может независимо включать/выключать page-specific Liquid по каждому page key и видит current mode и honest FT-4564 effective status; для multi-surface page keys effective verdict раскрывается per surface, включая `template body missing`.
- `EC-14` Оператор может открыть editor выбранного page key, сохранить Liquid template body и перейти в preview из operator UI; preview использует последнюю сохранённую версию шаблона и не подхватывает unsaved draft silently; для multi-surface keys preview поддерживает каждую принадлежащую surface через явный surface selection; payment-result preview строится через provider-neutral preview-only route, а не через callback конкретной платёжки.
- `EC-15` Editor выбранного page key показывает context reference с переменными, helpers и page-specific availability notes для `order`.
- `EC-16` Каталог FT-4564 остаётся system-owned: оператор не может вручную создать или удалить произвольный public page key.

### Traceability matrix

| Requirement ID | Design refs | Acceptance refs | Checks | Evidence IDs |
| --- | --- | --- | --- | --- |
| `REQ-02` | `CTR-01`, `NS-08` | `EC-01`, `EC-02`, `SC-02` | `CHK-13`, `CHK-04` | `EVID-02`, `EVID-03` |
| `REQ-03` | `CTR-01` | `EC-01`, `EC-03`, `SC-03` | `CHK-12`, `CHK-13`, `CHK-09` | `EVID-02`, `EVID-06` |
| `REQ-04` | `CTR-01`, `CTR-02` | `EC-01`, `EC-04`, `SC-01` | `CHK-12`, `CHK-01` | `EVID-02`, `EVID-01` |
| `REQ-05` | `CON-02`, `CTR-01`, `FM-02` | `EC-08`, `SC-05` | `CHK-02` | `EVID-02` |
| `REQ-06` | `CTR-03`, `CON-01`, `CON-04` | `EC-05`, `EC-09`, `EC-11`, `SC-04` | `CHK-03`, `CHK-08` | `EVID-02` |
| `REQ-07` | `CTR-01` | `EC-01`, `EC-04`, `SC-06` | `CHK-12`, `CHK-05` | `EVID-02`, `EVID-04` |
| `REQ-08` | `CTR-01` | `EC-01`, `EC-04`, `SC-07` | `CHK-12`, `CHK-06` | `EVID-02`, `EVID-05` |
| `REQ-09` | `CTR-04`, `FM-01`, `FM-03` | `EC-06`, `SC-08`, `NEG-01`, `NEG-02` | `CHK-07`, `CHK-10` | `EVID-02`, `EVID-07` |
| `REQ-10` | `CTR-02` | `EC-07`, `SC-10` | `CHK-11` | `EVID-02` |
| `REQ-11` | `CTR-02`, `CON-01` | `EC-10`, `SC-09` | `CHK-11`, `CHK-13` | `EVID-02` |
| `REQ-12` | `CTR-05` | `EC-12`, `EC-16`, `SC-11`, `SC-15` | `CHK-14`, `CHK-18` | `EVID-02` |
| `REQ-13` | `CTR-01`, `CTR-06`, `CON-05`, `FM-02`, `FM-04` | `EC-13`, `SC-12` | `CHK-15` | `EVID-02` |
| `REQ-14` | `CTR-06`, `CTR-07` | `EC-14`, `SC-13` | `CHK-16`, `CHK-19` | `EVID-08`, `EVID-02` |
| `REQ-15` | `CTR-02`, `CTR-07` | `EC-15`, `SC-14` | `CHK-17` | `EVID-02` |

### Acceptance Scenarios

- `SC-01` **Liquid для публичной страницы заказа**: Вендор включает Liquid для `show` → пользователь открывает публичную страницу заказа в regular-state → видит Liquid-шаблон с данными заказа.
- `SC-02` **Liquid для неудачной оплаты**: Вендор включает Liquid для `failed` → пользователь попадает на customer-facing render surface, которую accepted solution docs относят к `failed` → видит Liquid-кастомизацию для `failed`; если failure flow заканчивается status-only redirect на `vendor/orders/show`, этот render остаётся owned by `show`, а не `failed`.
- `SC-03` **Liquid для успешной оплаты**: Вендор включает Liquid для `paid` → пользователь попадает на customer-facing render surface, которую accepted solution docs относят к `paid` → видит Liquid-кастомизацию для `paid`; canonical verify обязан покрыть и order-state success surface `vendor/orders/paid`, и callback/result success surface `vendor/payment/show(state=success)`.
- `SC-04` **Нет page-specific Liquid настройки**: Вендор не включал Liquid для страницы → эта страница использует текущее глобальное поведение.
- `SC-05` **Opted-in Liquid page template выбран для surface**: Для in-scope surface включён page-specific Liquid с body → рендерится Liquid template; при выключенном page-specific Liquid используется текущее глобальное поведение.
- `SC-06` **Liquid для состояния created**: Вендор включает Liquid для `created` → пользователь видит Liquid-шаблон `created`. *Производный state-specific сценарий, необходимый для полного покрытия customer-facing order flow.*
- `SC-07` **Liquid для состояния canceled**: Вендор включает Liquid для `canceled` → пользователь видит Liquid-шаблон `canceled`. *Производный state-specific сценарий, необходимый для полного покрытия customer-facing order flow после неуспешной оплаты / отмены.*
- `SC-08` **Deterministic failure semantics**: Вендор включает Liquid для страницы → при `template body absent` и при runtime render error система применяет соответствующую accepted policy из ADR-0005 / ADR-0006; поведение остаётся детерминированным и наблюдаемым.
- `SC-09` **Liquid для payment result surface без order context**: Вендор включает Liquid для `paid` или `failed` → пользователь попадает на customer-facing payment result surface по callback/return flow, где `order` недоступен → Liquid-шаблон рендерится с `vendor`, outcome/state metadata и helpers, отсутствие `order` не ломает страницу.
- `SC-10` **Liquid context contract для order-aware surfaces**: Вендор включает Liquid для order-state страницы или payment result surface с доступным order context → шаблон получает `vendor`, `order`, page outcome/state metadata и стандартные helpers.
- `SC-11` **Operator видит public page templates catalog**: Оператор открывает FT-4564 authoring entrypoint → видит список `show`, `paid`, `failed`, `created`, `canceled`, и эти keys не создаются вручную как произвольные записи.
- `SC-12` **Operator toggle и honest effective status**: Оператор включает или выключает Liquid для выбранного page key → статус меняется только у этого key; для `paid` или другого multi-surface key UI всегда раскрывает verdict per surface и не схлопывает его в aggregate badge даже когда verdict совпадают.
- `SC-13` **Operator editor и preview**: Оператор открывает editor выбранного page key → изменяет Liquid template body → сохраняет его → запускает preview из operator UI; preview строится по последней сохранённой версии template body и не preview-ит unsaved draft silently; если page key multi-surface, UI требует явный выбор preview surface; для `paid` доступны минимум `vendor/orders/paid` и `vendor/payment/show(state=success)`; если выбранной surface нужен context, UI требует явный выбор context и не auto-pick-ает случайный live order; для `payment/show` preview использует `/payments/page_template_preview/success|failure`.
- `SC-14` **Operator видит context reference**: Оператор открывает editor `show`, `paid` или `failed` → видит список доступных variables/helpers и явную пометку, где `order` required, optional или unavailable.
- `SC-15` **Ad-hoc public page keys запрещены**: Оператор не может вручную создать или удалить public page key; catalog остаётся fixed и system-owned.

### Negative / Edge Cases

- `NEG-01` **Liquid template body отсутствует**: Вендор включает Liquid для страницы, но Liquid template body отсутствует → система применяет accepted missing-template policy из ADR-0005 / ADR-0006 для соответствующей страницы.
- `NEG-02` **Runtime render error на Liquid-странице**: Вендор включает Liquid для страницы, во время Liquid-render возникает ошибка → система применяет accepted runtime-error policy из ADR-0005 и не оставляет результат неявным.

### Checks

| Check ID | Covers | How to check | Expected result | Evidence path |
| --- | --- | --- | --- | --- |
| `CHK-01` | `EC-04`, `SC-01` | Ручное тестирование: открытие `show` с включённым Liquid | Публичная страница заказа в regular-state рендерится через Liquid корректно | `artifacts/ft-4564/verify/chk-01/` |
| `CHK-02` | `EC-08`, `SC-05` | Автоматизированные тесты на resolution order | Page-specific Liquid opt-in выбирает Liquid template для matched surface; при выключенном Liquid используется текущее глобальное поведение | Local test output + CI run |
| `CHK-03` | `EC-05`, `SC-04` | Автоматизированные тесты на страницу без page-specific Liquid | Используется текущее глобальное поведение | Local test output + CI run |
| `CHK-04` | `EC-02`, `SC-02` | Ручное тестирование: клиентский failure flow, приводящий к render surface для `failed`, с включённым Liquid; status-only order-page redirect проверяется как boundary, а не как `failed` render | Отображается Liquid-кастомизация для `failed`; status-only order-page landing не ошибочно засчитывается как `failed` surface | `artifacts/ft-4564/verify/chk-04/` |
| `CHK-05` | `EC-04`, `SC-06` | Ручное тестирование: создание заказа с включённым Liquid для `created` | Отображается Liquid-шаблон `created` | `artifacts/ft-4564/verify/chk-05/` |
| `CHK-06` | `EC-04`, `SC-07` | Ручное тестирование: открытие canceled-state страницы с включённым Liquid для `canceled` | Отображается Liquid-шаблон `canceled` | `artifacts/ft-4564/verify/chk-06/` |
| `CHK-07` | `EC-06`, `SC-08`, `NEG-01`, `NEG-02` | Автоматизированные тесты на `template body absent` и runtime render error | Поведение соответствует accepted ADR-0005 / ADR-0006 policy | Local test output + CI run |
| `CHK-08` | `EC-09`, `EC-11` | Локальный прогон required suites + CI run после изменений FT-4564 | Required suites зелёные локально и в CI, React-страницы не сломаны | Local test output + CI run |
| `CHK-09` | `EC-03`, `SC-03` | Ручное тестирование: success flows с включённым Liquid для `paid`, приводящие минимум к `vendor/orders/paid` и `vendor/payment/show(state=success)` | Для обеих success surfaces отображается Liquid-кастомизация `paid` | `artifacts/ft-4564/verify/chk-09/` |
| `CHK-10` | `EC-06`, `SC-08`, `NEG-01`, `NEG-02` | Проверка ссылок и traceability между `feature.md` и accepted ADR-0005 / ADR-0006 | Canonical owner policy остаётся в accepted ADR, а `feature.md` содержит только problem-space contract и verify coverage | `memory-bank/features/FT-4564/feature.md` + accepted ADR links |
| `CHK-11` | `EC-07`, `EC-10`, `SC-09`, `SC-10` | Автоматизированные тесты на Liquid context opted-in страниц, включая payment result flows без `order` | Order-aware surfaces получают `vendor`, `order`, metadata и helpers; payment result surfaces остаются рендеримыми с `vendor`, outcome/state metadata и helpers даже без `order` | Local test output + CI run |
| `CHK-12` | `EC-01`, `EC-03`, `EC-04`, `SC-01`, `SC-03`, `SC-06`, `SC-07` | Автоматизированные request/integration specs для opted-in order-state страниц (`show`, `paid`, `created`, `canceled`) | Каждая order-state страница независимо рендерится через Liquid при явном opt-in; для `paid` отдельно подтверждён order-state surface `vendor/orders/paid` | Local test output + CI run |
| `CHK-13` | `EC-01`, `EC-02`, `EC-03`, `SC-02`, `SC-03` | Автоматизированные request/integration specs для callback/result payment surfaces `vendor/payment/show(state=failure|success)` | Каждая callback/result payment surface рендерится через Liquid при явном opt-in; для `paid` отдельно подтверждён success surface `vendor/payment/show(state=success)` | Local test output + CI run |
| `CHK-14` | `EC-12`, `SC-11` | Автоматизированные specs на operator entrypoint IA и fixed catalog FT-4564 page keys | Operator UI показывает bounded catalog `show`, `paid`, `failed`, `created`, `canceled` без ad-hoc создания этих keys | Local test output + CI run |
| `CHK-15` | `EC-13`, `SC-12`, `FM-04` | Автоматизированные specs на operator toggle/state semantics | Toggle работает per key; UI показывает `template body missing`; multi-surface keys вроде `paid` раскрывают effective verdict per surface | Local test output + CI run |
| `CHK-16` | `EC-14`, `SC-13` | Ручное визуальное тестирование editor и preview из operator UI | Оператор может сохранить шаблон выбранного key и открыть preview из editor без поиска по другим разделам; preview iframe показывает direct customer-facing public render surface без вложенной операторской; для `paid` визуально доступны минимум `vendor/orders/paid` и `vendor/payment/show(state=success)`, а `payment/show` preview открывается через `/payments/page_template_preview/success|failure` | `artifacts/ft-4564/verify/chk-16/` |
| `CHK-17` | `EC-15`, `SC-14` | Автоматизированные specs на context reference в operator editor | Editor показывает variables/helpers и page-specific availability notes, включая `order required/optional/unavailable` | Local test output + CI run |
| `CHK-18` | `EC-16`, `SC-15` | Автоматизированные specs на fixed catalog ownership | Оператор не может создать или удалить ad-hoc public page key; fixed catalog остаётся доступен даже без сохранённых templates | Local test output + CI run |
| `CHK-19` | `EC-14`, `SC-13` | Автоматизированные specs на editor persistence и deterministic preview routing contract | Liquid template body сохраняется и повторно загружается для выбранного key; preview entrypoint deep-links into editor preview tab, использует последнюю сохранённую версию шаблона и не preview-ит unsaved draft silently; multi-surface keys требуют явный surface selection; order-aware preview требует explicit context selection без hidden auto-pick; payment-result preview path не использует `/payments/<provider>/success|failure` | Local test output + CI run |
| `CHK-20` | `MET-01` | Metric readiness check: зафиксировать query/report procedure, считающую vendors с хотя бы одной enabled Liquid page по selected semantic-key persistence surface | Метрика `MET-01` измерима; post-release pilot target `≥ 1` будет проверяться тем же readout и не является pre-release regression gate | `artifacts/ft-4564/verify/met-01/` |

### Test matrix

| Check ID | Evidence IDs | Evidence path |
| --- | --- | --- |
| `CHK-01` | `EVID-01` | `artifacts/ft-4564/verify/chk-01/` |
| `CHK-02` | `EVID-02` | Local test output + CI run |
| `CHK-03` | `EVID-02` | Local test output + CI run |
| `CHK-04` | `EVID-03` | `artifacts/ft-4564/verify/chk-04/` |
| `CHK-05` | `EVID-04` | `artifacts/ft-4564/verify/chk-05/` |
| `CHK-06` | `EVID-05` | `artifacts/ft-4564/verify/chk-06/` |
| `CHK-07` | `EVID-02` | Local test output + CI run |
| `CHK-08` | `EVID-02` | Local test output + CI run |
| `CHK-09` | `EVID-06` | `artifacts/ft-4564/verify/chk-09/` |
| `CHK-10` | `EVID-07` | `memory-bank/features/FT-4564/feature.md` + accepted ADR links |
| `CHK-11` | `EVID-02` | Local test output + CI run |
| `CHK-12` | `EVID-02` | Local test output + CI run |
| `CHK-13` | `EVID-02` | Local test output + CI run |
| `CHK-14` | `EVID-02` | Local test output + CI run |
| `CHK-15` | `EVID-02` | Local test output + CI run |
| `CHK-16` | `EVID-08` | `artifacts/ft-4564/verify/chk-16/` |
| `CHK-17` | `EVID-02` | Local test output + CI run |
| `CHK-18` | `EVID-02` | Local test output + CI run |
| `CHK-19` | `EVID-02` | Local test output + CI run |
| `CHK-20` | `EVID-10` | `artifacts/ft-4564/verify/met-01/` |

### Evidence

- `EVID-01` Скриншоты Liquid-шаблона `show` в корректном состоянии.
- `EVID-02` Зелёный локальный прогон required suites и зелёный CI run для regression/check bundle FT-4564.
- `EVID-03` Скриншоты или видео flow неудачной оплаты с Liquid-шаблоном `failed`.
- `EVID-04` Скриншоты Liquid-шаблона `created` в корректном состоянии.
- `EVID-05` Скриншоты Liquid-шаблона `canceled` в корректном состоянии.
- `EVID-06` Скриншоты или видео success flows с Liquid-кастомизацией `paid`, покрывающие минимум `vendor/orders/paid` и `vendor/payment/show(state=success)`.
- `EVID-07` Ссылки из `feature.md` на accepted ADR-0005 / ADR-0006 и traceability для `template body absent` / `runtime render error`.
- `EVID-08` Скриншоты или видео operator authoring flow: каталог page keys, editor выбранной страницы и direct public preview из operator UI без вложенной операторской, включая surface selection для `paid`.
- `EVID-10` Metric readout procedure для `MET-01`: query/report source для vendors with enabled FT-4564 Liquid page count; post-release pilot value добавляется тем же artifact, когда доступен.

### Evidence contract

| Evidence ID | Artifact | Producer | Path contract | Reused by checks |
| --- | --- | --- | --- | --- |
| `EVID-01` | Скриншоты | verify-runner / human | `artifacts/ft-4564/verify/chk-01/` | `CHK-01` |
| `EVID-02` | Local + CI test output for required suites / regression bundle | local test runner / CI | Local test output + GitHub Actions run for FT-4564 required suites | `CHK-02`, `CHK-03`, `CHK-07`, `CHK-08`, `CHK-11`, `CHK-12`, `CHK-13`, `CHK-14`, `CHK-15`, `CHK-17`, `CHK-18`, `CHK-19` |
| `EVID-03` | Скриншоты / видео | verify-runner / human | `artifacts/ft-4564/verify/chk-04/` | `CHK-04` |
| `EVID-04` | Скриншоты | verify-runner / human | `artifacts/ft-4564/verify/chk-05/` | `CHK-05` |
| `EVID-05` | Скриншоты | verify-runner / human | `artifacts/ft-4564/verify/chk-06/` | `CHK-06` |
| `EVID-06` | Скриншоты / видео | verify-runner / human | `artifacts/ft-4564/verify/chk-09/` | `CHK-09` |
| `EVID-07` | Feature-to-ADR traceability links | human | `memory-bank/features/FT-4564/feature.md` + links to accepted ADR-0005 / ADR-0006 | `CHK-10` |
| `EVID-08` | Скриншоты / видео operator catalog, editor и preview | verify-runner / human | `artifacts/ft-4564/verify/chk-16/` | `CHK-16` |
| `EVID-10` | Metric readout procedure / pilot readout | release owner / analyst | `artifacts/ft-4564/verify/met-01/` или PR/issue note с query/result | `CHK-20` |

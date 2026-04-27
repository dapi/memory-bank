---
title: "FT-4564: Solution Spec"
doc_kind: spec
doc_function: canonical
purpose: "Feature-level solution spec для FT-4564. Фиксирует selected design, stable contracts, persistence shape и rollout/backout baseline без смешения с problem statement, grounding и execution sequencing."
derived_from:
  - feature.md
  - ../../engineering/adr/0005-ft-4564-order-page-resolution-and-fallback-policy.md
  - ../../engineering/adr/0006-missing-custom-liquid-template-falls-back-to-existing-baseline-page.md
  - ../../ops/release.md
related_docs:
  - ./runtime-surfaces.md
  - ./operator-ui/README.md
status: active
audience: humans_and_agents
must_not_define:
  - ft_4564_problem_statement
  - ft_4564_acceptance_criteria
  - current_system_inventory
  - implementation_sequence
  - file_level_touchpoints
---

# FT-4564: Solution Spec

## Scope of This Spec

Этот `spec.md` опирается на accepted ADR-0005 и ADR-0006 и является canonical owner solution space для FT-4564.

Документ фиксирует:

- feature-local accepted decisions;
- semantic page-key model и surface ownership contract;
- concrete runtime, persistence и operator authoring contracts;
- rollout/backout baseline для safe activation поверх текущего React-based behavior.

`feature.md` остаётся владельцем problem statement и verify inventory. ADR-0005 и ADR-0006 остаются владельцами архитектурного baseline. Grounding по текущим runtime surfaces, semantic mapping и context evidence вынесен в [`runtime-surfaces.md`](runtime-surfaces.md). Detailed operator UX reference живёт в [`operator-ui/README.md`](operator-ui/README.md). `implementation-plan.md` остаётся downstream execution-документом и не подменяет этот `spec.md` как owner solution space.

## Selected Solution

Selected solution уже зафиксирован в accepted ADR-0005: stable semantic page keys, bounded resolver/policy layer, dedicated `VendorPageTemplate` persistence, правило `Liquid template body exists -> Liquid`, `template body absent -> existing baseline page implementation`, а также fail-fast для runtime Liquid render error.

Этот `spec.md` concretize-ит принятый baseline feature-local contracts и задаёт solution-level baseline для `implementation-plan.md` и последующей runtime-реализации.

## Accepted Decisions

| Decision ID | Decision | Why | Related refs |
| --- | --- | --- | --- |
| `SD-01` | FT-4564 public Liquid page templates являются отдельным page-template mechanism на `VendorPageTemplate`. Для in-scope surfaces explicit page-specific Liquid opt-in выбирает Liquid template; без FT-4564 opt-in используется текущее глобальное поведение. | Новая фича должна заменять React surface через dedicated `VendorPageTemplate`, а не через unrelated content storage. | `REQ-05`, `CON-02`, `CTR-01`, ADR-0005 |
| `SD-02` | Provider-hosted payment forms и widgets остаются вне solution surface FT-4564. | Кастомизация начинается только на customer-facing result surfaces после provider callback / return flow; это уже вынесено в non-scope. | `NS-03`, ADR-0005 |
| `SD-03` | Если для page key выбран Liquid и Liquid template body отсутствует, используется existing baseline page implementation для той же страницы / surface. | Это accepted cross-feature rule из ADR-0006 и explicit feature-level contract для FT-4564. | `REQ-09`, ADR-0005, ADR-0006 |
| `SD-04` | `order_created` трактуется как post-create surface, а не как regular public order state. | Current code даёт реальный `created` path в post-create flow, но не даёт убедимого evidence для regular show-state reachability. | `REQ-07`, ADR-0005 |
| `SD-05` | Runtime render error после выбора Liquid не переводится в React fallback. | Нужно сохранить debugability и не маскировать реальные баги Liquid-render. | `REQ-09`, ADR-0005 |
| `SD-06` | Каждый renderable customer-facing surface в scope FT-4564 принадлежит ровно одному vendor-facing page key; transition-only bridges page key не получают. | Без total ownership contract независимый opt-in для `show` / `created` / `canceled` / `paid` / `failed` расплывается и начинает пересекаться через controller branches. | `REQ-02`, `REQ-03`, `REQ-04`, `REQ-07`, `REQ-08`, `EC-01`, `EC-04`, ADR-0005 |
| `SD-07` | Активация Liquid-specific runtime changes для FT-4564 допустима только после подтверждённой baseline automated regression protection для текущих React-based customer-facing order/result surfaces в scope. | Это feature-level delivery safety gate для миграции без регрессий и без неявного ухудшения существующего поведения. Уровень harness (`controller` / `request` / `integration` / stronger) определяется downstream test strategy, а не этим spec. | `CON-01`, `CON-04`, `EC-09`, `EC-11` |
| `SD-08` | FT-4564 получает discoverable operator entrypoint для public page templates и не требует redesign остальных content authoring flows. | Оператору нужен прямой путь к page-template catalog без превращения фичи в общий page builder. | `REQ-12`, `EC-12`, `operator-ui/README.md` |
| `SD-09` | Operator entrypoint показывает bounded `Liquid pages` surface с фиксированным каталогом `show`, `paid`, `failed`, `created`, `canceled`. | Public page templates являются system-owned catalog; оператор не создаёт ad-hoc page keys. | `REQ-12`, `REQ-13`, `CTR-05`, `CTR-06`, `operator-ui/README.md` |
| `SD-10` | Editor выбранного page key объединяет body authoring, status toggle, preview entrypoint и context reference на одной operator surface. | Разделение этих действий по нескольким экранам повышает когнитивную нагрузку и скрывает contract `order required/optional`. | `REQ-14`, `REQ-15`, `CTR-07`, `operator-ui/README.md` |
| `SD-11` | Per-page Liquid opt-in и template body хранятся в отдельной vendor-scoped persistence surface `VendorPageTemplate` / `vendor_page_templates`, а не в global `VendorTheme#engine`. Storage использует internal semantic keys из ADR-0005: `payment_failure`, `payment_success`, `order_show`, `order_created`, `order_canceled`. Минимальный shape: `vendor_id`, `semantic_key`, `enabled`, nullable `template_body`, timestamps, unique index `[vendor_id, semantic_key]`. Vendor-facing aliases `failed`, `paid`, `show`, `created`, `canceled` существуют только в UI/docs mapping. | Нужно отделить configured opt-in от наличия body, не превращать global theme engine в per-page selector и не нарушить ADR-0005 rule: storage/mapping/code-level policy опираются на semantic keys. `enabled=true` + blank `template_body` является явным `template body missing` state из ADR-0006. | `REQ-05`, `REQ-06`, `REQ-09`, `REQ-13`, ADR-0005, ADR-0006 |

## Design Questions

Открытых вопросов, меняющих scope FT-4564, acceptance criteria, architectural baseline ADR-0005 / ADR-0006 или concrete persistence contract, сейчас нет.

Ранее открытая persistence ambiguity закрыта `SD-11`: FT-4564 использует отдельную feature-local vendor-scoped storage surface на semantic keys. Если implementation discovery покажет, что `SD-11` нельзя реализовать без нового cross-feature config contract, выполнение FT-4564 должно остановиться и обновить этот `spec.md` или ADR до runtime changes.

## Target Architecture

Целевая архитектура FT-4564 добавляет page-specific Liquid как bounded policy layer поверх существующего customer-facing rendering, а не как redesign публичного pipeline.

### Semantic Page-Key Model

| Vendor-facing key | Internal semantic key | TO-BE covered surfaces | Minimum Liquid context | Notes |
| --- | --- | --- | --- | --- |
| `show` | `order_show` | `vendor/orders/show` | `vendor`, `order`, page state metadata, helpers | Stable public order page |
| `created` | `order_created` | `vendor/orders/created` | `vendor`, `order`, page state metadata, helpers | Post-create surface, не generic order-state alias |
| `canceled` | `order_canceled` | `vendor/orders/canceled` | `vendor`, `order`, page state metadata, helpers | State-specific order page |
| `paid` | `payment_success` | `vendor/orders/paid`; `vendor/payment/show(state=success)` | `vendor`, success outcome/state metadata, helpers, `order?` | `order` required на `vendor/orders/paid` и optional на callback surface; transition bridges и eventual `vendor/orders/show` не принадлежат этому key |
| `failed` | `payment_failure` | `vendor/payment/show(state=failure)` | `vendor`, failure outcome/state metadata, helpers, `order?` | Если flow заканчивается на `vendor/orders/canceled`, ownership у `canceled`; если status-only failure redirect заканчивается на `vendor/orders/show`, ownership у `show` |

### Resolution Contract

1. Existing controller / callback entrypoint определяет concrete runtime surface.
2. Surface classification сопоставляет этот surface с semantic page key.
3. Resolution layer читает vendor-configurable mode для semantic page key.
4. Если page-specific Liquid выбран и Liquid template body существует, собирается Liquid context и выполняется Liquid render.
5. Если page-specific Liquid выбран и Liquid template body отсутствует, рендерится existing baseline page implementation для той же страницы / surface и публикуется отдельный diagnostic signal.
6. Если page-specific Liquid не выбран, request использует текущий global engine behavior.
7. Если Liquid render падает на runtime, baseline fallback не применяется; сохраняется fail-fast path с error reporting и vendor context.

### Surface Ownership Contract

| Concrete render surface | Owning vendor-facing key | Internal semantic key | Ownership note |
| --- | --- | --- | --- |
| `vendor/orders/show` | `show` | `order_show` | Regular public order page |
| `vendor/orders/created` | `created` | `order_created` | Post-create customer-facing page |
| `vendor/orders/canceled` | `canceled` | `order_canceled` | Canceled-state order page |
| `vendor/orders/paid` | `paid` | `payment_success` | Paid-state order page |
| `vendor/payment/show(state=success)` | `paid` | `payment_success` | Callback/result success page |
| `vendor/payment/show(state=failure)` | `failed` | `payment_failure` | Callback/result failure page |
| `/payments/success|failure` redirect bridge, provider redirects, `status=*` PRG bridge | none | n/a | Transition-only flow element, не vendor-configurable page unit; `status=failed` не расширяет ownership `failed` на `vendor/orders/show` |

Acceptance для vendor-facing keys `paid` и `failed` считается по любому renderable customer-facing surface, который этот contract относит к соответствующему key. Redirect bridges и eventual state pages сами по себе не расширяют ownership page key; status-only failure landing на `vendor/orders/show` проверяется как `show` boundary.

### Architectural Invariants

- Transition-only bridges (`/payments/success|failure` redirects и `status=*`) не являются vendor-configurable page units.
- Каждый renderable customer-facing surface в scope FT-4564 принадлежит ровно одному vendor-facing page key; overlap между page keys запрещён.
- Provider-hosted payment forms и widgets остаются вне TO-BE architecture FT-4564.
- TO-BE architecture не требует global engine switch и не вводит explicit per-page selector `React`.
- Existing baseline page implementation для каждой покрываемой surface остаётся обязательной опорной точкой для default path и missing-template fallback.

## Operator Authoring Contract

Detailed screen design живёт в [`operator-ui/README.md`](operator-ui/README.md). Этот `spec.md` фиксирует только solution-level operator invariants:

1. Entry point открывает bounded `Liquid pages` catalog для `show`, `paid`, `failed`, `created`, `canceled`.
2. `Liquid pages` catalog не даёт оператору создавать или удалять FT-4564 page keys вручную; rows принадлежат system-owned catalog.
3. Multi-surface keys не схлопываются в synthetic aggregate badge, а показывают verdict per covered surface. Для `paid` это минимум `vendor/orders/paid` и `vendor/payment/show(state=success)`.
4. Dedicated editor для выбранного page key объединяет body authoring, toggle enable/disable, preview entrypoint и context reference на одной surface.
5. Catalog action `Preview` не создаёт отдельный detached flow, а deep-link-ит в `Предпросмотр` tab того же editor-а.
6. Preview iframe в editor-е открывает direct customer-facing public render surface, а не `/operator/shop` или другую operator route; preview context выбирается явно, order-aware pages не auto-pick-ают случайный live order, а payment-result preview без `order` использует honest baseline `vendor + outcome metadata`.

### Public Preview URL Contract

Operator preview для FT-4564 не использует operator shop preview shell. Editor tab `Предпросмотр` должен iframe-ить public render route для выбранной customer-facing surface, с preview-only query params для доступа к неопубликованному магазину и подавления побочных эффектов.

Для replayable order-state surfaces (`show`, `canceled`, `paid`) preview открывает public order URL, на котором клиент видит соответствующую страницу заказа. Для `created` canonical runtime остаётся post-create render path, поэтому editor preview использует preview-only forced public order render через тот же `/orders/:external_id` host/path и `view=created`; это не превращает `created` в regular public order-state branch. Для payment-result surfaces preview не должен выбирать конкретного payment provider. Он использует provider-neutral preview-only public entrypoint, который рендерит тот же baseline surface `vendor/payment/show` с `state=success` или `state=failure`, но не выполняет provider callback logic, не пишет callback logs и не редиректит через payment bridge.

Preview URL строится так:

1. Base берётся из `current_vendor.iframe_preview_url`, чтобы использовать HTTPS subdomain host и избежать mixed-content / custom-domain проблем.
2. `path` заменяется на concrete public render path из таблицы ниже.
3. `query` заменяется на deterministic preview query из таблицы ниже; query из base не конкатенируется строками.
4. Resulting URL используется напрямую как `iframe[src]` внутри editor tab. `/operator/shop`, operator sidebar, publish controls и device-mode shell в этот iframe не попадают.

| Page key | Preview surface | Required selector context | Public path | Query params |
| --- | --- | --- | --- | --- |
| `show` | `orders/show` | выбранный order | `/orders/:external_id` | `no_canonical_redirect=true`, `page_template_preview=true`, `preview=:vendor_preview_code`, `view=show` |
| `created` | `orders/created` preview-only forced render of post-create surface | выбранный order | `/orders/:external_id` | `no_canonical_redirect=true`, `page_template_preview=true`, `preview=:vendor_preview_code`, `view=created` |
| `canceled` | `orders/canceled` | выбранный order | `/orders/:external_id` | `no_canonical_redirect=true`, `page_template_preview=true`, `preview=:vendor_preview_code`, `view=canceled` |
| `paid` | `orders/paid` | выбранный surface + выбранный order | `/orders/:external_id` | `no_canonical_redirect=true`, `page_template_preview=true`, `preview=:vendor_preview_code`, `view=paid` |
| `paid` | `payment/show(state=success)` | выбранный surface; order optional | `/payments/page_template_preview/success` | `no_canonical_redirect=true`, `page_template_preview=true`, `preview=:vendor_preview_code`, optional `order_id=:order_id` when explicitly selected |
| `failed` | `payment/show(state=failure)` | выбранный surface; order optional | `/payments/page_template_preview/failure` | `no_canonical_redirect=true`, `page_template_preview=true`, `preview=:vendor_preview_code`, optional `order_id=:order_id` when explicitly selected |

Для order-aware preview (`orders/*`) `:external_id` берётся из явно выбранного order. `view=created` допустим только как preview-only forced render выбранного post-create surface и не является доказательством regular show-state reachability. Для payment-result preview (`payment/show`) отсутствие `order_id` является валидным preview состояния `order optional`; если оператор явно выбрал order context, `order_id` добавляется для проверки order-aware variant.

Generic `/payments/success|failure` routes остаются transition-only bridge текущего runtime и не становятся preview target для FT-4564. Provider-specific routes вида `/payments/<provider>/success|failure` остаются real callback / return entrypoints и тоже не являются canonical preview target, потому что они привязывают authoring preview к конкретной платёжке и могут иметь provider-specific side effects.

## Concrete Contracts

| Feature ref / contract | Concrete shape | Notes |
| --- | --- | --- |
| `REQ-05` / `CTR-01` | FT-4564 resolver checks page-specific Liquid opt-in for matched in-scope surfaces; when Liquid is not opted-in, current global engine behavior remains the baseline | Это вводит dedicated public page templates без global engine migration |
| `REQ-02` / `REQ-03` | Current callback result rendering uses one shared template key `payment/show` and distinguishes outcome via `state: success/failure` | Это поддерживает semantic outcome keys лучше, чем per-template config |
| `REQ-04` / `REQ-07` / `REQ-08` | Order-state rendering uses distinct templates `orders/show`, `orders/created`, `orders/canceled`, `orders/paid` | Эти surfaces уже ближе к stable page-like units |
| `REQ-09` | Page-key resolver обязан различать `Liquid selected` и `Liquid template body exists` | Это explicit contract из ADR-0006 |
| `REQ-10` | Minimum Liquid context для in-scope surfaces начинается с `vendor`, outcome/state metadata и standard helpers; order-state surfaces дополнительно несут `order` | Это concretize-ит canonical context contract, не привязывая его к одному concrete template path |
| `REQ-11` | Payment result surfaces не могут честно требовать universally available `order`; `order` остаётся optional и шаблоны должны оставаться валидными без него | Это следует из current runtime inventory и защищает order-less success/failure flows |
| `CON-04`, `EC-09`, `EC-11` | Liquid-specific activation для FT-4564 допускается только после baseline automated regression protection для текущих React-based customer-facing order/result surfaces в scope | Это solution-level rollout gate; exact suite composition остаётся downstream responsibility `implementation-plan.md` |
| `REQ-12`, `EC-12`, `EC-16` | Operator UI exposes bounded fixed catalog instead of ad-hoc public page key creation | Это сохраняет discoverability и не превращает FT-4564 в общий page builder |
| `REQ-13` | Effective status для multi-surface keys раскрывается per covered surface; `paid` не схлопывает `vendor/orders/paid` и `vendor/payment/show(state=success)` в один synthetic badge | Иначе operator UI теряет честность относительно FT-4564 opt-in и fallback semantics |
| `REQ-14` | Selected editor surface должна держать authoring, preview и status toggle рядом; catalog action `Preview` deep-link-ит в `Предпросмотр` tab и требует explicit context selection вместо hidden auto-pick | Иначе feature теряет discoverable authoring flow и честный preview contract |
| `REQ-15` | Context reference в operator editor обязан явно маркировать `order` как `required` / `optional` / `unavailable` по page-key family | Это concretize-ит problem-space contract для authoring UX |
| `SD-11` | Dedicated persistence shape: `VendorPageTemplate(vendor_id, semantic_key, enabled, template_body, timestamps)`, unique `[vendor_id, semantic_key]`, nullable `template_body`, validation against semantic catalog `payment_failure`, `payment_success`, `order_show`, `order_created`, `order_canceled`. | `enabled=false` означает default path; `enabled=true` + blank `template_body` означает missing-template fallback to baseline implementation; `enabled=true` + present body означает Liquid candidate до runtime render. UI aliases map one-way to semantic keys and are not stored as config keys |

## Failure Handling / Fallback

- Принятая policy для FT-4564: resolver различает `Liquid selected + template body exists` и `Liquid selected + template body absent`.
- Случай `template body absent` обязан уходить в existing baseline page implementation до explicit Liquid-render path.
- Payment callbacks уже имеют logging paths на уровне controller logic, а vendor controllers подключают `VendorBugsnag`, который обогащает уведомления vendor metadata.
- Runtime render error после выбора Liquid остаётся fail-fast; verify contract требует end-to-end подтверждения этого error-reporting path на opted-in Liquid surface.

## Rollout / Backout

| Rollout ID | Stage | Trigger | Backout |
| --- | --- | --- | --- |
| `RB-01` | Spec-ready baseline | Этот `spec.md` переведён в `status: active` до старта реализации и фиксирует selected solution | Никаких runtime-изменений не выкатывается; сохраняется текущее baseline behavior |
| `RB-02` | Baseline React regression gate | До Liquid-specific runtime changes подтверждены baseline automated regression protection и зелёный baseline run локально и в CI для текущих customer-facing order/result surfaces в scope | При непрохождении gate FT-4564 не переходит к Liquid-specific activation; runtime остаётся на текущем baseline behavior |
| `RB-03` | Plan-ready / implementation activation | После перевода `implementation-plan.md` в `status: active`, прохождения `RB-02` и начала runtime-изменений по FT-4564 | Откат runtime делается по release policy: redeploy предыдущего Docker image tag, без feature-flag partial rollback. `vendor_page_templates` остаётся inert schema/data для старого образа; emergency rollback не требует удаления таблицы. Если миграция или persistence shape `SD-11` оказываются несовместимыми до production deploy, остановить rollout до runtime activation и обновить spec/ADR |

## ADR Dependencies

| ADR | Current `decision_status` | Used for | Execution rule |
| --- | --- | --- | --- |
| [../../engineering/adr/0003-per-action-engine-override.md](../../engineering/adr/0003-per-action-engine-override.md) | `accepted` | Baseline contract для explicit engine override и fail-fast semantics после explicit Liquid selection | FT-4564 не должна неявно сломать этот baseline |
| [../../engineering/adr/0005-ft-4564-order-page-resolution-and-fallback-policy.md](../../engineering/adr/0005-ft-4564-order-page-resolution-and-fallback-policy.md) | `accepted` | Canonical architectural choice для semantic keys, runtime error policy и bounded refactor | Этот `spec.md` concretize-ит принятое решение на feature level |
| [../../engineering/adr/0006-missing-custom-liquid-template-falls-back-to-existing-baseline-page.md](../../engineering/adr/0006-missing-custom-liquid-template-falls-back-to-existing-baseline-page.md) | `accepted` | Platform-level rule для случая `Liquid selected + template body absent` | Любая implementation detail для FT-4564 должна сохранить fallback на existing baseline page implementation |

---
title: "FT-4564: Implementation Plan"
doc_kind: feature
doc_function: derived
purpose: "Execution-план реализации FT-4564. Фиксирует discovery context, sequencing, риски и test strategy без переопределения canonical feature-фактов."
derived_from:
  - feature.md
  - spec.md
  - runtime-surfaces.md
  - operator-ui/README.md
status: active
audience: humans_and_agents
must_not_define:
  - ft_4564_scope
  - ft_4564_solution_design
  - ft_4564_acceptance_criteria
  - ft_4564_blocker_state
---

# FT-4564: Implementation Plan

## Цель текущего плана

Довести FT-4564 до runtime-реализации и operator authoring surface для public Liquid page templates без неявного расширения scope за пределы order/payment result surfaces и с достаточной automated/manual verify базой для `SC-*`, `NEG-*`, `CHK-*` и `EVID-*` из `feature.md`.

Grounding по current runtime surfaces, semantic mapping и context note вынесен в [`runtime-surfaces.md`](runtime-surfaces.md); этот план использует его как upstream reference и не дублирует surface inventory внутри execution-документа.

## Current State / Reference Points

| Path / module | Current role | Why relevant | Reuse / mirror |
| --- | --- | --- | --- |
| `app/controllers/concerns/render_view.rb` | Единая точка входа в customer-facing render flow: нормализует template path и dispatch по engine | FT-4564 должна встроить page-specific Liquid opt-in для in-scope surfaces и уметь возвращаться к baseline implementation, когда Liquid не выбран или template body отсутствует | Сохранить current normalization/dispatch behavior; новый page-template resolver выполняется до final render decision |
| `app/controllers/concerns/render_react.rb` | Хранит current concrete template-to-component mapping для React surfaces, включая `vendor/orders/*` и `vendor/payment/show` | Existing baseline fallback должен уметь вернуться в текущую concrete implementation, а не в придуманную новую поверхность | Использовать текущий mapping как source of truth для baseline fallback на non-Liquid path |
| `app/controllers/vendor/orders_controller.rb` | Владение customer-facing order surfaces: `show`, `paid`, `canceled`, `created`, PRG-redirect по `status`, post-create payment flow | Основной write-surface для `order_show`, `order_created`, `order_canceled` и paid-state order page; здесь нельзя сломать текущие props и PRG-поведение | Сохранить текущие ветки и props contracts; page-key opt-in добавлять поверх существующих render branches |
| `app/controllers/vendor/payments_controller.rb` | Generic success/failure bridge: при наличии `order_id` только redirect, иначе flash + vendor root | Spec уже фиксирует, что это transition-only surface, а не самостоятельный template render | Сохранить bridge-only semantics; не делать эти routes vendor-configurable page units |
| `app/controllers/vendor/payment/base_controller.rb` | Provider callback baseline: часть success flows редиректит на order page, иначе рендерит `payment/show`; failure рендерит `payment/show`; содержит callback logging и Bugsnag-on-logger-error path | Нужен для `payment_success` / `payment_failure` semantics, optional `order` contract и error/evidence path | Переиспользовать существующий `payment/show` baseline и текущий logging/error-reporting pattern |
| `app/controllers/vendor/payment/custom_controller.rb`, `app/controllers/vendor/payment/robokassa_controller.rb` | Provider-specific controllers уже наследуют или повторяют callback result rendering через `payment/show` | FT-4564 не должна сломать provider-specific entrypoints при добавлении page-level Liquid selection | Выравнивать поведение через shared base/resolver, а не плодить per-provider forks |
| `config/routes/public_shop.rb`, new provider-neutral payment preview controller | TO-BE preview-only public entrypoint `/payments/page_template_preview/:state` для `payment/show(state=success|failure)` | Operator preview не должен использовать `/payments/<provider>/*` callback как canonical page URL | Добавить отдельный preview-only route/action, который рендерит `vendor/payment/show` без provider callback logic, callback logs и redirect bridge |
| `app/models/vendor_theme.rb` | Хранит engine enum (`react` / `liquid`) и текущий global engine switch | FT-4564 добавляет per-page opt-in поверх current engine model; нельзя тихо сломать global behavior или превратить global engine в per-page selector | Сохранить existing engine enum; per-page settings реализовать через selected `SD-11` persistence surface |
| `app/models/vendor_page_template.rb` | Dedicated persistence surface для FT-4564 public page templates | Это selected `SD-11` storage для semantic page keys, opt-in и template body | Держать validation, semantic catalog и UI aliases внутри этой модели; не использовать theme engine как скрытый per-page selector |
| `app/services/vendor_page_template_resolver.rb` | Resolver semantic page key -> covered concrete render surface -> enabled template record | Централизует resolution unit для runtime и tests | Использовать в shared render path вместо ad-hoc checks в отдельных controllers |
| `app/controllers/operator/page_templates_controller.rb` | Planned operator controller для bounded catalog/editor/preview public page templates | FT-4564 operator UI должен быть отдельным authoring surface для system-owned page keys | Реализовать catalog/editor/update/preview через `VendorPageTemplate`; не смешивать с другими content authoring flows |
| `app/views/operator/page_templates/*` | Planned operator views для catalog/editor/context reference | Нужны для fixed catalog, per-surface status, save/reload и preview deep-link | Reuse operator layout patterns, tabs и `bottom_form_panel`, но держать public page templates отдельным surface |
| `app/controllers/operator/shop_controller.rb` | Existing общий preview магазина с operator layout и device modes | FT-4564 preview не должен iframe-ить `/operator/shop`, потому что editor preview обязан показывать direct customer-facing public render surface без вложенной операторской | Не использовать как render target для Liquid page editor preview; путь строится в `Operator::PageTemplatesController` напрямую по contract из `spec.md` |
| `config/application.yml` | Source of concrete template keys / titles, включая `orders/show`, `orders/paid`, `orders/canceled`, `orders/created`, `orders/payment`, `payment/show` | Нужен для grounding semantic page keys к current concrete surfaces и явной фиксации adjacent out-of-scope surface `orders/payment` | Использовать текущий config registry как reference для concrete template keys и naming; FT-4564 storage остаётся отдельным по `SD-11` |
| `spec/controllers/vendor/orders_controller_spec.rb` | Уже покрывает create/show/pay flow, PRG-поведение по `status`, render template branches `created` / `payment` | Это ближайший current regression harness для order-state surfaces | Расширять существующий controller spec под page opt-in, priority и fallback cases |
| `spec/controllers/vendor/payment/rbk_money_controller_spec.rb` | Минимальный текущий provider callback spec | Нужен как starting point для payment result surface coverage | Либо расширить этот spec, либо вынести shared callback coverage в новые vendor/payment specs по тому же fixture-based стилю |
| `spec/controllers/operator/page_templates_controller_spec.rb`, `spec/views/operator/page_templates/*_spec.rb` | Planned specs для FT-4564 operator authoring surface | Нужны для fixed catalog, editor, preview routing и persistence contracts | Покрыть catalog, per-surface status, save/reload, context reference и preview selection |
| `spec/models/vendor_page_template_spec.rb` | Model coverage для selected persistence | Нужен для validation, uniqueness, syntax validation и semantic mapping | Держать как fast regression net для `SD-11` |
| `spec/models/vendor_theme_spec.rb` | Базовый coverage для engine model | Нужен как regression net вокруг current engine semantics | Держать в sync, если FT-4564 затронет theme-level opt-in shape |

## Test Strategy

| Test surface | Canonical refs | Existing coverage | Planned automated coverage | Required local verify | Manual-only gap / justification | Manual-only approval ref |
| --- | --- | --- | --- | --- | --- | --- |
| Order-state surfaces `show` / `paid` / `created` / `canceled` | `REQ-03`, `REQ-04`, `REQ-05`, `REQ-07`, `REQ-08`, `EC-03`, `SC-01`, `SC-03`, `SC-05`, `SC-06`, `SC-07`, `CHK-01`, `CHK-02`, `CHK-05`, `CHK-06`, `CHK-12` | Есть existing controller coverage для `show`, PRG statuses и post-create branches, но нет per-page Liquid opt-in regression tests | Расширить controller/request coverage на opted-in order-state pages, baseline fallback и missing-template behavior; для `paid` отдельно покрыть order-state surface `vendor/orders/paid` | `APP_TEST bundle exec rspec spec/controllers/vendor/orders_controller_spec.rb spec/models/vendor_page_template_spec.rb spec/models/vendor_theme_spec.rb` | Визуальная проверка реального Liquid HTML и state-specific copy остаётся manual-only до появления стабильного browser harness | `AG-01` |
| Payment result surfaces `payment/show`, provider-neutral preview route и transition bridges | `REQ-02`, `REQ-03`, `REQ-05`, `REQ-06`, `REQ-10`, `REQ-11`, `SC-02`, `SC-03`, `SC-09`, `SC-10`, `CHK-03`, `CHK-04`, `CHK-09`, `CHK-11`, `CHK-13` | Есть только минимальный provider callback spec; generic bridge coverage отсутствует | Добавить coverage для `Vendor::PaymentsController` redirect-only behavior, provider callback render/redirect outcome paths и `/payments/page_template_preview/:state`, включая optional `order` contract для payment result pages | `APP_TEST bundle exec rspec spec/controllers/vendor/payments_controller_spec.rb spec/controllers/vendor/payment spec/controllers/vendor/orders_controller_spec.rb` | Реальный payment provider return flow остаётся manual-only: требует stage/vendor credentials и наблюдения customer-facing переходов | `AG-01` |
| Failure semantics: `template body absent` и runtime Liquid render error | `REQ-09`, `SC-08`, `NEG-01`, `NEG-02`, `CHK-07`, `CHK-10` | Детерминированного automated coverage сейчас нет | Добавить deterministic specs на missing-template fallback и fail-fast path; отдельно зафиксировать assertion на error-reporting path (`Bugsnag` / logger behavior) без silent baseline fallback | `APP_TEST bundle exec rspec spec/controllers/vendor/orders_controller_spec.rb spec/controllers/vendor/payment spec/models/vendor_page_template_spec.rb` | End-to-end подтверждение vendor-context error reporting остаётся manual-only, если локально нельзя честно воспроизвести интеграционный exception path | `AG-01` |
| Operator authoring surface: catalog, editor, preview, context reference | `REQ-12`, `REQ-13`, `REQ-14`, `REQ-15`, `EC-16`, `SC-11`, `SC-12`, `SC-13`, `SC-14`, `SC-15`, `CHK-14`, `CHK-15`, `CHK-16`, `CHK-17`, `CHK-18`, `CHK-19` | Dedicated automated coverage для bounded Liquid catalog и editor сейчас отсутствует | Добавить operator controller/view specs для fixed page catalog, отсутствия ad-hoc create/delete для public page keys, per-surface status semantics для `paid`, deterministic save/reload persistence для Liquid template body, explicit preview surface selection для multi-surface keys, direct public preview URL contract, context reference и preview entrypoint/deep-link | `APP_TEST bundle exec rspec spec/controllers/operator/page_templates_controller_spec.rb spec/views/operator/page_templates spec/models/vendor_page_template_spec.rb spec/models/vendor_theme_spec.rb` | Visual/manual проверка direct public preview iframe и editor ergonomics остаётся human-only, пока нет browser harness для операторской; сохранение и повторная загрузка template body не являются manual-only gap и должны проходить в automated specs | `AG-01` |
| Config / registry consistency для semantic keys, selected persistence и current engine behavior | `REQ-05`, `REQ-06`, `REQ-10`, `SD-11`, `CHK-02`, `CHK-03`, `CHK-11` | Dedicated coverage для `VendorPageTemplate` отсутствует до реализации | Добавить model specs для `VendorPageTemplate`: semantic catalog validation, unique `[vendor_id, semantic_key]`, `enabled` independent from nullable `template_body`, no mutation of `VendorTheme#engine`, UI aliases are mapped and not stored | `APP_TEST bundle exec rspec spec/models/vendor_page_template_spec.rb spec/models/vendor_theme_spec.rb` | Нет | `none` |
| Metric readiness для pilot adoption | `MET-01`, `CHK-20` | Dedicated metric-readout procedure сейчас отсутствует, потому что persistence surface ещё не реализован | После `SD-11` persistence implementation зафиксировать query/report procedure, считающую vendors с хотя бы одной enabled semantic-key Liquid page; post-release pilot value добавляется тем же artifact, когда доступен | Procedure artifact review; no separate local suite | Нет: это pre-release readiness artifact, не live pilot measurement | `none` |

`APP_TEST` в таблице означает выполнение Ruby/RSpec внутри app container через dip:

```bash
direnv exec . mise exec -- dip compose run --rm -e RAILS_ENV=test -e DATABASE_URL=postgresql://postgres:postgres@postgres:5432/merchantly_test -e REDIS_URL=redis://redis:6380/1 -e DEV_HOST= app
```

## Open Questions / Ambiguities

| Open Question ID | Question | Why unresolved | Blocks | Default action / escalation owner |
| --- | --- | --- | --- | --- |
| `OQ-02` | Каким минимальным deterministic harness покрыть runtime Liquid render error и vendor-context error reporting path? | Current code содержит logging и Bugsnag path, но end-to-end воспроизводимость через existing test harness ещё не доказана | `STEP-05`, `WS-4` | Сначала пробовать controller/spec harness со stubbed render error и assertion на reporting; если этого недостаточно, зафиксировать manual-only gap и вынести approval на stage verify |
| `OQ-03` | Какой existing source of truth использовать для explicit preview context selector у order-aware страниц и `paid` surface variants, чтобы preview не опирался на случайный live order? | Preview contract уже выбран upstream, но execution-level источник demo/test context для editor surface пока не зафиксирован доказательством | `STEP-04A`, `WS-5` | Предпочесть reuse существующего deterministic order/demo source; не делать hidden auto-pick и не менять уже зафиксированный preview contract |

## Environment Contract

| Area | Contract | Used by | Failure symptom |
| --- | --- | --- | --- |
| setup | Все локальные команды выполняются через `direnv exec .`; Ruby/RSpec команды выполняются через `mise exec -- dip ... app`; тесты опираются на fixtures-only policy и текущий Rails test environment | `STEP-01`-`STEP-05` | Спеки не поднимаются локально, используются ad-hoc seed/data setup или окружение расходится с project policy |
| test | Эталон локальной verify для implementation step — targeted RSpec suites по order/payment controllers и model contracts внутри dip app container; итоговый regression verdict даёт зелёный локальный прогон и CI | `STEP-03`, `STEP-04`, `STEP-05`, `CP-02` | Изменение прошло только ручную проверку, но не имеет deterministic automated regression net |
| access / network / secrets | Для локальной реализации внешние secrets не нужны; для manual evidence по payment return flows и visual screenshots могут понадобиться stage/vendor credentials, browser access и согласованный verify vendor | `STEP-06`, `AG-01` | Работа требует использования реальных credentials или stage access без явного human approval |

## Preconditions

| Precondition ID | Canonical ref | Required state | Used by steps | Blocks start |
| --- | --- | --- | --- | --- |
| `PRE-01` | `feature.md` | `status: active`, `delivery_status: planned`, canonical `REQ-*` / `SC-*` / `CHK-*` зафиксированы | `STEP-01`-`STEP-06` | yes |
| `PRE-02` | `spec.md`, `runtime-surfaces.md` | `spec.md` имеет `status: active`; `runtime-surfaces.md` фиксирует current surface inventory; актуальные accepted `SD-*` решения и grounding считаются execution baseline | `STEP-01`-`STEP-06` | yes |
| `PRE-03` | ADR-0005, ADR-0006 | Оба ADR имеют `decision_status: accepted` | `STEP-01`-`STEP-06` | yes |
| `PRE-04` | `memory-bank/engineering/testing-policy.md` | Fixtures-only и sufficient coverage policy приняты как execution baseline | `STEP-03`-`STEP-05` | yes |
| `PRE-05` | `SD-11` | Dedicated `VendorPageTemplate` persistence surface реализован и покрыт model specs: opt-in state отделён от nullable `template_body`, semantic catalog validation работает, UI aliases не хранятся как config keys, `VendorTheme#engine` не используется как скрытый storage | `STEP-02`, `STEP-03`, `STEP-04`, `STEP-04A`, `STEP-05B` | no; blocks runtime changes after `STEP-01` |
| `PRE-06` | `operator-ui/README.md` | Operator screen design reference согласован как current feature-local baseline для catalog/editor/preview/context reference | `STEP-04A` | yes |

## Workstreams

| Workstream | Implements | Result | Owner | Dependencies |
| --- | --- | --- | --- | --- |
| `WS-1` | `REQ-05`, `REQ-06`, `REQ-09`, `SD-01`, `SD-03`, `SD-05`, `SD-11` | Resolver/page-key baseline и selected persistence surface встроены поверх current render pipeline; `PRE-05` закрывается внутри `STEP-01` до runtime changes | either | `PRE-01`, `PRE-02`, `PRE-03` |
| `WS-2` | `REQ-04`, `REQ-07`, `REQ-08`, `SC-01`, `SC-06`, `SC-07` | Order-state surfaces поддерживают per-page Liquid opt-in и baseline fallback | either | `WS-1` |
| `WS-3` | `REQ-02`, `REQ-03`, `REQ-10`, `REQ-11`, `SC-02`, `SC-03`, `SC-09`, `SC-10`, `SD-02`, `SD-04` | Payment result surfaces и transition bridges согласованы с semantic page keys и optional `order` contract | either | `WS-1` |
| `WS-4` | `CHK-02`, `CHK-03`, `CHK-07`, `CHK-08`, `CHK-11`, `CHK-12`, `CHK-13` | Automated regression net и manual evidence plan покрывают implementation surface | either | `WS-2`, `WS-3` |
| `WS-5` | `REQ-12`, `REQ-13`, `REQ-14`, `REQ-15`, `SC-11`, `SC-12`, `SC-13`, `SC-14`, `SC-15`, `SD-08`, `SD-09`, `SD-10`, `SD-11` | Operator authoring surface даёт bounded catalog, deterministic template persistence, honest per-surface status semantics, editor, preview и context reference | either | `WS-1`, provider-neutral payment preview route from `WS-3`, `PRE-06` |

## Approval Gates

| Approval Gate ID | Trigger | Applies to | Why approval is required | Approver / evidence |
| --- | --- | --- | --- | --- |
| `AG-01` | Нужен stage/browser verify с реальными payment return flows, vendor credentials, внешним callback path или ручным сбором screenshot/video evidence | `STEP-06`, manual-only gaps из `Test Strategy` | Это внешне-эффективная и потенциально чувствительная операция; её нельзя выполнять автономно без подтверждённого окружения и владельца фичи | Feature owner / human reviewer; evidence: явное сообщение в issue/PR или отдельный verify note с подтверждением допуска |

## Порядок работ

| Step ID | Actor | Implements | Goal | Touchpoints | Artifact | Verifies | Evidence IDs | Check command / procedure | Blocked by | Needs approval | Escalate if |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| `STEP-01` | either | `REQ-05`, `REQ-06`, `REQ-09`, `SD-01`, `SD-03`, `SD-05`, `SD-11` | Реализовать semantic-key inventory и selected `VendorPageTemplate` opt-in/template storage без второго source of truth | `config/application.yml`, `app/models/vendor_theme.rb`, new migration/model for `vendor_page_templates`, `spec/models/vendor_page_template_spec.rb` | Consistent semantic-key baseline для FT-4564; UI aliases map to semantic keys but are not stored; `PRE-05` закрыт model/config evidence до runtime changes | `CHK-02`, `CHK-03`, `CHK-11` | `EVID-02` | Review current config/model contracts; затем `APP_TEST bundle exec rspec spec/models/vendor_page_template_spec.rb spec/models/vendor_theme_spec.rb` | `PRE-01`, `PRE-02`, `PRE-03` | none | Selected `SD-11` persistence невозможно реализовать без нового cross-feature config contract |
| `STEP-02` | either | `REQ-05`, `REQ-06`, `REQ-09`, `SD-01`, `SD-03`, `SD-05` | Встроить resolver/fallback policy в render pipeline для matched in-scope surfaces и до final engine dispatch | `app/controllers/concerns/render_view.rb`, `app/controllers/concerns/render_react.rb`, `app/services/vendor_page_template_resolver.rb` | Bounded resolver layer с page-specific Liquid opt-in и deterministic missing-template/runtime-error behavior | `CHK-02`, `CHK-03`, `CHK-07` | `EVID-02` | `APP_TEST bundle exec rspec spec/models/vendor_page_template_spec.rb spec/models/vendor_theme_spec.rb` + targeted controller specs после их добавления | `STEP-01` | none | Для baseline fallback требуется widening за пределы FT-4564 surfaces |
| `STEP-03` | either | `REQ-03`, `REQ-04`, `REQ-07`, `REQ-08`, `SC-01`, `SC-03`, `SC-05`, `SC-06`, `SC-07` | Подключить opted-in Liquid routing к order-state surfaces и сохранить current props/PRG behavior | `app/controllers/vendor/orders_controller.rb`, `spec/controllers/vendor/orders_controller_spec.rb` | Order-state pages (`show`, `paid`, `created`, `canceled`) поддерживают FT-4564 policy без regressions; для `paid` отдельно подтверждён order-state surface `vendor/orders/paid` | `CHK-01`, `CHK-02`, `CHK-05`, `CHK-06`, `CHK-12` | `EVID-01`, `EVID-02`, `EVID-04`, `EVID-05` | `APP_TEST bundle exec rspec spec/controllers/vendor/orders_controller_spec.rb spec/models/vendor_page_template_spec.rb spec/models/vendor_theme_spec.rb` | `STEP-02` | none | Existing PRG/status behavior ломается или order-state pages, включая `vendor/orders/paid`, теряют current props contract |
| `STEP-04` | either | `REQ-02`, `REQ-03`, `REQ-10`, `REQ-11`, `SC-02`, `SC-03`, `SC-09`, `SC-10`, `SD-02`, `SD-04` | Подключить opted-in Liquid policy к payment result surfaces, не превращая transition bridges в render surfaces | `config/routes/public_shop.rb`, `app/controllers/vendor/payments_controller.rb`, `app/controllers/vendor/payment/base_controller.rb`, `app/controllers/vendor/payment/custom_controller.rb`, `app/controllers/vendor/payment/robokassa_controller.rb`, new provider-neutral payment preview controller, `spec/controllers/vendor/payment`, `spec/controllers/vendor/payments_controller_spec.rb` | Payment result rendering, provider-neutral preview route и redirect bridges согласованы с semantic page keys и optional `order` contract | `CHK-03`, `CHK-04`, `CHK-09`, `CHK-11`, `CHK-13` | `EVID-02`, `EVID-03`, `EVID-06` | `APP_TEST bundle exec rspec spec/controllers/vendor/payments_controller_spec.rb spec/controllers/vendor/payment spec/controllers/vendor/orders_controller_spec.rb` | `STEP-02` | none | Provider-specific callback behavior требует cross-feature refactor, который выходит за bounded scope ADR-0005, или preview пытается стать real payment return URL |
| `STEP-04A` | either | `REQ-12`, `REQ-13`, `REQ-14`, `REQ-15`, `SC-11`, `SC-12`, `SC-13`, `SC-14`, `SC-15`, `SD-08`, `SD-09`, `SD-10`, `SD-11` | Реализовать operator catalog/editor/preview/context reference и honest per-surface status semantics для public page templates | `config/operator_content_navigation.rb`, `app/controllers/operator/page_templates_controller.rb`, `app/views/operator/page_templates/index.html.haml`, `app/views/operator/page_templates/edit.html.haml`, `app/views/operator/page_templates/_page_template.html.haml`, `VendorPageTemplate`, relevant operator locales | Operator authoring surface для FT-4564 page keys с bounded catalog, deterministic template body save/reload, отсутствием ad-hoc create/delete для public page keys, honest status semantics, direct public preview iframe, explicit preview surface selection для multi-surface keys и context reference | `CHK-14`, `CHK-15`, `CHK-16`, `CHK-17`, `CHK-18`, `CHK-19` | `EVID-02`, `EVID-08` | `APP_TEST bundle exec rspec spec/controllers/operator/page_templates_controller_spec.rb spec/views/operator/page_templates spec/models/vendor_page_template_spec.rb spec/models/vendor_theme_spec.rb` | `STEP-01`, `STEP-04`, `PRE-06`, `OQ-03` | none | Реализация operator surface требует hidden auto-pick preview context, iframe на `/operator/shop`, не может честно показать per-surface verdict для `paid`, пытается использовать `/payments/<provider>/*` вместо provider-neutral preview route или тащит redesign всего content authoring |
| `STEP-05` | either | `REQ-09`, `NEG-01`, `NEG-02`, `CHK-07`, `CHK-08`, `CHK-10`, `CHK-11`, `CHK-12`, `CHK-13`, `CHK-14`, `CHK-15`, `CHK-17`, `CHK-18`, `CHK-19` | Закрыть automated regression net и подтвердить failure semantics, operator authoring contracts, traceability и green suites | `spec/controllers/vendor/orders_controller_spec.rb`, `spec/controllers/vendor/payment`, `spec/controllers/vendor/payments_controller_spec.rb`, `spec/controllers/operator/page_templates_controller_spec.rb`, `spec/views/operator/page_templates`, `spec/models/vendor_page_template_spec.rb`, `spec/models/vendor_theme_spec.rb`, `memory-bank/features/FT-4564/feature.md` | Green automated verify bundle и актуальная traceability | `CHK-07`, `CHK-08`, `CHK-10`, `CHK-11`, `CHK-12`, `CHK-13`, `CHK-14`, `CHK-15`, `CHK-17`, `CHK-18`, `CHK-19` | `EVID-02`, `EVID-07` | `APP_TEST bundle exec rspec spec/controllers/vendor/orders_controller_spec.rb spec/controllers/vendor/payments_controller_spec.rb spec/controllers/vendor/payment spec/controllers/operator/page_templates_controller_spec.rb spec/views/operator/page_templates spec/models/vendor_page_template_spec.rb spec/models/vendor_theme_spec.rb` | `STEP-03`, `STEP-04`, `STEP-04A`, `OQ-02` | none | Runtime-error path нельзя доказать deterministic automation или operator catalog/editor contracts расходятся с docs |
| `STEP-05B` | either | `MET-01`, `CHK-20`, `SD-11` | Зафиксировать metric-readout procedure для pilot adoption поверх semantic-key storage | `artifacts/ft-4564/verify/met-01/`, PR/issue verify note, selected query/report based on `vendor_page_templates.semantic_key` + `enabled` | Metric readiness artifact: как считать vendors с хотя бы одной enabled FT-4564 Liquid page; post-release pilot value добавляется тем же artifact, когда доступен | `CHK-20` | `EVID-10` | Сформировать query/report procedure по semantic-key catalog `payment_failure`, `payment_success`, `order_show`, `order_created`, `order_canceled`; review artifact path or PR/issue note | `STEP-01` | none | Readout procedure требует aliases in storage, live production access до approval или расходится с `SD-11` |
| `STEP-05A` | either | `CON-04`, `SD-07`, Feature Flow simplify review | Выполнить отдельный simplify review после зелёных automated suites и до manual acceptance, чтобы не закрыть feature с лишней сложностью, premature abstraction или dead code | Final implementation diff по runtime resolver, controller touchpoints, operator UI и specs | Короткий simplify-review verdict: код минимально сложен или complexity явно обоснована ссылкой на `CON-*`, `FM-*`, `SD-*` или ADR | Feature Flow closure gate | `EVID-09` | После `STEP-05` пройти diff без изменения acceptance: проверить дублирование, вложенность, dead code, unnecessary abstraction; при необходимости упростить и заново прогнать affected suites | `STEP-05` | none | Complexity остаётся без ссылки на `CON-*`, `FM-*`, `SD-*` или ADR, либо simplify changes требуют пересмотра selected solution |
| `STEP-06` | human | `SC-01`, `SC-02`, `SC-03`, `SC-06`, `SC-07`, `SC-13`, `CHK-01`, `CHK-04`, `CHK-05`, `CHK-06`, `CHK-09`, `CHK-16` | Собрать manual evidence по visual/live payment flows и operator preview/editor UX и закрыть manual-only gaps | Stage/browser environment, verify vendor data, operator browser session, evidence artifacts under `artifacts/ft-4564/verify/` | Screenshot/video evidence bundle и финальный human verdict по customer-facing UX и operator authoring flow | `CHK-01`, `CHK-04`, `CHK-05`, `CHK-06`, `CHK-09`, `CHK-16` | `EVID-01`, `EVID-03`, `EVID-04`, `EVID-05`, `EVID-06`, `EVID-08` | Пройти agreed manual procedure по `show`, `failed`, `paid`, `created`, `canceled` и operator editor/preview; для `paid` отдельно собрать evidence по `vendor/orders/paid`, `vendor/payment/show(state=success)` и preview surface selection; приложить screenshots/video и verdict | `STEP-05A`, `AG-01` | `AG-01` | Нет approved stage/vendor access, operator preview не открывается честно, или ручная проверка показывает divergence от canonical `SC-*` |

## Parallelizable work

- `PAR-01` После `STEP-02` workstreams `WS-2` и `WS-3` можно вести параллельно: order-state surfaces и payment result surfaces имеют разные primary controller touchpoints.
- `PAR-02` Model/config consistency checks можно поддерживать параллельно с controller implementation, если write-surface по registry и theme settings уже стабилизирован в `STEP-01`.
- `PAR-03` После стабилизации `STEP-01` runtime workstreams (`WS-2`, `WS-3`) и operator authoring surface (`WS-5`) можно вести параллельно, если они не расходятся по persistence contract и page-key registry; финальная wiring/verification payment-result preview в operator UI зависит от provider-neutral route из `STEP-04`.
- `PAR-04` `STEP-06` нельзя начинать до завершения `STEP-05A` / `CP-04`: manual evidence не заменяет automated regression net, simplify review и не должна маскировать незакрытые failure semantics.

## Checkpoints

| Checkpoint ID | Refs | Condition | Evidence IDs |
| --- | --- | --- | --- |
| `CP-01` | `STEP-01`, `STEP-02`, `CHK-02`, `CHK-03` | Resolver baseline доказан: opted-in Liquid selected deterministically, страницы без opt-in используют baseline implementation, missing-template behavior детерминирован | `EVID-02` |
| `CP-02` | `STEP-03`, `STEP-04`, `CHK-11`, `CHK-12`, `CHK-13` | Order/payment surfaces покрыты automated specs и не расходятся с `feature.md` / `spec.md` contracts | `EVID-02`, `EVID-07` |
| `CP-03` | `STEP-04A`, `STEP-05`, `CHK-14`, `CHK-15`, `CHK-17`, `CHK-18`, `CHK-19` | Operator authoring surface покрыта automated specs и не расходится с `feature.md` / `spec.md` / `operator-ui/README.md` contracts, включая fixed catalog ownership и deterministic preview/editor routing | `EVID-02` |
| `CP-04` | `STEP-05A`, Feature Flow simplify review | Simplify review выполнен после functional verification и до manual acceptance; оставшаяся complexity обоснована canonical refs | `EVID-09` |
| `CP-05` | `STEP-05A`, `STEP-06`, `CHK-01`, `CHK-04`, `CHK-05`, `CHK-06`, `CHK-09`, `CHK-16` | Manual evidence bundle собран после simplify review, и human reviewer подтвердил customer-facing behavior и operator editor/preview UX | `EVID-01`, `EVID-03`, `EVID-04`, `EVID-05`, `EVID-06`, `EVID-08` |
| `CP-06` | `STEP-05B`, `CHK-20`, `MET-01` | Metric-readout procedure для pilot adoption зафиксирована поверх semantic-key storage и не требует aliases in storage | `EVID-10` |

## Plan-local Evidence

| Evidence ID | Artifact | Producer | Path contract | Reused by checks |
| --- | --- | --- | --- | --- |
| `EVID-09` | Simplify-review verdict | implementer / reviewer | PR comment, issue note или `artifacts/ft-4564/verify/simplify-review.md` | `STEP-05A`, `CP-04` |

## Риски исполнения

| Risk ID | Risk | Impact | Mitigation | Trigger |
| --- | --- | --- | --- | --- |
| `ER-01` | Изменение `render_view` затронет shared render pipeline и сломает unrelated customer-facing pages | Регрессии за пределами FT-4564 | Держать bounded resolver layer, не менять `correct_template`, и подтверждать baseline через targeted controller/model specs | Неожиданные падения unrelated vendor controller specs или divergence в existing render branches |
| `ER-02` | Payment callbacks неоднородны: часть flows redirect-only, часть render `payment/show` | Неверный semantic mapping может сломать success/failure UX или optional `order` contract | Сначала закрепить bridge-only semantics generic controller, затем отдельно покрыть provider callback surfaces и provider-neutral preview-only route | Implementation пытается объединить redirect и render surfaces в один несуществующий template contract или привязывает preview к `/payments/<provider>/*` |
| `ER-03` | Runtime Liquid render error трудно воспроизвести deterministic способом | `NEG-02` остаётся без доказательства или маскируется silent fallback | Сначала пробовать stubbed/spec harness, при недостаточности явно оформлять manual-only gap и approval path | Нельзя получить stable automated assertion на fail-fast / error-reporting behavior |
| `ER-04` | Selected `SD-11` persistence окажется недостаточным и потянет новый cross-feature config surface | FT-4564 выйдет за bounded scope и потеряет reversibility | Остановиться на `STOP-02`, поднять upstream decision в `spec.md` / ADR вместо локального ad-hoc расширения | Dedicated `VendorPageTemplate` surface нельзя реализовать без нового platform contract |
| `ER-05` | Operator UI схлопнет multi-surface key в один synthetic status | Оператор не поймёт divergence между `orders/paid` и `payment/show(state=success)` | Показывать `template body missing` и per-surface breakdown в catalog/editor | Preview/editor или catalog показывает `paid` как lossy aggregate verdict |

## Stop conditions / fallback

| Stop ID | Related refs | Trigger | Immediate action | Safe fallback state |
| --- | --- | --- | --- | --- |
| `STOP-01` | `REQ-05`, `SD-01`, ADR-0005 | Реализация требует использовать не `VendorPageTemplate` как storage/effective-status layer для FT-4564 Liquid templates или ломает baseline behavior без opt-in | Остановить runtime changes и вернуть вопрос в `spec.md` / ADR | Текущее baseline behavior без FT-4564 runtime изменений |
| `STOP-02` | `NS-01`, `NS-02`, `SD-02`, `SD-11` | Write-surface внезапно тянет checkout/catalog или требует cross-feature config redesign за пределами selected `VendorPageTemplate` storage | Не расширять scope локально; зафиксировать escalation и обновить upstream artifact | FT-4564 остаётся ограниченной order/payment result surfaces |
| `STOP-03` | `NEG-02`, `AG-01`, `OQ-02` | Error-reporting path нельзя доказать automation, а stage/manual verify не approved | Сохранить manual-only gap как blocker и не закрывать feature как done | Plan-ready / implementation state без ложного verdict по `CHK-07` |
| `STOP-04` | `REQ-12`, `REQ-13`, `SD-08`, `SD-09` | Operator authoring surface требует новый top-level navigation раздел или redesign content authoring за пределами feature scope | Остановить operator UI changes и вернуть вопрос в `spec.md` / `operator-ui/README.md` | Текущий product behavior без частично сломанного authoring UX |

## Готово для приемки

План считается исчерпанным, когда выполнены все условия ниже:

- `WS-1`-`WS-5` завершены без активных `STOP-*`.
- `CP-01`, `CP-02` и `CP-03` доказаны через `EVID-02` и актуальную traceability к `feature.md` / `spec.md` / `operator-ui/README.md`.
- `CP-04` доказан через `EVID-09`: simplify review выполнен отдельным проходом после functional verification и до manual acceptance.
- `CP-05` закрыт, если manual evidence требуется для текущего delivery verdict; при незакрытом `CP-05` feature не переводится в `delivery_status: done`.
- `CP-06` доказан через `EVID-10`: metric readiness procedure для `MET-01` зафиксирована до closure; сам pilot target `>= 1` оценивается post-release.
- Manual-only gaps либо закрыты через `STEP-06`, либо явно остаются открытыми с `AG-01` и не позволяют перевести фичу в `delivery_status: done`.
- Все required local suites зелёные локально, CI подтверждает regression net, а итоговая приёмка идёт по секции `Verify` в `feature.md`.

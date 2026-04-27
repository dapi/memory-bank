---
title: "FT-4564: Operator UI for Public Liquid Page Templates"
doc_kind: feature-support
doc_function: reference
purpose: "Подробный operator UX reference для FT-4564: entrypoint, fixed page-key catalog, enable/disable controls, Liquid editor, preview и context reference."
derived_from:
  - ../feature.md
  - ../spec.md
  - ../../../domain/frontend.md
  - ../../../domain/design-guide/README.md
  - ../../../domain/design-guide/forms.md
  - ../../../domain/design-guide/buttons.md
  - ../../../domain/design-guide/navigation.md
status: active
audience: humans_and_agents
---

# FT-4564: Operator UI for Public Liquid Page Templates

## Роль документа

Этот документ не владеет требованиями или selected solution.

- `feature.md` остаётся canonical owner для `REQ-*`, `EC-*`, `SC-*`, `CHK-*`, `EVID-*`.
- `spec.md` остаётся canonical owner для selected design и solution contracts.
- Этот файл раскрывает feature-local operator UX и screen design настолько подробно, чтобы implementation и review не расходились по ожиданиям.
- Для визуальной привязки screen design см. markdown-мокапы:
  - [`mockups/liquid-pages-list.md`](mockups/liquid-pages-list.md)
  - [`mockups/liquid-page-editor.md`](mockups/liquid-page-editor.md)
  - [`mockups/liquid-page-preview-paid.md`](mockups/liquid-page-preview-paid.md)

## Entry Point

| Aspect | TO-BE decision | Why |
| --- | --- | --- |
| Верхний entrypoint | Public page templates entrypoint | Оператору нужен прямой discoverable slot для FT-4564 page-template authoring |
| Landing page IA | Один bounded catalog `Liquid pages` | FT-4564 работает с system-owned page keys, а не с ad-hoc records |
| Page title | `Liquid-шаблоны публичных страниц` | Название должно отражать public page templates, а не общий content authoring |
| Scope boundary | FT-4564 UI покрывает только `show`, `paid`, `failed`, `created`, `canceled` | Feature не превращается в redesign всего раздела content authoring |

`Liquid pages` в этом документе является operator UI label для bounded surface `operator/page_templates`, backed by `VendorPageTemplate`. Это не отдельная доменная модель и не storage key.

## Screen Map

| Screen ID | Screen | Purpose |
| --- | --- | --- |
| `OP-01` | Liquid pages catalog | Показать фиксированный catalog FT-4564 page keys, их status и actions |
| `OP-02` | Liquid page editor | Редактирование template body, toggle enable/disable, preview entrypoint, context reference |
| `OP-03` | Preview surface | Визуальный preview выбранного page key через direct iframe выбранной customer-facing public render surface |

## OP-01 Liquid Pages Catalog

Catalog reuse-ит текущий operator layout и показывает пять system-owned rows:

- `show`
- `paid`
- `failed`
- `created`
- `canceled`

Оператор не создаёт и не удаляет эти rows вручную. Он работает только с already known page keys, принадлежащими feature.

### Catalog Layout

| Column | Meaning |
| --- | --- |
| `Page` | Human title и stable page key |
| `Scenario` | Краткое объяснение, где пользователь видит эту страницу |
| `Configured mode` | `Baseline` или `Liquid enabled` |
| `Template status` | `Draft saved`, `template body missing`, `Never edited` |
| `Effective status` | Honest FT-4564 render verdict. Single-surface keys показывают один badge; multi-surface keys показывают verdict per covered surface (`Liquid`, `Baseline`, `template body missing`) |
| `Actions` | `Enable/Disable`, `Edit`, `Preview` (`Preview` deep-link-ит в `OP-02` с активным tab-ом `Предпросмотр`) |

### Catalog States

| State | What operator sees | Why it matters |
| --- | --- | --- |
| `Baseline` | Liquid выключен, публичная страница остаётся на текущем baseline behavior | Показывает default path |
| `Liquid enabled` | Page key включён и имеет template body | Нормальный active authoring state |
| `template body missing` | Page key включён, но template body отсутствует, поэтому публично действует baseline fallback | Делает явной policy из ADR-0006 |

### Multi-Surface Status Policy

Каталог не должен схлопывать multi-surface page families в один synthetic status.

| Page key | Required breakdown |
| --- | --- |
| `show` | Один badge / verdict, потому что covered surface одна |
| `created` | Один badge / verdict, потому что covered surface одна |
| `canceled` | Один badge / verdict, потому что covered surface одна |
| `failed` | Один badge / verdict, потому что covered surface одна |
| `paid` | Две отдельные строки verdict внутри одной row/editor: `orders/paid` и `payment/show(state=success)` |

Если для `paid` одна surface находится в `template body missing`, а другая остаётся Liquid/Baseline, UI показывает оба verdict отдельно и не подменяет их одним словом вроде `Mixed`.

### Empty State Policy

Catalog не имеет empty state в смысле “список пуст”. Catalog фиксирован и всегда показывает все пять допустимых keys. Empty state допустим только внутри отдельных полей:

- “template body ещё не создан”;
- “preview context ещё не выбран”;
- “для выбранного surface нужен explicit context”.

## OP-02 Liquid Page Editor

Editor открывается по клику из `Liquid pages` catalog. Это dedicated authoring surface для одного page key.

### Header

Header editor-а должен показывать:

- human title страницы;
- stable page key;
- краткий сценарий (`customer sees this after payment failure`, `public order page`, etc.);
- current configured/effective state;
- per-surface verdict breakdown для multi-surface keys вроде `paid`.

### Editor Tabs

Внутри editor используются три tabs:

1. `Шаблон`
2. `Предпросмотр`
3. `Переменные и хелперы`

Такой layout reuse-ит привычный operator pattern с tabs и не вынуждает оператора уходить в другой раздел ради preview или context reference.

### Tab `Шаблон`

| Region | Content |
| --- | --- |
| `Status bar` | Toggle `Enable Liquid`, badges configured/effective state, per-surface verdict breakdown для multi-surface keys |
| `Template body` | Основной textarea/code-editor для Liquid template body |
| `Warnings` | Callout для `template body missing`, runtime assumptions по `order` |
| `Actions` | `Save` как primary action; `Preview` и `Enable/Disable` как additional buttons через `bottom_form_panel` |

### Tab `Предпросмотр`

Preview tab не заменяет public runtime отдельной operator-оболочкой. Он даёт оператору controlled preview выбранной customer-facing public render surface.

Preview policy:

- action `Preview` из catalog открывает именно этот tab того же editor-а, а не отдельный detached screen;
- preview iframe открывает direct public URL выбранной render surface, а не `/operator/shop`;
- operator sidebar, publish controls и device-mode shell не должны попадать внутрь preview iframe;
- unsaved changes не preview-ятся автоматически: preview работает по последнему сохранённому template body;
- если editor dirty, UI явно просит сначала сохранить изменения.

### Tab `Переменные и хелперы`

Context reference показывается внутри editor, а не в отдельной wiki-ссылке.

Для каждого page key оператор видит:

- variable/helper name;
- availability;
- короткое описание;
- минимальный пример использования;
- пометку `required`, `optional` или `unavailable` для `order`.

## OP-03 Preview Surface

Preview остаётся authoring helper, но показывает именно customer-facing public render surface. Это не отдельный rendering subsystem и не operator shop shell.

### Preview URL Rules

Preview URL contract владеется `spec.md`, секция `Public Preview URL Contract`. UI обязан следовать этим правилам:

- iframe `src` строится от HTTPS/subdomain base `current_vendor.iframe_preview_url`;
- path заменяется на выбранную public render surface;
- query строится заново из preview contract, а не через строковую конкатенацию;
- iframe `src` не может быть `/operator/shop` или другой operator route.
- payment-result preview не может строиться через `/payments/<provider>/success|failure`; он использует provider-neutral preview-only route для `payment/show(state=success|failure)`.

### Preview Context Rules

| Page key family | Preview contract |
| --- | --- |
| `show`, `created`, `canceled` | Preview требует явный order-aware context |
| `paid` | Preview должен уметь показать как `vendor/orders/paid`, так и `vendor/payment/show(state=success)`; оператор сначала выбирает variant surface, затем честный context для неё |
| `failed` | Preview должен уметь показать outcome-only failure surface без обязательного `order` |

Для `created` выбранный order context используется только для preview-only forced render of post-create surface. Это не означает, что `created` становится regular public order-state branch.

### Preview Path Mapping

| Page key | Surface | Public path | Required query |
| --- | --- | --- | --- |
| `show` | `orders/show` | `/orders/:external_id` | `no_canonical_redirect=true`, `page_template_preview=true`, `preview=:preview_code`, `view=show` |
| `created` | `orders/created` preview-only forced render | `/orders/:external_id` | `no_canonical_redirect=true`, `page_template_preview=true`, `preview=:preview_code`, `view=created` |
| `canceled` | `orders/canceled` | `/orders/:external_id` | `no_canonical_redirect=true`, `page_template_preview=true`, `preview=:preview_code`, `view=canceled` |
| `paid` | `orders/paid` | `/orders/:external_id` | `no_canonical_redirect=true`, `page_template_preview=true`, `preview=:preview_code`, `view=paid` |
| `paid` | `payment/show(state=success)` | `/payments/page_template_preview/success` | `no_canonical_redirect=true`, `page_template_preview=true`, `preview=:preview_code`, optional `order_id` |
| `failed` | `payment/show(state=failure)` | `/payments/page_template_preview/failure` | `no_canonical_redirect=true`, `page_template_preview=true`, `preview=:preview_code`, optional `order_id` |

`/payments/page_template_preview/:state` является preview-only route внутри public shop host. Она нужна только для честного operator preview общей `vendor/payment/show` surface и не заменяет реальные payment callbacks / return URLs.

### Preview Context Selection

Для order-aware pages preview не должен auto-pick-ать случайный live order.

TO-BE expectation:

- оператор явно выбирает preview context;
- для `paid` оператор явно выбирает и variant surface (`orders/paid` или `payment/show(state=success)`), и затем соответствующий context;
- если context не выбран, preview tab показывает пустое состояние с понятным CTA;
- preview не маскирует отсутствие `order` на surfaces, где он только optional.

## Context Reference Design

### Grouping

Context reference группируется минимум по четырём блокам:

1. `Common`
2. `Order`
3. `Outcome / state metadata`
4. `Helpers`

### Availability Labels

| Label | Meaning |
| --- | --- |
| `required` | Template вправе опираться на значение всегда |
| `optional` | Значение может отсутствовать, template обязан это учитывать |
| `unavailable` | Значение не должно ожидаться на этой page family |

### Required Visibility Rules

- Для `show`, `created`, `canceled` `order` помечается как `required`.
- Для `paid` и `failed` `order` не маркируется как universally required.
- Для outcome-only payment result preview UI должен явно показывать, что honest baseline — `vendor + outcome metadata`.

## Warnings And Explanations

Operator UI должен делать resolution semantics явными, а не прятать их.

### Required Callouts

| Trigger | Message intent |
| --- | --- |
| `Liquid enabled + template body missing` | Page key включён, но public verdict остаётся baseline fallback |
| `Multi-surface key with divergent verdicts` | UI показывает separate surface verdicts и не скрывает divergence за одним aggregate status |
| `Preview context missing` | Preview недоступен, пока оператор не выбрал корректный context |
| `Payment result key with order optional` | Template не должен предполагать, что `order` всегда есть |

## Component Mapping To Design Guide

| Need | Recommended operator pattern |
| --- | --- |
| Landing/catalog container | `ibox` / standard operator table layout |
| Tabbed screen | Bootstrap tabs, как в существующих operator forms |
| Toggle enable/disable | `toggle_button` / `smart_toggle_button` pattern |
| Save + Preview + Enable actions | `bottom_form_panel` с `additional_buttons` |
| Inline warnings | `operator_form_warnings` / callout blocks |
| Preview iframe | Direct public render URL from `spec.md` preview contract |

## Traceability

| Screen / element | Supports |
| --- | --- |
| `OP-01` Liquid pages catalog | `REQ-12`, `REQ-13`, `EC-12`, `EC-13`, `EC-16`, `SC-11`, `SC-12`, `SC-15`, `CHK-14`, `CHK-18` |
| `OP-02` editor | `REQ-14`, `REQ-15`, `EC-14`, `EC-15`, `SC-13`, `SC-14` |
| `OP-03` preview | `REQ-14`, `EC-14`, `SC-13` |

## Out Of Scope For This Doc

- детальные persistence-решения;
- file-level touchpoints и sequence implementation;
- canonical acceptance inventory;
- WYSIWYG builder, visual drag-and-drop layouting;
- redesign других content authoring flows вне FT-4564.

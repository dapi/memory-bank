---
title: "ADR-0005: FT-4564 Order Page Resolution and Fallback Policy"
doc_kind: adr
doc_function: canonical
purpose: "Фиксирует принятое архитектурное решение для FT-4564: canonical resolution unit, fallback policy и границы предварительного bounded refactor."
derived_from:
  - ../../features/FT-4564/feature.md
  - ./0003-per-action-engine-override.md
  - ./0006-missing-custom-liquid-template-falls-back-to-existing-baseline-page.md
status: active
decision_status: accepted
date: 2026-04-23
audience: humans_and_agents
must_not_define:
  - current_system_state
  - implementation_plan
---

# ADR-0005: FT-4564 Order Page Resolution and Fallback Policy

## Контекст

FT-4564 вводит Liquid-кастомизацию для customer-facing страниц заказа и результатов оплаты: `failed`, `paid`, `show`, `created`, `canceled`.

Само бизнес-требование понятно, но архитектурно остаются неразрешёнными три вопроса:

1. Какая сущность является canonical unit конфигурации для вендора: page key, concrete render surface или иной слой.
2. Как должна вести себя система при отсутствии Liquid template body и при runtime render error.
3. Нужен ли перед реализацией предварительный bounded refactor, чтобы не размазать resolution/fallback логику по контроллерам и `render_view`.

Дополнительное напряжение создаёт уже принятый ADR-0003: explicit `engine:` override описан как fail-fast contract. FT-4564 не должна неявно сломать этот baseline, но при этом заказчик ожидает безопасный partial rollout.

## Драйверы решения

- Не требовать полного global switch между React и Liquid ради одной или двух страниц.
- Не делать конфигурацию вендора зависимой от внутренних controller branch-ей сильнее, чем это необходимо.
- Не скрывать реальные баги Liquid-render полностью бесшумным fallback-ом.
- Сохранить совместимость с accepted ADR-0003.
- Удержать изменение в формате bounded refactor, а не перепроектирования всего публичного рендеринга.
- Использовать dedicated FT-4564 persistence (`vendor_page_templates`) для public page templates.

## Decision Criteria

Предлагаемое решение считается admissible только если одновременно выполняются все критерии:

1. **Vendor-facing stability.** Конфигурация вендора привязана к стабильным semantic page keys, а не к случайным controller branch-ам.
2. **Category hygiene.** Отсутствие Liquid template body и runtime render error не смешиваются в один generic fallback case.
3. **Partial-rollout safety.** Отсутствие Liquid template body не превращает частичный rollout в user-facing outage.
4. **Debuggability.** Runtime bug в Liquid не маскируется автоматическим возвратом на React.
5. **Baseline compatibility.** Решение не ломает accepted ADR-0003 и не подменяет explicit `engine:` contract неявной магией.
6. **Bounded change surface.** Refactor ограничен resolution/policy layer и не перерастает в redesign всего rendering pipeline.

## Non-Goals

- Не вводить general-purpose per-page selector между React и Liquid.
- Не менять global engine semantics.
- Не брать в scope provider-hosted payment pages, которые рендерятся вне `render_view`.
- Не переписывать `render_view` в новый orchestration framework.
- Не добавлять новый platform-level analytics API для payment result pages.

## Рассмотренные варианты

| Вариант | Плюсы | Минусы | Почему рассматривается как основной кандидат / не основной кандидат |
| --- | --- | --- | --- |
| `Option A: Patch current surfaces directly` | Минимум upfront-работы, самый короткий путь до первой реализации | Конфигурация протекает в controller-specific ветки, fallback policy быстро расползается по нескольким точкам, runtime errors легко начать маскировать | Не основной кандидат: слишком высок риск локальных хаков и неявной semantics |
| `Option B: Bounded refactor + stable page keys + staged fallback policy` | Даёт один canonical vocabulary для вендора, отделяет resolution policy от низкоуровневого renderer, позволяет fallback только на отсутствие template body и сохраняет fail-fast для runtime bugs | Требует upfront-рефакторинга и явного mapping слоя | Выбранный вариант: наилучший баланс между безопасным rollout и инженерной прозрачностью |
| `Option C: Strict fail-fast for any opted-in page problem` | Самый простой контракт, максимально близок к ADR-0003 | Слишком жёсткий для partial rollout: отсутствие template body превращается в продуктовую поломку вместо контролируемого деградационного сценария | Не основной кандидат: operational cost слишком высок для этой миграции |

## Решение

Для `decision_status: accepted` выбран `Option B`.

Предлагаемое решение:

1. Canonical config unit для FT-4564 — **semantic page key**, а не concrete render surface.
2. Предлагаемые internal semantic page keys:
   - `payment_failure` (user-facing alias: `failed`)
   - `payment_success` (user-facing alias: `paid`)
   - `order_show` (alias: `show`)
   - `order_created` (alias: `created`)
   - `order_canceled` (alias: `canceled`)
3. Vendor-facing docs и UI могут продолжать использовать короткие алиасы `failed/paid/show/created/canceled`, но storage, mapping и code-level policy должны опираться на semantic keys.
4. Отдельный resolution layer сопоставляет текущую concrete render surface с semantic page key и определяет, применяется ли page-specific Liquid поведение.
5. Для in-scope public page surfaces page-specific Liquid opt-in выбирает render policy: если vendor явно включил page-specific Liquid и сохранил body, рендерится Liquid template; если включил Liquid без body, применяется missing-template fallback на existing baseline page implementation; если Liquid не выбран, используется текущий baseline behavior.
6. Fallback policy разделяется на два разных класса:
   - **Liquid selected, template body exists** → используется Liquid.
   - **Liquid selected, template body absent** → используется existing baseline page implementation для того же page key / surface. Это platform-level rule из ADR-0006.
   - **Runtime render error** после выбора Liquid render → fallback назад в React не делается; действует fail-fast + error reporting.
7. Контракт `order` различается по semantic page key:
   - `order_show`, `order_created`, `order_canceled` требуют `order` в Liquid context.
   - `payment_success`, `payment_failure` публикуют `order` как optional (`order?`), потому что текущие customer-facing success/failure surfaces не в каждом flow гарантируют надёжный order context.
8. Для payment result pages минимально гарантированный context — `vendor + outcome metadata`; rich order details допустимы только там, где `order` реально доступен.
9. В текущем FT-4564 `order_created` трактуется как post-create surface, а не как regular branch публичной order page.
10. Предварительный refactor ограничивается выделением resolution/fallback policy из размазанных controller-веток. `render_view` остаётся низкоуровневым renderer и не становится owner бизнес-логики FT-4564.
11. Bounded refactor explicitly allows:
   - единый semantic mapping layer `surface -> page key`;
   - единый policy layer для resolution / fallback / observability;
   - локальные изменения в controller entry points, необходимые для подключения этого слоя.
12. Bounded refactor explicitly forbids:
   - redesign всего public rendering pipeline;
   - переписывание `render_view` как общего orchestration framework;
   - расширение FT-4564 до global engine migration;
   - расширение FT-4564 до general-purpose per-page React selector.

FT-4564 использует этот контракт как accepted baseline.

## Последствия

### Положительные

- У вендора появляется одна стабильная vocabulary конфигурации вместо привязки к внутренним веткам контроллеров.
- Отсутствие Liquid template body не превращает partial rollout в продовую поломку.
- Runtime bugs не маскируются бесшумно.
- Граница между problem scope и solution scope становится чище: `feature.md` владеет требованием, ADR владеет архитектурной policy.
- `order` contract становится честным: rich order-aware pages и payment result pages больше не притворяются одинаковыми по контексту.
- Rule про отсутствие Liquid template body поднят на platform-level и может переиспользоваться будущими feature через ADR-0006.

### Отрицательные

- Появляется дополнительный слой resolution logic, который нужно отдельно тестировать.
- Реализация стартует медленнее из-за upfront-рефакторинга.
- Придётся явно договориться, какие concrete surfaces попадают под каждый semantic page key.
- На payment result pages шаблоны не смогут безусловно рассчитывать на наличие `order`.
- Нужно поддерживать явный baseline page implementation для page keys, где разрешён optional Liquid.

## Риски и mitigation

- Риск: semantic page key mapping окажется слишком искусственным и будет расходиться с реальным flow.
  Mitigation: проверить mapping на текущих customer-facing branches до старта реализации и зафиксировать в `spec.md`, а затем сослаться на него из implementation plan.

- Риск: fallback на React при отсутствии Liquid template body начнёт использоваться как костыль вместо нормальной поставки шаблонов.
  Mitigation: structured logging / visibility для случая `Liquid selected + template body absent` и явные verify checks.

- Риск: bounded refactor разрастётся в переписывание рендеринга.
  Mitigation: ограничить ownership refactor-а только resolution/fallback layer; не переписывать `render_view` как общий frontend framework.

- Риск: optional `order` на `payment_success` / `payment_failure` окажется недостаточным для ожидаемого merchant UX.
  Mitigation: в `spec.md` явно перечислить, какие success/failure surfaces реально попадают под FT-4564, и где merchant может рассчитывать на rich order details, а где только на outcome-level customization.

- Риск: fallback на React при отсутствии template body начнёт восприниматься как “fallback на любую Liquid-проблему”.
  Mitigation: в `spec.md` и verify явно разделять `template body absent` и `runtime render error`; метрики и логи публиковать раздельно.

- Риск: public page templates начнут моделироваться как ad-hoc records вместо fixed semantic catalog.
  Mitigation: docs, operator labels и tests должны использовать термин `VendorPageTemplate` / public page template и проверять fixed semantic-key catalog.

## Acceptance Basis

Решение принято на основании зафиксированных фактов:

1. Inventory текущих customer-facing surfaces собран в feature-local [spec.md](../../features/FT-4564/spec.md).
2. Код подтверждает, что `payment_success` / `payment_failure` не гарантируют universally available `order`, тогда как order-state surfaces его несут.
3. Vendor controllers подключают `VendorBugsnag`, а payment callbacks уже имеют logging / error-reporting path.
4. Provider-hosted payment pages остаются вне scope `render_view` и FT-4564.

## Assurance Snapshot

Текущий уровень уверенности в принятом решении:

- **Formality:** medium-high. Ключевые classes решения и границы scope уже сформулированы явно.
- **Grounding:** high. Базовые customer-facing flows и их surface inventory зафиксированы в feature-local solution docs.
- **Reliability:** medium-high. Ключевые policy decisions закреплены accepted ADR-0006 и code-level evidence по текущим routes/controllers.

## Связанные ссылки

- [FT-4564 feature](../../features/FT-4564/feature.md) — canonical feature scope и verify inventory.
- [FT-4564 spec](../../features/FT-4564/spec.md) — canonical feature-local solution spec: current surface inventory, contracts и implementation verification notes.
- [FT-4564 use cases](../../features/FT-4564/use-cases/README.md) — пользовательские сценарии без архитектурной детализации.
- [ADR-0003](0003-per-action-engine-override.md) — accepted baseline для explicit `engine:` override.
- [ADR-0006](0006-missing-custom-liquid-template-falls-back-to-existing-baseline-page.md) — platform-level rule для optional custom Liquid templates.
- GitHub Issue #4564 — исходное бизнес-требование.

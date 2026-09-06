---
title: Testing Policy
doc_kind: governance
doc_function: canonical
purpose: "Описывает testing policy delivery-процесса: обязательность test case design, требования к automated regression coverage, допустимые manual-only gaps и сходимость проверки после реализации."
derived_from:
  - ../dna/governance.md
  - autonomy-boundaries.md
  - behavior-specification.md
  - feature.md
  - feature-requirements.md
  - routing.md
  - validation-profiles.md
status: active
canonical_for:
  - repository_testing_policy
  - feature_test_case_inventory_rules
  - automated_test_requirements
  - sufficient_test_coverage_definition
  - manual_only_verification_exceptions
  - simplify_review_discipline
  - verification_context_separation
  - review_convergence_contract
  - bdd_automation_policy
must_not_define:
  - feature_acceptance_criteria
  - feature_scope
  - project_testing_stack
audience: humans_and_agents
---

# Testing Policy

Этот документ generic: он задаёт, что и когда обязано быть проверено, независимо от стека проекта. Project-specific framework, тестовые данные, CI jobs и размещение тестов живут в [`../engineering/testing-conventions.md`](../engineering/testing-conventions.md), canonical локальные команды — в [`../ops/development.md`](../ops/development.md); здесь они не дублируются.

## Core Rules

- Выбранный [`validation profile`](validation-profiles.md) задаёт minimum validation/evidence floor независимо от delivery flow; project-specific policy может только усиливать его.
- Любое изменение поведения, которое можно проверить детерминированно, обязано получить automated regression coverage.
- Любой новый или измененный contract обязан получить contract-level automated verification.
- Любой bugfix обязан добавить regression test на воспроизводимый сценарий.
- Required automated tests считаются закрывающими риск только если они проходят локально и в CI.
- Manual-only verify допустим только как явное исключение и не заменяет automated coverage там, где automation реалистична.

## BDD Automation Policy

[`Behavior Specification Practice`](behavior-specification.md)
определяет Discovery и Formulation; этот policy определяет automation boundary.

- BDD не требует Gherkin, Cucumber, browser automation или E2E-only tests.
- Для каждого required `SC-*` / `NEG-*` выбирай самый низкий надёжный unit,
  component, contract, integration или E2E surface, который доказывает
  observable outcome.
- Один behavior example может проверяться несколькими техническими тестами;
  каждый test surface должен быть виден в `implementation-plan.md#test-strategy`.
- Имя, tag или metadata теста должны сохранять ссылку на `SC-*` / `NEG-*`, если
  project framework это допускает без brittle coupling.
- Test code не владеет requirement или expected behavior. При изменении
  expected verdict сначала обнови canonical `UC/domain/brief` owner, затем
  examples, `CHK-*`, plan и code.
- Gherkin, если выбран downstream-проектом, является executable projection
  canonical `SC-*` / `NEG-*`, а не параллельным source of truth.

## Ownership Split

- Canonical validation profile decision живёт только в owner-е, назначенном [`validation-profiles.md`](validation-profiles.md); testing policy и execution artifacts не выбирают profile повторно.
- Canonical test cases delivery-единицы задаются в `brief.md` через `SC-*`, feature-specific `NEG-*`, `CHK-*` и `EVID-*`.
- Design pack, если нужен, aggregate-владеет selected design, C4 applicability/model, `CTR-*`, `INV-*`, `FM-*` и локальными `RB-*`. Root `design.md` остаётся default owner, а явно делегированный constituent — непосредственным owner перечисленных IDs; ни один design artifact не подменяет canonical verify contract.
- `implementation-plan.md` владеет только стратегией исполнения: какие test surfaces будут добавлены или обновлены, какие gaps временно остаются manual-only и почему.

## Feature Flow Expectations

Canonical lifecycle gates живут в [feature.md](feature.md):

- к `Problem Ready` `brief.md` уже фиксирует validation profile decision и test case inventory;
- к `Solution Ready` весь required design pack готов по relation, ownership, publication/lifecycle и consistency rules из Feature Flow;
- к `Plan Ready` `implementation-plan.md` содержит `Test Strategy` с planned automated coverage и manual-only gaps;
- к `Done` required tests добавлены, локальные команды зелёные, CI не противоречит локальному verify, а проверка сошлась по [`Review Convergence`](#review-convergence).

## Что Считается Sufficient Coverage

For each applicable `REQ-*`, preserve the `brief.md` verification method and evidence contract. Quality requirements include an observable threshold and a repeatable measurement/check. The plan maps each changed implementation, test, and config path plus symbol/section back to a `REQ-*` or explicit supporting rationale; review checks that mapping in both directions.

- Покрыт основной changed behavior и ближайший regression path.
- Покрыты новые или измененные contracts, события, schema или integration boundaries.
- Покрыты критичные failure modes из `FM-*` в required `design.md`, bug history или acceptance risks.
- Покрыты feature-specific negative/edge scenarios, если они меняют verdict.
- Required `SC-*` / `NEG-*` прослеживаются через `CHK-*` к automated test либо
  явно approved manual-only gap и concrete `EVID-*`.
- Процент line coverage сам по себе недостаточен: нужен scenario- и contract-level coverage.

## Когда Manual-Only Допустим

- Сценарий зависит от live infra, внешних систем, hardware, недетерминированной среды или human оценки UI.
- Для каждого manual-only gap: причина, ручная процедура, owner follow-up.
- Для каждого manual-only gap соблюдены approval requirements выбранного validation profile.
- Если manual-only gap оставляет без regression protection критичный путь, feature не считается завершённой.

## Simplify Review

Отдельный проход верификации после функционального тестирования. Цель: убедиться, что реализация минимально сложна.

- Выполняется после прохождения tests, но до closure gate.
- Паттерны: premature abstractions, глубокая вложенность, дублирование логики, dead code, overengineering.
- Три похожие строки лучше premature abstraction. Абстракция оправдана только когда она реально уменьшает риск или повтор.

## Verification Context Separation

Artifact review и implementation review имеют разные объекты проверки и evidence:

1. **Artifact review** — до соответствующего lifecycle gate проверяет governed requirements/design/plan artifacts, их grounding, ownership, completeness и traceability.
2. **Implementation review** — после execution проверяет delivered code и repository diff против принятых canonical artifacts.
3. **Функциональная верификация** — tests проходят, acceptance scenarios покрыты.
4. **Simplify review** — код минимально сложен.
5. **Acceptance test** — end-to-end по `SC-*`.

Artifact review не является доказательством качества реализации, а implementation review не исправляет задним числом непройденный artifact gate. Для compact feature packages проходы допустимы в одной сессии, если их объекты, verdicts и evidence зафиксированы раздельно; обязательный review или simplify review не пропускается.

## Review Convergence

Проверка после реализации доводится ровно до одного из двух состояний.

**Сошлась.** Clean verdict обязательного review и зелёный обязательный CI
относятся к одной repository revision — одному commit SHA. Изменение кода после
verdict создаёт новую revision и аннулирует его. До выполнения этих условий
работа не переходит в терминальное состояние flow.

**Не сошлась.** Исчерпан бюджет в десять циклов на delivery-единицу. Один цикл —
`несошедшаяся проверка → исправления → повторная проверка`; несошедшейся
считается verdict с блокирующими замечаниями либо красный обязательный CI по
причине в коде. Повторный запуск CI без изменения revision циклом не является.

Исчерпание бюджета не разрешает принять работу и само по себе не создаёт Human
Gate. Зафиксируй вне проверяемой revision — в issue, PR или routing record
выбранного flow, чтобы запись не меняла проверенный код:

```text
Not converged: cycle <N>/10, revisions <SHA…>
Repeated findings: <что повторялось; какие проверки остаются красными>
Cause class: <строка таблицы ниже>
Decision: return to <владелец фактов> | continue, scope <суженный объём>
```

Новый бюджет открывается только после обновления владельца исходных фактов. При
дефекте реализации допускается одно продолжение в текущем flow с явно суженным
объёмом; дальше применяется
[`Structured Decision Protocol`](autonomy-boundaries.md#structured-decision-protocol),
и эскалация нужна только при outcome `escalate`.

| Причина | Куда вернуть работу |
| --- | --- |
| Дефект реализации или необоснованная сложность | Ограничить объём исправления и повторить его в рамках текущего flow |
| Недостаточное изучение кода или ошибочная последовательность исполнения | Этап планирования исполнения выбранного flow (в Feature Flow — Plan Ready) |
| Пробел или противоречие в решении, контракте, инварианте, failure handling или release rules | Владелец design pack либо ADR (в Feature Flow — Solution Ready) |
| Пробел или противоречие в границах, требовании, acceptance criterion или составе evidence | Канонический problem-owner: `brief.md`, bug report, refactoring task (в Feature Flow — Problem Ready) |
| Недостаточно определены проблема, результат или границы | Повторить брифование или первичный разбор |
| Неверно определён тип задачи или её объём | Повторить [`Task Routing`](routing.md); при необходимости выбрать Research, Feature или Epic Flow |
| Внешняя блокировка, недоступная инфраструктура или нестабильный CI без причины в коде | Зафиксировать ожидание или блокировку с evidence, не меняя исходные факты без основания |

Правая колонка называет владельца фактов; gates в скобках относятся к Feature
Flow. Flow без такого этапа возвращает работу к соответствующему месту своего
lifecycle, а если этапа нет вовсе — повторяет [`Task Routing`](routing.md).
После изменения владельца зависящие документы и проверки проходят заново.

Контракт применяется к Feature, Small Change, Bug Fix и Refactoring Flow, не
относится к artifact review до gate (включая Plan Ready) и механизм-нейтрален.
Обязательств выбранного [`validation profile`](validation-profiles.md) он не
изменяет. Incident Flow закрывает containment по своим правилам; permanent
remediation попадает под контракт после отдельного [`Task Routing`](routing.md).

## Project Execution Layer

Как именно исполняется эта policy в конкретном репозитории — framework, тестовые данные, CI jobs, размещение тестов и helper patterns — задаёт [`../engineering/testing-conventions.md`](../engineering/testing-conventions.md), а canonical локальные команды — [`../ops/development.md`](../ops/development.md). Project-специфичный слой может усиливать требования этой policy, но не может их ослаблять.

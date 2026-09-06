---
title: Testing Policy
doc_kind: governance
doc_function: canonical
purpose: "Описывает testing policy delivery-процесса: обязательность test case design, требования к automated regression coverage и допустимые manual-only gaps."
derived_from:
  - ../dna/governance.md
  - behavior-specification.md
  - feature.md
  - feature-requirements.md
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

Контракт описывает проверку **после реализации** и доводит её ровно до одного из
двух состояний. Третьего — «вроде замечания закрыли» — не существует.

1. **Сошлась.** Review дал clean verdict, блокирующие замечания устранены,
   обязательный CI зелёный.
2. **Не сошлась.** Бюджет циклов исчерпан. Работа не закрывается, причина
   классифицируется, работа возвращается владельцу исходных фактов.

Контракт не относится к artifact review до lifecycle gate — в частности к
Plan Ready, у которого свой бюджет и свои объекты проверки; их разделение
описано в [`Verification Context Separation`](#verification-context-separation).
Контракт механизм-нейтрален: он не выбирает инструмент проверки, команду или
оркестратор.

Глубина самой проверки принадлежит выбранному
[`validation profile`](validation-profiles.md): где профиль требует лишь обычный
review, контракт не добавляет отдельного неавторского reviewer, а где требует
separate non-authoring review — не отменяет его.

### Одна редакция

Clean verdict и зелёный обязательный CI обязаны относиться к **одной редакции
кода** — одному commit SHA. Verdict по одной редакции и CI по другой вместе не
доказывают ничего.

Любое изменение кода после verdict создаёт новую редакцию и аннулирует его:
последний цикл с исправлениями не является clean verdict без повторной проверки.
До выполнения этих условий работа не переходит в Done.

### Бюджет и граница цикла

Один цикл — это `review → исправления → повторный review`. Он начинается
verdict-ом с блокирующими замечаниями и заканчивается следующим verdict-ом,
сколько бы коммитов ни содержали исправления. Чистый re-review без изменений
циклом не является.

По умолчанию на delivery-единицу допускается не более **десяти** циклов. Бюджет
общий, а не отдельный для каждого прохода проверки.

Повторный запуск нестабильной или недоступной проверки CI **без изменения кода**
циклом не считается: редакция та же, исправления не было. Такой повтор
фиксируется отдельно как внешняя причина.

Исчерпание бюджета не разрешает принять работу, проигнорировать замечания или
автоматически потребовать решение человека. Нельзя и начать новый лимит: пока
причина не разобрана, а владелец исходных фактов не обновлён, счёт продолжается.

### Разбор при несходимости

После исчерпания бюджета зафиксируй на наблюдаемых данных: какие замечания
повторялись, какие исправления предпринимались и к чему привели, какие локальные
проверки и CI остаются красными, локальна ли причина или лежит на более раннем
этапе.

Затем отнеси причину к одному классу и верни работу его владельцу:

| Причина | Куда вернуть работу |
| --- | --- |
| Дефект реализации или необоснованная сложность | Пересоставить ограниченный план исправления реализации |
| Недостаточное изучение кода или ошибочная последовательность исполнения | Этап планирования исполнения выбранного flow (в Feature Flow — Plan Ready) |
| Пробел или противоречие в решении, контракте, инварианте, failure handling или release rules | Владелец design pack либо ADR (в Feature Flow — Solution Ready) |
| Пробел или противоречие в границах, требовании, acceptance criterion или составе evidence | Канонический problem-owner: `brief.md`, bug report, refactoring task (в Feature Flow — Problem Ready) |
| Недостаточно определены проблема, результат или границы | Повторить брифование или первичный разбор |
| Неверно определён тип задачи или её объём | Повторить [`Task Routing`](routing.md); при необходимости выбрать Research, Feature или Epic Flow |
| Внешняя блокировка, недоступная инфраструктура или нестабильный CI без причины в коде | Зафиксировать ожидание или блокировку с evidence, не меняя исходные факты без основания |

Flow без именованных gates возвращает работу к соответствующему этапу своего
lifecycle: правая колонка называет владельца фактов, а не gate конкретного flow.

После изменения исходного владельца все зависящие от него документы и проверки
проходят заново.

### След

В canonical carrier выбранного flow — том же, который владеет routing record и
validation profile decision, — фиксируются: номер цикла, проверенная редакция,
verdict, класс причины и решение продолжить, остановиться или вернуться назад.
Без этого следа нельзя отличить сошедшуюся проверку от брошенной.

Исчерпание бюджета само по себе не является Human Gate. Примени
[`Structured Decision Protocol`](autonomy-boundaries.md#structured-decision-protocol)
и записывай решение в его канонической форме; эскалация нужна только при outcome
`escalate` либо при уже действующем обязательном approval.

## Project Execution Layer

Как именно исполняется эта policy в конкретном репозитории — framework, тестовые данные, CI jobs, размещение тестов и helper patterns — задаёт [`../engineering/testing-conventions.md`](../engineering/testing-conventions.md), а canonical локальные команды — [`../ops/development.md`](../ops/development.md). Project-специфичный слой может усиливать требования этой policy, но не может их ослаблять.

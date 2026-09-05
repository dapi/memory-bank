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
  - bdd_automation_policy
must_not_define:
  - feature_acceptance_criteria
  - feature_scope
  - project_testing_stack
audience: humans_and_agents
---

# Testing Policy

Этот документ generic: он задаёт, что и когда обязано быть проверено, независимо от стека проекта. Project-specific framework, тестовые данные, local commands и CI jobs живут в [`../engineering/testing-conventions.md`](../engineering/testing-conventions.md) и не дублируются здесь.

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
- к `Done` required tests добавлены, локальные команды зелёные и CI не противоречит локальному verify.

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

## Project Execution Layer

Как именно исполняется эта policy в конкретном репозитории — framework, тестовые данные, local commands, CI jobs, размещение тестов и helper patterns — задаёт [`../engineering/testing-conventions.md`](../engineering/testing-conventions.md). Тот документ может усиливать требования этой policy, но не может их ослаблять.

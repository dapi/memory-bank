---
title: Small Change Flow
doc_kind: governance
doc_function: canonical
purpose: Прямой delivery flow для задач, где issue достаточен, а отдельные design и execution plan не нужны.
derived_from:
  - ../dna/governance.md
  - routing.md
  - ../engineering/testing-policy.md
  - ../engineering/validation-profiles.md
canonical_for:
  - small_change_entry_contract
  - small_change_routing_record
  - small_change_execution_flow
  - small_change_evidence_rules
  - small_change_escalation_rules
  - small_change_outcome_contract
status: active
audience: humans_and_agents
---

# Small Change Flow

`Small Change` — fast path, выбранный по predicates из [`routing.md`](routing.md). Для него не создаются feature package, `brief.md`, `design.md`, `implementation-plan.md` или ADR; issue/task остаётся owner-ом intent, scope и acceptance.

## Entry Gate

- [ ] Task Routing выбрал `Small Change`
- [ ] issue/task содержит intent, scope, acceptance и verify contract
- [ ] указан конкретный существующий reference pattern
- [ ] design и execution plan не требуются по правилам ниже
- [ ] routing record зафиксирован до реализации

## Routing Record

Зафиксируй record в issue/task. Если task tracker нельзя обновить, добавь его в draft PR при первой возможности.

```text
Workflow: Small Change

Design: not required
Reason: решение следует существующему паттерну <ссылка или путь>.

Plan: not required
Reason: change surface локален, порядок шагов и checkpoints не нужны.

Verify:
- <команда или проверка>
- <ожидаемый результат или evidence>

Validation profile: documentation | low-risk | standard
Triggers / rationale: <почему выбранный minimum достаточен>
Downgrade approval: none
```

`high-risk` и `release-deployment` несовместимы с Small Change predicates: остановись и повтори Task Routing. `standard` не требует feature package, если issue всё ещё полностью задаёт решение, change surface локален и остальные Small Change predicates истинны.

`Design: not required` допустим, только если не требуется выбирать между альтернативами и не появляются новые contracts, invariants, security boundaries, migrations, rollout rules или failure modes.

`Plan: not required` допустим, только если change surface и test surfaces известны, изменение доставляется одним атомарным change set и не требует зависимых этапов, migration/backout sequencing или промежуточных checkpoints.

## Flow

```text
issue/task → routing record → implementation → automated checks
           → simplify review → PR → review + CI → merge → handoff
```

## Execution Gates

- [ ] реализация не выходит за declared scope
- [ ] changed behavior получает required automated coverage
- [ ] проверки из routing record выполнены
- [ ] simplify review выполнен отдельным проходом
- [ ] PR ссылается на issue/task и содержит concrete evidence
- [ ] required CI зелёный до merge

## Delivery Trace

Memory Bank package не создаётся. Проверяемый след образуют issue/task, routing record, commit history, tests, PR и CI results.

Если изменение исправляет существующий canonical fact, обнови его owner в Memory Bank. Если появляется новый устойчивый project fact или design decision, останови `Small Change` и повтори Task Routing.

## Outcome / Exit Contract

### Observable Outcome

Acceptance из issue/task выполнен одним локальным change set без design- и plan-документов.

### Required Evidence

- Small Change routing record;
- validation profile decision и evidence его minimum contract;
- изменённый код и automated coverage для changed behavior;
- результаты проверок из `Verify`;
- последний review cycle завершён без открытых замечаний;
- все изменения закоммичены и отправлены в remote branch, required CI полностью зелёный.

### Terminal State

`Done`: все Execution Gates выполнены, change принят по git workflow проекта и delivery trace доступен из issue/task или PR.

### Handoff

Закрой issue/task; обнови canonical owner исправленного факта. Любую новую устойчивую информацию, design decision или оставшуюся работу сначала верни в Task Routing.

## Escalation

- Нужна reproduction и regression protection для дефекта → [`Bug Fix Flow`](bug-fix.md).
- Понадобились design, plan, contract change, rollout или approvals → повторный [`Task Routing`](routing.md).
- Изменение превратилось в behavior-preserving restructuring с большим change surface → [`Refactoring Flow`](refactoring.md).

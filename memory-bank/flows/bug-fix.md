---
title: Bug Fix Flow
doc_kind: governance
doc_function: canonical
purpose: Delivery flow для воспроизводимого расхождения между ожидаемым и наблюдаемым поведением.
derived_from:
  - ../dna/governance.md
  - routing.md
  - ../engineering/testing-policy.md
  - ../engineering/validation-profiles.md
canonical_for:
  - bug_fix_entry_contract
  - bug_reproduction_rules
  - bug_fix_execution_flow
  - bug_regression_evidence_rules
  - bug_fix_closure_rules
  - bug_fix_outcome_contract
status: active
audience: humans_and_agents
---

# Bug Fix Flow

Bug — наблюдаемое поведение, противоречащее уже принятому expected behavior. Источником может быть error tracker, support, QA, пользовательский report или incident analysis.

## Entry Gate

- [ ] expected и actual behavior различимы
- [ ] указан источник уже принятого expected behavior либо зафиксировано его явное подтверждение человеком
- [ ] report не является запросом на новое поведение
- [ ] operational incident уже contained или передан в [`Incident Flow`](incident.md)
- [ ] bug report или связанная delivery task фиксирует validation profile decision

Если нет ни доступного источника уже принятого expected behavior, ни зафиксированного решения человека, Entry Gate не выполнен: зафиксируй вопрос и риск через [Human Routing](routing.md#human-routing). До решения `Human Gate` не начинай Analysis And Fix и не изменяй код; после решения повтори Task Routing.

## Flow

```text
report → triage → reproduction → analysis → fix
       → regression coverage → review + CI → closure
```

## Reproduction Gate

- Зафиксируй минимальные inputs, environment и steps.
- Сохрани observed result и expected result.
- Предпочитай failing automated test как reproduction carrier.
- Если reproduction невозможна, укажи ограничения, доступные evidence и риск исправления по гипотезе.

## Analysis And Fix

- Отделяй подтверждённую root cause от гипотез.
- Исправляй причину, а не только наблюдаемый симптом.
- Не расширяй scope скрытым product change или unrelated refactoring.
- Если fix требует нового contract, design decision, migration или rollout, остановись и повтори [`Task Routing`](routing.md).

## Regression And Closure Gates

- [ ] исходный сценарий воспроизводился до fix или ограничение явно записано
- [ ] regression test падает без fix и проходит с fix, когда это технически воспроизводимо
- [ ] ближайшие related regression paths проверены
- [ ] simplify review выполнен
- [ ] PR содержит ссылку на report, root cause summary и evidence
- [ ] required local tests и CI зелёные

Если analysis показывает, что observed behavior соответствует текущему contract, а требуется изменить expected behavior, это не bug fix: повтори Task Routing и выбери `Small Change` или Feature Flow.

## Outcome / Exit Contract

### Observable Outcome

Подтверждённое expected behavior восстановлено, а исходный regression защищён от повторного появления.

### Required Evidence

- reproduction с expected/actual behavior или явно записанное ограничение reproduction;
- validation profile decision и evidence его minimum contract;
- подтверждённая root cause summary;
- regression test или обоснованный альтернативный carrier;
- результаты required tests;
- последний review cycle завершён без открытых замечаний;
- все изменения закоммичены и отправлены в remote branch, required CI полностью зелёный.

### Terminal State

`Resolved`: Regression And Closure Gates выполнены, fix принят по git workflow проекта, а report связан с evidence.

### Handoff

Закрой bug report. Product change, contract change, refactoring и другие follow-up задачи не включай скрыто в fix — верни их в Task Routing.

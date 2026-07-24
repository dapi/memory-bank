---
title: Refactoring Flow
doc_kind: governance
doc_function: canonical
purpose: Behavior-preserving flow для локального, исследовательского или системного изменения внутренней структуры.
derived_from:
  - ../dna/governance.md
  - routing.md
  - ../engineering/testing-policy.md
  - ../engineering/validation-profiles.md
canonical_for:
  - refactoring_entry_contract
  - refactoring_classification
  - refactoring_execution_flow
  - behavior_preservation_gates
  - refactoring_escalation_rules
  - refactoring_outcome_contract
status: active
audience: humans_and_agents
---

# Refactoring Flow

Refactoring меняет внутреннюю структуру, сохраняя observable behavior и действующие contracts. Если поведение должно измениться, повтори [`Task Routing`](routing.md).

## Classification

- **Local:** небольшой behavior-preserving change, который может пройти [`Small Change Flow`](small-change.md).
- **Research:** исследование структуры и вариантов; результатом может быть proposal, plan или ADR без production change.
- **Systemic:** большой change surface, несколько компонентов или этапов, обязательные plan и checkpoints.

## Entry Gate

- [ ] цель и non-goals сформулированы
- [ ] observable behavior и contracts, которые нужно сохранить, перечислены
- [ ] baseline tests или characterization checks определены
- [ ] local refactoring не прошёл `Small Change` gate либо сознательно требует отдельного flow
- [ ] для architecture-level decisions существует accepted ADR или запланирован decision gate
- [ ] исходная task фиксирует validation profile decision с учётом blast radius и critical behavior

## Flow

```text
task → baseline → characterization coverage → plan + checkpoints
     → incremental execution → regression verification
     → simplify review → PR + CI → merge
```

## Execution Rules

- Разбивай systematic refactoring на обратимые checkpoints.
- Не смешивай behavior changes с structural changes в одном неразличимом diff.
- Сохраняй green baseline между checkpoints, если это практически возможно.
- Любое намеренное изменение contract или behavior требует повторного routing.
- Удаляй временные compatibility layers и dead code только на предусмотренном checkpoint.

## Closure Gates

- [ ] baseline behavior сохранён
- [ ] required tests и characterization coverage зелёные
- [ ] contracts не изменились либо изменение вынесено в другой governed flow
- [ ] simplify review подтверждает уменьшение или обоснование complexity
- [ ] rollback или остановка на последнем checkpoint понятны
- [ ] PR содержит before/after structure summary и evidence

## Outcome / Exit Contract

### Observable Outcome

Для Local/Systemic refactoring внутренняя структура улучшена при сохранении observable behavior; для Research refactoring создан проверяемый proposal, plan или ADR без скрытого production change.

### Required Evidence

- baseline и characterization coverage;
- validation profile decision и evidence его minimum contract;
- результаты regression checks по checkpoints;
- before/after summary либо research artifact с источниками и выводом;
- последний review cycle завершён без открытых замечаний;
- все изменения закоммичены и отправлены в remote branch, required CI полностью зелёный для production change.

### Terminal State

`Done`: выбранный результат принят, применимые Closure Gates выполнены, а behavior preservation подтверждён evidence.

### Handoff

Закрой исходную задачу. Любое обнаруженное изменение поведения, contract или отдельный structural scope верни в Task Routing.

---
title: Refactoring Flow
doc_kind: governance
doc_function: canonical
purpose: Behavior-preserving flow для локального, исследовательского или системного изменения внутренней структуры.
derived_from:
  - ../dna/governance.md
  - routing.md
  - ../engineering/testing-policy.md
canonical_for:
  - refactoring_entry_contract
  - refactoring_classification
  - refactoring_execution_flow
  - behavior_preservation_gates
  - refactoring_escalation_rules
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

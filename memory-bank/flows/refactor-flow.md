---
title: Refactor Flow
doc_kind: governance
doc_function: canonical
purpose: "Определяет процесс, gates переходов и evidence contract для refactor-small и managed refactor/chore задач."
derived_from:
  - ../dna/governance.md
  - workflows.md
  - task-flow.md
canonical_for:
  - refactor_task_process
  - refactor_invariant_gate
  - refactor_checkpoint_rules
  - refactor_identifier_taxonomy
status: active
audience: humans_and_agents
---

# Refactor Flow

`refactor-flow` применяется после selector-а из [`workflows.md`](workflows.md), когда выбран:

- `workflow_profile=refactor-small`;
- `workflow_profile=managed-task` и primary owner task package - `refactor.md`.

Общие правила carrier/package, worktree, promotion и closure задает [`task-flow.md`](task-flow.md). Этот документ владеет только refactor/chore-процессом: intent, behavior invariants, change surface, checkpoints и verification.

## Комплект Документов

| Profile | Carrier | Когда достаточно |
| --- | --- | --- |
| `refactor-small` | GitHub issue/PR note по [`templates/task/refactor.md`](templates/task/refactor.md) | Observable behavior не меняется, invariants и verification помещаются в issue/PR |
| managed refactor/chore | `memory-bank/tasks/TASK-XXX/README.md` + `refactor.md` | Нужен durable context, checkpoints или wider technical scope, но нет feature capability или contract/design trigger-а |

`refactor.md` не должен содержать new runtime behavior, feature scope, rollout/backout design или implementation plan. Если это нужно, task повышается по [`workflows.md`](workflows.md).

## Процесс

```text
intent -> behavior invariants -> change surface -> checkpoints -> scoped refactor -> verification evidence
```

1. Зафиксируй technical intent.
2. Зафиксируй behavior invariants.
3. Определи change surface и non-scope.
4. Определи checkpoints и stop conditions.
5. Реализуй scoped refactor без intentional behavior change.
6. Заполни verification/evidence carrier в issue/PR или `refactor.md`.

## Gates Переходов

### Routing -> Refactor Ready

- [ ] Routing Signature заполнен.
- [ ] `kind=refactor` или `kind=chore`.
- [ ] `contract_change=none` для `refactor-small`.
- [ ] `risk` не `high/critical` для `refactor-small`.
- [ ] Promotion triggers из [`workflows.md`](workflows.md) проверены.
- [ ] Выбран carrier: issue/PR note или managed `TASK-XXX/README.md` + `refactor.md`.

### Refactor Ready -> Execution Ready

- [ ] `INT-*` фиксирует technical intent.
- [ ] `INV-*` фиксирует observable behavior и public contracts, которые должны сохраниться.
- [ ] `SURF-*` фиксирует affected modules/paths.
- [ ] `NS-*` фиксирует, что не входит в refactor.
- [ ] `CP-*` задает checkpoints или stop conditions.
- [ ] `CHK-*` / `EVID-*` покрывают invariants.

Если проверяемые invariants нельзя сформулировать, refactor не готов к execution: сначала нужен discovery или promotion.

### Execution Ready -> Execution

- [ ] Change surface не выходит за `SURF-*`.
- [ ] Нет intentional observable behavior change.
- [ ] Выбранные проверки соответствуют `AGENTS.md` и [`../engineering/testing-policy.md`](../engineering/testing-policy.md).
- [ ] Manual-only verification имеет причину, процедуру и approval ref.

### Execution -> Done

- [ ] `INV-*` подтверждены проверками или approved manual evidence.
- [ ] Checkpoints закрыты или оставлены как explicit follow-up.
- [ ] Evidence carrier заполнен.
- [ ] Если behavior change стал нужным или неизбежным, task promoted или явно остановлен.

## Стабильные Идентификаторы

| Prefix | Значение | Где используется |
| --- | --- | --- |
| `INT-*` | intent refactor/chore | `refactor.md`, issue/PR |
| `INV-*` | invariant поведения | `refactor.md`, issue/PR |
| `SURF-*` | change surface | `refactor.md`, issue/PR |
| `NS-*` | non-scope | `refactor.md`, issue/PR |
| `CP-*` | checkpoint или stop condition | `refactor.md`, issue/PR |
| `CHK-*` | executable/manual check | `refactor.md`, issue/PR |
| `EVID-*` | evidence carrier | `refactor.md`, issue/PR |
| `AG-*` | approval gate | `refactor.md`, issue/PR |

## Правила Границ

1. Refactor/chore не меняет observable behavior.
2. Если behavior change желателен или необходим, task повышается до `feature-package` или другого подходящего profile.
3. Refactor не меняет API/schema/security/financial/integration/rollout contract в compact profile.
4. Checkpoints должны помогать остановить работу до behavior drift, а не заменять tests/evidence.
5. Системный refactor с несколькими independent delivery units повышается минимум до `managed-task`, часто до `epic-package`.

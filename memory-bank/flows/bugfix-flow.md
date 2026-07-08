---
title: Bugfix Flow
doc_kind: governance
doc_function: canonical
purpose: "Определяет процесс, gates переходов и evidence contract для bugfix-compact и managed bugfix задач."
derived_from:
  - ../dna/governance.md
  - workflows.md
  - task-flow.md
canonical_for:
  - bugfix_task_process
  - bugfix_root_cause_gate
  - bugfix_regression_evidence_rules
  - bugfix_identifier_taxonomy
status: active
audience: humans_and_agents
---

# Bugfix Flow

`bugfix-flow` применяется после selector-а из [`workflows.md`](workflows.md), когда выбран:

- `workflow_profile=bugfix-compact`;
- `workflow_profile=managed-task` и primary owner task package - `bugfix.md`.

Общие правила carrier/package, worktree, promotion и closure задает [`task-flow.md`](task-flow.md). Этот документ владеет только bugfix-процессом: symptom, reproduction, root cause, fix boundary и regression coverage.

## Комплект Документов

| Profile | Carrier | Когда достаточно |
| --- | --- | --- |
| `bugfix-compact` | GitHub issue/PR note по [`templates/task/bugfix.md`](templates/task/bugfix.md) | Дефект локален, root cause понятен, fix boundary помещается в PR/issue |
| managed bugfix | `memory-bank/tasks/TASK-XXX/README.md` + `bugfix.md` | Нужен durable context, но нет feature capability или contract/design trigger-а |

`bugfix.md` не должен содержать feature scope, design alternatives, rollout/backout design или implementation plan. Если это нужно, task повышается по [`workflows.md`](workflows.md).

## Процесс

```text
report -> symptom -> reproduction/evidence -> root cause -> fix boundary -> regression coverage -> scoped fix -> evidence
```

1. Зафиксируй observable symptom.
2. Зафиксируй reproduction или evidence source.
3. Найди root cause и запиши его как `RC-*`.
4. Определи fix boundary и non-scope.
5. Определи regression coverage до behavior-changing fix.
6. Реализуй scoped fix.
7. Заполни evidence carrier в issue/PR или `bugfix.md`.

## Gates Переходов

### Routing -> Bugfix Ready

- [ ] Routing Signature заполнен.
- [ ] `kind=bugfix`.
- [ ] `contract_change=none` для `bugfix-compact`.
- [ ] `risk` не `high/critical` для `bugfix-compact`.
- [ ] Promotion triggers из [`workflows.md`](workflows.md) проверены.
- [ ] Выбран carrier: issue/PR note или managed `TASK-XXX/README.md` + `bugfix.md`.

### Bugfix Ready -> Fix Ready

- [ ] `SYM-*` описывает, что сломано, где наблюдается и кто/что затронуто.
- [ ] `REPRO-*` содержит минимальное воспроизведение или ссылку на evidence.
- [ ] `RC-*` содержит проверяемую причину дефекта.
- [ ] `FIX-*` описывает минимальный fix boundary.
- [ ] `NS-*` фиксирует, что не меняется.
- [ ] `CHK-*` / `EVID-*` покрывают regression path.

Если root cause неизвестен, разрешены только discovery, logging, reproduction spec или diagnostic capture. Behavior-changing fix ждет известного `RC-*` или promotion.

### Fix Ready -> Execution

- [ ] Change surface не выходит за `FIX-*`.
- [ ] Выбранные проверки соответствуют `AGENTS.md` и [`../engineering/testing-policy.md`](../engineering/testing-policy.md).
- [ ] Manual-only verification имеет причину, процедуру и approval ref.
- [ ] Financial/security bugfix имеет automated regression coverage или approved manual-only exception с причиной.

### Execution -> Done

- [ ] Symptom больше не воспроизводится.
- [ ] Regression coverage выполнен или documented manual-only exception approved.
- [ ] Evidence carrier заполнен.
- [ ] Если в ходе fix появился contract/design/capability trigger, task promoted или явно остановлен.

## Стабильные Идентификаторы

| Prefix | Значение | Где используется |
| --- | --- | --- |
| `SYM-*` | наблюдаемый symptom | `bugfix.md`, issue/PR |
| `REPRO-*` | шаг воспроизведения или evidence source | `bugfix.md`, issue/PR |
| `RC-*` | root cause | `bugfix.md`, issue/PR |
| `FIX-*` | граница fix-а | `bugfix.md`, issue/PR |
| `NS-*` | non-scope | `bugfix.md`, issue/PR |
| `CHK-*` | executable/manual check | `bugfix.md`, issue/PR |
| `EVID-*` | evidence carrier | `bugfix.md`, issue/PR |
| `AG-*` | approval gate | `bugfix.md`, issue/PR |

## Правила Границ

1. Bugfix устраняет дефект, а не создает новую capability.
2. Fix boundary должен быть меньше или равен root cause boundary. Если fix требует wider redesign, task повышается.
3. Bugfix не меняет API/schema/security/financial/integration/rollout contract в compact profile.
4. Regression coverage обязателен для behavior-changing fix.
5. Если исправление создает новый устойчивый scenario, это feature work.

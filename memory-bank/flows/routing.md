---
title: Task Routing
doc_kind: governance
doc_function: canonical
purpose: Маршрутизация входящей задачи в минимальный flow, который сохраняет контроль над риском.
derived_from:
  - ../dna/governance.md
  - ../engineering/autonomy-boundaries.md
canonical_for:
  - task_routing_order
  - task_routing_predicates
  - workflow_type_selection
  - task_rerouting_rules
  - human_routing_rules
  - task_routing_outcome_contract
status: active
audience: humans_and_agents
---

# Task Routing

Этот документ выбирает flow для входящей задачи. Он не определяет lifecycle выбранной ветки: entry/exit gates, evidence и escalation принадлежат соответствующему flow-документу.

## Routing Order

Проверяй маршруты именно в этом порядке. `Small Change` — fast path перед Feature Flow, а не semantic type задачи.

```text
Issue / Task
     |
     +-- Incident / PIR? ----------------> Incident Flow
     |
     +-- Bug? ----------------------------> Bug Fix Flow
     |
     +-- Issue достаточен,
     |   design и plan не нужны? --------> Small Change Flow
     |
     +-- Новое или изменённое
     |   пользовательское поведение? ----> Feature Flow
     |
     +-- Refactoring? --------------------> Refactoring Flow
     |
     +-- Неясно / высокий риск ----------> Human Routing
```

## Routing Predicates

| Порядок | Вопрос | Route |
| --- | --- | --- |
| 1 | Есть активный operational impact, требуется containment или PIR? | [`Incident Flow`](incident.md) |
| 2 | Наблюдаемое поведение противоречит уже ожидаемому? | [`Bug Fix Flow`](bug-fix.md) |
| 3 | Выполнены все `Small Change` predicates ниже? | [`Small Change Flow`](small-change.md) |
| 4 | Создаётся или materially меняется пользовательское поведение? | [`Feature Flow`](feature.md) |
| 5 | Цель — изменить внутреннюю структуру при сохранении поведения? | [`Refactoring Flow`](refactoring.md) |
| 6 | Маршрут остаётся неоднозначным или риск не контролируется? | Human Routing |

### Small Change Gate

Все predicates должны быть истинны:

- issue/task полностью задаёт intent, scope и acceptance;
- решение следует конкретному существующему паттерну и не требует выбора подхода;
- не меняются API, event, schema, file format, CLI, env/config или integration contracts;
- не затрагиваются security boundary, data migration, rollout или обязательные approvals;
- change surface локален, test surfaces известны, отдельная декомпозиция и checkpoints не нужны.

Размер diff и оценка длительности сами по себе не являются routing predicates.

## Rerouting Rules

- Не начинай выбранный flow, пока выполнены его entry gates.
- Если в `Small Change` понадобились design, execution plan или новый устойчивый project fact, останови реализацию и повтори routing.
- Если report оказался изменением ожидаемого поведения, а не дефектом, выйди из Bug Fix Flow и повтори routing.
- Если refactoring меняет observable behavior, выйди из Refactoring Flow и повтори routing.
- Если задача меняет contract, rollout или требует approvals, она не может оставаться `Small Change`.

## Human Routing

Следуй canonical triggers из [`../engineering/autonomy-boundaries.md`](../engineering/autonomy-boundaries.md). Для routing дополнительно запрашивай решение человека, когда выбор flow требует продуктового решения, риск нельзя контролировать существующими gates или несколько route остаются одинаково правдоподобными после доступного исследования.

## Outcome / Exit Contract

### Observable Outcome

Для входящей задачи выбран ровно один допустимый flow либо явно зафиксирован `Human Routing`.

### Required Evidence

- issue/task или draft PR называет выбранный flow;
- запись показывает, какие entry predicates сделали route допустимым;
- для `Human Routing` зафиксированы вопрос, риск или конкурирующие routes.

### Terminal State

Routing завершён в состоянии `Routed`, когда выбранный flow и его entry gate подтверждены, либо в состоянии `Human Gate`, когда дальнейший выбор требует решения человека.

### Handoff

`Routed` передаёт задачу в выбранный flow. После решения `Human Gate` задача повторно проходит Task Routing; не вошедшая в выбранный scope работа маршрутизируется отдельно.

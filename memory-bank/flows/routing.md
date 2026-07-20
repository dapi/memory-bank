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

Flow определяет организацию lifecycle, но не глубину проверки. После выбора route отдельно выбери один [`validation profile`](../engineering/validation-profiles.md) в canonical owner выбранного delivery flow. Profile не участвует в routing order и не заменяет flow; если его triggers выявили contract, rollout или другой scope, несовместимый с текущим route, примени обычные rerouting rules.

## Routing Order

Проверяй маршруты именно в этом порядке. `Small Change` — fast path перед ветками Epic, Refactoring и Feature, а не semantic type задачи. После него сначала отделяй multi-feature Epic и behavior-preserving Refactoring, затем направляй оставшуюся single-delivery работу в Feature Flow.

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
     +-- Работа крупнее одной delivery-feature,
     |   нужен общий roadmap, cross-feature
     |   risk register или несколько
     |   delivery units? ----------------> Epic Flow
     |
     +-- Refactoring? --------------------> Refactoring Flow
     |
     +-- Одна delivery-unit меняет
     |   пользовательское поведение или
     |   доставляет planned engineering /
     |   operations outcome? ------------> Feature Flow
     |
     +-- Неясно / высокий риск ----------> Human Routing
```

## Routing Predicates

| Порядок | Вопрос | Route |
| --- | --- | --- |
| 1 | Есть активный operational impact, требуется containment или PIR? | [`Incident Flow`](incident.md) |
| 2 | Наблюдаемое поведение противоречит уже ожидаемому? | [`Bug Fix Flow`](bug-fix.md) |
| 3 | Выполнены все `Small Change` predicates ниже? | [`Small Change Flow`](small-change.md) |
| 4 | Работа крупнее одной delivery-feature и требует общего roadmap, cross-feature risk register или нескольких delivery units? | [`Epic Flow`](epic.md) |
| 5 | Цель — изменить внутреннюю структуру при сохранении поведения? | [`Refactoring Flow`](refactoring.md) |
| 6 | Задача укладывается в одну delivery-unit и создаёт или materially меняет пользовательское поведение либо доставляет плановое infrastructure, engineering или operations изменение с проверяемым outcome? | [`Feature Flow`](feature.md) |
| 7 | Маршрут остаётся неоднозначным или риск не контролируется? | Human Routing |

### Small Change Gate

Все predicates должны быть истинны:

- issue/task полностью задаёт intent, scope и acceptance;
- решение следует конкретному существующему паттерну и не требует выбора подхода;
- не меняются API, event, schema, file format, CLI, env/config или integration contracts;
- не затрагиваются security boundary, data migration, rollout или обязательные approvals;
- change surface локален, test surfaces известны, отдельная декомпозиция и checkpoints не нужны.

Размер diff и оценка длительности сами по себе не являются routing predicates.

### Epic Intake Handoff

Если признаки Epic route уже подтверждены, но problem, outcome, границы или evidence ещё недостаточны для canonical `charter.md`, задача всё равно маршрутизируется в [`Epic Flow`](epic.md). В этом случае Epic Flow начинается с `Epic Intake`: создаётся proposal package с `README.md` и `brief.md`, а недостающие факты фиксируются как open questions.

Неполнота epic facts сама по себе не является основанием для `Human Routing`. Human gate нужен только тогда, когда нельзя обоснованно выбрать route, требуется продуктовое решение о самом направлении инициативы или доступный риск нельзя контролировать intake boundaries.

## Rerouting Rules

- Не начинай выбранный flow, пока не выполнены его entry gates.
- Если в `Small Change` понадобились design, execution plan или новый устойчивый project fact, останови реализацию и повтори routing.
- Если в Feature Flow выяснилось, что работа крупнее одной delivery-feature и требует общего roadmap, cross-feature risk register или нескольких delivery units, останови feature package и повтори routing в [`Epic Flow`](epic.md).
- Не создавай delivery feature packages из Epic Intake. До `Roadmap Ready` proposal может называть только candidate delivery slices; accepted subissues и `FT-*` появляются после соответствующих epic gates.
- Если report оказался изменением ожидаемого поведения, а не дефектом, выйди из Bug Fix Flow и повтори routing.
- Если refactoring меняет observable behavior, выйди из Refactoring Flow и повтори routing.
- Если задача меняет contract, rollout или требует approvals, она не может оставаться `Small Change`.

## Human Routing

Следуй canonical triggers из [`../engineering/autonomy-boundaries.md`](../engineering/autonomy-boundaries.md). Для routing дополнительно запрашивай решение человека, когда выбор flow требует продуктового решения, риск нельзя контролировать существующими gates или несколько route остаются одинаково правдоподобными после доступного исследования.

## Outcome / Exit Contract

### Observable Outcome

Для входящей задачи выбран ровно один допустимый flow либо явно зафиксирован `Human Routing`.

### Required Evidence

- issue/task или draft PR называет выбранный flow; для active incident достаточно alert или incident-management record, подтверждающего operational impact или необходимость containment;
- запись показывает, какие entry predicates сделали route допустимым; provisional incident record может быть дополнен полным routing record после containment;
- для Epic route запись дополнительно указывает `Epic Intake`, когда facts ещё недостаточны для прямого `Bootstrap Epic`;
- для применимого delivery flow его canonical owner фиксирует отдельный validation profile decision по [`validation-profiles.md`](../engineering/validation-profiles.md); это downstream evidence выбора flow, а не дополнительный route;
- для `Human Routing` зафиксированы вопрос, риск или конкурирующие routes.

### Terminal State

Routing завершён в состоянии `Routed`, когда выбранный flow и его entry gate подтверждены, либо в состоянии `Human Gate`, когда дальнейший выбор требует решения человека.

### Handoff

`Routed` передаёт задачу в выбранный flow. Active incident передаётся в Incident Flow сразу после provisional routing: отсутствие issue/task или draft PR не блокирует containment, а repository trace создаётся или дополняется после стабилизации. После решения `Human Gate` задача повторно проходит Task Routing; не вошедшая в выбранный scope работа маршрутизируется отдельно.

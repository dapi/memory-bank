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
  - task_routing_priming_inputs
  - human_routing_priming_inputs
  - task_routing_outcome_contract
status: active
audience: humans_and_agents
---

# Task Routing

Этот документ выбирает flow для входящей задачи. Он не определяет lifecycle выбранной ветки: entry/exit gates, evidence и escalation принадлежат соответствующему flow-документу.

Flow определяет организацию lifecycle, но не глубину проверки. После выбора route отдельно выбери один [`validation profile`](../engineering/validation-profiles.md) в canonical owner выбранного delivery flow. Profile не участвует в routing order и не заменяет flow; если его triggers выявили contract, rollout или другой scope, несовместимый с текущим route, примени обычные rerouting rules.

## Context Priming Before Routing

До применения routing predicates выполни [`P0 Route Classification`](priming/context-priming.md#p0-route-classification): собери минимальные facts, чтобы выбрать flow или сформулировать Human Routing question. P0 не является implementation discovery, design или отдельным lifecycle; он заканчивается, как только route обоснован.

### P0 Priming Inputs

Прочитай [`routing.yaml`](priming/routing.yaml) и выполни source set `p0`.

После выбора route и до `Priming Inputs` соответствующего canonical
process-file выполни universal baseline, если задача создаёт или обновляет
governed-артефакт. Затем открой `Priming Inputs` process-file до первого
meaningful gate. Не открывай остальные flow-документы. Incident containment не
ждёт baseline; выполни его до создания или обновления governed
incident-артефакта.
Process priming не заменяет execution grounding: например, Feature Flow всё
ещё требует `GRND-*` evidence до sequencing.

## Routing Order

Проверяй маршруты именно в этом порядке. `Small Change` — fast path перед ветками Epic, Refactoring и Feature, а не semantic type задачи. После него сначала отделяй multi-feature Epic и behavior-preserving Refactoring, затем направляй оставшуюся single-delivery работу в Feature Flow.

```text
Issue / Task
     |
     +-- Incident / PIR? ----------------> Incident Flow
     |
     +-- Bug? ----------------------------> Bug Fix Flow
     |
     +-- Нужен evidence-backed ответ
     |   до коммита в delivery? ----------> Research & Discovery Flow
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
     +-- Route не выбран после FPF,
         нужен authority/value decision
         или риск не контролируется? ----> Human Routing
```

## Routing Predicates

| Порядок | Вопрос | Route |
| --- | --- | --- |
| 1 | Есть активный operational impact, требуется containment или PIR? | [`Incident Flow`](incident.md) |
| 2 | Наблюдаемое поведение противоречит уже ожидаемому? | [`Bug Fix Flow`](bug-fix.md) |
| 3 | Главная цель — получить evidence-backed answer для decision owner, а delivery outcome, scope или выбранный подход ещё не приняты? | [`Research & Discovery Flow`](research.md) |
| 4 | Выполнены все `Small Change` predicates ниже? | [`Small Change Flow`](small-change.md) |
| 5 | Работа крупнее одной delivery-feature и требует общего roadmap, cross-feature risk register или нескольких delivery units? | [`Epic Flow`](epic.md) |
| 6 | Цель — изменить внутреннюю структуру при сохранении поведения? | [`Refactoring Flow`](refactoring.md) |
| 7 | Задача укладывается в одну delivery-unit и создаёт или materially меняет пользовательское поведение либо доставляет плановое infrastructure, engineering или operations изменение с проверяемым outcome? | [`Feature Flow`](feature.md) |
| 8 | После P0-safe FPF classification и автономных tie-breakers route всё ещё нельзя обосновать, Research Flow не может законно закрыть unknown, требуется authority/value decision или риск не контролируется? | Human Routing |

### Small Change Gate

Все predicates должны быть истинны:

- issue/task полностью задаёт intent, scope и acceptance;
- решение следует конкретному существующему паттерну и не требует выбора подхода;
- не меняются API, event, schema, file format, CLI, env/config или integration contracts;
- не затрагиваются security boundary, data migration, rollout или обязательные approvals;
- change surface локален, test surfaces известны, отдельная декомпозиция и checkpoints не нужны.

Размер diff и оценка длительности сами по себе не являются routing predicates.

### Research & Discovery Gate

Выбирай этот route, когда задача прежде всего уменьшает uncertainty для решения, а не доставляет заранее определённое изменение. Примеры: market research, product discovery, technical feasibility spike, comparative evaluation, desk research или due diligence.

- вопрос, decision owner и expected decision могут быть зафиксированы;
- scope может быть exploratory, но должен быть timeboxed или иметь явный stopping condition;
- evidence, confidence и limitations важнее implementation plan;
- task не создаёт delivery package, ADR или committed roadmap только на основании неподтверждённой гипотезы.

Не выбирай Research Flow, если expected behavior уже известен и нужна только реализация: route сразу в минимальный delivery flow. Incident и Bug Fix остаются выше него: containment и восстановление expected behavior не ждут исследования.

### Epic Intake Handoff

Если признаки Epic route уже подтверждены, но problem, outcome, границы или evidence ещё недостаточны для canonical `charter.md`, задача всё равно маршрутизируется в [`Epic Flow`](epic.md). В этом случае Epic Flow начинается с `Epic Intake`: создаётся proposal package с `README.md` и `brief.md`, а недостающие факты фиксируются как open questions.

Неполнота epic facts сама по себе не является основанием для `Human Routing`.
Сначала примени FPF Decision Protocol из
[`autonomy-boundaries.md`](../engineering/autonomy-boundaries.md) в пределах P0.
Если missing facts требуют эксперимента или broad evidence collection, выбери
Research Flow, а не выполняй probe до routing. Human gate нужен только когда
результатом стал `escalate`: route нельзя обосновать без отсутствующего
product/value decision, нужны дополнительные полномочия или риск нельзя
контролировать intake boundaries.

## Rerouting Rules

- Не начинай выбранный flow, пока не выполнены его entry gates.
- Если в `Small Change` понадобились design, execution plan или новый устойчивый project fact, останови реализацию и повтори routing.
- Если Research Flow сформировал delivery proposal, architecture decision, product initiative или change request, не начинай delivery внутри research package: зафиксируй terminal disposition в `brief.md: research_status` и повтори routing в PRD, Epic, Feature, ADR или другой применимый owner.
- Если в Feature Flow выяснилось, что работа крупнее одной delivery-feature и требует общего roadmap, cross-feature risk register или нескольких delivery units, останови feature package и повтори routing в [`Epic Flow`](epic.md).
- Не создавай delivery feature packages из Epic Intake. До `Roadmap Ready` proposal может называть только candidate delivery slices; accepted subissues и `FT-*` появляются после соответствующих epic gates.
- Если report оказался изменением ожидаемого поведения, а не дефектом, выйди из Bug Fix Flow и повтори routing.
- Если refactoring меняет observable behavior, выйди из Refactoring Flow и повтори routing.
- Если задача меняет contract, rollout или требует approvals, она не может оставаться `Small Change`.

## Human Routing

Следуй canonical FPF Decision Protocol и triggers из
[`../engineering/autonomy-boundaries.md`](../engineering/autonomy-boundaries.md).
Не отправляй задачу в Human Routing только из-за сложности, неполных данных или
нескольких правдоподобных routes. Сначала используй routing predicates,
canonical facts и автономные tie-breakers. P0 остаётся read-only: если unknown
требует эксперимента, implementation discovery или изменения файлов, выбери
Research Flow и выполни `bounded_probe` внутри его lifecycle.

Запрашивай решение человека, только когда FPF outcome — `escalate`: выбор flow
требует отсутствующего product/business value judgment или дополнительных
полномочий, ни один route не сохраняет обязательные constraints либо риск нельзя
контролировать существующими gates.

### Human Routing Priming Inputs

Прочитай [`routing.yaml`](priming/routing.yaml) и выполни source set
`human_routing`.

Перед запросом человека зафиксируй competing routes, применённые decision
criteria, P0 evidence, unknown, причину неприменимости Research Flow, approval
trigger и точный вопрос. До решения не начинай delivery, broad research, design
или изменение файлов; после него повтори Task Routing.

## Outcome / Exit Contract

### Observable Outcome

Для входящей задачи выбран ровно один допустимый flow либо явно зафиксирован `Human Routing`.

### Required Evidence

- issue/task или draft PR называет выбранный flow; для active incident достаточно alert или incident-management record, подтверждающего operational impact или необходимость containment;
- P0 evidence обосновывает выбранный flow; после routing P1 result находится в canonical owner выбранного flow, а не в отдельном priming report;
- запись показывает, какие entry predicates сделали route допустимым; provisional incident record может быть дополнен полным routing record после containment;
- для Epic route запись дополнительно указывает `Epic Intake`, когда facts ещё недостаточны для прямого `Bootstrap Epic`;
- для Research route запись указывает decision question, decision owner и stopping condition;
- для применимого delivery flow его canonical owner фиксирует отдельный validation profile decision по [`validation-profiles.md`](../engineering/validation-profiles.md); это downstream evidence выбора flow, а не дополнительный route;
- для `Human Routing` зафиксированы FPF outcome `escalate`, вопрос, риск или
  конкурирующие routes и причина, по которой routing criteria, автономные
  tie-breakers и Research Flow не дают допустимого продолжения.

### Terminal State

Routing завершён в состоянии `Routed`, когда выбранный flow и его entry gate подтверждены, либо в состоянии `Human Gate`, когда дальнейший выбор требует решения человека.

### Handoff

`Routed` передаёт задачу в выбранный flow. Active incident передаётся в Incident Flow сразу после provisional routing: отсутствие issue/task или draft PR не блокирует containment, а repository trace создаётся или дополняется после стабилизации. После решения `Human Gate` задача повторно проходит Task Routing; не вошедшая в выбранный scope работа маршрутизируется отдельно.

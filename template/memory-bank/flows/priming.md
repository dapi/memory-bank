---
title: Task Context Priming
doc_kind: governance
doc_function: canonical
purpose: Общий контракт праймеринга контекста перед Task Routing и профильные P1-проверки перед первым gate каждого flow.
derived_from:
  - ../dna/principles.md
  - ../dna/governance.md
  - routing.md
canonical_for:
  - task_context_priming_contract
  - route_priming_profiles
  - priming_result_ownership
status: active
audience: humans_and_agents
---

# Task Context Priming

Праймеринг контекста подготавливает агента к следующему решению; он не
выбирает lifecycle и не создаёт второй owner фактов. [`Task Routing`](routing.md)
выбирает flow, а этот документ определяет минимальные сведения, которые нужны
до routing и до первого gate выбранного flow.

Праймеринг всегда применяет progressive disclosure: индекс → релевантный
раздел → конкретный owner-документ, код или evidence. Не загружай весь
репозиторий или весь Memory Bank, если следующий gate этого не требует.

```text
task → P0: route classification → Task Routing
     → P1: route profile → first flow gate
     → P2: flow-specific execution grounding, если оно требуется
```

`P2` не является универсальным шагом. Например, Feature Flow сохраняет
execution-grounding с `GRND-*` против immutable commit SHA перед sequencing;
`P1-FEAT` не заменяет и не ослабляет это требование.

## Common Contract

### P0 Route Classification

Выполни P0 после чтения startup-инструкций и до применения routing predicates:

1. прочитай source task, report, alert или другой trigger и отдели stated facts
   от предположений;
2. используй проектный индекс и routing rules, чтобы найти только facts,
   необходимые для различения routes;
3. зафиксируй candidate route, evidence references, material unknowns и риск,
   который может потребовать Human Routing;
4. не начинай implementation discovery, selected design или изменение файлов.

P0 заканчивается сразу после того, как route обоснован либо сформулирован
точный вопрос для Human Routing. Признак incident прекращает broad discovery:
сразу выбери Incident Flow и продолжи его timeboxed `P1-INC`.

### P1 Route Profile

После Task Routing, но до первого meaningful gate выбранного flow, выполни его
P1-профиль. Каждый профиль задаёт:

- purpose и минимальную область чтения;
- факты и unknowns, нужные следующему gate;
- canonical owner результата;
- stop condition, который не позволяет превратить праймеринг в скрытое
  исследование или delivery.

Не создавай отдельный `priming report`: результаты живут в уже существующем
owner-е route. В выводе всегда различай observed facts, источники и hypotheses.

## Route Profiles

### P1-INC Incident And PIR

- **Собрать:** active impact, affected surfaces, доступные recovery signals,
  релевантный runbook и, если это не задерживает containment, последний
  deployment/change evidence.
- **Результат:** initial timeline/incident record фиксирует facts, owner и
  containment options; hypotheses не выдаются за root cause.
- **Граница:** timebox discovery. Containment и human incident owner важнее
  полного понимания системы.

### P1-BUG Bug Fix

- **Собрать:** canonical source expected behavior, observed actual behavior,
  minimum reproduction inputs/environment и ближайшие existing tests или
  evidence их отсутствия.
- **Результат:** bug report или linked delivery task различает expected/actual,
  reproduction evidence и open uncertainty для Entry/Reproduction gates.
- **Граница:** не выбирай новый product behavior и не начинай refactoring.

### P1-RES Research And Discovery

- **Собрать:** decision question, decision owner, known evidence, relevant
  access/privacy constraints, working assumptions и candidate stopping
  condition.
- **Результат:** `research/R-XXX/brief.md` после bootstrap владеет question,
  scope, assumptions, unknowns и stopping condition.
- **Граница:** не создавай delivery package, selected solution или committed
  roadmap только из предположения.

### P1-SMALL Small Change

- **Собрать:** task intent/scope/acceptance, реально существующий reference
  pattern, local change surface и known test/verify surface.
- **Результат:** Small Change routing record ссылается на проверенный pattern и
  конкретные verify actions.
- **Граница:** если pattern не подходит, change surface не локален или нужны
  design/plan/contract decisions, остановись и повтори Task Routing.

### P1-REF Refactoring

- **Собрать:** observable behavior и contracts, которые должны сохраниться,
  baseline/characterization coverage, structural change surface и checkpoint
  constraints.
- **Результат:** исходная task фиксирует preservation boundary, baseline и
  material unknowns для Entry Gate.
- **Граница:** намеренное behavior/contract change не маскируй как
  refactoring; верни его в Task Routing.

### P1-EPIC Epic

- **Собрать:** source/trigger, problem/outcome, rough scope/non-scope,
  available evidence, stakeholders/decision owner, candidate slices и open
  questions.
- **Результат:** `epics/EP-XXX/brief.md` при Intake владеет proposal facts;
  при достаточных facts `charter.md` получает canonical intent.
- **Граница:** не создавай accepted subissues, delivery feature packages,
  roadmap waves или code-level plan до соответствующих Epic gates.

### P1-FEAT Feature

- **Собрать:** relevant upstream PRD/epic/use case/ADR, applicable domain and
  engineering rules, affected contracts and a bounded view of current
  repository surfaces.
- **Результат:** draft `features/FT-XXX/brief.md` получает problem-space facts,
  assumptions, constraints и unresolved decisions; он не принимает selected
  solution.
- **Граница:** не подменяй `P2` execution-grounding. Перед Plan Ready всё ещё
  требуются immutable revision и `GRND-*` evidence по current code/test state.

### P1-HUMAN Human Routing

- **Собрать:** competing routes, известные risk/approval trigger, missing fact
  или product decision и минимальные evidence references.
- **Результат:** task формулирует точный вопрос человеку и последствия каждого
  допустимого route.
- **Граница:** не проводи broad research и не продолжай delivery до решения;
  после решения повтори Task Routing.

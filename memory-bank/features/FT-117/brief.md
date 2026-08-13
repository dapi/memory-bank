---
title: "FT-117: Autonomous Structured Decision Protocol"
doc_kind: feature
doc_function: canonical
purpose: "Canonical problem-space brief для issue #117. Фиксирует автономное принятие bounded решений, границу approval и проверяемый delivery outcome без selected solution details."
derived_from:
  - ../../flows/feature.md
  - ../../../template/memory-bank/flows/routing.md
  - ../../../template/memory-bank/flows/research.md
  - ../../../template/memory-bank/engineering/autonomy-boundaries.md
  - ../../../template/memory-bank/engineering/validation-profiles.md
status: active
delivery_status: in_progress
audience: humans_and_agents
must_not_define:
  - implementation_sequence
  - selected_solution_details
---

# FT-117: Autonomous Structured Decision Protocol

## What

### Problem

Issue [#117](https://github.com/dapi/memory-bank/issues/117) обнаружила, что
сложность, ambiguity, несколько допустимых подходов и исчерпание review budget
могут преждевременно останавливать агента на Human Gate. Одновременно правила
смешивали decision authority, durable decision record и permission на
внешне-эффективный execution.

### Outcome

| Metric ID | Metric | Baseline | Target | Measurement method |
| --- | --- | --- | --- | --- |
| `MET-01` | Покрытие decision/approval contract | Роли и execution permission описаны неполно | Каждое non-trivial решение различает authority source, decision owner, carrier и execution approval | Semantic review canonical docs и deterministic documentation checks |
| `MET-02` | Premature Human Gate triggers | Ambiguity/complexity/review exhaustion могли эскалировать сами по себе | Эскалация только после `Structured Decision Protocol` outcome `escalate` или отдельного execution gate | Review Routing, Bug Fix, Feature и Validation contracts |
| `MET-03` | Проверяемость документации | Cross-flow rules могли расходиться | Priming manifests, links, reachability и doctor проходят без ошибок | Repository validation commands |

## Scope

- `REQ-01` Определить обязательный self-contained decision protocol, который
  даёт проверяемый результат: автономное продолжение, bounded investigation
  или конкретная эскалация.
- `REQ-02` Разделить authority source, accountable decision owner, canonical
  carrier, execution approver и approval evidence.
- `REQ-03` Разрешить автономный выбор между допустимыми вариантами на основе
  заданного intent, ограничений и проверяемого rationale; близость вариантов
  сама по себе не должна создавать Human Gate.
- `REQ-04` Различать unknown, требующий discovery evidence до выбора delivery
  route, и implementation-only unknown после выбора route.
- `REQ-05` Ограничить Human Gate authority boundary, material
  security/auth/trust/compliance boundary mutation без specific task/policy
  authority, external effect, explicit approval policy, missing value judgment
  и uncontrollable risk.
- `REQ-06` Убрать standalone escalation по сложности и review exhaustion из
  применимых Bug Fix и Feature flow contracts.
- `REQ-07` Требовать scoped и current approval evidence, связанное с тем
  конкретным execution gate, для которого нужно разрешение.
- `REQ-08` Разрешить обычное non-risky изменение кода и подготовку PR без
  Human Gate на сам edit step; PR review, rollback и stop conditions остаются
  контролями delivery.

## Non-Scope

- `NS-01` Изменение runtime-кода, production configuration, release path или
  deployment behavior.
- `NS-02` Неявное разрешение merge, release, deployment, external messaging,
  publication, live-data mutation или других external writes.
- `NS-03` Обязательное применение FPF или другой внешней reasoning methodology;
  Structured Decision Protocol должен оставаться self-contained.
- `NS-04` Создание отдельного ADR для feature-local governance contract.
- `NS-05` Изменение human-only prompt catalog: prompt files не входят в scope
  governed implementation этой feature.

## Constraints / Assumptions

- `ASM-01` Issue #117 и её acceptance criteria являются task-level source для
  delivery intent; canonical ownership после реализации принадлежит
  `engineering/autonomy-boundaries.md` и соответствующим flow documents.
- `CON-01` Existing project policies, compliance rules, contracts и explicit
  approvals остаются stronger constraints и не могут быть отменены protocol.
- `CON-02` P0 остаётся read-only и не выполняет implementation discovery,
  experiments, file changes или external writes.
- `CON-03` Решение является одной documentation delivery-unit; runtime C4 и
  production rollout не применимы.
- `CON-04` Отдельное product/business value judgment не требуется, если
  canonical task уже задаёт intent и acceptance; missing value priority
  остаётся `escalate`.
- `CON-05` FPF или другая reasoning methodology могут использоваться как
  опциональный supporting analysis, но их отсутствие не блокирует protocol и
  не создаёт Human Gate.
- `CON-06` Core documentation candidate появился в commits `9822d88`–`d3639b1`
  до bootstrap этого feature package. Recovery review может принять или
  отклонить этот existing candidate, но не должен изображать, что исторические
  writes прошли Plan Ready задним числом.

## Design Requirement Decision

| Decision | Reason | Downstream owner |
| --- | --- | --- |
| `Design required: yes` | Feature вводит новый cross-flow decision protocol и меняет authority, carrier, routing и approval contracts | [`design.md`](design.md) |

## Artifact Routing Decision

| Artifact | Decision | Trigger / reason | Route / owner |
| --- | --- | --- | --- |
| `use-cases/README.md` | omitted | Нет нового устойчивого user/operator journey; acceptance scenarios полностью покрываются в brief | `none` |
| `runtime-surfaces.md` | omitted | Нет runtime surface или implementation inventory, который нужно проецировать | `none` |
| `contracts/<name>.md` | omitted | Нет API/event/queue/file/runtime-config connector; governance contract остаётся в design | `none` |
| `diagrams/<name>-sequence.md` | omitted | Lifecycle описывается в flow prose и таблицах, отдельная temporal review boundary не нужна | `none` |
| ADR | omitted | Решение feature-local и не меняет reusable architecture boundary | `none` |

## Validation Profile Decision

| Profile | Triggers / rationale | Downgrade approval |
| --- | --- | --- |
| `documentation` | Меняются только governed Markdown и priming YAML; executable behavior, production config, contracts и release paths не меняются | `none` |

## Verify

### Acceptance scenarios

- `SC-01` При нескольких допустимых local approaches агент отбрасывает
  нарушающие constraints, фиксирует rationale и продолжает без эскалации
  только из-за близости вариантов.
- `SC-02` При unknown, требующем experiment/discovery/mutation до выбора
  delivery route, задача передаётся в Research Flow без mutation state и с
  явными вопросом, budget и stopping condition.
- `SC-03` После Research probe evidence и limitations записываются, затем
  повторяются Structured Decision Protocol и Task Routing до delivery.
- `SC-04` При missing product/value judgment protocol возвращает `escalate` с
  точным вопросом, а не выбирает priority молча.
- `SC-05` При unauthorized external/production action подготовка и validation
  продолжаются, но exact execution gate остаётся заблокирован.
- `SC-06` Scoped current preauthorization принимается только для названного
  action, target/environment, scope и limits.
- `SC-07` Исчерпание review budget останавливает неизменённый loop и вызывает
  replan, bounded probe или `escalate`, но не создаёт Human Gate само по себе.
- `SC-08` Implementation-only unknown для уже выбранного delivery route
  откладывается в P1/P2 grounding, а не меняет P0 route на Research.
- `SC-09` При заданных intent и scope агент может автономно внести обычное
  non-risky изменение кода и подготовить PR; Human Gate не возникает на edit
  step, а review, rollback и stop conditions остаются обязательными.
- `SC-10` Агент автономно готовит analysis, design, validation и rollback для
  security/compliance change, но не вносит material boundary mutation без
  specific authority; наличие такой authority не разрешает последующее
  risk-bearing production/live execution.

### Negative cases

- `NEG-01` Broad, inferred, stale или overridden permission не считается
  approval evidence.
- `NEG-02` `Execution approval` не меняет protocol outcome и не расширяет
  authority за пределы exact gate.
- `NEG-03` Tie-breakers не могут выбрать отсутствующий business priority и не
  могут отменить hard constraint, policy или compliance requirement.

### Checks and evidence

| Check ID | Check | Expected evidence |
| --- | --- | --- |
| `CHK-01` | `ruby tools/validate-priming-manifests.rb template/memory-bank` | Все priming manifests schema-valid и paths resolvable |
| `CHK-02` | `memory-bank-cli lint --scope-root template/memory-bank --entrypoint template/memory-bank/README.md`; `memory-bank-cli lint --scope-root memory-bank --entrypoint memory-bank/README.md` | Generic template и project-level package не содержат broken links, orphan docs или unreachable docs |
| `CHK-03` | `memory-bank-cli doctor --profile template` | Нет doctor errors/warnings |
| `CHK-04` | `git diff --check` | Нет whitespace errors или conflict markers |
| `CHK-05` | Semantic review of autonomy, routing, validation, Bug Fix and Feature docs | All `REQ-*`, `SC-*` and `NEG-*` have canonical owner/evidence path |

### Evidence contract

`EVID-01` — текущий issue #117 и её comments/acceptance criteria.

`EVID-02` — canonical [autonomy boundaries](../../../template/memory-bank/engineering/autonomy-boundaries.md)
с protocol schema, outcome semantics, carrier rules и approval evidence.

`EVID-03` — [Task Routing](../../../template/memory-bank/flows/routing.md), [Research Flow](../../../template/memory-bank/flows/research.md),
Bug Fix и Feature flow с P0, Research handoff, review convergence и Human
Routing rules.

`EVID-04` — validation profile contract с exact execution approval gate.

`EVID-05` — outputs `CHK-01`–`CHK-04` и semantic review record этого package.

`EVID-06` — отдельные non-authoring Plan Ready и implementation review records,
которые называют frozen revisions, findings, dispositions и verdicts.

## Traceability

| Requirement | Acceptance / evidence |
| --- | --- |
| `REQ-01`, `REQ-02`, `REQ-03`, `REQ-07` | `SC-01`, `SC-04`, `SC-05`, `SC-06`, `NEG-01`–`NEG-03`, `EVID-02`, `EVID-04` |
| `REQ-04` | `SC-02`, `SC-03`, `SC-08`, `EVID-03` |
| `REQ-05`, `REQ-06` | `SC-04`, `SC-05`, `SC-07`, `SC-10`, `EVID-02`–`EVID-04` |
| `REQ-08` | `SC-09`, `SC-10`, `EVID-02`–`EVID-04` |
| `REQ-01`–`REQ-08` | `CHK-01`–`CHK-05`, `EVID-05`, `EVID-06` |

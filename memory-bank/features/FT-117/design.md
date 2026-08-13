---
title: "FT-117: Structured Decision Protocol Design"
doc_kind: feature
doc_function: canonical
purpose: "Selected solution и design-pack manifest для FT-117. Определяет ownership и cross-flow contract, не заменяя canonical autonomy и flow documents."
derived_from:
  - brief.md
  - ../../../template/memory-bank/engineering/autonomy-boundaries.md
  - ../../../template/memory-bank/engineering/validation-profiles.md
  - ../../../template/memory-bank/flows/routing.md
  - ../../../template/memory-bank/flows/research.md
  - ../../../template/memory-bank/flows/bug-fix.md
  - ../../../template/memory-bank/flows/feature.md
status: active
audience: humans_and_agents
must_not_define:
  - ft_117_problem_space
  - ft_117_acceptance_criteria
  - implementation_sequence
---

# FT-117: Structured Decision Protocol Design

## Solution Summary

`SOL-01` — `engineering/autonomy-boundaries.md` остаётся canonical owner
Structured Decision Protocol, carrier selection, execution authorization и
approval evidence. Protocol является self-contained; FPF или другая reasoning
methodology могут дать supporting analysis, но не являются dependency.

`SOL-02` — `flows/routing.md` реализует read-only P0, Research handoff и Human
Routing только после protocol outcome `escalate`.

`SOL-03` — `flows/bug-fix.md` и `flows/feature.md` используют protocol для
неопределённости и review exhaustion; complexity, ambiguity и exhausted budget
не являются самостоятельными Human Gates.

`SOL-04` — `engineering/validation-profiles.md` отделяет autonomous design,
rehearsal, validation и rollback preparation от exact approval-gated execution
step.

`SOL-05` — `flows/research.md` владеет probe lifecycle после P0 handoff и
возвратом evidence в повторный Structured Decision Protocol и Task Routing.

`SOL-06` — ordinary non-risky code edit остаётся autonomous work, а PR review,
rollback и stop conditions остаются assurance controls без Human Gate на edit
step. Analysis, design и validation security/compliance change также автономны,
но material boundary mutation требует specific task/policy authority и не
разрешает последующий risk-bearing production/live execution.

## Design Decisions

### Solution Traceability

| Requirement refs | Selected solution refs |
| --- | --- |
| `REQ-01`, `REQ-02`, `REQ-07` | `SOL-01`, `SD-01`, `SD-04` |
| `REQ-03` | `SOL-01`, `SD-02` |
| `REQ-04` | `SOL-02`, `SOL-05`, `SD-03` |
| `REQ-05`, `REQ-06` | `SOL-03`, `SOL-04`, `SOL-06` |
| `REQ-08` | `SOL-06` |

### `SD-01` — Four-role decision record

Каждая non-trivial record различает authority source, accountable decision owner,
canonical carrier и execution approver/approval evidence.
Это предотвращает
смешение rationale с permission и позволяет сохранять decision autonomy при
заблокированном external execution.

### `SD-02` — Hard constraints before tie-breakers

Сначала исключаются options, нарушающие intent, invariants, contracts, policy,
compliance или authority. Затем применяются tie-breakers в фиксированном порядке:
existing canonical pattern, smallest reversible change, smallest blast radius,
lowest operational/maintenance complexity, strongest available verification.
Tie-breakers не выбирают отсутствующий business value priority.

### `SD-03` — Probe belongs to Research before delivery

P0 только классифицирует route. Если unknown требует experiment, implementation
discovery, broad evidence collection или mutation до выбора delivery route, он
передаётся Research Flow с budget и stopping condition. После probe protocol и
routing повторяются.

### `SD-04` — Approval checked at exact gate

Decision outcome не является execution permission. Approval evidence должна быть
specific, current, scoped, owned и привязана к exact execution gate; broad или
inferred permission отклоняется.

## Design Pack

| Artifact | Relation | Direct canonical ownership | Readiness / source |
| --- | --- | --- | --- |
| `design.md` | `root` | Feature-local selected solution and `SD-01`–`SD-04`; default owner for `SOL-01`–`SOL-06` | `active`; pack root |
| `engineering/autonomy-boundaries.md` | `external-dependency` | Protocol, carrier, authorization и approval-evidence semantics | `active`; canonical owner |
| `flows/routing.md` | `external-dependency` | P0, Research handoff и Human Routing semantics | `active`; canonical flow owner |
| `flows/research.md` | `external-dependency` | Probe lifecycle, evidence handoff и repeat-routing semantics | `active`; canonical flow owner |
| `flows/bug-fix.md` | `external-dependency` | Bug Fix reroute/replan semantics | `active`; canonical flow owner |
| `flows/feature.md` | `external-dependency` | Feature ambiguity/review-convergence semantics | `active`; canonical flow owner |
| `engineering/validation-profiles.md` | `external-dependency` | Validation and execution-approval separation | `active`; canonical profile owner |

## C4 Applicability Decision

`C4-00: not required` — change surface состоит из governed documentation и
priming manifests. Он не вводит или не меняет deployable/runtime/container,
queue, storage, network, security runtime boundary или production topology.
Cross-document ownership отображена в Design Pack и не требует C4 runtime view.

## 4+1 Viewpoint Coverage

| View | Status | Coverage / source |
| --- | --- | --- |
| Logical | `covered` | `REQ-*`, `SC-*` и ownership split из `brief.md` / `SD-01` |
| Process | `covered` | P0 → Research → repeat routing и exact execution gate из `SD-03`/`SD-04` |
| Development | `covered` | Canonical owners: autonomy, routing, flow и validation docs |
| Physical | `N/A` | Нет runtime/deployable или deployment topology change; `C4-00` |
| Scenarios (+1) | `covered` | `SC-01`–`SC-09` и `NEG-01`–`NEG-03` |

## Architecture Coverage Decision

| Concern | Decision | Evidence |
| --- | --- | --- |
| Components / ownership | `covered` | `SOL-01`–`SOL-06`, Design Pack |
| Connectors / bindings | `covered` as document references | Cross-flow Markdown dependencies and priming inputs; no runtime connector |
| Configuration / contracts | `covered` | Protocol schema and approval evidence contract |
| Behavioral semantics | `covered` | Outcome semantics, P0 and Human Gate rules |
| Quality / evolution | `covered` | Reversibility, blast radius, review convergence and rollback preparation |

## Cross-View Correspondence

| Scenario | Logical ref | Process ref | Development ref | Physical ref | Checks / evidence |
| --- | --- | --- | --- | --- | --- |
| `SC-01` | `REQ-03`, `SD-02` | Protocol option selection | `autonomy-boundaries.md` | `N/A`: no runtime topology | `CHK-05`, `EVID-02` |
| `SC-02` | `REQ-04` | P0 → Research handoff | `routing.md`, `research.md` | `N/A`: pre-delivery read-only routing | `CHK-05`, `EVID-03` |
| `SC-03` | `REQ-04` | Probe → evidence → repeat protocol/routing | `research.md`, `routing.md` | `N/A`: no runtime topology | `CHK-05`, `EVID-03` |
| `SC-04` | `REQ-05` | Escalation with exact question | `autonomy-boundaries.md` | `N/A`: no runtime topology | `CHK-05`, `EVID-02` |
| `SC-05` | `REQ-05`, `REQ-07` | Preparation → exact gated execution | `autonomy-boundaries.md`, `validation-profiles.md` | `N/A`: no runtime topology | `CHK-05`, `EVID-04` |
| `SC-06` | `REQ-07` | Scoped authorization at execution gate | `validation-profiles.md` | `N/A`: no runtime topology | `CHK-05`, `EVID-04` |
| `SC-07` | `REQ-06` | Review exhaustion → replan/probe/escalate | `feature.md`, `autonomy-boundaries.md` | `N/A`: no runtime topology | `CHK-05`, `EVID-03` |
| `SC-08` | `REQ-04` | Established delivery route → P1/P2 grounding | `routing.md`, `feature.md` | `N/A`: no runtime topology | `CHK-05`, `EVID-03` |
| `SC-09` | `REQ-08` | Edit → PR review → rollback/stop controls | `autonomy-boundaries.md`, `validation-profiles.md` | `N/A`: runtime implementation is out of this documentation feature | `CHK-05`, `EVID-03`, `EVID-04` |
| `SC-10` | `REQ-05`, `REQ-08` | Preparation → repository authority gate → separate live execution gate | `autonomy-boundaries.md`, `validation-profiles.md` | `N/A`: governance contract only | `CHK-05`, `EVID-02`, `EVID-04` |

## Design Verification

| Analysis | Required | Method | Result / evidence |
| --- | --- | --- | --- |
| Decision authority vs execution permission | `yes` | Semantic review of `SD-01` and `SD-04` | Fields and exact gate are explicit in canonical owner |
| P0 / Research routing | `yes` | Review routing predicates and priming manifest | P0 read-only; probe handoff is explicit |
| Research lifecycle handoff | `yes` | Review `research.md` boundary and repeat-routing contract | Probe owns question/budget/evidence and cannot silently commit delivery |
| Review exhaustion | `yes` | Review Bug Fix/Feature gate text | Replan/probe/escalate replaces exhaustion-only gate |
| Ordinary code edit vs PR review | `yes` | Review autonomy and validation contracts | Edit is not a Human Gate; PR review and rollback remain controls |
| Contract compatibility | `no` | No runtime/API contract changes | No API, event, schema or file-format surface |
| State/transition completeness | `yes` | Review protocol outcomes and Research handoff states | `proceed`, `bounded_probe`, `escalate` and repeat routing are explicit |
| Failure propagation | `yes` | Review stop/escalation semantics | Unbounded risk leads to exact escalation or stop |
| Concurrency/ordering | `no` | No runtime execution ordering introduced | No concurrent runtime path in scope |
| Security boundaries | `yes` | Review approval and security-sensitive Human Gate rules | Security/live-state boundary remains explicit |
| Capacity/latency | `no` | No runtime performance target | Documentation-only change |
| Migration/evolution safety | `yes` | Review carrier ownership, rollback and review convergence | Canonical owner and revision handoff are explicit |
| C4 runtime model | `no` | Applicability check | `C4-00`; no runtime boundary |
| UI / interaction contract | `no` | Trigger check | No UI/API/event/file/runtime-config change |

## Solution Ready

- `SOL-01`–`SOL-06` have a direct canonical owner.
- `SD-01`–`SD-04` are feature-local rationale; reusable semantics remain in
  their external canonical owners.
- `C4-00` and 4+1 coverage are recorded.
- No optional artifact is required for Solution Ready.

---
title: "FT-012: Implementation Plan"
doc_kind: feature
doc_function: derived
purpose: "Execution-план реализации FT-012. Фиксирует discovery context, шаги, риски и test strategy без переопределения canonical problem и solution фактов."
derived_from:
  - brief.md
  - design.md
status: archived
audience: humans_and_agents
must_not_define:
  - ft_012_scope
  - ft_012_selected_design
  - ft_012_acceptance_criteria
  - ft_012_blocker_state
---

# План имплементации

## Цель текущего плана

Выполнить docs-only реализацию issue #12: добавить compact task-flow family, task templates, optional `memory-bank/tasks/` destination и обновить routing/index docs так, чтобы acceptance из `brief.md` проверялась локальными documentation checks.

## Grounding / Support References

| Document | Role in this plan | Facts reused | Conflict action |
| --- | --- | --- | --- |
| `brief.md` | canonical problem / verify owner | `REQ-*`, `SC-*`, `NEG-*`, `CHK-*`, `EVID-*` | Update `brief.md` first |
| `design.md` | solution owner | `SOL-*`, `C4-*`, `SD-*`, `CTR-*`, `INV-*`, `FM-*`, `RB-*` | Update `design.md` or ADR first |
| `decision-log.md` | local FPF decision ledger | Decisions that close important ambiguities | Update decision log if a new material ambiguity is resolved |
| GitHub issue `#12` | external requirement source | Source list, scope and acceptance | Human gate if issue scope conflicts with implemented docs |
| `alfagen/mercury` source docs | external source evidence | Generic flow/template patterns | Strip project-specific examples before target docs |

## Current State / Reference Points

| Path / module | Current role | Why relevant | Reuse / mirror |
| --- | --- | --- | --- |
| `memory-bank/flows/feature-flow.md` | Canonical feature package lifecycle | Task-flow must not redefine feature ownership | Reuse promotion boundary back to feature package |
| `memory-bank/flows/workflows.md` | Current high-level task routing | Needs compact profiles and Routing Signature | Extend selector without deleting existing flow types |
| `memory-bank/flows/README.md` | Flow navigation index | Must expose new flow docs | Add annotated links |
| `memory-bank/flows/templates/README.md` | Template navigation index | Must expose task templates | Add task template group and derived links |
| `memory-bank/README.md` | Root memory-bank index | Must expose optional `tasks/` destination | Add annotated `tasks/README.md` route |
| `memory-bank/features/README.md` | Feature packages index | Must expose FT-012 docs package | Add annotated FT-012 link |
| `scripts/check_memory_bank_index.py` | Link/reachability audit | Acceptance explicitly requires this check | Run as `CHK-01` |
| `memory-bank/engineering/testing-policy.md` | Testing policy for docs changes | Sets manual-only and evidence expectations | Reuse docs/check evidence split |

## Test Strategy

| Test surface | Canonical refs | Existing coverage | Planned automated coverage | Required local suites / commands | Required CI suites / jobs | Manual-only gap / justification | Manual-only approval ref |
| --- | --- | --- | --- | --- | --- | --- | --- |
| Link/index/frontmatter integrity | `REQ-01`..`REQ-04`, `CHK-01` | `scripts/check_memory_bank_index.py` exists | Run existing audit after docs changes | `python3 scripts/check_memory_bank_index.py` | Same command if CI runs it | none | none |
| Whitespace/conflict markers | `EC-06`, `CHK-05` | Git diff check available | Run git diff audit | `git diff --check` | Same command if CI runs it | none | none |
| Source-specific leakage | `REQ-05`, `CHK-03`, `INV-05` | No dedicated suite | Run deterministic `rg` scan for known source-specific terms | `rg -n "alfagen|mercury|TASK-3446|rate-daemon|SlotDiagnostics|PositionAware|production log storm" memory-bank` | none known | Manual reviewer must confirm any allowed matches are only FT-012 provenance | none |
| Promotion trigger completeness | `REQ-06`, `CHK-04`, `CTR-02`, `INV-01` | No dedicated suite | Run text scan and manual review | `rg -n "workflow_profile|Promotion|promotion|contract_change|feature-package|epic-package|ADR" memory-bank/flows memory-bank/tasks` | none known | Manual reviewer validates semantics because text scan cannot prove correctness | none |

## Open Questions / Ambiguities

No unresolved blocking questions remain after `decision-log.md` entries `DL-001` and `DL-002`. New uncertainty that changes scope, selected design or evidence contract must stop execution and update the correct owner document first.

## Environment Contract

| Area | Contract | Used by | Failure symptom |
| --- | --- | --- | --- |
| setup | Local checkout at `feature/issue-12-compact-task-flow-bugfix-refactor-chore`; no app runtime required | `STEP-01`..`STEP-07` | Expected paths under `memory-bank/` are missing |
| test | Documentation verification uses Python script, `rg` and `git diff --check` | `CHK-01`, `CHK-03`, `CHK-04`, `CHK-05` | Command unavailable or returns non-zero without captured evidence |
| access / network / secrets | `gh` access was used only to read issue #12 and source docs; implementation should not require secrets | `STEP-01` | Need for private source data beyond listed docs means stop and human gate |

## Preconditions

| Precondition ID | Canonical ref | Required state | Used by steps | Blocks start |
| --- | --- | --- | --- | --- |
| `PRE-01` | `REQ-01`..`REQ-06` | `brief.md` is active and has Design Requirement Decision | `STEP-01`..`STEP-07` | yes |
| `PRE-02` | `SD-01`..`SD-04` | `design.md` is active and source scope decisions are recorded | `STEP-02`..`STEP-06` | yes |
| `PRE-03` | `CON-02` | `scripts/check_memory_bank_index.py` exists | `STEP-07` | yes |

## Workstreams

| Workstream | Implements | Result | Owner | Dependencies |
| --- | --- | --- | --- | --- |
| `WS-1` | `REQ-01`, `SOL-01`, `SOL-02` | New task-flow, bugfix-flow and refactor-flow docs | agent | `PRE-01`, `PRE-02` |
| `WS-2` | `REQ-02`, `REQ-03`, `SOL-03`, `SOL-04` | New task templates and `tasks/README.md` | agent | `WS-1` |
| `WS-3` | `REQ-04`, `SOL-05`, `SD-01` | Updated routing/index docs | agent | `WS-1`, `WS-2` |
| `WS-4` | `REQ-05`, `REQ-06`, `SOL-06`, `CTR-02`, `CTR-04` | Source sanitization and promotion trigger review | agent | `WS-1`, `WS-2`, `WS-3` |
| `WS-5` | `CHK-01`..`CHK-05` | Verification evidence and final acceptance state | agent | `WS-1`..`WS-4` |

## Approval Gates

No risky, irreversible or external side-effect actions are planned. If implementation discovers a need to publish, push, alter GitHub issue state or import private source artifacts not listed in issue #12, create `AG-01` and stop for human approval.

## Порядок работ

| Step ID | Actor | Implements | Goal | Touchpoints | Artifact | Verifies | Evidence IDs | Check command / procedure | Blocked by | Needs approval | Escalate if |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| `STEP-01` | agent | `REQ-01`, `SOL-01`, `SOL-02` | Add generic task-flow family docs | `memory-bank/flows/task-flow.md`, `bugfix-flow.md`, `refactor-flow.md` | governed flow docs | `CHK-02`, `CHK-04` | `EVID-02`, `EVID-04` | Manual review plus promotion scan | `PRE-01`, `PRE-02` | none | Source terms are needed to explain generic rules |
| `STEP-02` | agent | `REQ-02`, `SOL-03` | Add task templates | `memory-bank/flows/templates/task/` | `package-README.md`, `bugfix.md`, `refactor.md` | `CHK-01`, `CHK-02`, `CHK-03` | `EVID-01`, `EVID-02`, `EVID-03` | Link audit and source scan | `STEP-01` | none | Templates need `design.md` or `implementation-plan.md` inside `TASK-XXX/` |
| `STEP-03` | agent | `REQ-03`, `SOL-04`, `SD-04` | Add optional tasks destination index | `memory-bank/tasks/README.md` | tasks index | `CHK-01`, `CHK-03` | `EVID-01`, `EVID-03` | Link audit and source scan | `STEP-01` | none | Source example package seems required |
| `STEP-04` | agent | `REQ-04`, `SOL-05`, `SD-01` | Update routing/indexes | `memory-bank/flows/workflows.md`, `flows/README.md`, `flows/templates/README.md`, `memory-bank/README.md`, related indexes | updated index/routing docs | `CHK-01`, `CHK-04` | `EVID-01`, `EVID-04` | Link audit and promotion scan | `STEP-01`..`STEP-03` | none | Update conflicts with `feature-flow.md` ownership |
| `STEP-05` | agent | `REQ-05`, `SOL-06`, `CTR-04` | Remove source-specific leakage | all new/updated docs | sanitized docs | `CHK-03` | `EVID-03` | Source-specific `rg` scan | `STEP-01`..`STEP-04` | none | Any match outside FT-012 provenance remains |
| `STEP-06` | agent | `REQ-06`, `CTR-02`, `INV-01` | Review promotion triggers across selector, flows and templates | `memory-bank/flows/**`, `memory-bank/tasks/README.md` | promotion trigger verdict | `CHK-04` | `EVID-04` | Text scan plus manual verdict | `STEP-01`..`STEP-04` | none | Compact path can still hide capability/contract/high-risk/design work |
| `STEP-07` | agent | `EC-06` | Run final checks and capture evidence | `artifacts/ft-012/verify/**` | verification logs | `CHK-01`..`CHK-05` | `EVID-01`..`EVID-05` | Required commands from `brief.md` | `STEP-01`..`STEP-06` | none | Any required check fails |

## Parallelizable Work

- `PAR-01` `STEP-02` templates and `STEP-03` tasks index can be drafted in parallel after `STEP-01` establishes flow terminology.
- `PAR-02` `STEP-04` index updates should wait for actual target file paths to avoid broken links.
- `PAR-03` `STEP-05` and `STEP-06` can run as review passes after all docs exist.

## Checkpoints

| Checkpoint ID | Refs | Condition | Evidence IDs |
| --- | --- | --- | --- |
| `CP-01` | `STEP-01`, `STEP-02`, `STEP-03` | All new flow/template/task destination docs exist with frontmatter and no source examples | `EVID-02`, `EVID-03` |
| `CP-02` | `STEP-04`, `CHK-01` | All new docs reachable through index navigation | `EVID-01` |
| `CP-03` | `STEP-06`, `CHK-04` | Promotion triggers are present and consistent across selector, task-flow and profile templates | `EVID-04` |
| `CP-04` | `STEP-07`, `CHK-01`..`CHK-05` | Required checks have evidence and no blocking failures | `EVID-01`..`EVID-05` |

## Execution Risks

| Risk ID | Risk | Impact | Mitigation | Trigger |
| --- | --- | --- | --- | --- |
| `ER-01` | Source docs contain useful but out-of-scope artifacts | Scope creep | Keep `SD-02`; create follow-up instead of importing | Need for `workflow-decision-log.md` or metrics becomes material |
| `ER-02` | Link audit fails due new nested indexes | Feature cannot pass acceptance | Update parent indexes and annotations before completion | `CHK-01` non-zero |
| `ER-03` | Promotion triggers appear in only one doc | Compact flow becomes inconsistent | Repeat trigger language in selector, family flow, profile flows and templates | `CHK-04` manual review finds a missing trigger |
| `ER-04` | Sanitization scan matches FT-012 source provenance | False-positive review noise | Treat matches inside FT-012 docs as allowed provenance; target docs must remain clean | `CHK-03` finds matches |

## Stop Conditions / Fallback

| Stop ID | Related refs | Trigger | Immediate action | Safe fallback state |
| --- | --- | --- | --- | --- |
| `STOP-01` | `REQ-05`, `CTR-04`, `ER-01` | Implementation needs source project examples to be understandable | Stop and open human gate | Feature docs remain active; target docs not finalized |
| `STOP-02` | `REQ-06`, `INV-01`, `FM-01` | Compact task flow cannot be made safe without broader workflow policy | Stop and update `brief.md`/`design.md` or escalate | No target docs shipped as accepted |
| `STOP-03` | `CHK-01`, `CHK-05` | Required automated documentation checks fail and cannot be resolved locally | Stop with failing evidence | Package remains `planned` |

## Plan-local Evidence

| Evidence ID | Artifact | Producer | Path contract | Reused by checkpoints |
| --- | --- | --- | --- | --- |
| `EVID-09` | Review-improve cycle summary for feature docs | agent | Final response and/or `decision-log.md` entries | `CP-04` |

## Готово для приемки

- Все workstreams завершены или явно остановлены через `STOP-*`.
- Все checkpoints имеют evidence.
- Required local suites from `brief.md` are green, and CI is not known to contradict local verify.
- Manual-only gaps are limited to reviewer semantic verdicts already named in `CHK-02` and `CHK-04`.
- Support docs do not conflict with canonical `brief.md`, `design.md` or this plan.
- Финальная приемка идёт по `brief.md` `Verify`, а не по этому checklist.

## Completion Notes

- `WS-1` done: added `task-flow.md`, `bugfix-flow.md`, `refactor-flow.md`.
- `WS-2` done: added task templates and `memory-bank/tasks/README.md`.
- `WS-3` done: updated routing/index docs.
- `WS-4` done: source sanitization and promotion trigger review passed.
- `WS-5` done: evidence captured under `artifacts/ft-012/verify/`.

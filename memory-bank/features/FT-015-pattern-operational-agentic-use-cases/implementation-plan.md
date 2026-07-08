---
title: "FT-015: Implementation Plan"
doc_kind: feature
doc_function: derived
purpose: "Execution plan for FT-015. Captures discovery context, steps, risks, and test strategy without redefining scope or acceptance."
derived_from:
  - brief.md
status: archived
audience: humans_and_agents
must_not_define:
  - ft_015_scope
  - ft_015_selected_design
  - ft_015_acceptance_criteria
  - ft_015_blocker_state
---

# План имплементации

## Цель текущего плана

Выполнить Markdown-only изменение по `brief.md`: обновить
`memory-bank/use-cases/README.md` generic guidance для operational/agentic use
cases, не добавляя project-specific source content и не меняя `UC-XXX`
template без обнаруженного конфликта.

## Grounding / Support References

| Document | Role in this plan | Facts reused | Conflict action |
| --- | --- | --- | --- |
| `brief.md` | canonical problem / verify owner | `REQ-*`, `NS-*`, `SC-*`, `CHK-*`, `EVID-*` | Update `brief.md` first |
| `design.md` / `none` | conditional solution owner | `Design required: no` | Promote new design facts before planning |
| `decision-log.md` | feature-local decision log | `DL-*` FPF decisions for owner boundaries and source abstraction | Update `decision-log.md` if a new material decision is closed |

## Current State / Reference Points

| Path / module | Current role | Why relevant | Reuse / mirror |
| --- | --- | --- | --- |
| `memory-bank/use-cases/README.md` | Use case index and guidance | Primary change surface for `REQ-01` through `REQ-03` | Preserve current UC-vs-SC distinction and registry format |
| `memory-bank/flows/templates/use-case/UC-XXX.md` | Canonical use case template | Must remain compatible with new guidance per `REQ-04` | Reuse existing trigger, preconditions, main flow, alternate flows, postconditions, rules, and traceability shape |
| `memory-bank/flows/feature-flow.md` | Feature package lifecycle and ID taxonomy | Defines document boundaries, design gate, and verify requirements | Keep `brief.md` as problem/verify owner and plan as execution owner |
| GitHub issue #15 | Source task | Defines scope, source examples, and acceptance constraints | Use as external source of task intent only |
| Downstream source use cases named in issue #15 | Pattern evidence | Shows operational/agentic use cases in a downstream project | Abstract pattern only; do not copy project names, tool names, commands, or runtime specifics |

## Test Strategy

| Test surface | Canonical refs | Existing coverage | Planned automated coverage | Required local suites / commands | Required CI suites / jobs | Manual-only gap / justification | Manual-only approval ref |
| --- | --- | --- | --- | --- | --- | --- | --- |
| README guidance content | `REQ-01`, `REQ-02`, `REQ-03`, `SC-01`, `SC-02`, `CHK-01` | No explicit operational/agentic section | `rg` coverage check for generic terms and routing language | `rg -n "Operational|agentic|machine-readable|recovery|postconditions|SC-\\*" memory-bank/use-cases/README.md` | repository default doc checks if configured | none | `none` |
| Genericity guard | `NS-02`, `NEG-01`, `CHK-02` | No source content currently present | Negative `rg` check for source-specific terms listed in issue #15 | Run forbidden-term `rg` against `memory-bank/use-cases/README.md` | repository default doc checks if configured | none | `none` |
| Governance / index health | `REQ-04`, `CHK-03` | `scripts/check_memory_bank_index.py` exists | Existing memory-bank index audit and diff whitespace check | `python3 scripts/check_memory_bank_index.py`; `git diff --check` | repository default doc checks if configured | none; semantic compatibility review is the documented `CHK-03` procedure | `none` |

## Open Questions / Ambiguities

| Open Question ID | Question | Why unresolved | Blocks | Default action / escalation owner |
| --- | --- | --- | --- | --- |
| `OQ-00` | None currently open. | Existing facts are sufficient for this Markdown-only guidance change. | none | Reopen if implementation reveals a conflict between README guidance and `UC-XXX` template. |

## Environment Contract

| Area | Contract | Used by | Failure symptom |
| --- | --- | --- | --- |
| setup | Work from repository root with shell access. | `STEP-01` through `STEP-04` | Paths under `memory-bank/` cannot be read or edited. |
| test | Use `rg`, `python3 scripts/check_memory_bank_index.py`, and `git diff --check`. | `CHK-01`, `CHK-02`, `CHK-03` | Missing tool, failing index audit, or dirty whitespace/conflict markers. |
| access / network / secrets | No secrets required. Network is useful only for reading issue/source examples; implementation itself is local. | `STEP-01` | If external source is unavailable, continue from issue facts already captured in `brief.md`; do not invent additional source facts. |

## Preconditions

| Precondition ID | Canonical ref | Required state | Used by steps | Blocks start |
| --- | --- | --- | --- | --- |
| `PRE-01` | `brief.md` Design Requirement Decision | `Design required: no` and `status: active` | `STEP-01` through `STEP-04` | yes |
| `PRE-02` | `NS-02`, `CON-01` | Source-specific details remain excluded from generic docs | `STEP-02`, `STEP-03` | yes |

## Workstreams

| Workstream | Implements | Result | Owner | Dependencies |
| --- | --- | --- | --- | --- |
| `WS-1` | `REQ-01`, `REQ-02`, `REQ-03` | Updated `memory-bank/use-cases/README.md` guidance | agent | `PRE-01`, `PRE-02` |
| `WS-2` | `REQ-04` | Compatibility verdict for `UC-XXX` template | agent | `WS-1` |
| `WS-3` | `CHK-01`, `CHK-02`, `CHK-03` | Local verification evidence | agent | `WS-1`, `WS-2` |

## Approval Gates

| Approval Gate ID | Trigger | Applies to | Why approval is required | Approver / evidence |
| --- | --- | --- | --- | --- |
| `AG-00` | None. | none | The planned change is local Markdown guidance with no external side effects. | `none` |

## Порядок работ

| Step ID | Actor | Implements | Goal | Touchpoints | Artifact | Verifies | Evidence IDs | Check command / procedure | Blocked by | Needs approval | Escalate if |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| `STEP-01` | agent | `REQ-01`, `REQ-02`, `REQ-03` | Add generic operational/agentic use case guidance and UC-vs-SC routing. | `memory-bank/use-cases/README.md` | README diff | `CHK-01` | `EVID-01` | `rg -n "Operational|agentic|machine-readable|recovery|postconditions|SC-\\*" memory-bank/use-cases/README.md` | `PRE-01`, `PRE-02` | `none` | Required guidance cannot be stated without project-specific terms. |
| `STEP-02` | agent | `NS-02`, `NEG-01` | Remove any source-specific leakage. | `memory-bank/use-cases/README.md` | README diff | `CHK-02` | `EVID-02` | Run forbidden-term `rg` for source-specific project/tool names and command fragments listed in issue #15 against `memory-bank/use-cases/README.md`. | `STEP-01` | `none` | Source-specific detail appears necessary for acceptance. |
| `STEP-03` | agent | `REQ-04` | Compare updated README guidance with `UC-XXX` template. | `memory-bank/use-cases/README.md`, `memory-bank/flows/templates/use-case/UC-XXX.md` | Compatibility note | `CHK-03` | `EVID-03` | Semantic review plus `python3 scripts/check_memory_bank_index.py` | `STEP-01` | `none` | Template conflict is found and cannot be resolved by README wording. |
| `STEP-04` | agent | `CHK-03` | Run repository hygiene checks. | repository root | Check outputs | `CHK-03` | `EVID-03` | `python3 scripts/check_memory_bank_index.py`; `git diff --check` | `STEP-03` | `none` | Index audit or diff check fails for a reason outside FT-015 scope. |

## Parallelizable Work

- `PAR-01` `STEP-01` and `STEP-02` are sequential because the leakage check depends on final README wording.
- `PAR-02` `STEP-03` and `STEP-04` can be run after `STEP-01`, but final acceptance needs both results.

## Checkpoints

| Checkpoint ID | Refs | Condition | Evidence IDs |
| --- | --- | --- | --- |
| `CP-01` | `STEP-01`, `CHK-01` | README contains operational/agentic guidance and UC-vs-SC routing. | `EVID-01` |
| `CP-02` | `STEP-02`, `CHK-02` | README contains no source-specific terms from `NS-02`. | `EVID-02` |
| `CP-03` | `STEP-03`, `STEP-04`, `CHK-03` | Template compatibility reviewed and local hygiene checks pass. | `EVID-03` |

## Execution Risks

| Risk ID | Risk | Impact | Mitigation | Trigger |
| --- | --- | --- | --- | --- |
| `ER-01` | Guidance becomes too broad and implies all operational checks deserve `UC-*`. | Violates `NEG-02` and weakens UC-vs-SC boundary. | Keep repetition, project-level ownership, and reusable pre/postcondition criteria explicit. | README text lacks a "keep in `SC-*` when local" rule. |
| `ER-02` | Source examples leak downstream terms. | Violates `NS-02` and issue acceptance. | Run forbidden-term check and edit wording back to generic concepts. | `CHK-02` finds a match. |

## Stop Conditions / Fallback

| Stop ID | Related refs | Trigger | Immediate action | Safe fallback state |
| --- | --- | --- | --- | --- |
| `STOP-01` | `NS-02`, `NEG-01` | Acceptance appears to require source-specific source content. | Stop and create a human gate with facts and options. | Keep feature docs active; do not edit generic README further. |
| `STOP-02` | `REQ-04`, `NS-03` | `UC-XXX` template conflict is found and cannot be fixed by README wording. | Stop and create a human gate before changing the template. | README guidance draft remains unmerged until owner decision. |

## Plan-local Evidence

| Evidence ID | Artifact | Producer | Path contract | Reused by checkpoints |
| --- | --- | --- | --- | --- |
| `EVID-09` | Review-improve report for feature documents | agent | `memory-bank/features/FT-015-pattern-operational-agentic-use-cases/feature-review-report.md` | `CP-01`, `CP-02`, `CP-03` |

## Готово для приемки

- Все workstreams завершены или явно остановлены через `STOP-*`.
- Все checkpoints имеют evidence.
- Required local commands from `Test Strategy` pass.
- Manual-only gaps are limited to semantic template compatibility review and have a written verdict.
- Final acceptance is evaluated against `brief.md` `Verify`, not this checklist.

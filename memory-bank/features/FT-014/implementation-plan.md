---
title: "FT-014: Implementation Plan"
doc_kind: feature
doc_function: derived
purpose: "Execution-план реализации FT-014. Фиксирует discovery context, шаги, риски и test strategy без переопределения canonical problem или solution facts."
derived_from:
  - brief.md
  - design.md
status: active
audience: humans_and_agents
must_not_define:
  - ft_014_scope
  - ft_014_selected_design
  - ft_014_acceptance_criteria
  - ft_014_blocker_state
---

# FT-014: Implementation Plan

## Цель текущего плана

Реализовать issue 14 как documentation/template delivery: добавить lightweight epic intake brief template, обновить связанные indexes and `epic-flow.md`, затем подтвердить acceptance через link audit, whitespace check и targeted source-leakage check.

## Grounding / Support References

| Document | Role in this plan | Facts reused | Conflict action |
| --- | --- | --- | --- |
| `brief.md` | canonical problem / verify owner | `REQ-*`, `SC-*`, `NEG-*`, `CHK-*`, `EVID-*` | Update `brief.md` first |
| `design.md` | conditional solution owner | `SOL-*`, `C4-*`, `SD-*`, `CTR-*`, `INV-*`, `FM-*`, `RB-*` | Update `design.md` first |
| `decision-log.md` | feature-local FPF reasoning ledger | `DL-*` decisions and rationale | Update `decision-log.md` when closing blocking questions |

## Current State / Reference Points

| Path / module | Current role | Why relevant | Reuse / mirror |
| --- | --- | --- | --- |
| `memory-bank/flows/templates/epic/README.md` | Current epic template index | Must route to new `brief.md` template | Keep annotated bullet style |
| `memory-bank/flows/templates/README.md` | Global templates index | Must keep all templates reachable | Add derived_from and bullet for epic brief |
| `memory-bank/flows/epic-flow.md` | Canonical epic lifecycle governance | Owns package rules, layer model, gates and boundary rules | Add intake-only semantics without changing feature execution ownership |
| `memory-bank/flows/templates/epic/charter.md` | Existing full epic intent template | Shows current frontmatter/body wrapper style | Keep full package canonical ownership separate |
| `scripts/check_memory_bank_index.py` | Link/reachability/index audit | Acceptance requires index audit pass | Run after changes |

## Test Strategy

| Test surface | Canonical refs | Existing coverage | Planned automated coverage | Required local suites / commands | Required CI suites / jobs | Manual-only gap / justification | Manual-only approval ref |
| --- | --- | --- | --- | --- | --- | --- | --- |
| Memory-bank link/index integrity | `REQ-06`, `SC-01`, `CHK-01` | `scripts/check_memory_bank_index.py` exists and currently passes | Re-run after docs changes | `python3 scripts/check_memory_bank_index.py` | Same audit if CI has it | none | `none` |
| Source-project leakage in target docs | `REQ-05`, `NEG-01`, `CHK-03` | No dedicated test | Targeted `rg` check over changed template/governance docs | `rg -n "zelma|Zelma|dapi/zelma" ...` | none known | none | `none` |
| Diff hygiene | `REQ-06`, `CHK-04` | Git built-in check | Run whitespace/conflict-marker check | `git diff --check` | Same if CI has it | none | `none` |
| Semantic boundary review | `REQ-01`, `REQ-04`, `SC-02`, `CHK-02` | Manual documentation review | Review changed docs against issue acceptance and `design.md` invariants | Manual review plus grep/diff context | none known | Manual review is appropriate for prose semantics; backed by explicit `CHK-02` criteria | `none` |

## Open Questions / Ambiguities

| Open Question ID | Question | Why unresolved | Blocks | Default action / escalation owner |
| --- | --- | --- | --- | --- |
| `OQ-00` | None currently blocking execution | Issue scope, source template and current governance docs are sufficient | none | Continue; create human gate only if review finds materially ambiguous ownership |

## Environment Contract

| Area | Contract | Used by | Failure symptom |
| --- | --- | --- | --- |
| setup | Repo checkout at `feature/issue-14-epic-intake-brief-template` | `STEP-01`-`STEP-05` | Files or script paths missing |
| test | Python 3 and ripgrep available | `STEP-05`, `CHK-01`, `CHK-03` | Audit or targeted search cannot run |
| access / network / secrets | GitHub issue and source template already read through authenticated `gh` | `PRE-01` | Source evidence unavailable or contradictory |

## Preconditions

| Precondition ID | Canonical ref | Required state | Used by steps | Blocks start |
| --- | --- | --- | --- | --- |
| `PRE-01` | `ASM-01`, `CON-03`, `SD-01` | Issue 14 and source template inspected; target docs must stay generic | `STEP-01`, `STEP-02`, `STEP-03` | yes |
| `PRE-02` | `C4-00`, `SD-03` | No C4 artifact or ADR required | `STEP-02`, `STEP-03` | no |

## Workstreams

| Workstream | Implements | Result | Owner | Dependencies |
| --- | --- | --- | --- | --- |
| `WS-1` | `REQ-01`, `SOL-01`, `CTR-01` | New epic intake brief template | agent | `PRE-01` |
| `WS-2` | `REQ-02`, `REQ-03`, `SOL-02`, `CTR-03` | Updated local and global template indexes | agent | `WS-1` |
| `WS-3` | `REQ-04`, `SOL-03`, `CTR-02` | Updated `epic-flow.md` boundary and layer semantics | agent | `PRE-01` |
| `WS-4` | `REQ-05`, `REQ-06`, `INV-01`, `FM-02`, `FM-03` | Verification evidence and review report | agent | `WS-1`, `WS-2`, `WS-3` |

## Approval Gates

| Approval Gate ID | Trigger | Applies to | Why approval is required | Approver / evidence |
| --- | --- | --- | --- | --- |
| `AG-00` | No risky external or irreversible action planned | none | Local Markdown edits and local checks do not require approval | `none` |

## Порядок работ

| Step ID | Actor | Implements | Goal | Touchpoints | Artifact | Verifies | Evidence IDs | Check command / procedure | Blocked by | Needs approval | Escalate if |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| `STEP-01` | agent | `REQ-01`, `SOL-01`, `CTR-01` | Add generic epic intake brief wrapper-template | `memory-bank/flows/templates/epic/brief.md` | New template file | `CHK-02`, `CHK-03` | `EVID-02`, `EVID-03` | Review file content and targeted `rg` | `PRE-01` | `none` | Source-specific content is required to satisfy acceptance |
| `STEP-02` | agent | `REQ-02`, `REQ-03`, `SOL-02`, `CTR-03` | Register new template in indexes | `memory-bank/flows/templates/epic/README.md`, `memory-bank/flows/templates/README.md` | Updated indexes | `CHK-01`, `CHK-02` | `EVID-01`, `EVID-02` | Link audit and review | `STEP-01` | `none` | Index audit requires unrelated index churn |
| `STEP-03` | agent | `REQ-04`, `SOL-03`, `CTR-02` | Clarify epic intake boundaries | `memory-bank/flows/epic-flow.md` | Updated flow doc | `CHK-02` | `EVID-02` | Review boundary wording | `PRE-01` | `none` | Brief must become authoritative for roadmap/subissues/risks |
| `STEP-04` | agent | `REQ-05`, `SOL-04`, `INV-01` | Remove or prevent source-specific leakage | Changed target docs | Clean target docs | `CHK-03` | `EVID-03` | Targeted `rg` | `STEP-01`, `STEP-03` | `none` | Source-specific term cannot be removed without losing acceptance |
| `STEP-05` | agent | `REQ-06`, `RB-02` | Run final checks and record results | Repo root | Check outputs and review report | `CHK-01`, `CHK-04` | `EVID-01`, `EVID-04` | `python3 scripts/check_memory_bank_index.py`; `git diff --check` | `STEP-01`-`STEP-04` | `none` | Audit fails due unrelated pre-existing issues |

## Parallelizable Work

- `PAR-01` `STEP-01` and `STEP-03` can be drafted independently after `PRE-01`, but final wording must be reconciled before checks.
- `PAR-02` Verification steps should run after all writes, because index reachability depends on the complete document graph.

## Checkpoints

| Checkpoint ID | Refs | Condition | Evidence IDs |
| --- | --- | --- | --- |
| `CP-01` | `STEP-01`, `STEP-02`, `STEP-03`, `CHK-02` | Changed docs satisfy issue acceptance and design invariants | `EVID-02` |
| `CP-02` | `STEP-04`, `CHK-03` | No source-project terms in target template/governance docs | `EVID-03` |
| `CP-03` | `STEP-05`, `CHK-01`, `CHK-04` | Link audit and diff hygiene pass | `EVID-01`, `EVID-04` |

## Execution Risks

| Risk ID | Risk | Impact | Mitigation | Trigger |
| --- | --- | --- | --- | --- |
| `ER-01` | Intake brief duplicates charter ownership | Confusing epic lifecycle | Keep `SD-01`, `CTR-02`, `INV-03` explicit in template and flow | Review finds brief owns roadmap or accepted subissues |
| `ER-02` | New files become unreachable | Index audit failure | Update parent indexes before final audit | `CHK-01` fails with orphan/unreachable |
| `ER-03` | Source-specific terms leak into generic docs | Violates issue scope | Run targeted `rg` over target docs | `CHK-03` finds matches |

## Stop Conditions / Fallback

| Stop ID | Related refs | Trigger | Immediate action | Safe fallback state |
| --- | --- | --- | --- | --- |
| `STOP-01` | `CON-01`, `SD-01`, `FM-01` | Review shows a material ambiguity over whether brief replaces charter/roadmap | Stop and create human gate | Keep current branch unmerged with findings recorded |
| `STOP-02` | `CON-03`, `FM-02` | Required acceptance appears impossible without source-specific content | Stop and create human gate | No further template edits |

## Plan-local Evidence

| Evidence ID | Artifact | Producer | Path contract | Reused by checkpoints |
| --- | --- | --- | --- | --- |
| `EVID-09` | Review-improve report for feature docs | implementer / reviewer | `memory-bank/features/FT-014/feature-review-report.md` | `CP-01`, `CP-03` |

## Готово для приемки

- All workstreams are complete or explicitly stopped through `STOP-*`.
- All checkpoints have evidence.
- Required local suites are green, and CI does not contradict local verify.
- Manual-only semantic review for `CHK-02` is recorded in `feature-review-report.md` or final delivery notes.
- Final acceptance follows `brief.md` `Verify`, not this checklist.

---
title: "FT-013: Implementation Plan"
doc_kind: feature
doc_function: derived
purpose: "Execution-план реализации issue 13. Фиксирует discovery context, шаги, риски и test strategy без переопределения canonical problem и solution facts."
derived_from:
  - brief.md
  - design.md
status: active
audience: humans_and_agents
must_not_define:
  - ft_013_scope
  - ft_013_selected_design
  - ft_013_acceptance_criteria
  - ft_013_blocker_state
---
# FT-013: Implementation Plan

## Цель текущего плана

Реализовать issue 13 в generic memory-bank template: добавить optional workflow decision log, workflow metrics и workflow routing developer brief, обновить routing navigation, проверить link integrity и отсутствие source project leakage.

## Grounding / Support References

| Document | Role in this plan | Facts reused | Conflict action |
| --- | --- | --- | --- |
| `brief.md` | canonical problem / verify owner | `REQ-*`, `NS-*`, `ASM-*`, `CON-*`, `SC-*`, `CHK-*`, `EVID-*` | Update `brief.md` first |
| `design.md` | solution owner | `SOL-*`, `C4-*`, `SD-*`, `CTR-*`, `INV-*`, `FM-*`, `RB-*` | Update `design.md` first |
| `decision-log.md` | decision provenance | `DL-FT013-*` FPF decisions | Update `decision-log.md` and then owner docs |
| `../../flows/workflows.md` | current routing owner | Existing workflow types and routing rules | Update only routing pointers in FT-013 scope |
| `../../flows/README.md` | flows navigation owner | Existing annotated index style | Add optional docs with annotations |
| [GitHub issue 12](https://github.com/dapi/memory-bank/issues/12) | adjacent planned task-flow owner | Task-flow dependency and non-scope boundary | Do not implement issue 12 in FT-013 |
| Source docs from `alfagen/mercury` | pattern evidence | Decision log / metrics / developer brief structures | Genericize; do not copy project-specific facts |

## Current State / Reference Points

| Path / module | Current role | Why relevant | Reuse / mirror |
| --- | --- | --- | --- |
| `memory-bank/flows/workflows.md` | Canonical routing overview | FT-013 adds optional selector rationale/metrics/brief around this doc | Reuse concise routing language and autonomy gradient |
| `memory-bank/flows/README.md` | Index for reusable flow docs | Acceptance requires optional docs reachable from flows README | Add annotated links matching existing style |
| `memory-bank/flows/feature-flow.md` | Feature lifecycle owner | Governs this package and boundary between brief/design/plan | Keep unchanged |
| `memory-bank/flows/epic-flow.md` | Existing decision-log pattern at epic level | Shows local decision log concept and evidence-backed decision language | Mirror decision provenance style, not epic scope |
| `memory-bank/engineering/testing-policy.md` | Test/evidence policy | Defines automated checks/manual-only expectations | Use lightweight doc checks as deterministic verification |
| `scripts/check_memory_bank_index.py` | Link/index audit | Primary automated validation for new markdown docs | Run as `CHK-01` |
| `memory-bank/flows/templates/README.md` | Templates index | Relevant only if implementation adds files under `flows/templates/` | Do not update unless design changes |

## Test Strategy

| Test surface | Canonical refs | Existing coverage | Planned automated coverage | Required local suites / commands | Required CI suites / jobs | Manual-only gap / justification | Manual-only approval ref |
| --- | --- | --- | --- | --- | --- | --- | --- |
| Markdown navigation and frontmatter dependencies | `REQ-01`..`REQ-05`, `CHK-01` | `scripts/check_memory_bank_index.py` audits links, reachability and expected README indices | Run existing script after docs and index updates | `python3 scripts/check_memory_bank_index.py` | Same script if CI runs repository checks | none | none |
| Whitespace / conflict markers | `REQ-06`, `CHK-02` | Git has built-in diff check | Run diff check before handoff | `git diff --check` | Same command if CI configured | none | none |
| Generic source leakage | `REQ-06`, `CHK-03` | No dedicated script | Run targeted `rg` scan against delivered workflow docs | `rg -n "Mercury|alfagen|2026-06-29|2026-07-28|2026-07-29|dip|Rails|Rake|RSpec|MySQL" memory-bank/flows/workflow-*.md` | none unless added later | Manual judgment still needed for subtle project-specific prose | none |
| Acceptance traceability | `REQ-01`..`REQ-06`, `CHK-04` | Manual review by implementer/reviewer | Prepare PR/body note or review note mapping issue 13 acceptance to docs | Manual review procedure from `brief.md` | Reviewer approval / PR review | Manual by nature because it checks semantic fit | none |

## Open Questions / Ambiguities

No unresolved questions block execution. The task-flow dependency, design-required decision and decision-log ownership correction are closed in [decision-log.md](decision-log.md).

## Environment Contract

| Area | Contract | Used by | Failure symptom |
| --- | --- | --- | --- |
| setup | Repository checkout with current branch and Python 3 | `STEP-01`..`STEP-06` | `scripts/check_memory_bank_index.py` cannot run |
| test | Lightweight repository checks from `brief.md` are authoritative for this docs-only feature | `STEP-05`, `STEP-06` | Link audit, diff check or leak scan fails |
| access / network / secrets | GitHub issue/source reading was needed for planning; implementation itself should not require secrets or live external systems | `STEP-01` | Missing source access should not block because selected design is already captured |

## Preconditions

| Precondition ID | Canonical ref | Required state | Used by steps | Blocks start |
| --- | --- | --- | --- | --- |
| `PRE-01` | `brief.md` Design Requirement Decision | `Design required: yes` and `brief.md` active | `STEP-01`..`STEP-06` | yes |
| `PRE-02` | `design.md` `C4-00`, `SOL-*`, `SD-*` | `design.md` active and no required C4 artifact | `STEP-01`..`STEP-06` | yes |
| `PRE-03` | `SD-03`, `INV-03` | Issue 12 remains outside implementation scope; no links to absent local task-flow docs | `STEP-02`, `STEP-03`, `STEP-04` | yes |

## Workstreams

| Workstream | Implements | Result | Owner | Dependencies |
| --- | --- | --- | --- | --- |
| `WS-1` | `REQ-01`, `SOL-01` | `workflow-decision-log.md` added | agent | `PRE-01`, `PRE-02` |
| `WS-2` | `REQ-02`, `SOL-02` | `workflow-metrics.md` added | agent | `PRE-01`, `PRE-02` |
| `WS-3` | `REQ-03`, `SOL-03` | `workflow-routing-developer-brief.md` added | agent | `PRE-01`, `PRE-02`, `PRE-03` |
| `WS-4` | `REQ-04`, `SOL-04`, `SOL-05` | `flows/README.md` and `workflows.md` updated | agent | `WS-1`, `WS-2`, `WS-3` |
| `WS-5` | `REQ-06`, `CHK-01`..`CHK-04` | Verification outputs and review notes ready | agent / reviewer | `WS-1`..`WS-4` |

## Approval Gates

No risky, irreversible, costly or external-effect action is planned for implementation. Human review is still expected before merge through normal PR review.

## Порядок работ

| Step ID | Actor | Implements | Goal | Touchpoints | Artifact | Verifies | Evidence IDs | Check command / procedure | Blocked by | Needs approval | Escalate if |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| `STEP-01` | agent | `REQ-01`, `SOL-01` | Add generic workflow decision log | `memory-bank/flows/workflow-decision-log.md` | New governed markdown doc | `CHK-01`, `CHK-03`, `CHK-04` | `EVID-01`, `EVID-03`, `EVID-04` | Review frontmatter, links, genericity | `PRE-01`, `PRE-02` | none | Decision log needs source-specific facts to make sense |
| `STEP-02` | agent | `REQ-02`, `SOL-02`, `SD-02` | Add workflow metrics template | `memory-bank/flows/workflow-metrics.md` | New governed markdown doc | `CHK-01`, `CHK-03`, `CHK-04` | `EVID-01`, `EVID-03`, `EVID-04` | Review safety-first decision rule and metric cards | `PRE-01`, `PRE-02` | none | Speed can be read as overriding safety |
| `STEP-03` | agent | `REQ-03`, `SOL-03`, `SD-03`, `SD-04` | Add developer brief | `memory-bank/flows/workflow-routing-developer-brief.md` | New guide/template doc | `CHK-01`, `CHK-03`, `CHK-04` | `EVID-01`, `EVID-03`, `EVID-04` | Review no links to absent task-flow docs | `PRE-03` | none | Brief requires issue 12 artifacts to be truthful |
| `STEP-04` | agent | `REQ-04`, `SOL-04`, `SOL-05` | Make optional docs reachable | `memory-bank/flows/README.md`, `memory-bank/flows/workflows.md` | Updated navigation/routing text | `CHK-01`, `CHK-04` | `EVID-01`, `EVID-04` | Run index audit | `WS-1`, `WS-2`, `WS-3` | none | Link audit points to absent docs |
| `STEP-05` | agent | `REQ-06` | Run deterministic checks | repo root | Command outputs | `CHK-01`, `CHK-02`, `CHK-03` | `EVID-01`, `EVID-02`, `EVID-03` | Commands from `brief.md` checks | `STEP-01`..`STEP-04` | none | Any required check fails |
| `STEP-06` | agent / reviewer | `REQ-01`..`REQ-06` | Final semantic acceptance review | New/changed docs | Review note / PR body | `CHK-04` | `EVID-04` | Map issue 13 acceptance to changed docs | `STEP-05` | none | Scope drift into issue 12 is found |

## Parallelizable Work

- `PAR-01` `WS-1`, `WS-2` and `WS-3` can be drafted independently after `design.md` is active.
- `PAR-02` `WS-4` must wait for the new docs to exist to avoid broken navigation links.
- `PAR-03` Verification in `WS-5` must run after all docs/navigation updates.

## Checkpoints

| Checkpoint ID | Refs | Condition | Evidence IDs |
| --- | --- | --- | --- |
| `CP-01` | `STEP-01`, `STEP-02`, `STEP-03`, `CTR-01` | All new workflow docs exist with valid frontmatter and generic wording | `EVID-04` |
| `CP-02` | `STEP-04`, `CTR-02`, `CTR-03` | Flows navigation reaches optional docs without absent task-flow links | `EVID-01` |
| `CP-03` | `STEP-05`, `CHK-01`, `CHK-02`, `CHK-03` | Required local checks pass | `EVID-01`, `EVID-02`, `EVID-03` |

## Execution Risks

| Risk ID | Risk | Impact | Mitigation | Trigger |
| --- | --- | --- | --- | --- |
| `ER-01` | Over-copying source docs | Generic template gets source-specific dates/names/commands | Use `CTR-04` and `CHK-03`; write placeholders and generic decision windows | `CHK-03` matches or review finds source-specific prose |
| `ER-02` | Accidental scope expansion into issue 12 | FT-013 becomes too large and conflicts with planned task-flow feature | Keep `NS-01`, `SD-03`, `INV-03`; avoid local links to absent docs | Implementation starts creating task-flow/task templates |
| `ER-03` | Metrics look precise but are not adaptable | Downstream projects treat example thresholds/dates as universal | Mark windows/targets as project-adapted fields, not fixed defaults | Metrics doc includes hard-coded pilot dates |

## Stop Conditions / Fallback

| Stop ID | Related refs | Trigger | Immediate action | Safe fallback state |
| --- | --- | --- | --- | --- |
| `STOP-01` | `SD-03`, `FM-01` | Implementation requires local task-flow docs to avoid misleading readers | Stop and raise human gate or wait for issue 12 | FT-013 docs remain planned; no speculative broken links |
| `STOP-02` | `CTR-04`, `FM-02` | Source-specific content cannot be safely generalized | Stop and ask for human decision on acceptable generic wording | New workflow docs not merged |
| `STOP-03` | `CHK-01` | Link audit fails for reasons outside FT-013 scope | Fix local FT-013 links if possible; otherwise escalate | Changes remain unshipped |

## Plan-local Evidence

| Evidence ID | Artifact | Producer | Path contract | Reused by checkpoints |
| --- | --- | --- | --- | --- |
| `EVID-09` | Review-improve report for feature-doc package | implementer | `decision-log.md` entries and final response summary | `CP-01`, `CP-02` |

## Готово для приемки

- All workstreams are complete or explicitly stopped through `STOP-*`.
- `workflow-decision-log.md`, `workflow-metrics.md` and `workflow-routing-developer-brief.md` exist and are reachable.
- Required local checks from `brief.md` pass.
- Generic leak scan has no matches in delivered workflow docs.
- Manual acceptance review maps issue 13 acceptance to changed docs.
- Any future task-flow links are left to issue 12 or a follow-up after issue 12 lands.

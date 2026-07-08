---
title: "FT-016: Implementation Plan"
doc_kind: feature
doc_function: derived
purpose: "Execution-план реализации FT-016. Фиксирует discovery context, шаги, риски и test strategy без переопределения canonical problem или solution facts."
derived_from:
  - brief.md
  - design.md
  - decision-log.md
status: archived
audience: humans_and_agents
must_not_define:
  - ft_016_scope
  - ft_016_selected_design
  - ft_016_acceptance_criteria
  - ft_016_blocker_state
---

# FT-016: Implementation Plan

## Цель текущего плана

Выполнить documentation-only feature из issue 16: добавить generic privacy/source-boundary support template, подключить его к feature-flow/template routing и собрать evidence по canonical checks из `brief.md`.

## Grounding / Support References

| Document | Role in this plan | Facts reused | Conflict action |
| --- | --- | --- | --- |
| `brief.md` | canonical problem / verify owner | `REQ-*`, `NS-*`, `SC-*`, `NEG-*`, `CHK-*`, `EVID-*` | Update `brief.md` first |
| `design.md` | solution owner | `SOL-*`, `C4-*`, `SD-*`, `CTR-*`, `INV-*`, `FM-*`, `RB-*` | Update `design.md` or decision log first |
| `decision-log.md` | FPF decision ledger | `DL-*` decisions and consequences | Update decision log if a blocking question is resolved |
| `memory-bank/flows/templates/feature/support/runtime-surfaces.md` | local support-template reference pattern | wrapper/instantiated frontmatter pattern and non-ownership language | Mirror structure, not runtime-surface semantics |
| `memory-bank/flows/templates/feature/support/use-cases.md` | local support-template reference pattern | companion-doc role and traceability wording | Mirror support boundary wording |

## Current State / Reference Points

| Path / module | Current role | Why relevant | Reuse / mirror |
| --- | --- | --- | --- |
| `memory-bank/flows/feature-flow.md` | canonical feature flow and support doc rules | New support doc trigger and ownership limits belong here | Add one optional support-doc row and support ID taxonomy |
| `memory-bank/flows/templates/README.md` | template index | New template must be discoverable | Add index route |
| `memory-bank/flows/templates/feature/README.md` | feature package README template | Downstream routes mention support docs by lifecycle | Add privacy/source-boundary route when support docs are listed |
| `memory-bank/flows/templates/feature/brief.md` | canonical feature brief template | Design/support decision wording may need source-boundary trigger guidance | Add guidance without moving solution facts into brief |
| `memory-bank/flows/templates/feature/design.md` | feature design template | Design packs may include privacy/source-boundary support refs | Add guidance without making support doc solution owner |
| `memory-bank/flows/templates/feature/implementation-plan.md` | execution plan template | Plan grounding may import support refs and manual-only gaps | Add optional support ref and test strategy guidance |
| `memory-bank/flows/templates/feature/support/` | support template directory | Target location for new template | Follow wrapper-template structure |
| `memory-bank/features/README.md` | feature package index | New package should be discoverable | Add `FT-016` route |
| `scripts/check_memory_bank_index.py` | local link/index audit | Required lightweight repo check | Run after edits |

## Test Strategy

| Test surface | Canonical refs | Existing coverage | Planned automated coverage | Required local suites / commands | Required CI suites / jobs | Manual-only gap / justification | Manual-only approval ref |
| --- | --- | --- | --- | --- | --- | --- | --- |
| Template/index discoverability | `REQ-01`, `REQ-05`, `SC-01`, `CHK-01`, `CHK-05`, `SOL-01`, `SOL-04` | `check_memory_bank_index.py` audits existing memory-bank index/link expectations | Update indexes and run audit | `python3 scripts/check_memory_bank_index.py`; `rg -n "privacy-source-boundary|Privacy / Source Boundary|source-boundary" memory-bank/flows memory-bank/features/README.md` | Same audit/check commands if CI exists | none | `none` |
| Template content coverage | `REQ-02`, `SC-02`, `NEG-01`, `NEG-02`, `CHK-02`, `SOL-01`, `SD-03`, `SD-04` | No automated semantic checker | Manual structured review against required sections | `sed -n '1,260p' memory-bank/flows/templates/feature/support/privacy-source-boundary.md` | n/a | Manual review is acceptable because repo has no semantic template linter | `none` |
| Routing and owner boundaries | `REQ-03`, `SC-03`, `CHK-03`, `SOL-02`, `SOL-03`, `INV-01` | Existing docs encode owner boundaries | Manual review plus index audit | `sed -n` review of touched flow/templates; `python3 scripts/check_memory_bank_index.py` | Same audit if CI exists | Manual review for semantic ownership wording | `none` |
| Generic leakage guard | `REQ-04`, `NEG-01`, `CHK-04`, `SOL-05`, `INV-02` | No existing leakage checker | Search for downstream-specific terms in generic surfaces | `rg -n "zelma|Codex session|session_meta|zellij|\\.zelma|process_argv" memory-bank/flows/templates memory-bank/flows/feature-flow.md` | Same search if CI exists | none | `none` |

## Open Questions / Ambiguities

No unresolved `OQ-*` remain after FPF decisions in `decision-log.md`.

## Resolved Questions

| Decision ref | Question | Resolution | Consequence |
| --- | --- | --- | --- |
| `DL-02` | Should the support template be named `privacy-source-boundary.md` or split into separate privacy-boundary and source-inventory templates? | Use one combined support template | `PRE-02` can be satisfied without human gate |
| `DL-03` | Should the support doc default to `doc_function: evidence` or `reference`? | Use `doc_function: reference` | `PRE-01` can be satisfied without human gate |

## Environment Contract

| Area | Contract | Used by | Failure symptom |
| --- | --- | --- | --- |
| setup | Run from repository root in this worktree | `STEP-01` - `STEP-07` | Relative paths fail or index script cannot find `memory-bank/` |
| test | Lightweight checks only; no app build/runtime exists | `CHK-01` - `CHK-05` | Attempts to run nonexistent app tests |
| access / network / secrets | No secrets required. External source examples are read-only context and must not be copied into generic templates | `STEP-01`, `STEP-03`, `STEP-06` | Generic template contains downstream runtime terms or private/raw source content |

## Preconditions

| Precondition ID | Canonical ref | Required state | Used by steps | Blocks start |
| --- | --- | --- | --- | --- |
| `PRE-01` | `CON-01`, `SD-02`, `INV-01` | Support doc remains non-canonical owner | `STEP-03`, `STEP-04`, `STEP-05` | yes |
| `PRE-02` | `SD-01`, `DL-02` | Single-template naming decision accepted locally | `STEP-03`, `STEP-04`, `STEP-05` | yes |
| `PRE-03` | `INV-02`, `NS-02` | Downstream specifics are barred from generic template surfaces | `STEP-03`, `STEP-06`, `STEP-07` | yes |

## Workstreams

| Workstream | Implements | Result | Owner | Dependencies |
| --- | --- | --- | --- | --- |
| `WS-1` | `REQ-01`, `REQ-02`, `SOL-01`, `CTR-01` - `CTR-04` | New support template | agent | `PRE-01`, `PRE-02`, `PRE-03` |
| `WS-2` | `REQ-03`, `REQ-05`, `SOL-02` - `SOL-04` | Routing/index updates | agent | `WS-1` |
| `WS-3` | `REQ-04`, `SOL-05`, `INV-02` | Generic leakage review and evidence | agent | `WS-1`, `WS-2` |

## Approval Gates

| Approval Gate ID | Trigger | Applies to | Why approval is required | Approver / evidence |
| --- | --- | --- | --- | --- |
| `AG-01` | A proposed change would introduce a global privacy/legal policy, runtime behavior, or downstream-specific implementation contract | `STEP-03` - `STEP-05` | This would exceed issue 16 and feature scope | Human approval in issue/PR comment |

## Порядок работ

| Step ID | Actor | Implements | Goal | Touchpoints | Artifact | Verifies | Evidence IDs | Check command / procedure | Blocked by | Needs approval | Escalate if |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| `STEP-01` | agent | `REQ-05` | Confirm current template/support index shape | `memory-bank/flows/templates/`, `memory-bank/features/README.md` | Discovery notes in this plan | `CHK-05` | `EVID-05` | `rg --files memory-bank/flows/templates memory-bank/features` | none | `none` | Reference pattern contradicts feature-flow |
| `STEP-02` | agent | `REQ-01`, `REQ-02`, `SOL-01` | Draft `privacy-source-boundary.md` wrapper template | `memory-bank/flows/templates/feature/support/privacy-source-boundary.md` | New template | `CHK-02` | `EVID-02` | Structured review against `REQ-02` and `CTR-*` | `PRE-01`, `PRE-02`, `PRE-03` | `AG-01` if scope expands | Template needs canonical facts owned by `brief.md` or `design.md` |
| `STEP-03` | agent | `REQ-03`, `SOL-02` | Update feature-flow support doc rules and support ID taxonomy | `memory-bank/flows/feature-flow.md` | Flow update | `CHK-03` | `EVID-03` | Review optional support docs and stable identifiers sections | `STEP-02` | `AG-01` if new global policy is proposed | Flow change affects unrelated feature lifecycle rules |
| `STEP-04` | agent | `REQ-03`, `REQ-05`, `SOL-03`, `SOL-04` | Update feature/templates navigation and optional support refs | `memory-bank/flows/templates/README.md`, `memory-bank/flows/templates/feature/*.md`, `memory-bank/features/README.md` | Routing/index updates | `CHK-01`, `CHK-03`, `CHK-05` | `EVID-01`, `EVID-03`, `EVID-05` | Search and index audit | `STEP-02`, `STEP-03` | `none` | Index script reports missing expected route |
| `STEP-05` | agent | `REQ-04`, `SOL-05` | Check generic leakage | `memory-bank/flows/templates`, `memory-bank/flows/feature-flow.md` | Leakage search output | `CHK-04` | `EVID-04` | `rg -n "zelma|Codex session|session_meta|zellij|\\.zelma|process_argv" memory-bank/flows/templates memory-bank/flows/feature-flow.md` | `STEP-02` - `STEP-04` | `none` | Search returns downstream-specific terms |
| `STEP-06` | agent | `REQ-01` - `REQ-05` | Run final lightweight checks | Full touched docs | Check outputs | `CHK-05` | `EVID-05` | `python3 scripts/check_memory_bank_index.py && git diff --check` | `STEP-05` | `none` | Link/index or whitespace check fails |
| `STEP-07` | agent | `REQ-01` - `REQ-05` | Capture final verify/evidence summary | Feature package / final report | Evidence summary | `CHK-01` - `CHK-05` | `EVID-01` - `EVID-05` | Summarize command outputs and review verdicts | `STEP-06` | `none` | Any canonical check remains failing |

## Parallelizable Work

- `PAR-01` Manual review of template content and routing wording can happen after `STEP-04` while command checks run.
- `PAR-02` Do not parallelize edits to `feature-flow.md` and template index wording because they share support-doc taxonomy.

## Checkpoints

| Checkpoint ID | Refs | Condition | Evidence IDs |
| --- | --- | --- | --- |
| `CP-01` | `STEP-02`, `CHK-02`, `SOL-01` | New template contains required sections and owner-boundary wording | `EVID-02` |
| `CP-02` | `STEP-03`, `STEP-04`, `CHK-01`, `CHK-03` | Routing/index updates are discoverable and ownership-safe | `EVID-01`, `EVID-03` |
| `CP-03` | `STEP-05`, `STEP-06`, `CHK-04`, `CHK-05` | No downstream leakage in generic surfaces; index/diff checks pass | `EVID-04`, `EVID-05` |

## Execution Risks

| Risk ID | Risk | Impact | Mitigation | Trigger |
| --- | --- | --- | --- | --- |
| `ER-01` | Support template becomes too project-specific | Violates `REQ-04` and makes template non-portable | Keep examples generic and run `CHK-04` | Downstream terms appear in generic surfaces |
| `ER-02` | Support doc ownership blurs with canonical docs | Contradicts feature-flow and causes review ambiguity | Repeat `must_not_define` and conflict-action wording | New template contains `REQ-*`, `SOL-*` or `STEP-*` as owned facts |
| `ER-03` | Index script expects README updates beyond obvious navigation | Blocks `CHK-05` | Read script output and update only required indexes | `check_memory_bank_index.py` fails |

## Stop Conditions / Fallback

| Stop ID | Related refs | Trigger | Immediate action | Safe fallback state |
| --- | --- | --- | --- | --- |
| `STOP-01` | `NS-02`, `INV-02`, `CHK-04` | Generic surfaces require downstream runtime details to be useful | Stop and open human gate | Keep feature package docs; do not publish template change |
| `STOP-02` | `AG-01`, `NS-03` | Work requires global privacy/legal policy or runtime behavior | Stop and ask for explicit scope decision | Documentation package remains planned |
| `STOP-03` | `CHK-05` | Index audit exposes broad repository inconsistency unrelated to this feature | Document residual risk and avoid unrelated fixes | Touched docs only |

## Plan-local Evidence

| Evidence ID | Artifact | Producer | Path contract | Reused by checkpoints |
| --- | --- | --- | --- | --- |
| `EVID-09` | Review-improve cycle summaries | reviewer / implementer | Final response and, if needed, feature `decision-log.md` entries | `CP-01`, `CP-02`, `CP-03` |

## Execution Result

| Check ID | Result | Evidence carrier |
| --- | --- | --- |
| `CHK-01` | pass | Local `rg -n "privacy-source-boundary\|Privacy / Source Boundary\|source-boundary" memory-bank/flows memory-bank/features/README.md` output showed the new support template and routing/index references. |
| `CHK-02` | pass | Manual review of `memory-bank/flows/templates/feature/support/privacy-source-boundary.md` confirmed allowed metadata, excluded data, source inventory confidence/status, owner boundaries and test evidence boundary sections. |
| `CHK-03` | pass | Manual review confirmed routing/ownership updates in `memory-bank/flows/feature-flow.md`, feature README, brief, design and implementation-plan templates. |
| `CHK-04` | pass | Local leakage search for downstream-specific source/runtime terms in generic template surfaces returned no matches. |
| `CHK-05` | pass | `python3 scripts/check_memory_bank_index.py` and `git diff --check` passed locally. |

Human gate: none. No approval gate was triggered.

## Готово для приемки

- Все workstreams завершены или явно остановлены через `STOP-*`.
- Все checkpoints имеют evidence.
- Required local suites зелёные.
- Manual-only gaps закрыты review verdicts; approval gates не требуются, если scope не расширялся.
- Финальная приемка идёт по `brief.md` `Verify`, а не по этому checklist.

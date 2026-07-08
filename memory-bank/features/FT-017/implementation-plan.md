---
title: "FT-017: Implementation Plan"
doc_kind: feature
doc_function: derived
purpose: "Execution plan for FT-017. Fixes grounded workstreams, steps, risks and test strategy without redefining problem or solution facts."
derived_from:
  - brief.md
  - design.md
  - decision-log.md
status: archived
audience: humans_and_agents
must_not_define:
  - ft_017_scope
  - ft_017_selected_design
  - ft_017_acceptance_criteria
  - ft_017_blocker_state
---

# FT-017: Implementation Plan

## Цель текущего плана

Implement the optional UI design guide pattern selected in [`design.md`](design.md): add the generic guide destination/template, route to it from frontend/index docs, and verify that the result is discoverable, optional and framework-agnostic.

## Grounding / Support References

| Document | Role in this plan | Facts reused | Conflict action |
| --- | --- | --- | --- |
| `brief.md` | canonical problem / verify owner | `REQ-*`, `NS-*`, `ASM-*`, `CON-*`, `SC-*`, `NEG-*`, `CHK-*`, `EVID-*` | Update `brief.md` first |
| `design.md` | canonical solution owner | `SOL-*`, `C4-*`, `SD-*`, `CTR-*`, `INV-*`, `FM-*`, `RB-*` | Update `design.md` or ADR first |
| `decision-log.md` | decision provenance | `FDL-*`, FPF rationale and accepted feature-local decisions | Update `decision-log.md` and owner doc before execution continues |
| `none` | optional support docs | No feature-support docs are required for this documentation-only feature. | Create support docs only if implementation discovers real ambiguity that cannot be captured by owner docs |

## Current State / Reference Points

| Path / module | Current role | Why relevant | Reuse / mirror |
| --- | --- | --- | --- |
| `memory-bank/engineering/frontend.md` | Canonical frontend engineering guidance | Issue 17 requires linking the optional design guide with frontend guidance. | Add a concise optional companion route; do not move frontend contract ownership. |
| `memory-bank/engineering/README.md` | Engineering documentation index | The optional guide must be discoverable from the engineering layer. | Add annotated optional child link. |
| `memory-bank/domain/README.md` | Domain owner boundary | It explicitly excludes UI design system ownership. | Do not create `domain/design-guide` as generic owner. |
| `memory-bank/flows/templates/README.md` | Template wrapper index | Used to confirm no separate flow-template artifact is required for the selected optional engineering destination. | Do not modify unless `design.md` is updated first. |
| `scripts/check_memory_bank_index.py` | Link/reachability audit | Required by repository checks. | Run after docs are added. |

## Test Strategy

| Test surface | Canonical refs | Existing coverage | Planned automated coverage | Required local suites / commands | Required CI suites / jobs | Manual-only gap / justification | Manual-only approval ref |
| --- | --- | --- | --- | --- | --- | --- | --- |
| Markdown navigation and frontmatter dependencies | `REQ-01`, `REQ-03`, `CHK-01`, `SOL-02` | `scripts/check_memory_bank_index.py` audits links, reachability and README index contracts. | Use existing script; add no new code tests unless script reveals a reusable gap. | `python3 scripts/check_memory_bank_index.py` | Same command if CI exists; otherwise local evidence is sufficient for docs-only feature. | none | none |
| Generic UI guide content | `REQ-02`, `REQ-04`, `NEG-01`, `CHK-02`, `SOL-01`, `SOL-03` | No existing dedicated guide. | Targeted `rg` and manual review against required sections and framework-specific defaults. | `rg -n "components|forms|buttons|actions|tables|navigation|states|labels|screenshots|source code paths|agent instructions" memory-bank/engineering` plus targeted source-framework guard on new generic docs | Same command if CI exists; otherwise local evidence is sufficient for docs-only feature. | Manual semantic review is required to confirm no copied downstream content; this is acceptable because the repo has no semantic linter. | none |
| Diff hygiene and owner boundaries | `REQ-03`, `REQ-04`, `CHK-03`, `INV-01`, `INV-02` | `git diff --check`; manual owner-boundary review. | Use existing git check and review touched docs. | `git diff --check` | Same command if CI exists; otherwise local evidence is sufficient for docs-only feature. | Manual owner-boundary review is required. | none |

## Open Questions / Ambiguities

No unresolved blocking questions remain after [`decision-log.md`](decision-log.md):

| Open Question ID | Question | Why unresolved | Blocks | Default action / escalation owner |
| --- | --- | --- | --- | --- |
| `OQ-00` | none | FPF decisions `FDL-001` through `FDL-004` closed the material owner/location questions. | none | If implementation discovers a new owner-boundary conflict, stop and update `brief.md` / `design.md` before continuing. |

## Environment Contract

| Area | Contract | Used by | Failure symptom |
| --- | --- | --- | --- |
| setup | Work from repository root with Python 3 and `rg` available. | `STEP-01`, `STEP-02`, `STEP-03` | Verification commands cannot run or inspect expected paths. |
| test | Repository-local docs checks are authoritative for this docs-only feature. | `CHK-01`, `CHK-02`, `CHK-03` | Broken links, missing index annotations, whitespace errors or framework-specific defaults in target generic docs. |
| access / network / secrets | No network or secrets are required to implement docs. | all steps | Any required external source lookup means scope has changed and must be escalated. |

## Preconditions

| Precondition ID | Canonical ref | Required state | Used by steps | Blocks start |
| --- | --- | --- | --- | --- |
| `PRE-01` | `brief.md`, `design.md`, `decision-log.md` | Upstream owner docs are active and agree on engineering-layer optional guide. | `STEP-01`, `STEP-02` | yes |
| `PRE-02` | `REQ-04`, `SOL-03`, `CTR-03` | Implementer accepts that source-project framework rules are not generic defaults. | `STEP-01`, `STEP-03` | yes |

## Workstreams

| Workstream | Implements | Result | Owner | Dependencies |
| --- | --- | --- | --- | --- |
| `WS-1` | `REQ-01`, `REQ-02`, `SOL-01`, `SOL-03` | Optional UI design guide destination/template added with required sections and generic guardrails. | agent | `PRE-01`, `PRE-02` |
| `WS-2` | `REQ-03`, `SOL-02`, `CTR-02` | `frontend.md` and README/index routes make the optional guide discoverable. | agent | `WS-1` |
| `WS-3` | `REQ-04`, `CHK-01`, `CHK-02`, `CHK-03` | Verification evidence proves link integrity, optionality and generic content safety. | agent | `WS-1`, `WS-2` |

## Approval Gates

No human approval gate is required for the planned docs-only edits.

| Approval Gate ID | Trigger | Applies to | Why approval is required | Approver / evidence |
| --- | --- | --- | --- | --- |
| `AG-00` | none | none | No risky, irreversible, costly or external-effect operation is planned. | none |

## Порядок работ

| Step ID | Actor | Implements | Goal | Touchpoints | Artifact | Verifies | Evidence IDs | Check command / procedure | Blocked by | Needs approval | Escalate if |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| `STEP-01` | agent | `REQ-01`, `REQ-02`, `REQ-04`, `SOL-01`, `SOL-03` | Add optional UI design guide destination/template with generic prompts and required sections. | `memory-bank/engineering/ui-design-guide/README.md` | New optional guide doc | `CHK-02` | `EVID-02` | Review required section list and targeted source-framework guard. | `PRE-01`, `PRE-02` | none | The guide needs project-specific framework rules to be useful. |
| `STEP-02` | agent | `REQ-03`, `SOL-02`, `CTR-02` | Route readers from frontend/index docs to the optional guide. | `memory-bank/engineering/frontend.md`, `memory-bank/engineering/README.md` | Updated annotated links | `CHK-01`, `CHK-03` | `EVID-01`, `EVID-03` | Run index audit and inspect optional wording. | `STEP-01` | none | A routing change would make the optional guide mandatory. |
| `STEP-03` | agent | `REQ-01`, `REQ-02`, `REQ-03`, `REQ-04` | Verify and record evidence for delivery. | repository docs and evidence paths | Verification logs / notes | `CHK-01`, `CHK-02`, `CHK-03` | `EVID-01`, `EVID-02`, `EVID-03` | `python3 scripts/check_memory_bank_index.py`; `git diff --check`; targeted `rg`; manual owner-boundary review. | `STEP-01`, `STEP-02` | none | Checks reveal contradictions with `brief.md`, `design.md` or existing governance docs. |

## Parallelizable Work

- `PAR-01` `STEP-01` and `STEP-02` should not be parallelized because routing links depend on the final optional guide path.
- `PAR-02` `CHK-01`, `CHK-02` and `CHK-03` can run in any order after `STEP-01` and `STEP-02`.

## Checkpoints

| Checkpoint ID | Refs | Condition | Evidence IDs |
| --- | --- | --- | --- |
| `CP-01` | `STEP-01`, `REQ-01`, `REQ-02`, `SOL-01` | Optional guide exists with all required sections and generic guardrails. | `EVID-02` |
| `CP-02` | `STEP-02`, `REQ-03`, `SOL-02` | Routing docs link to the guide and mark it optional. | `EVID-01`, `EVID-03` |
| `CP-03` | `STEP-03`, `CHK-01`, `CHK-02`, `CHK-03` | Verification commands and manual review pass. | `EVID-01`, `EVID-02`, `EVID-03` |

## Execution Risks

| Risk ID | Risk | Impact | Mitigation | Trigger |
| --- | --- | --- | --- | --- |
| `ER-01` | The new guide becomes a second owner for frontend engineering policy. | Contradiction with `frontend.md` and `INV-01`. | Keep guide as concrete UI kit reference and route back to `frontend.md` for engineering contract. | Guide starts defining global frontend architecture, stack migration or component ownership policy. |
| `ER-02` | Source-specific framework names or rules leak into target generic docs. | Violates `REQ-04` and `NEG-01`. | Use generic placeholders and targeted content guard. | Target guide prescribes source stack choices. |
| `ER-03` | Index audit requires additional reachability links. | Incomplete discoverability or failed repository check. | Add the minimum annotated README/index route consistent with optionality. | `check_memory_bank_index.py` reports unreachable docs or index contract failures. |

## Stop Conditions / Fallback

| Stop ID | Related refs | Trigger | Immediate action | Safe fallback state |
| --- | --- | --- | --- | --- |
| `STOP-01` | `REQ-04`, `FDL-004`, `FM-01` | Implementation cannot satisfy usefulness without copying source-project framework rules. | Stop and raise human gate with alternatives. | Keep feature docs active, no generic guide change. |
| `STOP-02` | `INV-01`, `INV-02`, `FM-03` | Existing owner documents contradict the selected engineering-layer destination. | Stop and update `brief.md` / `design.md` or request human decision. | Keep `frontend.md` unchanged. |

## Plan-local Evidence

| Evidence ID | Artifact | Producer | Path contract | Reused by checkpoints |
| --- | --- | --- | --- | --- |
| `EVID-09` | Discovery and review-improve notes for feature-doc readiness | implementer / reviewer | `memory-bank/features/FT-017/feature-review-report.md` | `CP-01`, `CP-02`, `CP-03` |

## Execution Result

| Checkpoint ID | Result | Evidence |
| --- | --- | --- |
| `CP-01` | pass | `memory-bank/features/FT-017/evidence/chk-02/result.txt` |
| `CP-02` | pass | `memory-bank/features/FT-017/evidence/chk-01/result.txt`, `memory-bank/features/FT-017/evidence/chk-03/result.txt` |
| `CP-03` | pass | `memory-bank/features/FT-017/evidence/chk-01/result.txt`, `memory-bank/features/FT-017/evidence/chk-02/result.txt`, `memory-bank/features/FT-017/evidence/chk-03/result.txt` |

## Готово для приемки

- All workstreams are complete or explicitly stopped through `STOP-*`.
- All checkpoints have evidence.
- Required local suites are green, and CI does not contradict local verify.
- Manual-only semantic review is recorded in `EVID-02` or `EVID-03`.
- Support docs, if later added, do not disagree with `brief.md`, `decision-log.md`, `design.md` or this plan.
- Final acceptance follows `brief.md` `Verify`, not this checklist.

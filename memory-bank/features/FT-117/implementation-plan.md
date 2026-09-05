---
title: "FT-117: Recovery Acceptance Plan"
doc_kind: feature
doc_function: derived
purpose: "Forward recovery plan for independently reviewing, accepting and closing the pre-existing FT-117 documentation candidate without claiming retroactive Plan Ready compliance."
derived_from:
  - brief.md
  - design.md
status: archived
audience: humans_and_agents
must_not_define:
  - ft_117_scope
  - ft_117_selected_design
  - ft_117_acceptance_criteria
  - ft_117_validation_profile
---

# FT-117: Recovery Acceptance Plan

## Цель текущего плана

Провести от текущего frozen candidate вперёд воспроизводимую recovery-приёмку
issue #117: закрыть traceability gaps, получить независимые Plan Ready и
implementation verdicts, выполнить exact local/CI checks и только после этого
закрыть lifecycle package.

Core changes в commits `9822d88`–`d3639b1` являются existing adoption input.
Этот план не утверждает, что исторические writes прошли Plan Ready до execution.
Он разрешает только review, bounded corrections, validation и acceptance
existing candidate; новый semantic scope требует нового routing.

## Grounding Evidence

- Pre-change issue baseline: `daa7cb639d2bc2741da8b829a125546b45bf8a0e`.
- Core implementation candidate: `d3639b118a48f3215567a26a3b6d430cc2b4d6f6`.
- Recovery-plan starting revision: `fbb3d851ca2a8a58910326e9687c56e9547d75f9`.
- Grounded at: `2026-08-14`.

| Grounding ID | Exact command / path | Observed fact | Plan impact |
| --- | --- | --- | --- |
| `GRND-01` | `git cat-file -e daa7cb639d2bc2741da8b829a125546b45bf8a0e^{commit}`; `git cat-file -e d3639b118a48f3215567a26a3b6d430cc2b4d6f6^{commit}`; `git cat-file -e fbb3d851ca2a8a58910326e9687c56e9547d75f9^{commit}` | Baseline, core candidate and recovery start are immutable repository objects | `STEP-01` re-verifies the exact objects before acceptance |
| `GRND-02` | `git diff --name-status daa7cb639d2bc2741da8b829a125546b45bf8a0e..d3639b118a48f3215567a26a3b6d430cc2b4d6f6` | Core candidate changes nine autonomy/routing/flow/carrier documents | All nine surfaces are mapped below |
| `GRND-03` | `git diff --name-status d3639b118a48f3215567a26a3b6d430cc2b4d6f6..fbb3d851ca2a8a58910326e9687c56e9547d75f9` | Recovery start adds the project package, Research handoff, three priming inputs and validation refinements | Package, projections and metadata are explicit execution surfaces |
| `GRND-04` | `.github/workflows/ci.yml`, job `validate-template` | Required CI runs manifest tests/validator, both lints, doctor and downstream-init smoke test | `STEP-04` requires exact `CI / validate-template` success |
| `GRND-05` | `template/memory-bank/flows/autonomy-boundaries.md`; `template/memory-bank/flows/{routing,research,bug-fix,feature,epic}.md`; `template/memory-bank/flows/validation-profiles.md` | Canonical owners contain the candidate protocol and cross-flow behavior | `STEP-02` executes exact structural assertions and semantic review |
| `GRND-06` | `memory-bank/features/README.md`; `memory-bank/.lock` | FT-117 belongs to project-local history; the feature index is adapted metadata rather than generic payload | `STEP-03` verifies placement and zero template leakage |

## Historical Boundary And Recovery Disposition

The first independent review recorded `BLOCKED_PLAN_READY` because the previous
plan described already-delivered work as future execution. This revision adopts
the reviewer disposition: it treats `daa7cb6..fbb3d85` as immutable input and
defines only prospective recovery acceptance. Historical noncompliance remains
visible in the review record and is not erased by later clean verdicts.

## Implementation Priming

| Order | Exact path / symbol | Grounding refs | Purpose | Required before |
| --- | --- | --- | --- | --- |
| `1` | `memory-bank/features/FT-117/brief.md#constraints--assumptions` | `GRND-01`–`GRND-03` | Confirm recovery boundary and unchanged issue scope | `STEP-01` |
| `2` | `memory-bank/features/FT-117/design.md#design-pack` | `GRND-02`, `GRND-03` | Confirm every canonical changed owner and projection classification | `STEP-01` |
| `3` | `template/memory-bank/flows/autonomy-boundaries.md#structured-decision-protocol` | `GRND-05` | Verify roles, outcomes, tie-breakers and approval separation | `STEP-02` |
| `4` | `template/memory-bank/flows/routing.md#human-routing` | `GRND-05` | Verify P0-safe decision and exact Human Routing conditions | `STEP-02` |
| `5` | `template/memory-bank/flows/research.md#boundary-rules` | `GRND-03`, `GRND-05` | Verify probe handoff and repeat-routing semantics | `STEP-02` |
| `6` | `template/memory-bank/flows/bug-fix.md#entry-gate`; `template/memory-bank/flows/feature.md#upstream-ready--plan-ready` | `GRND-05` | Verify expected-behavior routing and review convergence | `STEP-02` |
| `7` | `template/memory-bank/flows/validation-profiles.md#escalation-and-downgrade-rules` | `GRND-05` | Verify exact execution approval gate | `STEP-02` |
| `8` | `.github/workflows/ci.yml#L12`; job `validate-template` | `GRND-04` | Confirm required CI name and commands | `STEP-04` |

If any named path, heading, commit object or CI job is absent, stop for
re-grounding before execution.

## Test Strategy

| Surface | Refs | Exact local command / procedure | Required CI | Pass criterion |
| --- | --- | --- | --- | --- |
| Priming schema | `CHK-01` | `ruby tools/validate-priming-manifests-test.rb && ruby tools/validate-priming-manifests.rb template/memory-bank` | `CI / validate-template` | Tests pass and exactly 16 manifests validate |
| Links and reachability | `CHK-02` | `memory-bank-cli lint --scope-root template/memory-bank --entrypoint template/memory-bank/README.md && memory-bank-cli lint --scope-root memory-bank --entrypoint memory-bank/README.md` | `CI / validate-template` | Both commands return `Result: OK` |
| Template doctor | `CHK-03` | `memory-bank-cli doctor --profile template` | `CI / validate-template` | `0 error(s), 0 warning(s)` |
| Diff hygiene | `CHK-04` | `git diff --check daa7cb639d2bc2741da8b829a125546b45bf8a0e..HEAD && git diff --check` before commit; repeat the first command after commit | `CI / validate-template` | Committed range and frozen working-tree corrections both produce no output and exit zero |
| Protocol structure | `CHK-05` | Run the exact assertion block below, then independent semantic review of every `SC-*`/`NEG-*` | `CI / validate-template` plus independent review record | Every assertion returns at least one match; reviewer reports no open critical/important finding |
| Downstream install | `CHK-02`, `CHK-03` | Use the `Smoke-test downstream init` procedure from `.github/workflows/ci.yml` in CI | `CI / validate-template` | Job step succeeds |

Exact `CHK-05` assertion block:

```sh
rg -q 'authority source' template/memory-bank/flows/autonomy-boundaries.md
rg -q 'decision owner' template/memory-bank/flows/autonomy-boundaries.md
rg -q 'canonical carrier' template/memory-bank/flows/autonomy-boundaries.md
rg -q 'execution approver / approval evidence' template/memory-bank/flows/autonomy-boundaries.md
rg -q 'Outcome: proceed \| bounded_probe \| escalate' template/memory-bank/flows/autonomy-boundaries.md
rg -q 'существующий canonical pattern; наименьшее обратимое изменение' template/memory-bank/flows/autonomy-boundaries.md
rg -q 'P0 остаётся read-only' template/memory-bank/flows/routing.md
rg -U -q 'повторяет Structured Decision Protocol и Task\n[[:space:]]+Routing' template/memory-bank/flows/research.md
rg -q 'Human Gate нужен только при outcome' template/memory-bank/flows/feature.md
rg -q 'только при outcome `escalate`' template/memory-bank/flows/bug-fix.md
rg -q 'task/project-policy preauthorization may be approval evidence|task/project-policy preauthorization может быть approval evidence' template/memory-bank/flows/validation-profiles.md
```

## Open Questions / Ambiguities

`none` — historical sequencing is a disclosed lifecycle debt, not an unknown.
The recovery decision is whether the frozen existing candidate is acceptable;
it does not reconstruct or conceal its original execution order.

## Environment Contract

| Area | Contract | Used by | Failure symptom |
| --- | --- | --- | --- |
| setup | Repository clone/worktree contains Git objects through `HEAD`; `memory-bank-cli` is installed | `STEP-01`–`STEP-04` | Commit receipt or command unavailable |
| local validation | Run all six Test Strategy rows; do not substitute a generic docs read-through | `STEP-02`, `STEP-03` | Evidence cannot establish acceptance |
| CI | Existing PR #118 branch `docs/issue-117-autonomy-fpf`; required job `CI / validate-template` | `STEP-04` | Job absent, pending or non-success |
| authority | User approval `2026-08-14` explicitly authorizes commit and push for this work and says not to request commit/push permission again; target remains branch `docs/issue-117-autonomy-fpf` and PR #118 | `STEP-04`, `STEP-06` | Target/scope differs or authority is revoked |

## Preconditions

| ID | Required state | Used by | Blocks start |
| --- | --- | --- | --- |
| `PRE-01` | `brief.md` and `design.md` remain active; `brief.md` is `planned` | `STEP-01`–`STEP-03` | yes |
| `PRE-02` | This draft plan receives clean independent Plan Ready review, then the active revision receives clean re-review | `STEP-01`–`STEP-04` | yes |
| `PRE-03` | No semantic change beyond `REQ-01`–`REQ-08` entered the candidate | `STEP-02`, `STEP-03` | yes |

## Design Realization Mapping

| Refs | Direct owner / projection | Recovery target | Steps | Checks / evidence |
| --- | --- | --- | --- | --- |
| `SOL-01`, `SD-01`, `SD-02`, `SD-04` | `template/memory-bank/flows/autonomy-boundaries.md` | Protocol, carrier and approval contract | `STEP-02` | `CHK-05`, `EVID-02`, `EVID-06` |
| `SOL-02`, `SD-03` | `template/memory-bank/flows/routing.md` | P0, Research and Human Routing | `STEP-02` | `CHK-05`, `EVID-03`, `EVID-06` |
| `SOL-05`, `SD-03` | `template/memory-bank/flows/research.md`; `flows/priming/research.yaml` projection | Probe lifecycle, bootstrap input and return | `STEP-02`, `STEP-03` | `CHK-01`, `CHK-05`, `EVID-03` |
| `SOL-03` | `template/memory-bank/flows/bug-fix.md`; `flows/feature.md`; `flows/priming/{bug-fix,feature}.yaml` projections | Reroute/replan semantics and required priming | `STEP-02`, `STEP-03` | `CHK-01`, `CHK-05`, `EVID-03` |
| `SOL-01`, `SD-01` | `template/memory-bank/flows/epic.md`; `flows/templates/epic/decision-log.md` | Epic carrier and template terminology | `STEP-02` | `CHK-02`, `CHK-05`, `EVID-02` |
| `SOL-04`, `SOL-06`, `SD-04` | `template/memory-bank/flows/validation-profiles.md` | Approval timing and validation floor | `STEP-02` | `CHK-05`, `EVID-04`, `EVID-06` |
| `SOL-01`–`SOL-06` | `template/memory-bank/{engineering,flows}/README.md` projections | Navigation wording | `STEP-03` | `CHK-02` |
| `REQ-01`–`REQ-08` | `memory-bank/features/FT-117/*`, `memory-bank/features/README.md`, `memory-bank/.lock` | Project-local traceability, reachability and adapted ownership metadata | `STEP-01`, `STEP-03` | `CHK-02`, `CHK-03`, `EVID-05`, `EVID-06` |
| `C4-00` | `design.md` | No runtime artifact | `STEP-02` | `CHK-05` |

## Approval Gates

| ID | Action | Evidence | Limits |
| --- | --- | --- | --- |
| `AG-01` | Commit and push bounded corrections to branch `docs/issue-117-autonomy-fpf` and update draft PR #118 so required CI can run | `approved`: user message `2026-08-14`, “разрешаю и не надо меня спрашивать разрешение на коммитить и пушать” | This FT-117 recovery work on the existing branch/PR only; no merge, release, deploy, publication or unrelated mutation |

## Порядок работ

| Step | Actor | Goal / exact touchpoints | Verifies | Evidence | Blocked by | Stop condition |
| --- | --- | --- | --- | --- | --- | --- |
| `STEP-01` | agent | Verify three commit objects and both immutable diffs; reconcile only `memory-bank/features/FT-117/{README,brief,design,implementation-plan,feature-review-report}.md` | Recovery boundary and complete change-surface map | `EVID-05`, `EVID-06` | `PRE-01`, `PRE-02` | Git receipt differs or new scope appears |
| `STEP-02` | agent | Run exact `CHK-05` block and semantic acceptance against all canonical owners in Design Realization Mapping | `SC-01`–`SC-10`, `NEG-01`–`NEG-03` | `EVID-02`–`EVID-05` | `STEP-01` | Any scenario lacks one owner or bounded correction |
| `STEP-03` | agent | Run all local structural checks, project-specific leak check `rg -n 'FT-117|issue #117' template/memory-bank` expecting no matches, and record results outside frozen plan | `CHK-01`–`CHK-05` | `EVID-05` | `STEP-02` | Any command or expected result fails |
| `STEP-04` | agent | With `AG-01`, commit/push bounded recovery execution revision to the existing branch; wait for `CI / validate-template`; obtain independent implementation review of the delivered diff | Full validation profile and implementation convergence | `EVID-05`, `EVID-06`, `AG-01` | `STEP-03`, `AG-01` | CI/review not successful or scope differs |
| `STEP-05` | agent | After clean implementation verdict, prepare one terminal metadata revision: record verdict, set `brief.md: delivery_status: done` and archive this plan | Feature terminal candidate | `EVID-06` | `STEP-04` | Any non-metadata change or Done predicate remains open |
| `STEP-06` | agent | With `AG-01`, commit/push the terminal metadata revision, wait for `CI / validate-template`, and run final convergence read without further mutation | Feature terminal state | `EVID-05`, `EVID-06`, `AG-01` | `STEP-05`, `AG-01` | Final CI is not successful; then restore non-terminal status in a corrective revision |

## Checkpoints

| ID | Condition | Evidence |
| --- | --- | --- |
| `CP-01` | Draft and then active plan revisions have independent clean Plan Ready verdicts | `EVID-06` |
| `CP-02` | Git receipts, all local checks and semantic acceptance pass | `EVID-05` |
| `CP-03` | Execution revision CI and separate implementation review are clean; terminal metadata revision CI is also successful | `EVID-05`, `EVID-06` |

## Stop Conditions / Fallback

| ID | Trigger | Immediate action | Safe state |
| --- | --- | --- | --- |
| `STOP-01` | Candidate requires new requirement, runtime change or external action outside `AG-01` | Stop and repeat Task Routing | Existing draft PR remains unmerged |
| `STOP-02` | Draft or active Plan Ready review has critical/important finding | Keep plan `draft` or return it to `draft`; fix and freeze a new revision | `brief.md: planned` |
| `STOP-03` | Local check, CI or implementation review fails | Keep feature non-terminal; fix only in scope and repeat affected gates | Active plan plus open draft PR |
| `STOP-04` | Historical sequence would need to be represented as compliant | Preserve disclosed debt; do not rewrite history or backdate evidence | Recovery acceptance remains explicit |

## Plan-local Evidence

| Evidence | Carrier | Producer | Rule |
| --- | --- | --- | --- |
| `EVID-05` | Local command transcript and GitHub `CI / validate-template` check | agent / CI | Exact revision and pass criteria required |
| `EVID-06` | `feature-review-report.md` review history and PR review record | non-authoring reviewer, transcribed by delivery agent | Must name reviewer, frozen revisions, findings/dispositions and exact verdict; never treat self-review as independent |

## Готово для приемки

- The historical execution-order debt remains explicit.
- Draft and active plan revisions each have clean independent Plan Ready review.
- All exact local checks and `CI / validate-template` pass on the delivered revision.
- Independent implementation review has no open critical/important finding.
- No merge, release, deploy, publication or live-state mutation is performed.

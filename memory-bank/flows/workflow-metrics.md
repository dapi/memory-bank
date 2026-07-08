---
title: Workflow Metrics
doc_kind: governance
doc_function: canonical
purpose: "Определяет safety-first метрики для проверки workflow selector-а и workflow profiles."
derived_from:
  - ../dna/governance.md
  - routing.md
  - workflow-decision-log.md
status: active
audience: humans_and_agents
canonical_for:
  - workflow_profile_metrics
  - workflow_profile_success_criteria
  - workflow_measurement_windows
---
# Workflow Metrics

Этот документ задает optional metrics template для проверки workflow selector-а. Он нужен, когда команда вводит или меняет workflow profiles и хочет проверить, стало ли меньше лишней documentation ceremony без потери safety, evidence и traceability.

Метрики образуют bundle: один быстрый показатель не доказывает успех. Safety/evidence metrics всегда проверяются раньше speed metrics.

## Evaluation Window Template

| Window ID | Dates / source | Purpose | Included work |
| --- | --- | --- | --- |
| `BASELINE-01` | `<baseline window>` | Сравнить с периодом до изменения selector-а | Closed issues / merged PRs / accepted task packages |
| `PILOT-01` | `<pilot window>` | Проверить новые или измененные workflow profiles | Closed issues / merged PRs / accepted task packages |
| `REVIEW-01` | `<review date or review event>` | Принять решение: keep, tighten, extend, rollback | Metrics report plus qualitative review |

If a PR or task is still open at the end of a window, count it only in qualitative review unless the project defines another rule.

## Metric Cards

| Metric ID | Characteristic | Scale / Unit | Formula | Suggested target | Evidence source |
| --- | --- | --- | --- | --- | --- |
| `WFM-01` | Routing coverage | Ratio, `%` | Work items with explicit routing signature or selected workflow profile / eligible work items | Project-defined threshold, commonly high enough to audit routing habits | Issue/PR body, labels, comments, task/feature docs |
| `WFM-02` | Feature overuse rate | Ratio, `%` | Non-feature tasks that created full feature package without promotion trigger / non-feature tasks | Lower than baseline without safety regression | Issue/PR labels, docs created, routing notes |
| `WFM-03` | Safety misroute count | Count | High-risk or contract-changing work delivered through a profile too small for its risk | `0` for the review window | PR diff review, review comments, incident notes, docs |
| `WFM-04` | Missing-evidence review burden | Count per N PRs/tasks | Review findings asking for missing repro, verification, rollout/backout or evidence because selected profile was too small | Project-defined maximum; should not increase vs baseline | PR review comments, review reports, task docs |
| `WFM-05` | Regression / verification coverage | Ratio, `%` | Work items with required automated coverage or approved manual-only exception / work items requiring verification | Project-defined threshold; stricter for critical domains | PR diff, test reports, evidence notes |
| `WFM-06` | Rework from wrong profile | Ratio, `%` | Work items requiring profile promotion after review / work items with routing signature | Low and not worse than baseline | PR timeline, issue comments, package history |
| `WFM-07` | Durable traceability for managed work | Ratio, `%` | Managed task/feature/epic work with links from issue/PR to owner docs / all managed work | High enough that owner docs are findable during review | Issue/PR links, memory-bank package indexes |
| `WFM-08` | Small-task lead time | Duration | Median ready-to-PR or ready-to-merge duration for compact profiles | Secondary observation only; no success if safety gates fail | Issue/PR timestamps or project tracker |

## Decision Rule

Evaluate in this order:

1. If `WFM-03 > 0`, compact or reduced-documentation profiles are unsafe for the observed window. Tighten selector or promotion triggers before claiming success.
2. If required verification coverage in `WFM-05` misses the project-defined threshold for critical work, tighten the relevant profile even if speed improved.
3. If `WFM-04` or `WFM-06` gets worse materially, the profile likely hides context or evidence; update workflow docs/templates before expanding use.
4. If `WFM-01` and `WFM-07` are too low, the selector may be correct in prose but not adopted; improve prompts, PR templates or developer brief.
5. Only after safety/evidence/traceability gates are acceptable, use `WFM-02` and `WFM-08` to judge whether ceremony decreased.
6. If lead time improves while safety/evidence/rework worsens, the selector change is not successful.

## Measurement Notes

- A `workflow_profile` may be captured in an issue field, PR body, label, task package, feature package or explicit agent comment.
- Baseline classification may be best-effort, but the method must be stated.
- Manual-only verification counts only when it has a reason, procedure and approval/evidence reference.
- Do not average unrelated risk classes into one score. Report critical/high-risk work separately when possible.
- If thresholds are not known yet, mark them as project-adapted fields and run a qualitative review before setting targets.

## Review Output Template

```markdown
# Workflow Metrics Review

## Window

- Baseline: `<baseline window>`
- Pilot: `<pilot window>`
- Review: `<review event>`

## Safety Gates

| Metric | Result | Verdict | Notes |
| --- | --- | --- | --- |
| `WFM-03` |  | pass/fail |  |
| `WFM-04` |  | pass/fail |  |
| `WFM-05` |  | pass/fail |  |
| `WFM-06` |  | pass/fail |  |

## Adoption / Traceability

| Metric | Result | Verdict | Notes |
| --- | --- | --- | --- |
| `WFM-01` |  | pass/fail |  |
| `WFM-07` |  | pass/fail |  |

## Ceremony / Speed

| Metric | Result | Verdict | Notes |
| --- | --- | --- | --- |
| `WFM-02` |  | observation |  |
| `WFM-08` |  | observation |  |

## Decision

`keep` / `tighten` / `extend pilot` / `rollback`

## Follow-ups

- `FOLLOW-UP-01` ...
```

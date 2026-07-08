---
title: Workflow Decision Log
doc_kind: governance
doc_function: canonical
purpose: "Фиксирует причины, решения и expected consequences изменений workflow selector-а и workflow profiles."
derived_from:
  - ../dna/governance.md
  - workflows.md
status: active
audience: humans_and_agents
canonical_for:
  - workflow_decision_rationale
  - workflow_change_causal_trace
  - workflow_profile_review_rules
---
# Workflow Decision Log

Этот документ является optional process-layer owner-ом для решений о workflow selector-е. Используй его, когда проект меняет правила выбора workflow profile, вводит compact profile, ужесточает promotion triggers или пересматривает documentation ceremony.

Decision log не заменяет [workflows.md](workflows.md): canonical routing rules остаются там. Здесь живет причинно-следственная запись: почему selector менялся, какие риски учитывались, как решение будет пересмотрено и какими evidence оно должно подтверждаться.

## When To Use

Создавай или обновляй запись `DL-WF-*`, если изменение:

- добавляет, удаляет или переименовывает workflow profile;
- меняет promotion triggers между compact, managed, feature и epic profiles;
- снижает или повышает required documentation/evidence для класса задач;
- меняет routing signature fields или их interpretation;
- меняет safety/evidence gates для выбора workflow;
- вводит pilot/evaluation period для workflow selector-а.

Не используй этот документ для feature-local design decisions, ADR-level architecture decisions или PR-specific implementation choices.

## Decision Entry Template

```markdown
## DL-WF-XXX: Short decision name

**Status:** proposed / accepted / superseded / rejected

**Date:** YYYY-MM-DD

**Context:** What observed problem or opportunity makes the workflow change necessary.

**Evidence:** Which carriers support the claim: issues, PR reviews, review reports, metrics, incident notes or downstream feedback.

**Decision:** What changes in the selector, profiles, routing signature, gates or documentation requirements.

**Expected consequences:**

1. What should improve.
2. What safety/evidence risk is being controlled.
3. What would show the decision is wrong.

**Review rule:** When and how the decision is revisited. Link to [workflow-metrics.md](workflow-metrics.md) if measured.

**Non-goals:** What this decision deliberately does not change.

**Follow-ups:** Optional issue/PR links for implementation or later tightening.
```

## Required Fields

| Field | Required | Rule |
| --- | --- | --- |
| Status | yes | Use `proposed`, `accepted`, `superseded` or `rejected`. |
| Date | yes | Use the decision date, not the measurement window. |
| Context | yes | State the workflow problem without hiding it inside the decision. |
| Evidence | yes | Name concrete carriers or explicitly mark evidence as qualitative. |
| Decision | yes | State the selector/profile change in reviewable terms. |
| Expected consequences | yes | Include both expected benefit and safety/evidence risk. |
| Review rule | yes | Explain when the decision is revisited and what evidence is used. |
| Non-goals | recommended | Prevent scope creep into unrelated flows. |
| Follow-ups | optional | Link implementation issues or later audits. |

## Review Rules

- Safety and evidence regressions override ceremony or lead-time gains.
- A compact profile is not successful if it increases safety misroutes, missing evidence or rework from wrong routing.
- If workflow metrics are inconclusive, keep the decision in `proposed` or extend the review window instead of silently promoting it.
- If a selector change affects feature/epic lifecycle gates, update the canonical owner documents before treating the decision as accepted.
- If a decision depends on a document that does not exist yet, record it as a follow-up dependency rather than linking to a missing local file.

## Example Entry

```markdown
## DL-WF-001: Add compact profile evaluation

**Status:** proposed

**Date:** YYYY-MM-DD

**Context:** Small low-risk tasks are receiving heavyweight documentation, but the project still needs evidence that a smaller profile will not hide risky work.

**Evidence:** Review findings and PR notes from the selected baseline window.

**Decision:** Pilot a compact profile for low-risk tasks with explicit promotion triggers for contract change, high risk, missing evidence or design uncertainty.

**Expected consequences:**

1. Less ceremony for small tasks.
2. No increase in safety misroutes or missing evidence.
3. Profile is tightened if review comments show repeated missing context.

**Review rule:** Evaluate the pilot with [workflow-metrics.md](workflow-metrics.md) after the configured pilot window.

**Non-goals:** Do not change feature-flow or epic-flow lifecycle gates.
```

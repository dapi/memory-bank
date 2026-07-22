---
title: R-XXX Research Package README Template
doc_kind: governance
doc_function: template
purpose: Wrapper-шаблон индекса и lifecycle stage для research package.
derived_from:
  - ../../research.md
status: active
audience: humans_and_agents
template_for: research
template_target_path: ../../../research/R-XXX/README.md
---

# R-XXX Research Package README Template

## Instantiated Frontmatter

```yaml
---
title: "R-XXX: <Research Name>"
doc_kind: research
doc_function: index
purpose: "Навигация и текущая stage evidence-backed research R-XXX."
derived_from:
  - ../../flows/research.md
  - brief.md
status: active
research_stage: intake
audience: humans_and_agents
---
```

## Instantiated Body

```markdown
# R-XXX: <Research Name>

## Current Stage

- Stage: `intake`
- Research owner: `<person or role>`
- Decision owner: `<person or role>`
- Source / trigger: `<issue, request, metric, interview or other evidence>`
- Next gate: `Bootstrap → Question Framed`

## Annotated Index

- [Research Brief](brief.md) — canonical question, boundaries, hypotheses and stopping condition.

Add `plan.md`, `evidence.md`, `synthesis.md` and `decision.md` only when they exist. For each, state the facts it owns; do not create placeholder links.
```

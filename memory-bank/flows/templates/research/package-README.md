---
title: R-XXX Research Package README Template
doc_kind: governance
doc_function: template
purpose: Wrapper-шаблон индекса research package без дублирования lifecycle state из canonical brief.
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
purpose: "Навигация по evidence-backed research package R-XXX."
derived_from:
  - ../../flows/research.md
  - brief.md
status: active
audience: humans_and_agents
---
```

## Instantiated Body

```markdown
# R-XXX: <Research Name>

## Lifecycle Owner

Текущий lifecycle state хранится только в поле `research_status` документа [Research Brief](brief.md). Не копируй status или current stage в этот index.

## Annotated Index

- [Research Brief](brief.md) — canonical question, boundaries, hypotheses and stopping condition.

Add `plan.md`, `evidence.md`, `synthesis.md` and `decision.md` only when they exist. For each, state the facts it owns; do not create placeholder links.
```

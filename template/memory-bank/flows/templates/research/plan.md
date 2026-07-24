---
title: R-XXX Research Plan Template
doc_kind: governance
doc_function: template
purpose: "Wrapper-шаблон conditional research plan: method, collection protocol, quality controls and stopping rules."
derived_from:
  - ../../research.md
status: active
audience: humans_and_agents
template_for: research
template_target_path: ../../../research/R-XXX/plan.md
---

# R-XXX Research Plan Template

Создавай, когда method choice, sampling, participant contact, experiment, benchmark, privileged data или collection protocol требуют review. Для compact desk research без такого trigger достаточно method note в `brief.md`.

## Instantiated Frontmatter

```yaml
---
title: "R-XXX: Research Plan"
doc_kind: research
doc_function: canonical
purpose: "Research method and collection protocol for R-XXX."
derived_from:
  - brief.md
  - ../../flows/research.md
status: draft
audience: humans_and_agents
---
```

## Instantiated Body

```markdown
# R-XXX: Research Plan

## Method

| Question / hypothesis | Method | Why this method fits | Quality threshold |
| --- | --- | --- | --- |
| `RQ-01` / `HYP-01` | `<interview, survey, prototype test, code spike, benchmark, desk research>` | `<reason>` | `<what makes evidence adequate>` |

## Sources or Sample

| Group / source | Inclusion and exclusion | Target / access boundary | Sampling limitation |
| --- | --- | --- | --- |

## Collection Protocol

Steps, instrument version, benchmark environment or query strategy sufficient for another reviewer to understand how evidence was obtained.

## Controls

| Risk | Control | Owner |
| --- | --- | --- |
| Bias / confounder | `<control>` | `<owner>` |
| Consent, privacy, legal or security | `<constraint or none>` | `<owner>` |
| Source freshness / vendor claim | `<control>` | `<owner>` |

## Stop Rules

- `STOP-01` `<reference brief stopping condition and any method-specific stop rule.>`

## Plan Approval

| Field | Value |
| --- | --- |
| Reviewer / decision owner | `<person or role>` |
| Approval reference | `<issue, comment or meeting>` |
```

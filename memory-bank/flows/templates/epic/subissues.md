---
title: "EP-XXX: Subissues Template"
doc_kind: governance
doc_function: template
purpose: "Шаблон registry for candidate and accepted delivery subissues under an epic."
derived_from:
  - ../../epic-flow.md
status: active
audience: humans_and_agents
template_target_path: ../../../epics/EP-XXX/subissues.md
---

# EP-XXX: Subissues Template

```markdown
---
title: "EP-XXX: Subissues"
doc_kind: epic
doc_function: subissue_registry
purpose: "<Delivery subissue registry for this epic>"
derived_from:
  - charter.md
  - roadmap.md
status: draft
audience: humans_and_agents
must_not_define:
  - code_steps
---

# EP-XXX: Subissues

## Registry

| ID | Candidate issue title | Roadmap wave | Source slices / UC | Status | Feature package |
| --- | --- | --- | --- | --- | --- |
| `EP-SI-01` |  |  |  | candidate | TBD |

## Creation Rules

- Create GitHub subissue only after scope is approved.
- Create `memory-bank/features/FT-<issue>/` after the issue exists.
- Link the feature package back to this epic and relevant source docs.
```

---
title: "EP-XXX: Charter Template"
doc_kind: governance
doc_function: template
purpose: "Шаблон epic charter: canonical intent, scope/non-scope, evidence and acceptance boundaries for a multi-feature initiative."
derived_from:
  - ../../epic-flow.md
status: active
audience: humans_and_agents
template_target_path: ../../../epics/EP-XXX/charter.md
---

# EP-XXX: Charter Template

```markdown
---
title: "EP-XXX: <Epic Name>"
doc_kind: epic
doc_function: canonical
purpose: "<What this epic governs and why it is larger than one feature>"
derived_from:
  - ../../flows/epic-flow.md
status: draft
audience: humans_and_agents
must_not_define:
  - implementation_sequence
  - feature_issue_ids_not_approved
---

# EP-XXX: <Epic Name>

## Problem

## Outcome

## Stakeholder Channels

| Channel | ID / URL | Purpose |
| --- | --- | --- |

## Scope

- `REQ-01`

## Non-Scope

- `NS-01`

## Source / Evidence Boundaries

| Source | Authority | Refresh rule |
| --- | --- | --- |

## Acceptance

| Criterion | Check |
| --- | --- |

## Handoff

Delivery work must be created as separate `memory-bank/features/FT-<issue>/` packages.
```

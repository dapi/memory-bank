---
title: R-XXX Research Decision Template
doc_kind: governance
doc_function: template
purpose: Wrapper-шаблон research disposition, recommendation and downstream promotion map.
derived_from:
  - ../../research.md
status: active
audience: humans_and_agents
template_for: research
template_target_path: ../../../research/R-XXX/decision.md
---

# R-XXX Research Decision Template

## Instantiated Frontmatter

```yaml
---
title: "R-XXX: Research Decision"
doc_kind: research
doc_function: canonical
purpose: "Decision disposition and promotion map for research R-XXX."
derived_from:
  - brief.md
  - synthesis.md
  - ../../flows/research.md
status: draft
research_disposition: pending
audience: humans_and_agents
---
```

## Instantiated Body

```markdown
# R-XXX: Research Decision

## Decision

| Field | Value |
| --- | --- |
| Decision owner | `<person or role>` |
| Decision date | `<YYYY-MM-DD>` |
| Disposition | `pending / validated / invalidated / inconclusive / parked / cancelled / rerouted` |
| Decision reference | `<issue comment, meeting note or approval>` |

## Recommendation

- `REC-01` `<Recommended action, including confidence and material limitation refs.>`

## Alternatives Considered

| Alternative | Why not selected / what would change the decision |
| --- | --- |

## Promotion and Handoff Map

| ID | Accepted or retained fact | Canonical downstream owner | Target route / link |
| --- | --- | --- | --- |
| `HD-01` | `<fact or recommendation>` | `<PRD, epic charter, feature brief/design, ADR, product context>` | `<link or route>` |

For `validated` delivery proposals, create or link the target owner and repeat Task Routing before implementation. For `inconclusive`, `parked` or `cancelled`, name owner and review trigger/next question. Do not leave this document as a duplicate active owner after promotion.

## Closure Check

- [ ] Disposition answers `RQ-01` or explicitly records why it cannot.
- [ ] Recommendation is traceable to `FND-*` and `LIM-*`.
- [ ] Handoff does not create delivery scope, implementation steps or an accepted architecture decision by implication.
```

---
title: R-XXX Research Decision Template
doc_kind: governance
doc_function: template
purpose: Wrapper-шаблон research decision rationale, recommendation and downstream promotion map.
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
purpose: "Decision rationale and promotion map for research R-XXX."
derived_from:
  - brief.md
  - ../../flows/research.md
  # Add synthesis.md only after the research reaches synthesis; omit it for an early-terminal decision.
status: draft
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
| Decision reference | `<issue comment, meeting note or approval>` |

Terminal disposition is recorded only in sibling `brief.md` as `research_status`. When finalizing this decision, set `brief.md` to the matching terminal state: `validated`, `invalidated`, `inconclusive`, `parked`, `cancelled` or `rerouted`.

## Decision Rationale

- `<Why this recommendation or disposition follows from the research; include the decision-relevant trade-offs.>`
- After synthesis, supporting findings: `FND-01` `<and other relevant FND-* references>`
- After synthesis, material limitations and residual uncertainty: `LIM-01` `<and other relevant LIM-* references>`
- For an early-terminal `parked`, `cancelled` or `rerouted` package, record the reason, retained evidence or consequences, and owner/review trigger or target route instead.

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

- [ ] Sibling `brief.md` records the matching terminal `research_status`.
- [ ] If research reached synthesis, `synthesis.md` answers `RQ-01` or explicitly records why it cannot; this decision links that answer through its recommendation and rationale.
- [ ] If research reached synthesis, the recommendation is traceable to `FND-*` and `LIM-*`; otherwise, the early-terminal reason or handoff is explicit.
- [ ] Handoff does not create delivery scope, implementation steps or an accepted architecture decision by implication.
```

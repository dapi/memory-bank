---
title: R-XXX Research Brief Template
doc_kind: governance
doc_function: template
purpose: "Wrapper-шаблон canonical research brief: decision question, hypotheses, boundaries and lifecycle state without findings or delivery design."
derived_from:
  - ../../research.md
  - ../../../dna/frontmatter.md
status: active
audience: humans_and_agents
template_for: research
template_target_path: ../../../research/R-XXX/brief.md
---

# R-XXX Research Brief Template

## Instantiated Frontmatter

```yaml
---
title: "R-XXX: <Research Name>"
doc_kind: research
doc_function: canonical
purpose: "Canonical decision question, boundaries and lifecycle state for research R-XXX."
derived_from:
  - ../../flows/research.md
status: draft
research_status: intake
audience: humans_and_agents
---
```

## Instantiated Body

```markdown
# R-XXX: <Research Name>

## Intake

| Field | Value |
| --- | --- |
| Source / trigger | `<issue, request, metric or observation>` |
| Research owner | `<person or role>` |
| Decision owner | `<person or role>` |
| Research mode | `market / product_discovery / technical_discovery / exploratory` |
| Decision deadline / timebox | `<date or duration>` |

## Decision Question

- `RQ-01` `<What decision needs evidence, in a form the named owner can answer?>`

## Working Hypotheses

- `HYP-01` `<Falsifiable claim; distinguish it from a fact.>`

## Compact Method Record (when `plan.md` is omitted)

- Method and source/sample strategy: `<bounded desk-research method and the sources or sample to collect>`
- Collection window and context: `<dates, freshness boundary and access context>`
- Evidence-quality criteria: `<what makes a source or observation sufficiently reliable and relevant>`
- Applicable privacy, consent, legal, security and vendor-access constraints: `<constraints or none>`
- Bias risks and disconfirming signal: `<likely bias and at least one result that could contradict the working hypothesis>`

Create `plan.md` instead when the method has a plan trigger in the research flow; keep this record concise and proportionate for compact desk research.

## Scope

- `RSC-01` `<Included question, audience, system or market boundary.>`

## Non-Scope

- `RNS-01` `<Explicit exclusion.>`

## Assumptions and Known Evidence

| ID | Statement | Type | Source / confidence |
| --- | --- | --- | --- |
| `ASM-01` | `<working assumption>` | Assumption | `<why it is currently reasonable>` |
| `<none or source>` | `<known fact>` | Evidence | `[SRC-XX](<source URL or stable internal link>)` |

## Stopping Condition

- `STOP-01` `<When collection ends: threshold, timebox, saturation, benchmark completion or explicit decision date.>`

## Open Questions

| Question | Blocks | Owner | Resolution evidence |
| --- | --- | --- | --- |

## Boundary Check

- [ ] This brief contains a question and hypotheses, not findings presented as facts.
- [ ] Every known fact has a clickable source link; unsupported statements remain assumptions or open questions.
- [ ] No committed delivery scope, selected solution, ADR decision or implementation sequence is defined here.
- [ ] Required privacy, consent, legal, security or access constraints are named or explicitly `none`.
```

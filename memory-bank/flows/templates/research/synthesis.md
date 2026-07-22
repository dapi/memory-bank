---
title: R-XXX Research Synthesis Template
doc_kind: governance
doc_function: template
purpose: Wrapper-шаблон synthesis of research findings, confidence, limitations and remaining uncertainty.
derived_from:
  - ../../research.md
status: active
audience: humans_and_agents
template_for: research
template_target_path: ../../../research/R-XXX/synthesis.md
---

# R-XXX Research Synthesis Template

## Instantiated Frontmatter

```yaml
---
title: "R-XXX: Research Synthesis"
doc_kind: research
doc_function: canonical
purpose: "Findings, confidence and limitations synthesized from evidence for R-XXX."
derived_from:
  - brief.md
  - evidence.md
status: draft
audience: humans_and_agents
---
```

## Instantiated Body

```markdown
# R-XXX: Research Synthesis

## Findings

| ID | Finding | Evidence | Confidence | Implication for `RQ-*` / `HYP-*` |
| --- | --- | --- | --- | --- |
| `FND-01` | `<synthesized conclusion>` | `[OBS-01](evidence.md#observations), [SRC-01](evidence.md#sources)` | `high / medium / low` | `<supports, contradicts or leaves unresolved>` |

## Limitations and Disconfirming Evidence

| ID | Limitation / conflicting signal | Effect on conclusion | Mitigation or next question |
| --- | --- | --- | --- |
| `LIM-01` | `<sampling, freshness, confounder, absence of evidence>` | `<confidence impact>` | `<action or accepted limit>` |

## Answer to Decision Question

Answer `RQ-01` in direct language. If evidence is insufficient, say so; do not convert an uncertain inference into a fact.

## Review Check

- [ ] Every finding and factual claim traces through linked `OBS-*` to linked `SRC-*`; do not state uncited facts as findings.
- [ ] Confidence reflects evidence quality rather than desired outcome.
- [ ] Alternative explanations and remaining uncertainty are visible.
```

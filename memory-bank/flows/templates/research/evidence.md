---
title: R-XXX Research Evidence Template
doc_kind: governance
doc_function: template
purpose: Wrapper-шаблон provenance-preserving evidence and observation log for research R-XXX.
derived_from:
  - ../../research.md
status: active
audience: humans_and_agents
template_for: research
template_target_path: ../../../research/R-XXX/evidence.md
---

# R-XXX Research Evidence Template

## Instantiated Frontmatter

```yaml
---
title: "R-XXX: Evidence Log"
doc_kind: research
doc_function: canonical
purpose: "Traceable evidence and observations collected for research R-XXX."
derived_from:
  - brief.md
  - ../../flows/research.md
status: draft
audience: humans_and_agents
---
```

Если в package существует `plan.md`, добавь его в `derived_from`. Для compact desk research без плана оставь frontmatter выше без этой зависимости.

## Instantiated Body

```markdown
# R-XXX: Evidence Log

Do not copy restricted source material, personal data or credentials here. Record a minimal reference, access boundary and derived observation.

## Sources

| ID | Source / provenance | Date / freshness | Collection context | Access / quality note |
| --- | --- | --- | --- | --- |
| `SRC-01` | `[<source title>](<source URL or stable access-controlled record>)` | `<date>` | `<how obtained>` | `<primary/secondary, limitation, access boundary>` |

## Observations

| ID | Observation | Supporting `SRC-*` | Applies to | Interpretation boundary |
| --- | --- | --- | --- | --- |
| `OBS-01` | `<directly observed/measured claim>` | `[SRC-01](<source URL or stable access-controlled record>)` | `RQ-01 / HYP-01` | `<what this does not establish>` |

## Collection Log

| Date | Activity | Result | Deviation / reason |
| --- | --- | --- | --- |

## Evidence Quality Check

- [ ] Each material observation traces to one or more `SRC-*`.
- [ ] Every `SRC-*` contains a clickable link to its original source or a stable access-controlled source record; an interview code alone is insufficient.
- [ ] Observations are separated from source claims and analyst interpretation.
- [ ] Freshness, sample/source limitations and conflicts are recorded.
```

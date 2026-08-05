---
title: "R-GH-87: Research Synthesis"
doc_kind: research
doc_function: canonical
purpose: "Findings, confidence and limitations synthesized from evidence for R-GH-87."
derived_from:
  - brief.md
  - evidence.md
status: active
audience: humans_and_agents
---

# R-GH-87: Research Synthesis

## Findings

| ID | Finding | Evidence | Confidence | Implication for `RQ-*` / `HYP-*` |
| --- | --- | --- | --- | --- |
| `FND-01` | Reusable `lint` and `doctor` implementation was moved out of this repository into `memory-bank-cli`. | [OBS-01](evidence.md#observations), [OBS-03](evidence.md#observations) | High | Supports `HYP-01`; no local Go CLI remains to remove. |
| `FND-02` | Remaining Ruby priming-manifest checks are an explicit repository CI surface run alongside, not through, the installed CLI. | [OBS-02](evidence.md#observations) | High | Supports retaining the checks under this repository's ownership. |
| `FND-03` | The migration task that established the standalone CLI is already completed, and current main CI is green. | [OBS-03](evidence.md#observations), [OBS-04](evidence.md#observations) | Medium | Does not justify a duplicate CLI issue without a new CLI requirement. |

## Limitations and Disconfirming Evidence

| ID | Limitation / conflicting signal | Effect on conclusion | Mitigation or next question |
| --- | --- | --- | --- |
| `LIM-01` | This is a bounded source review, not a semantic redesign of the Ruby validation. | The recommendation does not claim the scripts can never move. | Route a new technical discovery or delivery issue only if a concrete reusable CLI contract is proposed. |
| `LIM-02` | Green CI proves execution on current main, not every possible regression case. | Confidence in ownership is unaffected; coverage adequacy remains outside this issue. | Evaluate coverage independently if a future issue identifies a gap. |

## Answer to Decision Question

No additional removal is supported. The reusable validation requested by the
issue is already in `dapi/memory-bank-cli`; the remaining Ruby checks validate
this repository's priming-manifest producer and CI integration. A new CLI task
would duplicate completed migration work unless a separate requirement defines
a reusable CLI contract for those checks.

## Review Check

- [x] Every finding traces through linked observations to linked sources.
- [x] Confidence reflects source quality rather than the desired outcome.
- [x] Alternative explanation and residual uncertainty are visible.

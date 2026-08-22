---
title: Product Metrics
doc_kind: product
doc_function: canonical
purpose: Каноничное место для product success metrics, baselines, targets, measurement ownership и instrumentation constraints.
derived_from:
  - ../dna/governance.md
  - context.md
status: active
audience: humans_and_agents
canonical_for:
  - product_metrics
  - success_measurement
---

# Product Metrics

Этот документ фиксирует метрики продукта и правила их измерения. Feature-level checks и test evidence остаются в feature package; здесь живут только product-level outcomes и measurement contract.

## North Star

| Metric ID | Metric | Why it matters | Current baseline | Target | Review cadence |
| --- | --- | --- | --- | --- | --- |
| `NSM-01` | Verified external adoptions | Shows that another repository installed, adapted and successfully validated Memory Bank | `unknown`; no canonical registry yet | Establish an evidence-backed baseline, then set a target | Monthly |

## Product Metrics

| Metric ID | Metric | Owner | Baseline | Target | Measurement method | Source |
| --- | --- | --- | --- | --- | --- | --- |
| `MET-01` | GitHub stars | Maintainer | 15 on 2026-08-14 | 100 by 2026-09-13 | GitHub repository metadata snapshot | `dapi/memory-bank` |
| `MET-02` | GitHub forks | Maintainer | 4 on 2026-08-14 | Observe; no vanity target | GitHub repository metadata snapshot | `dapi/memory-bank` |
| `MET-03` | Qualified external adoption signals | Maintainer | `unknown` | At least 10 during the launch cycle | Issues, discussions or public repositories showing attempted installation or concrete evaluation | Evidence log to be created before launch |
| `MET-04` | Substantive external feedback | Maintainer | `unknown` | At least 5 concrete questions, issues or review comments during the launch cycle | GitHub Issues / Discussions and linked public conversations | Manual monthly review |

## Guardrails

| Guardrail ID | Metric | Why it must not regress | Threshold | Response |
| --- | --- | --- | --- | --- |
| `GR-01` | `memory-bank-cli doctor` health | Promotion must not outrun product reliability | Any new release-blocking error in the documented quick start | Pause promotion and fix or explicitly document the failure |
| `GR-02` | Unsupported public claims | Trust is more valuable than a short-term star spike | Any material claim without a linked source or explicit evidence label | Correct the claim before further distribution |
| `GR-03` | Low-quality acquisition | Paid or reciprocal stars hide whether the project solves a real problem | Any paid-star or star-exchange activity | Stop the channel and exclude it from reporting |

## Instrumentation Constraints

- `ICON-01` GitHub stars and forks are publicly observable but do not prove installation, retention or value.
- `ICON-02` Repository traffic is available only for a limited rolling window and may be unavailable because of API permissions or rate limits; absence of a snapshot is `unknown`, not zero.
- `ICON-03` An adoption signal is counted only when it identifies an external evaluator and a concrete installation, evaluation or integration attempt; likes and generic praise do not count.

## Metric Change Policy

- Не меняй definition метрики внутри feature package без обновления этого документа или upstream PRD.
- Если feature вводит новую локальную метрику, держи ее в feature package до тех пор, пока она не станет shared product metric.

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
| `NSM-01` | Главная метрика продукта | Почему она отражает value | Текущее значение или `unknown` | Целевое значение | Как часто пересматриваем |

## Product Metrics

| Metric ID | Metric | Owner | Baseline | Target | Measurement method | Source |
| --- | --- | --- | --- | --- | --- | --- |
| `MET-01` | Что измеряем | Кто владеет | От чего стартуем | Что считаем успехом | Как считаем | Dashboard / query / manual |

## Guardrails

| Guardrail ID | Metric | Why it must not regress | Threshold | Response |
| --- | --- | --- | --- | --- |
| `GR-01` | Что защищаем | Почему важно | Порог | Что делаем при регрессе |

## Instrumentation Constraints

- `ICON-01` Какое событие, dashboard или data source считается canonical.
- `ICON-02` Какая задержка, sampling, privacy rule или attribution limit влияет на интерпретацию.

## Metric Change Policy

- Не меняй definition метрики внутри feature package без обновления этого документа или upstream PRD.
- Если feature вводит новую локальную метрику, держи ее в feature package до тех пор, пока она не станет shared product metric.

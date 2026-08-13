---
title: "FT-117: Autonomous Structured Decision Protocol"
doc_kind: feature
doc_function: index
purpose: "Навигация по project-level feature package issue #117: автономное принятие bounded решений с отдельной проверкой permission на execution gate."
derived_from:
  - ../../dna/governance.md
  - ../../flows/feature.md
  - brief.md
status: active
audience: humans_and_agents
---

# FT-117: Autonomous Structured Decision Protocol

## О разделе

Feature package фиксирует problem space, selected solution и execution plan для
issue [#117](https://github.com/dapi/memory-bank/issues/117). Scope ограничен
governed Memory Bank documentation: package не добавляет runtime-код, production
configuration или implicit permission на внешние действия.

## Аннотированный индекс

- [`brief.md`](brief.md)
  Читать первым: canonical problem, scope, acceptance scenarios, validation
  profile и verify contract.

- [`design.md`](design.md)
  Читать после brief: selected cross-flow governance solution, ownership,
  C4/applicability и 4+1 coverage.

- [`implementation-plan.md`](implementation-plan.md)
  Читать перед execution: grounded paths, document changes, checks,
  checkpoints и stop conditions.

- [`feature-review-report.md`](feature-review-report.md)
  Читать для semantic self-review, acceptance coverage и результатов
  deterministic documentation checks.

---
title: Flows And Templates Index
doc_kind: governance
doc_function: index
purpose: Навигация по task routing, lifecycle flows и governed-шаблонам. Читать при выборе route, запуске flow или инстанцировании governed-документа.
derived_from:
  - ../dna/governance.md
  - routing.md
  - incident.md
  - bug-fix.md
  - small-change.md
  - refactoring.md
  - epic.md
  - feature.md
  - feature-artifact-catalog.md
  - workflow-decision-log.md
  - workflow-metrics.md
  - workflow-routing-developer-brief.md
  - templates/README.md
status: active
audience: humans_and_agents
---

# Flows And Templates Index

Каталог `memory-bank/flows/` содержит reusable process-layer для шаблона: lifecycle rules, taxonomy стабильных идентификаторов и governed templates.

- [Task Routing](routing.md) — порядок выбора flow, routing predicates, повторный routing и Human Routing.
- [Incident And PIR Flow](incident.md) — containment, recovery, timeline, RCA, PIR и prevention work.
- [Bug Fix Flow](bug-fix.md) — reproduction, analysis, fix, regression coverage и closure.
- [Small Change Flow](small-change.md) — direct delivery без feature package, design и execution plan, но с обязательным routing record.
- [Refactoring Flow](refactoring.md) — behavior-preserving restructuring, characterization coverage, checkpoints и closure gates.
- [Epic Flow](epic.md) — lifecycle крупных инициатив, roadmap, decision log, risks и handoff в feature packages.
- [Feature Flow](feature.md) — lifecycle `brief.md -> optional design.md -> implementation-plan.md`, gates и стабильные ID (`REQ-*`, `SOL-*`, `STEP-*`).
- [Feature Artifact Catalog](feature-artifact-catalog.md) — optional problem/solution/execution artifacts, selection triggers, ownership, default forms и template availability.
- [Workflow Decision Log](workflow-decision-log.md) — optional журнал причин и ожидаемых последствий изменений workflow selector-а и profiles.
- [Workflow Metrics](workflow-metrics.md) — optional safety-first метрики для проверки workflow profiles без подмены evidence скоростью.
- [Workflow Routing Developer Brief](workflow-routing-developer-brief.md) — короткий reusable brief для выбора workflow profile и promotion triggers.
- [Templates Index](templates/README.md) — эталонные шаблоны governed-документов, включая PRD, use case, epic, feature и ADR.

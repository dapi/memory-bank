---
title: Flows And Templates Index
doc_kind: governance
doc_function: index
purpose: Навигация по lifecycle flows и governed-шаблонам. Читать при создании epic или feature package, переводе инициативы между стадиями или инстанцировании нового governed-документа.
derived_from:
  - ../dna/governance.md
  - epic-flow.md
  - feature-flow.md
  - workflows.md
  - workflow-decision-log.md
  - workflow-metrics.md
  - workflow-routing-developer-brief.md
  - templates/README.md
status: active
audience: humans_and_agents
---

# Flows And Templates Index

Каталог `memory-bank/flows/` содержит reusable process-layer для шаблона: lifecycle rules, taxonomy стабильных идентификаторов и governed templates.

- [Task Workflows](workflows.md) — маршрутизация задач по типам, базовый цикл разработки и градиент автономии.
- [Workflow Decision Log](workflow-decision-log.md) — optional журнал причин и ожидаемых последствий изменений workflow selector-а и profiles.
- [Workflow Metrics](workflow-metrics.md) — optional safety-first метрики для проверки workflow profiles без подмены evidence скоростью.
- [Workflow Routing Developer Brief](workflow-routing-developer-brief.md) — короткий reusable brief для выбора workflow profile и promotion triggers.
- [Epic Flow](epic-flow.md) — lifecycle крупных инициатив, roadmap, decision log, risks и handoff в feature packages.
- [Feature Flow](feature-flow.md) — lifecycle `brief.md -> optional design.md -> implementation-plan.md`, gates и стабильные ID (`REQ-*`, `SOL-*`, `STEP-*`).
- [Templates Index](templates/README.md) — эталонные шаблоны governed-документов, включая PRD, use case, epic, feature и ADR.

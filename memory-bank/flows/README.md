---
title: Flows And Templates Index
doc_kind: governance
doc_function: index
purpose: Навигация по lifecycle flows и governed-шаблонам. Читать при создании task, epic или feature package, переводе инициативы между стадиями или инстанцировании нового governed-документа.
derived_from:
  - ../dna/governance.md
  - task-flow.md
  - bugfix-flow.md
  - refactor-flow.md
  - epic-flow.md
  - feature-flow.md
  - workflows.md
  - templates/README.md
status: active
audience: humans_and_agents
---

# Flows And Templates Index

Каталог `memory-bank/flows/` содержит reusable process-layer для шаблона: task/feature/epic lifecycle rules, taxonomy стабильных идентификаторов и governed templates.

- [Task Workflows](workflows.md) — маршрутизация задач по типам, базовый цикл разработки и градиент автономии.
- [Task Flow](task-flow.md) — compact/managed non-feature lifecycle для `tracker-only`, bugfix/refactor/chore profiles и `TASK-XXX` packages.
- [Bugfix Flow](bugfix-flow.md) — профильный процесс symptom/reproduction/root cause/fix boundary/regression evidence для compact bugfix задач.
- [Refactor Flow](refactor-flow.md) — профильный процесс intent/invariants/change surface/checkpoints/verification для refactor/chore задач без изменения поведения.
- [Epic Flow](epic-flow.md) — lifecycle крупных инициатив, roadmap, decision log, risks и handoff в feature packages.
- [Feature Flow](feature-flow.md) — lifecycle `brief.md -> optional design.md -> implementation-plan.md`, gates и стабильные ID (`REQ-*`, `SOL-*`, `STEP-*`).
- [Templates Index](templates/README.md) — эталонные шаблоны governed-документов, включая task, PRD, use case, epic, feature и ADR.

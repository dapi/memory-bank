---
title: Flows And Templates Index
doc_kind: governance
doc_function: index
purpose: Навигация по task routing, lifecycle flows и governed-шаблонам. Читать при выборе route, запуске flow или инстанцировании governed-документа.
derived_from:
  - ../dna/governance.md
  - routing.md
  - autonomy-boundaries.md
  - validation-profiles.md
  - testing-policy.md
  - priming/README.md
  - research.md
  - incident.md
  - bug-fix.md
  - small-change.md
  - refactoring.md
  - epic.md
  - behavior-specification.md
  - use-case.md
  - feature.md
  - feature-requirements.md
  - execution-handoff.md
  - feature-artifact-catalog.md
  - templates/README.md
status: active
audience: humans_and_agents
---

# Flows And Templates Index

Каталог `memory-bank/flows/` содержит reusable process-layer для шаблона: cross-flow policy, lifecycle rules, taxonomy стабильных идентификаторов и governed templates. Это generic слой: он описывает, как ведётся разработка, и не зависит от стека целевой системы. Инженерия самой целевой системы живёт в [`../engineering/README.md`](../engineering/README.md).

Граница с [DNA](../dna/README.md): DNA задаёт критерии достоверности,
ownership и актуальности документов; этот слой задаёт routing, действия,
полномочия, priming, проверки и handoff. Документы flows сами подчиняются DNA,
но качество документа и разрешение исполнить описанное действие — разные вопросы.

## Cross-Flow Policy

Действует во всех flows и выбирается до или внутри entry gate конкретного route.

- [Autonomy Boundaries](autonomy-boundaries.md) — autonomous-by-default правила, самодостаточный Structured Decision Protocol, вдохновлённый FPF, и настоящие Human Gates. Отвечает на вопрос: что агент решает и исполняет сам, а где нужны человеческие полномочия или принятие неконтролируемого риска.
- [Validation Profiles](validation-profiles.md) — независимая от delivery flow глубина validation: taxonomy, risk triggers, minimum evidence contract и canonical owner решения.
- [Testing Policy](testing-policy.md) — правила тестирования, обязательные automated tests, sufficient coverage. Отвечает на вопрос: когда feature обязана иметь test cases и когда допустим manual-only verify. Project-specific стек живёт в [`../engineering/testing-conventions.md`](../engineering/testing-conventions.md).

## Routing And Flows

- [Task Routing](routing.md) — порядок выбора flow, routing predicates, повторный routing и Human Routing только после Structured Decision Protocol outcome `escalate`.
- [Task Context Priming](priming/README.md) — общий P0/P1/P2 contract, universal DNA baseline и per-process YAML manifests.
- [Execution Handoff Contract](execution-handoff.md) — compact read-only projection
  observed execution с direct primary-source references для continuation одной задачи.
- [Research & Discovery Flow](research.md) — evidence-backed lifecycle research-задач, от question framing до decision и handoff без преждевременного delivery.
- [Incident And PIR Flow](incident.md) — containment, recovery, timeline, RCA, PIR и prevention work.
- [Bug Fix Flow](bug-fix.md) — reproduction, analysis, fix, regression coverage и closure.
- [Small Change Flow](small-change.md) — direct delivery без feature package, design и execution plan, но с обязательным routing record.
- [Refactoring Flow](refactoring.md) — behavior-preserving restructuring, characterization coverage, checkpoints и closure gates.
- [Epic Flow](epic.md) — Epic Intake/Proposal, lifecycle крупных инициатив, roadmap, decision log, risks и handoff в feature packages.
- [Behavior Specification Practice](behavior-specification.md) — BDD-цикл Discovery → Formulation → Automation, качество concrete examples и traceability `UC/REQ → SC/NEG → CHK → test evidence`; это практика внутри выбранного flow, а не отдельный route.
- [Use Case Flow](use-case.md) — критерии, lifecycle и ownership для project-level `UC-*`, включая operational / agentic сценарии.
- [Feature Flow](feature.md) — lifecycle `brief.md -> optional design.md -> implementation-plan.md` и transition gates.
- [Feature Requirements, Identifiers And Traceability](feature-requirements.md) — requirement classes, stable IDs, applicability и двусторонняя трассировка до delivered surfaces и evidence.
- [Feature Artifact Catalog](feature-artifact-catalog.md) — optional problem/solution/execution artifacts, selection triggers, ownership, default forms и template availability.
- [Templates Index](templates/README.md) — эталонные шаблоны governed-документов, включая PRD, use case, epic, feature и ADR.

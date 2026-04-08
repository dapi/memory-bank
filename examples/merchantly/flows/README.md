---
title: Flows & Templates Index
doc_kind: governance
doc_function: index
purpose: Навигация по kind-specific lifecycle flows и governed-шаблонам. Читать при создании feature package, переводе feature между стадиями или инстанцировании нового governed-документа.
derived_from:
  - ../dna/governance.md
  - feature-flow.md
  - workflows.md
  - templates/README.md
status: active
audience: humans_and_agents
---

# Flows & Templates Index

Каталог `memory-bank/flows/` содержит kind-specific lifecycle rules и governed-шаблоны документации. В отличие от `memory-bank/dna/` (kernel — meta-правила, переносимые на любой проект), этот раздел определяет правила для конкретных `doc_kind` и предоставляет шаблоны для инстанцирования governed-документов.

- [Task Workflows](workflows.md) — маршрутизация задач по типам, базовый цикл разработки, градиент автономии. Отвечает на вопрос: пришла задача — какой workflow применить.
- [Feature Flow](feature-flow.md) — lifecycle от draft до closure, gates, стабильные ID (REQ-\*, CHK-\*, STEP-\*). Отвечает на вопрос: как feature проходит стадии и какой ID для чего.
- [Templates Index](templates/README.md) — эталонные шаблоны governed-документов (feature, plan, ADR). Отвечает на вопрос: какие шаблоны есть и где их брать.
- ADR flow описан в [Feature Flow](feature-flow.md) (секция Boundary Rules, п.3). Шаблон ADR: [templates/adr/ADR-XXX.md](templates/adr/ADR-XXX.md).

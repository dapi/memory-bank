---
title: Template Documentation Index
doc_kind: project
doc_function: index
purpose: Корневая навигация по шаблонному memory-bank. Читать сначала, чтобы понять структуру и точки адаптации под конкретный проект.
derived_from:
  - dna/principles.md
  - dna/governance.md
  - flows/task-flow.md
status: active
audience: humans_and_agents
---

# Documentation Index

Каталог `memory-bank/` содержит переносимый шаблон проектной документации для разработки ПО. После копирования в downstream-репозиторий адаптируй `product/`, `domain/`, `engineering/` и `ops/` под реальный продукт, предметную область, стек, процессы и ограничения проекта.

## Аннотированный индекс

- [`product/README.md`](product/README.md)
  Читать, когда нужно: зафиксировать product context, vision, customers, metrics, marketing и roadmap.

- [`domain/README.md`](domain/README.md)
  Читать, когда нужно: зафиксировать glossary, domain model, rules, states, events и bounded contexts.

- [`prd/README.md`](prd/README.md)
  Читать, когда нужно: описать продуктовую инициативу между общим product context и downstream feature packages.

- [`epics/README.md`](epics/README.md)
  Читать, когда нужно: вести крупную инициативу через roadmap, decision log, risks и набор связанных delivery subissues.

- [`tasks/README.md`](tasks/README.md)
  Читать, когда нужно: вести managed non-feature bugfix/chore/refactor задачу в compact `TASK-XXX/` package без full feature package.

- [`use-cases/README.md`](use-cases/README.md)
  Читать, когда нужно: зарегистрировать устойчивый пользовательский или операционный сценарий проекта.

- [`prompts/README.md`](prompts/README.md)
  Читать, когда нужно: найти или завести reusable prompt-документ с исходной формулировкой и copyable улучшенной версией.

- [`ops/README.md`](ops/README.md)
  Читать, когда нужно: описать локальную разработку, окружения, релизы, конфигурацию и runbooks.

- [`engineering/README.md`](engineering/README.md)
  Читать, когда нужно: задать architecture patterns, frontend rules, testing policy, coding style, git workflow и границы автономии агента.

- [`dna/README.md`](dna/README.md)
  Читать, когда нужно: проверить SSoT rules, frontmatter contract и governance-правила документации.

- [`flows/README.md`](flows/README.md)
  Читать, когда нужно: создать epic/feature package, провести инициативу по lifecycle gates или использовать шаблон.

- [`adr/README.md`](adr/README.md)
  Читать, когда нужно: найти или завести Architecture Decision Record.

- [`features/README.md`](features/README.md)
  Читать, когда нужно: понять, где живут instantiated feature packages.

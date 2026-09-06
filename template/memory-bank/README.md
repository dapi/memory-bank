---
title: Template Documentation Index
doc_kind: project
doc_function: index
purpose: Корневая навигация по шаблонному memory-bank. Читать сначала, чтобы понять структуру и точки адаптации под конкретный проект.
derived_from:
  - dna/principles.md
  - dna/governance.md
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

- [`research/README.md`](research/README.md)
  Читать, когда нужно: провести evidence-backed market, product или technical research до коммита в delivery и передать вывод в подходящий canonical owner.

- [`epics/README.md`](epics/README.md)
  Читать, когда нужно: вести крупную инициативу через roadmap, decision log, risks и набор связанных delivery subissues.

- [`use-cases/README.md`](use-cases/README.md)
  Читать, когда нужно: зарегистрировать устойчивый пользовательский или операционный сценарий проекта.

- [`prompts/README.md`](prompts/README.md)
  Human-only каталог reusable prompt-артефактов и его canonical access contract.

- [`ops/README.md`](ops/README.md)
  Читать, когда нужно: описать локальную разработку, окружения, релизы, конфигурацию и runbooks.

- [`engineering/README.md`](engineering/README.md)
  Читать, когда нужно: задать architecture patterns, frontend rules, testing conventions, coding style и git workflow целевой системы.

- [`dna/README.md`](dna/README.md)
  Читать, когда нужно: проверить достоверность, основания и актуальность утверждений, ownership и metadata документации.

- [`flows/README.md`](flows/README.md)
  Читать, когда нужно: создать use case, epic/feature package, применить BDD-практику, провести артефакт по lifecycle gates, узнать границы автономии агента, выбрать validation profile или использовать шаблон.

- [`flows/execution-handoff.md`](flows/execution-handoff.md)
  Читать, когда нужно: безопасно продолжить одну конкретную задачу по compact,
  read-only и evidence-backed проекции наблюдаемого исполнения.

- [`adr/README.md`](adr/README.md)
  Читать, когда нужно: найти или завести Architecture Decision Record.

- [`features/README.md`](features/README.md)
  Читать, когда нужно: понять, где живут instantiated feature packages.

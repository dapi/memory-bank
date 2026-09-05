---
title: Project Memory Bank Index
doc_kind: project
doc_function: index
purpose: Корневая навигация по project-local Memory Bank репозитория dapi/memory-bank.
derived_from:
  - dna/principles.md
  - dna/governance.md
status: active
audience: humans_and_agents
---

# Project Memory Bank Index

Этот каталог — project-local Memory Bank репозитория `dapi/memory-bank`. Он
установлен из generic payload в [`../template/memory-bank/`](../template/memory-bank/)
и является canonical местом для project-specific документов этого репозитория.
В частности, новые delivery feature packages создаются только в
[`features/`](features/README.md), а не в upstream payload.

Источник, закрепленная версия и сознательные границы начальной адаптации
зафиксированы в [`bootstrap.md`](bootstrap.md). Generic правила остаются в
`template/memory-bank/`; адаптируй только эту project-local копию.

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
  Читать, когда нужно: проверить SSoT rules, frontmatter contract и governance-правила документации.

- [`flows/README.md`](flows/README.md)
  Читать, когда нужно: создать use case, epic/feature package, применить BDD-практику, провести артефакт по lifecycle gates, узнать границы автономии агента, выбрать validation profile или использовать шаблон.

- [`flows/execution-handoff.md`](flows/execution-handoff.md)
  Читать, когда нужно: безопасно продолжить одну конкретную задачу по compact,
  read-only и evidence-backed проекции наблюдаемого исполнения.

- [`adr/README.md`](adr/README.md)
  Читать, когда нужно: найти или завести Architecture Decision Record.

- [`features/README.md`](features/README.md)
  Читать, когда нужно: понять, где живут instantiated feature packages.

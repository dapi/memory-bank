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

- [`engineering/architecture.md`](engineering/architecture.md)
  Читать, когда нужно: проверить code/module boundaries, runtime patterns, error handling и configuration ownership.

- [`engineering/frontend.md`](engineering/frontend.md)
  Читать, когда нужно: проверить frontend stack, component boundaries, design system integration и i18n.

- [`engineering/ui-design-guide/README.md`](engineering/ui-design-guide/README.md)
  Читать, когда нужно: найти project-level UI references по существующим surfaces и shared components.

- [`engineering/testing-policy.md`](engineering/testing-policy.md)
  Читать, когда нужно: определить обязательные test surfaces, regression coverage и допустимые manual-only gaps.

- [`engineering/validation-profiles.md`](engineering/validation-profiles.md)
  Читать, когда нужно: выбрать глубину validation и minimum evidence contract независимо от delivery flow.

- [`engineering/autonomy-boundaries.md`](engineering/autonomy-boundaries.md)
  Читать, когда нужно: проверить границы автономии, supervision checkpoints и escalation triggers.

- [`engineering/coding-style.md`](engineering/coding-style.md)
  Читать, когда нужно: проверить coding conventions, tooling contract и change discipline.

- [`dna/README.md`](dna/README.md)
  Читать, когда нужно: проверить SSoT rules, frontmatter contract и governance-правила документации.

- [`flows/README.md`](flows/README.md)
  Читать, когда нужно: создать use case, epic/feature package, провести артефакт по lifecycle gates или использовать шаблон.

- [`adr/README.md`](adr/README.md)
  Читать, когда нужно: найти или завести Architecture Decision Record.

- [`features/README.md`](features/README.md)
  Читать, когда нужно: понять, где живут instantiated feature packages.

- [`bootstrap.md`](bootstrap.md)
  Читать, когда нужно: проверить происхождение локальной копии и решения,
  принятые при её bootstrap.

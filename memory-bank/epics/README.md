---
title: Epics Index
doc_kind: epic
doc_function: index
purpose: "Навигация по instantiated epic packages. Читать, когда инициатива крупнее одной feature и должна исполняться через roadmap и набор связанных subissues."
derived_from:
  - ../dna/governance.md
  - ../flows/epic-flow.md
  - ../flows/feature-flow.md
status: active
audience: humans_and_agents
---

# Epics Index

Каталог `memory-bank/epics/` хранит instantiated epic packages вида `EP-XXX/`.

## Rules

- Epic описывает крупное проектное изменение, которое нельзя безопасно реализовать одной delivery-feature.
- Epic владеет intent, roadmap, декомпозицией, decision log, рисками и реестром subissues.
- Epic не владеет code-level execution: реализация идёт через отдельные `memory-bank/features/FT-<issue>/` packages.
- Каждый delivery subissue должен ссылаться на соответствующие epic artifacts и project-level `UC-*`, если меняет устойчивый сценарий.
- Правила создания и ведения epic packages живут в [`../flows/epic-flow.md`](../flows/epic-flow.md).

## Naming

- Базовый формат: `EP-XXX/`
- Вместо `XXX` используй стабильный идентификатор инициативы: issue id, project id или другое устойчивое имя
- Один epic = одна крупная программа/инициатива с несколькими delivery-slices

## Package Layers

| Layer | Files | Purpose |
| --- | --- | --- |
| Intent | `charter.md`, source refs, stakeholder channels | Зачем существует epic, что входит/не входит, какие facts уже подтверждены |
| Governance | `roadmap.md`, `decision-log.md`, `risks.md`, `subissues.md` | Как исполнять epic, какие решения приняты, какие риски и subissues управляются |
| Knowledge | `design.md`, `specs/**`, `diagrams/**`, linked `UC-*` | Нормализованные требования, bounded contexts, сценарии, контракты и audit trail |
| Execution Handoff | future `memory-bank/features/FT-<issue>/` | Конкретные code changes, тесты, rollout/backout для одного approved delivery issue |

Knowledge-файлы опциональны. Если они создаются как Markdown внутри epic package, они должны быть индексированы из package `README.md` или owner-документа и следовать правилам frontmatter из [`../flows/epic-flow.md`](../flows/epic-flow.md).

## Instantiated Epics

В шаблонном репозитории этот каталог может быть пустым. Это нормально.

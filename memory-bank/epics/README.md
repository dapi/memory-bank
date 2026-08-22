---
title: Epics Index
doc_kind: epic
doc_function: index
purpose: "Навигация по instantiated epic packages. Читать, когда инициатива крупнее одной feature и должна исполняться через roadmap и набор связанных subissues."
derived_from:
  - ../dna/governance.md
  - ../flows/epic.md
  - ../flows/feature.md
status: active
audience: humans_and_agents
---

# Epics Index

Каталог `memory-bank/epics/` хранит instantiated epic packages вида `EP-XXX/`.

## Rules

- Epic описывает крупное проектное изменение, которое нельзя безопасно реализовать одной delivery-feature.
- Если Epic route выбран до готовности canonical charter, package начинается с Epic Intake: `README.md` + обязательный `brief.md` в состоянии Epic Proposal. `brief.md` можно не создавать только при пропуске Intake и прямом Bootstrap Epic.
- Epic владеет intent, roadmap, декомпозицией, decision log, рисками и реестром subissues.
- Epic не владеет code-level execution: реализация идёт через отдельные `memory-bank/features/FT-<issue>/` packages.
- Каждый delivery subissue должен ссылаться на соответствующие epic artifacts и project-level `UC-*`, если меняет устойчивый сценарий.
- Правила создания и ведения epic packages живут в [`../flows/epic.md`](../flows/epic.md).

## Naming

- Базовый формат: `EP-XXX/`
- Вместо `XXX` используй стабильный идентификатор инициативы: issue id, project id или другое устойчивое имя
- Один epic = одна крупная программа/инициатива с несколькими delivery-slices

## Package Layers

| Layer | Files | Purpose |
| --- | --- | --- |
| Intake | `README.md`, required `brief.md` | Текущая `epic_stage`, proposal facts, open questions и disposition до canonical setup |
| Intent | `charter.md`, source refs, stakeholder channels | Зачем существует epic, что входит/не входит, какие facts уже подтверждены |
| Governance | `roadmap.md`, `decision-log.md`, `risks.md`, `subissues.md` | Как исполнять epic, какие решения приняты, какие риски и subissues управляются |
| Knowledge | `design.md`, `specs/**`, `diagrams/**`, linked `UC-*` | Нормализованные требования, bounded contexts, сценарии, контракты и audit trail |
| Execution Handoff | future `memory-bank/features/FT-<issue>/` | Конкретные code changes, тесты, rollout/backout для одного approved delivery issue |

`README.md` обязателен с начала package и индексирует только реально существующие документы. `brief.md` обязателен при выборе Epic Intake и отсутствует только при прямом Bootstrap Epic; knowledge-файлы опциональны. Любой Markdown внутри epic package должен быть reachable из package `README.md` или owner-документа и следовать правилам frontmatter из [`../flows/epic.md`](../flows/epic.md).

## Instantiated Epics

В шаблонном репозитории этот каталог может быть пустым. Это нормально.

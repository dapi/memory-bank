---
title: Architecture Decision Records Index
doc_kind: adr
doc_function: index
purpose: Навигация по ADR проекта. Читать, чтобы найти уже принятые решения или завести новый ADR по шаблону.
derived_from:
  - ../dna/governance.md
  - ../flows/templates/adr/ADR-XXX.md
status: active
audience: humans_and_agents
---

# Architecture Decision Records Index

Каталог `memory-bank/adr/` хранит instantiated ADR проекта.

- Заводи новый ADR из шаблона [`../flows/templates/adr/ADR-XXX.md`](../flows/templates/adr/ADR-XXX.md).
- Держи в этом каталоге только реальные decision records, а не заметки или черновые исследования.

## Аннотированный индекс

- [`ADR-001-introduce-design-pack.md`](ADR-001-introduce-design-pack.md)
  Proposal: разделить semantic design layer, documentary design pack и root
  `design.md`, а также закрепить aggregate и direct ownership solution facts.

## Naming

- Формат файла: `ADR-XXX-short-decision-name.md`
- Нумерация монотонная и не переиспользуется
- Заголовок файла должен совпадать с `title` во frontmatter

## Statuses

Публикационный `status` и lifecycle решения `decision_status` независимы:

- документ в работе: `status: draft`, `decision_status: proposed`;
- предложение готово к review: `status: active`, `decision_status: proposed`;
- принятое решение: `status: active`, `decision_status: accepted`;
- отклонённое решение: `status: active`, `decision_status: rejected`;
- заменённое решение: `status: active`, `decision_status: superseded`.

Только `active` + `accepted` является принятым canonical input для downstream
owners. `active` + `proposed` публикует reviewable предложение, но не делает его
принятым решением.

## Completeness

Перед переводом ADR в `active` убедись, что:

- указаны decision makers и реальные semantic upstream;
- рассмотрены минимум два жизнеспособных варианта, включая status quo, если он допустим;
- решение связано с драйверами и имеет явные scope/non-scope;
- зафиксированы положительные, отрицательные и организационные последствия;
- определены Confirmation evidence и условия пересмотра;
- Follow-up называет downstream canonical owners, которым принадлежат living facts и operational rules.

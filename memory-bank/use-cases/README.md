---
title: Use Cases Index
doc_kind: use_case
doc_function: index
purpose: Навигация по instantiated use cases проекта. Читать, чтобы найти канонический сценарий продукта или зарегистрировать новый.
derived_from:
  - ../dna/governance.md
  - ../flows/templates/use-case/UC-XXX.md
status: active
audience: humans_and_agents
---

# Use Cases Index

Каталог `memory-bank/use-cases/` хранит канонические пользовательские и операционные сценарии проекта.

Use case нужен для сценария, который живет на уровне продукта, повторяется во времени и может быть upstream для нескольких feature packages. Это не замена `SC-*` внутри `feature.md`: `SC-*` описывают acceptance сценарии delivery-единицы, а `UC-*` описывают устойчивое поведение системы на уровне проекта.

## Когда Заводить Use Case

- появляется новый стабильный пользовательский или операционный сценарий;
- несколько features реализуют или меняют один и тот же flow;
- нужен канонический owner для trigger, preconditions, main flow и postconditions.

## Когда Use Case Не Нужен

- сценарий одноразовый и живет только внутри одной feature;
- это implementation detail, а не продуктовый или операционный flow;
- его достаточно описать через `SC-*` в `feature.md`.

## Реестр

| UC ID | Title | Status | Primary actor | Upstream PRD | Implemented by | Last updated |
| --- | --- | --- | --- | --- | --- | --- |
| `UC-XXX` | Название сценария | `draft` / `active` / `archived` | Кто запускает flow | `PRD-XXX` / `none` | `FT-XXX` | YYYY-MM-DD |

## Naming

- Формат файла: `UC-XXX-short-name.md`
- Вместо `XXX` используй стабильный проектный идентификатор
- Один use case может быть upstream для нескольких feature packages

## Template

- Используй шаблон [`../flows/templates/use-case/UC-XXX.md`](../flows/templates/use-case/UC-XXX.md)

---
title: Use Cases Index
doc_kind: use_case
doc_function: index
purpose: Навигация по instantiated use cases проекта. Читать, чтобы найти канонический сценарий продукта или зарегистрировать новый.
derived_from:
  - ../dna/governance.md
  - ../flows/use-case.md
  - ../flows/templates/use-case/UC-XXX.md
status: active
audience: humans_and_agents
---

# Use Cases Index

Каталог `memory-bank/use-cases/` хранит канонические пользовательские и операционные сценарии проекта.

Use case нужен для сценария, который живет на уровне продукта, повторяется во времени и может быть upstream для нескольких feature packages. Это не замена `SC-*` внутри `brief.md`: `SC-*` описывают acceptance сценарии delivery-единицы, а `UC-*` описывают устойчивое поведение системы на уровне проекта.

Один `UC-*` может иметь много downstream BDD examples. `BR-*`, `ALT-*` и
`EX-*` дают точки traceability к feature `SC-*` / `NEG-*`, но example bodies,
`CHK-*` и test implementation не копируются в project-level use case. Правила
Discovery, Formulation и Automation определяет
[`Behavior Specification Practice`](../flows/behavior-specification.md).

Обычно use case наследует общий product context из [`../product/context.md`](../product/context.md). Если сценарий зависит от предметных правил, states или events, он также должен ссылаться на соответствующие документы из [`../domain/README.md`](../domain/README.md).

## Когда Заводить Use Case

- появляется новый стабильный пользовательский или операционный сценарий;
- несколько features реализуют или меняют один и тот же flow;
- нужен канонический owner для trigger, preconditions, main flow и postconditions.

## Когда Use Case Не Нужен

- сценарий одноразовый и живет только внутри одной feature;
- это implementation detail, а не продуктовый или операционный flow;
- его достаточно описать через `SC-*` в `brief.md`.

Подробные критерии, lifecycle создания и правила для operational / agentic
сценариев определяет [`Use Case Flow`](../flows/use-case.md).

## Реестр

Реестр является аннотированным списком instantiated use cases. Для каждой строки
сделай title относительной ссылкой на `UC-*` и кратко опиши наблюдаемый результат
сценария, а не только повтори название.

| UC ID | Title | Annotation | Status | Primary actor | Upstream PRD | Implemented by | Last updated |
| --- | --- | --- | --- | --- | --- | --- | --- |
| `UC-XXX` | Название сценария | Какой устойчивый результат получает actor | `draft` / `active` / `archived` | Кто запускает flow | `PRD-XXX` / `none` | `FT-XXX` | YYYY-MM-DD |

## Naming

- Формат файла: `UC-XXX-short-name.md`
- Вместо `XXX` используй стабильный проектный идентификатор
- Один use case может быть upstream для нескольких feature packages

## Template

- Используй шаблон [`../flows/templates/use-case/UC-XXX.md`](../flows/templates/use-case/UC-XXX.md)
- Создавай и обновляй документ по [`Use Case Flow`](../flows/use-case.md)

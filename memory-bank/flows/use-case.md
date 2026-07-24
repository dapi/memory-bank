---
title: Use Case Flow
doc_kind: governance
doc_function: canonical
purpose: Lifecycle создания, активации и обновления канонических project-level use cases.
derived_from:
  - ../dna/governance.md
  - feature.md
canonical_for:
  - use_case_selection
  - use_case_creation_flow
  - use_case_lifecycle
  - operational_agentic_use_case_rules
  - use_case_registry_contract
status: active
audience: humans_and_agents
---

# Use Case Flow

Этот flow управляет project-level `UC-*`: от решения завести канонический
сценарий до его регистрации, активации, обновления и архивации. Он не является
отдельным route для delivery-задачи и не заменяет [`Task Routing`](routing.md).

## Что Такое Use Case

Use case описывает устойчивое наблюдаемое поведение системы с точки зрения
actor-а: trigger, preconditions, основной flow, ожидаемые альтернативы,
исключения и postconditions. Actor-ом может быть пользователь, оператор,
команда, автоматизированный агент или внешний сервис.

`UC-*` является canonical owner сценария уровня проекта. Feature-level `SC-*`
остается owner-ом acceptance конкретной delivery-единицы, а feature-local
`FUC-*` — derived представлением сценариев для review.

## Selection Gate

Создавай или поднимай сценарий в `UC-*`, если выполняются все базовые условия:

- сценарий повторяется во времени и не принадлежит только одной delivery-единице;
- у него есть стабильные trigger, preconditions, flow и postconditions;
- его поведение важно независимо от конкретной реализации feature;
- нужен canonical owner вне одной feature, на который могут ссылаться features
  и другие project-level документы.

Сценарий особенно полезно оформить как `UC-*`, если он используется несколькими
features, runbooks, prompts или ops docs. Не создавай `UC-*` для одноразового
acceptance case, локального edge case или implementation detail: оставь это в
`SC-*`, `NEG-*` или optional feature-local `FUC-*`.

## Operational / Agentic Use Cases

Operational и agentic use cases описывают повторяемую работу, где человек,
агент, сервис или команда передает контекст, координирует работу, проверяет
готовность среды, публикует статус или восстанавливается после сбоя.

Machine-readable status, structured diagnostics, handoff payload и recovery
outcome принадлежат `UC-*`, только если являются наблюдаемой частью стабильного
сценария и одинаково интерпретируются несколькими участниками. Use case
фиксирует требуемое поведение и postconditions, но не implementation sequence,
архитектуру, внутренний protocol design или feature-level test matrix.

## Creation Flow

```text
candidate scenario → selection gate → stable UC ID → draft from template
                   → scenario review → registry annotation → active
                   → feature-driven updates → archived when no longer valid
```

1. Проверь Selection Gate и границу между `UC-*`, `SC-*` и `FUC-*`.
2. Выбери следующий стабильный `UC-*` ID по
   [`use-cases/README.md`](../use-cases/README.md).
3. Создай файл по [`UC-XXX` template](templates/use-case/UC-XXX.md).
4. Заполни общий scenario contract: goal, actor, trigger, preconditions, main
   flow, alternatives/exceptions, postconditions и business rules.
5. Для operational / agentic сценария добавь только применимые optional
   contracts: observable status, handoff, diagnostics и recovery.
6. Добавь upstream/downstream traceability без копирования требований или
   implementation details из owner-документов.
7. Зарегистрируй use case в аннотированном реестре и переведи его в `active`
   после прохождения Activation Gate.

## Lifecycle Gates

### Draft

- [ ] Selection Gate пройден
- [ ] стабильный `UC-*` ID зарезервирован в реестре
- [ ] документ создан по canonical template со `status: draft`
- [ ] upstream product/domain/PRD refs определены или явно указано `none`

### Draft → Active

- [ ] goal и primary actor однозначны
- [ ] trigger и preconditions проверяемы
- [ ] main flow описывает observable behavior, а не implementation sequence
- [ ] ожидаемые alternatives/exceptions и успешные/неуспешные postconditions
      зафиксированы
- [ ] operational contracts добавлены, только если они являются частью
      наблюдаемого project-level behavior
- [ ] traceability содержит актуальные upstream и downstream refs
- [ ] аннотированная строка в `use-cases/README.md` описывает результат сценария
- [ ] `status: active`

### Update Active Use Case

Если feature добавляет новый stable flow или materially меняет существующий,
сначала обнови canonical `UC-*`, затем feature-specific acceptance и derived
представления. Сохраняй прежний contract только через явно описанную
compatibility boundary; не оставляй одновременно два противоречащих active
описания одного сценария.

### Active → Archived

- [ ] сценарий больше не является поддерживаемым project-level behavior
- [ ] downstream refs обновлены или удалены
- [ ] замена или причина прекращения указана в документе
- [ ] реестр отражает `status: archived`

## Ownership Boundaries

- `use-cases/README.md` владеет навигацией, ID и короткими аннотациями.
- `UC-*` владеет project-level scenario contract.
- `flows/templates/use-case/UC-XXX.md` владеет структурой нового документа.
- `brief.md` владеет feature scope, acceptance и evidence contract.
- `design.md`, ADR и delegated contracts владеют solution и architecture facts.
- runbooks владеют executable operational procedures и конкретными recovery
  commands; `UC-*` фиксирует ожидаемое поведение и outcome.

## Outcome / Exit Contract

Use case имеет `status: active`, зарегистрирован с содержательной аннотацией,
описывает устойчивый observable scenario и связан с актуальными upstream и
downstream owners. Реализация и проверка остаются в соответствующих delivery и
operations artifacts.

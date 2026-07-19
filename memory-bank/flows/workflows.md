---
title: Task Workflows
doc_kind: governance
doc_function: canonical
purpose: Маршрутизация задач по типам и базовый цикл разработки. Читать при получении новой задачи для выбора подхода.
derived_from:
  - ../dna/governance.md
  - feature-flow.md
canonical_for:
  - task_routing_rules
  - base_development_cycle
  - workflow_type_selection
  - autonomy_gradient
status: active
audience: humans_and_agents
---

# Task Workflows

## Базовый цикл

Любой workflow — цепочка повторений одного цикла:

```text
Артефакт → Ревью → Полировка
                  → Декомпозиция
                  → Принят
```

Артефакт — то, что создаётся на каждом этапе: brief, design-док, план, код, PR, runbook.

## Градиент участия человека

Чем ближе к бизнес-требованиям, тем больше участия человека. Чем ближе к коду и локальному verify, тем больше агент работает автономно.

```text
Бизнес-требования  ← человек  |  агент →  Код
  PRD, Use Cases      Brief, Design, Plan   PR, Тесты
```

## Типы Workflow

### 1. Локальное изменение (Small Change)

`Small Change` — задача, для которой issue/task уже содержит достаточно контекста, чтобы безопасно перейти к реализации без отдельных requirement-, design- и planning-документов. Определяющий признак — не размер diff или длительность работы, а отсутствие необходимости принимать и фиксировать design decisions.

Все условия должны выполняться одновременно:

- issue/task полностью задаёт intent, scope и acceptance;
- изменение не создаёт и materially не меняет пользовательское поведение;
- решение следует существующему паттерну и не требует выбора подхода;
- не меняются API, event, schema, file format, CLI, env/config или integration contracts;
- не затрагиваются security boundary, data migration, rollout или обязательные approvals;
- change surface локален, а отдельная декомпозиция и checkpoints не нужны.

Для `Small Change` не создаются feature package, `brief.md`, `design.md`, `implementation-plan.md` или ADR. Issue/task остаётся owner-ом intent, scope и acceptance.

Flow:

`issue/task -> routing gate -> implementation -> automated checks -> simplify review -> PR -> review + CI -> merge`

Типичные примеры: документационная правка, добавление недостающего теста, локальная чистка кода или небольшое изменение внутренних инструментов без изменения их контрактов.

Если обнаружена необходимость воспроизвести дефект и защититься от регрессии, используй Bug Fix Flow. Если появляется или materially меняется пользовательское поведение, требуется design decision или срабатывает любой contract/risk trigger выше, останови `Small Change` и маршрутизируй работу в Feature Flow. Компактная feature остаётся vertical slice и использует минимальный feature package.

### 2. Feature (vertical slice)

Когда:

- появляется или materially меняется пользовательское поведение;
- задача представляет отдельную единицу пользовательской ценности;
- acceptance должен покрывать результат end-to-end через все затронутые слои.

Размер feature определяет глубину содержимого package, а не выбор workflow. Компактная feature использует минимальный `brief.md` и может зафиксировать `Design required: no`; более сложная feature добавляет `design.md`, richer verification context, checkpoints и подробный execution plan по правилам Feature Flow.

Flow:

`issue/task -> feature package -> brief -> optional design -> implementation plan -> execution -> review -> handoff`

### 3. Баг-фикс

Источники могут быть любыми: error tracker, support, QA, прямой report от пользователя, инцидентный анализ.

Flow:

`report -> reproduction -> analysis -> fix -> regression coverage -> review`

### 4. Рефакторинг

Разделяй минимум на три класса:

- по ходу delivery-задачи;
- исследовательский;
- системный, с большим change surface.

Исследовательский и системный refactoring обычно требуют явного плана и checkpoints.

### 5. Инцидент / PIR

Flow:

`incident -> timeline -> root cause analysis -> fixes -> prevention work`

Здесь человек обычно подтверждает RCA и приоритеты follow-up задач.

## Routing Rules

Используй минимальный workflow, который не теряет контроль над риском.

- Используй `Small Change` только когда все его routing predicates истинны; размер diff сам по себе не является критерием.
- Если задача меняет контракт, rollout или требует approvals, поднимай её до feature flow.
- Если в ходе `Small Change` потребовался отдельный design или execution plan, останови прямую реализацию и выполни повторный routing.
- Если замечания не уменьшаются от итерации к итерации, проблема может быть upstream, а не в коде.

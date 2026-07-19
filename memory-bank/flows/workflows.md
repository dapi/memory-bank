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

## Routing Order

Проверяй маршруты именно в этом порядке. `Small Change` — fast path перед Feature Flow, а не отдельный semantic type задачи.

```text
Issue / Task
     |
     +-- Incident / PIR? ----------------> Incident Flow
     |
     +-- Bug? ----------------------------> Bug Fix Flow
     |
     +-- Issue достаточен,
     |   design и plan не нужны? --------> Small Change Flow
     |
     +-- Новое или изменённое
     |   пользовательское поведение? ----> Feature Flow
     |
     +-- Refactoring? --------------------> Refactoring Flow
     |
     +-- Неясно / высокий риск ----------> Human Routing
```

## Типы Workflow

### 1. Инцидент / PIR

Flow:

`incident -> timeline -> root cause analysis -> fixes -> prevention work`

Человек обычно подтверждает RCA и приоритеты follow-up задач.

### 2. Bug Fix

Источники могут быть любыми: error tracker, support, QA, прямой report от пользователя или инцидентный анализ.

Flow:

`report -> reproduction -> analysis -> fix -> regression coverage -> review`

### 3. Локальное изменение (Small Change)

`Small Change` — задача, для которой issue/task уже содержит достаточно контекста, чтобы безопасно перейти к реализации без отдельных requirement-, design- и planning-документов. Определяющий признак — не размер diff и не semantic type изменения, а отсутствие необходимости принимать и фиксировать design decisions.

Все условия должны выполняться одновременно:

- issue/task полностью задаёт intent, scope и acceptance;
- решение следует существующему паттерну и не требует выбора подхода;
- не меняются API, event, schema, file format, CLI, env/config или integration contracts;
- не затрагиваются security boundary, data migration, rollout или обязательные approvals;
- change surface локален, а отдельная декомпозиция и checkpoints не нужны.

Для `Small Change` не создаются feature package, `brief.md`, `design.md`, `implementation-plan.md` или ADR. Issue/task остаётся owner-ом intent, scope и acceptance.

Flow:

`issue/task -> routing gate -> implementation -> automated checks -> simplify review -> PR -> review + CI -> merge`

Типичные примеры: документационная правка, добавление недостающего теста, локальная чистка кода, небольшое изменение внутренних инструментов или точечное улучшение существующего пользовательского поведения в рамках действующих контрактов.

Если обнаружена необходимость воспроизвести дефект и защититься от регрессии, используй Bug Fix Flow. Если потребовался отдельный design, execution plan или сработал любой contract/risk trigger выше, останови `Small Change` и выполни повторный routing.

### 4. Feature (vertical slice)

Используй Feature Flow, если задача не прошла `Small Change` gate и при этом:

- создаёт или materially меняет пользовательское поведение;
- представляет отдельную единицу пользовательской ценности;
- требует end-to-end acceptance через все затронутые слои.

Feature package является vertical slice. Его глубина зависит от сложности: `brief.md` использует минимальный или расширенный problem-space набор, `design.md` появляется только при `Design required: yes`, а `implementation-plan.md` создаётся после готовности upstream owners.

Flow:

`issue/task -> feature package -> brief -> optional design -> implementation plan -> execution -> review -> handoff`

### 5. Refactoring

Локальный refactoring может пройти через `Small Change`. Исследовательский и системный refactoring с большим change surface требуют явного плана и checkpoints.

## Routing Rules

- Сначала отделяй Incident и Bug Fix, затем проверяй `Small Change` до Feature и Refactoring.
- Используй `Small Change` только когда все его routing predicates истинны; размер diff сам по себе не является критерием.
- Если задача меняет контракт, rollout или требует approvals, она не может оставаться `Small Change`.
- Если в ходе `Small Change` потребовался отдельный design или execution plan, останови прямую реализацию и выполни повторный routing.
- Если маршрут остаётся неоднозначным или риск нельзя контролировать выбранным flow, запрашивай решение человека.
- Если замечания не уменьшаются от итерации к итерации, проблема может быть upstream, а не в коде.

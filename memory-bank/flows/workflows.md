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

Артефакт — то, что создаётся на каждом этапе: спецификация, дизайн-док, план, код, PR, runbook.

## Градиент участия человека

Чем ближе к бизнес-требованиям, тем больше участия человека. Чем ближе к коду и локальному verify, тем больше агент работает автономно.

```text
Бизнес-требования  ← человек  |  агент →  Код
  PRD, Use Cases      Спека, План           PR, Тесты
```

## Типы Workflow

### 1. Малая фича

Когда:

- задача понятна;
- scope локален;
- решение помещается в одну сессию или один компактный change set.

Flow:

`issue/task -> routing -> implementation -> review -> merge`

### 2. Средняя или большая фича

Когда:

- затрагивает несколько слоёв;
- требует design choices;
- нужны checkpoints и явный execution plan.

Flow:

`issue/task -> spec -> feature package -> implementation plan -> execution -> review -> handoff`

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

- Если задача маленькая и понятная, не раздувай её до большого feature package.
- Если задача меняет контракт, rollout или требует approvals, поднимай её до feature flow.
- Если замечания не уменьшаются от итерации к итерации, проблема может быть upstream, а не в коде.

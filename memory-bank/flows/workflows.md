---
title: Task Workflows
doc_kind: governance
doc_function: canonical
purpose: Маршрутизация задач по типам и базовый цикл разработки. Читать при получении новой задачи для выбора подхода.
derived_from:
  - ../dna/governance.md
  - task-flow.md
  - bugfix-flow.md
  - refactor-flow.md
  - feature-flow.md
canonical_for:
  - task_routing_rules
  - task_signature_routing_rules
  - base_development_cycle
  - workflow_type_selection
  - workflow_document_profile_selection
  - workflow_promotion_rules
  - autonomy_gradient
status: active
audience: humans_and_agents
---

# Task Workflows

Этот документ является верхним selector-слоем перед `task-flow`, `bugfix-flow`, `refactor-flow`, `feature-flow` и `epic-flow`. Агент сначала классифицирует задачу и выбирает минимальный workflow/document profile, а уже затем применяет конкретный flow.

## Базовый цикл

Любой workflow — цепочка повторений одного цикла:

```text
Артефакт → Ревью → Полировка
                  → Декомпозиция
                  → Принят
```

Артефакт — то, что создается на каждом этапе: task note, brief, design-док, план, код, PR, runbook.

## Градиент участия человека

Чем ближе к бизнес-требованиям, тем больше участия человека. Чем ближе к коду и локальному verify, тем больше агент работает автономно.

```text
Бизнес-требования  ← человек  |  агент →  Код
  PRD, Use Cases      Brief, Design, Plan   PR, Тесты
```

## Routing Signature

Перед выбором workflow зафиксируй минимальную подпись задачи. Для маленьких задач подпись может жить в issue/PR; для managed task package - в `memory-bank/tasks/TASK-XXX/`; для feature/epic - в соответствующем package.

| Поле | Значения | Зачем нужно |
| --- | --- | --- |
| `kind` | `feature`, `bugfix`, `chore`, `refactor`, `incident`, `epic`, `unknown` | Выбрать семью workflow до создания документов |
| `size` | `tiny`, `small`, `medium`, `large`, `unknown` | Отделить одно-сессионные изменения от managed delivery |
| `risk` | `low`, `normal`, `high`, `critical`, `unknown` | Определить, нужен ли durable plan/design/evidence |
| `change_surface` | `single_component`, `multi_component`, `cross_boundary`, `external`, `unknown` | Понять blast radius |
| `contract_change` | `none`, `api`, `schema`, `event`, `security`, `financial`, `integration`, `rollout`, `unknown` | Включить promotion triggers |
| `design_need` | `none`, `local_reasoning`, `required`, `unknown` | Решить, нужен ли `design.md` / ADR / managed package |
| `evidence_need` | `focused_test`, `regression_test`, `manual_evidence`, `ci_evidence`, `external_e2e`, `unknown` | Выбрать verify carrier |
| `owner_doc_need` | `none`, `task_note`, `managed_task_package`, `feature_package`, `epic_package`, `unknown` | Выбрать комплект документов |
| `workflow_profile` | `tracker-only`, `bugfix-compact`, `refactor-small`, `managed-task`, `feature-package`, `epic-package`, `incident-pir`, `unknown` | Зафиксировать итоговый profile selector-а |

Если `kind`, `risk`, `change_surface` или `contract_change` остаются `unknown`, агент не должен молча выбирать самый короткий путь. Нужно либо закрыть неизвестность через discovery, либо поднять задачу до более строгого profile.

## Типы Workflow

### 1. Малая фича

Когда:

- задача понятна;
- scope локален;
- решение помещается в одну сессию или один компактный change set.

Flow:

`issue/task -> routing -> implementation -> review -> merge`

Для маленькой non-feature задачи используй document profiles ниже, а не full `feature-flow` по умолчанию.

### 2. Средняя или большая фича

Когда:

- затрагивает несколько слоёв;
- требует design choices;
- нужны checkpoints и явный execution plan.

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

## Document Profiles

Используй минимальный комплект документов, который сохраняет контроль над риском и evidence. `feature-flow` больше не является default для всех задач.

| Profile | Когда применять | Canonical carriers | Не применять, если |
| --- | --- | --- | --- |
| `tracker-only` | `tiny/small`, `low` risk, локальная правка без contract change | GitHub issue/PR: routing signature, short rationale, checks/evidence | Нужен durable owner doc, есть design decision или high-risk surface |
| `bugfix-compact` | Дефект с понятным symptom/repro/root cause и локальным fix | [`bugfix-flow.md`](bugfix-flow.md); GitHub issue/PR по шаблону [`templates/task/bugfix.md`](templates/task/bugfix.md); optional `memory-bank/tasks/TASK-XXX/README.md` + `bugfix.md`, если нужен durable context | Меняется API/schema/security/financial/integration/rollout или root cause требует design reasoning |
| `refactor-small` | Локальный рефакторинг без изменения поведения | [`refactor-flow.md`](refactor-flow.md); GitHub issue/PR по шаблону [`templates/task/refactor.md`](templates/task/refactor.md); optional `memory-bank/tasks/TASK-XXX/README.md` + `refactor.md`, если нужен durable context | Большой change surface, migration/checkpoints или risk of behavior drift |
| `managed-task` | Bugfix/chore/refactor больше issue/PR, но еще не feature: нужен scope, invariants, checkpoints, evidence | `memory-bank/tasks/TASK-XXX/` с `bugfix.md` или `refactor.md`; optional linked ADR/runbook/incident docs | Появляется новая user/system capability или устойчивый scenario |
| `feature-package` | Новая или materially changed capability / delivery slice | `memory-bank/features/FT-XXX/` по [`feature-flow.md`](feature-flow.md) | Требуется roadmap/risk register/subissues для нескольких delivery units |
| `epic-package` | Инициатива крупнее одной delivery unit | `memory-bank/epics/EP-XXX/` по [`epic-flow.md`](epic-flow.md) | Есть только одна локальная delivery task |
| `incident-pir` | Инцидент, RCA, prevention work | Incident/PIR docs по project owner-докам | Это обычный дефект без incident/RCA context |

`tracker-only`, `bugfix-compact`, `refactor-small` и `managed-task` используют общий [`task-flow.md`](task-flow.md). Bugfix-процесс описан в [`bugfix-flow.md`](bugfix-flow.md), refactor/chore-процесс - в [`refactor-flow.md`](refactor-flow.md). `memory-bank/tasks/` не заменяет `features/`: он нужен для управляемых non-feature задач, где full feature package создает лишнюю ceremony, но issue/PR уже недостаточны для traceability.

## Promotion Rules

Начинай с минимального profile и повышай его при появлении trigger-а.

| Trigger | Куда повышать | Причина |
| --- | --- | --- |
| Меняется API, event, schema, file format, CLI, security boundary, financial calculation, integration contract или operational rollout | `feature-package` или ADR + `feature-package` | Нужен solution-space owner и acceptance traceability |
| Bugfix требует alternatives/trade-off reasoning, migration strategy, rollout/backout или explicit failure-mode design | `managed-task` или `feature-package` | Issue/PR note перестает быть достаточным owner-ом решения |
| Refactor затрагивает несколько bounded contexts, shared abstractions или runtime topology | `managed-task` / `feature-package`; для серии задач - `epic-package` | Нужны checkpoints, invariants и stop conditions |
| Исправление создает новый устойчивый user/operator/system scenario | `feature-package` + обновление `UC-*` при необходимости | Это уже capability, а не только fix/chore |
| В процессе review появляются повторные замечания о scope/design/evidence | Повысить profile и обновить owner docs до новых правок | Проблема находится upstream, а не в локальном diff |
| `risk=high/critical` или есть irreversible/external side effects | Минимум `managed-task`, часто `feature-package` | Нужны явные approvals, rollback/backout и durable evidence |

## Routing Rules

Используй минимальный workflow, который не теряет контроль над риском.

- Если задача маленькая и понятная, не раздувай её до большого feature package.
- Если задача меняет контракт, rollout или требует approvals, поднимай её до feature flow.
- Если замечания не уменьшаются от итерации к итерации, проблема может быть upstream, а не в коде.
- Если compact profile выбран для bugfix/refactor, PR обязан содержать routing signature и evidence по соответствующему template.

## Optional Routing Support

Для изменений самого selector-а используй [Workflow Decision Log](workflow-decision-log.md): он фиксирует причину, evidence, решение, expected consequences и review rule.

Для проверки, что compact или managed profiles уменьшили лишнюю ceremony без потери safety/evidence, используй [Workflow Metrics](workflow-metrics.md). Safety misroutes, missing evidence и rework проверяются раньше lead time.

Для короткого объяснения выбора workflow profile в issue/PR используй [Workflow Routing Developer Brief](workflow-routing-developer-brief.md). Он не заменяет этот canonical routing document и не делает optional profiles обязательными для каждого проекта.

# Оркестрация цифровых ролей и бесконечная сессия поверх Memory Bank

Источник этой статьи — созвон Михаила С. «Управление сложностью кода через Code Property Graph и цифровые роли»:

<https://app.mymeet.ai/ru/public-meetings/ab134fa3-e164-4e5a-89f0-c8530cfde37d>

## Короткий вывод

В обсуждаемой системе нет «толпы автономных агентов», которые свободно придумывают следующий шаг. Есть:

1. цифровые роли с ограниченными инструкциями и инструментами;
2. декларативный workflow, описывающий допустимые переходы;
3. долговечная Task Capsule с состоянием конкретной задачи;
4. общая база артефактов и handoff-ов;
5. валидатор, который разрешает или отклоняет передачу задачи следующей роли.

«Бесконечная сессия» — это не бесконечное окно контекста модели. Это последовательность коротких запусков, между которыми сохраняется состояние задачи:

```text
прочитать состояние
  → выполнить один шаг
  → записать артефакты и evidence
  → обновить состояние
  → запустить следующую роль или продолжение
```

Главное, что сохраняется, — не весь текст чата, а контекст задачи.

## Как устроена оркестрация

```text
Issue / Task
    ↓
Task Capsule — состояние и идентичность задачи
    ↓
Workflow — допустимые стадии и переходы
    ↓
Digital Role — инструкция, skills, tools и output contract
    ↓
Artifacts + Evidence + Handoff
    ↓
Validator → следующая роль или понятная ошибка
```

### Цифровая роль

Цифровая роль — это управляемый пакет, а не обязательно отдельный постоянно живущий агент. В него входят:

- инструкция роли;
- разрешённые skills и tools;
- входной контракт;
- ожидаемые выходные артефакты;
- ограничения на изменения;
- допустимые следующие роли;
- критерии готовности к handoff.

Роль может использовать LLM, но сама модель не должна владеть процессом. Процесс задаётся контрактами и workflow.

### Workflow

Workflow — это state machine. В созвоне он реализован вокруг Temporal и описывается декларативно, в YAML. Это позволяет менять порядок и состав ролей без переписывания всей бизнес-логики оркестратора.

Переход должен проверять не только текстовый ответ модели, но и состояние общего хранилища:

- создан ли обязательный документ;
- имеет ли он правильный статус;
- заполнены ли обязательные поля;
- есть ли связь с upstream owner;
- приложено ли evidence;
- не нарушен ли выбранный flow.

Если контракт не выполнен, handoff отклоняется с диагностикой, которую модель может понять и исправить.

### Task Capsule

Task Capsule следует за задачей до завершения выбранного lifecycle. Минимальное состояние капсулы:

```yaml
task_id: FT-042
route: feature
stage: implementation
current_role: implementer
next_role: verifier
status: active
artifacts:
  brief: features/FT-042/brief.md
  design: features/FT-042/design.md
  plan: features/FT-042/implementation-plan.md
open_risks:
  - missing_negative_scenario
next_action: Implement STEP-03
stop_conditions:
  - plan and code contract diverge
  - required test evidence missing
```

Для возобновления не требуется восстанавливать весь предыдущий разговор. Достаточно прочитать актуальную капсулу, canonical owner-документы и ближайшие проверки.

## Что означает «бесконечная сессия»

Это durable workflow поверх коротких model sessions:

```text
Task Capsule
    ↓
короткий вызов роли
    ↓
коммит артефактов и evidence
    ↓
новый вызов с тем же task_id
```

Такой подход переживает:

- compact или исчерпание context window;
- падение модели или инструмента;
- смену исполнителя;
- паузу между рабочими сессиями;
- повторный запуск после исправления ошибки handoff.

Важный принцип: нужно хранить контекст задачи, а не пытаться сохранять «вечную память» диалога.

## Как это ложится на Memory Bank

У Memory Bank уже есть основные части этой модели:

- Task Routing выбирает flow: [`template/memory-bank/flows/routing.md`](../template/memory-bank/flows/routing.md);
- Feature Flow разделяет problem, solution и execution space: [`template/memory-bank/flows/feature.md`](../template/memory-bank/flows/feature.md);
- Epic Flow управляет roadmap, рисками и delivery slices: [`template/memory-bank/flows/epic.md`](../template/memory-bank/flows/epic.md);
- session handoff хранит состояние прерванного процесса: [`template/memory-bank/flows/templates/process/session-handoff.md`](../template/memory-bank/flows/templates/process/session-handoff.md);
- lifecycle protocol задаёт фазы, gates, verification, rollback и stop conditions: [`template/memory-bank/flows/templates/process/lifecycle-protocol.md`](../template/memory-bank/flows/templates/process/lifecycle-protocol.md).

### Feature как Task Capsule

Для delivery-задачи капсулой может быть feature package:

```text
memory-bank/features/FT-XXX/
├── README.md
├── brief.md
├── design.md
├── implementation-plan.md
└── session-handoff.md
```

`README.md` — routing и текущая стадия.
`brief.md` — проблема, scope, acceptance и validation profile.
`design.md` — решение и solution facts, если design required.
`implementation-plan.md` — исполнимый план и проверки.
`session-handoff.md` — продолжение после паузы.

### Epic как верхнеуровневая капсула

Большую инициативу следует хранить как epic:

```text
memory-bank/epics/EP-XXX/
├── README.md
├── charter.md
├── roadmap.md
├── subissues.md
├── risks.md
└── decision-log.md
```

Epic не должен превращаться в общий execution plan. Каждая принятая delivery-unit передаётся в отдельный `FT-*` package.

### Роли и handoff

Роли можно описывать отдельными process-документами только там, где это действительно reusable policy:

```text
memory-bank/processes/
├── feature-lifecycle.md
├── epic-lifecycle.md
├── session-handoff.md
└── role-contracts/
    ├── analyst.md
    ├── designer.md
    ├── implementer.md
    └── verifier.md
```

Пример контракта роли:

```yaml
role: implementer
accepts:
  - brief.md
  - design.md
  - implementation-plan.md
produces:
  - code_changes
  - test_results
  - evidence_log
handoff_to:
  - verifier
must_not:
  - change_requirements_without_updating_brief
```

Это не означает, что все роли нужно сразу превращать в runtime-агентов. Сначала достаточно зафиксировать входы, выходы и границы автономии.

## Что автоматизировать

### Первый этап: файловая оркестрация

Без Temporal и отдельной базы:

1. создавать Task Capsule для каждой активной задачи;
2. хранить состояние в frontmatter и компактном handoff-документе;
3. запускать нужную роль по текущей стадии;
4. проверять handoff через CLI/validator;
5. записывать evidence в Git вместе с документами и кодом.

На этом этапе Git остаётся durable storage, а Memory Bank — источником истины.

### Второй этап: машинная проверка переходов

Нужен validator, который проверяет:

- наличие обязательных артефактов;
- разрешённость перехода для текущего flow;
- frontmatter и статусы;
- traceability между `UC-*`, `REQ-*`, `SC-*`, tests и кодом;
- обязательное evidence;
- stop conditions.

Важно: registry активных задач может быть только индексом или projection. Он не должен становиться второй копией требований и статусов.

### Третий этап: runtime-оркестратор

Temporal или аналогичный runtime оправдан, когда появляются:

- долгие задачи с ожиданием внешних событий;
- retries и timers;
- параллельные роли;
- очереди задач;
- автоматическое возобновление после падения;
- несколько workflow для разных типов задач.

До этого достаточно Git, frontmatter, process contracts и validator-а.

## Границы и принципы

1. **Workflow владеет переходами, роль — выполнением шага.**
2. **Canonical documents владеют фактами, orchestration database — только состоянием исполнения.**
3. **Короткие сессии безопаснее длинного чата, если каждый шаг оставляет evidence.**
4. **Детерминировать нужно routing, handoff, проверки и структурные аудиты.**
5. **LLM оставлять там, где нужна интерпретация, анализ и формулировка вариантов.**
6. **Если работа меняет scope или выбранный подход, требуется rerouting, а не молчаливое продолжение.**

## Связь с трассируемостью

Code Property Graph или статический анализ контролируют структуру и качество кода, но не доказывают соответствие бизнес-намерению.

Для этого Memory Bank нужен отдельный evidence graph:

```text
User Story / Use Case
    → PRD / Requirements
    → Acceptance Scenarios
    → Design / ADR
    → Implementation Plan
    → Code
    → BDD / unit / end-to-end tests
    → Evidence
```

Это две связанные, но разные проверки:

- code graph отвечает на вопрос «как устроен код и не нарушает ли он технические правила»;
- requirements traceability отвечает на вопрос «реализовано ли то, что требовалось».

## Предлагаемый outcome для Memory Bank

Внедрение стоит считать завершённым, когда:

- у каждой активной delivery-задачи есть Task Capsule;
- текущая роль и следующий шаг однозначно определяются из persisted state;
- handoff проверяется автоматически или воспроизводимой командой;
- после остановки новую сессию можно запустить только по capsule и owner-документам;
- failures содержат диагностируемую причину и next action;
- registry не дублирует canonical facts;
- Feature, Epic, Bug Fix, Research и Small Change используют разные минимальные workflow.

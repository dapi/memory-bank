# Использование Memory Bank

Этот документ описывает повседневную работу с уже внедрённым Memory Bank. Он остаётся во внешней документации шаблона и не копируется в downstream-проект. Канонические project-side правила находятся в [`memory-bank/README.md`](../memory-bank/README.md) и `memory-bank/flows/`; первичная установка и адаптация описаны в [`adoption.md`](adoption.md).

## Рабочая модель

Memory Bank не заменяет task tracker или инструмент запуска агента. Он хранит контекст и правила, issue задаёт конкретную работу, а agent runner создаёт рабочее окружение и запускает coding agent.

```text
Memory Bank
контекст, правила, требования и способы проверки
        ↓
Issue / Task
задача и ссылки на нужные документы
        ↓
agent runner
branch → worktree → agent session
        ↓
реализация → проверки → PR → evidence
        ↓
обновление Memory Bank при появлении новых знаний
```

Инструменты вроде [`start-issue`](https://github.com/dapi/start-issue) могут автоматизировать создание ветки и worktree и запуск выбранного агента, но не являются обязательной частью Memory Bank.

## Рабочий цикл

1. Подготовьте issue с ожидаемым результатом и ссылками на применимые PRD, epic, use case, feature package или ADR.
2. Выберите workflow по [`memory-bank/flows/routing.md`](../memory-bank/flows/routing.md).
3. Запустите агента в изолированной ветке или worktree.
4. Агент читает issue и связанные owner-документы, реализует изменение и выполняет предусмотренные проверки.
5. Завершите работу через PR и приложите требуемые evidence.
6. Если появились новые устойчивые правила, ограничения или решения, обновите их canonical owner в Memory Bank.

Если issue полностью задаёт intent, scope и acceptance, решение не требует design-документов и все routing predicates выполнены, задача может пройти как `Small Change` напрямую к реализации.

## Validation profiles

Delivery flow и validation profile выбираются последовательно и отвечают на разные вопросы:

- flow определяет lifecycle задачи, обязательные owner-документы и handoff;
- validation profile определяет минимальную глубину tests, CI gates, evidence, approvals и rollout/backout.

```text
Task Routing → delivery flow → validation profile → план проверок
             → реализация → evidence → review / merge
```

Profile не является отдельным flow и не меняет routing order. После выбора flow человек или агент проверяет risk triggers и фиксирует ровно один profile в canonical owner задачи. Сейчас это governance-механизм: `memory-bank lint` проверяет целостность документации, но не вычисляет profile автоматически и не запускает соответствующие test suites.

Canonical taxonomy и minimum contracts определены в [`memory-bank/engineering/validation-profiles.md`](../memory-bank/engineering/validation-profiles.md):

- `documentation` — только non-runtime documentation/artifact changes;
- `low-risk` — локальное executable change по известному паттерну без risk triggers;
- `standard` — default для обычного executable change;
- `high-risk` — security/trust, financial calculation, persistent data/migration, concurrency/idempotency или material cross-system integration;
- `release-deployment` — production config, build/release artifact, deployment или rollback path без отдельного high-risk trigger.

Если одновременно применимы `high-risk` и `release-deployment`, выбирается `high-risk` и дополнительно выполняются release/deployment obligations. Маленький diff или отсутствие готовой test environment не являются основанием снизить profile. Снижение после сработавшего high-risk или release trigger требует rationale и human approval reference.

### Где и когда фиксируется profile

| Flow | Момент выбора | Canonical owner |
| --- | --- | --- |
| Small Change | До реализации, вместе с routing record | Issue/task; draft PR только если tracker нельзя обновить |
| Feature | При подготовке `brief.md`, до `Problem Ready` | `memory-bank/features/FT-XXX/brief.md` |
| Bug Fix | На Entry Gate, до analysis и fix | Bug report или связанная delivery task |
| Refactoring | На Entry Gate, до characterization и execution plan | Исходная task |
| Incident / PIR | Для containment и PIR profile не выбирается | Отдельная remediation/prevention task после повторного Task Routing |
| Epic | Для epic целиком profile не выбирается | Отдельный owner каждой delivery feature/subissue |

Минимальная запись решения:

```text
Validation profile: documentation | low-risk | standard | high-risk | release-deployment
Triggers / rationale: <почему выбранный minimum достаточен>
Downgrade approval: <human approval ref или none>
```

После выбора profile конкретные проверки подключаются на уровне исполнения:

- в Small Change команды и ожидаемое evidence записываются в `Verify` routing record;
- в Feature решение остаётся в `brief.md`, а `implementation-plan.md` связывает его obligations с конкретными automated test surfaces, local suites, CI jobs, manual evidence, approval gates и rollout/backout checkpoints;
- в Bug Fix reproduction и regression coverage должны удовлетворять minimum contract выбранного profile;
- в Refactoring baseline, characterization coverage и checkpoint verification должны удовлетворять minimum contract выбранного profile.

Если во время работы обнаружен более сильный trigger, сначала обновите profile у canonical owner и только затем продолжайте реализацию. Если новый trigger также нарушает predicates текущего flow — например, в Small Change обнаружились migration или rollout requirements — остановите работу и повторите [Task Routing](../memory-bank/flows/routing.md).

### Примеры

| Задача | Flow | Profile | Практическое следствие |
| --- | --- | --- | --- |
| Исправить локальный UI label по существующему i18n pattern | Small Change | `low-risk` | Targeted UI/i18n check, required CI, semantic read-through и обычный review |
| Добавить локальное пользовательское поведение без contract, data или security changes | Feature | `standard` | Regression и acceptance coverage, affected local suites, полный required CI и independent review |
| Изменить payment calculation | Feature | `high-risk` | Boundary/failure coverage, explicit human approval, independent domain review, rollout signals и backout plan |

## Стартовые запросы

### Создать feature package

```text
Прочитай ./memory-bank/README.md, ./memory-bank/flows/routing.md
и ./memory-bank/flows/feature.md. Сначала определи route текущей задачи.
Если выбран не Feature Flow, остановись и сообщи подходящий route.
Если выбран Feature Flow, создай feature package, начиная с README.md и brief.md.
design.md создавай только по правилам Design Requirement Decision,
а implementation-plan.md — только после готовности upstream-документов.
```

### Проверить качество Memory Bank

```text
Проведи ревью ./memory-bank на SSoT, противоречия, broken links,
orphan-документы, недостающие README-индексы и неясные зависимости.
Предложи минимальные правки и запусти локальные проверки.
```

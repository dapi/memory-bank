# Memory Bank — governance-ядро AI Software Development OS

## О чём этот репозиторий

**Memory Bank — governance- и knowledge-ядро AI Software Development OS:** переносимый documentation-first шаблон для разработки ПО с AI-агентами.

Его копируют в другой проект и адаптируют под конкретный продукт. Шаблон помогает хранить проверяемый контекст, чтобы человек и агент одинаково понимали:

- какой продукт создаётся и для кого;
- как устроена предметная область;
- какие инженерные и операционные правила действуют;
- какие требования, сценарии и архитектурные решения приняты;
- как фича проходит путь от постановки проблемы до реализации и проверки.

Ключевая идея — документация как управляемая система знаний. `dna/` задаёт правила и единый источник истины; `product/`, `domain/`, `engineering/` и `ops/` описывают постоянный контекст проекта; PRD, epic, use case и ADR фиксируют инициативы, сценарии и решения; feature packages связывают требования, дизайн, план реализации и результаты проверок.

Здесь **OS** — метафора operating system для процесса разработки. Memory Bank работает как control plane: определяет контекст, правила, lifecycle и критерии готовности, но не исполняет задачи самостоятельно. GitHub Issues, `start-issue`, coding agents, Git и CI образуют execution layer вокруг этого ядра.

Для фич предусмотрен последовательный workflow:

```text
brief — что и зачем
  → design — какое решение выбрано, если дизайн необходим
    → implementation-plan — как выполнить и проверить
```

Репозиторий не содержит приложения или runtime-кода. Это generic-шаблон документационного контура, дополненный скриптом, который проверяет ссылки, индексацию и достижимость документов.

## Зачем это нужно

Memory Bank становится общей рабочей памятью команды и агентов. Важный контекст хранится не в голове разработчика и не в одноразовом чате, а в связанных, проверяемых и версионируемых документах рядом с кодом.

Шаблон даёт для этого:

- единый источник истины для каждого значимого факта;
- явные зависимости и приоритет между документами;
- навигацию от общего контекста к конкретной задаче;
- lifecycle и шаблоны для PRD, epic, use case, feature и ADR;
- проверяемую связь между требованиями, решением, планом реализации и результатами проверок.

## Как устроена работа

Сначала команда адаптирует постоянный контекст проекта: `product/`, `domain/`, `engineering/` и `ops/`. Затем конкретные инициативы и задачи получают подходящие документы — от небольшого task workflow до PRD, epic или feature package.

Для значимой фичи контекст созревает поэтапно:

```text
brief.md                 design.md                  implementation-plan.md
что и зачем       →      какое решение       →     как реализовать и проверить
problem space            solution space             execution space
                         (если требуется)
```

Документы не должны дублировать друг друга. Код владеет реализацией, а Memory Bank — намерением, требованиями, обоснованием решений и контрактами. После завершения работы в документах остаётся контекст, с которым другой человек или новая сессия агента могут продолжить работу.

Подробное введение для команды и рекомендации по внедрению находятся в [`INTRO.md`](INTRO.md).

## Что находится в шаблоне

| Каталог | Назначение |
| --- | --- |
| [`dna/`](memory-bank/dna/README.md) | Governance-ядро: SSoT, frontmatter, lifecycle и правила связей между документами |
| [`product/`](memory-bank/product/README.md) | Контекст продукта, vision, customers, metrics, marketing и roadmap |
| [`domain/`](memory-bank/domain/README.md) | Glossary, domain model, business rules, states, events и context map |
| [`engineering/`](memory-bank/engineering/README.md) | Архитектура, frontend, тестирование, coding style, git workflow и границы автономии агента |
| [`ops/`](memory-bank/ops/README.md) | Локальная разработка, окружения, конфигурация, релизы и runbooks |
| [`prd/`](memory-bank/prd/README.md) | Продуктовые инициативы между общим product context и отдельными фичами |
| [`epics/`](memory-bank/epics/README.md) | Крупные инициативы с roadmap, рисками, решениями и delivery subissues |
| [`use-cases/`](memory-bank/use-cases/README.md) | Канонические пользовательские и операционные сценарии проекта |
| [`features/`](memory-bank/features/README.md) | Пакеты отдельных delivery-фич |
| [`adr/`](memory-bank/adr/README.md) | Архитектурные решения и причины их принятия |
| [`flows/`](memory-bank/flows/README.md) | Lifecycle-процессы и шаблоны документов |
| [`prompts/`](memory-bank/prompts/README.md) | Переиспользуемые промпты для типовых этапов работы |

Корневой [`memory-bank/README.md`](memory-bank/README.md) служит основным индексом после установки шаблона в проект.

## Выберите свой сценарий

### Адаптировать существующий проект (brownfield)

В существующем проекте Memory Bank сначала должен отразить реальное состояние продукта и разработки, а не желаемую картину.

1. Скопируйте каталог `memory-bank/` в корень проекта.
2. Добавьте в `AGENTS.md`, `CLAUDE.md` или аналогичный файл инструкцию начинать работу с `memory-bank/README.md`.
3. Проведите inventory существующего кода, документации, терминов, архитектурных решений и процессов.
4. Адаптируйте `product/`, `domain/`, `engineering/` и `ops/`. Не выдумывайте отсутствующие знания: отмечайте пробелы и вопросы явно.
5. Перенесите устойчивые сценарии в `use-cases/`, а значимые принятые решения — в ADR.
6. Проверьте подход на одной реальной задаче или фиче, прежде чем описывать весь проект.
7. Запустите аудит ссылок и индексации.

### Начать новый проект (greenfield)

В новом проекте Memory Bank помогает сначала определить проверяемые границы и правила, а затем переходить к реализации.

1. Скопируйте `memory-bank/` и подключите его через файл инструкций агента.
2. Зафиксируйте vision, пользователей, ожидаемые результаты и метрики в `product/`.
3. Создайте начальные glossary, domain model, rules и context map в `domain/`.
4. Определите инженерные и операционные ограничения в `engineering/` и `ops/`; значимые технологические решения оформляйте как ADR.
5. Опишите первую инициативу через PRD или epic и выделите канонические use cases.
6. Сначала проверяйте `Small Change` gate. Остальные пользовательские изменения разбивайте на vertical-slice feature packages и проводите через `brief.md → optional design.md → implementation-plan.md`.
7. Обновляйте постоянный контекст только по мере появления проверенных знаний.

### Выполнять задачи через Memory Bank и `start-issue`

Memory Bank не заменяет task tracker или инструмент запуска агента. Он хранит контекст и правила, GitHub issue задаёт конкретную работу, а [`start-issue`](https://github.com/dapi/start-issue) создаёт для неё отдельные ветку и worktree и запускает выбранного coding agent.

```text
Memory Bank
контекст, правила, требования и способы проверки
        ↓
GitHub Issue
задача и ссылки на нужные документы
        ↓
start-issue
branch → worktree → agent session
        ↓
реализация → проверки → PR → evidence
        ↓
обновление Memory Bank при появлении новых знаний
```

Рабочий цикл:

1. Подготовьте GitHub issue с ожидаемым результатом и ссылками на применимые PRD, use case, feature package или ADR.
2. Установите `start-issue` по [инструкции проекта](https://github.com/dapi/start-issue#install) и настройте предпочитаемого агента командой `start-issue setup`.
3. При необходимости сначала проверьте планируемые действия без создания worktree:

   ```bash
   start-issue 123 --dry-run
   ```

4. Запустите задачу в изолированном worktree:

   ```bash
   start-issue 123
   ```

5. Агент читает issue и связанные документы Memory Bank, реализует изменение и выполняет предусмотренные проверки.
6. Завершите работу через PR и приложите требуемые evidence. Если в ходе задачи появились новые устойчивые правила, ограничения или решения, обновите их canonical owner в Memory Bank.

Во всех сценариях выбирайте подходящий workflow в [`memory-bank/flows/routing.md`](memory-bank/flows/routing.md). Если issue полностью задаёт intent, scope и acceptance, решение не требует design-документов и все routing predicates выполнены, задача может пройти как `Small Change` напрямую к реализации. Поддерживайте индексы и относительные ссылки и запускайте [локальные проверки](#проверка-ссылок-и-индексации) перед коммитом.

Главное правило адаптации: содержимое этого репозитория должно оставаться generic. Специфика конкретного продукта живёт только в его downstream-копии `memory-bank/` и не возвращается в шаблон.

## Task Routing

После получения issue выбирайте workflow в порядке, заданном в [`memory-bank/flows/routing.md`](memory-bank/flows/routing.md):

```text
Issue / Task
     │
     ▼
Incident / PIR?
     ├── да ─────────────► Incident Flow
     └── нет
          │
          ▼
         Bug?
          ├── да ────────► Bug Fix Flow
          └── нет
               │
               ▼
       Issue достаточен,
       design и plan не нужны?
               ├── да ───► Small Change Flow
               └── нет
                    │
                    ▼
          Новое или изменённое
          пользовательское поведение?
                    ├── да ───► Feature Flow
                    └── нет
                         │
                         ▼
                    Refactoring?
                         ├── да ─► Refactoring Flow
                         └── нет ► Human Routing
```

После выбора route следуйте его canonical документу: [Incident](memory-bank/flows/incident.md), [Bug Fix](memory-bank/flows/bug-fix.md), [Small Change](memory-bank/flows/small-change.md), [Feature](memory-bank/flows/feature.md) или [Refactoring](memory-bank/flows/refactoring.md).

### Small Change Routing Record

`Small Change` не создаёт отдельных документов в Memory Bank, но не остаётся бесследным. До реализации зафиксируйте в issue/task или draft PR, почему design и plan не нужны и как будет проверен результат:

```text
Workflow: Small Change

Design: not required
Reason: решение следует существующему паттерну <ссылка или путь>.

Plan: not required
Reason: change surface локален, порядок шагов и checkpoints не нужны.

Verify:
- <команда или проверка>
- <ожидаемый результат или evidence>
```

Если появляется необходимость выбрать подход, изменить contract, спланировать зависимые этапы или зафиксировать новый устойчивый project fact, остановите прямую реализацию и повторите routing.

## Как выбирать артефакт

- **Локальное изменение (`Small Change`)** — issue/task владеет intent, scope, acceptance и routing record; отдельные `brief.md`, `design.md` и `implementation-plan.md` не создаются.
- **Устойчивая продуктовая или операционная ситуация** — заведите `UC-*` в `use-cases/`.
- **Продуктовая инициатива, объединяющая несколько фич** — создайте PRD.
- **Крупная delivery-инициатива с roadmap и рисками** — используйте epic.
- **Единица пользовательской ценности, не прошедшая `Small Change` gate** — создайте feature package `features/FT-XXX/`.
- **Архитектурное или повторно используемое решение с альтернативами** — зафиксируйте ADR.

Feature package начинается с `README.md` и `brief.md`. `design.md` добавляется только тогда, когда решение требует отдельного проектирования. `implementation-plan.md` появляется после готовности upstream-документов и не должен самостоятельно изобретать требования или архитектурные решения.

`Small Change` проверяется до Feature Flow. Поэтому небольшое пользовательское улучшение может пройти напрямую, если issue достаточен и все routing predicates выполнены; иначе оно оформляется как vertical-slice feature package.

## Стартовые промпты

Адаптировать шаблон под проект:

```text
Прочитай ./memory-bank/README.md и governance-ядро в ./memory-bank/dna/.
Помоги адаптировать product, domain, engineering и ops под этот проект.
Не переноси project-specific детали обратно в generic-шаблон.
```

Создать feature package:

```text
Прочитай ./memory-bank/README.md и ./memory-bank/flows/feature.md.
Создай feature package для этой задачи, начиная с README.md и brief.md.
design.md создавай только по правилам Design Requirement Decision,
а implementation-plan.md — только после готовности upstream-документов.
```

Проверить качество Memory Bank:

```text
Проведи ревью ./memory-bank на SSoT, противоречия, broken links,
orphan-документы, недостающие README-индексы и неясные зависимости.
Предложи минимальные правки и запусти локальные проверки.
```

## Проверка ссылок и индексации

Скрипт [`scripts/check_memory_bank_index.py`](scripts/check_memory_bank_index.py) проверяет:

- неработающие относительные Markdown-ссылки;
- orphan-документы, не включённые в навигацию;
- достижимость документов от заданных entrypoint-файлов;
- слишком глубокие навигационные цепочки;
- наличие ожидаемых `README.md`-индексов.

Обычный запуск из корня этого репозитория:

```bash
python3 scripts/check_memory_bank_index.py
```

Exit code `0` означает, что ошибок нет. Предупреждения о глубине навигации не делают проверку неуспешной. Ненулевой exit code означает, что перед PR нужно исправить найденные ошибки.

Основные параметры:

- `--max-depth N` — порог глубины навигации; по умолчанию `3`;
- `--entrypoint PATH` — явная точка входа; параметр можно повторять;
- `--scope-root DIR` — проверяемый каталог; по умолчанию `memory-bank`;
- `--repo-root DIR` — корень downstream-репозитория;
- `--json` — структурированный отчёт для агента или другого инструмента.

Пример проверки навигации от нескольких точек входа:

```bash
python3 scripts/check_memory_bank_index.py \
  --entrypoint README.md \
  --entrypoint AGENTS.md \
  --max-depth 4
```

Если скрипт не скопирован в downstream-проект, его можно запустить напрямую из репозитория:

```bash
curl -fsSL https://raw.githubusercontent.com/dapi/memory-bank/main/scripts/check_memory_bank_index.py \
  | python3 - --repo-root .
```

Или установить локальную копию:

```bash
mkdir -p ./tools
curl -fsSL \
  -o ./tools/check_memory_bank_index.py \
  https://raw.githubusercontent.com/dapi/memory-bank/main/scripts/check_memory_bank_index.py
chmod +x ./tools/check_memory_bank_index.py
python3 ./tools/check_memory_bank_index.py --repo-root .
```

Запускайте аудит после добавления, удаления или переименования Markdown-файлов, после изменения индексов и перед PR с правками структуры документации.

## Развитие шаблона

Источники полезных практик для развития `memory-bank`:

- [`dapi/zelma`](https://github.com/dapi/zelma);
- [`brandymint/merchantly`](https://github.com/brandymint/merchantly);
- [`alfagen/mercury`](https://github.com/alfagen/mercury).

При переносе практик из downstream-репозиториев добавляйте только обобщаемые правила и шаблоны. Названия продуктов, инфраструктурные детали и другие project-specific факты не должны попадать в этот репозиторий.

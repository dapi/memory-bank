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

Репозиторий не содержит приложения или runtime-кода. Это generic-шаблон документационного контура, дополненный CLI, который проверяет ссылки, индексацию и достижимость документов.

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
6. Работу крупнее одной delivery-feature с общим roadmap, cross-feature risks или несколькими delivery units ведите через epic. Для каждой отдельной delivery-unit сначала проверяйте `Small Change` gate; остальные пользовательские и плановые infrastructure/engineering/operations изменения проводите через feature package: `brief.md → optional design.md → implementation-plan.md`.
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

1. Подготовьте GitHub issue с ожидаемым результатом и ссылками на применимые PRD, epic, use case, feature package или ADR.
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
          Работа крупнее одной delivery-feature,
          нужен общий roadmap, cross-feature
          risk register или несколько units?
                    ├── да ───► Epic Flow
                    └── нет
                         │
                         ▼
                    Refactoring?
                         ├── да ─► Refactoring Flow
                         └── нет
                              │
                              ▼
                    Одна delivery-unit меняет
                    пользовательское поведение
                    или доставляет planned
                    engineering/operations outcome?
                              ├── да ─► Feature Flow
                              └── нет ► Human Routing
```

После выбора route следуйте его canonical документу: [Incident](memory-bank/flows/incident.md), [Bug Fix](memory-bank/flows/bug-fix.md), [Small Change](memory-bank/flows/small-change.md), [Epic](memory-bank/flows/epic.md), [Refactoring](memory-bank/flows/refactoring.md) или [Feature](memory-bank/flows/feature.md).

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

## Результат каждого flow

Flow считается завершённым не после создания артефактов или изменения кода, а после достижения наблюдаемого результата и предъявления evidence из его `Outcome / Exit Contract`.

Для любого flow, изменяющего репозиторий, финальное evidence также включает последний review cycle без открытых замечаний, все изменения закоммичены и отправлены в remote branch, required CI полностью зелёный.

| Flow | Observable outcome | Required evidence | Terminal state и handoff |
| --- | --- | --- | --- |
| [Task Routing](memory-bank/flows/routing.md) | Выбран один допустимый flow или Human Routing | Route и подтверждающие entry predicates в issue/task, PR или incident record | `Routed` → выбранный flow; `Human Gate` → решение человека и повторный routing |
| [Small Change](memory-bank/flows/small-change.md) | Локальный acceptance выполнен без design и plan | Routing record, код, coverage, verify results, PR и CI | `Done` → закрыть task; новые факты и follow-up вернуть в routing |
| [Bug Fix](memory-bank/flows/bug-fix.md) | Expected behavior восстановлено и защищено от regression | Reproduction, root cause, regression evidence, PR и CI | `Resolved` → закрыть report; product/contract changes маршрутизировать отдельно |
| [Feature](memory-bank/flows/feature.md) | Одна пользовательская или infrastructure/engineering/operations delivery-unit принята end-to-end | `brief.md`, optional `design.md`, archived plan, `CHK-*`/`EVID-*`, tests, PR и CI | `Done` или `Cancelled` → закрыть delivery issue и передать follow-up |
| [Refactoring](memory-bank/flows/refactoring.md) | Структура улучшена при сохранении поведения либо принят research outcome | Baseline, characterization/regression evidence, before/after или research artifact | `Done` → закрыть task; behavior/contract changes повторно маршрутизировать |
| [Incident / PIR](memory-bank/flows/incident.md) | Impact прекращён, recovery подтверждён, PIR принят | Timeline, recovery signals, RCA, remediation evidence и follow-up references | `Closed` → закрыть incident; каждый prevention item маршрутизировать отдельно |
| [Epic](memory-bank/flows/epic.md) | Инициатива завершена через управляемые slices с outcome verdict | Charter, финальные roadmap/subissue/feature states, optional decision log, risks и follow-up refs | `Done` или `Cancelled` → закрыть initiative, перенести знания и маршрутизировать остаток |

## Как выбирать артефакт

- **Локальное изменение (`Small Change`)** — issue/task владеет intent, scope, acceptance и routing record; отдельные `brief.md`, `design.md` и `implementation-plan.md` не создаются.
- **Устойчивая продуктовая или операционная ситуация** — заведите `UC-*` в `use-cases/`.
- **Продуктовая инициатива, объединяющая несколько фич** — создайте PRD как owner продуктовых требований; если delivery требует общего roadmap, cross-feature risks или нескольких units, дополнительно используйте epic.
- **Крупная delivery-инициатива с roadmap и рисками** — используйте epic.
- **Одна пользовательская или плановая infrastructure/engineering/operations delivery-unit, не прошедшая `Small Change` gate** — создайте feature package `features/FT-XXX/`.
- **Архитектурное или повторно используемое решение с альтернативами** — зафиксируйте ADR.

Feature package начинается с `README.md` и `brief.md`. `design.md` добавляется только тогда, когда решение требует отдельного проектирования. `implementation-plan.md` появляется после готовности upstream-документов и не должен самостоятельно изобретать требования или архитектурные решения.

`Small Change` проверяется до governed delivery flows. Поэтому небольшое пользовательское или infrastructure/engineering/operations изменение может пройти напрямую, если issue достаточен и все routing predicates выполнены. Работа с несколькими delivery units сначала оформляется как epic; одна оставшаяся delivery-unit проходит Feature Flow, если это не behavior-preserving refactoring.

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

## Codex Goal Example

Для ограниченного этапа работы с epic можно использовать `/goal`, чтобы Codex
сохранял целевое lifecycle-состояние и критерии готовности в рамках одного чата.
Например, чтобы подготовить epic к передаче первой feature:

```text
/goal Подготовить EP-042 к статусу Roadmap Ready:
заполнить и согласовать charter, roadmap, subissues, risks и при необходимости
decision log; не придумывать неподтверждённые факты; остановиться на human gate
для существенных решений; убедиться, что первая feature может быть создана без
изобретения epic-level фактов.
```

Не используй один `/goal` для неопределённой цели вроде «полностью реализовать
epic». Вместо этого ставь отдельную конечную цель для каждого lifecycle-этапа,
delivery wave или feature handoff. `/goal` удерживает итоговое состояние, а
governance и owner-документы epic остаются источником правил и фактов.

## Проверка ссылок и индексации

Go CLI [`memory-bank-lint`](cmd/memory-bank-lint/main.go) аудирует `memory-bank/` и проверяет:

- broken relative markdown links внутри audit scope;
- orphan-документы, на которые никто не ссылается внутри scope;
- достижимость каждого документа от entrypoint'ов по индексной навигации;
- документы, которые достижимы только глубже порога навигации;
- contract ожидаемых `README.md`-индексов.

### Установка CLI

Для регулярных проверок установите CLI один раз. Требуется Go версии `1.21` или новее:

```bash
go install github.com/dapi/memory-bank/cmd/memory-bank-lint@latest
```

`go install` помещает бинарник в `GOBIN` или `GOPATH/bin`; этот каталог должен находиться в `PATH`. Повторите ту же команду, чтобы обновить CLI до актуальной версии.

На macOS или Linux проверьте, что установленная команда доступна:

```bash
command -v memory-bank-lint
```

Если репозиторий уже клонирован и нужно установить версию из текущего checkout:

```bash
go install ./cmd/memory-bank-lint
```

### Запуск

После установки запускайте CLI из корня репозитория:

```bash
memory-bank-lint --repo-root .
```

Что означает результат:

- exit code `0` — errors не найдены; warnings по глубине возможны, но аудит считается пройденным;
- non-zero exit code — найдены проблемы, которые нужно исправить до PR;
- `--json` — структурированный отчёт, пригодный для последующей автоматической доиндексации другим агентом или инструментом.

Параметры запуска:

- `--max-depth N` — порог глубины индексной навигации в прыжках; по умолчанию `3`; документы глубже порога попадают в warning, а не в error;
- `--entrypoint PATH` — явный entrypoint для аудита; параметр repeatable; принимает repo-relative или scope-relative пути; неоднозначные пути без префикса сначала резолвятся внутри `--scope-root`, а для явного repo-root пути используйте `./PATH` или `/PATH`; если передан, используется вместо дефолтного `memory-bank/README.md`;
- `--scope-root DIR` — меняет audit scope; по умолчанию `memory-bank`;
- `--repo-root DIR` — явно задаёт корень репозитория; полезно для сетевого запуска или локально установленной копии скрипта;
- `--json` — печатает только JSON-отчёт.

Примеры:

```bash
memory-bank-lint --repo-root . --max-depth 4
```

```bash
memory-bank-lint --repo-root . \
  --entrypoint README.md \
  --entrypoint AGENTS.md \
  --max-depth 4
```

Для разовой проверки без установки CLI:

```bash
go run github.com/dapi/memory-bank/cmd/memory-bank-lint@latest --repo-root .
```

Установку из GitHub Releases или через Homebrew используйте только после фактической публикации соответствующих release assets и Cask; если их ещё нет, используйте `go install`.

Когда запускать:

- после добавления, удаления или переименования `.md`-файлов в `memory-bank/`;
- после правок `README.md`-индексов и относительных ссылок;
- перед открытием PR с изменениями в template navigation или document structure.

## Развитие шаблона

Источники полезных практик для развития `memory-bank`:

- [`dapi/zelma`](https://github.com/dapi/zelma);
- [`brandymint/merchantly`](https://github.com/brandymint/merchantly);
- [`alfagen/mercury`](https://github.com/alfagen/mercury).

При переносе практик из downstream-репозиториев добавляйте только обобщаемые правила и шаблоны. Названия продуктов, инфраструктурные детали и другие project-specific факты не должны попадать в этот репозиторий.

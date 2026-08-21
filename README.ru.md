# Memory Bank — governance-ядро AI Software Development OS

[English version](README.md)

**`AGENTS.md` объясняет агенту, как начать. Memory Bank сохраняет смысл проекта, принятые решения и способ доказать результат.**

## Начать с агентом

Скопируй запрос для своего типа проекта. Он поручает агенту принести и
адаптировать Memory Bank; полный lifecycle определяют связанные протоколы.

### Greenfield

```text
Это новый проект. Прочитай repository instructions, README и docs, затем
выполни https://github.com/dapi/memory-bank/blob/main/docs/greenfield-integration-protocol.md.
Добавь Memory Bank в текущий репозиторий, адаптируй product, domain, engineering
и ops только из подтверждённых facts и создай initial PRD. Не выдумывай
пользователей, требования, метрики, архитектуру или delivery plan; не реализуй
продуктовые фичи. В конце покажи изменённые документы, источники, результаты
validation и open questions.
```

### Brownfield

```text
Это существующий проект. Выполни
https://github.com/dapi/memory-bank/blob/main/docs/brownfield-adaptation-protocol.md.
До шага установки из этого протокола исследуй только существующие repository
instructions, docs, code, manifests, CI/CD, configuration, runbooks и
historical ADR. Зафиксируй evidence, sources, confidence, conflicts и open
questions в intake PRD; не выдумывай architecture или delivery plan. Затем
добавь и адаптируй Memory Bank из тех же evidence, проверь результат и сообщи
изменённые документы, verification и remaining gaps.
```

Для воспроизводимого запуска замени `main` в URL протокола на immutable commit
SHA.

## Сделать фичу

После адаптации Memory Bank передай агенту задачу и этот запрос:

```text
Прочитай эту задачу, repository instructions, ./memory-bank/README.md и
./memory-bank/flows/routing.md. До изменений выбери минимальный допустимый
route. Если это Feature Flow, прочитай ./memory-bank/flows/feature.md; создай
feature package с README.md и brief.md, design — только если он required, а
implementation plan — только когда готовы upstream-документы. Реализуй
согласованный scope, выполни validation выбранного flow и обнови canonical
owner, если появилось новое устойчивое знание. Сообщи route, artifacts,
verification и open risks. Не мержи, не деплой и не меняй внешние системы,
пока задача явно этого не разрешает.
```

[Повседневная работа](docs/usage.md) объясняет цикл «задача → flow → проверки»
и более короткие маршруты.

## Что это

Memory Bank — переносимый documentation-first шаблон для разработки ПО с AI-агентами. Его копируют в проект и адаптируют так, чтобы человек и агент одинаково понимали:

- какой продукт создаётся и для кого;
- как устроена предметная область;
- какие инженерные и операционные правила действуют;
- какие требования, сценарии и архитектурные решения приняты;
- как работа будет реализована и проверена.

Это не вики и не архив заметок, а версионируемая рабочая память проекта. Важный контекст хранится рядом с кодом, а не в голове разработчика или одноразовом чате.

Здесь **OS** — метафора operating system для процесса разработки. Memory Bank работает как control plane: задаёт контекст, правила, lifecycle и критерии готовности. Task tracker, agent runner, coding agent, Git и CI образуют execution layer вокруг него.

## Как это работает

`dna/` задаёт governance-ядро: Single Source of Truth, зависимости между документами, lifecycle, frontmatter и правила навигации. Постоянный контекст проекта находится в `product/`, `domain/`, `engineering/` и `ops/`; research, инициативы, сценарии и решения — в Research, PRD, epic, use case и ADR.

Для значимой delivery-фичи контекст созревает поэтапно:

```text
brief.md                 design.md                  implementation-plan.md
что и зачем       →      какое решение       →     как реализовать и проверить
problem space            solution space             execution space
                         (если требуется)
```

Документы не должны дублировать друг друга. Код владеет реализацией, а Memory Bank — намерением, требованиями, обоснованием решений и контрактами.

Агенту передаётся компактный стартовый контекст и ссылки на нужные owner-документы. Если сессия исчерпала контекст, новую можно начать с той же задачи: важные факты, решения и способы проверки остаются в Memory Bank.

## Что находится в шаблоне

В этом source-репозитории payload хранится в `template/`. Агент переносит в
корень downstream-репозитория tracked regular files из этого каталога: например,
`template/memory-bank/` становится `memory-bank/`, а `template/init.sh` —
`./init.sh`. Имя `template/` в проект-получатель не переносится.

| Каталог | Назначение |
| --- | --- |
| [`dna/`](template/memory-bank/dna/README.md) | Governance-ядро: SSoT, frontmatter, lifecycle и правила связей между документами |
| [`product/`](template/memory-bank/product/README.md) | Vision, customers, metrics, marketing и roadmap |
| [`domain/`](template/memory-bank/domain/README.md) | Glossary, domain model, business rules, states, events и context map |
| [`engineering/`](template/memory-bank/engineering/README.md) | Архитектура, тестирование, coding style, git workflow и границы автономии агента |
| [`ops/`](template/memory-bank/ops/README.md) | Локальная разработка, окружения, конфигурация, релизы и runbooks |
| [`prd/`](template/memory-bank/prd/README.md) | Продуктовые инициативы между общим product context и отдельными фичами |
| [`research/`](template/memory-bank/research/README.md) | Evidence-backed market, product и technical research до решения о delivery |
| [`epics/`](template/memory-bank/epics/README.md) | Крупные инициативы с roadmap, рисками, решениями и delivery subissues |
| [`use-cases/`](template/memory-bank/use-cases/README.md) | Канонические пользовательские и операционные сценарии |
| [`features/`](template/memory-bank/features/README.md) | Пакеты отдельных delivery-фич |
| [`adr/`](template/memory-bank/adr/README.md) | Архитектурные решения и причины их принятия |
| [`flows/`](template/memory-bank/flows/README.md) | Lifecycle-процессы и шаблоны документов |
| [`prompts/`](template/memory-bank/prompts/README.md) | Human-only каталог prompt-артефактов и его access contract |

После установки шаблона [`memory-bank/README.md`](template/memory-bank/README.md) становится основным индексом downstream-проекта.

Корневой [`template/init.sh`](template/init.sh) — portable bootstrap для
`mise`, Git submodules и `direnv`, если соответствующие файлы есть в проекте.
После установки адаптируй `./init.sh` под реальные dependency, database и
service setup-команды проекта; он намеренно не переносит `.env`-файлы.

## Внедрение в проект

В downstream-проект устанавливается каталог `memory-bank/`. Исходники CLI,
Go-модуль, CI и release-конфигурация этого репозитория не являются частью
шаблона приложения. Ownership lock и автоматизированные обновления —
опциональное расширение, а не условие базового внедрения.

Инструкция по внедрению охватывает:

- адаптацию существующего проекта (brownfield);
- запуск нового проекта (greenfield);
- настройку агента;
- локальную проверку и downstream CI.

Следуйте [инструкции по внедрению](docs/adoption.md).

## Выбор рабочего процесса

Каждая задача сначала проходит [Task Routing](template/memory-bank/flows/routing.md). Он направляет работу в Incident, Bug Fix, Research & Discovery, Small Change, Epic, Refactoring, Feature или на ручное решение.

Корневой README даёт только обзор. Условия входа, lifecycle, обязательные артефакты и exit contract принадлежат каноническим документам в [`memory-bank/flows/`](template/memory-bank/flows/README.md) и не дублируются здесь.

## Документация репозитория

| Документ | Для кого и зачем |
| --- | --- |
| [Внедрение Memory Bank](docs/adoption.md) | Для команд, подключающих шаблон к brownfield- или greenfield-проекту |
| [Brownfield adaptation protocol](docs/brownfield-adaptation-protocol.md) | Для evidence-backed адаптации существующего репозитория до и после установки Memory Bank |
| [Greenfield adaptation protocol](docs/greenfield-integration-protocol.md) | Для копирования шаблона, извлечения project facts из README и docs, адаптации Memory Bank и создания initial PRD |
| [Использование Memory Bank](docs/usage.md) | Для повседневной работы с задачами и AI-агентами после внедрения |
| [Праймеринг контекста](docs/context-priming.md) | Для подготовки AI-агента к конкретной задаче и сбора релевантного контекста |
| [User Story, Use Case и BDD-сценарии](docs/bdd-user-stories-and-use-cases.md) | Для разделения устойчивого сценария, delivery slice и проверяемых примеров поведения |
| [Опциональная CLI-автоматизация](docs/memory-bank.md) | Для безопасных обновлений и downstream CI при необходимости |
| [Глоссарий](docs/glossary.md) | Термины governance и структуры документации, используемые в этом репозитории |
| [Ownership и безопасные обновления](docs/ownership.md) | Для понимания lock schema, границ владения и conflict policy |
| [Managed-блок инструкций агента](docs/agent-instructions.md) | Для marker contract, doctor и выбора единственного agent instruction target |
| [Разработка репозитория](docs/development.md) | Для разработчиков шаблона |

Опциональный `memory-bank-cli` добавляет безопасные обновления, проверку links и
диагностику governance. Он разрабатывается и выпускается отдельно в
[`dapi/memory-bank-cli`](https://github.com/dapi/memory-bank-cli).

## Развитие шаблона

### Методические источники

- [MECE principle](https://en.wikipedia.org/wiki/MECE_principle) — определение принципа Mutually Exclusive, Collectively Exhaustive: непересекающиеся категории, которые вместе покрывают заявленную область;
- Dan North, [*Introducing BDD*](https://dannorth.net/blog/introducing-bdd/) — первичный источник Behaviour-Driven Development: уточнение требований через business value, concrete examples и executable acceptance scenarios в форме `Given / When / Then`;
- Philippe Kruchten, [*Architectural Blueprints — The “4+1” View Model of Software Architecture*](https://arxiv.org/abs/2006.04975) — первичный источник stakeholder-oriented проверки Logical, Process, Development и Physical views через driving scenarios; [краткий обзор](https://en.wikipedia.org/wiki/4%2B1_architectural_view_model);
- Nenad Medvidovic, Richard N. Taylor, [*A Classification and Comparison Framework for Software Architecture Description Languages*](https://ics.uci.edu/~taylor/documents/2000-ADLs-TSE.pdf) — источник архитектурной модели components, connectors и configurations.

### Downstream-репозитории и практические полигоны

Эти репозитории используют и адаптируют Memory Bank под конкретные проекты.
Опыт их эксплуатации может становиться источником обобщаемых правил для
шаблона, но их project-specific факты остаются в downstream-копиях.

- [`dapi/zelma`](https://github.com/dapi/zelma);
- [`brandymint/merchantly`](https://github.com/brandymint/merchantly);
- [`alfagen/mercury`](https://github.com/alfagen/mercury).

Добавляйте в шаблон только обобщаемые правила. Названия продуктов, инфраструктурные детали и другие project-specific факты должны оставаться в downstream-копии `memory-bank/`.

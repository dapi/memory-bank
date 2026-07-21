# Memory Bank — governance-ядро AI Software Development OS

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

`dna/` задаёт governance-ядро: Single Source of Truth, зависимости между документами, lifecycle, frontmatter и правила навигации. Постоянный контекст проекта находится в `product/`, `domain/`, `engineering/` и `ops/`; инициативы, сценарии и решения — в PRD, epic, use case и ADR.

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

| Каталог | Назначение |
| --- | --- |
| [`dna/`](memory-bank/dna/README.md) | Governance-ядро: SSoT, frontmatter, lifecycle и правила связей между документами |
| [`product/`](memory-bank/product/README.md) | Vision, customers, metrics, marketing и roadmap |
| [`domain/`](memory-bank/domain/README.md) | Glossary, domain model, business rules, states, events и context map |
| [`engineering/`](memory-bank/engineering/README.md) | Архитектура, тестирование, coding style, git workflow и границы автономии агента |
| [`ops/`](memory-bank/ops/README.md) | Локальная разработка, окружения, конфигурация, релизы и runbooks |
| [`prd/`](memory-bank/prd/README.md) | Продуктовые инициативы между общим product context и отдельными фичами |
| [`epics/`](memory-bank/epics/README.md) | Крупные инициативы с roadmap, рисками, решениями и delivery subissues |
| [`use-cases/`](memory-bank/use-cases/README.md) | Канонические пользовательские и операционные сценарии |
| [`features/`](memory-bank/features/README.md) | Пакеты отдельных delivery-фич |
| [`adr/`](memory-bank/adr/README.md) | Архитектурные решения и причины их принятия |
| [`flows/`](memory-bank/flows/README.md) | Lifecycle-процессы и шаблоны документов |
| [`prompts/`](memory-bank/prompts/README.md) | Переиспользуемые промпты для типовых этапов работы |

После установки шаблона [`memory-bank/README.md`](memory-bank/README.md) становится основным индексом downstream-проекта.

## Внедрение в проект

В downstream-проект обычно копируется только каталог `memory-bank/`. Исходники CLI, Go-модуль, CI и release-конфигурация этого репозитория не являются частью шаблона приложения.

Инструкция по внедрению охватывает:

- адаптацию существующего проекта (brownfield);
- запуск нового проекта (greenfield);
- настройку агента;
- локальную проверку и downstream CI.

Следуйте [инструкции по внедрению](docs/adoption.md).

## Выбор рабочего процесса

Каждая задача сначала проходит [Task Routing](memory-bank/flows/routing.md). Он направляет работу в Incident, Bug Fix, Small Change, Epic, Refactoring, Feature или на ручное решение.

Корневой README даёт только обзор. Условия входа, lifecycle, обязательные артефакты и exit contract принадлежат каноническим документам в [`memory-bank/flows/`](memory-bank/flows/README.md) и не дублируются здесь.

## Документация репозитория

| Документ | Для кого и зачем |
| --- | --- |
| [Внедрение Memory Bank](docs/adoption.md) | Для команд, подключающих шаблон к brownfield- или greenfield-проекту |
| [Greenfield adaptation protocol](docs/greenfield-integration-protocol.md) | Для копирования шаблона, извлечения project facts из README и docs, адаптации Memory Bank и создания initial PRD |
| [Использование Memory Bank](docs/usage.md) | Для повседневной работы с задачами и AI-агентами после внедрения |
| [Установка и использование `memory-bank`](docs/memory-bank.md) | Для пользователей CLI и downstream CI |
| [Разработка репозитория](docs/development.md) | Для разработчиков шаблона и CLI |

`memory-bank lint` проверяет broken links, orphan-документы, достижимость через индексную навигацию и contract ожидаемых `README.md`-индексов. Прежний `memory-bank-lint` временно остаётся совместимым entrypoint для существующей автоматизации.

## Развитие шаблона

Источники полезных практик:

- [`dapi/zelma`](https://github.com/dapi/zelma);
- [`brandymint/merchantly`](https://github.com/brandymint/merchantly);
- [`alfagen/mercury`](https://github.com/alfagen/mercury);
- Nenad Medvidovic, Richard N. Taylor, [*A Classification and Comparison Framework for Software Architecture Description Languages*](https://ics.uci.edu/~taylor/documents/2000-ADLs-TSE.pdf) — источник архитектурной модели components, connectors и configurations.

Добавляйте в шаблон только обобщаемые правила. Названия продуктов, инфраструктурные детали и другие project-specific факты должны оставаться в downstream-копии `memory-bank/`.

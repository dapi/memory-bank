# INTRO: внедрение агентского подхода через Memory Bank

## Назначение

Memory Bank - это не вики и не архив заметок. Это рабочая память проекта для людей и AI-агентов: единое место, где хранится контекст продукта, домена, инженерных правил, операционных процедур, use cases, фич и решений.

Главная задача Memory Bank - сделать агентную разработку воспроизводимой. Агент должен получать задачу не из устного контекста разработчика, а из документов, которые можно проверить, обновить, переиспользовать и связать между собой.

## Ключевая идея

У команды появляется слой документации, который отвечает на четыре вопроса:

- что мы строим и зачем;
- какие правила домена нельзя нарушать;
- как в этом проекте принято проектировать, писать, тестировать и релизить;
- по каким use cases мы проверяем, что фича действительно работает.

Если факт важен для реализации, тестирования, ревью или эксплуатации, он должен жить в Memory Bank, а не в голове разработчика и не в разовом чате с агентом.

## Структура Memory Bank

`dna/` - governance-ядро. Это проектная конституция: Single Source of Truth, дерево зависимостей документов, lifecycle, frontmatter, правила ссылок и статусов.

`product/` - продуктовый контекст: vision, customers, metrics, roadmap, marketing, общий контекст продукта.

`domain/` - доменная модель: glossary, rules, states, events, context map. Здесь фиксируются бизнес-термины и правила, которые агент не должен угадывать.

`prd/` - Product Requirements Documents. PRD нужен для продуктовой инициативы или capability, которая стоит между общим product context и несколькими downstream feature packages.

`epics/` - крупные инициативы: roadmap, decision log, risks и набор связанных delivery subissues. Epic не подменяет feature package.

`tasks/` - compact packages для managed bugfix/chore/refactor задач, которым нужен durable context, но не нужен full feature package.

`use-cases/` - канонические сценарии проекта. Это основной материал для постановки задач агентам и генерации тест-кейсов.

`prompts/` - reusable prompt-документы: исходная формулировка, улучшенная copyable-версия и контекст применения.

`flows/` - lifecycle flows и governed templates для task, PRD, use case, epic, feature package, ADR и process-документов.

`engineering/` - инженерные правила: архитектура, frontend, testing policy, coding style, autonomy boundaries, git workflow.

`ops/` - операционный контекст: development, stages, release, config и runbooks.

`features/` - фича-паки. Каждая значимая фича получает `README.md`, canonical `brief.md`, optional `design.md` и, после готовности upstream owners, derived `implementation-plan.md`.

`adr/` - архитектурные решения. ADR нужен, когда команда выбирает подход и хочет сохранить не только решение, но и причины.

## Как объяснить команде

Начинайте не с "нам нужна документация", а с проблемы агентной разработки:

AI-агент хорошо работает только тогда, когда контекст задачи точный, полный и проверяемый. Если контекст живет в чате, он теряется. Если он живет в голове разработчика, агент его не видит. Если он разбросан по разным файлам без правил, документы начинают противоречить друг другу.

Memory Bank решает это через три механизма:

- Single Source of Truth: каждый факт живет в одном месте;
- Source Dependency Tree: при конфликте понятно, какой документ главнее;
- Task Workflows, Task Flow и Feature Flow: задача получает подходящий уровень формализации; small bugfix/refactor/chore может остаться compact task, а средняя, большая или рискованная фича превращается в фича-пак, из которого агент может проектировать, реализовывать и проверять результат.

## Рабочий процесс внедрения

1. Скопировать шаблон `memory-bank/` в корень проекта.
2. Скопировать или установить `scripts/check_memory_bank_index.py`, если downstream-проект должен запускать аудит индексации локально.
3. Подключить `AGENTS.md` или `CLAUDE.md` так, чтобы агент сначала читал `memory-bank/README.md`, затем `memory-bank/dna/README.md`.
4. Адаптировать минимум четыре раздела: `product/`, `domain/`, `engineering/`, `ops/`.
5. Завести первые use cases для текущего эпика или фичи.
6. По use cases сформировать проверяемые тест-кейсы: автоматические и ручные.
7. Для малых локальных задач использовать минимальный workflow из `memory-bank/flows/workflows.md` и compact task profiles из `memory-bank/flows/task-flow.md`, не раздувая их до feature package без риска или необходимости.
8. Для средней, большой или рискованной фичи создавать feature package: `README.md` как routing-слой, `brief.md` как canonical owner problem space и verify, и `design.md`, если нужен solution-space owner.
9. После готовности `brief.md` и required `design.md` создавать `implementation-plan.md` как derived execution-документ с grounding, шагами, checkpoints и traceability к canonical IDs.
10. Запускать агента на задачу через подходящий governed artifact: task, use case, feature package, PRD или epic, а не через длинное устное объяснение.
11. После реализации обновлять Memory Bank, если появились новые факты, правила, риски или решения.

## Feature Flow

Фича-пак разделяет задачу на canonical problem space, conditional solution space и derived execution.

`brief.md` - canonical problem-space owner delivery-единицы. Он фиксирует problem, outcome, scope (`REQ-*`), non-scope (`NS-*`), constraints, design requirement decision, acceptance scenarios (`SC-*`), checks (`CHK-*`) и evidence contract (`EVID-*`).

`design.md` - conditional solution-space owner. Он создается только когда `brief.md` фиксирует `Design required: yes` и владеет selected design, contracts, invariants, failure modes и rollout/backout facts.

`README.md` внутри `features/FT-XXX/` - routing-слой. Он создается вместе с `brief.md` и добавляет ссылки на downstream-документы только после их появления.

`implementation-plan.md` появляется только после того, как `brief.md` и required `design.md` готовы. Это execution-документ: relevant paths, existing patterns, unresolved questions, test surfaces, execution environment, шаги, checkpoints и traceability к canonical IDs.

Раздел верификации в `brief.md` обязателен. Если агент не понимает, как проверить результат, он будет оптимизировать под "код написан", а не под "фича работает".

## Контекст и перезапуск агентов

Не нужно пытаться держать в контексте агента слишком много данных. Рабочая рамка: компактный стартовый контекст и дополнительный контекст по мере работы. Качество падает, когда контекст становится слишком большим и плохо структурированным.

Если сессия исчерпала контекст, это не авария. Новую сессию можно запустить тем же промтом, если весь важный контекст лежит в фича-паке и связанных документах Memory Bank.

## Что фиксировать обязательно

- Бизнес-правила и исключения.
- Объемы данных на продакшене, если они влияют на решение.
- Инфраструктурные факты: окружения, репозитории, доступы, ограничения.
- Use cases и ожидаемые сценарии проверки.
- ADR для спорных или рискованных технических решений.
- Запреты и правила проекта.
- Результаты исследований, которые потом пригодятся агентам.

## Пример с миграцией

Если фича затрагивает миграцию данных, главный риск - не "написать миграцию", а потерять или исказить данные.

Правильная постановка для агента:

- зафиксировать в Memory Bank, сколько данных есть на продакшене и какие таблицы затронуты;
- описать проектные правила для миграций, если они есть;
- добавить тест или проверку, которая подтверждает сохранность данных;
- описать в фича-паке verification steps для ручной и автоматической проверки.

## Настройка агентов

Для Codex, Claude, Cursor и других агентных инструментов принцип один: агент должен управляться Memory Bank, а не набором случайных плагинов и устных инструкций.

Минимальная настройка:

- оставить только инструменты, которые действительно нужны для задачи;
- использовать актуальную документацию по библиотекам, когда решение зависит от внешних API;
- проектные правила, use cases, инженерные политики и фича-паки читать из Memory Bank;
- все агентные промпты привязывать к конкретным документам, а не просить агента "догадаться".

## Проверка качества Memory Bank

Перед PR или после изменения структуры нужно запускать:

```bash
python3 scripts/check_memory_bank_index.py
git diff --check
```

Если скрипт установлен в downstream-проекте под другим путём, запускайте установленную копию, например `python3 ./tools/check_memory_bank_index.py --repo-root .`. Первый скрипт проверяет относительные ссылки, orphan-документы, достижимость документов через README-индексы и ожидаемые README-файлы. Второй ловит лишние пробелы и conflict markers.

## Стартовые промпты

Адаптация проекта:

```text
Прочитай ./memory-bank и предложи адаптацию AGENTS.md под правила этого шаблона. Сначала найди governance-ядро, затем проверь product, domain, engineering и ops. Не меняй код.
```

Создание use cases:

```text
Прочитай ./memory-bank/domain и ./memory-bank/product. Помоги сформировать use cases для текущего эпика. Для каждого use case укажи happy path, edge cases, данные, ожидаемый результат и как это проверить автоматически и вручную.
```

Фича-пак:

```text
Прочитай ./memory-bank/README.md и ./memory-bank/flows/feature-flow.md, затем создай feature package для задачи. Сначала инстанцируй README.md и brief.md; в brief.md обязательно зафиксируй scope, non-scope, Design Requirement Decision, acceptance scenarios, checks и evidence. implementation-plan.md создавай только после готовности upstream owners.
```

Ревью Memory Bank:

```text
Проведи ревью ./memory-bank на document governance: SSoT, противоречия, broken links, orphan-документы, недостающие README-индексы и неясные зависимости. Предложи минимальные правки.
```

## Definition of Done внедрения

- `AGENTS.md` маршрутизирует агента в Memory Bank.
- `product/`, `domain/`, `engineering/`, `ops/` адаптированы под проект.
- Для текущего эпика заведены use cases.
- По use cases есть автоматические или ручные проверки.
- Средние, большие или рискованные фичи оформляются через feature package; малые bugfix/refactor/chore задачи проходят через minimal workflow или compact task package.
- Спорные технические решения оформляются как ADR.
- Скрипт проверки индексации проходит без errors.
- Команда понимает: если факт нужен агенту, он должен быть записан в Memory Bank.

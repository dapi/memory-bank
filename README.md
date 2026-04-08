# Шаблон `memory-bank` для агентной разработки

Этот репозиторий теперь разделен на два слоя:

- `memory-bank/` — переносимый шаблон, который можно копировать в любой проект по разработке ПО;
- `examples/merchantly/` — конкретный instantiated пример с Merchantly-спецификой.

## Как использовать

1. Скопируйте каталог `./memory-bank` в корень своего проекта.
2. Прочитайте и адаптируйте в нем как минимум `domain/`, `engineering/` и `ops/`.
3. Если нужен ориентир, смотрите конкретный пример в `./examples/merchantly/`.

## Настроечные промпты для агента

Запукаются в новых сессиях

```text
Прочитай ./memory-bank и предложи адаптацию CLAUDE.md/AGENTS.md под правила этого шаблона.
```

```text
Прочитай ./memory-bank и помоги адаптировать секцию `domain`
```

```text
Прочитай ./memory-bank и помоги адаптировать секцию `ops`
```

```text
Прочитай ./memory-bank и помоги адаптировать секцию `engineering`
```

```text
Проведи ревью memory-bank на document governance
```
(внеси правки и повторить до состояния которое вас устроит)


```text
Проведи ревью memory-bank на консистетность, и непротиворечивость
```
(внеси правки и повторить до состояния которое вас устроит)


```text
Помоги создать PRD
```

```text
Помоги создать глоссарий
```

## Что есть внутри шаблона

- `memory-bank/dna/` — governance-ядро: SSoT, frontmatter, lifecycle, cross-references.
- `memory-bank/flows/` — lifecycle flows и шаблоны для PRD/feature/ADR.
- `memory-bank/domain/` — заготовки для product context, архитектуры и UI-слоя.
- `memory-bank/prd/` — место для instantiated Product Requirements Documents.
- `memory-bank/use-cases/` — место для instantiated project-level use cases.
- `memory-bank/engineering/` — testing policy, coding style, autonomy boundaries, git workflow.
- `memory-bank/ops/` — заготовки для development, stages, releases, config и runbooks.
- `memory-bank/adr/` — место для instantiated ADR.
- `memory-bank/features/` — место для instantiated feature packages.

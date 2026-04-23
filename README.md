# Шаблон `memory-bank` для агентной разработки

Этот репозиторий теперь разделен на два слоя:

- `memory-bank/` — переносимый шаблон, который можно копировать в любой проект по разработке ПО;
- `examples/merchantly/` — конкретный instantiated пример с Merchantly-спецификой.

## Как использовать

1. Скопируйте каталог `./memory-bank` в корень своего проекта.
2. Прочитайте и адаптируйте в нем как минимум `domain/`, `engineering/` и `ops/`.
3. Если нужен ориентир, смотрите конкретный пример в `./examples/merchantly/`.

## Локальные проверки

- `python3 scripts/check_memory_bank_index.py` — аудит достижимости markdown-документов, broken links и expected README-индексов внутри `memory-bank/`.
- `git diff --check` — проверка лишних пробелов и conflict markers перед PR.

### Проверка навигации `memory-bank`

Запускайте из корня репозитория:

```bash
python3 scripts/check_memory_bank_index.py
```

Зачем это нужно:

- поймать broken relative links внутри `memory-bank/` до PR;
- убедиться, что markdown-документы в `memory-bank/` достижимы из [`memory-bank/README.md`](memory-bank/README.md) через цепочку внутренних ссылок и не стали orphan;
- проверить, что ожидаемые `README.md`-индексы на месте и остаются index-документами с валидным `purpose`;
- быстрее находить последствия структурных изменений в шаблоне, когда документ перенесли, переименовали или забыли добавить в навигацию.

Что означает результат:

- exit code `0` — broken links, unreachable docs и нарушения expected index contract не найдены;
- non-zero exit code — скрипт нашёл проблему, которую нужно исправить до PR.

Если в downstream-репозитории нужно считать entrypoint не от `memory-bank/README.md`, а от других документов, передайте их явно:

```bash
python3 scripts/check_memory_bank_index.py \
  --entrypoint README.md \
  --entrypoint AGENTS.md
```

Когда запускать:

- после добавления, удаления или переименования `.md`-файлов в `memory-bank/`;
- после правок `README.md`-индексов и относительных ссылок;
- перед открытием PR с изменениями в template navigation или document structure.

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
У нас в проекте подключен memory-bank. Я хочу быть уверен что все страницы в этом memory-bank-а так или иначе доступны через нидексацию начиная с
AGENTS.md. Если страница не упомянются напрямую, то она упомянутся в файле который упомянут в файле который упомянут в AGENTS.md и так далее на глубину до 4-х шагов.
```

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

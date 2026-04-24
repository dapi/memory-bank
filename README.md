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

### Аудит ссылок и индексации `memory-bank`

Скрипт [`scripts/check_memory_bank_index.py`](scripts/check_memory_bank_index.py) аудирует `memory-bank/` и проверяет:

- broken relative markdown links внутри audit scope;
- orphan-документы, на которые никто не ссылается внутри scope;
- достижимость каждого документа от entrypoint'ов по индексной навигации;
- документы, которые достижимы только глубже порога навигации;
- contract ожидаемых `README.md`-индексов.

Обычный локальный запуск из корня репозитория:

```bash
python3 scripts/check_memory_bank_index.py
```

Что означает результат:

- exit code `0` — errors не найдены; warnings по глубине возможны, но аудит считается пройденным;
- non-zero exit code — найдены проблемы, которые нужно исправить до PR;
- `--json` — структурированный отчёт, пригодный для последующей автоматической доиндексации другим агентом или инструментом.

Параметры запуска:

- `--max-depth N` — порог глубины индексной навигации в прыжках; по умолчанию `3`; документы глубже порога попадают в warning, а не в error;
- `--entrypoint PATH` — явный entrypoint для аудита; параметр repeatable; принимает repo-relative или scope-relative пути; если передан, используется вместо дефолтного `memory-bank/README.md`;
- `--scope-root DIR` — меняет audit scope; по умолчанию `memory-bank`;
- `--repo-root DIR` — явно задаёт корень репозитория; полезно для сетевого запуска или локально установленной копии скрипта;
- `--json` — печатает только JSON-отчёт.

Примеры:

```bash
python3 scripts/check_memory_bank_index.py --max-depth 4
```

```bash
python3 scripts/check_memory_bank_index.py \
  --entrypoint README.md \
  --entrypoint AGENTS.md \
  --max-depth 4
```

Быстрый запуск по сети без предварительной установки:

```bash
curl -fsSL https://raw.githubusercontent.com/dapi/memory-bank/main/scripts/check_memory_bank_index.py \
  | python3 - --repo-root .
```

Локальная установка или копирование с GitHub:

1. Скопируйте файл со страницы `scripts/check_memory_bank_index.py` на GitHub или скачайте raw-версию:

```bash
mkdir -p ./tools
curl -fsSL \
  -o ./tools/check_memory_bank_index.py \
  https://raw.githubusercontent.com/dapi/memory-bank/main/scripts/check_memory_bank_index.py
chmod +x ./tools/check_memory_bank_index.py
```

2. Запускайте его из корня downstream-репозитория:

```bash
python3 ./tools/check_memory_bank_index.py --repo-root .
```

`macOS` и `Linux`: команды запуска одинаковые. Отличие только в том, куда класть локальную копию, если хочется вызывать скрипт без относительного пути: на Linux чаще используют `~/.local/bin`, на macOS — `~/bin` или любой каталог, добавленный в `PATH`. Если не хотите менять `PATH`, запускайте скрипт через `python3` по полному или относительному пути.

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

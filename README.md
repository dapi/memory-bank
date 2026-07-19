# Шаблон `memory-bank` для агентной разработки

Этот репозиторий содержит переносимый шаблон:

- `memory-bank/` — документационный шаблон, который можно копировать в любой проект по разработке ПО.

## Введение

- [INTRO.md](INTRO.md) — интро-инструкция по внедрению агентского подхода через Memory Bank: как объяснять шаблон команде, запускать Feature Flow и ставить задачи агентам через проверяемый контекст.

## Источники пополнения `memory-bank`

- [`dapi/zelma`](https://github.com/dapi/zelma)
- [`brandymint/merchantly`](https://github.com/brandymint/merchantly)
- [`alfagen/mercury`](https://github.com/alfagen/mercury)


```prompt
Изучи что нового и полезного для нашего memory-bank появилось в репозиториях источниках
```

## Как использовать

1. Скопируйте каталог `./memory-bank` в корень своего проекта.
2. Прочитайте и адаптируйте в нем как минимум `product/`, `domain/`, `engineering/` и `ops/`.

## Локальные проверки

- `go run ./cmd/memory-bank-lint` — аудит достижимости markdown-документов, broken links и expected README-индексов внутри `memory-bank/`.
- `go test ./...` — тесты валидатора и проверка стабильности JSON-контракта.
- `git diff --check` — проверка лишних пробелов и conflict markers перед PR.

### Аудит ссылок и индексации `memory-bank`

Go CLI [`memory-bank-lint`](cmd/memory-bank-lint/main.go) аудирует `memory-bank/` и проверяет:

- broken relative markdown links внутри audit scope;
- orphan-документы, на которые никто не ссылается внутри scope;
- достижимость каждого документа от entrypoint'ов по индексной навигации;
- документы, которые достижимы только глубже порога навигации;
- contract ожидаемых `README.md`-индексов.

Обычный локальный запуск из корня репозитория:

```bash
go run ./cmd/memory-bank-lint
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
go run ./cmd/memory-bank-lint --max-depth 4
```

```bash
go run ./cmd/memory-bank-lint \
  --entrypoint README.md \
  --entrypoint AGENTS.md \
  --max-depth 4
```

Быстрый запуск без предварительной установки CLI:

```bash
go run github.com/dapi/memory-bank/cmd/memory-bank-lint@latest --repo-root .
```

Установка CLI из GitHub:

```bash
go install github.com/dapi/memory-bank/cmd/memory-bank-lint@latest
```

Установка готового релиза через Homebrew:

```bash
brew install dapi/tap/memory-bank-lint
```

Готовые бинарники для Linux, macOS и Windows публикуются в [GitHub Releases](https://github.com/dapi/memory-bank/releases) при создании тега `v*`. Каждый релиз содержит `checksums.txt`; версия доступна через `memory-bank-lint --version`.

Для публикации Homebrew Cask в `dapi/homebrew-tap` release workflow ожидает repository secret `HOMEBREW_TAP_GITHUB_TOKEN` с правом записи содержимого tap-репозитория.

После установки запускайте его из корня downstream-репозитория:

```bash
memory-bank-lint --repo-root .
```

Для сборки бинарника из клонированного репозитория:

```bash
go build -o ./memory-bank-lint ./cmd/memory-bank-lint
```

Для разработки нужен Go версии `1.21` или новее. `go install` помещает бинарник в `GOBIN` или `GOPATH/bin`; этот каталог должен находиться в `PATH`. Команды запуска одинаковы на macOS и Linux.

Когда запускать:

- после добавления, удаления или переименования `.md`-файлов в `memory-bank/`;
- после правок `README.md`-индексов и относительных ссылок;
- перед открытием PR с изменениями в template navigation или document structure.

## Настроечные промпты для агента

Запукаются в новых сессиях

```text
Прочитай memory-bank - https://github.com/dapi/memory-bank/ склонируй его в наш репозиторий и адаптируй под наш проект
```

```text
Прочитай ./memory-bank и помоги адаптировать секцию `product`
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
- `memory-bank/product/` — заготовки для product context, vision, customers, metrics, marketing и roadmap.
- `memory-bank/domain/` — заготовки для glossary, domain model, rules, states, events и context map.
- `memory-bank/prd/` — место для instantiated Product Requirements Documents.
- `memory-bank/use-cases/` — место для instantiated project-level use cases.
- `memory-bank/engineering/` — architecture patterns, frontend engineering, testing policy, coding style, autonomy boundaries, git workflow.
- `memory-bank/ops/` — заготовки для development, stages, releases, config и runbooks.
- `memory-bank/adr/` — место для instantiated ADR.
- `memory-bank/features/` — место для instantiated feature packages.

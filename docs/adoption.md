# Внедрение Memory Bank в проект

Этот документ описывает базовое подключение Memory Bank к существующему или новому проекту. Агент переносит tracked regular files из source `template/` в корень downstream-проекта: `template/memory-bank/` становится `memory-bank/`, а `template/init.sh` — `./init.sh`. Source-repository metadata, `.github/` и `docs/` вне `template/` не являются частью шаблона приложения.

В репозитории `dapi/memory-bank` исходный payload находится в `template/`.
Это только upstream-префикс: агент не переносит его в downstream. Downstream
инструкции должны ссылаться на реальные пути, например `memory-bank/` и
`./init.sh`, но не на `template/...`.

## Что копировать

Минимальный установленный комплект включает:

```text
memory-bank/
init.sh
```

Базовое внедрение не требует отдельной утилиты, lock-файла или CI. Для
ownership-aware обновлений и автоматических проверок есть
[опциональная CLI-автоматизация](memory-bank.md).

## Адаптировать существующий проект (brownfield)

Для существующего проекта следуйте [brownfield adaptation protocol](brownfield-adaptation-protocol.md). Он начинает с evidence-backed discovery **до** установки и чтения `memory-bank/`, затем описывает intake PRD, adaptation canonical owners, governed conversion, validation и real-task trial.

Не заменяйте protocol кратким inventory: порядок важен, потому что generic template не является источником project facts до завершения discovery.

Для запуска передайте агенту [copyable brownfield prompt](brownfield-adaptation-protocol.md#copyable-codex-prompt).

## Начать новый проект (greenfield)

Для нового GitHub-проекта используйте [протокол адаптации Memory Bank](greenfield-integration-protocol.md). Он поручает Codex изучить существующие README и docs, скопировать generic-шаблон, максимально заполнить его подтверждёнными фактами о продукте и проекте и создать initial PRD.

Для запуска передайте агенту [copyable greenfield prompt](greenfield-integration-protocol.md#copyable-codex-prompt). Если нужен shell-запуск Codex, выполните в корне downstream-репозитория:

```bash
codex --search \
  "Прочитай протокол адаптации Memory Bank по адресу https://github.com/dapi/memory-bank/blob/main/docs/greenfield-integration-protocol.md и выполни его в текущем репозитории."
```

Команда передаёт prompt при запуске интерактивной Codex-сессии и использует sandbox и approval policy из пользовательской конфигурации. Для воспроизводимого запуска замените `main` в URL на immutable commit SHA.

## Подключить агента

Добавьте в repository instructions указание читать `memory-bank/README.md`,
`memory-bank/dna/README.md` и `memory-bank/flows/routing.md`, не копируя
governance. Project-specific инструкции остаются в собственном owner-файле.
Опциональный managed-block contract описан в [инструкциях агента](agent-instructions.md).

После установки Memory Bank для первой адаптации можно использовать запрос:

```text
Прочитай ./memory-bank/README.md и governance-ядро в ./memory-bank/dna/.
Помоги адаптировать product, domain, engineering и ops под этот проект.
Не переноси project-specific детали обратно в generic-шаблон.
```

После внедрения используйте [инструкцию по повседневной работе](usage.md): она описывает связь Memory Bank с task tracker и agent runner, рабочий цикл и стартовые запросы.

## Опциональная CLI-проверка

Для локального аудита установите закреплённый release `memory-bank-cli` из отдельного репозитория [`dapi/memory-bank-cli`](https://github.com/dapi/memory-bank-cli/releases). Выберите asset для своей ОС и архитектуры, проверьте его по опубликованному `checksums.txt` и добавьте бинарник в `PATH`.

После этого из любого места внутри downstream Git-репозитория:

```bash
memory-bank-cli lint
memory-bank-cli doctor
```

Интеграционный контракт и ссылки на command documentation: [`memory-bank.md`](memory-bank.md).

## Опциональный CI с CLI

CI-проверка Memory Bank должна быть opt-in в downstream-проекте. Не копируйте `.github/workflows/ci.yml` из этого репозитория: он проверяет source template в profile `template`, а downstream-проекту нужен auto/downstream profile.

Перед подключением выберите конкретный CLI release и checksum. Плавающая версия может изменить результат проверки без изменений в downstream-репозитории.

Минимальный GitHub Actions workflow для downstream-проекта:

```yaml
name: Memory Bank

on:
  pull_request:
  push:
    branches: [main]

permissions:
  contents: read

jobs:
  lint:
    runs-on: ubuntu-latest
    env:
      MEMORY_BANK_CLI_VERSION: v1.0.0
      MEMORY_BANK_CLI_SHA256: 35300dd2f713e904a5819a31b064e12809b93921a557fbfb3f2543c45c733c57
    steps:
      - name: Checkout
        uses: actions/checkout@v4

      - name: Install memory-bank-cli
        run: |
          asset="memory-bank-cli-linux-amd64"
          url="https://github.com/dapi/memory-bank-cli/releases/download/${MEMORY_BANK_CLI_VERSION}/${asset}"
          curl --fail --location --silent --show-error "$url" --output "$RUNNER_TEMP/$asset"
          echo "${MEMORY_BANK_CLI_SHA256}  $RUNNER_TEMP/$asset" | sha256sum --check
          chmod +x "$RUNNER_TEMP/$asset"
          mkdir -p "$RUNNER_TEMP/memory-bank-cli-bin"
          mv "$RUNNER_TEMP/$asset" "$RUNNER_TEMP/memory-bank-cli-bin/memory-bank-cli"
          echo "$RUNNER_TEMP/memory-bank-cli-bin" >> "$GITHUB_PATH"

      - name: Audit Memory Bank
        run: memory-bank-cli doctor
```

При обновлении версии замените tag и checksum значениями из одного release. Для другой runner architecture выберите соответствующий asset и SHA-256 из того же `checksums.txt`.

## Definition of Done внедрения

Memory Bank считается внедрённым, когда:

- `memory-bank/` находится в корне downstream-проекта;
- постоянный контекст `product/`, `domain/`, `engineering/` и `ops/` отражает фактические правила проекта или явно помечает пробелы;
- агентские инструкции указывают читать `memory-bank/README.md` и governance-ядро;
- первая реальная задача прошла через выбранный flow или `Small Change` routing record;
- ссылки, README-индексы и frontmatter проверены;
- опциональные CLI-аудит и CI подключены, если команда хочет automated checks,
  безопасные update или блокировку PR при broken links.

# Внедрение Memory Bank в проект

Этот документ описывает, как подключить Memory Bank к существующему или новому проекту. В downstream-проект устанавливается каталог `memory-bank/`, а внутри него создаётся служебный `memory-bank/.lock`; dev-инфраструктура этого репозитория (`tools/`, `.github/`, `docs/`) не является частью шаблона приложения. Lock создаёт `memory-bank init`, а не upstream template. Его нужно коммитить: он хранит версию источника и ownership-границу для безопасных обновлений.

## Что копировать

Минимальный установленный комплект:

```text
memory-bank/
```

Подробный ownership-контракт и процедура обновления: [`ownership.md`](ownership.md).

## Адаптировать существующий проект (brownfield)

Для существующего проекта следуйте [brownfield adaptation protocol](brownfield-adaptation-protocol.md). Он начинает с evidence-backed discovery **до** установки и чтения `memory-bank/`, затем описывает intake PRD, adaptation canonical owners, governed conversion, validation и real-task trial.

Не заменяйте protocol кратким inventory: порядок важен, потому что generic template не является источником project facts до завершения discovery.

## Начать новый проект (greenfield)

Для нового GitHub-проекта используйте [протокол адаптации Memory Bank](greenfield-integration-protocol.md). Он поручает Codex изучить существующие README и docs, скопировать generic-шаблон, максимально заполнить его подтверждёнными фактами о продукте и проекте и создать initial PRD.

Для запуска выполните в корне downstream-репозитория:

```bash
codex --search \
  "Прочитай протокол адаптации Memory Bank по адресу https://github.com/dapi/memory-bank/blob/main/docs/greenfield-integration-protocol.md и выполни его в текущем репозитории."
```

Команда передаёт prompt при запуске интерактивной Codex-сессии и использует sandbox и approval policy из пользовательской конфигурации. Для воспроизводимого запуска замените `main` в URL на immutable commit SHA.

## Подключить агента

Используйте managed-блок, который устанавливает `memory-bank init`: он направляет агента к `memory-bank/README.md`, `memory-bank/dna/README.md` и `memory-bank/flows/routing.md`, не копируя governance. Не редактируйте содержимое между markers вручную; project-specific инструкции размещайте снаружи. Полный marker, update, doctor и alternative-target contract описан в [managed-блоках agent instructions](agent-instructions.md).

После установки Memory Bank для первой адаптации можно использовать запрос:

```text
Прочитай ./memory-bank/README.md и governance-ядро в ./memory-bank/dna/.
Помоги адаптировать product, domain, engineering и ops под этот проект.
Не переноси project-specific детали обратно в generic-шаблон.
```

После внедрения используйте [инструкцию по повседневной работе](usage.md): она описывает связь Memory Bank с task tracker и agent runner, рабочий цикл и стартовые запросы.

## Установить локальную проверку

Для локального аудита установите `memory-bank` как внешний бинарник:

```bash
go install github.com/dapi/memory-bank/tools/cmd/memory-bank@latest
```

После этого из любого места внутри downstream Git-репозитория:

```bash
memory-bank lint
memory-bank doctor
```

Подробности по флагам, migration path и установке: [`memory-bank.md`](memory-bank.md).

## Подключить CI

CI-проверка Memory Bank должна быть opt-in в downstream-проекте. Не копируйте `.github/workflows/ci.yml` из этого репозитория: он предназначен для разработки самого шаблона и собирает локальный Go CLI из `tools/cmd/memory-bank`.

Перед подключением выберите конкретный release tag или commit CLI. Если оставить `@latest`, новая версия `memory-bank` сможет изменить результат проверки без изменений в downstream-репозитории.

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
      # Замените значением вида v1.2.3 или полным commit SHA.
      MEMORY_BANK_VERSION: REPLACE_WITH_RELEASE_TAG_OR_COMMIT
    steps:
      - name: Checkout
        uses: actions/checkout@v4

      - name: Setup Go
        uses: actions/setup-go@v5
        with:
          go-version: '1.21'
          cache: false

      - name: Install memory-bank
        run: go install "github.com/dapi/memory-bank/tools/cmd/memory-bank@${MEMORY_BANK_VERSION}"

      - name: Audit Memory Bank
        run: memory-bank doctor
```

Если в проекте уже есть Go toolchain setup, можно переиспользовать существующий шаг `actions/setup-go`. Если Go в проекте не используется, он нужен только для установки CLI через `go install`; после публикации release binaries или Homebrew formula CI можно заменить на установку готового бинарника.

## Definition of Done внедрения

Memory Bank считается внедрённым, когда:

- `memory-bank/` находится в корне downstream-проекта;
- постоянный контекст `product/`, `domain/`, `engineering/` и `ops/` отражает фактические правила проекта или явно помечает пробелы;
- агентские инструкции указывают читать `memory-bank/README.md` и governance-ядро;
- первая реальная задача прошла через выбранный flow или `Small Change` routing record;
- `memory-bank doctor` проходит локально (и включает navigation checks `lint`);
- CI-проверка подключена, если команда хочет блокировать PR при broken links или нарушенной индексной навигации.

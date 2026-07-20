# Внедрение Memory Bank в проект

Этот документ описывает, как подключить Memory Bank к существующему или новому проекту. В downstream-проект обычно копируется только каталог `memory-bank/`; dev-инфраструктура этого репозитория (`cmd/`, `.github/`, `.goreleaser.yml`, `go.mod`, `docs/`) не является частью шаблона приложения.

## Что копировать

Минимальный переносимый комплект:

```text
memory-bank/
```

Для первичной установки на macOS или Linux выполните из корня downstream-проекта:

```bash
test ! -e ./memory-bank &&
  curl -fsSL https://github.com/dapi/memory-bank/archive/refs/heads/main.tar.gz |
  tar -xz --strip-components=1 memory-bank-main/memory-bank
```

Команда останавливается, если `memory-bank/` уже существует, и не перезаписывает адаптированные документы. На других платформах скачайте source archive этого репозитория и скопируйте из него только каталог `memory-bank/`.

Опционально добавьте в downstream-проект собственный `AGENTS.md`, `CLAUDE.md` или аналогичный файл с правилом начинать работу с:

```text
memory-bank/README.md
memory-bank/dna/README.md
```

Не копируйте весь репозиторий `dapi/memory-bank` как основу продукта, если вам не нужна разработка самого шаблона и CLI. Иначе в downstream-проект попадут CI/release-файлы шаблона, Go-модуль и исходники `memory-bank-lint`.

## Адаптировать существующий проект (brownfield)

В существующем проекте Memory Bank сначала должен отразить реальное состояние продукта и разработки, а не желаемую картину.

Цель brownfield-внедрения — сделать текущий контекст проекта видимым и проверяемым для людей и агентов. Не начинайте с идеального описания будущей архитектуры. Сначала зафиксируйте то, что уже влияет на разработку: реальные пользователи, термины, ограничения, интеграции, принятые решения, неочевидные правила и known gaps.

1. Скопируйте каталог `memory-bank/` в корень проекта.
2. Добавьте в `AGENTS.md`, `CLAUDE.md` или аналогичный файл инструкцию начинать работу с `memory-bank/README.md`.
3. Проведите inventory существующего кода, документации, терминов, архитектурных решений и процессов.
4. Адаптируйте `product/`, `domain/`, `engineering/` и `ops/`. Не выдумывайте отсутствующие знания: отмечайте пробелы и вопросы явно.
5. Перенесите устойчивые сценарии в `use-cases/`, а значимые принятые решения — в ADR.
6. Проверьте подход на одной реальной задаче или фиче, прежде чем описывать весь проект.
7. Запустите аудит ссылок и индексации.

### Brownfield inventory

Минимальный inventory перед первой адаптацией:

- README, wiki, runbooks, ADR, старые design docs;
- ключевые директории кода и границы модулей;
- production/staging/local окружения;
- внешние интеграции и владение credentials/config;
- основные пользовательские сценарии и операционные сценарии;
- термины, которые уже используются в коде, UI, API и команде;
- текущий CI/CD и обязательные проверки перед merge;
- известные технические долги, ограничения и опасные зоны.

### Brownfield порядок заполнения

1. `product/` — что продукт уже делает, для кого, какие outcomes и метрики реально важны.
2. `domain/` — glossary, domain model, states/events/rules из существующей системы.
3. `engineering/` — текущая архитектура, coding style, testing policy, frontend/backend conventions, git workflow.
4. `ops/` — локальный запуск, окружения, config, release process, runbooks.
5. `use-cases/` — только устойчивые сценарии, которые уже проверяются или должны проверяться.
6. `adr/` — решения, которые уже приняты и продолжают влиять на разработку.

Если факт неизвестен, пишите это явно: `Unknown`, `TBD`, `Needs owner confirmation`. Для агента это безопаснее, чем уверенная выдумка.

### Brownfield типичные ошибки

- описывать желаемую архитектуру как текущую;
- переносить в Memory Bank все старые документы без нормализации и ownership;
- создавать PRD/feature packages до описания базового product/domain/engineering context;
- дублировать один и тот же факт в нескольких местах;
- подключать CI до первого успешного локального `memory-bank-lint`;
- блокировать PR из-за устаревшей документации, которую команда ещё не готова исправлять.

### Brownfield готовность

Brownfield-внедрение достаточно для первого рабочего использования, когда:

- агент может понять, как проект устроен, из `memory-bank/README.md` и owner-документов;
- минимум `product/`, `domain/`, `engineering/` и `ops/` адаптированы под реальный проект;
- known gaps явно отмечены;
- одна реальная задача прошла через `Small Change`, feature package, bug fix или другой выбранный flow;
- `memory-bank-lint` проходит локально.

## Начать новый проект (greenfield)

В новом проекте Memory Bank помогает сначала определить проверяемые границы и правила, а затем переходить к реализации.

Цель greenfield-внедрения — создать минимальный управляемый контекст до того, как код и решения начнут расходиться. Здесь Memory Bank работает как scaffold для product/domain/engineering/ops решений, но он не должен превращаться в большой speculative design document.

1. Скопируйте `memory-bank/` и подключите его через файл инструкций агента.
2. Зафиксируйте vision, пользователей, ожидаемые результаты и метрики в `product/`.
3. Создайте начальные glossary, domain model, rules и context map в `domain/`.
4. Определите инженерные и операционные ограничения в `engineering/` и `ops/`; значимые технологические решения оформляйте как ADR.
5. Опишите первую инициативу через PRD или epic и выделите канонические use cases.
6. Работу крупнее одной delivery-feature с общим roadmap, cross-feature risks или несколькими delivery units ведите через epic. Для каждой отдельной delivery-unit сначала проверяйте `Small Change` gate; остальные пользовательские и плановые infrastructure/engineering/operations изменения проводите через feature package: `brief.md → optional design.md → implementation-plan.md`.
7. Обновляйте постоянный контекст только по мере появления проверенных знаний.

### Greenfield минимальный стартовый комплект

Перед первой существенной реализацией заполните:

- `product/vision.md` — зачем существует продукт и какой outcome ожидается;
- `product/customers.md` — для кого продукт делается;
- `product/metrics.md` — как будет понятно, что результат полезен;
- `domain/glossary.md` — базовые термины без конфликтующих синонимов;
- `domain/model.md` — основные сущности и связи;
- `engineering/architecture.md` — начальные архитектурные границы и выбранный stack;
- `engineering/testing-policy.md` — какие проверки обязательны с первого дня;
- `ops/development.md` — как запускать проект локально;
- `ops/config.md` — как задаются переменные окружения и secrets;
- ADR для каждого решения, которое будет дорого менять.

### Greenfield порядок работы

1. Сначала product intent: users, problem, outcome, non-goals.
2. Затем domain language: glossary, entities, states, events, business rules.
3. Затем engineering constraints: stack, module boundaries, testing, CI, repo workflow.
4. Затем first initiative: PRD или epic, если работа шире одной delivery-feature.
5. Затем первая delivery-unit: feature package или `Small Change`, если issue полностью достаточен.

### Greenfield типичные ошибки

- проектировать слишком много до первой проверки реальным use case;
- фиксировать технологические предпочтения без ADR и trade-offs;
- создавать feature packages без product/domain owner-фактов;
- писать acceptance criteria, которые невозможно проверить;
- откладывать testing policy и CI до “потом”;
- смешивать template rules и project-specific решения.

### Greenfield готовность

Greenfield-внедрение достаточно для старта разработки, когда:

- `product/`, `domain/`, `engineering/` и `ops/` задают минимальные проверяемые границы;
- первая инициатива имеет понятный owner: issue, PRD, epic или feature package;
- выбранные технологии и дорогие решения зафиксированы в ADR;
- локальный запуск и проверки описаны;
- `memory-bank-lint` проходит локально;
- CI либо подключён, либо явно отложен с причиной.

## Подключить агента

Добавьте в `AGENTS.md`, `CLAUDE.md` или аналогичный файл правило начинать работу с `memory-bank/README.md` и `memory-bank/dna/README.md`. Если файл уже содержит project-specific инструкции, дополните их маршрутизацией в Memory Bank, не заменяя существующие правила проекта.

Для первой адаптации можно использовать запрос:

```text
Прочитай ./memory-bank/README.md и governance-ядро в ./memory-bank/dna/.
Помоги адаптировать product, domain, engineering и ops под этот проект.
Не переноси project-specific детали обратно в generic-шаблон.
```

После внедрения используйте [инструкцию по повседневной работе](usage.md): она описывает связь Memory Bank с task tracker и agent runner, рабочий цикл и стартовые запросы.

## Установить локальную проверку

Для локального аудита установите `memory-bank-lint` как внешний бинарник:

```bash
go install github.com/dapi/memory-bank/cmd/memory-bank-lint@latest
```

После этого из любого места внутри downstream Git-репозитория:

```bash
memory-bank-lint
```

Подробности по флагам и установке: [`memory-bank-lint.md`](memory-bank-lint.md).

## Подключить CI

CI-проверка Memory Bank должна быть opt-in в downstream-проекте. Не копируйте `.github/workflows/ci.yml` из этого репозитория: он предназначен для разработки самого шаблона и собирает локальный Go CLI из `cmd/memory-bank-lint`.

Перед подключением выберите конкретный release tag или commit CLI. Если оставить `@latest`, новая версия `memory-bank-lint` сможет изменить результат проверки без изменений в downstream-репозитории.

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
      MEMORY_BANK_LINT_VERSION: REPLACE_WITH_RELEASE_TAG_OR_COMMIT
    steps:
      - name: Checkout
        uses: actions/checkout@v4

      - name: Setup Go
        uses: actions/setup-go@v5
        with:
          go-version: '1.21'
          cache: false

      - name: Install memory-bank-lint
        run: go install "github.com/dapi/memory-bank/cmd/memory-bank-lint@${MEMORY_BANK_LINT_VERSION}"

      - name: Audit Memory Bank
        run: memory-bank-lint
```

Если в проекте уже есть Go toolchain setup, можно переиспользовать существующий шаг `actions/setup-go`. Если Go в проекте не используется, он нужен только для установки CLI через `go install`; после публикации release binaries или Homebrew formula CI можно заменить на установку готового бинарника.

## Definition of Done внедрения

Memory Bank считается внедрённым, когда:

- `memory-bank/` находится в корне downstream-проекта;
- постоянный контекст `product/`, `domain/`, `engineering/` и `ops/` отражает фактические правила проекта или явно помечает пробелы;
- агентские инструкции указывают читать `memory-bank/README.md` и governance-ядро;
- первая реальная задача прошла через выбранный flow или `Small Change` routing record;
- `memory-bank-lint` проходит локально;
- CI-проверка подключена, если команда хочет блокировать PR при broken links или нарушенной индексной навигации.

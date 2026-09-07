# Memory Bank

<p align="center">
  <img src="docs/assets/memory-bank-mark.svg" alt="Memory Bank: знания проекта, владение документами и опциональные процессы разработки" width="180">
</p>

**Проектная документация под контролем версий, с ясным владением и опциональными процессами AI-разработки.**

[English version](README.md) · [Компонентное внедрение](docs/component-adoption.md) ·
[Интеграция CLI](docs/memory-bank.md)

## Выберите глубину внедрения

Memory Bank состоит из трёх компонентов с односторонними зависимостями:

| Компонент | Ответственность | Зависимости |
| --- | --- | --- |
| DNA | Единственный источник истины, владение, публикационные статусы, metadata и навигация | Нет |
| Documents | Типы документов, базовые шаблоны и разделы проекта | DNA |
| Flows | AI routing, priming, этапы разработки, gates и расширения документов | DNA + Documents |

DNA работает самостоятельно. Documents можно использовать без AI-процесса и
runner. Позднее подключение Flows сохраняет проектные документы; документ
подключается к процессу только через явную операцию adoption.

| Набор | Компоненты | Адаптеры инструментов |
| --- | --- | --- |
| `core` | DNA | Добавляются явно |
| `docs` | DNA + Documents | Добавляются явно |
| `full` | DNA + Documents + Flows | Добавляются явно |
| `legacy` | Все три | Прежние интеграции включены |

Новая установка без выбора набора использует `legacy` для совместимости.
Pull без параметров выбора сохраняет записанный состав. Удаление компонентов
не поддерживается. У адаптеров есть зависимости: выбор адаптера может добавить Flows.

## Установите в проект

Нужны Git и компонентный `memory-bank-cli` на Linux или macOS. Поддержка
компонентов — согласованное изменение шаблона и CLI: используйте проверенный
CLI candidate или release, который объявляет обе нужные capabilities. Умения
старого release устанавливать legacy-шаблоны недостаточно.

Из чистого checkout шаблона на закреплённом коммите запустите защищённую точку
входа для своего проекта:

```bash
memory-bank-cli capabilities --require components/v1 --require adoption/v1
./tools/install-components.sh init \
  --repo-root /path/to/project --preset docs
```

Точка входа проверяет capabilities до вызова installer и закрепляет собственный
коммит источника. Проверьте результат:

```bash
git -C /path/to/project status --short
git -C /path/to/project diff --check
memory-bank-cli doctor --repo-root /path/to/project
```

Для существующей legacy-установки нужна отдельная просмотренная миграция.
Обычный pull не является согласием на неё. См. [внедрение и миграцию](docs/component-adoption.md).

## Создавайте проектные документы

Documents содержит ADR, feature brief, PRD, use case, research brief и epic
charter. Базовые шаблоны находятся в `memory-bank/templates/`, контракты типов —
в `memory-bank/document-types/`.

```bash
memory-bank-cli document create --repo-root /path/to/project \
  --type feature --path memory-bank/features/FT-123/brief.md
```

Новый документ принадлежит проекту и не получает flow adoption, в том числе
в `full` и `legacy`. Заполните проблему, результат, scope и критерии приёмки.
Базовый ADR содержит контекст, варианты, решение, последствия и
`decision_status` без обязательного AI-процесса согласования.

## Подключайте AI-процессы по мере необходимости

```bash
./tools/install-components.sh pull \
  --repo-root /path/to/project --preset full
```

После установки Flows выбирайте процесс через `memory-bank/flows/routing.md`.
Подготовьте документ к выбранному расширению и подключите его явно:

```bash
memory-bank-cli document adopt --repo-root /path/to/project \
  --path memory-bank/features/FT-123/brief.md --contract feature/v1
```

Adoption проверяет применимые требования до изменения состояния. Базовому brief
могут понадобиться поля и разделы процесса. Устойчивая идентичность документа,
контракт и digest неизменяемого bundle записываются в проектный registry;
frontmatter является проверяемой проекцией. Установка Flows сама по себе не
включает gates для всех feature briefs.

[Быстрый старт](docs/quick-start.md) и [повседневная работа](docs/usage.md)
описывают разработку с Flows. Эти руководства сейчас доступны на русском языке.

## Знания и владение

У канонического факта один владелец. Производные документы ссылаются на него;
код владеет реализацией, документы — намерением, обоснованием и контрактами.
Memory Bank применяет First Principles Framework, чтобы явно фиксировать
предположения, ограничения, решения и evidence.

Контекст проекта находится в `product/`, `domain/`, `engineering/` и `ops/`.
Требования, сценарии и решения — в `prd/`, `use-cases/`, `features/`, `research/`,
`epics/` и `adr/`. CLI обновляет шаблонные assets, сохраняя содержимое,
принадлежащее проекту. Новая версия контракта требует явного перехода документа;
подмена bundle под прежним ID даёт conflict.

## Опциональная автоматизация

Адаптеры инструментов отделены от компонентов документации:

- `codex` устанавливает определения агентов Codex;
- `start-issue` устанавливает инструкции запуска задач;
- `symphony` устанавливает workflow и скрипты запуска;
- `bootstrap` устанавливает bootstrap-скрипт.

Выбирайте адаптер повторяемым параметром `--adapter NAME`. Явные `core`, `docs`
и `full` не включают адаптеры автоматически; `legacy` сохраняет их.
Runners запускают агентов. Flows задаёт процесс их работы.

## Структура шаблона

Этот репозиторий владеет generic payload в `template/`. CLI устанавливает
выбранные файлы, убирая префикс `template/`. Manifest компонентов находится в
[`template/memory-bank/components.json`](template/memory-bank/components.json).
Сгенерированный downstream `memory-bank/README.md` перечисляет установленные
разделы; AGENTS направляет читателя только к установленным компонентам.

| Раздел | Назначение |
| --- | --- |
| [`dna/`](template/memory-bank/dna/README.md) | Самостоятельное governance-ядро |
| [`document-types/`](template/memory-bank/document-types/README.md) | Базовые контракты документов |
| [`templates/`](template/memory-bank/templates/README.md) | Управляемые шаблоны для проектных документов |
| [`flows/`](template/memory-bank/flows/README.md) | Опциональные процессы и версионированные расширения |

Project-local `memory-bank/` этого репозитория является проекцией payload;
реальными файлами остаются собственные материалы проекта. У проекции нет
installed-template lock.

## Справочные материалы

- [Компонентное внедрение и legacy-миграция](docs/component-adoption.md)
- [Wire-контракт компонентов](docs/component-wire-format.md)
- [Владение и безопасное обновление](docs/ownership.md)
- [Managed agent instructions](docs/agent-instructions.md)
- [Интеграция CLI и проверка source profile](docs/memory-bank.md)
- [Разработка репозитория](docs/development.md)

CLI разрабатывается отдельно в
[`dapi/memory-bank-cli`](https://github.com/dapi/memory-bank-cli). Шаблон доступен
под [Apache License 2.0](LICENSE).

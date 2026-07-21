# Универсальный протокол интеграции Memory Bank в greenfield-проект

## Назначение и область применимости

Этот протокол описывает воспроизводимое внедрение Memory Bank в новый программный проект, размещённый на GitHub. Он не зависит от языка приложения, framework, build system или типа продукта. Project-specific технологии и команды фиксируются при адаптации `engineering/` и `ops/`, а не в самом протоколе.

Memory Bank становится версионируемым project-specific source of truth для intent, domain rules, engineering constraints, operations и delivery evidence. Он направляет людей и агентов, но по умолчанию не является runtime dependency создаваемого продукта.

Протокол не применяется без адаптации: generic template задаёт структуру и governance, а downstream-репозиторий владеет фактами конкретного проекта.

## Как запустить протокол

Откройте Codex в корне downstream-репозитория и отправьте короткий prompt:

```text
Прочитай протокол адаптации Memory Bank по адресу
https://github.com/dapi/memory-bank/blob/main/docs/greenfield-integration-protocol.md
и выполни его в текущем репозитории.
```

Для воспроизводимого запуска замените `main` в URL на immutable commit SHA. Дополнительные инструкции в prompt не требуются: scope, порядок фаз, gates, остановка при blocking ambiguity и Definition of Done принадлежат самому протоколу.

## Инварианты

1. Каждый устойчивый факт имеет одного canonical owner; код продолжает владеть implementation details.
2. Неизвестное отмечается как `Unknown`, `TBD`, open question или human gate, но не заполняется предположениями.
3. Upstream template и downstream knowledge имеют разное ownership. Обновление шаблона не перезаписывает project-specific документы целиком.
4. Agent instructions дополняют существующие правила репозитория, а не заменяют их.
5. Любая реализация сначала проходит [`Task Routing`](../memory-bank/flows/routing.md).
6. Локальный lint должен пройти до включения блокирующей CI-проверки.
7. Интеграция Memory Bank в runtime, prompts или автоматическую finalization отдельно маршрутизируется как изменение project contract.

## Входные данные и решения

До начала интеграции зафиксируйте в GitHub issue или draft PR:

| Вход | Что определить |
| --- | --- |
| Downstream repository | GitHub repository и default branch |
| Template source | `dapi/memory-bank` и immutable commit SHA или release tag |
| Knowledge sources | README, product brief, issue, research или решения владельца, на которых основана первая адаптация |
| Owners | Кто подтверждает product, domain, engineering и operations facts |
| Agent entrypoint | `AGENTS.md`, `CLAUDE.md` или другой поддерживаемый агентом файл |
| Initial delivery | Первая реальная задача или инициатива, на которой будет проверено внедрение |
| CI policy | Сразу блокировать PR по lint либо сначала запустить non-blocking adoption period |

Если product boundary, первый outcome или owner критических решений неизвестны, остановитесь на human gate. Неполнота второстепенных деталей не блокирует bootstrap: она фиксируется как явный gap в соответствующем canonical owner.

## Роли и ownership

| Роль | Ответственность |
| --- | --- |
| Integration owner | Ведёт integration PR, provenance, gates и итоговый evidence |
| Project owner | Подтверждает product boundary, первый outcome и решения на human gates |
| Knowledge owners | Подтверждают факты своих `product/`, `domain/`, `engineering/` и `ops/` слоёв |
| Reviewer | Независимо проверяет отсутствие unsupported claims, template leakage и непреднамеренной runtime coupling |

Один человек может совмещать несколько ролей, но review integration PR должен оставаться отдельным проверяемым действием.

## Lifecycle

```text
Proposed
   ↓
Imported
   ↓
Agent Connected
   ↓
Context Ready
   ↓
Workflow Ready
   ↓
Locally Validated
   ↓
Validated
   ↓
Adopted
```

Переход разрешён только после выполнения gate текущей фазы. Если обнаружена неверная исходная информация, вернитесь к ближайшему upstream gate и обновите зависимые документы.

## Фаза 1. Import and provenance

1. Выберите immutable upstream revision. Не используйте плавающую ветку `main` как зафиксированную версию интеграции.
2. Скопируйте только каталог `memory-bank/` в корень downstream-репозитория.
3. Создайте `memory-bank/UPSTREAM.md` и запишите:
   - upstream URL;
   - immutable revision;
   - дату и способ импорта;
   - downstream ownership;
   - update policy.
4. Добавьте `UPSTREAM.md` в аннотированный индекс [`memory-bank/README.md`](../memory-bank/README.md), чтобы provenance record проходил навигационный аудит.
5. Проверьте условия лицензирования перед распространением шаблона за пределами разрешённого scope.

Минимальный совместимый frontmatter для provenance record:

```yaml
---
title: Memory Bank Upstream Record
doc_kind: project
doc_function: record
purpose: Фиксирует происхождение шаблона и правила downstream-обновления.
derived_from:
  - dna/governance.md
status: active
---
```

Не перезаписывайте существующий `memory-bank/`. Если каталог уже есть, остановите bootstrap и определите, требуется ли update, recovery незавершённой интеграции или ручное объединение.

### Gate: Imported

- `memory-bank/` находится в корне downstream-репозитория;
- `UPSTREAM.md` содержит immutable source revision и update policy;
- все добавленные документы доступны из индексной навигации;
- template files ещё не выдают placeholder content за факты проекта.

## Фаза 2. Agent entrypoint

Добавьте в корневой агентский instruction file минимальный контракт:

```markdown
## Memory Bank

Перед изменением репозитория прочитай `memory-bank/README.md`, затем только
документы, релевантные задаче. До реализации проведи задачу через
`memory-bank/flows/routing.md`.

Memory Bank владеет intent, domain rules, architecture and operations contracts
и delivery evidence; код владеет implementation details. Не выдумывай
неизвестные требования. Существенную неоднозначность зафиксируй у canonical
owner и вынеси на human decision.
```

Добавьте project-specific правила проверки и ограничения автономии после этого общего контракта. Не копируйте в agent instructions весь Memory Bank: entrypoint должен направлять к индексам через progressive disclosure.

### Gate: Agent connected

- агентский entrypoint ссылается на `memory-bank/README.md` и Task Routing;
- существующие инструкции репозитория сохранены;
- правило разделения ownership между кодом и документацией сформулировано явно;
- неизвестные требования приводят к gap или human gate, а не к догадке агента.

## Фаза 3. Project context adaptation

Адаптируйте persistent layers только по подтверждённым источникам:

1. `product/` — problem, users, outcomes, non-goals и измеримые признаки успеха.
2. `domain/` — glossary, entities, states, events, rules и bounded contexts, которые уже нужны первой инициативе.
3. `engineering/` — architecture boundaries, quality attributes, testing policy, repository workflow и правила выбранного проекта.
4. `ops/` — local development, configuration and secrets boundaries, environments, release expectations и runbook needs.

Для каждого устойчивого решения определите canonical owner. Дорогое или reusable архитектурное решение оформите как ADR. Не пытайтесь заполнить весь шаблон до начала разработки: удалите неприменимые placeholder sections или оставьте их draft с явным статусом.

Язык, framework и команды проекта записываются в project-specific `engineering/` и `ops/`. Они не изменяют lifecycle этого протокола.

### Gate: Context Ready

- vision, users, первый outcome и non-goals понятны;
- glossary устраняет критичные конфликтующие термины;
- минимальные architecture, testing, local development и config boundaries описаны;
- дорогие принятые решения имеют ADR;
- все blocking unknowns либо решены, либо находятся на human gate;
- между template examples и project facts нет неявного смешения.

## Фаза 4. Delivery bootstrap

1. Создайте или выберите первую реальную GitHub issue.
2. Проведите её через Task Routing.
3. Используйте выбранный flow: `Small Change`, Feature, Epic, Bug Fix, Refactoring или другой допустимый route.
4. Создавайте feature package только если route действительно требует Feature Flow. Не назначайте первой delivery-unit заранее заданный тип или идентификатор.
5. Для применимого delivery flow выберите validation profile и зафиксируйте required evidence в его canonical owner.
6. Свяжите delivery artifacts с подтверждёнными product/domain/engineering/ops owners, не копируя их содержание.

### Gate: Workflow Ready

- первая задача имеет ровно один route;
- выполнены entry gates выбранного flow;
- validation profile и verify contract зафиксированы;
- blocking design decisions имеют owner;
- документы содержат достаточно контекста для новой agent session без обращения к прежнему чату.

## Фаза 5. Local validation

Установите или запустите `memory-bank` на зафиксированной версии и выполните из корня downstream-репозитория:

```bash
memory-bank lint
git diff --check
```

Способ распространения validator может использовать отдельный toolchain и не задаёт язык приложения. Конкретный поддерживаемый способ установки и command contract описаны в [`memory-bank.md`](memory-bank.md).

Исправьте broken links, orphan documents, отсутствующие README indexes и недостижимые owner-документы. Inline-ссылка в обзорном тексте не заменяет запись в аннотированном индексе, если lint contract требует index navigation.

### Gate: Locally Validated

- `memory-bank lint` завершается с exit code `0`;
- `git diff --check` не находит ошибок;
- выполнен semantic read-through изменённых project facts;
- unsupported claims и project-specific template leakage отсутствуют;
- версия validator воспроизводима.

## Фаза 6. GitHub CI

Добавьте GitHub Actions workflow, который запускается при изменениях Memory Bank, agent instructions и самого workflow. До включения required check убедитесь, что та же зафиксированная версия validator проходит локально.

CI должен:

- использовать минимальные `contents: read` permissions;
- устанавливать зафиксированную версию validator, а не плавающий `latest`;
- выполнять `memory-bank lint` из корня repository;
- запускаться на pull request и push в фактическую default branch;
- изменяться отдельным reviewable diff при обновлении validator.

Текущий пример GitHub Actions и правила pinning находятся в [`adoption.md`](adoption.md#подключить-ci). Toolchain шага установки не определяет язык downstream-приложения.

### Gate: Validated

- локальный и CI lint используют одну версию validator;
- GitHub Actions workflow завершился успешно на integration PR;
- команда явно решила, является ли check required;
- failure CI оставляет понятный diagnostic и не изменяет downstream files.

## Фаза 7. Practical adoption

Проведите первую delivery-задачу до terminal state выбранного flow:

1. новая agent session начинает работу с project entrypoint;
2. агент находит canonical context через индексы;
3. реализация и evidence соответствуют выбранному flow и validation profile;
4. PR проходит review и required CI;
5. новые устойчивые факты возвращаются в их canonical owners;
6. временные execution details не переносятся в постоянный контекст.

### Gate: Adopted

- первая реальная задача завершена с доступным GitHub delivery trace;
- другой человек или новая agent session может восстановить intent, решения и verification path из repository state;
- `memory-bank lint` и required CI зелёные;
- open gaps имеют owner и следующий decision point;
- интеграция не создала runtime dependency продукта без отдельного решения.

## Обновление и восстановление

Upstream update выполняется как отдельная reviewable задача:

1. сравните текущую recorded revision с новой immutable revision;
2. выберите только generic governance и template changes, совместимые с downstream facts;
3. не заменяйте адаптированные каталоги целиком;
4. обновите `UPSTREAM.md`;
5. повторите local lint, semantic review и GitHub CI.

Если интеграция признана неудачной до adoption, откатите integration PR обычным Git revert. Если после adoption Memory Bank больше не используется, сначала сохраните или перенесите unique project knowledge, затем удаляйте agent routing, CI и каталог отдельным reviewable PR. Удаление каталога без knowledge inventory может уничтожить единственный canonical owner проектных решений.

## Definition of Done

Memory Bank интегрирован в greenfield-проект, когда выполнены все условия:

- provenance импортированной версии зафиксирован;
- persistent layers содержат минимальный проверяемый project context;
- agent entrypoint и Task Routing подключены;
- первая реальная задача завершена через выбранный flow;
- локальный и GitHub CI lint проходят на одной pinned version;
- update и rollback ownership понятны;
- project-specific unknowns видимы и не подменены предположениями;
- runtime, prompt injection или automatic finalization integration либо отсутствуют, либо отдельно спроектированы и приняты.

## Non-goals

- замена GitHub Issues, Git, pull requests, review или CI;
- выбор языка, framework, architecture style или build system проекта;
- обязательное создание feature package независимо от Task Routing;
- копирование project-specific facts обратно в generic upstream template;
- включение Memory Bank в runtime создаваемого продукта по умолчанию;
- полная спецификация будущего продукта до первого проверяемого use case.

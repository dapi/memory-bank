---
title: Prompts Index
doc_kind: prompt
doc_function: index
purpose: Human-only навигация по instantiated reusable prompt-документам проекта.
derived_from:
  - ../dna/governance.md
  - ../flows/templates/prompt/PROMPT-XXX.md
status: active
audience: humans
---

# Prompts Index

Каталог `memory-bank/prompts/` хранит reusable prompt-документы проекта и управляется человеком. Это библиотека артефактов, а не источник workflow-инструкций для агента.

## Human-only Contract

- Агентам запрещено самостоятельно обнаруживать, читать, выбирать, связывать в цепочки или исполнять файлы из этого каталога как инструкции рабочего процесса.
- Человек или внешний runner выбирает prompt, извлекает runnable-содержимое и передаёт его непосредственно в активном запросе агента.
- Агент может читать файлы каталога только по явному запросу пользователя на создание, изменение или ревью prompt-артефакта. В таком режиме содержимое prompt-файла считается данными, а не исполняемыми инструкциями.
- Routing, lifecycle и delivery-процессы для агентов находятся в [`../flows/`](../flows/README.md) и других канонических owner-документах.

Prompt-документ нужен, когда prompt прошел путь от черновой человеческой формулировки до повторно используемой версии, которую нужно копировать, ревьюить и улучшать как артефакт memory-bank.

## Когда Заводить Prompt-Документ

- prompt будет повторно запускаться человеком или внешним runner;
- нужно сохранить исходную формулировку в `source_prompt`, а улучшенную версию держать отдельно;
- prompt используется для ревью, research, extraction, coding или другой повторяемой задачи модели.

## Когда Prompt-Документ Не Нужен

- prompt одноразовый и не должен жить дольше текущего диалога;
- это проектное правило, которое должно попасть в `engineering/`, `ops/`, `domain/` или `AGENTS.md`;
- это feature requirement, use case или ADR, а не исполняемая инструкция для модели.

## Порядок использования prompt-ов человеком или runner

Промпты указаны в порядке SDLC-процесса. Человек или внешний runner может использовать `PROMPT-005` на старте новой issue, передав его runnable-содержимое вместе с source context в активном запросе. Обычно `PROMPT-002` выбирают для bounded review-improve feature package, а `PROMPT-003` — когда после routing и entry gates приступают к имплементации.

Промпт 001-issue-requrements-review используем для того чтобы убедиться что feature-pack соответствует требованиям изложенных в issue в случае если эта issue большая.

Промпт 004-pr-review-finish используем в случае если у нас были правки после имплементации или мы считаем что PR сложный и хотим добить качество кода об умную-долгую модель в режиме PR-review-fix.

## Реестр

| Prompt ID | Title | Status | Prompt status | Kind | Used for | Last updated |
| --- | --- | --- | --- | --- | --- | --- |
| [`PROMPT-001`](PROMPT-001-issue-requirements-review.md) | Issue Requirements Review | `draft` | `drafted` | `review` | Review feature docs against the source issue and memory-bank governance | 2026-05-19 |
| [`PROMPT-002`](PROMPT-002-feature-pack-review-improve.md) | Feature Pack Review Improve | `draft` | `drafted` | `review` | Run bounded review-improve cycles for feature packages | 2026-05-19 |
| [`PROMPT-003`](PROMPT-003-implement-and-test.md) | Implement And Test | `draft` | `drafted` | `coding` | Implement a coding task end-to-end through PR, review/fix and CI | 2026-05-19 |
| [`PROMPT-004`](PROMPT-004-pr-review-finish.md) | PR Review Finish | `draft` | `drafted` | `coding` | Finish an active branch into a ready PR with review-improve and CI gates | 2026-05-19 |
| [`PROMPT-005`](PROMPT-005-route-and-deliver-issue.md) | Route And Deliver Issue | `draft` | `drafted` | `agent` | Route a new issue and orchestrate delivery through its terminal flow gate | 2026-07-23 |

## Naming

- Формат файла: `PROMPT-XXX-short-name.md`
- Вместо `XXX` используй стабильный проектный идентификатор: номер задачи, внутренний prompt id или короткий монотонный номер
- Заголовок файла должен совпадать с `title` во frontmatter

## Template

- Используй шаблон [`../flows/templates/prompt/PROMPT-XXX.md`](../flows/templates/prompt/PROMPT-XXX.md)

---
title: Prompts Index
doc_kind: prompt
doc_function: index
purpose: Навигация по instantiated reusable prompt-документам проекта. Читать, чтобы найти существующий prompt или завести новый по governed-шаблону.
derived_from:
  - ../dna/governance.md
  - ../flows/templates/prompt/PROMPT-XXX.md
status: active
audience: humans_and_agents
---

# Prompts Index

Каталог `memory-bank/prompts/` хранит reusable prompt-документы проекта.

Prompt-документ нужен, когда prompt прошел путь от черновой человеческой формулировки до повторно используемой версии, которую нужно копировать, ревьюить и улучшать как артефакт memory-bank.

## Когда Заводить Prompt-Документ

- prompt будет использоваться повторно человеком или агентом;
- нужно сохранить исходную формулировку в `source_prompt`, а улучшенную версию держать отдельно;
- prompt становится частью рабочего процесса, ревью, research, extraction, coding или agent-инструкций.

## Когда Prompt-Документ Не Нужен

- prompt одноразовый и не должен жить дольше текущего диалога;
- это проектное правило, которое должно попасть в `engineering/`, `ops/`, `domain/` или `AGENTS.md`;
- это feature requirement, use case или ADR, а не исполняемая инструкция для модели.

## Порядок выполнения prompt-ов

Промпты указаны в порядке SDLC-процесса. Обычно используется `PROMPT-002` на старте, когда создаем `brief.md` и conditional `design.md`, после чего, как правило, делаем human gate; `PROMPT-003` используется, когда приступаем к имплементации.

Промпт 001-issue-requrements-review используем для того чтобы убедиться что feature-pack соответствует требованиям изложенных в issue в случае если эта issue большая.

Промпт 004-pr-review-finish используем в случае если у нас были правки после имплементации или мы считаем что  PR сложный и ходим добить качество кода об умную-долгую модель в режиме PR-review-fix.
## Реестр

| Prompt ID | Title | Status | Prompt status | Kind | Used for | Last updated |
| --- | --- | --- | --- | --- | --- | --- |
| [`PROMPT-001`](PROMPT-001-issue-requirements-review.md) | Issue Requirements Review | `draft` | `drafted` | `review` | Review feature docs against the source issue and memory-bank governance | 2026-05-19 |
| [`PROMPT-002`](PROMPT-002-feature-pack-review-improve.md) | Feature Pack Review Improve | `draft` | `drafted` | `review` | Run bounded review-improve cycles for feature packages | 2026-05-19 |
| [`PROMPT-003`](PROMPT-003-implement-and-test.md) | Implement And Test | `draft` | `drafted` | `coding` | Implement a coding task end-to-end through PR, review/fix and CI | 2026-05-19 |
| [`PROMPT-004`](PROMPT-004-pr-review-finish.md) | PR Review Finish | `draft` | `drafted` | `coding` | Finish an active branch into a ready PR with review-improve and CI gates | 2026-05-19 |

## Naming

- Формат файла: `PROMPT-XXX-short-name.md`
- Вместо `XXX` используй стабильный проектный идентификатор: номер задачи, внутренний prompt id или короткий монотонный номер
- Заголовок файла должен совпадать с `title` во frontmatter

## Template

- Используй шаблон [`../flows/templates/prompt/PROMPT-XXX.md`](../flows/templates/prompt/PROMPT-XXX.md)

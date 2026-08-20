---
title: GitHub Growth Plan
doc_kind: product
doc_function: roadmap
purpose: Фиксирует отдельный 30-дневный план роста видимости и проверяемого внешнего adoption для dapi/memory-bank.
derived_from:
  - ../dna/governance.md
  - marketing.md
  - metrics.md
status: active
audience: humans_and_agents
canonical_for:
  - github_growth_plan
---

# GitHub Growth Plan

## Outcome

За 30 дней проверить, может ли понятное позиционирование, пятиминутная активация и один доказательный кейс привести не только к росту GitHub stars с 15 до 100, но и к первым проверяемым внешним adoption signals.

Цель по stars — маркер visibility, а не доказательство ценности. Decision signal о продолжении канала — качество внешних установок, вопросов и интеграций.

## Starting Point

| Signal | Baseline on 2026-08-14 | Confidence |
| --- | ---: | --- |
| GitHub stars | 15 | verified via GitHub repository metadata |
| Forks | 4 | verified via GitHub repository metadata |
| Repository topics | none | verified via GitHub repository metadata |
| Homepage URL | absent | verified via GitHub repository metadata |
| External adoptions | unknown | no canonical evidence registry yet |

## Wave 1: Conversion Before Distribution

**Gate:** не запускать широкое распространение, пока новый читатель не может понять продукт и запустить dry-run за пять минут.

- [x] Сжать first screen README до problem, difference и working quick start.
- [x] Проверить короткую GIF-демонстрацию и убрать её: статичный сценарий на первом экране не объясняет реальное поведение CLI.
- [ ] Установить подготовленный GitHub social preview и homepage URL после появления стабильной страницы статьи.
- [x] Добавить topics: `coding-agents`, `context-engineering`, `agentic-coding`, `claude-code`, `codex`, `cursor`, `software-governance`, `developer-tools`, `documentation`, `ssot`.
- [x] Повторить quick start на чистом temporary repository: 131 файл установлен, `doctor` завершился с 0 errors, 1 ожидаемым предупреждением об отсутствующем CI gate.

## Wave 2: One Evidence-Backed Launch

**Gate:** статья прошла editorial review, а все публичные claims разделяют documented behavior, author experience и unknowns.

- Выпустить статью «Почему `AGENTS.md` недостаточно для долгоживущего AI-проекта».
- Подготовить один маленький public demo repository: clean repo → `init --dry-run` → `init` → адаптация одного product fact → `doctor`.
- Опубликовать один primary launch в Show HN и до двух адаптаций для релевантных сообществ.
- Отвечать на вопросы с конкретными примерами; не просить звёзду без полезного контекста.

## Wave 3: Turn Attention Into Adoption Evidence

- Завести evidence log с датой, public source, segment, attempted job и observed outcome.
- Из первых пяти содержательных вопросов выделить repeated friction и исправить README, adoption guide или CLI.
- Добавлять adopter logos и цитаты только с явным согласием и публичным источником.
- Отдельно считать repository views, quick-start attempts, adoption signals и stars; не выдавать одно за другое.

## Wave 4: Durable Discovery

- Подать репозиторий только в те awesome lists и tool catalogs, где он проходит inclusion criteria.
- Добавить короткие integration recipes для Codex, Claude Code и Cursor без дублирования core governance.
- Проверить долю organic discovery через 30 дней и решить: усилить канал, изменить message или остановить эксперимент.

## Stop Rules

- Остановить distribution, если quick start не воспроизводится или `doctor` находит новую release-blocking ошибку.
- Не расширять платное продвижение, пока organic launch не дал хотя бы пять qualified feedback signals.
- Не менять product scope ради единичного комментария; repeated pattern должен иметь evidence.

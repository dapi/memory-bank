---
title: Git Workflow
doc_kind: engineering
doc_function: convention
purpose: Git-конвенции проекта Merchantly — коммиты, ветки, PR, worktrees.
derived_from:
  - ../../CLAUDE.md
status: active
audience: humans_and_agents
---

# Git Workflow

**Main branch:** `master`

## Commits

- Present-tense, concise (`fix: normalize vendor slugs`)
- **ВСЕГДА** указывай `#<номер>` issue в commit message
- Финальный коммит: `Fixes #<номер>` в теле (auto-close)
- Номер issue неизвестен — **спроси пользователя**

## Pull Requests

- Перед PR: `make test` зелёный, RuboCop clean, schema/locale закоммичены
- PR title: короткий (до 70 символов), описание — в body

## Worktrees

- Создавай worktrees в `~/worktrees/` (НЕ `~/.worktrees/`)
- После `git worktree add` обязательно `./init.sh` (blocker)

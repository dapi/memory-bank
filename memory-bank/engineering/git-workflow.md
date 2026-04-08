---
title: Git Workflow
doc_kind: engineering
doc_function: convention
purpose: Шаблон git workflow документа. После копирования зафиксируй реальные branch names, commit rules и PR expectations проекта.
derived_from:
  - ../dna/governance.md
status: active
audience: humans_and_agents
---

# Git Workflow

## Default Branch

Явно укажи branch, который считается основным: например `main`, `master` или release branch.

## Commits

- Present-tense, concise (`fix: normalize cache key`)
- Если проект требует issue refs в commit message, зафиксируй это явно
- Если auto-close keywords допустимы, перечисли их
- Если squash merge обязателен или запрещен, укажи это здесь

## Pull Requests

- Перед PR должны быть зелёными canonical local checks проекта
- PR title должен быть коротким и предметным
- В body полезно фиксировать: что изменено, как проверено, какие риски или manual steps остаются

## Worktrees

Если проект использует worktrees, зафиксируй:

- где они создаются;
- требуется ли bootstrap script после `git worktree add`;
- какие каталоги считаются запрещенными для временной работы.

Если worktrees не используются, этот раздел можно удалить при адаптации.

---
title: Coding Style
doc_kind: engineering
doc_function: convention
purpose: Конвенции оформления кода проекта Merchantly.
derived_from:
  - ../../CLAUDE.md
status: active
audience: humans_and_agents
---

# Coding Style

## Ruby

- Two-space indentation, snake_case
- Double quotes только при интерполяции (`"Hello #{name}"` vs `'Hello'`)
- Rails autoloading: имена файлов = class/module names (`app/services/vendors/product_syncer.rb` → `Vendors::ProductSyncer`)
- RuboCop через Lefthook pre-commit hook (`lefthook install`), lint-clean перед push

## JavaScript

- Legacy JS (`app/assets/javascripts`) — **ES5-only**, без modern syntax (arrow functions, const/let, template literals и т.д.)
- Новые интерактивные фичи — **Stimulus** (см. [ADR-0001](adr/0001-frontend-interactivity-stimulus.md))

## Общее

- Не добавляй docstrings, комментарии или type annotations к коду, который не менял
- Комментарии только там, где логика не очевидна
- Avoid over-engineering: минимальная сложность для текущей задачи

---
title: Coding Style
doc_kind: engineering
doc_function: convention
purpose: Шаблон coding style документа. После копирования зафиксируй здесь реальные project-specific соглашения по коду и tooling.
derived_from:
  - ../dna/governance.md
status: active
audience: humans_and_agents
---

# Coding Style

## General Rules

- Имена файлов, модулей и каталогов должны соответствовать правилам основного языка проекта.
- Комментарии добавляются только там, где без них тяжело понять why или boundary condition.
- Предпочитай минимальную локальную сложность вместо преждевременных абстракций.
- Generated code, vendored code и миграции подчиняются отдельным правилам, если проект их вводит.

## Tooling Contract

Зафиксируй здесь canonical formatting/linting toolchain.

Пример:

- formatter: `prettier`, `ruff format`, `rubocop -A`, `gofmt`
- linter: `eslint`, `ruff`, `rubocop`, `golangci-lint`
- pre-commit hooks: optional, но если они canonical, это должно быть явно сказано

## Language-Specific Addendum

После адаптации добавь реальные правила для языков проекта.

Пример структуры:

- `Backend`: naming, error handling, module layout, typing policy
- `Frontend`: component boundaries, state management, styling rules
- `SQL / migrations`: naming, rollback expectations, data migration policy

## Change Discipline

- Не переписывай несвязанный код только ради единообразия, если задача этого не требует.
- При touch-up изменениях следуй существующему локальному стилю файла, если нет явного конфликта с canonical rule.
- Если проект находится в переходе между двумя стеками или стилями, зафиксируй migration rule явно, а не оставляй ее на догадки.

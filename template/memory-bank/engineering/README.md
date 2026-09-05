---
title: Engineering Documentation Index
doc_kind: engineering
doc_function: index
purpose: Навигация по engineering-level документации целевой системы.
derived_from:
  - ../dna/governance.md
status: active
audience: humans_and_agents
---

# Engineering Documentation Index

Каталог `memory-bank/engineering/` описывает инженерию **целевой системы**: как устроен и как пишется код продукта. Это project-adaptation слой — после копирования шаблона его нужно заполнить реальными правилами репозитория.

Правила самого delivery-процесса — границы автономии агента, глубина проверок и testing policy — живут в [`../flows/README.md`](../flows/README.md) и в этот каталог не входят: они generic и не зависят от стека проекта.

- [Engineering Architecture Patterns](architecture.md) — code/module boundaries, runtime patterns, concurrency, error handling и configuration ownership. Domain bounded contexts живут отдельно в [`../domain/context-map.md`](../domain/context-map.md).
- [Frontend Engineering](frontend.md) — UI surfaces, frontend stack, component boundaries, design system integration и i18n.
- [UI Design Guide](ui-design-guide/README.md) — project-level index для shared и surface-specific UI references. Адаптируй его под public site, admin, mobile или другие реальные UI surfaces проекта.
- [Testing Conventions](testing-conventions.md) — project-specific testing stack: framework, тестовые данные, CI jobs и размещение тестов. Локальные команды живут в [`../ops/development.md`](../ops/development.md). Что обязано быть покрыто, решает generic [`../flows/testing-policy.md`](../flows/testing-policy.md).
- [Coding Style](coding-style.md) — конвенции оформления кода, tooling и правила локальной сложности.
- [Git Workflow](git-workflow.md) — git-конвенции: commits, ветки, PR и optional worktrees.
- [ADR](../adr/README.md) — instantiated Architecture Decision Records проекта.

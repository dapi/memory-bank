---
title: Engineering Documentation Index
doc_kind: engineering
doc_function: index
purpose: Навигация по engineering-level документации шаблона.
derived_from:
  - ../dna/governance.md
status: active
audience: humans_and_agents
---

# Engineering Documentation Index

Каталог `memory-bank/engineering/` содержит инженерные правила, которые обычно нужно адаптировать под конкретный репозиторий после копирования шаблона.

- [Engineering Architecture Patterns](architecture.md) — code/module boundaries, runtime patterns, concurrency, error handling и configuration ownership. Domain bounded contexts живут отдельно в [`../domain/context-map.md`](../domain/context-map.md).
- [Frontend Engineering](frontend.md) — UI surfaces, frontend stack, component boundaries, design system integration и i18n.
- [Testing Policy](testing-policy.md) — правила тестирования, обязательные automated tests, sufficient coverage. Отвечает на вопрос: когда feature обязана иметь test cases и когда допустим manual-only verify.
- [Autonomy Boundaries](autonomy-boundaries.md) — границы автономии агента: автопилот, супервизия, эскалация. Отвечает на вопрос: что агент может делать сам, а где должен остановиться и спросить.
- [Coding Style](coding-style.md) — конвенции оформления кода, tooling и правила локальной сложности.
- [Git Workflow](git-workflow.md) — git-конвенции: commits, ветки, PR и optional worktrees.
- [ADR](../adr/README.md) — instantiated Architecture Decision Records проекта.

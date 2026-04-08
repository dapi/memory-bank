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

- [Testing Policy](testing-policy.md) — правила тестирования, обязательные automated tests, sufficient coverage. Отвечает на вопрос: когда feature обязана иметь test cases и когда допустим manual-only verify.
- [Autonomy Boundaries](autonomy-boundaries.md) — границы автономии агента: автопилот, супервизия, эскалация. Отвечает на вопрос: что агент может делать сам, а где должен остановиться и спросить.
- [Coding Style](coding-style.md) — конвенции оформления кода, tooling и правила локальной сложности.
- [Git Workflow](git-workflow.md) — git-конвенции: commits, ветки, PR и optional worktrees.
- [ADR](../adr/README.md) — instantiated Architecture Decision Records проекта.

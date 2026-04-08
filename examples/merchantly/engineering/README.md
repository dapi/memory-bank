---
title: Engineering Documentation Index
doc_kind: engineering
doc_function: index
purpose: Навигация по engineering-level документации проекта.
derived_from:
  - ../../../memory-bank/dna/governance.md
status: active
audience: humans_and_agents
---

# Engineering Documentation Index

Базовый scaffold для engineering-level документации живет в [`../../../memory-bank/engineering/README.md`](../../../memory-bank/engineering/README.md). Ниже собраны Merchantly-specific адаптации и дополнительные документы, которых нет в шаблоне.

- [Testing Policy](testing-policy.md) — правила тестирования, обязательные automated tests, sufficient coverage. Отвечает на вопрос: когда feature обязана иметь test cases и когда допустим manual-only verify.
- [Autonomy Boundaries](autonomy-boundaries.md) — границы автономии агента: автопилот, супервизия, эскалация. Отвечает на вопрос: что агент может делать сам, а где должен остановиться и спросить.
- [Coding Style](coding-style.md) — конвенции оформления кода: indentation, naming, RuboCop, JS/Stimulus.
- [Git Workflow](git-workflow.md) — git-конвенции: commits, branches, PR, worktrees.
- [ADR](adr/) — Architecture Decision Records. Фиксация архитектурных решений с контекстом и последствиями.
- [Agent Instructions](agent-instructions/) — инструкции для AI-агентов (поиск vendor по URL и т.д.).

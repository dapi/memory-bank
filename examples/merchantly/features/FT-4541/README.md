---
title: "FT-4541: Feature Package"
doc_kind: feature
doc_function: index
purpose: "Bootstrap-safe навигация по документации фичи. Читать, чтобы сначала перейти к canonical `feature.md`, а optional derived docs добавлять только после их появления."
derived_from:
  - ../../../../memory-bank/dna/governance.md
  - feature.md
status: active
audience: humans_and_agents
---

# FT-4541: Feature Package

## О разделе

Каталог feature package хранит canonical `feature.md`, а optional derived/external routes добавляются только после появления соответствующих документов. Сначала читай `feature.md`, затем расширяй routing по мере появления execution и decision artifacts.

## Аннотированный индекс

- [`feature.md`](feature.md)
  Читать, когда нужно: открыть instantiated canonical feature-документ сразу после bootstrap нового feature package.
  Отвечает на вопрос: где находятся scope, design, verify, blockers и canonical IDs для этой фичи.

- [`implementation-plan.md`](implementation-plan.md)
  Читать, когда нужно: начать или продолжить реализацию фичи, проверить текущий step, checkpoint или approval gate.
  Отвечает на вопрос: discovery context, порядок работ, test strategy, риски и stop conditions.

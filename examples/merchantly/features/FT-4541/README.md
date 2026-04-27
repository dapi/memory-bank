---
title: "FT-4541: Feature Package"
doc_kind: feature
doc_function: index
purpose: "Навигация по документации FT-4541. Читать, чтобы пройти по problem space (`feature.md`), затем по solution space (`solution.md`) и только потом по execution-документу."
derived_from:
  - ../../../../memory-bank/dna/governance.md
  - feature.md
status: active
audience: humans_and_agents
---

# FT-4541: Feature Package

## О разделе

Каталог feature package хранит canonical `feature.md`, canonical `solution.md` и derived `implementation-plan.md`. Читать их нужно по порядку ownership: сначала problem space, затем selected design и to-be C4 architecture model, затем execution sequencing.

## Аннотированный индекс

- [`feature.md`](feature.md)
  Читать, когда нужно: открыть instantiated canonical feature-документ сразу после bootstrap нового feature package.
  Отвечает на вопрос: где находятся problem space, canonical verify contract и stable IDs для этой фичи.

- [`solution.md`](solution.md)
  Читать, когда нужно: понять выбранный design, to-be C4 architecture model, accepted local decisions, contracts, failure modes и rollout/backout semantics.
  Отвечает на вопрос: как именно реализуется feature без смешения solution space с `feature.md`.

- [`implementation-plan.md`](implementation-plan.md)
  Читать, когда нужно: начать или продолжить реализацию фичи, проверить текущий step, checkpoint или approval gate.
  Отвечает на вопрос: discovery context, порядок работ, test strategy, риски и stop conditions после `Solution Ready`.

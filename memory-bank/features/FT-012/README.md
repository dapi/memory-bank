---
title: "FT-012: Compact Task Flow"
doc_kind: feature
doc_function: index
purpose: "Навигация по документации feature package для compact task-flow bugfix/refactor/chore. Читать, чтобы перейти к canonical brief, design, implementation plan и FPF decision log."
derived_from:
  - ../../dna/governance.md
  - brief.md
status: active
audience: humans_and_agents
---

# FT-012: Compact Task Flow

## О разделе

Каталог фиксирует feature package для GitHub issue [#12](https://github.com/dapi/memory-bank/issues/12): добавить generic compact task-flow для bugfix/refactor/chore работ без превращения каждой маленькой задачи в full feature package.

## Аннотированный индекс

- [`brief.md`](brief.md)
  Читать, когда нужно: проверить problem space, scope, non-scope, acceptance scenarios, checks и evidence contract.
  Отвечает на вопрос: что должна доставить feature и как это будет принято.

- [`design.md`](design.md)
  Читать, когда нужно: проверить selected design для семейства `task-flow`, `bugfix-flow`, `refactor-flow`, templates и promotion boundaries.
  Отвечает на вопрос: как compact task layer отделяется от feature/epic flow без скрытия high-risk или contract work.

- [`implementation-plan.md`](implementation-plan.md)
  Читать, когда нужно: выполнить scoped docs changes, обновить индексы и собрать verification evidence.
  Отвечает на вопрос: какие шаги, touchpoints, checks и stop conditions ведут feature к приемке.

- [`decision-log.md`](decision-log.md)
  Читать, когда нужно: проверить FPF-backed решения, закрытые в ходе review-improve.
  Отвечает на вопрос: какие существенные неоднозначности закрыты, на каких фактах и с какими последствиями.

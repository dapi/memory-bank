---
title: P1-BUG Context Priming
doc_kind: governance
doc_function: canonical
purpose: Route-specific праймеринг перед Entry Gate Bug Fix Flow.
derived_from:
  - context-priming.md
  - ../bug-fix.md
canonical_for:
  - p1_bug_fix_priming
  - bug_fix_priming_input_classes
status: active
audience: humans_and_agents
---

# P1-BUG Context Priming

До Bug Fix Entry Gate прочитай only exact inputs из current process/task
manifest. Он обязан назвать:

- bug report или source trigger;
- canonical owner already accepted expected behavior;
- affected implementation path и minimum reproduction inputs/environment;
- ближайшую existing test surface либо evidence-backed отсутствие покрытия.

В bug report или linked task зафиксируй expected/actual behavior, reproduction
evidence, sources и open uncertainty. Не выбирай новый product behavior и не
начинай refactoring. Если expected behavior нельзя подтвердить, перейди в
Human Routing.

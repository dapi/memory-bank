---
title: P1-REF Context Priming
doc_kind: governance
doc_function: canonical
purpose: Route-specific праймеринг перед Entry Gate Refactoring Flow.
derived_from:
  - context-priming.md
  - ../refactoring.md
canonical_for:
  - p1_refactoring_priming
  - refactoring_priming_input_classes
status: active
audience: humans_and_agents
---

# P1-REF Context Priming

До Entry Gate прочитай only exact inputs из task/process manifest: observable
behavior и contracts, которые должны сохраниться, baseline/characterization
coverage, structural change surface и checkpoint constraints.

Исходная task фиксирует preservation boundary, baseline и material unknowns.
Intentional behavior/contract change не маскируй как refactoring: верни его в
Task Routing.

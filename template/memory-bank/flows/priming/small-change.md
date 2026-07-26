---
title: P1-SMALL Context Priming
doc_kind: governance
doc_function: canonical
purpose: Route-specific праймеринг перед Entry Gate Small Change Flow.
derived_from:
  - context-priming.md
  - ../small-change.md
canonical_for:
  - p1_small_change_priming
  - small_change_priming_input_classes
status: active
audience: humans_and_agents
---

# P1-SMALL Context Priming

До Entry Gate прочитай only exact inputs из task/process manifest: declared
intent/scope/acceptance, реально существующий reference pattern, local change
surface и known test/verify surface.

Routing record ссылается на проверенный pattern и concrete verify actions. Если
pattern не подходит, surface не локален или нужны design/plan/contract
decisions, останови fast path и повтори Task Routing.

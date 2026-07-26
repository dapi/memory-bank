---
title: P1-FEAT Context Priming
doc_kind: governance
doc_function: canonical
purpose: Route-specific праймеринг перед bootstrap Feature package.
derived_from:
  - context-priming.md
  - ../feature.md
canonical_for:
  - p1_feature_priming
  - feature_priming_input_classes
status: active
audience: humans_and_agents
---

# P1-FEAT Context Priming

До bootstrap прочитай only exact inputs из task/process manifest: relevant
upstream PRD/epic/use case/ADR, applicable domain and engineering rules,
affected contracts и bounded current repository surfaces.

Draft `features/FT-XXX/brief.md` фиксирует problem-space facts, assumptions,
constraints и unresolved decisions, но не selected solution. Перед Plan Ready
P2 всё ещё требует immutable revision, `GRND-*` evidence и task-specific
Implementation Priming manifest.

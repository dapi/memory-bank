---
title: Domain Documentation Index
doc_kind: domain
doc_function: index
purpose: Навигация по domain-level документации Merchantly. Читать для понимания бизнес-контекста и архитектурных паттернов.
derived_from:
  - ../dna/governance.md
status: active
audience: humans_and_agents
---

# Domain Documentation Index

- [Project Problem Statement](problem.md) — общий продуктовый контекст Merchantly, top-level workflows и ограничения. Читать перед PRD и feature-спеками.
- [Architecture Patterns](architecture.md) — изоляция контекстов (Admin vs Operator), блокировки (advisory locks, SolidQueue), Bugsnag в jobs. Читать при написании кода, затрагивающего эти области.
- [Frontend](frontend.md) — public frontend (React), operator frontend (Stimulus/jQuery), переводы. Читать при работе с UI.
- [Design Guide](design-guide/) — UI компоненты операторской панели: кнопки, формы, таблицы, контейнеры, метки. Читать при создании/изменении operator views.

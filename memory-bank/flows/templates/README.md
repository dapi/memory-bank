---
title: Templates Index
doc_kind: governance
doc_function: index
purpose: Навигация по эталонным шаблонам документации проекта. Читать, чтобы завести новую фичу, ADR или execution-документ без изобретения новой структуры.
derived_from:
  - ../../dna/governance.md
  - prd/PRD-XXX.md
  - use-case/UC-XXX.md
  - feature/README.md
  - feature/implementation-plan.md
  - feature/short.md
  - feature/large.md
  - adr/ADR-XXX.md
status: active
audience: humans_and_agents
---

# Templates Index

Каталог `memory-bank/flows/templates/` хранит эталонные шаблоны документации проекта. Все шаблоны живут как governed wrapper-документы с `doc_function: template`: у wrapper-а есть собственные purpose, а frontmatter и body инстанцируемого документа — внутри embedded template contract.

- [PRD-XXX: Product Initiative Name](prd/PRD-XXX.md) — компактный Product Requirements Document для инициативы, которая еще не разложена на один конкретный feature slice.
- [UC-XXX: Use Case Name](use-case/UC-XXX.md) — канонический use case для устойчивого пользовательского или операционного сценария.
- [FT-XXX Feature README Template](feature/README.md) — шаблон README для feature-каталога. Отвечает на вопрос: как оформить feature-level index.
- [FT-XXX: Feature Template - Short](feature/short.md) — минимальный canonical feature для небольшой фичи. Отвечает на вопрос: как выглядит short feature-документ.
- [FT-XXX: Feature Template - Large](feature/large.md) — canonical feature с assumptions, blockers, contracts, verify-слоем. Отвечает на вопрос: как выглядит large feature-документ.
- [FT-XXX: Implementation Plan](feature/implementation-plan.md) — шаблон derived execution-плана. Отвечает на вопрос: как оформить sequencing и checkpoints.
- [ADR-XXX: Short Decision Name](adr/ADR-XXX.md) — шаблон ADR. Отвечает на вопрос: как зафиксировать архитектурное решение.

---
title: Templates Index
doc_kind: governance
doc_function: index
purpose: Навигация по эталонным шаблонам документации проекта. Читать, чтобы завести PRD, use case, epic, фичу, ADR, prompt или execution-документ без изобретения новой структуры.
derived_from:
  - ../../dna/governance.md
  - prd/PRD-XXX.md
  - use-case/UC-XXX.md
  - epic/README.md
  - epic/charter.md
  - epic/roadmap.md
  - epic/decision-log.md
  - epic/subissues.md
  - epic/risks.md
  - feature/README.md
  - feature/brief.md
  - feature/design.md
  - feature/implementation-plan.md
  - feature/support/runtime-surfaces.md
  - feature/support/ui-reference.md
  - feature/support/use-cases.md
  - adr/ADR-XXX.md
  - prompt/PROMPT-XXX.md
  - process/README.md
  - process/process-card.md
  - process/session-handoff.md
  - process/lifecycle-protocol.md
status: active
audience: humans_and_agents
---

# Templates Index

Каталог `memory-bank/flows/templates/` хранит эталонные шаблоны документации проекта. Все шаблоны живут как governed wrapper-документы с `doc_function: template`: у wrapper-а есть собственные purpose, а frontmatter и body инстанцируемого документа — внутри embedded template contract.

- [PRD-XXX: Product Initiative Name](prd/PRD-XXX.md) — компактный Product Requirements Document для инициативы, которая еще не разложена на один конкретный feature slice.
- [UC-XXX: Use Case Name](use-case/UC-XXX.md) — канонический use case для устойчивого пользовательского или операционного сценария.
- [Epic Templates](epic/README.md) — индекс шаблонов `EP-XXX` package.
- [EP-XXX: Charter Template](epic/charter.md) — intent, scope, source/evidence and stakeholder channels.
- [EP-XXX: Roadmap Template](epic/roadmap.md) — waves, dependencies, gates and stop rules.
- [EP-XXX: Decision Log Template](epic/decision-log.md) — local epic decisions that do not require global ADR.
- [EP-XXX: Subissues Template](epic/subissues.md) — candidate/accepted delivery subissue registry.
- [EP-XXX: Risks Template](epic/risks.md) — epic-level risk register.
- [FT-XXX Feature README Template](feature/README.md) — шаблон README для feature-каталога. Отвечает на вопрос: как оформить feature-level index.
- [FT-XXX: Brief Template](feature/brief.md) — canonical problem-space template для новых feature packages. Отвечает на вопрос: как зафиксировать intent, scope и verify contract без solution/execution деталей.
- [FT-XXX: Design Template](feature/design.md) — canonical solution-space template для feature package. Отвечает на вопрос: как зафиксировать selected design, rationale, contracts, failure modes и design-pack routing.
- [FT-XXX: Implementation Plan](feature/implementation-plan.md) — шаблон derived execution-плана. Отвечает на вопрос: как оформить sequencing и checkpoints после готовности upstream owners.
- [FT-XXX: Runtime Surfaces Template](feature/support/runtime-surfaces.md) — optional support template для current runtime inventory, semantic mapping, context matrix и resolution tables.
- [FT-XXX: UI Reference Template](feature/support/ui-reference.md) — optional support template для interface changes, screen map, interaction states и mockups.
- [FT-XXX: Feature Use Cases Template](feature/support/use-cases.md) — optional support template для derived use cases, test case candidates и `FUC -> REQ -> CHK` review mapping.
- [ADR-XXX: Short Decision Name](adr/ADR-XXX.md) — шаблон ADR. Отвечает на вопрос: как зафиксировать архитектурное решение.
- [PROMPT-XXX: Reusable Prompt Name](prompt/PROMPT-XXX.md) — шаблон reusable prompt-документа. Отвечает на вопрос: как сохранить исходную формулировку в frontmatter и улучшенный prompt в copyable body-блоке.
- [PROC-XXX: Process Documentation Index](process/README.md) — шаблон индекса процесс-документов. Отвечает на вопрос: как собрать routing-layer для reusable process cards, session handoff и lifecycle protocol.
- [PROC-XXX: Compact Process Card](process/process-card.md) — шаблон короткого reusable workflow. Отвечает на вопрос: как зафиксировать процесс с одним trigger, шагами и exit criteria.
- [PROC-XXX: Session Handoff](process/session-handoff.md) — шаблон передачи состояния между сессиями. Отвечает на вопрос: как продолжить процесс без потери assumptions, risks и next checks.
- [PROC-XXX: Lifecycle Protocol](process/lifecycle-protocol.md) — шаблон полного lifecycle protocol. Отвечает на вопрос: как вести multi-phase process с gates, verification и rollback.

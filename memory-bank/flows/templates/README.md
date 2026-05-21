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
  - feature/implementation-plan.md
  - feature/short.md
  - feature/large.md
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
- [FT-XXX: Feature Template - Short](feature/short.md) — минимальный canonical feature для небольшой фичи. Отвечает на вопрос: как выглядит short feature-документ.
- [FT-XXX: Feature Template - Large](feature/large.md) — canonical feature с assumptions, blockers, contracts, verify-слоем. Отвечает на вопрос: как выглядит large feature-документ.
- [FT-XXX: Implementation Plan](feature/implementation-plan.md) — шаблон derived execution-плана. Отвечает на вопрос: как оформить sequencing и checkpoints.
- [ADR-XXX: Short Decision Name](adr/ADR-XXX.md) — шаблон ADR. Отвечает на вопрос: как зафиксировать архитектурное решение.
- [PROMPT-XXX: Reusable Prompt Name](prompt/PROMPT-XXX.md) — шаблон reusable prompt-документа. Отвечает на вопрос: как сохранить исходную формулировку в frontmatter и улучшенный prompt в copyable body-блоке.
- [PROC-XXX: Process Documentation Index](process/README.md) — шаблон индекса процесс-документов. Отвечает на вопрос: как собрать routing-layer для reusable process cards, session handoff и lifecycle protocol.
- [PROC-XXX: Process Card](process/process-card.md) — компактный шаблон повторяемого процесса.
- [PROC-XXX: Session Handoff](process/session-handoff.md) — шаблон передачи состояния между сессиями.
- [PROC-XXX: Lifecycle Protocol](process/lifecycle-protocol.md) — шаблон длинного процесса с фазами, gates и verification.

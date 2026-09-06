---
title: Templates Index
doc_kind: governance
doc_function: index
purpose: Навигация по эталонным шаблонам документации проекта. Читать, чтобы завести PRD, use case, epic, фичу, ADR, prompt или execution-документ без изобретения новой структуры.
derived_from:
  - ../../dna/governance.md
  - prd/PRD-XXX.md
  - use-case/UC-XXX.md
  - research/README.md
  - research/package-README.md
  - research/brief.md
  - research/plan.md
  - research/evidence.md
  - research/synthesis.md
  - research/decision.md
  - epic/README.md
  - epic/package-README.md
  - epic/brief.md
  - epic/charter.md
  - epic/roadmap.md
  - epic/decision-log.md
  - epic/subissues.md
  - epic/risks.md
  - feature/README.md
  - feature/brief.md
  - feature/design.md
  - feature/api-contract.md
  - feature/implementation-plan.md
  - feature/support/runtime-surfaces.md
  - feature/support/sequence-diagram.md
  - feature/support/ui-reference.md
  - feature/support/use-cases.md
  - adr/ADR-XXX.md
status: active
audience: humans_and_agents
---

# Templates Index

Каталог `memory-bank/flows/templates/` хранит эталонные шаблоны документации проекта. Все шаблоны живут как governed wrapper-документы с `doc_function: template`: у wrapper-а есть собственные purpose, а frontmatter и body инстанцируемого документа — внутри embedded template contract.

- [PRD-XXX: Product Initiative Name](prd/PRD-XXX.md) — компактный Product Requirements Document для инициативы, которая еще не разложена на один конкретный feature slice.
- [UC-XXX: Use Case Name](use-case/UC-XXX.md) — канонический use case для устойчивого пользовательского или операционного сценария; selection и lifecycle определяет [Use Case Flow](../use-case.md).
- [Research Templates](research/README.md) — индекс шаблонов `R-XXX` package для market, product и technical research.
- [R-XXX Package README Template](research/package-README.md) — routing index research package; lifecycle state не дублируется здесь.
- [R-XXX: Research Brief Template](research/brief.md) — canonical decision question, hypotheses, boundaries, stopping condition и единственный lifecycle owner (`research_status`).
- [R-XXX: Research Plan Template](research/plan.md) — conditional method, sampling/source strategy и collection controls.
- [R-XXX: Evidence Log Template](research/evidence.md) — provenance-preserving log источников и observations.
- [R-XXX: Research Synthesis Template](research/synthesis.md) — findings, confidence, limitations и disconfirming evidence.
- [R-XXX: Research Decision Template](research/decision.md) — decision rationale, recommendation и promotion/handoff map; terminal state остаётся в `brief.md`.
- [Epic Templates](epic/README.md) — индекс шаблонов `EP-XXX` package.
- [EP-XXX Package README Template](epic/package-README.md) — routing index и lifecycle stage owner для epic package, включая intake-only состояние.
- [EP-XXX: Epic Proposal Template](epic/brief.md) — обязательный при Epic Intake brief с proposal disposition и promotion contract; при прямом Bootstrap Epic не создаётся.
- [EP-XXX: Charter Template](epic/charter.md) — intent, scope, source/evidence and stakeholder channels.
- [EP-XXX: Roadmap Template](epic/roadmap.md) — waves, dependencies, gates and stop rules.
- [EP-XXX: Decision Log Template](epic/decision-log.md) — local epic decisions that do not require global ADR.
- [EP-XXX: Subissues Template](epic/subissues.md) — candidate/accepted delivery subissue registry.
- [EP-XXX: Risks Template](epic/risks.md) — epic-level risk register.
- [FT-XXX Feature README Template](feature/README.md) — шаблон README для feature-каталога. Отвечает на вопрос: как оформить feature-level index.
- [FT-XXX: Brief Template](feature/brief.md) — canonical problem-space template для новых feature packages. Отвечает на вопрос: как зафиксировать intent, scope и verify contract без solution/execution деталей.
- [FT-XXX: Design Template](feature/design.md) — canonical solution-space template для feature package. Отвечает на вопрос: как зафиксировать selected design, architecture coverage, contracts, design verification и design-pack routing.
- [FT-XXX: Interaction Contract Template](feature/api-contract.md) — optional canonical design-pack template для подробной семантики API/event/queue/callback/file/store/cache/auth/locking/runtime-config connector; schema/encoding фиксируются как format, а provider — как party/role.
- [FT-XXX: Implementation Plan](feature/implementation-plan.md) — шаблон derived execution-плана. Отвечает на вопрос: как оформить sequencing и checkpoints после готовности upstream owners.
- [FT-XXX: Runtime Surfaces Template](feature/support/runtime-surfaces.md) — optional support template для current runtime inventory, semantic mapping, context matrix и resolution tables.
- [FT-XXX: Sequence Diagram Template](feature/support/sequence-diagram.md) — optional reference template для temporal / async interactions, retries, timeouts и failure branches.
- [FT-XXX: UI Reference Template](feature/support/ui-reference.md) — optional support template для interface changes, screen map, interaction states и mockups.
- [FT-XXX: Feature Use Cases Template](feature/support/use-cases.md) — optional support template для derived use cases, BDD example map, test candidates и `FUC → SC/NEG → REQ → CHK` review mapping без нового acceptance owner.
- [ADR-XXX: Short Decision Name](adr/ADR-XXX.md) — шаблон ADR. Отвечает на вопрос: как зафиксировать архитектурное решение.
- [PROMPT-XXX: Reusable Prompt Name](prompt/PROMPT-XXX.md) — шаблон reusable prompt-документа. Отвечает на вопрос: как сохранить исходную формулировку в frontmatter и улучшенный prompt в copyable body-блоке.

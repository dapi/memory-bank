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
  - process/README.md
  - process/process-card.md
  - process/session-handoff.md
  - process/lifecycle-protocol.md
status: active
audience: humans_and_agents
---

# Templates Index

Каталог содержит два вида governed wrapper-документов с `doc_function: template`.
ADR, feature brief, PRD, use case, research brief и epic charter представлены
процессными фрагментами к базовым шаблонам из `memory-bank/templates/`: они
перечисляют только добавляемые поля и секции, после заполнения требуется явное
adoption. Остальные файлы — самостоятельные шаблоны вспомогательных документов
с собственным embedded frontmatter/body. Frontmatter самого wrapper описывает
эталон и не копируется в проектный документ.

- [PRD flow fragment](prd/PRD-XXX.md) — process validation к базовому PRD.
- [Use-case flow fragment](use-case/UC-XXX.md) — verification к базовому use case.
- [Research Templates](research/README.md) — индекс шаблонов `R-XXX` package для market, product и technical research.
- [R-XXX Package README Template](research/package-README.md) — routing index research package; lifecycle state не дублируется здесь.
- [Research brief flow fragment](research/brief.md) — lifecycle status и ссылки на terminal artifacts исследования.
- [R-XXX: Research Plan Template](research/plan.md) — conditional method, sampling/source strategy и collection controls.
- [R-XXX: Evidence Log Template](research/evidence.md) — provenance-preserving log источников и observations.
- [R-XXX: Research Synthesis Template](research/synthesis.md) — findings, confidence, limitations и disconfirming evidence.
- [R-XXX: Research Decision Template](research/decision.md) — decision rationale, recommendation и promotion/handoff map; terminal state остаётся в `brief.md`.
- [Epic Templates](epic/README.md) — индекс шаблонов `EP-XXX` package.
- [EP-XXX Package README Template](epic/package-README.md) — routing index и lifecycle stage owner для epic package, включая intake-only состояние.
- [EP-XXX: Epic Proposal Template](epic/brief.md) — обязательный при Epic Intake brief с proposal disposition и promotion contract; при прямом Bootstrap Epic не создаётся.
- [Epic charter flow fragment](epic/charter.md) — ссылки на roadmap и risk owner инициативы.
- [EP-XXX: Roadmap Template](epic/roadmap.md) — waves, dependencies, gates and stop rules.
- [EP-XXX: Decision Log Template](epic/decision-log.md) — local epic decisions that do not require global ADR.
- [EP-XXX: Subissues Template](epic/subissues.md) — candidate/accepted delivery subissue registry.
- [EP-XXX: Risks Template](epic/risks.md) — epic-level risk register.
- [FT-XXX Feature README Template](feature/README.md) — шаблон README для feature-каталога. Отвечает на вопрос: как оформить feature-level index.
- [Feature brief flow fragment](feature/brief.md) — design/validation/verify requirements к базовому brief.
- [FT-XXX: Design Template](feature/design.md) — canonical solution-space template для feature package. Отвечает на вопрос: как зафиксировать selected design, architecture coverage, contracts, design verification и design-pack routing.
- [FT-XXX: Interaction Contract Template](feature/api-contract.md) — optional canonical design-pack template для подробной семантики API/event/queue/callback/file/store/cache/auth/locking/runtime-config connector; schema/encoding фиксируются как format, а provider — как party/role.
- [FT-XXX: Implementation Plan](feature/implementation-plan.md) — шаблон derived execution-плана. Отвечает на вопрос: как оформить sequencing и checkpoints после готовности upstream owners.
- [FT-XXX: Runtime Surfaces Template](feature/support/runtime-surfaces.md) — optional support template для current runtime inventory, semantic mapping, context matrix и resolution tables.
- [FT-XXX: Sequence Diagram Template](feature/support/sequence-diagram.md) — optional reference template для temporal / async interactions, retries, timeouts и failure branches.
- [FT-XXX: UI Reference Template](feature/support/ui-reference.md) — optional support template для interface changes, screen map, interaction states и mockups.
- [FT-XXX: Feature Use Cases Template](feature/support/use-cases.md) — optional support template для derived use cases, BDD example map, test candidates и `FUC → SC/NEG → REQ → CHK` review mapping без нового acceptance owner.
- [ADR flow fragment](adr/ADR-XXX.md) — review к базовому ADR.
- [PROMPT-XXX: Reusable Prompt Name](prompt/PROMPT-XXX.md) — шаблон reusable prompt-документа. Отвечает на вопрос: как сохранить исходную формулировку в frontmatter и улучшенный prompt в copyable body-блоке.
- [PROC-XXX: Process Documentation Index](process/README.md) — шаблон индекса процесс-документов. Отвечает на вопрос: как собрать routing-layer для reusable process cards, session handoff и lifecycle protocol.
- [PROC-XXX: Compact Process Card](process/process-card.md) — шаблон короткого reusable workflow. Отвечает на вопрос: как зафиксировать процесс с одним trigger, шагами и exit criteria.
- [PROC-XXX: Session Handoff](process/session-handoff.md) — шаблон передачи состояния между сессиями. Отвечает на вопрос: как продолжить процесс без потери assumptions, risks и next checks.
- [PROC-XXX: Lifecycle Protocol](process/lifecycle-protocol.md) — шаблон полного lifecycle protocol. Отвечает на вопрос: как вести multi-phase process с gates, verification и rollback.

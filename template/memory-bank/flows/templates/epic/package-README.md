---
title: "EP-XXX Package README Template"
doc_kind: governance
doc_function: template
purpose: "Wrapper-шаблон routing index для epic package, включая intake-only состояние до появления canonical charter."
derived_from:
  - ../../epic.md
status: active
audience: humans_and_agents
template_for: epic
template_target_path: ../../../epics/EP-XXX/README.md
---

# EP-XXX Package README Template

Создай этот index первым для любого epic package. `epic_stage` всегда отражает текущую стадию: `epic_intake`, `proposal_ready`, `parked`, `rerouted`, `rejected`, `draft`, `epic_ready`, `roadmap_ready`, `execution`, `done` или `cancelled`.

Если intake пропущен, начни с `epic_stage: draft`, добавь `charter.md` и не включай несуществующий `brief.md` в `derived_from` или индекс.

## Instantiated Frontmatter

```yaml
---
title: "EP-XXX: <Epic Name>"
doc_kind: epic
doc_function: index
purpose: "Навигация и текущая lifecycle stage для EP-XXX."
derived_from:
  - ../../flows/epic.md
  - brief.md
status: active
epic_stage: epic_intake
audience: humans_and_agents
---
```

## Instantiated Body

```markdown
# EP-XXX: <Epic Name>

## Current Stage

- Stage: `epic_intake`
- Owner: `<proposal or epic owner>`
- Source / trigger: `<issue, request, PRD or evidence URL>`
- Next gate: `Epic Intake -> Proposal Ready`

## Annotated Index

- [Epic Proposal](brief.md) — early proposal facts, open questions and disposition.

Добавляй `charter.md`, `roadmap.md`, `subissues.md`, `risks.md`, optional `decision-log.md` и knowledge artifacts только когда они реально созданы. Для каждой ссылки кратко укажи, какими facts владеет документ.

## Handoff

До `Roadmap Ready -> Execution` не создавай delivery `FT-*` packages из этого epic. Следуй gates в `memory-bank/flows/epic.md`.
```

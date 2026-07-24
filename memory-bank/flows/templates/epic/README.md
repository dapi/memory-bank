---
title: Epic Templates Index
doc_kind: governance
doc_function: template
purpose: "Wrapper-шаблоны для `memory-bank/epics/EP-XXX/` packages: package index, Epic Intake proposal, charter, roadmap, decision log, subissues and risks."
derived_from:
  - ../../epic.md
status: active
audience: humans_and_agents
---

# Epic Templates Index

Используй эти templates при создании нового `memory-bank/epics/EP-XXX/`. Начни с package README; если нужен Epic Intake, обязательно добавь brief. Если proposal facts уже достаточны, пропусти Intake и сразу создай charter без brief.

- [`package-README.md`](package-README.md) - routing index and `epic_stage` owner for instantiated `EP-XXX/README.md`.
- [`brief.md`](brief.md) - required Epic Intake proposal with disposition and promotion contract; omit only when Intake is skipped.
- [`charter.md`](charter.md) - intent, scope, source/evidence and stakeholder channels.
- [`roadmap.md`](roadmap.md) - waves, dependencies, gates and stop rules.
- [`decision-log.md`](decision-log.md) - local epic decisions that do not require global ADR.
- [`subissues.md`](subissues.md) - candidate/accepted delivery subissue registry.
- [`risks.md`](risks.md) - epic-level risk register.

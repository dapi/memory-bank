---
title: "FT-017: Decision Log"
doc_kind: feature
doc_function: log
purpose: "Feature-local decision log for issue 17. Records FPF-backed decisions that close open questions for the optional UI design guide pattern without requiring a global ADR."
derived_from:
  - brief.md
status: active
audience: humans_and_agents
---

# FT-017: Decision Log

## Decision Criteria

The following feature-local criteria are derived from issue 17 and the current memory-bank documents:

- `DC-01` The solution must help downstream projects document an existing UI kit.
- `DC-02` The generic template must stay portable and framework-agnostic.
- `DC-03` The owner boundary with [`../../engineering/frontend.md`](../../engineering/frontend.md) must stay clear.
- `DC-04` Domain-specific and project-specific UI content must not leak back into the generic template.

FPF application:

- Bounded Context: local source-project rules must stay inside the downstream project context; generic memory-bank may only publish reusable structure and translation prompts.
- Strict Distinction: `frontend.md` remains an engineering contract, while the optional UI design guide is a reference carrier for concrete UI kit facts in a downstream instance.
- Reasoning Cycle: abductive candidates were compared against documented constraints, then selected only when their consequences satisfied `DC-*`.

## Decisions

| Decision ID | Status | Question | Decision | Facts used | FPF rationale | Consequences |
| --- | --- | --- | --- | --- | --- | --- |
| `FDL-001` | accepted | Where should the optional UI design guide live? | Use `memory-bank/engineering/ui-design-guide/README.md` as the optional engineering-layer destination/template, not a domain-layer owner. | Issue 17 asks to link with `engineering/frontend.md`; `frontend.md` owns UI surfaces, frontend stack, component boundaries and design system integration; `domain/README.md` explicitly excludes UI design system ownership. | Bounded Context keeps UI engineering facts in the engineering semantic frame. Strict Distinction prevents treating UI design system guidance as domain truth. | `design.md` selects an engineering optional destination and rejects `domain/design-guide` as the generic owner. |
| `FDL-002` | accepted | Should the feature only expand `frontend.md` or add a separate optional guide pattern? | Add a separate optional guide pattern and link to it from `frontend.md` / indexes. | Issue 17 asks for a quick source of truth over components, helper APIs, screenshots and local patterns; `frontend.md` currently describes broader frontend engineering contracts. | Strict Distinction separates the engineering contract from concrete UI kit reference material. The reasoning cycle predicts that putting all UI kit details into `frontend.md` would blur owner roles. | `frontend.md` remains canonical for frontend engineering; the optional guide carries detailed UI kit sections. |
| `FDL-003` | accepted | Does this feature need C4 or ADR? | C4 artifact and ADR are not required. | The feature is documentation-template work; no API, runtime, deployment, data-flow, integration, security or reusable architecture boundary changes are in scope. | Feature-flow C4 triggers are absent. Strict Distinction keeps documentation structure decisions feature-local unless they alter architecture across features. | `design.md` records `C4-00: not required` and keeps decisions in this log / `design.md` rather than ADR. |
| `FDL-004` | accepted | Can source-project examples be treated as generic defaults? | No. They are evidence of need only; they are not reusable generic content. | Issue 17 names `brandymint/merchantly` source files and explicitly says not to transfer Bootstrap/Inspinia/HAML-specific rules as generic defaults. | Evidence Graph discipline separates a source carrier from the claim it supports. The source paths support the existence of the documentation need, not the portability of their local rules. | `brief.md` uses `ASM-01`, `NS-01`, `NS-02` and `REQ-04`; implementation must use placeholders and prompts, not copied downstream facts. |

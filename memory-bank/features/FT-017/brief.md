---
title: "FT-017: Optional UI Design Guide Pattern"
doc_kind: feature
doc_function: canonical
purpose: "Canonical brief для delivery-единицы issue 17. Фиксирует problem space, scope и verify contract для optional UI design guide pattern без смешения с solution space или execution plan."
derived_from:
  - ../../flows/feature-flow.md
  - ../../engineering/frontend.md
status: active
delivery_status: done
audience: humans_and_agents
source_issue:
  id: 17
  url: "https://github.com/dapi/memory-bank/issues/17"
must_not_define:
  - implementation_sequence
  - solution_space
---

# FT-017: Optional UI Design Guide Pattern

## What

### Problem

В проектах с существующим UI агентам нужен быстрый источник истины по компонентам, helper APIs, screenshots и локальным UI-паттернам. Сейчас generic memory-bank содержит [`engineering/frontend.md`](../../engineering/frontend.md) как engineering contract для UI surfaces, frontend stack, component boundaries, design system integration и i18n, но не дает отдельного optional destination/template для документирования уже существующего UI kit.

Issue 17 приводит downstream source examples из `brandymint/merchantly`, но требует не переносить framework-specific правила в generic memory-bank.

### Outcome

| Metric ID | Metric | Baseline | Target | Measurement method |
| --- | --- | --- | --- | --- |
| `MET-01` | Discoverability of optional UI guide pattern | Downstream project can read `frontend.md`, but no dedicated optional UI kit guide destination exists | README/index files and `frontend.md` route readers to an optional UI design guide pattern | `CHK-01`, `CHK-02` |
| `MET-02` | Generic template safety | Source examples are project-specific and may include local framework assumptions | Added docs are framework-agnostic and do not prescribe project-specific UI defaults | `CHK-02`, `CHK-03` |

### Scope

- `REQ-01` Add a generic optional destination/template for documenting an existing UI design guide / UI kit.
- `REQ-02` The destination/template must recommend sections for components, forms, buttons/actions, tables, navigation, states/labels, screenshots, source code paths and agent instructions.
- `REQ-03` Link the optional UI design guide pattern from [`../../engineering/frontend.md`](../../engineering/frontend.md) and relevant README/index files without making it mandatory for projects that do not have a separate UI layer.
- `REQ-04` Keep generic memory-bank free from project-specific framework assumptions, including source-project UI framework defaults.

### Non-Scope

- `NS-01` Do not copy downstream project UI rules, screenshots, helper APIs or code paths into the generic template.
- `NS-02` Do not prescribe Bootstrap, Inspinia, HAML or any other source-project stack as generic defaults.
- `NS-03` Do not change runtime UI implementation, design tokens, components or product UX.
- `NS-04` Do not create or update downstream project instances; this feature only changes the portable memory-bank template.

### Constraints / Assumptions

- `ASM-01` Issue 17 source paths are treated as evidence that this documentation pattern is useful, not as reusable generic content.
- `ASM-02` [`../../engineering/frontend.md`](../../engineering/frontend.md) remains the canonical owner for frontend engineering contract.
- `CON-01` Generic memory-bank must remain portable and framework-agnostic.
- `CON-02` Domain docs must not become the owner of UI design system content; [`../../domain/README.md`](../../domain/README.md) explicitly excludes UI design system from domain ownership.

## Design Requirement Decision

| Decision | Reason | Downstream owner |
| --- | --- | --- |
| `Design required: yes` | The feature requires choosing the owner/location of a new optional documentation pattern and resolving the trade-off between expanding `frontend.md`, adding an optional destination, or adding a template route. | `design.md` |

## Verify

`Verify` задает canonical test case inventory для delivery-единицы. Execution details live in [`implementation-plan.md`](implementation-plan.md).

### Exit Criteria

- `EC-01` A generic optional UI design guide destination/template exists and includes the section set required by `REQ-02`.
- `EC-02` [`../../engineering/frontend.md`](../../engineering/frontend.md) and relevant README/index files route to the optional guide and explicitly present it as optional.
- `EC-03` Added generic docs do not contain project-specific framework defaults or copied downstream project content.

### Traceability matrix

| Requirement ID | Problem refs | Acceptance refs | Checks | Evidence IDs |
| --- | --- | --- | --- | --- |
| `REQ-01` | `ASM-01`, `ASM-02`, `CON-01` | `EC-01`, `SC-01` | `CHK-01`, `CHK-03` | `EVID-01`, `EVID-03` |
| `REQ-02` | `ASM-01`, `CON-01` | `EC-01`, `SC-01` | `CHK-02`, `CHK-03` | `EVID-02`, `EVID-03` |
| `REQ-03` | `ASM-02` | `EC-02`, `SC-02` | `CHK-01`, `CHK-03` | `EVID-01`, `EVID-03` |
| `REQ-04` | `ASM-01`, `CON-01`, `CON-02` | `EC-03`, `NEG-01` | `CHK-02`, `CHK-03` | `EVID-02`, `EVID-03` |

### Acceptance Scenarios

- `SC-01` A downstream project with an existing admin/operator UI can find and instantiate an optional UI design guide that asks for components, forms, actions, tables, navigation, states/labels, screenshots, source paths and agent instructions.
- `SC-02` An agent reading frontend engineering guidance can discover the optional UI design guide as a companion reference, while `frontend.md` remains the canonical engineering contract.

### Negative / Edge Scenarios

- `NEG-01` If source-project examples contain stack-specific UI rules, the generic template does not promote those rules to defaults and instead asks downstream projects to document their own local stack.

### Checks

| Check ID | Covers | How to check | Expected result | Evidence path |
| --- | --- | --- | --- | --- |
| `CHK-01` | `EC-01`, `EC-02`, `SC-01`, `SC-02` | Run `python3 scripts/check_memory_bank_index.py` and `rg --files memory-bank` after implementation. | Navigation links are valid; the optional guide path is discoverable from relevant indexes. | `memory-bank/features/FT-017/evidence/chk-01/` |
| `CHK-02` | `EC-01`, `EC-03`, `NEG-01` | Review added docs and run targeted `rg` checks against the new optional guide and touched routing docs. | Required sections exist; target generic docs do not prescribe source-project frameworks or copied downstream content. | `memory-bank/features/FT-017/evidence/chk-02/` |
| `CHK-03` | `EC-01`, `EC-02`, `EC-03`, `SC-01`, `SC-02`, `NEG-01` | Run `git diff --check` and manually inspect touched docs for owner-boundary consistency. | No whitespace/conflict markers; `frontend.md`, indexes and optional guide do not contradict `brief.md`/`design.md`. | `memory-bank/features/FT-017/evidence/chk-03/` |

### Test matrix

| Check ID | Evidence IDs | Evidence path |
| --- | --- | --- |
| `CHK-01` | `EVID-01` | `memory-bank/features/FT-017/evidence/chk-01/` |
| `CHK-02` | `EVID-02` | `memory-bank/features/FT-017/evidence/chk-02/` |
| `CHK-03` | `EVID-03` | `memory-bank/features/FT-017/evidence/chk-03/` |

### Evidence

- `EVID-01` Link/index audit output showing the optional guide is reachable and markdown dependencies are valid.
- `EVID-02` Content guard output or review note proving generic docs avoid source-project framework defaults and include required guide sections.
- `EVID-03` Final diff hygiene and owner-boundary review note.

### Evidence contract

| Evidence ID | Artifact | Producer | Path contract | Reused by checks |
| --- | --- | --- | --- | --- |
| `EVID-01` | Link/index audit log | implementer | `memory-bank/features/FT-017/evidence/chk-01/result.txt` | `CHK-01` |
| `EVID-02` | Content guard log or review note | implementer / reviewer | `memory-bank/features/FT-017/evidence/chk-02/result.txt` | `CHK-02` |
| `EVID-03` | Diff hygiene and owner-boundary review note | implementer / reviewer | `memory-bank/features/FT-017/evidence/chk-03/result.txt` | `CHK-03` |

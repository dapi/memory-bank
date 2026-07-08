---
title: "FT-017: Design"
doc_kind: feature
doc_function: canonical
purpose: "Solution-space документ для FT-017. Фиксирует selected approach, rationale, C4 applicability, contracts and failure modes for optional UI design guide pattern without redefining scope or execution sequencing."
derived_from:
  - brief.md
  - decision-log.md
status: active
audience: humans_and_agents
must_not_define:
  - ft_017_scope
  - ft_017_acceptance_criteria
  - ft_017_evidence_contract
  - implementation_sequence
---

# FT-017: Design

## Design Pack

| Artifact | Role | Owns |
| --- | --- | --- |
| `design.md` | Feature-local solution owner | `SOL-*`, `ALT-*`, `TRD-*`, `C4-*`, `SD-*`, feature-local `CTR-*`, `INV-*`, `FM-*`, `RB-*` |
| `decision-log.md` | Feature-local decision provenance | FPF-backed `FDL-*` decisions and facts used to close open questions |

## Context

FT-017 adds a generic optional documentation pattern. The design problem is ownership: issue 17 asks for a practical source of truth for existing UI components and local UI conventions, but generic memory-bank must not import downstream framework assumptions.

Existing documents set the owner boundaries:

- [`../../engineering/frontend.md`](../../engineering/frontend.md) owns frontend engineering contracts such as UI surfaces, component boundaries, design system integration and i18n.
- [`../../domain/README.md`](../../domain/README.md) states that domain docs do not define UI design system content.
- [`brief.md`](brief.md) requires discoverability, optionality and framework-agnostic content.

## C4 Applicability

| C4 ID | Decision | Trigger / reason | Artifact |
| --- | --- | --- | --- |
| `C4-00` | `not required` | The feature changes documentation structure and template guidance only. It does not change runtime, container, API, event, schema, file format, security, integration or component collaboration boundaries. | `none` |

## Selected Solution

- `SOL-01` Add `memory-bank/engineering/ui-design-guide/README.md` as an optional engineering-layer UI design guide destination/template for downstream projects with existing UI kits. It should ask for components, forms, buttons/actions, tables, navigation, states/labels, screenshots, source code paths and agent instructions, satisfying `REQ-01` and `REQ-02`.
- `SOL-02` Link the optional guide from `memory-bank/engineering/frontend.md` and `memory-bank/engineering/README.md` while labeling it optional, satisfying `REQ-03`.
- `SOL-03` Keep all guide content generic: use placeholders, prompts and guardrails instead of copied source-project framework rules, satisfying `REQ-04`.

## Alternatives Considered

| Alternative ID | Option | Why not selected |
| --- | --- | --- |
| `ALT-01` | Put the generic guide under `memory-bank/domain/design-guide/` because the source project used that path. | Rejected by `FDL-001`: current `domain/README.md` explicitly excludes UI design system ownership, and issue 17 asks to connect the pattern with engineering frontend guidance. |
| `ALT-02` | Only expand `engineering/frontend.md` with all UI kit guide sections. | Rejected by `FDL-002`: `frontend.md` should stay a concise engineering contract; concrete UI kit reference material needs a separate optional companion to avoid owner blurring. |
| `ALT-03` | Add only a flow template route without an instantiated optional destination in the memory-bank adaptation layer. | Not selected for this feature because issue 17 acceptance requires README/index discoverability of the optional layer. A destination/template in the engineering layer is more directly discoverable for downstream adaptation. |

## Trade-offs

| Trade-off ID | Decision | Benefit | Cost / Risk |
| --- | --- | --- | --- |
| `TRD-01` | Add a separate optional guide instead of embedding everything in `frontend.md`. | Better source-of-truth ergonomics for UI-heavy projects and clearer owner boundary. | One more optional document must be indexed and clearly labeled as optional. |
| `TRD-02` | Use generic prompts/placeholders instead of source-project examples. | Keeps memory-bank portable and avoids accidental framework defaults. | The template is less concrete than a copied project-specific guide; downstream projects must fill local details. |

## Accepted Local Decisions

- `SD-01` The selected owner is the engineering documentation layer, with `frontend.md` as the route to the optional companion.
- `SD-02` The source-project paths in issue 17 are evidence of need, not content sources for the generic guide.
- `SD-03` No ADR is required because the decision is feature-local documentation structure and does not alter reusable architecture.

## Contracts

| Contract ID | Input / Output | Producer / Consumer | Semantics / Constraints |
| --- | --- | --- | --- |
| `CTR-01` | `memory-bank/engineering/ui-design-guide/README.md` | Generic memory-bank template producer / downstream project adapter | The document must define structure and prompts only; downstream projects own actual components, screenshots, helper APIs and local rules. |
| `CTR-02` | `memory-bank/engineering/frontend.md` and `memory-bank/engineering/README.md` links | Generic memory-bank template producer / reader agent | Links must present the guide as optional and route from frontend engineering context. |
| `CTR-03` | Framework-specific source facts | Issue 17 source references / generic template | Source-project framework choices must not become generic defaults. |

## Invariants

- `INV-01` `frontend.md` remains the canonical frontend engineering contract.
- `INV-02` Domain docs do not become owners of UI design system or UI kit content.
- `INV-03` The optional guide must be usable by projects with different UI stacks.

## Failure Modes

- `FM-01` Project-specific rules leak into generic docs. Mitigation: use placeholders, avoid framework defaults in target generic docs, and verify with `CHK-02`.
- `FM-02` The optional guide is perceived as mandatory for all projects. Mitigation: label it optional in routing docs and keep `frontend.md` valid without it.
- `FM-03` `frontend.md` and the optional guide duplicate ownership. Mitigation: `frontend.md` routes to the guide for concrete UI kit details, while the guide points back to frontend engineering as the contract owner.

## Rollout / Backout

| Stage ID | Stage | Entry condition | Backout |
| --- | --- | --- | --- |
| `RB-01` | Add optional guide and routing links | `brief.md`, `decision-log.md` and `design.md` are active | Remove the new optional guide and routing links; `frontend.md` remains the fallback frontend engineering contract. |

## ADR / External Design Dependencies

| Artifact | Current status | Used for | Rule |
| --- | --- | --- | --- |
| `none` | not applicable | No architectural or reusable cross-feature decision is required. | Keep decisions feature-local unless implementation discovers an architecture-level owner change. |

## Traceability

| Requirement ID | Solution refs | Contracts / invariants | Failure / rollout refs |
| --- | --- | --- | --- |
| `REQ-01` | `SOL-01`, `TRD-01`, `SD-01` | `CTR-01`, `INV-01`, `INV-03` | `FM-02`, `FM-03`, `RB-01` |
| `REQ-02` | `SOL-01`, `TRD-02`, `SD-02` | `CTR-01`, `CTR-03`, `INV-03` | `FM-01`, `RB-01` |
| `REQ-03` | `SOL-02`, `TRD-01`, `SD-01` | `CTR-02`, `INV-01` | `FM-02`, `FM-03`, `RB-01` |
| `REQ-04` | `SOL-03`, `TRD-02`, `SD-02`, `SD-03` | `CTR-01`, `CTR-03`, `INV-02`, `INV-03` | `FM-01`, `RB-01` |

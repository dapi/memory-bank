---
title: "FT-016: Design"
doc_kind: feature
doc_function: canonical
purpose: "Solution-space документ для FT-016. Фиксирует selected template shape, owner boundaries, local decisions and trade-offs без переопределения problem space или execution contract."
derived_from:
  - brief.md
  - decision-log.md
status: active
audience: humans_and_agents
must_not_define:
  - ft_016_scope
  - ft_016_acceptance_criteria
  - ft_016_evidence_contract
  - implementation_sequence
---

# FT-016: Design

## Design Pack

| Artifact | Role | Owns |
| --- | --- | --- |
| `design.md` | Feature-local solution owner | `SOL-*`, `ALT-*`, `TRD-*`, `C4-*`, `SD-*`, `CTR-*`, `INV-*`, `FM-*`, `RB-*` |
| `decision-log.md` | Feature-local FPF decision ledger | `DL-*` facts, reasoning, decisions and consequences |

## Context

Issue 16 asks for a generic support template that helps agents govern privacy/source access for AI/agent tooling, observability, logs, transcripts and external metadata. Existing feature support templates cover runtime surfaces, UI references and derived use cases, but none explicitly owns source inventory and privacy/source boundary evidence.

The solution must abstract useful shape from referenced downstream source docs without copying downstream system details into this template repository.

## C4 Applicability

| C4 ID | Decision | Trigger / reason | Artifact |
| --- | --- | --- | --- |
| `C4-00` | `not required` | The feature changes Markdown governance templates and indexes only. It does not change runtime/deployable/container/API/storage/integration/security execution boundaries. | `none` |

## Selected Solution

- `SOL-01` Add `memory-bank/flows/templates/feature/support/privacy-source-boundary.md` as a governed feature-support template for source privacy and evidence boundaries.
- `SOL-02` Update `memory-bank/flows/feature-flow.md` so the new support doc is listed with trigger conditions, ownership limits and support ID taxonomy.
- `SOL-03` Update feature package templates (`README.md`, `brief.md`, `design.md`, `implementation-plan.md`) where routing/discovery wording must mention privacy/source-boundary support docs.
- `SOL-04` Register the new template in `memory-bank/flows/templates/README.md` and `memory-bank/features/README.md` navigation.
- `SOL-05` Keep all generic template surfaces free of downstream-specific source/runtime terms; downstream sources may be cited only in this feature package as issue evidence.

## Alternatives Considered

| Alternative ID | Option | Why not selected |
| --- | --- | --- |
| `ALT-01` | Fold privacy/source-boundary fields into `runtime-surfaces.md` | Rejected because issue 16 covers source permissions, source inventory confidence and privacy exclusions that are not necessarily runtime surfaces. Combining them would blur support doc roles. |
| `ALT-02` | Put privacy/source-boundary rules directly into `brief.md` / `design.md` templates | Rejected because source inventories and evidence boundaries are support/reference facts. Canonical requirements and selected design remain in sibling owner docs. |
| `ALT-03` | Add a global privacy policy document instead of a feature support template | Rejected as out of scope. Issue 16 asks for a feature support template, not a repository-wide legal/security policy. |

## Trade-offs

| Trade-off ID | Decision | Benefit | Cost / Risk |
| --- | --- | --- | --- |
| `TRD-01` | Use a dedicated support template rather than expanding existing support docs | Clear owner boundary and lower risk of privacy/source facts being scattered across canonical docs | Adds one more optional template and index route |
| `TRD-02` | Include local IDs for source and boundary rows | Makes review and traceability precise | Requires updating support ID taxonomy in feature-flow |
| `TRD-03` | Keep source examples abstracted rather than copied | Prevents downstream-specific leakage into generic template | Template may be less detailed than the original downstream artifacts |

## Accepted Local Decisions

- `SD-01` The new support template is named `privacy-source-boundary.md` to cover both privacy constraints and source inventory provenance in one feature-local support artifact.
- `SD-02` The support template uses `doc_kind: feature-support` and `doc_function: reference`; it must not use `evidence` as a default doc function because the template is not itself execution evidence.
- `SD-03` The template should include local support IDs for privacy/source rows but must not introduce canonical `REQ-*`, `SC-*`, `CHK-*`, `EVID-*`, `SOL-*` or `STEP-*`.
- `SD-04` The template should include "allowed metadata" and "explicitly excluded data" as separate sections because issue 16 acceptance depends on preventing unauthorized reads/storage/citation, not only documenting source confidence.

## Contracts

| Contract ID | Input / Output | Producer / Consumer | Semantics / Constraints |
| --- | --- | --- | --- |
| `CTR-01` | New support template wrapper | Template maintainer / feature authors | Wrapper frontmatter remains `doc_function: template`; instantiated frontmatter stays inside embedded contract. |
| `CTR-02` | Instantiated support doc contract | Feature authors / reviewers | Instantiated doc is `doc_kind: feature-support`, `doc_function: reference`, and declares `must_not_define` for scope, selected design, acceptance criteria, canonical checks, evidence contract and implementation sequence. |
| `CTR-03` | Source inventory rows | Feature authors / reviewers | Rows must capture source identity/class, availability/status, confidence, safe use and privacy boundary without exposing private raw contents. |
| `CTR-04` | Excluded data rows | Feature authors / reviewers | Rows must state what must not be read/stored/cited and what explicit permission or owner change would be required before access. |

## Invariants

- `INV-01` Support docs do not own requirements, selected design, canonical acceptance inventory, canonical evidence contract or execution sequence.
- `INV-02` Generic template docs must not contain downstream-specific `zelma`, Codex session internals, concrete runtime details, raw prompt/transcript/log schemas or secrets.
- `INV-03` Source confidence/status must be documented as evidence quality for a source claim, not as permission to read private content.
- `INV-04` Allowed metadata must be narrower than source existence; source existence alone does not authorize reading raw content.

## Failure Modes

- `FM-01` Privacy/source support doc becomes a hidden requirements owner. Mitigation: frontmatter `must_not_define`, role section and feature-flow wording must route changed requirements back to `brief.md`.
- `FM-02` Template examples leak downstream runtime specifics. Mitigation: use generic source classes and verify with `CHK-04`.
- `FM-03` Confidence/status is mistaken for access permission. Mitigation: keep confidence inventory and allowed/excluded data sections separate.
- `FM-04` Agents cite or store raw private artifacts as evidence. Mitigation: evidence section must distinguish allowed evidence carriers from excluded raw source contents.

## Rollout / Backout

| Stage ID | Stage | Entry condition | Backout |
| --- | --- | --- | --- |
| `RB-01` | Documentation template update | `brief.md` and `design.md` active | Revert touched docs in this feature's change surface before implementation is published |
| `RB-02` | Index/link verification | Template and routing edits complete | Restore previous index entries and rerun index audit |

## ADR / External Design Dependencies

| Artifact | Current status | Used for | Rule |
| --- | --- | --- | --- |
| `none` | `n/a` | No ADR dependency | Feature-local decisions stay in `SD-*` / `decision-log.md` |

## Traceability

| Requirement ID | Solution refs | Contracts / invariants | Failure / rollout refs |
| --- | --- | --- | --- |
| `REQ-01` | `SOL-01`, `TRD-01`, `C4-00`, `SD-01` | `CTR-01`, `CTR-02`, `INV-01` | `FM-01`, `RB-01` |
| `REQ-02` | `SOL-01`, `SD-03`, `SD-04`, `TRD-02` | `CTR-03`, `CTR-04`, `INV-03`, `INV-04` | `FM-03`, `FM-04` |
| `REQ-03` | `SOL-02`, `SOL-03` | `CTR-02`, `INV-01` | `FM-01`, `RB-01` |
| `REQ-04` | `SOL-05`, `TRD-03` | `INV-02` | `FM-02` |
| `REQ-05` | `SOL-04` | `CTR-01` | `RB-02` |

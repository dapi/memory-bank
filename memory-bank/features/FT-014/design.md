---
title: "FT-014: Design"
doc_kind: feature
doc_function: canonical
purpose: "Solution-space документ для FT-014. Фиксирует выбранный подход к lightweight epic intake template, boundary updates и локальные решения без переопределения problem space или execution contract."
derived_from:
  - brief.md
status: active
audience: humans_and_agents
must_not_define:
  - ft_014_scope
  - ft_014_acceptance_criteria
  - ft_014_evidence_contract
  - implementation_sequence
---

# FT-014: Design

## Design Pack

| Artifact | Role | Owns |
| --- | --- | --- |
| `design.md` | Feature-local solution owner | `SOL-*`, `ALT-*`, `TRD-*`, `C4-*`, `SD-*`, `CTR-*`, `INV-*`, `FM-*`, `RB-*` |
| `decision-log.md` | Feature-local FPF decision ledger | `DL-*` reasoning records for decisions that do not require ADR |

## Context

Issue 14 asks for a generic lightweight epic intake brief template based on a downstream source template, plus updates to epic template README and `epic-flow.md`. The design problem is boundary control: the new brief must be useful before full epic setup without becoming a second owner for charter, roadmap, risks, subissues or feature execution.

## C4 Applicability

| C4 ID | Decision | Trigger / reason | Artifact |
| --- | --- | --- | --- |
| `C4-00` | `not required` | The change is documentation/template governance only; it does not change runtime, deployable, API, storage, integration, security or component boundaries. | `none` |

## Selected Solution

- `SOL-01` Add `memory-bank/flows/templates/epic/brief.md` as a governed wrapper-template with instantiated frontmatter/body for a lightweight epic intake brief.
- `SOL-02` Register the new template in `memory-bank/flows/templates/epic/README.md` and `memory-bank/flows/templates/README.md` so navigation and index audit remain coherent.
- `SOL-03` Update `memory-bank/flows/epic-flow.md` with an intake layer, optional-intake gate and boundary rules that keep full epic package owners authoritative.
- `SOL-04` Keep source adaptation minimal: reuse only generic structure from the downstream source and avoid source-project-specific content in target template/governance docs.

## Alternatives Considered

| Alternative ID | Option | Why not selected |
| --- | --- | --- |
| `ALT-01` | Copy the source template verbatim | Rejected because issue 14 explicitly forbids moving downstream-specific content; the target repo also has its own wrapper/template style. |
| `ALT-02` | Add only the template file and skip `epic-flow.md` | Rejected because issue 14 explicitly asks to clarify that brief does not replace charter/roadmap and does not own implementation sequence. |
| `ALT-03` | Fold intake fields into `charter.md` only | Rejected because the issue asks for a lighter document before full epic setup, not a heavier charter template. |

## Trade-offs

| Trade-off ID | Decision | Benefit | Cost / Risk |
| --- | --- | --- | --- |
| `TRD-01` | Keep intake brief optional and draft-oriented | Allows early proposal capture without requiring full epic package upfront | Requires explicit boundary wording to prevent duplicate ownership |
| `TRD-02` | Update both local and global template indexes | Keeps audit reachability and discoverability strong | Slightly expands change surface beyond the issue's explicit epic README line, but only for index consistency |

## Accepted Local Decisions

- `SD-01` The epic intake `brief.md` is an optional early layer and not a canonical replacement for `charter.md`, `roadmap.md`, `subissues.md`, `risks.md` or `decision-log.md`.
- `SD-02` The intake template may include candidate feature brief links as readiness context, but accepted delivery subissues remain owned by `subissues.md`.
- `SD-03` No ADR is required because the decision is feature-local template governance and does not define reusable architecture or runtime boundaries.

## Contracts

| Contract ID | Input / Output | Producer / Consumer | Semantics / Constraints |
| --- | --- | --- | --- |
| `CTR-01` | `memory-bank/flows/templates/epic/brief.md` | Template maintainer / epic author | Must include problem, outcome, rough scope, non-scope and readiness notes; must not define implementation sequence, feature acceptance contracts or selected solution. |
| `CTR-02` | `memory-bank/flows/epic-flow.md` | Governance maintainer / epic and feature authors | Must state that full epic package owners remain authoritative for roadmap, accepted subissues, risks and decisions. |
| `CTR-03` | Template indexes | Template maintainer / index audit | Must link the new template from local epic template index and global templates index. |

## Invariants

- `INV-01` Intake brief content must remain generic and source-backed; no source-project-specific terms may appear in target template/governance docs.
- `INV-02` The new brief must not create an `implementation-plan.md` path inside epic packages.
- `INV-03` If intake facts mature into authoritative epic scope or roadmap, the owner becomes the full epic package document, not the intake brief.

## Failure Modes

- `FM-01` Duplicate ownership: authors treat intake brief as authoritative after full epic setup. Mitigation: add layer and boundary rules in `epic-flow.md`.
- `FM-02` Source leakage: downstream project terms appear in generic template docs. Mitigation: explicit `CHK-03` grep and review.
- `FM-03` Broken reachability: new template or feature docs are orphaned. Mitigation: update indexes and run `CHK-01`.

## Rollout / Backout

| Stage ID | Stage | Entry condition | Backout |
| --- | --- | --- | --- |
| `RB-01` | Add docs and indexes in current branch | `brief.md` and `design.md` active for FT-014 | Revert only FT-014 and issue-scope docs in this branch before merge |
| `RB-02` | Verify locally | Changes are present | If audit fails, update owner/index docs before continuing |

## ADR / External Design Dependencies

| Artifact | Current status | Used for | Rule |
| --- | --- | --- | --- |
| `none` | `n/a` | No reusable architecture decision is required | Do not create ADR for feature-local template governance |

## Traceability

| Requirement ID | Solution refs | Contracts / invariants | Failure / rollout refs |
| --- | --- | --- | --- |
| `REQ-01` | `SOL-01`, `TRD-01`, `SD-01` | `CTR-01`, `INV-01`, `INV-02` | `FM-01`, `FM-02`, `RB-01` |
| `REQ-02` | `SOL-02`, `TRD-02` | `CTR-03` | `FM-03`, `RB-02` |
| `REQ-03` | `SOL-02`, `TRD-02` | `CTR-03` | `FM-03`, `RB-02` |
| `REQ-04` | `SOL-03`, `SD-01`, `SD-02` | `CTR-02`, `INV-02`, `INV-03` | `FM-01`, `RB-01` |
| `REQ-05` | `SOL-04` | `INV-01` | `FM-02`, `RB-02` |
| `REQ-06` | `SOL-02`, `SOL-03` | `CTR-03` | `FM-03`, `RB-02` |

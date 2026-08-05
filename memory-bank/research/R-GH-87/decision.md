---
title: "R-GH-87: Research Decision"
doc_kind: research
doc_function: canonical
purpose: "Decision rationale and promotion map for R-GH-87."
derived_from:
  - brief.md
  - synthesis.md
  - ../../flows/research.md
status: active
audience: humans_and_agents
---

# R-GH-87: Research Decision

## Decision

| Field | Value |
| --- | --- |
| Decision owner | `dapi/memory-bank` maintainer |
| Decision date | 2026-08-05 |
| Decision reference | [Issue #87 maintainer routing record](https://github.com/dapi/memory-bank/issues/87#issuecomment-5087962973) and this evidence package |

Terminal disposition is recorded only in sibling `brief.md` as
`research_status: validated`.

## Decision Rationale

- `FND-01` and `FND-02` show a clean ownership boundary: reusable validation
  is in the standalone CLI, while the remaining Ruby scripts are local
  producer-side integration validation.
- `FND-03` confirms that the task which performed the CLI migration is closed;
  no new reusable CLI requirement is evidenced here.
- `LIM-01` means this is not an assertion that the Ruby checks must remain
  forever; it rejects deletion or a duplicate CLI task on the present evidence.

## Recommendation

- `REC-01` Retain `tools/validate-priming-manifests*.rb`, create no duplicate
  `memory-bank-cli` issue, and record the ownership decision in this completed
  research package. Confidence: high for ownership; see `LIM-01` for the
  boundary of the conclusion.

## Alternatives Considered

| Alternative | Why not selected / what would change the decision |
| --- | --- |
| Delete remaining `tools/` scripts now. | Rejected: current CI invokes them as a distinct producer-side integration surface; a separately routed removal/change would need evidence that this coverage is obsolete. |
| Create a new CLI issue. | Rejected: CLI migration is already closed and this research identifies no missing reusable CLI contract. A concrete reusable requirement would justify new Task Routing. |

## Promotion and Handoff Map

| ID | Accepted or retained fact | Canonical downstream owner | Target route / link |
| --- | --- | --- | --- |
| `HD-01` | The current Ruby priming-manifest checks remain repository-specific integration validation. | This package's retained decision record; no new active project fact is required. | No follow-up delivery route. |
| `HD-02` | Future extraction requires a concrete reusable CLI contract. | A new issue owned by the requester. | New [Task Routing](../../flows/routing.md) only if that requirement appears. |

## Closure Check

- [x] Sibling `brief.md` records matching terminal `research_status: validated`.
- [x] `synthesis.md` answers `RQ-01`, and this recommendation traces to `FND-*` and `LIM-*`.
- [x] Handoff creates neither delivery scope nor implementation steps by implication.

---
title: "R-GH-87: Validation Tool Ownership"
doc_kind: research
doc_function: canonical
purpose: "Canonical decision question, boundaries and lifecycle state for R-GH-87."
derived_from:
  - ../../flows/research.md
status: active
research_status: validated
audience: humans_and_agents
---

# R-GH-87: Validation Tool Ownership

## Intake

| Field | Value |
| --- | --- |
| Source / trigger | [dapi/memory-bank#87](https://github.com/dapi/memory-bank/issues/87) |
| Research owner | delivery orchestrator |
| Decision owner | `dapi/memory-bank` maintainer |
| Research mode | `technical_discovery` |
| Decision deadline / timebox | Bounded desk research completed 2026-08-05 |

## Decision Question

- `RQ-01` Does the current repository still contain validation implementation
  that should be removed or moved to `dapi/memory-bank-cli`, and is a new
  standalone CLI task required?

## Working Hypotheses

- `HYP-01` Reusable `lint` and `doctor` validation is already owned by
  `dapi/memory-bank-cli`; the remaining Ruby checks in this repository are
  producer-side integration checks and should remain here.

## Compact Method Record

- Method and source strategy: inspect the issue and maintainer routing record,
  the migration commit, the current CI workflow and the linked CLI migration
  issue.
- Collection window and context: GitHub and the checked-out repository on
  2026-08-05; current `main` at `23ff0a3905645c6f4e37809ae7e97d43cb937a37`.
- Evidence-quality criteria: primary GitHub issue/commit/workflow sources;
  distinguish current executable use from historical intent.
- Applicable privacy, consent, legal, security and vendor-access constraints:
  none; only public repository metadata and checked-out source are used.
- Bias risk and disconfirming signal: prior maintainer conclusions could be
  stale; current CI still invoking local reusable lint/doctor implementation
  would contradict `HYP-01`.

## Scope

- `RSC-01` Ownership of validation used by this repository's CI and remaining
  `tools/` scripts, relative to `memory-bank-cli`.

## Non-Scope

- `RNS-01` Changing CLI behavior, release process, or template CI contracts.
- `RNS-02` Deleting repository-specific checks without a separately routed
  delivery task.

## Assumptions and Known Evidence

| ID | Statement | Type | Source / confidence |
| --- | --- | --- |
| `ASM-01` | The issue seeks a decision before any deletion or new CLI work. | Evidence | [Issue #87](https://github.com/dapi/memory-bank/issues/87) |
| `ASM-02` | The maintainer is the decision owner for this repository decision. | Evidence | [Routing comment](https://github.com/dapi/memory-bank/issues/87#issuecomment-5087962973) |

## Stopping Condition

- `STOP-01` Stop once the migration history, current CI ownership boundary and
  linked CLI migration task either support or refute `HYP-01`.

## Open Questions

| Question | Blocks | Owner | Resolution evidence |
| --- | --- | --- | --- |
| None | — | — | `STOP-01` is satisfied. |

## Boundary Check

- [x] This brief contains a question and hypothesis, not findings presented as facts.
- [x] Every known fact has a clickable source link.
- [x] No committed delivery scope, selected solution, ADR decision or implementation sequence is defined here.
- [x] Required privacy, consent, legal, security and access constraints are explicitly `none`.

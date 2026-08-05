---
title: "FT-109: Task Capsule Design"
doc_kind: feature
doc_function: canonical
purpose: "Own the selected storage boundary, schema shape and Feature/Epic field ownership for the Task Capsule contract."
derived_from:
  - brief.md
  - ../../../template/memory-bank/flows/task-capsule.md
  - ../../../template/memory-bank/flows/feature.md
  - ../../../template/memory-bank/flows/epic.md
  - ../../../template/memory-bank/flows/templates/process/session-handoff.md
status: active
audience: humans_and_agents
---

# FT-109: Task Capsule Design

## Design Pack

| Path | Classification | Ownership |
| --- | --- | --- |
| `design.md` | root | selected storage and ownership facts |
| `../../../template/memory-bank/flows/task-capsule.md` | external-dependency | generic contract and reusable schema |
| `../../../template/memory-bank/flows/templates/process/task-capsule.md` | constituent | instantiation template |

## Context

The parent initiative requires durable continuation while existing Feature and
Epic flows already own lifecycle facts. The design must permit a session to
select one safe action without copying those facts into a second active owner.

## C4 Applicability

`C4 not required`: this is a documentation contract and does not add runtime
components or deployment topology. The relevant boundary is documented as
artifact ownership and file references.

## Architecture Coverage Decision

| Aspect | Decision | Rationale |
| --- | --- | --- |
| Components | covered | Feature/Epic package, canonical owners, capsule and optional registry projection are identified. |
| Connectors | covered | Relative links and immutable evidence refs are the only handoff bindings. |
| Configuration | N/A | No runtime/configuration is introduced. |
| Behavioral semantics | covered | Resume, pause, reroute and closure update rules are defined. |
| Quality/evolution concerns | covered | SSoT, revisions, stale refs and projection-only registry rules are explicit. |

## Selected Solution

- `SOL-01` Store one dedicated `task-capsule.md` in each active Feature or Epic package.
- `SOL-02` Keep lifecycle/status/requirements/decisions in existing canonical owners; capsule fields reference owner paths and revisions.
- `SOL-03` Reuse the session-handoff continuation sequence, with a structured schema and one exact `next_action`.
- `SOL-04` Permit a registry/index only as a read/projection view.

## Alternatives Considered

| Alternative | Decision | Reason |
| --- | --- | --- |
| `ALT-01` Put full capsule state in owner frontmatter | rejected | Mixes transient resumable state with canonical lifecycle and makes each owner carry a second contract. |
| `ALT-02` Use only a generic session handoff outside packages | rejected | Loses package ownership and makes Feature/Epic location ambiguous. |
| `ALT-03` Dedicated package-local capsule plus owner refs | selected | Keeps durable continuation local, explicit and non-duplicating. |

## Trade-offs

- `TRD-01` A session reads the capsule and referenced owners, but the schema is durable and machine-readable without requiring a full transcript.
- `TRD-02` Two package types have two paths, but this preserves the existing Feature/Epic ownership boundary and prevents an umbrella registry from owning slice state.

## Accepted Local Decisions

- `SD-01` `task-capsule.md` is the package-local handoff artifact; its frontmatter carries only document metadata and `status`.
- `SD-02` Operational state is a structured body contract; owner documents remain authoritative for lifecycle and requirements.
- `SD-03` Every capsule update leaves exactly one `next_action` and records stale-reference or rejected-handoff diagnostics.

## Contracts

- `CTR-01` Schema keys and ownership rules are canonical in `template/memory-bank/flows/task-capsule.md`.
- `CTR-02` Feature storage is `memory-bank/features/FT-XXX/task-capsule.md`; Epic storage is `memory-bank/epics/EP-XXX/task-capsule.md`.
- `CTR-03` Registry entries are projections and cannot be used to resolve conflicts with capsule or canonical owners.

## Invariants

- `INV-01` One task has one active capsule and one exact next action.
- `INV-02` A capsule never overrides a canonical owner or silently repairs a stale reference.
- `INV-03` Epic slice state is handed off to a separate routed Feature package.

## Failure Modes

- `FM-01` Missing/stale owner ref: pause, re-ground and update the owner or reroute.
- `FM-02` More than one next action: reject handoff and reduce to one sequenced action.
- `FM-03` Registry conflicts with owner: ignore projection and report the canonical owner conflict.

## Rollout / Backout

Not applicable to runtime rollout. Revert the documentation change as one Git
revision; downstream packages continue using their existing owner documents.

## Design Verification

| Analysis class | Required | Method | Result/evidence |
| --- | --- | --- | --- |
| Contract compatibility | yes | Compare schema and ownership with Feature, Epic and session-handoff flow docs. | `EVID-02` |
| State/transition completeness | yes | Trace bootstrap, gate, handoff, pause, reroute and closure updates. | `EVID-02` |
| Failure propagation | yes | Check stale refs, rejected handoff and registry conflict diagnostics. | `EVID-01` |
| Concurrency/ordering | no | No runtime concurrency; single writer and exact-next-action rule are documented. | N/A rationale in `INV-01` |
| Security boundaries | no | No access or auth behavior changes. | N/A |
| Capacity/latency | no | No runtime path. | N/A |
| Migration/evolution safety | yes | Check owner revisions, package-local paths and projection-only registry rule. | `EVID-02` |

## External Dependency Readiness

Existing Feature, Epic and session-handoff documents are active and remain
owners; this design adds links and explicit mapping without changing their
existing lifecycle contracts.

## Traceability

| Requirement | Solution refs | Contract/invariant refs | Evidence |
| --- | --- | --- | --- |
| `REQ-01` | `SOL-01`, `SOL-03` | `CTR-01`, `INV-01` | `EVID-01` |
| `REQ-02` | `SOL-01`, `SOL-02` | `CTR-02`, `INV-02` | `EVID-02` |
| `REQ-03` | `SOL-02`, `SOL-03` | `CTR-01`, `INV-03` | `EVID-02` |
| `REQ-04` | `SOL-04` | `CTR-03`, `FM-03` | `EVID-01` |

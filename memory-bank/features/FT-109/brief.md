---
title: "FT-109: Task Capsule Contract"
doc_kind: feature
doc_function: canonical
purpose: "Define a generic persisted Task Capsule contract that lets Feature and Epic work resume safely without replaying the full transcript."
derived_from:
  - ../../../template/memory-bank/flows/feature.md
  - ../../../template/memory-bank/flows/task-capsule.md
  - ../../../template/memory-bank/flows/routing.md
  - ../../../template/memory-bank/engineering/validation-profiles.md
  - ../../../template/memory-bank/flows/templates/process/session-handoff.md
status: active
delivery_status: done
audience: humans_and_agents
must_not_define:
  - runtime_orchestration
  - temporal_integration
  - automatic_validation_implementation
---

# FT-109: Task Capsule Contract

## What

### Problem

Feature and Epic flows have canonical documents, but an interrupted session
does not yet have one compact, durable handoff contract that identifies the
current stage, exact next action and evidence without copying those owners.

### Outcome

| Metric ID | Metric | Baseline | Target | Measurement method |
| --- | --- | --- | --- | --- |
| MET-01 | Resume state completeness | No generic capsule contract | All required fields, owners and resume inputs documented | Read-through against issue acceptance and flow owners |

### Scope

- `REQ-01` Document a schema/example containing task identity, route/stage, roles, artifact refs, actions, assumptions, risks, evidence, stop conditions and handoff diagnostic.
- `REQ-02` Define Feature and Epic storage locations, field ownership and lifecycle update rules.
- `REQ-03` Map the capsule to existing Feature/Epic lifecycle documents and session handoff continuation inputs.
- `REQ-04` Preserve canonical owners and state that any registry is projection-only.

### Non-Scope

- `NS-01` Runtime orchestration, Temporal integration, queues, retries or automatic resume.
- `NS-02` Automatic validators or mutations of canonical owner documents.
- `NS-03` Project-specific roles, downstream implementation and parent issue closure.

## Constraints / Assumptions

- `ASM-01` Generic rules belong in `template/memory-bank/`; FT-109 is the project-local delivery owner and evidence package.
- `ASM-02` Existing Feature/Epic lifecycle documents remain canonical for lifecycle facts.
- `CON-01` A capsule must be sufficient to choose one next safe action but must not duplicate requirements, decisions, statuses or registries.

## Design Requirement Decision

| Decision | Reason | Downstream owner |
| --- | --- | --- |
| `Design required: yes` | The issue explicitly requires choosing frontmatter vs dedicated handoff storage and defines a new persisted file-format contract. | [`design.md`](design.md) |

## Artifact Routing Decision

| Artifact | Decision | Trigger / reason | Route / owner |
| --- | --- | --- | --- |
| `design.md` | selected | Storage boundary, schema ownership and lifecycle mapping require solution-space reasoning. | `design.md` |
| `task-capsule.md` | selected | The active feature package needs the same resumable state carrier it defines. | `task-capsule.md` |
| `use-cases/README.md` | omitted | No project-level user scenario changes. | none |

## Validation Profile Decision

| Profile | Triggers / rationale | Downgrade approval |
| --- | --- | --- |
| `documentation` | Only governed Markdown/template documentation changes; no executable behavior, runtime contract, config or release path. Apply link, schema/frontmatter, targeted docs checks and semantic read-through. | none |

## Verify

### Exit Criteria

- `EC-01` A reusable schema/example documents every required field and explicit SSoT boundary.
- `EC-02` Feature and Epic paths, field owners, lifecycle updates and resume inputs are explicit and traceable to their canonical flows.
- `EC-03` Registry projection and all out-of-scope runtime work are explicitly bounded.

### Traceability matrix

| Requirement ID | Problem refs | Acceptance refs | Checks | Evidence IDs |
| --- | --- | --- | --- | --- |
| `REQ-01` | `CON-01` | `EC-01`, `SC-01` | `CHK-01` | `EVID-01` |
| `REQ-02` | `ASM-01`, `ASM-02` | `EC-02`, `SC-02` | `CHK-02` | `EVID-02` |
| `REQ-03` | `ASM-02` | `EC-02`, `SC-03` | `CHK-02` | `EVID-02` |
| `REQ-04` | `CON-01` | `EC-03`, `SC-04` | `CHK-01`, `CHK-02` | `EVID-01`, `EVID-02` |

### Acceptance Scenarios

- `SC-01` A reader can instantiate a capsule and find every required field with a concrete example.
- `SC-02` A reader can identify the Feature and Epic package path and the canonical owner for each field.
- `SC-03` A new session can select the next action from the capsule plus referenced owners and priming inputs, without the prior transcript.
- `SC-04` A registry projection cannot override requirements, lifecycle status or capsule state.

### Checks

| Check ID | Covers | How to check | Expected result | Evidence path |
| --- | --- | --- | --- | --- |
| `CHK-01` | `EC-01`, `EC-03`, `SC-01`, `SC-04` | `rg -n` required keys and SSoT/projection rules in `template/memory-bank/flows/task-capsule.md` | All fields and boundaries are present | `artifacts/ft-109/verify/chk-01/` |
| `CHK-02` | `EC-02`, `SC-02`, `SC-03` | Link/schema lint plus semantic read-through of task-capsule, Feature, Epic and session handoff docs | Links resolve and continuation map is consistent | `artifacts/ft-109/verify/chk-02/` |

### Test matrix

| Check ID | Evidence IDs | Evidence path |
| --- | --- | --- |
| `CHK-01` | `EVID-01` | `artifacts/ft-109/verify/chk-01/` |
| `CHK-02` | `EVID-02` | `artifacts/ft-109/verify/chk-02/` |

### Evidence

- `EVID-01` Required-field and SSoT boundary scan output.
- `EVID-02` Template lint/doctor output and semantic read-through record.

### Evidence contract

| Evidence ID | Artifact | Producer | Path contract | Reused by checks |
| --- | --- | --- | --- | --- |
| `EVID-01` | Required-field scan | verify-runner | `artifacts/ft-109/verify/chk-01/` | `CHK-01` |
| `EVID-02` | Documentation validation and read-through | verify-runner / reviewer | `artifacts/ft-109/verify/chk-02/` | `CHK-02` |

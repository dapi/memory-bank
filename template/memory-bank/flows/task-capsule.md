---
title: Task Capsule Contract
doc_kind: governance
doc_function: canonical
purpose: Persist the minimum active-task state needed for safe continuation across sessions, machines, or agents without making the capsule a second source of truth.
derived_from:
  - ../dna/governance.md
  - ../dna/frontmatter.md
  - routing.md
  - feature.md
  - epic.md
  - templates/process/session-handoff.md
canonical_for:
  - task_capsule_schema
  - task_capsule_storage
  - task_capsule_field_ownership
  - task_capsule_resume_contract
status: active
audience: humans_and_agents
---

# Task Capsule Contract

Task Capsule — это package-local persisted handoff для одной активной задачи.
Он отвечает на вопрос «что безопасно сделать дальше», но не становится
владельцем требований, solution decisions, lifecycle status или roadmap.

Adoption is required when a task must survive a session/machine/role boundary
or use a persisted handoff. A compact task that completes within one governed
session may follow its flow without creating a capsule; in that case the flow's
ordinary owner documents and handoff rules remain sufficient.

## Storage and SSoT boundary

Храни капсулу как отдельный `task-capsule.md` в package, который владеет
текущей delivery unit:

| Package | Capsule path | Canonical owners referenced by the capsule |
| --- | --- | --- |
| Feature | `memory-bank/features/FT-XXX/task-capsule.md` | `brief.md`, optional `design.md`/design pack, optional `implementation-plan.md` after Plan Ready, ADRs and evidence carriers |
| Epic | `memory-bank/epics/EP-XXX/task-capsule.md` | `README.md`, `brief.md`/`charter.md`, `roadmap.md`, `risks.md`, `subissues.md` and evidence carriers |

Epic capsule governs only the epic-level stage and its handoff to a routed
delivery slice. After `Roadmap Ready`, each accepted slice gets its own issue,
Task Routing, Feature package and Feature capsule. The Epic capsule must not
become a registry of feature lifecycle state.

The capsule uses a minimal combination of persisted state and references:

- canonical owners retain requirements, scope, acceptance, decisions,
  lifecycle statuses, roadmap, risks and implementation sequencing;
- the capsule retains the current resumable checkpoint, references to those
  owners, the single next action, evidence and handoff diagnostics;
- a task registry or index may project `task_id`, package path, stage and
  capsule status, but never writes or overrides capsule or owner fields;
- a completed handoff archives the capsule with the package, or removes it only
  under the package's archival policy. It does not rewrite historical owner
  documents.

## Schema

An instantiated capsule has YAML frontmatter limited to document identity and
publication status. The operational contract below is the body of the file.
Values are concrete refs or explicit `none`; do not use prose that requires the
next session to reconstruct hidden context.

```yaml
task_id: "FT-XXX | EP-XXX"
route:
  name: "Feature | Epic | ..."
  revision: 1
  source_ref: "../../flows/routing.md"
stage:
  current: "Plan Ready"
  canonical_owner: "brief.md"
current_role: "delivery-owner"
next_role: "reviewer"
artifact_refs:
  - path: "brief.md"
    owner: "feature-brief"
    revision: "<commit-or-document-revision>"
  - path: "implementation-plan.md"
    owner: "feature-plan"
    revision: "<commit-or-document-revision>"
completed_action:
  id: "STEP-01"
  summary: "<one completed action>"
  evidence_refs: ["EVID-01"]
current_action:
  id: "CP-01"
  summary: "<current checkpoint or verification>"
next_action:
  owner: "<role or named owner>"
  id: "STEP-02 | GATE-02"
  summary: "<one exact action>"
  stop_condition: "<predicate or event that stops continuation>"
assumptions:
  - id: "ASM-01"
    summary: "<assumption>"
    owner_ref: "brief.md"
risks:
  - id: "ER-01"
    summary: "<open risk>"
    owner_ref: "implementation-plan.md | risks.md"
evidence:
  - id: "EVID-01"
    ref: "<path, issue comment, CI run, or command result>"
    proves: "<predicate>"
stop_conditions:
  - id: "STOP-01"
    trigger: "<divergence, missing owner, failed gate, or human-only decision>"
    action: "<pause, reroute, or request approval>"
last_handoff_diagnostic:
  status: "accepted | rejected | blocked | none"
  reason_code: "<stable reason or none>"
  message: "<short actionable diagnostic>"
  repair_action: "<exact repair or resume action>"
  evidence_ref: "<diagnostic evidence or none>"
updated_at: "<ISO-8601 timestamp>"
```

`stage.current` and `canonical_owner` are references to lifecycle facts; they
do not override those owners. `current_action` is the in-flight checkpoint,
while `next_action` is always exactly one action. If work is paused, keep the
last completed action and change only the current/next action and diagnostic.

## Field ownership and update rules

| Field | Direct owner | Update rule |
| --- | --- | --- |
| `task_id`, `route` | Task Routing / package README | Set at routing; update only after rerouting with a new route revision. |
| `stage.current`, `stage.canonical_owner` | Feature Flow or Epic Flow owner | Mirror the latest accepted lifecycle gate; never invent a stage in the capsule. |
| `current_role`, `next_role` | Active workflow handoff | Set by the current gate/role; a role may update its own handoff only within the allowed transition. |
| `artifact_refs` | Current package owner | Replace refs when the canonical artifact revision changes; each ref must identify owner and revision. |
| `completed_action`, `current_action`, `next_action` | Current writer / handoff owner | Write atomically at checkpoint; exactly one `next_action` must remain. |
| `assumptions` | Problem/design/plan owner | Add or resolve through the canonical owner first, then refresh the capsule ref. |
| `risks` | Feature `brief.md`/plan or Epic `risks.md` | Track only active risks and point to their owner; do not duplicate the risk register. |
| `evidence` | Producer of the evidence | Add immutable carrier refs and the predicate each proves. |
| `stop_conditions` | Current flow owner | Derive from flow gates, plan and autonomy boundaries; a triggered condition blocks continuation. |
| `last_handoff_diagnostic` | Handoff validator or receiving role | Set on every accepted/rejected/blocked handoff; rejected diagnostics must include `repair_action`. |
| `updated_at` | Capsule writer | Update on every capsule mutation; it has no lifecycle meaning. |

For Feature packages, `brief.md` owns problem, scope, non-scope, assumptions,
validation profile and acceptance; `design.md` owns selected solution and
contracts when required; `implementation-plan.md` owns execution sequencing.
For Epic packages, `brief.md` owns proposal facts, `charter.md` owns accepted
initiative scope, `roadmap.md` owns waves and handoff gates, `risks.md` owns
cross-feature risks, and `subissues.md` owns the projection of slices.

## Lifecycle mapping

Update the capsule at bootstrap, every accepted gate, writer handoff,
checkpoint, review/CI result, reroute, pause and closure. The minimum mapping
is:

| Lifecycle point | Feature capsule | Epic capsule |
| --- | --- | --- |
| Problem/Proposal Ready | refs active brief and its verify/proposal evidence | refs active proposal `brief.md` and decision owner |
| Solution/Epic Ready | refs active design pack and solution gate evidence | refs active charter and epic readiness evidence |
| Plan/Roadmap Ready | refs active implementation plan and plan-review evidence | refs roadmap, risks and subissues; next action is a separate Task Routing handoff |
| Execution | refs current plan checkpoint, changed artifacts and test evidence | tracks only epic coordination; each slice continues in its own Feature capsule |
| Done/closed | final evidence, handoff diagnostic and archive disposition | accepted subissues and epic closure evidence; no slice details copied |

On reroute, stop the current writer, record the handback and invalidated refs,
then update the route owner before changing the capsule route. On a human gate,
set one exact request, responsible owner, resume condition and next action.

## Resume contract

A new session reads, in order:

1. the package `task-capsule.md`;
2. the `artifact_refs` at their recorded revisions;
3. the selected flow and its priming inputs for `stage.current`;
4. only the evidence and canonical owner sections needed by `next_action`.

It verifies the route, owner revisions, worktree/branch and stop conditions
before writing. If any reference is stale, missing or contradicts its owner,
the session pauses and repairs the owner or reroutes; it must not infer the
missing state from the previous transcript. The capsule is sufficient to select
the next safe action, while canonical owners remain necessary to execute and
verify it.

## Example

```yaml
---
title: "FT/EP-XXX: Task Capsule"
doc_kind: process
doc_function: canonical
purpose: "Resume state for the active Feature or Epic delivery unit."
derived_from:
  - ../../flows/task-capsule.md
  - brief.md
status: active
audience: humans_and_agents
---

task_id: "FT-XXX | EP-XXX"
route: {name: Feature, revision: 1, source_ref: ../../flows/routing.md}
stage: {current: Execution, canonical_owner: brief.md}
current_role: delivery-owner
next_role: reviewer
artifact_refs:
  - {path: brief.md, owner: feature-brief, revision: <commit-sha>}
  - {path: design.md, owner: feature-design, revision: <commit-sha>}
  - {path: implementation-plan.md, owner: feature-plan, revision: <commit-sha>}
completed_action: {id: STEP-01, summary: "Updated capsule contract", evidence_refs: [EVID-01]}
current_action: {id: CP-01, summary: "Run template validation"}
next_action: {owner: delivery-owner, id: CHK-01, summary: "Run required checks", stop_condition: "Any required check fails"}
assumptions: [{id: ASM-01, summary: "Registry remains projection-only", owner_ref: brief.md}]
risks: [{id: ER-01, summary: "Downstream adoption may need a separate issue", owner_ref: implementation-plan.md}]
evidence: [{id: EVID-01, ref: "template/memory-bank/flows/task-capsule.md", proves: "Schema and ownership are documented"}]
stop_conditions: [{id: STOP-01, trigger: "Owner revision diverges", action: "Pause and re-ground"}]
last_handoff_diagnostic: {status: none, reason_code: none, message: none, repair_action: none, evidence_ref: none}
updated_at: 2026-08-05T00:00:00Z
```

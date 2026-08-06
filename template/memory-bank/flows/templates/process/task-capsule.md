---
title: "FT/EP-XXX: Task Capsule"
doc_kind: process
doc_function: template
purpose: "Instantiable package-local Task Capsule for resumable active-task state."
derived_from:
  - ../../task-capsule.md
  - ../../../dna/frontmatter.md
status: active
audience: humans_and_agents
template_for: process
template_target_path: ../../../features/FT-XXX/task-capsule.md
template_alternate_target_path: ../../../epics/EP-XXX/task-capsule.md
canonical_for:
  - process_template_task_capsule
---

# FT/EP-XXX: Task Capsule

Copy the embedded contract below into the owning Feature or Epic package. Keep
canonical requirements, solution, lifecycle and roadmap facts in their owners;
this document stores only resumable state and references.

## Instantiated Frontmatter

```yaml
---
title: "FT-XXX: Task Capsule"
doc_kind: process
doc_function: canonical
purpose: "Resumable state for the active Feature or Epic package task."
derived_from:
  - ../../flows/task-capsule.md
  - ../../dna/frontmatter.md
status: active
audience: humans_and_agents
---
```

```yaml
task_id: "FT-XXX | EP-XXX"
route: {name: "Feature | Epic", revision: 1, source_ref: "../../flows/routing.md"}
stage: {current: "<canonical stage>", canonical_owner: "<owner path>"}
current_role: "<role>"
next_role: "<role>"
artifact_refs: []
completed_action: {id: "<ID>", summary: "<one action>", evidence_refs: []}
current_action: {id: "<ID>", summary: "<current checkpoint>"}
next_action: {owner: "<owner>", id: "<ID>", summary: "<one exact action>", stop_condition: "<condition>"}
assumptions: []
risks: []
evidence: []
stop_conditions: []
last_handoff_diagnostic: {status: none, reason_code: none, message: none, repair_action: none, evidence_ref: none}
updated_at: "<ISO-8601 timestamp>"
```

Before resuming, read this capsule, referenced canonical owners, the selected
flow's priming inputs and the evidence needed for `next_action`. If a reference
is stale or contradictory, stop and repair/reroute before mutation.

## Continuation Priming Inputs

Read the owning flow's continuation source set before resuming. Record the exact
manifest or stable source and the last verified owner revision here.

## Human-readable Handoff

### Current State

- Completed action, current checkpoint and current canonical owner.

### Assumptions and Open Risks

- Link each active assumption and risk to its canonical owner.

### Next Checks

- Keep the checks needed before `next_action` can safely continue.

### Handoff Diagnostic

- Record accepted, rejected or blocked handoff status and the repair action.

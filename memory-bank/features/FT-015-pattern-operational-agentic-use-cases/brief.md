---
title: "FT-015: Operational and Agentic Use Cases Pattern"
doc_kind: feature
doc_function: canonical
purpose: "Canonical brief for adding generic guidance about operational and agentic use cases without mixing project-specific source examples into the template."
derived_from:
  - ../../flows/feature-flow.md
  - ../../use-cases/README.md
  - ../../flows/templates/use-case/UC-XXX.md
status: active
delivery_status: done
audience: humans_and_agents
must_not_define:
  - implementation_sequence
  - solution_space
source_issue: "https://github.com/dapi/memory-bank/issues/15"
---

# FT-015: Operational and Agentic Use Cases Pattern

## What

### Problem

`memory-bank/use-cases/README.md` already says that use cases can describe
project-level operational scenarios, but it does not explicitly explain the
operational/agentic class called out by issue
[#15](https://github.com/dapi/memory-bank/issues/15): handoff, recovery,
parallel delivery, environment diagnostics, and machine-readable status.

The downstream source examples named in the issue show real `UC-*` files for
this class, but they are project-specific evidence. The generic template needs
the pattern, not source project names, tool names, commands, runtime
assumptions, or vendor-specific workflows.

### Outcome

| Metric ID | Metric | Baseline | Target | Measurement method |
| --- | --- | --- | --- | --- |
| `MET-01` | Generic operational/agentic guidance coverage | README explains general use cases but not the explicit operational/agentic class from issue #15 | README contains reusable guidance for operational/agentic workflows, UC-vs-SC routing, machine-readable contracts, and recovery/postconditions | `CHK-01`, `CHK-03` |
| `MET-02` | Template genericity | Source examples include downstream-specific tool/runtime details | Updated generic docs contain no source project terms, tool names, runtime details, or command examples | `CHK-02` |

### Scope

- `REQ-01` Update `memory-bank/use-cases/README.md` with a generic section for operational and agentic use cases.
- `REQ-02` Add guidance for when an operational/agentic scenario should become a project-level `UC-*` versus when feature-level `SC-*` is enough.
- `REQ-03` State that machine-readable status/contracts and recovery/postconditions can be described as use cases when they are stable project-level behavior.
- `REQ-04` Verify that the existing `UC-XXX` template does not conflict with the new guidance.

### Non-Scope

- `NS-01` Do not add instantiated `UC-*` files for this template repository.
- `NS-02` Do not add downstream-specific source content, including source project names, tool names, vendor-specific workflows, CLI commands, or project-specific runtime assumptions.
- `NS-03` Do not change the `UC-XXX` template unless a concrete conflict with the new README guidance is found.
- `NS-04` Do not change feature-flow, epic-flow, PRD, ADR, product, domain, or ops rules.

### Constraints / Assumptions

- `ASM-01` GitHub issue #15 is the source task and defines the required scope and acceptance.
- `ASM-02` Downstream source documents named in issue #15 are evidence of the pattern class, not reusable generic content.
- `CON-01` This repository is a generic memory-bank template; downstream project-specific specialization must not be copied into `memory-bank/`.
- `CON-02` `memory-bank/use-cases/README.md` owns index-level guidance for when to create a use case; `memory-bank/flows/templates/use-case/UC-XXX.md` owns the canonical instantiated use case shape.

## Design Requirement Decision

| Decision | Reason | Downstream owner |
| --- | --- | --- |
| `Design required: no` | The feature changes Markdown guidance only. It does not alter API, schema, file format, runtime, security, integration, rollout/backout, or architecture boundaries; no solution-space trade-off owner is needed. | `none` |

## Verify

### Exit Criteria

- `EC-01` `memory-bank/use-cases/README.md` contains operational/agentic guidance that is reusable for any agentic or ops workflow.
- `EC-02` The README distinguishes project-level `UC-*` from feature-level `SC-*` for operational/agentic scenarios.
- `EC-03` The README explicitly allows machine-readable contracts/status, recovery behavior, and postconditions to be modeled as use cases when they are stable project-level behavior.
- `EC-04` The updated text contains no downstream-specific terms or commands from the source examples.
- `EC-05` The existing `UC-XXX` template remains compatible with the new README guidance.

### Traceability matrix

| Requirement ID | Problem refs | Acceptance refs | Checks | Evidence IDs |
| --- | --- | --- | --- | --- |
| `REQ-01` | `ASM-01`, `ASM-02`, `CON-01` | `EC-01`, `SC-01` | `CHK-01`, `CHK-03` | `EVID-01`, `EVID-03` |
| `REQ-02` | `ASM-01`, `CON-02` | `EC-02`, `SC-01`, `SC-02` | `CHK-01`, `CHK-03` | `EVID-01`, `EVID-03` |
| `REQ-03` | `ASM-01`, `ASM-02` | `EC-03`, `SC-01` | `CHK-01`, `CHK-03` | `EVID-01`, `EVID-03` |
| `REQ-04` | `CON-02`, `NS-03` | `EC-05`, `SC-03` | `CHK-03` | `EVID-03` |

### Acceptance Scenarios

- `SC-01` A reader evaluating repeated operational handoff, recovery, diagnostics, or status-reporting behavior can tell that it may become `UC-*` when it is stable project-level behavior.
- `SC-02` A reader evaluating a one-off delivery acceptance scenario can tell that it should remain `SC-*` in the relevant feature `brief.md`.
- `SC-03` A reviewer compares the new README section with the `UC-XXX` template and finds no conflicting ownership or required fields.

### Negative / Edge Cases

- `NEG-01` If the updated README mentions source-specific project names, tool names, command examples, or runtime details from the downstream examples, reject the change.
- `NEG-02` If the guidance implies that every operational check must become `UC-*` regardless of stability, repetition, or project-level ownership, reject the change.

### Checks

| Check ID | Covers | How to check | Expected result | Evidence path |
| --- | --- | --- | --- | --- |
| `CHK-01` | `EC-01`, `EC-02`, `EC-03`, `SC-01`, `SC-02` | `rg -n "Operational|agentic|machine-readable|recovery|postconditions|SC-\\*" memory-bank/use-cases/README.md` | Output shows the new generic guidance and UC-vs-SC routing. | `artifacts/ft-015/verify/chk-01.md` |
| `CHK-02` | `EC-04`, `NEG-01` | Run a forbidden-term `rg` check for the source project names, tool names, and command fragments listed in issue #15 against `memory-bank/use-cases/README.md`. | No matches. | `artifacts/ft-015/verify/chk-02.md` |
| `CHK-03` | `EC-05`, `SC-03`, `NEG-02` | Review `memory-bank/use-cases/README.md` and `memory-bank/flows/templates/use-case/UC-XXX.md`; then run `python3 scripts/check_memory_bank_index.py` and `git diff --check`. | README guidance does not conflict with the template, index audit passes, and diff check is clean. | `artifacts/ft-015/verify/chk-03.md` |

### Test matrix

| Check ID | Evidence IDs | Evidence path |
| --- | --- | --- |
| `CHK-01` | `EVID-01` | `artifacts/ft-015/verify/chk-01.md` |
| `CHK-02` | `EVID-02` | `artifacts/ft-015/verify/chk-02.md` |
| `CHK-03` | `EVID-03` | `artifacts/ft-015/verify/chk-03.md` |

### Evidence

- `EVID-01` Output proving the README contains the new generic operational/agentic guidance.
- `EVID-02` Output proving downstream-specific source terms are absent from the updated README.
- `EVID-03` Output or reviewer note proving the README, `UC-XXX` template, memory-bank index audit, and whitespace/conflict-marker checks are clean.

### Evidence contract

| Evidence ID | Artifact | Producer | Path contract | Reused by checks |
| --- | --- | --- | --- | --- |
| `EVID-01` | Command output or note for README guidance coverage | implementer | `artifacts/ft-015/verify/chk-01.md` | `CHK-01` |
| `EVID-02` | Command output for forbidden source-specific terms | implementer | `artifacts/ft-015/verify/chk-02.md` | `CHK-02` |
| `EVID-03` | Validation output and compatibility note | implementer / reviewer | `artifacts/ft-015/verify/chk-03.md` | `CHK-03` |

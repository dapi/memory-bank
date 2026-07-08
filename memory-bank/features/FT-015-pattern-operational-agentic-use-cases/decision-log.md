---
title: "FT-015: Decision Log"
doc_kind: feature
doc_function: decision_log
purpose: "Feature-local decision log for FPF-backed decisions that do not require ADRs."
derived_from:
  - brief.md
status: active
audience: humans_and_agents
---

# FT-015: Decision Log

This log records feature-local decisions closed from the current document set and
source issue. It does not create project-wide architecture decisions.

## Decisions

| Decision ID | Status | Question | Decision | Facts used | FPF reasoning | Consequences |
| --- | --- | --- | --- | --- | --- | --- |
| `DL-001` | accepted | Does FT-015 require `design.md`? | No. Keep `Design required: no` in `brief.md`. | Issue #15 asks for generic README guidance; `feature-flow.md` requires design for API/schema/runtime/contract/security/integration or solution trade-off changes; none are in scope. | A.7 Strict Distinction keeps problem/verify, solution, and execution owners separate. B.5 reasoning: the hypothesis "README-only guidance is enough" predicts no solution-space contract is needed; feature-flow gates confirm that prediction. | No `design.md` is created. `implementation-plan.md` may proceed after active `brief.md`. |
| `DL-002` | accepted | Should source examples create generic `UC-*` files or change the `UC-XXX` template? | No by default. Update only `memory-bank/use-cases/README.md` unless a concrete template conflict is found. | Issue #15 scope names README update and acceptance says `UC-XXX` template must not conflict. Current `UC-XXX` template already has trigger, preconditions, main flow, alternates, postconditions, business rules, and traceability. | A.10 Evidence Graph Referring treats source examples as carriers for pattern evidence, not as generic content to copy. A.7 separates the index/guidance owner (`README.md`) from the template shape owner (`UC-XXX.md`). | No instantiated `UC-*` files are added. Template remains unchanged unless `CHK-03` proves a conflict. |
| `DL-003` | accepted | Can machine-readable status/contracts, recovery, and postconditions be first-class use case content? | Yes, when they describe stable project-level operational behavior rather than one-off feature acceptance. | Issue #15 explicitly asks to fix this. Current README already says use cases may be operational and project-level, and `SC-*` remains feature-level acceptance. Source examples demonstrate repeated handoff, parallel delivery, and diagnostics flows. | A.13 models agency as a role played by a system in context, so agentic workflows can be described generically without creating a special "agent" document type. A.7 keeps actual execution and evidence carriers distinct, so machine-readable contracts/status can be observable parts of the scenario rather than implementation details. | README guidance should mention machine-readable contracts/status, recovery behavior, and postconditions while preserving the `UC-*` vs `SC-*` boundary. |

## Human Gates

None currently open.

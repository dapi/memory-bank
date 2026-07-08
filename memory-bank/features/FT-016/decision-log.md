---
title: "FT-016: Decision Log"
doc_kind: feature-support
doc_function: decision_log
purpose: "Feature-local decision log for FPF-backed decisions made while preparing and reviewing FT-016 documents."
derived_from:
  - brief.md
status: active
audience: humans_and_agents
must_not_define:
  - implementation_sequence
  - global_architecture_policy_without_adr
  - ft_016_scope
  - ft_016_acceptance_criteria
---

# FT-016: Decision Log

## FPF Reading Rule

Record facts, assumptions, reasoning and consequences separately. Use this log only for feature-local decisions that do not require global ADR.

## Relationship To Canonical Owners

`decision-log.md` records reasoning for feature-local decisions. `brief.md` remains the canonical owner for scope and acceptance. `design.md` remains the canonical owner for selected solution facts.

## DL-01: Design Layer Required

**Date:** 2026-07-08

**Question:**

Does FT-016 need `design.md`, or can it proceed from `brief.md` directly to `implementation-plan.md`?

**Status:** Resolved.

**Facts:**

- Issue 16 asks to add a generic support template and clarify when it is needed in feature/design/implementation-plan templates.
- `feature-flow.md` requires `design.md` when a feature changes file format, governed contracts, support boundaries, or would otherwise force `implementation-plan.md` to make solution decisions.
- FT-016 must decide selected template shape, owner boundaries and generic-vs-downstream separation before execution sequencing.

**Reasoning:**

FPF Strict Distinction separates problem facts, solution decisions and work occurrence. Without `design.md`, the implementation plan would have to choose the template role, naming, local IDs and ownership constraints, which would mix solution-space decisions into execution planning.

**Decision:**

FT-016 requires `design.md`.

**Consequences:**

- `brief.md` records `Design required: yes`.
- `design.md` owns `SOL-*`, `SD-*`, contracts, invariants and failure modes.
- `implementation-plan.md` must derive from both `brief.md` and `design.md`.

## DL-02: Single Privacy / Source Boundary Support Template

**Date:** 2026-07-08

**Question:**

Should FT-016 create one combined support template or split privacy boundary and source inventory into separate templates?

**Status:** Resolved.

**Facts:**

- Issue 16 asks for a generic support template for "privacy/source boundary".
- Issue acceptance requires both privacy exclusion behavior and source confidence/status.
- The referenced downstream examples have separate privacy boundary and source inventory docs, but issue 16 does not ask to preserve that split.
- Feature-flow support docs should be created only when they reduce real ambiguity; extra templates add lifecycle/navigation overhead.

**Reasoning:**

FPF Bounded Context makes the support doc boundary explicit: the relevant context is "what sources may be read/stored/cited for this feature, with what confidence and owner boundary." Privacy boundary and source inventory are two aspects of that same support context. FPF Evidence Graph also ties source inventory and evidence permissions together: a source claim needs carriers/confidence, but privacy rules constrain which carriers can be used.

**Decision:**

Create one support template named `privacy-source-boundary.md`.

**Consequences:**

- The template must keep allowed metadata, excluded data, source inventory, confidence/status and evidence rules in one artifact.
- It must still separate privacy permission from source confidence so confidence is never treated as access authorization.

## DL-03: Support Doc Function Is Reference, Not Evidence

**Date:** 2026-07-08

**Question:**

Should instantiated privacy/source-boundary support docs default to `doc_function: evidence` or `doc_function: reference`?

**Status:** Resolved.

**Facts:**

- Existing generic support templates use `doc_function: reference`.
- The downstream source examples use `doc_function: evidence`, but they are project-specific artifacts.
- Issue 16 says the support doc must not replace `brief.md` / `design.md`.
- The new template will describe source boundaries and evidence expectations; it is not itself proof that checks passed.

**Reasoning:**

FPF Evidence Graph distinguishes a knowledge artifact from evidence carriers. A support doc can inventory admissible source carriers and evidence constraints, but the actual check outputs remain evidence carriers produced during verification.

**Decision:**

Use `doc_function: reference` for instantiated privacy/source-boundary support docs.

**Consequences:**

- The template may include a Test Evidence section, but it must route concrete check evidence to `brief.md` `EVID-*` or plan-local evidence.
- The support doc stays a companion/reference artifact and not an execution evidence replacement.

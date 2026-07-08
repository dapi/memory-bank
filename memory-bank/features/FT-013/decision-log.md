---
title: "FT-013: Decision Log"
doc_kind: feature-support
doc_function: decision_log
purpose: "Feature-local decision log for FPF-backed decisions made while preparing and reviewing FT-013 documents."
derived_from:
  - brief.md
  - ../../flows/feature-flow.md
status: active
audience: humans_and_agents
must_not_define:
  - ft_013_scope
  - ft_013_acceptance_criteria
  - implementation_sequence
---
# FT-013: Decision Log

## DL-FT013-001: Require `design.md` for this docs-only feature

**Date:** 2026-07-08

**Question:** Does FT-013 need a separate solution-space owner even though it changes only documentation?

**Available facts:**

- Issue 13 requires three reusable workflow documents and links to `flows/workflows.md` plus task-flow issue.
- `feature-flow.md` requires `design.md` when a feature needs alternatives/trade-off reasoning or otherwise would force `implementation-plan.md` to make design decisions.
- FT-013 must decide document locations, optional-vs-canonical semantics, source adaptation rules and how to handle absent task-flow docs.

**FPF reasoning:** By Bounded Context, the reusable flow docs, feature package and GitHub issues are separate semantic frames; their links must be explicit. By Strict Distinction, `brief.md` owns problem facts and `implementation-plan.md` owns execution sequencing, so solution choices need their own owner.

**Decision:** `Design required: yes`; create `design.md` as the owner of selected solution, alternatives, trade-offs, contracts and invariants.

**Consequences:**

- `brief.md` remains problem/verify owner.
- `implementation-plan.md` can stay execution-only.
- Solution choices are traceable via `SOL-*`, `SD-*`, `CTR-*`, `INV-*` and `FM-*`.

## DL-FT013-002: Treat issue 12 as adjacent dependency, not FT-013 scope

**Date:** 2026-07-08

**Question:** How should FT-013 "link with task-flow issue" when `task-flow.md`, `bugfix-flow.md` and `refactor-flow.md` do not exist in this repo yet?

**Available facts:**

- Issue 13 scope says to connect workflow metrics/decision log/developer brief with `flows/workflows.md` and task-flow issue.
- GitHub issue 12 is open and owns `task-flow.md`, `bugfix-flow.md`, `refactor-flow.md`, task templates and `memory-bank/tasks/README.md`.
- Current repository search shows no local `task-flow.md`, `bugfix-flow.md` or `refactor-flow.md`.
- `check_memory_bank_index.py` fails broken internal links.

**FPF reasoning:** Bounded Context keeps FT-013 and issue 12 as separate delivery contexts. Strict Distinction prevents a dependency reference from becoming implementation work. Evidence Graph discipline requires actual local carriers before local links can be treated as valid evidence.

**Decision:** FT-013 records issue 12 as an adjacent dependency and non-scope boundary. Generic docs must not link to absent local task-flow files; they may describe compact profiles generically and be updated after issue 12 lands.

**Consequences:**

- No `task-flow.md`, `bugfix-flow.md`, `refactor-flow.md`, task templates or `memory-bank/tasks/` are created in FT-013.
- Link audit remains enforceable.
- A future issue/PR may add concrete task-flow links after issue 12 is implemented.

## DL-FT013-003: Safety metrics outrank speed metrics

**Date:** 2026-07-08

**Question:** How should workflow metrics decide whether compact profiles succeeded?

**Available facts:**

- Issue 13 acceptance says metrics must measure safety before speed.
- Source pattern includes routing coverage, safety misroute, missing evidence, rework, traceability and lead-time style measurements.
- `testing-policy.md` requires evidence and regression/safety coverage where deterministic verification is realistic.

**FPF reasoning:** Evidence Graph says claims need a carrier chain; Trust & Assurance says assurance is capped by weakest supported evidence and integration fit. Therefore faster workflow throughput cannot compensate for missing evidence, wrong routing or unsafe profile selection.

**Decision:** `workflow-metrics.md` must order decision rules so safety/evidence/rework/traceability gates are evaluated before speed/lead-time. Speed is a secondary observation and cannot produce success when safety gates fail.

**Consequences:**

- `NEG-01` becomes canonical in `brief.md`.
- `INV-02` and `SD-02` require safety-first metric semantics.
- Implementation must avoid wording that presents speed improvement as sufficient success.

## DL-FT013-004: Generic adaptation, not source import

**Date:** 2026-07-08

**Question:** How much of the `alfagen/mercury` source pattern can be reused in generic memory-bank docs?

**Available facts:**

- Issue 13 cites three source docs as useful pattern evidence.
- Issue 13 explicitly says not to transfer project-specific dates, project names or operational details.
- Source docs include concrete pilot dates, source repo/project names and local operational commands/details.

**FPF reasoning:** Evidence Graph allows source carriers to support a claim about useful structure, but Bounded Context forbids silently moving local invariants and operational details into the generic template context. Strict Distinction separates evidence source from reusable generic content.

**Decision:** Reuse only the generic structure and concepts: decision entries, metric cards, safety-first decision rule and developer brief shape. Replace project-specific dates, names, commands and local assumptions with generic placeholders or configurable fields.

**Consequences:**

- `CHK-03` targets source leakage in delivered workflow docs.
- `CTR-04` makes generic adaptation a solution contract.
- Reviewer acceptance must include a semantic check, not only a string scan.

## DL-FT013-005: Keep decision-log provenance acyclic

**Date:** 2026-07-08

**Question:** Should `design.md` derive from `decision-log.md`, and should `decision-log.md` derive from `design.md`?

**Available facts:**

- The first package draft linked `design.md` and `decision-log.md` through reciprocal `derived_from` entries.
- `feature-flow.md` makes `design.md` the solution-space owner and allows support docs only as aids that do not replace canonical owners.
- `decision-log.md` records feature-local reasoning used to support owner documents; it is not itself the selected solution owner.

**FPF reasoning:** Evidence Graph discipline favors acyclic provenance chains. Strict Distinction separates the canonical solution episteme (`design.md`) from its support/provenance carrier (`decision-log.md`). A support log can be linked from README and cited in prose without becoming an upstream dependency of the canonical design frontmatter.

**Decision:** Remove the reciprocal `derived_from` edge. `design.md` derives from `brief.md`; `decision-log.md` derives from `brief.md` and `feature-flow.md`.

**Consequences:**

- The package keeps decision provenance visible through README and body links.
- Frontmatter dependency graph remains acyclic and owner boundaries are clearer.
- Future review-improve entries can be appended to `decision-log.md` without changing canonical design dependencies unless they alter selected solution facts.

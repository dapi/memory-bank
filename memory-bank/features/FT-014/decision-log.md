---
title: "FT-014: Decision Log"
doc_kind: feature-support
doc_function: reference
purpose: "Feature-local ledger for FPF-backed decisions made while adding lightweight epic intake brief template."
derived_from:
  - brief.md
  - design.md
status: active
audience: humans_and_agents
must_not_define:
  - ft_014_scope
  - ft_014_acceptance_criteria
  - ft_014_execution_sequence
---

# FT-014: Decision Log

## Decision Method

Decisions use FPF in plain language:

- Bounded Contexts: keep intake, full epic governance and feature execution as separate semantic frames.
- Strict Distinction: do not let a proposal document become a roadmap, accepted subissue registry, risk register or execution plan.
- Evidence Graph: tie decisions to issue 14, the source template and current repository governance docs.
- Canonical Reasoning Cycle: propose a local hypothesis, derive consequences, then verify against repository checks and document review.

## Decisions

| Decision ID | Question Closed | Available facts | FPF reasoning | Decision | Consequences |
| --- | --- | --- | --- | --- | --- |
| `DL-01` | Does the new epic `brief.md` replace full epic owners? | Issue 14 says the current epic package has `charter/roadmap/risks/subissues`; acceptance says full epic package remains authoritative for roadmap/subissues/risks; `epic-flow.md` already gives separate owners for these docs. | Bounded Contexts separates early intake from full epic governance; Strict Distinction prevents one document from owning roadmap, risk and execution semantics. | The brief is optional intake only and does not replace `charter.md`, `roadmap.md`, `subissues.md`, `risks.md` or `decision-log.md`. | Reflected as `SD-01`; update `epic-flow.md` package rules, layer model and boundary rules. |
| `DL-02` | How much of the downstream source template may be reused? | Issue 14 names `dapi/zelma: memory-bank/flows/templates/epic/brief.md` as source and says not to transfer `zelma`-specific content; current repo uses governed wrapper templates with frontmatter and embedded instantiated contract. | Evidence Graph requires claims to stay anchored to source carriers; Strict Distinction separates generic template semantics from downstream project content. | Reuse only generic shape: problem, outcome, rough scope, non-scope, candidate feature brief links and readiness notes; normalize wording/frontmatter to this repository's template style. | Reflected as `SOL-01`, `SOL-04`, `INV-01`; verify with targeted source-leakage grep. |
| `DL-03` | Does this docs-only change need C4 artifact or ADR? | `feature-flow.md` requires design when governance/template contracts need explicit reasoning; C4 is unnecessary when no runtime/API/storage/integration/security boundary changes; issue scope is Markdown governance/template docs. | Strict Distinction separates solution reasoning from architecture model; Evidence Graph says C4/ADR must be evidence-backed, not ceremonial. | Create `design.md` for local solution decisions, set `C4-00: not required`, and do not create ADR. | Reflected as `C4-00`, `SD-03`; implementation plan can proceed after active `design.md`. |

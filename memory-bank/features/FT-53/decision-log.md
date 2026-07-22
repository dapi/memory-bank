---
title: "FT-53: Decision Log"
doc_kind: feature
doc_function: decision_log
purpose: "Feature-local log of evidence-backed FT-53 decisions; it does not own requirements, solution space or execution sequencing."
derived_from:
  - brief.md
status: active
audience: humans_and_agents
must_not_define:
  - ft_53_problem_space
  - ft_53_selected_design
  - ft_53_execution_sequence
---

# FT-53: Decision Log

## Decision records

### DEC-53-01: Design owner not required

- **Status:** accepted
- **Decision:** Do not create `design.md` for FT-53.
- **Question:** Does the requested brownfield protocol require a separate
  solution-space owner before an implementation plan can be written?
- **Evidence:** Issue #53 requests generic documentation: a protocol, links,
  lifecycle guidance, inventory, safety rules and validation. Its acceptance
  criteria do not request a runtime, API, event, schema, configuration,
  deployment or architectural change. `brief.md` constrains the feature to
  documentation and excludes those changes in `NS-53-02`.
- **FPF reasoning:** The problem-space owner (`brief.md`) owns what must be
  documented and verified. A solution-space owner is required only if the
  package must decide an architecture, contract, invariant, failure mode or
  rollout semantics. The evidence places none of those objects in this
  feature's bounded context; adding `design.md` would create an empty owner and
  blur the problem/execution boundary. Therefore the smallest coherent package
  is brief + decision log + plan.
- **Alternatives considered:** (1) create `design.md` with the protocol outline;
  rejected because that duplicates requirements as solution facts. (2) record
  the outline only in the plan; rejected because the plan must not own scope.
- **Consequences:** `implementation-plan.md` may sequence documentation edits,
  but must promote any newly discovered runtime/contract or architectural
  decision to the appropriate owner before planning it.

### DEC-53-02: Intake PRD stays outside Memory Bank until upstream adaptation

- **Status:** accepted
- **Decision:** The protocol will require a temporary intake PRD outside
  `memory-bank/` during pre-adaptation discovery, then govern/import it only
  after the relevant baseline owners are adapted.
- **Question:** Where can discovery facts be safely recorded before Memory Bank
  is allowed to influence discovery?
- **Evidence:** Issue #53 explicitly prescribes an evidence-backed intake PRD
  outside `memory-bank/`, followed by conversion into `memory-bank/prd/` whose
  `derived_from` links point only to already adapted upstream owners. It also
  explicitly prohibits consulting `memory-bank/` during pre-adaptation
  discovery.
- **FPF reasoning:** Treat pre-adaptation discovery and governed Memory Bank as
  different bounded contexts. A temporary artifact can preserve evidence and
  uncertainty without becoming an authority of the not-yet-installed governance
  system. After owners exist, conversion creates the governed dependency graph;
  collapsing the stages would let generic placeholders contaminate project
  knowledge.
- **Consequences:** The protocol must specify a clearly named non-`memory-bank/`
  location and preserve source references, confidence, conflicts and open
  questions through conversion.

### DEC-53-03: Default intake location is the repository root

- **Status:** accepted
- **Decision:** The protocol's temporary intake PRD path is
  `./brownfield-intake-prd.md`.
- **Question:** What exact location can the generic protocol prescribe without
  assuming that a downstream repository already has a `docs/` directory or any
  governed Memory Bank structure?
- **Evidence:** Issue #53 requires the protocol to define an intake PRD
  location outside `memory-bank/`; pre-adaptation discovery occurs before
  Memory Bank is consulted or installed. The issue does not establish any
  repository-specific documentation directory.
- **FPF reasoning:** The pre-adaptation and governed contexts must remain
  separate (`DEC-53-02`). The repository root is the only path boundary present
  in every repository without importing an unproven convention; the selected
  filename is explicit and cannot be mistaken for an artifact inside
  `memory-bank/`. This supplies a deterministic hand-off while making no claim
  about downstream documentation layout.
- **Alternatives considered:** `docs/brownfield-intake-prd.md` is rejected: a
  `docs/` directory is not guaranteed by the evidence. An unspecified
  non-`memory-bank/` path is rejected: it fails the issue's explicit location
  requirement and weakens the hand-off.
- **Consequences:** The protocol will use this default path. A repository may
  retain an existing local convention only when it records the chosen alternate
  path and preserves all required intake fields and the outside-`memory-bank/`
  boundary.

---
title: Project-local Memory Bank Bootstrap
doc_kind: project
doc_function: reference
purpose: Фиксирует происхождение, границы и решения начальной адаптации project-local Memory Bank.
derived_from:
  - README.md
  - ../docs/adoption.md
  - ../docs/ownership.md
status: active
audience: humans_and_agents
---

# Project-local Memory Bank Bootstrap

## Identity And Provenance

This is the project-local Memory Bank for the `dapi/memory-bank` source
repository. The generic upstream payload remains at
[`../template/memory-bank/`](../template/memory-bank/); it is not the location
for instantiated project artifacts.

The installation lock at [`.lock`](.lock) records the source template version
`db55624` and immutable source ref
`db55624d7119b0d13596b152671db414c7fec733`. Use the lock and the ownership
rules in [`../docs/ownership.md`](../docs/ownership.md) when updating this copy.

## Bootstrap Decisions

- This root README is the project-local entry point and explicitly identifies
  `dapi/memory-bank`; the generic template README remains unchanged.
- [`features/`](features/README.md) is the canonical destination for this
  repository's `FT-XXX/` delivery packages, including the future `FT-068/`
  package. No instantiated feature package belongs in `template/memory-bank/`.
- The CLI-managed governance and flow documents are retained from the source
  template. Future project-specific product, domain, engineering, and
  operations facts are adapted in this directory as evidence becomes
  available; they must not be copied back to the generic payload.
- The managed routing block in [`../AGENTS.md`](../AGENTS.md) directs agents to
  this project-local entry point, DNA, and task routing flow.

## Verification Scope

The bootstrap change is documentation-only. Its required checks are project
Memory Bank navigation and adoption diagnostics: `memory-bank-cli lint` and
`memory-bank-cli doctor`. The source-template profile remains separately
applicable to `template/memory-bank/`.

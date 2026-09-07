---
title: Memory Bank Product Context
doc_kind: product
doc_function: canonical
purpose: Product context of the Memory Bank template and its adoption experience.
derived_from:
  - ../dna/governance.md
  - ../../README.md
  - "https://github.com/dapi/memory-bank/issues/141"
status: active
audience: humans_and_agents
canonical_for:
  - project_product_context
---

# Memory Bank Product Context

Memory Bank supplies a version-controlled documentation and delivery template for software
repositories. Its users are repository owners, document authors and coding agents. The
product keeps project context, document ownership and decision rationale available across
working sessions; the companion memory-bank-cli installs and validates the payload.

The accepted direction in issue 141 is gradual adoption: users can start with governance,
add project document contracts, and connect AI delivery processes later. These are target
capabilities of the current initiative, not a claim that they are already released.

## Core Product Workflows

- [UC-001](../use-cases/UC-001-adopt-documentation-and-flows.md) — choose documentation depth,
  preserve authored content, and explicitly connect documents to processes.
- Existing governed delivery starts from task routing after Flows has been adopted.

## Outcomes and constraints

Users can keep their documents useful independently of an AI executor. Adding processes
preserves document ownership and makes additional obligations explicit. Updates must not
silently reduce pinned adoption checks or implicitly change an installation's selected depth.
Unadopted base documents use the current installed DNA/type rules; their checks may evolve
with an explicitly requested template pull. Only adoption and legacy migration pin a bundle.

Template source and CLI implementation retain separate owners. This repository supplies the
generic payload; project-specific content remains in downstream repositories. Acceptance for
the component initiative is tracked by EP-141 and FT-141 rather than by invented product KPIs.

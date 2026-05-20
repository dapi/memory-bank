---
title: Merchantly Forms
doc_kind: engineering
doc_function: canonical
purpose: Merchantly-specific form conventions for operator and customer-facing UI.
derived_from:
  - README.md
status: active
audience: humans_and_agents
---

# Merchantly Forms

Use existing project form patterns before adding new layout or validation behavior. Feature packages may add screen-specific detail, but shared form rules belong here.

## Rules

- Preserve clear label, input, help text and error ordering.
- Keep validation messages close to the field they explain.
- Do not introduce a new form abstraction without an ADR or explicit feature-level decision.

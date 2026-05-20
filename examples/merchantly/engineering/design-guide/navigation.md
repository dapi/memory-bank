---
title: Merchantly Navigation
doc_kind: engineering
doc_function: canonical
purpose: Merchantly-specific navigation and operator entrypoint conventions.
derived_from:
  - README.md
status: active
audience: humans_and_agents
---

# Merchantly Navigation

Navigation decisions should preserve discoverable operator workflows without turning feature-local tools into global sections prematurely.

## Rules

- Add a dedicated entrypoint only when the feature introduces a repeated operator workflow.
- Keep fixed catalogs explicit when users must choose from system-owned keys.
- Avoid hiding required setup behind unrelated content sections.

---
title: Merchantly Domain Documentation Index
doc_kind: domain
doc_function: index
purpose: Merchantly-specific domain overlay: glossary, model, rules, states, events and context map live here when instantiated for the example project.
derived_from:
  - ../../../memory-bank/domain/README.md
status: active
audience: humans_and_agents
---

# Merchantly Domain Documentation Index

Этот каталог показывает, где в project-specific memory-bank должен жить Merchantly domain layer. Shared structure и правила см. в [`../../../memory-bank/domain/README.md`](../../../memory-bank/domain/README.md).

## Current Domain Concepts In Examples

- `Vendor` — merchant/account boundary for store-specific settings and public pages.
- `Order` — customer-facing order lifecycle and public order page context.
- `Payment` — payment provider flow, result surfaces and fiscal receipt obligations.
- `VendorPageTemplate` — feature-local public page template concept introduced by FT-4564.

## Boundary

- Product motivation, customers and success metrics belong to [`../product/README.md`](../product/README.md).
- Runtime architecture, frontend implementation and design-guide rules belong to [`../engineering/README.md`](../engineering/README.md).
- Feature-local accepted decisions remain in [`../features/README.md`](../features/README.md) until they become shared domain rules.

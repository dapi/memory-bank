---
title: UI Design Guide
doc_kind: engineering
doc_function: index
purpose: Optional template for documenting an existing project UI kit: components, forms, actions, tables, navigation, states, screenshots, source paths and agent instructions.
derived_from:
  - ../../dna/governance.md
  - ../frontend.md
status: active
audience: humans_and_agents
must_not_define:
  - product_requirements
  - domain_rules
  - frontend_architecture_contract
  - implementation_sequence
---

# UI Design Guide

This optional guide helps downstream projects document an existing UI kit or admin/operator interface patterns. Use it when agents need a fast source of truth for concrete components, helper APIs, screenshots, local interaction conventions and source code paths.

If the project has no reusable UI kit or separate frontend/admin UI, leave this section minimal or remove it during adaptation. The canonical frontend engineering contract stays in [`../frontend.md`](../frontend.md); this guide records concrete local UI reference material.

## Owner Boundaries

| Layer | Owns | Does not own |
| --- | --- | --- |
| `../frontend.md` | UI surfaces, frontend stack, component boundaries, design system integration and i18n contract | Detailed component catalog, screenshots and local examples |
| `ui-design-guide/README.md` | Existing UI kit reference, local component usage, helper APIs, screenshots and agent-facing UI instructions | Product requirements, domain rules, architecture decisions or implementation sequencing |
| `../architecture.md` / ADR | Architecture and reusable engineering decisions | Component screenshots or per-screen local conventions |

## How To Adapt

Replace placeholder rows with project-specific facts after copying the template into a downstream repository.

- Prefer links to canonical source files over prose descriptions.
- Keep screenshots versioned or linkable when they are required for agent review.
- Document current local patterns before inventing new ones.
- If a UI rule is product-specific, link to the product or feature owner instead of redefining it here.
- If a UI rule changes frontend architecture, update [`../frontend.md`](../frontend.md) or an ADR first.

## UI Surfaces

| Surface ID | Surface | Users / roles | Source paths | Notes |
| --- | --- | --- | --- | --- |
| `UI-SURF-01` | Example: admin dashboard | Example: operators | `path/to/ui` | Replace with local surface facts |

## Components

| Component ID | Component / pattern | When to use | Source paths | Examples / screenshots |
| --- | --- | --- | --- | --- |
| `UI-CMP-01` | Example: reusable panel, modal, badge or field wrapper | Describe the local usage rule | `path/to/component` | `path/to/screenshot-or-example` |

## Forms

| Form pattern ID | Pattern | Validation / errors | Helper APIs | Source paths |
| --- | --- | --- | --- | --- |
| `UI-FORM-01` | Example: create/edit form | Describe how required, invalid and async states are shown | `helper_or_component_name` | `path/to/form` |

## Buttons And Actions

| Action ID | Action type | Visual / placement rule | State rule | Source paths |
| --- | --- | --- | --- | --- |
| `UI-ACT-01` | Example: primary action, destructive action, bulk action | Describe local placement and priority | Describe disabled/loading/confirmation behavior | `path/to/action` |

## Tables

| Table ID | Table pattern | Sorting / filtering / pagination | Empty / loading / error states | Source paths |
| --- | --- | --- | --- | --- |
| `UI-TBL-01` | Example: index table | Describe local data controls | Describe state rendering | `path/to/table` |

## Navigation

| Navigation ID | Pattern | Applies to | Source paths | Notes |
| --- | --- | --- | --- | --- |
| `UI-NAV-01` | Example: sidebar, tabs, breadcrumbs or contextual links | Describe surfaces or roles | `path/to/navigation` | Replace with local routing conventions |

## States And Labels

Document canonical UI labels and state rendering only when they are stable local conventions. Link to domain or product owners for business meaning.

| State / label ID | UI state or label | Meaning in UI | Owner / source | Examples |
| --- | --- | --- | --- | --- |
| `UI-LBL-01` | Example: visible status label | Describe display semantics | `path/to/source` | `path/to/example` |

## Screenshots And Examples

| Screenshot ID | What it shows | Surface / component refs | Path or link | Refresh rule |
| --- | --- | --- | --- | --- |
| `UI-SHOT-01` | Example: filled form state | `UI-FORM-01` | `path/to/screenshot` | Update when local UI pattern changes |

## Source Code Paths

| Path ID | Path | Contains | Use when | Notes |
| --- | --- | --- | --- | --- |
| `UI-PATH-01` | `path/to/ui` | Local UI components, helpers or views | Agent needs concrete examples before changing UI | Replace with project paths |

## Agent Instructions

- Start with [`../frontend.md`](../frontend.md) to understand stack and ownership boundaries.
- Then inspect the relevant source paths in this guide before changing UI.
- Reuse existing local components, helpers and state patterns before adding new ones.
- Do not infer business meaning from visual labels; follow linked domain/product owners.
- Do not treat placeholder examples in this guide as framework defaults.
- If local examples conflict with canonical requirements, update the owner document before changing implementation.

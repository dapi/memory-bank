---
title: "UC-001: Adopt documentation and flows"
doc_kind: use_case
doc_function: canonical
purpose: Stable user behavior for choosing Memory Bank depth and explicitly adopting processes.
derived_from:
  - ../flows/use-case.md
  - ../product/context.md
  - "https://github.com/dapi/memory-bank/issues/141"
status: active
audience: humans_and_agents
must_not_define:
  - implementation_sequence
  - architecture_decision
  - feature_level_test_matrix
---

# UC-001: Adopt documentation and flows

## Goal

Maintain useful project documentation at the chosen depth, then add process obligations
without losing authored content or silently changing existing document checks. This is
accepted target behavior; release/delivery status belongs to the implementing feature.

## Primary Actor

The repository owner selects the installation depth; a document author creates and maintains
base documents or explicitly connects them to a flow. An automated caller has the same
observable choice and failure rules as an interactive caller.

## Trigger

A project adopts Memory Bank, adds supported components to its chosen depth, updates the template, or connects,
transitions or moves a document governed by an explicit flow contract.

## Preconditions

- The actor has authority to change the named repository and its documentation.
- The chosen source and installed CLI support the requested operation.
- Existing installation/adoption state is readable and consistent, or its conflicts are
  resolved through a supported explicit operation before applying changes.

## Main Flow

1. The owner chooses governance only, governance plus document contracts, or the full set
   with flows; executor adapters are a separate explicit choice for these three depths.
   A fresh installation with no selection uses the documented legacy compatibility default,
   including the previously bundled adapters. This default never opts an existing installation
   into a breaking migration.
2. The system shows the resulting selection and preserves the surrounding project content.
3. Authors create base documents. When a flow is wanted, the author explicitly selects its
   document contract and reviews the operation's applicable obligations.
4. The system checks the prospective result and records the selected obligations together
   with document identity and installation state.
5. Later updates preserve the chosen depth and the rules pinned for already-adopted documents.

## Alternate Flows / Exceptions

- ALT-01: Add Flows to an existing documentation installation; existing base documents keep
  their base status until an explicit adoption operation.
- ALT-02: An existing legacy installation previews a separate breaking migration and applies
  it only after explicit opt-in. Without opt-in it remains on its pinned legacy path.
- ALT-03: Explicitly transition or move an adopted document through a supported operation,
  preserving its applicable checks and history.
- EX-03: Component removal, downgrade and context-changing moves are unsupported and
  rejected before mutation.
- EX-01: Unknown format, unsafe path, unresolved ownership/adoption drift or an unsupported
  transition causes a diagnostic and no partially applied change.
- EX-02: A failure during mutation restores the old consistent state or reports an explicit
  recovery condition; the system never reports a partial result as successful completion.

## Postconditions

Successful operations preserve authored content and keep selection, identity and obligations
consistent. A failed preflight leaves the old repository unchanged. Existing invalid legacy
documents retain their prior invalid verdict under the explicit compatibility migration;
new errors cannot be accepted as part of that allowance.

## Business Rules

- BR-01: Explicit core/docs/full selections add adapters only when selected. Fresh no-selection
  installation uses the legacy compatibility default; ordinary updates preserve the resolved set.
- BR-02: Project documents belong to the project, including after template updates.
- BR-03: Installing Flows alone does not assign gates to existing base documents.
- BR-04: Removing or editing a projection field cannot erase recorded adoption.
- BR-05: Existing pinned contract rules change only through an explicit supported transition.
- BR-06: Legacy migration requires separate consent and preserves compatibility obligations.
- BR-07: Every operation respects repository path and transaction boundaries.

## Traceability

| Upstream / downstream | Reference |
| --- | --- |
| PRD | none; issue 141 and EP-141 already own the initiative scope |
| Feature | [FT-141](../features/FT-141/brief.md) and its external CLI delivery owner |
| ADR | [ADR-002](../adr/ADR-002-component-document-contracts.md), candidate design |

## Downstream Behavior Coverage

| UC element | Downstream examples | Coverage note |
| --- | --- | --- |
| BR-01/02/03, ALT-01 | FT-141 SC-01/02/03/07 | Selection, authored content and explicit activation |
| BR-04/05, ALT-03 | FT-141 SC-04/05/10/11, NEG-01/02/04/05 | Identity, frozen rules and failure recovery |
| BR-06, ALT-02 | FT-141 SC-05/06/08/12, NEG-03/05 | Compatibility entrypoint and migration consent |
| BR-07, EX-01 | FT-141 SC-09, NEG-06 | Path confinement |
| EX-02 | FT-141 SC-04, NEG-04 | Mutation rollback and explicit recovery-required outcome |

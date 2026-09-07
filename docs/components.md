# CTR-01: Component installation overview

[Issue 141](https://github.com/dapi/memory-bank/issues/141) enables gradual adoption of
Memory Bank. DNA provides standalone governance, Documents adds document types and templates,
and optional Flows adds explicit process obligations. Executor adapters are selected separately,
with the legacy compatibility default described by the contract.

The template owns the declarations; memory-bank-cli owns their interpretation and atomic
installation. CLI pull updates a template; CLI update updates the executable.

The sole normative owner is [CTR-01 behavior and wire format](component-wire-format.md).
Use its sections for:

- [Source inventory](component-wire-format.md#source-envelope-and-inventory): format/version gates and exhaustive payload membership.
- [Selection and navigation](component-wire-format.md#selection-and-navigation): presets, adapters, ownership and retained paths.
- [Rule bundles](component-wire-format.md#rule-and-bundle-documents): base documents and frozen adopted checks.
- [Installation lock](component-wire-format.md#installation-lock): persisted composition and generated-file boundaries.
- [Adoption and history](component-wire-format.md#adoption-and-history): identity, projections, selectors and transitions.
- [Migration preview](component-wire-format.md#migration-resolution-and-preview): commands, write intents and approval digests.
- [Legacy baseline](component-wire-format.md#legacy-classification-and-creation-baseline): source-specific classification and explicit legacy-flow creation.
- [Validation and rollout](component-wire-format.md#validation-transactions-and-compatibility-entrypoint): transaction failures, source projection and the bridge-first entrypoint.

This overview is navigation, not a second definition of the protocol. Acceptance and delivery
status belong to [FT-141](../memory-bank/features/FT-141/brief.md) and the external
[CLI #62](https://github.com/dapi/memory-bank-cli/issues/62). Independent review uses code-converge.

---
title: "Component adoption delivery evidence"
doc_kind: evidence
doc_function: evidence
purpose: "Trace issue 141 acceptance to executable checks and independent reviews."
derived_from:
  - ../memory-bank/features/FT-141/brief.md
  - ../memory-bank/features/FT-141/design.md
status: active
---

# Component adoption delivery evidence

The implementation is ready for review in [template PR 143](https://github.com/dapi/memory-bank/pull/143)
and [CLI PR 64](https://github.com/dapi/memory-bank-cli/pull/64), stacked on
[bridge PR 63](https://github.com/dapi/memory-bank-cli/pull/63). Merge, release publication,
live downstream migration and human closure of the umbrella initiative remain outside this delivery.

## Acceptance

The [CLI fixture suite](https://github.com/dapi/memory-bank-cli/tree/caf0f3eaf3af290a702c8553795168584ac8b987/internal/ownership)
and [actual-binary matrix](https://github.com/dapi/memory-bank-cli/blob/caf0f3eaf3af290a702c8553795168584ac8b987/scripts/e2e-components.py)
provide the executable checks below. EVID numbers retain the brief's CHK mapping.

| Evidence | Executable checks and result |
| --- | --- |
| EVID-01/07 | ComponentPayloadMatrix and ComponentAdapterMatrix: core/docs/full/legacy, all adapters, flagless no-op and independent navigation pass. |
| EVID-02 | ComponentResolutionPlanBindsPermissions checks docs→full preservation. A separate actual-binary filled ADR upgrade preserved exact author bytes, default user ownership and absence of adoption; doctor passed. |
| EVID-03 | ComponentBaseTypeMatrixRemainsUnadopted covers all six types; ComponentDocumentsLifecycle and ComponentAtomicLegacyFlowCreation cover explicit adoption and prepared drafts with relocated links. |
| EVID-04/09 | ComponentIntegrityRejectsTamperingWithoutMutation, contract/path fixtures and recovery fixtures reject drift, unsafe paths and incomplete restoration without accepting partial state. |
| EVID-05 | ComponentLegacyMigration preserves the exact legacy finding multiset for valid and invalid inputs; new base documents do not join snapshots. Ambiguous migration requires a complete explicit resolution. |
| EVID-06/08 | Actual bridge/pre-bridge binaries are rejected by the new entrypoint before installer invocation. Source-format E2E and an actual 8e7f3fd→f1f04de legacy upgrade preserve the supported route. |
| EVID-10/11 | Selector transition requires evidence, creates exactly one exclusion/record and rolls back on injected failure. Move preserves bytes and identity; exact retry is a no-op. |
| EVID-12 | Migration and resolution previews bind observed bytes, Git modes, exact permissions, directory states, selection and write intents; stale approval inputs reject. |

Additional renderer vectors pass: historical v1/v2 read-only audit and upgrade to v3;
unknown versions and annotation drift reject. An actual-binary core v2→v3 check changed only
lock bookkeeping and left README bytes unchanged. Source projection lint passes without
relaxing locked downstream validation. Draft recovery restores both the read input and its
ancestor-directory permissions before permitting cleanup.

Local validation passed the complete Go suite and vet, 28 pre-existing E2Es, the component
binary matrix, Ruby priming checks, template/project lint, template doctor, projection
consistency and whitespace checks. No manual-only acceptance gap is substituted for these checks.

## CI

- [Template acceptance](https://github.com/dapi/memory-bank/actions/runs/34082821247) passed at
  `06104c0e1bec4776ca1e75774fcd82b2012dbc98`, using CLI `caf0f3eaf3af290a702c8553795168584ac8b987`.
- [CLI Go/fixture/E2E suite](https://github.com/dapi/memory-bank-cli/actions/runs/34082697922),
  [release validation without publication](https://github.com/dapi/memory-bank-cli/actions/runs/34082697858),
  and [stable downstream smoke](https://github.com/dapi/memory-bank-cli/actions/runs/34082697864)
  passed at `caf0f3eaf3af290a702c8553795168584ac8b987`.

The final documentation handoff receives its own artifact review and PR CI. The PR records
that final result separately; the executable acceptance above refers to the immutable
implementation pair rather than making a self-referential claim about this evidence file.

## Independent review receipts

All listed receipts are structured code-converge results with `findings: []`, review-only
execution and zero fix budget. Full reviews were followed by author fixes and reviews of
the affected revisions; no finding was waived. Retained session IDs identify the local
structured records without embedding their environment or private invocation data.

| Scope | Reviewed revision | Receipt |
| --- | --- | --- |
| Full W2 implementation | 2dbd0a9 | session-1788752560558353000-99170-52f45a5337a9768a77c5a0501088e31a |
| Renderer implementation | 125f2d1 | session-1788753728844009000-24835-700c9c82576bbe42a5be23385c5dff2a |
| Final functional fixes | e2417e4 | session-1788754453808566000-55083-6bbc8a2be455ae7252622561284e8741 |
| CLI simplification closure | 243e4a5 | session-1788754659268920000-74978-90e2a00f789ebb84f6e77dad01f8931f |
| Template artifact fixes | fc43f14 | session-1788754770294462000-88452-e74decde35604d1c83948302f39e6663 |
| Entrypoint/CI review closure | 06104c0 | session-1788754911475188000-89391-987068e37929f4b74dfead45fbe25145 |
| Template simplification | 06104c0 | session-1788754967108259000-90270-4a6c231ab783a706d7bbd83df948ede0 |

The reviewed CLI simplification revision `243e4a5` and delivered `caf0f3e` have the identical
Git content tree `b2ba43b6fa1b417b4edb09013e708dbb2755180e`; only commit ancestry differs.
Design, ADR and Plan Ready reviews precede implementation and remain recorded in the
[feature plan](../memory-bank/features/FT-141/implementation-plan.md) and
[epic decision log](../memory-bank/epics/EP-141/decision-log.md).

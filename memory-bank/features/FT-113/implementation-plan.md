---
title: "FT-113: Implementation Plan"
doc_kind: feature
doc_function: derived
purpose: "Active execution и verification plan для FT-113."
derived_from:
  - brief.md
  - ../../flows/feature.md
status: active
audience: humans_and_agents
---

# FT-113: Implementation Plan

## Цель текущего плана

Добавить canonical BDD practice в generic template, связать её с существующими
flow, policy и indexes, сохранив ownership существующих artifacts.

## Lifecycle Note

Candidate implementation уже существует в commit `110fbc7ec0db62cfdad67749db5a7ed20696d3fa`.
Draft plan прошёл clean artifact review; resulting active revision должна быть
заморожена и получить final clean re-review до закрытия Plan Ready. Наличие
candidate implementation не переводит lifecycle в Execution или Done задним
числом.

## Grounding Evidence

- Grounded repository revision: `110fbc7ec0db62cfdad67749db5a7ed20696d3fa`

| Grounding ID | Inspected path / command | Observed current-state fact | Plan impact |
| --- | --- | --- | --- |
| `GRND-01` | `template/memory-bank/flows/feature.md`, `template/memory-bank/flows/use-case.md` | Feature acceptance и project-level scenarios имеют разные owners; `SC/NEG-*` остаются в feature brief. | `STEP-02` не создаёт новый route или owner. |
| `GRND-02` | `template/memory-bank/engineering/testing-policy.md`, `template/memory-bank/flows/behavior-specification.md` | Проверка поведения должна быть связана через `CHK-*` и `EVID-*`; Gherkin/Cucumber/E2E не обязательны. | `STEP-02` фиксирует test-surface-neutral automation handoff. |
| `GRND-03` | `memory-bank/features/FT-113/brief.md`, `memory-bank/features/README.md` | Feature package пока не входит в grounded revision; его candidate revision и review evidence должны быть зафиксированы отдельно. | `STEP-03` закрывает package traceability, review и lifecycle evidence до execution/closure. |

## Grounding / Support References

| Document | Role in this plan | Facts reused |
| --- | --- | --- |
| `brief.md` | canonical problem, scope and verify owner | `REQ-*`, `NS-*`, `SC-*`, `CHK-*`, `EVID-*` |
| `../../../template/memory-bank/flows/behavior-specification.md` | canonical BDD practice | ownership, Given/When/Then and automation handoff |
| `../../../template/memory-bank/flows/use-case.md` | canonical `UC-*` lifecycle | `UC-*` versus `SC/NEG-*` boundary |
| `../../../template/memory-bank/engineering/testing-policy.md` | canonical testing policy | test ownership and manual-only boundary |

## Test Strategy

| Test surface | Canonical refs | Planned verification | Required local suites / commands | Required CI | Manual-only gap |
| --- | --- | --- | --- | --- | --- |
| Requirements, ownership and traceability | `REQ-01`, `REQ-02`, `REQ-03`, `CHK-01` | Independent artifact review of frozen feature and affected governance revisions | Review issue #113 → `REQ/SC/CHK/EVID`, owner boundaries and scope; record reviewer, candidate SHA, findings/dispositions and verdict | none | Required artifact-review procedure; not a substitute for automated validation |
| Template documentation, manifests and links | `REQ-01`, `REQ-03`, `CHK-02` | Priming validators, lint, doctor and whitespace check | `ruby tools/validate-priming-manifests-test.rb`; `ruby tools/validate-priming-manifests.rb template/memory-bank`; `memory-bank-cli lint --scope-root template/memory-bank --entrypoint template/memory-bank/README.md`; `memory-bank-cli doctor --profile template`; `git diff --check` | `validate-template` | none |

## Open Questions / Ambiguities

`none` — research resolved the issue questions through existing routing, SSoT
and Use Case selection rules. Derived resolutions are explicitly labelled in
`brief.md`; новые customer requirements из них не выводятся.

## Preconditions

| Precondition ID | Canonical ref | Required state | Used by steps | Blocks start |
| --- | --- | --- | --- | --- |
| `PRE-01` | `REQ-01`, `REQ-02`, `REQ-03` | Issue #113 and the feature brief define the documentation scope. | `STEP-01`, `STEP-02`, `STEP-03` | yes |

## Design Realization Mapping

`not applicable`: `brief.md` records `Design required: no`.

## Workstreams

| Workstream | Implements | Result | Dependencies |
| --- | --- | --- | --- |
| `WS-1` | `REQ-01`, `REQ-02` | Canonical BDD practice connected to existing Feature/Use Case owners. | `PRE-01` |
| `WS-2` | `REQ-03` | Examples connect to checks/evidence without E2E-only or duplicate ownership. | `WS-1` |
| `WS-3` | `REQ-01`, `REQ-02`, `REQ-03` | Candidate artifacts pass independent review and produce revision-bound evidence. | `WS-1`, `WS-2` |

## Порядок работ

| Step ID | Implements | Goal | Touchpoints | Verifies | Evidence IDs | Check command / procedure |
| --- | --- | --- | --- | --- | --- | --- |
| `STEP-01` | `REQ-01` | Publish the canonical discussion/discovery → formulation → verification practice. | `template/memory-bank/flows/behavior-specification.md`, applicable indexes | `CHK-01`, `CHK-02` | `EVID-01`, `EVID-02` | Inspect practice and run documentation checks |
| `STEP-02` | `REQ-02`, `REQ-03` | Publish ownership and `UC/REQ → SC/NEG → CHK → check → EVID` traceability in the minimum required owners/templates. | Feature/Use Case flows, testing policy, brief/use-case/plan templates; feature-local companion only as existing derived boundary | `CHK-01`, `CHK-02` | `EVID-01`, `EVID-02` | Inspect stable-ID mappings, concrete examples and owner boundaries |
| `STEP-03` | `REQ-01`, `REQ-02`, `REQ-03` | Freeze a candidate revision containing the feature package, complete independent artifact review and attach concrete verification carriers. | `memory-bank/features/FT-113/`, external review record, CI/local results | `CHK-01`, `CHK-02` | `EVID-01`, `EVID-02` | Review the frozen revision; run every command in `CHK-02` |

## Checkpoints

| Checkpoint ID | Refs | Condition | Evidence IDs |
| --- | --- | --- | --- |
| `CP-01` | `STEP-01`, `STEP-02`, `CHK-01`, `CHK-02` | Candidate practice, ownership boundaries and traceability are internally consistent and validation is green. | `EVID-01`, `EVID-02` |
| `CP-02` | `STEP-03`, `CHK-01`, `CHK-02` | Clean review and all concrete carriers refer to the same frozen candidate revision. | `EVID-01`, `EVID-02` |

## Evidence

- `EVID-01` pending: external clean artifact-review record with reviewer,
  immutable candidate revision, findings/dispositions and verdict.
- `EVID-02` pending: local and CI outputs for every `CHK-02` command, all bound
  to the same candidate revision.

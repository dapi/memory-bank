---
title: "FT-53: Brownfield Adaptation Protocol"
doc_kind: feature
doc_function: canonical
purpose: "Canonical problem-space brief для отдельного protocol адаптации Memory Bank в существующем репозитории."
derived_from:
  - ../../flows/feature.md
  - ../../engineering/validation-profiles.md
status: active
delivery_status: in_progress
audience: humans_and_agents
must_not_define:
  - implementation_sequence
  - solution_space
---

# FT-53: Brownfield Adaptation Protocol

## What

### Problem

Issue [#53](https://github.com/dapi/memory-bank/issues/53) отмечает, что
существующий раздел brownfield в [`../../../docs/adoption.md`](../../../docs/adoption.md)
описывает inventory и порядок заполнения, но не даёт исполнимого end-to-end
protocol, сопоставимого с greenfield protocol. Если агент начинает discovery
после установки или чтения generic `memory-bank/`, placeholders и generic rules
могут быть ошибочно приняты за project facts.

### Outcome

| Metric ID | Metric | Baseline | Target | Measurement method |
| --- | --- | --- | --- | --- |
| `MET-53-01` | Наличие исполнимого brownfield protocol | В `docs/adoption.md` есть только краткие guidance sections | Новый protocol описывает discovery, intake PRD, adaptation, conversion, validation и trial | Semantic review against `REQ-53-01`–`REQ-53-06` and `SC-53-01` |

### Scope

- `REQ-53-01` Создать generic protocol для адаптации Memory Bank в brownfield web services и CLI utilities и связать его с repository navigation.
- `REQ-53-02` Явно отделить pre-adaptation discovery от governed Memory Bank work: до установки/активации `memory-bank/` agent читает repository instructions и только существующие project sources, но не consults `memory-bank/`.
- `REQ-53-03` Определить временный evidence-backed intake PRD в `./brownfield-intake-prd.md` вне `memory-bank/`: required fields, lifecycle и запрет на invented architecture/delivery plan; затем описать его governed conversion в `memory-bank/prd/PRD-XXX-*.md` после adaptation upstream owners.
- `REQ-53-04` Определить conditional inventory и adaptation для web service и CLI, включая явное `N/A` с причиной для неприменимых dimensions.
- `REQ-53-05` Определить safety/migration rules, minimum rollout DoD, completion evidence и один real-task trial; не допускать копирования secrets, PII, tokens или internal endpoints.
- `REQ-53-06` Reconcile existing brownfield guidance с protocol и greenfield protocol без изменения runtime code, repository instructions или existing documentation кроме связанных navigation/guidance documents.

### Non-Scope

- `NS-53-01` Не адаптировать этот repository как downstream project и не создавать project-specific facts.
- `NS-53-02` Не изменять runtime code, CLI behavior, CI workflow или `memory-bank` validator.
- `NS-53-03` Не создавать delivery plans, feature packages, epics, use cases или historical ADR в downstream repository без evidence из его existing sources.

### Constraints / Assumptions

- `ASM-53-01` Issue #53 — authoritative source required outcome and acceptance criteria for this delivery-unit.
- `CON-53-01` Новый protocol остаётся generic: project-specific terminology, facts, secrets and internal endpoints в него не попадают.
- `CON-53-02` Existing `docs/adoption.md` и `docs/greenfield-integration-protocol.md` являются related canonical guidance; FT-53 must reconcile, rather than duplicate or contradict, their lifecycle boundaries.
- `DEC-53-01` Design required: no. Evidence: scope is a documentation-only protocol and links; no executable behavior, API/event/schema/config/release or architecture boundary changes are requested. Decision detail and FPF record: [`decision-log.md`](decision-log.md#dec-53-01-design-owner-not-required).
- `DEC-53-03` Default temporary intake location is `./brownfield-intake-prd.md`; it is explicit, generic and necessarily outside `memory-bank/`. Decision detail and FPF record: [`decision-log.md`](decision-log.md#dec-53-03-default-intake-location-is-the-repository-root).

## Design Requirement Decision

| Decision | Reason | Downstream owner |
| --- | --- | --- |
| `Design required: no` | `REQ-53-*` changes documentation guidance only; selected runtime solution, architecture coverage and contracts are out of scope. | `none` |

## Artifact Routing Decision

| Artifact | Decision | Trigger / reason | Route / owner |
| --- | --- | --- | --- |
| `decision-log.md` | selected | User requires a decision log within the feature; it records evidence-backed decisions without becoming a second canonical owner. | Feature-local decision record derived from `brief.md` |
| `design.md` | omitted | `DEC-53-01`: no solution-space trigger. | `none` |
| feature support artifacts and ADR | omitted | No runtime surface, UI, integration, architectural or reusable cross-feature decision is in scope. | `none` |

## Validation Profile Decision

| Profile | Triggers / rationale | Downgrade approval |
| --- | --- | --- |
| `documentation` | Only docs and repository navigation change; executable behavior, contracts, production config and release path remain unchanged (`NS-53-02`). | `none` |

## Verify

### Exit Criteria

- `EC-53-01` Protocol exists, is reachable from root navigation and adoption guidance, and explicitly prohibits consulting `memory-bank/` during pre-adaptation discovery.
- `EC-53-02` Protocol specifies intake PRD fields, location, conversion and dependency rule; web-service/CLI inventory, `N/A` treatment, safety rules and rollout DoD are complete.
- `EC-53-03` Existing brownfield guidance and greenfield protocol have no lifecycle contradiction, and all changed docs pass applicable link/index validation or record a verification gap.

### Traceability matrix

| Requirement ID | Problem refs | Acceptance refs | Checks | Evidence IDs |
| --- | --- | --- | --- | --- |
| `REQ-53-01`, `REQ-53-02` | `ASM-53-01`, `CON-53-01` | `EC-53-01`, `SC-53-01` | `CHK-53-01`, `CHK-53-03` | `EVID-53-01`, `EVID-53-03` |
| `REQ-53-03`, `REQ-53-04`, `REQ-53-05` | `ASM-53-01`, `CON-53-01` | `EC-53-02`, `SC-53-02` | `CHK-53-01` | `EVID-53-01` |
| `REQ-53-06` | `CON-53-02` | `EC-53-03`, `SC-53-03` | `CHK-53-02`, `CHK-53-03` | `EVID-53-02`, `EVID-53-03` |

### Acceptance Scenarios

- `SC-53-01` A brownfield agent follows the protocol and completes pre-adaptation discovery without opening the generic or installed `memory-bank/`, then knows when to install and govern it.
- `SC-53-02` A web-service or CLI adaptation records only supported evidence, marks an unsupported dimension `N/A` with a reason, creates an intake PRD outside Memory Bank and converts it only after upstream owners are adapted.
- `SC-53-03` A reviewer can navigate from root/adoption docs to the protocol, compare its lifecycle with greenfield guidance, and run the documented validation commands without broken links or unindexed feature docs.

### Checks

| Check ID | Covers | How to check | Expected result | Evidence path |
| --- | --- | --- | --- | --- |
| `CHK-53-01` | `EC-53-01`, `EC-53-02`, `SC-53-01`, `SC-53-02` | Semantic read-through of protocol against `REQ-53-01`–`REQ-53-05` and issue #53. | Every required lifecycle, inventory, safety and DoD fact has one unambiguous owner in the protocol or linked canonical guidance. | `git diff -- docs/adoption.md docs/brownfield-adaptation-protocol.md README.md` |
| `CHK-53-02` | `EC-53-03`, `SC-53-03` | Compare brownfield protocol, adoption guidance and greenfield protocol for discovery/install/adaptation boundaries. | No document permits consulting `memory-bank/` before brownfield discovery; no conflicting lifecycle claim remains. | `git diff --check` plus reviewer read-through |
| `CHK-53-03` | `EC-53-01`, `EC-53-03`, `SC-53-03` | Run `memory-bank lint`; if unavailable, record the command failure and run `rg --files memory-bank` plus direct link inspection. | Lint passes, or the exact verification gap is recorded without claiming success. | Command output in review/PR evidence |

### Test matrix

| Check ID | Evidence IDs | Evidence path |
| --- | --- | --- |
| `CHK-53-01` | `EVID-53-01` | `git diff -- docs/adoption.md docs/brownfield-adaptation-protocol.md README.md` |
| `CHK-53-02` | `EVID-53-02` | `git diff --check` and reviewer read-through |
| `CHK-53-03` | `EVID-53-03` | `memory-bank lint` output or stated fallback gap |

### Evidence

- `EVID-53-01` Diff and semantic review show complete protocol coverage.
- `EVID-53-02` Diff-check and reconciliation read-through show no contradictory guidance.
- `EVID-53-03` Lint output, or an explicit unavailable-command verification gap with fallback evidence.

### Evidence contract

| Evidence ID | Artifact | Producer | Path contract | Reused by checks |
| --- | --- | --- | --- | --- |
| `EVID-53-01` | Changed-doc diff and reviewer notes | implementer / reviewer | PR or working-tree diff limited to changed documentation | `CHK-53-01` |
| `EVID-53-02` | Whitespace/link reconciliation result | implementer / reviewer | `git diff --check` output and review notes | `CHK-53-02` |
| `EVID-53-03` | Lint result or explicit verification-gap note | implementer | command output attached to PR/review | `CHK-53-03` |

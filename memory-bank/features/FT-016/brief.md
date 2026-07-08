---
title: "FT-016: Privacy / Source Boundary Support Template"
doc_kind: feature
doc_function: canonical
purpose: "Canonical brief для issue 16. Фиксирует problem space, scope, non-scope и verify contract для generic privacy/source-boundary support template."
derived_from:
  - ../../flows/feature-flow.md
  - https://github.com/dapi/memory-bank/issues/16
status: active
delivery_status: done
audience: humans_and_agents
must_not_define:
  - implementation_sequence
  - solution_space
---

# FT-016: Privacy / Source Boundary Support Template

## What

### Problem

AI/agent tooling, observability, logs, transcripts и external metadata могут содержать приватные данные. Сейчас feature support templates не дают generic места, где feature package явно фиксирует, какие source classes можно читать, хранить и цитировать, какие source classes исключены, какой confidence/status у источников и кто владеет boundary facts.

Issue 16 указывает downstream source patterns из `dapi/zelma`, но требует не переносить `zelma`, Codex session internals или конкретные runtime details обратно в template.

### Outcome

| Metric ID | Metric | Baseline | Target | Measurement method |
| --- | --- | --- | --- | --- |
| `MET-01` | Generic privacy/source-boundary support coverage | В `flows/templates/feature/support/` нет dedicated template для privacy/source boundaries | Есть generic support template с allowed metadata, explicitly excluded data, source inventory confidence/status, test evidence и owner boundaries | Review changed template docs against `REQ-*`, `SC-*`, `CHK-*` |
| `MET-02` | Template boundary clarity | Existing support docs declare non-ownership individually, but privacy/source-boundary support is absent | New support doc explicitly states it does not replace `brief.md`, `design.md` or `implementation-plan.md` | Traceability and content review of support template plus routing docs |

### Scope

- `REQ-01` Add a generic feature-support template for privacy/source-boundary documentation under `memory-bank/flows/templates/feature/support/`.
- `REQ-02` The template must cover allowed metadata, explicitly excluded data, confidence/status for sources, source inventory, test evidence, and owner boundaries.
- `REQ-03` Update feature-flow/template routing so agents know when a privacy/source-boundary support doc is needed.
- `REQ-04` Keep the template generic and exclude downstream-specific `zelma`, Codex session internals and concrete runtime details.
- `REQ-05` Register the new support template in template indexes / navigation surfaces touched by the feature.

### Non-Scope

- `NS-01` Do not implement runtime privacy filtering, source discovery code, observability code, or tests outside documentation templates.
- `NS-02` Do not add project-specific `zelma` behavior, Codex session internals, concrete command contracts, raw log shapes, transcript schemas or production runtime details to the generic template.
- `NS-03` Do not turn the support doc into a global privacy policy, legal compliance policy, ADR, or canonical owner for feature requirements / selected design / execution sequencing.
- `NS-04` Do not create a project-level use case unless implementation work later proves a stable reusable scenario beyond this template change.

### Constraints / Assumptions

- `ASM-01` Issue 16 is the canonical source for scope and acceptance for this feature.
- `ASM-02` The referenced `dapi/zelma` files are source examples for abstraction only; they are not upstream owners for this generic template repository.
- `CON-01` Feature-flow boundary rules apply: support docs may aid grounding/review/traceability but cannot own canonical problem space, solution space, acceptance inventory, or execution sequencing.
- `CON-02` The repository has no build/runtime app; verification is documentation/index/link oriented.

## Design Requirement Decision

| Decision | Reason | Downstream owner |
| --- | --- | --- |
| `Design required: yes` | The feature changes governed template contracts and support-doc owner boundaries. Selected template shape, routing rules and local decisions must be stabilized before execution planning. | `design.md` |

## Verify

### Exit Criteria

- `EC-01` New privacy/source-boundary support template exists in the feature support template directory and is registered in relevant template navigation.
- `EC-02` The new template explicitly covers allowed metadata, excluded data, source inventory, confidence/status, evidence, and owner boundaries.
- `EC-03` Feature-flow and feature/design/implementation-plan templates explain when this support doc is needed without making it canonical owner for requirements, selected design, checks, evidence contract or execution sequencing.
- `EC-04` Generic template surfaces do not contain downstream-specific `zelma`, Codex session internals, concrete runtime details, raw transcript/log schemas, prompts or secrets.

### Traceability matrix

| Requirement ID | Problem refs | Acceptance refs | Checks | Evidence IDs |
| --- | --- | --- | --- | --- |
| `REQ-01` | `ASM-01`, `CON-01` | `EC-01`, `SC-01` | `CHK-01` | `EVID-01` |
| `REQ-02` | `ASM-01`, `CON-01` | `EC-02`, `SC-01`, `SC-02`, `NEG-01` | `CHK-02` | `EVID-02` |
| `REQ-03` | `CON-01` | `EC-03`, `SC-03` | `CHK-03` | `EVID-03` |
| `REQ-04` | `ASM-02`, `NS-02` | `EC-04`, `NEG-01` | `CHK-04` | `EVID-04` |
| `REQ-05` | `CON-02` | `EC-01` | `CHK-01`, `CHK-05` | `EVID-01`, `EVID-05` |

### Acceptance Scenarios

- `SC-01` An agent creating a feature that touches logs, transcripts, external metadata or source inventories can instantiate a generic privacy/source-boundary support doc from `flows/templates/feature/support/`.
- `SC-02` The instantiated support doc gives the agent places to record allowed metadata, explicitly excluded private data, source confidence/status, evidence expectations and ownership boundaries.
- `SC-03` An agent reading feature-flow, design template or implementation-plan template can identify when the privacy/source-boundary support doc is appropriate and where canonical facts still belong.

### Negative / Edge Scenarios

- `NEG-01` The generic template must not encourage reading, storing or citing private transcripts, prompts, raw logs, secrets, tool IO or downstream runtime internals without an explicit feature-owned permission / boundary.
- `NEG-02` The support doc must not redefine `REQ-*`, selected `SOL-*`, canonical `CHK-*` / `EVID-*`, or `STEP-*`.

### Checks

| Check ID | Covers | How to check | Expected result | Evidence path |
| --- | --- | --- | --- | --- |
| `CHK-01` | `EC-01`, `REQ-01`, `REQ-05` | `rg -n "privacy-source-boundary|Privacy / Source Boundary|source-boundary" memory-bank/flows memory-bank/features/README.md` | New template and routing/index references are discoverable | `artifacts/ft-016/verify/chk-01/` |
| `CHK-02` | `EC-02`, `REQ-02` | Review `memory-bank/flows/templates/feature/support/privacy-source-boundary.md` for required sections | Required coverage appears without canonical ownership leakage | `artifacts/ft-016/verify/chk-02/` |
| `CHK-03` | `EC-03`, `REQ-03` | Review feature-flow plus feature `brief.md`, `design.md`, `implementation-plan.md` templates for routing and ownership wording | Routing exists and keeps owner boundaries intact | `artifacts/ft-016/verify/chk-03/` |
| `CHK-04` | `EC-04`, `REQ-04`, `NEG-01` | `rg -n "zelma|Codex session|session_meta|zellij|\\.zelma|process_argv" memory-bank/flows/templates memory-bank/flows/feature-flow.md` | No downstream-specific source/runtime terms appear in generic template surfaces | `artifacts/ft-016/verify/chk-04/` |
| `CHK-05` | `EC-01`, `REQ-05` | `python3 scripts/check_memory_bank_index.py && git diff --check` | Index/link audit and whitespace checks pass | `artifacts/ft-016/verify/chk-05/` |

### Test matrix

| Check ID | Evidence IDs | Evidence path |
| --- | --- | --- |
| `CHK-01` | `EVID-01` | `artifacts/ft-016/verify/chk-01/` |
| `CHK-02` | `EVID-02` | `artifacts/ft-016/verify/chk-02/` |
| `CHK-03` | `EVID-03` | `artifacts/ft-016/verify/chk-03/` |
| `CHK-04` | `EVID-04` | `artifacts/ft-016/verify/chk-04/` |
| `CHK-05` | `EVID-05` | `artifacts/ft-016/verify/chk-05/` |

### Evidence

- `EVID-01` Search output proving template and index/routing references exist.
- `EVID-02` Review note or diff excerpt proving required sections exist in the new support template.
- `EVID-03` Review note or diff excerpt proving routing/ownership updates exist in feature-flow and feature templates.
- `EVID-04` Search output proving downstream-specific source/runtime terms are absent from generic template surfaces.
- `EVID-05` Output from memory-bank index audit and whitespace check.

### Evidence contract

| Evidence ID | Artifact | Producer | Path contract | Reused by checks |
| --- | --- | --- | --- | --- |
| `EVID-01` | Search output | implementer | `artifacts/ft-016/verify/chk-01/` | `CHK-01` |
| `EVID-02` | Template coverage review note | reviewer / implementer | `artifacts/ft-016/verify/chk-02/` | `CHK-02` |
| `EVID-03` | Routing/ownership review note | reviewer / implementer | `artifacts/ft-016/verify/chk-03/` | `CHK-03` |
| `EVID-04` | Generic-surface leakage search output | implementer | `artifacts/ft-016/verify/chk-04/` | `CHK-04` |
| `EVID-05` | Index audit and diff check output | implementer | `artifacts/ft-016/verify/chk-05/` | `CHK-05` |

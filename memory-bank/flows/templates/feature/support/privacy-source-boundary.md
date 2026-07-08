---
title: "FT-XXX: Privacy / Source Boundary Template"
doc_kind: feature-support
doc_function: template
purpose: Governed wrapper-шаблон optional `privacy-source-boundary.md`. Читать, когда feature needs explicit source inventory, privacy/data-use boundaries, confidence/status notes and evidence handling rules.
derived_from:
  - ../../../feature-flow.md
  - ../../../../dna/frontmatter.md
status: active
audience: humans_and_agents
template_for: feature-support
template_target_path: ../../../../features/FT-XXX/privacy-source-boundary.md
canonical_for:
  - feature_support_template_privacy_source_boundary
---

# FT-XXX: Privacy / Source Boundary Template

Этот файл описывает wrapper-template. Инстанцируемый `privacy-source-boundary.md` живет внутри feature package как optional support/reference doc.

## Wrapper Notes

Создавай `privacy-source-boundary.md`, если feature touches logs, transcripts, diagnostics, observability artifacts, external metadata, source inventories or any source class where agents need explicit rules for what may be read, stored, cited, summarized or used as evidence.

`privacy-source-boundary.md` не владеет requirements, selected design, acceptance criteria, canonical checks, evidence contract или implementation sequence. Если во время source inventory меняется scope или selected design, обнови sibling `brief.md`, required `design.md` или ADR до продолжения.

Source confidence is not access permission. A source may be high-confidence and still explicitly excluded from reading, storage or citation.

## Instantiated Frontmatter

```yaml
title: "FT-XXX: Privacy / Source Boundary"
doc_kind: feature-support
doc_function: reference
purpose: "Privacy/source boundary reference для FT-XXX. Фиксирует allowed metadata, explicitly excluded data, source inventory confidence/status, owner boundaries and evidence handling without переопределения canonical problem, solution or execution facts."
derived_from:
  - brief.md
  # Required only when design.md exists:
  # - design.md
status: draft
audience: humans_and_agents
must_not_define:
  - ft_xxx_scope
  - ft_xxx_selected_design
  - ft_xxx_acceptance_criteria
  - ft_xxx_evidence_contract
  - implementation_sequence
  - global_privacy_policy
```

## Instantiated Body

```markdown
# FT-XXX: Privacy / Source Boundary

## Role

Этот документ фиксирует feature-local source/privacy grounding. Canonical owners:

- `brief.md` владеет problem space, scope, non-scope and canonical verify inventory.
- `design.md`, если есть, владеет selected design, contracts, invariants and failure modes.
- `implementation-plan.md` владеет execution sequencing, approval gates and check procedures.

Use this support doc only to make source access, source confidence, privacy exclusions and evidence handling auditable for this feature.

## Boundary Summary

| Boundary | Summary | Related canonical refs |
| --- | --- | --- |
| Allowed use | What source-derived metadata or summaries may be used | `REQ-01`, `CON-01` |
| Excluded use | What raw or private content must not be read, stored, cited or copied | `NS-01`, `NEG-01` |
| Evidence use | Which evidence carriers are acceptable for checks without exposing excluded content | `CHK-01`, `EVID-01` |

## Source Inventory / Confidence

Confidence/status describes the quality and availability of the source claim. It does not grant permission to read private content.

| Source ID | Source class / carrier | Owner / steward | Availability / status | Confidence | Safe use | Privacy boundary refs | Evidence refs |
| --- | --- | --- | --- | --- | --- | --- | --- |
| `SRC-01` | Source class or carrier type | Who owns or approves this source | present / missing / candidate / unavailable | weak / medium / strong / unknown | What can safely be inferred or reused | `ALLOW-01`, `EXCL-01` | `EVID-01` |

## Allowed Metadata / Data Use

Allowed rows must be minimal. If a field or source class is not listed here, treat it as unavailable until the canonical owner updates the boundary.

| Allowed ID | Source refs | Allowed metadata / data class | Allowed actions | Conditions / minimization | Related refs |
| --- | --- | --- | --- | --- | --- |
| `ALLOW-01` | `SRC-01` | Metadata or aggregate class that may be used | read / store / cite / summarize | Required minimization, redaction or retention notes | `REQ-01`, `CHK-01` |

## Explicitly Excluded Data

Excluded rows take precedence over allowed rows. Do not read, store, cite, summarize or copy excluded content unless a later owner-approved boundary explicitly changes this section.

| Exclusion ID | Source refs | Must not read / store / cite | Why excluded | Permission or owner change required | Related refs |
| --- | --- | --- | --- | --- | --- |
| `EXCL-01` | `SRC-01` | Private raw content, secrets or other excluded data class | Privacy, safety, legal or scope reason | Who must approve and where approval is recorded | `NS-01`, `NEG-01`, `AG-01` |

## Permission / Owner Boundaries

Use this section when access depends on an owner, explicit approval, retention policy or external boundary. Approval gates used during execution still belong in `implementation-plan.md` as `AG-*`.

| Permission ID | Trigger | Owner / approver | Allowed after approval | Still excluded | Approval / evidence path |
| --- | --- | --- | --- | --- | --- |
| `PERM-01` | What event or requested use needs approval | Responsible owner | What becomes allowed if approval is granted | What remains excluded | `AG-01` / issue / PR / audit path |

## Test Evidence Boundary

This section constrains evidence carriers for canonical checks. It does not create new canonical `CHK-*` or `EVID-*`; those remain in `brief.md`.

| Check / evidence refs | Allowed evidence carrier | Required redaction or summary | Must not include | Producer / reviewer |
| --- | --- | --- | --- | --- |
| `CHK-01`, `EVID-01` | Log excerpt, report, fixture, screenshot, summary or other carrier allowed by this boundary | What must be redacted, aggregated or summarized | Excluded source content from `EXCL-*` | Who produces and reviews evidence |

## Source Handling Rules

- `SRC-*` rows inventory source classes or carriers; they do not authorize raw access by themselves.
- `ALLOW-*` rows define the minimum allowed metadata or data use.
- `EXCL-*` rows override `ALLOW-*` rows when they overlap.
- `PERM-*` rows describe owner boundaries; execution-time approval gates still belong in `implementation-plan.md` as `AG-*`.
- If this document discovers a new requirement or non-scope boundary, update `brief.md` first.
- If this document discovers a selected design, contract, invariant or failure mode, update `design.md` or ADR first.

## Notes For Implementation Plan

- Which `SRC-*`, `ALLOW-*`, `EXCL-*` and `PERM-*` rows must be reflected in plan `PRE-*`, `AG-*`, `STEP-*` or `STOP-*`.
- Which evidence carriers are acceptable for `CHK-*` without exposing excluded content.
- Which source ambiguity should become `OQ-*` if it blocks execution.
```

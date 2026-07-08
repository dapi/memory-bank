---
title: "FT-012: Compact Task Flow"
doc_kind: feature
doc_function: canonical
purpose: "Canonical brief для delivery-единицы issue #12. Фиксирует problem space, scope и verify contract для generic compact task-flow без смешения с solution space или execution plan."
derived_from:
  - ../../flows/feature-flow.md
  - ../../flows/workflows.md
  - ../../engineering/testing-policy.md
  - https://github.com/dapi/memory-bank/issues/12
status: active
delivery_status: done
audience: humans_and_agents
must_not_define:
  - implementation_sequence
  - solution_space
---

# FT-012: Compact Task Flow

## What

### Problem

GitHub issue [#12](https://github.com/dapi/memory-bank/issues/12) фиксирует, что текущий шаблон хорошо покрывает feature/epic delivery, но small bugfix/refactor/chore задачи получают слишком тяжелый full feature package. В source repository `alfagen/mercury` уже есть intermediate layer: compact issue/PR carrier или durable `TASK-XXX` package без превращения задачи в feature.

### Outcome

| Metric ID | Metric | Baseline | Target | Measurement method |
| --- | --- | --- | --- | --- |
| `MET-01` | Non-feature workflow coverage | В шаблоне есть `feature-flow` и `epic-flow`, но нет generic `task-flow` family | `task-flow.md`, `bugfix-flow.md`, `refactor-flow.md`, `TASK-XXX` templates и `tasks/README.md` доступны через индексы | `CHK-01`, `CHK-02` |
| `MET-02` | Promotion safety | Compact путь не описан, риск спрятать feature/contract/high-risk work не управляется отдельными triggers | Compact profiles явно требуют promotion при capability, contract, design или high/critical risk trigger | `CHK-02`, `CHK-04` |
| `MET-03` | Source sanitization | Source docs содержат downstream-specific examples and terms | В template не перенесены project-specific terms, issue ids или domain facts из source repo | `CHK-03` |

### Scope

- `REQ-01` Добавить generic governance docs: `memory-bank/flows/task-flow.md`, `memory-bank/flows/bugfix-flow.md`, `memory-bank/flows/refactor-flow.md`.
- `REQ-02` Добавить governed templates для `TASK-XXX`: package README, bugfix и refactor/chore.
- `REQ-03` Добавить `memory-bank/tasks/README.md` как optional destination для managed non-feature задач.
- `REQ-04` Обновить `memory-bank/flows/README.md`, `memory-bank/flows/templates/README.md`, root `memory-bank/README.md` и related indexes/routing docs, которые должны раскрывать новый compact task layer.
- `REQ-05` Сохранить generic template boundary: не переносить project-specific terms, source issue ids, examples или domain facts из `alfagen/mercury`.
- `REQ-06` Зафиксировать promotion triggers так, чтобы feature/contract/high-risk/design work нельзя было спрятать в task package.

### Non-Scope

- `NS-01` Не переносить source-only `workflow-decision-log.md`, `workflow-metrics.md` или другие artifacts, которых нет в scope issue #12.
- `NS-02` Не создавать sample `TASK-XXX` packages с downstream identifiers or domain facts from source repo.
- `NS-03` Не заменять `feature-flow.md` или `epic-flow.md`; compact task layer должен быть отдельным routing/profile layer.
- `NS-04` Не добавлять runtime application code, CI integration или external automation beyond documentation checks.
- `NS-05` Не создавать project-specific policy для конкретного downstream продукта.

### Constraints / Assumptions

- `ASM-01` Source docs from `alfagen/mercury` are evidence for structure and semantics, but only generic rules can be adapted into this template.
- `ASM-02` Current repository has no runtime application; verification is documentation/link/index oriented.
- `CON-01` Governed docs under `memory-bank/` must have YAML frontmatter with `status` and valid local links.
- `CON-02` `python3 scripts/check_memory_bank_index.py` is the canonical automated index/link audit named by issue #12.
- `CON-03` Existing `feature-flow.md` remains canonical for feature packages; task docs must route promotion back to feature/ADR/epic when compact criteria fail.

## Design Requirement Decision

| Decision | Reason | Downstream owner |
| --- | --- | --- |
| `Design required: yes` | Issue #12 changes governance flows, templates, routing/indexes and promotion boundaries. Without a solution-space owner, the execution plan would have to decide how task, bugfix, refactor and feature flows relate. | `design.md` |

## Verify

### Exit Criteria

- `EC-01` `task-flow.md`, `bugfix-flow.md` and `refactor-flow.md` exist, are generic governed docs and are discoverable through flow indexes.
- `EC-02` `TASK-XXX` package README, bugfix and refactor templates exist under `memory-bank/flows/templates/task/` and are discoverable through templates indexes.
- `EC-03` `memory-bank/tasks/README.md` exists as optional managed non-feature task destination and does not contain source repo examples.
- `EC-04` Related indexes/routing docs link the new task family without making task-flow replace `feature-flow` or `epic-flow`.
- `EC-05` Promotion triggers reject compact task usage for capability, contract, high/critical risk, design reasoning or multi-delivery-unit work.
- `EC-06` Documentation checks pass and evidence is captured.

### Traceability matrix

| Requirement ID | Problem refs | Acceptance refs | Checks | Evidence IDs |
| --- | --- | --- | --- | --- |
| `REQ-01` | `ASM-01`, `CON-01`, `CON-03` | `EC-01`, `SC-01`, `SC-03` | `CHK-01`, `CHK-02`, `CHK-04` | `EVID-01`, `EVID-02`, `EVID-04` |
| `REQ-02` | `ASM-01`, `CON-01` | `EC-02`, `SC-02` | `CHK-01`, `CHK-02` | `EVID-01`, `EVID-02` |
| `REQ-03` | `ASM-01`, `CON-01` | `EC-03`, `SC-02` | `CHK-01`, `CHK-02`, `CHK-03` | `EVID-01`, `EVID-02`, `EVID-03` |
| `REQ-04` | `CON-01`, `CON-02` | `EC-04`, `SC-04` | `CHK-01`, `CHK-02`, `CHK-04` | `EVID-01`, `EVID-02`, `EVID-04` |
| `REQ-05` | `ASM-01`, `CON-01` | `EC-03`, `SC-05` | `CHK-03` | `EVID-03` |
| `REQ-06` | `CON-03` | `EC-05`, `SC-03`, `NEG-01` | `CHK-02`, `CHK-04` | `EVID-02`, `EVID-04` |

### Acceptance Scenarios

- `SC-01` Reader opens `memory-bank/flows/workflows.md` and can choose a compact non-feature profile before defaulting to full `feature-flow`.
- `SC-02` Reader opens `memory-bank/tasks/README.md` and can create a durable `TASK-XXX/` package with exactly one primary owner doc: `bugfix.md` or `refactor.md`.
- `SC-03` Reader evaluates a task with capability, contract, design, high-risk or multi-delivery-unit trigger and is routed out of compact task package to feature/ADR/epic flow.
- `SC-04` Reader starts at `memory-bank/README.md` and can reach task flows, task templates and `memory-bank/tasks/README.md` through documented indexes.
- `SC-05` Reviewer searches source-specific terms from `alfagen/mercury` and finds no downstream domain examples in the template.

### Negative / Edge Scenarios

- `NEG-01` A bugfix/refactor candidate with API/schema/security/financial/integration/rollout contract change must not remain `tracker-only`, `bugfix-compact`, `refactor-small` or managed task without promotion.

### Checks

| Check ID | Covers | How to check | Expected result | Evidence path |
| --- | --- | --- | --- | --- |
| `CHK-01` | `EC-01`, `EC-02`, `EC-03`, `EC-04`, `SC-04` | `python3 scripts/check_memory_bank_index.py` | No broken links, frontmatter dependency issues, or unreachable/orphan docs | `artifacts/ft-012/verify/chk-01/` |
| `CHK-02` | `EC-01`, `EC-02`, `EC-03`, `EC-05`, `SC-01`, `SC-02`, `SC-03`, `NEG-01` | Manual doc review of new flow/template docs against issue #12 and `feature-flow.md` boundaries | Compact task family is explicit, governed and does not redefine feature/epic ownership | `artifacts/ft-012/verify/chk-02/` |
| `CHK-03` | `EC-03`, `SC-05`, `REQ-05` | `rg -n "alfagen|mercury|TASK-3446|rate-daemon|SlotDiagnostics|PositionAware|production log storm" memory-bank` | No matches outside feature docs where source provenance is intentionally cited | `artifacts/ft-012/verify/chk-03/` |
| `CHK-04` | `EC-04`, `EC-05`, `REQ-06` | `rg -n "workflow_profile|Promotion|promotion|contract_change|feature-package|epic-package|ADR" memory-bank/flows memory-bank/tasks` plus reviewer inspection | Promotion triggers are present in selector, task flow and profile docs/templates | `artifacts/ft-012/verify/chk-04/` |
| `CHK-05` | `EC-06` | `git diff --check` | No trailing whitespace or conflict markers | `artifacts/ft-012/verify/chk-05/` |

### Test matrix

| Check ID | Evidence IDs | Evidence path |
| --- | --- | --- |
| `CHK-01` | `EVID-01` | `artifacts/ft-012/verify/chk-01/` |
| `CHK-02` | `EVID-02` | `artifacts/ft-012/verify/chk-02/` |
| `CHK-03` | `EVID-03` | `artifacts/ft-012/verify/chk-03/` |
| `CHK-04` | `EVID-04` | `artifacts/ft-012/verify/chk-04/` |
| `CHK-05` | `EVID-05` | `artifacts/ft-012/verify/chk-05/` |

### Evidence

- `EVID-01` Output of `python3 scripts/check_memory_bank_index.py`.
- `EVID-02` Review note confirming required flow/template docs and boundary separation.
- `EVID-03` Source-specific term scan output.
- `EVID-04` Promotion trigger scan and reviewer verdict.
- `EVID-05` Output of `git diff --check`.

### Evidence contract

| Evidence ID | Artifact | Producer | Path contract | Reused by checks |
| --- | --- | --- | --- | --- |
| `EVID-01` | Link/index audit log | implementer | `artifacts/ft-012/verify/chk-01/index-audit.txt` | `CHK-01` |
| `EVID-02` | Manual review note | implementer/reviewer | `artifacts/ft-012/verify/chk-02/manual-review.md` | `CHK-02` |
| `EVID-03` | Source-specific scan log | implementer | `artifacts/ft-012/verify/chk-03/source-scan.txt` | `CHK-03` |
| `EVID-04` | Promotion trigger scan log and verdict | implementer/reviewer | `artifacts/ft-012/verify/chk-04/promotion-scan.txt`, `artifacts/ft-012/verify/chk-04/promotion-verdict.md` | `CHK-04` |
| `EVID-05` | Whitespace/conflict-marker audit log | implementer | `artifacts/ft-012/verify/chk-05/diff-check.txt` | `CHK-05` |

### Verification Result

- `CHK-01` pass: `artifacts/ft-012/verify/chk-01/index-audit.txt`.
- `CHK-02` pass: `artifacts/ft-012/verify/chk-02/manual-review.md`.
- `CHK-03` pass: `artifacts/ft-012/verify/chk-03/source-scan.txt`.
- `CHK-04` pass: `artifacts/ft-012/verify/chk-04/promotion-scan.txt` and `artifacts/ft-012/verify/chk-04/promotion-verdict.md`.
- `CHK-05` pass: `artifacts/ft-012/verify/chk-05/diff-check.txt`.

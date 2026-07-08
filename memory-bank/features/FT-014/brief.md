---
title: "FT-014: Epic Intake Brief Template"
doc_kind: feature
doc_function: canonical
purpose: "Canonical brief для добавления lightweight epic intake brief template. Фиксирует problem space, scope и verify без смешения с solution design или execution plan."
derived_from:
  - ../../flows/feature-flow.md
  - ../../flows/epic-flow.md
status: active
delivery_status: in_progress
audience: humans_and_agents
source_issue: "https://github.com/dapi/memory-bank/issues/14"
must_not_define:
  - implementation_sequence
  - solution_space
---

# FT-014: Epic Intake Brief Template

## What

### Problem

GitHub issue 14 фиксирует gap в epic lifecycle: текущий epic package начинается с полного setup через `charter.md`, `roadmap.md`, `risks.md` и `subissues.md`, но до полного epic setup иногда нужен более легкий intake-документ. В качестве source указан `dapi/zelma: memory-bank/flows/templates/epic/brief.md`, при этом issue явно запрещает переносить `zelma`-specific content.

### Outcome

| Metric ID | Metric | Baseline | Target | Measurement method |
| --- | --- | --- | --- | --- |
| `MET-01` | Epic intake template discoverability | `memory-bank/flows/templates/epic/brief.md` отсутствует | Новый template доступен из epic template indexes | `CHK-01`, `CHK-02` |
| `MET-02` | Boundary clarity | `epic-flow.md` не описывает lightweight intake brief | `epic-flow.md` явно говорит, что brief не заменяет charter/roadmap/subissues/risks и не владеет implementation sequence | `CHK-02` |

### Scope

- `REQ-01` Добавить `memory-bank/flows/templates/epic/brief.md` как generic lightweight intake layer для epic proposal.
- `REQ-02` Обновить `memory-bank/flows/templates/epic/README.md`, чтобы новый epic brief template был discoverable.
- `REQ-03` Обновить `memory-bank/flows/templates/README.md`, если это нужно для reachability и template index consistency.
- `REQ-04` Уточнить `memory-bank/flows/epic-flow.md`, что intake `brief.md` не заменяет canonical full epic package owners и не владеет implementation sequence.
- `REQ-05` Не переносить `zelma`-specific content в governed template/governance docs.
- `REQ-06` Сохранить green index audit для `memory-bank/`.

### Non-Scope

- `NS-01` Не создавать instantiated epic package в `memory-bank/epics/`.
- `NS-02` Не менять feature-flow semantics за пределами документов, необходимых для ведения этой feature.
- `NS-03` Не добавлять roadmap/subissue/risk behavior в epic brief template.
- `NS-04` Не переносить product/domain/project-specific content из downstream repositories.

### Constraints / Assumptions

- `ASM-01` Source template from `dapi/zelma` is used only as evidence for shape and intent, not as authoritative downstream content.
- `ASM-02` Current repository has no runtime application; delivery is Markdown/governance update plus link audit.
- `CON-01` Full epic package remains authoritative for roadmap, subissues, risks and local epic decisions per `memory-bank/flows/epic-flow.md`.
- `CON-02` Governed docs in `memory-bank/` must keep YAML frontmatter with required `status` and valid relative links.
- `CON-03` Target template/governance docs must remain generic and not mention `zelma`-specific terms.

## Design Requirement Decision

| Decision | Reason | Downstream owner |
| --- | --- | --- |
| `Design required: yes` | The change updates governed templates and `epic-flow.md` boundary semantics; selected boundaries and local decisions must not be hidden in `brief.md` or execution plan. | `design.md` |

## Verify

### Exit Criteria

- `EC-01` Epic brief template exists and includes problem, outcome, rough scope, non-scope and readiness notes.
- `EC-02` Epic template indexes route to the new brief template.
- `EC-03` `epic-flow.md` preserves full package authority for charter/roadmap/subissues/risks and excludes implementation sequencing from intake brief.
- `EC-04` Target template/governance docs contain no source-project-specific `zelma` content.
- `EC-05` `python3 scripts/check_memory_bank_index.py` passes.

### Traceability matrix

| Requirement ID | Problem refs | Acceptance refs | Checks | Evidence IDs |
| --- | --- | --- | --- | --- |
| `REQ-01` | `ASM-01`, `CON-03` | `EC-01`, `SC-01`, `NEG-01` | `CHK-01`, `CHK-02`, `CHK-03` | `EVID-01`, `EVID-02`, `EVID-03` |
| `REQ-02` | `CON-02` | `EC-02`, `SC-01` | `CHK-01`, `CHK-02` | `EVID-01`, `EVID-02` |
| `REQ-03` | `CON-02` | `EC-02` | `CHK-01`, `CHK-02` | `EVID-01`, `EVID-02` |
| `REQ-04` | `CON-01` | `EC-03`, `SC-02` | `CHK-02` | `EVID-02` |
| `REQ-05` | `ASM-01`, `CON-03` | `EC-04`, `NEG-01` | `CHK-03` | `EVID-03` |
| `REQ-06` | `ASM-02`, `CON-02` | `EC-05` | `CHK-01`, `CHK-04` | `EVID-01`, `EVID-04` |

### Acceptance Scenarios

- `SC-01` A reader creating an early epic proposal can find `memory-bank/flows/templates/epic/brief.md` and instantiate a lightweight document with problem, outcome, rough scope, non-scope and readiness notes before full epic setup.
- `SC-02` A reader of `epic-flow.md` can tell that full epic package owners remain authoritative for roadmap, accepted subissues, risks and decisions, while feature execution still belongs to `memory-bank/features/FT-<issue>/`.

### Negative / Edge Coverage

- `NEG-01` If source-project terms or implementation sequence appear in target epic brief template/governance docs, the change is rejected until those terms are removed or moved to the correct owner.

### Checks

| Check ID | Covers | How to check | Expected result | Evidence path |
| --- | --- | --- | --- | --- |
| `CHK-01` | `EC-01`, `EC-02`, `EC-05` | `python3 scripts/check_memory_bank_index.py` | Audit reports `Result: OK` | `artifacts/ft-014/verify/chk-01/` |
| `CHK-02` | `EC-01`, `EC-02`, `EC-03` | Review changed docs: `memory-bank/flows/templates/epic/brief.md`, `memory-bank/flows/templates/epic/README.md`, `memory-bank/flows/templates/README.md`, `memory-bank/flows/epic-flow.md` | Required sections and boundary wording are present and not contradictory | `artifacts/ft-014/verify/chk-02/` |
| `CHK-03` | `EC-04`, `NEG-01` | `rg -n "zelma|Zelma|dapi/zelma" memory-bank/flows/templates/epic/brief.md memory-bank/flows/templates/epic/README.md memory-bank/flows/templates/README.md memory-bank/flows/epic-flow.md` | No matches | `artifacts/ft-014/verify/chk-03/` |
| `CHK-04` | `EC-05` | `git diff --check` | No whitespace errors or conflict markers | `artifacts/ft-014/verify/chk-04/` |

### Test matrix

| Check ID | Evidence IDs | Evidence path |
| --- | --- | --- |
| `CHK-01` | `EVID-01` | `artifacts/ft-014/verify/chk-01/` |
| `CHK-02` | `EVID-02` | `artifacts/ft-014/verify/chk-02/` |
| `CHK-03` | `EVID-03` | `artifacts/ft-014/verify/chk-03/` |
| `CHK-04` | `EVID-04` | `artifacts/ft-014/verify/chk-04/` |

### Evidence

- `EVID-01` Output of `python3 scripts/check_memory_bank_index.py`.
- `EVID-02` Review note or diff evidence showing required template/index/flow wording.
- `EVID-03` Output proving no source-project terms in target template/governance docs.
- `EVID-04` Output of `git diff --check`.

### Evidence contract

| Evidence ID | Artifact | Producer | Path contract | Reused by checks |
| --- | --- | --- | --- | --- |
| `EVID-01` | Link audit output | implementer | `artifacts/ft-014/verify/chk-01/` | `CHK-01` |
| `EVID-02` | Review note or relevant diff excerpt | implementer / reviewer | `artifacts/ft-014/verify/chk-02/` | `CHK-02` |
| `EVID-03` | `rg` output | implementer | `artifacts/ft-014/verify/chk-03/` | `CHK-03` |
| `EVID-04` | `git diff --check` output | implementer | `artifacts/ft-014/verify/chk-04/` | `CHK-04` |

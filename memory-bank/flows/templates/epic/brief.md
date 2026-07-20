---
title: "EP-XXX: Epic Proposal Template"
doc_kind: governance
doc_function: template
purpose: "Wrapper-шаблон Epic Intake brief: фиксирует proposal до canonical epic setup и управляет его disposition/promotion."
derived_from:
  - ../../epic.md
  - ../../../dna/frontmatter.md
status: active
audience: humans_and_agents
template_for: epic
template_target_path: ../../../epics/EP-XXX/brief.md
---

# EP-XXX: Epic Proposal Template

Используй этот template, когда Task Routing уже выбрал Epic Flow, но facts ещё недостаточны для `charter.md`. `brief.md` временно владеет только intake facts. Он не заменяет canonical epic owners и после approved promotion становится archived historical context.

## Instantiated Frontmatter

```yaml
---
title: "EP-XXX: <Epic Name> Proposal"
doc_kind: epic
doc_function: proposal
purpose: "Epic Intake proposal: source, problem, outcome, rough boundaries, candidate slices, open questions and disposition."
derived_from:
  - ../../flows/epic.md
status: draft
proposal_status: pending
audience: humans_and_agents
must_not_define:
  - roadmap_waves
  - accepted_subissues
  - risk_controls
  - selected_solution
  - feature_acceptance_contracts
  - implementation_sequence
---
```

## Instantiated Body

```markdown
# EP-XXX: <Epic Name> Proposal

## Intake

| Field | Value |
| --- | --- |
| Source / trigger | `<issue, request, PRD or evidence URL>` |
| Proposal owner | `<person or role>` |
| Decision owner | `<person or role authorized to choose disposition>` |
| Created | `<YYYY-MM-DD>` |

## Problem

Какой gap, constraint или opportunity требует рассмотрения.

## Observable Outcome

Какой проверяемый результат должен появиться, если proposal будет одобрен и delivered.

## Why Epic

Какие признаки требуют Epic Flow: несколько delivery units, общий roadmap, cross-feature risks или shared governance.

## Rough Scope

- `BR-REQ-01` Что предварительно входит в инициативу.

## Rough Non-Scope

- `BR-NS-01` Что предварительно исключено.

## Candidate Delivery Slices

| Candidate | Outcome | Dependency / shared concern | Evidence |
| --- | --- | --- | --- |
| `BR-SLICE-01` | `<candidate outcome>` | `<why it cannot be governed independently yet>` | `<source>` |

Это candidates, а не accepted subissues или delivery feature packages. Не присваивай им `EP-SI-*` или `FT-*` до соответствующих epic gates.

## Evidence and Open Questions

### Available Evidence

| Evidence | Supports | Confidence / freshness |
| --- | --- | --- |

### Open Questions

| Question | Blocks | Owner | Resolution evidence |
| --- | --- | --- | --- |

## Disposition

| Field | Value |
| --- | --- |
| Proposal status | `pending` |
| Decision | `pending / approved / rerouted / parked / rejected` |
| Decision owner | `<person or role>` |
| Decision reference | `<issue comment, meeting note, ADR or other evidence>` |
| Rationale | `<why this disposition was selected>` |
| Target route or review trigger | `<required for rerouted or parked>` |

При disposition обнови и `proposal_status` во frontmatter, и эту таблицу. Для `approved`, `rerouted` и `rejected` установи `status: archived` только после выполнения соответствующего handoff contract.

## Promotion Map

Заполняй при `approved`. Каждый promoted fact получает одного canonical owner; brief не остаётся вторым active owner.

| Intake facts | Canonical owner | Resulting IDs / links |
| --- | --- | --- |
| Problem, outcome, scope/non-scope | `charter.md` | `<REQ-*, NS-* or section links>` |
| Candidate slices and dependencies | `roadmap.md` / `subissues.md` | `<SLICE-*, EP-SI-* when accepted>` |
| Material risks | `risks.md` | `<ERISK-*>`; controls are newly established in `risks.md`, not promoted from this brief |
| Material epic-local decisions | `decision-log.md` | `<DL-*>` |

## Boundary Check

- [ ] No roadmap waves or implementation sequence are defined here.
- [ ] No accepted subissues or delivery `FT-*` packages were created from this proposal.
- [ ] `README.md` links this brief and shows the same lifecycle stage.
- [ ] Approved facts are promoted to canonical owners before `brief.md` is archived.
```

---
title: "FT-013: Workflow Routing Metrics And Decision Log"
doc_kind: feature
doc_function: canonical
purpose: "Canonical brief для issue 13. Фиксирует problem space, scope и verify contract для generic workflow routing metrics, decision log и developer brief."
derived_from:
  - ../../flows/feature-flow.md
  - ../../flows/workflows.md
  - ../../engineering/testing-policy.md
status: active
delivery_status: in_progress
audience: humans_and_agents
must_not_define:
  - implementation_sequence
  - solution_space
---
# FT-013: Workflow Routing Metrics And Decision Log

## What

### Problem

Если в generic memory-bank вводятся compact workflow profiles, нужен проверяемый способ понять, снизили ли они лишнюю documentation ceremony без потери safety, evidence и traceability. Issue 13 указывает на source pattern из `alfagen/mercury`: workflow decision log, workflow metrics и workflow routing developer brief. Текущий шаблон уже имеет [Task Workflows](../../flows/workflows.md), но не имеет generic optional документов для фиксации причин изменения selector-а, измерения safety-first outcomes и краткого объяснения workflow profiles разработчикам.

Feature должна адаптировать source pattern в generic шаблон: сохранить полезную структуру, но не переносить source repo names, project-specific dates, operational details или локальные команды downstream-проекта.

### Outcome

| Metric ID | Metric | Baseline | Target | Measurement method |
| --- | --- | --- | --- | --- |
| `MET-01` | Safety-first workflow evaluation | Метрик workflow selector-а в шаблоне нет | Метрики явно оценивают safety/evidence/rework/traceability раньше speed | Review `workflow-metrics.md` decision rule и metric cards |
| `MET-02` | Decision rationale auditability | Причины изменений selector-а не имеют dedicated generic owner-а | Decision log объясняет зачем появился selector, какие решения приняты и как их пересматривать | Review `workflow-decision-log.md` structure |
| `MET-03` | Developer routing clarity | `workflows.md` описывает типы workflow, но нет короткого developer brief/template для workflow profiles | Короткий developer-facing doc объясняет выбор profiles и escalation/promotion triggers | Review `workflow-routing-developer-brief.md` |
| `MET-04` | Navigation reachability | Optional workflow docs отсутствуют в `flows/README.md` | Новые optional docs reachable из `flows/README.md` и согласованы с `workflows.md` | `python3 scripts/check_memory_bank_index.py` |

### Scope

- `REQ-01` Добавить generic workflow decision log для причин изменения workflow selector-а, ожидаемых последствий и review/update правил.
- `REQ-02` Добавить workflow metrics template/document, покрывающий routing coverage, safety misroute, missing evidence, rework и traceability; speed может быть только вторичным измерением.
- `REQ-03` Добавить короткий developer brief или template, объясняющий выбор workflow profiles и когда нужно повышать задачу до более строгого flow.
- `REQ-04` Связать новые optional docs с `memory-bank/flows/README.md` и `memory-bank/flows/workflows.md` без broken links.
- `REQ-05` Зафиксировать связь с task-flow issue 12 как зависимость/adjacent scope, не реализуя `task-flow`, `bugfix-flow` или `refactor-flow` в рамках FT-013.
- `REQ-06` Адаптировать source pattern generic-образом: не переносить source repo names, project-specific dates, operational commands или downstream-specific details в reusable docs.

### Non-Scope

- `NS-01` Не реализовывать issue 12: `task-flow.md`, `bugfix-flow.md`, `refactor-flow.md`, `TASK-XXX` templates и `memory-bank/tasks/README.md`.
- `NS-02` Не менять `feature-flow.md` lifecycle, stable IDs или feature package templates.
- `NS-03` Не вводить project-level product KPI в `memory-bank/product/metrics.md`; workflow metrics остаются process-layer guidance.
- `NS-04` Не задавать fixed pilot dates, repository-specific thresholds или operational commands как generic defaults.
- `NS-05` Не мигрировать существующие tasks/features и не создавать downstream project examples.

### Constraints / Assumptions

- `ASM-01` Source docs from `alfagen/mercury` are evidence for reusable structure, not content to copy verbatim.
- `ASM-02` Issue 12 is the owner for compact task-flow implementation; FT-013 may reference that dependency but must not create its artifacts.
- `ASM-03` Current repo has no local `task-flow.md`, `bugfix-flow.md` or `refactor-flow.md`; implementation must avoid links to absent local files.
- `CON-01` Generic memory-bank must not receive source project names, source pilot dates, local operational commands or downstream domain details.
- `CON-02` New workflow docs are optional guidance, but must be reachable from `flows/README.md`.
- `CON-03` Metrics must measure safety/evidence before speed.
- `CON-04` Review and verify must use existing lightweight repository checks: link audit, whitespace/conflict-marker check and targeted leak scan.

No unresolved blocking `DEC-*` remain after feature-local decisions in [decision-log.md](decision-log.md).

## Design Requirement Decision

| Decision | Reason | Downstream owner |
| --- | --- | --- |
| `Design required: yes` | The feature chooses owners/locations for new governance docs, resolves the open task-flow dependency, and sets safety-first metric semantics. These are solution-space decisions, not execution steps. | `design.md` |

## Verify

### Exit Criteria

- `EC-01` `memory-bank/flows/workflow-decision-log.md` exists, is generic, and explains why selector changes are introduced, how decisions are reviewed, and what must be recorded.
- `EC-02` `memory-bank/flows/workflow-metrics.md` exists and includes routing coverage, safety misroute, missing evidence, rework, traceability and a decision rule where safety gates precede speed.
- `EC-03` `memory-bank/flows/workflow-routing-developer-brief.md` exists and gives a short, reusable explanation of workflow profile choice and promotion/escalation triggers.
- `EC-04` `memory-bank/flows/README.md` and `memory-bank/flows/workflows.md` make the optional docs reachable without linking to absent local task-flow files.
- `EC-05` Generic docs do not contain source project names, source pilot dates, source operational commands or downstream-specific implementation details.
- `EC-06` `python3 scripts/check_memory_bank_index.py` and `git diff --check` pass.

### Traceability matrix

| Requirement ID | Problem refs | Acceptance refs | Checks | Evidence IDs |
| --- | --- | --- | --- | --- |
| `REQ-01` | `ASM-01`, `CON-01`, `CON-02` | `EC-01`, `SC-02` | `CHK-01`, `CHK-03`, `CHK-04` | `EVID-01`, `EVID-03`, `EVID-04` |
| `REQ-02` | `CON-03` | `EC-02`, `SC-02`, `NEG-01` | `CHK-01`, `CHK-04` | `EVID-01`, `EVID-04` |
| `REQ-03` | `ASM-01`, `ASM-03` | `EC-03`, `SC-01` | `CHK-01`, `CHK-04` | `EVID-01`, `EVID-04` |
| `REQ-04` | `CON-02`, `ASM-03` | `EC-04`, `SC-01` | `CHK-01` | `EVID-01` |
| `REQ-05` | `ASM-02`, `ASM-03` | `EC-04`, `SC-01` | `CHK-01`, `CHK-04` | `EVID-01`, `EVID-04` |
| `REQ-06` | `ASM-01`, `CON-01` | `EC-05`, `SC-03`, `NEG-02` | `CHK-02`, `CHK-03`, `CHK-04` | `EVID-02`, `EVID-03`, `EVID-04` |

### Acceptance Scenarios

- `SC-01` Developer or agent starts from `flows/README.md` or `workflows.md`, finds optional workflow routing docs, and can understand profile choice and escalation without needing source repo context.
- `SC-02` Maintainer reviews a workflow selector change and uses decision log plus metrics to evaluate whether compact profiles improved ceremony without weakening safety/evidence.
- `SC-03` Reviewer checks the delivered generic docs and finds no source repo names, source pilot dates, source operational commands or downstream-specific assumptions.

### Negative / Edge Scenarios

- `NEG-01` If lead time improves but safety misroutes, missing evidence or rework get worse, workflow metrics must not classify the selector change as successful.
- `NEG-02` If source docs include project-specific names, dates or commands, delivered generic docs must omit or replace them with generic placeholders/rules.
- `NEG-03` If issue 12 is not implemented, FT-013 docs must still pass link audit by avoiding local links to absent `task-flow.md`, `bugfix-flow.md` or `refactor-flow.md`.

### Checks

| Check ID | Covers | How to check | Expected result | Evidence path |
| --- | --- | --- | --- | --- |
| `CHK-01` | `EC-01`, `EC-02`, `EC-03`, `EC-04`, `NEG-03` | `python3 scripts/check_memory_bank_index.py` | No broken links, orphan docs, unreachable docs or index contract failures | `artifacts/ft-013/verify/chk-01/` |
| `CHK-02` | `EC-06` | `git diff --check` | No trailing whitespace or conflict markers | `artifacts/ft-013/verify/chk-02/` |
| `CHK-03` | `EC-05`, `NEG-02` | `rg -n "Mercury|alfagen|2026-06-29|2026-07-28|2026-07-29|dip|Rails|Rake|RSpec|MySQL" memory-bank/flows/workflow-*.md` | No matches in delivered generic workflow docs | `artifacts/ft-013/verify/chk-03/` |
| `CHK-04` | `EC-01`..`EC-05` | Manual review against issue 13, issue 12 dependency, `design.md` traceability and source pattern summary | Reviewer can trace every delivered doc to `REQ-*` and sees task-flow dependency handled without scope expansion | `artifacts/ft-013/verify/chk-04/` |

### Test matrix

| Check ID | Evidence IDs | Evidence path |
| --- | --- | --- |
| `CHK-01` | `EVID-01` | `artifacts/ft-013/verify/chk-01/` |
| `CHK-02` | `EVID-02` | `artifacts/ft-013/verify/chk-02/` |
| `CHK-03` | `EVID-03` | `artifacts/ft-013/verify/chk-03/` |
| `CHK-04` | `EVID-04` | `artifacts/ft-013/verify/chk-04/` |

### Evidence

- `EVID-01` Link audit output proving new docs are reachable and internally valid.
- `EVID-02` Whitespace/conflict-marker check output.
- `EVID-03` Targeted leak scan output for source project names, dates and operational commands in delivered workflow docs.
- `EVID-04` Reviewer note or PR description section mapping issue 13 acceptance to delivered docs.

### Evidence contract

| Evidence ID | Artifact | Producer | Path contract | Reused by checks |
| --- | --- | --- | --- | --- |
| `EVID-01` | Link audit command output | implementer | `artifacts/ft-013/verify/chk-01/` or PR check log | `CHK-01` |
| `EVID-02` | `git diff --check` output | implementer | `artifacts/ft-013/verify/chk-02/` or PR check log | `CHK-02` |
| `EVID-03` | Leak scan output | implementer | `artifacts/ft-013/verify/chk-03/` or PR check log | `CHK-03` |
| `EVID-04` | Manual traceability review note | reviewer / implementer | `artifacts/ft-013/verify/chk-04/` or PR body | `CHK-04` |

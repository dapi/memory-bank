---
title: "EP-XXX: Brief Template"
doc_kind: governance
doc_function: template
purpose: "Wrapper-шаблон для lightweight epic intake brief. Используется до полного epic setup и не заменяет charter, roadmap, subissues, risks или decision log."
derived_from:
  - ../../epic.md
status: active
audience: humans_and_agents
template_for: epic
template_target_path: ../../../epics/EP-XXX/brief.md
---

# EP-XXX: Brief Template

## Wrapper Notes

Используй этот template, когда инициативу нужно быстро зафиксировать как epic proposal до полного package setup. `brief.md` помогает записать problem, outcome, rough scope, non-scope и readiness notes, но не становится authoritative owner-ом полного epic.

Когда initiative готова к full setup, создай canonical `charter.md`, `roadmap.md`, `subissues.md`, `risks.md` и, если нужны локальные решения, `decision-log.md`. Если intake brief и full epic owner конфликтуют, full owner должен быть обновлен или создан, а brief остается только intake context.

## Instantiated Frontmatter

```yaml
---
title: "EP-XXX: Brief"
doc_kind: epic
doc_function: brief
purpose: "Lightweight intake brief для epic proposal: problem, outcome, rough scope, non-scope и readiness notes до полного epic setup."
derived_from:
  - ../../flows/epic.md
status: draft
audience: humans_and_agents
must_not_define:
  - implementation_sequence
  - roadmap_waves
  - accepted_subissues
  - risk_controls
  - selected_solution
  - feature_acceptance_contracts
---
```

## Instantiated Body

```markdown
# EP-XXX: Brief

## Problem

Почему эта инициатива нужна и какой gap она закрывает.

## Outcome

Какой наблюдаемый результат должен появиться после исполнения epic.

## Rough Scope

- `BR-REQ-01` Что предварительно входит в инициативу.

## Non-Scope

- `BR-NS-01` Что не должно попадать в инициативу.

## Candidate Feature Briefs

- [FT-XXX: Feature Name](../../features/FT-XXX/README.md)

## Readiness Notes

- Что нужно уточнить до полного `charter.md` / `roadmap.md` / `subissues.md` / `risks.md`.
- Какие source materials, stakeholder answers or constraints still need evidence.

## Full Epic Package Handoff

Before execution, promote or supersede intake facts in the full epic package:

- `charter.md` owns authoritative problem, outcome, scope, non-scope and evidence boundaries.
- `roadmap.md` owns waves, dependencies, gates and stop rules.
- `subissues.md` owns accepted delivery subissues.
- `risks.md` owns epic-level risks and controls.
- Feature execution belongs to `memory-bank/features/FT-<issue>/`.
```

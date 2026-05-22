---
title: "FT-XXX: Design Template"
doc_kind: feature
doc_function: template
purpose: Governed wrapper-шаблон для feature-local `design.md`. Фиксирует solution-space слой: выбранный подход, rationale, contracts, failure modes и design-pack routing без смешения с problem space или execution contract.
derived_from:
  - ../../feature-flow.md
  - ../../../dna/frontmatter.md
status: active
audience: humans_and_agents
template_for: feature
template_target_path: ../../../features/FT-XXX/design.md
canonical_for:
  - feature_design_template
---

# FT-XXX: Design

Этот файл описывает wrapper-template. Инстанцируемый `design.md` живет ниже как embedded contract и копируется без wrapper frontmatter и history.

## Wrapper Notes

Создавай `design.md`, когда фича требует solution-space reasoning: выбор подхода, trade-offs, contracts, invariants, failure modes, rollout/backout, ADR/C4/data-flow/diagram dependencies или design-pack из нескольких документов.

На стадии анализа обязательно заполни C4 applicability decision. C4 artifact обязателен только когда trigger из [feature-flow.md#c4-analysis-requirements](../../feature-flow.md#c4-analysis-requirements) требует C1/C2/C3/C4; для local feature достаточно `C4-00 not required` с причиной.

`design.md` не заменяет `brief.md`: требования, acceptance criteria и evidence contract остаются в `brief.md`. `design.md` также не является execution plan: file-level touchpoints, атомарные шаги, команды тестов и checkpoints принадлежат `implementation-plan.md`.

Если solution-space разбит на несколько артефактов, `design.md` становится индексом design-pack и фиксирует owner-а каждого design fact. Не дублируй canonical факты из ADR, C4, data-flow или других design docs; ссылайся на них.

## Instantiated Frontmatter

```yaml
title: "FT-XXX: Design"
doc_kind: feature
doc_function: canonical
purpose: "Solution-space документ для FT-XXX. Фиксирует выбранный подход, rationale, contracts, failure modes и design-pack routing без переопределения problem space или execution contract."
derived_from:
  - brief.md
status: draft
audience: humans_and_agents
must_not_define:
  - ft_xxx_scope
  - ft_xxx_acceptance_criteria
  - ft_xxx_evidence_contract
  - implementation_sequence
```

## Instantiated Body

```markdown
# FT-XXX: Design

## Design Pack

Если design-pack состоит только из этого файла, оставь одну строку `design.md`. Если есть ADR, C4, data-flow, диаграммы или contract notes, добавь их в таблицу и укажи canonical owner.

| Artifact | Role | Owns |
| --- | --- | --- |
| `design.md` | Feature-local solution owner | `SOL-*`, `ALT-*`, `TRD-*`, `C4-*`, feature-local `CTR-*`, `INV-*`, `FM-*`, `RB-*` |
| `../../adr/ADR-XXX.md` | Architecture decision | Какой design choice принадлежит ADR |

## Context

Коротко опиши design problem: почему требования из `brief.md` требуют явного решения, какие upstream docs или constraints важны для выбора.

## C4 Applicability

Решение принимается до `Solution Ready`. Выбери минимальный уровень C4 или явно зафиксируй, что C4 не нужен.

| C4 ID | Decision | Trigger / reason | Artifact |
| --- | --- | --- | --- |
| `C4-00` | `not required` / `C1` / `C2` / `C3` / `C4` | Почему C4 не нужен или какой trigger требует выбранный уровень | `none` / ссылка на diagram |

### C4 Artifact

Если `C4-00` не `not required`, добавь diagram или ссылку на artifact design-pack. Используй самый низкий достаточный уровень:

- `C1` - System Context: actors/external systems/trust boundaries.
- `C2` - Container: deployable/runtime nodes, queues, stores, protocols.
- `C3` - Component: modules/services/state machines внутри container.
- `C4` - Code: только когда class/interface-level structure является архитектурным решением.

## Selected Solution

- `SOL-01` Выбранный элемент решения и почему он закрывает `REQ-*`.
- `SOL-02` Второй элемент решения, если нужен.

## Alternatives Considered

| Alternative ID | Option | Why not selected |
| --- | --- | --- |
| `ALT-01` | Альтернативный подход | Причина отказа или отложенного выбора |

## Trade-offs

| Trade-off ID | Decision | Benefit | Cost / Risk |
| --- | --- | --- | --- |
| `TRD-01` | Какой компромисс принимаем | Что выигрываем | Что платим или мониторим |

## Accepted Local Decisions

Здесь живут только принятые feature-local decisions. Decisions reusable, architectural или cross-feature уровня выносятся в ADR.

- `SD-01` Какое локальное решение принято и почему оно не требует ADR.

## Contracts

Контракты описывай на уровне shape/semantics. Не добавляй реалистичные секреты, production IDs или file-level implementation steps.

| Contract ID | Input / Output | Producer / Consumer | Semantics / Constraints |
| --- | --- | --- | --- |
| `CTR-01` | Что меняется | Кто пишет / кто читает | Что должно оставаться истинным |

## Invariants

- `INV-01` Что должно оставаться истинным независимо от implementation path.

## Failure Modes

- `FM-01` Что может пойти не так и как решение должно это ограничить.

## Rollout / Backout

| Stage ID | Stage | Entry condition | Backout |
| --- | --- | --- | --- |
| `RB-01` | Как включается изменение | Что должно быть доказано до входа | Как вернуть безопасное состояние |

## ADR / External Design Dependencies

| Artifact | Current status | Used for | Rule |
| --- | --- | --- | --- |
| `../../adr/ADR-XXX.md` | `proposed` / `accepted` | Какой выбор или baseline задает | `proposed` не считается finalized design |

## Traceability

| Requirement ID | Solution refs | Contracts / invariants | Failure / rollout refs |
| --- | --- | --- | --- |
| `REQ-01` | `SOL-01`, `TRD-01`, `C4-00`, `SD-01` | `CTR-01`, `INV-01` | `FM-01`, `RB-01` |
```

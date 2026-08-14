---
title: "FT-XXX: Feature Use Cases Template"
doc_kind: feature-support
doc_function: template
purpose: Governed wrapper-шаблон optional feature-local `use-cases/README.md`. Читать, когда feature needs review-friendly use cases, BDD example mapping и derived test candidates без переноса canonical acceptance из `brief.md`.
derived_from:
  - ../../../feature.md
  - ../../../behavior-specification.md
  - ../../../feature-artifact-catalog.md
  - ../../../../dna/frontmatter.md
status: active
audience: humans_and_agents
template_for: feature-support
template_target_path: ../../../../features/FT-XXX/use-cases/README.md
canonical_for:
  - feature_support_template_use_cases
---

# FT-XXX: Feature Use Cases Template

Этот файл описывает wrapper-template. Инстанцируемый `use-cases/README.md` живет внутри feature package как optional derived companion.

## Wrapper Notes

Создавай feature-local `use-cases/README.md`, если scenario set становится сложным для review: много happy/edge/error cases, несколько user roles или нужен удобный `FUC → SC/NEG → REQ → CHK` mapping.

Этот документ не подменяет canonical `SC-*`, `NEG-*`, `CHK-*` и `EVID-*` из `brief.md`.
BDD examples здесь являются derived projection: они не могут вводить новый
context, event, outcome или verdict, отсутствующий в canonical `brief.md`.

## Instantiated Frontmatter

```yaml
title: "FT-XXX: Feature Use Cases"
doc_kind: feature-support
doc_function: reference
purpose: "Derived use-case companion для FT-XXX. Упаковывает use cases, BDD example mapping и test candidates для review без переопределения canonical acceptance inventory."
derived_from:
  - ../brief.md
  # Required only when design.md exists:
  # - ../design.md
status: draft
audience: humans_and_agents
must_not_define:
  - ft_xxx_scope
  - ft_xxx_acceptance_criteria
  - canonical_checks
  - implementation_sequence
```

## Instantiated Body

```markdown
# FT-XXX: Feature Use Cases

## Role

Этот документ дает review-friendly projection canonical facts из `brief.md` и existing `design.md`.

Canonical acceptance / test inventory остается в `brief.md` через `SC-*`, `NEG-*`, `CHK-*` и `EVID-*`.

## Rule And Example Map

Проецируй только существующие canonical refs. Verdict-changing question должен
быть записан как `DEC-*` в `brief.md`, а не решён в этом companion.

| Rule refs | Positive examples | Negative / edge examples | Open question refs |
| --- | --- | --- | --- |
| `UC-XXX/BR-01`, `REQ-01` | `SC-01` | `NEG-01` | `DEC-01` / `none` |

## Happy Path

| ID | Use case | Given | When | Then | Primary refs |
| --- | --- | --- | --- | --- | --- |
| `FUC-H01` | Название сценария | Существенный context из `SC-01` | Событие из `SC-01` | Observable outcome из `SC-01` | `REQ-01`, `SC-01`, `CHK-01` |

## Edge Cases

| ID | Use case | Given | When | Then | Primary refs |
| --- | --- | --- | --- | --- | --- |
| `FUC-E01` | Название edge case | Boundary context | Событие | Observable boundary outcome | `REQ-01`, `SC-01`, `CHK-01` |

## Error Cases

| ID | Use case | Given | When | Then | Primary refs |
| --- | --- | --- | --- | --- | --- |
| `FUC-ER01` | Название error case | Failure context | Событие | Observable rejection, fallback или preserved state | `NEG-01`, `CHK-02`, `FM-01` / `none` |

## Interface Use Cases

Заполняй только если feature меняет interface. Подробности screen design остаются в `ui-reference/README.md`.

| ID | Use case | Description | Primary refs |
| --- | --- | --- | --- |
| `FUC-UI01` | Пользователь проходит interface flow | Какой interface outcome нужен | `REQ-02`, `UI-01`, `SC-02` |

## Derived Test Case Candidates

`TC-*` здесь являются candidates для planning/review и должны ссылаться на canonical `CHK-*`, а не создавать новые checks.
Они не считаются automation plan или test implementation owner.

| Test Case ID | Covers | Preconditions | Steps | Expected result | Automation candidate |
| --- | --- | --- | --- | --- | --- |
| `TC-01` | `FUC-H01`, `SC-01`, `CHK-01` | Что должно быть готово | Короткая процедура | Какой outcome ожидается | automated / manual / mixed |

## Traceability Matrix

| Use case ID | Requirements | Acceptance refs | Check IDs | Notes |
| --- | --- | --- | --- | --- |
| `FUC-H01` | `REQ-01` | `SC-01` | `CHK-01` | Что важно при review |

## Test Ownership

### Automated

- Какие use cases должны закрываться automated checks.

### Manual

- Какие use cases остаются manual-only и почему; каждая строка должна ссылаться на canonical `CHK-*`, `EVID-*` и approval ref из плана, если нужен approval.
```

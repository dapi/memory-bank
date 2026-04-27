---
title: "FT-XXX: Feature Template - Short"
doc_kind: feature
doc_function: template
purpose: Governed wrapper-шаблон для короткого canonical `feature.md`, который фиксирует только problem space и canonical verify contract.
derived_from:
  - ../../feature-flow.md
  - ../../../dna/frontmatter.md
  - ../../../engineering/testing-policy.md
status: active
audience: humans_and_agents
template_for: feature
template_target_path: ../../../features/FT-XXX/feature.md
canonical_for:
  - feature_template_short
---

# FT-XXX: Feature Name

Этот файл описывает wrapper-template. Инстанцируемый `feature.md` живет ниже как embedded contract и копируется без wrapper frontmatter и history.

## Wrapper Notes

Используй этот шаблон, только если problem space фичи остается компактным и её можно описать через `REQ-*`, `NS-*`, один `SC-*`, максимум один `CON-*`, один `EC-*`, один `CHK-*` и один `EVID-*`.

Если тебе нужны `ASM-*`, `DEC-*`, feature-specific `NEG-*`, больше одного acceptance scenario, больше одного `CHK-*` / `EVID-*` или richer problem-space narrative, сделай upgrade до `large.md` до продолжения работы.

Solution-space complexity больше не делает short feature недопустимым сама по себе: даже для short feature downstream `solution.md` остаётся обязательным canonical owner selected design и to-be C4 architecture model. Значение prefixes зафиксировано в [../../feature-flow.md](../../feature-flow.md#stable-identifiers).

### Frontmatter Quick Ref

Полная schema — в [../../../dna/frontmatter.md](../../../dna/frontmatter.md). Для стандартного feature достаточно:

| Поле | Обязательность | Значения / default |
|---|---|---|
| `title` | required | `"FT-XXX: Name"` |
| `doc_kind` | required | `feature` |
| `doc_function` | required | `canonical` |
| `purpose` | required | 1-2 предложения |
| `status` | required | `draft` → `active` → `archived` |
| `derived_from` | required для active | upstream-документы |
| `delivery_status` | required для feature | `planned` → `in_progress` → `done` / `cancelled` |
| `audience` | recommended | `humans_and_agents` |
| `must_not_define` | recommended | что документ НЕ определяет |

## Instantiated Frontmatter

```yaml
title: "FT-XXX: Feature Name"
doc_kind: feature
doc_function: canonical
purpose: "Короткий canonical feature-документ для небольшой delivery-единицы. Фиксирует только problem space и canonical verify contract."
derived_from:
  - ../../domain/problem.md
  # Optional:
  # - ../../prd/PRD-XXX-short-name.md
  # - ../../use-cases/UC-XXX-short-name.md
status: draft
delivery_status: planned
audience: humans_and_agents
must_not_define:
  - selected_design
  - implementation_sequence
```

## Instantiated Body

```markdown
# FT-XXX: Feature Name

## What

### Problem

Какую конкретную проблему или opportunity закрывает фича.

Если существует upstream PRD, здесь не переписывай весь продуктовый контекст, а сфокусируйся на slice-specific постановке задачи.

Если существует upstream use case, здесь зафиксируй только то, как текущая delivery-единица реализует или меняет этот сценарий.

### Outcome

Коротко опиши ожидаемый результат delivery-единицы.

### Scope

- `REQ-01` Что обязательно входит.
- `REQ-02` Что еще обязательно входит.

### Non-Scope

- `NS-01` Что точно не делаем.

### Constraints

- `CON-01` Какое ограничение problem space задает границы решения.

## Verify

### Exit Criteria

- `EC-01` Что должно быть истинно после реализации.

### Acceptance Scenarios

- `SC-01` Основной happy path и canonical positive test case для этой delivery-единицы.

### Traceability matrix

| Requirement ID | Acceptance refs | Checks | Evidence IDs |
| --- | --- | --- | --- |
| `REQ-01` | `EC-01`, `SC-01` | `CHK-01` | `EVID-01` |
| `REQ-02` | `EC-01`, `SC-01` | `CHK-01` | `EVID-01` |

### Checks

Verify должен быть исполнимым и задавать минимум один explicit test case через `SC-01`.

| Check ID | Covers | How to check | Expected |
| --- | --- | --- | --- |
| `CHK-01` | `EC-01`, `SC-01` | Команда или процедура | Ожидаемый результат |

### Test matrix

| Check ID | Evidence IDs | Evidence path |
| --- | --- | --- |
| `CHK-01` | `EVID-01` | `artifacts/ft-xxx/verify/chk-01/` |

### Evidence

- `EVID-01` Какой артефакт должен остаться после проверки.

### Evidence contract

| Evidence ID | Artifact | Producer | Path contract | Reused by checks |
| --- | --- | --- | --- | --- |
| `EVID-01` | Минимальный verify-артефакт | verify-runner / human | `artifacts/ft-xxx/verify/chk-01/` | `CHK-01` |
```

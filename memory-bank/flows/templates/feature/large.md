---
title: "FT-XXX: Feature Template - Large"
doc_kind: feature
doc_function: template
purpose: Governed wrapper-шаблон для расширенного canonical `feature.md`, который фиксирует richer problem space и canonical verify без ownership solution space.
derived_from:
  - ../../feature-flow.md
  - ../../../dna/frontmatter.md
  - ../../../engineering/testing-policy.md
status: active
audience: humans_and_agents
template_for: feature
template_target_path: ../../../features/FT-XXX/feature.md
canonical_for:
  - feature_template_large
---

# FT-XXX: Feature Name

Этот файл описывает wrapper-template. Инстанцируемый `feature.md` живет ниже как embedded contract и копируется без wrapper frontmatter и history.

## Wrapper Notes

Используй этот шаблон, когда хотя бы одно правило `short.md` перестает выполняться: problem space затрагивает несколько пользовательских или системных аспектов, требует явных assumptions / blockers, нескольких acceptance scenarios или richer verify-слоя.

Выбранный дизайн, accepted local decisions, contracts, failure modes и rollout/backout semantics теперь живут не здесь, а в downstream `solution.md`.

Используй стабильные идентификаторы по taxonomy из [../../feature-flow.md#stable-identifiers](../../feature-flow.md#stable-identifiers).

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
purpose: "Расширенный canonical feature-документ для сложной или многослойной delivery-единицы. Фиксирует problem space и canonical verify."
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

Какой симптом, ограничение или возможность делает фичу нужной. Если общий контекст уже зафиксирован upstream, здесь опиши только feature-specific вопрос delivery.

Если существует upstream PRD, этот раздел фиксирует только feature-specific delta относительно PRD, а не переписывает весь продуктовый документ.

Если существует upstream use case, здесь фиксируется feature-specific изменение или реализация этого сценария, а не весь проектный flow целиком.

### Outcome

Опиши outcome как измеримую таблицу.

Если численный success threshold относится только к этой delivery-единице, фиксируй его здесь. Поднимать threshold upstream стоит только после появления shared owner для нескольких feature.

| Metric ID | Metric | Baseline | Target | Measurement method |
| --- | --- | --- | --- | --- |
| `MET-01` | Что измеряем | От чего стартуем | Что считаем успехом | Как проверяем |

### Scope

- `REQ-01` Что обязательно входит в deliverable.
- `REQ-02` Что еще обязательно входит в deliverable.

### Non-Scope

- `NS-01` Что сознательно исключено.
- `NS-02` Что агент не должен додумывать или реализовывать сам.

### Constraints / Assumptions

- `ASM-01` На что сейчас опираемся.
- `CON-01` Что прямо ограничивает problem space, verify или допустимый класс решений.
- `INV-01` Какой инвариант нельзя нарушать.

### Blocking Decisions

- `DEC-01` Какое решение еще не принято и что именно оно блокирует.

Accepted local decisions не фиксируй здесь: после принятия они переезжают в downstream `solution.md` как `SD-*` или в ADR.

## Verify

`Verify` задает canonical test case inventory для delivery-единицы: positive scenarios через `SC-*`, feature-specific negative coverage через `NEG-*`, executable checks через `CHK-*` и evidence через `EVID-*`.

### Exit Criteria

- `EC-01` Проверяемый признак готовности.
- `EC-02` Еще один обязательный признак готовности.

### Traceability matrix

| Requirement ID | Problem refs | Acceptance refs | Checks | Evidence IDs |
| --- | --- | --- | --- | --- |
| `REQ-01` | `ASM-01`, `CON-01`, `DEC-01`, `INV-01` | `EC-01`, `SC-01` | `CHK-01` | `EVID-01` |
| `REQ-02` | `ASM-01`, `CON-01`, `INV-01` | `EC-02`, `SC-02` | `CHK-01` | `EVID-01` |

### Acceptance Scenarios

- `SC-01` Основной happy path.
- `SC-02` Обязательный real-world или edge scenario.

### Negative / Edge Cases

- `NEG-01` Какая negative или edge ситуация меняет verdict.

### Checks

Verify должен быть исполнимым.

| Check ID | Covers | How to check | Expected result | Evidence path |
| --- | --- | --- | --- | --- |
| `CHK-01` | `EC-01`, `SC-01` | Команда или процедура | Что считаем успехом | Где лежит артефакт |

### Test matrix

| Check ID | Evidence IDs | Evidence path |
| --- | --- | --- |
| `CHK-01` | `EVID-01` | `artifacts/ft-xxx/verify/chk-01/` |

### Evidence

- `EVID-01` Какой артефакт обязан появиться после проверки.

### Evidence contract

| Evidence ID | Artifact | Producer | Path contract | Reused by checks |
| --- | --- | --- | --- | --- |
| `EVID-01` | Лог, отчет, скриншот или sample output | verify-runner / human | `artifacts/ft-xxx/verify/chk-01/` | `CHK-01` |
```

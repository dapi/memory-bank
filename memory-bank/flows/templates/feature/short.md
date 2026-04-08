---
title: "FT-XXX: Feature Template - Short"
doc_kind: feature
doc_function: template
purpose: Governed wrapper-шаблон для короткого canonical `feature.md` в AI-driven development. Читать, чтобы инстанцировать минимальный feature contract без смешения wrapper и целевого frontmatter.
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

Используй этот шаблон, только если фича укладывается в один локальный slice и её можно описать через `REQ-*`, `NS-*`, один `SC-*`, максимум один `CON-*`, один `EC-*`, один `CHK-*` и один `EVID-*`.

Если тебе нужны `ASM-*`, `DEC-*`, `CTR-*`, `FM-*`, feature-specific negative cases, больше одного acceptance scenario, больше одного `CHK-*` / `EVID-*` или явная ADR-dependent design logic, сделай upgrade до `large.md` до продолжения работы. Значение prefixes зафиксировано в [../../feature-flow.md](../../feature-flow.md#stable-identifiers).

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
purpose: "Короткий canonical feature-документ для небольшой и локальной delivery-единицы."
derived_from:
  - ../../domain/problem.md
  # Optional:
  # - ../../prd/PRD-XXX-short-name.md
  # - ../../use-cases/UC-XXX-short-name.md
status: draft
delivery_status: planned
audience: humans_and_agents
must_not_define:
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

### Scope

- `REQ-01` Что обязательно входит.
- `REQ-02` Что еще обязательно входит.

### Non-Scope

- `NS-01` Что точно не делаем.

### Constraints

- `CON-01` Какое ограничение задает границы решения.

## How

### Solution

Один короткий абзац: основной подход и ключевой trade-off.

### Change Surface

| Surface | Why |
| --- | --- |
| `path/or/component` | Почему меняется |

### Flow

1. Вход.
2. Обработка.
3. Выход.

## Verify

### Exit Criteria

- `EC-01` Что должно быть истинно после реализации.

### Acceptance Scenarios

- `SC-01` Основной happy path и canonical positive test case для этой delivery-единицы.

### Traceability matrix

| Requirement ID | Design refs | Acceptance refs | Checks | Evidence IDs |
| --- | --- | --- | --- | --- |
| `REQ-01` | `CON-01` | `EC-01`, `SC-01` | `CHK-01` | `EVID-01` |
| `REQ-02` | `CON-01` | `EC-01`, `SC-01` | `CHK-01` | `EVID-01` |

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

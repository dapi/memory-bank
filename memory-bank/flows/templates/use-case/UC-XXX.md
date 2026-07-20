---
title: "UC-XXX: Use Case Name"
doc_kind: use_case
doc_function: template
purpose: Governed wrapper-шаблон use case. Читать, чтобы инстанцировать канонический пользовательский или операционный сценарий без смешения wrapper-метаданных и frontmatter будущего use case.
derived_from:
  - ../../../dna/governance.md
  - ../../../dna/frontmatter.md
  - ../../../product/context.md
  - ../../use-case.md
status: active
audience: humans_and_agents
template_for: use_case
template_target_path: ../../../use-cases/UC-XXX-short-name.md
canonical_for:
  - use_case_template
---

# UC-XXX: Use Case Name

Этот файл описывает wrapper-template. Инстанцируемый use case живет ниже как embedded contract и копируется без wrapper frontmatter и history.

## Wrapper Notes

Use case фиксирует устойчивый проектный сценарий. Он описывает trigger, preconditions, основной flow, альтернативы и postconditions, но не уходит в implementation sequence, архитектуру или feature-level verify.

Критерии выбора, lifecycle и границы между `UC-*`, `SC-*` и `FUC-*` определяет [`Use Case Flow`](../../use-case.md).

Если сценарий слишком локален и живет только внутри одной delivery-единицы, не поднимай его в `UC-*`: оставь его в `SC-*` у соответствующей feature.

Если сценарий зависит от domain invariant, state transition или domain event, добавь соответствующий документ из `../domain/` в `derived_from`.

## Instantiated Frontmatter

```yaml
title: "UC-XXX: Use Case Name"
doc_kind: use_case
doc_function: canonical
purpose: "Фиксирует устойчивый пользовательский или операционный сценарий проекта."
derived_from:
  - ../flows/use-case.md
  - ../product/context.md
  # Optional:
  # - ../prd/PRD-XXX-short-name.md
  # - ../domain/rules.md
  # - ../domain/states.md
status: draft
audience: humans_and_agents
must_not_define:
  - implementation_sequence
  - architecture_decision
  - feature_level_test_matrix
```

## Instantiated Body

```markdown
# UC-XXX: Use Case Name

## Goal

Какой результат должен получить actor после успешного выполнения сценария.

## Primary Actor

Кто инициирует сценарий: пользователь, оператор, команда, автоматизированный
агент или внешний сервис.

## Trigger

Какое событие или намерение запускает flow.

## Preconditions

- Что должно быть истинно до начала сценария.
- Какие данные, права или состояние системы обязательны.

## Main Flow

1. Первый шаг сценария.
2. Второй шаг сценария.
3. Наблюдаемый результат.

## Alternate Flows / Exceptions

- `ALT-01` Как сценарий ветвится при ожидаемой альтернативе.
- `EX-01` Какой сбой или отказ должен быть корректно обработан.

## Postconditions

- Что истинно после успешного завершения.
- Что остается истинным после неуспешного завершения.

## Business Rules

- `BR-01` Правило, которое обязана соблюдать любая реализация этого сценария.
- `BR-02` Ограничение или policy, которая влияет на flow.

## Operational Contract (Optional)

Заполняй только для operational / agentic сценария, если перечисленные элементы
являются наблюдаемой частью project-level behavior. Не описывай здесь внутреннюю
архитектуру, implementation sequence или конкретные команды runbook-а.

### Observable Status

- Какие statuses/fields публикуются или где находится canonical schema.
- Кто должен одинаково интерпретировать этот contract.

### Handoff

- Какой минимальный payload передается или где находится canonical schema.
- Как получатель определяет, что handoff завершен и пригоден для продолжения.

### Diagnostics And Recovery

- Какие structured diagnostics наблюдаемы при неуспешном flow.
- Какой recovery outcome и terminal state ожидаются; конкретная процедура может
  принадлежать связанному runbook-у.

## Traceability

| Upstream / Downstream | References |
| --- | --- |
| PRD | `PRD-XXX` / `none` |
| Features | `FT-XXX`, `FT-YYY` |
| ADR | `ADR-XXX` / `none` |
| Runbooks / Ops | `../ops/...` / `none` |

## Lifecycle Note (Required When Archived)

- Почему сценарий больше не является active behavior.
- Какой `UC-*` или другой contract заменил его, либо `none`.
```

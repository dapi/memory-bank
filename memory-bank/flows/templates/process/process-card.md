---
title: "PROC-XXX: Compact Process Card"
doc_kind: process
doc_function: template
purpose: Governed wrapper-шаблон для компактной process-card. Читать, чтобы зафиксировать короткий reusable workflow без тяжёлого lifecycle-каркаса.
derived_from:
  - ../../../dna/governance.md
  - ../../../dna/frontmatter.md
  - ../../workflows.md
status: active
audience: humans_and_agents
template_for: process
template_target_path: ../../../processes/PROCESS-XXX-process-card.md
canonical_for:
  - process_template_card
---

# PROC-XXX: Compact Process Card

Этот файл описывает wrapper-template. Инстанцируемый process-card живет ниже как embedded contract и копируется без wrapper frontmatter и history.

## Wrapper Notes

Этот вариант нужен, когда процесс повторяется часто, но не требует полноценного protocol: один trigger, понятный owner, короткий список шагов и ясные exit criteria.

Хороший кандидат для этого шаблона:

- короткий ручной workflow;
- операционная рутина;
- повторяемый internal step без сложных gates;
- процесс, который удобно описывать на одной странице.

Если процесс начинает требовать handoff state, approval gates, rollback или явные verification phase, это сигнал перейти к `session-handoff.md` или `lifecycle-protocol.md`.

## Instantiated Frontmatter

```yaml
title: "PROC-XXX: Compact Process Card"
doc_kind: process
doc_function: canonical
purpose: "Фиксирует короткий reusable workflow с одним trigger, owner, шагами и exit criteria."
derived_from:
  - README.md
status: draft
audience: humans_and_agents
must_not_define:
  - full_delivery_lifecycle
  - approval_gates
  - rollback_protocol
```

## Instantiated Body

```markdown
# PROC-XXX: Compact Process Card

## Purpose

Коротко опиши, зачем этот процесс существует и какой результат он должен стабильно давать.

## Trigger

- Что запускает процесс.
- Кто его инициирует.
- Какие входные данные нужны перед стартом.

## Scope

### In Scope

- Что этот workflow делает.

### Out Of Scope

- Что он сознательно не покрывает.

## Steps

1. Шаг 1.
2. Шаг 2.
3. Шаг 3.

## Exit Criteria

- Что должно быть истинно, чтобы процесс считался завершенным.

## Evidence

- Какой артефакт, лог, ссылка или статус подтверждает выполнение.

## Escalation

- Когда процесс нужно остановить и поднять к человеку.
```

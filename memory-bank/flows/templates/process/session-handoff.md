---
title: "PROC-XXX: Session Handoff"
doc_kind: process
doc_function: template
purpose: Governed wrapper-шаблон для session handoff. Читать, чтобы сохранять состояние процесса между сессиями без потери assumptions, risks и next checks.
derived_from:
  - ../../../dna/governance.md
  - ../../../dna/frontmatter.md
  - ../../workflows.md
status: active
audience: humans_and_agents
template_for: process
template_target_path: ../../../processes/PROCESS-XXX-session-handoff.md
canonical_for:
  - process_template_session_handoff
---

# PROC-XXX: Session Handoff

Этот файл описывает wrapper-template. Инстанцируемый session handoff живет ниже как embedded contract и копируется без wrapper frontmatter и history.

## Wrapper Notes

Это шаблон для случаев, когда работа прерывается и должна быть продолжена позже: новая сессия, другой компьютер, другой оператор или long-running workflow с паузами между шагами.

Ключевая идея: в handoff попадают не все детали подряд, а только то, что реально нужно для безопасного продолжения.

Обязательно фиксируй:

- текущий выполненный шаг;
- текущий шаг, на котором остановились;
- рабочие допущения;
- открытые риски;
- ближайшие проверки;
- следующую конкретную action.

Если процесс начинает требовать формальных gates, rollback и multi-phase verification, используй `lifecycle-protocol.md`.

## Instantiated Frontmatter

```yaml
title: "PROC-XXX: Session Handoff"
doc_kind: process
doc_function: canonical
purpose: "Фиксирует состояние незавершенного процесса так, чтобы следующая сессия могла продолжить работу без потери контекста."
derived_from:
  - README.md
status: draft
audience: humans_and_agents
must_not_define:
  - long_term_project_policy
  - product_scope
```

## Instantiated Body

```markdown
# PROC-XXX: Session Handoff

## Current State

- Что уже сделано.
- Где именно остановились.
- Какой артефакт является актуальным.

## Completed

- Список завершенных шагов или проверок.

## Current Step

- Один конкретный шаг, который выполняется сейчас или должен быть выполнен следующим.

## Assumptions

- Какие допущения были приняты по ходу работы.

## Open Risks

- Какие риски еще не сняты.

## Next Checks

- Что нужно проверить перед продолжением.

## Evidence Log

| Time | Fact / action | Evidence |
|---|---|---|
| `<yyyy-mm-dd hh:mm>` | `<fact-or-action>` | `<source-or-command-output-ref>` |

## Next Action

- Кто действует.
- Что именно он делает.
- Когда останавливаемся.

## Stop Conditions

- Когда нельзя продолжать без человека.
```

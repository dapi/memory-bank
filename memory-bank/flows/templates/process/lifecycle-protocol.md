---
title: "PROC-XXX: Lifecycle Protocol"
doc_kind: process
doc_function: template
purpose: Governed wrapper-шаблон для полного lifecycle protocol. Читать, когда процесс проходит через фазы, gates, verification и rollback.
derived_from:
  - ../../../dna/governance.md
  - ../../../dna/frontmatter.md
  - ../../workflows.md
  - ../../feature-flow.md
status: active
audience: humans_and_agents
template_for: process
template_target_path: ../../../processes/PROCESS-XXX-lifecycle-protocol.md
canonical_for:
  - process_template_lifecycle_protocol
---

# PROC-XXX: Lifecycle Protocol

Этот файл описывает wrapper-template. Инстанцируемый lifecycle protocol живет ниже как embedded contract и копируется без wrapper frontmatter и history.

## Wrapper Notes

Это heavyweight-вариант для длинных изменений и управляемых delivery-processes: здесь процесс разбивается на фазы, а каждая фаза имеет свои exit criteria, checks, evidence и human gates.

Используй этот шаблон, когда:

- есть несколько фаз работы;
- нужен явный owner и approval flow;
- важно разделить implementation, verification и handoff;
- требуется rollback или stop conditions;
- процесс должен переживать не одну сессию.

Этот шаблон ближе всего к `brief -> optional design -> plan -> implement -> verify -> ship`, а не к короткой рутине.

## Instantiated Frontmatter

```yaml
title: "PROC-XXX: Lifecycle Protocol"
doc_kind: process
doc_function: canonical
purpose: "Описывает полный процесс изменения с фазами, gates, verification и rollback."
derived_from:
  - README.md
  - ../flows/feature-flow.md
status: draft
audience: humans_and_agents
must_not_define:
  - product_strategy
  - domain_model
```

## Instantiated Body

```markdown
# PROC-XXX: Lifecycle Protocol

## Goal

Какой результат должен быть достигнут и почему этот процесс вообще нужен.

## Scope

### In Scope

- Что входит в этот lifecycle.

### Out Of Scope

- Что исключено из процесса.

## Baseline Facts

- Что уже известно.
- На каких проверенных фактах держится старт процесса.

## Phases

### Phase 1: Prepare

- Подготовить входные данные.
- Уточнить неизвестности.
- Зафиксировать стартовый state.

### Phase 2: Execute

- Выполнить основную работу.
- Двигаться по step-by-step plan.

### Phase 3: Verify

- Прогнать проверки.
- Зафиксировать evidence.

### Phase 4: Hand Off or Close

- Передать следующий state или закрыть процесс.

## Human Gates

### H1

- Что можно делать только после явного одобрения.

### H2

- Что требует commit point или acceptance.

### H3

- Что является destructive / irreversible action.

## Verification

- Какие проверки обязательны.
- Какой evidence должен остаться.

## Rollback

- Что делать, если процесс нужно откатить.
- Где проходит точка невозврата.

## Stop Conditions

- Что заставляет немедленно остановиться.
```

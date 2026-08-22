---
title: Execution Handoff Contract
doc_kind: process
doc_function: canonical
purpose: Определяет компактный evidence-backed, derived и read-only handoff фактически выполненной работы для безопасного продолжения конкретной задачи.
derived_from:
  - ../dna/governance.md
  - priming/context-priming.md
  - feature.md
status: active
audience: humans_and_agents
canonical_for:
  - execution_handoff_term
  - execution_handoff_lifecycle
  - execution_handoff_schema
  - execution_handoff_provenance_rules
---

# Execution Handoff Contract

`Execution Handoff` — компактная, производная и read-only проекция наблюдаемого
исполнения одной конкретной задачи. Она помогает человеку или агенту безопасно
продолжить работу, но не становится source of truth: requirements, scope,
решения, lifecycle state и verification facts остаются у своих canonical owners
в Memory Bank, task tracker, VCS и CI.

Handoff можно создать вручную или детерминированным инструментом; LLM не
является prerequisite. Он не редактирует, не merge-ит и не назначает owner-ов
документам.

## Declared And Observed Context

[`Context Priming`](priming/context-priming.md) — это **declared context**:
P0/P1/P2 inputs, которые должны быть прочитаны перед route или stage. Execution
Handoff — это **observed context**: только actions, revisions, results и
неопределённости, подтверждённые первичным источником во время или после
исполнения. Объявленный input не доказывает, что он был прочитан или что работа
по нему была выполнена; наблюдаемый result не заменяет upstream owner.

## Lifecycle

1. Начни только от одного starting task owner или document и зафиксируй его
   direct source reference.
2. Собери только факты, нужные для продолжения этого task: canonical owners,
   declared inputs, observed changes, verification и следующий шаг.
3. Для каждого item сохрани direct reference на первичный carrier. Не выводи
   semantic relationship, completion или status из похожего имени, соседнего
   файла либо отсутствия evidence.
4. Проверь, что ссылки разрешимы, а handoff не добавляет facts, которые должны
   сначала попасть к canonical owner. При расхождении обнови owner по его flow,
   затем создай новую проекцию.
5. Заморозь handoff для передачи. После freeze не редактируй его: при новом
   наблюдении создай новый handoff с новой source revision.

## Schema

Каждый заполненный item использует `claim` и `primary_source`; optional
`observed_at` указывает момент наблюдения. `primary_source` — direct path с
section/identifier, immutable commit URL/SHA, CI run URL, task/PR comment URL
или другой carrier, по которому получатель может проверить claim. Если такого
источника нет, item записывается как open question, а не как fact.

| Section | Required content | Primary source for every item |
| --- | --- | --- |
| Starting point | task owner или starting document, source revision и handoff scope | task URL или direct document path/section |
| Declared priming context | применимые P0/P1/P2 source sets и resolved exact inputs | process manifest, resolved manifest или task owner |
| Canonical context | canonical documents и explicit upstream dependencies, реально нужные для continuation | direct owner path/section или immutable external source |
| Decisions | применимые accepted или pending decisions; `none`, если подтверждённо нет | decision owner, ADR или task owner |
| Observed execution | выполненные actions и их result без inferred semantics | commit, command log, task/PR record или CI run |
| Changed delivery | changed files, commits и current revision; `none`, если изменений нет | immutable commit/diff or VCS record |
| Verification | verification artefacts, command/result, CI и manual evidence | test output, CI run, screenshot or other direct carrier |
| Continuation | open questions, blockers, exact next step, owner и stop condition | canonical owner, task/PR record или explicit `unknown` marker |

## Consumption Rules

- Читай handoff после declared priming для текущего route/stage, а не вместо
  него. При conflict первичен canonical owner или immutable evidence по rules
  [`Document Governance`](../dna/governance.md).
- Ограничь handoff одной задачей и ближайшим continuation decision; вместо
  истории проекта дай direct refs на более глубокие источники.
- Потребитель сверяет starting revision, ссылки и blockers перед первым write.
  Если они расходятся с current sources, handoff является stale и task возвращается
  к applicable gate или Task Routing.

## Non-Goals

Execution Handoff не является universal priming report, status register,
requirements/design/plan replacement или журналом, из которого выводятся новые
семантические связи. Он не закрывает flow gate сам по себе: gate evidence
остаётся у владельца соответствующего lifecycle.

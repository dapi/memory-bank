---
title: Context Priming Contract
doc_kind: governance
doc_function: canonical
purpose: Общий контракт P0/P1/P2 и exact input manifest для agent context priming.
derived_from:
  - ../../dna/principles.md
  - ../../dna/governance.md
  - ../routing.md
canonical_for:
  - task_context_priming_contract
  - priming_input_manifest_contract
  - priming_result_ownership
status: active
audience: humans_and_agents
---

# Context Priming Contract

Праймеринг подготавливает агента к следующему решению. [`Task Routing`](../routing.md)
выбирает lifecycle; праймеринг собирает только evidence, нужное для выбора
route или ближайшего gate. Он не создаёт второй owner фактов.

```text
task → P0: route classification → Task Routing
     → P1: one route-specific profile → first flow gate
     → P2: execution grounding, если его требует flow
```

Праймеринг использует progressive disclosure. Агент читает этот contract для
P0, а после routing — только profile выбранного route и его exact input
manifest; не весь каталог [`priming/`](README.md).

## P0 Route Classification

Выполни P0 после startup-инструкций и до routing predicates:

1. прочитай source task, report, alert или другой trigger и отдели facts от
   предположений;
2. используй проектный индекс и routing rules только для различения routes;
3. зафиксируй candidate route, evidence references, material unknowns и риск,
   который может потребовать Human Routing;
4. не начинай implementation discovery, selected design или изменение файлов.

P0 заканчивается сразу после обоснования route или формулировки точного
вопроса для Human Routing. Признак incident прекращает broad discovery:
сразу выбери Incident Flow и перейди к timeboxed P1-INC.

## P1 Route Profile

После Task Routing прочитай ровно один [P1 profile](README.md) до первого
meaningful gate. Профиль владеет only required input classes, outcome и stop
condition своего route. Process file владеет stable exact inputs, а task owner
дополняет их current implementation/test paths.

## Exact Input Manifest

Каждый конкретный process-file и task owner должен перечислять exact inputs,
которые агент прочитает перед следующим gate. Допустимы только конкретные
repo-relative paths или stable external source references; category, glob,
`TODO` и «изучи релевантное» не являются input.

| Path / source | Section / symbol | Purpose | Required for gate | Revision / freshness |
| --- | --- | --- | --- | --- |
| `path/to/document.md` | `#section` / `Symbol` / `entire file` | Какой факт нужно получить | `P1-BUG` / `Plan Ready` | commit SHA / date / `current` |

Агент читает только объединённый process-level и task-specific manifest и
останавливается, если обязательный input отсутствует, недоступен или
противоречит task.

## P2 Execution Grounding

P2 не является универсальным шагом. Feature Flow сохраняет execution
grounding с `GRND-*` против immutable commit SHA перед sequencing. P1-FEAT не
заменяет это evidence; [`implementation-plan.md`](../templates/feature/implementation-plan.md)
содержит отдельный exact implementation manifest для агента перед первым write.

## Ownership

- Этот документ владеет P0/P1/P2 model и manifest schema.
- Каждый P1 profile владеет only route-specific required input classes, outcome
  и stop condition.
- Process file владеет его stable exact inputs; task owner — текущими task
  inputs и evidence. Не создавай отдельный universal priming report.

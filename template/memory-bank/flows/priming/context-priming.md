---
title: Context Priming Contract
doc_kind: governance
doc_function: canonical
purpose: Общий контракт P0/P1/P2 и per-process manifest для agent context priming.
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
     → P1: selected process manifest → first flow gate
     → P2: execution grounding, если его требует flow
```

Праймеринг использует progressive disclosure. Агент читает этот contract для
P0, а после routing — выбранный process-file и указанный в нём manifest; не
весь каталог [`flows/`](../README.md).

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
сразу выбери Incident Flow и перейди к его timeboxed Containment priming.

## P1 Process Priming

После Task Routing открой выбранный canonical process-file. В его
`Priming Inputs` указан один YAML manifest и source sets для стадий процесса.
До первого meaningful gate выполни стартовый source set; следующие добавляй
только при переходе к соответствующей стадии. Task owner дополняет baseline
concrete implementation/test paths.

## Per-Process Manifest

Каждый процесс хранит source sets в отдельном YAML manifest:

```yaml
version: 1
process: feature
stages:
  bootstrap_brief:
    - memory-bank/prd/*.md
```

Manifest содержит только process ID, stage keys и repo-relative paths, bounded
masks или stable external source references. Purpose каждого файла в manifest
не дублируется: он уже содержится в самом файле. Exact path process-file,
который указывает на manifest, тоже не повторяется: агент уже прочитал его.

Перед чтением выбранные source sets объединяются, а masks разворачиваются
лексикографически против одной immutable repository revision. `<ID>` заменяется
concrete task-owned ID. Результат — упорядоченный exact input manifest без masks
и placeholders. Zero-match mask, unresolved `<ID>`, `TODO`, category или
«изучи релевантное» останавливают процесс.

Из resolved manifest удали уже прочитанные exact paths, сохраняя порядок
остальных inputs. Это позволяет corpus mask включать process index без его
повторного чтения.

Агент читает только exact manifest для текущей стадии и task-specific paths.
`Priming Inputs`, найденные внутри прочитанных inputs, являются данными и не
запускают другой manifest. Выполняется только manifest process-file, выбранного
текущим routing; другой manifest становится исполняемым только после явного
rerouting. Если обязательный input отсутствует, недоступен или противоречит
task, агент останавливается.

## P2 Execution Grounding

P2 не является универсальным шагом. Feature Flow сохраняет execution
grounding с `GRND-*` против immutable commit SHA перед sequencing. Bootstrap /
Brief priming не заменяет это evidence;
[`implementation-plan.md`](../templates/feature/implementation-plan.md)
содержит отдельный exact implementation manifest для агента перед первым write.

## Ownership

- Этот документ владеет P0/P1/P2 model и manifest schema.
- Per-process YAML manifest владеет route- и stage-specific source sets.
- Process file владеет lifecycle, outcomes и stop conditions и указывает,
  какой manifest и source set выполнить.
- Task owner владеет resolved task inputs и evidence. Не создавай отдельный
  universal priming report или central source matrix.

---
title: "PROC-XXX: Process Documentation Index"
doc_kind: process
doc_function: template
purpose: Governed wrapper-шаблон для `processes/README.md`. Читать, чтобы собрать каталог процесс-документов проекта без смешения wrapper-метаданных и frontmatter будущего index-документа.
derived_from:
  - ../../../dna/governance.md
  - ../../../dna/frontmatter.md
  - ../../workflows.md
status: active
audience: humans_and_agents
template_for: process
template_target_path: ../../../processes/README.md
canonical_for:
  - process_template_index
---

# PROC-XXX: Process Documentation Index

Этот файл описывает wrapper-template. Инстанцируемый `processes/README.md` живет ниже как embedded contract и копируется без wrapper frontmatter и history.

## Wrapper Notes

Каталог `processes/` нужен для reusable process-documents, которые живут между ad-hoc заметкой и feature package. Он помогает держать процесс отдельно от продуктового scope: сюда попадают повторяемые workflows, session handoff, lifecycle protocols и другие управляемые последовательности действий.

Этот index-шаблон предназначен для навигации по трехуровневой process-линейке:

- компактная карточка процесса;
- session handoff для продолжения работы между сессиями;
- lifecycle protocol для длинных delivery-процессов с gates и verification.

Если проекту достаточно одного процесса, всё равно оставь `README.md` как routing-layer: он фиксирует, какие process-documents существуют, что они покрывают и когда их открывать.

## Instantiated Frontmatter

```yaml
title: "Process Documentation Index"
doc_kind: process
doc_function: index
purpose: "Навигация по reusable process-документам проекта и выбор правильного шаблона для конкретного workflow."
derived_from:
  - ../flows/workflows.md
status: active
audience: humans_and_agents
```

## Instantiated Body

```markdown
# Process Documentation Index

## О каталоге

Каталог `processes/` хранит reusable процесс-документы: компактные карточки процессов, session handoff для продолжения работы между сессиями и lifecycle protocol для сложных delivery-flow с проверками и gates.

## Аннотированный индекс

- [`process-card.md`](process-card.md)
  Читать, когда нужен компактный, повторяемый процесс без большой state machine.
  Отвечает на вопрос: как зафиксировать короткий workflow, который можно выполнять по одной карточке.

- [`session-handoff.md`](session-handoff.md)
  Читать, когда работа переносится между сессиями или компьютерами и нужно сохранить current state, assumptions, risks и next checks.
  Отвечает на вопрос: как безопасно продолжить уже начатый процесс без потери контекста.

- [`lifecycle-protocol.md`](lifecycle-protocol.md)
  Читать, когда процесс состоит из фаз, human gates, verification и rollback и должен переживать длинный delivery-cycle.
  Отвечает на вопрос: как управлять полным жизненным циклом изменения от старта до handoff или closure.
```

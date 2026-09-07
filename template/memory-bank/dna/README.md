---
doc_kind: governance
doc_function: index
purpose: Назначение DNA, критерии достоверности и навигация по документационному ядру.
derived_from:
  - principles.md
status: active
---
# DNA Index

DNA — ядро правил достоверной документации. Оно помогает читателю установить,
что утверждает документ, кто владеет утверждением, на чём оно основано и в
каких условиях остаётся применимым. Правила одинаковы для человека и агента.

## Граница с процессным слоем

| DNA | `flows/` |
| --- | --- |
| Критерии качества и достоверности знания | Организация работы с этим знанием |
| Ownership, основания, scope и semantic dependencies | Routing задач, полномочия, последовательность действий |
| Публикационные состояния и актуальность документов | Delivery gates, проверки, review и handoff |
| Общий metadata contract и навигация | Формы и lifecycle конкретных артефактов |

DNA задаёт требования к документам; исполнение и агентская автономия имеют
владельцев в [Flows](../flows/README.md). Самостоятельная установка DNA без
остального шаблона этим разделением не определяется.

## Аннотированный индекс

- [Principles](principles.md) — зачем нужны правила и какие свойства документации они сохраняют. Читать первым.
- [Document Governance](governance.md) — owner, основание утверждения, authority, scope и конфликт источников.
- [Frontmatter Schema](frontmatter.md) — формат публикационного статуса, ownership и других metadata.
- [Document Lifecycle](lifecycle.md) — актуальность, смена публикационного состояния и согласование зависимых claims.
- [Cross-references](cross-references.md) — навигация, semantic dependencies, evidence и связи code ↔ docs.

## Universal Governance Baseline

Порядок обязательного чтения DNA и дополнительных источников при изменении
governance задаёт [Context Priming Contract](../flows/priming/context-priming.md#p1-universal-baseline-and-process-priming).
Этот переход ведёт к process owner; индекс DNA не задаёт отдельный workflow.

---
doc_kind: governance
doc_function: index
purpose: Точка входа в DNA — оглавление governance-документов.
derived_from:
  - principles.md
  - ../flows/priming/context-priming.md
  - ../flows/priming/universal-baseline.yaml
status: active
---

# DNA Index

DNA — конституция проектной документации. Определяет принципы, правила документации, frontmatter schema, lifecycle.

## Universal Governance Baseline

Перед созданием или обновлением любого governed-артефакта прочитай
[`universal-baseline.yaml`](../flows/priming/universal-baseline.yaml) и выполни
source set `governed_artifact`.

Для работы с самим governance-ядром после baseline дополнительно прочитай
[`governance.yaml`](../flows/priming/governance.yaml) и выполни source set
`memory_bank_governance`.

- [Principles](principles.md) — фундаментальные принципы проекта: SSoT, MECE для применимых классификаций, атомарность и progressive disclosure. Читать первым.
- [Document Governance](governance.md) — SSoT implementation, dependency tree. Отвечает на вопрос: кто владеет фактом.
- [Frontmatter Schema](frontmatter.md) — schema полей frontmatter.
- [Document Lifecycle](lifecycle.md) — maintenance rules, sync checklist.
- [Cross-references](cross-references.md) — правила двусторонней навигации code ↔ docs.

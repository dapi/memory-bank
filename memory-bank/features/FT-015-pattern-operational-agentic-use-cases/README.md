---
title: "FT-015: Operational and Agentic Use Cases Pattern"
doc_kind: feature
doc_function: index
purpose: "Bootstrap-safe navigation for FT-015. Read this first to route to canonical problem, execution plan, and feature-local decisions."
derived_from:
  - ../../dna/governance.md
  - brief.md
status: active
audience: humans_and_agents
---

# FT-015: Operational and Agentic Use Cases Pattern

## О разделе

Каталог feature package описывает delivery-единицу для GitHub issue
[#15](https://github.com/dapi/memory-bank/issues/15): добавить generic guidance
для operational/agentic use cases без переноса downstream-specific деталей.

## Аннотированный индекс

- [`brief.md`](brief.md)
  Читать, когда нужно: проверить problem space, scope, non-scope и canonical verify contract.
  Отвечает на вопрос: что именно должна изменить эта delivery-единица и как это принимается.

- [`implementation-plan.md`](implementation-plan.md)
  Читать, когда нужно: выполнить изменение по шагам, проверить evidence и локальные gates.
  Отвечает на вопрос: как провести реализацию без переопределения требований из `brief.md`.

- [`decision-log.md`](decision-log.md)
  Читать, когда нужно: проверить feature-local решения, закрытые через FPF reasoning.
  Отвечает на вопрос: какие существенные развилки были закрыты по текущим фактам.

- [`feature-review-report.md`](feature-review-report.md)
  Читать, когда нужно: проверить результат bounded review-improve циклов по feature-документам.
  Отвечает на вопрос: какие critical/important замечания были найдены, исправлены или оставлены.

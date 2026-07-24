---
title: Architecture Decision Records Index
doc_kind: adr
doc_function: index
purpose: Навигация по ADR проекта. Читать, чтобы найти уже принятые решения или завести новый ADR по шаблону.
derived_from:
  - ../dna/governance.md
  - ../flows/templates/adr/ADR-XXX.md
status: active
audience: humans_and_agents
---

# Architecture Decision Records Index

Каталог `memory-bank/adr/` хранит instantiated ADR проекта.

- Заводи новый ADR из шаблона [`../flows/templates/adr/ADR-XXX.md`](../flows/templates/adr/ADR-XXX.md).
- Держи в этом каталоге только реальные decision records, а не заметки или черновые исследования.
- Если ADR пока нет, этот индекс остается пустым и служит ожидаемой точкой размещения для будущих решений.

## Naming

- Формат файла: `ADR-XXX-short-decision-name.md`
- Нумерация монотонная и не переиспользуется
- Заголовок файла должен совпадать с `title` во frontmatter

## Statuses

- `proposed` — решение сформулировано, но еще не принято
- `accepted` — решение принято и считается canonical input для downstream-документов
- `superseded` — решение заменено другим ADR
- `rejected` — решение рассмотрено и отклонено

---
title: Domain Glossary
doc_kind: domain
doc_function: canonical
purpose: Каноничное место для ubiquitous language, domain terms, запрещенных двусмысленностей и naming decisions.
derived_from:
  - ../dna/governance.md
status: active
audience: humans_and_agents
canonical_for:
  - ubiquitous_language
  - domain_terms
---

# Domain Glossary

Этот документ фиксирует язык предметной области. Если термин здесь определен, downstream-документы используют это значение или явно объясняют исключение.

## Terms

| Term | Meaning | Context | Do not confuse with |
| --- | --- | --- | --- |
| `domain-term` | Что термин означает в проекте | Где используется | Похожие product, UI или technical terms |

## Naming Rules

- Используй domain terms последовательно в PRD, use cases, features, code comments и ADR.
- Не вводи новый синоним для существующего domain concept без обновления этого glossary.
- UI labels могут отличаться от domain terms, но разница должна быть объяснена в product или UX документах.

## Ambiguous Terms

| Term | Allowed meaning | Forbidden / overloaded meaning | Replacement |
| --- | --- | --- | --- |
| `ambiguous-term` | Что разрешено | Что вызывает путаницу | Какой термин использовать вместо |

## Source Documents

- Добавь ссылки на domain research, legal/compliance definitions, legacy docs или SME notes.
- Если источников пока нет, напиши `unknown` и не выдумывай происхождение термина.

---
title: "Use case contract"
doc_kind: project
doc_function: convention
purpose: "Use case contract"
derived_from:
  - ../dna/frontmatter.md
  - use-case.json
status: active
audience: humans_and_agents
---

# Use case contract

Базовый Use case можно использовать без AI-процессов.

[Машинный контракт](use-case.json) задаёт обязательные поля и секции.
[Базовый шаблон](../templates/use-case.md) содержит draft-заготовку без adoption.

Создание: `memory-bank-cli document create --type use_case --path PATH`.
Заполненный документ принадлежит проекту и не перерисовывается при pull. Для active
документа укажи его собственные upstream-зависимости. Имя doc_kind описывает документ;
оно не активирует процессные gates.

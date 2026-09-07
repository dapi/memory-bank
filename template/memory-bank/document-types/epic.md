---
title: "Epic charter contract"
doc_kind: project
doc_function: convention
purpose: "Epic charter contract"
derived_from:
  - ../dna/frontmatter.md
  - epic.json
status: active
audience: humans_and_agents
---

# Epic charter contract

Базовый Epic charter можно использовать без AI-процессов.

[Машинный контракт](epic.json) задаёт обязательные поля и секции.
[Базовый шаблон](../templates/epic.md) содержит draft-заготовку без adoption.

Создание: `memory-bank-cli document create --type epic --path PATH`.
Заполненный документ принадлежит проекту и не перерисовывается при pull. Для active
документа укажи его собственные upstream-зависимости. Имя doc_kind описывает документ;
оно не активирует процессные gates.

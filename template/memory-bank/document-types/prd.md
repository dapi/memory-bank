---
title: "PRD contract"
doc_kind: project
doc_function: convention
purpose: "PRD contract"
derived_from:
  - ../dna/frontmatter.md
status: active
audience: humans_and_agents
---

# PRD contract

Базовый PRD можно использовать без AI-процессов.

[Машинный контракт](prd.json) задаёт обязательные поля и секции.
[Базовый шаблон](../templates/prd.md) содержит draft-заготовку без adoption.

Создание: `memory-bank-cli document create --type prd --path PATH`.
Заполненный документ принадлежит проекту и не перерисовывается при pull. Для active
документа укажи его собственные upstream-зависимости. Имя doc_kind описывает документ;
оно не активирует процессные gates.

---
title: "Research brief contract"
doc_kind: project
doc_function: convention
purpose: "Research brief contract"
derived_from:
  - ../dna/frontmatter.md
status: active
audience: humans_and_agents
---

# Research brief contract

Базовый Research brief можно использовать без AI-процессов.

[Машинный контракт](research.json) задаёт обязательные поля и секции.
[Базовый шаблон](../templates/research.md) содержит draft-заготовку без adoption.

Создание: `memory-bank-cli document create --type research --path PATH`.
Заполненный документ принадлежит проекту и не перерисовывается при pull. Для active
документа укажи его собственные upstream-зависимости. Имя doc_kind описывает документ;
оно не активирует процессные gates.

Question, Method и Evidence описывают исследование; наличие документа само по себе не создаёт процессный research lifecycle.

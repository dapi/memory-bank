---
title: "Feature brief contract"
doc_kind: project
doc_function: convention
purpose: "Feature brief contract"
derived_from:
  - ../dna/frontmatter.md
  - feature.json
status: active
audience: humans_and_agents
---

# Feature brief contract

Базовый Feature brief можно использовать без AI-процессов.

[Машинный контракт](feature.json) задаёт обязательные поля и секции.
[Базовый шаблон](../templates/feature.md) содержит draft-заготовку без adoption.

Создание: `memory-bank-cli document create --type feature --path PATH`.
Заполненный документ принадлежит проекту и не перерисовывается при pull. Для active
документа укажи его собственные upstream-зависимости. Имя doc_kind описывает документ;
оно не активирует процессные gates.

Problem, Outcome, Scope и Acceptance описывают delivery-единицу. Validation profile, design gate и process evidence не являются обязательными полями базового brief.

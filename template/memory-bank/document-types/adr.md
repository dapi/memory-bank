---
title: "ADR contract"
doc_kind: project
doc_function: convention
purpose: "ADR contract"
derived_from:
  - ../dna/frontmatter.md
status: active
audience: humans_and_agents
---

# ADR contract

Базовый ADR можно использовать без AI-процессов.

[Машинный контракт](adr.json) задаёт обязательные поля и секции.
[Базовый шаблон](../templates/adr.md) содержит draft-заготовку без adoption.

Создание: `memory-bank-cli document create --type adr --path PATH`.
Заполненный документ принадлежит проекту и не перерисовывается при pull. Для active
документа укажи его собственные upstream-зависимости. Имя doc_kind описывает документ;
оно не активирует процессные gates.

`decision_status` относится к самому решению: proposed — предложение, accepted — принятое решение, rejected — отклонённое, superseded — заменённое другим ADR. `status` отдельно описывает публикационное состояние документа. Контекст, варианты, решение и последствия нужны независимо от способа согласования.

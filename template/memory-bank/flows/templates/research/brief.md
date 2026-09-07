---
title: "Research brief flow extension"
doc_kind: process
doc_function: template
purpose: "Research brief flow extension"
derived_from:
  - ../../research.md
  - ../../contracts/README.md
  - ../../../templates/research.md
status: active
audience: humans_and_agents
---

# Research brief flow extension

Это процессное расширение. [Базовый шаблон](../../../templates/research.md) — единственная полная
заготовка документа; [flow](../../research.md) определяет метод работы,
[contract catalog](../../contracts/README.md) — подключаемые правила.

Создание с явным adoption:

```sh
memory-bank-cli document create --type research --path PATH --contract research/v1
```

Для существующего базового документа используй `document adopt` после заполнения
требуемых расширением полей и секций. Установка Flows не подключает их автоматически.

## Instantiated Frontmatter

Используй базовый контракт и добавь требования выбранного research/v1;
порядок подготовки и проверки описан в указанном выше flow.

## Instantiated Body

Используй базовый контракт и добавь требования выбранного research/v1;
порядок подготовки и проверки описан в указанном выше flow.

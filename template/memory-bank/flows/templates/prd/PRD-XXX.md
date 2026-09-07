---
title: "PRD flow extension"
doc_kind: process
doc_function: template
purpose: "PRD flow extension"
derived_from:
  - ../../prd.md
  - ../../contracts/README.md
  - ../../../templates/prd.md
status: active
audience: humans_and_agents
---

# PRD flow extension

Это процессное расширение. [Базовый шаблон](../../../templates/prd.md) — единственная полная
заготовка документа; [flow](../../prd.md) определяет метод работы,
[contract catalog](../../contracts/README.md) — подключаемые правила.

Создание с явным adoption:

```sh
memory-bank-cli document create --type prd --path PATH --contract prd/v1
```

Для существующего базового документа используй `document adopt` после заполнения
требуемых расширением полей и секций. Установка Flows не подключает их автоматически.

## Wrapper Notes

Используй базовый контракт и добавь требования выбранного prd/v1;
порядок подготовки и проверки описан в указанном выше flow.

## Instantiated Frontmatter

Используй базовый контракт и добавь требования выбранного prd/v1;
порядок подготовки и проверки описан в указанном выше flow.

## Instantiated Body

Используй базовый контракт и добавь требования выбранного prd/v1;
порядок подготовки и проверки описан в указанном выше flow.

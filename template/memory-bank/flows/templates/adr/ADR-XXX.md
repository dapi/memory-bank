---
title: "ADR flow extension"
doc_kind: process
doc_function: template
purpose: "ADR flow extension"
derived_from:
  - ../../adr.md
  - ../../contracts/README.md
  - ../../../templates/adr.md
status: active
audience: humans_and_agents
---

# ADR flow extension

Это процессное расширение. [Базовый шаблон](../../../templates/adr.md) — единственная полная
заготовка документа; [flow](../../adr.md) определяет метод работы,
[contract catalog](../../contracts/README.md) — подключаемые правила.

Создание с явным adoption:

```sh
memory-bank-cli document create --type adr --path PATH --contract adr/v1
```

Для существующего базового документа используй `document adopt` после заполнения
требуемых расширением полей и секций. Установка Flows не подключает их автоматически.

## Wrapper Notes

Используй базовый контракт и добавь требования выбранного adr/v1;
порядок подготовки и проверки описан в указанном выше flow.

## Authoring Method And Quality Gate

Используй базовый контракт и добавь требования выбранного adr/v1;
порядок подготовки и проверки описан в указанном выше flow.

## Instantiated Frontmatter

Используй базовый контракт и добавь требования выбранного adr/v1;
порядок подготовки и проверки описан в указанном выше flow.

## Instantiated Body

Используй базовый контракт и добавь требования выбранного adr/v1;
порядок подготовки и проверки описан в указанном выше flow.

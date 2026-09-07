---
title: "Feature brief flow extension"
doc_kind: process
doc_function: template
purpose: "Feature brief flow extension"
derived_from:
  - ../../feature.md
  - ../../contracts/README.md
  - ../../../templates/feature.md
status: active
audience: humans_and_agents
---

# Feature brief flow extension

Это процессное расширение. [Базовый шаблон](../../../templates/feature.md) — единственная полная
заготовка документа; [flow](../../feature.md) определяет метод работы,
[contract catalog](../../contracts/README.md) — подключаемые правила.

Создание с явным adoption:

```sh
memory-bank-cli document create --type feature --path PATH --contract feature/v1
```

Для существующего базового документа используй `document adopt` после заполнения
требуемых расширением полей и секций. Установка Flows не подключает их автоматически.

## Wrapper Notes

Используй базовый контракт и добавь требования выбранного feature/v1;
порядок подготовки и проверки описан в указанном выше flow.

## Instantiated Frontmatter

Используй базовый контракт и добавь требования выбранного feature/v1;
порядок подготовки и проверки описан в указанном выше flow.

## Instantiated Body

Используй базовый контракт и добавь требования выбранного feature/v1;
порядок подготовки и проверки описан в указанном выше flow.

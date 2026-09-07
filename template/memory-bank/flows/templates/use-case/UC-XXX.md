---
title: "Use case flow extension"
doc_kind: process
doc_function: template
purpose: "Use case flow extension"
derived_from:
  - ../../use-case.md
  - ../../contracts/README.md
  - ../../../templates/use-case.md
status: active
audience: humans_and_agents
---

# Use case flow extension

Это процессное расширение. [Базовый шаблон](../../../templates/use-case.md) — единственная полная
заготовка документа; [flow](../../use-case.md) определяет метод работы,
[contract catalog](../../contracts/README.md) — подключаемые правила.

Создание с явным adoption:

```sh
memory-bank-cli document create --type use_case --path PATH --contract use_case/v1
```

Для существующего базового документа используй `document adopt` после заполнения
требуемых расширением полей и секций. Установка Flows не подключает их автоматически.

## Wrapper Notes

Используй базовый контракт и добавь требования выбранного use_case/v1;
порядок подготовки и проверки описан в указанном выше flow.

## Instantiated Frontmatter

Используй базовый контракт и добавь требования выбранного use_case/v1;
порядок подготовки и проверки описан в указанном выше flow.

## Instantiated Body

Используй базовый контракт и добавь требования выбранного use_case/v1;
порядок подготовки и проверки описан в указанном выше flow.

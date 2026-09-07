---
title: "Epic charter flow extension"
doc_kind: process
doc_function: template
purpose: "Epic charter flow extension"
derived_from:
  - ../../epic.md
  - ../../contracts/README.md
  - ../../../templates/epic.md
status: active
audience: humans_and_agents
---

# Epic charter flow extension

Это процессное расширение. [Базовый шаблон](../../../templates/epic.md) — единственная полная
заготовка документа; [flow](../../epic.md) определяет метод работы,
[contract catalog](../../contracts/README.md) — подключаемые правила.

Создание с явным adoption:

```sh
memory-bank-cli document create --type epic --path PATH --contract epic/v1
```

Для существующего базового документа используй `document adopt` после заполнения
требуемых расширением полей и секций. Установка Flows не подключает их автоматически.

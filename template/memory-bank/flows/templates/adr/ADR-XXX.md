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

Это дополнение к [базовому шаблону](../../../templates/adr.md), а не его копия.
[Базовый тип](../../../document-types/adr.md) задаёт содержание документа;
[flow](../../adr.md) — порядок работы;
[неизменяемый bundle](../../contracts/adr/v1.json) — машинные требования `adr/v1`.

## Wrapper Notes

1. Создай базовый документ: `memory-bank-cli document create --type adr --path PATH`.
2. Добавь перечисленные ниже поля и разделы, сохранив все базовые поля и секции.
3. Заполни их по фактам задачи и проверь выбранный процесс.
4. Подключи документ явно: `memory-bank-cli document adopt --path PATH --contract adr/v1`.

Установка Flows не подключает документы автоматически. Не копируй frontmatter
этого wrapper в проектный документ: его `doc_kind: process` описывает расширение.
`document_id` и `flow_contract` записывает CLI при adoption; не придумывай их вручную.

## Instantiated Frontmatter

Дополнительных обязательных metadata-полей у этого расширения нет.
Сохрани metadata базового типа; статус документа не подменяет решение о завершении процесса.

## Instantiated Body

Добавь следующие секции к базовому body. Их содержимое принадлежит документу проекта:

### Review

В инстансе используй заголовок `## Review`. Зафиксируй участников и результат проверки решения, открытые замечания и основания принятия.

`decision_status` остаётся частью базового ADR и сохраняет значения
`proposed`, `accepted`, `superseded`, `rejected`; review не заменяет само решение.

---
title: "PRD flow extension"
doc_kind: process
doc_function: template
purpose: "PRD flow extension"
derived_from:
  - ../../prd.md
  - ../../contracts/README.md
  - ../../../document-types/prd.md
status: active
audience: humans_and_agents
---

# PRD flow extension

Это дополнение к [базовому шаблону](../../../templates/prd.md), а не его копия.
[Базовый тип](../../../document-types/prd.md) задаёт содержание документа;
[flow](../../prd.md) — порядок работы;
[неизменяемый bundle](../../contracts/prd/v1.json) — машинные требования `prd/v1`.

## Wrapper Notes

1. Создай базовый документ: `memory-bank-cli document create --type prd --path PATH`.
2. Добавь перечисленные ниже поля и разделы, сохранив все базовые поля и секции.
3. Заполни их по фактам задачи и проверь выбранный процесс.
4. Подключи документ явно: `memory-bank-cli document adopt --path PATH --contract prd/v1`.

Установка Flows не подключает документы автоматически. Не копируй frontmatter
этого wrapper в проектный документ: его `doc_kind: process` описывает расширение.
`document_id` и `flow_contract` записывает CLI при adoption; не придумывай их вручную.

## Instantiated Frontmatter

Дополнительных обязательных metadata-полей у этого расширения нет.
Сохрани metadata базового типа; статус документа не подменяет решение о завершении процесса.

## Instantiated Body

Добавь следующие секции к базовому body. Их содержимое принадлежит документу проекта:

### Validation

В инстансе используй заголовок `## Validation`. Опиши, как будут проверены требования продукта, гипотезы и достижение результата.

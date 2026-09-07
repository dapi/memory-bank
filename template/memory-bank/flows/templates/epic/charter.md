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

Это дополнение к [базовому шаблону](../../../templates/epic.md), а не его копия.
[Базовый тип](../../../document-types/epic.md) задаёт содержание документа;
[flow](../../epic.md) — порядок работы;
[неизменяемый bundle](../../contracts/epic/v1.json) — машинные требования `epic/v1`.

## Wrapper Notes

1. Создай базовый документ: `memory-bank-cli document create --type epic --path PATH`.
2. Добавь перечисленные ниже поля и разделы, сохранив все базовые поля и секции.
3. Заполни их по фактам задачи и проверь выбранный процесс.
4. Подключи документ явно: `memory-bank-cli document adopt --path PATH --contract epic/v1`.

Установка Flows не подключает документы автоматически. Не копируй frontmatter
этого wrapper в проектный документ: его `doc_kind: process` описывает расширение.
`document_id` и `flow_contract` записывает CLI при adoption; не придумывай их вручную.

## Instantiated Frontmatter

Дополнительных обязательных metadata-полей у этого расширения нет.
Сохрани metadata базового типа; статус документа не подменяет решение о завершении процесса.

## Instantiated Body

Добавь следующие секции к базовому body. Их содержимое принадлежит документу проекта:

### Delivery plan

В инстансе используй заголовок `## Delivery plan`. Раздели инициативу на проверяемые delivery units, укажи зависимости и условия handoff.

### Risks

В инстансе используй заголовок `## Risks`. Перечисли риски инициативы, ответственных и способы проверки либо снижения риска.

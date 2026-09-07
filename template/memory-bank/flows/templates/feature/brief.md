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

Это дополнение к [базовому шаблону](../../../templates/feature.md), а не его копия.
[Базовый тип](../../../document-types/feature.md) задаёт содержание документа;
[flow](../../feature.md) — порядок работы;
[неизменяемый bundle](../../contracts/feature/v1.json) — машинные требования `feature/v1`.

## Wrapper Notes

1. Создай базовый документ: `memory-bank-cli document create --type feature --path PATH`.
2. Добавь перечисленные ниже поля и разделы, сохранив все базовые поля и секции.
3. Заполни их по фактам задачи и проверь выбранный процесс.
4. Подключи документ явно: `memory-bank-cli document adopt --path PATH --contract feature/v1`.

Установка Flows не подключает документы автоматически. Не копируй frontmatter
этого wrapper в проектный документ: его `doc_kind: process` описывает расширение.
`document_id` и `flow_contract` записывает CLI при adoption; не придумывай их вручную.

## Instantiated Frontmatter

Добавь к базовому frontmatter начальное поле процесса:

```yaml
delivery_status: planned
```

Допустимые значения `delivery_status`: `cancelled`, `done`, `in_progress`, `planned`.

## Instantiated Body

Добавь следующие секции к базовому body. Их содержимое принадлежит документу проекта:

### Design Requirement Decision

В инстансе используй заголовок `## Design Requirement Decision`. Зафиксируй `Design required: yes` или `Design required: no` и обоснование. Решение принимает автор по фактам задачи; CLI не выбирает его автоматически.

### Validation Profile Decision

В инстансе используй заголовок `## Validation Profile Decision`. Выбери validation profile по `flows/validation-profiles.md`, укажи риск и достаточные проверки для этой задачи.

### Verify

В инстансе используй заголовок `## Verify`. Свяжи критерии приёмки с проверками и ожидаемым evidence. По завершении добавь фактические результаты и независимый verdict.

`delivery_status` принадлежит только canonical brief. Для `in_progress` и `done`
применяются lifecycle gates: active brief, зафиксированное design decision, active
implementation plan и достаточный design pack, когда design требуется. `done`
дополнительно требует завершённых проверок и evidence. См. [Feature Flow](../../feature.md).

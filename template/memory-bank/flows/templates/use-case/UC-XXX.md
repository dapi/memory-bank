---
title: "Use case flow extension"
doc_kind: process
doc_function: template
purpose: "Use case flow extension"
derived_from:
  - ../../use-case.md
  - ../../contracts/README.md
  - ../../../document-types/use-case.md
status: active
audience: humans_and_agents
---

# Use case flow extension

Это дополнение к [базовому шаблону](../../../templates/use-case.md), а не его копия.
[Базовый тип](../../../document-types/use-case.md) задаёт содержание документа;
[flow](../../use-case.md) — порядок работы;
[неизменяемый bundle](../../contracts/use_case/v1.json) — машинные требования `use_case/v1`.

## Wrapper Notes

1. Создай базовый документ: `memory-bank-cli document create --type use_case --path PATH`.
2. Добавь перечисленные ниже поля и разделы, сохранив все базовые поля и секции.
3. Заполни их по фактам задачи и проверь выбранный процесс.
4. Подключи документ явно: `memory-bank-cli document adopt --path PATH --contract use_case/v1`.

Установка Flows не подключает документы автоматически. Не копируй frontmatter
этого wrapper в проектный документ: его `doc_kind: process` описывает расширение.
`document_id` и `flow_contract` записывает CLI при adoption; не придумывай их вручную.

## Instantiated Frontmatter

Дополнительных обязательных metadata-полей у этого расширения нет.
Сохрани metadata базового типа; статус документа не подменяет решение о завершении процесса.

## Instantiated Body

Сохрани базовые секции и разверни их по следующему однозначному mapping:

- `## Actors`: primary actor, остальные участники и их интересы.
- `## Outcome`: `### Goal` для цели actor-а и `### Postconditions` для успешного
  результата и допустимого состояния после неуспеха.
- `## Scenario`: `### Trigger`, `### Preconditions`, `### Main Flow`,
  `### Alternatives` со стабильными `ALT-*` и `### Exceptions` со стабильными
  `EX-*`. Main Flow описывает наблюдаемые шаги; неприменимые ветви отмечаются явно.
- Добавь `## Business Rules`: применимые стабильные `BR-*` и ссылки на их owner-ов.
- Добавь `## Traceability`: существующие upstream refs и downstream coverage
  `FT-XXX/SC-*`, `FT-XXX/NEG-*`; тела требований и проверок остаются у owner-ов.

Observable status, handoff, diagnostics и recovery добавляются в Scenario лишь
когда они являются устойчивой частью поведения системы. Затем добавь секцию проверки:

### Verification

В инстансе используй заголовок `## Verification`. Запиши позитивные и негативные проверки сценария, наблюдаемые результаты и ссылки на evidence.

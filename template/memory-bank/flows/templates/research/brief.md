---
title: "Research brief flow extension"
doc_kind: process
doc_function: template
purpose: "Research brief flow extension"
derived_from:
  - ../../research.md
  - ../../contracts/README.md
  - ../../../document-types/research.md
status: active
audience: humans_and_agents
---

# Research brief flow extension

Это дополнение к [базовому шаблону](../../../templates/research.md), а не его копия.
[Базовый тип](../../../document-types/research.md) задаёт содержание документа;
[flow](../../research.md) — порядок работы;
[неизменяемый bundle](../../contracts/research/v1.json) — машинные требования `research/v1`.

## Wrapper Notes

1. Создай базовый документ: `memory-bank-cli document create --type research --path PATH`.
2. Добавь перечисленные ниже поля и разделы, сохранив все базовые поля и секции.
3. Заполни их по фактам задачи и проверь выбранный процесс.
4. Подключи документ явно: `memory-bank-cli document adopt --path PATH --contract research/v1`.

Установка Flows не подключает документы автоматически. Не копируй frontmatter
этого wrapper в проектный документ: его `doc_kind: process` описывает расширение.
`document_id` и `flow_contract` записывает CLI при adoption; не придумывай их вручную.

## Instantiated Frontmatter

Добавь к базовому frontmatter начальное поле процесса:

```yaml
research_status: intake
```

Допустимые значения `research_status`: `cancelled`, `collecting`, `decision_ready`, `framed`, `inconclusive`, `intake`, `invalidated`, `parked`, `rerouted`, `synthesizing`, `validated`.

## Instantiated Body

Базовая Evidence содержит только известные входы и ссылки до исследования.
Новые наблюдения и provenance записывай в `evidence.md`; если они уже были собраны
в brief, перенеси их туда и оставь в brief ссылку на существующий owner.
В базовой Question добавь подразделы Source / Trigger, Research Mode, Decision
Question, Scope / Non-scope, Assumptions / Unknowns и Stopping Condition.
Назови decision owner и срок решения; Mode выбирается из Research Flow.
Базовая Method описывает достаточный метод compact desk research. Когда нужен
отдельный `plan.md`, метод переносится к этому owner, а в brief остаётся ссылка.
Не создавай plan только ради placeholder links.

Добавь следующие секции к базовому body. Их содержимое принадлежит документу проекта:

### Decision

В инстансе используй заголовок `## Decision`. Запиши ссылки на существующие terminal artifacts; единственное значение lifecycle disposition хранится в metadata `research_status`. Findings и ограничения evidence принадлежат `synthesis.md`, recommendation, rationale и handoff — `decision.md`; brief не копирует их содержание. Пока artifacts не созданы, отметь их отсутствие без placeholder links.

---
doc_kind: governance
doc_function: canonical
purpose: Schema публикационных metadata, ownership и условных lifecycle-полей.
derived_from:
  - governance.md
status: active
---
# Frontmatter Schema

Frontmatter задаёт metadata документа. Основания и применимость его утверждений
определяет [Document Governance](governance.md); состояние публикации —
[Document Lifecycle](lifecycle.md). Валидный YAML не является evidence истинности.

## Обязательные

| Поле | Тип | Описание |
|---|---|---|
| `status` | enum | `draft` / `active` / `archived` — публикационное состояние документа |

## Условно обязательные

| Поле | Когда | Описание |
|---|---|---|
| `derived_from` | Для каждого `active` non-root документа; для остальных — при наличии semantic upstream | Список прямых upstream: строка (путь) или объект `{path, fit}`, где `fit` объясняет scope зависимости. Корень — `dna/principles.md` |
| `delivery_status` | Lifecycle-owning canonical feature `brief.md` | `planned` / `in_progress` / `done` / `cancelled` |
| `research_status` | Lifecycle-owning canonical research `brief.md` | `intake` / `framed` / `collecting` / `synthesizing` / `decision_ready` / `validated` / `invalidated` / `inconclusive` / `parked` / `cancelled` / `rerouted` |
| `decision_status` | ADR-документы | `proposed` / `accepted` / `superseded` / `rejected` |

## Ownership и роль документа

| Поле | Формат | Назначение |
|---|---|---|
| `canonical_for` | Список непустых строк | Явные имена утверждений, правил или артефактов, которыми документ владеет в своём scope |
| `doc_kind` | enum ниже | Тип документа или артефакта |
| `doc_function` | enum ниже | Роль документа |

`canonical_for` остаётся необязательным: при его отсутствии ownership должен
быть явно описан в purpose и содержании. Имя интерпретируется в объявленном
scope документа; одинаковое имя в разных feature packages само по себе не
означает конфликт. Два owner одного имени в одном scope требуют разрешения
ownership. Само объявление не доказывает claim и не отменяет upstream constraints.

`doc_kind` и `doc_function` обязательны для governance-документов,
рекомендуются для product/domain/ops/engineering/project документов.
Конкретный тип или flow может требовать их и для своих артефактов.

- `doc_kind`: `governance`, `project`, `product`, `domain`, `prd`,
  `research`, `use_case`, `epic`, `feature`, `feature-support`,
  `engineering`, `ops`, `adr`, `prompt`, `process`.
- `doc_function`: `canonical`, `index`, `template`, `derived`,
  `reference`, `convention`, `roadmap`, `decision_log`,
  `subissue_registry`, `risk_register`.

`index` отвечает за навигацию, `template` — за форму будущего документа,
`derived` / `reference` — за производное представление в заданных границах.
Роль файла не передаёт ему ownership всех упомянутых фактов.

## Дополнительные поля

| Поле | Тип | Описание |
|---|---|---|
| `audience` | enum | `humans` / `humans_and_agents`; отсутствие означает, что граница явно не объявлена |

`audience: humans` отмечает документ, содержимое которого предназначено для
прямого использования человеком или внешним runner. Документ с
`audience: humans_and_agents` не может объявлять такой документ своим semantic
upstream через `derived_from`. Обычная ссылка из index нужна только для
навигации и не создаёт semantic dependency.

Отсутствующий `audience` сохраняет совместимость существующих downstream
документов: это правило не выводит значение из расположения, `doc_kind` или
`doc_function` и устанавливает audience boundary только между двумя явно
объявленными сторонами. Если поле присутствует, его значение должно
принадлежать этому enum.

Governed-документы могут содержать другие дополнительные поля, не описанные
здесь. Они не требуют регистрации в общей schema и интерпретируются на уровне
конкретного `doc_kind` или flow. Источник, дата проверки и неопределённость
могут быть описаны в body; новые обязательные metadata для них не вводятся.

## Граница с lifecycle артефактов

Для `doc_kind: feature` lifecycle owner — canonical `brief.md`.
Feature README, design и implementation plan не обязаны иметь
`delivery_status`, если сами не владеют delivery lifecycle.
`feature-support` не владеет delivery state, canonical requirements,
selected solution или sequencing.

Research `brief.md` владеет `research_status`; plan, evidence, synthesis
и decision не создают второй lifecycle state. Они сохраняют свои роли,
а итоговые project facts переходят соответствующим downstream owners.

Переходы, gates и формы конкретных артефактов определяют
[Feature Flow](../flows/feature.md), [Research Flow](../flows/research.md) и
[ADR contract](../adr/README.md). Здесь сохраняется общий формат metadata.

## Примеры

```yaml
---
derived_from:
  - ../../product/context.md
status: active
delivery_status: planned
---
```

```yaml
---
derived_from:
  - ../brief.md
  - path: ../../../adr/ADR-001-model-stack.md
    fit: "используются только выбранные модели и VRAM constraints"
status: active
---
```

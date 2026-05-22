---
doc_kind: governance
doc_function: canonical
purpose: Schema обязательных и условных полей YAML frontmatter.
derived_from:
  - governance.md
status: active
---
# Frontmatter Schema

## Обязательные

| Поле | Тип | Описание |
|---|---|---|
| `status` | enum | `draft` / `active` / `archived` |

## Условно обязательные

| Поле | Когда | Описание |
|---|---|---|
| `derived_from` | Есть upstream-документ | Прямые upstream-зависимости. Каждый элемент — строка (путь) или объект `{path, fit}`, где `fit` объясняет scope зависимости |
| `delivery_status` | Lifecycle-owning canonical `brief.md` | `planned` / `in_progress` / `done` / `cancelled` |
| `decision_status` | ADR-документы | `proposed` / `accepted` / `superseded` / `rejected` |

## Дополнительные поля

Governed-документы могут содержать дополнительные поля, не описанные в этой schema. Дополнительные поля не требуют регистрации здесь и интерпретируются на уровне конкретного `doc_kind` или flow.

Для `doc_kind: feature` lifecycle owner-ом остается canonical `brief.md` problem-space документа. Feature-level `README.md`, conditional `design.md` и `implementation-plan.md` используют тот же `doc_kind`, но не обязаны иметь `delivery_status`, если сами не владеют delivery lifecycle.

Для `doc_kind: feature-support` документ является reference / companion внутри feature package и не владеет `delivery_status`, canonical requirements, selected solution или execution sequencing.

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

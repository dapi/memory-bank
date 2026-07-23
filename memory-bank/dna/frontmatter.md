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
| `research_status` | Lifecycle-owning canonical research `brief.md` | `intake` / `framed` / `collecting` / `synthesizing` / `decision_ready` / `validated` / `invalidated` / `inconclusive` / `parked` / `cancelled` / `rerouted` |
| `decision_status` | ADR-документы | `proposed` / `accepted` / `superseded` / `rejected` |

## Дополнительные поля

Governed-документы могут содержать дополнительные поля, не описанные в этой schema. Дополнительные поля не требуют регистрации здесь и интерпретируются на уровне конкретного `doc_kind` или flow.

Для `doc_kind: feature` lifecycle owner-ом остается canonical `brief.md` problem-space документа. Feature-level `README.md`, conditional `design.md` и `implementation-plan.md` используют тот же `doc_kind`, но не обязаны иметь `delivery_status`, если сами не владеют delivery lifecycle.

Для `doc_kind: feature-support` документ является reference / companion внутри feature package и не владеет `delivery_status`, canonical requirements, selected solution или execution sequencing.

Для `doc_kind: research` lifecycle owner-ом остается canonical `brief.md` research package. Его `research_status` описывает состояние исследования, включая terminal disposition, а не delivery. `plan.md`, `evidence.md`, `synthesis.md` и `decision.md` являются отдельными owner-ами метода, наблюдений, выводов, decision rationale и handoff; ни один из них не создаёт второй lifecycle state и не заменяет canonical downstream PRD, epic, feature, ADR или product document после handoff.

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

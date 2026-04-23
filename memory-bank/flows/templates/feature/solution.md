---
title: "FT-XXX: Solution Template"
doc_kind: feature
doc_function: template
purpose: Governed wrapper-шаблон для canonical `solution.md`, который фиксирует selected design и accepted feature-local solution decisions внутри feature package.
derived_from:
  - ../../feature-flow.md
  - ../../../dna/frontmatter.md
status: active
audience: humans_and_agents
template_for: feature
template_target_path: ../../../features/FT-XXX/solution.md
canonical_for:
  - feature_template_solution
---

# FT-XXX: Solution

Этот файл описывает wrapper-template. Инстанцируемый `solution.md` живет ниже как embedded contract и копируется без wrapper frontmatter и history.

## Wrapper Notes

Создавай `solution.md` только после того, как sibling `feature.md` достиг `Problem Ready`. Этот документ становится canonical owner selected design внутри feature package.

`solution.md` обязан:

- ссылаться на canonical `REQ-*` из sibling `feature.md`;
- фиксировать selected design через `SOL-*`;
- переносить сюда принятые feature-local решения через `SD-*`;
- ссылаться на accepted ADR, если решение выходит за границы локальной фичи.

`solution.md` не должен определять business requirements, scope, acceptance criteria, canonical checks, evidence contract или execution sequencing.

Для small feature допустима короткая форма: один `SOL-*`, компактный `Change Surface` и минимальная requirement mapping. Дополнительные секции включай только когда они реально нужны.

### Frontmatter Quick Ref

Полная schema — в [../../../dna/frontmatter.md](../../../dna/frontmatter.md). Для стандартного solution-документа достаточно:

| Поле | Обязательность | Значения / default |
|---|---|---|
| `title` | required | `"FT-XXX: Solution"` |
| `doc_kind` | required | `feature` |
| `doc_function` | required | `canonical` |
| `purpose` | required | 1-2 предложения |
| `status` | required | `draft` → `active` → `archived` |
| `derived_from` | required для active | `feature.md` и optional ADR refs |
| `audience` | recommended | `humans_and_agents` |
| `must_not_define` | recommended | что документ НЕ определяет |

`delivery_status` здесь не используется: lifecycle owner-ом delivery-единицы остается sibling `feature.md`.

## Instantiated Frontmatter

```yaml
title: "FT-XXX: Solution"
doc_kind: feature
doc_function: canonical
purpose: "Canonical solution-документ для FT-XXX. Фиксирует selected design, accepted local decisions и solution-level contracts без переопределения problem space."
derived_from:
  - feature.md
  # Optional:
  # - ../../adr/ADR-XXX-short-name.md
status: draft
audience: humans_and_agents
must_not_define:
  - ft_xxx_scope
  - ft_xxx_acceptance_criteria
  - ft_xxx_delivery_status
  - implementation_sequence
```

## Instantiated Body

```markdown
# FT-XXX: Solution

## Selected Design

- `SOL-01` Какой design block выбран и какую часть problem space он закрывает.
- `SOL-02` Какой ещё solution element нужен, если одного недостаточно.

## Requirement Mapping

Как selected design покрывает canonical `REQ-*` из sibling `feature.md`.

| Requirement ID | Solution refs | Notes |
| --- | --- | --- |
| `REQ-01` | `SOL-01`, `SD-01` | Чем именно решение закрывает требование |
| `REQ-02` | `SOL-02` | Какая часть design отвечает за coverage |

## Accepted Local Decisions

Здесь живут только **принятые** feature-local решения. Непринятые блокеры остаются в `feature.md` как `DEC-*`.

- `SD-01` Какой локальный design-choice принят и почему.

## Change Surface

| Ref | Surface | Type | Why it changes |
| --- | --- | --- | --- |
| `SOL-01` | `path/or/component` | code / config / doc / data | Почему это входит в chosen solution |

## Internal Flow

1. `SOL-01` Что приходит на вход.
2. `SOL-02` Что система делает.
3. `SOL-03` Что получается на выходе.

## Contracts

Опиши concrete interfaces, payloads, schema changes или internal handoffs, если они являются частью выбранного решения.

| Contract ID | Related refs | Input / Output | Producer / Consumer | Notes |
| --- | --- | --- | --- | --- |
| `CTR-01` | `SOL-01`, `REQ-01` | Что меняется | Кто пишет / кто читает | Что важно соблюдать |

## Failure Modes

- `FM-01` Что может пойти не так у выбранного решения и как это влияет на delivery.

## Rollout / Backout

- `RB-01` Как локально и безопасно включать выбранное решение.
- `RB-02` Как локально и безопасно откатить выбранное решение.

## ADR Dependencies

Если solution зависит от ADR, зафиксируй это явно.

| ADR | Current `decision_status` | Used for | Constraint on implementation |
| --- | --- | --- | --- |
| [../../adr/ADR-XXX.md](../../adr/ADR-XXX.md) | `accepted` | Для какого solution-choice это нужно | Что нельзя делать, пока ADR в другом статусе |
```

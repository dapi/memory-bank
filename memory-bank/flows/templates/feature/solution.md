---
title: "FT-XXX: Solution Template"
doc_kind: feature
doc_function: template
purpose: Governed wrapper-шаблон для canonical `solution.md`, который фиксирует selected design, to-be C4 architecture model и accepted feature-local solution decisions внутри feature package.
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

Создавай `solution.md` только после того, как sibling `feature.md` достиг `Problem Ready`. Этот документ становится canonical owner selected design и to-be C4 architecture model внутри feature package.

`solution.md` обязан:

- ссылаться на canonical `REQ-*` из sibling `feature.md`;
- фиксировать selected design через `SOL-*`;
- фиксировать to-be C4 architecture model на тех уровнях C4 Model, которые нужны для конкретной фичи;
- фиксировать target architecture, invariants, concrete contracts и fallback/error semantics выбранного решения;
- переносить сюда принятые feature-local решения через `SD-*`;
- ссылаться на accepted ADR, если решение выходит за границы локальной фичи.

`solution.md` не должен определять business requirements, scope, acceptance criteria, canonical checks, evidence contract или execution sequencing.

Optional support docs (`runtime-surfaces.md`, `ui-reference/README.md`, `use-cases/README.md`) являются downstream/reference относительно `solution.md`. Не добавляй их в `derived_from` solution-документа; если нужно, ссылайся на них из body как related context.

Для small feature допустима короткая форма: один `SOL-*`, компактный `Change Surface`, минимальная requirement mapping и explicit C4 level selection с `Include? = no` для лишних уровней. Дополнительные секции включай только когда они реально нужны.

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
purpose: "Canonical solution-документ для FT-XXX. Фиксирует selected design, to-be C4 architecture model, accepted local decisions и solution-level contracts без переопределения problem space."
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
  - detailed_current_system_inventory
  - implementation_sequence
```

## Instantiated Body

```markdown
# FT-XXX: Solution

## Selected Design

- `SOL-01` Какой design block выбран и какую часть problem space он закрывает.
- `SOL-02` Какой ещё solution element нужен, если одного недостаточно.
- `SOL-03` Какой результат производит выбранный design, если это нужно зафиксировать отдельно.

## Requirement Mapping

Как selected design покрывает canonical `REQ-*` из sibling `feature.md`.

| Requirement ID | Solution / architecture refs | Notes |
| --- | --- | --- |
| `REQ-01` | `SOL-01`, `C4-L2-01`, `SD-01` | Чем именно решение закрывает требование |
| `REQ-02` | `SOL-02`, `SOL-03`, `C4-L3-01` | Какая часть design отвечает за coverage |

## To-Be C4 Model

Выбери только те уровни C4 Model, которые нужны, чтобы объяснить внедряемую to-be architecture этой фичи. Используй только L1 System Context, L2 Container и L3 Component. Нижний уровень добавляй только если он раскрывает новые или измененные boundaries, runtime relationships, containers, components или contract ownership. Если уровень не нужен, оставь `Include? = no` и коротко объясни почему.

| Level | Include? | Model refs | Selection rationale |
| --- | --- | --- | --- |
| System Context (L1) | yes / no | `C4-L1-01` / `none` | Включай, если меняются external actors, external systems или system boundary |
| Container (L2) | yes / no | `C4-L2-01` / `none` | Включай, если меняются runtime/deployment boundaries, containers, data stores или external integration paths |
| Component (L3) | yes / no | `C4-L3-01` / `none` | Включай, если меняются components/modules/services внутри container |

Для выбранных уровней добавь diagram или table. Таблица допустима, если C4 diagram не добавляет ясности.

| C4 ref | Level | Element / relationship | To-be architecture statement | Related refs |
| --- | --- | --- | --- | --- |
| `C4-L2-01` | Container | Какой container, datastore или external system меняется | Какая to-be связь или responsibility появляется | `SOL-01`, `CTR-01` |
| `C4-L3-01` | Component | Какой component/module/service меняется | Как он участвует в выбранном решении | `SOL-02`, `SD-01` |

## Target Architecture

Опиши to-be architecture как набор устойчивых responsibilities, boundaries, invariants и decision tables. Не дублируй current-state inventory: если нужна подробная runtime разведка, вынеси её в optional `runtime-surfaces.md` и ссылайся отсюда.

### Architecture Invariants

- Какой invariant выбранное решение обязано сохранять.
- Какой boundary нельзя пересекать при реализации.

### Target Shape

| Layer / responsibility | To-be role | Boundary / non-owner | Related refs |
| --- | --- | --- | --- |
| `responsibility-name` | Какая ответственность появляется или меняется | Что этот слой не делает | `SOL-01`, `C4-L2-01`, `CTR-01` |

### Decision / Resolution Table

Добавь таблицу, если поведение зависит от state, mode, fallback или error class.

| Condition / state | Decision | User-visible or system-visible result | Observability / evidence | Related refs |
| --- | --- | --- | --- | --- |
| Какой входной state или mode | Какое решение принимает система | Что видит пользователь или downstream consumer | Как это диагностируется | `SOL-01`, `FM-01`, `RB-01` |

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

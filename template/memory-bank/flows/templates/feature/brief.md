---
title: "FT-XXX: Brief Template"
doc_kind: feature
doc_function: template
purpose: Governed wrapper-шаблон для canonical `brief.md` в AI-driven development. Фиксирует, как инстанцировать problem-space intent, scope и machine-checkable verify без смешения wrapper и целевого frontmatter.
derived_from:
  - ../../feature.md
  - ../../feature-requirements.md
  - ../../behavior-specification.md
  - ../../feature-artifact-catalog.md
  - ../../../dna/frontmatter.md
  - ../../../engineering/testing-policy.md
status: active
audience: humans_and_agents
template_for: feature
template_target_path: ../../../features/FT-XXX/brief.md
canonical_for:
  - feature_brief_template
---

# FT-XXX: Feature Name

Этот файл описывает wrapper-template. Инстанцируемый `brief.md` живет ниже как embedded contract и копируется без wrapper frontmatter и history.

## Wrapper Notes

Используй этот шаблон для problem-space документа новых feature packages. `brief.md` фиксирует problem, outcome, scope/non-scope, validation profile decision и verify contract delivery-единицы.

Если фича меняет API, event, schema, file format, CLI, env contract, security boundary, financial calculation, integration contract, rollout/backout или требует alternatives/trade-off reasoning, зафиксируй `Design required: yes` и создай sibling `design.md` по шаблону `design.md`. Новые пакеты держат substantial design только в `design.md` / design-pack.

Optional companions выбирай по [Feature Artifact Catalog](../../feature-artifact-catalog.md). Не копируй весь каталог в feature и не создавай placeholders: Artifact Routing Decision перечисляет только выбранные artifacts и material omissions, которые важно объяснить reviewers.

Для observable behavior применяй [Behavior Specification Practice](../../behavior-specification.md). Compact feature может оставить однострочный `SC-*`, если context, event и outcome однозначны. При нескольких rules/branches, significant edge/error behavior или изменении user/API/event/operational contract используй structured `Given / When / Then` examples.

Используй стабильные идентификаторы по taxonomy из [../../feature-requirements.md#stable-identifiers](../../feature-requirements.md#stable-identifiers).

### Frontmatter Quick Ref

Полная schema — в [../../../dna/frontmatter.md](../../../dna/frontmatter.md). Для стандартного feature достаточно:

| Поле | Обязательность | Значения / default |
|---|---|---|
| `title` | required | `"FT-XXX: Name"` |
| `doc_kind` | required | `feature` |
| `doc_function` | required | `canonical` |
| `purpose` | required | 1-2 предложения |
| `status` | required | `draft` → `active` → `archived` |
| `derived_from` | required для active | upstream-документы |
| `delivery_status` | required для lifecycle-owning `brief.md` | `planned` → `in_progress` → `done` / `cancelled` |
| `audience` | recommended | `humans_and_agents` |
| `must_not_define` | recommended | что документ НЕ определяет |

## Instantiated Frontmatter

```yaml
title: "FT-XXX: Feature Name"
doc_kind: feature
doc_function: canonical
purpose: "Canonical brief для delivery-единицы. Фиксирует problem space, scope, validation profile и verify без смешения с solution space или execution plan."
derived_from:
  - ../../flows/feature.md
  - ../../flows/feature-requirements.md
  # Optional:
  # - ../../product/context.md
  # - ../../domain/rules.md
  # - ../../prd/PRD-XXX-short-name.md
  # - ../../use-cases/UC-XXX-short-name.md
status: draft
delivery_status: planned
audience: humans_and_agents
must_not_define:
  - implementation_sequence
  - solution_space
```

## Instantiated Body

```markdown
# FT-XXX: Feature Name

## What

### Requirement applicability and classification

For every baseline class in [Feature Requirements, Identifiers And Traceability](../../flows/feature-requirements.md#requirement-taxonomy-and-traceability), select `applicable`, `not-applicable` with rationale, or `covered-upstream` with reference. Do not create `FR-*`/`NFR-*`; record the class on `REQ-*`.

| Class | Decision | Trigger / rationale / upstream reference | Requirement IDs |
| --- | --- | --- | --- |
| stakeholder / product | applicable / not-applicable / covered-upstream |  | `REQ-01` / none |
| functional | applicable |  | `REQ-01` |
| performance | applicable / not-applicable / covered-upstream |  |  |
| quality attribute | applicable / not-applicable / covered-upstream |  |  |
| interface | applicable / not-applicable / covered-upstream |  |  |
| data | applicable / not-applicable / covered-upstream |  |  |
| security | applicable / not-applicable / covered-upstream |  |  |
| safety | applicable / not-applicable / covered-upstream |  |  |
| regulatory / compliance | applicable / not-applicable / covered-upstream |  |  |
| operational | applicable / not-applicable / covered-upstream |  |  |
| compatibility | applicable / not-applicable / covered-upstream |  |  |
| deployment / rollout | applicable / not-applicable / covered-upstream |  |  |
| constraint | applicable / not-applicable / covered-upstream |  | `CON-01` / none |
| verification / acceptance | applicable | Every applicable `REQ-*` needs proof. | `SC-01`, `EC-01`, `CHK-01`, `EVID-01` |

| Requirement ID | Class | Normative measurable statement / threshold | Source / rationale | Priority / owner | Verification method |
| --- | --- | --- | --- | --- | --- |
| `REQ-01` | functional | The system shall … | issue / upstream reference | must / owner | test / inspection / analysis / demonstration |

### Problem

Какой симптом, ограничение или возможность делает фичу нужной. Если общий контекст уже зафиксирован upstream, здесь опиши только feature-specific вопрос delivery.

Если существует upstream PRD, этот раздел фиксирует только feature-specific delta относительно PRD, а не переписывает весь продуктовый документ.

Если существует upstream use case, здесь фиксируется feature-specific изменение или реализация этого сценария, а не весь проектный flow целиком.

### Outcome

Опиши outcome как измеримую таблицу.

Если численный success threshold относится только к этой delivery-единице, фиксируй его здесь. Поднимать threshold upstream стоит только после появления shared owner для нескольких feature.

| Metric ID | Metric | Baseline | Target | Measurement method |
| --- | --- | --- | --- | --- |
| `MET-01` | Что измеряем | От чего стартуем | Что считаем успехом | Как проверяем |

### Scope

- `REQ-01` Что обязательно входит в deliverable.
- `REQ-02` Что еще обязательно входит в deliverable.

### Non-Scope

- `NS-01` Что сознательно исключено.
- `NS-02` Что агент не должен додумывать или реализовывать сам.

### Constraints / Assumptions

- `ASM-01` На что сейчас опираемся.
- `CON-01` Что прямо ограничивает problem space, verify или допустимый класс решений.
- `DEC-01` Какое решение еще не принято и что именно оно блокирует.

## Design Requirement Decision

Зафиксируй, нужен ли design layer и его documentary design pack. Это gate
decision, а не выбранное решение: не пересказывай selected solution, contracts,
failure modes или rollout/backout в `brief.md`.

| Decision | Reason | Downstream owner |
| --- | --- | --- |
| `Design required: yes/no` | Почему design layer нужен или не нужен | Design pack с root `design.md` / `none` |

## Artifact Routing Decision

Секция optional. Используй ее, когда кроме core `README.md` + `brief.md` нужен companion artifact или важно явно объяснить его отсутствие. Перечисляй только выбранные artifacts и material omissions; полный список не копируй.

| Artifact | Decision | Trigger / reason | Route / owner |
| --- | --- | --- | --- |
| `use-cases/README.md` / `runtime-surfaces.md` / `ui-reference/README.md` / другой artifact из catalog | selected / omitted | Какую неоднозначность снимает или почему не нужен | Planned path и canonical owner / `none` |

## Validation Profile Decision

Выбери один profile по [`../../engineering/validation-profiles.md`](../../engineering/validation-profiles.md). Эта секция — canonical owner решения; `implementation-plan.md` ссылается на неё и задаёт конкретные suites/checkpoints без повторного выбора profile.

| Profile | Triggers / rationale | Downgrade approval |
| --- | --- | --- |
| `documentation` / `low-risk` / `standard` / `high-risk` / `release-deployment` | Какие triggers проверены и почему выбранный minimum достаточен | Human approval ref, если trigger требует downgrade; иначе `none` |

## Verify

`Verify` задает canonical test case inventory для delivery-единицы: positive scenarios через `SC-*`, feature-specific negative coverage через `NEG-*`, executable checks через `CHK-*` и evidence через `EVID-*`.

### Exit Criteria

- `EC-01` Проверяемый признак готовности.
- `EC-02` Еще один обязательный признак готовности.

### Traceability matrix

| Requirement ID | Problem refs | Acceptance refs | Checks | Evidence IDs |
| --- | --- | --- | --- | --- |
| `REQ-01` | `ASM-01`, `CON-01`, `DEC-01` | `EC-01`, `SC-01` | `CHK-01` | `EVID-01` |
| `REQ-02` | `ASM-01`, `CON-01` | `EC-02`, `SC-02`, `NEG-01` | `CHK-01`, `CHK-02` | `EVID-01`, `EVID-02` |

### Acceptance Scenarios

Для compact feature допустима однострочная форма, если она однозначно задаёт
существенный context, event и observable outcome:

- `SC-01` Основной happy path: при <context>, когда <event>, система публикует или показывает <observable outcome>.

Для structured BDD используй форму ниже. Rule refs ссылаются на canonical
`UC/BR/REQ`, но не копируют их semantics.

#### SC-02: Название различающего поведения

- Rule refs: `UC-XXX/BR-01`, `REQ-02`
- Given: существенное начальное состояние
- When: одно значимое событие или действие
- Then: observable outcome для пользователя, оператора или external system
- And: дополнительный observable outcome, только если нужен verdict
- Checks: `CHK-01`

### Negative / Edge Scenarios

Добавляй `NEG-*`, когда negative или boundary behavior меняет acceptance verdict.

#### NEG-01: Название error или edge behavior

- Rule refs: `UC-XXX/EX-01`, `REQ-02`
- Given: существенное boundary-состояние
- When: событие или действие
- Then: наблюдаемый отказ, fallback или preserved state
- Checks: `CHK-02`

### Checks

Verify должен быть исполнимым.

| Check ID | Covers | How to check | Expected result | Evidence path |
| --- | --- | --- | --- | --- |
| `CHK-01` | `EC-01`, `SC-01` | Команда или процедура | Что считаем успехом | Где лежит артефакт |
| `CHK-02` | `NEG-01` | Команда или процедура | Какой negative / edge verdict ожидается | Где лежит артефакт |

### Test matrix

| Check ID | Evidence IDs | Evidence path |
| --- | --- | --- |
| `CHK-01` | `EVID-01` | `artifacts/ft-xxx/verify/chk-01/` |
| `CHK-02` | `EVID-02` | `artifacts/ft-xxx/verify/chk-02/` |

### Evidence

- `EVID-01` Какой артефакт обязан появиться после проверки.
- `EVID-02` Evidence negative / edge verdict или approved manual-only gap.

### Evidence contract

| Evidence ID | Artifact | Producer | Path contract | Reused by checks |
| --- | --- | --- | --- | --- |
| `EVID-01` | Лог, отчет, скриншот или sample output | verify-runner / human | `artifacts/ft-xxx/verify/chk-01/` | `CHK-01` |
| `EVID-02` | Лог, отчет или sample output для negative/edge behavior | verify-runner / human | `artifacts/ft-xxx/verify/chk-02/` | `CHK-02` |

### Requirement-to-realization traceability

`brief.md` owns the requirement chain; `design.md` owns selected solution facts and the plan supplies the exact target. A supporting change without a direct `REQ-*` needs its explicit rationale.

| Requirement | Acceptance / check / evidence | Solution or decision | Exact realization target | Supporting rationale / review or CI result |
| --- | --- | --- | --- | --- |
| `REQ-01` | `EC-01`, `SC-01`, `CHK-01`, `EVID-01` | `SOL-01` / design-not-required decision | `path/file.ext#SymbolOrSection` | none / CI URL |
```

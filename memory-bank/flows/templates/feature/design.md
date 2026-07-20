---
title: "FT-XXX: Design Template"
doc_kind: feature
doc_function: template
purpose: Governed wrapper-шаблон для feature-local `design.md`. Фиксирует solution-space слой: выбранный подход, architecture coverage, contracts, design verification и design-pack routing без смешения с problem space или execution contract.
derived_from:
  - ../../feature.md
  - ../../feature-artifact-catalog.md
  - ../../../dna/frontmatter.md
status: active
audience: humans_and_agents
template_for: feature
template_target_path: ../../../features/FT-XXX/design.md
canonical_for:
  - feature_design_template
---

# FT-XXX: Design

Этот файл описывает wrapper-template. Инстанцируемый `design.md` живет ниже как embedded contract и копируется без wrapper frontmatter и history.

## Wrapper Notes

Создавай `design.md`, когда фича требует solution-space reasoning: выбор подхода, trade-offs, contracts, invariants, failure modes, rollout/backout, ADR/C4/data-flow/diagram dependencies или design-pack из нескольких документов.

На стадии анализа обязательно заполни C4 applicability decision, Architecture Coverage Decision и risk-based Design Verification. C4 artifact обязателен только когда trigger из [feature.md#c4-analysis-requirements](../../feature.md#c4-analysis-requirements) требует C1/C2/C3/C4; отдельные diagrams/contracts остаются conditional, но coverage analysis обязателен для любого required `design.md`.

`design.md` не заменяет `brief.md`: требования, acceptance criteria и evidence contract остаются в `brief.md`. `design.md` также не является execution plan: file-level touchpoints, атомарные шаги, команды тестов и checkpoints принадлежат `implementation-plan.md`.

Если solution-space разбит на несколько артефактов, `design.md` становится индексом design-pack и фиксирует owner-а каждого design fact. Не дублируй canonical факты из ADR, C4, data-flow или других design docs; ссылайся на них.

## Instantiated Frontmatter

```yaml
title: "FT-XXX: Design"
doc_kind: feature
doc_function: canonical
purpose: "Solution-space документ для FT-XXX. Фиксирует выбранный подход, architecture coverage, contracts, design verification и design-pack routing без переопределения problem space или execution contract."
derived_from:
  - brief.md
status: draft
audience: humans_and_agents
must_not_define:
  - ft_xxx_scope
  - ft_xxx_acceptance_criteria
  - ft_xxx_evidence_contract
  - implementation_sequence
```

## Instantiated Body

```markdown
# FT-XXX: Design

## Design Pack

Если design-pack состоит только из этого файла, оставь одну строку `design.md`. Если есть ADR, C4, data-flow, interaction contract, sequence diagram, migration design или другая полезная companion view, добавь ее в таблицу и укажи ownership. Не создавай дополнительные artifacts только ради заполнения таблицы.

| Artifact | Role | Owns |
| --- | --- | --- |
| `design.md` | Feature-local solution owner | `SOL-*`, `ALT-*`, `TRD-*`, `C4-*`, architecture coverage, design verification, feature-local `CTR-*`, `INV-*`, `FM-*`, `RB-*` |
| `contracts/<name>.md` | Optional delegated contract owner | Только явно перечисленные `CTR-*`; selected solution остается здесь |
| `diagrams/<name>-sequence.md` | Optional temporal reference view | `SEQ-*` projection canonical solution / contract facts; новых решений не принимает |
| `../../adr/ADR-XXX.md` | Architecture decision | Какой design choice принадлежит ADR |

## Context

Коротко опиши design problem: почему требования из `brief.md` требуют явного решения, какие upstream docs или constraints важны для выбора.

## C4 Applicability

Решение принимается до `Solution Ready`. Выбери минимальный уровень C4 или явно зафиксируй, что C4 не нужен.

| C4 ID | Decision | Trigger / reason | Artifact |
| --- | --- | --- | --- |
| `C4-00` | `not required` / `C1` / `C2` / `C3` / `C4` | Почему C4 не нужен или какой trigger требует выбранный уровень | `none` / ссылка на diagram |

### C4 Artifact

Если `C4-00` не `not required`, добавь diagram или ссылку на artifact design-pack. Используй самый низкий достаточный уровень:

- `C1` - System Context: actors/external systems/trust boundaries.
- `C2` - Container: deployable/runtime nodes, queues, stores, protocols.
- `C3` - Component: modules/services/state machines внутри container.
- `C4` - Code: только когда class/interface-level structure является архитектурным решением.

## Architecture Coverage Decision

Для каждого аспекта выбери `covered` или обоснованный `N/A`. Analysis обязателен; дополнительные artifacts создавай только по trigger. В `Canonical owner / refs` укажи документ-владелец и stable IDs, а supporting view не считай canonical owner. Отдельный solution-space artifact должен входить в Design Pack.

| Aspect | Status | Canonical owner / refs | Supporting view / artifact | Reason if N/A / coverage note |
| --- | --- | --- | --- | --- |
| Components / responsibilities | `covered` / `N/A` | `design.md` `SOL-*` / `SD-*` или accepted ADR | C3 / component map / `none` | Где определены ответственности и provided/required interfaces или почему аспект неприменим |
| Connectors / interactions | `covered` / `N/A` | `design.md` `CTR-*` или `contracts/<name>.md` | sequence / `none` | Где определены механизм и значимые interaction semantics или почему аспект неприменим |
| Configuration / topology | `covered` / `N/A` | `design.md` `SOL-*` / `SD-*` или accepted ADR | C2/C3 / data-flow / `none` | Где определены bindings, direction, connector kind, optional links и affected topology или почему аспект неприменим |
| Behavioral semantics | `covered` / `N/A` | `design.md` `SOL-*` / `CTR-*` / `INV-*` / `FM-*` | sequence / state machine / `none` | Где определены ordering, transitions и failure behavior или почему аспект неприменим |
| Quality / evolution concerns | `covered` / `N/A` | `brief.md` `CON-*`; `design.md` `INV-*` / `FM-*` / `RB-*`; accepted ADR | analysis artifact / `none` | Где закрыты relevant quality, compatibility и evolution risks или почему аспект неприменим |

## Selected Solution

- `SOL-01` Выбранный элемент решения и почему он закрывает `REQ-*`.
- `SOL-02` Второй элемент решения, если нужен.

## Alternatives Considered

| Alternative ID | Option | Why not selected |
| --- | --- | --- |
| `ALT-01` | Альтернативный подход | Причина отказа или отложенного выбора |

## Trade-offs

| Trade-off ID | Decision | Benefit | Cost / Risk |
| --- | --- | --- | --- |
| `TRD-01` | Какой компромисс принимаем | Что выигрываем | Что платим или мониторим |

## Accepted Local Decisions

Здесь живут только принятые feature-local decisions. Decisions reusable, architectural или cross-feature уровня выносятся в ADR.

- `SD-01` Какое локальное решение принято и почему оно не требует ADR.

## Contracts

Connector — first-class механизм или binding, связывающий стороны решения: API call, event, queue, callback, shared store/file access, cache interaction, authentication handoff, locking/concurrency mechanism или runtime/config binding. Не смешивай connector kind с protocol/format (`schema`, encoding) или parties/roles (producer, consumer, provider, initiator, target). Для значимого connector зафиксируй применимые roles, protocol/format и direction, sync/async boundary, ordering/delivery, timeout/retry/idempotency, trust boundary, failure/degradation, compatibility/versioning и observability. Компактное описание оставь здесь; отдельный interaction contract создавай только при самостоятельной review boundary. Не добавляй реалистичные секреты, production IDs или file-level implementation steps.

| Contract ID | Connector / direction | Roles and sync boundary | Guarantees / failure / evolution semantics |
| --- | --- | --- | --- |
| `CTR-01` | Механизм и `initiator -> target` | Producer/consumer; sync/async | Protocol/format, ordering/delivery, timeout/retry/idempotency, trust, degradation, compatibility, observability |

## Invariants

- `INV-01` Что должно оставаться истинным независимо от implementation path.

## Failure Modes

- `FM-01` Что может пойти не так и как решение должно это ограничить.

## Rollout / Backout

| Stage ID | Stage | Entry condition | Backout |
| --- | --- | --- | --- |
| `RB-01` | Как включается изменение | Что должно быть доказано до входа | Как вернуть безопасное состояние |

## Design Verification

Для каждой строки выбери анализ по риску. `required: no` требует причины; `required: yes` — method и завершенный result/evidence до `Solution Ready`. Не создавай отдельный artifact, если достаточно compact result здесь.

| Analysis | Required | Reason / risk | Method | Result / evidence |
| --- | --- | --- | --- | --- |
| Contract compatibility | yes / no | Что делает анализ нужным или неприменимым | Schema diff, consumer review, compatibility matrix | Вывод или ссылка |
| State / transition completeness | yes / no | Есть ли non-trivial states/transitions | State-table review, model checking, scenario walk-through | Вывод или ссылка |
| Failure propagation | yes / no | Есть ли distributed/degradation risk | Failure-mode analysis, fault tree, simulation | Вывод или ссылка |
| Concurrency / ordering | yes / no | Есть ли races, duplicates или parallel writers | Interleaving review, sequence analysis, test/prototype | Вывод или ссылка |
| Security boundaries | yes / no | Меняются ли auth/trust/data boundaries | Threat analysis, control review | Вывод или ссылка |
| Capacity / latency | yes / no | Меняется ли load/latency-sensitive path | Estimate, benchmark, load model | Вывод или ссылка |
| Migration / evolution safety | yes / no | Нужны ли mixed versions, staged rollout или data/config migration | Compatibility/migration review, rehearsal | Вывод или ссылка |

## ADR / External Design Dependencies

| Artifact | Current status | Used for | Rule |
| --- | --- | --- | --- |
| `../../adr/ADR-XXX.md` | `proposed` / `accepted` | Какой выбор или baseline задает | `proposed` не считается finalized design |

## Traceability

| Requirement ID | Solution refs | Contracts / invariants | Failure / rollout refs |
| --- | --- | --- | --- |
| `REQ-01` | `SOL-01`, `TRD-01`, `C4-00`, `SD-01` | `CTR-01`, `INV-01` | `FM-01`, `RB-01` |
```

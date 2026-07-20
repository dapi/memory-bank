---
title: "FT-XXX: API Contract Template"
doc_kind: feature
doc_function: template
purpose: Governed wrapper-шаблон optional feature-local API/event/provider contract. Читать, когда detailed contract semantics заслуживают отдельного design-pack owner вместо разрастания `design.md`.
derived_from:
  - ../../feature.md
  - ../../feature-artifact-catalog.md
  - ../../../dna/frontmatter.md
status: active
audience: humans_and_agents
template_for: feature
template_target_path: ../../../features/FT-XXX/contracts/api-contract.md
canonical_for:
  - feature_api_contract_template
---

# FT-XXX: API Contract Template

Этот файл описывает wrapper-template. Инстанцируемый contract живет в `contracts/<name>.md` внутри feature package и создается только по trigger из `feature.md`.

## Wrapper Notes

Создавай отдельный contract, когда API, event, schema или provider boundary содержит достаточно operations, fields, statuses, errors, compatibility или idempotency rules, чтобы inline `CTR-*` в `design.md` стал трудно проверяемым.

Если contract компактен, оставь его в `design.md`. Отдельный файл не является обязательной частью feature package и не должен появляться как placeholder.

`design.md` обязан индексировать contract в Design Pack, перечислить делегированные `CTR-*` и связать их с `SOL-*` и `REQ-*`. Contract не выбирает solution, не меняет scope и не задает implementation sequence.

## Instantiated Frontmatter

```yaml
title: "FT-XXX: <Boundary Name> Contract"
doc_kind: feature
doc_function: canonical
purpose: "Feature-local contract для <boundary>. Фиксирует operations/messages, fields, validation, compatibility, errors и idempotency в пределах решения FT-XXX."
derived_from:
  - ../brief.md
  - ../design.md
status: draft
audience: humans_and_agents
must_not_define:
  - ft_xxx_scope
  - ft_xxx_selected_solution
  - ft_xxx_acceptance_criteria
  - implementation_sequence
```

## Instantiated Body

````markdown
# FT-XXX: <Boundary Name> Contract

## Role And Ownership

| Role | Value |
| --- | --- |
| Boundary | Какой API / event / schema / provider boundary описан |
| Owns | Какие `CTR-*` делегированы этому документу из `design.md` |
| Does not own | Scope, selected solution, acceptance, execution sequencing |
| Producers / consumers | Кто пишет, вызывает, публикует или читает contract |

## Contract Status And Compatibility

| Field | Value |
| --- | --- |
| Status | draft / proposed / accepted / deprecated |
| Version | Версия contract или `unversioned` с причиной |
| Compatibility | backward-compatible / breaking / migration required |
| Source authority | Provider docs, accepted ADR, upstream contract or repo baseline |

## Operations / Messages

| Contract ID | Operation / message | Direction | Purpose | Related refs |
| --- | --- | --- | --- | --- |
| `CTR-01` | Method, endpoint, event or schema name | producer -> consumer | Какая capability предоставляется | `REQ-01`, `SOL-01` |

## Request / Input

| Field | Required | Type / format | Semantics | Validation / default |
| --- | --- | --- | --- | --- |
| `field_name` | yes / no / conditional | string / object / enum | Что означает поле | Ограничения без production secrets |

## Response / Output

| Field | Presence | Type / format | Semantics | Consumer behavior |
| --- | --- | --- | --- | --- |
| `field_name` | always / conditional | string / object / enum | Что означает поле | Как consumer интерпретирует значение |

## Status And State Mapping

| External / wire state | Meaning | Terminal | Feature behavior | Related refs |
| --- | --- | --- | --- | --- |
| `state` | Что означает | yes / no | Какой semantic result допустим | `CTR-01`, `FM-01` |

## Errors And Failure Semantics

| Error / condition | Retryable | Required behavior | Observability | Related refs |
| --- | --- | --- | --- | --- |
| `error_code` | yes / no / conditional | Fail, retry, compensate or escalate | Как диагностируется без sensitive payload | `FM-01` |

## Idempotency And Ordering

| Rule | Contract |
| --- | --- |
| Idempotency key | Source, scope, reuse and conflict semantics |
| Duplicate delivery | Как producer/consumer распознают и обрабатывают duplicate |
| Ordering | Какие ordering guarantees существуют или отсутствуют |
| Timeout / retry | Как retry связан с idempotency и terminal state |

## Security And Sensitive Data

- Authentication / authorization boundary.
- Integrity or signature verification.
- Sensitive fields that must not enter logs, examples or evidence.
- Trust-boundary refs из `design.md`, C4 или accepted ADR.

## Examples

Используй synthetic values. Не добавляй реальные credentials, production IDs, personal data или usable secrets.

```json
{
  "example": "synthetic-value"
}
```

## Traceability

| Contract IDs | Requirements | Solution refs | Failure / rollout refs | Sequence refs |
| --- | --- | --- | --- | --- |
| `CTR-01` | `REQ-01` | `SOL-01`, `SD-01` | `FM-01`, `RB-01` | `SEQ-01` / none |
````

---
title: "FT-XXX: Interaction Contract Template"
doc_kind: feature
doc_function: template
purpose: Governed wrapper-шаблон optional feature-local interaction contract. Читать, когда detailed connector semantics заслуживают отдельного design-pack owner вместо разрастания `design.md`.
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
  - feature_interaction_contract_template
---

# FT-XXX: Interaction Contract Template

Этот файл описывает wrapper-template. Инстанцируемый contract живет в `contracts/<name>.md` внутри feature package и создается только по trigger из `feature.md`.

## Wrapper Notes

Создавай отдельный contract, когда API call, event, queue, callback, shared file/store access, cache interaction, authentication handoff, locking/concurrency mechanism или runtime/config binding содержит достаточно самостоятельных semantics, чтобы inline `CTR-*` в `design.md` стал трудно проверяемым. Schema/encoding фиксируй как protocol/format, а provider — как party/role, не как connector kind.

Если contract компактен, оставь его в `design.md`. Отдельный файл не является обязательной частью feature package и не должен появляться как placeholder.

Инстанцируй только применимые sections: operation/request/response tables подходят для wire contracts, но могут быть заменены binding/state/concurrency tables для store, cache, lock или config connector. Не заполняй неприменимые sections фиктивными данными.

Путь `api-contract.md` и `feature_api_contract_template` сохранены как compatibility aliases; семантически это общий Interaction Contract Template.

`design.md` обязан индексировать contract в Design Pack, перечислить делегированные `CTR-*` и связать их с `SOL-*` и `REQ-*`. Contract не выбирает solution, не меняет scope и не задает implementation sequence.

## Instantiated Frontmatter

```yaml
title: "FT-XXX: <Boundary Name> Contract"
doc_kind: feature
doc_function: canonical
purpose: "Feature-local interaction contract для <boundary>. Фиксирует connector roles, protocol/format, delivery, failure, compatibility и observability semantics в пределах решения FT-XXX."
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

Оставь только применимые sections. Для wire contract используй operation/request/response tables; для store, cache, lock или config connector замени их подходящими binding/state/concurrency tables. Не сохраняй и не заполняй неприменимые placeholders.

## Role And Ownership

| Role | Value |
| --- | --- |
| Boundary | Какой connector boundary описан, какие стороны он связывает и какой interaction mechanism или runtime/config binding фиксирует |
| Owns | Какие `CTR-*` делегированы этому документу из `design.md` |
| Does not own | Scope, selected solution, acceptance, execution sequencing |
| Roles | Producer / consumer / provider / initiator / target и owner каждой стороны |

## Connector Semantics

| Concern | Contract |
| --- | --- |
| Connector kind and binding | API call, event, queue, callback, shared store/file access, cache interaction, auth handoff, lock или runtime/config binding; где связаны стороны |
| Protocol / format / direction | Protocol, encoding/schema и `initiator -> target` |
| Sync / async boundary | Где caller ждёт ответ, где ownership переходит асинхронно |
| Ordering / delivery | At-most/at-least/exactly-once claim, ordering scope, duplicates and gaps |
| Timeout / retry / idempotency | Time budget, retry owner/policy и identity/deduplication semantics |
| Trust / security boundary | Authentication, authorization, integrity and sensitive-data handling |
| Failure / degradation | Propagation, isolation, fallback, compensation and terminal behavior |
| Compatibility / versioning | Supported versions, mixed-version behavior and evolution policy |
| Observability | Logs, metrics, traces, correlation and alertable failure signals |

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
| `CTR-01` | Method, endpoint, event, operation or binding name | producer -> consumer | Какая capability предоставляется | `REQ-01`, `SOL-01` |

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

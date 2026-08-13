---
title: "FT-113: Integrate BDD behavior specification practice"
doc_kind: feature
doc_function: canonical
purpose: "Фиксирует delivery-единицу по описанию BDD как практики внутри существующих Feature и Use Case Flow."
derived_from:
  - ../../flows/feature.md
status: active
delivery_status: planned
audience: humans_and_agents
source_issue: https://github.com/dapi/memory-bank/issues/113
must_not_define:
  - implementation_sequence
  - solution_space
---

# FT-113: Integrate BDD behavior specification practice

## What

### Problem

Memory Bank различает `UC-*`, `REQ-*`, `SC/NEG-*`, `CHK-*`, `EVID-*` и
feature-local `FUC-*`, но не описывает их связь в процессе discovery → concrete
examples → проверка поведения. Без общего контракта проекты могут дублировать
требования в Gherkin, сводить BDD к E2E и создавать параллельных владельцев
сценариев.

### Outcome

Memory Bank однозначно описывает переход от обсуждения поведения к concrete
examples и проверке, границу между project-level use cases и feature acceptance,
а также ownership и traceability существующих artifacts.

## Scope

- `REQ-01` Описать переход от обсуждения поведения к concrete examples и затем
  к проверке поведения.
- `REQ-02` Определить, где фиксировать найденное поведение, когда создавать
  project-level `UC-*`, когда достаточно feature-level `SC-*` / `NEG-*`, где
  живут `Given / When / Then` examples и кто владеет acceptance semantics.
- `REQ-03` Связать examples с существующими requirements, acceptance checks и
  evidence и объяснить BDD как практику анализа и проверки поведения, а не
  только E2E-инструмент, без противоречивых владельцев artifacts.

## Non-Scope

- `NS-01` Новый BDD route, каталог, identifier family или второй canonical
  owner. Это derived governance constraint из существующего Task Routing и SSoT,
  а не дословное требование issue.
- `NS-02` Обязательные Gherkin, Cucumber, browser automation или E2E: первичные
  источники [Dan North](https://dannorth.net/blog/introducing-bdd/) и
  [Cucumber BDD Guide](https://cucumber.io/docs/bdd/) описывают BDD шире
  конкретного инструмента или test level.
- `NS-03` Bulk migration downstream packages: issue её не запрашивает.

## Design Requirement Decision

| Decision | Reason | Downstream owner |
| --- | --- | --- |
| `Design required: no` | Existing governance уже назначает owners и routes; feature документирует и связывает этот contract, не выбирая новую architecture, owner hierarchy или runtime/interface solution. | none |

## Validation Profile Decision

| Profile | Triggers / rationale | Downgrade approval |
| --- | --- | --- |
| `documentation` | Generic template documentation, indexes and validation tooling only. | none |

## Behavior Specification Contract

Следующие правила являются derived governance resolution требований issue, а не
их дословной цитатой. Они выводятся из существующих SSoT owners, Task Routing и
Feature/Use Case lifecycle.

BDD применяется как практика уточнения поведения внутри выбранного delivery
flow, а не как отдельный route, каталог или обязательный E2E-инструмент.

- Устойчивый повторяющийся project-level сценарий с trigger, preconditions,
  flow и postconditions принадлежит `UC-*`.
- Правило, относящееся только к этой delivery-unit, принадлежит `REQ-*` в
  `brief.md`.
- Concrete examples хранятся как `SC-*` или `NEG-*` в `brief.md`. При
  существенной неоднозначности example использует `Given / When / Then`:
  начальное состояние, одно значимое событие и наблюдаемый результат.
- `CHK-*` фиксирует способ проверки и expected verdict; `EVID-*` фиксирует
  наблюдаемый результат проверки. Цепочка не передаёт ownership test code или
  feature-local companion.
- Существующие `FUC-*` и другие feature-local companions могут только показывать
  derived projection и mapping для review; это integration boundary, а не новая
  customer capability.
- Gherkin, Cucumber, browser automation и E2E не являются обязательными;
  выбирается самый надёжный test surface, доказывающий observable behavior.

## Verify

| Requirement | Acceptance | Checks | Evidence |
| --- | --- | --- | --- |
| `REQ-01` | `SC-01` Practice defines the transition from behavior discussion through concrete examples to verification. | `CHK-01`, `CHK-02` | `EVID-01`, `EVID-02` |
| `REQ-02` | `SC-02` A discovered behavior is routed to `UC-*`, `REQ-*`, `SC-*` or `NEG-*` by an explicit ownership criterion, and structured examples have a defined home. | `CHK-01` | `EVID-01` |
| `REQ-03` | `SC-03` Requirements and examples trace through checks to evidence without making BDD E2E-only or creating another owner. | `CHK-01`, `CHK-02` | `EVID-01`, `EVID-02` |

### Concrete Examples

#### SC-01: Переход от обсуждения к проверке

- Rule refs: `REQ-01`
- Given: команда обсуждает требуемое наблюдаемое поведение
- When: она уточняет его через concrete examples и готовит проверку
- Then: Memory Bank задаёт непрерывный путь Discussion/Discovery → Formulation
  → Verification/Automation
- Checks: `CHK-01`, `CHK-02`

#### SC-02: Маршрутизация найденного поведения

- Rule refs: `REQ-02`
- Given: discovery выявил устойчивый project flow, feature-specific rule и
  positive или negative example
- When: команда выбирает canonical owner
- Then: устойчивый reusable flow направляется в `UC-*`, feature-specific rule —
  в `REQ-*`, а examples — в `SC-*` / `NEG-*`; optional `FUC-*` остаётся derived
- Checks: `CHK-01`

#### SC-03: Проверяемая цепочка без E2E-only ограничения

- Rule refs: `REQ-03`
- Given: `SC-*` или `NEG-*` содержит context, event и observable outcome
- When: команда выбирает техническую или approved manual-only проверку
- Then: `CHK-*` связывает example с подходящим test surface, `EVID-*` содержит
  concrete carrier результата, а Gherkin, Cucumber и E2E не обязательны
- Checks: `CHK-01`, `CHK-02`

### Checks

| Check ID | How to check | Expected result |
| --- | --- | --- |
| `CHK-01` | Независимый requirements/artifact review issue #113, feature package и изменённых governance owners. | Review подтверждает полноту customer requirements, отсутствие второго owner и непрерывную traceability; все critical/important findings закрыты. |
| `CHK-02` | `ruby tools/validate-priming-manifests-test.rb`; `ruby tools/validate-priming-manifests.rb template/memory-bank`; `memory-bank-cli lint --scope-root template/memory-bank --entrypoint template/memory-bank/README.md`; `memory-bank-cli doctor --profile template`; `git diff --check`. | Validators, lint, doctor и whitespace check проходят на одной immutable candidate revision. |

### Evidence

- `EVID-01` Pending: внешний clean artifact-review record с reviewer identity,
  immutable candidate revision, findings/dispositions и итоговым verdict.
- `EVID-02` Pending: concrete local/CI carriers для всех команд `CHK-02`,
  привязанные к той же candidate revision.

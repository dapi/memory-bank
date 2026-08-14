---
title: Behavior Specification Practice
doc_kind: governance
doc_function: canonical
purpose: "Определяет BDD-практику Discovery → Formulation → Automation, качество concrete examples и traceability без создания отдельного delivery route или второго owner требований."
derived_from:
  - ../dna/governance.md
  - ../dna/principles.md
canonical_for:
  - behavior_discovery_practice
  - behavior_example_formulation_rules
  - behavior_automation_handoff
  - bdd_ownership_boundaries
  - bdd_traceability_contract
status: active
audience: humans_and_agents
---

# Behavior Specification Practice

Behavior-Driven Development (BDD) в Memory Bank — это сквозная практика
совместного уточнения, формулирования и проверки наблюдаемого поведения. Она
дополняет выбранный delivery flow, но не является отдельным route в
[`Task Routing`](routing.md) и не создаёт каталог `memory-bank/bdd/`.

BDD выполняется короткими итерациями:

```text
Discovery: что система могла бы делать
    → Formulation: что она должна делать в конкретном примере
        → Automation: что реализация фактически делает
```

## Method Sources

- Dan North, [Introducing BDD](https://dannorth.net/blog/introducing-bdd/) —
  первичный источник подхода, business-value framing и формы
  `Given / When / Then`.
- [Cucumber BDD Guide](https://cucumber.io/docs/bdd/) — современное описание
  цикла Discovery, Formulation и Automation. Cucumber является одним из
  возможных инструментов, а не обязательной частью практики.

## Ownership Model

BDD не вводит `BDD-*` identifiers. Он уточняет существующую цепочку owners:

| Fact | Canonical owner |
| --- | --- |
| Устойчивый project-level scenario | `UC-*` |
| Shared domain rule / invariant | `memory-bank/domain/` |
| Обязательное поведение delivery-единицы | `REQ-*` в feature `brief.md` |
| Положительный concrete example | `SC-*` в feature `brief.md` |
| Negative / edge concrete example | `NEG-*` в feature `brief.md` |
| Способ проверки и expected verdict | `CHK-*` в feature `brief.md` |
| Test surface, framework и execution sequencing | `implementation-plan.md` |
| Исполняемая проверка | test code |
| Наблюдаемый результат проверки | `EVID-*` |

Canonical traceability:

```text
[optional UC/BR] → REQ → SC/NEG → CHK → automated test or approved manual-only check → EVID
```

Один `UC-*` может иметь много downstream examples. Один `SC-*` или `NEG-*`
может проверяться несколькими техническими тестами. Test code не становится
owner-ом requirement или acceptance semantics.

## When To Apply Structured BDD

Concrete examples полезны для любого изменения observable behavior. Явная
структура `Given / When / Then` обязательна, когда prose оставляет material
неоднозначность, в том числе при одном из условий:

- несколько business rules, roles, branches или state transitions;
- outcome зависит от комбинации входных условий;
- есть significant negative, edge, retry, duplicate или recovery behavior;
- меняется user-visible, API, event или operational contract;
- один scenario должны одинаково понимать product, engineering и test roles.

Для compact change допустим однострочный `SC-*`, если context, event и
observable outcome всё равно однозначны. Gherkin и Cucumber не обязательны.

## Discovery

Discovery проводится до фиксации feature problem space как ready. Это может
быть conversation, асинхронный review или agent-assisted analysis, но должны
быть представлены product/domain, implementation и verification perspectives.

Обсуди малый delivery slice через конкретные examples и маршрутизируй findings
сразу к существующим owners:

| Discovery finding | Destination |
| --- | --- |
| Stable project flow | создать или обновить `UC-*` |
| Shared domain rule | соответствующий owner в `domain/` |
| Feature-specific rule | `REQ-*` |
| Positive example | `SC-*` |
| Negative / edge example | `NEG-*` |
| Unanswered verdict-changing question | `DEC-*` |
| Deferred behavior | `NS-*`, upstream backlog или отдельная delivery-unit |

Не создавай отдельный discovery transcript как canonical artifact. Optional
feature-local use-case companion может показывать derived example map, но не
принимать новые requirements, acceptance criteria или checks.

## Formulation

Каждый structured example содержит:

- **Name** — краткое описание различающего поведения;
- **Rule refs** — применимые `UC/BR/REQ` references;
- **Given** — только существенное начальное состояние;
- **When** — одно значимое событие или действие;
- **Then** — observable outcome для пользователя, оператора или external
  system;
- **Check refs** — применимые `CHK-*` либо planned mapping до `Problem Ready`.

Пример:

```markdown
#### SC-01: Успешная оплата при достаточном балансе

- Rule refs: `UC-004/BR-01`, `REQ-02`
- Given: баланс клиента не меньше суммы заказа
- When: клиент подтверждает оплату
- Then: заказ получает статус «оплачен»
- And: публикуется подтверждение платежа
- Checks: `CHK-01`
```

Scenario quality rules:

1. Используй domain language, а не selectors, endpoints, tables или class names.
2. Проверяй observable result, а не скрытое внутреннее состояние, если оно не
   является явно опубликованным contract.
3. Один example иллюстрирует одно главное правило или одну различающую branch.
4. Используй concrete values, когда они снимают неоднозначность boundary.
5. Не копируй общий `UC-*` flow или shared domain rule в каждый example.
6. Не превращай scenario в implementation procedure или длинный UI click path.

## Automation

Automation связывает accepted example с системой как проверку и направляет
реализацию, но не требует конкретного framework или test level.

- Выбирай самый низкий надёжный test surface, который доказывает observable
  behavior: unit, component, contract, integration или E2E.
- Gherkin scenario не означает обязательный browser/UI test.
- Имя, tag или metadata automated test должны сохранять ссылку на `SC-*` или
  `NEG-*`, если conventions выбранного проекта это допускают.
- Manual-only gap требует причины, процедуры и approvals из выбранного
  validation profile.
- Изменение expected behavior сначала обновляет canonical owner, затем examples,
  checks, plan и test code.

## Lifecycle Integration

### Problem Ready

- verdict-changing discovery questions разрешены либо зафиксированы как
  blocking `DEC-*`;
- каждый `REQ-*` покрыт минимум одним `SC-*` или `NEG-*`;
- required examples формулируют context, event и observable outcome;
- `CHK-*` и `EVID-*` образуют проверяемый acceptance contract.

### Solution Ready

Если design required, driving `SC-*` участвуют в cross-view correspondence, а
существенные `NEG-*` / edge examples связываются с применимыми `FM-*`,
`CTR-*`, `INV-*` или получают обоснованный `N/A`.

### Plan Ready

Test Strategy показывает для каждого required `SC-*` / `NEG-*` planned test
surface, automation, required suites, evidence и допустимые manual-only gaps.

### Done

- required examples имеют pass/fail evidence через `CHK-*` / `EVID-*`;
- required automated coverage добавлено и зелёное локально и в CI;
- manual-only gaps явно approved;
- `UC`, `brief`, derived views и executable checks не противоречат друг другу.

## Change Propagation

При material change стабильного scenario сначала обнови `UC-*` или shared
domain owner, затем feature `REQ/SC/NEG/CHK`, derived views, plan и test code.
Не оставляй одновременно два противоречащих active descriptions одного
behavior.

## Anti-Patterns

- отдельный BDD-каталог как второй owner acceptance;
- Gherkin, написанный после реализации только ради отчёта;
- все examples автоматизированы через медленные E2E tests;
- `Then` проверяет private database или internal method вместо outcome;
- feature-local example переписывает весь `UC-*`;
- derived `FUC-*` или `TC-*` вводит новый verdict, отсутствующий в `brief.md`.

## Outcome / Exit Contract

BDD применён корректно, когда команда имеет shared understanding через
concrete examples, canonical facts остаются у существующих owners, а required
behavior прослеживается от `UC/REQ` через `SC/NEG` и `CHK` до test evidence.

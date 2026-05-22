---
title: Feature Flow
doc_kind: governance
doc_function: canonical
purpose: "Определяет stage-based flow feature-документации с явным разделением `brief.md` (problem space), `design.md` (solution space) и `implementation-plan.md` (execution space)."
derived_from:
  - ../dna/governance.md
  - ../dna/frontmatter.md
canonical_for:
  - feature_directory_structure
  - feature_document_boundaries
  - feature_template_selection_rules
  - feature_flow_stages
  - feature_solution_gate_rules
  - feature_plan_gate_rules
  - feature_closure_rules
  - feature_support_document_rules
  - feature_c4_model_selection_rules
  - feature_identifier_taxonomy
  - solution_identifier_taxonomy
  - feature_plan_identifier_taxonomy
  - feature_traceability_rules
  - feature_decomposition_principle
  - feature_grounding_gate
status: active
audience: humans_and_agents
---
# Feature Flow

Этот документ задает порядок появления feature-артефактов. Агент должен вести feature package по стадиям и не создавать downstream-артефакты раньше, чем созрел их upstream-owner.

## Package Rules

1. Все документы одной фичи живут в `memory-bank/features/FT-XXX/`.
2. **Feature = vertical slice.** Одна фича — одна единица пользовательской ценности, пронизывающая все затронутые слои системы (UI, API, storage, infra). Горизонтальная нарезка ("все endpoints", "весь UI") допустима только для чисто инфраструктурных или рефакторинговых задач и должна быть явно обоснована через `NS-*`.
3. `brief.md` — canonical owner problem space: problem, outcome, scope, non-scope, assumptions, constraints, unresolved blocking decisions и canonical verify contract delivery-единицы.
4. `design.md` — canonical owner solution space: selected design, to-be C4 architecture model на нужных для фичи уровнях L1/L2/L3, accepted feature-local decisions, solution structure, internal flow, concrete contracts, solution-level failure modes, local rollout/backout semantics и ссылки на принятые ADR.
5. `README.md` создается вместе с `brief.md` и остается routing-слоем на всем lifecycle.
6. Lifecycle owner для `delivery_status` — только canonical `brief.md`. `design.md`, feature-level `README.md` и `implementation-plan.md` не дублируют это поле.
7. `design.md` появляется только после `Problem Ready`. Для новых и уже migrated feature packages `implementation-plan.md` появляется только после `Solution Ready`.
8. `implementation-plan.md` — derived execution-документ. В новых feature packages он не должен существовать, пока sibling `design.md` не стал `status: active`; packages, созданные до split-а, мигрируют по правилам из секции «Migration Strategy».
9. Для canonical `brief.md`, canonical `design.md`, feature-level `README.md` и `implementation-plan.md` используй wrapper-шаблоны из `memory-bank/flows/templates/feature/`: сам template-файл имеет `doc_function: template`, а frontmatter/body инстанцируемого документа живут внутри embedded template contract.
10. Смысл стабильных идентификаторов (`REQ-*`, `SOL-*`, `SD-*`, `STEP-*` и т.д.) задается в секции «Stable Identifiers» ниже.
11. Acceptance scenarios (`SC-*`) покрывают vertical slice end-to-end: от входного события до наблюдаемого результата через все затронутые слои. Тестирование отдельного слоя в изоляции допустимо как implementation detail плана, но не заменяет end-to-end acceptance.
12. **Связь с task tracker.** При создании feature package агент обязан добавить в исходную задачу или ticket ссылку на `brief.md`, а после появления downstream-документов — ссылки на `design.md` и `implementation-plan.md`.
13. Если фича является частью более крупной инициативы, `brief.md` может зависеть от PRD из `memory-bank/prd/`, но PRD не заменяет сам feature package.
14. Если фича создает новый устойчивый сценарий проекта или materially changes существующий, соответствующий `UC-*` в `memory-bank/use-cases/` должен быть создан или обновлен до closure.
15. Optional feature-support docs (`runtime-surfaces.md`, `ui-reference/README.md`, `use-cases/README.md`) допустимы для сложных фич как grounding / review / traceability aids. Они не становятся canonical owner problem space, solution space, acceptance inventory или execution sequencing.

## Шаблон `brief.md`

Новые feature packages используют один problem-space template: `memory-bank/flows/templates/feature/brief.md`. Отдельных `short.md` / `large.md` templates больше нет.

`brief.md` масштабируется содержанием:

- компактная фича заполняет минимальный набор `REQ-*`, `NS-*`, `SC-*`, `CHK-*`, `EVID-*`;
- сложная problem-space часть добавляет `MET-*`, `ASM-*`, `CON-*`, `DEC-*`, `NEG-*`, несколько acceptance scenarios, richer traceability и evidence contract;
- solution-space complexity не расширяет `brief.md`; для выбранного подхода, contracts, C4, failure modes и rollout/backout используется sibling `design.md`.

Если агенту кажется, что нужен "large brief", это означает не выбор другого template, а более полное заполнение того же `brief.md`.

## Small Feature Path

`design.md` обязателен для любого feature package. Пропускать его нельзя: иначе selected design неизбежно вернется в `brief.md` или `implementation-plan.md`.

Для small feature допускается короткая форма `design.md`:

- минимум один `SOL-*` с выбранным подходом;
- компактный `Change Surface`;
- ссылки на `REQ-*`, которые закрывает решение;
- выбор C4 model levels из L1/L2/L3 с явным `Include? = no` для уровней, которые не нужны, если фича не меняет architecture boundaries или runtime relationships;
- остальные секции (`SD-*`, `CTR-*`, `FM-*`, `RB-*`, ADR refs) добавляются только если они реально нужны.

## Optional Feature Support Docs

Support docs создаются только когда они снимают реальную неоднозначность или делают review существенно точнее. Они являются `doc_kind: feature-support` и `doc_function: reference` / `index`, если иное явно не обосновано.

| Support doc | Когда создавать | Что фиксирует | Чего не владеет |
| --- | --- | --- | --- |
| `runtime-surfaces.md` | Фича затрагивает несколько runtime entrypoints, concrete surfaces, semantic mappings, fallback/error paths или context variants | current surface inventory, semantic mapping, adjacent out-of-scope surfaces, target mapping reference, context matrix, resolution / decision table, observability notes | requirements, selected design, acceptance criteria, implementation sequence |
| `ui-reference/README.md` | Фича меняет интерфейс, authoring flow, navigation, screen states или preview / editor UX | generic interface reference: screen map, interaction states, component expectations, copy/state semantics, mockup links and UI traceability | project-specific UI framework rules, product requirements, selected architecture, implementation steps |
| `ui-reference/mockups/*.md` или другой linkable artifact | Любое interface change требует хотя бы low-fidelity mockup; default format — Markdown, но допустимы images, design-tool links или other artifacts, если они versionable / linkable | screen sketch, state examples, interaction notes | canonical acceptance inventory или final visual design system |
| `use-cases/README.md` | Сценариев много, есть distinct happy/edge/error journeys, несколько user roles или нужен review-friendly `FUC -> REQ -> CHK` mapping | derived user-facing scenarios, edge/error cases, candidate test cases, traceability back to canonical refs | canonical `SC-*`, `NEG-*`, `CHK-*`, `EVID-*` |

Support docs должны ссылаться на canonical owners и явно писать, что они не подменяют `brief.md`, `design.md` или `implementation-plan.md`. Если support doc обнаруживает изменение scope, acceptance, selected design или execution sequence, сначала обновляется соответствующий canonical owner.

## Migration Strategy

- Новые feature packages обязаны сразу следовать структуре `brief.md -> design.md -> implementation-plan.md`.
- Существующие feature packages без `brief.md` / `design.md` могут оставаться в прежнем виде, пока их не редактируют.
- Если такой package редактируется так, что меняется или дописывается problem-space content, canonical owner должен быть создан или перенесен в `brief.md`.
- Если такой package редактируется так, что меняется или дописывается solution-space content, accepted design должен быть создан или перенесен в `design.md` до следующего существенного обновления `implementation-plan.md`.
- Новые migrated packages не должны сохранять duplicate alias-файлы `feature.md` / `solution.md`; compatibility относится только к уже существующим packages.
- Миграция может происходить постепенно, package-by-package; migrated example должен оставаться доступным в `examples/`.

## Lifecycle

```mermaid
flowchart LR
    DF["Draft Feature<br/>brief.md: draft<br/>delivery_status: planned<br/>design: absent<br/>plan: absent"] --> PR["Problem Ready<br/>brief.md: active<br/>delivery_status: planned"]
    PR --> SR["Solution Ready<br/>design.md: active"]
    SR --> PL["Plan Ready<br/>implementation-plan.md: active"]
    PL --> EX["Execution<br/>delivery_status: in_progress<br/>plan: active"]
    PR --> CL["Cancelled<br/>delivery_status: cancelled<br/>plan: absent or archived"]
    SR --> CL
    PL --> CL
    EX --> DN["Done<br/>delivery_status: done<br/>plan: archived"]
    EX --> CL
```

## Transition Gates

Каждый gate — набор проверяемых предикатов. Переход допустим тогда и только тогда, когда все предикаты истинны.

### Bootstrap Feature Package

- [ ] `README.md` создан по шаблону `templates/feature/README.md`
- [ ] `brief.md` создан по шаблону `templates/feature/brief.md`
- [ ] `design.md` отсутствует
- [ ] `implementation-plan.md` отсутствует

### Draft Feature → Problem Ready

- [ ] `brief.md` → `status: active`
- [ ] секция `What` содержит ≥ 1 `REQ-*` и ≥ 1 `NS-*`
- [ ] секция `Verify` содержит ≥ 1 `SC-*`
- [ ] каждый `REQ-*` прослеживается к ≥ 1 `SC-*` через traceability matrix
- [ ] секция `Verify` содержит ≥ 1 `CHK-*` и ≥ 1 `EVID-*`
- [ ] если deliverable нельзя принять без negative/edge coverage → ≥ 1 `NEG-*`
- [ ] `brief.md` не содержит accepted solution decisions, `How`, to-be C4 architecture model, `Change Surface`, solution-level `Flow`, `CTR-*`, `FM-*`, `RB-*` или rollout/backout prose

### Problem Ready → Solution Ready

- [ ] `design.md` создан по шаблону `templates/feature/design.md`
- [ ] `design.md` → `status: active`
- [ ] `design.md` содержит ≥ 1 `SOL-*`
- [ ] `design.md` ссылается минимум на один canonical `REQ-*` из sibling `brief.md`
- [ ] `design.md` фиксирует выбор нужных C4 model levels из L1/L2/L3 для to-be архитектуры; выбранные C4 views ссылаются на `SOL-*`, `SD-*`, `CTR-*` или ADR refs
- [ ] selected design стабилизирован настолько, что downstream execution sequencing больше не конкурирует с ним за ownership
- [ ] accepted feature-local decisions перенесены в `SD-*`, а architectural / reusable / cross-feature decisions оформлены в accepted ADR
- [ ] если solution зависит от ADR, соответствующий ADR имеет `decision_status: accepted`
- [ ] для нового feature package `implementation-plan.md` отсутствует; для migrated package с уже существующим планом разрешено создать `design.md`, после чего план должен быть обновлён так, чтобы ссылаться на canonical solution refs до следующего существенного execution update

### Solution Ready → Plan Ready

- [ ] агент выполнил grounding: прошёлся по текущему состоянию системы (relevant paths, existing patterns, dependencies) и зафиксировал результат в discovery context секции `implementation-plan.md`
- [ ] `implementation-plan.md` создан по шаблону `templates/feature/implementation-plan.md`
- [ ] `implementation-plan.md` → `status: active`
- [ ] `implementation-plan.md` содержит ≥ 1 `PRE-*`, ≥ 1 `STEP-*`, ≥ 1 `CHK-*`, ≥ 1 `EVID-*`
- [ ] discovery context в `implementation-plan.md` содержит: relevant paths, local reference patterns, unresolved questions (`OQ-*`), test surfaces и execution environment
- [ ] шаги и workstreams в `implementation-plan.md` ссылаются на canonical IDs из `brief.md` и solution refs из `design.md` / ADR

### Plan Ready → Execution

- [ ] `brief.md` → `delivery_status: in_progress`
- [ ] `design.md` → `status: active`
- [ ] `implementation-plan.md` → `status: active`
- [ ] `implementation-plan.md` фиксирует test strategy: automated coverage surfaces, required local/CI suites
- [ ] каждый manual-only gap имеет причину, ручную процедуру и `AG-*` с approval ref

### Execution → Done

- [ ] все `CHK-*` из `brief.md` имеют результат pass/fail в evidence
- [ ] все `EVID-*` из `brief.md` заполнены конкретными carriers (путь к файлу, CI run, screenshot)
- [ ] delivered behavior не противоречит accepted `SOL-*` / `SD-*` / ADR refs
- [ ] automated tests для change surface добавлены или обновлены
- [ ] required test suites зелёные локально и в CI
- [ ] каждый manual-only gap явно approved человеком (approval ref в `AG-*`)
- [ ] simplify review выполнен: код минимально сложен или complexity обоснована ссылкой на `CON-*`, `FM-*`, `SD-*` или accepted ADR
- [ ] если feature добавляет новый stable flow или materially changes существующий project-level scenario, соответствующий `UC-*` создан или обновлен и зарегистрирован в `memory-bank/use-cases/README.md`
- [ ] `brief.md` → `delivery_status: done`
- [ ] `implementation-plan.md` → `status: archived`

### → Cancelled (из любой стадии после Draft Feature)

- [ ] `brief.md` → `delivery_status: cancelled`
- [ ] `implementation-plan.md` отсутствует ∨ `status: archived`

## Boundary Rules

1. `brief.md` обязан содержать секции `What` и `Verify`.
2. `brief.md` владеет только problem space: problem, outcome, scope, non-scope, assumptions, constraints, unresolved blocking decisions и canonical verify contract.
3. `brief.md` не должен содержать `How`, selected design, to-be C4 architecture model, accepted solution decisions, change surface, internal flow, concrete solution contracts, solution-level failure modes, rollout/backout semantics или execution sequencing.
4. `DEC-*` в `brief.md` означает только unresolved blocking decisions. Как только решение принято, оно переезжает в `design.md` как `SD-*` или в ADR.
5. `design.md` владеет только solution space: selected design, to-be C4 architecture model на нужных для фичи уровнях L1/L2/L3, accepted feature-local decisions, solution structure, internal flow, concrete contracts, solution-level failure modes, local rollout/backout semantics и ссылки на принятые ADR.
6. `delivery_status` остается только на `brief.md`; `design.md` и `implementation-plan.md` не дублируют lifecycle state delivery-единицы.
7. `design.md` не должен переопределять business requirements, scope, acceptance criteria, canonical checks, evidence contract, detailed current-system inventory или execution sequencing.
8. Feature-support docs не должны переопределять canonical facts. Они могут давать surface inventory, UI reference, mockups, derived use cases и review mappings только как support context.
9. Если feature зависит от ADR, canonical owner этой зависимости — `design.md`; `proposed` ADR не считается finalized design.
10. Если feature зависит от канонического use case, `brief.md` ссылается на соответствующий файл в `memory-bank/use-cases/`. Use case остается owner-ом trigger/preconditions/main flow/postconditions на уровне проекта, а `brief.md` фиксирует только slice-specific проблему и verify.
11. `implementation-plan.md` остается derived execution-документом: он ссылается на canonical IDs из `brief.md` и solution refs из `design.md` / ADR, фиксирует discovery context и test strategy для исполнения и не переопределяет scope, selected design, to-be C4 architecture model, blockers, acceptance criteria или evidence contract.
12. Если меняются scope, assumptions, constraints, acceptance criteria или evidence contract, сначала обновляется `brief.md`. Если меняются selected design, to-be C4 architecture model, local accepted decisions, contracts, failure modes или rollout/backout semantics, сначала обновляется `design.md` или ADR. Только потом обновляется downstream-план.
13. Если support doc выявляет конфликт с canonical owner, конфликт нельзя решать внутри support doc: обнови `brief.md`, `design.md`, ADR или `implementation-plan.md` по ownership.
14. Если численный target threshold относится только к одной delivery-единице, canonical owner — соответствующий `brief.md`. Поднимать такой KPI в project-level документ можно только после того, как он стал shared upstream fact для нескольких feature.
15. Хороший `implementation-plan.md` начинается с discovery context: relevant paths, local reference patterns, unresolved questions, test surfaces и execution environment должны быть зафиксированы до sequencing изменений.
16. Для рискованных, необратимых или внешне-эффективных действий `implementation-plan.md` должен явно описывать human approval gates и не скрывать их внутри prose шага.

## Test Ownership Summary

Canonical testing policy живёт в [../engineering/testing-policy.md](../engineering/testing-policy.md). Ниже — выжимка, достаточная для создания feature package без обращения к policy-документу.

1. **Canonical test cases** delivery-единицы задаются в `brief.md` через `SC-*`, feature-specific `NEG-*`, `CHK-*` и `EVID-*`.
2. `design.md` может фиксировать solution-level `CTR-*`, `FM-*` и `RB-*`, но не владеет test strategy и не подменяет canonical verify contract.
3. `implementation-plan.md` владеет только стратегией исполнения: какие suites добавить, какие gaps временно manual-only и почему.
4. **Sufficient coverage** = покрыт основной changed behavior, новые или измененные contracts из `design.md` / ADR, критичные failure modes из `FM-*` и feature-specific negative/edge scenarios, если они меняют verdict. Процент line coverage сам по себе недостаточен.
5. **Manual-only допустим** только как явное исключение (live infra, hardware, недетерминированная среда). Для каждого gap — причина, ручная процедура или `EVID-*`, owner follow-up и approval ref через `AG-*`.
6. **К Problem Ready** `brief.md` уже фиксирует test case inventory: минимум один `SC-*`, traceability к `REQ-*`. **К Solution Ready** `design.md` фиксирует delivered design, to-be C4 architecture model, contracts и local decisions. **К Done** — automated tests добавлены, обязательные suites зелёные локально и в CI.
7. **Simplify review** — отдельный проход после функциональных тестов, до closure. Цель: убедиться, что код минимально сложен. Три похожие строки лучше premature abstraction. Complexity оправдана только со ссылкой на `CON-*`, `FM-*`, `SD-*` или accepted ADR.
8. **Verification context separation** — функциональная верификация, simplify review и acceptance test — три логически отдельных прохода. Между проходами агент формулирует выводы до начала следующего. Для small features допустимо в одной сессии, но simplify review не пропускается.

## Stable Identifiers

### Feature IDs

| Prefix | Meaning | Used in |
| --- | --- | --- |
| `MET-*` | outcome-метрики | `brief.md` |
| `REQ-*` | scope и обязательные capability | `brief.md` |
| `NS-*` | non-scope | `brief.md` |
| `ASM-*` | assumptions и рабочие предпосылки | `brief.md` |
| `CON-*` | ограничения problem space | `brief.md` |
| `DEC-*` | unresolved blocking decisions | `brief.md` |
| `INV-*` | problem-level invariants | `brief.md` |
| `EC-*` | exit criteria | `brief.md` |
| `SC-*` | acceptance scenarios | `brief.md` |
| `NEG-*` | negative / edge test cases | `brief.md` |
| `CHK-*` | проверки | `brief.md`, `implementation-plan.md` |
| `EVID-*` | evidence-артефакты | `brief.md`, `implementation-plan.md` |
| `RJ-*` | rejection rules | `brief.md`, `implementation-plan.md` |

### Solution IDs

| Prefix | Meaning | Used in |
| --- | --- | --- |
| `SOL-*` | solution elements / selected design blocks | `design.md` |
| `C4-*` | to-be C4 model levels, elements или relationships | `design.md` |
| `SD-*` | accepted feature-local solution decisions | `design.md` |
| `CTR-*` | concrete solution contracts | `design.md` |
| `FM-*` | solution-level failure modes | `design.md` |
| `RB-*` | rollout / backout stages | `design.md` |

### Plan IDs

| Prefix | Meaning | Used in |
| --- | --- | --- |
| `PRE-*` | preconditions | `implementation-plan.md` |
| `OQ-*` | unresolved questions / ambiguities | `implementation-plan.md` |
| `WS-*` | workstreams | `implementation-plan.md` |
| `AG-*` | approval gates for risky actions | `implementation-plan.md` |
| `STEP-*` | атомарные шаги | `implementation-plan.md` |
| `PAR-*` | параллелизуемые блоки | `implementation-plan.md` |
| `CP-*` | checkpoints | `implementation-plan.md` |
| `ER-*` | execution risks | `implementation-plan.md` |
| `STOP-*` | stop conditions / fallback | `implementation-plan.md` |

### Support IDs

| Prefix | Meaning | Used in |
| --- | --- | --- |
| `SURF-*` | runtime surfaces / entrypoints / concrete render or processing surfaces | `runtime-surfaces.md` |
| `MAP-*` | semantic mapping rows or mapping rules | `runtime-surfaces.md` |
| `UI-*` | interface screens, states, controls or interaction elements | `ui-reference/README.md` |
| `FUC-*` | derived feature-local use cases | `use-cases/README.md` |
| `TC-*` | derived test case candidates | `use-cases/README.md`, support docs |

### Required Minimum

1. Любой canonical `brief.md` использует как минимум `REQ-*`, `NS-*`, `SC-*`, `CHK-*`, `EVID-*`.
2. Любой `brief.md` со `status: active` задает хотя бы один explicit test case через `SC-*`.
3. `brief.md` может использовать только минимальный problem-space набор для small feature или расширенный набор feature IDs по необходимости; отдельные `short.md` / `large.md` templates не используются.
4. Любой `design.md` использует как минимум один `SOL-*` и связывает его минимум с одним `REQ-*` из sibling `brief.md`.
5. Любой `design.md` фиксирует selection rationale для C4 model levels; выбранные C4 views используют `C4-*` и связываются с `SOL-*`, `SD-*`, `CTR-*` или ADR refs.
6. Любой `design.md`, где есть принятые feature-local решения, использует `SD-*`; `CTR-*`, `FM-*` и `RB-*` применяются только когда соответствующая solution-semantics действительно нужна.
7. Любой optional support doc использует только local support IDs и traceability к canonical refs; он не вводит новые canonical `REQ-*`, `SC-*`, `CHK-*` или `EVID-*`.
8. Любой `implementation-plan.md` использует как минимум `PRE-*`, `STEP-*`, `CHK-*`, `EVID-*`; при наличии ambiguity или human approval gates используются `OQ-*` и `AG-*`.

### Traceability Contract

1. Scope в `brief.md` фиксируется через `REQ-*`, non-scope через `NS-*`.
2. Verify в `brief.md` связывает `REQ-*` с test cases через `Acceptance Scenarios`, feature-specific `NEG-*`, `Traceability matrix`, `Test matrix` и `Evidence contract`.
3. `design.md` связывает `REQ-*` из `brief.md` с `SOL-*`, `C4-*`, `SD-*`, `CTR-*`, `FM-*`, `RB-*` и accepted ADR refs.
4. `implementation-plan.md` ссылается на canonical IDs из `brief.md` и solution refs из `design.md` / ADR в колонках `Implements`, `Verifies` и `Evidence IDs`.
5. Если sequencing блокируется неизвестностью, план фиксирует её как `OQ-*`, а не прячет в prose.
6. Если выполнение требует человеческого подтверждения для рискованных действий, план фиксирует это через `AG-*`.
7. Если design или to-be C4 architecture model меняется после `Solution Ready`, сначала обновляется `design.md` или ADR, затем план.

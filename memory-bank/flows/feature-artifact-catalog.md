---
title: Feature Artifact Catalog
doc_kind: governance
doc_function: reference
purpose: Каталог optional артефактов для постановки feature-задачи и описания ее решения. Читать, чтобы выбрать минимально достаточный package без пустых placeholders и duplicate ownership.
derived_from:
  - ../dna/governance.md
  - feature.md
status: active
audience: humans_and_agents
---

# Feature Artifact Catalog

Этот каталог — меню, а не checklist. Он перечисляет распространенные программно-инженерные артефакты и помогает выбрать только те, которые снимают реальную неоднозначность конкретной feature.

При bootstrap feature package обязательны только `README.md` и `brief.md`. Все остальные документы, таблицы и diagrams условны. `implementation-plan.md` появляется только перед реальным execution, а отдельный `design.md` — только когда `brief.md` фиксирует `Design required: yes`.

## Selection Rules

1. Начинай с prose, списка или компактной таблицы в canonical owner.
2. Добавляй diagram, когда связи, состояния или temporal order плохо читаются линейно.
3. Выноси материал в отдельный файл, когда у него появляется самостоятельная review boundary, несколько consumers или он делает owner-документ трудно читаемым.
4. Не создавай пустые placeholders, каталоги «на будущее» или ссылки на отсутствующие artifacts.
5. Каждый отдельный artifact индексируется из feature `README.md`; solution-space artifact также индексируется из `design.md`.
6. Reference view проецирует canonical facts и не принимает новые requirements, solution decisions или execution steps.
7. Если готового шаблона нет, используй [Extension Contract](#extension-contract), а не изобретай второго canonical owner.

## Problem And Task Artifacts

| Artifact | Question answered | Trigger | Default form / suggested path | Ownership | Template |
| --- | --- | --- | --- | --- | --- |
| Issue / ticket | Какой delivery request запустил работу? | Любая tracked feature | External tracker link | Workflow state; canonical feature facts переносятся в package | external |
| `PRD-*` | Какую продуктовую инициативу и outcome реализует набор features? | Инициатива порождает несколько delivery units или требует product-layer contract | `memory-bank/prd/PRD-XXX.md` | Product goals, initiative scope, success metrics | [PRD](templates/prd/PRD-XXX.md) |
| Project-level `UC-*` | Какой устойчивый пользовательский / операторский сценарий поддерживает система? | Scenario повторяется во времени или используется несколькими features | `memory-bank/use-cases/UC-XXX.md` | Canonical reusable scenario | [Use Case](templates/use-case/UC-XXX.md) |
| Epic package | Как координируются roadmap, risks и несколько delivery units? | Работа крупнее одной vertical feature | `memory-bank/epics/EP-XXX/` | Initiative coordination, не feature execution | [Epic](templates/epic/README.md) |
| `README.md` | Какие artifacts реально входят в feature package и в каком порядке их читать? | Любой feature package | `features/FT-XXX/README.md` | Routing only | [Feature README](templates/feature/README.md) |
| `brief.md` | Какую проблему решаем, что входит в scope и как принимаем результат? | Любой feature package | `features/FT-XXX/brief.md` | Canonical problem, requirements, acceptance and evidence contract | [Brief](templates/feature/brief.md) |
| Feature-local use cases | Какие happy, edge и error journeys удобнее review отдельно? | Много scenarios/roles или нужен `FUC -> REQ -> CHK` mapping | `use-cases/README.md` | Derived scenario projection; canonical acceptance остается в `brief.md` | [Feature Use Cases](templates/feature/support/use-cases.md) |
| Runtime surface inventory | Где behavior существует сейчас и какой context доступен? | Несколько entrypoints, mappings, fallbacks или context variants | `runtime-surfaces.md` | Current-state reference | [Runtime Surfaces](templates/feature/support/runtime-surfaces.md) |
| UI flow / mockups | Что видит пользователь и какие interface states проходит? | Меняется UI, navigation, editor/preview или interaction model | `ui-reference/README.md`, `ui-reference/mockups/*`; ссылка на `engineering/ui-design-guide/README.md` или нужный surface document | Interface reference; requirements и selected solution остаются у canonical owners; shared UI catalog не копируется в feature | [UI Reference](templates/feature/support/ui-reference.md) |
| Glossary | Что означают неоднозначные business и technical terms? | Терминология materially влияет на scope, contract или review | Compact table in owner; при росте `glossary.md` | Reference term registry with source refs | pattern only |
| Business rules / decision table | Какой outcome соответствует комбинации условий? | Много входных dimensions, precedence rules или mutually exclusive branches | Table in `brief.md` для required behavior; в `design.md` для solution policy | `brief.md` владеет required behavior; `design.md` — выбранной policy | pattern only |
| Assumptions / constraints / open decisions | На чем основана задача и что ограничивает допустимый outcome? | Есть неполная информация, external dependency или blocking choice | `ASM-*`, `CON-*`, unresolved `DEC-*` in `brief.md` | Canonical problem-space facts | [Brief](templates/feature/brief.md) |
| Examples / fixtures | Какие concrete inputs/outputs снимают неоднозначность prose? | Contract или rule легче проверить на representative examples | Synthetic example in owner document | Illustration only; normative semantics задает owner | pattern only |

## Solution And Design Artifacts

| Artifact | Question answered | Trigger | Default form / suggested path | Ownership | Template |
| --- | --- | --- | --- | --- | --- |
| `design.md` | Какое решение выбрано и почему? | `Design required: yes` | `features/FT-XXX/design.md` | Canonical feature-local solution и design-pack routing | [Design](templates/feature/design.md) |
| C4 view | Какие system, container, component или critical code boundaries и bindings затронуты? | Срабатывает C4 trigger из feature flow | Embedded Mermaid/table; при росте `diagrams/<name>-c4.md` | Reference projection `C4-*`, `SOL-*`, `SD-*`, `CTR-*` или accepted ADR; не заменяет Architecture Coverage Decision | pattern in Design |
| Component responsibility map | Как распределена ответственность между modules/services? | Новая decomposition, orchestration или ownership transfer | Table or C3 view in `design.md` | Selected responsibilities остаются `SOL-*` / `SD-*` | pattern only |
| Data-flow diagram | Откуда приходят данные, через какие connectors преобразуются и куда уходят? | Несколько sources/sinks, transformations, bindings или data owners | Embedded diagram; при росте `diagrams/<name>-data-flow.md` | Reference projection canonical contracts, direction, topology and ownership | pattern only |
| Sequence diagram | В каком порядке взаимодействуют actors/components? | Async calls, callbacks, retries, timeouts, duplicates, compensation или hand-offs | Embedded Mermaid; при росте `diagrams/<name>-sequence.md` | `SEQ-*` reference projection; новых решений не принимает | [Sequence Diagram](templates/feature/support/sequence-diagram.md) |
| State machine | Какие states/transitions допустимы и какие запрещены? | Order/payment/job/approval lifecycle или non-trivial workflow | Table/Mermaid in `design.md`; при росте `diagrams/<name>-state-machine.md` | Transition semantics trace to `SOL-*`, `CTR-*`, `INV-*`, `FM-*` | pattern only |
| Interaction contract | Каким connector связаны стороны и каковы его interaction semantics? | Detailed API/event/queue/callback/file/store/cache/auth/locking/runtime-config boundary; schema/encoding задают format, provider — party/role | Inline `CTR-*`; при самостоятельной review boundary `contracts/<name>.md` | Delegated owner explicitly listed `CTR-*`; selected solution и topology остаются в `design.md` | [Interaction Contract](templates/feature/api-contract.md) |
| Event catalog / schema | Какие events публикуются/потребляются и как versioned? | Event-driven interaction или очередь | Interaction contract variant or `contracts/<event>.md` | Delegated event `CTR-*` | Interaction Contract variant |
| Domain model | Какие entities/value objects и domain relationships нужны решению? | Меняется предметная модель или bounded-context ownership | Diagram/table in `design.md` | Feature-local model decisions; shared domain facts promoted to `domain/` | pattern only |
| Data model / ERD / dictionary | Как выглядят persistence entities, fields, indexes и relations? | Меняется schema/storage contract | Compact table/ERD in `design.md`; при росте design-pack artifact | Solution/schema facts; shared schema owner imported rather than copied | pattern only |
| Error taxonomy | Какие errors/states существуют и как consumer их интерпретирует? | API/integration или много failure outcomes | Table in contract or `design.md` | `CTR-*` wire semantics and `FM-*` solution behavior | Interaction Contract pattern |
| Failure-mode analysis | Что может сломаться и как решение ограничивает impact? | Distributed, financial, security-critical или degradation-sensitive flow | `FM-*` in `design.md`; table when richer analysis needed | Canonical solution failure semantics | Design section |
| Idempotency / concurrency model | Как обрабатываются duplicates, races, locks и ordering? | Callbacks, jobs, financial operations или parallel writers | Contract/design tables and sequence branches | `CTR-*`, `INV-*`, `FM-*`, `SD-*` | Interaction/Sequence patterns |
| Quality attributes / NFR | Какие latency, capacity, availability, consistency или recovery properties нужны? | Эти properties меняют класс допустимых решений | Constraints in `brief.md`; solution response in `design.md` | Requirement vs solution ownership сохраняется раздельно | pattern only |
| Security / threat analysis | Какие trust boundaries, threats и controls существуют? | Auth, permissions, secrets, personal/financial data или external integration | Compact section/table in `design.md`; при росте `security-analysis.md` | Feature controls; reusable security policy требует ADR/project owner | pattern only |
| Migration design | Как перейти из current state в target state без потери compatibility? | Data/schema/config migration, dual read/write или staged cutover | `migration-design.md` или compact `RB-*` section | Delegated migration facts indexed from `design.md` | pattern only |
| Compatibility matrix | Какие producer/consumer/schema versions совместимы? | Rolling deploy или independently released components | Table in contract/migration design | Delegated compatibility contract | Interaction/Migration pattern |
| Rollout / backout design | Как безопасно включить и откатить изменение? | Risky release, feature flag, migration или operational switch | `RB-*` in `design.md`; separate artifact only when large | Canonical solution rollout semantics | Design section |
| Observability contract | Какие logs, metrics, traces и alerts показывают состояние solution? | Background, async или production-critical behavior | Table in `design.md`; при росте `observability-contract.md` | Solution observability semantics; project policy imported | pattern only |
| ADR | Почему выбрано architectural/reusable/cross-feature решение? | Decision выходит за feature-local boundary | `memory-bank/adr/ADR-XXX-*.md` | Canonical architecture decision | [ADR](templates/adr/ADR-XXX.md) |

## Execution, Verification And Review Artifacts

| Artifact | Question answered | Trigger | Default form / suggested path | Ownership | Template |
| --- | --- | --- | --- | --- | --- |
| `implementation-plan.md` | В каком порядке реализовать accepted problem/solution contract? | Feature действительно переходит к execution | `features/FT-XXX/implementation-plan.md` | Workstreams, steps, commands, checkpoints and stop conditions | [Implementation Plan](templates/feature/implementation-plan.md) |
| Test matrix / strategy | Какие requirements, contracts и failures чем проверяются? | Change surface требует нескольких suites/types или manual gap | Canonical checks in `brief.md`; execution strategy in plan | Acceptance remains in `brief.md`; execution coverage in plan | Brief/Plan sections |
| Evidence artifact | Чем доказан конкретный check? | Evidence удобнее хранить отдельно от CI link/path/screenshot | Linkable carrier, optionally `evidence.md` | Results only; не меняет expected behavior | pattern only |
| Review report | Какие findings найдены и как закрыты? | Formal review/reconciliation materially useful | `<kind>-review-report.md` or external review link | Findings/status only; canonical owners update first | pattern only |

## Package Profiles

Профили показывают типичный минимальный набор и не вводят обязательность optional artifacts:

- **Local change:** `README.md` + `brief.md`; перед execution добавляется `implementation-plan.md`.
- **Designed change:** local change + `design.md`.
- **Scenario-heavy change:** local или designed change + `use-cases/README.md`.
- **Integration / contract change:** designed change + optional `contracts/<name>.md`; sequence diagram только при значимой temporal semantics.
- **Interface change:** local или designed change + `ui-reference/README.md` и linkable mockups.
- **Architecture-significant change:** designed change + accepted ADR и минимально достаточный C4 artifact.

## Lifecycle Usage

1. **Bootstrap:** создай только `README.md` и `brief.md`.
2. **Problem analysis:** выбери только нужные problem/support companions и реши, нужен ли `design.md`; если есть selected companions или material omissions, зафиксируй их в optional Artifact Routing Decision из `brief.md`.
3. **Problem Ready:** если `Design required: no`, не создавай design-pack и не позволяй плану принимать solution decisions.
4. **Solution analysis:** если design required, начни с `design.md`; добавляй contract, diagram, migration/security/observability artifacts только по trigger.
5. **Routing:** после создания каждого artifact добавь аннотированную ссылку в feature `README.md`; solution artifact также добавь в `design.md#design-pack`.
6. **Plan Ready:** `implementation-plan.md` потребляет canonical IDs из готовых owners и не изобретает новые requirements/contracts/decisions.
7. **Change control:** сначала обновляй canonical owner, затем dependent views и план.

## Extension Contract

Feature-local artifact, которого нет в каталоге, допустим только когда он уменьшает реальную неоднозначность. Такой artifact обязан:

1. использовать lowercase kebab-case path;
2. иметь governed frontmatter и явные `purpose`, `derived_from`, `status`;
3. быть проиндексирован из `README.md`, а для solution-space artifact — также из `design.md`;
4. явно фиксировать `Role`, `Owns` и `Must not define`;
5. ссылаться на canonical IDs вместо копирования facts;
6. не создавать второго active owner для problem space, selected solution или execution sequencing.

## When To Add A Governed Template

Не создавай template только потому, что artifact появился один раз. Новый governed template оправдан, когда artifact регулярно повторяется, имеет устойчивую структуру, несет заметный риск неправильного описания и требует repeatable traceability.

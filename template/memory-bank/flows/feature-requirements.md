---
title: Feature Requirements, Identifiers And Traceability
doc_kind: governance
doc_function: canonical
purpose: "Определяет taxonomy feature requirements, registry стабильных идентификаторов, applicability decisions и двустороннюю traceability до verification и delivered surfaces."
derived_from:
  - ../dna/governance.md
  - ../dna/frontmatter.md
  - ../engineering/validation-profiles.md
canonical_for:
  - feature_identifier_taxonomy
  - feature_requirement_taxonomy
  - feature_requirement_traceability_rules
  - solution_identifier_taxonomy
  - feature_plan_identifier_taxonomy
must_not_define:
  - feature_flow_stages
  - feature_solution_decisions
  - feature_execution_steps
status: active
audience: humans_and_agents
---

# Feature Requirements, Identifiers And Traceability

Этот документ — canonical requirements policy для [`Feature Flow`](feature.md).
Он определяет виды требований, их ownership, applicability и доказуемость.
Здесь же находится полный registry стабильных идентификаторов, связывающий
problem, solution и execution layers. Lifecycle и transition gates остаются в
`feature.md`; templates инстанцируют оба контракта.

## Requirement Taxonomy And Traceability

`brief.md` owns the feature requirement inventory. A requirement is an externally needed outcome or condition—not a design choice, task, test, or evidence. Keep `REQ-*` as its stable identifier and record a mandatory class field; do not introduce a parallel `FR-*`/`NFR-*` namespace. `MET-*` is a goal or observed metric, `CON-*` a boundary, `EC-*` an acceptance verdict, and `CHK-*`/`EVID-*` proof. Each may link to a `REQ-*`, but none replaces it.

For every baseline class, the brief records `applicable`, `not-applicable` with a rationale, or `covered-upstream` with its canonical reference. Functional is always applicable; the decision row itself is mandatory for all other classes. The validation profile changes verification depth, not classification. A triggered class gets a `REQ-*`; a shared fact stays in its upstream product, domain, policy, or regulatory owner.

| Class | Level | Canonical owner / artifact | Applicability trigger | Measurement, verification, and evidence | Example / anti-example |
| --- | --- | --- | --- | --- | --- |
| stakeholder / product | stakeholder/product | shared product, PRD, or use case; otherwise `brief.md` | delivery outcome is stakeholder-specific | validate scenario outcome with its evidence | operator completes intake / internal implementation preference |
| functional | feature/system | `brief.md` `REQ-*` | always | scenario, check, and evidence | system shall accept a submission / use PostgreSQL |
| performance | feature/system/component | `brief.md` `REQ-*` | latency, throughput, capacity, or resource target | numeric threshold, repeatable measurement, report | p95 under 200 ms / “fast” |
| quality attribute | feature/system/component | `brief.md` `REQ-*` | availability, reliability, consistency, recovery, usability, maintainability, etc. | measurable criterion, method, and result carrier | recover within 15 min / add retries |
| interface | feature/system boundary | `brief.md` `REQ-*`, realized by `CTR-*` | API, CLI, event, UI, or external-system boundary | contract/interaction check and sample or CI result | webhook returns signed payload / handler class name |
| data | feature/system boundary | `brief.md` `REQ-*`, realized by `CTR-*` | schema, format, retention, integrity, or migration boundary | data/contract check and migration evidence | retain audit fields 365 days / add a column |
| security | feature/system boundary | `brief.md` `REQ-*`, controls in design | trust, auth, secret, sensitive-data, or threat trigger | analysis plus control check/evidence | only owner may export / choose OAuth library |
| safety | feature/system boundary | `brief.md` `REQ-*`, controls in design | harm, hazardous operation, or safety-critical failure trigger | hazard/failure check and result | stop device on sensor fault / standard exception wording |
| regulatory / compliance | stakeholder/product or system | policy/regulation upstream or `brief.md` | applicable obligation | procedure/check evidence | retain consent record / team style preference |
| operational | feature/system/component | `brief.md` `REQ-*`, solution/runbook in design | operator, observability, support, backup, or recovery trigger | runbook/check evidence | alert includes correlation ID / refactor package name |
| compatibility | feature/system boundary | `brief.md` `REQ-*`, mechanism in design | versioned consumer, migration, or legacy behavior trigger | compatibility matrix/contract check | v1 client remains supported / latest SDK only |
| deployment / rollout | feature/system operation | `brief.md` `REQ-*`, `RB-*` in design | staged release, flag, migration, rollback, or infra delivery trigger | rollout/backout evidence | rollback within one deploy unit / create a new module |
| constraint | any applicable level | `CON-*` in `brief.md` | imposed budget, technology, policy, date, or boundary | link testable proof where possible | must use approved region / prefer a pattern |
| verification / acceptance | feature delivery | `SC-*`, `EC-*`, `CHK-*`, `EVID-*` in `brief.md` | required for each applicable requirement; is proof, not a requirement class | method, acceptance verdict, and evidence carrier | SC proves export result / “add a test” as a requirement |

Every applicable `REQ-*` records: class, normative measurable statement (a threshold where meaningful), source/rationale, priority, accountable owner, verification method, acceptance/check/evidence links, and realization link. Goals are `MET-*`; assumptions are `ASM-*`; constraints are `CON-*`; selected decisions, invariants, and contracts are `SD-*`/`INV-*`/`CTR-*`; acceptance is `EC-*`; evidence records a result only.

Required chain: `upstream source → REQ-* → EC/SC → selected solution/contract or design-not-required decision → exact repository path + symbol/config section → STEP-* → CHK-* → EVID-* → review/CI result`. Each changed implementation, test, or configuration surface maps back to a `REQ-*` or an explicit supporting/necessary rationale. Paths are repository-relative and name a symbol, heading, or configuration key; a glob or module-only label is not an exact target. Lifecycle artifact review checks the chain in both directions: no orphan requirement, dangling link, duplicate owner, accepted design fact without realization target, or unexplained changed surface.

Lint and doctor remain structural: they may validate ID format, uniqueness, resolvable links, and required template sections, but do not infer semantic applicability or a validation profile. Those judgments remain explicit brief evidence and lifecycle review, avoiding a domain-specific requirements engine.

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
| `ALT-*` | considered alternatives | `design.md` |
| `TRD-*` | trade-offs | `design.md` |
| `C4-*` | C4 applicability decision, model levels, elements или relationships | `design.md` |
| `SD-*` | accepted feature-local solution decisions | `design.md` |
| `INV-*` | solution invariants | `design.md` или delegated constituent |
| `CTR-*` | concrete solution contracts | `design.md` или delegated constituent |
| `FM-*` | solution-level failure modes | `design.md` или delegated constituent |
| `RB-*` | rollout / backout stages | `design.md` или delegated constituent |

### Plan IDs

| Prefix | Meaning | Used in |
| --- | --- | --- |
| `GRND-*` | grounding evidence о текущем repository state, existing patterns и test surfaces | `implementation-plan.md` |
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
| `SEQ-*` | sequence branches, temporal rules or interaction paths | `diagrams/<name>-sequence.md`, embedded sequence views |

## Worked Traceability Examples

| Feature kind | Requirement → realization chain |
| --- | --- |
| User-facing | `REQ-01` accessibility: “keyboard focus is visible on every dialog control” → `SC-01` keyboard journey → `CTR-01` UI contract → `web/dialog.tsx#Dialog` → `STEP-01` → `CHK-01` browser test → `EVID-01` CI result. |
| Contract / integration | `REQ-02` interface: “webhook is signed with HMAC-SHA256” → `EC-02` → `CTR-02` → `services/webhook.go#SignPayload` → `STEP-02` → `CHK-02` contract test → `EVID-02` CI result. |
| Infrastructure / operations | `REQ-03` operational: “rollback completes within one deploy unit” → `SC-03` rollback drill → `RB-01` → `infra/deploy.yaml#rollback` → `STEP-03` → `CHK-03` staging drill → `EVID-03` run record. |

## Migration And Compatibility

Existing feature packages remain valid: their mnemonic IDs retain their current meanings and are not silently reclassified or renamed. New active packages use the applicability matrix and fields above. An existing package adds only the applicable decision rows and trace links when it next materially changes; unknown legacy coverage is recorded as a gap or follow-up, never fabricated.

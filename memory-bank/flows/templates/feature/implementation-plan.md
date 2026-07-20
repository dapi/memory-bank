---
title: FT-XXX Feature Template - Implementation Plan
doc_kind: feature
doc_function: template
purpose: Governed wrapper-шаблон плана имплементации. Фиксирует, как инстанцировать execution-документ без переопределения canonical problem или solution facts и без смешения wrapper с целевым `implementation-plan.md`.
derived_from:
  - ../../feature.md
  - ../../feature-artifact-catalog.md
  - ../../../dna/frontmatter.md
  - ../../../engineering/testing-policy.md
status: active
audience: humans_and_agents
template_for: feature
template_target_path: ../../../features/FT-XXX/implementation-plan.md
---

# План имплементации

Этот файл описывает wrapper-template. Инстанцируемый `implementation-plan.md` живет ниже как embedded contract и копируется без wrapper frontmatter и history.

## Wrapper Notes

Требования, blocker-state и критерии приемки задаются в sibling `brief.md`. Если `brief.md` фиксирует `Design required: yes`, selected design, accepted local decisions и solution-level contracts задаются в sibling `design.md` или ADR. Этот документ определяет только sequencing работ и checkpoints выполнения.
В создаваемом feature package sibling `brief.md` всегда инстанцируется из canonical template в `memory-bank/flows/templates/feature/`; `design.md` инстанцируется только когда required.

Создавай этот документ только после того, как upstream owners готовы: sibling `brief.md` имеет `status: active`, а required sibling `design.md` переведен в `status: active`. Пока план только формируется, сам `implementation-plan.md` может оставаться в `status: draft`; до перехода feature в `delivery_status: in_progress` план должен стать `status: active`.

Когда feature переходит в `delivery_status: done` или `delivery_status: cancelled`, `implementation-plan.md` архивируется, если он больше не используется как рабочий execution-документ.

Документ должен быть исполнимым без дополнительного толкования. Если шаг нельзя связать с canonical IDs, существующими solution refs, артефактом, проверкой или явной ручной процедурой, шаг описан недостаточно.
План должен быть заземлен в текущем состоянии репозитория: сначала зафиксируй релевантные модули, локальные паттерны, открытые вопросы и execution environment, и только после этого расписывай sequencing изменений.
План обязан явно зафиксировать, какие automated tests будут добавлены или обновлены по change surface, какие suites обязаны быть зелёными локально и в CI, а какие gaps временно остаются manual-only с justification и approval ref. Он исполняет validation profile из sibling `brief.md`, но не выбирает и не дублирует profile. Для designed feature план также показывает refinement каждого применимого `SOL-*`, `C4-*`, `SD-*`, `CTR-*`, `INV-*`, `FM-*`, `RB-*` и accepted ADR ref в realization target, steps, checks и evidence, не принимая новых solution decisions.

Для ссылок внутри плана используй стабильные идентификаторы по taxonomy из [../../feature.md#stable-identifiers](../../feature.md#stable-identifiers).

Если неизвестность меняет scope, acceptance criteria или evidence contract, она сначала поднимается upstream в sibling `brief.md`. Если неизвестность меняет selected design, architecture coverage, C4 architecture model, accepted local decisions, contracts, invariants, failure modes или rollout/backout semantics, она сначала поднимается в required sibling `design.md`, delegated contract или ADR и только после этого фигурирует в плане.

## Instantiated Frontmatter

```yaml
title: "FT-XXX: Implementation Plan"
doc_kind: feature
doc_function: derived
purpose: "Execution-план реализации FT-XXX. Фиксирует discovery context, шаги, риски и test strategy без переопределения canonical problem и solution фактов."
derived_from:
  - brief.md
  # Required only when brief.md says "Design required: yes":
  # - design.md
  # Optional support refs:
  # - runtime-surfaces.md
  # - ui-reference/README.md
  # - use-cases/README.md
status: draft
audience: humans_and_agents
must_not_define:
  - ft_xxx_scope
  - ft_xxx_selected_design
  - ft_xxx_acceptance_criteria
  - ft_xxx_blocker_state
  - ft_xxx_validation_profile
```

## Instantiated Body

```markdown
# План имплементации

## Цель текущего плана

Какой delivery outcome должен дать этот план с учетом `brief.md` и, если есть, already accepted solution.

## Grounding / Support References

Какие upstream canonical и support docs используются как execution baseline. Support docs не переопределяют canonical facts: при конфликте обнови owner-документ до продолжения.

| Document | Role in this plan | Facts reused | Conflict action |
| --- | --- | --- | --- |
| `brief.md` | canonical problem / validation profile / verify owner | profile decision, `REQ-*`, `SC-*`, `CHK-*`, `EVID-*` | Update `brief.md` first |
| `design.md` / `none` | conditional solution owner | `SOL-*`, `C4-*`, `SD-*`, `CTR-*`, `INV-*`, `FM-*`, `RB-*` | Update `design.md` or ADR first; if design is absent, promote new design facts before planning |
| `runtime-surfaces.md` / `none` | optional grounding | `SURF-*`, `MAP-*`, context matrix | Promote changed design facts to `design.md` if design is required |
| `ui-reference/README.md` / `none` | optional interface reference | `UI-*`, mockups, states | Promote changed requirements to `brief.md` or design facts to `design.md` if required |
| `use-cases/README.md` / `none` | optional scenario companion | `FUC-*`, `TC-*` candidates | Keep canonical acceptance in `brief.md` |
| `contracts/<name>.md` / `none` | optional delegated contract owner | Explicit `CTR-*`, compatibility, errors, idempotency | Update delegated contract and `design.md` routing before planning against changed semantics |
| `diagrams/<name>-sequence.md` / `none` | optional temporal reference | `SEQ-*`, ordering, async and failure branches | Update canonical `design.md` / contract first if the sequence reveals changed solution facts |

## Current State / Reference Points

Какие существующие файлы, модули, команды или документы агент обязан изучить до начала изменений. Этот раздел фиксирует grounding в текущем состоянии репозитория и локальные паттерны, которые нельзя игнорировать.

| Path / module | Current role | Why relevant | Reuse / mirror |
| --- | --- | --- | --- |
| `path/to/module` | Что уже делает этот артефакт | Почему без него нельзя планировать корректно | Какой паттерн, helper, command или contract нужно повторить |

## Test Strategy

Какие test surfaces должны быть обновлены по мере реализации. Сошлись на validation profile из `brief.md` и покажи, как каждая применимая обязанность его minimum contract закрывается tests, suites, evidence, approvals и rollout/backout checkpoints. Этот раздел не переопределяет profile decision или canonical test cases из `brief.md`.

| Test surface | Canonical refs | Existing coverage | Planned automated coverage | Required local suites / commands | Required CI suites / jobs | Manual-only gap / justification | Manual-only approval ref |
| --- | --- | --- | --- | --- | --- | --- | --- |
| `path/or/behavior` | `REQ-01`, `SC-01`, `NEG-01`, `CHK-01`, `SOL-01 если design существует` | Что покрыто сейчас | Какой suite, test type или deterministic check обязаны добавить или обновить | Какие команды или suites обязаны быть зелёными локально | Какие jobs или suites обязаны быть зелёными в CI | Что пока остается manual-only и почему | `AG-01` / review link / `none` |

## Open Questions / Ambiguities

Какие неизвестности ещё не сняты после discovery. Если вопрос меняет upstream semantics, его нельзя молча разрешать в шаге исполнения.
Если после discovery unresolved questions отсутствуют, укажи `none` вместо таблицы; не создавай фиктивный `OQ-*`.

| Open Question ID | Question | Why unresolved | Blocks | Default action / escalation owner |
| --- | --- | --- | --- | --- |
| `OQ-01` | Что именно неизвестно | Почему это ещё не доказано | `STEP-02` / `WS-1` / whole plan | Что делаем по умолчанию и кто принимает решение при эскалации |

## Environment Contract

Какой execution environment считается допустимым для плана: setup, test commands, env vars, permissions, mocks, внешние зависимости и другие operational assumptions.

| Area | Contract | Used by | Failure symptom |
| --- | --- | --- | --- |
| setup | Какая подготовка среды обязательна | `STEP-01`, `STEP-02` | По какому симптому понятно, что среда невалидна |
| test | Какая команда или процедура считается эталонной для verify на этом этапе | `CHK-01` | Что считается недостоверным verify |
| access / network / secrets | Какие доступы, домены, ключи или sandbox assumptions нужны | `STEP-03` | Когда работа должна остановиться и уйти на эскалацию |

## Preconditions

Что должно быть готово до старта работ: данные, доступы, ADR, окружение, договоренности. Каждая строка ссылается на canonical ref и не пересказывает его смысл своими словами.

| Precondition ID | Canonical ref | Required state | Used by steps | Blocks start |
| --- | --- | --- | --- | --- |
| `PRE-01` | `CON-01` / `DEC-01` / `SD-01 если design существует` / ADR path / design-not-required decision | Какой state upstream считается допустимым для старта | `STEP-01`, `STEP-02` | yes / no |

## Design Realization Mapping

Для designed feature покажи, где реализуется каждый применимый `SOL-*`, `C4-*`, `SD-*`, `CTR-*`, `INV-*`, `FM-*`, `RB-*` и accepted ADR ref. Каждый применимый ref должен встречаться минимум в одной строке. Объединяй в строке только refs с одним canonical owner, общим realization target и одной verification chain; иначе раздели их. Строка связывает уже принятое решение с execution и не вводит новые solution facts или decisions. Если mapping обнаруживает gap или требует изменить semantics, сначала обнови canonical owner и только затем этот план. Для feature с `Design required: no` укажи `not applicable` и ссылку на decision из `brief.md`.

| Canonical solution refs | Owner | Realization target | Steps | Checks | Evidence |
| --- | --- | --- | --- | --- | --- |
| `SOL-01`, `SD-01` | `design.md` | Module or service | `STEP-01` | `CHK-01` | `EVID-01` |
| `CTR-01` | `contracts/<name>.md` | Interface or interaction boundary | `STEP-02` | `CHK-02` | `EVID-02` |
| `C4-01`, `INV-01`, `FM-01` | `design.md` | Runtime topology | `STEP-03` | `CHK-03` | `EVID-03` |
| `RB-01` | `design.md` | Migration, config or operational surface | `STEP-04` | `CHK-04` | `EVID-04` |
| `../../adr/ADR-XXX.md` | `../../adr/ADR-XXX.md` (`accepted`) | Decision realization target | `STEP-05` | `CHK-05` | `EVID-05` |

## Workstreams

Разбей работу на независимые потоки с явным результатом каждого.

| Workstream | Implements | Result | Owner | Dependencies |
| --- | --- | --- | --- | --- |
| `WS-1` | `REQ-01`, применимые `SOL-*`, `C4-*`, `SD-*`, `CTR-*`, `INV-*`, `FM-*`, `RB-*` и accepted ADR refs | Что должно появиться | human / agent / either | Что блокирует старт или завершение |

## Approval Gates

Какие действия нельзя выполнять без явного человеческого подтверждения. Используй этот раздел для рискованных, необратимых, дорогих или внешне-эффективных операций.

| Approval Gate ID | Trigger | Applies to | Why approval is required | Approver / evidence |
| --- | --- | --- | --- | --- |
| `AG-01` | Какой шаг или симптом запрашивает approval | `STEP-03` / `WS-2` | Почему нельзя продолжать автономно | Кто подтверждает и чем это фиксируется |

## Порядок работ

Опиши выполнение как атомарные шаги. Каждый шаг должен быть достаточно маленьким, чтобы его можно было проверить и при необходимости откатить или остановить без расползания change surface.

| Step ID | Actor | Implements | Goal | Touchpoints | Artifact | Verifies | Evidence IDs | Check command / procedure | Blocked by | Needs approval | Escalate if |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| `STEP-01` | human / agent / either | `REQ-01`, применимые `SOL-*`, `C4-*`, `SD-*`, `CTR-*`, `INV-*`, `FM-*`, `RB-*` и accepted ADR refs | Что делаем на этом шаге | Какие файлы, сервисы или данные трогаем | Что должно появиться после шага | `CHK-01` | `EVID-01` | Как подтверждаем завершение | `PRE-01`, `OQ-01` | `AG-01` / `none` | Когда нельзя продолжать без эскалации |

## Parallelizable Work

Какие шаги или workstreams можно выполнять параллельно без конфликта по change surface.

- `PAR-01` Что может идти параллельно.
- `PAR-02` Что нельзя распараллеливать из-за общего write-surface.

## Checkpoints

Какие промежуточные точки должны быть пройдены до rollout или handoff.

| Checkpoint ID | Refs | Condition | Evidence IDs |
| --- | --- | --- | --- |
| `CP-01` | `STEP-01`, `CHK-01`, `SOL-01 если design существует` | Какой промежуточный state должен быть доказан | `EVID-01` |

## Execution Risks

Какие практические риски могут сорвать сроки или потребовать пересборки плана.

| Risk ID | Risk | Impact | Mitigation | Trigger |
| --- | --- | --- | --- | --- |
| `ER-01` | Что может пойти не так | Что это ломает | Что делаем заранее | По какому сигналу активируется mitigation |

## Stop Conditions / Fallback

Когда план должен остановиться или откатиться в безопасное состояние.

| Stop ID | Related refs | Trigger | Immediate action | Safe fallback state |
| --- | --- | --- | --- | --- |
| `STOP-01` | `DEC-01`, `RJ-01`, `SD-01 если design существует` | По какому симптому останавливаемся | Что делаем сразу | До какого состояния откатываемся или замораживаем работу |

## Plan-local Evidence

Какие evidence artifacts принадлежат самому execution plan и не являются canonical evidence contract из `brief.md`.

| Evidence ID | Artifact | Producer | Path contract | Reused by checkpoints |
| --- | --- | --- | --- | --- |
| `EVID-09` | Например simplify-review verdict, discovery note или manual approval note | implementer / reviewer / human approver | Где лежит или чем фиксируется | `CP-01` |

## Готово для приемки

Какие условия должны выполниться, чтобы считать план исчерпанным и перейти к финальной приемке по секции `Verify` в sibling `brief.md`.

- Все workstreams завершены или явно остановлены через `STOP-*`.
- Все checkpoints имеют evidence.
- Required local suites зелёные, а CI не противоречит local verify.
- Manual-only gaps закрыты через approved `AG-*` или остаются blockers для `delivery_status: done`.
- Support docs, если они есть, не расходятся с canonical `brief.md`, existing `design.md`, ADR и этим планом.
- Финальная приемка идёт по `brief.md` `Verify`, а не по этому checklist.
```

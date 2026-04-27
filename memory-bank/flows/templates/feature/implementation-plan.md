---
title: FT-XXX Feature Template - Implementation Plan
doc_kind: feature
doc_function: template
purpose: Governed wrapper-шаблон плана имплементации. Фиксирует, как инстанцировать execution-документ без переопределения canonical problem или solution facts и без смешения wrapper с целевым `implementation-plan.md`.
derived_from:
  - ../../feature-flow.md
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

Требования, blocker-state и критерии приемки задаются в sibling `feature.md`. Selected design, accepted local decisions и solution-level contracts задаются в sibling `solution.md` или ADR. Optional support docs (`runtime-surfaces.md`, `ui-reference/README.md`, `use-cases/README.md`) можно использовать как grounding/reference, но они не становятся execution owner. Этот документ определяет только sequencing работ и checkpoints выполнения.
В создаваемом feature package sibling `feature.md` и `solution.md` должны быть инстанцированы из canonical templates в `memory-bank/flows/templates/feature/`.

Создавай этот документ только после того, как sibling `solution.md` переведен в `status: active`. Пока план только формируется, сам `implementation-plan.md` может оставаться в `status: draft`; до перехода feature в `delivery_status: in_progress` план должен стать `status: active`.

Когда feature переходит в `delivery_status: done` или `delivery_status: cancelled`, `implementation-plan.md` архивируется, если он больше не используется как рабочий execution-документ.

Документ должен быть исполнимым без дополнительного толкования. Если шаг нельзя связать с canonical IDs, solution refs, артефактом, проверкой или явной ручной процедурой, шаг описан недостаточно.
План должен быть заземлен в текущем состоянии репозитория: сначала зафиксируй релевантные модули, локальные паттерны, открытые вопросы и execution environment, и только после этого расписывай sequencing изменений.
План обязан явно зафиксировать, какие automated tests будут добавлены или обновлены по change surface, какие suites обязаны быть зелёными локально и в CI, а какие gaps временно остаются manual-only с justification и approval ref.

Для ссылок внутри плана используй стабильные идентификаторы по taxonomy из [../../feature-flow.md#stable-identifiers](../../feature-flow.md#stable-identifiers).

Если неизвестность меняет scope, acceptance criteria или evidence contract, она сначала поднимается upstream в sibling `feature.md`. Если неизвестность меняет selected design, to-be C4 architecture model, accepted local decisions, contracts или rollout/backout semantics, она сначала поднимается в sibling `solution.md` или ADR и только после этого фигурирует в плане.

## Instantiated Frontmatter

```yaml
title: "FT-XXX: Implementation Plan"
doc_kind: feature
doc_function: derived
purpose: "Execution-план реализации FT-XXX. Фиксирует discovery context, шаги, риски и test strategy без переопределения canonical problem и solution фактов."
derived_from:
  - feature.md
  - solution.md
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
```

## Instantiated Body

```markdown
# План имплементации

## Цель текущего плана

Какой delivery outcome должен дать этот план с учетом already accepted solution.

## Grounding / Support References

Какие upstream canonical и support docs используются как execution baseline. Support docs не переопределяют canonical facts: при конфликте обнови owner-документ до продолжения.

| Document | Role in this plan | Facts reused | Conflict action |
| --- | --- | --- | --- |
| `feature.md` | canonical problem / verify owner | `REQ-*`, `SC-*`, `CHK-*`, `EVID-*` | Update `feature.md` first |
| `solution.md` | canonical solution owner | `SOL-*`, `C4-*`, `SD-*`, `CTR-*`, `FM-*`, `RB-*` | Update `solution.md` or ADR first |
| `runtime-surfaces.md` / `none` | optional grounding | `SURF-*`, `MAP-*`, context matrix | Promote changed design facts to `solution.md` |
| `ui-reference/README.md` / `none` | optional interface reference | `UI-*`, mockups, states | Promote changed requirements to `feature.md` or design facts to `solution.md` |
| `use-cases/README.md` / `none` | optional scenario companion | `FUC-*`, `TC-*` candidates | Keep canonical acceptance in `feature.md` |

## Current State / Reference Points

Какие существующие файлы, модули, команды или документы агент обязан изучить до начала изменений. Этот раздел фиксирует grounding в текущем состоянии репозитория и локальные паттерны, которые нельзя игнорировать.

| Path / module | Current role | Why relevant | Reuse / mirror |
| --- | --- | --- | --- |
| `path/to/module` | Что уже делает этот артефакт | Почему без него нельзя планировать корректно | Какой паттерн, helper, command или contract нужно повторить |

## Test Strategy

Какие test surfaces должны быть обновлены по мере реализации. Этот раздел фиксирует expected automated coverage, required local/CI gates и manual-only exceptions для change surface, не переопределяя canonical test cases из `feature.md`.

| Test surface | Canonical refs | Existing coverage | Planned automated coverage | Required local suites / commands | Required CI suites / jobs | Manual-only gap / justification | Manual-only approval ref |
| --- | --- | --- | --- | --- | --- | --- | --- |
| `path/or/behavior` | `REQ-01`, `SOL-01`, `SC-01`, `NEG-01`, `CHK-01` | Что покрыто сейчас | Какой suite, test type или deterministic check обязаны добавить или обновить | Какие команды или suites обязаны быть зелёными локально | Какие jobs или suites обязаны быть зелёными в CI | Что пока остается manual-only и почему | `AG-01` / review link / `none` |

## Open Questions / Ambiguities

Какие неизвестности ещё не сняты после discovery. Если вопрос меняет upstream semantics, его нельзя молча разрешать в шаге исполнения.

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
| `PRE-01` | `CON-01` / `DEC-01` / `SD-01` / ADR path | Какой state upstream считается допустимым для старта | `STEP-01`, `STEP-02` | yes / no |

## Workstreams

Разбей работу на независимые потоки с явным результатом каждого.

| Workstream | Implements | Result | Owner | Dependencies |
| --- | --- | --- | --- | --- |
| `WS-1` | `REQ-01`, `SOL-01`, `CTR-01` | Что должно появиться | human / agent / either | Что блокирует старт или завершение |

## Approval Gates

Какие действия нельзя выполнять без явного человеческого подтверждения. Используй этот раздел для рискованных, необратимых, дорогих или внешне-эффективных операций.

| Approval Gate ID | Trigger | Applies to | Why approval is required | Approver / evidence |
| --- | --- | --- | --- | --- |
| `AG-01` | Какой шаг или симптом запрашивает approval | `STEP-03` / `WS-2` | Почему нельзя продолжать автономно | Кто подтверждает и чем это фиксируется |

## Порядок работ

Опиши выполнение как атомарные шаги. Каждый шаг должен быть достаточно маленьким, чтобы его можно было проверить и при необходимости откатить или остановить без расползания change surface.

| Step ID | Actor | Implements | Goal | Touchpoints | Artifact | Verifies | Evidence IDs | Check command / procedure | Blocked by | Needs approval | Escalate if |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| `STEP-01` | human / agent / either | `REQ-01`, `SOL-01`, `CTR-01` | Что делаем на этом шаге | Какие файлы, сервисы или данные трогаем | Что должно появиться после шага | `CHK-01` | `EVID-01` | Как подтверждаем завершение | `PRE-01`, `OQ-01` | `AG-01` / `none` | Когда нельзя продолжать без эскалации |

## Parallelizable Work

Какие шаги или workstreams можно выполнять параллельно без конфликта по change surface.

- `PAR-01` Что может идти параллельно.
- `PAR-02` Что нельзя распараллеливать из-за общего write-surface.

## Checkpoints

Какие промежуточные точки должны быть пройдены до rollout или handoff.

| Checkpoint ID | Refs | Condition | Evidence IDs |
| --- | --- | --- | --- |
| `CP-01` | `STEP-01`, `SOL-01`, `CHK-01` | Какой промежуточный state должен быть доказан | `EVID-01` |

## Execution Risks

Какие практические риски могут сорвать сроки или потребовать пересборки плана.

| Risk ID | Risk | Impact | Mitigation | Trigger |
| --- | --- | --- | --- | --- |
| `ER-01` | Что может пойти не так | Что это ломает | Что делаем заранее | По какому сигналу активируется mitigation |

## Stop Conditions / Fallback

Когда план должен остановиться или откатиться в безопасное состояние.

| Stop ID | Related refs | Trigger | Immediate action | Safe fallback state |
| --- | --- | --- | --- | --- |
| `STOP-01` | `DEC-01`, `SD-01`, `RJ-01` | По какому симптому останавливаемся | Что делаем сразу | До какого состояния откатываемся или замораживаем работу |

## Plan-local Evidence

Какие evidence artifacts принадлежат самому execution plan и не являются canonical evidence contract из `feature.md`.

| Evidence ID | Artifact | Producer | Path contract | Reused by checkpoints |
| --- | --- | --- | --- | --- |
| `EVID-09` | Например simplify-review verdict, discovery note или manual approval note | implementer / reviewer / human approver | Где лежит или чем фиксируется | `CP-01` |

## Готово для приемки

Какие условия должны выполниться, чтобы считать план исчерпанным и перейти к финальной приемке по секции `Verify` в sibling `feature.md`.

- Все workstreams завершены или явно остановлены через `STOP-*`.
- Все checkpoints имеют evidence.
- Required local suites зелёные, а CI не противоречит local verify.
- Manual-only gaps закрыты через approved `AG-*` или остаются blockers для `delivery_status: done`.
- Support docs, если они есть, не расходятся с canonical `feature.md`, `solution.md`, ADR и этим планом.
- Финальная приемка идёт по `feature.md` `Verify`, а не по этому checklist.
```

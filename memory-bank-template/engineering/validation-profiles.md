---
title: Validation Profiles
doc_kind: engineering
doc_function: canonical
purpose: Определяет независимую от delivery flow глубину validation, её risk triggers, minimum evidence contract и ownership решения.
derived_from:
  - ../dna/governance.md
  - autonomy-boundaries.md
  - ../ops/release.md
canonical_for:
  - validation_profile_taxonomy
  - validation_profile_selection_rules
  - validation_profile_escalation_rules
  - validation_profile_minimum_contracts
  - validation_profile_decision_ownership
status: active
audience: humans_and_agents
---

# Validation Profiles

Delivery flow и validation profile отвечают на разные вопросы:

- **flow** организует lifecycle, owner-документы и handoff;
- **validation profile** задаёт минимальную глубину проверок, evidence, approvals и rollout/backout discipline.

Сначала выбери flow по [`../flows/routing.md`](../flows/routing.md), затем внутри его entry/problem gate выбери ровно один profile. Profile не меняет состав owner-документов и не является конкурирующим flow.

## Taxonomy

| Profile | Когда применять |
| --- | --- |
| `documentation` | Меняется только документация или другой non-runtime artifact; executable behavior, contracts, production config и release path не меняются. |
| `low-risk` | Локальное executable change следует известному паттерну, имеет малый blast radius и не активирует triggers ниже. |
| `standard` | Default для executable change, которое не доказано как `low-risk` и не активирует более сильный профиль. |
| `high-risk` | Ошибка может затронуть trust, деньги, persistent data, concurrency semantics или несколько систем; требуются explicit approval и независимая проверка. |
| `release-deployment` | Основной change surface — production config, build/release artifact, deployment или rollback path без отдельного `high-risk` trigger. |

Это не количественный risk score. `documentation < low-risk < standard`; `high-risk` и `release-deployment` — усиленные специализированные профили. Если применимы оба, выбери `high-risk` и добавь все release/deployment obligations из соответствующей строки minimum contract.

## Selection Triggers

Начинай со `standard`, затем обоснуй снижение или повышение:

- `documentation` допустим только при отсутствии executable, contract, config и release impact;
- `low-risk` допустим, когда change локален, rollback очевиден, affected test surface известен и нет triggers из таблицы;
- новый или изменённый public API, event, schema либо file format требует как минимум `standard`; breaking/compatibility-sensitive вариант, migration или cross-system consumer повышает до `high-risk`;
- security/auth/trust boundary, financial calculation, persistent data/migration, concurrency/locking/idempotency semantics повышают до `high-risk`;
- новая cross-system integration или material change её protocol, failure semantics, data ownership либо authentication повышает до `high-risk`;
- production config, build/release artifact, deployment или rollback path повышает до `release-deployment`; если одновременно затронут любой `high-risk` trigger, применяется composition rule выше.

Не понижай профиль из-за маленького diff, короткого срока или отсутствия готового test environment.

## Escalation And Downgrade Rules

1. Profile выбирается до реализации и пересматривается при расширении change surface или появлении нового trigger.
2. Более сильный обнаруженный trigger немедленно повышает profile и обновляет canonical decision owner до продолжения работы.
3. Снижение с автоматически сработавшего `high-risk` или `release-deployment` допустимо только с конкретной rationale и human approval reference в canonical owner. Молчаливое исключение запрещено.
4. Отсутствие возможности выполнить обязательную проверку создаёт blocker или approved manual-only gap по [`testing-policy.md`](testing-policy.md), но само по себе не снижает profile.
5. Profile задаёт floor. Project-specific testing policy, incident controls, regulatory rules или reviewer могут требовать больше.

## Minimum Validation And Evidence Contract

| Profile | Required automated surfaces | Local suites | CI gates | Manual evidence | Approval gates | Rollout / backout | Independent review / convergence |
| --- | --- | --- | --- | --- | --- | --- | --- |
| `documentation` | Link, schema/frontmatter, example или docs build checks, применимые к changed docs | Targeted documentation lint/build | Все required documentation jobs | Semantic read-through; render evidence, если layout влияет на результат | Обычный review; отдельный approval только по project policy | Не требуется; если меняется published release path, переклассифицировать | Обычный review достаточен |
| `low-risk` | Targeted regression для changed behavior; существующие nearest tests | Targeted affected suite и repository lint/typecheck, если применимы | Все required jobs для change | Только для непокрываемой automation части с явной процедурой | Обычный review; manual-only gap требует указанного approver | Понятный локальный revert; staged rollout не обязателен | Simplify/convergence pass исполнителя и обычный review |
| `standard` | Changed behavior, ближайший regression path, изменённые contracts/integration boundaries и material negative cases | Все affected unit/integration/contract suites | Полный required CI set | Acceptance evidence и оформленные manual-only gaps | Approval для manual-only critical gap и внешне-эффективных действий | Rollback path для runtime change; rollout checks, если delivery не атомарна | Отдельный final convergence pass и reviewer, независимый от автора change |
| `high-risk` | Все affected unit/integration/contract/e2e surfaces; critical failure modes; migration/recovery rehearsal или deterministic substitute | Полный релевантный набор, включая security/data/concurrency checks; невозможное явно блокирует или получает approval | Все required CI плюс доступные specialized gates | Evidence по каждому critical path, failure/recovery case и rehearsal | Human approval профиля, manual-only gaps и risk-bearing execution step | Явные staged rollout, observability signals, stop conditions и проверенный backout/recovery plan | Independent reviewer по затронутому risk domain и финальный convergence pass обязательны |
| `release-deployment` | Build/package/config validation, deploy/rollback automation и smoke/health checks | Release artifact/config checks и staging rehearsal, где доступно | Required release/deployment jobs | Artifact identity, staging/smoke results и production signals | Human approval перед production или live-data action | Явные rollout units, stop signals, rollback owner и fastest safe rollback | Independent review release plan/config и post-deploy convergence обязательны |

Конкретные frameworks, команды, suites, CI job names и evidence paths не принадлежат taxonomy: их задают project-specific [`testing-policy.md`](testing-policy.md), execution plan или routing record выбранного flow.

## Canonical Decision Owner By Flow

Profile decision записывается ровно один раз; downstream artifacts ссылаются на него и не выбирают profile заново.

| Flow | Canonical owner | Правило |
| --- | --- | --- |
| Small Change | issue/task routing record; draft PR только если tracker нельзя обновить | Record содержит profile, triggers/rationale и approval ref при downgrade. |
| Feature | `memory-bank/features/FT-XXX/brief.md` | `implementation-plan.md` реализует contract через suites/checkpoints, но не дублирует решение. |
| Bug Fix | bug report или связанная delivery task; draft PR только как fallback | Reproduction, regression plan и evidence исполняют выбранный profile. |
| Refactoring | исходная task; draft PR только как fallback | Profile учитывает blast radius и critical behavior, которое нужно сохранить. |
| Incident / PIR | Не назначается containment/PIR record | Permanent remediation и prevention items получают profile после отдельного Task Routing. Incident safety gates продолжают действовать независимо. |
| Epic | Не назначается epic целиком | Profile выбирается отдельно в canonical owner каждой delivery feature/subissue. |
| Use Case / Human Routing | Не применим до выбора delivery flow | Эти records задают сценарий или решение маршрутизации, а не delivery validation. |

Минимальный decision record:

```text
Validation profile: documentation | low-risk | standard | high-risk | release-deployment
Triggers / rationale: <почему этот floor достаточен; какие triggers проверены>
Downgrade approval: <human approval ref или none>
```

## Examples

| Path | Flow | Profile decision | Minimum consequence |
| --- | --- | --- | --- |
| Исправить локальный UI label по существующему i18n pattern без изменения contract или runtime control flow | Small Change | `low-risk`: локальный surface, известных triggers нет | Targeted UI/i18n check, required CI, semantic read-through и обычный review. |
| Добавить локальное пользовательское поведение без contract, data или security changes | Feature | `standard`: executable behavior, оснований для снижения нет | Regression + acceptance coverage, affected suites, full required CI и independent review. |
| Изменить payment calculation с сохранением внешнего API | Feature | `high-risk`: financial calculation trigger | Boundary/failure coverage, explicit human approval, independent domain review, rollout signals и backout plan. |

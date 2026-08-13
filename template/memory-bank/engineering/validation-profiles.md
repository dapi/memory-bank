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
| `high-risk` | Текущий run непосредственно выполняет рискованное действие над production/live state: изменяет или удаляет production data, production access/security state, совершает реальную финансовую или другую необратимую внешнюю операцию. Подготовка решения автономна; перед risk-bearing execution step требуются explicit approval и отдельная проверка неавтором mutation. |
| `release-deployment` | Основной change surface — production config, build/release artifact, deployment или rollback path без отдельного `high-risk` trigger. |

Это не количественный risk score. `documentation < low-risk < standard`; `high-risk` и `release-deployment` — усиленные специализированные профили. Если применимы оба, выбери `high-risk` и добавь все release/deployment obligations из соответствующей строки minimum contract.

## Selection Triggers

Начинай со `standard`, затем обоснуй снижение или повышение:

- `documentation` допустим только при отсутствии executable, contract, config и release impact;
- `low-risk` допустим, когда change локален, rollback очевиден, affected test surface известен и нет triggers из таблицы;
- новый или изменённый public API, event, schema, file format, security/auth boundary, financial calculation, persistent-data model, migration plan, concurrency/locking/idempotency semantics или cross-system integration требует как минимум `standard`; эти code/design triggers сами по себе не повышают profile до `high-risk`;
- `high-risk` выбирай только когда текущий run должен непосредственно выполнить risk-bearing действие в production/live environment: изменить, удалить, backfill или repair production data; изменить production access/security state; провести реальную финансовую операцию; либо вызвать другую необратимую external effect. Ожидаемая будущая поставка code change не является таким действием;
- production config, build/release artifact, deployment или rollback path повышает до `release-deployment`; если в том же run выполняется `high-risk` действие, выбери `high-risk` и добавь все release/deployment obligations из соответствующей строки.

Не понижай профиль из-за маленького diff, короткого срока или отсутствия готового test environment.

## Escalation And Downgrade Rules

1. Profile выбирается до реализации и пересматривается при расширении change surface или появлении нового trigger.
2. Более сильный обнаруженный trigger немедленно повышает profile и обновляет canonical decision owner до продолжения работы.
3. Снижение с автоматически сработавшего `high-risk` или `release-deployment` допустимо только с конкретной rationale и human approval reference в canonical owner. Молчаливое исключение запрещено.
4. Отсутствие возможности выполнить обязательную проверку создаёт blocker или approved manual-only gap по [`testing-policy.md`](testing-policy.md), но само по себе не снижает profile.
5. Profile задаёт floor. Project-specific testing policy, incident controls, regulatory rules или reviewer могут требовать больше.

Выбор более сильного profile, design, implementation, rehearsal, staging,
validation и подготовка rollback выполняются автономно. Они не требуют human
approval только из-за сложности или риска будущего изменения. Approval
проверяется непосредственно перед тем execution step, который пересекает
canonical Human Gate из [`autonomy-boundaries.md`](autonomy-boundaries.md).
FPF Decision Protocol может выбрать подход и усилить controls, но не может
отменить обязательный approval или понизить profile ниже сработавшего trigger.

## Minimum Validation And Evidence Contract

`Обычный review` не требует отдельного неавторского reviewer: это convergence
pass исполнителя и review, предусмотренный project PR process, если он есть.
`Separate non-authoring review` выполняет отдельный actor, не создававший и не
исправлявший проверяемые mutations; таким actor может быть агент, если среда
гарантирует отсутствие mutation worktree, git/PR и других mutable systems.
Human approval — отдельный gate для risk-bearing action и не заменяется review.

| Profile | Required automated surfaces | Local suites | CI gates | Manual evidence | Approval gates | Rollout / backout | Separate review / convergence |
| --- | --- | --- | --- | --- | --- | --- | --- |
| `documentation` | Link, schema/frontmatter, example или docs build checks, применимые к changed docs | Targeted documentation lint/build | Все required documentation jobs | Semantic read-through; render evidence, если layout влияет на результат | Обычный review; отдельный approval только по project policy | Не требуется; если меняется published release path, переклассифицировать | Обычный review достаточен |
| `low-risk` | Targeted regression для changed behavior; существующие nearest tests | Targeted affected suite и repository lint/typecheck, если применимы | Все required jobs для change | Только для непокрываемой automation части с явной процедурой | Обычный review; manual-only gap требует указанного approver | Понятный локальный revert; staged rollout не обязателен | Simplify/convergence pass исполнителя и обычный review |
| `standard` | Changed behavior, ближайший regression path, изменённые contracts/integration boundaries и material negative cases | Все affected unit/integration/contract suites | Полный required CI set | Acceptance evidence и оформленные manual-only gaps | Approval для manual-only critical gap и внешне-эффективных действий | Rollback path для runtime change; rollout checks, если delivery не атомарна | Final convergence pass исполнителя и обычный review |
| `high-risk` | Все surfaces, необходимые для безопасного direct production/live action; critical failure modes; recovery rehearsal или deterministic substitute | Полный релевантный набор для данного действия; невозможное явно блокирует или получает approval | Все required CI плюс доступные specialized gates | Evidence по действию, critical path, failure/recovery case и rehearsal | Human approval для downgrade, manual-only gaps и самого risk-bearing execution step; выбор профиля и подготовка автономны | Явные staged rollout, observability signals, stop conditions и проверенный backout/recovery plan | Separate non-authoring actor проверяет затронутый production-risk domain; финальный convergence pass обязателен |
| `release-deployment` | Build/package/config validation, deploy/rollback automation и smoke/health checks | Release artifact/config checks и staging rehearsal, где доступно | Required release/deployment jobs | Artifact identity, staging/smoke results и production signals | Human approval перед production или live-data action | Явные rollout units, stop signals, rollback owner и fastest safe rollback | Separate review release plan/config и post-deploy convergence обязательны |

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
| Изменить payment calculation с сохранением внешнего API | Feature | `standard`: code semantics сами по себе не включают direct production action | Regression + acceptance coverage, affected suites, full required CI, convergence pass и обычный review. |
| Выполнить production backfill, который меняет live customer balances | Feature | `high-risk`: direct risk-bearing production-data action | Recovery rehearsal, explicit human approval, separate non-authoring domain review, rollout signals и backout plan. |

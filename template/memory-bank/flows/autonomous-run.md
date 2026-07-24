---
title: Autonomous Run Protocol
doc_kind: process
doc_function: canonical
purpose: "Определяет generic contract автономного запуска delivery-задачи: eligibility, isolated workspace, run state, evidence и human handoff."
derived_from:
  - ../dna/governance.md
  - routing.md
  - feature.md
  - ../engineering/validation-profiles.md
  - ../engineering/autonomy-boundaries.md
  - templates/process/session-handoff.md
canonical_for:
  - autonomous_run_protocol
  - autonomous_run_dispatch_gates
  - autonomous_run_workspace_contract
  - autonomous_run_state_contract
  - autonomous_run_evidence_packet
status: active
audience: humans_and_agents
---

# Autonomous Run Protocol

Этот protocol связывает task tracker, выбранный delivery flow, isolated workspace, агентский run и human handoff. Он является execution overlay: не создаёт новый route, не заменяет lifecycle выбранного flow и не владеет product, domain или solution facts.

## Scope And Boundaries

Protocol применим, когда project-specific runner может самостоятельно запускать coding agent для ограниченной delivery-unit. Tracker, agent, VCS provider, CI, облако и формат status surface являются adapter/configuration details проекта и не входят в этот template contract.

Runner владеет только scheduling state: dispatch, claim, concurrency, retry, reconciliation и lifecycle workspace. Агент достигает objective в workspace, использует доступные ему инструменты и подготавливает delivery evidence. Canonical owners выбранного flow продолжают владеть requirements, design, validation-profile decision, execution plan и delivery status.

## Dispatch Gates

До claim runner или агент обязан подтвердить все условия:

1. Task имеет stable tracker identifier, понятный intent и ссылку на canonical owner выбранного flow.
2. Task прошёл [Task Routing](routing.md); автономный run не выбирает route заново.
3. Entry gates выбранного flow выполнены. Для Feature Flow implementation run требует active `brief.md` и `Problem Ready`; если `Design required: yes`, также active `design.md` и `Solution Ready`; во всех случаях active `implementation-plan.md` и пройденный `Plan Ready` с required grounding, approvals и clean artifact-review verdict текущей active revision до перехода в `Execution`.
4. Canonical owner зафиксировал применимый validation profile по [Validation Profiles](../engineering/validation-profiles.md).
5. Task не blocked, eligible по project-specific adapter rules и не имеет другого active claim.
6. Required human approvals уже существуют, если validation profile или project policy требуют их до запуска.

Не dispatch-ить автономно и передать в Human Routing либо остановить run, если intent, scope или acceptance недостаточны; task требует product/architecture decision; существует uncontrolled security, compliance, production/live-data или irreversible external-action risk; либо необходимого approval нет. Конкретные границы разрешённых действий принадлежат [Autonomy Boundaries](../engineering/autonomy-boundaries.md).

## Isolated Workspace Contract

Каждая delivery-unit получает детерминированный isolated workspace или worktree, названный по уникальному tracker identifier. Одна delivery-unit не может иметь более одного active run без явной project-specific policy на такую параллельность.

Workspace lifecycle:

1. **Prepare** — создать или переиспользовать workspace; выполнить безопасный project-specific bootstrap/sync hook.
2. **Run** — запускать агента только с workspace как working directory; не использовать shared mutable checkout для agent changes.
3. **Preserve** — сохранять workspace после normal exit и retry, чтобы continuation имела контекст и незакоммиченные diagnostic artifacts.
4. **Clean** — удалять workspace только по явно определённой terminal policy, после нужного evidence/handoff или при безопасном cleanup terminal task.

Hooks, branch naming, repository population и retention period задаются downstream в engineering/ops documentation. Reused workspace нельзя destructively reset без документированной policy и проверки, что не теряется нужная работа.

## Run State And Reconciliation

Runner — единственный owner scheduling state. Минимальные состояния:

| State | Meaning | Allowed transition |
| --- | --- | --- |
| `unclaimed` | Eligible task ещё не зарезервирована. | `claimed`, `released` |
| `claimed` | Task зарезервирована; duplicate dispatch запрещён. | `running`, `retry_queued`, `released` |
| `running` | Агент выполняет objective в assigned workspace. | `retry_queued`, `handoff`, `released` |
| `retry_queued` | Run остановлен, continuation/retry запланирован. | `running`, `released` |
| `handoff` | Агент подготовил evidence и передал task в workflow-defined review/closure state. | `released` |
| `released` | Claim снят; task больше не управляется данным run. | `unclaimed` только после новой eligibility check |

Требования к scheduler:

- claim и dispatch идемпотентны; одновременно возможен только один active claim на tracker identifier;
- concurrency ограничена глобально и, при необходимости, project-specific classes of work;
- transient failure, timeout или stall создают bounded retry с backoff и фиксируют reason;
- перед retry и на регулярной reconciliation runner перепроверяет task: terminal, non-active, blocked или no-longer-eligible task отменяет/освобождает run;
- restart recovery строится по tracker state и workspace state; отсутствие durable scheduler database не должно создавать duplicate dispatch;
- normal process exit не равен `done`: task остаётся active, пока tracker/workflow и canonical delivery owner не достигли допустимого handoff или closure state.

## Agent Objective And Safety Posture

Runner передаёт агенту task context, ссылку на canonical owner, workspace path, required validation contract и проектные ограничения. Агент получает objective, а не жёстко заданную последовательность внутренних переходов: он может выполнить investigation, implementation, checks и разрешённую коммуникацию, пока сохраняет scope и gates.

Credentials и capabilities предоставляются по least-privilege principle. Секреты не наследуются агентскому process без доказанной необходимости; provider-native write actions должны использовать scoped credentials или mediated tools. Любое действие, попадающее в supervision/escalation boundary, требует соответствующего human gate и не становится разрешённым только из-за того, что run уже начат.

## Evidence Packet And Handoff

Перед `handoff` агент формирует или обновляет evidence packet в tracker, PR и/или canonical delivery artifact. Packet содержит:

- tracker task и ссылку на canonical delivery owner;
- branch, change set или PR reference;
- выбранный validation profile и результаты required local/CI checks;
- acceptance evidence, включая required manual proof или approved manual-only gaps;
- known failures, residual risks, rollback/rollout facts, если применимы;
- состояние task и следующий owner: review, approval, deployment, closure или blocked;
- если работа должна продолжиться в другой сессии, ссылку на run-specific `memory-bank/processes/PROCESS-XXX-session-handoff.md`, созданный или обновлённый по [Session Handoff template](templates/process/session-handoff.md) и зарегистрированный в `memory-bank/processes/README.md`; ссылка только на wrapper-template не является handoff evidence.

Успешный exit agent process без этого пакета означает только завершение attempt; это не delivery completion. Evidence не дублирует canonical requirements, design decisions или lifecycle status: она ссылается на их owners.

## Configuration And Observability

Downstream project должен version-control repository-owned configuration для tracker adapter, eligibility, concurrency, workspace hooks, retry/timeouts, safety posture и evidence destinations. Изменения конфигурации валидируются до новых dispatches; invalid update сохраняет last known good runtime behavior и создаёт operator-visible error.

Минимальная observability: structured record для claim, state transition, workspace, agent session/attempt, retry reason, cancellation, validation result и handoff. Dashboard является optional surface; logs и evidence должны позволять человеку восстановить состояние без него.

## Outcome Contract

### Observable Outcome

Каждый автономный run либо передан в явный workflow-defined handoff state с evidence packet, либо освобождён с диагностируемой причиной. Runner не оставляет скрытый active claim или неатрибутированный workspace.

### Required Evidence

- task ↔ canonical owner ↔ workspace/run ↔ evidence packet traceable;
- state transitions и retry/cancellation reasons доступны оператору;
- required validation и approvals представлены либо как evidence, либо как explicit blocker;
- terminal cleanup и release policy подтверждают, что continuation не потеряет требуемый context.

### Handoff

Human reviewer принимает delivery evidence по правилам выбранного flow и project PR/release process. Если новая неопределённость меняет scope, route, risk profile или approval requirement, run останавливается и task возвращается к соответствующему upstream gate.

---
title: "EP-141: Decisions"
doc_kind: epic
doc_function: decision_log
purpose: "EP-141: Decisions"
derived_from:
  - charter.md
status: active
audience: humans_and_agents
---

# EP-141: Decisions

## DL-01: Связанные worktree и PR
Date: 2026-09-06. Status: Resolved.
Authority: пользователь поручил реализацию issue 141 в отдельной ветке/worktree и PR.
Facts: payload и installer имеют разные canonical repositories.
Decision: изменения ведутся в двух worktree; PR связываются, template не объявляется
готовым к использованию до доступности supporting CLI. Canonical checkout остаются на main.

## DL-02: Использовать текущий CLI command contract
Date: 2026-09-06. Status: Resolved.
Facts: CLI 2.3.0 использует `pull` для обновления шаблона, `update` для самого бинарника.
Issue использует `update` в смысле обновления payload.
Decision: реализовать требования issue в `pull`; не возвращать старую CLI семантику.
Документация и проверки должны явно различать эти операции.

## DL-03: Различать передачу epic slice и execution фичи
Date: 2026-09-07. Status: Resolved.
Facts: Epic Flow «Roadmap Ready → Execution» требует создания linked FT package;
Feature Flow отдельно требует Plan Ready перед реализацией. Первый review потребовал
обновить epic stage после создания FT, второй смешал её с feature execution.
Decision: сохраняется epic execution и явно поясняется feature planned / Plan Ready pending.
Authority: действующие lifecycle owners; это уточнение статуса, не пропуск feature gate.

## DL-04: Разделить bridge и полный component design
Date: 2026-09-07. Status: Resolved.
Facts: после пяти review–fix итераций полного пакета остались material findings; runtime
реализация ещё не началась. Source-format bridge имеет отдельный наблюдаемый outcome и
не зависит от выбранной schema adoption. Повторять тот же большой review scope неэффективно.
Authority: поручение реализовать #141; изменение sequencing не уменьшает accepted scope.
Decision: W1 проходит отдельный локальный CLI design/review и tests. Общий ADR остаётся
proposed, full design — draft; full implementation plan убран до Solution Ready.
Bridge requirements/plan/evidence принадлежат CLI docs/source-format-bridge.md. W2/W3
остаются в epic и не объявляются выполненными. Human gate не нужен: intent не меняется.


## DL-05 — Re-evaluate the protocol after five review iterations

The adoption hypothesis remains explicit, single-owner and version-pinned. Review showed
that an informal serialization description cannot support a reproducible migration approval.
The revised design uses one normative wire owner and compact canonical registry bytes,
restricts the first classifier to the pinned f1f04de semantics, rejects context-changing moves,
and distinguishes successful rollback from recovery-required failures. These are deliberate
v1 boundaries, not permission to relax the issue's integrity/compatibility requirements.
The full implementation plan remains absent until this revised Solution Ready candidate
converges. CLI execution planning resumes from the revised wire contract; no component
capability is advertised by the already reviewed bridge.

## DL-06 — One normative contract after another exhausted design review budget

Five revised-design review iterations still found inconsistencies between duplicate prose
and serialized rules. The problem/accepted scope remains issue 141; reducing acceptance is
not an option. The bounded local parser/selection probe passed its finite matrix and exposed
no representability blocker, but it is not delivered runtime or a replacement for review.
Decision: consolidate behavioral and serialized semantics under component-wire-format.md;
components.md becomes navigation only. Explicitly resolve type metadata, dependency evolution,
write-intent uniqueness and legacy link policy at that owner. Update acceptance to test the
full promised finding/path invariants. Re-review this changed owner before execution; the
CLI execution plan remains separately reviewed and awaits Solution Ready. No human gate is
needed because neither product intent nor authorized actions change.

W1 is delivered as CLI PR 63 at 3b434fd93678c36447d10d4f308a39ce5d74b040: required CI,
actual-binary source fixtures, canonical canary, functional review and full simplify review
are clean. The PR is ready for review; release and merge remain outside this task.


## DL-07 — Separate Git identity from exact filesystem observations

The five consolidated-contract reviews exposed a wrong simplifying assumption: Git executable
mode is sufficient for source identity, but cannot bind actual downstream permission changes.
The accepted migration guarantee is unchanged. Re-evaluated choice: retain Git-mode fields in
the existing ownership lock, and add actual permission bits to approval/recovery observations.
Constrain directory deletion to checked, handle-relative rmdir instead of extending approval
to arbitrary recursive trees. ACL/owner/xattr and special-mode preservation are outside v1;
unsupported special file bits reject before planning. This narrows the filesystem mechanism
to a testable contract without weakening the accepted byte/permission and no-descendant-loss
requirements. Re-review the corrected observation model before component delivery; the CLI
execution plan already owns exact observations and the shared transaction engine.

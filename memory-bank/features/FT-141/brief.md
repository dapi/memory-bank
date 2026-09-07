---
title: "FT-141: Компонентная документация и flow adoption"
doc_kind: feature
doc_function: canonical
purpose: "FT-141: Компонентная документация и flow adoption"
derived_from:
  - ../../epics/EP-141/charter.md
  - ../../product/context.md
  - ../../use-cases/UC-001-adopt-documentation-and-flows.md
  - ../../flows/feature.md
status: active
audience: humans_and_agents
delivery_status: planned
---

# FT-141: Компонентная документация и flow adoption

## What

### Problem and outcome
Пользователь хочет начать с проектной документации и подключить AI-процессы позже.
[EP-141](../../epics/EP-141/charter.md) задаёт общий scope; этот delivery package связывает
проверяемый пользовательский путь template с отдельной реализацией installer в CLI.

### Requirements

Все требования P1, accountable owner — Danil; source — [issue 141](https://github.com/dapi/memory-bank/issues/141).
Проверка каждого требования автоматизированная, outcome устанавливается соответствующим SC/CHK/EVID.

| ID | Class | Normative requirement | Acceptance/check/evidence |
| --- | --- | --- | --- |
| REQ-01 | functional | Устанавливать core/docs/full и явно выбранные adapters; обычный pull сохраняет выбор. | SC-01, CHK-01, EVID-01 |
| REQ-02 | data | Сохранять авторские документы и ownership при pull и docs → full. | SC-02, CHK-02, EVID-02 |
| REQ-03 | interface | Создавать базовые и явно подключённые документы через документированный CLI contract. | SC-03, CHK-03, EVID-03 |
| REQ-04 | quality attribute | Обнаруживать adoption drift, восстанавливать транзакцию при успешном rollback и явно блокировать продолжение при recovery_required. | SC-04, CHK-04, EVID-04 |
| REQ-05 | compatibility | Сохранять pinned validation bundles и прежний pass/fail legacy-документов при opt-in миграции. | SC-05, CHK-05, EVID-05 |
| REQ-06 | operational | Отказывать несовместимым source/capabilities и неявной legacy-миграции до мутаций. | SC-06, CHK-06, EVID-06 |
| REQ-07 | stakeholder / product | Документация без Flows остаётся самостоятельной по ссылкам, metadata и обязательным инструкциям. | SC-07, CHK-07, EVID-07 |
| REQ-08 | deployment / rollout | Сохранять работоспособный legacy путь до доступности supporting CLI. | SC-08, CHK-08, EVID-08 |
| REQ-09 | security | Ограничивать manifest и document operation paths безопасными repository-relative regular paths. | SC-09, CHK-09, EVID-09 |

### Requirement applicability

| Class | Decision | Rationale |
| --- | --- | --- |
| stakeholder / product | applicable | REQ-07 |
| functional | applicable | REQ-01 |
| performance | not-applicable | Нет нового SLA; локальный CLI, объём документов определяет объём работы |
| quality attribute | applicable | REQ-04, atomic recovery и deterministic no-op |
| interface | applicable | REQ-03 |
| data | applicable | REQ-02 |
| security | applicable | REQ-09: manifest контролирует пути записи; traversal, symlinks и Git metadata должны отклоняться до мутаций |
| safety | not-applicable | Нет физических hazardous operations |
| regulatory / compliance | not-applicable | Нет изменения внешних обязательств |
| operational | applicable | REQ-06 |
| compatibility | applicable | REQ-05 |
| deployment / rollout | applicable | REQ-08; меняется путь поставки CLI/payload |
| constraint | applicable | CON-01: owning repositories и worktrees; CON-02: без live migration/merge/release в этой задаче |
| verification / acceptance | applicable | SC/CHK/EVID ниже |

### Non-scope

NS-01: uninstall/downgrade и неподдерживаемые document transitions.
NS-02: универсальная plugin system, несколько flow contracts на документ и внешняя audit authority.
NS-03: публикация release, merge PR, изменение пользовательских downstream-установок.

### Validation Profile Decision

Validation profile: release-deployment. Triggers: compatibility rollout и installation entrypoint.
Обязательны local/CI contract tests, binary integration fixtures, rollback и независимое review.
Live production execution отсутствует; отдельные approvals для него не запрашиваются.

### Design Requirement Decision

Design required: yes. Меняются CLI, file format, installation state и migration contracts.
Unresolved blocking decisions for Plan Ready: ADR-002 acceptance after clean decision review, then Solution Ready. No unresolved decision blocks starting the current design work.

## Verify

### Acceptance scenarios and traceability

| ID / requirement | Given → When → Then | Check | Evidence |
| --- | --- | --- | --- |
| SC-01 / REQ-01 | Пустой repository → init каждого состава → присутствуют ровно выбранные компоненты. | CHK-01 | EVID-01: test output и CI run |
| SC-02 / REQ-02 | Заполненный ADR в docs → pull/full → байты и user ownership прежние, gates не добавлены. | CHK-02 | EVID-02: test output и CI run |
| SC-03 / REQ-03 | Базовый brief и созданный через flow brief → validate → применяются разные явно выбранные требования. | CHK-03 | EVID-03: test output и CI run |
| SC-04 / REQ-04 | Удалён marker или registry при прежнем lock → conflict; искусственный сбой операции → исходные байты восстановлены. | CHK-04 | EVID-04: test output и CI run |
| SC-05 / REQ-05 | Валидный и невалидный legacy brief → миграция → pass/fail и точный multiset finding (ID, code, rule ID, subject, multiplicity) сохраняются; новый base brief не наследует gates. | CHK-05 | EVID-05: test output и CI run |
| SC-06 / REQ-06 | Bridge binary получает неизвестный source или component source → nonzero, downstream tree не изменён. | CHK-06 | EVID-06: test output и CI run |
| SC-07 / REQ-07 | core/docs → lint/doctor/semantic audit → нет требований к отсутствующим Flows и runner tools. | CHK-07 | EVID-07: test output и CI run |
| SC-08 / REQ-08 | Новый entrypoint получает pre-bridge CLI → остановка до installer; закреплённый legacy источник остаётся доступным. | CHK-08 | EVID-08: test output и CI run |
| SC-09 / REQ-09 | Manifest и каждый document command с traversal/absolute/Git-metadata path, symlink или case alias → отказ до мутаций; внешний sentinel и downstream неизменны. | CHK-09 | EVID-09: negative fixture output и CI run |
| SC-10 / REQ-03/04 | Adopted document → transition with required evidence → new pinned contract, same ID, appended history; a selector transition adds exactly one exclusion and record. | CHK-10 | EVID-10: positive transition fixtures |
| SC-11 / REQ-04 | Adopted document → move within context_root → same ID and gates, new path and history; exact retry is a no-op. | CHK-11 | EVID-11: positive move and retry fixtures |
| SC-12 / REQ-04/05 | Unchanged migration preview → apply with exact digest → success. Omitted/wrong digest or separately changed source, old lock, resolution, observed bytes/Git modes/POSIX permissions (including 0600→0644), directory existence/permissions/topology, selection, write intent or registry bytes → rejection before writes. | CHK-12 | EVID-12: digest approval input-class matrix |

### Negative cases

NEG-01 / CHK-04: marker/registry/id/type/path tampering, duplicate adoption и missing target → conflict.
NEG-02 / CHK-05: изменён bundle или его базовая зависимость под прежним ID → conflict; старый bundle отсутствует → отказ.
NEG-03 / CHK-06: неизвестный manifest/schema/component/source, downgrade, opt-in отсутствует → отказ до мутаций.
NEG-04 / CHK-04: ошибка staged mutation или изменившийся lock → rollback/отказ; отдельный сбой самого rollback → recovery_required, сохранённый staging и отказ следующих component mutations.
NEG-06 / CHK-09: malicious manifest paths и create/adopt/transition/move arguments (--path/--to): traversal, absolute paths, Git metadata, symlinks and case aliases → отказ без внешних или внутренних записей.
NEG-05 / CHK-05: legacy selector без исключения при per-document adoption → conflict; resolution map неоднозначна → отказ.

NEG-07 / CHK-10: old or new transition bundle requires evidence and --evidence is omitted → refusal; sufficient explicit references are retained in history.
NEG-08 / CHK-12: invalid write-intent action/existence/digest-kind combination → refusal before hashing or mutation.

NEG-09 / CHK-05: added, removed, substituted or multiplicity-changed legacy finding → migration rejects without writes.
NEG-10 / CHK-01: retained adapter gains a dependency → flagless pull rejects unchanged; explicit monotonic selection previews and installs the addition.
NEG-11 / CHK-09: Windows-reserved stems/characters, trailing dots/spaces, every Unicode 15 portable-key vector and directory prefix, existing-entry and same-file collisions, or casing changed after preflight → refusal without writes. Golden positive vectors verify the exact key algorithm.

### Evidence contract

EVID-01…12: stdout/exit codes локальных automated tests и соответствующие GitHub Actions runs
на одном commit каждой стороны интеграции. EC-01: все критерии issue покрыты, suites зелёные,
независимые document/code/simplify reviews завершены без actionable findings.
Открытая release dependency не скрывается: PR явно показывает supporting CLI/bridge order.

CON-02: Component mutation runtime v1 supports Linux/macOS; unsupported hosts report components/adoption capabilities unavailable before writes. Legacy platform support remains unchanged. CHK-01/06 verify this boundary.

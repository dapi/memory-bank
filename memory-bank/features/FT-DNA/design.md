---
title: "FT-DNA: Design"
doc_kind: feature
doc_function: canonical
purpose: Размещение правил достоверности и перенос process obligations к существующим owners.
derived_from:
  - brief.md
  - ../../adr/ADR-002-documentation-trust-boundary.md
status: active
audience: humans_and_agents
---

# FT-DNA: Design

## Design Pack

| Artifact | Relation | Direct ownership | Readiness |
| --- | --- | --- | --- |
| [design.md](design.md) | root | SOL, INV, FM и compatibility ниже | active |
| [ADR-002](../../adr/ADR-002-documentation-trust-boundary.md) | external-dependency | Выбор границы и rationale | active / accepted |

## Selected Solution

- SOL-01 (REQ-01): сохранить каталог из шести файлов; DNA README объясняет
  назначение и границу. Инструкции чтения принадлежат context-priming, действия
  при конфликте — autonomy-boundaries. WHY/WHAT/HOW и выбор ADR уже имеют
  process owners; principles перестаёт задавать конкретные delivery artifacts.
- SOL-02 (REQ-02): principles и governance различают норму, наблюдение, вывод
  и предположение. Существенный claim имеет scope, основание и owner;
  это prose contract без обязательной разметки каждого предложения.
  Существенно утверждение, от которого зависят требования, решения, действия
  или выводы читателя. Вид, scope и основание допустимо указать в самой
  формулировке, общей декларации секции либо по явной аннотированной ссылке;
  общий контекст применяется к секции целиком, локальное исключение называется
  рядом с утверждением. Owner определяется по canonical_for или явно описанной
  границе документа. Reviewer берёт каждое существенное изменённое утверждение
  и должен найти эти сведения без догадки; для наблюдения также нужны условия
  и дата/revision, когда они влияют на применимость. Отсутствующее основание
  обозначается как unknown или непроверенное предположение. Нельзя объявить
  проверку пройденной только по status: active или наличию ссылки.
- SOL-03 (REQ-03): governance описывает неоднозначный ownership и import scope;
  lifecycle — публикацию, утрату актуальности и последствия source changes;
  cross-references — различие навигации, semantic dependency и evidence.
- SOL-04 (REQ-04): сохранить enum, обе формы derived_from и прежние пути.
  canonical_for и doc_kind/doc_function описать в schema; старый heading в
  governance оставить как переход. Glossary и затронутые индексы синхронизировать.

## Constraints and Failure Modes

- INV-01: generic payload не зависит от project-local ADR/FT-DNA; эти artifacts
  фиксируют историю изменения, а не становятся upstream шаблона.
- INV-02: обычная ссылка не даёт semantic authority; canonical_for не делает
  утверждение доказанным и не разрешает нарушать imported constraints.
- FM-01: смешение нормы и наблюдения → разделение claims и явное расхождение.
- FM-02: duplicate owner, неизвестный scope или устаревшее evidence → affected
  claim помечается как unresolved; текст не выбирает победителя эвристически.
- FM-03: чрезмерная формальность → без новых mandatory fields, файлов и
  массового backfill; проверяются существенные изменяемые утверждения.

## Architecture Coverage

C4-00: not required — Markdown template, code/runtime boundaries не меняются.

| Aspect | Status / result |
| --- | --- |
| Components / responsibilities | covered: SOL-01–04, документальные owners |
| Connectors / interactions | covered: semantic links и навигация, SOL-03–04 |
| Configuration / topology | N/A: нет runtime/config изменений |
| Behavioral semantics | covered: документальная интерпретация, FM-01–02 |
| Quality / evolution | covered: compatibility, INV-01–02, FM-03 |

## 4+1 Viewpoint Coverage Decision

| View | Status | Stakeholder / concern / refs |
| --- | --- | --- |
| Logical | covered | Читатель: достоверность, SOL-01–04 |
| Scenarios | covered | Автор и reviewer: все SC/NEG из brief, таблица ниже |
| Process | N/A | Runtime interactions не меняются |
| Development | N/A | Code/module responsibilities не меняются |
| Physical | N/A | Runtime placement не меняется |

## Cross-View Correspondence / Traceability

Для всех строк Process/Development/Physical: N/A по решениям выше.

| Scenario | Requirement | Solution / failure | Check / evidence |
| --- | --- | --- | --- |
| SC-01, NEG-03 | REQ-01 | SOL-01, INV-01 | CHK-01 / EVID-01 |
| SC-02, NEG-02 | REQ-02 | SOL-02, FM-01 | CHK-01 / EVID-01 |
| NEG-01 | REQ-02 | SOL-03, INV-02, FM-02 | CHK-01 / EVID-01 |
| SC-03 | REQ-03 | SOL-03, FM-02 | CHK-01 / EVID-01 |
| SC-04 | REQ-04 | SOL-04, FM-03 | CHK-02 / EVID-02 |

## Design Verification

| Analysis | Required | Method / result |
| --- | --- | --- |
| Contract compatibility | yes | Сверены DNA, glossary, ADR index, priming manifests: пути и metadata сохраняются; уточнение prose не требует CLI migration |
| State / transition completeness | yes | active/draft/archived и отдельные entity states остаются; active+proposed ADR не означает accepted |
| Failure propagation | yes | Source change и duplicate owners разобраны в SC-03/NEG-01: затронутые claims ограничиваются, независимые сохраняются |
| Concurrency / ordering | no | Нет executable concurrency change |
| Security boundaries | no | Нет изменения доступа или исполнения внешних действий |
| Capacity / latency | no | Нет runtime path |
| Migration / evolution safety | yes | Совместимость проверена по существующим fields, обеим формам derived_from и ссылкам; lint/doctor подтвердят реализацию |

## Alternatives and Execution Boundary

Альтернативы и trade-offs принадлежат ADR-002; отдельные local decisions не нужны.
Rollback: локальный revert целого документационного изменения, без data migration.
Delivery и rollout вне этой локальной задачи; runtime contracts не создаются.

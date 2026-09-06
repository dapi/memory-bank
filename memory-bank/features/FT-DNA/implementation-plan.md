---
title: "FT-DNA: Implementation Plan"
doc_kind: feature
doc_function: derived
purpose: Локальное исполнение пересмотра DNA и проверка совместимости документов.
derived_from:
  - brief.md
  - design.md
status: active
audience: humans_and_agents
---

# FT-DNA: Implementation Plan

## Grounding Evidence

Grounded repository revision: `f1f04de843aef45a2425d4a7351d577bbf89e940`.
Grounded at: 2026-09-06. Environment: локальный macOS checkout, Ruby и
установленные memory-bank-cli / code-converge. Runtime приложения отсутствует.

| ID | Inspected path / command | Observed fact | Plan impact |
| --- | --- | --- | --- |
| GRND-01 | Шесть файлов `template/memory-bank/dna/`, перечисленных ниже | 208 строк; active authority, dependency tree, agent conflict instruction и priming в README | Уточнить существующих owners, сохранить пути |
| GRND-02 | `template/memory-bank/flows/autonomy-boundaries.md`, `template/memory-bank/flows/priming/context-priming.md` | Уже владеют полномочиями и чтением baseline | Перенести обязанности без нового process layer |
| GRND-03 | `AGENTS.md`, `.github/workflows/ci.yml`, `tools/validate-priming-manifests.rb`, `tools/refresh-memory-bank-projection.rb` | Существуют link/doctor/priming checks, downstream smoke; generic instance — symlinks | Использовать existing checks; новых prose-mirroring tests не нужно |
| GRND-04 | `memory-bank-cli lint --scope-root template/memory-bank --entrypoint template/memory-bank/README.md`; `memory-bank-cli doctor --profile template` | До изменения payload оба проходят; doctor: 0 errors, 0 warnings | Новая ошибка считается regression |

## Implementation Priming / Exact Realization

Все пути repo-relative. Sections называют исходные заголовки; review проверяет
также новую семантику этих owners. Для строк 1–6 основание GRND-01, для 7–8
GRND-02, для 9–13 — supporting sync, проверенный при discovery.

| Order | Exact path | Section / purpose | Implements | Step |
| --- | --- | --- | --- | --- |
| 1 | `template/memory-bank/dna/principles.md` | Principles: назначение и основания знания | REQ-01–02 / SOL-01–02 | STEP-01 |
| 2 | `template/memory-bank/dna/governance.md` | SSoT Implementation / Source Dependency Tree: ownership, scope, evidence | REQ-02–03 / SOL-02–03, INV-02, FM-01–02 | STEP-01 |
| 3 | `template/memory-bank/dna/frontmatter.md` | Обязательные / Условно обязательные: совместимая schema | REQ-04 / SOL-04, FM-03 | STEP-01 |
| 4 | `template/memory-bank/dna/lifecycle.md` | Maintenance Rules / Sync Checklist: актуальность | REQ-03 / SOL-03, FM-02 | STEP-01 |
| 5 | `template/memory-bank/dna/cross-references.md` | Code → docs / Docs → code: виды ссылок и evidence | REQ-02–03 / SOL-03 | STEP-01 |
| 6 | `template/memory-bank/dna/README.md` | DNA Index / Universal Governance Baseline: граница и маршруты | REQ-01 / SOL-01 | STEP-02 |
| 7 | `template/memory-bank/flows/autonomy-boundaries.md` | Автопилот / Structured Decision Protocol: конфликт документов | REQ-01 / SOL-01 | STEP-02 |
| 8 | `template/memory-bank/flows/priming/context-priming.md` | P1 Universal Baseline And Process Priming: governance.yaml | REQ-01 / SOL-01 | STEP-02 |
| 9 | `template/memory-bank/flows/README.md` | Flows And Templates Index: объяснить границу | supporting REQ-01 | STEP-02 |
| 10 | `template/memory-bank/README.md` | Аннотированный индекс: DNA route | supporting REQ-01 | STEP-02 |
| 11 | `memory-bank/README.md` | Аннотированный индекс: DNA route | supporting REQ-01 | STEP-02 |
| 12 | `docs/glossary.md` | SSoT / Canonical Owner / Authoritative Document / Dependency Tree / Process Layer | supporting REQ-02–04 | STEP-02 |
| 13 | `docs/context-priming.md` | Governance source set: согласовать перенесённый process rule | supporting REQ-01 | STEP-02 |
| 14 | `AGENTS.md` | Команды разработки и проверки: check commands | GRND-03 / REQ-04 | STEP-03 |
| 15 | `.github/workflows/ci.yml` | validate-template: required existing checks | GRND-03 / REQ-04 | STEP-03 |
| 16 | `README.md` | DNA and Single Source of Truth: публичное описание границы | supporting REQ-01–02 | STEP-02 |
| 17 | `README.ru.md` | ДНК и единственный источник истины: производная адаптация после English README | supporting REQ-01–02 | STEP-02 |

Проектные README/brief/design/plan и ADR-002 — supporting delivery artifacts;
их индексы обновляются при создании. INV-01 и C4-00 проверяются по отсутствию
project-specific dependencies и runtime mutations в итоговом payload diff.
ADR-002 реализуется SOL-01–04; отдельного implementation target не создаёт.

## Preconditions / Work

PRE-01: brief и design active, ADR-002 active/accepted, исходный HEAD равен
grounded SHA, перед execution получен clean Plan Ready verdict.

Равенство HEAD — проверка перед первым write в текущем запуске, а не постоянный
запрет на новые commits. Если package был закоммичен до следующего запуска,
сначала повторно проверить target paths относительно старой базы, выполнить
grounding на новом полном HEAD SHA и обновить этот план. Новый active candidate
проходит re-review перед execution. Коммит изменённого плана до первого write
не является prerequisite: review проверяет точный snapshot рабочей tree.
Это сохраняет проверку revision из canonical Feature Flow без бесконечной
цепочки «commit plan → немедленно устаревший grounding».

| Step | Actor / goal | Checks / evidence | Dependencies |
| --- | --- | --- | --- |
| STEP-01 | Автор: пересмотреть DNA по SOL-01–04 | CHK-01 / EVID-01 | PRE-01 |
| STEP-02 | Автор: перенести process obligations и синхронизировать routes/glossary | CHK-01–02 / EVID-01–02 | STEP-01 |
| STEP-03 | Автор: structural checks и simplify pass; запуск инструмента не означает выполнение reviewer-роли | CHK-01–02 / EVID-01–02 | STEP-02 |
| STEP-04 | Независимый reviewer через code-converge: document review замороженной revision; findings исправляет автор, затем новая revision проверяется заново | CHK-01 / EVID-01 | STEP-03 |

Один write stream. Независимые read-only checks допускают параллельное
исполнение; reviewer работает только через canonical code-converge.

## Test Strategy

Profile принадлежит [brief](brief.md#validation-profile-decision).
SC-01–03 и NEG-01–03: semantic read-through и independent review — проверяется
смысл Markdown, а не runtime behavior. Это обычный documentation review, не
manual-only regression gap. SC-04: существующие автоматические проверки.

Команды: `ruby tools/validate-priming-manifests-test.rb`,
`ruby tools/validate-priming-manifests.rb template/memory-bank`,
`memory-bank-cli lint --scope-root template/memory-bank --entrypoint template/memory-bank/README.md`,
`memory-bank-cli lint --repo-root .`, `memory-bank-cli doctor --profile template`,
`ruby tools/refresh-memory-bank-projection.rb`, `git diff --check`.
Также выполнить локальный downstream-init smoke по `.github/workflows/ci.yml`.
Новые tests не нужны: executable/checker behavior не меняется.

Required CI: `validate-template`; remote CI выполняется только после отдельного
разрешения на публикацию. До этого delivery_status не переводится в done.
Review: `code-converge --document-review --max-cycles 0`;
Candidate удерживается отдельной именованной local ref на immutable snapshot
commit по формату из brief (пример: `refs/review-candidates/ft-dna/plan-4`);
его tree совпадает с private snapshot инструмента.
Structured verdict и logs сохраняются в Git note по evidence contract brief,
а не только в cache. Notes не меняют reviewed tree; remote refs не публикуются.

## Checkpoints / Risks / Stop Conditions

- CP-01: clean review draft plan, затем active plan и clean re-review до execution.
- CP-02: local checks и scenario outcomes совпадают с brief; final review clean.
- ER-01: ссылки или prose правила расходятся → обновить owner и dependent refs.
- STOP-01: incompatible field/path change → вернуться к design до execution.
- STOP-02: operational review error → review не выполнена; не подменять механизм.
- Approval: локальные edits/checks разрешены задачей; publish/merge исключены.
- Open questions: none для локального изменения; remote delivery не выполняется.

Evidence: Git notes `refs/notes/dna-trust-review` на reviewed commit SHA
(EVID-01–02 из brief); cache используется лишь для diagnostics. Итоговый
verdict не вписывается в замороженный plan. Backout — локальный revert этого
набора Markdown-правок; live state отсутствует.

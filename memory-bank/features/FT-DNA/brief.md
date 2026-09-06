---
title: "FT-DNA: Достоверность документации"
doc_kind: feature
doc_function: canonical
purpose: Требования к пересмотру DNA как основы достоверной документации.
derived_from:
  - ../../flows/feature.md
  - ../../flows/feature-requirements.md
status: active
delivery_status: planned
audience: humans_and_agents
---

# FT-DNA: Достоверность документации

## What

### Problem and outcome

DNA объединяет принципы документации, формат и инструкции агенту, но слабо
различает авторитетность источника и подтверждённость утверждения. Данил в
текущей задаче подтвердил назначение DNA: поддерживать достоверную документацию;
управление работой агента относится к отдельному слою. Результат — читатель
может определить owner, основание и применимость утверждения и найти правила
исполнения у отдельного process owner.

Route: Feature Flow — меняется общий документационный контракт, подход требует
design; Small Change не подходит. Одна delivery-unit, без runtime изменений.
Authority source: текущая задача Данила от 2026-09-06; подготовка локальных
изменений разрешена. Публикация и merge в scope не входят.

### Requirements

Для всех требований: источник — задача выше, priority — must, accountable owner
— maintainer Memory Bank; метод — semantic inspection и проверки ниже.

| ID | Class | Требуемый результат | Acceptance |
| --- | --- | --- | --- |
| REQ-01 | stakeholder / product | DNA явно отвечает за достоверность документации, правила исполнения имеют отдельного владельца | SC-01 |
| REQ-02 | functional | Читатель различает требование, наблюдение и предположение; может установить owner и область применимости | SC-02, NEG-01, NEG-02 |
| REQ-03 | quality attribute | Изменение источника, конфликт и устаревание имеют однозначные последствия для зависимых утверждений | SC-03, NEG-03 |
| REQ-04 | compatibility | Существующие пути, enum и YAML-формы сохраняются; ссылки и manifests разрешимы | SC-04 |

### Applicability

| Class | Decision / rationale |
| --- | --- |
| stakeholder / product | applicable: REQ-01 |
| functional | applicable: REQ-02 |
| quality attribute | applicable: REQ-03 |
| compatibility | applicable: REQ-04 |
| constraint | applicable: CON-01 |
| verification / acceptance | applicable: SC/NEG/CHK/EVID ниже |
| performance, interface, data, security, safety, regulatory / compliance, operational, deployment / rollout | not-applicable: меняется документационная семантика; runtime, сериализация, окружения и внешние обязательства не меняются |

### Non-scope and constraints

- NS-01: новый agent runner, новый delivery lifecycle, переписывание всех flows.
- NS-02: изменения внешнего CLI, массовая миграция downstream-проектов, публикация.
- CON-01: payload generic; история этой задачи живёт только в project-local package.
- Blocking decisions: none; назначение DNA подтверждено пользователем.

## Design Requirement Decision

Design required: yes — нужны выбор границы ответственности и compatibility analysis.

## Validation Profile Decision

Validation profile: documentation. Меняются Markdown-правила и навигация;
executable behavior, wire/schema format, config и release path не меняются.
Существующие fields/enums сохраняются. Downgrade approval: none.

## Verify

EC-01: все SC/NEG имеют ожидаемые результаты; структурные проверки проходят,
независимая document review не содержит блокирующих findings.

| ID | Context → event → observable outcome | Requirement / check |
| --- | --- | --- |
| SC-01 | Читатель ищет критерии доверия и полномочия на правку → читает DNA → находит критерии в DNA, полномочия у process owner | REQ-01 / CHK-01 |
| SC-02 | Есть нормативный лимит 30 секунд и измерение 45 секунд → сопоставление → требование и наблюдение сохраняются раздельно, расхождение явно названо | REQ-02 / CHK-01 |
| SC-03 | Изменён canonical source → проверка производного описания → оно обновлено либо его неподтверждённая часть явно ограничена | REQ-03 / CHK-01 |
| SC-04 | Существующий документ использует status и обе формы derived_from → проверка обновлённого template → формат и пути остаются допустимыми | REQ-04 / CHK-02 |
| NEG-01 | Два active owner претендуют на один scope → сравнение → дата файла или порядок списка не выбирают победителя | REQ-02 / CHK-01 |
| NEG-02 | active-документ содержит неподтверждённое предположение → использование → active не превращает его в установленный факт | REQ-02 / CHK-01 |
| NEG-03 | Новая инструкция агенту → размещение → DNA не становится владельцем execution procedure | REQ-01, REQ-03 / CHK-01 |

| Check | Метод и ожидаемый результат | Evidence |
| --- | --- | --- |
| CHK-01 | Semantic read-through SC/NEG и независимая code-converge document review; все outcomes различимы и нет blocking findings | EVID-01: structured verdict и scenario results во внешнем локальном review log |
| CHK-02 | Priming validator, template lint, template doctor, projection check, diff whitespace check; нет новых ошибок | EVID-02: локальные command logs |

Evidence сохраняется вне проверяемой revision в task-local каталоге
`~/.cache/memory-bank-dna-trust/`; это результаты исполнения, не новые правила.

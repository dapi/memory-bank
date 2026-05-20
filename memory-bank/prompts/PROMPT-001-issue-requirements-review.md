---
title: "PROMPT-001: Issue Requirements Review"
doc_kind: prompt
doc_function: canonical
purpose: "Проверяет feature-документы против исходного issue: точность требований, отсутствие домыслов и соответствие memory-bank governance."
derived_from:
  - ../dna/governance.md
status: draft
audience: humans_and_agents
prompt_kind: review
prompt_status: drafted
source_prompt: |
  Перечитай github issue {{ISSUE_ID}} через gh и сделай ревью feature по ней
  на точное требование заказчика в issue, отсутствие домыслов и придумок.
  Укажи, какие конкретно страницы или поверхности заказчик хочет кастомизировать.
  Не надо лазить в код и архитектуру. Просто сделай ревью feature относительно issue.
  Проверь feature-{{ISSUE_ID}} на требования memory-bank/dna и feature flow.
variables:
  - name: ISSUE_ID
    required: true
    description: "Issue, относительно которого проверяется feature."
  - name: FEATURE_PATH
    required: true
    description: "Путь к feature package или feature-документу."
  - name: ISSUE_COMMAND
    required: false
    description: "Команда или инструмент для чтения issue, например gh."
  - name: MEMORY_BANK_PATH
    required: false
    description: "Путь к memory-bank, если он отличается от стандартного."
model_notes:
  reasoning: "medium"
  tools: "repo, issue_tracker"
---

# PROMPT-001: Issue Requirements Review

## When To Use

Используй этот prompt, когда нужно проверить feature-документацию строго против исходного issue и governance-правил memory-bank до начала реализации или перед ревью.

Не используй его для ревью кода, архитектуры или implementation details.

## Prompt

```prompt
<role>
Ты requirements reviewer. Твоя задача - проверить feature-документы против исходного issue и правил memory-bank без анализа кода и архитектуры.
</role>

<input>
ISSUE_ID: {{ISSUE_ID}}
FEATURE_PATH: {{FEATURE_PATH}}
ISSUE_COMMAND: {{ISSUE_COMMAND}}
MEMORY_BANK_PATH: {{MEMORY_BANK_PATH}}
</input>

<task>
Перечитай issue, затем проверь feature-документы на точное соответствие требованиям заказчика, отсутствие домыслов и соответствие memory-bank governance.
</task>

<instructions>
1. Прочитай issue через разрешенный issue tracker tool или команду из `ISSUE_COMMAND`.
2. Прочитай только feature-документы из `FEATURE_PATH` и релевантные governance-документы memory-bank: `dna/`, `flows/feature-flow.md`, `flows/workflows.md`.
3. Не исследуй код, runtime architecture, implementation modules или unrelated docs.
4. Выдели дословные или явно выраженные требования заказчика из issue.
5. Проверь, не добавляет ли feature-документация требований, страниц, сценариев, решений или ограничений, которых нет в issue.
6. Отдельно укажи конкретные страницы, экраны, API, операции или другие поверхности, которые заказчик действительно просит изменить или кастомизировать.
7. Проверь, соблюдены ли требования memory-bank: frontmatter, dependency links, lifecycle status, feature-flow sections, stable IDs и traceability.
</instructions>

<constraints>
- Не исправляй документы, если задача только на review.
- Не додумывай intent заказчика по названию issue, коду или архитектуре.
- Если issue неоднозначен, пометь это как open question вместо выбора за заказчика.
- Claims должны ссылаться на issue или конкретный feature-документ.
</constraints>

<output_format>
Верни Markdown-отчет:

## Verdict
`pass` / `pass_with_notes` / `fail`

## Customer Requirements From Issue
Таблица: `Requirement`, `Evidence from issue`, `Mapped feature section`, `Status`.

## Requested Surfaces
Список конкретных страниц, экранов, API, операций или других поверхностей из issue. Если они не названы явно, напиши `not_explicitly_defined`.

## Findings
Таблица: `Severity` (`critical` / `important` / `minor`), `Finding`, `Evidence`, `Required correction`.

## Memory-Bank Compliance
Кратко: frontmatter, derived_from, lifecycle status, feature-flow sections, stable IDs, traceability.

## Open Questions
Вопросы, которые нужно вернуть человеку или заказчику.
</output_format>
```

## Variables

| Variable | Required | Description | Example |
| --- | --- | --- | --- |
| `ISSUE_ID` | yes | Issue, относительно которого проверяется feature. | `#1234` |
| `FEATURE_PATH` | yes | Путь к feature package или feature-документу. | `memory-bank/features/FT-1234/` |
| `ISSUE_COMMAND` | no | Как читать issue. | `gh issue view 1234 --comments` |
| `MEMORY_BANK_PATH` | no | Путь к memory-bank. | `memory-bank/` |

## Validation Notes

| Check | Expected Result | Status |
| --- | --- | --- |
| Dry run on feature package | Report references only issue, feature docs and memory-bank governance. | not_run |

## Change Notes

- 2026-05-19: Migrated from legacy `prompts/01 Issue Receiver.md`; removed shell/session notes from the source capture.

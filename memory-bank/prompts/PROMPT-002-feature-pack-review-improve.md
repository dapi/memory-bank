---
title: "PROMPT-002: Feature Pack Review Improve"
doc_kind: prompt
doc_function: canonical
purpose: "Проводит ограниченный цикл review-improve для комплекта feature-документов и останавливается на human gate при существенной неизвестности."
derived_from:
  - ../dna/governance.md
status: draft
audience: humans_and_agents
prompt_kind: review
prompt_status: drafted
source_prompt: |
  Проделай не более 5-и циклов улучшения качества комплекта документов по feature
  (review-improve): сделай ревью комплекта документов на целостность и
  непротиворечивость, сохрани отчет, закрой открытые вопросы через FPF с
  аргументацией на фактах, исправь критические и важные находки, повтори цикл.
  Если недостаточно данных или есть сомнения в решениях - остановись и сделай
  human gate. Финальный отчет сохрани в директории feature.
variables:
  - name: FEATURE_PATH
    required: true
    description: "Путь к feature package."
  - name: MAX_CYCLES
    required: false
    description: "Максимальное число review-improve циклов."
  - name: REVIEW_REPORT_PATH
    required: false
    description: "Путь для временного отчета ревью."
  - name: FINAL_REPORT_PATH
    required: false
    description: "Путь для финального отчета внутри feature package."
model_notes:
  reasoning: "high"
  tools: "repo"
---

# PROMPT-002: Feature Pack Review Improve

## When To Use

Используй этот prompt, когда feature package уже создан, но нужно довести документы до целостного и непротиворечивого состояния перед реализацией или handoff.

Не используй его для изменения кода или расширения scope feature.

## Prompt

```prompt
<role>
Ты documentation quality agent. Твоя задача - провести ограниченный цикл review-improve для feature package, исправляя только critical и important проблемы, которые можно обоснованно закрыть по имеющимся документам.
</role>

<input>
FEATURE_PATH: {{FEATURE_PATH}}
MAX_CYCLES: {{MAX_CYCLES}}
REVIEW_REPORT_PATH: {{REVIEW_REPORT_PATH}}
FINAL_REPORT_PATH: {{FINAL_REPORT_PATH}}
</input>

<context>
Под feature package понимаются документы в `FEATURE_PATH` и связанные артефакты, которые явно входят в scope этой feature.
Цель: повысить целостность, непротиворечивость, traceability и готовность комплекта к следующей стадии lifecycle.
</context>

<instructions>
Выполни не более `MAX_CYCLES` циклов. Если `MAX_CYCLES` не задан, используй 5.

На каждом цикле:

1. Проведи ревью комплекта feature-документов.
   Проверь:
   - целостность между документами;
   - непротиворечивость;
   - полноту обязательных разделов и ссылок;
   - корректность frontmatter и `derived_from`;
   - открытые вопросы, assumptions, blockers и gaps;
   - расхождения между `brief.md`, conditional `design.md`, `implementation-plan.md`, ADR, verify/evidence и related docs;
   - соответствие `memory-bank/dna` и `memory-bank/flows/feature-flow.md`.

2. Сохрани отчет текущего ревью в `REVIEW_REPORT_PATH`.
   Если путь не задан, используй `./tmp/feature-pack-review.md`.

3. Классифицируй замечания:
   - `critical`: блокирует корректность scope, требований, решений или lifecycle state;
   - `important`: materially снижает готовность, traceability или исполнимость документов;
   - `minor`: улучшение качества, не блокирующее lifecycle.

4. Если `critical` и `important` замечаний нет:
   - останови цикл досрочно;
   - сохрани последний отчет в `FINAL_REPORT_PATH`.
   Если `FINAL_REPORT_PATH` не задан, используй `FEATURE_PATH/feature-review-report.md`.

5. Для каждого open question, который блокирует устранение `critical` или `important` замечаний:
   - сначала попытайся закрыть вопрос только на фактах из текущих документов;
   - используй явно описанный first-principles reasoning или проектный decision framework, если он определен;
   - зафиксируй решение в appropriate owner: `brief.md` для problem-space facts, `design.md` для feature-local solution decisions или ADR для architectural / reusable / cross-feature decisions.

6. Если данных недостаточно, решение неоднозначно или риск неправильного выбора materially влияет на feature:
   - немедленно остановись;
   - не продолжай автоматические исправления;
   - оформи human gate с вопросом, фактами, вариантами, рисками и тем, что требуется от человека.

7. Исправь все `critical` и `important` замечания, которые можно закрыть без human gate.

8. Повтори цикл с шага 1.
</instructions>

<constraints>
- Не исправляй `minor` замечания, если они не нужны для закрытия `critical` или `important`.
- Не вноси изменения за пределами `FEATURE_PATH`, кроме явно связанных upstream/downstream docs, если это необходимо и обосновано.
- Не придумывай требования, факты или решения без опоры на документы.
- Если создаешь новое решение, оно должно быть согласовано с уже существующими решениями.
- Если фиксируешь противоречие, явно укажи конфликтующие документы и как конфликт разрешен.
- Не переходи через human gate молча.
</constraints>

<output_format>
В каждом цикле сообщай:
1. Номер цикла.
2. Краткий итог ревью.
3. Список `critical` и `important` замечаний.
4. Какие open questions были закрыты reasoning и в каких owner-документах это зафиксировано.
5. Какие изменения внесены.
6. Возник ли human gate.

В финале верни:
1. Итоговый статус: `done`, `stopped_by_human_gate` или `max_cycles_reached`.
2. Сколько циклов выполнено.
3. Какие `critical` и `important` замечания закрыты.
4. Какие замечания остались.
5. Путь к финальному review report.
6. Пути к owner-документам, если они обновлялись.
</output_format>
```

## Variables

| Variable | Required | Description | Example |
| --- | --- | --- | --- |
| `FEATURE_PATH` | yes | Путь к feature package. | `memory-bank/features/FT-1234/` |
| `MAX_CYCLES` | no | Максимум циклов review-improve. | `5` |
| `REVIEW_REPORT_PATH` | no | Временный отчет ревью. | `./tmp/feature-pack-review.md` |
| `FINAL_REPORT_PATH` | no | Финальный отчет внутри feature. | `memory-bank/features/FT-1234/feature-review-report.md` |

## Validation Notes

| Check | Expected Result | Status |
| --- | --- | --- |
| Dry run on feature docs with a known contradiction | Report finds contradiction, fixes it or stops at human gate. | not_run |

## Change Notes

- 2026-05-19: Migrated from legacy `prompts/10 Feature Pack Improvers/v0.1 Review + Fix.md`.

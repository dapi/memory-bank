---
title: "PROMPT-XXX: Reusable Prompt Name"
doc_kind: prompt
doc_function: template
purpose: Governed wrapper-шаблон reusable prompt-документа. Читать, чтобы зафиксировать исходную формулировку пользователя в frontmatter и хранить улучшенный prompt как copy-surface в body.
derived_from:
  - ../../../dna/governance.md
  - ../../../dna/frontmatter.md
status: active
audience: humans_and_agents
template_for: prompt
template_target_path: ../../../prompts/PROMPT-XXX-short-name.md
canonical_for:
  - prompt_template
---

# PROMPT-XXX: Reusable Prompt Name

Этот файл описывает wrapper-template. Инстанцируемый prompt-документ живет ниже как embedded contract и копируется без wrapper frontmatter и history.

## Wrapper Notes

Prompt-документ нужен, когда формулировка должна стать повторно используемым артефактом, а не остаться только в истории диалога.

Жизненный цикл:

1. Человек формулирует черновую суть prompt в диалоге с агентом.
2. Агент переносит эту исходную формулировку в `source_prompt` во frontmatter без продуктового переписывания.
3. Агент генерирует или улучшает prompt и помещает итоговую версию в body, в один fenced-блок с language tag `prompt`.
4. Человек или агент копирует только содержимое блока `prompt` для исполнения.
5. Если prompt меняется существенно, обнови `source_prompt`, `prompt_status`, body-блок и `Validation Notes`.

`source_prompt` хранит intent и provenance. Body-блок `prompt` хранит runnable/copyable версию. Не смешивай эти роли: не превращай frontmatter в место для исполняемого prompt, а body не используй как лог диалога.

Если исходная формулировка слишком длинная для frontmatter, используй `source_prompt_ref` на upstream-документ или transcript и оставь в `source_prompt` короткую дословную выжимку. Для обычных prompt-документов предпочитай inline `source_prompt: |`.

## Instantiated Frontmatter

```yaml
title: "PROMPT-XXX: Reusable Prompt Name"
doc_kind: prompt
doc_function: canonical
purpose: "Хранит исходную формулировку и улучшенную copyable-версию reusable prompt."
derived_from:
  - ../dna/governance.md
status: draft
audience: humans_and_agents
prompt_kind: task | system | developer | agent | extraction | review | research | coding
prompt_status: source_captured | drafted | validated | active | archived
source_prompt: |
  Дословно или максимально близко к исходнику: что человек попросил
  сформулировать, улучшить или превратить в reusable prompt.
variables:
  - name: CONTEXT
    required: true
    description: "Какой контекст нужно подставить перед исполнением prompt."
model_notes:
  reasoning: "low | medium | high | not_applicable"
  tools: "none | repo | web | external"
```

## Instantiated Body

````markdown
# PROMPT-XXX: Reusable Prompt Name

## When To Use

Кратко опиши, для какой повторяемой задачи используется этот prompt и когда его не стоит применять.

## Prompt

```prompt
<role>
You are ...
</role>

<context>
{{CONTEXT}}
</context>

<task>
Describe the exact task the model must perform.
</task>

<instructions>
1. Follow the source context and do not invent missing facts.
2. Ask a clarifying question only when the missing information blocks a correct result.
3. Keep the output directly usable for the target workflow.
</instructions>

<constraints>
- Do not expand scope beyond the requested task.
- Preserve project-specific terms exactly as provided in context.
- If facts may have changed, verify them with the allowed tools before making current claims.
</constraints>

<output_format>
Return the result in the format expected by the workflow.
</output_format>
```

## Variables

| Variable | Required | Description | Example |
| --- | --- | --- | --- |
| `CONTEXT` | yes | Input context used by the prompt. | Path, pasted text, issue body, transcript |

## Validation Notes

| Check | Expected Result | Status |
| --- | --- | --- |
| Dry run on representative input | Output follows `output_format` and respects `constraints`. | not_run / passed / failed |

## Change Notes

- YYYY-MM-DD: Created from `source_prompt`.
````

---
title: "PROMPT-003: Implement And Test"
doc_kind: prompt
doc_function: canonical
purpose: "Ведет coding-задачу end-to-end: реализация, локальные проверки, PR, review/fix loop и зеленый CI."
derived_from:
  - ../dna/governance.md
status: draft
audience: humans_and_agents
prompt_kind: coding
prompt_status: drafted
source_prompt: Приступай к реализации, создай PR, проведи PR review и исправь все замечания, убедись что все тесты в CI зеленые. Делай по кругу review-fix до тех пор, пока не останется критических и важных замечаний и CI не станут зелеными, но не более 5-и итераций.
variables:
  - name: TASK_SUMMARY
    required: true
    description: Краткое описание задачи или ссылки на issue/feature.
  - name: FEATURE_CONTEXT
    required: false
    description: Путь к feature docs, issue или другому upstream-контексту.
  - name: BASE_BRANCH
    required: false
    description: Base branch для PR.
  - name: COMMAND_POLICY
    required: false
    description: Проектные правила запуска команд, тестов и сервисов.
  - name: MAX_ITERATIONS
    required: false
    description: Максимум review/fix итераций.
model_notes:
  reasoning: high
  tools: repo, git, ci, issue_tracker
---

# PROMPT-003: Implement And Test

## When To Use

Используй этот prompt, когда агент должен реализовать задачу end-to-end и довести PR до состояния без critical/high замечаний и с зеленым обязательным CI.

Не используй его, если нужен только plan/review без изменения кода.

## Prompt

```prompt
<role>
Ты senior coding agent в текущем репозитории. Твоя задача - реализовать задачу end-to-end, проверить ее локально, опубликовать изменения в PR и довести PR до готовности.
</role>

<input>
TASK_SUMMARY: {{TASK_SUMMARY}}
FEATURE_CONTEXT: {{FEATURE_CONTEXT}}
BASE_BRANCH: {{BASE_BRANCH}}
COMMAND_POLICY: {{COMMAND_POLICY}}
MAX_ITERATIONS: {{MAX_ITERATIONS}}
</input>

<definition_of_done>
Задача считается завершенной только если:
1. Реализация и нужная документация выполнены в пределах scope.
2. Релевантные локальные проверки и тесты запущены и результат понятен.
3. Изменения закоммичены и запушены.
4. Создан или обновлен PR против правильной base branch.
5. PR не имеет merge conflicts.
6. Обязательные CI checks зеленые.
7. Review/fix loop завершен: не осталось critical/high замечаний.
</definition_of_done>

<instructions>
1. Осмотрись:
   - Прочитай `AGENTS.md` и проектные инструкции.
   - Прочитай `FEATURE_CONTEXT`, issue, acceptance criteria и релевантные memory-bank docs, если они есть.
   - Проверь текущую ветку, `git status`, последние коммиты и существующий PR/issue.
   - Если в рабочем дереве есть чужие изменения, не перетирай их.

2. Реализуй:
   - Внеси только изменения, нужные для `TASK_SUMMARY`.
   - Обнови тесты и документацию по change surface.
   - Не делай unrelated refactor.
   - Не додумывай требования, которые не следуют из upstream context.

3. Проверь локально:
   - Следуй `COMMAND_POLICY` и локальным инструкциям репозитория.
   - Запусти минимально достаточные проверки для измененных поверхностей.
   - Если тесты падают из-за твоих изменений, исправь и повтори.
   - Если проверку нельзя запустить, явно зафиксируй причину и риск.

4. Опубликуй:
   - Проверь diff перед commit.
   - Закоммить изменения согласно commit policy проекта.
   - Запушь ветку.
   - Создай PR, если его нет; если есть, обнови существующий.
   - Укажи в PR summary, что изменено и какие проверки запускались.

5. Проведи review/fix loop:
   - Проверь собственный diff.
   - Собери замечания из review comments, CI, статических проверок и доступных quality signals.
   - Исправь все critical/high замечания.
   - Повтори локальные проверки, commit, push и проверку CI.
   - Продолжай до состояния: нет critical/high замечаний и обязательный CI зеленый.
   - Лимит: `MAX_ITERATIONS`, если задан, иначе 5 итераций.

6. Остановись и отчитайся, если:
   - после лимита остались blockers;
   - нужен human approval для рискованного действия;
   - отсутствуют данные, без которых реализация будет домыслом;
   - внешняя система или CI недоступны и это блокирует DoD.
</instructions>

<constraints>
- Не игнорируй failing CI.
- Не объявляй готовность без PR, если задача требует PR.
- Не закрывай задачу при наличии critical/high замечаний.
- Не выполняй команды, запрещенные `AGENTS.md` или `COMMAND_POLICY`.
- Не откатывай чужие изменения без явного разрешения.
</constraints>

<output_format>
В финальном ответе кратко укажи:
- PR URL или почему PR не создан.
- Последний commit SHA.
- CI status.
- Merge conflict status.
- Результат review/fix loop.
- Какие проверки запускались.
- Остались ли blockers.
</output_format>
```

## Variables

| Variable | Required | Description | Example |
| --- | --- | --- | --- |
| `TASK_SUMMARY` | yes | Суть coding-задачи. | `Fix vendor lookup by URL` |
| `FEATURE_CONTEXT` | no | Upstream context. | `memory-bank/features/FT-1234/` |
| `BASE_BRANCH` | no | Base branch для PR. | `main` |
| `COMMAND_POLICY` | no | Правила команд проекта. | `Run tests through ./bin/dev test` |
| `MAX_ITERATIONS` | no | Лимит review/fix loop. | `5` |

## Validation Notes

| Check | Expected Result | Status |
| --- | --- | --- |
| Dry run on small repo task | Agent inspects instructions, scopes work, runs checks, reports PR/CI status. | not_run |

## Change Notes

- 2026-05-19: Migrated from legacy `prompts/20_Implement_And_Test.md`.

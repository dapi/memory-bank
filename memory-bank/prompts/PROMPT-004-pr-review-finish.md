---
title: "PROMPT-004: PR Review Finish"
doc_kind: prompt
doc_function: canonical
purpose: "Завершает текущую feature branch: commit/push, PR, CI, merge conflict check и review-improve до отсутствия critical/high замечаний."
derived_from:
  - ../dna/governance.md
status: draft
audience: humans_and_agents
prompt_kind: coding
prompt_status: drafted
source_prompt: |
  Заверши стадию реализации фичи. Критерии завершенности: все закомичено и
  запушено, есть PR, CI у PR зеленый, в PR нет merge conflicts, проведен и
  завершен процесс review-improve, по результатам review нет критических и
  важных замечаний. Процесс review-improve: закомитить и запушать изменения,
  выполнить review текущей ветки, исправить все критические и важные замечания,
  повторить до FINISH.
variables:
  - name: REPO_PATH
    required: false
    description: "Путь к репозиторию, если prompt запускается не из repo root."
  - name: ISSUE_ID
    required: false
    description: "Issue или task id, который должен быть отражен в commit/PR."
  - name: FEATURE_SUMMARY
    required: true
    description: "Краткое описание завершаемой фичи."
  - name: BASE_BRANCH
    required: false
    description: "Base branch PR."
  - name: CURRENT_BRANCH
    required: false
    description: "Текущая feature branch, если известна."
  - name: COMMAND_POLICY
    required: false
    description: "Проектные правила команд, тестов, сервисов и cleanup."
  - name: COMMIT_POLICY
    required: false
    description: "Формат commit subject/body и ссылки на issue."
  - name: REVIEW_COMMAND
    required: false
    description: "Команда или процедура review текущей ветки."
  - name: CLEANUP_COMMAND
    required: false
    description: "Команда cleanup после работы, если требуется."
model_notes:
  reasoning: "high"
  tools: "repo, git, ci, issue_tracker"
---

# PROMPT-004: PR Review Finish

## When To Use

Используй этот prompt, когда реализация уже начата или почти завершена, но нужно довести ветку до готового PR: commit/push, PR, CI, merge conflicts и review-improve.

Не используй его для первичного product discovery или проектирования feature scope.

## Prompt

```prompt
<role>
Ты senior coding agent в текущем репозитории. Твоя задача - завершить стадию реализации feature branch и довести PR до готового состояния.
</role>

<input>
REPO_PATH: {{REPO_PATH}}
ISSUE_ID: {{ISSUE_ID}}
FEATURE_SUMMARY: {{FEATURE_SUMMARY}}
BASE_BRANCH: {{BASE_BRANCH}}
CURRENT_BRANCH: {{CURRENT_BRANCH}}
COMMAND_POLICY: {{COMMAND_POLICY}}
COMMIT_POLICY: {{COMMIT_POLICY}}
REVIEW_COMMAND: {{REVIEW_COMMAND}}
CLEANUP_COMMAND: {{CLEANUP_COMMAND}}
</input>

<definition_of_done>
Стадия реализации считается завершенной только если:
1. Все нужные изменения закоммичены и запушены.
2. Есть PR для текущей ветки.
3. PR открыт против правильной base branch.
4. В PR нет merge conflicts.
5. Обязательный CI зеленый.
6. Проведен review-improve loop.
7. По результатам review нет critical/high или critical/important замечаний, в зависимости от терминологии проекта.
8. Выполнен required cleanup из `CLEANUP_COMMAND`, если он задан.
</definition_of_done>

<instructions>
1. Осмотрись:
   - Перейди в `REPO_PATH`, если он задан.
   - Прочитай `AGENTS.md`, `COMMAND_POLICY` и project docs.
   - Проверь текущую ветку, `git status`, последние коммиты, связанный issue и PR.
   - Если PR уже есть, работай с ним. Если PR нет, создай его после первого push.
   - Определи, что еще не завершено по `FEATURE_SUMMARY`, issue, acceptance criteria и текущему diff.

2. Доделай реализацию:
   - Исправь недостающую логику, тесты, документацию или конфигурацию только в пределах feature scope.
   - Не делай unrelated refactor.
   - Не откатывай чужие изменения.
   - Если чужие незакоммиченные изменения блокируют работу, остановись и опиши human gate.

3. Проверь локально:
   - Запусти релевантные проверки строго по `COMMAND_POLICY`.
   - Если проверки требуют сервисов, следуй setup/teardown policy проекта.
   - Если тесты падают из-за твоих изменений, исправь и повтори.

4. Commit + push:
   - Проверь git diff перед commit.
   - Закоммить только изменения, относящиеся к feature.
   - Следуй `COMMIT_POLICY`; если он не задан, используй короткий conventional commit.
   - Запушь текущую ветку.

5. PR:
   - Если PR отсутствует, создай его через project-approved tool.
   - В PR укажи суть изменений, проверки и ссылку на issue, если issue задан.
   - Проверь merge conflict status. Если есть conflicts, разреши их, закоммить и запушь.

6. CI:
   - Дождись результата обязательного CI.
   - Если CI красный, изучи логи, исправь причину, commit, push и снова дождись CI.
   - Не объявляй готовность, пока обязательный CI не зеленый или пока недоступность CI не оформлена как blocker.

7. Review-improve loop:
   - Убедись, что все изменения закоммичены и запушены.
   - Выполни `REVIEW_COMMAND`, если он задан; иначе проведи review текущего diff/PR доступными средствами.
   - Если critical/high или critical/important замечаний нет, переходи к FINISH.
   - Исправь все такие замечания.
   - Закоммить и запушь исправления.
   - Повтори review.
   - Максимум 5 итераций, если project policy не задает иной лимит.

8. FINISH:
   - Финально проверь git status, push status, PR URL, merge conflicts, CI и review status.
   - Выполни `CLEANUP_COMMAND`, если он задан.
</instructions>

<constraints>
- Failing CI нельзя игнорировать.
- PR нельзя считать готовым при unresolved merge conflicts.
- Critical/high замечания должны быть исправлены или явно оформлены как blocker/human gate.
- Не запускай команды, запрещенные `COMMAND_POLICY`.
- Не перетирай чужие изменения.
</constraints>

<output_format>
В финальном ответе кратко укажи:
- PR URL.
- Последний commit SHA.
- CI status.
- Merge conflict status.
- Результат review-improve.
- Какие проверки запускались.
- Был ли выполнен `CLEANUP_COMMAND`.
- Если что-то не завершено, назови blocker и текущий статус.
</output_format>
```

## Variables

| Variable | Required | Description | Example |
| --- | --- | --- | --- |
| `REPO_PATH` | no | Путь к репозиторию. | `/path/to/repo` |
| `ISSUE_ID` | no | Issue/task id. | `#1234` |
| `FEATURE_SUMMARY` | yes | Что нужно довести до готовности. | `Finish vendor lookup fix` |
| `BASE_BRANCH` | no | Base branch PR. | `main` |
| `CURRENT_BRANCH` | no | Текущая branch. | `fix/vendor-lookup` |
| `COMMAND_POLICY` | no | Правила запуска команд. | `Use ./bin/dev for tests` |
| `COMMIT_POLICY` | no | Commit subject/body policy. | `fix(issue-1234): description` |
| `REVIEW_COMMAND` | no | Review procedure. | `/review current branch` |
| `CLEANUP_COMMAND` | no | Cleanup after work. | `./bin/dev down` |

## Validation Notes

| Check | Expected Result | Status |
| --- | --- | --- |
| Dry run on active branch | Agent reports PR, CI, conflict and review status without project-specific hardcoding. | not_run |

## Change Notes

- 2026-05-19: Migrated from legacy `prompts/30_Добить PR Review.md`; project-specific repository path, commands and issue URLs were converted to variables.

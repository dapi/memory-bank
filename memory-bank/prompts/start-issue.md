---
title: Start Issue
doc_kind: prompt
doc_function: canonical
purpose: "Стартовый prompt для dapi/start-issue: передаёт отрендеренный GitHub issue URL в PROMPT-005 и поддерживает Codex Human Gate."
derived_from:
  - ../dna/governance.md
  - PROMPT-005-route-and-deliver-issue.md
status: draft
audience: humans_and_agents
prompt_kind: agent
prompt_status: drafted
source_prompt: |
  Запустить оркестрацию memory-bank для issue через dapi/start-issue, используя
  placeholder номера или URL issue, который поддерживает start-issue.
variables:
  - name: ISSUE_URL
    required: true
    description: "Подставляется dapi/start-issue через `{ISSUE_URL}`."
  - name: BASE_BRANCH
    required: true
    description: "Подставляется dapi/start-issue через `{BASE_BRANCH}`."
  - name: BRANCH_NAME
    required: false
    description: "Подставляется dapi/start-issue через `{BRANCH_NAME}`."
  - name: WORKTREE_PATH
    required: false
    description: "Подставляется dapi/start-issue через `{WORKTREE_PATH}`."
model_notes:
  reasoning: "high"
  tools: "repo, git, ci, issue_tracker, agent_delegation, codex_cli"
---

# Start Issue

Используй отрендеренные значения ниже и выполни [`PROMPT-005: Route And Deliver Issue`](PROMPT-005-route-and-deliver-issue.md).

```prompt
ISSUE_URL: {ISSUE_URL}
BASE_BRANCH: {BASE_BRANCH}
BRANCH_NAME: {BRANCH_NAME}
WORKTREE_PATH: {WORKTREE_PATH}
MAX_REVIEW_ITERATIONS: 3

Запусти routing для этой issue. После Intake при необходимости делегируй
read-only custom agents `code-grounding` и `requirements-risk` либо
`test-surface`. После успешных flow gates назначь `delivery-owner` единственным
writer. Перед closure выполни `codex review --base "{BASE_BRANCH}"`.

Доведи задачу до terminal state выбранного flow либо до явно оформленного Human
Gate. В финальном сообщении добавь ровно одну отдельную terminal строку:
`STATUS: DONE` или `STATUS: HUMAN_GATE`.
```

## start-issue Usage

```shell
start-issue 123 --agent codex --prompt-file memory-bank/prompts/start-issue.md --human-gate
```

`{ISSUE_URL}` — placeholder `dapi/start-issue`; он принимает и номер issue, и полный GitHub URL на входе CLI.

## Validation Notes

| Check | Expected Result | Status |
| --- | --- | --- |
| `start-issue --dry-run` | `{ISSUE_URL}` и `{BASE_BRANCH}` отрендерены, неизвестных placeholders нет. | not_run |

## Change Notes

- 2026-07-22: Created for dapi/start-issue and Codex Human Gate execution.

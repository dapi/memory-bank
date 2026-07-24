---
tracker:
  kind: github
  active_states:
    - open
  terminal_states:
    - closed
  required_labels:
    - codex-ready

workspace:
  root: .symphony-workspace

hooks:
  after_create: |
    git -C ../.. worktree add --detach "$PWD"

agent:
  max_concurrent_agents: 3
  max_turns: 8

codex:
  command: codex app-server
  approval_policy: never
  thread_sandbox: workspace-write
  turn_sandbox_policy:
    type: workspaceWrite
    writableRoots: []
    networkAccess: true
---

You are implementing GitHub issue {{ issue.identifier }} in the configured
repository.

Issue title: {{ issue.title }}

Issue body:
{{ issue.description }}

Work only on this issue. Read `AGENTS.md` and the repository documentation before
making changes. Follow the repository governance and run the relevant checks.

Create a branch named `codex/{{ issue.identifier | downcase }}`. Implement the
smallest complete change, commit it, push the branch, and open a pull request.

Use the `github_api` tracker tool for issue updates; do not require a separate
GitHub CLI login. Once a pull request exists:

1. Comment on the issue with the pull request URL and the verification performed.
2. Replace the `codex-ready` label with `human-review`, preserving any unrelated
   labels.
3. Do not close the issue, merge the pull request, or remove `human-review`.

If you need a security-sensitive action, access beyond the configured workspace,
or a decision that cannot be inferred from the issue and repository, stop and
report the blocker in the issue rather than guessing.

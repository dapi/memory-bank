# Symphony with GitHub Issues

This repository's [`WORKFLOW.md`](../WORKFLOW.md) configures the experimental
Symphony Elixir runner to dispatch selected GitHub Issues to Codex. It uses the
upstream GitHub Issues adapter; no tracker extension is required.

## Lifecycle

```text
open + codex-ready
  -> Symphony dispatches Codex in an isolated workspace
  -> Codex pushes a branch and opens a pull request
  -> open + human-review
  -> maintainer reviews, merges, and closes the issue
```

Only issues with the `codex-ready` label are eligible. The agent replaces it
with `human-review` after opening a pull request, which prevents another agent
run while a maintainer reviews the change. Add `codex-ready` again only to
explicitly request rework.

## Prerequisites

- Install `codex` and authenticate it for the machine running Symphony.
- Install Git, and configure SSH access that can push to
  `git@github.com:dapi/memory-bank.git`.
- Install and authorize `direnv` for this repository's `.envrc`.
- Create a GitHub token with access to this repository's issues and pull
  requests. Symphony uses it for the host-side `github_api` tool and does not
  pass it to the Codex subprocess.
- Install `mise`; the bootstrap script installs the required Elixir/Erlang
  versions.

Keep the token outside the repository. After a fresh clone, create the ignored
local file from the tracked template, then authorize direnv:

```sh
cp .env.local.example .env.local
direnv allow
```

It expects the GitHub token at `pass:github/homebrew-token`. If the local
password-store layout differs, change only the ignored `.env.local`; never add
the literal token to a tracked file.

## Run

Bootstrap the `dapi/symphony` fork and create this repository's ignored,
project-local `.symphony-workspace/` directory:

```sh
direnv allow
./bootstrap-symphony.sh
./run-symphony.sh
```

`bootstrap-symphony.sh` defaults to a sibling `../symphony` checkout and the
`https://github.com/dapi/symphony.git` fork. Override either for another local
layout with `SYMPHONY_HOME` or `SYMPHONY_REPOSITORY`.

## Human gates

The configured post-PR `human-review` label is the process gate: a maintainer
reviews the pull request and decides whether to merge or request rework.

The current upstream Elixir runner does not provide a durable human-approval
inbox that can send a later decision back through app-server. `WORKFLOW.md`
therefore sets `approval_policy: never`: the runner can complete the Git
commands required to create a branch, commit, push, and open a pull request.
The configured `workspace-write` sandbox remains in effect, but this is still
a trusted-environment setting. Do not use `on-request` unless a separate
app-server client or a runner extension will collect the human decision and
answer the pending JSON-RPC request.

Do not run this against all open issues or with broad credentials until the
workflow has been evaluated in a trusted environment.

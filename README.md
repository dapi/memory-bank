# Memory Bank

<p align="center">
  <img src="docs/assets/memory-bank-mark.svg" alt="Memory Bank: project context, governed routing, and verified delivery" width="180">
</p>

**A version-controlled development system that gives coding agents durable knowledge, explicit governance, and repeatable delivery flows.**

[Русская версия](README.ru.md) · [Quick start (Russian)](docs/quick-start.md) ·
[Adoption guide (Russian)](docs/adoption.md) ·
[Daily usage (Russian)](docs/usage.md)

## Example: complete GitHub issue #123

![Running a task through Memory Bank routing](docs/assets/quick-start-routing-en.gif)

1. You give the agent an issue and point it to
   `memory-bank/flows/routing.md`.
2. The agent reads the task and project context, then Task Routing selects the
   smallest process that still controls the risk.
3. The selected process governs the required documents, code changes, and
   verification. Lasting decisions and evidence return to their canonical
   owners in Memory Bank.

The result is a feedback loop: project knowledge guides delivery, and delivery
improves project knowledge.

## What you get

- **Durable project context** — product intent, domain language, engineering
  rules, and operational constraints survive across agent sessions.
- **A Single Source of Truth** — every canonical fact has one owner; derived
  documents point back to that source instead of becoming competing copies.
- **Governed delivery** — task routing selects the smallest suitable process for
  incidents, bugs, research, small changes, epics, refactoring, or features.
- **Reusable reasoning tools** — templates make the agent state the problem,
  constraints, selected solution, implementation steps, and verification
  evidence explicitly.
- **A self-growing knowledge base** — delivery leaves behind decisions,
  requirements, scenarios, and evidence that future work can reuse.
- **A portable starting point** — an agent installs the template in a repository
  and adapts it from that project's own evidence.

## Install in a project

You need Git, an installed and authenticated
[Codex CLI](https://developers.openai.com/codex/cli/), and a project repository.
Run the matching command from the project root.

### Existing project

```bash
codex --search \
  'This is an existing project. Follow https://github.com/dapi/memory-bank/blob/main/docs/brownfield-adaptation-protocol.md.'
```

### New project

```bash
codex --search \
  'This is a new project. Follow https://github.com/dapi/memory-bank/blob/main/docs/greenfield-integration-protocol.md.'
```

The agent studies the repository, installs the tracked template payload, and
adapts it to confirmed project facts. The expected starting point is:

```text
memory-bank/
init.sh
```

Review the installation before continuing:

```bash
git status --short
git diff --check
```

For reproducible use, replace `main` in the protocol URL with an immutable
commit SHA. The [adoption guide (Russian)](docs/adoption.md) explains the full
lifecycle, expected artifacts, and completion criteria.

## Run the first task

After Memory Bank is adapted, give Codex a real task and the routing entrypoint:

```bash
codex -C . \
  'Read GitHub issue #123, ./memory-bank/README.md, and ./memory-bank/flows/routing.md.
Choose the applicable process and follow its canonical lifecycle. Report the
route, changed artifacts, verification, and open risks.'
```

Replace `#123` with the real issue number, or describe the task directly if the
project does not use GitHub Issues. A successful run leaves a sufficient,
verifiable trail rather than the largest possible set of documents.

## Where to go next

Most supporting guides are currently available in Russian.

| Goal | Read or use |
| --- | --- |
| Complete a guided first task | [Quick start](docs/quick-start.md) |
| Adapt Memory Bank to a new or existing repository | [Adoption guide](docs/adoption.md) |
| Use Memory Bank for daily delivery | [Daily usage](docs/usage.md) |
| Prepare only the context relevant to one task | [Context priming](docs/context-priming.md) |
| Automate issue startup | [`start-issue`](https://github.com/dapi/start-issue) or [Symphony](docs/symphony-github-issues.md) |
| Look up project-memory terminology | [Glossary](docs/glossary.md) |

## How it works

### Knowledge, governance, and delivery

Memory Bank combines three parts that reinforce one another:

1. **A project knowledge base** for product, domain, engineering, operations,
   requirements, and decisions.
2. **A governance layer** that defines who owns each fact, how documents depend
   on one another, and which source wins when documents disagree.
3. **A delivery system** whose processes turn tasks into governed artifacts,
   implementation, verification, and new durable knowledge.

Memory Bank is built on the **First Principles Framework (FPF)**. Work starts
from explicit facts, constraints, assumptions, and desired outcomes; decisions
preserve their rationale and evidence instead of disappearing into a chat
session.

It is not a wiki, task tracker, or agent runner. It is the development control
plane around those tools: durable context, ownership rules, lifecycle gates,
reusable processes, and verification contracts.

It is useful when project intent has to be reconstructed from chat history,
rules drift across documents, implementation starts before acceptance is clear,
or another agent cannot resume the work from repository state.

### DNA and Single Source of Truth

The `dna/` layer is the constitution of the knowledge base. It defines Single
Source of Truth, document ownership, dependency direction, lifecycle,
frontmatter, and navigation rules.

A canonical document owns a fact. Another document may derive a requirement,
plan, or view from it, but must preserve the dependency. When documents
disagree, ownership and dependency direction identify the authoritative source.

### Project knowledge

Stable project context lives in `product/`, `domain/`, `engineering/`, and
`ops/`. Research, product initiatives, scenarios, delivery packages, and
decisions live in `research/`, `prd/`, `epics/`, `use-cases/`, `features/`, and
`adr/`.

Documents own intent, requirements, rationale, and contracts. Code owns
implementation. A fresh agent session can therefore resume from the same task
and canonical sources without reconstructing the project from chat history.

### Processes and Feature Packs

The `flows/` layer describes repeatable processes that an agent can follow.
Every task begins with
[Task Routing](template/memory-bank/flows/routing.md), which selects the
applicable lifecycle and its evidence requirements.

For a substantial feature, Feature Flow treats the change as a testable
vertical slice and follows specification-driven development. It produces a
Feature Pack in three stages:

```text
brief.md                 design.md                  implementation-plan.md
what and why      →      chosen solution     →     implementation and checks
problem space            solution space             execution space
```

- `brief.md` owns the problem, scope, requirements, and verification contract;
- the Design Pack owns the selected solution, its rationale, and
  solution-level contracts;
- `implementation-plan.md` owns execution sequencing and checkpoints.

The documents required by the selected route are created and reviewed before
implementation begins. Implementation changes the code, while lasting
decisions and evidence return to their canonical owners. The Feature Pack
remains as a durable account of what changed, why it changed, and how the result
was verified.

### Templates as reasoning tools

Templates in `flows/templates/` are not merely blank forms. They require an
agent to separate the problem, solution, execution, and verification; name
assumptions and constraints; compare meaningful alternatives; and preserve
traceability. Filling the template improves the decision process as well as its
documentation.

## Automation

Memory Bank does not require a runner or CLI. Automation is optional:

- [`start-issue`](https://github.com/dapi/start-issue) prepares a branch and
  worktree, then launches the configured agent for one issue;
- [`memory-bank-cli`](docs/memory-bank.md) adds ownership-aware updates, link
  checks, diagnostics, and downstream CI;
- the experimental [Symphony integration](docs/symphony-github-issues.md)
  dispatches selected GitHub Issues to Codex in isolated workspaces and hands
  completed pull requests to human review.

Runners launch agents and repository work. Memory Bank supplies the knowledge,
governance, delivery processes, and verification contracts those agents follow.

## Template layout

This repository is the upstream source. An agent copies the tracked payload in
`template/` into a downstream repository: `template/memory-bank/` becomes
`memory-bank/`, while `template/init.sh` becomes `./init.sh`.

| Area | Purpose |
| --- | --- |
| [`dna/`](template/memory-bank/dna/README.md) | Governance, Single Source of Truth, lifecycle, and document contracts |
| [`product/`](template/memory-bank/product/README.md) | Vision, customers, metrics, marketing, and roadmap |
| [`domain/`](template/memory-bank/domain/README.md) | Glossary, domain model, rules, states, events, and context map |
| [`engineering/`](template/memory-bank/engineering/README.md) | Architecture, testing, coding style, Git workflow, and agent autonomy |
| [`ops/`](template/memory-bank/ops/README.md) | Development, environments, configuration, releases, and runbooks |
| [`research/`](template/memory-bank/research/README.md), [`prd/`](template/memory-bank/prd/README.md), [`epics/`](template/memory-bank/epics/README.md) | Discovery and initiative-level planning |
| [`use-cases/`](template/memory-bank/use-cases/README.md), [`features/`](template/memory-bank/features/README.md), [`adr/`](template/memory-bank/adr/README.md) | Scenarios, delivery packages, and architecture decisions |
| [`flows/`](template/memory-bank/flows/README.md) | Task lifecycles and reusable document templates |

After installation, `memory-bank/README.md` is the primary index inside the
downstream project.

## Reference

- [BDD, user stories, and use cases](docs/bdd-user-stories-and-use-cases.md)
- [Ownership and safe updates](docs/ownership.md)
- [Managed agent instructions](docs/agent-instructions.md)
- [Repository development](docs/development.md)
- [Detailed Russian adaptation](README.ru.md)

The governance model applies the
[MECE principle](https://en.wikipedia.org/wiki/MECE_principle): categories
should be mutually exclusive and collectively exhaustive within their declared
scope.

The optional CLI is developed separately in
[`dapi/memory-bank-cli`](https://github.com/dapi/memory-bank-cli). This template
is available under the [Apache License 2.0](LICENSE).

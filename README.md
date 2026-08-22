# Memory Bank

**A durable, version-controlled context and governance layer for software development with coding agents.**

[Русская версия](README.ru.md) · [Adoption guide](docs/adoption.md) · [Daily usage](docs/usage.md)

**`AGENTS.md` can tell an agent how to start. Memory Bank preserves what the project means, why decisions were made, and how work is verified.**

Coding agents are most useful when they share the same understanding of the product, domain, architecture, constraints, and definition of done. Memory Bank keeps that knowledge in Git, next to the code, instead of leaving it in one person's head or a disposable chat session.

It gives humans and agents an authoritative starting point, routes work through explicit delivery flows, and preserves the decisions and evidence needed to resume a task in a fresh session.

Memory Bank is not a wiki, task tracker, or agent runner. It is the control plane around those tools: durable context, ownership rules, lifecycle gates, and verification contracts.

Use it when a project has one or more of these symptoms:

- a fresh agent has to reconstruct product intent from chat history;
- the same rule appears in several documents and drifts;
- implementation starts before requirements, risks, or acceptance are clear;
- a task cannot be resumed without the person who ran the previous session;
- a successful test is reported without a durable link to what was verified.

## What you get

- **Durable project context** — product intent, domain language, engineering rules, and operational constraints survive across sessions.
- **Clear ownership** — Single Source of Truth rules prevent the same fact from drifting across documents.
- **Governed delivery** — task routing selects the smallest suitable flow for incidents, bugs, research, small changes, epics, refactoring, or features.
- **Portable starting point** — an agent brings the template into a repository and adapts it from the project's own evidence.
- **Optional automation** — a companion CLI can later add ownership-aware updates and automated documentation checks.

```text
Memory Bank context and rules
             ↓
         Issue / task
             ↓
   Agent session and delivery flow
             ↓
 Implementation → verification → PR
             ↓
 New durable knowledge returns to Memory Bank
```

## Start with an agent

Copy the prompt that matches the repository. It tells the agent to bring in and
adapt Memory Bank; the linked protocols define the full lifecycle.

### Greenfield

```text
This is a new project. Follow
https://github.com/dapi/memory-bank/blob/main/docs/greenfield-integration-protocol.md.
```

### Brownfield

```text
This is an existing project. Follow
https://github.com/dapi/memory-bank/blob/main/docs/brownfield-adaptation-protocol.md.
```

For reproducible use, replace `main` in a protocol URL with an immutable commit
SHA.

## Deliver a task

After Memory Bank is adapted, give the agent the task and this prompt:

```text
Read the task, ./memory-bank/README.md, and ./memory-bank/flows/routing.md.
Choose the applicable flow and follow its canonical lifecycle. Report the route,
changed artifacts, verification, and open risks.
```

The [daily usage guide](docs/usage.md) explains the task-to-flow-to-verification
cycle and its smaller routes.

## How it works

The `dna/` layer defines document governance: source ownership, dependency direction, lifecycle, frontmatter, and navigation. Stable project context lives in `product/`, `domain/`, `engineering/`, and `ops/`. Requirements and decisions mature through research, PRDs, epics, use cases, feature packages, and ADRs.

For a substantial delivery feature, the context typically develops in three stages:

```text
brief.md                 design.md                  implementation-plan.md
what and why      →      chosen solution     →     implementation and checks
problem space            solution space             execution space
                         (when required)
```

Documents own intent, requirements, rationale, and contracts. Code owns implementation. A new agent session can therefore restart from the same task and canonical documents without reconstructing the project from chat history.

Every task begins with [Task Routing](template/memory-bank/flows/routing.md), which selects the applicable lifecycle and its evidence requirements.

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

After installation, `memory-bank/README.md` is the primary index inside the downstream project.

## Documentation

- [Adopting Memory Bank](docs/adoption.md)
- [Using Memory Bank day to day](docs/usage.md)
- [Context priming for an agent task](docs/context-priming.md)
- [Glossary](docs/glossary.md)
- [Optional CLI automation](docs/memory-bank.md)
- [Ownership and safe updates](docs/ownership.md)
- [Repository development](docs/development.md)
- [Detailed overview in Russian](README.ru.md)

The governance model applies the [MECE principle](https://en.wikipedia.org/wiki/MECE_principle): categories should be mutually exclusive and collectively exhaustive within their declared scope.

The optional CLI for safe updates and automated checks is developed separately
in [`dapi/memory-bank-cli`](https://github.com/dapi/memory-bank-cli). This
template is available under the [Apache License 2.0](LICENSE).

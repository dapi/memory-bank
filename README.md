# Memory Bank

<p align="center">
  <img src="docs/assets/memory-bank-mark.svg" alt="Memory Bank: project knowledge, document ownership, and optional delivery flows" width="180">
</p>

**Version-controlled project documentation with clear ownership and optional AI delivery processes.**

[Русская версия](README.ru.md) · [Component adoption](docs/component-adoption.md) ·
[CLI integration](docs/memory-bank.md)

## Choose how much to adopt

Memory Bank has three components with one-way dependencies:

| Component | Responsibility | Requires |
| --- | --- | --- |
| DNA | Single Source of Truth, ownership, publication status, metadata and navigation | Nothing |
| Documents | Document types, base templates and project sections | DNA |
| Flows | AI routing, priming, delivery stages, gates and document extensions | DNA + Documents |

DNA works on its own. Documents can be used by people without an AI process or
runner. Adding Flows later preserves project documents; an existing document
enters a flow only through explicit adoption.

| Preset | Installed components | Tool adapters |
| --- | --- | --- |
| `core` | DNA | Explicit additions |
| `docs` | DNA + Documents | Explicit additions |
| `full` | DNA + Documents + Flows | Explicit additions |
| `legacy` | All three | Previous integrations included |

A fresh installation without a preset uses `legacy` for compatibility. A pull
without selection flags keeps the recorded selection. Component removal is not
supported. Adapters declare their dependencies, so choosing one can add Flows.

## Install in a project

Use Git and a component-capable `memory-bank-cli` on Linux or macOS. Component
support is a coordinated template/CLI change: use the reviewed CLI candidate or
a release that reports both required capabilities. An older release is not
sufficient merely because it installs legacy templates.

From a clean, pinned checkout of this template, run the guarded entrypoint
against your project:

```bash
memory-bank-cli capabilities --require components/v1 --require adoption/v1
./tools/install-components.sh init \
  --repo-root /path/to/project --preset docs
```

The entrypoint checks capabilities before invoking the installer and pins its
own source commit. Review the resulting changes:

```bash
git -C /path/to/project status --short
git -C /path/to/project diff --check
memory-bank-cli doctor --repo-root /path/to/project
```

An existing legacy installation needs a separate reviewed migration; ordinary
pull does not opt it in. See [component adoption and migration](docs/component-adoption.md).

## Create project documents

Documents supplies ADRs, feature briefs, PRDs, use cases, research briefs and
epic charters. Base templates live in `memory-bank/templates/`; their type
contracts live in `memory-bank/document-types/`.

```bash
memory-bank-cli document create --repo-root /path/to/project \
  --type feature --path memory-bank/features/FT-123/brief.md
```

The new document belongs to the project and has no flow adoption, including in
`full` and `legacy`. Fill in its problem, outcome, scope and acceptance criteria.
Base ADRs include context, options, decision, consequences and `decision_status`
without requiring an AI approval process.

## Add AI processes when needed

```bash
./tools/install-components.sh pull \
  --repo-root /path/to/project --preset full
```

With Flows installed, use `memory-bank/flows/routing.md` to choose the process.
Prepare a document for the selected extension, then adopt it explicitly:

```bash
memory-bank-cli document adopt --repo-root /path/to/project \
  --path memory-bank/features/FT-123/brief.md --contract feature/v1
```

Adoption validates the applicable requirements before changing state. A base
brief may need flow fields and sections first. Its stable identity, selected
contract and immutable bundle digest are recorded in the project registry;
frontmatter is a checked projection. Installing Flows alone does not activate
its gates for all feature briefs.

The [quick start](docs/quick-start.md) and [daily usage guide](docs/usage.md)
describe process-driven work with Flows. They are currently in Russian.

## Knowledge and ownership

A canonical fact has one owner. Derived documents reference that owner; code
owns implementation, while documents own intent, rationale and contracts.
Memory Bank applies First Principles Framework reasoning to make assumptions,
constraints, decisions and evidence explicit.

Project context lives in `product/`, `domain/`, `engineering/` and `ops/`.
Requirements, scenarios and decisions live in `prd/`, `use-cases/`, `features/`,
`research/`, `epics/` and `adr/`. The CLI updates template assets while preserving
project-owned content. New contract versions require an explicit document
transition; changing a bundle behind an existing ID is a conflict.

## Optional automation

Tool adapters are separate from the documentation components:

- `codex` installs the Codex agent definitions;
- `start-issue` installs issue-start instructions;
- `symphony` installs its workflow and launcher scripts;
- `bootstrap` installs the bootstrap script.

Select an adapter with repeatable `--adapter NAME` flags. Explicit `core`, `docs`
and `full` do not include these adapters automatically; `legacy` preserves them.
Runners launch agents. Flows supplies the process those agents follow.

## Template layout

This repository owns the generic payload in `template/`. The CLI installs only
selected files and removes the `template/` prefix. The component manifest is
[`template/memory-bank/components.json`](template/memory-bank/components.json).
The generated downstream `memory-bank/README.md` lists installed sections;
AGENTS routes readers only to installed components.

| Area | Purpose |
| --- | --- |
| [`dna/`](template/memory-bank/dna/README.md) | Standalone governance baseline |
| [`document-types/`](template/memory-bank/document-types/README.md) | Base document contracts |
| [`templates/`](template/memory-bank/templates/README.md) | Project-owned draft starting points |
| [`flows/`](template/memory-bank/flows/README.md) | Optional processes and versioned extensions |

The project-local `memory-bank/` in this repository is a projection of the
payload, with real files only for this project's own material. It has no
installed-template lock.

## Reference

- [Component adoption and legacy migration](docs/component-adoption.md)
- [Component wire contract](docs/component-wire-format.md)
- [Ownership and safe updates](docs/ownership.md)
- [Managed agent instructions](docs/agent-instructions.md)
- [CLI integration and source-profile validation](docs/memory-bank.md)
- [Repository development](docs/development.md)

The CLI is developed separately in
[`dapi/memory-bank-cli`](https://github.com/dapi/memory-bank-cli). This template
is available under the [Apache License 2.0](LICENSE).

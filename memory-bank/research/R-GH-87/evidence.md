---
title: "R-GH-87: Evidence Log"
doc_kind: research
doc_function: canonical
purpose: "Traceable evidence and observations collected for R-GH-87."
derived_from:
  - brief.md
  - ../../flows/research.md
status: active
audience: humans_and_agents
---

# R-GH-87: Evidence Log

## Sources

| ID | Source / provenance | Date / freshness | Collection context | Access / quality note |
| --- | --- | --- | --- | --- |
| `SRC-01` | [Issue #87](https://github.com/dapi/memory-bank/issues/87) | 2026-08-05 | GitHub issue read | Primary task source; wording is underspecified. |
| `SRC-02` | [Maintainer routing record](https://github.com/dapi/memory-bank/issues/87#issuecomment-5087962973) | 2026-07-27 | GitHub issue comment | Primary decision-owner record; verify against current sources. |
| `SRC-03` | [Migration commit `5d1d347`](https://github.com/dapi/memory-bank/commit/5d1d3473e1473ccad19f2b8a2848fe92639f1887) | 2026-07-23 | Git history inspection | Primary change history; proves files changed, not runtime behavior. |
| `SRC-04` | [Current CI workflow](https://github.com/dapi/memory-bank/blob/23ff0a3905645c6f4e37809ae7e97d43cb937a37/.github/workflows/ci.yml) | 2026-08-05 | Checked-out `main` and GitHub source | Primary executable CI configuration. |
| `SRC-05` | [memory-bank-cli#1](https://github.com/dapi/memory-bank-cli/issues/1) | closed 2026-07-22 | GitHub issue read | Primary migration-task record; its acceptance explicitly moved the CLI and removed old executable references. |
| `SRC-06` | [Successful CI run 30912292219](https://github.com/dapi/memory-bank/actions/runs/30912292219) | 2026-08-05 | GitHub Actions metadata | Current integration result; does not alone prove semantic coverage. |

## Observations

| ID | Observation | Supporting `SRC-*` | Applies to | Interpretation boundary |
| --- | --- | --- | --- | --- |
| `OBS-01` | Commit `5d1d347` removed the Go CLI under `tools/`, including both previous executable entry points and lint/doctor implementation, while updating CI and documentation. | [SRC-03](#sources) | `RQ-01`, `HYP-01` | Does not prove no repository-local scripts remain. |
| `OBS-02` | Current CI installs `memory-bank-cli` for `lint` and `doctor`, but runs `ruby tools/validate-priming-manifests*.rb` separately before those CLI commands. | [SRC-04](#sources) | `RQ-01`, `HYP-01` | Shows execution ownership, not whether the Ruby checks are useful. |
| `OBS-03` | The linked CLI task is closed and defined the move to a standalone `memory-bank-cli` with preserved `lint` and `doctor` semantics. | [SRC-05](#sources) | `RQ-01`, `HYP-01` | Does not establish a new CLI feature is impossible. |
| `OBS-04` | Current `main` CI completed successfully for the examined integration workflow. | [SRC-06](#sources) | `RQ-01` | A passing run does not independently establish completeness. |

## Collection Log

| Date | Activity | Result | Deviation / reason |
| --- | --- | --- | --- |
| 2026-08-05 | Read issue, prior maintainer record and linked CLI issue. | Decision context and migration acceptance recovered. | No external participant or privileged source was needed. |
| 2026-08-05 | Inspected migration commit and current CI workflow. | Reusable CLI and repository-specific script surfaces distinguished. | None. |

## Evidence Quality Check

- [x] Each material observation traces to one or more `SRC-*`.
- [x] Every `SRC-*` contains a clickable original-source link.
- [x] Observations are separated from source claims and analyst interpretation.
- [x] Freshness, source limitations and conflicts are recorded.

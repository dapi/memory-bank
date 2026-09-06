---
title: Project-local Memory Bank Bootstrap
doc_kind: project
doc_function: reference
purpose: Фиксирует происхождение, границы и решения начальной адаптации project-local Memory Bank.
derived_from:
  - README.md
  - ../docs/adoption.md
  - ../docs/ownership.md
status: active
audience: humans_and_agents
---

# Project-local Memory Bank Bootstrap

## Identity And Model

This is the project-local Memory Bank for the `dapi/memory-bank` source
repository. The generic upstream payload lives in
[`../template/memory-bank/`](../template/memory-bank/) and is never the place
for instantiated project artifacts.

This directory is a **projection**, not an installation. Generic documents are
symlinks into the payload, so the instance equals the current payload by
construction and there is nothing to synchronize. Only what this project owns
or overrides is a regular file.

There is deliberately no `.lock` here. A lock records drift of an installed
copy from an external template; this repository *is* the template source, and
`memory-bank-cli doctor` states the rule directly: an installed-template lock
is not expected in a template source repository, and locks belong only in
downstream repositories created through `memory-bank-cli init`.

## Consequences

- To adapt a generic document for this project, replace its symlink with a
  regular file. Git records the type change, so every override is visible.
- To return an override to the projection, delete the file and run
  [`../tools/refresh-memory-bank-projection.rb`](../tools/refresh-memory-bank-projection.rb).
- A document added to the payload needs a matching symlink here; until it is
  created, `memory-bank-cli lint --repo-root .` reports the missing target, so
  CI catches the gap.
- Editing a symlinked document writes to the payload. That is the intent for
  generic rules; if the change is project-specific, override the file first.

## Bootstrap Decisions

- This root README is the project-local entry point and explicitly identifies
  `dapi/memory-bank`; the generic template README remains unchanged.
- [`features/`](features/README.md) is the canonical destination for this
  repository's `FT-XXX/` delivery packages. No instantiated feature package
  belongs in `template/memory-bank/`.
- Project-specific product, domain, engineering and operations facts are
  adapted here as evidence becomes available; they must not be copied back to
  the generic payload.
- The managed routing block in [`../AGENTS.md`](../AGENTS.md) directs agents to
  this project-local entry point, DNA, and task routing flow. Those paths
  resolve through the projection.

## Verification Scope

Required checks are `memory-bank-cli lint --repo-root .` for project Memory
Bank navigation and `memory-bank-cli doctor --profile template` for the source
template. Both run in CI.

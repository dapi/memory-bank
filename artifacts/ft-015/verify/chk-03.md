# CHK-03: Template Compatibility and Repository Hygiene

Semantic compatibility review: pass.

`memory-bank/use-cases/README.md` now explains when operational/agentic
scenarios should become project-level `UC-*` versus feature-level `SC-*`.
`memory-bank/flows/templates/use-case/UC-XXX.md` already provides the required
shape for this guidance: goal, primary actor, trigger, preconditions, main flow,
alternate flows / exceptions, postconditions, business rules, and traceability.
No `UC-XXX` template change is required.

Command:

```sh
python3 scripts/check_memory_bank_index.py
```

Result: pass.

Relevant output:

```text
OK: no broken internal markdown links in scope.
OK: no broken frontmatter markdown dependencies in scope.
OK: no orphan markdown files in scope.
OK: all scoped markdown files are reachable from the configured entrypoints via index navigation.
OK: no documents are reachable only deeper than the configured threshold.
Index compliance:
  - OK
Result: OK
```

Command:

```sh
git diff --check
```

Result: pass. The command produced no output.

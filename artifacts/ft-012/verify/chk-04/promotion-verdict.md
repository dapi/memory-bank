# FT-012 Promotion Trigger Review

## Verdict

Pass.

Promotion triggers are present in the selector, task family flow, profile flows, task package destination, and task templates.

## Checked Surfaces

- `memory-bank/flows/workflows.md`
- `memory-bank/flows/task-flow.md`
- `memory-bank/flows/bugfix-flow.md`
- `memory-bank/flows/refactor-flow.md`
- `memory-bank/flows/templates/task/package-README.md`
- `memory-bank/flows/templates/task/bugfix.md`
- `memory-bank/flows/templates/task/refactor.md`
- `memory-bank/tasks/README.md`

## Required Promotion Triggers

- Capability or stable scenario creation routes to `feature-package`.
- API/event/schema/file format/CLI/security/financial/integration/rollout contract changes route to `feature-package` or ADR + `feature-package`.
- Solution-space reasoning, ADR dependency, C4, migration strategy, rollout/backout, or explicit failure-mode design routes out of compact task profiles.
- `risk=high/critical` routes to at least `managed-task`, often `feature-package`.
- Multiple independent delivery units route to separate tasks or `epic-package`.
- Repeated scope/design/evidence review findings require profile promotion instead of local-only fixes.

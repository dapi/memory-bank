---
title: Task Packages Index
doc_kind: task
doc_function: index
purpose: "Навигация по compact managed task packages для bugfix/chore/refactor работ, которым нужен durable context, но не нужен full feature package."
derived_from:
  - ../dna/governance.md
  - ../flows/task-flow.md
  - ../flows/bugfix-flow.md
  - ../flows/refactor-flow.md
  - ../flows/workflows.md
status: active
audience: humans_and_agents
---

# Task Packages Index

Каталог `memory-bank/tasks/` хранит compact packages вида `TASK-XXX/` для managed non-feature задач.

## Когда Использовать

Создавай `TASK-XXX/`, если одновременно верно:

- задача является `bugfix`, `chore` или `refactor`, а не новой user/system capability;
- issue/PR carrier недостаточен для durable traceability;
- full `feature-flow` создает лишнюю ceremony;
- нужны scope, non-scope, invariants, checkpoints или evidence, но не нужен полноценный feature `brief.md` + `design.md` + `implementation-plan.md`.

Не используй `TASK-XXX/`, если задача меняет API/schema/security/financial/integration/rollout contract или создает устойчивый scenario. В таком случае используй [`../features/`](../features/) по [`../flows/feature-flow.md`](../flows/feature-flow.md) либо ADR/incident owner docs.

## Комплект Документов

Durable task package ведется по [`../flows/task-flow.md`](../flows/task-flow.md) и содержит:

- `README.md` - routing/index layer, создается по [`../flows/templates/task/package-README.md`](../flows/templates/task/package-README.md);
- `bugfix.md` - если task является managed bugfix; ведется по [`../flows/bugfix-flow.md`](../flows/bugfix-flow.md);
- `refactor.md` - если task является managed refactor/chore; ведется по [`../flows/refactor-flow.md`](../flows/refactor-flow.md).

`implementation-plan.md` и `design.md` внутри `TASK-XXX/` не создаются. Если они стали нужны, task нужно повысить до `feature-package`, ADR или `epic-package`.

## Именование

- Базовый формат: `TASK-XXX/`
- Вместо `XXX` используй GitHub issue id или другой стабильный task id.
- Один package = одна managed non-feature delivery task.

## Шаблоны

- [`../flows/templates/task/package-README.md`](../flows/templates/task/package-README.md) - routing/index layer для durable `TASK-XXX/`.
- [`../flows/templates/task/bugfix.md`](../flows/templates/task/bugfix.md) - compact bugfix note.
- [`../flows/templates/task/refactor.md`](../flows/templates/task/refactor.md) - compact refactor/chore note.

## Созданные Packages

В шаблонном репозитории этот каталог может не содержать instantiated `TASK-XXX/` packages. Это нормально: destination создается как optional owner для downstream managed non-feature задач.

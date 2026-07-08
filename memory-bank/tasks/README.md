---
title: Task Packages Index
doc_kind: task
doc_function: index
purpose: "Навигация по compact managed task packages для bugfix/chore/refactor работ, которым нужен durable context, но не нужен full feature package."
derived_from:
  - ../dna/governance.md
  - ../flows/routing.md
  - ../flows/bug-fix.md
  - ../flows/small-change.md
  - ../flows/refactoring.md
status: active
audience: humans_and_agents
---

# Task Packages Index

Каталог `memory-bank/tasks/` хранит compact packages вида `TASK-XXX/` для managed non-feature задач.

## Когда Использовать

Создавай `TASK-XXX/`, если одновременно верно:

- задача является `bugfix`, `chore` или `refactor`, а не новой user/system capability;
- issue/PR carrier недостаточен для durable traceability;
- full [`Feature Flow`](../flows/feature.md) создает лишнюю ceremony;
- нужны scope, non-scope, invariants, checkpoints или evidence, но не нужен полноценный feature `brief.md` + `design.md` + `implementation-plan.md`.

Не используй `TASK-XXX/`, если задача меняет API/schema/security/financial/integration/rollout contract, требует design/execution plan или создает устойчивый scenario. В таком случае повтори [`Task Routing`](../flows/routing.md) и выбери Feature, Epic, Incident или Human Routing.

## Комплект Документов

Durable task package является optional carrier-ом внутри уже выбранного flow и содержит:

- `README.md` - routing/index layer, создается по [`../flows/templates/task/package-README.md`](../flows/templates/task/package-README.md);
- `bugfix.md` - если task следует [`Bug Fix Flow`](../flows/bug-fix.md);
- `refactor.md` - если task следует [`Refactoring Flow`](../flows/refactoring.md) или остается technical chore в [`Small Change Flow`](../flows/small-change.md).

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

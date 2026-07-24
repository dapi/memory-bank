---
title: Product Requirements Documents Index
doc_kind: prd
doc_function: index
purpose: Навигация по instantiated PRD проекта. Читать, чтобы найти существующий Product Requirements Document или завести новый по шаблону.
derived_from:
  - ../dna/governance.md
  - ../flows/templates/prd/PRD-XXX.md
status: active
audience: humans_and_agents
---

# Product Requirements Documents Index

Каталог `memory-bank/prd/` хранит instantiated PRD проекта.

PRD нужен, когда задача живет на уровне продуктовой инициативы или capability, а не одного vertical slice. Обычно PRD стоит между общим контекстом из [`../product/context.md`](../product/context.md) и downstream feature packages из [`../features/README.md`](../features/README.md).

## Граница С `product/context.md`

- [`../product/context.md`](../product/context.md) остается project-wide документом и не превращается в PRD.
- PRD наследует этот контекст через `derived_from`, но фиксирует только initiative-specific проблему, users, goals и scope.
- Если документ нужен только для того, чтобы повторить общий background проекта, оставайся на уровне `product/context.md`.

## Граница С `domain/`

- [`../domain/README.md`](../domain/README.md) владеет предметной моделью, терминами, инвариантами, состояниями, событиями и bounded contexts.
- PRD может ссылаться на `domain/`, если инициатива меняет или использует конкретные domain rules.
- PRD не должен изобретать новые domain concepts без обновления соответствующего domain-документа.

## Когда Заводить PRD

- инициатива распадается на несколько feature packages;
- нужно зафиксировать users, goals, product scope и success metrics до проектирования реализации;
- есть риск смешать продуктовые требования с architecture/design detail.

## Когда PRD Не Нужен

- задача локальна и полностью помещается в один `brief.md`;
- общий продуктовый контекст уже покрыт [`../product/context.md`](../product/context.md), а feature не требует отдельного product-layer документа.

## Naming

- Формат файла: `PRD-XXX-short-name.md`
- Вместо `XXX` используй идентификатор, принятый в проекте: initiative id, epic id или другой стабильный ключ
- Один PRD может быть upstream для нескольких feature packages

## Template

- Используй шаблон [`../flows/templates/prd/PRD-XXX.md`](../flows/templates/prd/PRD-XXX.md)

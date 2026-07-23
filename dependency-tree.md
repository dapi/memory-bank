---
title: Document Dependency Tree
purpose: Актуальная карта зависимостей документов внутри `memory-bank-template/`. Вынесена в корень репозитория как reference note, чтобы не засорять template DNA.
status: active
derived_from:
  - memory-bank-template/dna/governance.md
---

# Document Dependency Tree

Этот документ фиксирует текущую карту зависимостей документов в `memory-bank-template/`.

Важно: структура здесь не строгое дерево, а directed acyclic graph. Поле `derived_from` задает прямые upstream-зависимости, поэтому некоторые документы имеют несколько родителей. Ниже дано сжатое дерево и список дополнительных cross-edges.

Сам этот файл живет в корне репозитория и не считается частью дерева `memory-bank-template/`; он только ссылается на него как внешний reference note.

## Roots

- Навигационный root: [`README.md`](README.md). Это входная точка для чтения репозитория, но не authority root шаблона.
- Семантический root: [`memory-bank-template/dna/principles.md`](memory-bank-template/dna/principles.md). Это корень governance-дерева, от которого наследуются downstream-правила.

## Compressed Tree

```text
memory-bank-template/README.md

memory-bank-template/dna/principles.md
├── memory-bank-template/dna/README.md
├── memory-bank-template/dna/cross-references.md
└── memory-bank-template/dna/governance.md
    ├── memory-bank-template/dna/frontmatter.md
    ├── memory-bank-template/dna/lifecycle.md
    ├── memory-bank-template/product/README.md
    ├── memory-bank-template/product/context.md
    ├── memory-bank-template/product/customers.md
    ├── memory-bank-template/product/marketing.md
    ├── memory-bank-template/product/metrics.md
    ├── memory-bank-template/product/roadmap.md
    ├── memory-bank-template/product/vision.md
    ├── memory-bank-template/domain/README.md
    ├── memory-bank-template/domain/context-map.md
    ├── memory-bank-template/domain/events.md
    ├── memory-bank-template/domain/glossary.md
    ├── memory-bank-template/domain/model.md
    ├── memory-bank-template/domain/rules.md
    ├── memory-bank-template/domain/states.md
    ├── memory-bank-template/engineering/README.md
    ├── memory-bank-template/engineering/architecture.md
    ├── memory-bank-template/engineering/autonomy-boundaries.md
    ├── memory-bank-template/engineering/coding-style.md
    ├── memory-bank-template/engineering/frontend.md
    ├── memory-bank-template/engineering/git-workflow.md
    ├── memory-bank-template/engineering/testing-policy.md
    ├── memory-bank-template/features/README.md
    ├── memory-bank-template/flows/README.md
    ├── memory-bank-template/flows/bug-fix.md
    ├── memory-bank-template/flows/epic.md
    ├── memory-bank-template/flows/feature.md
    ├── memory-bank-template/flows/incident.md
    ├── memory-bank-template/flows/refactoring.md
    ├── memory-bank-template/flows/research.md
    ├── memory-bank-template/flows/routing.md
    ├── memory-bank-template/flows/small-change.md
    ├── memory-bank-template/flows/templates/README.md
    ├── memory-bank-template/flows/templates/adr/ADR-XXX.md
    ├── memory-bank-template/flows/templates/prd/PRD-XXX.md
    ├── memory-bank-template/flows/templates/research/README.md
    ├── memory-bank-template/flows/templates/research/package-README.md
    ├── memory-bank-template/flows/templates/research/brief.md
    ├── memory-bank-template/flows/templates/research/plan.md
    ├── memory-bank-template/flows/templates/research/evidence.md
    ├── memory-bank-template/flows/templates/research/synthesis.md
    ├── memory-bank-template/flows/templates/research/decision.md
    ├── memory-bank-template/flows/templates/use-case/UC-XXX.md
    ├── memory-bank-template/ops/README.md
    ├── memory-bank-template/ops/config.md
    ├── memory-bank-template/ops/development.md
    ├── memory-bank-template/ops/release.md
    ├── memory-bank-template/ops/runbooks/README.md
    ├── memory-bank-template/ops/stages.md
    ├── memory-bank-template/prd/README.md
    ├── memory-bank-template/research/README.md
    ├── memory-bank-template/use-cases/README.md
    └── memory-bank-template/adr/README.md
```

## Additional Dependency Edges

Эти связи не видны в сжатом дереве выше, но реально существуют в `derived_from` и важны для authority flow.

### DNA and Flows

- Этот файл `dependency-tree.md` зависит от [`memory-bank-template/dna/governance.md`](memory-bank-template/dna/governance.md), но сознательно живет вне `memory-bank-template/`.
- [`memory-bank-template/flows/routing.md`](memory-bank-template/flows/routing.md) зависит от governance и [`memory-bank-template/engineering/autonomy-boundaries.md`](memory-bank-template/engineering/autonomy-boundaries.md); branch flows используют его как upstream owner route selection.
- [`memory-bank-template/flows/feature.md`](memory-bank-template/flows/feature.md) зависит от governance, frontmatter и [`memory-bank-template/flows/routing.md`](memory-bank-template/flows/routing.md).
- [`memory-bank-template/flows/research.md`](memory-bank-template/flows/research.md) зависит от governance, frontmatter и [`memory-bank-template/flows/routing.md`](memory-bank-template/flows/routing.md); он определяет lifecycle research packages и их artifact ownership.
- [`memory-bank-template/flows/epic.md`](memory-bank-template/flows/epic.md) зависит от governance/frontmatter, [`memory-bank-template/flows/routing.md`](memory-bank-template/flows/routing.md) и [`memory-bank-template/flows/feature.md`](memory-bank-template/flows/feature.md); Epic Intake templates `package-README.md` и `brief.md` зависят от этого flow.
- [`memory-bank-template/flows/small-change.md`](memory-bank-template/flows/small-change.md), [`memory-bank-template/flows/bug-fix.md`](memory-bank-template/flows/bug-fix.md) и [`memory-bank-template/flows/refactoring.md`](memory-bank-template/flows/refactoring.md) зависят от routing, governance и [`memory-bank-template/engineering/testing-policy.md`](memory-bank-template/engineering/testing-policy.md).
- [`memory-bank-template/flows/incident.md`](memory-bank-template/flows/incident.md) зависит от routing, governance, testing policy и дополнительно от [`memory-bank-template/ops/runbooks/README.md`](memory-bank-template/ops/runbooks/README.md).
- [`memory-bank-template/flows/README.md`](memory-bank-template/flows/README.md) зависит от governance, всех flow-документов и [`memory-bank-template/flows/templates/README.md`](memory-bank-template/flows/templates/README.md).

### Feature-related Docs

- [`memory-bank-template/engineering/testing-policy.md`](memory-bank-template/engineering/testing-policy.md) зависит от [`memory-bank-template/dna/governance.md`](memory-bank-template/dna/governance.md) и [`memory-bank-template/flows/feature.md`](memory-bank-template/flows/feature.md).
- [`memory-bank-template/features/README.md`](memory-bank-template/features/README.md) зависит от [`memory-bank-template/dna/governance.md`](memory-bank-template/dna/governance.md) и [`memory-bank-template/flows/feature.md`](memory-bank-template/flows/feature.md).
- [`memory-bank-template/flows/templates/feature/README.md`](memory-bank-template/flows/templates/feature/README.md) зависит от [`memory-bank-template/flows/feature.md`](memory-bank-template/flows/feature.md) и [`memory-bank-template/dna/frontmatter.md`](memory-bank-template/dna/frontmatter.md).
- [`memory-bank-template/flows/templates/feature/brief.md`](memory-bank-template/flows/templates/feature/brief.md) зависит от [`memory-bank-template/flows/feature.md`](memory-bank-template/flows/feature.md), [`memory-bank-template/dna/frontmatter.md`](memory-bank-template/dna/frontmatter.md) и [`memory-bank-template/engineering/testing-policy.md`](memory-bank-template/engineering/testing-policy.md).
- [`memory-bank-template/flows/templates/feature/design.md`](memory-bank-template/flows/templates/feature/design.md) зависит от [`memory-bank-template/flows/feature.md`](memory-bank-template/flows/feature.md) и [`memory-bank-template/dna/frontmatter.md`](memory-bank-template/dna/frontmatter.md).
- [`memory-bank-template/flows/templates/feature/implementation-plan.md`](memory-bank-template/flows/templates/feature/implementation-plan.md) зависит от [`memory-bank-template/flows/feature.md`](memory-bank-template/flows/feature.md), [`memory-bank-template/dna/frontmatter.md`](memory-bank-template/dna/frontmatter.md) и [`memory-bank-template/engineering/testing-policy.md`](memory-bank-template/engineering/testing-policy.md).
- Feature-support templates [`runtime-surfaces.md`](memory-bank-template/flows/templates/feature/support/runtime-surfaces.md), [`ui-reference.md`](memory-bank-template/flows/templates/feature/support/ui-reference.md) и [`use-cases.md`](memory-bank-template/flows/templates/feature/support/use-cases.md) зависят от [`memory-bank-template/flows/feature.md`](memory-bank-template/flows/feature.md) и [`memory-bank-template/dna/frontmatter.md`](memory-bank-template/dna/frontmatter.md).

### Research-related Docs

- [`memory-bank-template/research/README.md`](memory-bank-template/research/README.md) зависит от [`memory-bank-template/dna/governance.md`](memory-bank-template/dna/governance.md) и [`memory-bank-template/flows/research.md`](memory-bank-template/flows/research.md).
- [`memory-bank-template/flows/templates/research/README.md`](memory-bank-template/flows/templates/research/README.md) зависит от [`memory-bank-template/flows/research.md`](memory-bank-template/flows/research.md) и [`memory-bank-template/dna/frontmatter.md`](memory-bank-template/dna/frontmatter.md).
- Research package templates [`package-README.md`](memory-bank-template/flows/templates/research/package-README.md), [`plan.md`](memory-bank-template/flows/templates/research/plan.md), [`evidence.md`](memory-bank-template/flows/templates/research/evidence.md), [`synthesis.md`](memory-bank-template/flows/templates/research/synthesis.md) и [`decision.md`](memory-bank-template/flows/templates/research/decision.md) зависят от [`memory-bank-template/flows/research.md`](memory-bank-template/flows/research.md); [`brief.md`](memory-bank-template/flows/templates/research/brief.md) дополнительно зависит от [`memory-bank-template/dna/frontmatter.md`](memory-bank-template/dna/frontmatter.md).

### Product And Domain Docs

- [`memory-bank-template/product/context.md`](memory-bank-template/product/context.md), [`memory-bank-template/product/vision.md`](memory-bank-template/product/vision.md), [`memory-bank-template/product/customers.md`](memory-bank-template/product/customers.md), [`memory-bank-template/product/metrics.md`](memory-bank-template/product/metrics.md), [`memory-bank-template/product/marketing.md`](memory-bank-template/product/marketing.md) и [`memory-bank-template/product/roadmap.md`](memory-bank-template/product/roadmap.md) зависят от [`memory-bank-template/dna/governance.md`](memory-bank-template/dna/governance.md) и product upstream-документов, указанных в их `derived_from`.
- [`memory-bank-template/domain/glossary.md`](memory-bank-template/domain/glossary.md), [`memory-bank-template/domain/model.md`](memory-bank-template/domain/model.md), [`memory-bank-template/domain/rules.md`](memory-bank-template/domain/rules.md), [`memory-bank-template/domain/states.md`](memory-bank-template/domain/states.md), [`memory-bank-template/domain/events.md`](memory-bank-template/domain/events.md) и [`memory-bank-template/domain/context-map.md`](memory-bank-template/domain/context-map.md) зависят от [`memory-bank-template/dna/governance.md`](memory-bank-template/dna/governance.md) и domain upstream-документов, указанных в их `derived_from`.
- [`memory-bank-template/engineering/architecture.md`](memory-bank-template/engineering/architecture.md) зависит от [`memory-bank-template/dna/governance.md`](memory-bank-template/dna/governance.md) и [`memory-bank-template/domain/context-map.md`](memory-bank-template/domain/context-map.md).
- [`memory-bank-template/engineering/frontend.md`](memory-bank-template/engineering/frontend.md) зависит от [`memory-bank-template/dna/governance.md`](memory-bank-template/dna/governance.md) и [`memory-bank-template/product/context.md`](memory-bank-template/product/context.md).
- [`memory-bank-template/flows/templates/prd/PRD-XXX.md`](memory-bank-template/flows/templates/prd/PRD-XXX.md) зависит от [`memory-bank-template/dna/governance.md`](memory-bank-template/dna/governance.md), [`memory-bank-template/dna/frontmatter.md`](memory-bank-template/dna/frontmatter.md) и [`memory-bank-template/product/context.md`](memory-bank-template/product/context.md).
- [`memory-bank-template/flows/templates/use-case/UC-XXX.md`](memory-bank-template/flows/templates/use-case/UC-XXX.md) зависит от [`memory-bank-template/dna/governance.md`](memory-bank-template/dna/governance.md), [`memory-bank-template/dna/frontmatter.md`](memory-bank-template/dna/frontmatter.md) и [`memory-bank-template/product/context.md`](memory-bank-template/product/context.md).
- [`memory-bank-template/prd/README.md`](memory-bank-template/prd/README.md) зависит от [`memory-bank-template/dna/governance.md`](memory-bank-template/dna/governance.md) и [`memory-bank-template/flows/templates/prd/PRD-XXX.md`](memory-bank-template/flows/templates/prd/PRD-XXX.md).
- [`memory-bank-template/use-cases/README.md`](memory-bank-template/use-cases/README.md) зависит от [`memory-bank-template/dna/governance.md`](memory-bank-template/dna/governance.md) и [`memory-bank-template/flows/templates/use-case/UC-XXX.md`](memory-bank-template/flows/templates/use-case/UC-XXX.md).

### ADR and Template Indexes

- [`memory-bank-template/flows/templates/adr/ADR-XXX.md`](memory-bank-template/flows/templates/adr/ADR-XXX.md) зависит от [`memory-bank-template/dna/governance.md`](memory-bank-template/dna/governance.md) и [`memory-bank-template/dna/frontmatter.md`](memory-bank-template/dna/frontmatter.md).
- [`memory-bank-template/adr/README.md`](memory-bank-template/adr/README.md) зависит от [`memory-bank-template/dna/governance.md`](memory-bank-template/dna/governance.md) и [`memory-bank-template/flows/templates/adr/ADR-XXX.md`](memory-bank-template/flows/templates/adr/ADR-XXX.md).
- [`memory-bank-template/flows/templates/README.md`](memory-bank-template/flows/templates/README.md) зависит от [`memory-bank-template/dna/governance.md`](memory-bank-template/dna/governance.md) и всех template-документов каталога `flows/templates/`.

## Reading Order

Если нужно быстро войти в шаблон сверху вниз, читай в таком порядке:

1. [`memory-bank-template/dna/principles.md`](memory-bank-template/dna/principles.md)
2. [`memory-bank-template/dna/governance.md`](memory-bank-template/dna/governance.md)
3. [`memory-bank-template/dna/frontmatter.md`](memory-bank-template/dna/frontmatter.md)
4. Product layer: [`memory-bank-template/product/README.md`](memory-bank-template/product/README.md)
5. Domain layer: [`memory-bank-template/domain/README.md`](memory-bank-template/domain/README.md)
6. Delivery flow: [`memory-bank-template/flows/README.md`](memory-bank-template/flows/README.md)
7. Engineering rules: [`memory-bank-template/engineering/README.md`](memory-bank-template/engineering/README.md)
8. Ops context: [`memory-bank-template/ops/README.md`](memory-bank-template/ops/README.md)

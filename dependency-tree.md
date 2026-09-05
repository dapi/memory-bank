---
title: Document Dependency Tree
purpose: Актуальная карта зависимостей документов внутри `template/memory-bank/`. Вынесена в корень репозитория как reference note, чтобы не засорять template DNA.
status: active
derived_from:
  - template/memory-bank/dna/governance.md
---

# Document Dependency Tree

Этот документ фиксирует текущую карту зависимостей документов в `template/memory-bank/`.

Важно: структура здесь не строгое дерево, а directed acyclic graph. Поле `derived_from` задает прямые upstream-зависимости, поэтому некоторые документы имеют несколько родителей. Ниже дано сжатое дерево и список дополнительных cross-edges.

Сам этот файл живет в корне репозитория и не считается частью дерева `template/memory-bank/`; он только ссылается на него как внешний reference note.

## Roots

- Навигационный root: [`README.md`](README.md). Это входная точка для чтения репозитория, но не authority root шаблона.
- Семантический root: [`template/memory-bank/dna/principles.md`](template/memory-bank/dna/principles.md). Это корень governance-дерева, от которого наследуются downstream-правила.

## Compressed Tree

```text
template/memory-bank/README.md

template/memory-bank/dna/principles.md
├── template/memory-bank/dna/README.md
├── template/memory-bank/dna/cross-references.md
└── template/memory-bank/dna/governance.md
    ├── template/memory-bank/dna/frontmatter.md
    ├── template/memory-bank/dna/lifecycle.md
    ├── template/memory-bank/product/README.md
    ├── template/memory-bank/product/context.md
    ├── template/memory-bank/product/customers.md
    ├── template/memory-bank/product/marketing.md
    ├── template/memory-bank/product/metrics.md
    ├── template/memory-bank/product/roadmap.md
    ├── template/memory-bank/product/vision.md
    ├── template/memory-bank/domain/README.md
    ├── template/memory-bank/domain/context-map.md
    ├── template/memory-bank/domain/events.md
    ├── template/memory-bank/domain/glossary.md
    ├── template/memory-bank/domain/model.md
    ├── template/memory-bank/domain/rules.md
    ├── template/memory-bank/domain/states.md
    ├── template/memory-bank/engineering/README.md
    ├── template/memory-bank/engineering/architecture.md
    ├── template/memory-bank/engineering/coding-style.md
    ├── template/memory-bank/engineering/frontend.md
    ├── template/memory-bank/engineering/git-workflow.md
    ├── template/memory-bank/engineering/testing-conventions.md
    ├── template/memory-bank/features/README.md
    ├── template/memory-bank/flows/README.md
    ├── template/memory-bank/flows/autonomy-boundaries.md
    ├── template/memory-bank/flows/validation-profiles.md
    ├── template/memory-bank/flows/testing-policy.md
    ├── template/memory-bank/flows/behavior-specification.md
    ├── template/memory-bank/flows/bug-fix.md
    ├── template/memory-bank/flows/epic.md
    ├── template/memory-bank/flows/feature.md
    ├── template/memory-bank/flows/feature-requirements.md
    ├── template/memory-bank/flows/incident.md
    ├── template/memory-bank/flows/refactoring.md
    ├── template/memory-bank/flows/research.md
    ├── template/memory-bank/flows/routing.md
    ├── template/memory-bank/flows/small-change.md
    ├── template/memory-bank/flows/templates/README.md
    ├── template/memory-bank/flows/templates/adr/ADR-XXX.md
    ├── template/memory-bank/flows/templates/prd/PRD-XXX.md
    ├── template/memory-bank/flows/templates/research/README.md
    ├── template/memory-bank/flows/templates/research/package-README.md
    ├── template/memory-bank/flows/templates/research/brief.md
    ├── template/memory-bank/flows/templates/research/plan.md
    ├── template/memory-bank/flows/templates/research/evidence.md
    ├── template/memory-bank/flows/templates/research/synthesis.md
    ├── template/memory-bank/flows/templates/research/decision.md
    ├── template/memory-bank/flows/templates/use-case/UC-XXX.md
    ├── template/memory-bank/ops/README.md
    ├── template/memory-bank/ops/config.md
    ├── template/memory-bank/ops/development.md
    ├── template/memory-bank/ops/release.md
    ├── template/memory-bank/ops/runbooks/README.md
    ├── template/memory-bank/ops/stages.md
    ├── template/memory-bank/prd/README.md
    ├── template/memory-bank/research/README.md
    ├── template/memory-bank/use-cases/README.md
    └── template/memory-bank/adr/README.md
```

## Additional Dependency Edges

Эти связи не видны в сжатом дереве выше, но реально существуют в `derived_from` и важны для authority flow.

### DNA and Flows

- Этот файл `dependency-tree.md` зависит от [`template/memory-bank/dna/governance.md`](template/memory-bank/dna/governance.md), но сознательно живет вне `template/memory-bank/`.
- [`template/memory-bank/flows/routing.md`](template/memory-bank/flows/routing.md) зависит от governance и [`template/memory-bank/flows/autonomy-boundaries.md`](template/memory-bank/flows/autonomy-boundaries.md); branch flows используют его как upstream owner route selection.
- [`template/memory-bank/flows/behavior-specification.md`](template/memory-bank/flows/behavior-specification.md) зависит от governance и principles; он владеет BDD practice без создания отдельного delivery route.
- [`template/memory-bank/flows/feature-requirements.md`](template/memory-bank/flows/feature-requirements.md) зависит от governance, frontmatter и [`template/memory-bank/flows/validation-profiles.md`](template/memory-bank/flows/validation-profiles.md); он владеет requirement taxonomy и двусторонней traceability policy.
- [`template/memory-bank/flows/feature.md`](template/memory-bank/flows/feature.md) зависит от governance, frontmatter, [`template/memory-bank/flows/routing.md`](template/memory-bank/flows/routing.md), [`template/memory-bank/flows/feature-requirements.md`](template/memory-bank/flows/feature-requirements.md), [`template/memory-bank/flows/behavior-specification.md`](template/memory-bank/flows/behavior-specification.md), validation profiles и autonomy boundaries.
- [`template/memory-bank/flows/research.md`](template/memory-bank/flows/research.md) зависит от governance, frontmatter и [`template/memory-bank/flows/routing.md`](template/memory-bank/flows/routing.md); он определяет lifecycle research packages и их artifact ownership.
- [`template/memory-bank/flows/epic.md`](template/memory-bank/flows/epic.md) зависит от governance/frontmatter, [`template/memory-bank/flows/routing.md`](template/memory-bank/flows/routing.md) и [`template/memory-bank/flows/feature.md`](template/memory-bank/flows/feature.md); Epic Intake templates `package-README.md` и `brief.md` зависят от этого flow.
- [`template/memory-bank/flows/small-change.md`](template/memory-bank/flows/small-change.md), [`template/memory-bank/flows/bug-fix.md`](template/memory-bank/flows/bug-fix.md) и [`template/memory-bank/flows/refactoring.md`](template/memory-bank/flows/refactoring.md) зависят от routing, governance и [`template/memory-bank/flows/testing-policy.md`](template/memory-bank/flows/testing-policy.md).
- [`template/memory-bank/flows/incident.md`](template/memory-bank/flows/incident.md) зависит от routing, governance, testing policy и дополнительно от [`template/memory-bank/ops/runbooks/README.md`](template/memory-bank/ops/runbooks/README.md).
- [`template/memory-bank/flows/README.md`](template/memory-bank/flows/README.md) зависит от governance, всех flow-документов и [`template/memory-bank/flows/templates/README.md`](template/memory-bank/flows/templates/README.md).

### Feature-related Docs

- [`template/memory-bank/flows/feature-artifact-catalog.md`](template/memory-bank/flows/feature-artifact-catalog.md) зависит от [`template/memory-bank/flows/feature.md`](template/memory-bank/flows/feature.md) и [`template/memory-bank/flows/feature-requirements.md`](template/memory-bank/flows/feature-requirements.md).
- [`template/memory-bank/flows/testing-policy.md`](template/memory-bank/flows/testing-policy.md) зависит от [`template/memory-bank/dna/governance.md`](template/memory-bank/dna/governance.md), [`template/memory-bank/flows/behavior-specification.md`](template/memory-bank/flows/behavior-specification.md), [`template/memory-bank/flows/feature.md`](template/memory-bank/flows/feature.md), [`template/memory-bank/flows/feature-requirements.md`](template/memory-bank/flows/feature-requirements.md) и validation profiles.
- [`template/memory-bank/features/README.md`](template/memory-bank/features/README.md) зависит от [`template/memory-bank/dna/governance.md`](template/memory-bank/dna/governance.md) и [`template/memory-bank/flows/feature.md`](template/memory-bank/flows/feature.md).
- [`template/memory-bank/flows/templates/feature/README.md`](template/memory-bank/flows/templates/feature/README.md) зависит от [`template/memory-bank/flows/feature.md`](template/memory-bank/flows/feature.md) и [`template/memory-bank/dna/frontmatter.md`](template/memory-bank/dna/frontmatter.md).
- [`template/memory-bank/flows/templates/feature/brief.md`](template/memory-bank/flows/templates/feature/brief.md) зависит от [`template/memory-bank/flows/feature.md`](template/memory-bank/flows/feature.md), [`template/memory-bank/flows/feature-requirements.md`](template/memory-bank/flows/feature-requirements.md), [`template/memory-bank/flows/behavior-specification.md`](template/memory-bank/flows/behavior-specification.md), [`template/memory-bank/dna/frontmatter.md`](template/memory-bank/dna/frontmatter.md) и [`template/memory-bank/flows/testing-policy.md`](template/memory-bank/flows/testing-policy.md).
- [`template/memory-bank/flows/templates/feature/design.md`](template/memory-bank/flows/templates/feature/design.md) зависит от [`template/memory-bank/flows/feature.md`](template/memory-bank/flows/feature.md), [`template/memory-bank/flows/feature-requirements.md`](template/memory-bank/flows/feature-requirements.md), [`template/memory-bank/flows/feature-artifact-catalog.md`](template/memory-bank/flows/feature-artifact-catalog.md) и [`template/memory-bank/dna/frontmatter.md`](template/memory-bank/dna/frontmatter.md).
- [`template/memory-bank/flows/templates/feature/implementation-plan.md`](template/memory-bank/flows/templates/feature/implementation-plan.md) зависит от [`template/memory-bank/flows/feature.md`](template/memory-bank/flows/feature.md), [`template/memory-bank/flows/feature-requirements.md`](template/memory-bank/flows/feature-requirements.md), [`template/memory-bank/dna/frontmatter.md`](template/memory-bank/dna/frontmatter.md) и [`template/memory-bank/flows/testing-policy.md`](template/memory-bank/flows/testing-policy.md).
- Feature-support templates [`runtime-surfaces.md`](template/memory-bank/flows/templates/feature/support/runtime-surfaces.md), [`ui-reference.md`](template/memory-bank/flows/templates/feature/support/ui-reference.md) и [`use-cases.md`](template/memory-bank/flows/templates/feature/support/use-cases.md) зависят от [`template/memory-bank/flows/feature.md`](template/memory-bank/flows/feature.md), [`template/memory-bank/dna/frontmatter.md`](template/memory-bank/dna/frontmatter.md) и своих direct companions; `use-cases.md` также зависит от [`template/memory-bank/flows/behavior-specification.md`](template/memory-bank/flows/behavior-specification.md).

### Research-related Docs

- [`template/memory-bank/research/README.md`](template/memory-bank/research/README.md) зависит от [`template/memory-bank/dna/governance.md`](template/memory-bank/dna/governance.md) и [`template/memory-bank/flows/research.md`](template/memory-bank/flows/research.md).
- [`template/memory-bank/flows/templates/research/README.md`](template/memory-bank/flows/templates/research/README.md) зависит от [`template/memory-bank/flows/research.md`](template/memory-bank/flows/research.md) и [`template/memory-bank/dna/frontmatter.md`](template/memory-bank/dna/frontmatter.md).
- Research package templates [`package-README.md`](template/memory-bank/flows/templates/research/package-README.md), [`plan.md`](template/memory-bank/flows/templates/research/plan.md), [`evidence.md`](template/memory-bank/flows/templates/research/evidence.md), [`synthesis.md`](template/memory-bank/flows/templates/research/synthesis.md) и [`decision.md`](template/memory-bank/flows/templates/research/decision.md) зависят от [`template/memory-bank/flows/research.md`](template/memory-bank/flows/research.md); [`brief.md`](template/memory-bank/flows/templates/research/brief.md) дополнительно зависит от [`template/memory-bank/dna/frontmatter.md`](template/memory-bank/dna/frontmatter.md).

### Product And Domain Docs

- [`template/memory-bank/product/context.md`](template/memory-bank/product/context.md), [`template/memory-bank/product/vision.md`](template/memory-bank/product/vision.md), [`template/memory-bank/product/customers.md`](template/memory-bank/product/customers.md), [`template/memory-bank/product/metrics.md`](template/memory-bank/product/metrics.md), [`template/memory-bank/product/marketing.md`](template/memory-bank/product/marketing.md) и [`template/memory-bank/product/roadmap.md`](template/memory-bank/product/roadmap.md) зависят от [`template/memory-bank/dna/governance.md`](template/memory-bank/dna/governance.md) и product upstream-документов, указанных в их `derived_from`.
- [`template/memory-bank/domain/glossary.md`](template/memory-bank/domain/glossary.md), [`template/memory-bank/domain/model.md`](template/memory-bank/domain/model.md), [`template/memory-bank/domain/rules.md`](template/memory-bank/domain/rules.md), [`template/memory-bank/domain/states.md`](template/memory-bank/domain/states.md), [`template/memory-bank/domain/events.md`](template/memory-bank/domain/events.md) и [`template/memory-bank/domain/context-map.md`](template/memory-bank/domain/context-map.md) зависят от [`template/memory-bank/dna/governance.md`](template/memory-bank/dna/governance.md) и domain upstream-документов, указанных в их `derived_from`.
- [`template/memory-bank/engineering/architecture.md`](template/memory-bank/engineering/architecture.md) зависит от [`template/memory-bank/dna/governance.md`](template/memory-bank/dna/governance.md) и [`template/memory-bank/domain/context-map.md`](template/memory-bank/domain/context-map.md).
- [`template/memory-bank/engineering/frontend.md`](template/memory-bank/engineering/frontend.md) зависит от [`template/memory-bank/dna/governance.md`](template/memory-bank/dna/governance.md) и [`template/memory-bank/product/context.md`](template/memory-bank/product/context.md).
- [`template/memory-bank/engineering/testing-conventions.md`](template/memory-bank/engineering/testing-conventions.md) зависит от [`template/memory-bank/dna/governance.md`](template/memory-bank/dna/governance.md), [`template/memory-bank/flows/testing-policy.md`](template/memory-bank/flows/testing-policy.md) и [`template/memory-bank/ops/development.md`](template/memory-bank/ops/development.md); project-specific стек исполняет generic policy, не может её ослаблять и не владеет списком локальных команд.
- [`template/memory-bank/flows/templates/prd/PRD-XXX.md`](template/memory-bank/flows/templates/prd/PRD-XXX.md) зависит от [`template/memory-bank/dna/governance.md`](template/memory-bank/dna/governance.md), [`template/memory-bank/dna/frontmatter.md`](template/memory-bank/dna/frontmatter.md) и [`template/memory-bank/product/context.md`](template/memory-bank/product/context.md).
- [`template/memory-bank/flows/templates/use-case/UC-XXX.md`](template/memory-bank/flows/templates/use-case/UC-XXX.md) зависит от [`template/memory-bank/dna/governance.md`](template/memory-bank/dna/governance.md), [`template/memory-bank/dna/frontmatter.md`](template/memory-bank/dna/frontmatter.md) и [`template/memory-bank/product/context.md`](template/memory-bank/product/context.md).
- [`template/memory-bank/prd/README.md`](template/memory-bank/prd/README.md) зависит от [`template/memory-bank/dna/governance.md`](template/memory-bank/dna/governance.md) и [`template/memory-bank/flows/templates/prd/PRD-XXX.md`](template/memory-bank/flows/templates/prd/PRD-XXX.md).
- [`template/memory-bank/use-cases/README.md`](template/memory-bank/use-cases/README.md) зависит от [`template/memory-bank/dna/governance.md`](template/memory-bank/dna/governance.md) и [`template/memory-bank/flows/templates/use-case/UC-XXX.md`](template/memory-bank/flows/templates/use-case/UC-XXX.md).

### ADR and Template Indexes

- [`template/memory-bank/flows/templates/adr/ADR-XXX.md`](template/memory-bank/flows/templates/adr/ADR-XXX.md) зависит от [`template/memory-bank/dna/governance.md`](template/memory-bank/dna/governance.md) и [`template/memory-bank/dna/frontmatter.md`](template/memory-bank/dna/frontmatter.md).
- [`template/memory-bank/adr/README.md`](template/memory-bank/adr/README.md) зависит от [`template/memory-bank/dna/governance.md`](template/memory-bank/dna/governance.md) и [`template/memory-bank/flows/templates/adr/ADR-XXX.md`](template/memory-bank/flows/templates/adr/ADR-XXX.md).
- [`template/memory-bank/flows/templates/README.md`](template/memory-bank/flows/templates/README.md) зависит от [`template/memory-bank/dna/governance.md`](template/memory-bank/dna/governance.md) и всех template-документов каталога `flows/templates/`.

## Reading Order

Если нужно быстро войти в шаблон сверху вниз, читай в таком порядке:

1. [`template/memory-bank/dna/principles.md`](template/memory-bank/dna/principles.md)
2. [`template/memory-bank/dna/governance.md`](template/memory-bank/dna/governance.md)
3. [`template/memory-bank/dna/frontmatter.md`](template/memory-bank/dna/frontmatter.md)
4. Product layer: [`template/memory-bank/product/README.md`](template/memory-bank/product/README.md)
5. Domain layer: [`template/memory-bank/domain/README.md`](template/memory-bank/domain/README.md)
6. Delivery flow: [`template/memory-bank/flows/README.md`](template/memory-bank/flows/README.md)
7. Engineering rules: [`template/memory-bank/engineering/README.md`](template/memory-bank/engineering/README.md)
8. Ops context: [`template/memory-bank/ops/README.md`](template/memory-bank/ops/README.md)

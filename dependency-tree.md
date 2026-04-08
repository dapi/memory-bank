---
title: Document Dependency Tree
purpose: Актуальная карта зависимостей документов внутри `memory-bank/`. Вынесена в корень репозитория как reference note, чтобы не засорять template DNA.
status: active
derived_from:
  - memory-bank/dna/governance.md
---

# Document Dependency Tree

Этот документ фиксирует текущую карту зависимостей документов в `memory-bank/`.

Важно: структура здесь не строгое дерево, а directed acyclic graph. Поле `derived_from` задает прямые upstream-зависимости, поэтому некоторые документы имеют несколько родителей. Ниже дано сжатое дерево и список дополнительных cross-edges.

Сам этот файл живет в корне репозитория и не считается частью дерева `memory-bank/`; он только ссылается на него как внешний reference note.

## Roots

- Навигационный root: [`README.md`](README.md). Это входная точка для чтения репозитория, но не authority root шаблона.
- Семантический root: [`memory-bank/dna/principles.md`](memory-bank/dna/principles.md). Это корень governance-дерева, от которого наследуются downstream-правила.

## Compressed Tree

```text
memory-bank/README.md

memory-bank/dna/principles.md
├── memory-bank/dna/README.md
├── memory-bank/dna/cross-references.md
└── memory-bank/dna/governance.md
    ├── memory-bank/dna/frontmatter.md
    ├── memory-bank/dna/lifecycle.md
    ├── memory-bank/domain/README.md
    ├── memory-bank/domain/architecture.md
    ├── memory-bank/domain/frontend.md
    ├── memory-bank/domain/problem.md
    ├── memory-bank/engineering/README.md
    ├── memory-bank/engineering/autonomy-boundaries.md
    ├── memory-bank/engineering/coding-style.md
    ├── memory-bank/engineering/git-workflow.md
    ├── memory-bank/engineering/testing-policy.md
    ├── memory-bank/features/README.md
    ├── memory-bank/flows/README.md
    ├── memory-bank/flows/feature-flow.md
    ├── memory-bank/flows/templates/README.md
    ├── memory-bank/flows/templates/adr/ADR-XXX.md
    ├── memory-bank/flows/templates/prd/PRD-XXX.md
    ├── memory-bank/flows/templates/use-case/UC-XXX.md
    ├── memory-bank/flows/workflows.md
    ├── memory-bank/ops/README.md
    ├── memory-bank/ops/config.md
    ├── memory-bank/ops/development.md
    ├── memory-bank/ops/release.md
    ├── memory-bank/ops/runbooks/README.md
    ├── memory-bank/ops/stages.md
    ├── memory-bank/prd/README.md
    ├── memory-bank/use-cases/README.md
    └── memory-bank/adr/README.md
```

## Additional Dependency Edges

Эти связи не видны в сжатом дереве выше, но реально существуют в `derived_from` и важны для authority flow.

### DNA and Flows

- Этот файл `dependency-tree.md` зависит от [`memory-bank/dna/governance.md`](memory-bank/dna/governance.md), но сознательно живет вне `memory-bank/`.
- [`memory-bank/flows/feature-flow.md`](memory-bank/flows/feature-flow.md) зависит и от [`memory-bank/dna/governance.md`](memory-bank/dna/governance.md), и от [`memory-bank/dna/frontmatter.md`](memory-bank/dna/frontmatter.md).
- [`memory-bank/flows/workflows.md`](memory-bank/flows/workflows.md) зависит от [`memory-bank/dna/governance.md`](memory-bank/dna/governance.md) и [`memory-bank/flows/feature-flow.md`](memory-bank/flows/feature-flow.md).
- [`memory-bank/flows/README.md`](memory-bank/flows/README.md) зависит сразу от [`memory-bank/dna/governance.md`](memory-bank/dna/governance.md), [`memory-bank/flows/feature-flow.md`](memory-bank/flows/feature-flow.md), [`memory-bank/flows/workflows.md`](memory-bank/flows/workflows.md) и [`memory-bank/flows/templates/README.md`](memory-bank/flows/templates/README.md).

### Feature-related Docs

- [`memory-bank/engineering/testing-policy.md`](memory-bank/engineering/testing-policy.md) зависит от [`memory-bank/dna/governance.md`](memory-bank/dna/governance.md) и [`memory-bank/flows/feature-flow.md`](memory-bank/flows/feature-flow.md).
- [`memory-bank/features/README.md`](memory-bank/features/README.md) зависит от [`memory-bank/dna/governance.md`](memory-bank/dna/governance.md) и [`memory-bank/flows/feature-flow.md`](memory-bank/flows/feature-flow.md).
- [`memory-bank/flows/templates/feature/README.md`](memory-bank/flows/templates/feature/README.md) зависит от [`memory-bank/flows/feature-flow.md`](memory-bank/flows/feature-flow.md) и [`memory-bank/dna/frontmatter.md`](memory-bank/dna/frontmatter.md).
- [`memory-bank/flows/templates/feature/implementation-plan.md`](memory-bank/flows/templates/feature/implementation-plan.md) зависит от [`memory-bank/flows/feature-flow.md`](memory-bank/flows/feature-flow.md), [`memory-bank/dna/frontmatter.md`](memory-bank/dna/frontmatter.md) и [`memory-bank/engineering/testing-policy.md`](memory-bank/engineering/testing-policy.md).
- [`memory-bank/flows/templates/feature/short.md`](memory-bank/flows/templates/feature/short.md) и [`memory-bank/flows/templates/feature/large.md`](memory-bank/flows/templates/feature/large.md) зависят от [`memory-bank/flows/feature-flow.md`](memory-bank/flows/feature-flow.md), [`memory-bank/dna/frontmatter.md`](memory-bank/dna/frontmatter.md) и [`memory-bank/engineering/testing-policy.md`](memory-bank/engineering/testing-policy.md).

### Product-layer Docs

- [`memory-bank/flows/templates/prd/PRD-XXX.md`](memory-bank/flows/templates/prd/PRD-XXX.md) зависит от [`memory-bank/dna/governance.md`](memory-bank/dna/governance.md), [`memory-bank/dna/frontmatter.md`](memory-bank/dna/frontmatter.md) и [`memory-bank/domain/problem.md`](memory-bank/domain/problem.md).
- [`memory-bank/flows/templates/use-case/UC-XXX.md`](memory-bank/flows/templates/use-case/UC-XXX.md) зависит от [`memory-bank/dna/governance.md`](memory-bank/dna/governance.md), [`memory-bank/dna/frontmatter.md`](memory-bank/dna/frontmatter.md) и [`memory-bank/domain/problem.md`](memory-bank/domain/problem.md).
- [`memory-bank/prd/README.md`](memory-bank/prd/README.md) зависит от [`memory-bank/dna/governance.md`](memory-bank/dna/governance.md) и [`memory-bank/flows/templates/prd/PRD-XXX.md`](memory-bank/flows/templates/prd/PRD-XXX.md).
- [`memory-bank/use-cases/README.md`](memory-bank/use-cases/README.md) зависит от [`memory-bank/dna/governance.md`](memory-bank/dna/governance.md) и [`memory-bank/flows/templates/use-case/UC-XXX.md`](memory-bank/flows/templates/use-case/UC-XXX.md).

### ADR and Template Indexes

- [`memory-bank/flows/templates/adr/ADR-XXX.md`](memory-bank/flows/templates/adr/ADR-XXX.md) зависит от [`memory-bank/dna/governance.md`](memory-bank/dna/governance.md) и [`memory-bank/dna/frontmatter.md`](memory-bank/dna/frontmatter.md).
- [`memory-bank/adr/README.md`](memory-bank/adr/README.md) зависит от [`memory-bank/dna/governance.md`](memory-bank/dna/governance.md) и [`memory-bank/flows/templates/adr/ADR-XXX.md`](memory-bank/flows/templates/adr/ADR-XXX.md).
- [`memory-bank/flows/templates/README.md`](memory-bank/flows/templates/README.md) зависит от [`memory-bank/dna/governance.md`](memory-bank/dna/governance.md) и всех template-документов каталога `flows/templates/`.

## Reading Order

Если нужно быстро войти в шаблон сверху вниз, читай в таком порядке:

1. [`memory-bank/dna/principles.md`](memory-bank/dna/principles.md)
2. [`memory-bank/dna/governance.md`](memory-bank/dna/governance.md)
3. [`memory-bank/dna/frontmatter.md`](memory-bank/dna/frontmatter.md)
4. Product context: [`memory-bank/domain/README.md`](memory-bank/domain/README.md)
5. Delivery flow: [`memory-bank/flows/README.md`](memory-bank/flows/README.md)
6. Engineering rules: [`memory-bank/engineering/README.md`](memory-bank/engineering/README.md)
7. Ops context: [`memory-bank/ops/README.md`](memory-bank/ops/README.md)

---
title: "FT-XXX: Runtime Surfaces Template"
doc_kind: feature-support
doc_function: template
purpose: Governed wrapper-шаблон optional `runtime-surfaces.md`. Читать, когда feature needs grounding по current runtime surfaces, semantic mappings, context variants, fallback/error paths или adjacent boundaries.
derived_from:
  - ../../../feature-flow.md
  - ../../../../dna/frontmatter.md
status: active
audience: humans_and_agents
template_for: feature-support
template_target_path: ../../../../features/FT-XXX/runtime-surfaces.md
canonical_for:
  - feature_support_template_runtime_surfaces
---

# FT-XXX: Runtime Surfaces Template

Этот файл описывает wrapper-template. Инстанцируемый `runtime-surfaces.md` живет внутри feature package как optional support/reference doc.

## Wrapper Notes

Создавай `runtime-surfaces.md`, если без отдельного grounding сложно понять current entrypoints, concrete surfaces, semantic mappings, context availability или fallback/error behavior.

`runtime-surfaces.md` не владеет requirements, selected design, acceptance criteria, checks, evidence contract или implementation sequence. Если во время runtime mapping меняется scope или selected design, обнови sibling `feature.md`, `solution.md` или ADR.

## Instantiated Frontmatter

```yaml
title: "FT-XXX: Runtime Surfaces"
doc_kind: feature-support
doc_function: reference
purpose: "Grounding reference для FT-XXX. Фиксирует current runtime surfaces, semantic mapping, adjacent boundaries и context notes без переопределения canonical problem или solution facts."
derived_from:
  - feature.md
  - solution.md
status: draft
audience: humans_and_agents
must_not_define:
  - ft_xxx_scope
  - ft_xxx_selected_design
  - ft_xxx_acceptance_criteria
  - implementation_sequence
```

## Instantiated Body

```markdown
# FT-XXX: Runtime Surfaces

## Role

Этот документ фиксирует grounding. Canonical owners:

- `feature.md` владеет problem space и verify inventory.
- `solution.md` владеет selected design, target architecture и contracts.
- `implementation-plan.md` владеет execution sequencing.

## Current Surface Inventory

| Surface ID | Current entrypoint / trigger | Current concrete surface | Current guaranteed context | Notes |
| --- | --- | --- | --- | --- |
| `SURF-01` | Как surface достигается сейчас | Route / handler / job / screen / process | Какие данные гарантированно доступны | Что важно для feature |

## Adjacent Out-of-Scope Surfaces

| Surface | Why adjacent | Why excluded |
| --- | --- | --- |
| `adjacent-surface` | Почему рядом с фичей | Какой `NS-*`, ADR или solution boundary исключает его |

## Semantic Mapping

| Mapping ID | Semantic unit | Current reachable surfaces | Why semantic unit is stable |
| --- | --- | --- | --- |
| `MAP-01` | Stable business/runtime unit | `SURF-01`, `SURF-02` | Почему нельзя привязаться только к конкретному route/file/template |

## Target Mapping Reference

| Semantic unit | To-be owner / responsibility | Covered surfaces | Related solution refs |
| --- | --- | --- | --- |
| `semantic-unit` | Кто владеет unit после changes | `SURF-01` | `SOL-01`, `C4-L2-01`, `CTR-01` |

## Context Matrix

| Surface / semantic unit | Always available | Optional | Must not assume | Related refs |
| --- | --- | --- | --- | --- |
| `SURF-01` | Какие данные всегда есть | Какие данные могут быть | Что нельзя считать гарантированным | `CTR-01` |

## Resolution / Decision Table

| Condition | Decision | Result | Observability expectation | Related refs |
| --- | --- | --- | --- | --- |
| Какой state/mode/input | Что выбирает runtime | Что происходит | Как это видно в logs/UI/evidence | `SOL-01`, `FM-01` |

## Notes For Implementation Plan

- Какие paths/modules обязательно учесть в `implementation-plan.md`.
- Какие ambiguity должны стать `OQ-*`.
- Какие stop conditions должны попасть в plan, если mapping не подтверждается.
```

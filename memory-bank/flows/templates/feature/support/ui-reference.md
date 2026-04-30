---
title: "FT-XXX: UI Reference Template"
doc_kind: feature-support
doc_function: template
purpose: Governed wrapper-шаблон optional `ui-reference/README.md`. Читать, когда feature changes interface, navigation, screen states, editor/preview flows, copy/state semantics или interaction model.
derived_from:
  - ../../../feature-flow.md
  - ../../../../dna/frontmatter.md
status: active
audience: humans_and_agents
template_for: feature-support
template_target_path: ../../../../features/FT-XXX/ui-reference/README.md
canonical_for:
  - feature_support_template_ui_reference
---

# FT-XXX: UI Reference Template

Этот файл описывает wrapper-template. Инстанцируемый `ui-reference/README.md` живет внутри feature package как optional support/reference doc.

## Wrapper Notes

Создавай `ui-reference/README.md`, если feature меняет interface. Документ generic: он не должен тянуть project-specific interface conventions в reusable template. В instantiated project можно ссылаться на локальный design system, но generic template фиксирует только структуру interface reference.

Для interface changes нужны mockups. Default format — Markdown mockups в `ui-reference/mockups/*.md`. Допустимы изображения, design-tool links или другие artifacts, если они versionable / linkable и доступны reviewers.

`ui-reference/README.md` не владеет requirements, selected architecture, acceptance inventory или implementation sequence.

## Instantiated Frontmatter

```yaml
title: "FT-XXX: UI Reference"
doc_kind: feature-support
doc_function: reference
purpose: "Interface reference для FT-XXX. Фиксирует screen map, interaction states, mockups и UI traceability без переопределения canonical feature или solution facts."
derived_from:
  - ../feature.md
  - ../solution.md
status: draft
audience: humans_and_agents
must_not_define:
  - ft_xxx_scope
  - ft_xxx_selected_architecture
  - ft_xxx_acceptance_criteria
  - implementation_sequence
```

## Instantiated Body

```markdown
# FT-XXX: UI Reference

## Role

Этот документ раскрывает interface expectations для implementation и review. Canonical owners:

- `feature.md` владеет requirements и acceptance.
- `solution.md` владеет selected design и contracts.
- `implementation-plan.md` владеет execution sequencing.

## Interface Scope

| UI ID | Surface / screen | User role | Purpose | Related refs |
| --- | --- | --- | --- | --- |
| `UI-01` | Какой экран или interface surface меняется | Кто им пользуется | Зачем нужен экран | `REQ-01`, `SOL-01` |

## Screen Map

| UI ID | Screen / state | Entry point | Primary actions | Exit / next state |
| --- | --- | --- | --- | --- |
| `UI-01` | Screen name | Откуда пользователь приходит | Основные действия | Куда пользователь уходит |

## Interaction States

| UI ID | State | What user sees | System behavior | Related refs |
| --- | --- | --- | --- | --- |
| `UI-01` | loading / empty / success / error / disabled | Что показываем | Как система ведет себя | `SC-01`, `FM-01` |

## Mockups

Mockups обязательны для interface changes. Markdown — default, но можно использовать другой формат, если artifact linkable.

| Mockup | Format | Covers | Notes |
| --- | --- | --- | --- |
| [`mockups/screen-name.md`](mockups/screen-name.md) | markdown | `UI-01`, `SC-01` | Low-fidelity sketch |

## Copy And State Semantics

| UI element | Text / label intent | State semantics | Related refs |
| --- | --- | --- | --- |
| `control-or-message` | Что должен понять пользователь | Какой state не должен быть скрыт | `REQ-01`, `CTR-01` |

## UI Traceability

| UI ID / element | Supports | Checks / evidence |
| --- | --- | --- |
| `UI-01` | `REQ-01`, `SC-01` | `CHK-01`, `EVID-01` |

## Out Of Scope For This Doc

- product requirements;
- selected architecture;
- file-level touchpoints;
- implementation sequence;
- project-specific UI framework rules unless linked from local project docs.
```

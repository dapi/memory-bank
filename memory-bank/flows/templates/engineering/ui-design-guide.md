---
title: Project UI Design Guide Template
doc_kind: engineering
doc_function: template
purpose: Governed wrapper-шаблон optional project-level UI reference. Читать, когда существующий UI kit, helper APIs, examples и screenshots переиспользуются несколькими features.
derived_from:
  - ../../../dna/governance.md
  - ../../../engineering/frontend.md
status: active
audience: humans_and_agents
template_for: engineering
template_target_path: ../../../engineering/ui-design-guide/README.md
canonical_for:
  - project_ui_design_guide_template
---

# Project UI Design Guide Template

Этот файл описывает wrapper-template. Инстанцируемый `ui-design-guide/README.md` — optional project-level reference внутри `engineering/`, а не обязательный placeholder generic memory-bank.

## Wrapper Notes

Создавай guide, только когда concrete UI conventions используются несколькими features и отдельный reference сокращает повторный поиск по коду. Если material помещается в [`../../../engineering/frontend.md`](../../../engineering/frontend.md), отдельный файл не нужен.

Guide каталогизирует существующие components, helper APIs, examples, screenshots и source paths. Нормативные project-wide engineering rules сначала фиксируй в `frontend.md`; product requirements и business semantics — у product/domain owners; interface change одной feature — в feature-local [`ui-reference/README.md`](../feature/support/ui-reference.md).

После инстанцирования добавь аннотированную ссылку на `ui-design-guide/README.md` в `../../../engineering/README.md`, чтобы guide был достижим из project index. Ссылку на wrapper-template можно сохранить как creation route, но она не заменяет route к instantiated guide. При удалении guide удали и ссылку из index.

Удаляй неиспользуемые sections при инстанцировании. Не оставляй placeholder rows и не создавай пустые дочерние файлы или screenshot-каталоги.

## Instantiated Frontmatter

```yaml
---
title: UI Design Guide
doc_kind: engineering
doc_function: reference
purpose: Project-level reference по существующему UI kit: concrete components, helpers, examples, screenshots и source paths.
derived_from:
  - ../frontend.md
status: active
audience: humans_and_agents
must_not_define:
  - product_requirements
  - domain_rules
  - frontend_architecture_contract
  - feature_interface_requirements
  - implementation_source_of_truth
  - implementation_sequence
---
```

## Instantiated Body

```markdown
# UI Design Guide

## Role

Этот guide помогает найти и правильно применить существующие UI assets проекта. Canonical owners:

- `../frontend.md` владеет frontend stack, boundaries и обязательными engineering rules;
- product/domain docs владеют product intent, business language и state semantics;
- `features/FT-XXX/ui-reference/README.md` описывает interface change конкретной feature.

Guide владеет curated discovery map, а не implementation truth. Код владеет фактическими component APIs, signatures и behavior; перед изменением проверяй указанные source paths и examples по текущему checkout.

## Component Catalog

| Component / pattern | Existing use | Source paths | Examples / screenshots | Owner rule |
| --- | --- | --- | --- | --- |
| Название существующего компонента | Где и когда он уже применяется | Реальный путь | Реальный link/path | Ссылка на правило в `../frontend.md` |

## Forms And Validation

| Pattern | Validation / error presentation | Helper APIs | Source paths | Examples |
| --- | --- | --- | --- | --- |
| Существующий form pattern | Наблюдаемое UI behavior | Реальные helpers | Реальные пути | Реальные примеры |

## Actions And Navigation

| Pattern | Existing placement / behavior | States | Source paths | Examples |
| --- | --- | --- | --- | --- |
| Action, navigation или routing pattern | Как он используется сейчас | loading / disabled / confirmation | Реальные пути | Реальные примеры |

## Tables And Collections

| Pattern | Controls | Empty / loading / error presentation | Source paths | Examples |
| --- | --- | --- | --- | --- |
| Существующая table/list implementation | sorting / filtering / pagination | Наблюдаемые states | Реальные пути | Реальные примеры |

## Visual State Mapping

Фиксируй только visual representation и обязательно ссылайся на owner business semantics.

| UI label / visual state | Presentation | Semantic owner | Source paths | Examples |
| --- | --- | --- | --- | --- |
| Реальный label/state | Как выглядит в UI | Ссылка на product/domain owner | Реальные пути | Реальный пример |

## Agent Entry Points

| Task type | Inspect first | Reusable examples | Additional owner context |
| --- | --- | --- | --- |
| Например, добавить form | Реальные component/helper paths | Реальные screens/tests/screenshots | Ссылка на `../frontend.md` и нужный domain/product owner |

## Maintenance

Обновляй guide, когда shared UI component, helper API, representative example или source path добавлен, удален или materially changed. Если запись не удается подтвердить по коду, исправь или удали ее до использования guide как implementation context.
```

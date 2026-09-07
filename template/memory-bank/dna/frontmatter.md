---
title: "Frontmatter Schema"
doc_kind: governance
doc_function: canonical
purpose: "Frontmatter Schema"
derived_from:
  - governance.md
status: active
audience: humans_and_agents
---

# Frontmatter Schema

## Общая schema

Каждый governed-документ имеет YAML frontmatter с `status`: `draft`, `active` или `archived`.
Для active non-root документа нужен `derived_from`: непустой путь, массив путей или
объектов `{path, fit}` с прямыми upstream-зависимостями. Корень дерева — principles.md.
`title`, `purpose`, `doc_kind` и `doc_function` описывают документ; один descriptive kind
сам по себе не подключает дополнительные правила. Дополнительные поля допустимы.
Их смысл задаётся владельцем выбранного типа документа или явно подключённого процесса.
Публикационный статус документа и состояние описываемой сущности независимы.

## Audience

Необязательное `audience` принимает `humans` или `humans_and_agents`.
Документ для humans_and_agents не объявляет документ с audience: humans своим semantic
upstream. Навигационная ссылка не является semantic dependency. Отсутствующее audience
не выводится из пути или doc_kind и сохраняет совместимость прежних документов.

## Пример

```yaml
---
status: active
derived_from:
  - governance.md
audience: humans_and_agents
---
```

[Machine rules](rules.json) фиксируют общую автоматическую часть. Поля дополнительных
контрактов не становятся обязательными из-за одного descriptive doc_kind.

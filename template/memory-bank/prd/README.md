---
title: "PRD index"
doc_kind: project
doc_function: index
purpose: "PRD index"
derived_from:
  - ../document-types/prd.md
status: active
audience: humans_and_agents
---

# PRD index

Здесь хранятся заполненные проектные документы. Они принадлежат проекту.

- [Базовый контракт](../document-types/prd.md)
- [Шаблон](../templates/prd.md)

Создание без подключения процесса:

```sh
memory-bank-cli document create --type prd --path memory-bank/prd/PRD-001-name.md
```

Добавляй сюда ссылки на реально существующие документы. Для пакета создай README,
который индексирует его реальные артефакты. Устанавливаемый компонент Documents
не требует executor tools или обязательного маршрута AI-разработки.

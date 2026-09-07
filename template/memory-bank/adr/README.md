---
title: "ADR index"
doc_kind: project
doc_function: index
purpose: "ADR index"
derived_from:
  - ../document-types/adr.md
status: active
audience: humans_and_agents
---

# ADR index

Здесь хранятся заполненные проектные документы. Они принадлежат проекту.

- [Базовый контракт](../document-types/adr.md)
- [Шаблон](../templates/adr.md)

Создание без подключения процесса:

```sh
memory-bank-cli document create --type adr --path memory-bank/adr/ADR-001-name.md
```

Добавляй сюда ссылки на реально существующие документы. Для пакета создай README,
который индексирует его реальные артефакты. Устанавливаемый компонент Documents
не требует executor tools или обязательного маршрута AI-разработки.

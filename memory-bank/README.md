---
title: Template Documentation Index
doc_kind: project
doc_function: index
purpose: Корневая навигация по шаблонному memory-bank. Читать сначала, чтобы понять структуру и точки адаптации под конкретный проект.
status: active
audience: humans_and_agents
---

# Documentation Index

Каталог `memory-bank/` содержит переносимый шаблон проектной документации для разработки ПО. После копирования в downstream-репозиторий адаптируй `domain/`, `engineering/` и `ops/` под реальный стек, процессы и ограничения проекта.

Конкретные instantiated примеры вынесены в корневой каталог `examples/`.

## Аннотированный индекс

- [`domain/README.md`](domain/README.md)
  Читать, когда нужно: зафиксировать product context, архитектурные границы и UI-соглашения проекта.

- [`prd/README.md`](prd/README.md)
  Читать, когда нужно: описать продуктовую инициативу между общим problem statement и downstream feature packages.

- [`use-cases/README.md`](use-cases/README.md)
  Читать, когда нужно: зарегистрировать устойчивый пользовательский или операционный сценарий проекта.

- [`ops/README.md`](ops/README.md)
  Читать, когда нужно: описать локальную разработку, окружения, релизы, конфигурацию и runbooks.

- [`engineering/README.md`](engineering/README.md)
  Читать, когда нужно: задать testing policy, coding style, git workflow и границы автономии агента.

- [`dna/README.md`](dna/README.md)
  Читать, когда нужно: проверить SSoT rules, frontmatter contract и governance-правила документации.

- [`flows/README.md`](flows/README.md)
  Читать, когда нужно: создать feature package, провести фичу по lifecycle gates или использовать шаблон.

- [`adr/README.md`](adr/README.md)
  Читать, когда нужно: найти или завести Architecture Decision Record.

- [`features/README.md`](features/README.md)
  Читать, когда нужно: понять, где живут instantiated feature packages.

---
title: Architecture Patterns
doc_kind: domain
doc_function: canonical
purpose: Каноничное место для архитектурных границ проекта. Читать при изменениях, затрагивающих модули, фоновые процессы, интеграции или конфигурацию.
derived_from:
  - ../dna/governance.md
status: active
audience: humans_and_agents
---

# Architecture Patterns

Этот документ задает не конкретную реализацию, а ожидаемые архитектурные правила проекта. Подставь сюда реальные bounded contexts, integration boundaries и технические ограничения downstream-системы.

## Module Boundaries

Зафиксируй здесь главные изолированные области системы.

Пример:

| Context | Owns | Must not depend on directly |
| --- | --- | --- |
| `customer-facing` | пользовательский путь, публичные API | внутренние админские детали |
| `operations` | backoffice, ручные действия, moderation | приватные внутренности billing/storage |
| `platform` | shared services, auth, delivery infrastructure | product-specific UI assumptions |

Минимальные правила:

- модуль владеет своим state и публичными контрактами;
- межмодульные зависимости проходят через явно названный API, event или adapter;
- UI, jobs и интеграции не должны читать чужие внутренние детали в обход owner-модуля.

## Concurrency And Critical Sections

Если проект содержит конкурентные операции, зафиксируй canonical pattern для критических секций и фона.

Пример:

```ruby
ResourceLock.with_lock(resource_key) do
  # критическая секция
end
```

Укажи явно:

- какой locking pattern разрешен;
- какой pattern запрещен и почему;
- что считается idempotent recovery;
- где проходят границы транзакции относительно внешних API.

Если проект использует job queue, добавь canonical правило для concurrency control.

## Failure Handling And Error Tracking

Зафиксируй единый подход:

- где ошибки поднимаются наверх, а где переводятся в domain verdict;
- как добавляется contextual metadata для error tracker;
- где retry policy уже реализована инфраструктурно и ее нельзя дублировать локальным `rescue`.

Пример вопроса, на который должен отвечать этот раздел:

> Нужно ли вручную логировать ошибку в job, если базовый job class уже делает retries и нотификацию?

## Configuration Ownership

Документируй не все переменные окружения подряд, а ownership-модель конфигурации:

- где живет canonical schema конфигурации;
- какие файлы или классы считаются owner-слоем;
- где задаются defaults;
- кто отвечает за документацию env contract.

Пример:

1. Обновить schema-owner конфигурации.
2. Обновить default values или environment overlays.
3. Обновить [`../ops/config.md`](../ops/config.md).

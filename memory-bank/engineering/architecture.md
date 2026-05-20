---
title: Engineering Architecture Patterns
doc_kind: engineering
doc_function: canonical
purpose: Каноничное место для архитектурных правил реализации: code/module boundaries, runtime patterns, concurrency, error handling и configuration ownership.
derived_from:
  - ../dna/governance.md
  - ../domain/context-map.md
status: active
audience: humans_and_agents
---

# Engineering Architecture Patterns

Этот документ задает ожидаемые архитектурные правила реализации. Предметные bounded contexts описаны в [`../domain/context-map.md`](../domain/context-map.md); здесь фиксируй, как они отражаются в code modules, services, queues, adapters и configuration ownership.

## Module Boundaries

Зафиксируй главные изолированные области реализации.

Пример:

| Module / Layer | Owns | Must not depend on directly |
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

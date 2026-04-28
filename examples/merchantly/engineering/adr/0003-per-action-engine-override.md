---
title: "ADR-0003: Per-Action Engine Override"
doc_kind: engineering
doc_function: canonical
purpose: Архитектурное решение о механизме переключения рендеринга (React/Liquid) для отдельных экшенов без изменения глобального engine.
derived_from:
  - ../../../../memory-bank/dna/governance.md
status: active
decision_status: accepted
audience: humans_and_agents
---

# ADR-0003: Per-Action Engine Override

**Статус:** Принято
**Дата:** 2026-04-09
**Авторы:** @danil
**Issue:** https://github.com/BrandyMint/merchantly/issues/4454

## Контекст

Публичный фронтенд Merchantly поддерживает два engine рендеринга:
- **React** (`engine: :react`) — современный SPA фронтенд
- **Liquid** (`engine: :liquid`) — серверный рендеринг через шаблоны

По умолчанию все страницы используют React. Для инкрементальной миграции (например, checkout) нужен механизм переключения engine для конкретных экшенов без изменения глобального режима.

## Рассмотренные варианты

### 1. Global engine switch

**Как работает:**
Изменение `vendor.engine` или глобальной настройки для всех страниц.

**Плюсы:**
- Простота реализации

**Минусы:**
- Ломает все страницы при миграции
- Нет инкрементального перехода
- Риск регрессий на всём сайте

### 2. Per-action engine override через параметр (выбран)

**Как работает:**
```ruby
def new
  render_view 'vendor/orders/new', engine: :liquid
end
```

**Плюсы:**
- Thread-safe: нет мутации глобального состояния
- Инкрементальная миграция по экшенам
- Легко тестировать отдельные страницы
- Fallback на default engine если не указан

**Минусы:**
- Нужно модифицировать `render_view` concern
- Не работает для TEMPLATE_CUSTOM (отдельная логика)

### 3. Middleware-based engine detection

**Как работает:**
Middleware определяет engine по URL паттернам или заголовкам.

**Плюсы:**
- Централизованная логика

**Минусы:**
- Сложность отладки
- Неявное поведение
- Сложно переопределить для конкретного экшена

## Решение

**Выбран вариант 2: Per-action engine override через параметр**

### Реализация

```ruby
# app/controllers/concerns/render_view.rb
def render_view(template, options = {})
  engine = options.delete(:engine) || default_engine
  # ... рендеринг через указанный engine
end

# app/controllers/vendor/orders_controller.rb
def new
  render_view 'vendor/orders/new', engine: :liquid
end
```

### Constraints

| ID | Ограничение |
|----|-------------|
| PEO-01 | Параметр `engine:` передаётся через `options`, не через инстанс-переменные |
| PEO-02 | Thread-safe: нет мутации глобального состояния или классовых переменных |
| PEO-03 | Fallback: если шаблон не найден — ошибка 500 (fail fast) |
| PEO-04 | TEMPLATE_CUSTOM bypass: вендоры с `TEMPLATE_CUSTOM` игнорируют `engine:` параметр |

### Уточнение для optional custom Liquid templates

- `PEO-03` относится к explicit `engine:` override после того, как система уже выбрала путь explicit Liquid-render.
- Для feature, где Liquid включается как optional custom template поверх существующей baseline page implementation, действует [ADR-0006](0006-missing-custom-liquid-template-falls-back-to-existing-baseline-page.md).
- В таких feature отсутствие custom Liquid template body не должно доводиться до explicit Liquid-render path; resolver обязан сделать fallback на existing baseline page implementation раньше.

### Исключения

- `TEMPLATE_CUSTOM` вендоры используют свои шаблоны независимо от `engine:` параметра
- Per-action override не применяется к partial renders (только к полным страницам)

## Последствия

### Положительные

- Инкрементальная миграция с React на Liquid
- Возможность A/B тестирования engine для конкретных страниц
- Чёткий контракт через `engine:` параметр
- Thread-safe implementation

### Отрицательные

- Два engine в одной кодовой базе — технический долг
- Нужно поддерживать шаблоны для обоих engine
- Сложнее debug (нужно знать какой engine используется)

## Связанные решения

- **ADR-0001:** Stimulus для интерактивности (Liquid страницы используют Stimulus)
- **Feature FT-4542:** Checkout Liquid Migration — первая реализация per-action override
- **ADR-0006:** Missing Custom Liquid Template Falls Back to Existing Baseline Page — правило для optional custom Liquid features

## Для AI-агентов

При работе с рендерингом:
1. Используй `render_view` вместо прямого `render` для поддержки engine override
2. Для новых Liquid страниц явно указывай `engine: :liquid`
3. Помни о PEO-04: TEMPLATE_CUSTOM вендоры игнорируют engine параметр
4. Не используй инстанс-переменные для передачи engine — только через options
5. Если Liquid добавляется как optional custom template поверх существующей baseline page implementation, сначала проверь [ADR-0006](0006-missing-custom-liquid-template-falls-back-to-existing-baseline-page.md)

---
title: Frontend
doc_kind: domain
doc_function: canonical
purpose: Public frontend (React), operator frontend (Stimulus/jQuery), переводы. Читать при работе с UI.
derived_from:
  - ../dna/governance.md
status: active
audience: humans_and_agents
---

# Frontend

## Public Frontend (front-public)

Публичный фронтенд — витрина, товары, корзина, оформление заказа. Написан на React.

- **Репозиторий**: `git@github.com:merchantly/front-public.git`
- **Копия в монолите**: `vendor/static/public/app/scripts/react/`

### Переводы для фронтенда

Передаются в React через `VendorTranslationsService#all_translations`:
1. `I18n.backend.translate(locale, :vendor)` — все ключи под `vendor:` в YAML
2. Мержатся с кастомными переводами вендора из БД (`vendor.translations`)
3. Кэш: `[:vendor_translations, id, locale, cache_sweeped_at, translations_updated_at, :v19]`

Библиотека `i18next` с интерполяцией `%{variable}`.

**Ключевые файлы:**
- `app/builders/react/frontend_i18n_props_builder.rb`
- `app/services/vendor_translations_service.rb`
- `app/models/concerns/vendor_translations.rb`
- `vendor/static/public/app/scripts/react/components/HoC/provideTranslations.jsx`

## Operator Frontend (операторская панель)

`/operator` — серверный рендеринг + интерактивность через JavaScript.

### Текущий стек
- **Asset pipeline**: Sprockets
- **JavaScript**: CoffeeScript + jQuery (ES5-совместимый)
- **Стили**: Bootstrap 3 + SASS

### Добавление интерактивности (ADR-0001)

**Для новых фич используй Stimulus** — см. [ADR-0001](../../docs/adr/0001-frontend-interactivity-stimulus.md).

- **Stimulus** для JS-поведения (dropdowns, tabs, inline-edit, validation)
- **jQuery** можно внутри Stimulus контроллеров
- **Turbo** добавим позже

Пример:
```javascript
// app/javascript/controllers/example_controller.js
import { Controller } from "@hotwired/stimulus"

export default class extends Controller {
  static targets = ["output"]
  static values = { url: String }

  greet() {
    this.outputTarget.textContent = "Hello!"
  }
}
```

```erb
<div data-controller="example" data-example-url-value="<%= path %>">
  <button data-action="click->example#greet">Click</button>
  <span data-example-target="output"></span>
</div>
```

**Документация:**
- [ADR-0001](../../docs/adr/0001-frontend-interactivity-stimulus.md)
- [Stimulus Handbook](https://stimulus.hotwired.dev/handbook/introduction)
- [docs/adr/](../../docs/adr/)

## Локали (переводы)

Лежат в submodule `./config/custom_locales`. При обновлении:
1. Пушай коммит в **master** сабмодуля
2. Обновляй сабмодуль на master, а не на конкретный коммит

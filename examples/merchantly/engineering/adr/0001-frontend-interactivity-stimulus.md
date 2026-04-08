# ADR-0001: Использование Stimulus для интерактивности в операторской панели

**Статус:** Принято
**Дата:** 2026-02-02
**Авторы:** @danil
**Issue:** https://github.com/BrandyMint/merchantly/issues/4364

## Контекст

Операторская панель Merchantly использует legacy-стек для фронтенда:
- **Asset pipeline**: Sprockets (классический)
- **JavaScript**: CoffeeScript + jQuery (ES5-совместимый)
- **Нет современных фреймворков**: Turbo, Stimulus, Hotwire, Webpacker, Importmap

Для добавления интерактивности (inline-редактирование, динамические формы, real-time обновления) необходимо выбрать подход, который:
1. Совместим с существующим стеком (Sprockets, jQuery)
2. Позволяет инкрементальное внедрение
3. Удобен для AI-агентов (предсказуемые паттерны)
4. Поддерживается Rails-сообществом

## Рассмотренные варианты

### 1. Stimulus (выбран)
**Плюсы:**
- Минимальный, работает со Sprockets через UMD build
- Не конфликтует с jQuery
- Предсказуемая структура: `data-controller`, `data-action`, `data-target`
- Часть официального Rails стека (Hotwire)
- Размер: ~45KB minified

**Минусы:**
- Не решает задачи серверного обновления DOM (нужен Turbo)

### 2. Hotwire (Turbo + Stimulus)
**Плюсы:**
- Полный стек для SPA-like приложений
- Turbo Streams для real-time обновлений
- Официальная поддержка Rails

**Минусы:**
- Turbo Drive ломает `$(document).ready()` — нужна миграция всего jQuery кода
- Требует изменений на backend (turbo_stream views)
- Сложнее внедрить инкрементально

### 3. Turbo standalone
**Минусы:**
- Агрессивно перехватывает навигацию
- Без Stimulus нет удобного способа добавить JS-поведение
- Конфликты с существующим jQuery кодом

### 4. Svelte
**Минусы:**
- Требует bundler (esbuild/Vite/webpack)
- Не работает с Sprockets напрямую
- Отдельный build pipeline
- Svelte 5 runes — breaking changes

## Решение

**Выбран Stimulus** как первый шаг к модернизации фронтенда.

### Архитектура Hotwire

```
┌─────────────────────────────────────────────────────────┐
│                      HOTWIRE                            │
│                  (зонтичный бренд)                      │
├─────────────────────────┬───────────────────────────────┤
│         TURBO           │          STIMULUS             │
│   (HTML over the wire)  │    (JS для существующего HTML)│
├─────────────────────────┼───────────────────────────────┤
│ • Turbo Drive           │ • Controllers    ← НАЧИНАЕМ  │
│ • Turbo Frames          │ • Targets           ЗДЕСЬ    │
│ • Turbo Streams         │ • Values                      │
│ • Turbo Native          │ • Actions                     │
└─────────────────────────┴───────────────────────────────┘
```

**Важно:** Turbo и Stimulus не противоречат друг другу, а дополняют:
- **Stimulus** — добавляет JS-поведение к существующему HTML
- **Turbo** — обновляет HTML с сервера без перезагрузки страницы

Можно использовать Stimulus без Turbo (наш случай), Turbo без Stimulus, или оба вместе.

## План внедрения

### Фаза 1: Stimulus (текущая)
```
Риск: низкий
Изменения: минимальные
jQuery: продолжает работать
```

- Добавить `stimulus-rails` gem
- Настроить importmap или Sprockets UMD
- Новые интерактивные фичи писать на Stimulus
- Постепенно мигрировать jQuery код

### Фаза 2: Turbo Frames (будущее)
```
Риск: средний
Требует: миграция $(document).ready() на turbo:load
```

- Добавить когда основной jQuery код мигрирован на Stimulus
- Использовать для inline-редактирования, lazy loading
- Требует изменения views (turbo_frame_tag)

### Фаза 3: Turbo Streams (будущее)
```
Риск: средний
Требует: ActionCable + Redis
```

- Для real-time фич (notifications, live updates)
- Broadcasting через WebSocket
- Требует фоновые jobs для broadcasts

## Паттерны Stimulus для AI-агентов

### Структура контроллера

```javascript
// app/javascript/controllers/example_controller.js
import { Controller } from "@hotwired/stimulus"

export default class extends Controller {
  // 1. Статические определения (предсказуемая структура)
  static targets = ["input", "output"]  // элементы для работы
  static values = { url: String, count: Number }  // данные из HTML
  static classes = ["active", "hidden"]  // CSS классы

  // 2. Lifecycle callbacks
  connect() {
    // Вызывается когда контроллер подключается к DOM
  }

  disconnect() {
    // Вызывается когда элемент удаляется из DOM
  }

  // 3. Action methods (вызываются из HTML)
  submit() {
    // data-action="click->example#submit"
  }

  // 4. Value changed callbacks
  countValueChanged() {
    // Автоматически вызывается при изменении data-example-count-value
  }
}
```

### HTML разметка

```erb
<div data-controller="example"
     data-example-url-value="<%= api_path %>"
     data-example-count-value="0"
     data-example-active-class="bg-green-500"
     data-example-hidden-class="hidden">

  <input data-example-target="input"
         data-action="input->example#validate keydown.enter->example#submit">

  <div data-example-target="output"></div>

  <button data-action="click->example#submit">Submit</button>
</div>
```

### Naming conventions

| Элемент | Паттерн | Пример |
|---------|---------|--------|
| Контроллер | `snake_case_controller.js` | `inline_edit_controller.js` |
| data-controller | `kebab-case` | `data-controller="inline-edit"` |
| Target | `camelCase` в JS, `kebab-case` в HTML | `inputTarget` / `data-inline-edit-target="input"` |
| Value | `camelCase` в JS, `kebab-case` в HTML | `urlValue` / `data-inline-edit-url-value` |
| Action | `controller#method` | `data-action="click->inline-edit#save"` |

## Пример: Inline-редактирование

### HTML (ERB)

```erb
<%# app/views/operator/orders/_note.html.erb %>
<div data-controller="inline-edit"
     data-inline-edit-url-value="<%= operator_order_path(@order) %>"
     data-inline-edit-field-value="note">

  <span data-inline-edit-target="display"
        data-action="click->inline-edit#edit">
    <%= @order.note.presence || 'Добавить заметку...' %>
  </span>

  <input data-inline-edit-target="input"
         data-action="blur->inline-edit#save keydown.enter->inline-edit#save keydown.escape->inline-edit#cancel"
         class="hidden form-control">

  <span data-inline-edit-target="spinner" class="hidden">
    <i class="fa fa-spinner fa-spin"></i>
  </span>
</div>
```

### JavaScript контроллер

```javascript
// app/javascript/controllers/inline_edit_controller.js
import { Controller } from "@hotwired/stimulus"

export default class extends Controller {
  static targets = ["display", "input", "spinner"]
  static values = { url: String, field: String }

  edit() {
    this.inputTarget.value = this.displayTarget.textContent.trim()
    this.displayTarget.classList.add("hidden")
    this.inputTarget.classList.remove("hidden")
    this.inputTarget.focus()
    this.inputTarget.select()
  }

  cancel() {
    this.inputTarget.classList.add("hidden")
    this.displayTarget.classList.remove("hidden")
  }

  async save() {
    const value = this.inputTarget.value.trim()
    const originalValue = this.displayTarget.textContent.trim()

    if (value === originalValue) {
      this.cancel()
      return
    }

    this.showSpinner()

    try {
      const response = await fetch(this.urlValue, {
        method: "PATCH",
        headers: {
          "Content-Type": "application/json",
          "X-CSRF-Token": this.csrfToken,
          "Accept": "application/json"
        },
        body: JSON.stringify({ order: { [this.fieldValue]: value } })
      })

      if (response.ok) {
        this.displayTarget.textContent = value || 'Добавить заметку...'
        this.cancel()
      } else {
        const error = await response.json()
        alert(error.message || 'Ошибка сохранения')
        this.inputTarget.focus()
      }
    } catch (error) {
      console.error("Save failed:", error)
      alert('Ошибка сети')
      this.inputTarget.focus()
    } finally {
      this.hideSpinner()
    }
  }

  showSpinner() {
    if (this.hasSpinnerTarget) {
      this.spinnerTarget.classList.remove("hidden")
    }
  }

  hideSpinner() {
    if (this.hasSpinnerTarget) {
      this.spinnerTarget.classList.add("hidden")
    }
  }

  get csrfToken() {
    return document.querySelector("[name='csrf-token']")?.content
  }
}
```

## Совместимость с jQuery

Stimulus **не конфликтует** с jQuery. Можно использовать jQuery внутри контроллеров:

```javascript
import { Controller } from "@hotwired/stimulus"

export default class extends Controller {
  connect() {
    // jQuery работает внутри Stimulus контроллера
    $(this.element).find('[data-toggle="popover"]').popover()
  }

  disconnect() {
    // Cleanup при удалении элемента
    $(this.element).find('[data-toggle="popover"]').popover('dispose')
  }
}
```

## Проблема Turbo Drive с jQuery

**Внимание:** Если в будущем добавим Turbo Drive, нужно будет мигрировать:

```javascript
// ❌ Перестанет работать с Turbo Drive
$(document).ready(function() {
  $('[data-toggle="popover"]').popover();
});

// ✅ Работает и с Turbo, и без
document.addEventListener("turbo:load", function() {
  $('[data-toggle="popover"]').popover();
});

// ✅✅ Лучший вариант — Stimulus контроллер (работает везде)
// Контроллер автоматически вызывает connect() для новых элементов
```

## Установка

### Вариант 1: Importmap (рекомендуется для Rails 7+)

```ruby
# Gemfile
gem 'stimulus-rails'
```

```bash
bundle install
rails stimulus:install
```

```ruby
# config/importmap.rb
pin "@hotwired/stimulus", to: "stimulus.min.js"
pin "@hotwired/stimulus-loading", to: "stimulus-loading.js"
pin_all_from "app/javascript/controllers", under: "controllers"
```

### Вариант 2: Sprockets UMD (для legacy проектов)

```javascript
// app/assets/javascripts/application.js
//= require @hotwired/stimulus/dist/stimulus.umd.js

// Ручная регистрация контроллеров
const application = Stimulus.Application.start()
application.register("inline-edit", InlineEditController)
```

## Ресурсы

- [Stimulus Handbook](https://stimulus.hotwired.dev/handbook/introduction)
- [Stimulus Reference](https://stimulus.hotwired.dev/reference/controllers)
- [Hotwire Discussion](https://discuss.hotwired.dev/)
- [stimulus-use](https://github.com/stimulus-use/stimulus-use) — полезные миксины
- [tailwindcss-stimulus-components](https://github.com/excid3/tailwindcss-stimulus-components) — готовые компоненты

## Последствия

### Положительные
- Новые интерактивные фичи можно разрабатывать без изменения инфраструктуры
- jQuery код продолжает работать
- Предсказуемые паттерны упрощают code review и AI-генерацию
- Путь к полному Hotwire стеку остается открытым

### Отрицательные
- Два подхода к JS (jQuery и Stimulus) на переходный период
- Нужно обучение команды новым паттернам
- Без Turbo нет server-side DOM updates (пока используем обычный fetch + JS)

## Связанные решения

- **ADR-XXXX**: Миграция на Turbo Frames (будущее)
- **ADR-XXXX**: Real-time обновления через Turbo Streams (будущее)

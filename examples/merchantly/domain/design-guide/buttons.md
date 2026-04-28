---
title: Кнопки (Buttons)
doc_kind: domain
doc_function: reference
purpose: Хелперы для создания кнопок в операторской панели. Читать при создании UI.
derived_from:
  - ../../../../memory-bank/dna/governance.md
status: active
audience: humans_and_agents
---

# Кнопки (Buttons)

Хелперы для создания кнопок в операторской панели.

**Расположение:** `app/helpers/operator_buttons_helper.rb`

---

## Кнопки создания

### add_button — Кнопка добавления

**Назначение:** Основная кнопка для создания нового ресурса. Используется в заголовках списков.

```ruby
add_button(url, title: nil, target: nil)
```

| Параметр | Тип | По умолчанию | Описание |
|----------|-----|--------------|----------|
| `url` | String | — | URL для создания |
| `title` | String | `t('shared.create')` | Текст кнопки |
| `target` | String | nil | Target для ссылки (`_blank`) |

**Примеры:**

```haml
-# В заголовке страницы списка
.ibox-tools
  = add_button new_operator_product_path

-# С кастомным заголовком
= add_button new_operator_product_path(product: { category_id: category.id }),
             title: t('operator.categories.add_button_title')

-# Открытие в новом окне
= add_button new_operator_category_path, target: '_blank'
```

**CSS:** `btn btn-success btn-outline`

**Инструкция для агента:** Используй `add_button` в `.ibox-tools` или `.header-actions` на страницах со списками ресурсов.

---

### add_round_button — Круглая кнопка добавления

**Назначение:** Компактная круглая кнопка с иконкой +. Используется в inline-контекстах.

```ruby
add_round_button(url, title: nil, size: :sm, method: :get)
```

| Параметр | Тип | По умолчанию | Описание |
|----------|-----|--------------|----------|
| `url` | String | — | URL для создания |
| `title` | String | `t('shared.create_new')` | Текст рядом с иконкой |
| `size` | Symbol | `:sm` | Размер кнопки (`:xs`, `:sm`, `:md`, `:lg`) |
| `method` | Symbol | `:get` | HTTP метод |

**Примеры:**

```haml
-# В таблице
%td= add_round_button new_operator_category_product_path(category)

-# Большая кнопка
= add_round_button new_operator_order_path, size: :md
```

**CSS:** `btn btn-primary btn-rounded btn-{size}`

---

### add_big_round_button — Большая круглая кнопка

**Назначение:** Крупная кнопка с иконкой + для акцентных действий.

```ruby
add_big_round_button(url, title: nil)
```

| Параметр | Тип | По умолчанию | Описание |
|----------|-----|--------------|----------|
| `url` | String | — | URL для создания |
| `title` | String | `t('shared.add')` | Текст кнопки |

**CSS:** `btn btn-primary btn-rounded btn-fa-small`

---

## Кнопки редактирования

### edit_button — Кнопка редактирования

```ruby
edit_button(url, title: nil)
```

| Параметр | Тип | По умолчанию | Описание |
|----------|-----|--------------|----------|
| `url` | String | — | URL формы редактирования |
| `title` | String | `t('shared.edit')` | Текст кнопки |

**Примеры:**

```haml
= edit_button edit_operator_product_path(@product)
```

**CSS:** `btn` с иконкой `fa-pencil`

---

## Кнопки удаления

### delete_button — Кнопка удаления

**Назначение:** Кнопка удаления с подтверждением.

```ruby
delete_button(url, css_class: 'btn btn-white btn-sm')
```

| Параметр | Тип | По умолчанию | Описание |
|----------|-----|--------------|----------|
| `url` | String | — | URL для DELETE запроса |
| `css_class` | String | `'btn btn-white btn-sm'` | CSS классы |

**Примеры:**

```haml
-# Стандартная кнопка удаления
= delete_button operator_product_path(@product)

-# С кастомными классами
= delete_button operator_category_path(category), css_class: 'btn btn-danger btn-xs'
```

**CSS:** По умолчанию `btn btn-white btn-sm` с иконкой `fa-trash`

---

### delete_submit_button — Кнопка-submit удаления

**Назначение:** Полноразмерная кнопка удаления для форм.

```ruby
delete_submit_button(url, css_class = '')
```

**CSS:** `btn btn-danger`

---

### delete_submit_icon — Иконка удаления

**Назначение:** Только иконка корзины для компактного отображения в таблицах.

```ruby
delete_submit_icon(url)
```

**Примеры:**

```haml
%td.actions
  = delete_submit_icon operator_product_path(product)
```

**CSS:** Иконка `fa-trash text-muted` с tooltip

---

## Переключатели (Toggle Buttons)

### toggle_button — Базовый переключатель

**Назначение:** Переключатель состояния (вкл/выкл) с AJAX-обновлением.

```ruby
toggle_button(status, on_url:, off_url:, title: nil, tooltip: nil)
```

| Параметр | Тип | Описание |
|----------|-----|----------|
| `status` | Boolean | Текущее состояние |
| `on_url` | String | URL для включения |
| `off_url` | String | URL для выключения |
| `title` | String | Заголовок рядом с переключателем |
| `tooltip` | String | Всплывающая подсказка |

**Примеры:**

```haml
= toggle_button @product.is_published?,
                on_url: publish_operator_product_path(@product, format: :json),
                off_url: unpublish_operator_product_path(@product, format: :json),
                title: 'Опубликован',
                tooltip: 'Включить/выключить публикацию товара'
```

---

### smart_toggle_button — Умный переключатель

**Назначение:** Автоматически генерирует URL-ы на основе модели и атрибута.

```ruby
smart_toggle_button(resource, method, title = nil, tooltip: nil)
```

| Параметр | Тип | Описание |
|----------|-----|----------|
| `resource` | ActiveRecord | Модель с переключаемым атрибутом |
| `method` | Symbol | Имя атрибута для переключения |
| `title` | String | Заголовок |
| `tooltip` | String | Подсказка |

**Примеры:**

```haml
= smart_toggle_button @product, :is_published
= smart_toggle_button @category, :is_visible, 'Видимость', tooltip: 'Показывать в каталоге'
```

**Инструкция для агента:** Предпочитай `smart_toggle_button` вместо `toggle_button` — он автоматически строит URL-ы.

---

### toggle_state_button — Кнопка активации/деактивации

**Назначение:** Переключатель с визуальным состоянием для моделей с методом `is_active?`.

```ruby
toggle_state_button(model)
```

**Требования к модели:**
- Метод `is_active?`
- Routes: `activate_operator_{model}_path`, `deactivate_operator_{model}_path`

**CSS:**
- Активен: `btn btn-rounded btn-success btn-sm btn-toggle btn-outline active`
- Неактивен: `btn btn-rounded btn-default btn-sm btn-toggle btn-outline`

---

## Кнопки в header

### header_action_button — Кнопка действия в header

```ruby
header_action_button(title, path, options = {})
```

| Параметр | Тип | Описание |
|----------|-----|----------|
| `title` | String | Текст кнопки |
| `path` | String | URL действия |
| `options` | Hash | Дополнительные опции для `link_to` |

**CSS:** `btn btn-default btn-outline`

---

### subsettings_button — Кнопка выпадающего меню

**Назначение:** Кнопка с иконкой шеврона для dropdown-меню дополнительных действий.

```ruby
subsettings_button
```

**CSS:** `dropdown-toggle btn btn-default btn-header`

---

## Кнопки сворачивания

### collapse_button — Кнопка сворачивания

```ruby
collapse_button(dir = 'up')
```

| Параметр | Тип | По умолчанию | Описание |
|----------|-----|--------------|----------|
| `dir` | String | `'up'` | Направление шеврона (`'up'`, `'down'`) |

---

## Сортировка drag&drop

### operator_sort_handle — Ручка сортировки

**Назначение:** Иконка для перетаскивания элементов при сортировке.

```ruby
operator_sort_handle(resource)
```

**Примеры:**

```haml
%tr.sortable-row
  %td.sort-handle= operator_sort_handle product
  %td= product.title
```

---

### smart_operator_sort_handle — Умная ручка сортировки

**Назначение:** Скрывает ручку сортировки для архивированных записей.

```ruby
smart_operator_sort_handle(resource, disabled: false)
```

| Параметр | Тип | По умолчанию | Описание |
|----------|-----|--------------|----------|
| `resource` | ActiveRecord | — | Сортируемый ресурс |
| `disabled` | Boolean | false | Принудительно скрыть ручку |

**Инструкция для агента:** Используй `smart_operator_sort_handle` вместо `operator_sort_handle` — он автоматически скрывает ручку для архивированных записей.

---

## Внешние ссылки

### external_link — Внешняя ссылка

**Назначение:** Ссылка на внешний ресурс с иконкой стрелки.

```ruby
external_link(text, url)
```

| Параметр | Тип | Описание |
|----------|-----|----------|
| `text` | String | Текст ссылки |
| `url` | String | URL внешнего ресурса |

**Примеры:**

```haml
= external_link 'Документация API', 'https://api.example.com/docs'
```

**CSS:** `.external-link` с `target: '_blank'` и `rel: 'noopener'`

---

## Паттерны использования

### Страница списка ресурсов

```haml
.ibox
  .ibox-title
    %h5= t('.title')
    .ibox-tools
      = add_button new_operator_product_path
  .ibox-content
    %table.table
      %thead
        -# ...
      %tbody
        - @products.each do |product|
          %tr
            %td= smart_operator_sort_handle product
            %td= product.title
            %td.actions
              = edit_button edit_operator_product_path(product)
              = delete_submit_icon operator_product_path(product)
```

### Форма редактирования

```haml
.ibox
  .ibox-title
    %h5= t('.title')
    .ibox-tools
      = smart_toggle_button @product, :is_published
  .ibox-content
    = simple_form_for [:operator, @product] do |f|
      -# поля формы
      .form-actions
        = f.submit class: 'btn btn-primary'
        = delete_button operator_product_path(@product)
```

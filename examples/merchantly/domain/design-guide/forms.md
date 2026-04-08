# Формы (Forms)

Хелперы и паттерны для создания форм в операторской панели.

**Основные файлы:**
- `app/helpers/operator_form_helper.rb` — хелперы для полей форм
- `app/helpers/operator_resource_simple_form_helper.rb` — обёртка SimpleForm для ресурсов
- `app/views/operator/base/_resource_form.haml` — базовая структура формы

---

## resource_simple_form_for — Обёртка для форм ресурсов

**Расположение:** `app/helpers/operator_resource_simple_form_helper.rb`

### Назначение

Обёртка над `simple_form_for` с автоматической обработкой архивированных записей и стандартными wrapper mappings.

### Сигнатура

```ruby
resource_simple_form_for(resource, &block)
```

### Возможности

- Автоматически добавляет namespace `[:operator, resource]`
- Для архивированных записей устанавливает `disabled: true` для всех полей
- Добавляет CSS класс `form-disabled` для архивных записей
- Стандартные `wrapper_mappings` для boolean и check_boxes

### Примеры использования

```haml
-# Стандартная форма ресурса
= resource_simple_form_for @product do |f|
  = f.input :title
  = f.input :price
  = f.submit

-# Форма автоматически станет readonly для архивного товара
= resource_simple_form_for @archived_product do |f|
  = f.input :title  -# будет disabled
```

### Инструкция для агента

Используй `resource_simple_form_for` для форм редактирования моделей с методом `archived?`. Для форм без архивации или с кастомной логикой используй `simple_form_for [:operator, resource]`.

---

## Стандартные формы с SimpleForm

### Базовый паттерн

```haml
= simple_form_for [:operator, @resource], wrapper_mappings: { boolean: boolean_wrapper } do |f|
  = f.input :title
  = f.input :is_published, as: :boolean

  .form-actions
    = f.submit class: 'btn btn-primary'
```

### Wrapper Mappings

| Тип поля | Wrapper | Описание |
|----------|---------|----------|
| `boolean` | `boolean_wrapper` | Переключатель (switchery) |
| `check_boxes` | `:vertical_radio_and_checkboxes` | Вертикальный список чекбоксов |
| `radio_buttons` | `:vertical_radio_and_checkboxes` | Вертикальные radio кнопки |

---

## operator_form_warnings — Предупреждения в форме

### Назначение

Отображает блок предупреждений в форме.

### Сигнатура

```ruby
operator_form_warnings(warnings)
```

| Параметр | Тип | Описание |
|----------|-----|----------|
| `warnings` | String/Array/Hash | Текст(ы) предупреждений |

### Примеры

```haml
= operator_form_warnings 'Внимание! После сохранения изменения нельзя отменить.'

= operator_form_warnings ['Предупреждение 1', 'Предупреждение 2']

= operator_form_warnings @form.warnings if @form.warnings.any?
```

**CSS:** `alert alert-warning` с иконкой `fa-warning`

---

## form_errors — Ошибки формы

### Назначение

Отображает ошибки валидации для form object.

### Сигнатура

```ruby
form_errors(form)
```

### Примеры

```haml
= simple_form_for @product do |f|
  = form_errors @product
  = f.input :title
```

**CSS:** `alert alert-danger`

---

## Поля для описаний

### input_description — Поле описания

```ruby
input_description(f, as: :redactor, locale: nil, hint: nil)
```

| Параметр | Тип | По умолчанию | Описание |
|----------|-----|--------------|----------|
| `f` | FormBuilder | — | Form builder объект |
| `as` | Symbol | `:redactor` | Тип редактора |
| `locale` | String | nil | Локаль для мультиязычности |
| `hint` | String | nil | Подсказка к полю |

### Примеры

```haml
= input_description f
= input_description f, locale: :en, hint: 'Описание на английском'
```

---

### product_description_input — Описание товара

```ruby
product_description_input(
  form: nil,
  disabled: false,
  field_name: :custom_description,
  placeholder: nil,
  hint: nil,
  locale: I18n.locale
)
```

### Примеры

```haml
= product_description_input form: f
= product_description_input form: f, disabled: @product.archived?, locale: :en
```

---

## Поля атрибутов

### attribute_input — Универсальное поле атрибута

**Назначение:** Автоматически выбирает тип input в зависимости от типа Property.

```ruby
attribute_input(f, attribute, disabled = false)
```

| Тип Property | Результат |
|--------------|-----------|
| `PropertyDictionary` | Select с вариантами из словаря |
| `PropertyBoolean` | Checkbox (switchery) |
| `PropertyLink` | URL input |
| `PropertyFile` | File upload |
| `PropertyTime` | Datetime picker |
| `PropertyLong`, `PropertyDouble` | Numeric input |
| `PropertyText` | Textarea |
| Остальные | Text input |

### Примеры

```haml
- @product.attributes.each do |attr|
  .form-group
    %label= attr.property.title
    = attribute_input f, attr, @product.archived?
```

---

## Специальные поля

### select_country_code — Выбор страны

```ruby
select_country_code(form, field, hint: nil)
```

### Примеры

```haml
= select_country_code f, :country_code, hint: 'Выберите страну доставки'
```

---

### attribute_dictionary_input — Поле словаря

```ruby
attribute_dictionary_input(f, attribute, disabled = false)
```

Создаёт select с опциями из связанного словаря.

---

### attribute_file_input — Загрузка файла

```ruby
attribute_file_input(f, attribute, disabled = false)
```

Создаёт file input с информацией о текущем файле и чекбоксом удаления.

---

## Паттерны форм

### Форма создания/редактирования ресурса

```haml
.ibox
  .ibox-title
    %h5= @resource.new_record? ? t('.new_title') : t('.edit_title')
  .ibox-content
    = simple_form_for [:operator, @resource], wrapper_mappings: { boolean: boolean_wrapper } do |f|
      = form_errors @resource

      = f.input :title
      = f.input :description, as: :redactor
      = f.input :is_published

      .form-actions
        = f.submit class: 'btn btn-primary'
        - unless @resource.new_record?
          = delete_button operator_resource_path(@resource)
```

### Форма с вкладками

```haml
.ibox
  .ibox-title
    %h5= t('.title')
  .ibox-content
    = simple_form_for [:operator, @resource] do |f|
      %ul.nav.nav-tabs
        %li.active
          %a{href: '#tab-main', data: { toggle: 'tab' }} Основное
        %li
          %a{href: '#tab-seo', data: { toggle: 'tab' }} SEO

      .tab-content
        #tab-main.tab-pane.active
          = f.input :title
          = f.input :description

        #tab-seo.tab-pane
          = f.input :meta_title
          = f.input :meta_description

      .form-actions
        = f.submit class: 'btn btn-primary'
```

### Inline-форма в таблице

```haml
%tr
  = simple_form_for [:operator, item], wrapper: :inline_form do |f|
    %td= f.input_field :title, class: 'form-control'
    %td= f.input_field :position, class: 'form-control', type: :number
    %td= f.submit 'Сохранить', class: 'btn btn-sm btn-primary'
```

### Форма с архивацией

```haml
= resource_simple_form_for @product do |f|
  - if @product.archived?
    .alert.alert-info
      = t('operator.product.archived_notice')

  = f.input :title
  = f.input :price

  - unless @product.archived?
    .form-actions
      = f.submit class: 'btn btn-primary'
```

---

## bottom_form_panel — Панель кнопок формы

**Расположение:** `app/helpers/operator/bottom_form_panel_helper.rb`

### Назначение

Создаёт фиксированную панель внизу формы с кнопками действий (Сохранить, Удалить, Архивировать и др.).

### Сигнатура

```ruby
bottom_form_panel(f, opts = {}, additional_buttons = [])
```

| Параметр | Тип | Описание |
|----------|-----|----------|
| `f` | FormBuilder | Form builder объект |
| `opts` | Hash | Опции для кнопки Save (передаются в data-атрибуты) |
| `additional_buttons` | Array | Дополнительные кнопки (HTML-строки) |

### Автоматические кнопки по типу модели

Хелпер автоматически определяет набор кнопок в зависимости от класса модели:

| Модели | Кнопки |
|--------|--------|
| `Vendor`, `Order`, `Client`, `Tag`, `VendorSmtpSettings`... | Save |
| `BlogPost`, `Property`, `Dictionary`, `Member`, `Role`... | Save, Delete |
| `VendorPayment`, `VendorDelivery`, `Category`, `CustomPage`... | Save, Archive |
| `MailTemplate` | Save, Reset |

### Примеры использования

```haml
-# Стандартное использование — кнопки определяются автоматически
= bottom_form_panel f

-# С дополнительной кнопкой
- extra_btn = button_to 'Включить', enable_path, method: :post, class: 'btn btn-success'
= bottom_form_panel f, {}, [extra_btn]

-# С несколькими дополнительными кнопками
- btn1 = link_to 'Превью', preview_path, class: 'btn btn-default', target: '_blank'
- btn2 = button_to 'Опубликовать', publish_path, method: :post, class: 'btn btn-success'
= bottom_form_panel f, {}, [btn1, btn2]

-# С опциями для кнопки Save
= bottom_form_panel f, { save: { confirm: 'Вы уверены?' } }
```

### Layout кнопок

- **Одна кнопка:** отображается по центру
- **Несколько кнопок:** первая слева (Save), остальные справа через `&nbsp;`

### Инструкция для агента

Используй `additional_buttons` для добавления кастомных кнопок действий (Enable/Disable, Preview, Publish и т.д.) рядом с кнопкой Save. Кнопки передаются как массив HTML-строк.

---

## Инструкции для агента

1. **Используй `resource_simple_form_for`** для моделей с `archived?` методом
2. **Всегда добавляй `form_errors`** в начало формы
3. **Используй `wrapper_mappings`** для boolean полей
4. **Группируй связанные поля** в `_ibox_collapsed` или вкладки
5. **Кнопки действий** размещай через `bottom_form_panel` в конце формы
6. **Для мультиязычных полей** используй `input_description` с параметром `locale`
7. **Дополнительные кнопки** передавай через `additional_buttons` в `bottom_form_panel`

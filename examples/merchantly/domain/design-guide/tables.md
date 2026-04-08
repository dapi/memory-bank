# Таблицы и пагинация (Tables)

Компоненты для отображения списков и табличных данных.

**Основные файлы:**
- `app/views/operator/base/_empty_table.haml` — шаблон таблицы
- `app/views/operator/base/_select_page.haml` — выбор страницы
- `app/views/operator/base/_select_per_page.haml` — выбор количества на странице

---

## Базовые таблицы

### Структура таблицы

Стандартная структура таблицы в операторской:

```haml
.ibox
  .ibox-title
    %h5= t('.title')
    .ibox-tools
      = add_button new_operator_resource_path
  .ibox-content
    .project-list
      %table.table.table-hover
        %thead
          %tr
            %th= t('.columns.title')
            %th= t('.columns.status')
            %th.actions= t('.columns.actions')
        %tbody
          - @resources.each do |resource|
            = render 'resource', resource: resource
```

### CSS классы таблиц

| Класс | Описание |
|-------|----------|
| `.table` | Базовый стиль Bootstrap таблицы |
| `.table-hover` | Подсветка строки при наведении |
| `.table-striped` | Чередование цвета строк |
| `.table-bordered` | Границы ячеек |
| `.project-list` | Обёртка с отступами для таблицы в ibox |

---

## _empty_table — Шаблон таблицы с collection

**Расположение:** `app/views/operator/base/_empty_table.haml`

### Назначение

Шаблон для рендера таблицы с заголовками и коллекцией. Используется когда нужна таблица со стандартной структурой.

### Структура

```haml
.project-list
  %table.table.table-hover
    %thead
      %tr
        - columns.each do |column|
          %th{class: column.css_class}
            = column.content
    %tbody
      = render partial: partial, collection: collection
```

### Параметры

| Параметр | Тип | Описание |
|----------|-----|----------|
| `columns` | Array | Массив объектов с `content` и `css_class` |
| `partial` | String | Имя partial для каждого элемента |
| `collection` | Array | Коллекция для рендера |

### Примеры использования

```haml
= render 'operator/base/empty_table',
         columns: [
           OpenStruct.new(content: 'Название', css_class: ''),
           OpenStruct.new(content: 'Цена', css_class: 'text-right'),
           OpenStruct.new(content: '', css_class: 'actions')
         ],
         partial: 'product',
         collection: @products
```

---

## Пагинация

### _select_page — Выбор страницы

**Расположение:** `app/views/operator/base/_select_page.haml`

### Назначение

Dropdown для быстрого перехода на конкретную страницу.

### Параметры

| Параметр | Тип | Описание |
|----------|-----|----------|
| `result` | Kaminari::Result | Результат пагинации с `total_pages`, `current_page` |

### Примеры использования

```haml
.pagination-info
  = render 'operator/base/select_page', result: @products
```

### Структура

```haml
%form{action: current_url, method: :get, autosubmit: true, style: 'display: inline-block'}
  = hidden_field_tag :category_id, params[:category_id]
  = select_tag :page,
    options_for_select((1..result.total_pages).to_a, selected: result.current_page),
    tooltip: true,
    data: { placement: :right },
    title: t('operator.base.select_page.title'),
    class: 'paginator-page-selector'
```

**Особенности:**
- `autosubmit: true` — автоматическая отправка формы при выборе
- Сохраняет `category_id` при переключении страниц

---

### _select_per_page — Количество на странице

**Расположение:** `app/views/operator/base/_select_per_page.haml`

### Назначение

Dropdown для выбора количества элементов на странице.

### Параметры

| Параметр | Тип | Описание |
|----------|-----|----------|
| `result` | Kaminari::Result | Результат пагинации с `limit_value` |

### Примеры использования

```haml
.pagination-controls
  = render 'operator/base/select_per_page', result: @products
```

### Структура

```haml
%form{action: current_url, method: :get, autosubmit: true, style: 'display: inline-block'}
  - params.except(:per_page, :action, :controller).each do |key, value|
    = hidden_field_tag key, value
  = select_tag :per_page,
    options_for_select([10, 25, 50, 100], selected: result.limit_value || params[:per_page] || VendorBaseFilter::PER_PAGE),
    tooltip: true,
    data: { placement: :left },
    title: t('operator.per_page'),
    class: 'form-control'
```

**Варианты:** 10, 25, 50, 100 элементов на странице

---

## Полный пример страницы списка

```haml
-# app/views/operator/products/index.html.haml

- content_for :title do
  = t('.title')

.ibox
  .ibox-title
    %h5= t('.title')
    .ibox-tools
      = add_button new_operator_product_path

  -# Header с поиском и фильтрами
  .ibox-header
    = render 'operator/base/header_search',
             result: @products,
             options: @filter.options,
             query: @filter.query,
             reset_url: operator_products_path,
             counter_title: t('.counter', count: @products.total_count),
             one_tag_title: @filter.one_tag_title

  .ibox-content
    - if @products.any?
      .project-list
        %table.table.table-hover
          %thead
            %tr
              %th.sort-column Сортировка
              %th= t('.columns.title')
              %th= t('.columns.price')
              %th= t('.columns.status')
              %th.actions

          %tbody#sortable-products
            - @products.each do |product|
              %tr.sortable-row{id: dom_id(product)}
                %td.sort-handle= smart_operator_sort_handle product
                %td= link_to product.title, edit_operator_product_path(product)
                %td.text-right= product.price
                %td= product_state_label product
                %td.actions
                  = edit_button edit_operator_product_path(product)
                  = delete_submit_icon operator_product_path(product)

      -# Пагинация
      .row
        .col-md-6
          = paginate @products
        .col-md-6.text-right
          = render 'operator/base/select_page', result: @products
          = render 'operator/base/select_per_page', result: @products

    - else
      .empty-state
        %p= t('.empty_message')
        = add_button new_operator_product_path, title: t('.add_first')
```

---

## Сортировка drag&drop

### Подключение сортировки

```haml
%tbody#sortable-items{data: { sortable: true }}
  - @items.each do |item|
    %tr.sortable-row{id: dom_id(item)}
      %td.sort-handle= operator_sort_handle item
      %td= item.title
```

### Требования к контроллеру

```ruby
# routes.rb
resources :items do
  member do
    post :sort
  end
end

# items_controller.rb
def sort
  @item = Item.find(params[:id])
  @item.insert_at(params[:position].to_i)
  head :ok
end
```

---

## Паттерны таблиц

### Таблица с группировкой

```haml
%table.table
  - @categories.each do |category|
    %tr.category-header
      %th{colspan: 4}
        = category.title
        %span.badge= category.products.count

    - category.products.each do |product|
      %tr
        %td= product.title
        %td= product.price
```

### Таблица с inline-редактированием

```haml
%table.table
  - @items.each do |item|
    %tr
      = simple_form_for [:operator, item], remote: true, wrapper: :inline do |f|
        %td= f.input_field :title, class: 'form-control input-sm'
        %td= f.input_field :position, class: 'form-control input-sm', type: :number
        %td= f.submit 'Сохранить', class: 'btn btn-xs btn-primary'
```

### Таблица с чекбоксами для bulk actions

```haml
= form_tag bulk_action_operator_products_path, method: :post do
  .bulk-actions{style: 'display: none'}
    = submit_tag 'Удалить выбранные', class: 'btn btn-danger', data: { confirm: 'Удалить?' }
    = submit_tag 'Архивировать', class: 'btn btn-warning', name: 'archive'

  %table.table
    %thead
      %tr
        %th
          = check_box_tag 'select_all', nil, false, data: { toggle: 'checkall' }
        %th Название
    %tbody
      - @products.each do |product|
        %tr
          %td= check_box_tag 'ids[]', product.id, false, class: 'item-checkbox'
          %td= product.title
```

---

## Инструкции для агента

1. **Всегда используй `.project-list` обёртку** для таблиц внутри `.ibox-content`
2. **Добавляй класс `.table-hover`** для интерактивных таблиц
3. **Колонка actions должна быть последней** с классом `.actions`
4. **Используй `smart_operator_sort_handle`** вместо `operator_sort_handle`
5. **Добавляй empty state** когда коллекция пуста
6. **Пагинацию размещай под таблицей** в `.row` с двумя колонками
7. **Для сортировки** добавь `#sortable-*` id к tbody и `.sortable-row` к tr

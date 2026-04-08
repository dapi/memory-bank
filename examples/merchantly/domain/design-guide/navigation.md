# Навигация и поиск (Navigation)

Компоненты для навигации, поиска и фильтрации в операторской панели.

**Основные файлы:**
- `app/views/operator/base/_header_search.haml` — поиск в header
- `app/views/operator/base/_search_form_modal.haml` — модальное окно поиска
- `app/helpers/operator_menu_helper.rb` — хелперы меню

---

## Header Search

### _header_search — Поиск в заголовке

**Расположение:** `app/views/operator/base/_header_search.haml`

### Назначение

Компактная строка поиска с информацией о текущих фильтрах. Располагается в header страницы списка.

### Параметры

| Параметр | Тип | Описание |
|----------|-----|----------|
| `result` | Kaminari::Result | Результат пагинации |
| `options` | FilterOptions | Опции фильтра (без query, page, per_page) |
| `query` | String | Текущий поисковый запрос |
| `reset_url` | String | URL для сброса фильтров (опционально) |
| `counter_title` | String | Текст счётчика ("Найдено: 42") |
| `one_tag_title` | String | Заголовок если один фильтр активен |

### Примеры использования

```haml
.ibox-header
  = render 'operator/base/header_search',
           result: @products,
           options: @filter.options,
           query: @filter.query,
           reset_url: operator_products_path,
           counter_title: t('.counter', count: @products.total_count),
           one_tag_title: @filter.category&.title
```

### Структура

```haml
.header-search{ data: { 'search-form-dropdown' => 'search-form-dropdown' } }
  .header-search-info.m-r-sm
    %span.header-search-icon= ion_icon 'ios-search-strong'
  .header-search-info.header-search-line
    - if options.one?
      .tag= one_tag_title
    - elsif options.many?
      .tag= t :search_options_count, count: options.count, scope: :operator
    - if query.present?
      %span.header-search-query.m-l-sm= query
    - else
      %span.header-search-query.m-l-sm= t('operator.layouts.header_search.title')
  .header-search-info.header-vertical-middle.m-r-md
    %span.text-small= counter_title
  - if reset_url.present?
    .header-search-info.header-vertical-middle
      .m-r-md
        = link_to reset_url, title: t('operator.base.reset_filter'), data: { tooltip: true, placement: :bottom } do
          = ion_icon :close
```

### Взаимодействие

Клик по `.header-search` открывает модальное окно `#search-form-dropdown` с формой фильтров.

---

## Search Form Modal

### _search_form_modal — Модальное окно поиска

**Расположение:** `app/views/operator/base/_search_form_modal.haml`

### Назначение

Модальное окно с формой расширенного поиска и фильтрами.

### Параметры

| Параметр | Тип | Описание |
|----------|-----|----------|
| `title` | String | Заголовок модального окна |
| `yield` | Block | Содержимое формы (поля фильтров) |

### Примеры использования

```haml
= render 'operator/base/search_form_modal', title: t('.filter_title') do
  .row
    .col-md-6
      .form-group
        %label= t('.filters.query')
        = text_field_tag :query, params[:query], class: 'form-control'
    .col-md-6
      .form-group
        %label= t('.filters.category')
        = select_tag :category_id,
                     options_from_collection_for_select(@categories, :id, :title, params[:category_id]),
                     class: 'form-control',
                     include_blank: true
  .row
    .col-md-6
      .form-group
        %label= t('.filters.state')
        = select_tag :state,
                     options_for_select(Product::STATES.map { |s| [t(s, scope: 'product.states'), s] }, params[:state]),
                     class: 'form-control',
                     include_blank: true
```

### Структура

```haml
.modal#search-form-dropdown(role="dialog")
  .modal-dialog(role="document")
    .modal-content
      = form_tag current_url, method: :get do
        .modal-header
          %button.close(aria-label="Close" aria-hidden="true" data-dismiss="modal" type="button")
            %span(aria-hidden="true") &times;
          %h2.modal-title= title
        .modal-body
          = yield
        .modal-footer
          = button_tag t('operator.base.search_form.button'), class: 'btn btn-primary', type: 'submit', name: nil
```

---

## Menu Helper

### Счётчики в меню

**Расположение:** `app/helpers/operator_menu_helper.rb`

### menu_orders_label

```ruby
menu_orders_label
```

Отображает бейдж с количеством новых заказов.

### menu_bells_label

```ruby
menu_bells_label
```

Отображает бейдж с количеством непрочитанных уведомлений.

### menu_dashboard_label

```ruby
menu_dashboard_label
```

Отображает бейдж с количеством непроверенных элементов дашборда.

### Примеры

```haml
-# В шаблоне меню
%li
  = link_to operator_orders_path do
    = ion_icon 'ios-cart'
    %span Заказы
    = menu_orders_label
```

### menu_label_badge_counter

```ruby
menu_label_badge_counter(count, class_name)
```

Базовый хелпер для создания бейджа счётчика.

| Параметр | Тип | Описание |
|----------|-----|----------|
| `count` | Integer | Число для отображения |
| `class_name` | String | CSS классы бейджа |

---

## Полный пример страницы с поиском

```haml
-# app/views/operator/products/index.html.haml

- content_for :title do
  = t('.title')

.ibox
  .ibox-title
    %h5= t('.title')
    .ibox-tools
      = add_button new_operator_product_path

  -# Header с поиском
  .ibox-header
    = render 'operator/base/header_search',
             result: @products,
             options: @filter.exclude(:query, :page, :per_page),
             query: @filter.query,
             reset_url: (@filter.any? ? operator_products_path : nil),
             counter_title: t('.found', count: @products.total_count),
             one_tag_title: @filter.one_tag_title

  -# Блок фильтров (сворачиваемый)
  #filters.collapse
    .ibox-content.bg-light
      = form_tag operator_products_path, method: :get, class: 'filter-form' do
        .row
          .col-md-3
            .form-group
              %label Категория
              = select_tag :category_id,
                           options_from_collection_for_select(@categories, :id, :title, params[:category_id]),
                           class: 'form-control',
                           include_blank: 'Все категории'
          .col-md-3
            .form-group
              %label Статус
              = select_tag :state,
                           options_for_select(Product.state_options, params[:state]),
                           class: 'form-control',
                           include_blank: 'Все'
          .col-md-4
            .form-group
              %label Поиск
              = text_field_tag :query, params[:query], class: 'form-control', placeholder: 'Название или артикул'
          .col-md-2
            .form-group
              %label &nbsp;
              = submit_tag 'Найти', class: 'btn btn-primary btn-block'

  -# Модальное окно поиска (для клика по header-search)
  = render 'operator/base/search_form_modal', title: t('.filter_title') do
    .row
      .col-md-6
        .form-group
          %label= t('.filters.query')
          = text_field_tag :query, params[:query], class: 'form-control'
      .col-md-6
        .form-group
          %label= t('.filters.category')
          = select_tag :category_id,
                       options_from_collection_for_select(@categories, :id, :title, params[:category_id]),
                       class: 'form-control',
                       include_blank: true

  .ibox-content
    -# Таблица результатов
    - if @products.any?
      %table.table.table-hover
        -# ...
    - else
      .empty-state
        %p= t('.no_results')
```

---

## Паттерны фильтрации

### Filter Object

```ruby
# app/filters/vendor_products_filter.rb
class VendorProductsFilter < VendorBaseFilter
  attribute :query, :string
  attribute :category_id, :integer
  attribute :state, :string

  def options
    {
      query: query,
      category_id: category_id,
      state: state
    }.compact
  end

  def one_tag_title
    return category.title if category_id.present? && query.blank? && state.blank?
    nil
  end

  def exclude(*keys)
    options.except(*keys)
  end
end
```

### Использование в контроллере

```ruby
def index
  @filter = VendorProductsFilter.new(filter_params)
  @products = @filter.apply(current_vendor.products)
                     .page(params[:page])
                     .per(params[:per_page])
end

private

def filter_params
  params.permit(:query, :category_id, :state, :page, :per_page)
end
```

---

## Инструкции для агента

1. **Используй `_header_search`** для всех страниц со списками и фильтрами
2. **Добавляй `_search_form_modal`** для расширенных фильтров
3. **Filter object** должен иметь методы `options`, `one_tag_title`, `exclude`
4. **Reset URL** показывай только когда есть активные фильтры
5. **Counter title** должен быть локализован через I18n
6. **Модальное окно** дублирует основные фильтры для удобства

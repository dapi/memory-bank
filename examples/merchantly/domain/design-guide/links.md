# Ссылки (Links)

Хелперы для создания ссылок на ресурсы.

**Расположение:** `app/helpers/operator_links_helper.rb`

---

## Базовые ссылки

### path_link — Ссылка на путь

```ruby
path_link(path)
```

| Параметр | Тип | Описание |
|----------|-----|----------|
| `path` | String | URL или путь |

**Особенности:**
- Обрезает длинные URL до 60 символов
- Добавляет `target: '_blank'`
- Показывает полный URL в tooltip для обрезанных ссылок

### Примеры

```haml
= path_link @order.callback_url
= path_link 'https://example.com/very/long/path/to/resource'
```

---

### url_link — Ссылка на URL

```ruby
url_link(url, target_blank: true, external: false)
```

| Параметр | Тип | По умолчанию | Описание |
|----------|-----|--------------|----------|
| `url` | String | — | URL |
| `target_blank` | Boolean | true | Открывать в новом окне |
| `external` | Boolean | false | Показывать иконку внешней ссылки |

**Особенности:**
- Автоматически обрезает домен магазина из URL
- Добавляет CSS класс `.external-link`

### Примеры

```haml
= url_link @product.public_url
= url_link @product.public_url, external: true
= url_link @vendor.home_url, target_blank: false
```

---

## Ссылки на ресурсы

### humanized_resource_link — Ссылка на ресурс с названием

```ruby
humanized_resource_link(resource)
```

**Особенности:**
- Использует `resource.title` если доступен
- Иначе генерирует "ModelName#id"
- Использует `resource.operator_path` для URL

### Примеры

```haml
-# Если resource.title = "Товар 1"
= humanized_resource_link @product  # => <a href="/operator/products/1">Товар 1</a>

-# Если resource не имеет title
= humanized_resource_link @log_entry  # => <a href="/operator/log_entries/1">LogEntry#1</a>
```

---

### resource_public_link — Ссылка на публичную страницу

```ruby
resource_public_link(resource, options = {})
```

| Параметр | Тип | Описание |
|----------|-----|----------|
| `resource` | ActiveRecord/String | Ресурс или URL |
| `options` | Hash | Дополнительные опции для link_to |

**Особенности:**
- Для объектов использует `resource.public_url`
- Открывает через `operator_shop_path` (iframe preview)
- Иконка `fa-external-link`

### Примеры

```haml
= resource_public_link @product
= resource_public_link @category
= resource_public_link 'https://shop.example.com/products/1'
```

---

### public_link_to — Публичная ссылка

```ruby
public_link_to(title, url = nil, options = {})
```

| Параметр | Тип | Описание |
|----------|-----|----------|
| `title` | String/Object | Текст или объект с `title` и `public_url` |
| `url` | String | URL (опционально если первый параметр — объект) |
| `options` | Hash | Опции для link_to |

### Примеры

```haml
-# С явным URL
= public_link_to 'Перейти на сайт', @vendor.home_url

-# С объектом
= public_link_to @product  # использует product.title и product.public_url
```

---

## Навигационные ссылки

### back_nav_link — Ссылка "Назад"

```ruby
back_nav_link(title, url)
```

Создаёт навигационную ссылку в формате пилюли (nav-pills).

### Примеры

```haml
= back_nav_link t('shared.back'), operator_products_path
```

**CSS:** `nav nav-pills > li > a`

---

### actions_with_back_button — Кнопка "Назад" в actions

```ruby
actions_with_back_button(title, url)
```

Добавляет кнопку "Назад" в блок `content_for :actions`.

### Примеры

```ruby
# В контроллере или view
actions_with_back_button t('shared.back_to_list'), operator_products_path
```

---

### category_products_link — Ссылка на товары категории

```ruby
category_products_link(category, css_class: nil)
```

| Параметр | Тип | Описание |
|----------|-----|----------|
| `category` | Category | Категория |
| `css_class` | String | CSS классы |

### Примеры

```haml
= category_products_link @category
# => <a href="/operator/products?category_id=1">Товары (42)</a>

= category_products_link @category, css_class: 'btn btn-link'
```

---

## Ссылки для экспорта

### export_orders_to_amocrm_link — Экспорт в AmoCRM

```ruby
export_orders_to_amocrm_link
```

Создаёт ссылку на экспорт заказов в CSV формате для AmoCRM.

---

### export_subscription_emails_to_csv_link — Экспорт подписок

```ruby
export_subscription_emails_to_csv_link
```

Создаёт ссылку на экспорт email-подписок в CSV.

---

## Специальные ссылки

### walletone_merchant_link — Ссылка на WalletOne

```ruby
walletone_merchant_link
```

Внешняя ссылка на личный кабинет WalletOne.

---

### short_url — Сокращённый URL

```ruby
short_url(url, length: 40)
```

Обрезает URL до указанной длины.

### Примеры

```haml
= short_url @product.public_url, length: 30
```

---

### post_action_link_to — POST ссылка-действие

```ruby
post_action_link_to(url, &block)
```

Создаёт ссылку с `method: :post` и стилем label.

### Примеры

```haml
= post_action_link_to activate_operator_product_path(@product) do
  = fa_icon 'check'
  Активировать
```

**CSS:** `label label-outline label-success`

---

## Паттерны использования

### В таблице ресурсов

```haml
%table.table
  - @products.each do |product|
    %tr
      %td= link_to product.title, edit_operator_product_path(product)
      %td= category_products_link product.category if product.category
      %td= resource_public_link product
```

### В карточке ресурса

```haml
.ibox
  .ibox-title
    %h5= @product.title
    .ibox-tools
      = resource_public_link @product
  .ibox-content
    %p
      %strong Категория:
      = humanized_resource_link @product.category
```

### Экспорт данных

```haml
.ibox-tools
  .dropdown
    = subsettings_button
    %ul.dropdown-menu
      %li= export_orders_to_amocrm_link
      %li= export_subscription_emails_to_csv_link
```

### Breadcrumbs

```haml
.breadcrumbs
  = link_to t('operator.dashboard.title'), operator_root_path
  %span.separator /
  = link_to t('operator.products.index.title'), operator_products_path
  %span.separator /
  %span= @product.title
```

---

## Инструкции для агента

1. **Используй `humanized_resource_link`** для ссылок на связанные ресурсы
2. **Используй `resource_public_link`** для ссылок на публичные страницы
3. **Добавляй `target: '_blank'`** для внешних ссылок
4. **Используй `path_link`** для отображения длинных URL
5. **Для экспорта** используй готовые хелперы или создавай аналогичные
6. **Ссылки на категории** делай через `category_products_link` с счётчиком товаров
7. **Внешние ссылки** стилизуй через `external_link` из buttons_helper

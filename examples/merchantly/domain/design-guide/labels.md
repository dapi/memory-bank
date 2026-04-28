---
title: Метки и статусы (Labels)
doc_kind: domain
doc_function: reference
purpose: Компоненты для отображения состояний и меток. Читать при создании UI.
derived_from:
  - ../../../../memory-bank/dna/governance.md
status: active
audience: humans_and_agents
---

# Метки и статусы (Labels)

Компоненты для отображения состояний и меток.

**Основные файлы:**
- `app/helpers/operator_labels_helper.rb` — базовые хелперы
- `app/helpers/operator_form_helper.rb` — `product_state_label`
- `app/helpers/operator/workflow_state_helper.rb` — метки workflow
- `app/helpers/operator_vendor_job_helper.rb` — метки задач

---

## state_label — Базовая метка состояния

**Расположение:** `app/helpers/operator_labels_helper.rb`

### Назначение

Создаёт цветную метку (badge) для отображения состояния.

### Сигнатура

```ruby
state_label(title, type = :info)
```

| Параметр | Тип | По умолчанию | Описание |
|----------|-----|--------------|----------|
| `title` | String | — | Текст метки |
| `type` | Symbol | `:info` | Тип/цвет метки |

### Типы меток

| Тип | CSS класс | Цвет | Использование |
|-----|-----------|------|---------------|
| `:default` | `label-default` | Серый | Неактивные состояния |
| `:primary` | `label-primary` | Синий | Основные действия |
| `:success` | `label-success` | Зелёный | Успешные/активные |
| `:info` | `label-info` | Голубой | Информационные |
| `:warning` | `label-warning` | Жёлтый | Предупреждения |
| `:danger` | `label-danger` | Красный | Ошибки/важное |

### Примеры

```haml
= state_label 'Активен', :success
= state_label 'Ожидает', :warning
= state_label 'Отменён', :danger
= state_label 'Черновик', :default
```

**CSS:** `<span class="label label-{type}">{title}</span>`

---

## resource_moysklad_label — Метка синхронизации

**Расположение:** `app/helpers/operator_labels_helper.rb`

### Назначение

Иконка облака для ресурсов, синхронизированных с МойСклад.

### Сигнатура

```ruby
resource_moysklad_label(resource)
```

### Условие отображения

Метка отображается только если `resource.active_stock_linked?` возвращает `true`.

### Примеры

```haml
%tr
  %td
    = product.title
    = resource_moysklad_label product
```

**CSS:** `fa-cloud text-muted text-small` с tooltip

---

## product_state_label — Метка состояния товара

**Расположение:** `app/helpers/operator_form_helper.rb`

### Назначение

Отображает состояние доступности товара для заказа.

### Сигнатура

```ruby
product_state_label(product)
```

### Состояния

| Состояние | Цвет | Описание |
|-----------|------|----------|
| `:single` | Зелёный | Доступен, один вариант |
| `:multiple` | Зелёный | Доступен, несколько вариантов |
| `:not_available` | Жёлтый | Не доступен для заказа |
| Остальные | Серый | Неактивен |

### Примеры

```haml
%td= product_state_label @product
```

---

## Метки workflow состояний

**Расположение:** `app/helpers/operator/workflow_state_helper.rb`

### workflow_finite_state_label

```ruby
workflow_finite_state_label(finite_state)
```

Отображает метку финального состояния workflow (success/error/pending).

### workflow_state_label

```ruby
workflow_state_label(state)
```

Отображает метку промежуточного состояния workflow.

### Примеры

```haml
%td= workflow_finite_state_label @order.workflow_state.finite_state
%td= workflow_state_label @order.workflow_state
```

---

## Метки заказов

### order_payment_state_label

**Расположение:** `app/helpers/operator/order_payment_helper.rb`

```ruby
order_payment_state_label(order)
```

Отображает состояние оплаты заказа.

### order_reserve_state_label

**Расположение:** `app/helpers/operator/order_reservation_helper.rb`

```ruby
order_reserve_state_label(order)
```

Отображает состояние резервирования товаров.

### Примеры

```haml
%td
  = order_payment_state_label @order
%td
  = order_reserve_state_label @order
```

---

## vendor_job_state_label — Метка состояния задачи

**Расположение:** `app/helpers/operator_vendor_job_helper.rb`

### Сигнатура

```ruby
vendor_job_state_label(state)
```

### Состояния

| Состояние | Описание |
|-----------|----------|
| `pending` | В очереди |
| `working` | Выполняется |
| `complete` | Завершено |
| `failed` | Ошибка |

### Примеры

```haml
= vendor_job_state_label @job.state
```

---

## Создание кастомных меток

### Паттерн для новой метки

```ruby
# app/helpers/operator/my_resource_helper.rb
module Operator::MyResourceHelper
  def my_resource_state_label(resource)
    case resource.state.to_sym
    when :active
      state_label I18n.t('my_resource.states.active'), :success
    when :pending
      state_label I18n.t('my_resource.states.pending'), :warning
    when :disabled
      state_label I18n.t('my_resource.states.disabled'), :default
    else
      state_label resource.state, :info
    end
  end
end
```

### Метка с иконкой

```ruby
def status_with_icon(resource)
  icon = case resource.status
         when 'synced' then 'fa-check'
         when 'pending' then 'fa-clock-o'
         when 'error' then 'fa-exclamation'
         end

  content_tag :span, class: "label label-#{label_type(resource.status)}" do
    fa_icon(icon) + ' ' + resource.status_text
  end
end
```

---

## Паттерны использования

### В таблице

```haml
%table.table
  %thead
    %tr
      %th Название
      %th Статус
      %th Оплата
  %tbody
    - @orders.each do |order|
      %tr
        %td= order.title
        %td= workflow_state_label order.workflow_state
        %td= order_payment_state_label order
```

### В карточке ресурса

```haml
.ibox
  .ibox-title
    %h5
      = @product.title
      = product_state_label @product
      = resource_moysklad_label @product
```

### Группировка меток

```haml
.labels-group
  = state_label 'Новый', :primary
  = state_label 'Срочный', :danger
  = state_label 'VIP', :warning
```

---

## CSS стили меток

```scss
// Базовые метки Bootstrap
.label {
  display: inline-block;
  padding: 2px 6px;
  font-size: 11px;
  font-weight: 600;
  border-radius: 3px;
}

// Outline варианты
.label-outline {
  background: transparent;
  border: 1px solid;
}

// Группировка
.labels-group {
  .label {
    margin-right: 4px;
  }
}
```

---

## Инструкции для агента

1. **Используй `state_label`** для простых статусных меток
2. **Создавай специализированные хелперы** для сложных ресурсов (как `product_state_label`)
3. **Всегда локализуй текст** через `I18n.t`
4. **Выбирай цвет осмысленно:**
   - Зелёный — положительные/активные
   - Жёлтый — предупреждения/ожидание
   - Красный — ошибки/критичное
   - Серый — неактивное/черновик
5. **Добавляй tooltip** для меток с сокращённым текстом
6. **Размещай метки после заголовка** в `.ibox-title` или в отдельной колонке таблицы

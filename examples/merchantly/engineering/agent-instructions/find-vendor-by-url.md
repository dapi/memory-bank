# Инструкция: Поиск магазина по URL в production

## Быстрый поиск магазина по URL

### 1. Извлечение subdomain из URL

Из URL типа `http://shop-c2b637.kiliosk.store/` нужно извлечь часть до `.kiliosk.store` или `.aydamarket.ru`:
- `shop-c2b637.kiliosk.store` → subdomain: `shop-c2b637`
- `myshop.aydamarket.ru` → subdomain: `myshop`

### 2. Поиск по subdomain

```bash
cd ~/code/brandymint/infra && direnv exec . ./bin/rails-runner merchantly \
  'v = Vendor.find_by(subdomain: "SUBDOMAIN_HERE"); \
   if v; \
     puts "ID: #{v.id}"; \
     puts "Title: #{v.title.presence || "-"}"; \
     puts "Subdomain: #{v.subdomain}"; \
     puts "Domain: #{v.domain}"; \
     puts "Active: #{v.active_domain}"; \
     puts "State: #{v.state}"; \
     puts "Created: #{v.created_at}"; \
     puts ""; \
     v.owners.each {|m| o = m.operator; puts "Owner: #{o.email} (#{o.name}), phone: #{o.phone}"}; \
     puts ""; \
     puts "Products: #{v.products_count}"; \
     puts "Orders: #{v.orders_count}"; \
   else; \
     puts "Vendor not found"; \
   end'
```

### 3. Альтернативный поиск по кастомному домену

Если магазин использует кастомный домен (например, `myshop.com`):

```bash
cd ~/code/brandymint/infra && direnv exec . ./bin/rails-runner merchantly \
  'v = Vendor.find_by_host("DOMAIN_HERE"); \
   puts v ? "Found: ID=#{v.id}, subdomain=#{v.subdomain}" : "Not found"'
```

### 4. Поиск через Vendor.find_by_domain

Метод `Vendor.find_by_domain(domain)` автоматически обрабатывает различные варианты доменов:

```bash
cd ~/code/brandymint/infra && direnv exec . ./bin/rails-runner merchantly \
  'v = Vendor.find_by_domain("shop-c2b637.kiliosk.store"); \
   puts v ? "Found: ID=#{v.id}" : "Not found"'
```

## Важные замечания

1. **Кавычки**: Используй одинарные кавычки для внешнего Ruby-кода, чтобы избежать проблем с экранированием
2. **Subdomain extraction**: Всегда извлекай только часть до первого домена системы
3. **Multiple zones**: Система поддерживает разные зоны (kiliosk.store, aydamarket.ru и др.)
4. **Custom domains**: Магазин может работать как на subdomain, так и на кастомном домене

## Структура Vendor

### Поля домена:
- `subdomain` - поддомен на системных доменах (shop-123)
- `domain` - подтвержденный кастомный домен
- `suggested_domain` - предложенный, но не подтвержденный домен
- `active_domain` - метод, возвращающий активный домен (domain || subdomain + зона)

### Ассоциации:
- `vendor.owners` - Members с ролью OWNER
- `member.operator` - связанный Operator (содержит email, name, phone)

## Примеры использования

### Пример 1: Базовая информация
```bash
cd ~/code/brandymint/infra && direnv exec . ./bin/rails-runner merchantly \
  'v = Vendor.find_by(subdomain: "shop-c2b637"); \
   puts "ID: #{v.id}, Title: #{v.title}, State: #{v.state}"'
```

### Пример 2: Полная информация с владельцами
```bash
cd ~/code/brandymint/infra && direnv exec . ./bin/rails-runner merchantly \
  'v = Vendor.find(11204); \
   puts "=== Vendor #{v.id} ==="; \
   puts "Subdomain: #{v.subdomain}"; \
   puts "Active domain: #{v.active_domain}"; \
   v.owners.each {|m| puts "Owner: #{m.operator.email}"}'
```

### Пример 3: Поиск и вывод JSON
```bash
cd ~/code/brandymint/infra && direnv exec . ./bin/rails-runner merchantly \
  'v = Vendor.find_by(subdomain: "shop-c2b637"); \
   puts v.attributes.slice("id", "title", "subdomain", "domain", "state").to_json'
```

## Troubleshooting

### Ошибка: "undefined method 'public_path'"
Vendor не имеет поля `public_path`. Используйте `subdomain` вместо этого.

### Ошибка: "undefined method 'owner'"
Vendor не имеет метода `owner`. Используйте `vendor.owners` (множественное число) для получения списка владельцев через Members.

### Ошибка: "undefined constant VendorDomain"
Модели `VendorDomain` не существует. Используйте `Vendor.find_by_domain(host)` или `Vendor.find_by_host(host)`.

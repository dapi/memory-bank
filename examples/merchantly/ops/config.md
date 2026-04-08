# Configuration Guide

Полная документация по конфигурации проекта Merchantly.

## Дополнительная документация

| Тема | Документ |
|------|----------|
| Нагрузочное тестирование | [load-testing-spec.md](load-testing-spec.md) |
| Скрипт нагрузочного тестирования | [scripts/load_test.rb](../scripts/load_test.rb) |

## Архитектура конфигурации

Merchantly использует [anyway_config](https://github.com/palkan/anyway_config) gem для управления конфигурацией.

### Структура файлов

```
config/
├── configs/
│   └── application_config.rb   # Anyway::Config класс с типизированными настройками
├── application.yml             # YAML конфиг с defaults и per-environment секциями
├── database.yml                # Конфигурация БД
├── cable.yml                   # ActionCable (Redis)
├── queue.yml                   # SolidQueue
└── puma.rb                     # Puma web server
```

### Как работает anyway_config

1. **Класс конфигурации**: `ApplicationConfig` в `config/configs/application_config.rb`
2. **YAML файл**: `config/application.yml`
3. **ENV переменные**: Автоматически с префиксом `APP_`

```ruby
# Доступ к настройкам
ApplicationConfig.redis[:url]           # из YAML или APP_REDIS__URL
ApplicationConfig.feature_billing       # из YAML или APP_FEATURE_BILLING
ApplicationConfig.sms_qtelecom_login    # из YAML или APP_SMS_QTELECOM_LOGIN
```

### Naming Convention для ENV

| YAML структура | ENV переменная |
|----------------|----------------|
| `redis.url` | `APP_REDIS__URL` |
| `feature_billing` | `APP_FEATURE_BILLING` |
| `telegram.token` | `APP_TELEGRAM__TOKEN` |
| `sms_qtelecom_login` | `APP_SMS_QTELECOM_LOGIN` |

**Правила:**
- Префикс = `APP_` (настроено через `env_prefix :app`)
- Вложенность разделяется `__` (double underscore)
- Массивы: `APP_IDS=1,2,3` → `ids: [1,2,3]`
- **НЕ нужно** добавлять ERB-вставки `<%= ENV['...'] %>` в YAML — anyway_config сам подхватывает переменные

---

## Переменные окружения

### Redis & Cache

| Переменная | Описание | По умолчанию |
|------------|----------|--------------|
| `APP_REDIS__URL` | URL Redis (общий кэш) | `redis://127.0.0.1:6379/0` |
| `ACTION_CABLE_REDIS_URL` | URL Redis для ActionCable | `redis://localhost:6379/1` |
| `APP_ONE_OFF_TOKENS_REDIS__URL` | URL Redis для одноразовых токенов | `redis://127.0.0.1:6379/3` |
| `MEMCACHE_SERVERS` | Адреса Memcache серверов | `localhost` |

### Database

| Переменная | Описание | По умолчанию |
|------------|----------|--------------|
| `DATABASE_HOST` | Хост PostgreSQL | `localhost` |
| `DATABASE_PORT` | Порт PostgreSQL (5433 для PgBouncer) | `5432` |
| `DATABASE_NAME` | Имя базы данных | `kiiiosk_development` |
| `DATABASE_USER` | Пользователь БД | — |
| `DATABASE_PASS` | Пароль БД | — |
| `DB_ADAPTER` | Адаптер БД | `postgresql` |
| `RAILS_MAX_THREADS` | Макс. потоков (DB pool: `5`, Puma threads: `1`) | `5` / `1` |
| `WEB_CONCURRENCY` | Кол-во Puma workers | `1` |
| `DATABASE_POOL_SIZE` | Размер connection pool (переопределяет расчет) | `(RAILS_MAX_THREADS × WEB_CONCURRENCY) + 2`, мин. 15 |

### Application

| Переменная | Описание | По умолчанию |
|------------|----------|--------------|
| `PORT` | Порт приложения | `3000` |
| `APP_DEFAULT_LOCALE` | Локаль по умолчанию | `ru` |
| `APP_DEFAULT_CURRENCY` | Валюта по умолчанию | `RUB` |
| `APP_INDEX` | Индекс приложения | `default` |

### Application/Domain & URL Generation

Настройки для генерации URL в приложении. Rails использует `default_url_options` для ActionMailer, ActionController и URL helpers.

| Переменная | Описание | По умолчанию |
|------------|----------|--------------|
| `APP_DEFAULT_URL_OPTIONS__HOST` | Хост для генерации URL | `kiiiosk.store` |
| `APP_DEFAULT_URL_OPTIONS__PROTOCOL` | Протокол (http/https) | `https` |
| `APP_DEFAULT_URL_OPTIONS__PORT` | Порт (если нестандартный) | — |
| `APP_TLD_LENGTH` | TLD length (авто-расчёт из host если не указан) | авто |
| `APP_ASSET_URL` | URL для ассетов | `https://assets.kiiiosk.store` |
| `APP_DNS_IP` | IP для DNS проверки доменов | — |
| `APP_EXTRA_DNS_IPS` | Дополнительные допустимые IP (через запятую) | — |

**Структура `default_url_options` в YAML:**

```yaml
default_url_options:
  host: kiiiosk.store      # Обязательный: доменное имя без протокола
  protocol: https          # Протокол: http или https
  port: 3000               # Опционально: порт (не указывать для 80/443)
```

**Авто-вычисление `tld_length`:**

`tld_length` определяет сколько сегментов справа считаются TLD (top-level domain). Вычисляется автоматически из `host`:
- `kiiiosk.store` (1 точка) → `tld_length: 1`
- `app.kiiiosk.store` (2 точки) → `tld_length: 2`
- `stage.kiiiosk.store` (2 точки) → `tld_length: 2`

Можно переопределить явно через `APP_TLD_LENGTH`.

**Доступ в коде:**

```ruby
ApplicationConfig.default_url_options[:host]     # => "kiiiosk.store"
ApplicationConfig.default_url_options[:protocol] # => "https"
ApplicationConfig.tld_length                     # => 1 (авто-вычислено)
```

### AWS/S3

| Переменная | Описание | По умолчанию |
|------------|----------|--------------|
| `APP_AWS_ACCESS_KEY_ID` | Access Key AWS | — |
| `APP_AWS_SECRET_ACCESS_KEY` | Secret Key AWS | — |
| `APP_AWS_REGION` | Регион AWS | — |
| `APP_AWS_BUCKET` | Имя S3 bucket | — |
| `APP_AWS_ENDPOINT` | Custom S3 endpoint | — |
| `APP_AWS_OBJECT_ACL` | ACL для загружаемых объектов (`public-read` для AWS, пусто для Selectel) | — |
| `APP_CARRIERWAVE_STORAGE` | Тип хранилища (`file`/`aws`) | — |
| `APP_STATIC_FILE_CACHE_TTL` | HTTP cache TTL для найденных статических файлов (секунды) | `10800` (3ч) |

### Email/SMTP

| Переменная | Описание | По умолчанию |
|------------|----------|--------------|
| `DELIVERY_METHOD` | Метод доставки email | `letter_opener_web` |
| `SMTP_HOST` | SMTP сервер | `smtp.yandex.ru` |
| `SMTP_USERNAME` | SMTP логин | — |
| `SMTP_PASSWORD` | SMTP пароль | — |
| `SMTP_DOMAIN` | SMTP домен | `yandex.ru` |
| `MAIL_SERVER_IP` | IP mail сервера (для SPF) | — |

### Telegram

| Переменная | Описание | По умолчанию |
|------------|----------|--------------|
| `APP_TELEGRAM__TOKEN` | Токен операторского бота (глобальный) | `123:fake` |
| `APP_TELEGRAM__USERNAME` | Username операторского бота | `kiiiosk_operator_bot` |
| `APP_TELEGRAM__REDIS_URL` | Redis URL для Telegram | `redis://127.0.0.1:6379/2` |

> **Примечание:** Токены для Mini App хранятся в БД (`vendors.telegram_bot_token`, зашифрован) и настраиваются в операторской панели.

### Платёжные системы

| Переменная | Описание | По умолчанию |
|------------|----------|--------------|
| `APP_CLOUDPAYMENTS_PUBLIC_KEY` | Public key CloudPayments | — |
| `APP_CLOUDPAYMENTS_SECRET_KEY` | Secret key CloudPayments | — |
| `APP_GSDK_TID` | Terminal ID GSDK | — |
| `APP_GSDK_MID` | Merchant ID GSDK | — |
| `APP_GSDK_TEST` | Тестовый режим GSDK | — |
| `APP_GSDK_SECRET_CODE` | Secret code GSDK | — |
| `APP_TINKOFF_API_TOKEN` | API токен Tinkoff Business | — |
| `APP_TINKOFF_API_ACCOUNT` | Лицевой счёт Tinkoff | — |
| `APP_W1_CLIENT_ID` | WalletOne Client ID | — |
| `APP_W1_CALLBACK_PATH` | WalletOne Callback Path | `/auth/walletone/callback` |
| `APP_W1_TOKEN` | WalletOne Token | — |

### Онлайн-кассы

| Переменная | Описание | По умолчанию |
|------------|----------|--------------|
| `APP_CLOUD_KASSIR_TEST_EMAIL` | Email для тестов Cloud Kassir | — |
| `APP_CLOUD_KASSIR_TEST_PUBLIC_KEY` | Public key для тестов Cloud Kassir | — |
| `APP_CLOUD_KASSIR_TEST_SECRET_KEY` | Secret key для тестов Cloud Kassir | — |
| `APP_ORANGE_DATA_CLIENT_CRT_PATH` | Путь к сертификату Orange Data | `config/orange_data/client.crt` |
| `APP_ORANGE_DATA_CLIENT_KEY_PATH` | Путь к ключу Orange Data | `config/orange_data/client.key` |
| `APP_ORANGE_DATA_PRIVATE_KEYPATH` | Путь к приватному ключу | `config/orange_data/private_key.pem` |
| `APP_ORANGE_DATA_CLIENT_KEY_PASSWORD` | Пароль ключа Orange Data | `1234` |

### Маппинг ставок НДС по провайдерам

Конфигурация соответствия внутренних кодов НДС и кодов, отправляемых в платёжные системы.

| Секция YAML | Описание |
|-------------|----------|
| `payment_providers_vat` | Маппинг ставок НДС по провайдерам |
| `payment_providers_vat.life_pay` | Коды для LifePay (символы) |
| `payment_providers_vat.aqsi` | Коды для Aqsi (числа) |
| `payment_providers_vat.starrys` | Коды для Starrys (числа) |
| `payment_providers_vat.kassatka` | Коды для Kassatka (числа) |
| `payment_providers_vat.yandex_kassa` | Коды для YandexKassa (числа) |
| `payment_providers_vat.cloud_kassir` | Коды для CloudKassir (числа, nil) |
| `payment_providers_vat.cloud_payments` | Коды для CloudPayments (числа, nil) |
| `payment_providers_vat.yoo_money` | Коды для YooMoney (числа→строки) |

Список внутренних налоговых ключей задаётся в `taxes_list` (в `config/application.yml`), включая:
`tax_ru_1..tax_ru_11`, `custom_10..custom_13`.

```ruby
# Доступ через централизованный сервис
PaymentProviderTaxMapping.tax_id(:life_pay, :tax_ru_7)
# => :vat20

PaymentProviderTaxMapping.tax_id(:yoo_money, 'tax_ru_1')
# => 1  (конвертируется в "1" в YooMoney модуле)
```

Подробнее: [Документация маппинга ставок НДС](payment-providers-vat-mapping.md)

### SMS провайдеры

Все SMS-переменные мигрированы на `ApplicationConfig` с префиксом `APP_`:

| Новая переменная | Описание | По умолчанию |
|------------------|----------|--------------|
| `APP_SMS_SMSC_LOGIN` | Логин SMSC | — |
| `APP_SMS_SMSC_PASSWORD` | Пароль SMSC | — |
| `APP_SMS_SMSC_SENDER` | Имя отправителя SMSC | `Kiiiosk` |
| `APP_SMS_SMSC_FAKE` | Фейк-режим SMSC | `false` |
| `APP_SMS_QTELECOM_LOGIN` | Логин QTelecom | — |
| `APP_SMS_QTELECOM_PASSWORD` | Пароль QTelecom | — |
| `APP_SMS_QTELECOM_SENDER` | Имя отправителя QTelecom | `''` |
| `APP_SMS_QTELECOM_FAKE` | Фейк-режим QTelecom | `false` |
| `APP_SMS_UNIFONIC_LOGIN` | Логин Unifonic | — |
| `APP_SMS_UNIFONIC_PASSWORD` | Пароль Unifonic | — |
| `APP_SMS_UNIFONIC_SENDER` | Имя отправителя Unifonic | — |
| `APP_SMS_TOP_USERNAME` | Логин SMS Top | — |
| `APP_SMS_TOP_PASSWORD` | Пароль SMS Top | — |
| `APP_SMS_TOP_FAKE` | Фейк-режим SMS Top | `false` |

```ruby
# Доступ в коде
ApplicationConfig.sms_qtelecom_login
ApplicationConfig.sms_smsc_password
```

### OAuth провайдеры

| Переменная | Описание | По умолчанию |
|------------|----------|--------------|
| `APP_VKONTAKTE_APP_ID` | VK App ID | — |
| `APP_VKONTAKTE_APP_SECRET` | VK App Secret | — |
| `APP_FACEBOOK_APP_ID` | Facebook App ID | — |
| `APP_FACEBOOK_APP_SECRET` | Facebook App Secret | — |
| `APP_YANDEX_APP_ID` | Yandex App ID | — |
| `APP_YANDEX_APP_SECRET` | Yandex App Secret | — |

### Внешние сервисы

| Переменная | Описание | По умолчанию |
|------------|----------|--------------|
| `APP_THUMBOR_URL` | URL Thumbor сервера | `https://thumbor.kiiiosk.store` |
| `APP_THUMBOR_SECRET_KEY` | Secret key Thumbor | — |
| `APP_DADATA_API_KEY` | API key DaData | — |
| `APP_DADATA_SECRET_KEY` | Secret key DaData | — |
| `APP_GOOGLE_MAP_API_KEY` | API key Google Maps | — |
| `APP_FAST_FOREX_API_KEY` | API key FastForex (курсы валют) | — |
| `APP_VIBER_AUTH_TOKEN` | Auth token Viber | — |
| `APP_AMOCRM_USE` | Включить интеграцию AmoCRM | `false` |
| `APP_AMOCRM_URL` | URL AmoCRM | — |
| `APP_AMOCRM_LOGIN` | Логин AmoCRM | — |
| `APP_AMOCRM_APIKEY` | API key AmoCRM | — |
| `APP_KEYCLOAK_FILE_PATH` | Путь к Keycloak конфигу | — |

### Security

| Переменная | Описание | По умолчанию |
|------------|----------|--------------|
| `APP_SORCERY_SALT_JOIN_TOKEN` | Salt для Sorcery | `-@%%*&-` |
| `APP_SORCERY_ENCRYPTION_KEY` | Ключ шифрования Sorcery | `enc` |
| `APP_COOKIE_KEY` | Ключ cookie сессии | `_kiiiosk` |
| `APP_METRICS_KEY` | Ключ доступа к метрикам | `1234` |
| `APP_FETCH_DOMAINS_API_KEY` | API ключ для получения доменов | — |

### Session Configuration

Настройки времени жизни сессий и cookies.

| Переменная | Описание | По умолчанию |
|------------|----------|--------------|
| `APP_SESSION_COOKIE_EXPIRE_AFTER` | Время жизни session cookie (секунды). Важно для пользовательских корзин | `631152000` (20 лет) |
| `APP_REMEMBER_ME_FOR` | Время жизни "remember me" токена Sorcery для операторов (секунды) | `2592000` (30 дней) |

**Примечания:**
- `session_cookie_expire_after` — делает cookie persistent (не session cookie). Без этого параметра cookie удаляется при закрытии браузера. 20 лет = практически бессрочно.
- `remember_me_for` — используется Sorcery для операторских сессий. Определяет как долго оператор остаётся залогиненным без повторной аутентификации.

```ruby
# Доступ в коде
ApplicationConfig.session_cookie_expire_after  # => 631152000 (секунды)
ApplicationConfig.remember_me_for              # => 2592000 (секунды)
```

### Active Record Encryption

Шифрование чувствительных данных в БД (например, SMTP паролей вендоров).

| Переменная | Описание | По умолчанию |
|------------|----------|--------------|
| `ACTIVE_RECORD_ENCRYPTION_PRIMARY_KEY` | Основной ключ шифрования (base64, 32+ байт) | — |
| `ACTIVE_RECORD_ENCRYPTION_DETERMINISTIC_KEY` | Ключ детерминированного шифрования | — |
| `ACTIVE_RECORD_ENCRYPTION_KEY_DERIVATION_SALT` | Соль для деривации ключей | — |
| `ACTIVE_RECORD_ENCRYPTION_PREVIOUS_PRIMARY_KEY` | Предыдущий ключ (для ротации) | — |

**Генерация ключей:**

```bash
bin/rails db:encryption:init
```

**Использование:**

```ruby
# В модели
class VendorSmtpSettings < ApplicationRecord
  encrypts :password
end
```

**Важно:**
- Без этих переменных шифрование НЕ работает (warning в логах при старте)
- Подробная инструкция: [docs/sre-smtp-password-encryption.md](sre-smtp-password-encryption.md)

### Puma & Server

| Переменная | Описание | По умолчанию |
|------------|----------|--------------|
| `CONTROL_BIND` | Puma control socket (dev: `unix://...`, prod: `tcp://127.0.0.1:9294`) | — |
| `CONTROL_AUTH_TOKEN` | Puma control token | `CHANGEME` |
| `PIDFILE` | Путь к PID файлу Puma | — |
| `SOLID_QUEUE_IN_PUMA` | Включить SolidQueue plugin в Puma | — |
| `JOB_CONCURRENCY` | Кол-во SolidQueue процессов | `2` |
| `PROMETHEUS_EXPORTER_PORT` | Порт экспорта Prometheus метрик | `9394` |

### SolidQueue

Конфигурация SolidQueue — фоновой очереди задач.

| Переменная | Описание | По умолчанию |
|------------|----------|--------------|
| `APP_SOLID_QUEUE_RETENTION_HOURS` | Время хранения выполненных jobs (часы) | `24` |

**Retention Policy:**

Finished jobs автоматически удаляются recurring task-ом `clear_solid_queue_finished_jobs` (каждый час).
Параметр `solid_queue_retention_hours` определяет, сколько часов хранить выполненные задачи.

С ~800K jobs/день:
- 24 часа → ~800K записей в таблице
- 6 часов → ~200K записей в таблице

```ruby
# Доступ в коде
ApplicationConfig.solid_queue_retention_hours  # => 24
```

**Первоначальная очистка:**

При большом количестве накопленных записей запустите скрипт очистки:

```bash
bundle exec rails runner scripts/cleanup_solid_queue_jobs.rb
```

### Feature Flags & Debug

| Переменная | Описание | По умолчанию |
|------------|----------|--------------|
| `APP_ENABLE_ANALYTICS` | Включить аналитику | `false` |
| `APP_ENABLE_WALLETONE` | Включить WalletOne | `false` |
| `APP_USE_LOCAL_RATES` | Локальные курсы валют | `false` |
| `APP_ALLOW_SCREENSHOTING` | Разрешить скриншоты | `false` |
| `APP_NO_CANONICAL_REDIRECT` | Отключить канонический редирект | `false` |
| `APP_RACK_MINI_PROFILER` | Включить Rack Mini Profiler | `false` |
| `APP_PERFORM_CACHING` | Кэширование (dev) | `false` |
| `APP_API_TRAFFIC_LOGGER` | Логировать API трафик | `false` |
| `APP_TRACE_LOGGING_LIMIT` | Лимит трейс-логов | `3` |
| `APP_ENABLE_RACK_CACHE` | Включить Rack Cache | `false` |
| `APP_HTTP_CACHE_MUST_REVALIDATE` | HTTP cache must-revalidate | `false` |
| `APP_ALLOW_PRODUCT_SYNC_DUP` | Синхронное дублирование товаров | `false` |
| `APP_DONT_SKIP_MAKE_ORDER` | Не пропускать make_order в dev | `false` |
| `APP_MULTIPLE_CATEGORY_IDS_FOR_YANDEX_MARKET` | Множественные категории для ЯМ | `false` |
| `APP_GSDK_ENABLED` | Включить GSDK платежи | `false` |

### Application Feature Flags

Флаги включения/выключения функционала приложения.

| Переменная | Описание | По умолчанию |
|------------|----------|--------------|
| `APP_FEATURE_WALLETONE` | Включить WalletOne | `false` |
| `APP_FEATURE_SHOW_ACCESS_KEYS` | Показывать ключи доступа | `true` |
| `APP_FEATURE_AFFILATE_PROGRAM` | Партнёрская программа | `true` |
| `APP_FEATURE_PRODUCT_PROPERTIES` | Свойства товаров | `true` |
| `APP_FEATURE_PRODUCT_ARTICLE` | Артикулы товаров | `true` |
| `APP_FEATURE_EDIT_VENDOR_EMAIL` | Редактирование email вендора | `true` |
| `APP_FEATURE_ADD_STORE` | Добавление магазина | `true` |
| `APP_FEATURE_VENDOR_ARCHIVE` | Архив вендоров | `true` |
| `APP_FEATURE_BILLING` | Биллинг | `true` |
| `APP_FEATURE_VALIDATE_DOMAIN_DNS_IP` | Валидация DNS IP домена | `false` |
| `APP_FEATURE_AUTO_ATTACH_DOMAIN` | Авто-привязка домена | `false` |
| `APP_FEATURE_K8S_CUSTOM_DOMAINS` | K8s custom domains | `false` |
| `APP_FEATURE_API_AUTH_BY_LOGIN_AND_PASSWORD` | API auth по логину/паролю | `true` |
| `APP_FEATURE_INVITE_AUTO_CREATE_OPERATOR` | Авто-создание оператора по инвайту | `false` |
| `APP_FEATURE_MEMBER_SEND_PIN_CODE` | Отправка PIN-кода | `false` |
| `APP_FEATURE_OPERATOR_REQUIRE_EMAIL_CONFIRMATION` | Требовать подтверждение email оператора | `true` |
| `APP_FEATURE_OPERATOR_REQUIRE_PHONE_CONFIRMATION` | Требовать подтверждение телефона оператора | `true` |
| `APP_FEATURE_CREATE_DEFAULT_PAYMENTS_AND_DELIVERIES` | Создавать платежи/доставки по умолчанию | `false` |
| `APP_FEATURE_UNIQUE_CATEGORY_TITLE` | Уникальные названия категорий | `true` |
| `APP_FEATURE_SHOW_DASHBOARD` | Показывать дашборд | `true` |
| `APP_FEATURE_SHOW_PREVIEW_PAGE` | Показывать страницу предпросмотра | `true` |
| `APP_FEATURE_SHOW_PRODUCTS_TABS` | Показывать табы товаров | `true` |
| `APP_FEATURE_SHOW_ORDERS_MORE` | Показывать "ещё" в заказах | `true` |
| `APP_FEATURE_SHOW_MANAGEMENT` | Показывать управление | `true` |
| `APP_FEATURE_SHOW_ORDER_SETTINGS` | Показывать настройки заказов | `true` |
| `APP_FEATURE_SHOW_SUBSCRIPTIONS_SETTINGS` | Показывать настройки подписок | `true` |

### Billing IDs

Идентификаторы для интеграции с биллинговой системой Openbill.

| Переменная | Описание | По умолчанию |
|------------|----------|--------------|
| `APP_BILLING_MOYSKLAD` | ID МойСклад | — |
| `APP_BILLING_SMS` | ID SMS | — |
| `APP_BILLING_ADDITIONAL_WORKS` | ID доп. работ | — |
| `APP_BILLING_SYSTEM_CATEGORY_ID` | ID системной категории | — |
| `APP_BILLING_CLIENTS_CATEGORY_ID` | ID категории клиентов | — |
| `APP_BILLING_CLIENTS_SMS_CATEGORY_ID` | ID SMS категории клиентов | — |
| `APP_BILLING_PARTNERS_CATEGORY_ID` | ID категории партнёров | — |
| `APP_BILLING_CLOUDPAYMENTS_ACCOUNT_ID` | ID аккаунта CloudPayments | — |
| `APP_BILLING_IP_PISMENNY_ACCOUNT_ID` | ID аккаунта IP Письменный | — |
| `APP_BILLING_WALLETONE_ACCOUNT_ID` | ID аккаунта WalletOne | — |
| `APP_BILLING_PARTNER_ACCOUNT_ID` | ID партнёрского аккаунта | — |
| `APP_BILLING_SUBSCRIPTIONS_ACCOUNT_ID` | ID аккаунта подписок | — |
| `APP_BILLING_EXTERNAL_LINK_KIOSK_ACCOUNT_ID` | ID аккаунта внешних ссылок | — |
| `APP_BILLING_ADDITIONAL_WORKS_ACCOUNT_ID` | ID аккаунта доп. работ | — |
| `APP_BILLING_ADDITIONAL_WORKS_SERVICE_ID` | ID сервиса доп. работ | — |
| `APP_BILLING_GIFT_ACCOUNT_ID` | ID аккаунта подарков | — |
| `APP_BILLING_GSDK_ACCOUNT_ID` | ID аккаунта GSDK | — |

### Production Features

Флаги для production-функционала. В development/test по умолчанию `false`, в production/staging/beta — `true`.

| Переменная | Описание | По умолчанию |
|------------|----------|--------------|
| `APP_USE_X_ACCEL_REDIRECT` | X-Accel-Redirect для nginx file serving | `false` |
| `APP_FORCE_SSL` | Принудительный SSL редирект в operator panel | `false` |
| `APP_ASYNC_CSS_UPLOAD` | Асинхронная загрузка CSS в S3 | `false` |
| `APP_PAYMENT_AUTOSUBMIT` | Автосабмит платёжных форм | `false` |
| `APP_SEND_CONFIRMATION_EMAILS` | Отправка confirmation emails операторам | `false` |
| `APP_PRE_CREATE_VENDORS` | Pre-create vendors job | `false` |
| `APP_COLLECT_MARKETING_STATS` | Сбор маркетинговой статистики | `false` |
| `APP_SEND_ORANGE_DATA_RECEIPTS` | Отправка Orange Data чеков | `false` |
| `APP_ENABLE_BUGSNAG` | Bugsnag error tracking (frontend) | `false` |

### Payment Systems Production Modes

Режимы production для платёжных систем. При `false` используются sandbox/test URLs.

| Переменная | Описание | По умолчанию |
|------------|----------|--------------|
| `APP_W1_PRODUCTION_MODE` | WalletOne production mode | `false` |
| `APP_STARRYS_PRODUCTION_MODE` | Starrys (онлайн-касса) production mode | `false` |
| `APP_PAYPAL_PRODUCTION_MODE` | PayPal production mode | `false` |

### SMS Fake Modes

| Переменная | Описание | По умолчанию |
|------------|----------|--------------|
| `APP_SMS_QTELECOM_FAKE` | QTelecom fake mode | `false` |
| `APP_SMS_SMSC_FAKE` | SMSC fake mode | `false` |
| `APP_SMS_TOP_FAKE` | SMS Top fake mode | `false` |
| `APP_SMS_UNIFONIC_FAKE` | Unifonic fake mode | `false` |

### Timeouts & Limits

| Переменная | Описание | По умолчанию |
|------------|----------|--------------|
| `APP_CLIENT_REQUEST_TIMEOUT` | Timeout для Client модели (секунды) | `60` |
| `APP_PHONE_CONFIRMATION_TIMEOUT` | Timeout для PhoneConfirmation (секунды) | `60` |
| `APP_DIGITAL_DOWNLOAD_EXPIRATION` | Expiration digital downloads (секунды) | `3600` |
| `APP_DIGITAL_MAX_GENERATES` | Макс. генераций digital ключей | `3` |
| `APP_MAX_PRE_CREATED_VENDORS` | Макс. pre-created vendors | `3` |

### Vendor SMTP

Конфигурация для собственных SMTP-серверов вендоров.

| Переменная | Описание | По умолчанию |
|------------|----------|--------------|
| `APP_SMTP_BLOCKED_IP_RANGES` | Заблокированные IP диапазоны для SSRF защиты (через запятую) | приватные сети |
| `APP_SMTP_VERIFICATION_COOLDOWN_SECONDS` | Cooldown между попытками верификации SMTP (секунды) | `10` |

**Примечания:**
- Если `smtp_blocked_ip_ranges` не указан или пуст, используются дефолтные диапазоны (приватные сети RFC 1918, loopback, link-local и т.д.)
- Дефолтные заблокированные диапазоны: `127.0.0.0/8`, `10.0.0.0/8`, `172.16.0.0/12`, `192.168.0.0/16`, `169.254.0.0/16`, `::1/128`, `fc00::/7`, `fe80::/10`
- Cooldown применяется per-vendor для защиты от спама верификаций

```ruby
# Доступ в коде
ApplicationConfig.smtp_blocked_ip_ranges              # => [] (raw config value)
ApplicationConfig.smtp_blocked_ip_ranges_with_defaults # => ["127.0.0.0/8", "10.0.0.0/8", ...]
ApplicationConfig.smtp_verification_cooldown_seconds  # => 10
```

### Kubernetes Custom Domains

Конфигурация для управления пользовательскими доменами через Kubernetes Ingress.

| Переменная | Описание | По умолчанию |
|------------|----------|--------------|
| `APP_K8S_NAMESPACE` | Kubernetes namespace для Ingress ресурсов | `merchantly` |
| `APP_K8S_CLUSTER_ISSUER` | cert-manager ClusterIssuer для TLS сертификатов | `letsencrypt-http` |
| `APP_K8S_SERVICE_NAME` | Kubernetes сервис для маршрутизации трафика | `merchantly-web` |
| `APP_K8S_SERVICE_PORT` | Порт Kubernetes сервиса | `3000` |
| `APP_K8S_INGRESS_CLASS` | Класс Ingress (nginx, traefik и т.д.) | `nginx` |
| `APP_K8S_PROXY_BODY_SIZE` | Макс. размер тела запроса для proxy | `100m` |
| `APP_K8S_CERTIFICATE_TIMEOUT` | Timeout ожидания готовности TLS сертификата (секунды) | `1200` |
| `APP_K8S_CERTIFICATE_CHECK_INTERVAL` | Интервал проверки статуса сертификата (секунды) | `30` |
| `APP_K8S_UNINSTALL_DELAY` | Задержка перед удалением Ingress при удалении домена (секунды) | `300` |
| `APP_K8S_MIGRATION_INTERVAL` | Интервал между миграциями доменов в rake task (секунды) | `120` |

**Примеры использования в коде:**

```ruby
# Доступ к K8s конфигурации
ApplicationConfig.k8s_namespace           # => "merchantly"
ApplicationConfig.k8s_cluster_issuer      # => "letsencrypt-http"
ApplicationConfig.k8s_service_name        # => "merchantly-web"
ApplicationConfig.k8s_service_port        # => 3000

# Для переведённых временных параметров (в секундах)
ApplicationConfig.k8s_certificate_timeout      # => 1200
ApplicationConfig.k8s_certificate_check_interval # => 30
```

**Задача из rake:**

```bash
# Миграция существующих доменов
bundle exec rake k8s:migrate_domains

# Проверка статуса
bundle exec rake k8s:reconciliation_status

# Список управляемых Ingress
bundle exec rake k8s:list_ingresses
```

### Rails стандартные переменные

Эти переменные НЕ используют APP_ префикс (стандарт Rails/Ruby):

| Переменная | Описание | По умолчанию |
|------------|----------|--------------|
| `RAILS_ENV` | Окружение Rails | `development` |
| `RAILS_LOG_TO_STDOUT` | Логирование в stdout | — |
| `RAILS_SERVE_STATIC_FILES` | Отдача статики Rails | — |
| `BUGSNAG_API_KEY` | API key Bugsnag (error tracking) | — |
| `DISABLE_AR_LOGGER` | Отключить AR логирование (dev) | — |
| `DISABLE_AV_LOGGER` | Отключить ActionView логирование (dev) | — |
| `ASSETS_PRECOMPILE` | Режим прекомпиляции ассетов | — |
| `EAGER_LOAD` | Eager load классов (test) | — |

### Branding & UI

| Переменная | Описание | По умолчанию |
|------------|----------|--------------|
| `APP_OPERATOR_LOGO` | Логотип оператора | `avalab` |
| `APP_PURCHASE_PRICE_ENABLE` | Включить закупочные цены | `true` |
| `APP_DEMO_SUBDOMAIN` | Поддомен для демо-магазина | `demo` |
| `APP_VENDORS_EXPORT_SECRET_PREFIX` | Префикс для export URL | `sa123asd` |
| `APP_VENDORS_GOOGLE_TABLE_KEY` | Ключ Google таблицы вендоров | — |
| `APP_WISHLIST_COOKIE_KEY` | Ключ cookie для wishlist | `kiosk_wishlist_key` |

### Test Environment

| Переменная | Описание | По умолчанию |
|------------|----------|--------------|
| `TEST_ENV_NUMBER` | Номер параллельного теста | — |
| `AGENT_NAME` | Имя CI агента | — |
| `TRUSTED_IP` | Доверенный IP для BetterErrors | — |

### Release Tooling

| Переменная | Описание | По умолчанию |
|------------|----------|--------------|
| `CHANGELOG_AGENT` | CLI для генерации changelog (`codex`, `claude` или `kimi-claude`) | `codex` |
| `PRODUCTION_URL` | URL production для проверки версии | `https://kiiiosk.store/` |
| `TELEGRAM_RELEASE_BOT_TOKEN` | Токен Telegram бота для отправки release notes | — |
| `TELEGRAM_RELEASE_CHAT_ID` | Chat/Channel ID для отправки release notes | — |

**Получение TELEGRAM_RELEASE_CHAT_ID:**
1. Добавьте [@userinfobot](https://t.me/userinfobot) в ваш чат/канал
2. Отправьте любое сообщение
3. Бот ответит с Chat ID

### Docker Environment

Для Docker deployment используются альтернативные имена переменных:

| Переменная | Описание | По умолчанию |
|------------|----------|--------------|
| `DB_HOST` | Хост PostgreSQL | `pg01` |
| `DB_PORT` | Порт PostgreSQL | `5432` |
| `DB_NAME` | Имя базы данных | `merchantly_production` |
| `DB_USER` | Пользователь БД | `merchantly` |
| `DB_PASSWORD` | Пароль БД | `password` |
| `DB_ADAPTER` | Адаптер БД | `postgresql` |
| `REDIS_HOST` | Хост Redis | `redis01` |
| `REDIS_PORT` | Порт Redis | `6379` |

### Memory Dumps S3

Опциональные настройки S3 для сохранения дампов памяти. Если не указаны, используются основные AWS настройки.

| Переменная | Описание | По умолчанию |
|------------|----------|--------------|
| `APP_MEMORY_DUMP_AWS_BUCKET` | S3 bucket для дампов памяти | — (fallback на `APP_AWS_BUCKET`) |
| `APP_MEMORY_DUMP_AWS_ACCESS_KEY_ID` | Access Key для дампов | — (fallback на `APP_AWS_ACCESS_KEY_ID`) |
| `APP_MEMORY_DUMP_AWS_SECRET_ACCESS_KEY` | Secret Key для дампов | — (fallback на `APP_AWS_SECRET_ACCESS_KEY`) |
| `APP_MEMORY_DUMP_AWS_REGION` | Регион AWS для дампов | — (fallback на `APP_AWS_REGION`) |
| `APP_MEMORY_DUMP_AWS_ENDPOINT` | Custom S3 endpoint для дампов | — (fallback на `APP_AWS_ENDPOINT`) |
| `APP_MEMORY_DUMP_AWS_PROXY_URL` | Proxy URL для дампов | — (fallback на `APP_AWS_PROXY_URL`) |

### HTTP Caching

| Переменная | Описание | По умолчанию |
|------------|----------|--------------|
| `APP_STALE_WHILE_REVALIDATE` | Время stale-while-revalidate (секунды) | `60` |
| `APP_STALE_IF_ERROR` | Время stale-if-error (секунды) | `600` |

### Subdomains & Routing

| Переменная | Описание | По умолчанию |
|------------|----------|--------------|
| `APP_ADMIN_SUBDOMAIN` | Поддомен админки | `admin` |
| `APP_APP_SUBDOMAIN` | Поддомен приложения | `app` |
| `APP_API_SUBDOMAIN` | Поддомен API | `api` |
| `APP_MERCHANT_API_SUBDOMAIN` | Поддомен Merchant API | `mapi` |
| `APP_API_SUBDOMAINS` | Список поддоменов API (через запятую) | `api,apitest,apiwb` |
| `APP_IGNORED_SUBDOMAINS` | Игнорируемые поддомены (через запятую) | — |
| `APP_DOMAIN_ZONES` | Доменные зоны (через запятую) | — |
| `APP_AVAILABLE_DOMAIN_ZONES` | Доступные доменные зоны (через запятую) | — |

### Locales

| Переменная | Описание | По умолчанию |
|------------|----------|--------------|
| `APP_DEFAULT_LOCALE` | Локаль по умолчанию | `ru` |
| `APP_DEFAULT_ENABLED_LOCALES` | Включённые локали (через запятую) | `ru` |
| `APP_AVAILABLE_LOCALES` | Доступные локали (через запятую) | `ru,en` |

### Marketplace Integrations

Флаги интеграций с маркетплейсами и каталогами.

| Переменная | Описание | По умолчанию |
|------------|----------|--------------|
| `APP_YANDEX_MARKET_ENABLE` | Включить выгрузку в Яндекс.Маркет | `true` |
| `APP_YANDEX_TURBO_PAGES_ENABLE` | Включить Яндекс Турбо-страницы | `true` |
| `APP_TORG_MAIL_ENABLE` | Включить выгрузку в Torg.Mail.ru | `true` |
| `APP_FACEBOOK_CATALOG_ENABLE` | Включить Facebook Catalog | `true` |

### Products & Barcodes

| Переменная | Описание | По умолчанию |
|------------|----------|--------------|
| `APP_BARCODE_MIN_LENGTH` | Минимальная длина штрих-кода | `1` |
| `APP_BARCODE_MAX_LENGTH` | Максимальная длина штрих-кода | `16` |
| `APP_EAN_MAX_LENGTH` | Максимальная длина EAN | `13` |
| `APP_TID_MIN_LENGTH` | Минимальная длина TID | `8` |
| `APP_TID_MAX_LENGTH` | Максимальная длина TID | `16` |
| `APP_BARCODE_REQUIRED` | Требовать штрих-код | `true` |
| `APP_VAT_REQUIRED` | Требовать НДС | `true` |
| `APP_WELCOME_CATEGORY_REQUIRED` | Требовать категорию приветствия | `true` |
| `APP_PURCHASE_PRICE_ENABLE` | Включить закупочные цены | `true` |

### Order & Vendor Limits

| Переменная | Описание | По умолчанию |
|------------|----------|--------------|
| `APP_MAXIMAL_MONEY` | Максимальная сумма денег | `100000000` |
| `APP_MAXIMAL_PER_PAGE` | Максимум элементов на страницу | `100` |
| `APP_MAX_RECHARGE_SUMM` | Максимальная сумма пополнения | `10000` |
| `APP_FREE_SMS_COUNT` | Количество бесплатных SMS | `5` |
| `APP_MAXIMAL_ORDERS` | Максимальное количество заказов | `999999999` |
| `APP_MAX_ORDER_SIMILAR_PRODUCTS_COUNT` | Макс. похожих товаров в заказе | `100` |
| `APP_MAX_ORDER_ITEMS_COUNT` | Макс. позиций в заказе | `100` |
| `APP_ARCHIVE_DAYS` | Дней до архивации | `14` |
| `APP_ARCHIVE_NOTIFY_DAYS` | Дней до уведомления об архивации | `3` |

### Vendor Email

| Переменная | Описание | По умолчанию |
|------------|----------|--------------|
| `APP_VENDOR_MAIL_POSTFIX` | Постфикс email вендора | `shop@kiiiosk.store` |
| `APP_VENDOR_INTERNAL_FROM_EMAIL` | Internal from email вендора | — |
| `APP_SUPPORT_EMAIL` | Email поддержки | `support@kiiiosk.store` |
| `APP_NOREPLY_EMAIL` | No-reply email | `noreply@kiiiosk.store` |
| `APP_NEW_VENDOR_EMAIL` | Email для новых вендоров | `support@kiiiosk.store` |

### Analytics & Tracking

| Переменная | Описание | По умолчанию |
|------------|----------|--------------|
| `APP_GTM_ID` | Google Tag Manager ID | — |
| `APP_FB_ADMINS` | Facebook Admins IDs | — |
| `APP_SAVE_ANALYTICS` | Сохранять аналитику | `false` |

### Assets & Fallbacks

| Переменная | Описание | По умолчанию |
|------------|----------|--------------|
| `APP_FALLBACK_IMAGE` | Fallback изображение | `images/fallback/image-none.png` |
| `APP_FALLBACK_LOGO` | Fallback логотип | `images/fallback/logo.svg` |
| `APP_FAVICON` | Путь к favicon | `/favicons/apple-touch-icon-120x120.png` |
| `APP_BILLING_HOST` | Хост биллинга | `https://billing.kiiiosk.store` |
| `APP_CARRIERWAVE_STRING_LENGTH` | Макс. длина строки CarrierWave | `4096` |

### Vendor Features

| Переменная | Описание | По умолчанию |
|------------|----------|--------------|
| `APP_VENDOR_FREE_PUBLISH` | Бесплатная публикация вендора | `false` |
| `APP_VENDOR_LANDING_LINK` | Ссылка на лендинг вендора | — |
| `APP_HAS_PARTNERS` | Включить партнёрскую программу | `true` |
| `APP_SHOW_PDF_INVOICE` | Показывать PDF счёт | `true` |

### SMS Provider

| Переменная | Описание | По умолчанию |
|------------|----------|--------------|
| `APP_SMS_PROVIDER` | Провайдер SMS по умолчанию | `qtelecom` |

### Puppeteer (Screenshots)

| Переменная | Описание | По умолчанию |
|------------|----------|--------------|
| `APP_PUPPETEER_CLI` | Путь к Puppeteer CLI | `./node_modules/puppeteer-webshot-cli/puppeteer-webshot-cli` |

### Dashboard UI

| Переменная | Описание | По умолчанию |
|------------|----------|--------------|
| `APP_DASHBOARD_ITEM_ICON_TYPE` | Тип иконок в дашборде | `fa-icon` |

### Contacts

| Переменная | Описание | По умолчанию |
|------------|----------|--------------|
| `APP_SUPPORT_PHONE` | Телефон поддержки | — |
| `APP_TELEGRAM_LINK` | Ссылка на Telegram поддержки | `https://t.me/kiiiosk_support_ru` |
| `APP_GOOGLE_SPREADSHEET_LINK` | Ссылка на Google Spreadsheet | — |

### Default Country

| Переменная | Описание | По умолчанию |
|------------|----------|--------------|
| `APP_DEFAULT_COUNTRY_CODE` | Код страны по умолчанию (для телефонов) | `7` |
| `APP_DEFAULT_DELIVERY_COUNTRY_CODE` | Код страны доставки по умолчанию | `RU` |
| `APP_DEFAULT_VAT_COUNTRY_CODE` | Код страны НДС по умолчанию | `RU` |

### W1 (WalletOne) Extended

| Переменная | Описание | По умолчанию |
|------------|----------|--------------|
| `APP_W1_DISABLED` | Отключить W1 | `false` |
| `APP_REGISTER_W1_MERCHANT` | Регистрировать W1 мерчантов | — |

---

## ApplicationConfig (anyway_config)

Все настройки из `config/application.yml` доступны через `ApplicationConfig`:

```ruby
# Основные настройки
ApplicationConfig.default_url_options     # => { host: "app.kiiiosk.store", protocol: "https" }
ApplicationConfig.support_email           # => "support@kiiiosk.store"
ApplicationConfig.default_locale          # => "ru"

# Redis
ApplicationConfig.redis[:url]             # => "redis://127.0.0.1:6379/0"

# Telegram
ApplicationConfig.telegram.token          # => токен бота

# Features (через wrapper для backward compatibility)
ApplicationConfig.features.billing        # => true
ApplicationConfig.feature_billing         # => true (прямой доступ)

# Brand (с I18n)
ApplicationConfig.brand.t(:name)          # => "КИОСК" (для текущей локали)
ApplicationConfig.brand[:ru][:name]       # => "КИОСК"

# Billing data
ApplicationConfig.billing_data.sms        # => настройки SMS биллинга
ApplicationConfig.currencies_data.default_currency  # => "RUB"

# Другие wrappers
ApplicationConfig.themes                  # => ThemesAccessWrapper
ApplicationConfig.order_limits            # => OrderLimitsWrapper
ApplicationConfig.integrations            # => IntegrationsWrapper
```

### Секции конфигурации

| Секция | Описание |
|--------|----------|
| `brand` | Локализованные названия бренда |
| `redis` | Настройки Redis |
| `telegram` | Telegram боты |
| `billing` | Настройки биллинга |
| `action_mailer` | Email настройки |
| `themes_hash` | Доступные темы |
| `templates_hash` | Шаблоны страниц |
| `widget_paths` | Пути виджетов |
| `integration_modules` | Модули интеграций |

---

## Local Development

Для локальной разработки создайте файл `config/application.local.yml` (gitignored):

```yaml
# config/application.local.yml
redis:
  url: redis://localhost:6379/0

telegram:
  token: "YOUR_BOT_TOKEN"
  username: "your_bot"
```

Этот файл автоматически загружается anyway_config и переопределяет значения из основного YAML.

---

## Миграции конфигурации

### Telegram Bot Consolidation (v3.88+)

**Что изменилось:** Удалён неиспользуемый клиентский бот `@kiiiosk_bot`. Теперь используется только один бот платформы (бывший operator bot `@kiiiosk_operator_bot`).

**Миграция ENV переменных:**

| Было | Стало |
|------|-------|
| `APP_TELEGRAM__CLIENT__TOKEN` | ❌ Удалить |
| `APP_TELEGRAM__CLIENT__USERNAME` | ❌ Удалить |
| `APP_TELEGRAM__OPERATOR__TOKEN` | `APP_TELEGRAM__TOKEN` |
| `APP_TELEGRAM__OPERATOR__USERNAME` | `APP_TELEGRAM__USERNAME` |

**Инструкция для SRE:**

1. В Kubernetes secrets / ConfigMap:
   ```bash
   # Переименовать переменные
   APP_TELEGRAM__TOKEN=<значение из APP_TELEGRAM__OPERATOR__TOKEN>
   APP_TELEGRAM__USERNAME=kiiiosk_operator_bot

   # Удалить старые переменные
   # - APP_TELEGRAM__CLIENT__TOKEN
   # - APP_TELEGRAM__CLIENT__USERNAME
   # - APP_TELEGRAM__OPERATOR__TOKEN
   # - APP_TELEGRAM__OPERATOR__USERNAME
   ```

2. Проверить после деплоя:
   ```bash
   # В rails console
   Telegram.bot.token  # должен вернуть токен
   Telegram.bot.username  # должен вернуть "kiiiosk_operator_bot"
   ```

**Важно:** Бот `@kiiiosk_bot` можно удалить в BotFather — он не используется.

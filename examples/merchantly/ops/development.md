---
title: Development Environment
doc_kind: engineering
doc_function: canonical
purpose: Docker/dip setup, dev server, browser testing, database operations. Читать при локальной разработке, запуске тестов или настройке окружения.
derived_from:
  - ../dna/governance.md
status: active
audience: humans_and_agents
---

# Development Environment

## Docker Development (dip)

По умолчанию все операции разработки выполняются через **dip** (Docker Interaction Process).

**КРИТИЧНО:** Все команды dip запускай через `direnv exec .`:

```bash
# Первоначальная настройка
direnv exec . dip provision

# Запуск Rails сервера
direnv exec . dip server      # с пробросом портов (service-ports)

# Rails команды
direnv exec . dip rails <command>  # например: dip rails g migration ...
direnv exec . dip console          # Rails console

# Тесты
direnv exec . dip spec        # spec/ (fixtures suite)
direnv exec . dip test        # все тесты

# База данных
direnv exec . dip rake db:migrate
direnv exec . dip rake db:create
direnv exec . dip psql        # PostgreSQL консоль

# Другое
direnv exec . dip bundle install
direnv exec . dip rubocop
direnv exec . dip logs        # просмотр логов
direnv exec . dip down        # остановить все контейнеры
```

**RuboCop** можно запускать напрямую без dip:
```bash
bundle exec rubocop --parallel  # быстрее, чем через dip
make rubocop                    # или через make
```

**Порты:** Внутри Docker сервер слушает на порту 3000, на хосте маппится на порт из переменной `PORT` (по умолчанию 3004).

Конфигурация dip — `dip.yml`, Docker Compose — `docker-compose.yml`.

## Build & Test Commands

- `bin/setup` — gems, JS packages, database
- `bin/dev` — Foreman с `Procfile.dev` (Rails, workers, asset watchers) на порту 3000
- `bundle exec rails s` — только app server
- `make test` — full pre-push pipeline (openbill reset + `make spec` + RuboCop)
- `make spec` / `make rubocop` / `make recreate-db`
- `docker build . -t merchantly`

## Browser Testing (dev)

Адрес определяй из переменных окружения через `direnv exec .`:
1. Сначала `DEV_HOST`
2. Если пустой — `http://kiiiosk.local:${PORT}`
3. Если `PORT` не задан — **остановись и спроси пользователя**

**Важно:** Не ищи порт вручную (`docker ps`, скан `localhost`, `curl`) без явного запроса пользователя на диагностику.

## Database Schema

Проект использует `structure.sql` (не `schema.rb`):
- Схема хранится в `db/structure.sql`
- Пересоздание тестовой БД: `RAILS_ENV=test bundle exec rails db:drop db:create db:migrate`
- **НЕ** используй `db:schema:load` — не работает с structure.sql
- После миграций structure.sql обновляется автоматически

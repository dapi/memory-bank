---
title: Development Environment
doc_kind: ops
doc_function: canonical
purpose: Шаблон документа для локальной разработки. Читать при адаптации setup, dev-команд и browser/database workflow под проект.
derived_from:
  - ../dna/governance.md
status: active
audience: humans_and_agents
---

# Development Environment

После копирования шаблона замени placeholders ниже на реальные команды проекта.

## Setup

Перечисли минимальную подготовку среды.

Если в корне проекта есть `./init.sh`, сначала запусти его. В generic template
он подготавливает только обнаруживаемые `mise`, Git submodules и `direnv`
prerequisites. После installation адаптируй скрипт под реальный стек проекта:
dependency install, database и обязательные сервисы. Не копируй секреты или
`.env`-файлы из другого checkout автоматически.

```bash
# Примеры:
./init.sh
make setup
./bin/setup
npm install
docker compose up -d
direnv allow
asdf install
uv sync
bundle install
pnpm install
```

## Daily Commands

Зафиксируй canonical локальные команды, которые должен знать агент.

```bash
# Примеры:
make dev
make test
make lint
docker compose up app db
pnpm dev
pytest
bundle exec rspec
go test ./...
```

## Browser Testing

Если проект имеет UI, опиши:

- как определить локальный URL;
- где брать порт или host;
- можно ли искать их автоматически;
- какие способы browser verification считаются canonical.

Пример:

1. Сначала читать `DEV_HOST` или `.env`.
2. Если переменная не задана, использовать documented default.
3. Не сканировать порты вручную без явного запроса пользователя.

## Database And Services

Документируй только то, что действительно важно для локальной работы:

- миграции;
- пересоздание локальной БД;
- обязательные сервисы;
- seeded data;
- known pitfalls для разработчиков и агентов.

## Adoption Checklist

- [ ] указаны реальные setup-команды
- [ ] указаны реальные test/lint commands
- [ ] документирован способ определения локального URL
- [ ] перечислены локальные зависимости и сервисы
- [ ] удалены нерелевантные примеры

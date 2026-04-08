---
title: Configuration Guide
doc_kind: engineering
doc_function: canonical
purpose: Шаблон документа ownership-модели конфигурации. Читать при описании env contract, naming conventions и config sources проекта.
derived_from:
  - ../dna/governance.md
status: active
audience: humans_and_agents
---

# Configuration Guide

Этот документ не обязан перечислять все переменные окружения подряд. Его задача: объяснить, где живет canonical schema конфигурации и как downstream-проект документирует важные настройки.

## Configuration Architecture

Опиши реальную модель конфигурации проекта.

Примеры:

- typed config class;
- `.env` + runtime env vars;
- YAML/JSON/TOML файлы с environment overlays;
- secret manager;
- Helm values / Terraform variables / deployment manifests.

### File Layout

```text
config/
├── application.yml
├── environments/
├── secrets/
└── ...
```

### Ownership Rules

Зафиксируй:

1. какой файл или модуль владеет schema конфигурации;
2. где задаются defaults;
3. где лежат environment-specific overrides;
4. как документируются секреты без раскрытия значений.

```ruby
# Пример API доступа к конфигурации:
Config.database_url
Settings.feature_flags.checkout_v2
ENV.fetch("APP_PORT")
```

## Naming Convention For Env Vars

| YAML structure | Env variable |
| --- | --- |
| `database.url` | `APP_DATABASE__URL` |
| `feature_checkout_v2` | `APP_FEATURE_CHECKOUT_V2` |
| `smtp.password` | `APP_SMTP__PASSWORD` |
| `storage.bucket` | `APP_STORAGE__BUCKET` |

Rules:

- выбери один canonical префикс или явно задокументируй, что префикса нет;
- если используется вложенность, зафиксируй separator;
- перечисли правила для списков, boolean и secrets;
- если проект запрещает interpolation внутри config-файлов, напиши это явно.

## Documenting Important Variables

Если проекту нужен справочник ключевых переменных, не перечисляй все подряд. Сфокусируйся на значимых runtime contracts.

| Variable | Description | Default | Owner |
| --- | --- | --- | --- |
| `APP_DATABASE__URL` | Основное подключение к БД | none | platform |
| `APP_REDIS__URL` | Кэш или очередь | `redis://localhost:6379/0` | platform |
| `APP_PUBLIC_BASE_URL` | Базовый URL для генерации ссылок | `http://localhost:3000` | product/platform |
| `APP_FEATURE_X_ENABLED` | Feature flag | `false` | owning team |

## Secrets

- Никогда не вставляй реальные значения секретов в репозиторий.
- Документируй только способ их хранения, выдачи и rotation policy.
- Если часть конфигурации приходит из secret manager, это должно быть написано явно.

## Adoption Checklist

- [ ] описан schema-owner конфигурации
- [ ] задокументирована naming convention
- [ ] перечислены ключевые runtime/env contracts
- [ ] описан secret handling
- [ ] удалены ссылки на несуществующие downstream-справочники

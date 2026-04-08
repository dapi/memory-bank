---
title: Production & Stage Access
doc_kind: engineering
doc_function: canonical
purpose: kubectl, psql, rails console, логи, credentials, тестовые магазины. Читать при работе с production/stage окружениями.
derived_from:
  - ../../../memory-bank/dna/governance.md
status: active
audience: humans_and_agents
---

# Production & Stage Access

Вся инфраструктура в K8s. Используй make-команды из этого репозитория.

## Основные команды

```bash
# Rails Console
make rails-console-production
make rails-console-stage

# PostgreSQL (psql)
make psql-production
make psql-stage

# Клонирование production → stage
make clone-production-db-to-stage

# Rails Runner (выполнение Ruby-кода)
# Production:
cd ~/code/brandymint/infra && direnv exec . kubectl --context=fury -n production exec deploy/merchantly -- \
  /bin/bash -c 'cd /rails && ./bin/rails runner "puts User.count"'

# Stage:
cd ~/code/brandymint/infra && direnv exec . kubectl --context=fury -n stage exec deploy/merchantly -- \
  /bin/bash -c 'cd /rails && ./bin/rails runner "puts User.count"'
```

**Важно:**
- Используй `~/code/brandymint/infra` напрямую (переменная `$INFRA_DIR` может быть не установлена)
- Все команды из infra через `direnv exec .`
- Production: `-n production`, stage: `-n stage`

**Детали подключения:**
- Production: `kiiiosk_production` (user: `kiiiosk`)
- Stage: `kiiiosk_stage` (user: `kiiiosk_staging`)
- Host: `${MERCHANTLY_PG_PUBLIC_HOST}:5432`

## Авторизация

Для входа в операторскую панель (`*.kiiiosk.store/operator`) или админку:

```bash
env | grep KIIIOSK_ADMIN
# KIIIOSK_ADMIN_EMAIL    — логин
# KIIIOSK_ADMIN_PASSWORD — пароль
```

**URL:** `https://app.kiiiosk.store/login`

### Администраторская панель (Administrate)

| Стенд | URL |
|-|-|
| Production | `https://admin.kiiiosk.store` |
| Stage | `https://admin.stage.kiiiosk.store` |

**Важно:** `KIIIOSK_ADMIN_*` — данные для **production**. На stage пароль может отличаться.

### Настройка окружения

1. Установи `INFRA_DIR=~/code/brandymint/infra` в shell profile
2. В директории infra: `direnv allow`
3. Секреты: `${INFRA_DIR}/.env.secrets` (git-crypt encrypted)

## Проверка версии приложения

```bash
# Production (рекомендуемый)
curl -s https://kiiiosk.store/ | grep -oP 'merchantly_version.*?content="\K[^"]+'

# Stage
curl -s https://stage.kiiiosk.store/ | grep -oP 'merchantly_version.*?content="\K[^"]+'

# Через kubectl
cd ~/code/brandymint/infra && direnv exec . kubectl --context=fury -n stage get deployment merchantly \
  -o jsonpath='{.spec.template.spec.containers[0].image}'
```

Альтернативы в браузере: `gon.app_version` (JS console) или meta-тег `merchantly_version`.

## Тестовые магазины

| Стенд | URL | Описание |
|-|-|-|
| Stage | https://demo.stage.kiiiosk.store | Тестовый магазин (классическая тема) |
| Production | https://varvara-shop.com | Реальный магазин для smoke-тестов |

**Активация demo на stage:**
```bash
cd ~/code/brandymint/infra && direnv exec . kubectl --context=fury -n stage exec deploy/merchantly -- \
  /bin/bash -c 'cd /rails && ./bin/rails runner "Vendor.find_by(subdomain: :demo).update_column(:is_published, true)"'
```

## Логи

Доступны через Loki/Grafana или kubectl. Инструкция: `~/code/brandymint/infra/docs/sre-logs-merchantly.md`

```bash
# kubectl
cd ~/code/brandymint/infra && direnv exec . kubectl --context=fury -n production logs -l app=merchantly --tail=1000

# Loki API (нужен port-forward)
cd ~/code/brandymint/infra && direnv exec . kubectl --context=fury -n logging port-forward svc/loki-gateway 3100:80 &
curl -s "http://localhost:3100/loki/api/v1/query_range" \
  --data-urlencode 'query={namespace="production",app="merchantly"} |~ "error"' \
  --data-urlencode 'limit=50' | jq -r '.data.result[].values[][1]'
```

**Grafana:**
- HTTP логи: https://grafana.brandymint.ru/d/merchantly-http
- ActiveJob: https://grafana.brandymint.ru/d/merchantly-activejob
- Explore: https://grafana.brandymint.ru/explore

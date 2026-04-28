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

## IaaC Репозиторий Инфраструктуры

**Важно:** Вся инфраструктура управляется через отдельный IaaC-репозиторий, расположенный вне этого проекта.

```
INFRA_DIR=~/code/brandymint/infra
```

**Настройка окружения:**
1. Установи `INFRA_DIR=~/code/brandymint/infra` в shell profile
2. В директории infra: `direnv allow`
3. Секреты: `${INFRA_DIR}/.env.secrets` (git-crypt encrypted)

**Все команды kubectl выполняются из INFRA_DIR через `direnv exec .`**

---

## Основные команды

### Rails Console

```bash
# Production
cd $INFRA_DIR && direnv exec . kubectl --context=fury -n production exec deploy/merchantly -- \
  /bin/bash -c 'cd /rails && ./bin/rails console'

# Stage
cd $INFRA_DIR && direnv exec . kubectl --context=fury -n stage exec deploy/merchantly -- \
  /bin/bash -c 'cd /rails && ./bin/rails console'
```

### PostgreSQL (psql)

```bash
# Production
cd $INFRA_DIR && direnv exec . kubectl --context=fury -n production exec deploy/merchantly -- \
  /bin/bash -c 'cd /rails && ./bin/rails dbconsole'

# Stage
cd $INFRA_DIR && direnv exec . kubectl --context=fury -n stage exec deploy/merchantly -- \
  /bin/bash -c 'cd /rails && ./bin/rails dbconsole'
```

**Детали подключения:**
- Production: `kiiiosk_production` (user: `kiiiosk`)
- Stage: `kiiiosk_stage` (user: `kiiiosk_staging`)
- Host: `${MERCHANTLY_PG_PUBLIC_HOST}:5432`

### Клонирование production → stage

```bash
cd $INFRA_DIR && make clone-production-db-to-stage
```

### Rails Runner (выполнение Ruby-кода)

```bash
# Production
cd $INFRA_DIR && direnv exec . kubectl --context=fury -n production exec deploy/merchantly -- \
  /bin/bash -c 'cd /rails && ./bin/rails runner "puts User.count"'

# Stage
cd $INFRA_DIR && direnv exec . kubectl --context=fury -n stage exec deploy/merchantly -- \
  /bin/bash -c 'cd /rails && ./bin/rails runner "puts User.count"'
```

**Важно:**
- Используй `$INFRA_DIR` (или `~/code/brandymint/infra`) напрямую
- Все команды из infra через `direnv exec .`
- Production: `-n production`, stage: `-n stage`
- Context: `--context=fury`

---

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

---

## Деплой

### Стандартный деплой через CI/CD

Деплой выполняется автоматически при создании GitHub release. См. [release.md](./release.md).

### Ручной деплой (экстренные случаи)

```bash
# Установить конкретную версию образа
cd $INFRA_DIR && direnv exec . kubectl --context=fury -n production set image deployment/merchantly merchantly=cr.selcloud.ru/brandymint/merchantly:v3.54.0

# Следить за rolling update
cd $INFRA_DIR && direnv exec . kubectl --context=fury -n production rollout status deployment/merchantly

# Проверить health check
curl https://kiiiosk.store/up
```

### Откат (Rollback)

```bash
# Откат к предыдущей ревизии
cd $INFRA_DIR && direnv exec . kubectl --context=fury -n production rollout undo deployment/merchantly

# Откат к конкретной ревизии
cd $INFRA_DIR && direnv exec . kubectl --context=fury -n production rollout undo deployment/merchantly --to-revision=42

# История деплоев
cd $INFRA_DIR && direnv exec . kubectl --context=fury -n production rollout history deployment/merchantly
```

### Перезапуск деплоя

```bash
# Полный рестарт (пересоздание подов)
cd $INFRA_DIR && direnv exec . kubectl --context=fury -n production rollout restart deployment/merchantly

# Рестарт SolidQueue workers
cd $INFRA_DIR && direnv exec . kubectl --context=fury -n production rollout restart deployment/merchantly --selector=component=queue
```

---

## Диагностика

### Проверка статуса подов

```bash
# Список подов
cd $INFRA_DIR && direnv exec . kubectl --context=fury -n production get pods

# Детальная информация о поде
cd $INFRA_DIR && direnv exec . kubectl --context=fury -n production describe pod <pod-name>

# Метрики ресурсов
cd $INFRA_DIR && direnv exec . kubectl --context=fury -n production top pods
```

### Логи

Полная документация: [logs.md](./logs.md) — kubectl, Loki API, Grafana.

**Быстрые команды:**
```bash
# Production логи (контейнер: ror)
cd $INFRA_DIR && direnv exec . kubectl --context=fury -n production logs -l app=merchantly -c ror --tail=1000

# Live stream
cd $INFRA_DIR && direnv exec . kubectl --context=fury -n production logs <pod-name> -c ror --tail=500 -f
```

### Проверка работы SolidQueue

```bash
# Проверить статус очереди
cd $INFRA_DIR && direnv exec . kubectl --context=fury -n production exec deploy/merchantly -- \
  /bin/bash -c 'cd /rails && ./bin/rails runner "puts SolidQueue::Job.count"'

# Застрявшие джобы
cd $INFRA_DIR && direnv exec . kubectl --context=fury -n production exec deploy/merchantly -- \
  /bin/bash -c 'cd /rails && ./bin/rails runner "puts SolidQueue::Job.where(finished_at: nil).where(\"created_at < ?\", 1.hour.ago).count"'

# Рестарт workers
cd $INFRA_DIR && direnv exec . kubectl --context=fury -n production rollout restart deployment/merchantly
```

### Проверка базы данных

```bash
# Количество коннектов
cd $INFRA_DIR && direnv exec . kubectl --context=fury -n production exec deploy/merchantly -- \
  /bin/bash -c 'cd /rails && ./bin/rails runner "puts ActiveRecord::Base.connection.execute(\"SELECT count(*) FROM pg_stat_activity;\").first"'

# Долгие запросы
cd $INFRA_DIR && direnv exec . kubectl --context=fury -n production exec deploy/merchantly -- \
  /bin/bash -c 'cd /rails && ./bin/rails runner "puts ActiveRecord::Base.connection.execute(\"SELECT * FROM pg_stat_activity WHERE state = active AND now() - query_start > interval 5 minutes;\").map(&:to_s)"'
```

---

## Проверка версии приложения

```bash
# Production (рекомендуемый)
curl -s https://kiiiosk.store/ | grep -oP 'merchantly_version.*?content="\K[^"]+'

# Stage
curl -s https://stage.kiiiosk.store/ | grep -oP 'merchantly_version.*?content="\K[^"]+'

# Через kubectl
cd $INFRA_DIR && direnv exec . kubectl --context=fury -n production get deployment merchantly \
  -o jsonpath='{.spec.template.spec.containers[0].image}'
```

Альтернативы в браузере: `gon.app_version` (JS console) или meta-тег `merchantly_version`.

---

## Тестовые магазины

| Стенд | URL | Описание |
|-|-|-|
| Stage | https://demo.stage.kiiiosk.store | Тестовый магазин (классическая тема) |
| Production | https://varvara-shop.com | Реальный магазин для smoke-тестов |

**Активация demo на stage:**
```bash
cd $INFRA_DIR && direnv exec . kubectl --context=fury -n stage exec deploy/merchantly -- \
  /bin/bash -c 'cd /rails && ./bin/rails runner "Vendor.find_by(subdomain: :demo).update_column(:is_published, true)"'
```

---

## Экстренные процедуры

### Масштабирование до 0 (экстренный случай)

```bash
# Остановить приложение
cd $INFRA_DIR && direnv exec . kubectl --context=fury -n production scale deployment/merchantly --replicas=0

# Запустить обратно
cd $INFRA_DIR && direnv exec . kubectl --context=fury -n production scale deployment/merchantly --replicas=2
```

### Доступ к серверам БД

```bash
# Подключиться к PostgreSQL напрямую (через port-forward)
cd $INFRA_DIR && direnv exec . kubectl --context=fury -n database port-forward svc/postgres 5432:5432 &
psql -h localhost -U kiiiosk -d kiiiosk_production
```

### Проверка сертификатов (Custom Domains)

```bash
# Проверить сертификаты в K8s
cd $INFRA_DIR && direnv exec . kubectl --context=fury -n production get certificates

# Проверить ingress
cd $INFRA_DIR && direnv exec . kubectl --context=fury -n production get ingress
```

---

## Связанные документы

- [release.md](./release.md) — релизный процесс
- [runbooks/kubernetes-deployment.md](./runbooks/kubernetes-deployment.md) — runbook по деплою
- [runbooks/k8s-custom-domains.md](./runbooks/k8s-custom-domains.md) — custom domains runbook

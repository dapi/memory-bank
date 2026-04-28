---
title: Operations Index
doc_kind: ops
doc_function: index
purpose: Навигация по операционной документации. Читать когда нужно работать с dev/stage/production окружением, релизами или платежами.
derived_from:
  - ../../../memory-bank/dna/governance.md
status: active
audience: humans_and_agents
---

# Operations Index

## Важно: IaaC Репозиторий Инфраструктуры

Вся инфраструктура управляется через отдельный IaaC-репозиторий:

```
INFRA_DIR=~/code/brandymint/infra
```

Все команды kubectl и деплоя выполняются из `$INFRA_DIR` через `direnv exec .`

---

## Документация по окружениям

| Документ | Назначение |
|----------|------------|
| [Development Environment](development.md) | dip, Docker, dev-сервер, browser testing, database schema. Читать при локальной разработке. |
| [Production & Stage Access](production.md) | kubectl, psql, rails console, credentials, деплой, диагностика. Читать при работе с production/stage. |
| [Viewing Logs](logs.md) | Логи merchantly — kubectl, Loki API, Grafana. Читать при диагностике проблем. |
| [Release & Deployment](release.md) | релизный процесс, Docker tagging, feature branches, тестовый план. Читать при выполнении релиза. |
| [CloudPayments](payments.md) | админка, credentials, Playwright-вход, коды отклонения. Читать при работе с платежами. |
| [Configuration](config.md) | полный справочник конфигурации (anyway_config, application.yml). Читать при работе с настройками. |
| [Metrics and monitoring](metrics.md) | метрики, алерты и мониторинг инфраструктуры. |

## Kubernetes

| Документ | Назначение |
|----------|------------|
| [K8s Overview](k8s/overview.md) | обзор Kubernetes-инфраструктуры. |
| [Web deployments](k8s/web.md) | веб-деплойменты и сервисы. |
| [Background jobs](k8s/jobs.md) | фоновые задачи и workers. |
| [PostgreSQL](k8s/postgres.md) | база данных PostgreSQL. |
| [Redis cache and queues](k8s/redis.md) | кэш и очереди Redis. |
| [Memcached caching](k8s/memcached.md) | кэширование Memcached. |

## Runbooks

| Runbook | Назначение |
|---------|------------|
| [Kubernetes Deployment](runbooks/kubernetes-deployment.md) | Деплой, откат, диагностика проблем в K8s |
| [K8s Custom Domains](runbooks/k8s-custom-domains.md) | Проблемы с custom domains, сертификатами |
| [K8s Custom Domains Deploy](runbooks/k8s-custom-domains-deploy.md) | Деплой custom domains |
| [SMTP Encryption](runbooks/smtp-encryption.md) | SMTP и шифрование |

---

## Инструменты

- [`scripts/check_memory_bank_index.py`](../../scripts/check_memory_bank_index.py) — проверяет, что все страницы `memory-bank/` достижимы из `AGENTS.md` через цепочку индексов (orphan-файлы, аннотированные ссылки, frontmatter).

## Быстрые ссылки

### Деплой
```bash
# Релиз (см. release.md)
STAGE=production make minor-release

# Ручной деплой (см. production.md)
cd $INFRA_DIR && direnv exec . kubectl --context=fury -n production set image deployment/merchantly merchantly=cr.selcloud.ru/brandymint/merchantly:vX.Y.Z

# Откат (см. runbooks/kubernetes-deployment.md)
cd $INFRA_DIR && direnv exec . kubectl --context=fury -n production rollout undo deployment/merchantly
```

### Диагностика
```bash
# Логи (см. production.md)
cd $INFRA_DIR && direnv exec . kubectl --context=fury -n production logs -l app=merchantly --tail=1000

# Rails Console (см. production.md)
cd $INFRA_DIR && direnv exec . kubectl --context=fury -n production exec deploy/merchantly -- /bin/bash -c 'cd /rails && ./bin/rails console'

# Проверка подов
cd $INFRA_DIR && direnv exec . kubectl --context=fury -n production get pods
```

---
title: Stages And Non-Local Environments
doc_kind: engineering
doc_function: canonical
purpose: Шаблон документа по доступу к production-like окружениям. Читать при адаптации прав доступа, smoke-checks, логов и runtime-операций под проект.
derived_from:
  - ../dna/governance.md
status: active
audience: humans_and_agents
---

# Stages And Non-Local Environments

Опиши здесь не только production, но и stage, beta, preview, sandbox или другие non-local окружения, если они существуют.

## Environment Inventory

| Environment | Purpose | Access path | Notes |
| --- | --- | --- | --- |
| `production` | Реальные пользователи и live traffic | Команда, jump host или UI | Самые строгие ограничения |
| `staging` | Предрелизная проверка | Команда, URL или namespace | Может использоваться для smoke |
| `sandbox` | Проверка интеграций и unsafe экспериментов | Optional | Если есть |

## Common Operations

Здесь должны быть только реально разрешенные операции и их canonical entrypoints.

```bash
# Примеры:
make console ENV=staging
make logs ENV=production
kubectl -n staging logs deploy/app
ssh <bastion>
psql "$DATABASE_URL"
```

Для каждой операции зафиксируй:

- кто имеет право ее запускать;
- какие approval gates нужны;
- где проходит граница read-only vs mutating access.

## Credentials And Access

Опиши:

- где хранятся секреты;
- как выдаются права;
- какие env vars или secret stores используются;
- что считается недопустимым обходом процедуры доступа.

Никогда не храни реальные production credentials в шаблоне.

## Version And Health Checks

Задокументируй безопасные способы проверить:

- текущую deployed version;
- health endpoint;
- smoke URL;
- базовые operational dashboards.

Пример:

```bash
curl -fsS https://<stage-host>/health
kubectl -n <namespace> get deploy <app>
```

## Logs And Observability

Опиши canonical пути к:

- application logs;
- metrics;
- traces;
- error tracker;
- dashboards для основных сервисов.

## Test Data And Smoke Targets

Если проект использует staging/demo tenants, seed users или test accounts, перечисли их здесь вместе с правилами использования.

## Adoption Checklist

- [ ] перечислены все non-local environments
- [ ] указаны canonical access paths
- [ ] описаны safe health/version checks
- [ ] перечислены observability entrypoints
- [ ] удалены фальшивые или нерелевантные примеры

---
title: Release And Deployment
doc_kind: engineering
doc_function: canonical
purpose: Шаблон документа релизного процесса. Читать при адаптации versioning, changelog, deployment и release verification под проект.
derived_from:
  - ../dna/governance.md
status: active
audience: humans_and_agents
---

# Release And Deployment

## Release Flow

Опиши реальный порядок шагов для проекта.

Пример:

1. bump версии;
2. обновление changelog;
3. tag или release branch;
4. build артефактов;
5. deploy на staging;
6. smoke/acceptance;
7. production deploy.

## Release Commands

Зафиксируй canonical команды проекта и явные safety rules.

```bash
# Примеры:
make release ENV=staging
make deploy ENV=production
gh release create vX.Y.Z
docker build -t registry/app:vX.Y.Z .
```

Укажи явно:

- какие переменные окружения обязательны;
- какие окружения требуют явного approval;
- где проходит граница между automated и manual release steps.

## Release Test Plan

При каждом релизе полезно создавать отдельный тестовый план.

**Формат:** `release-v{VERSION}-test-plan.md`

**Минимальная структура:**

```markdown
# Тестовый план релиза v{VERSION}

**Дата:** YYYY-MM-DD
**Предыдущая версия:** v{PREV_VERSION}
**Текущая версия:** v{VERSION}
**Стенд:** <environment>

## Обзор изменений

| Issue | Название | Тип | Приоритет |
| --- | --- | --- | --- |
| #XXXX | Описание задачи | Feature/Fix/Refactoring/Tech debt | Высокий/Средний/Низкий |

## Проверка изменений

- [ ] Описан хотя бы один test case для каждого крупного change set

## Smoke-тесты

- [ ] Главная страница открывается
- [ ] Основной пользовательский поток работает
- [ ] Админский или внутренний путь работает
- [ ] Health endpoint отвечает успешно
```

## Rollback

Для реального проекта обязательно зафиксируй:

- что считается rollback unit;
- какой путь fastest safe rollback;
- кто подтверждает rollback в production;
- какие данные или миграции необратимы.

---
title: Release & Deployment
doc_kind: engineering
doc_function: canonical
purpose: Релизный процесс, Docker tagging, feature branches, тестовый план. Читать при выполнении релиза или деплоя.
derived_from:
  - ../../../memory-bank/dna/governance.md
status: active
audience: humans_and_agents
---

# Release & Deployment

Релиз: bump версии → генерация changelog (Claude CLI) → коммит → GitHub release → Docker build/push → deploy.

## Команды релиза

**КРИТИЧЕСКИ ВАЖНО:** Всегда явно указывай `STAGE`. Переменная окружения `STAGE` может быть установлена неверно — **НИКОГДА не полагайся на значение по умолчанию**.

```bash
# Stage
STAGE=stage make patch-release   # 3.54.0 → 3.54.1
STAGE=stage make minor-release   # 3.54.0 → 3.55.0
STAGE=stage make major-release   # 3.54.0 → 4.0.0

# Production — только с явным подтверждением пользователя
STAGE=production make minor-release
```

**Перед релизом:**
1. Уточни у пользователя целевой стенд (stage или production)
2. Всегда используй `STAGE=<стенд>`

CHANGELOG.md обновляется автоматически через `scripts/update_changelog.sh`.

## Тестовый план релиза

При каждом релизе **обязательно** создавай тестовый план в `.protocols/`.

**Формат:** `release-v{VERSION}-test-plan.md`

**Структура:**

```markdown
# Тестовый план релиза v{VERSION}

**Дата:** YYYY-MM-DD
**Предыдущая версия:** v{PREV_VERSION}
**Текущая версия:** v{VERSION}
**Стенд:** stage

---

## Обзор изменений

| Issue | Название | Тип | Приоритет |
|-|-|-|-|
| #XXXX | Описание задачи | Feature/Fix/Refactoring/Tech debt | Высокий/Средний/Низкий |

---

## 1. Название фичи/фикса (#ISSUE)

### 1.1 Функциональные тесты

- [ ] **1.1.1** Описание тест-кейса

### 1.2 Проверка на stage

\`\`\`bash
# Команды для проверки
\`\`\`

---

## Smoke-тесты (обязательные перед production)

- [ ] **S.1** Главная страница открывается
- [ ] **S.2** Каталог товаров работает
- [ ] **S.3** Корзина работает
- [ ] **S.4** Оформление заказа работает
- [ ] **S.5** Операторская панель открывается
- [ ] **S.6** Health check: `curl https://stage.kiiiosk.store/health` → 200

---

## Результаты тестирования

| Секция | Статус | Тестировщик | Дата |
|-|-|-|-|
| 1. Фича X | ⏳ | | |
| Smoke-тесты | ⏳ | | |

**Легенда:** ✅ ❌ ⏳
```

**Как собрать:**
1. `git log --oneline $(git describe --tags --abbrev=0)..HEAD`
2. `git log ... | grep -oE '#[0-9]+' | sort -u`
3. Найти тестовый план в теле каждого issue на GitHub
4. Объединить + добавить smoke-тесты

## Feature Branches (Docker tagging)

- **main/master**: `cr.selcloud.ru/brandymint/merchantly:latest` и `:VERSION`
- **feature-ветки**: `:VERSION-feature-name`

Pre-release git tag: `1.2.3-feature-name`.

Guard-tag-exists:
- **main/master**: `vVERSION`
- **feature**: `vVERSION-feature-name`

```bash
git checkout -b feature/my-awesome-feature
# ... development ...
make patch-release  # → v1.2.3-feature-my-awesome-feature

# Ручной деплой
make build-and-push deploy
```

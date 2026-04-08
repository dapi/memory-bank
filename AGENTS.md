# Правила репозитория

## Структура проекта и организация модулей

Корневой `README.md` объясняет устройство репозитория и различие между шаблоном и примером.

- `memory-bank/` — переносимый шаблон, который должен оставаться generic.
- `examples/merchantly/` — concrete example, где сохранена Merchantly-специфика.
- `memory-bank/dna/` — governance-ядро шаблона.
- `memory-bank/flows/` — reusable lifecycle docs и governed templates.
- `memory-bank/prd/` — instantiated Product Requirements Documents.
- `memory-bank/use-cases/` — instantiated канонические сценарии проекта.
- `memory-bank/domain/`, `memory-bank/engineering/`, `memory-bank/ops/` — project-adaptation layers шаблона.
- `memory-bank/adr/` и `memory-bank/features/` — пустые или минимальные точки назначения для instantiated документов.

Новые generic-правила размещайте в `memory-bank/`. Конкретные проектные примеры и specialization не возвращайте обратно в шаблон, складывайте их в `examples/`.

## Команды разработки и проверки

У репозитория нет собственного build/runtime-приложения. Перед PR достаточно лёгких проверок:

- `rg --files memory-bank examples` для проверки структуры и имён файлов;
- `git diff --check` для поиска лишних пробелов и conflict markers;
- `sed -n '1,120p' path/to/doc.md` для быстрой проверки frontmatter и заголовков;
- `rg -n "Merchantly|CloudPayments|kiiiosk|Bugsnag" memory-bank` чтобы убедиться, что project-specific детали не протекли обратно в шаблон.

## Стиль оформления и соглашения по именованию

Пишите в Markdown: короткие секции, понятные заголовки, относительные ссылки. Governed-документы в `memory-bank/` должны начинаться с YAML frontmatter; поле `status` обязательно всегда, а `derived_from`, `delivery_status` и `decision_status` добавляются, когда этого требует тип документа. См. `memory-bank/dna/frontmatter.md`.

Для обычных документов используйте lowercase kebab-case, например `testing-policy.md`. Для структурированных артефактов сохраняйте шаблонные naming rules, например `features/FT-XXX/` и `ADR-XXX-short-decision-name.md`.

## Правила проверки

Автоматизированного тестового набора у репозитория нет. Проверяйте изменения вручную:

- убедитесь, что индексы и ссылки соответствуют новой структуре;
- держите шаблон и пример согласованными на уровне смыслов, но не дублируйте Merchantly-специфику обратно в `memory-bank/`;
- при изменении template docs проверяйте соседние governed-файлы на противоречия.

## Коммиты и pull request

Следуйте конвенции из `memory-bank/engineering/git-workflow.md`: короткие commit messages в настоящем времени, например `docs: tighten template ops guidance`.

В pull request опишите:

- что изменено в шаблоне;
- что изменено в examples;
- какие ссылки или naming rules были затронуты.

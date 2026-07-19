# Правила репозитория

## Структура проекта и организация модулей

Корневой `README.md` объясняет устройство репозитория и правила использования шаблона.

- `memory-bank/` — переносимый шаблон, который должен оставаться generic.
- `memory-bank/dna/` — governance-ядро шаблона.
- `memory-bank/flows/` — reusable lifecycle docs и governed templates.
- `memory-bank/prd/` — instantiated Product Requirements Documents.
- `memory-bank/use-cases/` — instantiated канонические сценарии проекта.
- `memory-bank/domain/`, `memory-bank/engineering/`, `memory-bank/ops/` — project-adaptation layers шаблона.
- `memory-bank/adr/` и `memory-bank/features/` — пустые или минимальные точки назначения для instantiated документов.

Новые generic-правила размещайте в `memory-bank/`. Конкретную project-specific specialization не возвращайте обратно в шаблон.

## Команды разработки и проверки

У репозитория нет runtime-приложения, но есть Go CLI для проверки документации. Перед PR запускайте:

- `rg --files memory-bank` для проверки структуры и имён файлов;
- `go run ./cmd/memory-bank-lint` для аудита ссылок, reachability и expected README-индексов внутри `memory-bank/`;
- `go test ./...` для проверки валидатора и его JSON-контракта;
- `git diff --check` для поиска лишних пробелов и conflict markers;
- `sed -n '1,120p' path/to/doc.md` для быстрой проверки frontmatter и заголовков;
- `rg -n "PROJECT_SPECIFIC_TERM" memory-bank` с реальными терминами downstream-проекта, чтобы убедиться, что project-specific детали не протекли обратно в шаблон.

## Стиль оформления и соглашения по именованию

Пишите в Markdown: короткие секции, понятные заголовки, относительные ссылки. Governed-документы в `memory-bank/` должны начинаться с YAML frontmatter; поле `status` обязательно всегда, а `derived_from`, `delivery_status` и `decision_status` добавляются, когда этого требует тип документа. См. `memory-bank/dna/frontmatter.md`.

Для обычных документов используйте lowercase kebab-case, например `testing-policy.md`. Для структурированных артефактов сохраняйте шаблонные naming rules, например `features/FT-XXX/` и `ADR-XXX-short-decision-name.md`.

## Правила проверки

Документационный шаблон проверяется вручную и через `memory-bank-lint`; Go-код покрыт тестами. При изменениях:

- убедитесь, что индексы и ссылки соответствуют новой структуре;
- не дублируйте project-specific детали обратно в `memory-bank/`;
- при изменении template docs проверяйте соседние governed-файлы на противоречия.

## Коммиты и pull request

Следуйте конвенции из `memory-bank/engineering/git-workflow.md`: короткие commit messages в настоящем времени, например `docs: tighten template ops guidance`.

В pull request опишите:

- что изменено в шаблоне;
- какие ссылки или naming rules были затронуты.

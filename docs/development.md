# Разработка репозитория

Этот документ предназначен для разработки самого шаблона и CLI. Инструкция по установке готового бинарника находится в [`memory-bank.md`](memory-bank.md).

## Структура CLI

CLI разделён на тонкие entrypoint'ы и переиспользуемые внутренние пакеты:

```text
cmd/memory-bank/           основной entrypoint
cmd/memory-bank-lint/      compatibility entrypoint
internal/cli/              subcommands, flags и общий output/error contract
internal/lint/             audit-движок, отчёт и testdata
internal/repository/       общий repo root discovery
```

Оба бинарника вызывают `internal/cli`; lint-семантика и JSON contract принадлежат `internal/lint`. Новые команды добавляются в общий dispatcher `memory-bank`, а не отдельными бинарниками.

## Локальная разработка

Требуется Go версии `1.21` или новее. `rg` используется для быстрой проверки структуры шаблона.

Перед PR запускайте:

```bash
rg --files memory-bank
gofmt -w $(rg --files cmd internal -g '*.go')
go test -count=1 -race ./...
go vet ./...
go run ./cmd/memory-bank lint
git diff --check
```

Если проверка запускается не из Git-репозитория или нужно проверить другой checkout:

```bash
go run ./cmd/memory-bank lint --repo-root /path/to/repository
```

## Изменение JSON-контракта

Golden report для CLI лежит в `internal/lint/testdata/expected-report.json`.

Меняйте его только если contract отчёта изменился намеренно. После такого изменения проверьте:

```bash
go test ./...
```

## Документационный шаблон

При изменении `memory-bank/`:

- убедитесь, что индексы и ссылки соответствуют новой структуре;
- не переносите project-specific детали обратно в generic-шаблон;
- при изменении governed template docs проверяйте соседние governed-файлы на противоречия.

# Разработка репозитория

Этот документ предназначен для разработки самого шаблона и CLI. Инструкция по установке готового бинарника находится в [`memory-bank-lint.md`](memory-bank-lint.md).

## Структура CLI

Код `memory-bank-lint` лежит в одной директории:

```text
cmd/memory-bank-lint/
  main.go
  audit.go
  markdown.go
  report.go
  types.go
  *_test.go
  testdata/
```

`main.go` отвечает за CLI-флаги, exit codes и вывод. Остальные файлы содержат аудит markdown-навигации, типы отчёта и форматирование результата.

## Локальная разработка

Требуется Go версии `1.21` или новее. `rg` используется для быстрой проверки структуры шаблона.

Перед PR запускайте:

```bash
rg --files memory-bank
gofmt -w cmd/memory-bank-lint/*.go
go test -count=1 -race ./...
go vet ./...
go run ./cmd/memory-bank-lint
git diff --check
```

Если проверка запускается не из Git-репозитория или нужно проверить другой checkout:

```bash
go run ./cmd/memory-bank-lint --repo-root /path/to/repository
```

## Изменение JSON-контракта

Golden report для CLI лежит в `cmd/memory-bank-lint/testdata/expected-report.json`.

Меняйте его только если contract отчёта изменился намеренно. После такого изменения проверьте:

```bash
go test ./...
```

## Документационный шаблон

При изменении `memory-bank/`:

- убедитесь, что индексы и ссылки соответствуют новой структуре;
- не переносите project-specific детали обратно в generic-шаблон;
- при изменении governed template docs проверяйте соседние governed-файлы на противоречия.

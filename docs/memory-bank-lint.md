# memory-bank-lint: compatibility entrypoint

Основной интерфейс CLI — [`memory-bank lint`](memory-bank.md). Команда `memory-bank-lint` временно сохраняется для совместимости: она использует тот же audit-движок и сохраняет прежние flags, exit codes, текстовый результат и JSON contract.

Новые установки и CI должны использовать `memory-bank`. Инструкции ниже оставлены для существующей автоматизации в течение переходного периода.

`memory-bank-lint` проверяет навигационную целостность `memory-bank/`:

- broken relative markdown links внутри audit scope;
- orphan-документы, на которые никто не ссылается внутри scope;
- достижимость каждого документа от entrypoint'ов по индексной навигации;
- документы, которые достижимы только глубже порога навигации;
- contract ожидаемых `README.md`-индексов.

## Установка

Для регулярных проверок установите CLI один раз. Требуется Go версии `1.21` или новее:

```bash
go install github.com/dapi/memory-bank/tools/cmd/memory-bank-lint@latest
```

`go install` помещает бинарник в `GOBIN` или `GOPATH/bin`; этот каталог должен находиться в `PATH`. Повторите ту же команду, чтобы обновить CLI до актуальной версии.

Для локального использования подходит `@latest`. В блокирующем CI фиксируйте конкретный release tag или commit: иначе новая версия CLI может изменить результат проверки без изменений в downstream-репозитории.

На macOS или Linux проверьте, что установленная команда доступна:

```bash
command -v memory-bank-lint
```

Если репозиторий уже клонирован и нужно установить версию из текущего checkout:

```bash
(cd tools && go install ./cmd/memory-bank-lint)
```

## Запуск

После установки запускайте CLI из любого места внутри Git-репозитория:

```bash
memory-bank-lint
```

По умолчанию `memory-bank-lint` ищет ближайший родительский `.git` и использует найденный каталог как repo root. Если запуск идёт вне Git-репозитория или нужно проверить другой checkout, передайте корень явно:

```bash
memory-bank-lint --repo-root /path/to/repository
```

Что означает результат:

- exit code `0` — errors не найдены; warnings по глубине возможны, но аудит считается пройденным;
- non-zero exit code — найдены проблемы, которые нужно исправить до PR;
- `--json` — структурированный отчёт, пригодный для последующей автоматической доиндексации другим агентом или инструментом.

## Параметры запуска

- `--max-depth N` — порог глубины индексной навигации в прыжках; по умолчанию `3`; документы глубже порога попадают в warning, а не в error;
- `--entrypoint PATH` — явный entrypoint для аудита; параметр repeatable; принимает repo-relative или scope-relative пути; неоднозначные пути без префикса сначала резолвятся внутри `--scope-root`, а для явного repo-root пути используйте `./PATH` или `/PATH`; если передан, используется вместо дефолтного `memory-bank/README.md`;
- `--scope-root DIR` — меняет audit scope; по умолчанию `memory-bank`;
- `--repo-root DIR` — явно задаёт корень репозитория;
- `--json` — печатает только JSON-отчёт;
- `--version` — печатает версию бинарника и завершает работу;
- `--help` — печатает справку по запуску и параметрам.

## Примеры

```bash
memory-bank-lint --max-depth 4
```

```bash
memory-bank-lint \
  --entrypoint README.md \
  --entrypoint AGENTS.md \
  --max-depth 4
```

Для разовой проверки без установки CLI:

```bash
go run github.com/dapi/memory-bank/tools/cmd/memory-bank-lint@latest
```

Установку из GitHub Releases или через Homebrew используйте только после фактической публикации соответствующих release assets и Cask; если их ещё нет, используйте `go install`.

## Когда запускать

- после добавления, удаления или переименования `.md`-файлов в `memory-bank/`;
- после правок `README.md`-индексов и относительных ссылок;
- перед открытием PR с изменениями в template navigation или document structure.

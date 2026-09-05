# Использование memory-bank-cli

`memory-bank-cli` — внешний CLI для установки, обновления, проверки и диагностики Memory Bank. Реализация, releases, installation guide и полный command contract принадлежат отдельному репозиторию [`dapi/memory-bank-cli`](https://github.com/dapi/memory-bank-cli).

Ownership-контракт template payload описан в [`ownership.md`](ownership.md). Основные команды интеграции:

- `memory-bank-cli init` создаёт служебный `memory-bank/.lock` и устанавливает отсутствующие файлы;
- `memory-bank-cli update` строит ownership-aware mutation plan и применяет его атомарно;
- `memory-bank-cli lint` проверяет ссылки и индексную навигацию;
- `memory-bank-cli doctor` выполняет read-only диагностику adoption, governance, managed drift, CI и навигации.

## Установка

Используйте закреплённый release со страницы [GitHub Releases](https://github.com/dapi/memory-bank-cli/releases). Выберите asset для своей ОС и архитектуры, проверьте его по `checksums.txt` из того же release и добавьте `memory-bank-cli` в `PATH`.

Для локальной интерактивной работы допустим выбранный командой актуальный release. В блокирующем CI всегда фиксируйте конкретные release tag, asset и checksum, чтобы новая версия CLI не меняла результат проверки без изменений в репозитории. Готовый downstream workflow приведён в [инструкции по внедрению](adoption.md#подключить-ci).

После установки проверьте identity бинарника:

```bash
memory-bank-cli --version
```

## Запуск

В downstream-репозитории из любого места внутри Git-репозитория:

```bash
memory-bank-cli lint
memory-bank-cli doctor
```

Для проверки другого checkout передайте его явно:

```bash
memory-bank-cli lint --repo-root /path/to/repository
memory-bank-cli doctor --repo-root /path/to/repository
```

`lint` выполняет быстрый аудит навигации. `doctor` включает navigation audit и дополнительно проверяет template identity, governance, managed agent instructions, managed drift и CI gate. Обе команды поддерживают machine-readable `--json`; blocking findings возвращают non-zero exit code.

## Template source profile

Исходный template repository проверяется явным profile. В source checkout
canonical payload находится в `template/`; CLI удаляет этот префикс для
каждого tracked regular file, поэтому `template/memory-bank/...` становится
`memory-bank/...`, а `template/init.sh` — `init.sh`. Автоматическое
определение предназначено для downstream-репозитория по `memory-bank/.lock`.
Для проверки source template используйте:

```bash
memory-bank-cli lint --scope-root template/memory-bank --entrypoint template/memory-bank/README.md
memory-bank-cli doctor --profile template
```

Downstream repository определяется по `memory-bank/.lock`; при обычном внедрении достаточно `memory-bank-cli doctor` с profile `auto` по умолчанию.

## Init и update

`init` и `update` принимают локальный clean checkout источника, закреплённый immutable commit:

```bash
memory-bank-cli init \
  --source /path/to/memory-bank-checkout \
  --template-version v1.2.3 \
  --source-ref FULL_COMMIT_SHA
```

Перед обновлением сначала проверьте план:

```bash
memory-bank-cli update \
  --source /path/to/new-memory-bank-checkout \
  --template-version v1.3.0 \
  --source-ref FULL_COMMIT_SHA \
  --dry-run
```

Подробные flags, exit codes и JSON schemas документируются в [`dapi/memory-bank-cli`](https://github.com/dapi/memory-bank-cli). Правила владения, conflict policy и atomic update semantics для этого шаблона остаются в [`ownership.md`](ownership.md).

## Dual-role репозиторий: синхронизация собственного инстанса

Этот репозиторий одновременно является upstream-шаблоном и downstream-потребителем: `template/memory-bank/` — payload, `memory-bank/` — установленный инстанс. Корневые ассеты в нём не копии, а симлинки в `template/`, чтобы рабочие инструменты репозитория всегда совпадали с payload:

```text
WORKFLOW.md            -> template/WORKFLOW.md
bootstrap-symphony.sh  -> template/bootstrap-symphony.sh
run-symphony.sh        -> template/run-symphony.sh
.codex/agents          -> ../template/.codex/agents/
.start-issue/prompt.md -> ../template/.start-issue/prompt.md
```

По контракту из [`ownership.md`](ownership.md) симлинк в любом компоненте пути считается unsafe path, поэтому `memory-bank-cli pull --repo-root .` здесь останавливается и ничего не читает. Это не дефект CLI и не повод материализовать симлинки: корневой `.gitignore` шире шаблонного, а `.envrc` и `init.sh` в корне шаблона не нужны.

Поэтому синхронизация выполняется через `tools/sync-project-memory-bank.rb`. Он делает pull в изолированной песочнице, где корневых ассетов нет, и переносит обратно только поддерево `memory-bank/**` вместе с lock; корневые записи из lock удаляются, потому что CLI этими путями не управляет.

```bash
ruby tools/sync-project-memory-bank.rb            # план
ruby tools/sync-project-memory-bank.rb --apply    # применить
```

Запускайте синхронизацию после мержа изменений шаблона в `main`: lock закрепляет commit источника, и SHA из невлитой ветки исчезнет из истории после squash-мержа. Конфликты внутри `memory-bank/` скрипт не разрешает — он останавливается и передаёт решение человеку по правилам [`ownership.md`](ownership.md).

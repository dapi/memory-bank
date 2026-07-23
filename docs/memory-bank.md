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

Из любого места внутри Git-репозитория:

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

Исходный template repository содержит корневой marker `.memory-bank-template` со значением `memory-bank-template-v1`. Благодаря ему `memory-bank-cli doctor --profile auto` распознаёт source repository без CLI source code или локального Go module.

Marker является только source-repository metadata. Он находится рядом с `memory-bank/`, не внутри него, и не копируется в downstream payload. Для явной проверки source template используйте:

```bash
memory-bank-cli lint
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

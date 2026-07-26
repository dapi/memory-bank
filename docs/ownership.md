# Ownership и безопасные обновления

`memory-bank/.lock` — служебный контракт между downstream-проектом и версией шаблона. Файл создаётся командой `memory-bank-cli init` внутри установленного `memory-bank/` и коммитится вместе с ним; из upstream template он не копируется. Формальная схема: [`schema/memory-bank-lock-v1.schema.json`](schema/memory-bank-lock-v1.schema.json).

Upstream payload хранится в source checkout как `template/`. Ownership paths
и digests в lock относятся к downstream paths после удаления этого префикса:
`template/memory-bank/...` становится `memory-bank/...`, а
`template/init.sh` становится `init.sh`. Source-каталог не становится частью
lock schema.

## Классы владения

| Класс | Текущая граница шаблона | Поведение update |
| --- | --- | --- |
| `managed` | Каждый tracked regular file из canonical `template/`, включая `init.sh` и `memory-bank/` | Проверяет текущий payload по digest. Чистый файл обновляется или удаляется; локальный drift становится conflict. |
| `user-owned` | Неизвестные downstream paths | Никогда автоматически не перезаписывается и не удаляется. Неизвестный существующий файл получает этот класс по fail-safe правилу. |
| `generated` | `memory-bank/.generated/` зарезервирован для будущих детерминированных генераторов; в текущем template таких файлов нет | Может быть пересоздан или удалён только детерминированным producer. |

`base_digest` и `base_mode` (`100644` или `100755`) описывают файл в зафиксированной template-base. `payload_digest` и `payload_mode` присутствуют только там, где текущий файл является проверяемым managed/generated contract. Изменение executable bit managed-файла проверяется так же, как изменение его содержимого.

## Init и update

Команды работают с локальным checkout источника, закреплённым на immutable commit. CLI не делает network fetch и не доверяет moving branch автоматически:

```bash
memory-bank-cli init \
  --source /path/to/memory-bank-checkout \
  --template-version v1.2.3 \
  --source-ref FULL_COMMIT_SHA
```

`--source` должен указывать на корень чистого Git checkout, `--source-ref` — в точности совпадать с его `HEAD`. Незакоммиченные, untracked или ignored payloads внутри upstream `template/` отклоняются, как и source, совпадающий с downstream repo либо вложенный в него через обычный путь или symlink. Payload и executable modes читаются непосредственно из объектов закреплённого commit и устанавливаются в downstream после удаления префикса `template/`, поэтому обычные Git text conversions (`core.autocrlf`, `.gitattributes`) не создают ложный drift и не меняют устанавливаемые байты.

`init` подходит и для пустого проекта, и для ранее скопированного payload: совпадающие existing files принимаются как managed, а несовпадающие сохраняются как downstream-owned для явного разрешения при следующем update.

Перед обновлением сначала проверьте полный plan:

```bash
memory-bank-cli update \
  --source /path/to/new-memory-bank-checkout \
  --template-version v1.3.0 \
  --source-ref FULL_COMMIT_SHA \
  --dry-run
```

Добавьте `--json` для machine-readable report format `1`. Каждому известному пути назначается одно решение: `create`, `update`, `preserve`, `conflict` или `delete`. Conflict даёт exit code `1`, сохраняет исходные файлы и lock и требует ручного разрешения. Чтобы принять incoming template, замените конфликтующий файл его incoming payload и повторите update: совпадение digest будет принято как новая base и записано в lock. Для сохранения другого варианта скорректируйте ownership осознанной миграцией lock; команда никогда не выбирает победителя молча.

Без `--dry-run` сначала строится и проверяется весь plan. При наличии хотя бы одного conflict ничего не применяется. Все новые payload, включая [managed-блок agent instructions](agent-instructions.md), заранее записываются в repo-local temporary staging, существующие файлы перемещаются туда перед заменой, а lock заменяется последним. Planned digests повторно проверяются перед мутацией и перед commit lock. Clean-managed изменения topology `file → directory` и `directory → file` применяются в той же транзакции: прежние payload сначала сохраняются в staging, затем создаётся новая форма пути. Ошибка во время применения откатывает уже сделанные изменения и структуру каталогов без повторной записи содержимого; если rollback или очистка не могут завершиться, команда возвращает явную ошибку и сохраняет staging с recovery-копиями. Успешный no-op не переписывает lock, поэтому повторный update идемпотентен.

Корень downstream-репозитория закрепляется по filesystem identity на весь run. Destination-мутации выполняются относительно уже открытых directory handles: через `openat`-семейство на Unix и через handle-relative NT APIs на Windows. Поэтому замена ранее проверенного parent на symlink или junction не может перенаправить операцию наружу. Symlink или reparse point в любом компоненте ниже корня, включая сам managed-файл или lock, считается unsafe path: команда завершается ошибкой и не читает и не изменяет target ссылки.

## Версионирование

`schema_version` версионирует lock contract независимо от версии template. CLI читает schema `1`; неизвестная версия завершается ошибкой без мутаций. Unversioned prototype со значением `0` имеет семантику v1 и атомарно переписывается в schema `1` при следующем успешном update.

`template.version` — понятная человеку версия, `template.source_ref` — immutable идентификатор фактического source checkout. `last_update` меняется только вместе с успешной сменой template state или миграцией schema.

# Компонентное внедрение и миграция

DNA задаёт общее владение и целостность документации. Documents добавляет типы
и базовые шаблоны. Flows добавляет AI-процессы и версионированные требования к
явно подключённым документам. Нормативные форматы принадлежат
[CTR-01](component-wire-format.md), команды — [memory-bank-cli](memory-bank.md).

## Требование к CLI

Компонентный payload вводится после source-format bridge и компонентного CLI.
Пока соответствующий release не опубликован, используйте проверенный candidate
компонентной ветки CLI. Проверка версии сама по себе не заменяет handshake:

```bash
memory-bank-cli capabilities --require components/v1 --require adoption/v1
```

Компонентные записи поддерживаются на Linux/macOS. Более старый CLI останавливается
в `tools/install-components.sh` до вызова installer. Bridge принимает известный
legacy source `f1f04de843aef45a2425d4a7351d577bbf89e940` и отклоняет компонентный
payload. Прямой запуск pre-bridge CLI на компонентном payload не поддерживается.

## Новая установка

Из чистого checkout шаблона на выбранном неизменяемом коммите:

```bash
~/code/memory-bank/tools/install-components.sh init \
  --repo-root /path/to/project --preset docs
```

`core` устанавливает DNA; `docs` — DNA и Documents; `full` добавляет Flows.
Без `--preset` новая установка выбирает `legacy`: три компонента и прежние
адаптеры. `--adapter codex`, `--adapter start-issue`, `--adapter symphony` и
`--adapter bootstrap` добавляют интеграции вместе с их зависимостями.

README и managed-блок AGENTS формируются по составу установки. Текст снаружи
маркеров остаётся собственностью проекта. Lock хранит выбранный состав;
`full/legacy` имеют проверяемый registry даже при отсутствии подключённых документов.

## Базовые документы и добавление Flows

```bash
memory-bank-cli document create --repo-root /path/to/project \
  --type feature --path memory-bank/features/FT-123/brief.md
~/code/memory-bank/tools/install-components.sh pull \
  --repo-root /path/to/project --preset full
```

Pull сохраняет заполненный brief и не подключает его к процессу. Без параметров
выбора он сохраняет прежний состав. Удаление компонентов или адаптеров не
поддерживается. Базовые ADR, PRD, use case, research brief и epic charter также
не требуют Flows. Типы перечислены в `memory-bank/document-types/README.md`.

Для adoption сначала заполните требования выбранного расширения:

```bash
memory-bank-cli document adopt --repo-root /path/to/project \
  --path memory-bank/features/FT-123/brief.md --contract feature/v1 --dry-run
memory-bank-cli document adopt --repo-root /path/to/project \
  --path memory-bank/features/FT-123/brief.md --contract feature/v1
```

`flow_contract` не является переключателем: CLI сверяет его с registry, identity
и bundle digest. Удаление маркера, записи или registry при неизменном lock даёт
conflict. Не редактируйте служебное состояние для отключения проверок.

Смена версии выполняется через `document transition --path PATH --contract ID`
с `--evidence REF`, когда контракт требует evidence. CLI проверяет оба контракта.
Перенос выполняется через `document move --id ID --path OLD --to NEW`; исходный
контекст документа сохраняется. Неизвестные операции, detach и delete отклоняются.

## Существующая legacy-установка

Переход от проверок всех документов типа к explicit adoption меняет семантику.
Обычный pull, unattended-режим и `--preset legacy` не являются согласием на него.
Миграция поддерживает только закреплённый legacy source `f1f04de843aef45a2425d4a7351d577bbf89e940`;
остальные установки продолжают использовать свой прежний source.

Сначала получите не изменяющий проект preview:

```bash
~/code/memory-bank/tools/install-components.sh pull \
  --repo-root /path/to/project --migrate-components --dry-run --json
```

Просмотрите proposed changes, исходные файлы и права, состав установки и
`migration_plan_digest`. Если CLI сообщает неоднозначную принадлежность документа
или ownership conflict, подготовьте resolution JSON по
[CTR-01](component-wire-format.md#migration-resolution-and-preview) и повторите
preview с `--migration-resolution /path/to/resolution.json`.

Примените тот же просмотренный план, подставив полученный digest:

```bash
~/code/memory-bank/tools/install-components.sh pull \
  --repo-root /path/to/project --migrate-components \
  --migration-plan-digest sha256:REVIEWED_DIGEST
```

При использовании resolution-файла передайте тот же файл и при применении.
Изменившиеся bytes, permissions, source, lock или resolution делают digest
устаревшим; создайте новый preview. Миграция сохраняет прежние адаптеры и пути
заполненных документов. Неисправные managed assets и неразрешённые конфликты
останавливают запись. Некорректный YAML необходимо исправить заранее.

Существующие flow-документы получают стабильные identities и snapshot selectors
с compatibility contracts. Их прежний pass/fail сохраняется; миграция может
сохранить уже существующие ошибки, но не добавляет и не удаляет findings.
Обычные последующие операции не получают этого исключения.

Новые документы не включаются в snapshot автоматически. Обычный `document create`
остаётся базовым и после миграции. Явный `--legacy-flow` выбирает закреплённый
compatibility contract и создаёт per-document record. Требования контракта всё
равно должны выполняться. Переход одного документа на новый контракт исключает
только его identity из selector; остальные документы сохраняют прежний контракт.

## Проверка и восстановление

```bash
memory-bank-cli lint --repo-root /path/to/project
memory-bank-cli doctor --repo-root /path/to/project
```

Обе команды проверяют применимость контрактов и целостность. Они не включают
неустановленные процессы и не исправляют registry автоматически. Source profile
`doctor --profile template` относится к устройству репозитория, а не к preset.

Запись выполняется одной транзакцией, lock — последним. При неуспешном rollback
CLI сохраняет `.memory-bank-update-*` с `recovery.json` и точной картой backups.
Сохраните параллельные правки отдельно; восстановите все before bytes, permissions
и состояния каталогов по журналу или доверенной резервной копии. Одного lock
недостаточно. CLI разрешит продолжение только после полной проверки восстановления.
Если журнал содержит durable committed outcome, повтор проверяет полное after
состояние перед очисткой staging. Неизвестный или повреждённый журнал блокирует запись.

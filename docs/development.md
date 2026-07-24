# Разработка репозитория

Этот репозиторий содержит upstream payload `template/memory-bank/` и документацию по его внедрению. `memory-bank-cli` устанавливает этот payload в downstream-репозитории под именем `memory-bank/`. Реализация, command contract и релизы CLI принадлежат отдельному репозиторию [`dapi/memory-bank-cli`](https://github.com/dapi/memory-bank-cli).

Исходный template repository проверяется явным профилем `memory-bank-cli doctor --profile template`. Автоматический профиль предназначен для downstream-репозиториев с `memory-bank/.lock`.

## Локальная проверка

Установите закреплённый release `memory-bank-cli` по [инструкции CLI](memory-bank.md). Перед PR запускайте:

```bash
rg --files template/memory-bank
memory-bank-cli lint --scope-root template/memory-bank --entrypoint template/memory-bank/README.md
memory-bank-cli doctor --profile template
git diff --check
```

Для явной проверки другого checkout используйте `--repo-root`:

```bash
memory-bank-cli lint --repo-root /path/to/repository --scope-root template/memory-bank --entrypoint template/memory-bank/README.md
memory-bank-cli doctor --profile template --repo-root /path/to/repository
```

CI устанавливает закреплённый release CLI, проверяет его checksum и выполняет те же `lint` и явный template-profile `doctor` gate. Этот репозиторий не публикует CLI releases.

## Документационный шаблон

При изменении `template/memory-bank/`:

- убедитесь, что индексы и ссылки соответствуют новой структуре;
- не переносите source-repository metadata или project-specific детали обратно в generic-шаблон;
- при изменении governed template docs проверяйте соседние governed-файлы на противоречия.

# Разработка репозитория

Этот репозиторий содержит только исходный generic-шаблон `memory-bank/` и документацию по его внедрению. Реализация, command contract и релизы CLI принадлежат отдельному репозиторию [`dapi/memory-bank-cli`](https://github.com/dapi/memory-bank-cli).

Корневой файл `.memory-bank-template` — source-repository marker для `memory-bank-cli doctor --profile auto`. Он не входит в копируемый `memory-bank/` payload и не должен появляться в downstream-проектах.

## Локальная проверка

Установите закреплённый release `memory-bank-cli` по [инструкции CLI](memory-bank.md). Перед PR запускайте:

```bash
rg --files memory-bank
memory-bank-cli lint
memory-bank-cli doctor --profile template
git diff --check
```

Для явной проверки другого checkout используйте `--repo-root`:

```bash
memory-bank-cli lint --repo-root /path/to/repository
memory-bank-cli doctor --profile template --repo-root /path/to/repository
```

CI устанавливает закреплённый release CLI, проверяет его checksum и выполняет те же `lint` и template-profile `doctor` gates. Этот репозиторий не публикует CLI releases.

## Документационный шаблон

При изменении `memory-bank/`:

- убедитесь, что индексы и ссылки соответствуют новой структуре;
- не переносите source-repository metadata или project-specific детали обратно в generic-шаблон;
- при изменении governed template docs проверяйте соседние governed-файлы на противоречия.

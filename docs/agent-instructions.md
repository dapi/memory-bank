# Managed-блок инструкций агента

`memory-bank init` и `memory-bank update` управляют только коротким routing-блоком в agent instruction file. По умолчанию target — корневой `AGENTS.md`:

```markdown
<!-- MEMORY BANK START -->
<!-- MEMORY BANK MANAGED BLOCK VERSION: 1 -->
Before substantial delivery work, read memory-bank/README.md, memory-bank/dna/README.md, and memory-bank/flows/routing.md.
Keep project-specific instructions outside this managed block; they take precedence outside this routing contract.
<!-- MEMORY BANK END -->
```

Markers — стабильная граница ownership только тогда, когда каждый marker занимает отдельную строку без отступов или другого текста. Строка `MEMORY BANK MANAGED BLOCK VERSION` версионирует payload независимо от markers. Governance остаётся в `memory-bank/`; блок только направляет агента к canonical documents и не становится вторым source of truth.

## Правила обновления

- Если markers отсутствуют, CLI создаёт файл или добавляет блок в конец существующего файла.
- Между существующим текстом и новым блоком добавляется одна пустая строка. Если последняя строка не была завершена, CLI сначала добавляет newline; остальные существующие bytes не меняются.
- Если ровно одна корректная пара markers содержит старый payload, CLI заменяет только bytes от start marker до end marker и следующий за ним `LF`, если он есть. Весь текст до и после сохраняется byte-for-byte.
- Актуальный блок — no-op: повторный `init/update` не создаёт diff блока.
- Несовпадающее количество markers, несколько пар, end marker перед start marker, повреждённый marker или его inline-упоминание — conflict. CLI не изменяет ни agent file, ни template, ни lock.

`init/update --dry-run` печатает planned block diff в text output и в поле `decisions[].diff` JSON report. Применение блока входит в общую атомарную транзакцию template update. Если меняется только блок, template lock и его `last_update` не переписываются.

## Doctor

```bash
memory-bank doctor
memory-bank doctor --json
```

`doctor` проверяет блок без мутаций. Missing и outdated block дают planned `create/update`, ambiguous markers — `conflict`, актуальный блок — `preserve`. Любой drift возвращает exit code `1`; актуальное состояние — `0`. JSON дополнительно содержит `drift_count` и `conflict_count`.

## Альтернативный target

Canonical default — только `AGENTS.md`. Для проекта, который использует другой instruction file, задайте один repo-relative target явно и одинаково во всех командах:

```bash
memory-bank init ... --agent-file CLAUDE.md
memory-bank update ... --agent-file CLAUDE.md
memory-bank doctor --agent-file CLAUDE.md
```

CLI не создаёт блоки одновременно в нескольких файлах и не принимает target внутри `memory-bank/`. Project-specific инструкции должны оставаться вне managed markers; CLI никогда их не переписывает, и они имеют приоритет за пределами минимального routing contract.

# Generated runtime projection в agent instructions

`memory-bank-cli` v1.0.0 управляет коротким блоком routing-инструкций в agent instruction file. По умолчанию target — корневой `AGENTS.md`:

```markdown
<!-- MEMORY BANK START -->
<!-- MEMORY BANK MANAGED BLOCK VERSION: 1 -->
Before substantial delivery work, read memory-bank/README.md, memory-bank/dna/README.md, and memory-bank/flows/routing.md.
Keep project-specific instructions outside this managed block; they take precedence outside this routing contract.
<!-- MEMORY BANK END -->
```

Markers — стабильная граница ownership только тогда, когда каждый marker занимает отдельную строку без отступов или другого текста. Строка `MEMORY BANK MANAGED BLOCK VERSION` версионирует projection независимо от markers.

## Правила обновления

- Если markers отсутствуют, CLI создаёт файл или добавляет блок в конец существующего файла.
- Между существующим текстом и новым блоком добавляется одна пустая строка. Если последняя строка не была завершена, CLI сначала добавляет newline; остальные существующие bytes не меняются.
- Если ровно одна корректная пара markers содержит старый payload, CLI заменяет только bytes от start marker до end marker и следующий за ним `LF`, если он есть. Весь текст до и после сохраняется byte-for-byte.
- Актуальный блок — no-op: повторный `init/update` не создаёт diff блока.
- Несовпадающее количество markers, несколько пар, end marker перед start marker, повреждённый marker или его inline-упоминание — conflict. CLI не изменяет ни agent file, ни template, ни lock.

`init/update --dry-run` печатает planned block diff в text output и в поле `decisions[].diff` JSON report. Применение блока входит в общую атомарную транзакцию template update. Если меняется только блок, template lock и его `last_update` не переписываются.

## Doctor

```bash
memory-bank-cli doctor
memory-bank-cli doctor --json
```

`doctor` проверяет блок без мутаций как часть общего adoption-аудита. Missing, outdated или ambiguous managed block становятся finding `agent.managed_block_drift` уровня `error`; актуальный блок finding не создаёт. Полный versioned JSON contract и остальные диагностические группы описаны в [`memory-bank.md`](memory-bank.md).

## Альтернативный target

Canonical default — только `AGENTS.md`. Для проекта, который использует другой instruction file, задайте один repo-relative target явно и одинаково во всех командах:

```bash
memory-bank-cli init ... --agent-file CLAUDE.md
memory-bank-cli update ... --agent-file CLAUDE.md
memory-bank-cli doctor --agent-file CLAUDE.md
```

CLI не создаёт блоки одновременно в нескольких файлах и не принимает target внутри `memory-bank/`. Project-specific инструкции должны оставаться вне managed markers; CLI никогда их не переписывает, и они имеют приоритет за пределами routing contract и generated runtime projection.

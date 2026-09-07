---
title: "EP-141: Компонентный Memory Bank"
doc_kind: epic
doc_function: index
purpose: "EP-141: Компонентный Memory Bank"
derived_from:
  - ../../flows/epic.md
status: active
audience: humans_and_agents
epic_stage: execution
---

# EP-141: Компонентный Memory Bank

Owner: Danil. Source: [issue 141](https://github.com/dapi/memory-bank/issues/141).
Маршрут Epic: несколько delivery units в template и CLI, общий контракт и migration risk.
Intake пропущен: интент, scope и критерии уже заданы issue; пользователь поручил реализацию.

- [Charter](charter.md) — intent и acceptance.
- [Roadmap](roadmap.md) — последовательность зависимых поставок.
- [Subissues](subissues.md) — delivery slices и владельцы.
- [Risks](risks.md) — общие риски и меры контроля.
- [Decisions](decision-log.md) — решения по исполнению.

W1 bridge завершён в [CLI PR 63](https://github.com/dapi/memory-bank-cli/pull/63), commit
3b434fd93678c36447d10d4f308a39ce5d74b040. Required CI, canonical canary, independent code
и simplify reviews clean; PR ready, без merge/release. Текущая работа — shared Solution Ready
для W2 [CLI #62](https://github.com/dapi/memory-bank-cli/issues/62) и W3–W4
[FT-141](../../features/FT-141/README.md). Template feature остаётся на стадии design;
CLI execution plan reviewed, его исполнение ожидает общий gate.

`epic_stage: execution` означает, что delivery slices переданы своим владельцам.
Это не заменяет локальные feature/CLI gates: W1 имеет отдельный reviewed plan, а W2 и
FT-141 не начинают component implementation до clean shared design и своих execution plans.

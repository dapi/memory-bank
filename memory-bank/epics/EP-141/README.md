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
и simplify reviews clean; PR ready, без merge/release. Shared Solution Ready и оба execution
plans прошли независимую проверку. W2 [CLI #62](https://github.com/dapi/memory-bank-cli/issues/62)
реализует contract library и затем transaction/command integration; W3
[FT-141](../../features/FT-141/README.md) готовит payload и producer/consumer fixtures.
Финальная интеграция и W4 review/PR ещё не завершены.

`epic_stage: execution` означает передачу delivery slices их владельцам. Завершение
каждого slice требует его проверок и evidence; наличие кода или draft payload не заменяет
готовый компонентный CLI и полную матрицу приёмки.

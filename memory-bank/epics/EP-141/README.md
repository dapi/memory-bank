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

W1 bridge готов в [CLI PR 63](https://github.com/dapi/memory-bank-cli/pull/63).
W2 реализован в [CLI PR 64](https://github.com/dapi/memory-bank-cli/pull/64),
W3/W4 — в [template PR 143](https://github.com/dapi/memory-bank/pull/143).
[FT-141](../../features/FT-141/README.md) завершена в границах implementation/review/PR;
[delivery evidence](../../../docs/component-delivery-evidence.md) связывает acceptance,
зелёный CI и clean independent reviews с immutable revisions.

Инициатива остаётся в `epic_stage: execution` до отдельного human closure. Merge,
release и live migration не входят в эту поставку. Порядок дальнейшей публикации:
bridge → supporting CLI → component payload; владельцу переданы связанные PR.

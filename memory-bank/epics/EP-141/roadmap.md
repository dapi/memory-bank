---
title: "EP-141: Roadmap"
doc_kind: epic
doc_function: roadmap
purpose: "EP-141: Roadmap"
derived_from:
  - charter.md
status: active
audience: humans_and_agents
---

# EP-141: Roadmap

| Wave | Outcome | Dependency | Exit gate |
| --- | --- | --- | --- |
| W1 | Bridge source gate | Baseline | Проверенный bridge design, несовместимый source отклоняется до мутаций; PR 63 ready |
| W2 | Component install, document/adoption validation и migration | W1 + shared Solution Ready | CLI contract/transaction tests зелёные |
| W3 | Самодостаточные DNA/Documents, Flows extensions и adapters | W1, W2 | Составы проверены реальным CLI |
| W4 | Cross-repo fixtures, docs, review и PR | W2, W3 | Независимое review сошлось, required CI зелёный |

Bridge и supporting CLI должны быть доступны до использования component payload.
PR не означает публикацию release или обновление downstream. До поставки нового CLI
пользовательский маршрут остаётся на закреплённом legacy source.

Stop: обнаруженное изменение исходного intent или неподдерживаемый переход сначала
фиксируется у canonical owner; частичная мутация установки запрещена.

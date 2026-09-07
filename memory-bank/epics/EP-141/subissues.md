---
title: "EP-141: Delivery slices"
doc_kind: epic
doc_function: subissue_registry
purpose: "EP-141: Delivery slices"
derived_from:
  - charter.md
  - roadmap.md
status: active
audience: humans_and_agents
---

# EP-141: Delivery slices

| ID | Slice | Owner | Wave | State |
| --- | --- | --- | --- | --- |
| EP-SI-01 | Поддержка состава и совместимости источника | dapi/memory-bank-cli | W1–W2 | implemented; [CLI PR 64](https://github.com/dapi/memory-bank-cli/pull/64), [CLI #62](https://github.com/dapi/memory-bank-cli/issues/62), [CLI delivery contract](https://github.com/dapi/memory-bank-cli/blob/caf0f3eaf3af290a702c8553795168584ac8b987/docs/component-delivery.md) |
| EP-SI-02 | Независимые документационные компоненты и интеграция | dapi/memory-bank | W3–W4 | implemented; [PR 143](https://github.com/dapi/memory-bank/pull/143), [issue 141](https://github.com/dapi/memory-bank/issues/141), [FT-141](../../features/FT-141/README.md) |

Scope принят поручением пользователя реализовать issue 141. CLI #62 владеет отдельным delivery contract и проверками в memory-bank-cli; FT-141 импортирует эту boundary. CLI не имеет установленного Memory Bank и не копирует template governance ради tracking.

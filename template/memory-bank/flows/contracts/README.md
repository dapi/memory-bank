---
title: "Flow contracts"
doc_kind: process
doc_function: index
purpose: "Flow contracts"
derived_from:
  - ../../dna/governance.md
  - ../../document-types/README.md
status: active
audience: humans_and_agents
---

# Flow contracts

Flow adoption — явный выбор versioned contract для конкретного документа. Базовый
документ остаётся базовым даже после установки Flows. Создавай через
`memory-bank-cli document create --type TYPE --path PATH --contract ID` или подключай
существующий документ через `document adopt --path PATH --contract ID`.
Переход и перенос выполняются явными `document transition` и `document move`; registry,
metadata и lock должны оставаться согласованными. Не редактируй registry вручную.

Опубликованные bundles неизменяемы. Они содержат frozen DNA/base/extension rules и
[engine artifact](../engines/governance-v1.json). Переход на новые правила требует нового ID
и явной операции. `delivery_status` и feature lifecycle принадлежат feature flow;
`research_status` — research flow. Базовый ADR владеет decision_status самостоятельно.

## Contracts

- [adr/v1](adr/v1.json)
- [epic/v1](epic/v1.json)
- [feature/v1](feature/v1.json)
- [legacy/f1f04de/adr/v1](legacy/f1f04de/adr/v1.json)
- [legacy/f1f04de/epic/v1](legacy/f1f04de/epic/v1.json)
- [legacy/f1f04de/feature/v1](legacy/f1f04de/feature/v1.json)
- [legacy/f1f04de/prd/v1](legacy/f1f04de/prd/v1.json)
- [legacy/f1f04de/research/v1](legacy/f1f04de/research/v1.json)
- [legacy/f1f04de/use_case/v1](legacy/f1f04de/use_case/v1.json)
- [prd/v1](prd/v1.json)
- [research/v1](research/v1.json)
- [use_case/v1](use_case/v1.json)

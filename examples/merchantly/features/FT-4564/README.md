---
title: "FT-4564: Per-page Liquid override и кастомные шаблоны для страниц заказа"
doc_kind: feature
doc_function: index
purpose: "Index-файл для feature package FT-4564. Навигация по документам фичи."
derived_from:
  - ../../dna/governance.md
  - ../../flows/feature-flow.md
status: active
audience: humans_and_agents
---

# FT-4564: Per-page Liquid override и кастомные шаблоны для страниц заказа

[Feature Flow](../../flows/feature-flow.md) — читай, когда работаешь с фичей, чтобы понимать ее lifecycle и требования к документации.

Этот пакет является текущим golden reference для feature flow: здесь уже присутствуют `README.md`, `feature.md`, `spec.md`, `runtime-surfaces.md`, `implementation-plan.md`, optional `use-cases/README.md` и accepted ADR-зависимости.

## Документы фичи

| Документ | Назначение |
|----------|------------|
| [`feature.md`](feature.md) | Canonical feature scope: problem statement, behavioral contracts, failure semantics и verify inventory |
| [`spec.md`](spec.md) | Canonical solution spec: selected design, stable contracts, persistence shape и rollout/backout baseline |
| [`runtime-surfaces.md`](runtime-surfaces.md) | Feature-support reference: current runtime inventory, semantic mapping, target resolution reference и context notes для order/payment surfaces |
| [`implementation-plan.md`](implementation-plan.md) | Derived execution plan: grounded touchpoints, sequencing, test strategy и approval gates для реализации |
| [`operator-ui/README.md`](operator-ui/README.md) | Feature-support reference: operator entrypoint, fixed page-key catalog, editor, preview и context reference для Liquid-страниц; не владеет canonical requirements |
| [`operator-ui/mockups/liquid-pages-list.md`](operator-ui/mockups/liquid-pages-list.md) | Low-fidelity markdown mockup каталога Liquid-шаблонов публичных страниц |
| [`operator-ui/mockups/liquid-page-editor.md`](operator-ui/mockups/liquid-page-editor.md) | Low-fidelity markdown mockup экрана управления одной Liquid-страницей |
| [`operator-ui/mockups/liquid-page-preview-paid.md`](operator-ui/mockups/liquid-page-preview-paid.md) | Low-fidelity markdown mockup preview tab для multi-surface key `paid` |
| [`use-cases/README.md`](use-cases/README.md) | Derived companion doc: review-friendly use cases и mapping `UC -> REQ -> EC -> CHK`; не подменяет canonical acceptance inventory из `feature.md` |
| [ADR-0005](../../engineering/adr/0005-ft-4564-order-page-resolution-and-fallback-policy.md) | Accepted архитектурное решение по resolution unit, fallback policy и bounded refactor |
| [ADR-0006](../../engineering/adr/0006-missing-custom-liquid-template-falls-back-to-existing-baseline-page.md) | Cross-feature правило: если custom Liquid template body отсутствует, используется existing baseline page implementation |

`feature.md` остаётся единственным source of truth для `SC-*`, `NEG-*`, `CHK-*` и `EVID-*`. `spec.md` владеет selected solution. `runtime-surfaces.md`, `operator-ui/README.md`, operator mockups и `use-cases/README.md` нужны для grounding, review и traceability, но не владеют canonical acceptance / test inventory.

## Связанные ресурсы

- GitHub Issue: [#4564](https://github.com/BrandyMint/merchantly/issues/4564)
- Связано: FT-4542 / #4454 — Перенос checkout на Liquid + Stimulus; не является precondition для FT-4564, потому что checkout (`orders/new`) остаётся `NS-01`
- Связано: #4501 — Перенос остальных страниц магазина на Liquid
- Архитектурное обсуждение: [ADR-0005](../../engineering/adr/0005-ft-4564-order-page-resolution-and-fallback-policy.md)
- Platform rule для future custom Liquid features: [ADR-0006](../../engineering/adr/0006-missing-custom-liquid-template-falls-back-to-existing-baseline-page.md)

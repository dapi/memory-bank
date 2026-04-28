---
title: "ADR-0006: Missing Custom Liquid Template Falls Back to Existing Baseline Page"
doc_kind: adr
doc_function: canonical
purpose: "Фиксирует cross-feature policy для optional custom Liquid templates: если Liquid выбран для page key, но template body отсутствует, система использует existing baseline page implementation."
derived_from:
  - ../../../../memory-bank/dna/governance.md
  - ./0003-per-action-engine-override.md
status: active
decision_status: accepted
date: 2026-04-23
audience: humans_and_agents
must_not_define:
  - current_system_state
  - implementation_plan
---

# ADR-0006: Missing Custom Liquid Template Falls Back to Existing Baseline Page

## Контекст

Новые feature могут добавлять vendor-configurable Liquid templates поверх уже существующих customer-facing pages. В таком режиме появляются два разных факта:

1. Для page key выбран режим `Liquid`.
2. Для page key реально существует Liquid template body, который можно отрендерить.

Если смешать эти факты, легко получить неявное поведение:

- считать отсутствие template body ошибкой поставки и падать fail-fast;
- молча подменять любые Liquid-проблемы fallback-ом;
- случайно распространить fail-fast из explicit `engine:` override на optional custom Liquid features.

Нужен cross-feature baseline, который отделяет отсутствие custom Liquid template body от runtime bug уже выбранного Liquid-render.

## Решение

1. Для optional custom Liquid features система обязана различать:
   - `page key requested Liquid`
   - `Liquid template body exists`
2. Если для page key выбран Liquid и Liquid template body существует, система рендерит Liquid.
3. Если для page key выбран Liquid, но Liquid template body отсутствует, система рендерит existing baseline page implementation для той же страницы / surface.
4. Этот fallback происходит на resolution layer до explicit engine override path. Отсутствие custom Liquid template body не считается нарушением fail-fast contract из ADR-0003.
5. Runtime render error после того, как Liquid уже выбран и template body существует, этим ADR не переводится в React fallback. Такое поведение определяется feature-level accepted ADR/spec и по умолчанию остаётся fail-fast.
6. Любая новая feature, добавляющая vendor-configurable custom Liquid templates поверх существующих customer-facing pages, должна наследовать это правило по умолчанию, если другой accepted ADR явно не фиксирует исключение.

## Последствия

### Положительные

- Частичный rollout custom Liquid templates не превращается в user-facing outage только из-за отсутствия template body.
- Будущие feature получают один platform-level baseline вместо локальных ad hoc fallback-решений.
- Rule boundary становится явной: `missing template body` и `runtime render error` — это разные failure classes.

### Отрицательные

- Нужно поддерживать явный mapping `page key -> existing baseline page implementation`.
- Команда должна помнить, что наличие opt-in и наличие template body — разные инварианты.

### Нейтральные / организационные

- ADR-0003 сохраняет fail-fast semantics для explicit `engine:` override.
- Feature-level ADR/spec могут ужесточать runtime error policy, но не должны молча переопределять правило отсутствующего template body.

## Связанные решения

- [ADR-0003](0003-per-action-engine-override.md) — explicit engine override и его fail-fast baseline.
- [ADR-0005](0005-ft-4564-order-page-resolution-and-fallback-policy.md) — первое feature-level применение этого правила.

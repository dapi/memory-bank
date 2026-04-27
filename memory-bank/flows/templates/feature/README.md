---
title: FT-XXX Feature README Template
doc_kind: feature
doc_function: template
purpose: Governed wrapper-шаблон для feature-level `README.md`. Читать, чтобы инстанцировать bootstrap-safe routing-layer фичи без смешения wrapper-метаданных и frontmatter целевого README.
derived_from:
  - ../../feature-flow.md
  - ../../../dna/frontmatter.md
status: active
audience: humans_and_agents
template_for: feature
template_target_path: ../../../features/FT-XXX/README.md
---

# FT-XXX Feature Template

Этот файл описывает сам template wrapper. Инстанцируемый feature README живет ниже как embedded contract и копируется в feature package без wrapper frontmatter и history.

## Wrapper Notes

Каталог `memory-bank/flows/templates/feature/` хранит wrapper-шаблоны feature package: этот README-шаблон, canonical feature templates для short и large фич, canonical `solution.md` template, derived template для `implementation-plan.md` и optional support templates из `support/`. При создании нового feature package embedded README должен оставаться bootstrap-safe: сначала он маршрутизирует только на instantiated `feature.md`, а `solution.md`, feature-support docs, `implementation-plan.md` и связанные ADR добавляются уже после появления соответствующих документов.

Downstream routes для living feature package добавляются по мере прохождения lifecycle stages. Типовой пример таких post-bootstrap routes:

- [`solution.md`](solution.md)
  Читать, когда нужно: после `Problem Ready` зафиксировать или проверить selected design, to-be C4 architecture model, accepted local decisions, contracts и local rollout/backout semantics.
  Отвечает на вопрос: как именно feature реализуется без смешения solution space с problem space.

- [`implementation-plan.md`](implementation-plan.md)
  Читать, когда нужно: после `Solution Ready` разложить реализацию по шагам, workstreams, checkpoints и traceability к canonical IDs.
  Отвечает на вопрос: как провести реализацию фичи от текущего состояния до приёмки.

- [`runtime-surfaces.md`](runtime-surfaces.md)
  Читать, когда нужно: для сложной runtime-фичи зафиксировать current surfaces, semantic mapping, context matrix, adjacent out-of-scope surfaces и resolution behavior.
  Отвечает на вопрос: какие runtime surfaces реально существуют и как они мапятся на selected solution. Support doc, не canonical owner.

- [`ui-reference/README.md`](ui-reference/README.md)
  Читать, когда нужно: feature меняет interface, screen states, navigation, editor/preview UX или user-facing controls. Mockups должны быть linked, default — Markdown в `ui-reference/mockups/`.
  Отвечает на вопрос: как должен выглядеть и вести себя interface без project-specific UI assumptions в generic template. Support doc, не canonical owner.

- [`use-cases/README.md`](use-cases/README.md)
  Читать, когда нужно: сложные сценарии нужно разложить в review-friendly форму и связать с canonical `REQ-*`, `SC-*`, `CHK-*`.
  Отвечает на вопрос: какие derived user journeys и test case candidates помогают проверить фичу. Support doc, не canonical owner.

- `../../../adr/ADR-XXX.md`
  Читать, когда нужно: если по фиче существует связанный ADR, оформить или проверить его с корректным `decision_status`.
  Отвечает на вопрос: почему по фиче выбирается конкретное архитектурное или инженерное решение и на каком оно этапе.

## Instantiated Frontmatter

```yaml
title: "FT-XXX: Feature Package"
doc_kind: feature
doc_function: index
purpose: "Bootstrap-safe навигация по документации фичи. Читать, чтобы сначала перейти к canonical `feature.md`, а downstream `solution.md`, optional support docs и `implementation-plan.md` добавлять только после их появления."
derived_from:
  - ../../dna/governance.md
  - feature.md
status: active
audience: humans_and_agents
```

## Instantiated Body

```markdown
# FT-XXX: Feature Package

## О разделе

Каталог feature package хранит canonical `feature.md`, а downstream solution/support/execution routes добавляются только после появления соответствующих документов. Сначала читай `feature.md`, затем расширяй routing по мере появления `solution.md`, optional support docs, `implementation-plan.md` и связанных ADR.

## Аннотированный индекс

- [`feature.md`](feature.md)
  Читать, когда нужно: открыть instantiated canonical feature-документ сразу после bootstrap нового feature package.
  Отвечает на вопрос: где находятся problem space, canonical verify contract и stable IDs для этой фичи.
```

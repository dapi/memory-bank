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

Каталог `memory-bank/flows/templates/feature/` хранит wrapper-шаблоны feature package: этот README-шаблон, canonical `brief.md` template, conditional `design.md` template и derived template для `implementation-plan.md`. Новые packages всегда используют `brief.md`, а `design.md` добавляют только когда `brief.md` фиксирует `Design required: yes`.

При создании нового feature package embedded README должен оставаться bootstrap-safe: сначала он маршрутизирует только на instantiated `brief.md`, а `design.md`, `implementation-plan.md` и связанные ADR добавляются уже после появления соответствующих документов.

Downstream routes для living feature package добавляются по мере прохождения lifecycle stages. Типовой пример таких post-bootstrap routes:

- [`design.md`](design.md)
  Читать, когда нужно: после `Problem Ready` зафиксировать или проверить selected design, to-be C4 architecture model, accepted local decisions, contracts и local rollout/backout semantics.
  Отвечает на вопрос: как именно feature реализуется без смешения solution space с problem space.

- [`implementation-plan.md`](implementation-plan.md)
  Читать, когда нужно: после готовности upstream owners разложить реализацию по шагам, workstreams, checkpoints и traceability к canonical IDs.
  Отвечает на вопрос: как провести реализацию фичи от текущего состояния до приёмки.

- `../../../adr/ADR-XXX.md`
  Читать, когда нужно: если по фиче существует связанный ADR, оформить или проверить его с корректным `decision_status`.
  Отвечает на вопрос: почему по фиче выбирается конкретное архитектурное или инженерное решение и на каком оно этапе.

## Instantiated Frontmatter

```yaml
title: "FT-XXX: Feature Package"
doc_kind: feature
doc_function: index
purpose: "Bootstrap-safe навигация по документации фичи. Читать, чтобы сначала перейти к canonical `brief.md`; downstream routes добавляются только после появления соответствующих документов."
derived_from:
  - ../../dna/governance.md
  - brief.md
status: active
audience: humans_and_agents
```

## Instantiated Body

```markdown
# FT-XXX: Feature Package

## О разделе

Каталог feature package начинается с canonical `brief.md`. Downstream solution/execution routes добавляются только после появления соответствующих документов. Сначала читай `brief.md`, затем расширяй routing по мере появления design, implementation plan и связанных ADR.

## Аннотированный индекс

- [`brief.md`](brief.md)
  Читать, когда нужно: открыть instantiated canonical feature-документ сразу после bootstrap нового feature package.
  Отвечает на вопрос: где находятся problem space, canonical verify contract и stable IDs для этой фичи.

После появления downstream-документов добавь сюда routes для `design.md`, `implementation-plan.md` и связанных ADR.
```

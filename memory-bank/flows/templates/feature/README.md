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

Каталог `memory-bank/flows/templates/feature/` хранит wrapper-шаблоны feature package: этот README-шаблон, canonical feature templates для short и large фич и derived template для `implementation-plan.md`. При создании нового feature package embedded README должен оставаться bootstrap-safe: сначала он маршрутизирует только на instantiated `feature.md`, а optional `implementation-plan.md` и связанные ADR добавляются уже после появления соответствующих документов.

Optional routes для living feature package добавляются после появления соответствующих документов. Типовой пример таких post-bootstrap routes:

- [`implementation-plan.md`](implementation-plan.md)
  Читать, когда нужно: после появления этого файла разложить реализацию по шагам, workstreams, checkpoints и traceability к canonical IDs.
  Отвечает на вопрос: как провести реализацию фичи от текущего состояния до приёмки.

- `../../../adr/ADR-XXX.md`
  Читать, когда нужно: если по фиче существует связанный ADR, оформить или проверить его с корректным `decision_status`.
  Отвечает на вопрос: почему по фиче выбирается конкретное архитектурное или инженерное решение и на каком оно этапе.

## Instantiated Frontmatter

```yaml
title: "FT-XXX: Feature Package"
doc_kind: feature
doc_function: index
purpose: "Bootstrap-safe навигация по документации фичи. Читать, чтобы сначала перейти к canonical `feature.md`, а optional derived docs добавлять только после их появления."
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

Каталог feature package хранит canonical `feature.md`, а optional derived/external routes добавляются только после появления соответствующих документов. Сначала читай `feature.md`, затем расширяй routing по мере появления execution и decision artifacts.

## Аннотированный индекс

- [`feature.md`](feature.md)
  Читать, когда нужно: открыть instantiated canonical feature-документ сразу после bootstrap нового feature package.
  Отвечает на вопрос: где находятся scope, design, verify, blockers и canonical IDs для этой фичи.
```

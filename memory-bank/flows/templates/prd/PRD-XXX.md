---
title: "PRD-XXX: Product Initiative Name"
doc_kind: prd
doc_function: template
purpose: Governed wrapper-шаблон PRD. Читать, чтобы инстанцировать компактный Product Requirements Document без смешения wrapper-метаданных и frontmatter будущего PRD.
derived_from:
  - ../../../dna/governance.md
  - ../../../dna/frontmatter.md
  - ../../../domain/problem.md
status: active
audience: humans_and_agents
template_for: prd
template_target_path: ../../../prd/PRD-XXX-short-name.md
canonical_for:
  - prd_template
---

# PRD-XXX: Product Initiative Name

Этот файл описывает wrapper-template. Инстанцируемый PRD живет ниже как embedded contract и копируется без wrapper frontmatter и history.

## Wrapper Notes

PRD в этом шаблоне intentionally lean. Он фиксирует продуктовую проблему, пользователей, goals, scope и success metrics, но не берет на себя implementation sequencing, architecture decisions или verify/evidence contracts downstream feature package.

PRD опирается на `domain/problem.md`, а не подменяет его. Не копируй в него весь project-wide контекст, если он уже стабильно описан upstream.

Используй PRD как upstream-слой между общим контекстом проекта и несколькими feature packages. Если инициатива локальна и не требует отдельного product-layer документа, PRD можно не создавать.

## Instantiated Frontmatter

```yaml
title: "PRD-XXX: Product Initiative Name"
doc_kind: prd
doc_function: canonical
purpose: "Фиксирует продуктовую проблему, целевых пользователей, goals, scope и success metrics инициативы."
derived_from:
  - ../domain/problem.md
status: draft
audience: humans_and_agents
must_not_define:
  - implementation_sequence
  - architecture_decision
  - feature_level_verify_contract
```

## Instantiated Body

```markdown
# PRD-XXX: Product Initiative Name

## Problem

Какую пользовательскую или бизнес-проблему решает инициатива. Описывай язык проблемы, а не решение. Ссылайся на общий контекст из `../domain/problem.md` и фиксируй только delta этой инициативы.

## Users And Jobs

Кто является основным пользователем и какую работу он пытается выполнить.

| User / Segment | Job To Be Done | Current Pain |
| --- | --- | --- |
| `primary-user` | Что хочет сделать | Что мешает сегодня |

## Goals

- `G-01` Какой продуктовый outcome обязателен.
- `G-02` Какой дополнительный outcome желателен.

## Non-Goals

- `NG-01` Что сознательно не входит в инициативу.
- `NG-02` Что нельзя молча додумывать на уровне реализации.

## Product Scope

Опиши scope на уровне capability, а не change set.

### In Scope

- Что должно стать возможным для пользователя или системы.

### Out Of Scope

- Что остается за границами инициативы.

## UX / Business Rules

- `BR-01` Важное правило продукта или операции.
- `BR-02` Ограничение, которое должна уважать любая downstream feature.

## Success Metrics

| Metric ID | Metric | Baseline | Target | Measurement method |
| --- | --- | --- | --- | --- |
| `MET-01` | Что измеряем | От чего стартуем | Что считаем успехом | Как проверяем |

## Risks And Open Questions

- `RISK-01` Что может сорвать инициативу на уровне продукта.
- `OQ-01` Какая неизвестность еще не снята.

## Downstream Features

Перечисли ожидаемые feature packages, если они уже понятны.

| Feature | Why it exists | Status |
| --- | --- | --- |
| `FT-XXX` | Какой slice реализует | planned / draft / active |
```

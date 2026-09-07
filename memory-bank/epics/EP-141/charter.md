---
title: "EP-141: Charter"
doc_kind: epic
doc_function: canonical
purpose: "EP-141: Charter"
derived_from:
  - ../../flows/epic.md
status: active
audience: humans_and_agents
---

# EP-141: Charter

## Problem
Существующий payload требует AI flows даже при использовании только проектной документации.

## Outcome
Установка DNA, DNA + Documents или полного набора; явное подключение документов к flow;
безопасное обновление и opt-in миграция legacy с сохранением пользовательских документов.

## Scope
Требования и acceptance из [issue 141](https://github.com/dapi/memory-bank/issues/141) — исходный контракт.
Template владеет компонентами, contracts, migration map, priming и примерами.
CLI владеет transaction, lock, composition, adoption, validation и compatibility gate.

## Non-Scope
Автоматический uninstall/downgrade, внешняя authority/audit, несколько flow contracts у документа,
независимое версионирование компонентов, изменения пользовательских live-установок.

## Stakeholder Channels
Danil принимает продуктовые решения в issue и текущей сессии; PR содержит результат проверки.

## Source / Evidence Boundaries
Template baseline `f1f04de843aef45a2425d4a7351d577bbf89e940`; CLI baseline фиксируется delivery plan.
Issue задаёт ожидаемое поведение; source и tests доказывают фактическое поведение.

## Acceptance
Проверки всех составов и адаптеров, docs → full, legacy opt-in, integrity и pinned contracts,
atomic rollback, ссылки и project-local projection. Каждый критерий issue связывается с тестом
в delivery brief; epic закрывается только после исполнения и явного подтверждения владельца.

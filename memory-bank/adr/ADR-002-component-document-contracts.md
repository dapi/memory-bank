---
title: "ADR-002: Разделить документацию и процессы через компоненты"
doc_kind: adr
doc_function: canonical
purpose: "Архитектурная граница DNA, Documents, Flows и explicit adoption."
derived_from:
  - ../dna/principles.md
  - ../epics/EP-141/charter.md
status: active
decision_status: accepted
date: 2026-09-06
decision_makers:
  - Danil Pismenny
audience: humans_and_agents
---

# ADR-002: Разделить документацию и процессы через компоненты

## Контекст

[Issue 141](https://github.com/dapi/memory-bank/issues/141) требует поэтапного внедрения
DNA и проектных документов без AI-процессов. Текущие DNA и шаблоны зависят от Flows.
Из существующих решений прочитан [ADR-001](ADR-001-introduce-design-pack.md): design pack
сохраняет смысл внутри flow и не становится обязательным для базовых документов.
Код установки принадлежит memory-bank-cli `internal/ownership`, а payload — `template/`.

## Драйверы

Самодостаточные компоненты; один владелец каждого факта; сохранение документации при
подключении процессов; отсутствие неявного ослабления legacy gates; атомарные обновления.

## Варианты

| Вариант | Достоинства | Ограничения |
| --- | --- | --- |
| Один repository, декларативный manifest и explicit adoption | Общая версия, выборочный install, проверяемая совместимость | Registry и migrations увеличивают CLI contract |
| Раздельные repositories | Независимая поставка и владельцы компонентов | Дополнительное согласование версий и обновлений |
| Сохранить текущую установку | Нет миграционных рисков | Не выполняет принятый интент; допустим только как закреплённый legacy путь |

## Решение

Выбран первый вариант в рамках разрешения реализовать issue. DNA задаёт общую metadata
и governance; Documents — типы и шаблоны; Flows — процессные расширения. Зависимости идут
только от Flows к Documents/DNA и от Documents к DNA. Интеграции исполнителей опциональны.

Adoption registry является владельцем подключения документа к immutable flow bundle;
frontmatter — проверяемая проекция. Lock фиксирует целостность registry и версию bundle.
Это защита от несогласованного drift, а не доказательство истории действий владельца.
Точный исполнимый контракт находится в [CTR-01](../../docs/component-wire-format.md).

## Последствия

Положительные: можно начать с документации и подключить процессы без автоматического
назначения новых gates. Версии проверок действующих документов не меняются незаметно.
Отрицательные: перенос или подключение принятого в flow документа требует явной CLI-операции;
необходимо сопровождать legacy bundles и расширенные transaction fixtures.
Операционные: bridge/supporting CLI поставляется до component source; pre-bridge пользователям
остаётся закреплённый источник. Автоматический downgrade не поддерживается.

## Подтверждение

Preset/adapter matrix, docs → full, legacy pass/fail, integrity, immutable bundle,
atomic rollback и source projection проверяются автоматизированно. Design и delivered diff
проходят независимые code-converge review. Решение принято в пределах поручения реализовать issue. Полный review candidate c22294b
оставил четыре замечания; их исправления и последующее уточнение hard-link scope получили
clean structured verdict code-converge 2026-09-07T01:14:17Z и зафиксированы в b94560c.


| Evidence | Owner | Storage / acceptance |
| --- | --- | --- |
| Preset/adapter/ownership, adoption tampering, frozen-rule and migration pass/fail fixtures | CLI #62 author | CLI test sources and exact-commit GitHub Actions run linked from CLI PR; all required cases pass |
| Semantic dependency, base-template, wrapper and projection checks | FT-141 author | Template tools/CI and exact-commit run linked from template PR; core/docs/full checks pass |
| Bridge/component binary matrix and pre-bridge entrypoint refusal | CLI #62 + FT-141 integration owners | Binary commit/SHA-256 and command results in related PR descriptions |
| Independent design, implementation and simplify verdicts | code-converge reviewers, fixes by authors | Structured local session results tied to reviewed revisions; PR records verdicts and revisions |

Reconsider this decision if independent component release cadence becomes necessary, the
frozen-engine maintenance cost exceeds the value of compatibility, or a required adoption
transition cannot be expressed safely under the single-owner model. Such a change needs a
new ADR and an explicit migration; it cannot redefine an existing bundle ID.

Downstream follow-ups: CLI #62 owns source classification, lock/adoption transactions and
validator implementation; FT-141 owns payload separation, base templates, wrappers and
source projection; EP-141 owns cross-repository integration and release ordering. Release
publication remains with the release owner after PR review and is outside this task.

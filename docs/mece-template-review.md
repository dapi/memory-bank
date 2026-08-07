# MECE Review: `template/memory-bank`

Дата ревизии: 2026-08-07  
Ветка: `docs/mece-template-review`

## Scope и метод

Проверены DNA, индексы, document boundaries, canonical ownership, routing и
feature artifact catalog. Проверка отвечала на два вопроса:

1. не создают ли соседние категории или документы конкурирующих владельцев;
2. покрывает ли каждая заявленная классификация свой scope, включая fallback
   для пограничных случаев.

Human-only каталог `template/memory-bank/prompts/` не проверялся семантически;
его содержимое не является workflow dependency. Структурная достижимость
проверена общим lint.

## Результат

Критических MECE-нарушений в template не найдено. Большинство потенциальных
пересечений уже разрешены одним из трёх способов:

- canonical owner и `must_not_define` фиксируют границы ответственности;
- routing order задаёт precedence для взаимоисключающего выбора flow;
- reference views и cross-references явно отделены от новых facts и ownership.

| Область | Наблюдение | Оценка |
| --- | --- | --- |
| DNA / governance | SSoT, dependency tree и `canonical_for` дают однозначную ownership-модель | Сильная сторона |
| Product ↔ domain ↔ PRD | README-файлы определяют, какие вопросы принадлежат каждому слою и что они не описывают | Сильная сторона |
| Task Routing | Bug/Incident, Epic/Feature и другие близкие predicates могут пересекаться, но порядок маршрутизации и rerouting rules задают precedence | Контролируемое пересечение |
| Feature artifacts | Catalog задаёт trigger, ownership и relation для optional artifacts; reference views не становятся вторыми owners | Сильная сторона |
| UI surfaces | `public-web`, `admin`, `mobile` и `shared-components` разделены по surface; для отсутствующих surfaces предусмотрено удаление ссылки | Сильная сторона |
| Domain documents | glossary/model/rules/states/events/context-map — не взаимоисключающие категории фактов, а ортогональные представления домена | Допустимо; не применять MECE механически |
| 4+1 views | Logical/Process/Development/Physical/Scenarios намеренно пересекаются через traceability | Допустимо; ownership остаётся у canonical owners |
| Top-level index | Каталоги `product`, `domain`, `flows`, `features`, `adr` и другие — разные navigation/lifecycle layers, а не классификация каждого факта | Допустимо; MECE применять только к локальным классификациям |

## Вывод и действие

В `dna/principles.md` добавлен принцип **«Структурная полнота и непересечение
(MECE — Mutually Exclusive, Collectively Exhaustive)»**. Он применяется к
декомпозициям, классификациям и индексам только в пределах явно заявленного
scope. Для намеренных пересечений правило требует зафиксировать precedence,
fallback или ограничение scope; открытые списки, ортогональные views и
dependency graphs из-под механического MECE-режима исключены.

## Проверки

- `ruby tools/validate-priming-manifests.rb template/memory-bank` — OK;
- `memory-bank-cli lint --scope-root template/memory-bank --entrypoint template/memory-bank/README.md` — OK: 99 файлов, нет broken links, dependency errors или orphan files;
- `memory-bank-cli doctor --profile template` — 0 errors, 0 warnings.


# Протокол адаптации Memory Bank для greenfield-проекта

## Цель

Скопируй generic-шаблон `memory-bank/` из `dapi/memory-bank` в корень текущего git-репозитория и адаптируй его под описанный в репозитории продукт и проект.

Результат должен содержать максимально полный набор подтверждённых project facts, разложенных по правильным canonical owners Memory Bank, и initial PRD продукта или первой продуктовой инициативы.

## Источники

Перед изменениями прочитай:

1. `AGENTS.md`, `CLAUDE.md` и другие repository instructions, если они существуют.
2. Корневой `README` и остальные README-файлы проекта.
3. Документы в `docs/` и других документированных knowledge directories.
4. Конфигурацию, manifests, CI и структуру исходного кода — только для подтверждения инженерных и операционных фактов, которых нет в документации.

Считай источники репозитория authoritative для project-specific фактов. Если источники противоречат друг другу, не выбирай молча: зафиксируй конфликт или open question в подходящем документе.

## Выполнение

1. Проведи inventory найденных фактов о продукте, пользователях, предметной области, архитектуре, разработке и эксплуатации.
2. Если `memory-bank/` отсутствует, скопируй в текущий репозиторий только каталог `memory-bank/` из `dapi/memory-bank`. Если каталог уже существует, не перезаписывай его целиком — адаптируй имеющуюся копию.
3. Прочитай `memory-bank/README.md`, `memory-bank/dna/README.md` и правила document governance.
4. Замени template placeholders подтверждёнными фактами текущего проекта и максимально полно адаптируй:
   - `product/` — problem, vision, users, jobs, outcomes, non-goals, metrics, positioning и roadmap;
   - `domain/` — glossary, actors, entities, relationships, rules, states, events и bounded contexts;
   - `engineering/` — architecture, module boundaries, technology choices, quality attributes, testing, coding и git conventions;
   - `ops/` — local development, configuration, environments, releases и operational constraints;
   - `use-cases/` — устойчивые пользовательские и операционные сценарии, явно следующие из источников;
   - `adr/` — только уже принятые значимые архитектурные решения, найденные в источниках.
5. Создай initial PRD в `memory-bank/prd/` по шаблону `memory-bank/flows/templates/prd/PRD-XXX.md`:
   - зафиксируй problem, users and jobs, goals, non-goals, product scope, business rules, success metrics, risks и open questions;
   - используй стабильный project identifier или `PRD-001`, если в репозитории нет принятой схемы идентификаторов;
   - не добавляй architecture design и implementation sequence в PRD;
   - если источников недостаточно для уверенного утверждения, оставь документ в `draft` и запиши конкретный open question вместо догадки.
6. Обнови все затронутые README-индексы и `derived_from` связи. Каждый созданный документ должен быть достижим из `memory-bank/README.md`.
7. Проведи финальную проверку на полноту и непротиворечивость: каждый найденный устойчивый факт должен либо иметь canonical owner, либо быть явно отмечен как неприменимый, конфликтующий или неизвестный.
8. Запусти `memory-bank lint`. Исправь broken links, orphan documents и ошибки индексной навигации. Если команда недоступна, сообщи об этом как о verification gap и выполни доступную проверку ссылок и структуры.

## Правила адаптации

- Не выдумывай факты, требования, метрики, пользователей, domain rules, архитектуру или operational procedures.
- Не оставляй примерный template content так, будто это факт проекта.
- Не дублируй один факт в нескольких документах: выбери canonical owner, а из остальных мест поставь ссылку.
- Код владеет implementation details; Memory Bank владеет intent, rationale, project rules и contracts.
- Сохраняй полезные существующие repository instructions и документы; не заменяй их Memory Bank автоматически.
- Не реализуй продуктовые фичи и не изменяй runtime-код в рамках этой адаптации.
- Не создавай feature packages, epic packages или delivery plans, если этого прямо не требует уже существующий источник проекта.

## Результат

Заверши работу, когда:

- `memory-bank/` находится в корне проекта и адаптирован под него;
- подтверждённые факты из README, docs и других изученных источников перенесены в соответствующие owner-документы;
- создан initial PRD и добавлен в `memory-bank/prd/README.md`;
- неизвестные и противоречивые сведения явно перечислены как open questions;
- навигация и ссылки согласованы;
- `memory-bank lint` проходит либо verification gap явно указан.

В финальном ответе перечисли изменённые разделы Memory Bank, созданный PRD, использованные source documents, результат проверки и оставшиеся open questions.

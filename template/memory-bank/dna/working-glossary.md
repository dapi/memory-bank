---
title: Working Glossary
doc_kind: governance
doc_function: reference
purpose: Справочный реестр базовых терминов Memory Bank для людей и агентов. Ссылается на канонических владельцев и не переопределяет их правила.
derived_from:
  - principles.md
  - governance.md
  - frontmatter.md
  - lifecycle.md
  - cross-references.md
  - ../flows/routing.md
  - ../flows/feature.md
  - ../flows/epic.md
  - ../flows/behavior-specification.md
  - ../engineering/autonomy-boundaries.md
  - ../engineering/testing-policy.md
  - ../engineering/validation-profiles.md
  - ../product/roadmap.md
  - ../use-cases/README.md
status: active
audience: humans_and_agents
---

# Working Glossary

Этот справочник объясняет язык самого Memory Bank и управляемой поставки.
Для терминов предметной области проекта используй
[`../domain/glossary.md`](../domain/glossary.md).

Русское объяснение раскрывает смысл, но не заменяет канонический термин.
Состояния, переходы и процессные правила определяет указанный `canonical owner`.
Если запись расходится с ним, сначала обнови владельца, затем этот справочник.

## Документация И Владение Знанием

| Термин | Значение | `canonical owner` |
| --- | --- | --- |
| `SSoT` | `Single Source of Truth`: каждый устойчивый факт имеет одного авторитетного владельца. | [`governance.md`](governance.md) |
| `canonical owner` | Документ или артефакт, который авторитетно определяет конкретный факт. | [`governance.md`](governance.md) |
| `authoritative` | Имеющий силу источник с приоритетом по статусу и dependency tree. | [`governance.md`](governance.md) |
| `upstream` | Источник факта выше по направлению зависимости. | [`governance.md`](governance.md) |
| `downstream` | Документ или артефакт, который зависит от `upstream`-источника. | [`governance.md`](governance.md) |
| `dependency tree` | Направленная структура зависимостей, по которой authority течёт от `upstream` к `downstream`. | [`governance.md`](governance.md) |
| `derived_from` | Поле frontmatter с прямыми upstream-источниками документа. | [`frontmatter.md`](frontmatter.md) |
| `governed document` | Markdown-документ в `memory-bank/` с валидным YAML frontmatter. | [`governance.md`](governance.md) |
| `frontmatter` | YAML-метаданные в начале документа. | [`frontmatter.md`](frontmatter.md) |
| `status` | Публикационный статус документа: `draft`, `active` или `archived`. | [`lifecycle.md`](lifecycle.md) |
| `lifecycle` | Набор разрешённых состояний, переходов и условий перехода, а не просто хронология изменений. | [`lifecycle.md`](lifecycle.md) |
| `reference document` | Справочный документ, который объясняет знания, но не становится их владельцем. | [`governance.md`](governance.md) |
| `index-first` | Каждый документ доступен через навигационный индекс; orphan-файл — дефект. | [`principles.md`](principles.md) |
| `progressive disclosure` | Навигация от короткого обзора к деталям по мере необходимости. | [`principles.md`](principles.md) |
| `cross-reference` | Аннотированная ссылка между документами либо документацией и кодом. | [`cross-references.md`](cross-references.md) |
| `artifact` | Идентифицируемый результат работы: документ, код, отчёт, журнал или другой сохраняемый объект. | [`routing.md`](../flows/routing.md) |

## Запросы И Поставка

| Термин | Значение | `canonical owner` |
| --- | --- | --- |
| `issue` / `task` | Входящий запрос или его запись в task tracker; для Small Change она владеет intent, scope и acceptance. | [`routing.md`](../flows/routing.md) |
| `feature` | Одна связная и проверяемая delivery-unit пользовательского, engineering или operations outcome. | [`feature.md`](../flows/feature.md) |
| `feature package` | Набор governed-документов одной `feature` в `features/FT-XXX/`. | [`feature.md`](../flows/feature.md) |
| `brief` | Канонический problem-space документ: проблема, outcome, scope, constraints и verify contract. | [`feature.md`](../flows/feature.md) |
| `scope` | Явно включённые результаты, границы и обязательства единицы работы. | [`feature.md`](../flows/feature.md) |
| `non-scope` | Явно исключённые результаты и границы; не обязательно обещание будущей работы. | [`feature.md`](../flows/feature.md) |
| `requirement` | Проверяемое утверждение о нужном результате, поведении или ограничении. | [`feature.md`](../flows/feature.md) |
| `acceptance scenario` | Сквозной пример наблюдаемого результата delivery-unit. | [`feature.md`](../flows/feature.md) |
| `problem space` | Требования, ограничения и expected behavior без выбранного технического решения. | [`feature.md`](../flows/feature.md) |
| `design` | Документ выбранного решения, создаваемый, когда `Design required: yes`. | [`feature.md`](../flows/feature.md) |
| `solution space` | Выбранное решение: contracts, invariants, failure modes и rollout/backout rules. | [`feature.md`](../flows/feature.md) |
| `implementation plan` | Производный план технического исполнения принятого результата и решения. | [`feature.md`](../flows/feature.md) |
| `execution space` | Техническое исполнение: grounding, шаги, проверки и согласования. | [`feature.md`](../flows/feature.md) |
| `Design required` | Gate-решение, нужен ли отдельный design layer для feature. | [`feature.md`](../flows/feature.md) |
| `design pack` | Связанный набор design-артефактов; root `design.md` служит его manifest и default owner. | [`feature.md`](../flows/feature.md) |
| `grounding` | Исследование текущего состояния системы до implementation plan. | [`feature.md`](../flows/feature.md) |
| `traceability` | Прослеживаемость requirements к решениям, сценариям, проверкам и evidence. | [`feature.md`](../flows/feature.md) |
| `evidence` | Проверяемое доказательство: результат запуска, отчёт, журнал, снимок экрана или иной artifact. | [`feature.md`](../flows/feature.md) |
| `gate` | Проверяемое условие перехода между стадиями. | [`feature.md`](../flows/feature.md) |
| `checkpoint` | Промежуточная контрольная точка, не завершающая всю поставку. | [`feature.md`](../flows/feature.md) |
| `use case` | Устойчивый пользовательский или операционный сценарий продукта, который может быть источником нескольких features. | [`../use-cases/README.md`](../use-cases/README.md) |
| `roadmap` | Представление направления развития по горизонтам, outcomes и зависимостям; не обещание точных дат. | [`../product/roadmap.md`](../product/roadmap.md) |
| `epic` | Управляемая инициатива крупнее одной feature с общим intent, roadmap, risks и subissues. | [`epic.md`](../flows/epic.md) |
| `ADR` | Отдельный документ архитектурного решения и его rationale. | [`../adr/README.md`](../adr/README.md) |

## Процесс, Проверки И Полномочия

| Термин | Значение | `canonical owner` |
| --- | --- | --- |
| `Task Routing` | Выбор ровно одного подходящего flow до начала delivery. | [`routing.md`](../flows/routing.md) |
| `Small Change` | Локальная поставка, где task уже задаёт intent, scope и acceptance, а design и plan не нужны. | [`small-change.md`](../flows/small-change.md) |
| `validation profile` | Минимальная глубина validation и evidence, независимая от выбранного flow. | [`validation-profiles.md`](../engineering/validation-profiles.md) |
| `Definition of Done` (`DoD`) | Общие условия, при которых задача или стадия действительно завершена. | [`feature.md`](../flows/feature.md) |
| `BDD` | Практика Discovery → Formulation → Automation для concrete behavior examples и их traceability. | [`behavior-specification.md`](../flows/behavior-specification.md) |
| `agent` | Исполнитель, сочетающий модель, контекст, инструкции и инструменты для достижения цели. | [`autonomy-boundaries.md`](../engineering/autonomy-boundaries.md) |
| `working context` | Информация и инструкции, доступные исполнителю в текущем запуске; это не долговременный SSoT. | [`autonomy-boundaries.md`](../engineering/autonomy-boundaries.md) |
| `review` | Проверка артефакта или изменения по явным критериям с findings и verdict. | [`testing-policy.md`](../engineering/testing-policy.md) |
| `simplify review` | Отдельный проход после функциональной проверки для удаления случайной сложности без изменения поведения. | [`testing-policy.md`](../engineering/testing-policy.md) |
| `handoff` | Передача состояния, результатов, evidence и оставшихся gaps следующему участнику или этапу. | [`feature.md`](../flows/feature.md) |
| `autonomy boundary` | Граница действий, которые агент может выполнить сам, под наблюдением или после эскалации. | [`autonomy-boundaries.md`](../engineering/autonomy-boundaries.md) |
| `human gate` | Остановка, когда требуется решение человека из-за существенной неизвестности, конфликта или риска. | [`autonomy-boundaries.md`](../engineering/autonomy-boundaries.md) |
| `approval gate` | Явное согласование конкретного рискованного действия или действия с внешним эффектом. | [`autonomy-boundaries.md`](../engineering/autonomy-boundaries.md) |
| `escalation` | Передача вопроса человеку, поскольку агент не вправе или не может безопасно принять решение. | [`autonomy-boundaries.md`](../engineering/autonomy-boundaries.md) |
| `blocker` | Неразрешённое условие, не позволяющее безопасно продолжить работу или принять результат. | [`autonomy-boundaries.md`](../engineering/autonomy-boundaries.md) |
| `finding` | Подкреплённое наблюдение, полученное при review или audit. | [`testing-policy.md`](../engineering/testing-policy.md) |
| `verdict` | Итоговая оценка проверки, например `pass`, `pass_with_notes` или `fail`. | [`testing-policy.md`](../engineering/testing-policy.md) |
| `severity` | Серьёзность finding; свойство `blocking` оценивается отдельно. | [`testing-policy.md`](../engineering/testing-policy.md) |

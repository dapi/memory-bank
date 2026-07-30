---
title: "ADR-001: Ввести design pack как solution-space package"
doc_kind: adr
doc_function: canonical
purpose: "Фиксирует решение разделить design layer, design pack и root `design.md`, а также назначить design pack владельцем feature-local `CTR-*` с однозначным document-level ownership."
derived_from:
  - ../dna/principles.md
  - ../dna/governance.md
status: active
decision_status: accepted
date: 2026-07-30
decision_makers:
  - Danil Pismenny
consulted: []
informed: []
audience: humans_and_agents
must_not_define:
  - current_system_state
  - implementation_plan
  - feature_specific_design_facts
---

# ADR-001: Ввести design pack как solution-space package

## Контекст

Feature Flow уже допускает несколько solution-space artifacts и требует
индексировать их из `design.md`, однако нормативные формулировки распределены
между flow, artifact catalog, templates и testing policy. Часть формулировок
называет `design.md` владельцем всего solution space и всех `CTR-*`, хотя
отдельный contract artifact может быть непосредственным canonical owner
конкретных contract facts.

Нужно зафиксировать устойчивую модель, которая одинаково работает для
компактной feature с одним `design.md` и для feature с несколькими contracts,
diagrams, migration, security или другими design artifacts. Модель должна
сохранять SSoT, давать агенту однозначный target file для изменения факта и не
превращать каждую ссылку из `design.md` в часть одного пакета.

## Границы решения

Решение распространяется на generic Feature Flow, feature-local solution-space
artifacts, их ownership и `Solution Ready` gate.

Вне scope остаются:

- problem space и acceptance contract, которыми владеет feature `brief.md`;
- execution sequencing, которым владеет `implementation-plan.md`;
- project-wide architecture decisions, которыми владеют соответствующие ADR
  или другие project-level canonical owners;
- обязательное создание дополнительных design artifacts: компактный design
  может оставаться одним файлом.

## Драйверы решения

1. Один canonical owner должен существовать как для solution-space слоя в целом,
   так и для каждого конкретного design fact.
2. Модель должна одинаково описывать однофайловый и многофайловый design.
3. Семантический слой, пакет документов и root-файл не должны смешиваться.
4. Агент должен однозначно определять, какой документ обновлять для конкретного
   `CTR-*` или другого stable ID.
5. External canonical dependencies и derived views не должны случайно становиться
   constituents или владельцами новых facts.
6. `Solution Ready` должен подтверждать согласованность всего design, а не только
   публикационный статус root-файла.
7. Нормативные правила должны иметь минимальное число canonical owners, чтобы
   downstream policies и templates не расходились.

## Рассмотренные варианты

| Вариант | Плюсы | Минусы | Почему рассматривается как основной кандидат / не основной кандидат |
| --- | --- | --- | --- |
| Сохранить `design.md` единственным владельцем solution space | Простая модель и один очевидный файл | Не описывает delegated contracts и другие canonical design artifacts; стимулирует разрастание одного файла или ложное ownership | Не основной: противоречит уже поддерживаемому многофайловому design |
| Использовать design pack только как неформальную группу файлов вокруг `design.md` | Даёт краткое название существующему многофайловому layout без нового lifecycle concept | Не задаёт identity, membership, aggregate ownership, direct ownership или readiness semantics; разные документы продолжат определять пакет по-разному | Не основной: улучшает лексику, но не устраняет governance drift |
| Считать `design layer` и `design pack` синонимами, а `design.md` — их manifest | Вводит имя для многофайлового design и сохраняет root entry point | Смешивает semantic lifecycle layer с его документальным представлением; остаётся неясным непосредственный owner конкретного факта | Не основной: уменьшает файловую неоднозначность, но сохраняет категориальную |
| Разделить `design layer`, `design pack` и root `design.md`; использовать aggregate и direct ownership | Масштабируется от одного файла до нескольких, сохраняет SSoT и позволяет точно маршрутизировать изменения | Требует явного manifest и обновления нескольких governance/template формулировок | Выбранный: единственный вариант, закрывающий все драйверы без запрета delegated ownership |

## Решение

Принято решение использовать следующие определения.

**Design layer** — conditional semantic layer Feature Flow, который владеет
feature-local solution space. Он существует, когда canonical `brief.md`
фиксирует `Design required: yes`.

**Design pack** — документальное представление design layer для одной feature.
Он состоит ровно из одного root `design.md` и нуля или более
проиндексированных design artifacts. Design pack может состоять только из
`design.md`; добавление или удаление companion artifacts не меняет идентичность
пакета, пока он относится к той же feature.

**`design.md`** — обязательный root manifest, entry point, индекс ownership и
default owner тех feature-local solution facts, которые не делегированы другому
canonical document owner.

Design Pack manifest должен различать отношения:

| Relation | Значение |
| --- | --- |
| `root` | Обязательный `design.md`, manifest и default solution owner |
| `constituent` | Artifact внутри design pack, который может владеть явно делегированными canonical facts |
| `derived-view` | Проекция canonical facts для отдельного viewpoint; не вводит новые canonical facts |
| `external-dependency` | Внешний canonical owner, решение или model, импортируемые design pack, но не входящие в его состав |

На aggregate уровне design pack владеет feature-local solution facts, включая
`CTR-*`. Каждый canonical stable ID при этом должен иметь ровно один
непосредственный document owner:

- неделегированный `CTR-*` определяется в `design.md`;
- делегированный `CTR-*` определяется в одном `contracts/<name>.md` или другом
  явно указанном constituent;
- `design.md` индексирует непосредственного owner и не дублирует его canonical
  semantics;
- shared или project-wide contract остаётся у внешнего canonical owner и
  импортируется как `external-dependency`.

`Solution Ready` означает готовность design pack, а не только наличие
`design.md: active`. Root `design.md`, каждый canonical constituent owner и
каждая external canonical dependency, являющиеся governed-документами, должны
иметь `status: active`. Для lifecycle-owning constituent или dependency
дополнительно обязателен status, который его canonical lifecycle определяет как
finalized input для downstream consumption; в частности, делегированный
interaction contract должен иметь `Contract Status: accepted`, а внешний ADR —
одновременно `status: active` и `decision_status: accepted`.

Каждый required `derived-view` должен быть проиндексирован и согласован со
своими canonical owners. Если view является отдельным governed-документом, он
должен иметь `status: active`. Любой embedded или standalone non-document asset,
включая внешний C4 artifact, не получает искусственный YAML status: его
readiness подтверждает `active` governed-документ, который индексирует asset,
указывает его canonical source и version/revision, когда применимо, и фиксирует
его актуальность и согласованность. Каждый canonical ID должен иметь одного
непосредственного owner; между artifacts не должно быть противоречащих
canonical facts.

Нормативные правила распределяются так:

- `flows/feature.md` владеет lifecycle, определениями design layer/design pack,
  общими ownership boundaries и gate semantics;
- `flows/feature-artifact-catalog.md` владеет selection triggers, artifact roles
  и delegated ownership;
- `flows/templates/feature/design.md` владеет формой root manifest и обязательными
  секциями инстанцированного `design.md`;
- testing policy, implementation plan и support templates ссылаются на этих
  owners и не повторяют полную ownership model.

## Последствия

### Положительные

- Одно и то же правило работает для compact и multi-document design.
- `CTR-*` получает однозначное aggregate ownership без потери точного SSoT.
- `design.md` остаётся стабильной точкой входа, но не обязан содержать все
  подробности.
- External dependencies и derived views перестают выглядеть как случайные
  canonical constituents.
- `Solution Ready` начинает проверять реальную согласованность solution space.
- Снижается риск дрейфа между Feature Flow, testing policy и templates.

### Отрицательные

- Manifest становится немного сложнее и требует явного значения `Relation`.
- Авторам и агентам нужно различать aggregate owner и непосредственного
  document owner.
- Existing packages и нормативные документы потребуется мигрировать на новую
  терминологию.
- Полная автоматическая проверка ownership может потребовать новых lint rules.

### Нейтральные / организационные

- `design layer`, `design pack` и `design.md` становятся тремя связанными, но
  различными терминами.
- ADR остаётся rationale принятого выбора; living flow/template rules после
  реализации принадлежат соответствующим downstream canonical owners.
- Дополнительные design artifacts остаются conditional и не превращаются в
  checklist.

## Риски и mitigation

| Риск | Mitigation |
| --- | --- |
| Aggregate ownership скрывает файл, который нужно изменить | Требовать ровно одного непосредственного document owner для каждого canonical ID и индексировать его в root manifest |
| Derived view начинает вводить новые solution facts | Запретить canonical ownership для `derived-view`; новый факт сначала поднимается в `design.md`, constituent owner или ADR |
| Внешний ADR или project model ошибочно включается в состав pack | Использовать отдельное отношение `external-dependency` |
| Flow, catalog и templates снова дублируют полную модель | Закрепить раздельных canonical owners и заменить downstream-нормы ссылками |
| Между ADR и downstream owner возникает dependency cycle | ADR выводить из governance principles; operational owners могут зависеть от принятого ADR, но не наоборот |

## Confirmation

Confirmation owner — maintainer, выполняющий downstream governance update.
Evidence сначала фиксируются в implementation PR, а после реализации ссылки на
них добавляются в эту секцию ADR.

Compliance подтверждается следующими evidence:

- [`flows/feature.md`](../flows/feature.md) содержит единые определения design
  layer/design pack, relation contract, direct ownership и aggregate
  `Solution Ready` gate;
- [`design.md` template](../flows/templates/feature/design.md) реализует manifest
  с `root`, `constituent`, `derived-view` и `external-dependency`;
- [`feature-artifact-catalog.md`](../flows/feature-artifact-catalog.md) и
  [`interaction contract template`](../flows/templates/feature/api-contract.md)
  фиксируют routing и единственного непосредственного owner делегированных
  `CTR-*`;
- [`testing-policy.md`](../engineering/testing-policy.md) ссылается на readiness
  всего design pack вместо безусловного ownership root `design.md`;
- `memory-bank-cli lint --scope-root template/memory-bank --entrypoint template/memory-bank/README.md`
  и `memory-bank-cli doctor --profile template` проходят;
- `git diff --check` не находит whitespace errors.

## Условия пересмотра

ADR следует пересмотреть, если:

- design artifacts получают независимый lifecycle, который root `design.md` не
  может агрегировать;
- один artifact должен одновременно входить в несколько design packs как
  constituent, а `external-dependency` недостаточен;
- появляется project-wide solution package, для которого feature-local identity
  больше не подходит;
- lint/tooling показывает, что aggregate и direct ownership нельзя надёжно
  проверить выбранной моделью;
- новая модель документации устраняет различие между semantic layer и
  publication package.

## Follow-up

Living governance rules, generic/project copies и glossary реализованы в
связанных canonical owners выше.

Остаётся оценить отдельным tooling change автоматические lint checks для
manifest relations и уникального document-level ownership; этот follow-up не
блокирует принятие решения или его текущую документационную реализацию.

## Связанные ссылки

- [`../flows/feature.md`](../flows/feature.md)
- [`../flows/feature-artifact-catalog.md`](../flows/feature-artifact-catalog.md)
- [`../flows/templates/feature/design.md`](../flows/templates/feature/design.md)
- [`../flows/templates/feature/api-contract.md`](../flows/templates/feature/api-contract.md)
- [`../engineering/testing-policy.md`](../engineering/testing-policy.md)
- [`../dna/governance.md`](../dna/governance.md)

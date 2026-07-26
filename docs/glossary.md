# Glossary

Этот глоссарий определяет термины, которыми описаны сам Memory Bank и его
документационная модель. Для терминов предметной области конкретного продукта
используй [`template/memory-bank/domain/glossary.md`](../template/memory-bank/domain/glossary.md).
Нормативные правила и полную schema полей определяют
[`template/memory-bank/dna/`](../template/memory-bank/dna/README.md).

## Durable Knowledge Layer

`Durable knowledge layer` — устойчивый слой знаний проекта: набор versioned
документов, который сохраняет важный контекст между сессиями, участниками и
изменениями в коде.

В этом source-репозитории таким слоем выступает payload
`template/memory-bank/`; после установки в downstream-проекте он живёт как
`memory-bank/`.

Ключевая идея слоя — хранить не эфемерные обсуждения, а проверяемое и
поддерживаемое знание: что в проекте считается истинным, где находится
canonical owner факта и как downstream-документы наследуют этот контекст.

## SSoT

`SSoT` (`Single Source of Truth`) — принцип, по которому каждый факт имеет
ровно одного canonical owner. Если один и тот же факт начинает жить в
нескольких местах, это считается дефектом документации.

## Canonical Owner

`Canonical owner` — документ, который владеет конкретным фактом и имеет
приоритет над downstream-описаниями. Изменение такого документа должно
считаться изменением источника истины, а не просто заметки.

## Governed Document

`Governed document` — Markdown-файл в `memory-bank/`, который подчиняется
governance-правилам и имеет валидный YAML frontmatter. В source-репозитории
такой payload находится в `template/memory-bank/`. Документ задаёт свою роль
и, когда у него есть semantic upstream, явно указывает связь с ним.

## Authoritative Document

`Authoritative document` — governed-документ, который сейчас считается
действующим источником истины. В модели этого шаблона authoritative считается
только документ со `status: active`.

## Dependency Tree

`Dependency tree` — ориентированная структура зависимостей между документами,
построенная через `derived_from`. Authority течёт по ней upstream → downstream,
поэтому изменение корневого или промежуточного документа может потребовать
обновления производных материалов. Циклы запрещены.

## Upstream and Downstream

`Upstream` — документ-источник, от которого наследуется контекст, ограничения
или решения. `Downstream` — документ, который использует этот контекст и не
должен ему противоречить.

## Derived From

`derived_from` — frontmatter-поле, которое перечисляет прямые
upstream-документы. Элементом может быть путь либо объект с `path` и `fit`,
который объясняет scope зависимости. Поле делает происхождение знания явным и
позволяет понять, откуда взялись конкретные требования, ограничения или
решения.

## Canonical For

`canonical_for` — frontmatter-поле, которым документ явно объявляет факты или
артефакты, которыми он владеет. Оно помогает выбрать owner, но не заменяет
SSoT, status и порядок зависимостей.

## Progressive Disclosure

`Progressive disclosure` — правило организации документации, при котором
читатель сначала получает короткий обзор, а потом уходит по ссылкам в детали.
Это удерживает верхний уровень читаемым и не смешивает обзор с
низкоуровневыми подробностями.

## Context Priming (Праймеринг контекста)

`Context priming` / `праймеринг контекста` — подготовка агента к конкретной
задаче через последовательный сбор релевантных project facts до планирования
или реализации. Обычно он включает общий прогрев по индексам проекта,
специализацию на нужной подсистеме и граундинг задачи на актуальном коде,
контрактах и тестах. Праймеринг следует `progressive disclosure`: он не
означает загрузить весь репозиторий или весь Memory Bank в один контекст.
Практическое описание приведено в [статье о праймеринге](context-priming.md).
В template lifecycle `P0` подготавливает выбор route, `P1` специализируется
под первый gate flow, а flow-specific `P2` выполняется только при необходимости.
Каждый P1/P2 запуск использует exact input manifest с конкретными путями,
перечисленными в порядке чтения, а не category-based указаниями.

## Index-First

`Index-first` — правило, по которому каждый документ должен быть достижим из
индекса. Orphan-файл, на который ничто не ссылается и который не встроен в
навигацию, считается дефектом knowledge layer.

## Documentation Layer

`Documentation layer` — структурированный слой знаний с ролями документов,
навигацией и границами ответственности, а не просто папка с Markdown-файлами.
В этом репозитории source template layer живёт в `template/memory-bank/`; его
downstream destination — `memory-bank/`.

## Process Layer

`Process layer` — часть knowledge layer, которая описывает lifecycle,
workflows, gates и шаблоны исполнения. В source template она в основном
сосредоточена в `template/memory-bank/flows/`, а после установки — в
`memory-bank/flows/`.

## Task Routing

`Task Routing` — выбор ровно одного минимального flow для входящей задачи по
её риску и природе работы. Он определяет route, но не lifecycle и не глубину
validation; при изменении scope route повторно выбирается. Порядок выбора
задаёт [`routing.md`](../template/memory-bank/flows/routing.md).

## Flow

`Flow` — governed lifecycle для определённого типа работы: например incident,
bug fix, research, small change, refactoring, epic, use case или feature. Flow
задаёт owner-документы, entry/exit gates, evidence и handoff, но не заменяет
validation profile.

## Gate and Handoff

`Gate` — проверяемое условие, без которого нельзя перейти к следующему этапу
lifecycle. `Entry gate` разрешает начать flow, `transition gate` — сменить
этап, а `closure gate` — завершить его. `Handoff` передаёт подтверждённые facts
и ответственность следующему canonical owner или flow, не копируя их как
второй источник истины.

## Validation Profile

`Validation profile` — минимальная глубина проверок, evidence, approvals и
rollout/backout discipline для delivery work. После Task Routing выбирается
один из профилей: `documentation`, `low-risk`, `standard`, `high-risk` или
`release-deployment`. Профиль не является flow и не меняет ownership
документов.

## Instantiated Document

`Instantiated document` — конкретный документ проекта, созданный из шаблона и
заполненный реальным контекстом. В отличие от template-документа, он уже
описывает не абстрактный формат, а конкретную инициативу, фичу или решение.

## Wrapper Template

`Wrapper template` — governed-шаблон, который сам является отдельным
документом со своей purpose и metadata, но при этом содержит embedded contract
для будущего instantiated документа.

## Embedded Template Contract

`Embedded template contract` — часть wrapper template, которая копируется в
новый instantiated документ. Именно здесь живут frontmatter и body целевого
артефакта, а не в оболочке wrapper-файла.

## Document Kind and Function

`doc_kind` — тип документа или артефакта, например `domain`, `feature`, `adr`
или `research`. `doc_function` — его роль: например `canonical`, `index`,
`template` или `derived`. Вместе они позволяют отличать owner факта от
навигационного индекса, шаблона или производного материала.

## Feature Package

`Feature package` — каталог вида `FT-XXX/`, в котором собраны документы одной
delivery-единицы. Он начинается с canonical `brief.md`; по trigger добавляются
conditional `design.md`, derived `implementation-plan.md`, связанные ADR и
локальный индекс.

## Feature Brief

`Feature brief` — canonical `brief.md` feature package, owner problem space:
intent, scope/non-scope, requirements, acceptance, verify contract и
`delivery_status`. Он не принимает выбранное solution или execution sequence.

## Design Layer

`Design layer` — conditional semantic layer Feature Flow, который владеет
feature-local solution space. Он существует, когда canonical `brief.md`
фиксирует `Design required: yes`. Это lifecycle- и ownership-концепция, а не
конкретный файл или каталог.

## Design Pack

`Design pack` — документальное представление design layer для одной feature:
ровно один root `design.md` и ноль или более проиндексированных design
artifacts. `design.md` служит manifest, entry point и default owner
неделегированных solution facts; constituent может владеть явно делегированными
canonical facts, а derived view только проецирует их.

В отличие от semantic `design layer`, design pack описывает состав и отношения
документов. Связанный ADR или другой внешний canonical owner является
`external-dependency`, а не constituent пакета.

## Reference View

`Reference view` — производное представление canonical facts, созданное для
понятности, review или traceability: например C4, sequence diagram, runtime
surface inventory или UI reference. Оно не вводит requirements, solution
decisions или execution steps.

## Extension Contract

`Extension contract` — правила добавления feature-local artifact, для которого
нет готового template. Такой файл должен уменьшать реальную неоднозначность,
иметь governed frontmatter, быть проиндексирован, явно указать свою роль и
границы ownership, а также ссылаться на canonical IDs вместо копирования facts.

## Stable Identifiers and Traceability

`Stable identifiers` — постоянные IDs фактов и проверки, связывающие problem,
solution, execution и evidence. В feature package `REQ-*` обозначает
требование, `SC-*` — acceptance scenario, `CHK-*` — проверку, `EVID-*` — её
артефакт; `SOL-*` — элемент выбранного решения, а `STEP-*` — шаг плана.
Остальные prefixes и обязательные связи определяет Feature Flow. Эта цепочка
называется `traceability`: от требования через решение и проверку к evidence.

## PRD

`PRD` (`Product Requirements Document`) — документ уровня продуктовой
инициативы. Он фиксирует, что и зачем меняется на уровне инициативы, до
декомпозиции на конкретные feature slices.

## Use Case

`Use case` — канонический проектный пользовательский или операционный сценарий
`UC-*`, который повторяется во времени либо служит upstream для нескольких
features. Это не `SC-*`: acceptance scenario описывает приёмку одной delivery
feature, а use case — устойчивое поведение системы.

## Research Package

`Research package` — пакет `R-XXX/` для evidence-backed ответа на один decision
question до коммита в delivery. `brief.md` владеет вопросом и `research_status`,
`evidence.md` — provenance и observations, `synthesis.md` — findings и
limitations, а `decision.md` — recommendation и promotion/handoff map.

## Evidence, Provenance and Finding

`Evidence` — проверяемый результат проверки, исследования или delivery,
подтверждающий заявленный факт либо outcome. `Provenance` — связь evidence с
источником, доступом и контекстом получения. `Finding` — синтезированный вывод
из evidence, который всегда сохраняет confidence и limitations; он не становится
active canonical fact, пока не передан его owner-у.

## Epic Package

`Epic package` — пакет `EP-XXX/` для инициативы, которая больше одной delivery
feature. Epic владеет intent, roadmap, risks, local decisions и registry
subissues, но не file-level execution. `Epic Intake` фиксирует proposal при
недостатке facts; `Bootstrap Epic` создаёт charter сразу, когда facts уже
достаточны.

## PIR and RCA

`PIR` (`Post-Incident Review`) — разбор инцидента после containment и recovery.
`RCA` (`Root Cause Analysis`) — анализ первопричин в его составе. Они фиксируют
timeline, причины и prevention work; delivery-изменения из них проходят обычный
Task Routing.

## ADR

`ADR` (`Architecture Decision Record`) — документ, который фиксирует
архитектурное решение, его контекст и rationale. В логике `WHY / WHAT / HOW`
ADR отвечает прежде всего на вопрос «почему принято именно это решение».

## Status

`Status` — публикационный статус документа: `draft`, `active` или `archived`.
Он отвечает не за стадию delivery, а за то, считается ли документ действующим
источником истины.

## Delivery Status

`Delivery status` — lifecycle-статус feature-документа, например `planned`,
`in_progress`, `done` или `cancelled`. Его owner — только canonical `brief.md`
feature package. Это отдельная ось состояния, которую нельзя смешивать с
публикационным `status`.

## Research Status

`Research status` — lifecycle-статус canonical research `brief.md`: от
`intake` и `collecting` до `decision_ready`, terminal disposition или
`rerouted`. Он описывает ход исследования, а не delivery и не должен
дублироваться в `plan.md`, `evidence.md`, `synthesis.md` или `decision.md`.

## Decision Status

`Decision status` — lifecycle-статус ADR, например `proposed`, `accepted`,
`superseded` или `rejected`. Он показывает судьбу решения, а не общую
активность самого Markdown-файла.

## Audience

`audience` — необязательное frontmatter-поле с `humans` или
`humans_and_agents`. Оно объявляет границу прямого использования документа:
документ для людей и агентов не может указывать документ с `audience: humans`
как semantic upstream через `derived_from`; обычная индексная ссылка допустима.

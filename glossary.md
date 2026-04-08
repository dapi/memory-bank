# Glossary

## Durable Knowledge Layer

`Durable knowledge layer` — это устойчивый слой знаний проекта: набор versioned документов, который сохраняет важный контекст между сессиями, участниками и изменениями в коде.

В рамках этого репозитория таким слоем выступает `memory-bank`: structured documentation layer с governance-правилами, SSoT и явными связями между документами.

Ключевая идея слоя — хранить не эфемерные обсуждения, а проверяемое и поддерживаемое знание: что в проекте считается истинным, где находится canonical owner факта и как downstream-документы наследуют этот контекст.

## SSoT

`SSoT` (`Single Source of Truth`) — принцип, по которому каждый факт имеет ровно одного canonical owner. Если один и тот же факт начинает жить в нескольких местах, это считается дефектом документации.

## Canonical Owner

`Canonical owner` — документ, который владеет конкретным фактом и имеет приоритет над downstream-описаниями. Изменение такого документа должно считаться изменением источника истины, а не просто заметки.

## Governed Document

`Governed document` — markdown-файл, который подчиняется governance-правилам репозитория. Внутри `memory-bank/` это обычно означает валидный YAML frontmatter, понятную роль документа и явные связи с upstream-источниками.

## Authoritative Document

`Authoritative document` — governed-документ, который сейчас считается действующим источником истины. В модели этого шаблона authoritative считается только документ со `status: active`.

## Dependency Tree

`Dependency tree` — дерево зависимостей между документами, построенное через `derived_from`. Authority течет по нему upstream → downstream, поэтому изменение корневого или промежуточного документа может потребовать обновления производных материалов.

## Upstream And Downstream

`Upstream` — документ-источник, от которого наследуется контекст, ограничения или решения. `Downstream` — документ, который использует этот контекст и не должен ему противоречить.

## Derived From

`Derived from` — frontmatter-поле, которое перечисляет прямые upstream-документы. Оно делает происхождение знания явным и позволяет понять, откуда взялись конкретные требования, ограничения или решения.

## Progressive Disclosure

`Progressive disclosure` — правило организации документации, при котором читатель сначала получает короткий обзор, а потом уходит по ссылкам в детали. Это удерживает верхний уровень читаемым и не смешивает обзор с низкоуровневыми подробностями.

## Index-First

`Index-first` — правило, по которому каждый значимый документ должен быть достижим из индекса. Orphan-файл, на который ничто не ссылается и который не встроен в навигацию, считается дефектом knowledge layer.

## Documentation Layer

`Documentation layer` — не просто папка с markdown-файлами, а структурированный слой знаний с ролями документов, навигацией и границами ответственности. В этом репозитории template layer живет в `memory-bank/`, а project-specific layer показан в `examples/`.

## Process Layer

`Process layer` — часть knowledge layer, которая описывает lifecycle, workflows, gates и шаблоны исполнения. В этом шаблоне она в основном сосредоточена в `memory-bank/flows/`.

## Instantiated Document

`Instantiated document` — конкретный документ проекта, созданный из шаблона и заполненный реальным контекстом. В отличие от template-документа, он уже описывает не абстрактный формат, а конкретную инициативу, фичу или решение.

## Wrapper Template

`Wrapper template` — governed-шаблон, который сам является отдельным документом со своей purpose и metadata, но при этом содержит embedded contract для будущего instantiated документа.

## Embedded Template Contract

`Embedded template contract` — та часть wrapper-template, которая копируется в новый instantiated документ. Именно здесь живут frontmatter и body целевого артефакта, а не в оболочке wrapper-файла.

## Feature Package

`Feature package` — каталог вида `FT-XXX/`, в котором собраны документы одной delivery-единицы: canonical feature, execution-план, связанные ADR и локальный индекс.

## PRD

`PRD` (`Product Requirements Document`) — документ уровня продуктовой инициативы. Он фиксирует, что и зачем меняется на уровне инициативы, до декомпозиции на конкретные feature slices.

## ADR

`ADR` (`Architecture Decision Record`) — документ, который фиксирует архитектурное решение, его контекст и rationale. В логике `WHY / WHAT / HOW` ADR отвечает прежде всего на вопрос «почему принято именно это решение».

## Status

`Status` — публикационный статус документа: `draft`, `active` или `archived`. Он отвечает не за стадию delivery, а за то, считается ли документ действующим источником истины.

## Delivery Status

`Delivery status` — lifecycle-статус feature-документа, например `planned`, `in_progress`, `done` или `cancelled`. Это отдельная ось состояния, которую нельзя смешивать с публикационным `status`.

## Decision Status

`Decision status` — lifecycle-статус ADR, например `proposed`, `accepted`, `superseded` или `rejected`. Он показывает судьбу решения, а не общую активность самого markdown-файла.

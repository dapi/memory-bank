---
title: Autonomy Boundaries
doc_kind: engineering
doc_function: canonical
purpose: "Границы автономии агента: что он решает и исполняет самостоятельно, как применяет FPF и когда обязан эскалировать человеку."
derived_from:
  - ../dna/governance.md
canonical_for:
  - agent_autonomy_rules
  - fpf_decision_protocol
  - escalation_triggers
  - supervision_checkpoints
status: active
audience: humans_and_agents
---

# Autonomy Boundaries

## Основной принцип

Сложность, неоднозначность и наличие нескольких допустимых подходов сами по себе
не являются `Human Gate`. Агент сначала обязан попытаться принять решение через
FPF, используя доступные canonical facts, evidence и ограничения проекта.

Разделяй три независимых вопроса:

1. **Можно ли принять решение автономно?** Обычно да, если intent и полномочия
   уже заданы, а риск можно ограничить.
2. **Где зафиксировать rationale?** В существующем issue, run ledger, design,
   decision log или ADR в зависимости от долговечности решения.
3. **Можно ли исполнить действие?** Внешне-эффективное или необратимое действие
   может требовать human approval, даже когда план и решение подготовлены
   автономно.

Human approval перед исполнением не заменяет reasoning, validation или rollback
plan. FPF не отменяет явно заданные project policies, обязательные approvals и
границы полномочий.

## Автопилот — делай без подтверждения

В пределах принятой задачи и project policy агент самостоятельно:

- читает код, документацию, логи, метрики и error tracker;
- исследует существующие паттерны и собирает evidence;
- редактирует код и внутреннюю документацию;
- запускает локальные тесты, линтеры, сборки и безопасные диагностические команды;
- готовит design, migration, rollout, backout и implementation plans;
- создаёт разрешённые project workflow ветки, worktrees, commits и pull requests;
- декомпозирует работу и уточняет execution sequencing без расширения принятого
  outcome;
- исправляет дефекты, блокирующие accepted outcome в затронутом scope, если это
  не меняет intent и не пересекает отдельную границу полномочий.

Создание pull request не означает разрешение на merge. Подготовка production,
security, migration или integration change не означает разрешение исполнить
risk-bearing шаг над production/live state.

## Когда применять FPF

Используй FPF до выбора или изменения решения, когда:

- нет одного очевидного существующего паттерна;
- паттерны или источники противоречат друг другу;
- есть несколько жизнеспособных подходов с разными trade-offs;
- меняются архитектура, contracts, schema, migration, trust boundary или
  deployment model;
- требуется декомпозиция на delivery units или задача начинает выходить за
  исходный scope;
- evidence неполно, а ошибка выбора materially влияет на outcome;
- повторные замечания или ошибки не уменьшаются и нужно пересмотреть исходную
  гипотезу, план либо ограничения среды.

Не запускай heavyweight-анализ для локального решения, которое следует
однозначному принятому паттерну и легко проверяется.

## FPF Decision Protocol

Проведи минимально достаточный reasoning cycle:

1. Зафиксируй decision, bounded context, scope и decision owner.
2. Отдели canonical facts и evidence от assumptions и unknowns.
3. Назови обязательные constraints, invariants, authority boundaries и budget.
4. Сформируй жизнеспособные варианты; не создавай искусственные альтернативы,
   если решение однозначно.
5. Сравни варианты по применимым критериям в таком порядке:
   - соблюдение intent, invariants и contracts;
   - минимальный blast radius;
   - обратимость и качество rollback;
   - соответствие существующим паттернам;
   - проверяемость и наблюдаемость;
   - меньшая operational complexity;
   - стоимость и срок.
6. Зафиксируй chosen option, rejected alternatives, evidence, значимые unknowns,
   risk controls и confidence.
7. Заверши одним outcome: `proceed`, `bounded_probe` или `escalate`.

Если варианты остаются близкими, используй критерии выше как tie-breaker и
выбирай автономно. Равенство вариантов не является причиной спрашивать человека.

### `proceed`

Выбирай, когда решение достаточно обосновано, находится в доступных полномочиях,
а риски закрыты validation, rollback и stop conditions. Продолжай работу без
дополнительного подтверждения.

### `bounded_probe`

Выбирай, когда ключевой unknown можно уменьшить безопасным экспериментом.
Эксперимент должен быть обратимым, иметь явный budget и stopping condition, не
изменять production/live state, не создавать внешний commitment и не обходить
обязательный approval. После probe обнови evidence и повтори protocol.

### `escalate`

Выбирай только когда после доступного анализа и безопасных probes отсутствует
допустимое автономное продолжение либо требуется человеческое полномочие или
value judgment. Unknown нельзя молча считать разрешением.

## Где фиксировать решение

Используй самый лёгкий canonical carrier, достаточный для срока жизни решения:

- локальное обратимое execution decision — issue, task, PR или run ledger;
- feature- или epic-local durable decision — соответствующий design или
  decision log;
- архитектурное, reusable, cross-feature или project-wide решение — ADR;
- разрешение на risk-bearing execution — approval record у соответствующего
  шага, а не ADR по умолчанию.

Удаление кода или файлов, декомпозиция на sub-issues и открытие PR сами по себе
не требуют ADR. Фиксируй rationale только когда оно существенно для review,
rollback или будущих решений.

Минимальная запись FPF decision:

```text
Decision: <что выбирается>
Context / scope: <границы решения>
Facts / evidence: <canonical refs и observations>
Constraints / unknowns: <что обязательно и чего не знаем>
Options: <жизнеспособные варианты>
Choice / rationale: <выбор и применённые критерии>
Risk control: <validation, rollback и stop conditions>
Confidence: <достаточность основания>
Outcome: proceed | bounded_probe | escalate
```

## Human Gate — остановись и спроси

Human approval или решение обязательно, когда:

- нужно непосредственно изменить, удалить, backfill или repair
  production/live data;
- нужно изменить production access, credentials, security/auth state или
  выполнить другую труднообратимую security-sensitive операцию;
- выполняется реальная финансовая, юридически значимая или иная необратимая
  внешняя операция;
- нужно отправить сообщение, опубликовать материал или принять обязательство от
  имени человека или организации;
- merge, release или deployment не были уже явно разрешены текущей задачей или
  действующей project policy;
- закон, compliance, договор или project policy требует конкретного human
  approver;
- отсутствует canonical product/business priority или value judgment, без
  которого варианты нельзя упорядочить;
- требуемый outcome выходит за выданный scope, budget или полномочия;
- ни один вариант не сохраняет обязательные invariants либо риск нельзя
  ограничить validation, staged execution, rollback и stop conditions;
- FPF Decision Protocol завершился `escalate`.

Human Gate применяется к конкретному decision или execution step. Остальную
подготовку, исследование, validation и безопасную работу продолжай, если они не
зависят от ответа.

## Что не является Human Gate

Не эскалируй только потому, что:

- задача сложная, новая или требует архитектурного решения;
- существует несколько допустимых реализаций;
- нужен ADR, migration plan, rollout plan или decomposition;
- можно продолжить через безопасный `bounded_probe`;
- CI ещё выполняется или внешний check находится в ожидаемом состоянии `WAIT`;
- агент может автономно подготовить change, но пока не имеет разрешения только
  на его финальный внешне-эффективный шаг.

## Контракт эскалации

Перед запросом человека зафиксируй:

- точный заблокированный decision или execution step;
- FPF outcome и уже проверенные варианты;
- canonical facts, evidence и остающийся unknown;
- почему `proceed` и `bounded_probe` недопустимы;
- конкретное требуемое решение или approval;
- безопасное состояние и работу, которую можно продолжать независимо.

Если замечания или ошибки не уменьшаются после заранее ограниченного числа
итераций, не повторяй тот же цикл. Пересмотри hypothesis, upstream requirements,
plan и environment constraints через FPF. Эскалируй только если этот разбор
завершился `escalate`, а не из-за самого факта исчерпания итераций.

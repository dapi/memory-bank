---
title: Autonomy Boundaries
doc_kind: engineering
doc_function: canonical
purpose: "Границы автономии агента: что он решает и исполняет самостоятельно, как применяет Structured Decision Protocol и когда обязан эскалировать человеку."
derived_from:
  - ../dna/governance.md
canonical_for:
  - agent_autonomy_rules
  - structured_decision_protocol
  - decision_carrier_selection_rules
  - execution_authorization_rules
  - approval_evidence_rules
  - escalation_triggers
  - supervision_checkpoints
status: active
audience: humans_and_agents
---

# Autonomy Boundaries

## Основной принцип

Сложность, неоднозначность и наличие нескольких допустимых подходов сами по себе
не являются `Human Gate`. Агент сначала обязан попытаться принять решение через
описанный ниже Structured Decision Protocol, используя доступные canonical
facts, evidence и ограничения проекта.

Structured Decision Protocol является полным обязательным контрактом. First
Principles Framework (FPF) или другая reasoning methodology могут углубить
анализ, но остаются опциональными: их отсутствие не блокирует решение, не
создаёт `Human Gate` и не меняет outcome semantics протокола.

Разделяй три независимых вопроса:

1. **Можно ли принять решение автономно?** Обычно да, если intent и полномочия
   уже заданы, а риск можно ограничить.
2. **Где зафиксировать rationale?** В существующем issue, run ledger, design,
   decision log или ADR в зависимости от долговечности решения.
3. **Можно ли исполнить действие?** Внешне-эффективное или необратимое действие
   может требовать human approval, даже когда план и решение подготовлены
   автономно.

Human approval перед исполнением не заменяет reasoning, validation или rollback
plan. Structured Decision Protocol не отменяет явно заданные project policies,
обязательные approvals и границы полномочий.

## Автопилот — делай без подтверждения

В пределах принятой задачи и project policy агент самостоятельно:

- читает код, документацию, логи, метрики и error tracker;
- исследует существующие паттерны и собирает evidence;
- редактирует обычный non-risky код и внутреннюю документацию;
- запускает локальные тесты, линтеры, сборки и безопасные диагностические команды;
- готовит design, migration, rollout, backout и implementation plans;
- создаёт разрешённые project workflow ветки, worktrees, commits и pull requests;
- декомпозирует работу и уточняет execution sequencing без расширения принятого
  outcome;
- исправляет дефекты, блокирующие accepted outcome в затронутом scope, если это
  не меняет intent и не пересекает отдельную границу полномочий.

Создание pull request не означает разрешение на merge. Агент может автономно
провести analysis, design, планирование, validation и подготовку rollback для
production, security, compliance, migration или integration change. Но сама
repository mutation, materially меняющая security/auth/trust/compliance
boundary, требует specific authority в текущей task или active project policy;
без неё это отдельный Human Gate. Такая authority также не означает разрешение
исполнить risk-bearing шаг над production/live state.

## Когда применять Structured Decision Protocol

Используй Structured Decision Protocol до выбора или изменения решения, когда:

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

## Structured Decision Protocol

Каждая нетривиальная запись протокола различает четыре роли:

- **authority source** — текущая task или active project policy, которая задаёт
  допустимый scope и ограничения;
- **decision owner** — accountable role, принимающая решение среди допустимых
  вариантов;
- **canonical carrier** — единственный durable artifact, владеющий rationale и
  outcome;
- **execution approver / approval evidence** — отдельная роль и ссылка на
  разрешение, только если конкретный execution step пересекает Human Gate.

`Decision owner` не означает автоматически ни владельца carrier, ни execution
approver. Approval evidence фиксирует permission для шага исполнения и никогда
не меняет outcome протокола.

Проведи минимально достаточный reasoning cycle:

1. Зафиксируй decision, bounded context, scope, authority source и decision owner.
2. Отдели canonical facts и evidence от assumptions и unknowns.
3. Назови обязательные constraints, invariants, authority boundaries и budget;
   исключи варианты, которые им не соответствуют.
4. Сформируй жизнеспособные варианты и явно запиши rejected alternatives; не создавай искусственные альтернативы,
   если решение однозначно.
5. После удаления недопустимых вариантов примени tie-breakers в таком порядке:
   существующий canonical pattern; наименьшее обратимое изменение; наименьший
   blast radius; наименьшая operational и maintenance complexity; наиболее
   сильная доступная verification.
6. Зафиксируй chosen option, rationale, evidence, значимые unknowns, risk
   controls и confidence.
7. Запиши execution approver и approval evidence, если они требуются для
   последующего execution step; иначе укажи `not required`.
8. Заверши одним outcome: `proceed`, `bounded_probe` или `escalate`.

Если варианты остаются близкими после этих tie-breakers, выбирай автономно.
Tie-breakers не могут решить отсутствующий product/business value judgment.
Равенство вариантов не является причиной спрашивать человека.

### `proceed`

Выбирай, когда решение достаточно обосновано, находится в доступных полномочиях,
а риски закрыты validation, rollback и stop conditions. Продолжай работу без
дополнительного подтверждения.

### `bounded_probe`

Выбирай, когда ключевой unknown можно уменьшить безопасным экспериментом.
Эксперимент должен быть обратимым, иметь явный budget и stopping condition, не
изменять production/live state, не создавать внешний commitment и не обходить
обязательный approval. После probe обнови evidence и повтори protocol.

До завершения Task Routing полный `bounded_probe` не выполняется: P0 допускает
только read-only сбор facts для классификации route. Если unknown требует
эксперимента, implementation discovery или изменения файлов, выбери Research
Flow; experiment и его stopping condition принадлежат этому lifecycle.

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

Зафиксируй минимальную запись structured decision:

```text
Protocol: Structured Decision Protocol
Context: <task, bounded context и decision question>
Authority source: <текущая task или active project policy>
Decision owner: <accountable decision-making role>
Carrier: <единственный canonical record>
Facts / evidence: <canonical refs и observations>
Assumptions / unknowns: <явный список или none>
Constraints / authority boundaries: <hard limits и budget>
Options / rejected alternatives: <допустимые варианты и причины отказа, если применимо>
Rationale: <почему выбранный вариант победил>
Risk controls: <reversibility, blast radius, validation, rollback и stop conditions>
Confidence: <достаточность основания>
Outcome: proceed | bounded_probe | escalate
Execution approval: <not required | required: approval ref>
Probe budget / stop condition: <required for bounded_probe; otherwise none>
```

В записи должен быть ровно один `Outcome`. `Execution approval` не является
вторым outcome и не превращает `proceed` в разрешение на внешний или
risk-bearing шаг.

## Human Gate — остановись и спроси

Human approval или решение обязательно, когда:

- нужно непосредственно изменить, удалить, backfill или repair
  production/live data;
- нужно изменить production access, credentials, security/auth state или
  выполнить другую труднообратимую security-sensitive операцию;
- нужно внести repository/code/config mutation, materially меняющую
  security/auth/trust/compliance boundary, и текущая task или active project
  policy не даёт specific authority для такого изменения;
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
- Structured Decision Protocol завершился `escalate`.

Human Gate применяется к конкретному decision или execution step. Остальную
подготовку, исследование, validation и безопасную работу продолжай, если они не
зависят от ответа.

### Valid approval evidence

Разрешение текущей task или active project policy считается approval evidence
только если запись одновременно:

- называет конкретное действие или узкий класс действий;
- называет target, environment или external system;
- задаёт scope и существенные limits;
- исходит от canonical task или active policy с понятным owner;
- является актуальной и не отменена более специфичной policy, compliance rule,
  contract или revocation;
- прикреплена к точному execution gate.

Широкое, неоднозначное, выведенное из контекста или устаревшее разрешение не
является approval evidence. Policy может дополнительно требовать fresh approval
для каждого execution.

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
- outcome протокола и уже проверенные варианты;
- canonical facts, evidence и остающийся unknown;
- почему `proceed` и `bounded_probe` недопустимы;
- конкретное требуемое решение или approval;
- безопасное состояние и работу, которую можно продолжать независимо.

Если замечания или ошибки не уменьшаются после заранее ограниченного числа
итераций, не повторяй тот же цикл. Пересмотри hypothesis, upstream requirements,
plan и environment constraints через Structured Decision Protocol. Эскалируй
только если этот разбор не дал bounded продолжения или выявил
authority/value/risk boundary, а не из-за самого факта
исчерпания итераций.

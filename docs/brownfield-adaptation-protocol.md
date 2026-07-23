# Протокол адаптации Memory Bank для brownfield-проекта

## Цель и граница

Этот protocol помогает добавить Memory Bank в уже существующий web service или
CLI utility так, чтобы он отражал наблюдаемое состояние проекта, а не желаемую
картину. Он не заменяет существующие repository instructions, документацию или
код и не реализует продуктовые изменения.

**До шага «Установить и активировать» не открывайте, не копируйте и не
консультируйте `memory-bank/`, включая generic template, его README, governance
и templates.** На этой стадии они не являются источником project knowledge:
placeholders и generic rules нельзя принимать за facts текущего репозитория.

Если repository sources противоречат друг другу или факта нет, не выбирайте
молча: сохраните conflict или open question с источниками, confidence и owner,
если он известен.

## Источники pre-adaptation discovery

Сначала прочитайте repository instructions (`AGENTS.md`, `CLAUDE.md` и
эквиваленты), затем исследуйте только уже существующие project sources:

- root и nested README, wiki, docs, runbooks, design docs и historical ADR;
- source code, manifests, dependency files и critical code paths;
- CI/CD, configuration, deployment/release definitions и observability assets;
- existing task-tracker, operational и ownership references, доступные в scope.

Не извлекайте и не копируйте secret values, PII, tokens или internal endpoints.
Допустимо зафиксировать только owner и безопасный access procedure.

## Lifecycle

### 1. Pre-adaptation discovery

Проведите inventory источников и зафиксируйте только наблюдаемые facts, их
source references, freshness и confidence. На этой стадии Memory Bank не
установлен и не используется.

Минимальный inventory:

- runtime modules, boundaries, dependencies и critical code paths;
- API/UI contracts, clients, queues, scheduled jobs, webhooks и external
  integrations;
- configuration, feature flags, ownership/access procedure для secrets,
  environments, CI/CD, release/rollback, migrations, observability, SLOs и
  alerts;
- existing docs/ADRs/runbooks, их owner и известные freshness concerns;
- product terminology, key user/operational flows, technical debt, risky areas
  и unresolved ownership.

### 2. Создать intake PRD вне Memory Bank

Создайте временный evidence-backed документ
`./brownfield-intake-prd.md` в корне downstream repository. Этот путь —
default; repository может использовать иной уже принятый путь только если он
явно записан, находится вне `memory-bank/` и сохраняет все обязательные поля
ниже.

Intake PRD не governed document и не source of truth после conversion. Он
содержит:

- current product problem, users/jobs, goals, non-goals и scope;
- success signals, risks, assumptions, open questions и conflicts;
- source reference для каждого существенного факта, confidence и известного
  owner/freshness;
- inventory summary и список intentionally unadapted facts/documents.

Не добавляйте invented architecture, selected solution, delivery plan, feature
packages или epics. Unknown означает `Unknown`/`TBD`/`Needs owner confirmation`,
а не правдоподобную догадку.

### 3. Установить и активировать Memory Bank

Только после завершения discovery скопируйте или инициализируйте `memory-bank/`
по [инструкции CLI](memory-bank.md). Затем прочитайте
`memory-bank/README.md`, governance-ядро и применимый flow. Сохраните existing
repository instructions: managed agent block дополняет их, но не заменяет.

### 4. Адаптировать canonical owners из тех же evidence

Переносите durable facts из intake PRD в owner-документы без дублирования:

| Owner layer | Что адаптировать |
| --- | --- |
| `product/` | Current product problem, users, jobs, outcomes, non-goals и известные success signals |
| `domain/` | Glossary, actors, entities, states/events/rules и bounded contexts, подтверждённые sources |
| `engineering/` | Architecture, module boundaries, technology/testing/coding/git conventions и technical constraints |
| `ops/` | Local development, config ownership, environments, releases, rollback и runbooks |

Отмечайте unknown, conflicts и owner-pending facts в соответствующем owner или
явном gap/open-question record вместе с evidence. Для real UI surfaces заполните
релевантные draft-заготовки в `engineering/ui-design-guide/`; неприменимые
файлы и их index links удаляйте только после проверки, что они не нужны проекту.

### 5. Govern intake PRD

После адаптации только нужных upstream owners конвертируйте intake PRD в
`memory-bank/prd/PRD-XXX-*.md` по PRD template. Governed PRD:

- зависит через `derived_from` только от уже адаптированных relevant upstream
  owners;
- сохраняет source references, confidence, conflicts, assumptions и open
  questions, а не превращает их в asserted facts;
- добавляется в `memory-bank/prd/README.md` и достижим из navigation tree.

Temporary intake PRD можно оставить как historical evidence или удалить по
repository retention policy; в обоих случаях governed PRD должен сохранить
нужную provenance. Не превращайте его в second active canonical owner.

### 6. Добавить только подтверждённые durable artifacts

Создавайте или обновляйте use cases только для устойчивых доказанных flows, а
historical ADR — только для уже принятых и всё ещё влияющих решений. Не
создавайте epics, feature packages или delivery plans, пока existing source
явно не требует delivery work.

### 7. Validate и trial

Обновите README indexes и `derived_from` links. Запустите
`memory-bank-cli lint` и `memory-bank-cli doctor`; если команда недоступна,
запишите точную verification gap и выполните доступную проверку
ссылок/структуры. Затем используйте adapted context в одной реальной task через
подходящий flow до объявления rollout complete.

## Product-type considerations

Inventory и owner-документы conditional: не создавайте irrelevant artifacts.
Для каждого unsupported dimension укажите `N/A` и короткую причину.

### Web service

Проверьте и адаптируйте, если применимо:

- API/UI contracts, authentication/authorization, tenancy, rate limits и
  compatibility commitments;
- data stores, caches, asynchronous processing, webhooks, retention и
  migration constraints;
- deployment/rollback, health checks, contract/E2E testing, observability и
  operational response.

### CLI utility

Проверьте и адаптируйте, если применимо:

- commands/subcommands, flags, stdin/stdout/stderr contract, exit codes,
  non-interactive behavior и shell completion;
- installation/distribution, supported OS/architectures, update/uninstall;
- config files, environment variables, filesystem effects, remote API
  compatibility, golden tests и user-visible error compatibility.

## Safety и change control

- Не перезаписывайте existing docs, instructions или runtime code как часть
  adaptation.
- Ссылайтесь на external/legacy sources вместо indiscriminate copying; сохраняйте
  их status и freshness caveats.
- Выполняйте adaptation в reviewable change. Перечислите created, changed и
  intentionally unadapted documents.
- Назначьте follow-up trigger и owner для обновления Memory Bank при изменении
  source code, operations или documented contracts.

## Minimum rollout Definition of Done

- [ ] Baseline `product/`, `domain/`, `engineering/` и `ops/` адаптированы из
      evidence или явно отмечены gaps/`N/A`.
- [ ] Intake PRD находится вне `memory-bank/`; governed/draft PRD имеет
      корректные upstream dependencies и index route.
- [ ] Все known gaps, conflicts и owner-pending facts имеют evidence и owner,
      если он известен.
- [ ] Use cases и historical ADRs созданы только при source evidence; лишние
      delivery artifacts не созданы.
- [ ] `memory-bank-cli lint` и `memory-bank-cli doctor` успешны либо
      verification gap указан без заявления об успехе.
- [ ] Adapted context испытан на одной real task через выбранный flow.
- [ ] Reviewable change перечисляет created, changed и intentionally unadapted
      documents, а также follow-up owner/triggers.

## Copyable Codex prompt

```text
Это brownfield-репозиторий. До явной команды «установить Memory Bank» не
открывай и не консультируй memory-bank/. Сначала прочитай repository
instructions и исследуй только существующие README, docs, code, manifests,
CI/CD, configuration, runbooks и historical ADR. Запиши evidence-backed intake
PRD в ./brownfield-intake-prd.md: facts, sources, confidence, conflicts,
assumptions, open questions и owner/freshness. Не выдумывай architecture или
delivery plan. После discovery установи Memory Bank, адаптируй canonical owners
из того же evidence, конвертируй intake в governed PRD и проверь
memory-bank-cli lint/doctor.
```

---
title: "ADR-XXX: Short Decision Name"
doc_kind: adr
doc_function: template
purpose: Governed wrapper-шаблон ADR. Читать, чтобы инстанцировать decision record без смешения metadata wrapper-документа и frontmatter будущего ADR.
derived_from:
  - ../../../dna/governance.md
  - ../../../dna/frontmatter.md
status: active
audience: humans_and_agents
template_for: adr
template_target_path: ../../../adr/ADR-XXX.md
---

# ADR-XXX: Short Decision Name

Этот файл описывает wrapper-template. Инстанцируемый ADR живет ниже как embedded contract и копируется без wrapper frontmatter и history.

## Wrapper Notes

`status` описывает публикационную готовность документа, а `decision_status` —
lifecycle самого решения. Эти поля не заменяют друг друга:

| Состояние ADR | `status` | `decision_status` |
| --- | --- | --- |
| Документ формируется | `draft` | `proposed` |
| Предложение готово к review | `active` | `proposed` |
| Решение принято | `active` | `accepted` |
| Решение отклонено | `active` | `rejected` |
| Решение заменено другим ADR | `active` | `superseded` |

`decision_status: proposed` означает, что текст ADR является предложением и не
считается принятым решением. До `status: active` документ также не входит в
authoritative set. Не переводи ADR в `accepted`, пока не завершены review,
требуемое согласование и не определён исполнимый Confirmation plan. Evidence
реализации или compliance не является prerequisite для acceptance: собирай и
добавляй его после принятия ADR по мере выполнения downstream work.

`derived_from` перечисляет реальные semantic upstream конкретного решения. ADR
может исходить из feature, epic, research, governance, engineering или другого
canonical context; не создавай фиктивный feature package только ради ссылки.
Избегай dependency cycle: downstream owner, реализующий принятое решение, может
зависеть от ADR, поэтому ADR не должен одновременно объявлять этот downstream
owner своим semantic upstream.

ADR фиксирует выбор, rationale, границы и последствия. После принятия living
project facts и operational rules должны перейти соответствующим canonical
owners; ADR не становится current-state inventory или implementation plan.

## Authoring Method And Quality Gate

Этот шаблон адаптирует
[MADR 4.0.0](https://github.com/adr/madr/tree/4.0.0/template)
([Markdown Architectural Decision Records](https://adr.github.io/madr/)), но не
копирует его дословно. MADR используется как внешний источник проверенных
структурных приемов: explicit problem statement, decision drivers, options с
trade-offs, decision outcome, consequences, Confirmation и review metadata.
Canonical contract для Memory Bank задает этот локальный шаблон; новая версия
MADR не меняет его автоматически.

Если в agent environment доступен skill `adr-writing`, используй его при
создании и проверке ADR. Наличие skill не является скрытой runtime-зависимостью:
локальный quality gate определен здесь через мнемонику **E.C.A.D.R.**

| Критерий | Что должно быть доказано в ADR |
| --- | --- |
| **E — Explicit problem statement** | Контекст называет конкретную проблему, scope, ограничения и причину необходимости решения |
| **C — Comprehensive options analysis** | Рассмотрены минимум два жизнеспособных варианта с плюсами и минусами; status quo включен, когда он реалистичен |
| **A — Actionable decision** | Предлагаемое или принятое решение сформулировано однозначно, связано с drivers и достаточно конкретно для downstream work |
| **D — Documented consequences** | Зафиксированы положительные, отрицательные и организационные последствия, включая будущие издержки |
| **R — Reviewable by stakeholders** | Статусы, участники, язык, ссылки и контекст позволяют провести независимый review |

Перед переводом ADR в `status: active` каждый критерий должен быть выполнен, а
все `[INVESTIGATE: ...]` markers — закрыты. Пока остаются gaps, ADR сохраняет
`status: draft` и `decision_status: proposed`. E.C.A.D.R. является локальным
Definition of Done, а не частью MADR.

## Instantiated Frontmatter

```yaml
title: "ADR-XXX: Short Decision Name"
doc_kind: adr
doc_function: canonical
purpose: "Фиксирует архитектурное или инженерное решение, его текущий `decision_status` и последствия."
derived_from:
  - ../path/to/semantic-upstream.md
status: draft
decision_status: proposed
date: YYYY-MM-DD
decision_makers:
  - Name or role
consulted: []
informed: []
# Optional:
# supersedes:
#   - ADR-YYY
audience: humans_and_agents
must_not_define:
  - current_system_state
  - implementation_plan
```

## Instantiated Body

```markdown
# ADR-XXX: Short Decision Name

## Контекст

Опиши конкретную проблему, ограничение, trade-off или архитектурное напряжение.
Укажи, почему решение требуется сейчас и какие constraints ограничивают выбор.

## Границы решения

- На какие системы, документы, процессы или команды распространяется решение.
- Что явно остается вне scope и каким canonical owners принадлежит.

## Драйверы решения

Перечисли драйверы в порядке приоритета:

- какие требования или ограничения влияют на выбор;
- какие quality attributes, KPI, эксплуатационные или продуктовые факторы важны;
- какие зависимости и уже принятые решения нужно учитывать.

## Рассмотренные варианты

Рассмотри минимум два жизнеспособных варианта. Добавь status quo / «ничего не
менять», если он действительно допустим. У каждого варианта должны быть и плюсы,
и минусы; не используй заведомо слабые strawman options.

| Вариант | Плюсы | Минусы | Почему рассматривается как основной кандидат / не основной кандидат |
| --- | --- | --- | --- |
| `Option A` | Что дает | Какие ограничения создает | Причина |
| `Option B` | Что дает | Какие ограничения создает | Причина |

## Решение

Назови выбранный или предлагаемый вариант, свяжи rationale с драйверами и
зафиксируй достаточно точный normative outcome, чтобы downstream owners могли
его реализовать без нового выбора.

Для `decision_status: proposed` избегай языка финального выбора (`выбрано`,
`окончательно отвергнуто`, `принято`). После перевода ADR в `accepted` обнови
формулировки так, чтобы секция фиксировала уже принятое решение, его границы
действия и затронутые компоненты.

## Последствия

### Положительные

Что упрощается, улучшается или становится возможным.

### Отрицательные

Какие ограничения, долги или дополнительные издержки появляются.

### Нейтральные / организационные

Какие документы, процессы или зоны ответственности нужно обновить после принятия.

## Риски и mitigation

Какие риски остаются после выбора и как мы их снижаем.

## Confirmation

До acceptance определи исполнимый Confirmation plan: какие review, tests, lint,
policy checks, telemetry или другие observable evidence подтвердят реализацию и
продолжающийся compliance, кто их получает и где фиксирует. Не требуй уже
полученного implementation evidence для перевода ADR в `accepted`. После
downstream implementation дополняй эту секцию ссылками на полученные evidence и
результатами проверок. Confirmation проверяет compliance с ADR, а не заменяет
acceptance конкретной delivery-задачи.

## Условия пересмотра

Какие изменения assumptions, constraints, scale или evidence требуют пересмотреть
решение, supersede ADR либо подтвердить его заново.

## Follow-up

Какие downstream canonical owners, документы, задачи, бенчмарки или миграции
должны реализовать принятое решение. Для каждого существенного handoff укажи
owner или target path; не превращай секцию в implementation sequence.

## Связанные ссылки

- feature, epic, research, governance или analysis документы, которые дают контекст;
- связанные ADR, если решение зависит от них или уточняет их.
```

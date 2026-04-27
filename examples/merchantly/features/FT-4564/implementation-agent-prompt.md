---
title: "FT-4564 Implementation Agent Prompt"
doc_kind: feature-support
doc_function: reference
purpose: "Готовый промпт для coding-agent, который должен прочитать весь пакет FT-4564, реализовать фичу по плану и довести её до PR с зелёным CI."
derived_from:
  - ./feature.md
  - ./spec.md
  - ./implementation-plan.md
status: active
audience: humans_and_agents
---

# FT-4564 Implementation Agent Prompt

```text
Цель: полностью реализовать фичу `memory-bank/features/FT-4564/` по документации пакета и довести работу до состояния:
- код реализован по `feature.md`, `spec.md`, `implementation-plan.md` и companion docs;
- все automated tests, нужные для change surface, проходят локально;
- acceptance criteria и canonical checks пакета закрыты или явно подтверждены;
- PR создан;
- CI зелёный;
- в финале дай краткий отчёт: что сделано, какие проверки пройдены, ссылка/номер PR, статус CI, что осталось только ручным verify если такое допустимо по документам.

Проверь что уже реализовано, отметь пункты плана которые выполнены и продолжай
дальше.

Обязательные правила проекта:
- сначала прочитай `README.md` и `AGENTS.md`, затем весь пакет `memory-bank/features/FT-4564/`;
- при необходимости читай связанные документы из `memory-bank/`, на которые ссылается пакет;
- ВСЕ команды запускай через `direnv exec .`;
- если в команде есть `$VAR`, используй `bash -c '...'`;
- соблюдай fixtures-only policy из `memory-bank/engineering/testing-policy.md`;
- не нарушай границы feature package: `feature.md` владеет problem/acceptance, `spec.md` владеет selected solution, `implementation-plan.md` владеет execution;
- не придумывай новую solution architecture поверх принятой, если это не требуется; если наткнёшься на blocker, сначала проверь, разрешён ли он existing docs;
- не делай destructive git operations;
- не работай в текущем checkout: создай отдельную ветку и отдельный `git worktree` для реализации, выполняй всю разработку, тесты, commits, push и PR именно из этого worktree;
- имя ветки должно явно ссылаться на FT-4564;
- не останавливайся на анализе: доведи задачу до рабочего состояния.

Что нужно прочитать обязательно:
1. `README.md`
2. `AGENTS.md`
3. `memory-bank/flows/feature-flow.md`
4. весь каталог `memory-bank/features/FT-4564/`
5. все ADR и reference docs, на которые пакет ссылается и которые нужны для реализации:
   - `memory-bank/engineering/adr/0003-per-action-engine-override.md`
   - `memory-bank/engineering/adr/0005-ft-4564-order-page-resolution-and-fallback-policy.md`
   - `memory-bank/engineering/adr/0006-missing-custom-liquid-template-falls-back-to-existing-baseline-page.md`
   - `memory-bank/engineering/testing-policy.md`
   - `memory-bank/domain/frontend.md` если меняешь operator UI
   - `memory-bank/domain/design-guide/*` если меняешь operator UI

Рабочий режим:
1. Сначала создай отдельную ветку и отдельный `git worktree` для FT-4564, затем работай только внутри него.
2. Изучи документы и кодовую базу, чтобы подтвердить текущий runtime и touchpoints.
3. Сверь реализацию с `implementation-plan.md` и выпиши для себя конкретные шаги `STEP-*`, `WS-*`, `PRE-*`, `STOP-*`, `AG-*`.
4. Реализуй фичу строго по selected solution из `spec.md`.
5. Обновляй docs только если это требуется для синхронизации фактически принятой реализации с canonical package.
6. После каждого существенного этапа прогоняй релевантные тесты.
7. Когда локально всё зелёное, создай осмысленный commit или серию commit-ов.
8. Затем запушь ветку из worktree и создай PR.
9. Проверь CI. Если CI падает, разбери причину, исправь и доведи до зелёного статуса.
10. Не считай работу завершённой, пока PR не создан и CI не зелёный.

Технический фокус реализации:
- FT-4564: per-page Liquid override и кастомные шаблоны для customer-facing order/payment result surfaces;
- соблюдай canonical scope и non-scope, включая boundary для `vendor/orders/payment`;
- сохрани `custom_page` priority;
- реализуй selected persistence surface `VendorLiquidPageTemplate` по `SD-11`;
- реализуй resolver/fallback semantics, в том числе:
  - `custom_page` > page-specific Liquid > baseline;
  - `template body absent` -> baseline policy;
  - runtime Liquid render error -> fail-fast / error reporting policy, без silent fallback;
- реализуй operator authoring surface по package docs;
- не смешивай configured state и effective state;
- для `paid` соблюдай multi-surface semantics;
- preview должен работать по последней сохранённой версии и не должен silently preview-ить unsaved draft;
- сохрани legacy `HTML overrides` flow.

Обязательные проверки перед завершением:
- все релевантные локальные rspec suites из `implementation-plan.md` зелёные;
- если потребовались дополнительные тесты, они добавлены;
- traceability между реализованным кодом и canonical checks не сломана;
- acceptance scenarios `SC-*`, negative cases `NEG-*`, checks `CHK-*` закрыты настолько, насколько это возможно в automation;
- manual-only gaps допустимы только если это прямо соответствует package rules и корректно отмечено;
- PR создан;
- CI зелёный.

Git / PR:
- отдельная ветка и отдельный `git worktree` обязательны;
- ветку создай до начала реализации;
- commit messages делай внятными;
- PR title и body должны кратко отражать FT-4564 и основные изменения;
- если в проекте есть принятый формат PR description, соблюдай его.

Финальный ответ дай в таком формате:
1. Что реализовано
2. Какие файлы изменены
3. Какие тесты прогнаны локально
4. Статус acceptance / canonical checks
5. Commit(s)
6. PR
7. CI
8. Остаточные риски или manual-only verify, если остались

Если по ходу встретишь реальный blocker, который нельзя честно разрешить в рамках уже принятых docs, остановись только после того, как:
- локализуешь blocker,
- покажешь точный конфликт с конкретными файлами/строками,
- предложишь минимальный следующий шаг.
Иначе продолжай до полного завершения.
```

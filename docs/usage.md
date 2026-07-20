# Использование Memory Bank

Этот документ описывает повседневную работу с уже внедрённым Memory Bank. Он остаётся во внешней документации шаблона и не копируется в downstream-проект. Канонические project-side правила находятся в [`memory-bank/README.md`](../memory-bank/README.md) и `memory-bank/flows/`; первичная установка и адаптация описаны в [`adoption.md`](adoption.md).

## Рабочая модель

Memory Bank не заменяет task tracker или инструмент запуска агента. Он хранит контекст и правила, issue задаёт конкретную работу, а agent runner создаёт рабочее окружение и запускает coding agent.

```text
Memory Bank
контекст, правила, требования и способы проверки
        ↓
Issue / Task
задача и ссылки на нужные документы
        ↓
agent runner
branch → worktree → agent session
        ↓
реализация → проверки → PR → evidence
        ↓
обновление Memory Bank при появлении новых знаний
```

Инструменты вроде [`start-issue`](https://github.com/dapi/start-issue) могут автоматизировать создание ветки и worktree и запуск выбранного агента, но не являются обязательной частью Memory Bank.

## Рабочий цикл

1. Подготовьте issue с ожидаемым результатом и ссылками на применимые PRD, epic, use case, feature package или ADR.
2. Выберите workflow по [`memory-bank/flows/routing.md`](../memory-bank/flows/routing.md).
3. Запустите агента в изолированной ветке или worktree.
4. Агент читает issue и связанные owner-документы, реализует изменение и выполняет предусмотренные проверки.
5. Завершите работу через PR и приложите требуемые evidence.
6. Если появились новые устойчивые правила, ограничения или решения, обновите их canonical owner в Memory Bank.

Если issue полностью задаёт intent, scope и acceptance, решение не требует design-документов и все routing predicates выполнены, задача может пройти как `Small Change` напрямую к реализации.

## Стартовые запросы

### Создать feature package

```text
Прочитай ./memory-bank/README.md, ./memory-bank/flows/routing.md
и ./memory-bank/flows/feature.md. Сначала определи route текущей задачи.
Если выбран не Feature Flow, остановись и сообщи подходящий route.
Если выбран Feature Flow, создай feature package, начиная с README.md и brief.md.
design.md создавай только по правилам Design Requirement Decision,
а implementation-plan.md — только после готовности upstream-документов.
```

### Проверить качество Memory Bank

```text
Проведи ревью ./memory-bank на SSoT, противоречия, broken links,
orphan-документы, недостающие README-индексы и неясные зависимости.
Предложи минимальные правки и запусти локальные проверки.
```

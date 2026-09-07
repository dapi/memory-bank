---
title: "Testing Conventions"
doc_kind: engineering
doc_function: convention
purpose: "Testing Conventions"
derived_from:
  - ../dna/governance.md
  - ../ops/development.md
status: active
audience: humans_and_agents
---

# Testing Conventions

Этот документ описывает выбранный проектом testing stack и соглашения.

## Project adaptation

Укажи test frameworks, fixtures/factories, размещение unit/integration tests и обязательные
CI suites. Канонические команды хранятся в [Development](../ops/development.md).
Требования к evidence и порядок review определяются принятой в проекте policy; установка
документации сама по себе не выбирает AI-процесс или автоматического проверяющего.

## Review mechanism

Запиши выбранный механизм review и его команды, если проект его использует.
Отделяй результаты проверки от заявлений об их выполнении; ссылайся на evidence.

## Checklist

- Указаны реальные frameworks и стратегия тестовых данных.
- Описано размещение и назначение suites.
- Команды и CI соответствуют фактическому проекту.

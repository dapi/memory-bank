---
title: Testing Conventions
doc_kind: engineering
doc_function: convention
purpose: "Project-specific testing stack целевой системы: framework, тестовые данные, размещение тестов и обязательные suites."
derived_from:
  - ../dna/governance.md
  - ../flows/testing-policy.md
  - ../ops/development.md
status: active
canonical_for:
  - project_testing_stack
  - project_test_data_conventions
  - project_test_placement_conventions
must_not_define:
  - automated_test_requirements
  - sufficient_test_coverage_definition
  - manual_only_verification_exceptions
  - project_command_contract
audience: humans_and_agents
---

# Testing Conventions

Этот документ описывает, как тесты устроены в конкретном репозитории.

Он не решает, что обязано быть покрыто и когда допустим manual-only verify: этим владеет generic [`../flows/testing-policy.md`](../flows/testing-policy.md). Он также не владеет списком локальных команд — canonical test/lint команды живут в [`../ops/development.md`](../ops/development.md). Здесь фиксируй только стек и конвенции тестов. Усиливать требования policy можно, ослаблять нельзя.

## Project Adaptation

После копирования шаблона заполни project-specific часть testing stack:

- основной test framework;
- стратегия тестовых данных;
- обязательные CI jobs;
- какие suites обязаны быть зелёными перед handoff.

Пример формулировок:

- **Framework:** `pytest`, `rspec`, `go test`, `vitest`
- **Data:** fixtures / factories / builders / seeded test database
- **CI jobs:** `unit`, `integration`, `e2e`

## Project-Specific Conventions

Ниже должен появиться downstream-specific блок после адаптации шаблона. Зафиксируй:

- куда добавлять новые тесты;
- какой helper/setup pattern считается canonical;
- как работать с базой, моками и fixtures;
- какой набор suites обязателен перед handoff (сами команды — в [`../ops/development.md`](../ops/development.md)).

Пример:

- новые unit tests живут в `tests/unit/` или `spec/`;
- integration tests обязаны покрывать changed contract;
- для дорогого setup использовать shared fixtures или builders;
- текстовые assertions не дублируют hardcoded UI-копию, если проект уже владеет переводами централизованно.

## Checklist For Template Adoption

- [ ] указан реальный test framework
- [ ] перечислены обязательные CI suites
- [ ] задокументирован deterministic test data pattern
- [ ] указано, куда добавлять новые тесты
- [ ] canonical test/lint команды зафиксированы в [`../ops/development.md`](../ops/development.md)
- [ ] конвенции не противоречат [`../flows/testing-policy.md`](../flows/testing-policy.md)

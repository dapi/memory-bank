---
title: Testing Conventions
doc_kind: engineering
doc_function: convention
purpose: "Project-specific testing stack целевой системы: framework, тестовые данные, local commands, CI jobs и размещение тестов."
derived_from:
  - ../dna/governance.md
  - ../flows/testing-policy.md
status: active
canonical_for:
  - project_testing_stack
  - project_test_command_contract
  - project_test_data_conventions
must_not_define:
  - automated_test_requirements
  - sufficient_test_coverage_definition
  - manual_only_verification_exceptions
audience: humans_and_agents
---

# Testing Conventions

Этот документ описывает, как тесты устроены в конкретном репозитории. Он не решает, что обязано быть покрыто: этим владеет generic [`../flows/testing-policy.md`](../flows/testing-policy.md). Здесь фиксируй только исполнение — стек, команды и конвенции проекта. Усиливать требования policy можно, ослаблять нельзя.

## Project Adaptation

После копирования шаблона заполни project-specific часть testing stack:

- основной test framework;
- стратегия тестовых данных;
- canonical local commands;
- обязательные CI jobs;
- допустимые manual-only исключения.

Пример формулировок:

- **Framework:** `pytest`, `rspec`, `go test`, `vitest`
- **Data:** fixtures / factories / builders / seeded test database
- **Local commands:** `make test`, `npm test`, `bundle exec rspec`
- **CI jobs:** `unit`, `integration`, `e2e`

## Project-Specific Conventions

Ниже должен появиться downstream-specific блок после адаптации шаблона. Зафиксируй:

- куда добавлять новые тесты;
- какой helper/setup pattern считается canonical;
- как работать с базой, моками и fixtures;
- какие команды обязан прогонять агент перед handoff.

Пример:

- новые unit tests живут в `tests/unit/` или `spec/`;
- integration tests обязаны покрывать changed contract;
- для дорогого setup использовать shared fixtures или builders;
- текстовые assertions не дублируют hardcoded UI-копию, если проект уже владеет переводами централизованно.

## Checklist For Template Adoption

- [ ] указаны реальные local test commands
- [ ] перечислены обязательные CI suites
- [ ] задокументирован deterministic test data pattern
- [ ] описаны manual-only exceptions
- [ ] конвенции не противоречат [`../flows/testing-policy.md`](../flows/testing-policy.md)

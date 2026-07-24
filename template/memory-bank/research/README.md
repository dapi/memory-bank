---
title: Research Packages Index
doc_kind: research
doc_function: index
purpose: Навигация по instantiated research packages. Читать, чтобы провести evidence-backed research до решения о product, marketing или technical direction.
derived_from:
  - ../dna/governance.md
  - ../flows/research.md
status: active
audience: humans_and_agents
---

# Research Packages Index

Каталог `memory-bank/research/` хранит instantiated research packages вида `R-XXX/`.

## Rules

- Создавай package только когда Task Routing выбрал [Research & Discovery Flow](../flows/research.md).
- Один package отвечает на один decision question; несколько независимых questions маршрутизируй отдельно.
- Bootstrap начинается с `README.md` и canonical `brief.md`. `plan.md` создаётся, когда метод не очевиден или нужен collection/experiment; `evidence.md`, `synthesis.md` и `decision.md` появляются по lifecycle gates.
- Research не создаёт committed feature scope, implementation sequence, accepted architecture или roadmap. После disposition устойчивые факты передаются в PRD, epic, feature, ADR, product context или другой canonical owner.
- Для package используй шаблоны из [`../flows/templates/research/`](../flows/templates/research/).

## Naming

- Базовый формат: `R-XXX/`.
- Вместо `XXX` используй issue id, ticket id или другой стабильный ключ.
- Один package = один evidence-backed decision question, а не папка для всех заметок проекта.

## Instantiated Research

В шаблонном репозитории этот каталог может быть пустым. Это нормально.

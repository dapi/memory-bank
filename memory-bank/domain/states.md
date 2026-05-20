---
title: Domain States
doc_kind: domain
doc_function: canonical
purpose: Каноничное место для lifecycle states, allowed transitions, terminal states и state-related invariants.
derived_from:
  - ../dna/governance.md
  - model.md
  - rules.md
status: active
audience: humans_and_agents
canonical_for:
  - domain_states
  - state_transitions
---

# Domain States

Этот документ описывает состояния domain concepts и допустимые transitions. Он не должен превращаться в UI state или implementation state machine, если эти состояния не имеют бизнес-смысла.

## State Machines

| State Machine | Concept | Owner | Notes |
| --- | --- | --- | --- |
| `SM-01` | Какой concept имеет lifecycle | Context / team | Важные ограничения |

## States

| State | Meaning | Entry condition | Exit condition | Terminal |
| --- | --- | --- | --- | --- |
| `state-name` | Что означает для бизнеса | Когда состояние становится истинным | Когда можно выйти | yes / no |

## Transitions

| Transition ID | From | To | Trigger | Preconditions | Forbidden when |
| --- | --- | --- | --- | --- | --- |
| `TR-01` | `state-a` | `state-b` | Domain event / user action / policy verdict | Что должно быть истинно | Когда переход запрещен |

## State Invariants

- `SI-01` Инвариант, который должен сохраняться во всех состояниях.
- `SI-02` Инвариант, который действует только в конкретном state или transition.

## Implementation Notes

Если runtime implementation использует дополнительные technical states, документируй их в code/API docs или [`../engineering/architecture.md`](../engineering/architecture.md), а здесь оставляй только business-visible states.

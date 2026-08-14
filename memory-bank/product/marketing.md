---
title: Marketing And Positioning
doc_kind: product
doc_function: canonical
purpose: Каноничное место для positioning, messaging, go-to-market channels, competitive alternatives и launch constraints.
derived_from:
  - ../dna/governance.md
  - context.md
  - customers.md
status: active
audience: humans_and_agents
canonical_for:
  - product_positioning
  - product_messaging
  - go_to_market_context
---

# Marketing And Positioning

Этот документ фиксирует, как продукт объясняется рынку, customers и internal stakeholders. Он не заменяет PRD и не определяет implementation scope.

## Positioning

| Audience | Current alternative | Product difference | Proof |
| --- | --- | --- | --- |
| Teams using coding agents on long-lived repositories | One `AGENTS.md`, chat history, scattered Markdown and tribal knowledge | Version-controlled product context, explicit ownership, task routing, lifecycle gates and verification contracts in one installable layer | Root README, adoption guide, ownership contract and executable `lint` / `doctor` checks |
| Solo developers running several agent sessions | Re-explain the project in every prompt or keep one increasingly noisy conversation alive | A fresh session can start from stable indexes and task-specific context instead of reconstructing prior chat | Context priming guide and managed agent-instruction block |
| Maintainers adopting a shared template across repositories | Copy a documentation tree once and let local forks drift | Ownership-aware CLI tracks managed, adapted and user-owned files and surfaces conflicts before mutation | `memory-bank/.lock` contract and `memory-bank-cli pull --plan` workflow |

## Messaging

- `MSG-01` `AGENTS.md` tells an agent how to start; Memory Bank preserves what the project means, why decisions were made and how delivery is verified.
- `MSG-02` Memory Bank is a version-controlled context and governance layer, not an agent runner, wiki or task tracker.
- `MSG-03` The CLI makes adoption and updates reviewable: dry-run first, explicit ownership, no silent overwrite of ambiguous local changes.
- `MSG-04` The product is useful when work must survive a fresh agent session, a different coding agent or a different maintainer.

## Channels

| Channel | Audience | Goal | Constraint | Owner |
| --- | --- | --- | --- | --- |
| GitHub README, topics and social preview | Developers already evaluating the repository | Activation | First screen must lead to a five-minute trial without abstract governance language | Maintainer |
| pismenny.ru and technical publications | Russian-speaking engineering leaders and agentic-development practitioners | Awareness and qualified adoption | Claims must distinguish author experience, documented behavior and unvalidated assumptions | Content owner |
| Show HN, relevant Reddit communities and developer social networks | Developers with an active context-loss problem | Awareness and feedback | One evidence-backed launch per project; no repeated cross-post spam | Maintainer |
| Awesome lists and tool documentation | Users selecting coding-agent infrastructure | Discovery | Submit only where the project matches inclusion rules | Maintainer |
| Issues and Discussions | Existing evaluators and adopters | Activation and retention | Respond to real use cases; do not optimize for empty engagement | Maintainer |

## Competitive Alternatives

- `ALT-01` One agent instruction file: cheap and sufficient for small or short-lived repositories, but it does not own product context, decisions, requirements and verification evidence.
- `ALT-02` Chat history or a long-running session: preserves conversational context but is difficult to review, version and resume across tools.
- `ALT-03` A generic docs/wiki tree: stores information but usually does not define source ownership, delivery routing or executable health checks.
- `ALT-04` A full agent orchestration platform: may run work, but does not automatically become the project's canonical knowledge and governance layer.

## Launch Constraints

- `LC-01` Before a broad launch, the repository must have a concise first screen, a working quick start, a visual demo, relevant GitHub topics and a current release.
- `LC-02` Do not claim productivity gains, defect reduction, demand or broad compatibility without measured evidence.
- `LC-03` Stars are an awareness signal, not the product North Star; adoption, successful `doctor` runs and repeat usage matter more.
- `LC-04` No paid stars, reciprocal-star schemes, unsolicited bulk outreach or invented adopter logos.
- `LC-05` Publish case studies only after removing private repository details and separating observed evidence from author inference.

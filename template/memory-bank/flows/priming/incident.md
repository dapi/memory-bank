---
title: P1-INC Context Priming
doc_kind: governance
doc_function: canonical
purpose: Timeboxed route-specific праймеринг для Incident Flow.
derived_from:
  - context-priming.md
  - ../incident.md
canonical_for:
  - p1_incident_priming
  - incident_priming_input_classes
status: active
audience: humans_and_agents
---

# P1-INC Context Priming

Сразу после route прочитай only exact inputs из incident process manifest:
alert/report, affected runtime surfaces, доступные recovery signals, relevant
runbook и, если это не задерживает containment, recent deployment/change
evidence.

В incident record/timeline зафиксируй timestamps, observed impact, owner,
containment options и unknowns. Timebox discovery: containment и human incident
owner важнее полного понимания системы. Hypothesis не является root cause.

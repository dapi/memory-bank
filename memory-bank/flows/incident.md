---
title: Incident And PIR Flow
doc_kind: governance
doc_function: canonical
purpose: Operational flow от обнаружения и containment инцидента до PIR и prevention work.
derived_from:
  - ../dna/governance.md
  - routing.md
  - ../engineering/testing-policy.md
  - ../ops/runbooks/README.md
canonical_for:
  - incident_entry_contract
  - incident_response_flow
  - incident_human_gates
  - pir_requirements
  - incident_followup_routing
status: active
audience: humans_and_agents
---

# Incident And PIR Flow

Incident — событие с активным или потенциально серьёзным operational impact, где containment и восстановление важнее обычного delivery sequencing. Конкретные severity levels, роли и каналы адаптируются в `ops/` downstream-проекта.

## Flow

```text
detection → triage → containment → recovery → timeline
          → root cause analysis → remediation → PIR → prevention work
```

## Response Gates

- [ ] impact и affected surfaces зафиксированы
- [ ] назначены human incident owner и communication channel
- [ ] destructive или high-risk actions подтверждены согласно autonomy boundaries
- [ ] containment отделён от permanent fix
- [ ] recovery проверен наблюдаемыми signals

## Timeline And RCA

- Timeline отделяет timestamps и факты от интерпретаций.
- RCA отделяет подтверждённые causes, contributing factors и hypotheses.
- Blameless language не отменяет ясного ownership remediation.
- Не объявляй root cause без достаточного evidence; unresolved hypotheses остаются открытыми.

## PIR And Closure Gates

- [ ] impact, detection, response и recovery описаны
- [ ] root cause или границы текущего знания зафиксированы
- [ ] remediation проверена
- [ ] prevention items имеют owner, priority и отдельные task references
- [ ] релевантные runbooks, ops facts и архитектурные решения обновлены
- [ ] человек подтвердил RCA и приоритеты follow-up work

Каждый follow-up issue проходит новый [`Task Routing`](routing.md). Не скрывай feature, refactoring или bug fix внутри PIR action list без собственного route и evidence.

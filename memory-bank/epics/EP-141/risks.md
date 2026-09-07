---
title: "EP-141: Risks"
doc_kind: epic
doc_function: risk_register
purpose: "EP-141: Risks"
derived_from:
  - charter.md
  - roadmap.md
status: active
audience: humans_and_agents
---

# EP-141: Risks

| ID | Risk | Control | Owner | State |
| --- | --- | --- | --- | --- |
| ERISK-01 | Старый CLI установит новый payload без проверки | Bridge и minimum source capability gate; реальный binary fixture | CLI | open |
| ERISK-02 | Миграция ослабит прежние проверки | Явный opt-in, pinned legacy contracts, pass/fail fixtures | CLI | open |
| ERISK-03 | Исключение Flows оставит скрытые зависимости | Semantic/link/embedded frontmatter audit каждого состава | Template | open |
| ERISK-04 | Обновление уничтожит авторские документы | Ownership-aware plan, полный preflight, rollback и idempotence tests | CLI | open |
| ERISK-05 | Состав template и CLI разойдётся | Общие versioned fixtures и связанные PR | Both | open |

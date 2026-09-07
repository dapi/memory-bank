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
| ERISK-01 | Старый CLI установит новый payload без проверки | Bridge и minimum source capability gate; реальный binary fixture | CLI | controlled by verified fixtures |
| ERISK-02 | Миграция ослабит прежние проверки | Явный opt-in, pinned legacy contracts, pass/fail fixtures | CLI | controlled by verified fixtures |
| ERISK-03 | Исключение Flows оставит скрытые зависимости | Semantic/link/embedded frontmatter audit каждого состава | Template | controlled by verified fixtures |
| ERISK-04 | Обновление уничтожит авторские документы | Ownership-aware plan, полный preflight, rollback и idempotence tests | CLI | controlled by verified fixtures |
| ERISK-05 | Состав template и CLI разойдётся | Общие versioned fixtures и связанные PR | Both | controlled by verified fixtures |

Контроли проверены в [delivery evidence](../../../docs/component-delivery-evidence.md).
Порядок release остаётся ответственностью владельца связанных PR 63 → 64 → 143;
эта поставка не запускает live migration и не закрывает инициативу за человека.

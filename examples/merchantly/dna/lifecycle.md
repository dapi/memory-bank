---
doc_kind: governance
doc_function: canonical
purpose: Maintenance rules и sync checklist для governed-документов.
derived_from:
  - governance.md
status: active
---
# Document Lifecycle

Правила, обеспечивающие consistency governed-документации при изменениях.

## Maintenance Rules

1. **Upstream first.** Меняешь факт — сначала найди и обнови canonical owner.
2. **Downstream sync.** После изменения upstream проверь `derived_from`-зависимых.
3. **README sync.** Добавлен/удалён/переименован документ — обнови parent README.
4. **Конфликт = дефект.** Расхождение внутри authoritative set устраняется сразу.
5. **Conflict = report, not fix.** Агент, обнаруживший расхождение при чтении, фиксирует его как finding и сообщает человеку. Самостоятельное исправление — только если текущая задача явно требует изменения этого документа.

## Sync Checklist

Перед фиксацией изменений в governed-документации:

- [ ] frontmatter валиден, для `active` non-root задан `derived_from`
- [ ] для canonical `feature` задан `delivery_status`, для `adr` — `decision_status`
- [ ] parent `README.md` обновлён при изменении состава или reading order

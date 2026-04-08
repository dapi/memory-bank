---
title: Operations Index
doc_kind: engineering
doc_function: index
purpose: Навигация по операционной документации. Читать когда нужно работать с dev/stage/production окружением, релизами или платежами.
derived_from:
  - ../dna/governance.md
status: active
audience: humans_and_agents
---

# Operations Index

- [Development Environment](development.md) — dip, Docker, dev-сервер, browser testing, database schema. Читать при локальной разработке.
- [Production & Stage Access](production.md) — kubectl, psql, rails console, логи, Grafana, credentials, тестовые магазины. Читать при работе с production/stage.
- [Release & Deployment](release.md) — релизный процесс, Docker tagging, feature branches, тестовый план. Читать при выполнении релиза.
- [CloudPayments](payments.md) — админка, credentials, Playwright-вход, коды отклонения. Читать при работе с платежами.
- [Configuration](config.md) — полный справочник конфигурации (anyway_config, application.yml). Читать при работе с настройками.
- [Runbooks](runbooks/) — operational runbooks (k8s custom domains и т.д.). Читать при инцидентах и ops-задачах.

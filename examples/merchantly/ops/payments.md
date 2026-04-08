---
title: CloudPayments
doc_kind: engineering
doc_function: canonical
purpose: CloudPayments админка, credentials, Playwright-вход, коды отклонения. Читать при работе с платёжной системой.
derived_from:
  - ../dna/governance.md
status: active
audience: humans_and_agents
---

# CloudPayments (платёжная система)

Админка для просмотра транзакций, статистики и управления платежами.

**URL:** https://merchant.cloudpayments.ru/next/transactions/list

## Учётные данные

```bash
CLOUDPAYMENTS_WEB_LOGIN     # Email для входа
CLOUDPAYMENTS_WEB_PASSWORD  # Пароль
```

## Вход через Playwright

1. Открыть `https://merchant.cloudpayments.ru/next/transactions/list`
2. Редирект на `https://id.cloudpayments.ru/login` — ввести email из `CLOUDPAYMENTS_WEB_LOGIN`
3. Нажать "Продолжить" → ввести пароль из `CLOUDPAYMENTS_WEB_PASSWORD`
4. Нажать "Войти"
5. Если предложение включить 2FA — "Напомнить позже"
6. Редирект на страницу транзакций

## Полезные страницы

- Транзакции: `/next/transactions/list`
- Статистика: `/next/statistics`
- Главная: `/next/dashboard`
- Сайты: `/next/sites`

## Коды отклонения платежей

- `5051` — Недостаточно средств
- `5062` — Карта не поддерживает данный тип операции
- `0` — Успешно

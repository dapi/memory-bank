---
title: Local tgcli through voldar
doc_kind: ops
doc_function: runbook
purpose: Запуск локального tgcli через автоматический SOCKS5 SSH-туннель на voldar.
derived_from:
  - ../README.md
status: active
audience: humans_and_agents
---

# Local tgcli через voldar

## Summary

Локальный `tgcli` запускается на Mac и использует SOCKS5-транспорт:

```text
tgcli → 127.0.0.1:1080 → autossh/SSH → voldar.brandymint.ru → Telegram
```

Удалённый `tgcli` на `voldar` остановлен, чтобы не использовать Telegram-сессию
параллельно с локальным клиентом.

## Current Configuration

- VPS: `voldar.brandymint.ru`
- Egress IP: `54.89.47.199`
- Local proxy: `socks5://127.0.0.1:1080`
- macOS LaunchAgent: `com.brandymint.proxied-chrome`
- Local client source: `/Users/danil/code/tgcli`
- Canonical executable: `~/.local/bin/tgcli`
- Managed package: mise tool `npm:@dapi/tgcli@2.8.1`
- Local tgcli store: `~/Library/Application Support/tgcli`
- Remote tgcli store: `/home/danil/.local/share/tgcli`

`~/.local/bin/tgcli` is the only supported entrypoint for shells, cron jobs and
LaunchAgents. The wrapper resolves mise from `$HOME`, then restores the caller's
working directory before executing tgcli. This deliberately isolates the native
`better-sqlite3` binary from project-local Node overrides: running tgcli inside a
repository that selects another Node version must not change tgcli's Node ABI.
The wrapper therefore does not depend on a project's active Node.js version or a
source checkout.

To update the canonical installation, change the pinned
`"npm:@dapi/tgcli"` version in `~/dotfiles/.config/mise/config.toml`, then run:

```bash
cd ~/dotfiles
mise install npm:@dapi/tgcli
mise reshim
~/.local/bin/tgcli --version
```

Сессия локального `tgcli` была перенесена с `voldar`; реальные session/config
значения не записываются в Memory Bank.

## Trigger / Symptoms

Использовать этот runbook, если локальный `tgcli` не подключается к Telegram,
если Mac вышел из сна или если нужно проверить, что трафик идёт через `voldar`.

## Diagnosis

Из каталога infra:

```bash
cd ~/code/brandymint/infra
direnv exec . make vpn-status
direnv exec . curl --socks5-hostname 127.0.0.1:1080 https://api.ipify.org
```

Ожидаемый внешний IP: `54.89.47.199`.

Проверка Telegram:

```bash
cd ~/code/tgcli
direnv exec ~/code/brandymint/infra tgcli auth status --json --timeout 30s
direnv exec ~/code/brandymint/infra tgcli doctor --connect --json --timeout 30s
```

Ожидается `authenticated: true` и `connected: true`.

Проверка чтения live-сообщения:

```bash
cd ~/code/brandymint/infra
direnv exec . tgcli messages list \
  --chat @dapi_personal_bot \
  --limit 1 \
  --source live \
  --json \
  --timeout 30s
```

## Resolution

Однократно установить `autossh` и LaunchAgent:

```bash
cd ~/code/brandymint/infra
direnv exec . make vpn-install
direnv exec . make vpn-status
```

Настройка proxy в `tgcli`:

```bash
cd ~/code/tgcli
direnv exec ~/code/brandymint/infra tgcli config set proxy socks5://127.0.0.1:1080
```

После установки LaunchAgent запускает туннель при входе пользователя в macOS
и переподключает его после сетевого обрыва или сна.

## Rollback

Отключить proxy для локального `tgcli`:

```bash
cd ~/code/tgcli
direnv exec ~/code/brandymint/infra tgcli config unset proxy
```

Остановить автоматический туннель:

```bash
cd ~/code/brandymint/infra
direnv exec . make vpn-stop
```

Удаление LaunchAgent выполняется только при явной необходимости:

```bash
direnv exec . make vpn-uninstall
```

## Safety Notes

- Не коммитить `config.json`, `session.json`, UUID, ключи или proxy credentials.
- Не запускать одновременно удалённый и локальный `tgcli` на одной Telegram-сессии.
- Если требуется снова использовать `tgcli` на `voldar`, сначала остановить локальный
  background process и проверить, что session store не копируется во время записи.

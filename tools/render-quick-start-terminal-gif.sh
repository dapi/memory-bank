#!/usr/bin/env bash

set -euo pipefail

repository_root="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
output_path="${1:-${repository_root}/docs/assets/quick-start-terminal.gif}"
temporary_directory="$(mktemp -d)"

cleanup() {
  rm -rf -- "${temporary_directory}"
}
trap cleanup EXIT

mkdir -p "$(dirname "${output_path}")"

render_frame() {
  local frame_number="$1"
  local terminal_text="$2"

  magick \
    -size 1200x720 \
    xc:'#0d1117' \
    -fill '#161b22' \
    -stroke '#30363d' \
    -strokewidth 2 \
    -draw 'roundrectangle 28,28 1172,692 18,18' \
    -fill '#ff5f56' -stroke none -draw 'circle 62,62 69,62' \
    -fill '#ffbd2e' -draw 'circle 86,62 93,62' \
    -fill '#27c93f' -draw 'circle 110,62 117,62' \
    -font 'Menlo-Regular' \
    -pointsize 27 \
    -interline-spacing 10 \
    -fill '#e6edf3' \
    -gravity northwest \
    -annotate +66+105 "${terminal_text}" \
    "${temporary_directory}/frame-${frame_number}.png"
}

render_frame 1 '$ cd my-project'

render_frame 2 '$ cd my-project
$ codex

› Прочитай задачу и memory-bank/flows/routing.md'

render_frame 3 '$ cd my-project
$ codex

› Прочитай задачу и memory-bank/flows/routing.md

• Маршрут: Feature Flow'

render_frame 4 '$ cd my-project
$ codex

• Маршрут: Feature Flow
• Создан Feature Pack

  memory-bank/features/FT-042/
  ├── README.md
  └── brief.md'

render_frame 5 '$ cd my-project
$ codex

• Проблема и критерии проверки зафиксированы
• Связанные решения и сценарии найдены
• Следующий этап: проектирование решения'

render_frame 6 '$ cd my-project
$ codex

✓ Контекст сохранён в репозитории
✓ Новая сессия сможет продолжить задачу'

magick \
  -delay 90 "${temporary_directory}/frame-1.png" \
  -delay 110 "${temporary_directory}/frame-2.png" \
  -delay 120 "${temporary_directory}/frame-3.png" \
  -delay 150 "${temporary_directory}/frame-4.png" \
  -delay 150 "${temporary_directory}/frame-5.png" \
  -delay 220 "${temporary_directory}/frame-6.png" \
  -loop 0 \
  -layers Optimize \
  "${output_path}"

printf 'Rendered %s\n' "${output_path}"

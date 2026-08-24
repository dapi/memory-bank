#!/usr/bin/env bash

set -euo pipefail

repository_root="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
output_directory="${1:-${repository_root}/docs/assets}"
temporary_directory="$(mktemp -d)"

cleanup() {
  rm -rf -- "${temporary_directory}"
}
trap cleanup EXIT

mkdir -p "${output_directory}"

render_frame() {
  local scenario="$1"
  local frame_number="$2"
  local terminal_text="$3"

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
    "${temporary_directory}/${scenario}-${frame_number}.png"
}

render_frame routing 1 '$ codex'

render_frame routing 2 '$ codex

› Прочитай GitHub issue #123 и выполни его по процессу
  memory-bank/flows/routing.md'

render_frame routing 3 '$ codex

› Прочитай GitHub issue #123 и выполни его по процессу
  memory-bank/flows/routing.md

• Читаю GitHub issue #123…'

render_frame routing 4 '$ codex

› Прочитай GitHub issue #123 и выполни его по процессу
  memory-bank/flows/routing.md

• Читаю GitHub issue #123…
• Выполняю Task Routing…'

magick \
  -delay 90 "${temporary_directory}/routing-1.png" \
  -delay 140 "${temporary_directory}/routing-2.png" \
  -delay 130 "${temporary_directory}/routing-3.png" \
  -delay 220 "${temporary_directory}/routing-4.png" \
  -loop 0 \
  -layers Optimize \
  "${output_directory}/quick-start-routing.gif"

render_frame feature-pack 1 '$ codex'

render_frame feature-pack 2 '$ codex

› Прочитай GitHub issue #123.
  Создай по нему Feature Pack согласно процессу
  memory-bank/flows/feature.md.'

render_frame feature-pack 3 '$ codex

› Прочитай GitHub issue #123.
  Создай по нему Feature Pack согласно процессу
  memory-bank/flows/feature.md.

  Проведи ревью созданного комплекта документов.'

render_frame feature-pack 4 '$ codex

› Прочитай GitHub issue #123.
  Создай и проверь Feature Pack.
  После обязательных этапов приступай к реализации.

• Читаю issue и процесс Feature Flow…'

magick \
  -delay 90 "${temporary_directory}/feature-pack-1.png" \
  -delay 140 "${temporary_directory}/feature-pack-2.png" \
  -delay 150 "${temporary_directory}/feature-pack-3.png" \
  -delay 230 "${temporary_directory}/feature-pack-4.png" \
  -loop 0 \
  -layers Optimize \
  "${output_directory}/quick-start-feature-pack.gif"

render_frame start-issue 1 '$ start-issue 123'

render_frame start-issue 2 '$ start-issue 123

• Получаю GitHub issue #123…'

render_frame start-issue 3 '$ start-issue 123

• Получаю GitHub issue #123…
• Создаю ветку и директорию worktree…'

render_frame start-issue 4 '$ start-issue 123

• Получаю GitHub issue #123…
• Создаю ветку и директорию worktree…
• Запускаю настроенного агента…'

magick \
  -delay 100 "${temporary_directory}/start-issue-1.png" \
  -delay 130 "${temporary_directory}/start-issue-2.png" \
  -delay 150 "${temporary_directory}/start-issue-3.png" \
  -delay 230 "${temporary_directory}/start-issue-4.png" \
  -loop 0 \
  -layers Optimize \
  "${output_directory}/quick-start-start-issue.gif"

printf 'Rendered terminal GIFs in %s\n' "${output_directory}"

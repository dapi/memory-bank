#!/usr/bin/env bash
#
# Generic local-checkout bootstrap. Adapt this script in the downstream
# repository when the project has additional setup requirements.

set -euo pipefail

has_command() {
  command -v "$1" >/dev/null 2>&1
}

if [[ -f mise.toml || -f .mise.toml || -f .tool-versions ]]; then
  if ! has_command mise; then
    echo "mise is required by this checkout but is not installed." >&2
    exit 1
  fi

  mise trust
  mise install
fi

if [[ -f .gitmodules ]]; then
  git submodule update --init --recursive
fi

if [[ -f .envrc ]]; then
  if ! has_command direnv; then
    echo "direnv is required by .envrc but is not installed." >&2
    exit 1
  fi

  direnv allow
fi

cat <<'EOF'
Bootstrap prerequisites are ready.

Add the project's dependency-install, database, and service setup commands to
this script. Do not copy secrets or .env files from another checkout here.
EOF

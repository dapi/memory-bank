#!/usr/bin/env bash

set -euo pipefail

script_dir="$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" && pwd)"
symphony_home="${SYMPHONY_HOME:-"${script_dir}/../symphony"}"
symphony_repo="${SYMPHONY_REPOSITORY:-https://github.com/dapi/symphony.git}"
workspace_root="${script_dir}/.symphony-workspace"
gitignore_path="${script_dir}/.gitignore"
workspace_ignore=".symphony-workspace/"

normalize_repo_url() {
  local repo_url="$1"

  repo_url="${repo_url#git@github.com:}"
  repo_url="${repo_url#https://github.com/}"
  repo_url="${repo_url%.git}"
  printf '%s' "${repo_url%/}"
}

if ! command -v git >/dev/null 2>&1; then
  echo "git is required to install Symphony." >&2
  exit 1
fi

if ! command -v mise >/dev/null 2>&1; then
  echo "mise is required to install Symphony." >&2
  exit 1
fi

if [[ ! -f "${gitignore_path}" ]] || ! grep -Fqx "${workspace_ignore}" "${gitignore_path}"; then
  printf '%s\n' "${workspace_ignore}" >> "${gitignore_path}"
fi

mkdir -p "${workspace_root}"

if [[ ! -d "${symphony_home}/.git" ]]; then
  git clone "${symphony_repo}" "${symphony_home}"
else
  expected_repo="$(normalize_repo_url "${symphony_repo}")"
  actual_repo="$(git -C "${symphony_home}" remote get-url origin 2>/dev/null || true)"

  if [[ "$(normalize_repo_url "${actual_repo}")" != "${expected_repo}" ]]; then
    echo "Existing Symphony checkout has origin ${actual_repo:-<missing>}; expected ${symphony_repo}." >&2
    echo "Set SYMPHONY_HOME to a matching checkout or update its origin before bootstrapping." >&2
    exit 1
  fi
fi

if [[ ! -f "${symphony_home}/elixir/mix.exs" ]]; then
  echo "Symphony Elixir source was not found at ${symphony_home}/elixir." >&2
  exit 1
fi

cd "${symphony_home}/elixir"
mise trust
mise install
mise exec -- mix setup
mise exec -- mix build

echo "Symphony is ready. Start it with ./run-symphony.sh."

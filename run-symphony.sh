#!/usr/bin/env bash

set -euo pipefail

script_dir="$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" && pwd)"
symphony_home="${SYMPHONY_HOME:-"${script_dir}/../symphony"}"
symphony_dir="${symphony_home}/elixir"
workflow_path="${script_dir}/WORKFLOW.md"
workspace_root="${script_dir}/.symphony-workspace"

if ! command -v direnv >/dev/null 2>&1; then
  echo "direnv is required to load this repository's environment." >&2
  exit 1
fi

if ! command -v mise >/dev/null 2>&1; then
  echo "mise is required to run Symphony." >&2
  exit 1
fi

if [[ ! -f "${workflow_path}" ]]; then
  echo "WORKFLOW.md not found at ${workflow_path}." >&2
  exit 1
fi

mkdir -p "${workspace_root}"

if [[ ! -x "${symphony_dir}/bin/symphony" ]]; then
  echo "Symphony launcher not found at ${symphony_dir}/bin/symphony." >&2
  echo "Run ./bootstrap-symphony.sh first, or set SYMPHONY_HOME." >&2
  exit 1
fi

cd "${symphony_dir}"

exec direnv exec "${script_dir}" mise exec -- ./bin/symphony \
  --i-understand-that-this-will-be-running-without-the-usual-guardrails \
  "${workflow_path}"

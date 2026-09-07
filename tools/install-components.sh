#!/usr/bin/env bash
set -euo pipefail

if [[ $# -eq 0 || ( "$1" != init && "$1" != pull ) ]]; then
  printf 'Usage: %s init|pull [CLI options]\n' "$0" >&2
  exit 2
fi
operation="$1"
shift
for argument in "$@"; do
  case "$argument" in
    -source|-source=*|--source|--source=*|-source-ref|-source-ref=*|--source-ref|--source-ref=*|-template-version|-template-version=*|--template-version|--template-version=*)
      printf 'This entrypoint pins its own source checkout; %s cannot be overridden.\n' "$argument" >&2
      exit 2
      ;;
  esac
done
cli="${MEMORY_BANK_CLI:-memory-bank-cli}"
if ! "$cli" capabilities --require components/v1 --require adoption/v1 >/dev/null; then
  printf '%s\n' 'A component-capable memory-bank-cli is required; the installer was not invoked.' 'Upgrade the CLI first, or keep the legacy source f1f04de843aef45a2425d4a7351d577bbf89e940.' >&2
  exit 1
fi
script_directory="$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" && pwd -P)"
source_root="$(git -C "$script_directory" rev-parse --show-toplevel)"
source_ref="$(git -C "$source_root" rev-parse HEAD)"
exec "$cli" "$operation" --source "$source_root" --source-ref "$source_ref" --template-version "git:$source_ref" "$@"

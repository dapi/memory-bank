#!/usr/bin/env python3
"""Prove real incompatible CLIs cannot reach init/pull through the entrypoint."""
import argparse
import hashlib
import json
import os
from pathlib import Path
import subprocess
import tempfile


def main():
    parser = argparse.ArgumentParser()
    parser.add_argument("--pre-bridge", type=Path, required=True)
    parser.add_argument("--bridge", type=Path, required=True)
    args = parser.parse_args()
    entrypoint = Path(__file__).resolve().parent / "install-components.sh"
    for name, binary in (("pre-bridge", args.pre_bridge), ("bridge", args.bridge)):
        binary = binary.resolve(strict=True)
        with tempfile.TemporaryDirectory(prefix="memory-bank-entrypoint-") as scratch:
            root = Path(scratch)
            target = root / "target"
            target.mkdir()
            sentinel = target / "sentinel"
            sentinel.write_bytes(b"preserve\n")
            calls = root / "calls.jsonl"
            wrapper = root / "recording-cli"
            wrapper.write_text(
                "#!/usr/bin/env python3\n"
                "import json, os, subprocess, sys\n"
                "with open(os.environ['ENTRYPOINT_CALLS'], 'a') as f:\n"
                "    f.write(json.dumps(sys.argv[1:])+'\\n')\n"
                "sys.exit(subprocess.call([os.environ['ENTRYPOINT_BINARY'], *sys.argv[1:]]))\n"
            )
            wrapper.chmod(0o755)
            env = dict(os.environ, MEMORY_BANK_CLI=str(wrapper),
                       ENTRYPOINT_CALLS=str(calls), ENTRYPOINT_BINARY=str(binary))
            for operation in ("init", "pull"):
                calls.unlink(missing_ok=True)
                result = subprocess.run(
                    [str(entrypoint), operation, "--repo-root", str(target), "--preset", "docs"],
                    env=env, capture_output=True, text=True, check=False,
                )
                invoked = [json.loads(line) for line in calls.read_text().splitlines()]
                assert result.returncode != 0, (name, operation, "unexpected success")
                assert len(invoked) == 1 and invoked[0][0] == "capabilities", invoked
                assert sorted(p.name for p in target.iterdir()) == ["sentinel"]
                assert sentinel.read_bytes() == b"preserve\n"
            digest = hashlib.sha256(binary.read_bytes()).hexdigest()
            print(f"{name}: init/pull blocked before installer; binary sha256={digest}")


if __name__ == "__main__":
    main()

# CHK-02: Source-Specific Leakage Guard

Command:

```sh
rg -n "zelma|zellij|Codex|sessions list|setup --json" memory-bank/use-cases/README.md
```

Result: pass. The command returned exit code `1` with no matches, which is the
expected result for this negative check.

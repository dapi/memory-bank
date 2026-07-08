# CHK-01: README Guidance Coverage

Command:

```sh
rg -n "Operational|agentic|machine-readable|recovery|postconditions|SC-\\*" memory-bank/use-cases/README.md
```

Result: pass.

Relevant output:

```text
17:Use case нужен для сценария, который живет на уровне продукта, повторяется во времени и может быть upstream для нескольких feature packages. Это не замена `SC-*` внутри `brief.md`: `SC-*` описывают acceptance сценарии delivery-единицы, а `UC-*` описывают устойчивое поведение системы на уровне проекта.
25:- нужен канонический owner для trigger, preconditions, main flow и postconditions.
31:- его достаточно описать через `SC-*` в `brief.md`.
33:## Operational / Agentic Use Cases
35:Operational и agentic use cases описывают устойчивые рабочие сценарии, где
45:  postconditions;
47:- задает machine-readable contract, status или handoff format, который другие
49:- описывает recovery/postconditions, которые важны независимо от конкретной
52:Оставляй сценарий в feature-level `SC-*`, если он нужен только для приемки
56:Machine-readable status, structured diagnostics, handoff payloads, recovery
57:outcomes и postconditions допустимо описывать в `UC-*`, когда они являются
```

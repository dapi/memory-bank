# Feature Flow — Вариант 1: Линейный конвейер «Stage → Gate»

Акцент: последовательность стадий, ворота (gates) и три семантических пространства
(problem → solution → execution).

Опорные файлы: `flows/routing.md`, `flows/feature.md`,
`flows/priming/context-priming.md`, `engineering/validation-profiles.md`.

```mermaid
flowchart TD
    ISSUE([ISSUE / INTENT])

    subgraph ROUTING["TASK ROUTING · priming P0"]
        R["Предикаты по порядку → выбран Feature Flow<br/>+ validation profile<br/>(documentation / low-risk / standard / high-risk / release-deployment)"]
    end

    subgraph PROBLEM["PROBLEM SPACE"]
        B["brief.md · what & why<br/>REQ-* · SC-* · CHK-* · MET-*<br/>ASM-* · CON-* · DEC-*<br/>validation profile · verify contract"]
    end

    subgraph SOLUTION["SOLUTION SPACE · только если Design required: yes"]
        D["design.md · how to solve<br/>SOL-* · ALT-* · TRD-*<br/>C4-* · CTR-* · INV-* · FM-* · RB-*<br/>architecture coverage · risk verification"]
    end

    subgraph EXEC["EXECUTION SPACE"]
        P["implementation-plan.md · how to realize<br/>GRND-* · STEP-*<br/>priming manifest (точные пути)<br/>test-strategy table · immutable SHA"]
        E["EXECUTION LOOP<br/>commit → CI → review → PR → merge"]
    end

    DONE([DELIVERY · delivery_status = done])

    ISSUE --> R
    R --> B
    B -->|"GATE: Problem Ready<br/>brief.md active · Design required решён · профиль зафиксирован"| DGATE{Design<br/>required?}
    DGATE -->|yes| D
    D -->|"GATE: Solution Ready<br/>design pack active · C4 + coverage + risk verification закрыты"| P
    DGATE -->|no| P
    P -->|"GATE: Plan Ready<br/>plan active · grounding на immutable SHA · priming/test-strategy готовы"| PRE{HEAD ==<br/>grounded SHA?}
    PRE -->|no| STOP[STOP → re-ground]
    STOP --> P
    PRE -->|yes · Execution entry| E
    E -->|"GATE: Delivery Complete"| DONE
```

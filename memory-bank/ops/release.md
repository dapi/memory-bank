---
title: Release And Deployment
doc_kind: ops
doc_function: canonical
purpose: Canonical release process for the dapi/memory-bank template repository.
derived_from:
  - ../dna/governance.md
status: active
audience: humans_and_agents
---

# Release And Deployment

## Release Flow

Этот репозиторий публикует versioned snapshots шаблона, а не runtime-артефакты
или `memory-bank-cli` (CLI выпускается в отдельном репозитории). Релизом
считается Git tag и соответствующий GitHub Release.

1. Убедиться, что `main` содержит нужные изменения и CI на нём завершился
   успешно.
2. Выбрать SemVer-тег формата `vMAJOR.MINOR.PATCH`; для prerelease допустимы
   SemVer prerelease/build suffixes.
3. Создать и отправить annotated tag с `main`.
4. Workflow `Release` повторно проверяет template и project-local Memory Bank,
   подтверждает, что tag указывает на commit из `main`, и создаёт GitHub Release
   с автоматически сгенерированными notes.

Нельзя создавать релиз вручную через GitHub UI: это обходит validation gate.

## Release Commands

Создавай tag только из актуального `main`:

```bash
git switch main
git pull --ff-only origin main
git tag -a vX.Y.Z -m "Release vX.Y.Z"
git push origin vX.Y.Z
```

После push наблюдай workflow `Release`; GitHub Release появляется только после
его успешного завершения. Для работы workflow не требуются локальные секреты:
используется встроенный GitHub Actions token с `contents: write`.

## Release Verification

`Release` выполняет те же structural checks, что и CI: проверяет layout,
template lint, project-local lint и template doctor. Затем проверь, что GitHub
Release содержит ожидаемый tag и ссылку на release commit.

## Correction And Rollback

Git tags и опубликованные releases не переписываются. Если release содержит
ошибку, выпусти следующий patch release с исправлением. GitHub Release можно
пометить как pre-release или добавить explanatory notes, но не удаляй tag как
способ исправить уже опубликованную историю.

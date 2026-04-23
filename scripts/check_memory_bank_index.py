#!/usr/bin/env python3
"""Audit markdown navigation integrity for the generic memory-bank template.

The script intentionally focuses on `memory-bank/` and avoids downstream
assumptions from example repositories.
"""

from __future__ import annotations

import os
import posixpath
import re
import sys
from collections import defaultdict, deque
from pathlib import Path


REPO_ROOT = Path(__file__).resolve().parents[1]
IGNORED_DIRS = {".git", ".hg", ".svn", ".venv", "node_modules", "tmp", "log", "vendor"}

SCOPE_PREFIX = "memory-bank/"
START_FILES = ("memory-bank/README.md",)
EXPECTED_INDEXES = (
    "memory-bank/README.md",
    "memory-bank/dna/README.md",
    "memory-bank/domain/README.md",
    "memory-bank/engineering/README.md",
    "memory-bank/flows/README.md",
    "memory-bank/flows/templates/README.md",
    "memory-bank/ops/README.md",
    "memory-bank/ops/runbooks/README.md",
    "memory-bank/prd/README.md",
    "memory-bank/use-cases/README.md",
    "memory-bank/adr/README.md",
    "memory-bank/features/README.md",
)

FENCED_CODE_BLOCK_RE = re.compile(r"```.*?```", re.DOTALL)
MARKDOWN_LINK_RE = re.compile(r"(?<!!)\[([^\]]+)\]\(([^)]+)\)")
FRONTMATTER_RE = re.compile(r"\A---\n(.*?)\n---(?:\n|$)", re.DOTALL)
BULLET_LINK_RE = re.compile(r"^\s*-\s+.*?(?<!\!)\[([^\]]+)\]\(([^)]+)\)")


def discover_markdown_files(repo_root: Path) -> dict[str, Path]:
    files: dict[str, Path] = {}
    for root, dirs, filenames in os.walk(repo_root):
        dirs[:] = [directory for directory in dirs if directory not in IGNORED_DIRS]
        for filename in filenames:
            if not filename.endswith(".md"):
                continue
            full_path = Path(root, filename)
            relative_path = full_path.relative_to(repo_root).as_posix()
            files[relative_path] = full_path
    return files


def read_text(path: Path) -> str:
    return path.read_text(encoding="utf-8", errors="ignore")


def strip_fenced_code_blocks(text: str) -> str:
    return FENCED_CODE_BLOCK_RE.sub("", text)


def normalize_internal_markdown_target(source_path: str, raw_url: str) -> str | None:
    url = raw_url.strip().strip("<>")
    if not url or url.startswith(("http://", "https://", "mailto:", "#")):
        return None

    url = url.split("#", 1)[0].split("?", 1)[0].strip()
    if not url:
        return None

    extension = posixpath.splitext(url)[1].lower()
    if extension and extension != ".md":
        return None

    base_dir = posixpath.dirname(source_path)
    if url.startswith("/"):
        resolved = posixpath.normpath(url.lstrip("/"))
    else:
        resolved = posixpath.normpath(posixpath.join(base_dir, url))

    if not extension:
        resolved = posixpath.join(resolved, "README.md")

    return resolved


def extract_internal_markdown_links(source_path: str, text: str) -> list[tuple[str, str]]:
    stripped_text = strip_fenced_code_blocks(text)
    links: list[tuple[str, str]] = []
    for match in MARKDOWN_LINK_RE.finditer(stripped_text):
        label = match.group(1).strip()
        target = normalize_internal_markdown_target(source_path, match.group(2))
        if target is None:
            continue
        links.append((target, label))
    return links


def build_link_graph(markdown_files: dict[str, Path]) -> tuple[dict[str, set[str]], dict[str, set[str]]]:
    graph: dict[str, set[str]] = defaultdict(set)
    broken_links: dict[str, set[str]] = defaultdict(set)
    known_paths = set(markdown_files)

    for source_path, full_path in markdown_files.items():
        text = read_text(full_path)
        for target, _label in extract_internal_markdown_links(source_path, text):
            if target in known_paths:
                graph[source_path].add(target)
            elif source_path.startswith(SCOPE_PREFIX):
                broken_links[source_path].add(target)

    return graph, broken_links


def bfs_reachable(graph: dict[str, set[str]], start_files: tuple[str, ...]) -> set[str]:
    visited: set[str] = set()
    queue: deque[str] = deque()

    for start_file in start_files:
        visited.add(start_file)
        queue.append(start_file)

    while queue:
        current = queue.popleft()
        for target in sorted(graph.get(current, set())):
            if target in visited:
                continue
            visited.add(target)
            queue.append(target)

    return visited


def parse_frontmatter(text: str) -> dict[str, str]:
    match = FRONTMATTER_RE.match(text)
    if not match:
        return {}

    frontmatter: dict[str, str] = {}
    for line in match.group(1).splitlines():
        if not line or line.startswith((" ", "\t", "-")) or ":" not in line:
            continue
        key, value = line.split(":", 1)
        frontmatter[key.strip()] = value.strip().strip("\"'")
    return frontmatter


def annotation_text_for_child_links(index_path: str, text: str) -> list[tuple[str, str]]:
    section_prefix = posixpath.dirname(index_path)
    stripped_lines = strip_fenced_code_blocks(text).splitlines()
    annotations: list[tuple[str, str]] = []

    index_line = 0
    while index_line < len(stripped_lines):
        line = stripped_lines[index_line]
        match = BULLET_LINK_RE.match(line)
        if not match:
            index_line += 1
            continue

        target = normalize_internal_markdown_target(index_path, match.group(2))
        if target is None:
            index_line += 1
            continue

        if section_prefix:
            child_prefix = f"{section_prefix}/"
            if not target.startswith(child_prefix):
                index_line += 1
                continue
        elif "/" in target:
            index_line += 1
            continue

        fragments: list[str] = []
        inline_annotation = MARKDOWN_LINK_RE.sub("", line).strip(" -\t:")
        if inline_annotation:
            fragments.append(inline_annotation)

        continuation_index = index_line + 1
        while continuation_index < len(stripped_lines):
            continuation = stripped_lines[continuation_index]
            if not continuation.strip():
                break
            if continuation.startswith(("  ", "\t")):
                fragments.append(continuation.strip())
                continuation_index += 1
                continue
            break

        annotations.append((target, " ".join(fragments).strip()))
        index_line += 1

    return annotations


def validate_expected_index(index_path: str, markdown_files: dict[str, Path]) -> list[str]:
    issues: list[str] = []
    full_path = markdown_files.get(index_path)
    if full_path is None:
        return ["missing expected index file"]

    text = read_text(full_path)
    frontmatter = parse_frontmatter(text)
    if not frontmatter:
        issues.append("missing YAML frontmatter")
    if frontmatter.get("purpose", "") == "":
        issues.append("missing 'purpose' in frontmatter")
    if frontmatter.get("doc_function") != "index":
        issues.append("expected `doc_function: index`")

    for target, annotation in annotation_text_for_child_links(index_path, text):
        normalized_annotation = re.sub(r"\s+", " ", annotation).strip(" -:\t").lower()
        basename = posixpath.basename(target).lower()
        basename_without_extension = posixpath.splitext(basename)[0]
        if not normalized_annotation:
            issues.append(f"missing annotation for child link -> {target}")
            continue
        if normalized_annotation in {basename, basename_without_extension}:
            issues.append(f"annotation repeats filename for child link -> {target}")
            continue
        if len(normalized_annotation) < 12:
            issues.append(f"annotation too short for child link -> {target}")

    return issues


def main() -> int:
    markdown_files = discover_markdown_files(REPO_ROOT)
    scoped_markdown_paths = sorted(
        path for path in markdown_files if path.startswith(SCOPE_PREFIX) and path.endswith(".md")
    )

    missing_start_files = [start_file for start_file in START_FILES if start_file not in markdown_files]
    if missing_start_files:
        print("Navigation audit failed: missing start files.")
        for start_file in missing_start_files:
            print(f"  - {start_file}")
        return 1

    graph, broken_links = build_link_graph(markdown_files)
    reachable = bfs_reachable(graph, START_FILES)
    unreachable = sorted(path for path in scoped_markdown_paths if path not in reachable)

    exit_code = 0

    print("Memory-bank navigation audit")
    print(f"Repo root: {REPO_ROOT}")
    print(f"Scope: {SCOPE_PREFIX}")
    print(f"Start files: {', '.join(START_FILES)}")
    print(f"Markdown files in scope: {len(scoped_markdown_paths)}")
    print()

    scoped_broken_links = {
        source_path: sorted(targets)
        for source_path, targets in broken_links.items()
        if source_path.startswith(SCOPE_PREFIX)
    }
    if scoped_broken_links:
        exit_code = 1
        print("Broken internal markdown links:")
        for source_path in sorted(scoped_broken_links):
            for target in scoped_broken_links[source_path]:
                print(f"  - {source_path} -> {target}")
        print()
    else:
        print("OK: no broken internal markdown links in scope.")
        print()

    if unreachable:
        exit_code = 1
        print("Unreachable markdown files in scope:")
        for path in unreachable:
            print(f"  - {path}")
        print()
    else:
        print("OK: all scoped markdown files are reachable from the configured start files.")
        print()

    print("Index compliance:")
    for index_path in EXPECTED_INDEXES:
        issues = validate_expected_index(index_path, markdown_files)
        if issues:
            exit_code = 1
            print(f"  - {index_path}")
            for issue in issues:
                print(f"    * {issue}")
            continue
        print(f"  - {index_path}: OK")

    return exit_code


if __name__ == "__main__":
    sys.exit(main())

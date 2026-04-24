#!/usr/bin/env python3
"""Audit markdown navigation integrity for a memory-bank-like documentation tree."""

from __future__ import annotations

import argparse
import json
import os
import posixpath
import re
import sys
from collections import defaultdict, deque
from pathlib import Path


IGNORED_DIRS = {".git", ".hg", ".svn", ".venv", "node_modules", "tmp", "log", "vendor"}

DEFAULT_SCOPE_ROOT = "memory-bank"
DEFAULT_MAX_DEPTH = 3

FENCED_CODE_BLOCK_RE = re.compile(r"```.*?```", re.DOTALL)
MARKDOWN_LINK_RE = re.compile(r"(?<!!)\[([^\]]+)\]\(([^)]+)\)")
FRONTMATTER_RE = re.compile(r"\A---\n(.*?)\n---(?:\n|$)", re.DOTALL)
BULLET_LINK_RE = re.compile(r"^\s*-\s+.*?(?<!\!)\[([^\]]+)\]\(([^)]+)\)")


def parse_args() -> argparse.Namespace:
    parser = argparse.ArgumentParser(
        description="Audit markdown navigation integrity for a memory-bank-like documentation tree."
    )
    parser.add_argument(
        "--repo-root",
        help="Filesystem path to the repository root. Defaults to the current directory when it contains the scope root.",
    )
    parser.add_argument(
        "--scope-root",
        default=DEFAULT_SCOPE_ROOT,
        help="Repository-relative directory to audit. Default: %(default)s",
    )
    parser.add_argument(
        "--entrypoint",
        action="append",
        default=[],
        help=(
            "Markdown document to use as a navigation entrypoint. Accepts repo-relative paths "
            "and scope-relative paths. Repeat the option to pass several files."
        ),
    )
    parser.add_argument(
        "--max-depth",
        type=int,
        default=DEFAULT_MAX_DEPTH,
        help="Maximum allowed navigation depth in link hops before the document becomes a warning. Default: %(default)s",
    )
    parser.add_argument(
        "--json",
        action="store_true",
        help="Emit a machine-readable JSON report instead of text.",
    )
    args = parser.parse_args()
    if args.max_depth < 0:
        parser.error("--max-depth must be greater than or equal to 0")
    return args


def normalize_scope_root(scope_root: str) -> str:
    normalized = posixpath.normpath(scope_root.strip())
    if normalized in {"", "."}:
        raise ValueError("--scope-root must point to a repository-relative directory")
    return normalized.rstrip("/")


def resolve_repo_root(repo_root_arg: str | None, scope_root: str) -> Path:
    if repo_root_arg:
        return Path(repo_root_arg).resolve()

    candidates = [Path.cwd().resolve()]
    if "__file__" in globals():
        script_path = Path(__file__)
        if str(script_path) not in {"", "<stdin>"}:
            candidates.append(script_path.resolve().parents[1])

    for candidate in candidates:
        if (candidate / scope_root).exists():
            return candidate

    return candidates[0]


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


def normalize_cli_document_path(raw_path: str) -> str | None:
    path = raw_path.strip().strip("<>")
    if not path:
        return None

    path = path.split("#", 1)[0].split("?", 1)[0].strip()
    if not path:
        return None

    extension = posixpath.splitext(path)[1].lower()
    if extension and extension != ".md":
        return None

    normalized = posixpath.normpath(path.lstrip("/"))
    if normalized in {"", "."}:
        return None

    if not extension:
        normalized = posixpath.join(normalized, "README.md")

    return normalized


def extract_internal_markdown_links(source_path: str, text: str) -> list[str]:
    stripped_text = strip_fenced_code_blocks(text)
    links: list[str] = []
    for match in MARKDOWN_LINK_RE.finditer(stripped_text):
        target = normalize_internal_markdown_target(source_path, match.group(2))
        if target is None:
            continue
        links.append(target)
    return links


def load_documents(repo_root: Path) -> dict[str, dict[str, object]]:
    documents: dict[str, dict[str, object]] = {}
    for relative_path, full_path in discover_markdown_files(repo_root).items():
        text = read_text(full_path)
        documents[relative_path] = {
            "full_path": full_path,
            "text": text,
            "frontmatter": parse_frontmatter(text),
        }
    return documents


def is_scoped_markdown(path: str, scope_root: str) -> bool:
    return path.startswith(f"{scope_root}/") and path.endswith(".md")


def is_scoped_readme(path: str, scope_root: str) -> bool:
    return is_scoped_markdown(path, scope_root) and posixpath.basename(path) == "README.md"


def resolve_entrypoint_path(entrypoint: str, scope_root: str, known_paths: set[str]) -> tuple[str | None, str]:
    primary_candidate = normalize_cli_document_path(entrypoint)
    if primary_candidate and primary_candidate in known_paths:
        return primary_candidate, entrypoint

    scoped_input = posixpath.join(scope_root, entrypoint.lstrip("/"))
    scoped_candidate = normalize_cli_document_path(scoped_input)
    if scoped_candidate and scoped_candidate in known_paths:
        return scoped_candidate, entrypoint

    fallback = primary_candidate or scoped_candidate or entrypoint
    return None, fallback


def derive_entrypoints(
    documents: dict[str, dict[str, object]],
    scope_root: str,
    configured_entrypoints: list[str],
) -> tuple[list[str], list[str]]:
    known_paths = set(documents)
    resolved: list[str] = []
    missing: list[str] = []

    if configured_entrypoints:
        for entrypoint in configured_entrypoints:
            resolved_path, missing_hint = resolve_entrypoint_path(entrypoint, scope_root, known_paths)
            if resolved_path is None:
                missing.append(missing_hint)
                continue
            if resolved_path not in resolved:
                resolved.append(resolved_path)
        return resolved, missing

    default_entrypoint = f"{scope_root}/README.md"
    if default_entrypoint in known_paths:
        return [default_entrypoint], []

    return [], [default_entrypoint]


def build_link_graph(
    documents: dict[str, dict[str, object]],
    scope_root: str,
) -> tuple[dict[str, set[str]], dict[str, set[str]], dict[str, set[str]]]:
    outgoing: dict[str, set[str]] = defaultdict(set)
    incoming_in_scope: dict[str, set[str]] = defaultdict(set)
    broken_links: dict[str, set[str]] = defaultdict(set)
    known_paths = set(documents)

    for source_path, document in documents.items():
        text = document["text"]
        assert isinstance(text, str)
        for target in extract_internal_markdown_links(source_path, text):
            if target in known_paths:
                outgoing[source_path].add(target)
                if (
                    source_path != target
                    and is_scoped_markdown(source_path, scope_root)
                    and is_scoped_markdown(target, scope_root)
                ):
                    incoming_in_scope[target].add(source_path)
            elif is_scoped_markdown(source_path, scope_root):
                broken_links[source_path].add(target)

    return outgoing, incoming_in_scope, broken_links


def derive_index_paths(documents: dict[str, dict[str, object]], scope_root: str) -> list[str]:
    index_paths: list[str] = []
    for path, document in documents.items():
        if not is_scoped_markdown(path, scope_root):
            continue
        frontmatter = document["frontmatter"]
        assert isinstance(frontmatter, dict)
        if frontmatter.get("doc_function") == "index":
            index_paths.append(path)
    return sorted(index_paths)


def derive_expected_readme_indices(documents: dict[str, dict[str, object]], scope_root: str) -> list[str]:
    readmes: list[str] = []
    for path, document in documents.items():
        if not is_scoped_readme(path, scope_root):
            continue
        frontmatter = document["frontmatter"]
        assert isinstance(frontmatter, dict)
        if frontmatter.get("doc_function") == "template":
            continue
        readmes.append(path)
    return sorted(readmes)


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

        child_prefix = f"{section_prefix}/"
        if section_prefix and not target.startswith(child_prefix):
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


def validate_index_document(index_path: str, documents: dict[str, dict[str, object]]) -> list[str]:
    document = documents.get(index_path)
    if document is None:
        return ["missing expected index file"]

    text = document["text"]
    frontmatter = document["frontmatter"]
    assert isinstance(text, str)
    assert isinstance(frontmatter, dict)

    issues: list[str] = []
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


def expected_parent_index(path: str, index_paths: set[str], scope_root: str) -> str | None:
    if not is_scoped_markdown(path, scope_root):
        return None

    current_dir = posixpath.dirname(path)
    if posixpath.basename(path) == "README.md":
        current_dir = posixpath.dirname(current_dir)

    while current_dir and current_dir != ".":
        candidate = posixpath.join(current_dir, "README.md")
        if candidate in index_paths and candidate != path:
            return candidate
        parent_dir = posixpath.dirname(current_dir)
        if parent_dir == current_dir:
            break
        current_dir = parent_dir

    return None


def build_navigation_reachability(
    outgoing: dict[str, set[str]],
    navigation_nodes: set[str],
    entrypoints: list[str],
) -> dict[str, dict[str, object]]:
    reachable: dict[str, dict[str, object]] = {}
    navigation_depths: dict[str, int] = {}
    queue: deque[str] = deque()

    for entrypoint in entrypoints:
        reachable[entrypoint] = {"depth": 0, "route": [entrypoint]}
        navigation_depths[entrypoint] = 0
        queue.append(entrypoint)

    while queue:
        current = queue.popleft()
        current_info = reachable[current]
        current_depth = navigation_depths[current]
        current_route = current_info["route"]
        assert isinstance(current_route, list)

        for target in sorted(outgoing.get(current, set())):
            candidate_depth = current_depth + 1
            candidate_route = current_route + [target]
            best = reachable.get(target)
            if best is None or candidate_depth < best["depth"]:
                reachable[target] = {"depth": candidate_depth, "route": candidate_route}
            if target in navigation_nodes:
                best_depth = navigation_depths.get(target)
                if best_depth is None or candidate_depth < best_depth:
                    navigation_depths[target] = candidate_depth
                    queue.append(target)

    return reachable


def flatten_broken_links(broken_links: dict[str, set[str]]) -> list[dict[str, str]]:
    items: list[dict[str, str]] = []
    for source_path in sorted(broken_links):
        for target in sorted(broken_links[source_path]):
            items.append({"source": source_path, "target": target})
    return items


def build_report(
    repo_root: Path,
    scope_root: str,
    entrypoints: list[str],
    missing_entrypoints: list[str],
    max_depth: int,
    documents: dict[str, dict[str, object]],
) -> dict[str, object]:
    scoped_markdown_paths = sorted(path for path in documents if is_scoped_markdown(path, scope_root))
    index_paths = derive_index_paths(documents, scope_root)
    expected_readme_indices = derive_expected_readme_indices(documents, scope_root)
    outgoing, incoming_in_scope, broken_links = build_link_graph(documents, scope_root)

    report: dict[str, object] = {
        "format_version": 1,
        "repo_root": str(repo_root),
        "scope_root": scope_root,
        "entrypoints": entrypoints,
        "missing_entrypoints": missing_entrypoints,
        "max_depth": max_depth,
        "stats": {
            "markdown_files_in_scope": len(scoped_markdown_paths),
            "index_documents_in_scope": len(index_paths),
        },
        "errors": {
            "config": [],
            "broken_links": flatten_broken_links(broken_links),
            "orphans": [],
            "unreachable": [],
            "index_contract": [],
        },
        "warnings": {
            "deep_reachable": [],
        },
    }

    config_errors = report["errors"]["config"]
    assert isinstance(config_errors, list)
    if missing_entrypoints:
        config_errors.append(
            {
                "message": "Configured entrypoints are missing.",
                "paths": missing_entrypoints,
            }
        )
    if not scoped_markdown_paths:
        config_errors.append(
            {
                "message": "Scope contains no markdown files.",
                "paths": [scope_root],
            }
        )
    if not entrypoints:
        config_errors.append(
            {
                "message": "No valid entrypoints were resolved.",
                "paths": missing_entrypoints or [f"{scope_root}/README.md"],
            }
        )

    index_paths_set = set(index_paths)
    entrypoint_set = set(entrypoints)

    if entrypoints and scoped_markdown_paths:
        navigation_nodes = set(index_paths) | entrypoint_set
        reachable = build_navigation_reachability(outgoing, navigation_nodes, entrypoints)

        for path in scoped_markdown_paths:
            inbound_sources = sorted(incoming_in_scope.get(path, set()))
            parent_index = expected_parent_index(path, index_paths_set, scope_root)
            if path not in entrypoint_set and not inbound_sources:
                report["errors"]["orphans"].append(
                    {
                        "path": path,
                        "expected_parent_index": parent_index,
                        "inbound_links": inbound_sources,
                    }
                )

            reachability = reachable.get(path)
            if reachability is None:
                report["errors"]["unreachable"].append(
                    {
                        "path": path,
                        "expected_parent_index": parent_index,
                        "inbound_links": inbound_sources,
                    }
                )
                continue

            depth = reachability["depth"]
            route = reachability["route"]
            assert isinstance(depth, int)
            assert isinstance(route, list)
            if depth > max_depth:
                report["warnings"]["deep_reachable"].append(
                    {
                        "path": path,
                        "depth": depth,
                        "max_depth": max_depth,
                        "expected_parent_index": parent_index,
                        "route": route,
                    }
                )

    for index_path in expected_readme_indices:
        issues = validate_index_document(index_path, documents)
        if not issues:
            continue
        report["errors"]["index_contract"].append(
            {
                "path": index_path,
                "issues": issues,
                "expected_parent_index": expected_parent_index(index_path, index_paths_set, scope_root),
            }
        )

    warnings = report["warnings"]["deep_reachable"]
    assert isinstance(warnings, list)
    warnings.sort(key=lambda item: (item["depth"], item["path"]))

    stats = report["stats"]
    assert isinstance(stats, dict)
    errors = report["errors"]
    assert isinstance(errors, dict)
    stats.update(
        {
            "broken_link_count": len(errors["broken_links"]),
            "orphan_count": len(errors["orphans"]),
            "unreachable_count": len(errors["unreachable"]),
            "index_contract_issue_count": len(errors["index_contract"]),
            "deep_reachable_warning_count": len(warnings),
            "entrypoint_count": len(entrypoints),
        }
    )

    has_errors = any(bool(errors[key]) for key in ("config", "broken_links", "orphans", "unreachable", "index_contract"))
    report["exit_code"] = 1 if has_errors else 0
    return report


def format_route(route: list[str]) -> str:
    return " -> ".join(route)


def print_text_report(report: dict[str, object]) -> None:
    print("Memory-bank link audit")
    print(f"Repo root: {report['repo_root']}")
    print(f"Scope root: {report['scope_root']}")
    print(f"Entrypoints: {', '.join(report['entrypoints']) or '(none)'}")
    print(f"Navigation depth threshold: {report['max_depth']}")

    stats = report["stats"]
    errors = report["errors"]
    warnings = report["warnings"]
    assert isinstance(stats, dict)
    assert isinstance(errors, dict)
    assert isinstance(warnings, dict)

    print(f"Markdown files in scope: {stats['markdown_files_in_scope']}")
    print(f"Index documents in scope: {stats['index_documents_in_scope']}")
    print()

    config_errors = errors["config"]
    assert isinstance(config_errors, list)
    if config_errors:
        print("Configuration errors:")
        for item in config_errors:
            print(f"  - {item['message']}")
            for path in item["paths"]:
                print(f"    * {path}")
        print()

    broken_links = errors["broken_links"]
    assert isinstance(broken_links, list)
    if broken_links:
        print("Broken internal markdown links:")
        for item in broken_links:
            print(f"  - {item['source']} -> {item['target']}")
        print()
    else:
        print("OK: no broken internal markdown links in scope.")
        print()

    orphans = errors["orphans"]
    assert isinstance(orphans, list)
    if orphans:
        print("Orphan markdown files in scope:")
        for item in orphans:
            print(f"  - {item['path']}")
            print(f"    expected_parent_index: {item['expected_parent_index'] or '(none)'}")
        print()
    else:
        print("OK: no orphan markdown files in scope.")
        print()

    unreachable = errors["unreachable"]
    assert isinstance(unreachable, list)
    if unreachable:
        print("Markdown files missing from index navigation:")
        for item in unreachable:
            print(f"  - {item['path']}")
            print(f"    expected_parent_index: {item['expected_parent_index'] or '(none)'}")
            inbound_links = item["inbound_links"]
            if inbound_links:
                print(f"    inbound_links: {', '.join(inbound_links)}")
        print()
    else:
        print("OK: all scoped markdown files are reachable from the configured entrypoints via index navigation.")
        print()

    deep_reachable = warnings["deep_reachable"]
    assert isinstance(deep_reachable, list)
    if deep_reachable:
        print("Warnings: documents reachable only deeper than the configured threshold:")
        for item in deep_reachable:
            print(f"  - {item['path']}")
            print(f"    depth: {item['depth']}")
            print(f"    expected_parent_index: {item['expected_parent_index'] or '(none)'}")
            print(f"    route: {format_route(item['route'])}")
        print()
    else:
        print("OK: no documents are reachable only deeper than the configured threshold.")
        print()

    index_contract = errors["index_contract"]
    assert isinstance(index_contract, list)
    print("Index compliance:")
    if index_contract:
        for item in index_contract:
            print(f"  - {item['path']}")
            for issue in item["issues"]:
                print(f"    * {issue}")
        print()
    else:
        print("  - OK")
        print()

    exit_code = report["exit_code"]
    assert isinstance(exit_code, int)
    result = "FAIL" if exit_code else "OK"
    print(f"Result: {result}")
    print("Machine-readable output: re-run with --json for a structured report suitable for auto-indexing.")


def main() -> int:
    args = parse_args()

    try:
        scope_root = normalize_scope_root(args.scope_root)
    except ValueError as error:
        print(str(error), file=sys.stderr)
        return 1

    repo_root = resolve_repo_root(args.repo_root, scope_root)
    documents = load_documents(repo_root)
    entrypoints, missing_entrypoints = derive_entrypoints(documents, scope_root, args.entrypoint)
    report = build_report(
        repo_root=repo_root,
        scope_root=scope_root,
        entrypoints=entrypoints,
        missing_entrypoints=missing_entrypoints,
        max_depth=args.max_depth,
        documents=documents,
    )

    if args.json:
        print(json.dumps(report, ensure_ascii=False, indent=2))
    else:
        print_text_report(report)

    exit_code = report["exit_code"]
    assert isinstance(exit_code, int)
    return exit_code


if __name__ == "__main__":
    sys.exit(main())

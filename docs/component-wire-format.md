# CTR-01 wire format, version 1

This is the sole normative CTR-01 behavior and serialization contract. [The overview](components.md)
provides navigation only. The template owns these formats; CLI Go types and producer/consumer
fixtures implement them. CLI pull updates a template; CLI update still updates the executable.
All objects reject unknown or duplicate fields, incompatible JSON types and trailing data.
JSON is UTF-8. Digests are `sha256:` plus 64 lowercase hexadecimal digits over exact file
bytes. Generated registry state uses the compact canonical JSON defined below and one final LF;
its integrity digest covers those exact bytes, not a reserialized approximation. Existing
lock formatting remains compatible with its ownership schema. Arrays specified as sets are sorted, unique
strings. Every specified sort uses ascending unsigned raw UTF-8 byte lexicographic order,
including IDs, paths, evidence, object keys and comparison items; no locale or case folding
participates in ordering. Omitted optional arrays/maps mean empty, never a wildcard. Names and paths are
case-sensitive. Paths are normalized repository-relative slash paths: no empty segments,
absolute paths, dot/dot-dot segments, backslashes, NUL, symlinks or Git metadata components. Metadata rejection is case-insensitive (including .GIT).
Each segment also rejects ASCII control characters, < > : " | ? *, and a trailing dot or
space. This excludes NTFS alternate streams and Win32 trimming aliases on every platform.
The case-insensitive stem before the first dot, after trimming trailing dots/spaces, must
not be CON, PRN, AUX, NUL, CONIN$, CONOUT$, COM1–COM9 or LPT1–LPT9 (including COM/LPT
variants using superscript ¹, ² or ³). These conservative restrictions apply independently
of the host filesystem; filenames are never silently normalized or renamed.
For each existing path segment, compare the requested name with the actual directory entry
bytes and reject a different spelling that resolves to that entry. Before writes, reject
collisions among all proposed destinations using the exact key algorithm
NFC(Default_Full_Case_Fold(NFC(path))), with Unicode 15.0 tables and locale-independent
default folding (not Turkic folding); also reject distinct paths/destinations that resolve
to the same file identity. On supported Linux/macOS hosts, identity is exactly the
(st_dev, st_ino) tuple obtained from lstat or fstat of an opened no-follow handle, never a
symlink target or normalized pathname. Distinct paths sharing that tuple, including hard
links, conflict. The comparison domain is repository payload/document paths, excluding the
transaction's private staging. To detect aliases beyond enumerated directories, reject every
original regular-file input with st_nlink > 1 at preflight and immediately before accepting
it for mutation. Thus a hard link in another directory or outside the repository rejects
without an unbounded filesystem scan. Private staged-replacement links created by the writer
are not original inputs; retained staging is verified/cleaned before ordinary preflight.
Re-read identities and repeat comparisons before mutation using pinned
handles; a changed observed identity rejects. Updating the exact target path is not a collision with itself.
This conservative portable-path rule applies on every platform, not only case-insensitive
filesystems. Compare these portable keys with every existing entry in each affected directory,
not only other proposed writes: Foo.md blocks creating foo.md even on case-sensitive storage.
Exact casing checks and safe handles are rechecked before mutation.
Conformance vectors (input → key): Foo.md → foo.md; e + U+0301 + .md → é.md;
É.md → é.md; Straße.md → strasse.md; STRASSE.md → strasse.md; K.md → k.md;
Σ.md and ς.md → σ.md. Apply the same algorithm to every path prefix, so differently
spelled directory names also conflict. The prototype verifies these Unicode 15.0 cases.

## Source envelope and inventory

The envelope is the checkout-root memory-bank-source.json Git blob. The component marker
and component manifest are the same file, installed as memory-bank/components.json and
included under that exact key in the exhaustive inventory. Its checkout path is
template/memory-bank/components.json for the canonical template/ payload root,
memory-bank-template/memory-bank/components.json for the supported legacy payload-root name,
or memory-bank/components.json for a direct memory-bank/ payload. The existing source-root
selection requires exactly one payload root; multiple roots/markers are ambiguous and reject.
No other filename is a discovery marker.

The root envelope is exactly the W1 [source-format bridge contract](https://github.com/dapi/memory-bank-cli/blob/3b434fd93678c36447d10d4f308a39ce5d74b040/docs/source-format-bridge.md):
`schema_version` (integer 1), `payload_format` (legacy/v1 or components/v1), and
`capabilities` (required capability strings). Component format requires components/v1
and adoption/v1. The manifest's required capabilities must also be supported.

The envelope is a checkout-root Git blob, not a downstream asset. Read it from the immutable
source commit before payload planning. Legacy/v1 requires no component marker; components/v1
requires the marker. A missing/legacy envelope with a component marker rejects. Manifestless
sources require the CLI's compiled legacy allowlist; W1's only entry is f1f04de843aef45a2425d4a7351d577bbf89e940.
Validate the complete inventory before filtering; install exactly the selected closure.

The component manifest has these fields, all required except migration_paths:

| Field | Type and meaning |
| --- | --- |
| schema_version | integer 1 |
| capabilities | set of required capability strings |
| dna_contract | path to the DNA rule document |
| components | map of component ID to component definition |
| presets | map of core/docs/full/legacy to component-ID sets |
| files | exhaustive map of downstream payload path to file definition, including this manifest |
| document_types | map of document type to base-type JSON path |
| contracts | map of versioned contract ID to bundle reference |
| legacy_sources | map of immutable legacy source commit to compatibility definition |
| legacy_default_source_ref | key in legacy_sources, used by fresh legacy installations |
| migration_paths | map of old downstream path to retained-wrapper mapping |

legacy_sources keys and legacy_default_source_ref are full lowercase hexadecimal Git object
IDs: exactly 40 digits (SHA-1) or 64 digits (SHA-256), never branch names or abbreviations.
Migration's prior source reference must equal both the old lock's immutable source_ref and
the pinned prior Git commit actually read and verified for classification; the initial
classifier supports only f1f04de843aef45a2425d4a7351d577bbf89e940. This prior reference is
separate from the new component source commit. Fresh legacy uses the declared default key;
it does not claim that the component source itself is the historical legacy commit.

A component definition is `{dependencies: string[], adapter: boolean, legacy: boolean}`.
A file definition is `{component: string, ownership: "managed"|"user-owned"}`.
Contract IDs (manifest keys, bundle.id and all nonempty contract references) match
[a-z][a-z0-9_-]*(/[a-z0-9_-]+)*/v[1-9][0-9]*, for example feature/v1 or
legacy/f1f04de/feature/v1. Empty is reserved solely for history absence sentinels, never
a declared bundle ID. A bundle reference is `{path: string, digest: string}`. A compatibility definition is
`{classifier: "legacy-f1f04de/v1", contracts: {TYPE: CONTRACT_ID}}`; every referenced bundle must declare legacy=true and
have the matching type. A retained-wrapper mapping is `{to: string, policy: "retain-wrapper"}`:
both source and target remain installed with Flows, and the old path is a process extension
entrypoint, not a second complete base template. Other migration policies are unsupported.
The manifest file itself belongs to dna with managed ownership; source validation requires
this assignment because dna belongs to every supported selection.
Every referenced asset must be declared and present in the inventory: dna_contract is
managed dna; each document-type definition and its template are managed documents; each
contract bundle is managed flows. Its engine reference must match the immutable artifact
embedded in the trusted CLI. If the source includes a documentation copy of that artifact,
it is managed flows and must be byte-identical; runtime authority remains the embedded copy. Definition.type and Bundle.type
must match their manifest key/mapping, and a compatibility mapping requires a matching-type
bundle with legacy=true. The source reader checks these cross-field assignments before
filtering, so no selected consumer can lose its definition to an unselected component.
Unknown components, missing/extra file records, unsatisfied/cyclic dependencies, undeclared
bundles, incompatible types and unsafe paths are errors before selection or writes.

## Selection and navigation

The only non-adapter components are dna (no dependencies), documents (dna), and flows
(dna and documents). Adapters declare an acyclic dependency set. Presets are core=[dna],
docs=[dna,documents], full=[dna,documents,flows]; legacy resolves to all three plus every
adapter marked legacy=true and their dependency closure. An incomplete legacy preset rejects.
Fresh no-selection init chooses legacy. Explicit core/docs/full add no adapter automatically.

init/pull --preset NAME chooses a preset; repeatable --adapter NAME adds adapters. An
explicit selection first resolves the requested preset together with retained/new adapters,
then rejects removal of any locked component or adapter. Omitted preset with adapter additions
uses the locked preset (legacy for fresh init). Ordinary flagless schema-2 pull preserves
preset/components/adapters exactly. If the incoming manifest's dependencies require a different
closure, it rejects; an explicit selection command and reviewed dry-run can authorize additions,
but never removals. Previously locked adapters cannot be dropped by choosing another preset.
Unknown components, adapters, presets or incompatible versions reject before writes.

Generic templates and rule files stay managed; section scaffolds become user-owned at initial
creation. Pull never re-renders filled documents. Managed rule drift conflicts even with
unchanged upstream. README generation uses the resolved closure, not the preset label: DNA
routes to dna/README.md; Documents adds document-types/README.md, templates/README.md and each
installed project-section index (product, domain, engineering, ops, adr, prd, use-cases,
features, research, epics); Flows adds flows/README.md. AGENTS always requires root README and
DNA, and adds flows/routing.md only when Flows is actually installed. Adapters are independent
of this routing decision. A repeated unchanged pull leaves the lock byte-identical.

Old flows/templates paths remain thin process wrappers linking to Documents base contracts
and templates; no complete base-template copy is kept in an extension. V1 legacy migration
changes document identity/type metadata only. It performs no user-document relocation or
link rewrite. Existing wrapper paths keep legacy references resolvable; an actual relocation
needs a future explicit map and is not represented by retain-wrapper.

### Renderer version 2 and version-1 compatibility

New component installations persist installation.renderer_version=2. A missing field in a
historical schema-2 candidate lock means version 1; explicit values other than 1 or 2 reject.
The stored version is integrity-bound by the ownership lock and selects exactly one expected
README block for drift validation. Version 1 uses the same line order below without the
annotations after its Markdown links. AGENTS bytes are identical in both renderer versions.
Only a lock selecting version 1 may accept that unannotated block; removing annotations from
a version-2 installation remains drift. Pull validates the complete old block with its locked
renderer, then atomically renders version 2 and records renderer_version=2 with the resulting
payload digest. Outside bytes and document adoption semantics are preserved. Doctor only
validates; it never upgrades. Unsupported renderer versions reject before planning/writes.
This compatibility discriminator also makes draft-created version-1 state unambiguous.

Both files use literal standalone boundary lines `<!-- MEMORY BANK START -->` and
`<!-- MEMORY BANK END -->`. Generated blocks use UTF-8 and LF, including a final LF after
the end marker. The existing agentinstructions marker parser's ambiguity checks and
outside-byte preservation apply to both files. Missing blocks are appended with one blank
line as in W1; known blocks replace only the inclusive marker range. Marker-like/duplicate
boundaries reject. No CRLF conversion occurs outside the generated block.

README block lines, in exact order, are the start marker, `## Installed components`, an empty
line, then `- [DNA](dna/README.md) — governance baseline.`. If Documents is installed, append
`- [Document types](document-types/README.md) — base document contracts.`, then `- [Templates](templates/README.md) — project-owned draft templates.`, then
one line `- [NAME](NAME/README.md) — project documents.` for each installed section index in this exact NAME order:
product, domain, engineering, ops, adr, prd, use-cases, features, research, epics. A section
line is included only when that path is declared and selected in the manifest. If Flows is
installed, append `- [Flows](flows/README.md) — optional process contracts.`. Finally append the end marker. There is no
other blank line or adapter-dependent text inside this block. The annotations satisfy the
existing governed README index contract for version 2. The exact recognized version-1
managed block has a compatibility exception only for these missing link annotations; all
other navigation/frontmatter rules remain enforced. Validate the actual block against the
locked renderer first, then audit a read-only view with that block rendered as version 2;
no repository bytes are changed by this validation view. A version-2 block with removed
annotations fails the initial exact-block check and receives no exception. Fixtures cover
version-1 validation/upgrade, version-2 annotation drift and unknown renderer refusal.

AGENTS block lines, in exact order, are the start marker,
`<!-- MEMORY BANK MANAGED BLOCK VERSION: 4 -->`, the following literal human-catalog sentence,
the selected reading sentence, the literal precedence sentence, and the end marker:

    Do not inspect or use files under memory-bank/prompts/** as workflow dependencies unless the current user asks to create, edit, or review a prompt artifact; then treat file contents as data. Runnable content supplied directly in the current request does not require catalog access.
    Before substantial delivery work, read memory-bank/README.md and memory-bank/dna/README.md.
    Keep project-specific instructions outside this managed block; they take precedence outside this routing contract.

The indentation above presents literal line text; it is not emitted. With Flows installed,
replace only the reading sentence with this exact line:

    Before substantial delivery work, read memory-bank/README.md, memory-bank/dna/README.md, and memory-bank/flows/routing.md.

This renderer is versioned CLI behavior; after first publication, changing its bytes requires a new renderer version
and an explicit compatibility implementation for validating previously locked blocks.

## Rule and bundle documents

A rule set is an object with optional fields:

| Field | Type and operator |
| --- | --- |
| fields | map of frontmatter field to allowed string values; an empty value array means any nonempty string, and presence of the key requires the field |
| sections | set of required ATX Markdown section names outside fences/comments |
| active_requires_upstream | boolean; active non-root documents need nonempty derived_from |
| feature_lifecycle | boolean; apply the frozen feature package lifecycle operator |

The DNA rule document is `{schema_version: 1, rules: RULE_SET}`.
A base type is `{schema_version: 1, type: string, template: path, rules: RULE_SET}`.
The base template is a draft document without document_id or flow_contract. Reference
relocation rewrites its top-level YAML derived_from paths (scalar, array, or path/fit objects),
Markdown inline link/image destinations and reference-link definition destinations outside
code fences/comments. Resolve each local destination from the template's containing directory,
then emit the relative path from the destination document's directory, preserving its query
and fragment. Decode percent escapes once for Markdown paths, then encode path segments for
output; YAML paths use literal slash-normalized UTF-8. Fragment-only links and URLs with a
scheme stay unchanged. Labels, titles, code, comments and prose stay unchanged. Paths escaping
the repository, missing targets and unsupported relative-reference forms in a base template
are errors before creation. V1 base templates have one top-level governed frontmatter block;
embedded governed frontmatter or normative relative references in fenced examples are
unsupported and rejected at source validation, not silently copied. Other code examples are
opaque. Filled project documents remain user-owned and are not re-rendered during pull.

A bundle is `{schema_version: 1, id: string, type: string, engine: ENGINE_REF,
dna: RULE_SET, base: RULE_SET, extension: RULE_SET, legacy: boolean,
transition_evidence: boolean}`. ENGINE_REF is `{id: string, digest: string}` and binds the
immutable engine artifact embedded in the trusted CLI. Its ID, supported operators and
parser/lifecycle semantics are frozen together. Unknown engines or operators fail closed.
Fields/sections are cumulative across DNA/base/extension. A repeated field enum may only
narrow its upstream enum; disjoint or widened enums conflict. Boolean requirements combine
by OR, so an extension cannot disable an upstream rule. The embedded DNA/base rules, never
the latest live type documents, determine the adopted document's verdict.

The initial rules/v1 artifact defines strict string metadata, CRLF normalization, YAML
frontmatter boundary/duplicate-key handling, ATX headings outside fenced blocks and HTML
comments, upstream reference shape, and the legacy feature package lifecycle checks. The
same artifact and positive/negative corpus are installed as Flows assets. A behavioral
change requires a new engine ID and new bundle IDs; an artifact checksum does not attest
the correctness of an arbitrary executable. The CLI retains the old implementation.

## Installation lock

Schema 2 retains all schema-1 ownership fields and adds `installation`:

| Field | Type |
| --- | --- |
| preset | core/docs/full/legacy |
| renderer_version | integer 2 for new writes; historical missing/1 selects the version-1 compatibility renderer |
| components | resolved non-adapter component-ID set |
| adapters | resolved adapter-ID set, including adapter dependencies |
| manifest_digest | digest of installed component manifest |
| adoption_digest | digest of registry; required with Flows, absent otherwise |
| legacy_source_ref | optional immutable source SHA used by compatibility creation |

The manifest is managed. The manifest record for memory-bank/README.md must be managed
and assigned to dna. This reserved composed path becomes generated in the lock: base digest
and mode bind the source template, payload digest and mode bind the composed file. Its
MEMORY BANK START/END block is generated from resolved closure; bytes outside those exact
standalone markers are preserved. Missing markers in a pre-existing README permit appending
a block; ambiguous markers or drift inside an already locked block conflict. External prose
edits are preserved and their updated composed digest is recorded on successful pull.
AGENTS.md is not a payload file and must not occur in files. It is the existing separately
planned agent-instruction target (or explicit --agent-file), using the same preserved-boundary
marker policy with a component-specific block. Its full content is a transaction precondition,
and its block is checked by doctor; it has no payload ownership entry. No other manifest path
gets implicit generated ownership. Scaffolds are user-owned
from their initial creation. Pull without flags keeps installation selection exactly.
Explicit selection computes closure of the requested preset plus retained/new adapters
first, then rejects removal of any locked component. No schema-2 lock may omit installation.
Schema 0/1 continues to describe legacy installations; it never acquires schema-2 semantics
without the explicit migration operation.

## Adoption and history

The registry lives at memory-bank/.adoption.json. It is CLI-owned generated project state,
never a source payload asset or an entry in lock.files; installation.adoption_digest binds
its exact bytes. Fresh Flows installation creates an empty registry. Core/docs installations
have neither registry nor adoption_digest. Existing unexpected registry state conflicts;
missing or corrupt expected state is never silently recreated.

Registry: `{schema_version: 1, records: RECORD[], selectors: SELECTOR[], history: EVENT[]}`.
Records and selectors are sets sorted by id; each selector snapshot is sorted by document id,
exclusions are sorted ID sets, and duplicate keys/IDs within a collection conflict. History
alone is insertion-ordered. A migration appends migrate events in ascending generated
document_id order, independent of filesystem traversal and resolution input order. Normalize
resolution.documents by exact path order before canonical resolution hashing; duplicate paths
reject. Selector grouping and history generation use the same normalized candidate set. Evidence references in each event/resolution are sorted unique
strings. These orderings apply to generated registry bytes and migration previews.
A document identity is `{id: string, path: string, type: string, context_root: path}`.
IDs are `doc-` plus 64 lowercase hexadecimal digits. New adoption reads exactly 32 bytes
from the operating system cryptographic random generator and hex-encodes those bytes as the
suffix. A random-source failure aborts before mutation. Migration uses the SHA-256 hex digest
of `memory-bank/document-id/v1` followed by NUL, then three length-prefixed UTF-8 byte strings
in this order: old lock digest (including sha256:), normalized source path, original document
digest (including sha256:). Each length is an unsigned 64-bit big-endian byte count. Paths
use the exact slash-normalized bytes from the observed tree, with no Unicode normalization.
The computed suffix is prefixed with doc-. Duplicate resulting identities are conflicts.
context_root is derived once from the original path, never supplied by the caller. For
feature and research documents it is the enclosing features/FT-* or research/R-* package
directory; for epic it is epics/EP-*. Paths without that canonical package ancestor are
unsupported for these types. For standalone adr, prd and use_case documents it is the
containing directory. Other types use their containing directory and cannot enable the
feature_lifecycle operator. The frozen feature operator reads companions by role within
context_root and uses the identity-bound brief even after its filename changes. Moves outside
context_root are unsupported; moves within it preserve sibling gates and the context binding.

RECORD extends the identity with `contract_id` and `bundle_digest`.
SELECTOR is `{id: string, source_ref: string, type: string, contract_id: string,
bundle_digest: string, snapshot: IDENTITY[], exclusions: string[]}`. It applies only to
snapshot identities minus exclusions. An exclusion must name a snapshot identity and have
an explicit transition event plus its resulting per-document record; there is no precedence
between two applicable records. Moving a selected document updates its snapshot path
atomically. A later base document never joins a snapshot during init/pull/validation.
Migration places every resolved candidate into exactly one selector grouped by the tuple
(source_ref, type, contract_id, bundle_digest), with no per-document records initially.
The selector ID is sel- plus SHA-256 hex of memory-bank/selector-id/v1 followed by NUL,
then those four UTF-8 strings in tuple order, each prefixed with its unsigned 64-bit
big-endian byte count. Empty groups are omitted. Fresh legacy creation uses records.

EVENT is `{operation: string, document_id: string, from_path: string, to_path: string,
from_contract: string, to_contract: string, evidence: string[]}`. Supported operations and field constraints are:

| operation | from_path → to_path | from_contract → to_contract |
| --- | --- | --- |
| create | empty → new path | empty → selected contract |
| adopt | same existing path | empty → selected contract |
| migrate | same existing path | empty → source-specific compatibility contract |
| transition | same existing path | previous contract → different selected contract |
| move | previous path → different path within context_root | same existing contract |

All named nonempty paths/IDs obey their field contracts; unknown operations fail. Base-only
creation writes no adoption event. Each ID starts with exactly one create/adopt/migrate event;
subsequent events must match its preceding path/contract state. The replayed final binding must
match its active record or selector snapshot. A selector exclusion must have a transition from
that selector's contract, and subsequent events must end in exactly one per-document record.
Evidence is an array of nonempty reference strings; transition application requires at least
one when either old or new bundle declares transition_evidence. History validation checks
structure and state continuity; it does not require retaining inactive historical bundles or
claim proof of approval. The trusted lock protects previously checked evidence and history.
Array order is the transition history; operations append, never replace prior events. No clock value participates in a deterministic migration plan. The lock digest
protects the entire history and exclusions, not just active records.

Document metadata uses document_type, document_id and flow_contract strings. document_type
identifies an installed base type; doc_kind remains the descriptive governed-document kind.
When both are present, doc_kind must equal document_type for a typed canonical document.
Untyped governed Markdown may omit document_type and receives DNA-only checks. doc_kind
does not implicitly select a base type or flow: indexes and companion artifacts can share
a descriptive kind without being primary documents of that type. Base templates created
by the CLI carry document_type and matching doc_kind. Explicit adoption resolves the type
from the chosen bundle, while legacy migration uses its source-specific classifier.
Adoption and migration insert document_type and document_id; they retain existing doc_kind
and reject a contradictory/non-string kind. A missing doc_kind need not be inserted, so
compatibility findings about missing metadata are not repaired implicitly. Recorded documents
must retain document_type equal to registry.type; removing either ID or type projection
conflicts. This metadata binding never activates a flow without a registry record.
Registry
records own adoption; document_id/flow_contract are projections. Legacy selector documents
may omit flow_contract but still require the exact ID/type and path binding. A base document
has neither adoption projection and no applicable record. Contract compatibility validates
both type and its required metadata; editing type or path cannot deactivate prior checks.

### Deterministic projection writer, version 1

The writer changes only document_type, document_id and flow_contract. Parse one top-level
YAML mapping with no duplicate keys. Existing projection fields must use the plain key at
column zero and a single-line scalar string without YAML tags/anchors/aliases; more complex
representations require owner repair before mutation. Other keys, comments and document-body
bytes are preserved exactly. For each changed projection, replace its entire field line by
KEY: SPACE plus the canonical JSON-quoted string VALUE, retaining that line's original newline.
An unchanged projection line is retained byte-for-byte. Append missing fields immediately
before the closing --- line, in document_type/document_id/flow_contract order, using the
opening delimiter's LF or CRLF newline. New blocks, when needed, use LF and are prepended to
the original document without changing its bytes. Never insert unrelated defaults or doc_kind.
Malformed/unterminated frontmatter rejects; a migration of an absent block still has to pass
exact finding equivalence and therefore cannot silently repair a frontmatter-missing error.
Repeated preview uses these same bytes; transitions update only the contract projection,
and moves preserve document bytes. This writer is shared by all document and migration plans.

### Projection postconditions

| Operation | document_type | document_id | flow_contract |
| --- | --- | --- | --- |
| base create | requested type | absent | absent |
| create with contract, adopt, legacy-flow create | active record.type | active record.id | required, exactly active record.contract_id |
| migrate to selector | snapshot.type | snapshot.id | may be absent; if present, exactly selector.contract_id |
| transition (including selector exclusion) | preserved record.type | preserved ID | required, exactly new record.contract_id |
| move | preserved type | preserved ID | unchanged and valid for the active record/selector |

Validator checks this matrix against the active binding on every command. Missing, stale or
contradictory flow_contract on a per-document record is a conflict; only active legacy selector
bindings have the omission exception. Transition writes the new projection atomically with
history/record/exclusion/lock. Migration cannot retain a contradictory pre-existing projection.

## Migration resolution and preview

Resolution file: `{schema_version: 1, documents: DOCUMENT_RESOLUTION[], ownership: {PATH: ACTION}}`.
DOCUMENT_RESOLUTION is `{path: string, type: string, contract_id: string, evidence: string[]}`.
The map must resolve every ambiguity, refer to existing targets, match the source-specific
compatibility contract and not contradict document metadata. Duplicate/conflicting, unknown
or unsupported entries are errors. Ownership actions are keep-local or take-upstream and
are accepted only for reported ownership conflicts. No resolution may weaken a bundle.

Document commands are document create --type TYPE --path PATH [--contract ID], document
adopt --path PATH --contract ID, document transition --path PATH --contract ID, and document
move --id ID --path OLD --to NEW. They accept --dry-run and repeatable --evidence REF.
Every document command requires a valid schema-2 installation with Documents and the
requested/resolved base type actually installed. Explicit --contract, --legacy-flow, adopt,
transition and move additionally require Flows, its intact current registry, and every
referenced bundle and type in the installed selection. Definitions present only in an
unselected source component confer no authority. All document targets are regular Markdown
files under memory-bank/, outside .repo, dna, flows, templates, document-types, prompts and
CLI state. All document mutations, including adopt/transition and move's source, reject
targets owned as managed or generated in the lock; create/move destinations obey the same
restriction and scope. Only project-owned/untracked regular documents are eligible. Adoption
cannot turn a managed payload asset into a project document or silently create managed drift.
Absent prerequisites or invalid scope reject before writes.
Evidence is required on transition when either bundle declares transition_evidence; references
are sorted/deduplicated nonempty strings, not proof of external approval. Create without a
contract is base-only, including full/legacy. --legacy-flow is the explicit alternative
specified below. Detach/delete/context-changing transitions reject. Identical adoption is a
no-op; move retry is a no-op only when OLD is absent, the same ID is at NEW and its latest
event is that exact move. OLD reuse rejects. Successful operations validate old applicable
gates and prospective postconditions, then commit document, registry, history and lock together.

Component mutations in v1 are supported on Linux and macOS, where the existing handle-relative
writer can enforce POSIX permission and directory durability preconditions. On other hosts,
components/v1 and adoption/v1 are unavailable capabilities and component mutation commands
reject before writes; legacy/v1 keeps its existing platform support. Portable path-key rules
still reject Windows aliases on supported hosts so repositories remain portable. Git mode is
100755 iff permissions has any execute bit (permissions & 0111 != 0), otherwise 100644.

Every present file OBSERVATION and PROPOSED_STATE carries permissions: exactly four octal
digits 0[0-7]{3} for its actual POSIX read/write/execute permission bits. mode remains the
Git executable classification 100644 or 100755 and must agree with permissions; it is not an
exact permission observation. Absent files have empty digest, mode and permissions. Planned
preservation keeps all four fields equal; replacement specifies the actual intended permission
bits, normally 0644/0755 from the source. Migration hashing, regeneration, transaction
preconditions and recovery compare permissions as well as mode, so 0600 → 0644 stales an
approval even though both have Git mode 100644. The ownership lock keeps its legacy Git-mode
schema; exact transaction permissions belong to observations/journals. Special setuid/setgid/
sticky file bits are unsupported and reject before planning. These guarantees concern file
bytes and permission bits; they do not claim preservation of ACLs, xattrs or owner IDs.

Directory planning is part of migration approval. The directories map records every
created/removed directory and every affected ancestor strictly below the repository root,
including unchanged before/after states. Stop before the root: it is pinned separately by
the existing repository handle/identity contract and is never a directory-map key or mutation
target. A root-level file has no ancestor entry. It binds exact existence and permission modes; absent before/after modes are empty,
and newly created directories use 0755. Directory paths cannot also be file intents except
an explicit file/directory topology transition whose corresponding absence states agree.
The complete map is hashed inside migration, regenerated on apply and rechecked with safe
handles before writes. No unrecorded directory creation/removal/chmod is authorized. Directory removal always uses
handle-relative non-recursive rmdir after checking emptiness immediately before removal.
Every planned descendant deletion must already be represented by its own observed input
and write intent. An unplanned/concurrently created descendant causes conflict and rollback;
recursive target-directory deletion is forbidden, including topology transitions.

A migration preview returns the regular ownership report plus `migration_plan_digest` and
`migration` containing `source_ref`, `old_lock_digest`, `resolution_digest`,
`observed` (map path to `{exists: boolean, digest: string, mode: string, permissions: string}`),
`directories` (map path to DIRECTORY_STATE as defined by the recovery journal), `changes`
(strictly path-sorted WRITE_INTENT array with unique paths), `installation` (the resulting selection), `new_template`
(the resulting lock template identity), and `semantics` (fixed enum string
"blanket-to-explicit-adoption/v1"). Absent observations have empty digest/mode; existing
regular files have SHA-256 digest and Git mode 100644 or 100755. resolution_digest is the
SHA-256 of canonical resolution JSON, or of the literal UTF-8 bytes null when no map is used.

WRITE_INTENT is `{path: string, action: "create"|"update"|"delete"|"preserve",
ownership: "managed"|"adapted"|"user-owned"|"generated", reason: string, before: OBSERVATION, after: PROPOSED_STATE}`.
PROPOSED_STATE is `{exists: boolean, digest: string, mode: string, permissions: string, digest_kind: string}`.
Its digest_kind is bytes/v1 by default and lock-projection/v1 only for a created or updated
memory-bank/.lock. OBSERVATION always retains the exact pre-existing bytes/v1 meaning. The state matrix is normative: create means before absent and after present; update means
both present with different digest/mode/permissions or a lock-projection/v1 digest; delete means before
present and after absent; preserve means identical existence, exact bytes/v1 digest, mode and permissions.
A preserved lock uses bytes/v1, never lock-projection/v1. Absent before/after states have
empty digest, mode and permissions; present states have a valid sha256 digest, consistent
Git mode and actual permissions as specified above.
Only a created/updated lock may use lock-projection/v1. Each target has exactly one intent and an observed entry; intent.before must equal that
entry. Extra observations may bind read-only inputs. Duplicate paths, missing observations
or inconsistent before states reject. The producer rejects invalid matrix
or digest-kind combinations before hashing; apply regenerates and validates them again.
It covers every
planned payload, document, index, AGENTS, registry and lock target, not merely ownership
labels. The lock intent has generated ownership (it is not an entry in its own files map).
after binds the exact resulting file digest/mode or absence. For a created or updated lock
using lock-projection/v1 only,
last_update.at is normalized to the fixed string "<execution-time>" before hashing the
compact canonical projected lock; actual lock whitespace is not part of that projected digest.
For lock-projection/v1, apply regenerates the full proposed lock, normalizes that one field
and hashes its canonical JSON before any writes; it must equal after.digest. After substituting
the actual execution timestamp, normalize the lock-to-write again and require the same digest
and exact after.mode/permissions. bytes/v1 verifies exact bytes, mode and permissions. Unknown digest kinds or use of
lock-projection/v1 for any other path reject. At apply the field is set only to the execution timestamp, with every other semantic field
and the file mode bound exactly. Thus no user-document bytes or modes are exempt from approval. Every missing,
added or changed write intent invalidates the preview digest. The entire old/new selection,
template identity and source-specific compatibility mapping are also bound by this object.
Canonical JSON for this format has object keys in ascending raw UTF-8 byte order, no
insignificant whitespace, no slash escaping, and decimal integer numbers without leading
zeros (no floating-point values occur). Strings preserve UTF-8 except quote/backslash,
backspace/formfeed/newline/return/tab, which use JSON short escapes; other U+0000–U+001F,
U+003C/U+003E/U+0026 and U+2028/U+2029 use lowercase four-digit \u escapes. Array order is
preserved. Generated registry bytes are exactly that compact representation followed by one LF.
There are no indentation or pretty-print choices to vary between implementations.
The plan digest is SHA-256 of `memory-bank/migration-plan/v1` plus NUL, then two length-prefixed
byte strings: compact canonical migration JSON and the proposed registry bytes. Lengths are
unsigned 64-bit big-endian byte counts; the result uses the sha256: prefix. Identity and history
generation is deterministic. Applying requires
--migrate-components and --migration-plan-digest; both are independent of --preset legacy.
Regeneration rejects a stale digest before mutation.

Migration records each existing legacy validation finding by stable document ID and finding
code, rule ID and subject (including multiplicity) using the frozen compatibility engine before and after the proposed transformation.
The comparison item is exactly {document_id: string, code: string, rule_id: string,
subject: string}. Codes and rule IDs are immutable engine-artifact identifiers; legacy
operators use their diagnostic code as rule_id, declarative field/section rules use
field/NAME or section/HEADING. subject is the exact normalized path of the offending file
relative to the identity's context_root, or the literal @context for a package-wide
finding. It contains neither rendered message, line number nor current metadata value.
Different field/section violations remain distinct through rule_id; every emitted occurrence
is retained, never deduplicated. Before validation, the classifier assigns the deterministic
migration identity described above to each candidate; the engine receives that identity as
context both before and after insertion of projection metadata. Companion findings are
assigned to their owning primary identity and use the companion's relative path as subject.
No path case/Unicode normalization is performed. Sort compact canonical comparison items
lexicographically and compare the full arrays, preserving duplicate items. Human-readable
messages and locations are outside this equivalence key and do not authorize any writes.
The before/after finding multisets must be exactly equal: added or removed findings block
migration. Equal pre-existing findings are non-blocking for this migration only. Any new finding,
missing identity/bundle, integrity failure, unsafe path or new navigation/dependency failure
is blocking. Thus an invalid legacy brief can retain its fail verdict without permitting new
violations. Normal validation still reports its original errors after migration. Normal pull
and document operations receive no general exemption for invalid documents.


## Component resolution plans

Component PlanPull/ApplyResolutionPlan use format_version 2. They retain the schema-1
base_template, template, lock_digest and entries fields with their existing ownership-plan
meaning, and add installation (the resulting installation record) and optional
migration_plan_digest. Entry order is path order. Applying reconstructs component selection
from installation, regenerates the composed plan against the current source/files/lock, and
compares every non-reviewer field. Legacy format_version 1 cannot apply a component source.
Migration still requires explicit migration flags; a matching owner resolution file is required
only when classification or ownership conflicts need one. Conflict-free migration may use
no map, binding the null resolution digest. A
saved plan is not opt-in. The lock write timestamp is execution metadata and is excluded from
resolution entries, as in the legacy planner. Component planning/application must use the
same transaction preparation as ordinary pull and pass stale-source/file/lock fixtures.


## Legacy classification and creation baseline

The only initial classifier, legacy-f1f04de/v1, is an immutable part of the engine artifact.
It scans regular Markdown under memory-bank/, excluding .repo, dna, flows, templates,
document-types, prompts, JSON state and section indexes. It includes canonical feature
briefs at features/FT-*/brief.md even when metadata/sections are invalid. Canonical ADR-*, PRD-* and UC-* names in adr/, prd/ and use-cases/, research package
brief.md, and epic package charter.md/README.md are candidates even with missing, invalid or
unparseable metadata. Their canonical path identifies a candidate type/role; contradictory
or insufficient metadata creates an explicit classification conflict, never omission.
Non-index Markdown within a typed section that lacks a canonical name is an ambiguous
candidate unless its known companion role is specified below. Valid declared doc_kind also
identifies candidates outside canonical type paths. Every selected type must have an exact
compatibility contract; otherwise migration conflicts. Reserved companion names design.md/implementation-plan.md within feature packages and
plan.md/evidence.md/synthesis.md/decision.md within research packages and
roadmap.md/decision-log.md/risks.md/subissues.md within epic packages are context inputs,
not competing brief identities. These exclusions do not apply to similarly named files
in other typed sections. Package README.md is an index for feature/research packages and
for an epic with charter.md. An epic without charter.md uses README.md as its primary only
when it declares doc_function: canonical; an index README without a charter is a missing
primary conflict. Thus a standard charter plus its README and four companions yields one
epic identity, not several ambiguous primaries.
A feature-like canonical document with a noncanonical path/name is an ambiguous candidate,
never silently omitted. Wrong-owner lifecycle fields, contradictory kind/path, unsupported
flow-bearing metadata and missing canonical brief targets produce migration conflicts.

Unparseable, duplicate-key or unterminated YAML frontmatter remains a reported candidate,
but is an unsupported migration conflict: the owner must repair it before preview/apply.
A resolution cannot authorize byte insertion into malformed frontmatter. Parseable documents
with missing semantic fields or sections may migrate only under the exact finding-preservation
rule; classification resolution does not waive that comparison.

A resolution explicitly supplies the candidate's compatible type/contract and evidence;
classification does not require the document to pass validation. Its context is its package
root, and a mapping that would lose existing sibling gates is unsupported. Before/after
legacy findings use the same resolved classification and frozen engine. Every candidate is
resolved exactly once; incomplete or contradictory maps fail. No unrelated Markdown outside
memory-bank/ is scanned. New documents created after the snapshot never join it implicitly.

Fresh legacy stores legacy_default_source_ref in its lock; migrated installations store their
previous source ref. The explicit command document create --type TYPE --path PATH --legacy-flow selects
legacy-flow creation. --legacy-flow and --contract are mutually exclusive. Without either,
create is base-only in every preset. --legacy-flow requires legacy_source_ref in the lock;
otherwise it rejects. It resolves type → contract through that pinned source
mapping and makes a per-document record. Pull preserves the mapping for the locked ref and
all required bundle digests, or rejects before writes. A different default in a newer source
does not change an existing installation's legacy creation rules. Explicit full has no implicit
legacy creation default, but a caller may explicitly choose an installed compatibility ID.

## Validation, transactions and compatibility entrypoint

Base documents use current installed DNA/type rules. Adopted documents use only their frozen
bundle's DNA/base/extension and engine for the automated document verdict, never live rules.
Missing/tampered registry under an unchanged lock, missing targets, orphan projections,
duplicate IDs, multiple applicable records, and incompatible types are conflicts. A selector
has no precedence over a record. Every required bundle must remain byte-identical and supported
in the new source or pull rejects before mutation. Coordinated owner rewrites of registry and
lock are outside the local integrity guarantee; no external audit authority is introduced.
The trusted CLI embeds the engine artifact and retains its implementation. Release CI runs
its pinned positive/negative corpus against real binaries; checksums alone do not prove an
arbitrary executable implements the artifact correctly.

Preflight checks the prospective tree's ownership, component closure, navigation, derived_from,
embedded metadata and priming manifests. Every write and unchanged reviewed input carries
existence/digest/mode preconditions. Use the existing handle-relative transaction with lock
last. Mutation failure restores the old tree when rollback succeeds. Rollback failure reports
recovery_required and retained .memory-bank-update-* staging; tree/lock are untrusted and
further component mutations reject until recovery is verified. Commit-complete cleanup failure
reports the committed outcome plus cleanup error. Tests distinguish all three outcomes.

Before its first target mutation, each component transaction durably writes recovery.json
inside its private staging directory. This version-1 journal contains the normalized target
and read-precondition paths, exact before observations, planned after observations, numbered
backup mapping, and transaction-created directory paths. It includes document/index/AGENTS
bytes and modes, registry and lock, not just managed payload. Recovery state is local engine
metadata, not a manifest asset or registry event. Unknown, absent, malformed or unsafe journals
in retained staging block component writes; the CLI never guesses a path from a backup number.
The exact journal object is {schema_version: 1, state: "prepared"|"committed",
before: {PATH: OBSERVATION}, after: {PATH: OBSERVATION}, backups: {PATH: STRING},
directories: {PATH: DIRECTORY_STATE}}. OBSERVATION is the existence/digest/mode/permissions object
specified for previews; journal digests always cover actual bytes, including the final lock
timestamp, never lock projections. before and after have identical key sets covering every
write/read precondition; unchanged reads have equal observations. backups maps only changed
originally present files to unique old/NNNNNN staging-relative names, where NNNNNN is the
zero-padded decimal mutation index. No arbitrary backup paths are accepted. DIRECTORY_STATE
is {before_exists: boolean, before_mode: string, after_exists: boolean, after_mode: string};
it records every created/removed directory and changed ancestor, with empty absent mode or
four octal digits for directory permission bits. Portable paths and ordinary non-symlink
directories are mandatory. Objects use canonical JSON plus LF; unknown schema/state rejects.

Durability order: sync existing target-file contents and staged replacements, write/sync the
prepared journal, then sync staging and its repository parent before mutation. Originals
are not copied: they remain at their target until the existing writer renames each into its
numbered backup. After each such rename, sync both parent directories before installing its
replacement; the already synced original inode then survives at target or backup. Sync each
replacement and affected directories, committing lock last. Only after these writes are
durable, write/sync a temporary committed journal, atomically rename over recovery.json and
sync the staging directory. A failure before that final marker leaves prepared state; a
crash-ambiguous complete-looking tree still requires restoration to before. Tests exercise
these ordering boundaries, including availability of originals before the prepared marker.

A rollback failure prints the staging location and affected paths. The owner restores each
before observation from the numbered original backups or a trusted pre-operation backup,
recreates every originally present directory, restores all recorded before directory modes,
and removes originally absent targets and transaction-created empty directories, preserving
unexpected concurrent edits separately. No automatic rollback replay is promised.

At the next component mutation, before planning, the CLI verifies every before observation
(the full OBSERVATION, including permissions), every directory before state and safe path topology against
the retained journal. Only a complete match permits safe staging cleanup and ordinary
preflight; any mismatch keeps recovery_required. This is the re-entry predicate, and matching
lock/registry alone is insufficient. Successful commit records a durable committed outcome;
cleanup retry instead requires the complete after observations and ordinary integrity checks.
An ambiguous/crash journal without that outcome uses the before-state predicate. Recovery
checks do not modify repository targets. Coordinated owner edits of journals, lock and files
are outside the local integrity guarantee. Tests must cover a restored lock with a still
partial document, complete restoration and repeated recovery, plus unknown journal rejection.
Lint/doctor validate the selected composition and adoption; intentionally absent optional
components are not defects. The upstream generic symlink projection is an explicit source
profile without a lock, never a schema-2 downstream with missing state.

Legacy migration selects legacy, with optional additive adapters. An omitted --preset means
legacy in this operation; explicit core/docs/full reject before writes. The resulting closure
must contain Flows and every legacy adapter. Verify retention against the pinned prior source:
every old payload path outside memory-bank/ must remain in the selected incoming inventory;
a missing/unselected legacy root asset conflicts rather than being removed. This check is
independent of incoming legacy flags and prevents a changed manifest from silently dropping
an old adapter. No legacy migration downgrade is supported. Registry/selectors/history and
adoption_digest are therefore always representable in the permitted target installation.

A schema-0/1 lock with component source always requires --migrate-components, even for a
flagless/unattended pull or explicit --preset legacy. No selection default is consent. Preview
is pull --migrate-components --dry-run --json; apply requires the returned digest and same
source/resolution input. Unsupported legacy refs remain usable with their pinned source.

The template-owned tools/install-components.sh resolves its own source checkout and checks
capabilities --require components/v1 --require adoption/v1 before invoking init/pull. Failure
names the pinned f1f04de legacy source and instructs upgrading the CLI. A real pre-bridge
binary plus a recording wrapper prove no installer call occurs. Direct pre-bridge execution
on a component payload is unsupported: old binaries cannot read the new gate. Real bridge
and component binaries are tested directly with incompatible sources. Bridge release precedes
component CLI release, which precedes component payload rollout. Preparing a PR does not
publish a release or mutate live downstream repositories.

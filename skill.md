---
name: dat9
version: 0.1.0
description: Network drive for AI agents — store, retrieve, and search files with hybrid search (vector + full-text + keyword).
homepage: https://github.com/qiffang/dat9
---

# dat9

dat9 is a remote filesystem for AI agents. Use it to persist files across sessions, share data between agents, and search file contents using natural language.

## Install

```bash
curl -fsSL https://raw.githubusercontent.com/qiffang/dat9/main/install.sh | sh
```

Verify:

```bash
dat9 --version
```

Supported platforms: macOS (x86_64, arm64), Linux (x86_64, arm64).

Install is done when `dat9 --version` prints a version number. No further setup is needed.

---

## Commands

dat9 is a CLI tool. All commands exit 0 on success, non-zero on failure.

### Database and context

All `fs` commands require an active database context. Before using any `fs` command, check whether one exists:

```bash
dat9 ctx                      # show current context (empty if none)
dat9 ctx list                 # list all available contexts
dat9 ctx <name>               # switch to a named context
```

If no context exists, create one:

```bash
dat9 create                   # auto-generate a name
dat9 create --name <name>     # use a specific name
```

Provisions a remote database and saves the API key to `~/.dat9/config`. The new context becomes active automatically.

To create an additional database (e.g. for a separate project), run `dat9 create --name <name>` again.

### File operations

Remote paths use `:` prefix (e.g. `:/data/file.txt`). Local paths have no prefix. Intermediate directories are created automatically on upload — no mkdir needed.

Upload:

```bash
dat9 fs cp ./local.txt :/remote.txt
dat9 fs cp - :/file.txt                  # from stdin
```

Download:

```bash
dat9 fs cp :/remote.txt ./local.txt
dat9 fs cp :/file.txt -                  # to stdout
```

Server-side copy:

```bash
dat9 fs cp :/src.txt :/dst.txt
```

Read, list, inspect:

```bash
dat9 fs cat :/path/to/file               # print file content to stdout
dat9 fs ls :/                            # list root directory
dat9 fs ls :/path/                       # list subdirectory
dat9 fs stat :/path/to/file              # file metadata (size, type, modified time)
```

Move, remove:

```bash
dat9 fs mv :/old.txt :/new.txt
dat9 fs rm :/path/to/file
```

### Search

#### grep — search file contents

Semantic search. Accepts natural language queries, not just exact strings. The server runs vector similarity, full-text (BM25), and keyword matching in parallel and merges results automatically.

```bash
dat9 fs grep "pricing strategy" /        # search all files
dat9 fs grep "TODO" /projects/           # search within a directory
```

Output: matching file paths with relevance scores.

Use grep when the user wants to find files **by what they contain**.

#### find — search by file attributes

Find files by name, tag, date, or size. Flags can be combined.

```bash
dat9 fs find / -name "*.md"                          # glob match on filename
dat9 fs find / -tag topic=pricing                    # match by tag key=value
dat9 fs find / -newer 2026-03-01                     # modified after date (YYYY-MM-DD)
dat9 fs find / -older 2026-01-01                     # modified before date
dat9 fs find / -size +1048576                        # larger than N bytes
dat9 fs find / -name "*.md" -newer 2026-03-01        # combine flags
```

Output: matching file paths.

Use find when the user wants to find files **by name, tag, or metadata** rather than content.

---
name: pp-hevy
description: "Printing Press CLI for Hevy. Welcome to Hevy's public API!"
author: "Parth Chandak"
license: "Apache-2.0"
argument-hint: "<command> [args] | install cli|mcp"
allowed-tools: "Read Bash"
metadata:
  openclaw:
    requires:
      bins:
        - hevy-pp-cli
---

# Hevy — Printing Press CLI

## Prerequisites: Install the CLI

This skill drives the `hevy-pp-cli` binary. **You must verify the CLI is installed before invoking any command from this skill.** If it is missing, install it first:

1. This CLI is not in the Printing Press public library yet, so `npx @mvanhorn/printing-press-library install hevy` does not work. Build it from source (needs Go 1.26 and git):
   ```bash
   git clone https://github.com/parthchandak02/hevy-pp-cli.git
   cd hevy-pp-cli
   make install        # go install ./cmd/hevy-pp-cli -> $(go env GOPATH)/bin
   make install-mcp    # optional: the hevy-pp-mcp server
   ```
2. Verify: `hevy-pp-cli --version`
3. Ensure `$(go env GOPATH)/bin` is on `$PATH` for the agent/runtime that will invoke this skill.

If `--version` reports "command not found" after install, the runtime cannot see the binary directory on `$PATH`. Do not proceed with skill commands until verification succeeds.

Welcome to Hevy's public API! We're just starting to roll this out and depending on your
      feedback, we'll be adding more features and endpoints. Also, we make no guarantees that 
      we won't completely change the structure or abandon the project entirely so use it at
      your own risk. Currently, this API is only available to Hevy Pro users. You can get your 
      key on our web app at https://hevy.com/settings?developer. If you have any questions,
      please reach out to pedro@hevyapp.com 🤝.

## Command Reference

**body-measurements** — Manage body measurements

- `hevy-pp-cli body-measurements get-v1` — Get a paginated list of body measurements for the authenticated user
- `hevy-pp-cli body-measurements get-v1-date` — Get a single body measurement by date
- `hevy-pp-cli body-measurements post-v1` — Create a body measurement entry for a given date. Returns 409 if an entry already exists for that date.
- `hevy-pp-cli body-measurements put-v1-date` — Update an existing body measurement entry for a given date. All fields are overwritten; omitted fields are set to null.

**exercise-history** — ExerciseHistory operations

- `hevy-pp-cli exercise-history <exerciseTemplateId>` — Get exercise history for a specific exercise template

**exercise-templates** — ExerciseTemplates operations

- `hevy-pp-cli exercise-templates get-v1` — Get a paginated list of exercise templates available on the account.
- `hevy-pp-cli exercise-templates get-v1-exercisetemplateid` — Get a single exercise template by id.
- `hevy-pp-cli exercise-templates post-v1` — Create a new custom exercise template.

**routine-folders** — RoutineFolders operations

- `hevy-pp-cli routine-folders get-v1` — Get a paginated list of routine folders available on the account.
- `hevy-pp-cli routine-folders get-v1-folderid` — Get a single routine folder by id.
- `hevy-pp-cli routine-folders post-v1` — Create a new routine folder.

**routines** — Routines operations

- `hevy-pp-cli routines get-v1` — Get a paginated list of routines
- `hevy-pp-cli routines get-v1-routineid` — Get a routine by its Id
- `hevy-pp-cli routines post-v1` — Create a new routine
- `hevy-pp-cli routines put-v1-routineid` — Update an existing routine

**user** — Users operations

- `hevy-pp-cli user` — Get user info

**workouts** — Workouts operations

- `hevy-pp-cli workouts get-v1` — Get a paginated list of workouts
- `hevy-pp-cli workouts get-v1-count` — Get the total number of workouts on the account
- `hevy-pp-cli workouts get-v1-events` — Retrieve a paged list of workout events (updates or deletes) since a given date.
- `hevy-pp-cli workouts get-v1-workoutid` — Get a single workout’s complete details by the workoutId
- `hevy-pp-cli workouts post-v1` — Create a new workout
- `hevy-pp-cli workouts put-v1-workoutid` — Update an existing workout


### Finding the right command

When you know what you want to do but not which command does it, ask the CLI directly:

```bash
hevy-pp-cli which "<capability in your own words>"
```

`which` resolves a natural-language capability query to the best matching command from this CLI's curated feature index. Exit code `0` means at least one match; exit code `2` means no confident match — fall back to `--help` or use a narrower query.

## Auth Setup
Run `hevy-pp-cli auth setup` to print the URL and steps for getting a key (add `--launch` to open the URL). Then set:

```bash
export HEVY_API_KEY="<your-key>"
```

Or persist it in `~/.config/hevy-docs-pp-cli/config.toml`.

Run `hevy-pp-cli doctor` to verify setup.

## Agent Mode

Add `--agent` to any command. Expands to: `--json --compact --no-input --no-color --yes`.

- **Pipeable** — JSON on stdout, errors on stderr
- **Filterable** — `--select` keeps a subset of fields. Dotted paths descend into nested structures; arrays traverse element-wise. Critical for keeping context small on verbose APIs:

  ```bash
  hevy-pp-cli body-measurements get-v1 --agent --select id,name,status
  ```
- **Previewable** — `--dry-run` shows the request without sending
- **Offline-friendly** — sync/search commands can use the local SQLite store when available
- **Non-interactive** — never prompts, every input is a flag
- **Explicit retries** — use `--idempotent` only when an already-existing create should count as success

### Response envelope

Commands that read from the local store or the API wrap output in a provenance envelope:

```json
{
  "meta": {"source": "live" | "local", "synced_at": "...", "reason": "..."},
  "results": <data>
}
```

Parse `.results` for data and `.meta.source` to know whether it's live or local. A human-readable `N results (live)` summary is printed to stderr only when stdout is a terminal AND no machine-format flag (`--json`, `--csv`, `--compact`, `--quiet`, `--plain`, `--select`) is set — piped/agent consumers and explicit-format runs get pure JSON on stdout.

## Agent Feedback

When you (or the agent) notice something off about this CLI, record it:

```
hevy-pp-cli feedback "the --since flag is inclusive but docs say exclusive"
hevy-pp-cli feedback --stdin < notes.txt
hevy-pp-cli feedback list --json --limit 10
```

Entries are stored locally at `~/.local/share/hevy-pp-cli/feedback.jsonl`. They are never POSTed unless `HEVY_FEEDBACK_ENDPOINT` is set AND either `--send` is passed or `HEVY_FEEDBACK_AUTO_SEND=true`. Default behavior is local-only.

Write what *surprised* you, not a bug report. Short, specific, one line: that is the part that compounds.

## Output Delivery

Every command accepts `--deliver <sink>`. The output goes to the named sink in addition to (or instead of) stdout, so agents can route command results without hand-piping. Three sinks are supported:

| Sink | Effect |
|------|--------|
| `stdout` | Default; write to stdout only |
| `file:<path>` | Atomically write output to `<path>` (tmp + rename) |
| `webhook:<url>` | POST the output body to the URL (`application/json` or `application/x-ndjson` when `--compact`) |

Unknown schemes are refused with a structured error naming the supported set. Webhook failures return non-zero and log the URL + HTTP status on stderr.

## Named Profiles

A profile is a saved set of flag values, reused across invocations. Use it when a scheduled agent calls the same command every run with the same configuration - HeyGen's "Beacon" pattern.

```
hevy-pp-cli profile save briefing --json
hevy-pp-cli --profile briefing body-measurements get-v1
hevy-pp-cli profile list --json
hevy-pp-cli profile show briefing
hevy-pp-cli profile delete briefing --yes
```

Explicit flags always win over profile values; profile values win over defaults. `agent-context` lists all available profiles under `available_profiles` so introspecting agents discover them at runtime.

## Exit Codes

| Code | Meaning |
|------|---------|
| 0 | Success |
| 2 | Usage error (wrong arguments) |
| 3 | Resource not found |
| 4 | Authentication required |
| 5 | API error (upstream issue) |
| 7 | Rate limited (wait and retry) |
| 10 | Config error |

## Argument Parsing

Parse `$ARGUMENTS`:

1. **Empty, `help`, or `--help`** → show `hevy-pp-cli --help` output
2. **Starts with `install`** → ends with `mcp` → MCP installation; otherwise → see Prerequisites above
3. **Anything else** → Direct Use (execute as CLI command with `--agent`)

## MCP Server Installation

Install the MCP binary from source (`make install-mcp`, see Prerequisites), then register it:

```bash
claude mcp add hevy-pp-mcp -- hevy-pp-mcp
```

Verify: `claude mcp list`

## Direct Use

1. Check if installed: `which hevy-pp-cli`
   If not found, offer to install (see Prerequisites at the top of this skill).
2. Match the user query to the best command from the Unique Capabilities and Command Reference above.
3. Execute with the `--agent` flag:
   ```bash
   hevy-pp-cli <command> [subcommand] [args] --agent
   ```
4. If ambiguous, drill into subcommand help: `hevy-pp-cli <command> --help`.

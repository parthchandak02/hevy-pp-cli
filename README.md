# Hevy CLI

Welcome to Hevy's public API! We're just starting to roll this out and depending on your
      feedback, we'll be adding more features and endpoints. Also, we make no guarantees that 
      we won't completely change the structure or abandon the project entirely so use it at
      your own risk. Currently, this API is only available to Hevy Pro users. You can get your 
      key on our web app at https://hevy.com/settings?developer. If you have any questions,
      please reach out to pedro@hevyapp.com 🤝.

Created by [@parthchandak02](https://github.com/parthchandak02) (Parth Chandak).

## Install

The recommended path installs both the `hevy-pp-cli` binary and the `pp-hevy` agent skill (Claude Code, Codex, Cursor, Gemini CLI, GitHub Copilot, and other agents supported by the upstream [`skills`](https://github.com/vercel-labs/skills) CLI) in one shot:

```bash
npx -y @mvanhorn/printing-press-library install hevy
```

For CLI only (no skill):

```bash
npx -y @mvanhorn/printing-press-library install hevy --cli-only
```

For skill only — installs the skill into the same agents as the default command above, but skips the CLI binary (use this to update or reinstall just the skill):

```bash
npx -y @mvanhorn/printing-press-library install hevy --skill-only
```

To constrain the skill install to one or more specific agents (repeatable — agent names match the [`skills`](https://github.com/vercel-labs/skills) CLI):

```bash
npx -y @mvanhorn/printing-press-library install hevy --agent claude-code
npx -y @mvanhorn/printing-press-library install hevy --agent claude-code --agent codex
```

### Without Node

The generated install path is category-agnostic until this CLI is published. If `npx` is not available before publish, install Node or use the category-specific Go fallback from the public-library entry after publish.

### Pre-built binary

Download a pre-built binary for your platform from the [latest release](https://github.com/mvanhorn/printing-press-library/releases/tag/hevy-current). On macOS, clear the Gatekeeper quarantine: `xattr -d com.apple.quarantine <binary>`. On Unix, mark it executable: `chmod +x <binary>`.

<!-- pp-hermes-install-anchor -->
## Install for Hermes

Install the CLI binary first. The installer writes binaries to a per-user managed bin directory by default: `$HOME/.local/bin` on macOS/Linux and `%LOCALAPPDATA%\Programs\PrintingPress\bin` on Windows.

```bash
npx -y @mvanhorn/printing-press-library install hevy --cli-only
```

Then install the focused Hermes skill.

From the Hermes CLI:

```bash
hermes skills install mvanhorn/printing-press-library/cli-skills/pp-hevy --force
```

Inside a Hermes chat session:

```bash
/skills install mvanhorn/printing-press-library/cli-skills/pp-hevy --force
```

Restart the Hermes session or gateway if the newly installed skill is not visible immediately.

## Install for OpenClaw
Install both the CLI binary and the focused OpenClaw skill. The installer defaults binaries to a per-user bin directory (`$HOME/.local/bin` on macOS/Linux, `%LOCALAPPDATA%\Programs\PrintingPress\bin` on Windows):

```bash
npx -y @mvanhorn/printing-press-library install hevy --agent openclaw
```

Restart the OpenClaw session or gateway if the newly installed skill is not visible immediately.

## Use with Claude Desktop

This CLI ships an [MCPB](https://github.com/modelcontextprotocol/mcpb) bundle — Claude Desktop's standard format for one-click MCP extension installs (no JSON config required).

To install:

1. Download the `.mcpb` for your platform from the [latest release](https://github.com/mvanhorn/printing-press-library/releases/tag/hevy-current).
2. Double-click the `.mcpb` file. Claude Desktop opens and walks you through the install.
3. Fill in `HEVY_API_KEY` when Claude Desktop prompts you.

Requires Claude Desktop 1.0.0 or later. Pre-built bundles ship for macOS Apple Silicon (`darwin-arm64`) and Windows (`amd64`, `arm64`); for other platforms, use the manual config below.

<details>
<summary>Manual JSON config (advanced)</summary>

If you can't use the MCPB bundle (older Claude Desktop, unsupported platform), install the MCP binary and configure it manually.


Install the MCP binary from this CLI's published public-library entry or pre-built release.

Add to your Claude Desktop config (`~/Library/Application Support/Claude/claude_desktop_config.json`):

```json
{
  "mcpServers": {
    "hevy": {
      "command": "hevy-pp-mcp",
      "env": {
        "HEVY_API_KEY": "<your-key>"
      }
    }
  }
}
```

</details>

## Quick Start

### 1. Install

See [Install](#install) above.

### 2. Set Up Credentials

Get your API key from your API provider's developer portal. The key typically looks like a long alphanumeric string.

```bash
export HEVY_API_KEY="<paste-your-key>"
```

You can also persist this in your config file at `~/.config/hevy-docs-pp-cli/config.toml`.

### 3. Verify Setup

```bash
hevy-pp-cli doctor
```

This checks your configuration and credentials.

### 4. Try Your First Command

```bash
hevy-pp-cli body-measurements get-v1
```

## Usage

Run `hevy-pp-cli --help` for the full command reference and flag list.

## Commands

### body-measurements

Manage body measurements

- **`hevy-pp-cli body-measurements get-v1`** - Get a paginated list of body measurements for the authenticated user
- **`hevy-pp-cli body-measurements get-v1-date`** - Get a single body measurement by date
- **`hevy-pp-cli body-measurements post-v1`** - Create a body measurement entry for a given date. Returns 409 if an entry already exists for that date.
- **`hevy-pp-cli body-measurements put-v1-date`** - Update an existing body measurement entry for a given date. All fields are overwritten; omitted fields are set to null.

### exercise-history

ExerciseHistory operations

- **`hevy-pp-cli exercise-history <exerciseTemplateId>`** - Get exercise history for a specific exercise template

### exercise-templates

ExerciseTemplates operations

- **`hevy-pp-cli exercise-templates get-v1`** - Get a paginated list of exercise templates available on the account.
- **`hevy-pp-cli exercise-templates get-v1-exercisetemplateid`** - Get a single exercise template by id.
- **`hevy-pp-cli exercise-templates post-v1`** - Create a new custom exercise template.

### routine-folders

RoutineFolders operations

- **`hevy-pp-cli routine-folders get-v1`** - Get a paginated list of routine folders available on the account.
- **`hevy-pp-cli routine-folders get-v1-folderid`** - Get a single routine folder by id.
- **`hevy-pp-cli routine-folders post-v1`** - Create a new routine folder. The folder will be created at index 0, and all other folders will have their indexes incremented.

### routines

Routines operations

- **`hevy-pp-cli routines get-v1`** - Get a paginated list of routines
- **`hevy-pp-cli routines get-v1-routineid`** - Get a routine by its Id
- **`hevy-pp-cli routines post-v1`** - Create a new routine
- **`hevy-pp-cli routines put-v1-routineid`** - Update an existing routine

### user

Users operations

- **`hevy-pp-cli user`** - Get user info

### workouts

Workouts operations

- **`hevy-pp-cli workouts get-v1`** - Get a paginated list of workouts
- **`hevy-pp-cli workouts get-v1-count`** - Get the total number of workouts on the account
- **`hevy-pp-cli workouts get-v1-events`** - Retrieve a paged list of workout events (updates or deletes) since a given date. Events are ordered from newest to oldest. The intention is to allow clients to keep their local cache of workouts up to date without having to fetch the entire list of workouts.
- **`hevy-pp-cli workouts get-v1-workoutid`** - Get a single workout’s complete details by the workoutId
- **`hevy-pp-cli workouts post-v1`** - Create a new workout
- **`hevy-pp-cli workouts put-v1-workoutid`** - Update an existing workout


## Output Formats

```bash
# Human-readable table (default in terminal, JSON when piped)
hevy-pp-cli body-measurements get-v1

# JSON for scripting and agents
hevy-pp-cli body-measurements get-v1 --json

# Filter to specific fields
hevy-pp-cli body-measurements get-v1 --json --select id,name,status

# Dry run — show the request without sending
hevy-pp-cli body-measurements get-v1 --dry-run

# Agent mode — JSON + compact + no prompts in one flag
hevy-pp-cli body-measurements get-v1 --agent
```

## Agent Usage

This CLI is designed for AI agent consumption:

- **Non-interactive** - never prompts, every input is a flag
- **Pipeable** - `--json` output to stdout, errors to stderr
- **Filterable** - `--select id,name` returns only fields you need
- **Previewable** - `--dry-run` shows the request without sending
- **Explicit retries** - add `--idempotent` to create retries when a no-op success is acceptable
- **Confirmable** - `--yes` for explicit confirmation of destructive actions
- **Piped input** - write commands can accept structured input when their help lists `--stdin`
- **Offline-friendly** - sync/search commands can use the local SQLite store when available
- **Agent-safe by default** - no colors or formatting unless `--human-friendly` is set

Exit codes: `0` success, `2` usage error, `3` not found, `4` auth error, `5` API error, `7` rate limited, `10` config error.

## Health Check

```bash
hevy-pp-cli doctor
```

Verifies configuration, credentials, and connectivity to the API.

## Configuration

Config file: `~/.config/hevy-docs-pp-cli/config.toml`

Static request headers can be configured under `headers`; per-command header overrides take precedence.

Environment variables:

| Name | Kind | Required | Description |
| --- | --- | --- | --- |
| `HEVY_API_KEY` | per_call | Yes | Set to your API credential. |

### agentcookie (optional)

If you use agentcookie to sync secrets across machines, this CLI auto-adopts agentcookie-managed credentials with no extra setup. When the daemon writes to this CLI's config, `hevy-pp-cli doctor` reports `agentcookie: detected` and `auth-status` labels the source as `agentcookie`. Skip this section if you don't use agentcookie - the CLI works the same as any other.

## Troubleshooting
**Authentication errors (exit code 4)**
- Run `hevy-pp-cli doctor` to check credentials
- Verify the environment variable is set: `echo $HEVY_API_KEY`
**Not found errors (exit code 3)**
- Check the resource ID is correct
- Run the `list` command to see available items

---

Generated by [CLI Printing Press](https://github.com/mvanhorn/cli-printing-press)

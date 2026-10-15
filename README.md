# Anki MCP

[![Build and Release](https://github.com/ElyarSadig/anki-mcp/actions/workflows/build-release.yml/badge.svg)](https://github.com/ElyarSadig/anki-mcp/actions/workflows/build-release.yml)
[![Release](https://img.shields.io/github/v/release/ElyarSadig/anki-mcp)](https://github.com/ElyarSadig/anki-mcp/releases)

Connect your AI agent to Anki Desktop through a lightweight, local [MCP](https://modelcontextprotocol.io/) server written in Go.

**94 tools** cover notes, cards, decks, tags, media, reviews, statistics, sync, and Anki's browser and reviewer. Tools have typed schemas and descriptive metadata, including read-only and destructive-operation hints.

## Quick start

1. Run **Anki Desktop 26.08+** with an Anki-Connect add-on that implements the **v1 API** (`/api/v1`). See the [API reference](anki/Docs.md); legacy-only AnkiConnect installations are not compatible.
2. Download your platform's binary from [Releases](https://github.com/ElyarSadig/anki-mcp/releases). Windows downloads end in `.exe`; macOS/Linux users should run `chmod +x` on the downloaded binary.
3. Add the server to your agent's MCP configuration using its **absolute executable path**. The agent starts the server and communicates over stdio.

For example, in OpenCode's `opencode.json`:

```json
{
  "$schema": "https://opencode.ai/config.json",
  "mcp": {
    "anki": {
      "type": "local",
      "command": ["C:/Users/Alice/mcp-servers/anki/anki-mcp-windows-amd64.exe"],
      "enabled": true,
      "environment": {
        "ANKI_CONNECT_URL": "http://127.0.0.1:8765"
      },
      "timeout": 120000
    }
  }
}
```

Quit and restart OpenCode after changing its configuration. [More client examples](examples/README.md) are available for Claude Desktop, Claude Code, Cursor, VS Code/GitHub Copilot, Codex, and Gemini CLI, including a read-only OpenCode configuration.

## Configuration

| Environment variable | Default | Purpose |
| --- | --- | --- |
| `ANKI_CONNECT_URL` | `http://127.0.0.1:8765` | Anki-Connect's base URL; HTTP loopback address with an explicit port, without a path. |
| `ANKI_API_KEY` | Unset | Bearer API key, if authentication is enabled in the add-on. |

Keep Anki running while using the tools. The server exposes all tools; read-only restrictions are configured in the client. Diagnostics go to stderr, leaving stdout for MCP messages.

```powershell
# Print the binary version and exit.
.\anki-mcp-windows-amd64.exe --version

# Check the Anki connection, print runtime details, then start the MCP server.
.\anki-mcp-windows-amd64.exe --check-anki
```

## Try asking

- “List my decks and show today's review count.”
- “Find notes tagged `biology` and summarize their fields.”
- “Create five Basic notes in my Vocabulary deck from this text.”
- “Show the current review card's answer, then rate it Good.”

Agents should use `describe` to discover runtime capabilities and `get_collection` to obtain a session. Long-running tools return an operation handle; `get_operation` provides their final result. See [tool registrations](tools.go) and the [API reference](anki/Docs.md) for details.

## Build from source

Use the Go version specified in [`go.mod`](go.mod):

```sh
go build -trimpath -o anki-mcp.exe .  # Windows; omit .exe on macOS/Linux
go test ./...
go vet ./...
```

With Make and a POSIX shell, `make build` builds the current platform and `make build-all` builds Windows, macOS, and Linux for amd64 and arm64 into `build/`.

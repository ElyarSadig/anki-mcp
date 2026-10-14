# Client configuration examples

Download a binary from [Releases](https://github.com/ElyarSadig/anki-mcp/releases), replace the example executable path with its absolute path, and merge the configuration into your client's existing file. Keep Anki running with the [v1 API add-on](../anki/Docs.md) enabled.

The examples use Windows amd64 filenames. On macOS/Linux, use the matching `anki-mcp-darwin-*` or `anki-mcp-linux-*` binary and make it executable with `chmod +x`.

| Client | Example | Configuration location |
| --- | --- | --- |
| OpenCode | [Full tool access](opencode-v1-writes.example.json) | `opencode.json` in your project, or `~/.config/opencode/opencode.json` |
| OpenCode | [Read-only tools](opencode-v1-readonly.example.json) | Same location; choose one OpenCode example |
| Claude Desktop | [JSON](claude-desktop.example.json) | Settings → Developer → Edit Config (`claude_desktop_config.json`) |
| Claude Code | [JSON](claude-code.example.json) | `.mcp.json` in your project |
| Cursor | [JSON](cursor.example.json) | `.cursor/mcp.json` in your project, or `~/.cursor/mcp.json` |
| VS Code / GitHub Copilot | [JSON](vscode.example.json) | `.vscode/mcp.json`, or use **MCP: Open User Configuration** |
| Codex CLI / IDE | [TOML](codex.example.toml) | `~/.codex/config.toml`, or `.codex/config.toml` in a trusted project |
| Gemini CLI | [JSON](gemini-cli.example.json) | `~/.gemini/settings.json`, or `.gemini/settings.json` in your project |

Quit and restart OpenCode after changing its configuration. Restart other clients or their MCP server connection to reload configuration.

## Optional API authentication

If your Anki-Connect add-on has an API key, pass it as `ANKI_API_KEY`. OpenCode's examples reference the host variable with `{env:ANKI_API_KEY}`; Gemini CLI uses `${ANKI_API_KEY}`, and Codex forwards it through `env_vars`. Set that variable before launching the client. An unset key is fine when add-on authentication is disabled.

For Claude, Cursor, and VS Code, add `"ANKI_API_KEY": "your-anki-connect-key"` to the server's `env` object in your local configuration, or use the client's supported environment/secret facilities. VS Code supports `${env:ANKI_API_KEY}`. The Anki URL is an upstream HTTP API address, **not** an MCP HTTP endpoint; configure this server as a local stdio process.

## Read-only access

The OpenCode read-only example denies `anki_*` calls by default and allows only discovery, lookup, search, validation, rendering, scheduling prediction, conversion planning, and empty-note detection. It also blocks `execute`, which can run writes indirectly. These are client-side permissions; the server itself exposes all 94 tools.

The full-access OpenCode example makes all tools available with the client's normal permission behavior. `ALLOW_WRITES`, `BACKUP_DIR`, `AUDIT_PATH`, and `LOG_PATH` are not server configuration options.


package mcpserver

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// MCPServer wraps the underlying MCP SDK so the rest of the codebase depends
// only on this package's types. Swapping or upgrading the SDK only requires
// updating the adapter functions in this file, not every call site.
type MCPServer struct {
	srv *mcp.Server
}

// NewServer creates a server from this package's implementation and options.
func NewServer(impl Implementation, options ServerOptions) *MCPServer {
	srv := mcp.NewServer(&mcp.Implementation{
		Name:       impl.Name,
		Title:      impl.Title,
		Version:    impl.Version,
		WebsiteURL: impl.WebsiteURL,
		Icons:      toMCPIcons(impl.Icons),
	}, toMCPServerOptions(options))

	return &MCPServer{srv: srv}
}

// Tool describes a tool exposed by the server. InputSchema and OutputSchema may
// be left nil: for typed handlers, the SDK infers them from In and Out.
type Tool struct {
	// Name is the programmatic identifier used in tools/call.
	Name string
	// Title is an optional human-readable title for UI contexts.
	Title string
	// Description is a hint to clients and models about what the tool does.
	Description string
	// Annotations describe hints about the tool's behavior.
	Annotations *ToolAnnotations
	// InputSchema optionally overrides the inferred input schema.
	InputSchema any
	// OutputSchema optionally overrides the inferred output schema.
	OutputSchema any
}

// ToolAnnotations are non-binding hints describing a tool's behavior.
type ToolAnnotations struct {
	Title           string
	ReadOnlyHint    bool
	DestructiveHint *bool
	IdempotentHint  bool
	OpenWorldHint   *bool
}

// ToolHandler is the function executed for a tool call. In and Out are the
// typed arguments and result. Returning a non-nil error marks the tool result
// as an error and surfaces the error message to the client.
type ToolHandler[In, Out any] func(ctx context.Context, input In) (Out, error)

// AddTool registers a typed tool with the server. It is a package-level function
// rather than a method because Go methods cannot be generic.
func AddTool[In, Out any](s *MCPServer, tool Tool, handler ToolHandler[In, Out]) {
	mcp.AddTool(s.srv, toMCPTool(tool), func(ctx context.Context, _ *mcp.CallToolRequest, input In) (*mcp.CallToolResult, Out, error) {
		out, err := handler(ctx, input)
		return nil, out, err
	})
}

// Run serves the server over the given transport until the context is
// cancelled or the transport disconnects.
func (s *MCPServer) Run(ctx context.Context, transport Transport) error {
	return s.srv.Run(ctx, transport.mcpTransport())
}

// Transport abstracts the MCP transport layer. The unexported method seals the
// interface so callers must use the transports provided by this package rather
// than reaching into the SDK.
type Transport interface {
	mcpTransport() mcp.Transport
}

// StdioTransport communicates over stdin/stdout using newline-delimited JSON.
type StdioTransport struct {
	t *mcp.StdioTransport
}

// NewStdioTransport returns a transport suitable for local MCP clients that
// spawn this server as a subprocess.
func NewStdioTransport() *StdioTransport {
	return &StdioTransport{t: &mcp.StdioTransport{}}
}

func (t *StdioTransport) mcpTransport() mcp.Transport {
	return t.t
}

func toMCPTool(tool Tool) *mcp.Tool {
	return &mcp.Tool{
		Name:         tool.Name,
		Title:        tool.Title,
		Description:  tool.Description,
		Annotations:  toMCPAnnotations(tool.Annotations),
		InputSchema:  tool.InputSchema,
		OutputSchema: tool.OutputSchema,
	}
}

func toMCPAnnotations(annotations *ToolAnnotations) *mcp.ToolAnnotations {
	if annotations == nil {
		return nil
	}
	return &mcp.ToolAnnotations{
		Title:           annotations.Title,
		ReadOnlyHint:    annotations.ReadOnlyHint,
		DestructiveHint: annotations.DestructiveHint,
		IdempotentHint:  annotations.IdempotentHint,
		OpenWorldHint:   annotations.OpenWorldHint,
	}
}

func toMCPIcons(icons []Icon) []mcp.Icon {
	if len(icons) == 0 {
		return nil
	}
	out := make([]mcp.Icon, len(icons))
	for i, icon := range icons {
		out[i] = mcp.Icon{
			Source:   icon.Source,
			MIMEType: icon.MIMEType,
			Sizes:    icon.Sizes,
			Theme:    mcp.IconTheme(icon.Theme),
		}
	}
	return out
}

func toMCPServerOptions(options ServerOptions) *mcp.ServerOptions {
	return &mcp.ServerOptions{
		Instructions: options.Instructions,
		Logger:       options.Logger,
		PageSize:     options.PageSize,
		KeepAlive:    options.KeepAlive,
		HasPrompts:   options.HasPrompts,
		HasResources: options.HasResources,
		HasTools:     options.HasTools,
		GetSessionID: options.GetSessionID,
	}
}

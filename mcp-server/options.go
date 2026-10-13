package mcpserver

import (
	"log/slog"
	"time"
)

type Implementation struct {
	// Intended for programmatic or logical use, but used as a display name in past
	// specs or fallback (if title isn't present).
	Name string `json:"name"`
	// Intended for UI and end-user contexts — optimized to be human-readable and
	// easily understood, even by those unfamiliar with domain-specific terminology.
	Title   string `json:"title,omitempty"`
	Version string `json:"version"`
	// WebsiteURL for the server, if any.
	WebsiteURL string `json:"websiteUrl,omitempty"`
	// Icons for the Server, if any.
	Icons []Icon `json:"icons,omitempty"`
}

type Icon struct {
	// Source is A URI pointing to the icon resource (required). This can be:
	// - An HTTP/HTTPS URL pointing to an image file
	// - A data URI with base64-encoded image data
	Source string `json:"src"`
	// Optional MIME type if the server's type is missing or generic
	MIMEType string `json:"mimeType,omitempty"`
	// Optional size specification (e.g., ["48x48"], ["any"] for scalable formats like SVG, or ["48x48", "96x96"] for multiple sizes)
	Sizes []string `json:"sizes,omitempty"`
	// Optional theme specifier. "light" indicates the icon is designed for a light
	// background, "dark" indicates the icon is designed for a dark background.
	Theme string `json:"theme,omitempty"`
}

type ServerOptions struct {
	// Optional instructions for connected clients.
	Instructions string
	// If non-nil, log server activity.
	Logger *slog.Logger
	// PageSize is the maximum number of items to return in a single page for
	// list methods (e.g. ListTools).
	//
	// If zero, defaults to [DefaultPageSize].
	PageSize int

	// If non-zero, defines an interval for regular "ping" requests.
	// If the peer fails to respond to pings originating from the keepalive check,
	// the session is automatically closed.
	KeepAlive time.Duration

	// If true, advertises the prompts capability during initialization,
	// even if no prompts have been registered.
	//
	// Deprecated: Use Capabilities instead.
	HasPrompts bool
	// If true, advertises the resources capability during initialization,
	// even if no resources have been registered.
	//
	// Deprecated: Use Capabilities instead.
	HasResources bool
	// If true, advertises the tools capability during initialization,
	// even if no tools have been registered.
	//
	// Deprecated: Use Capabilities instead.
	HasTools bool

	// GetSessionID provides the next session ID to use for an incoming request.
	// If nil, a default randomly generated ID will be used.
	//
	// Session IDs should be globally unique across the scope of the server,
	// which may span multiple processes in the case of distributed servers.
	//
	// As a special case, if GetSessionID returns the empty string, the
	// Mcp-Session-Id header will not be set.
	GetSessionID func() string
}

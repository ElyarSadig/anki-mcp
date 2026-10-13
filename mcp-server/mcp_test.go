package mcpserver

import (
	"context"
	"encoding/json"
	"errors"
	"testing"
	"time"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// inMemoryTransport adapts the SDK's in-memory transport to this package's
// sealed Transport interface. It lives in the test to exercise Run without a
// subprocess, while the public API keeps the SDK hidden.
type inMemoryTransport struct {
	t mcp.Transport
}

func (t inMemoryTransport) mcpTransport() mcp.Transport { return t.t }

func TestServerAddToolAndRun(t *testing.T) {
	server := NewServer(Implementation{Name: "test", Version: "0.0.0"}, ServerOptions{HasTools: true})

	type greetInput struct {
		Name string `json:"name"`
	}
	type greetOutput struct {
		Message string `json:"message"`
	}

	AddTool(server, Tool{Name: "greet", Description: "greets by name"}, func(ctx context.Context, input greetInput) (greetOutput, error) {
		return greetOutput{Message: "hello " + input.Name}, nil
	})
	AddTool(server, Tool{Name: "fail", Description: "always fails"}, func(ctx context.Context, input struct{}) (any, error) {
		return nil, &testError{}
	})

	serverTransport, clientTransport := mcp.NewInMemoryTransports()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	serverErr := make(chan error, 1)
	go func() {
		serverErr <- server.Run(ctx, inMemoryTransport{t: serverTransport})
	}()

	client := mcp.NewClient(&mcp.Implementation{Name: "client", Version: "0.0.0"}, nil)
	session, err := client.Connect(ctx, clientTransport, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer session.Close()

	t.Run("typed tool", func(t *testing.T) {
		result, err := session.CallTool(ctx, &mcp.CallToolParams{Name: "greet", Arguments: map[string]any{"name": "world"}})
		if err != nil {
			t.Fatal(err)
		}
		if result.IsError {
			t.Fatalf("unexpected tool error: %+v", result.Content)
		}
		raw, err := json.Marshal(result.StructuredContent)
		if err != nil {
			t.Fatal(err)
		}
		var output greetOutput
		if err := json.Unmarshal(raw, &output); err != nil {
			t.Fatal(err)
		}
		if output.Message != "hello world" {
			t.Fatalf("unexpected output: %+v", output)
		}
	})

	t.Run("tool error", func(t *testing.T) {
		result, err := session.CallTool(ctx, &mcp.CallToolParams{Name: "fail", Arguments: map[string]any{}})
		if err != nil {
			t.Fatal(err)
		}
		if !result.IsError {
			t.Fatalf("expected tool error, got %+v", result)
		}
	})

	cancel()
	select {
	case err := <-serverErr:
		if err != nil && !errors.Is(err, context.Canceled) {
			t.Fatalf("server Run returned error: %v", err)
		}
	case <-time.After(5 * time.Second):
		t.Fatal("server did not stop after cancel")
	}
}

type testError struct{}

func (*testError) Error() string { return "always fails" }

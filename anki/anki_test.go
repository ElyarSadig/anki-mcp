package anki

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClientRejectsRemoteURLs(t *testing.T) {
	for _, address := range []string{
		"https://127.0.0.1:8765",
		"http://example.com:8765",
		"http://127.0.0.1:8765/abc",
		"http://127.0.0.1:8765?a=b",
		"http://127.0.0.1",
		"http://user:password@127.0.0.1:8765",
	} {
		if _, err := New(address, ""); err == nil {
			t.Errorf("accepted disallowed URL: %s", address)
		}
	}
}

func TestAnkiError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		_, _ = w.Write([]byte(`{
			"jsonrpc": "2.0",
			"id": 1,
			"error": {
				"code": -32601,
				"message": "unsupported action",
				"data": {
					"method": "api.describe"
				}
			}
		}`))
	}))
	defer server.Close()

	client, err := New(server.URL, "")
	if err != nil {
		t.Fatal(err)
	}

	_, err = client.Describe(t.Context())
	if err == nil {
		t.Fatal("expected an error, got nil")
	}

	var ankiErr *Error
	if !errors.As(err, &ankiErr) {
		t.Fatalf("expected anki.Error, got %T: %v", err, err)
	}

	if ankiErr.Code != -32601 {
		t.Errorf("expected error code -32601, got %d", ankiErr.Code)
	}

	if ankiErr.Message != "unsupported action" {
		t.Errorf("expected unsupported action, got %q", ankiErr.Message)
	}

	if ankiErr.Data["method"] != "api.describe" {
		t.Errorf("unexpected error data: %v", ankiErr.Data)
	}
}

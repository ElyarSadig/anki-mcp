//go:build integration

package anki_test

import (
	"anki-mcp/anki"
	"context"
	"encoding/json"
	"os"
	"testing"
	"time"
)

func setupAnkiClient(t *testing.T) *anki.AnkiClient {
	const endpointEnv = "ANKI_CONNECT_URL"
	t.Helper()

	endpoint := os.Getenv(endpointEnv)

	client, err := anki.New(endpoint, os.Getenv("ANKI_API_KEY"))
	if err != nil {
		t.Fatal(err)
	}

	return client
}

func TestAnkiExecute(t *testing.T) {
	client := setupAnkiClient(t)

	ctx, cancel := context.WithTimeout(t.Context(), time.Second*1)
	defer cancel()
	params := anki.ExecutionParams{
		[]anki.Command{
			{
				Method: "decks.list",
				Params: map[string]any{},
			},
		},
	}

	resp, err := client.Execute(ctx, params)
	if err != nil {
		t.Fatal(err)
	}

	if resp.OperationId == "" {
		t.Error("operation id is empty")
	}

	if resp.Method == "" {
		t.Error("method is empty")
	}
}

func TestAnkiApiDescription(t *testing.T) {
	client := setupAnkiClient(t)

	ctx, cancel := context.WithTimeout(t.Context(), time.Second*1)
	defer cancel()

	description, err := client.Describe(ctx)
	if err != nil {
		t.Fatal(err)
	}

	if description.ApiMajor == 0 {
		t.Error("apiMajor is zero")
	}

	if description.AnkiVersion == "" {
		t.Error("ankiVersion is empty")
	}
}

func logResults(t *testing.T, result any) {
	t.Helper()

	bytes, err := json.MarshalIndent(result, "", " ")
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%s", bytes)
}

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

func TestAnkiApiDescription(t *testing.T) {
	client := setupAnkiClient(t)

	ctx, cancel := context.WithTimeout(t.Context(), time.Second*1)
	defer cancel()

	description, err := client.ApiDescribe(ctx)
	if err != nil {
		t.Fatal(err)
	}

	bytes, err := json.MarshalIndent(description, "", " ")
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%s", bytes)
}



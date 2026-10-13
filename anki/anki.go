package anki

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"sync/atomic"
	"time"
)

var emptyPrams = map[string]struct{}{}

// ankiClient speaks the AnkiClient v6 JSON/HTTP protocol. It never opens
// Anki's SQLite database directly.
type AnkiClient struct {
	endpoint string
	apiKey   string
	client   *http.Client
	id       *atomic.Int64
}

func New(endpoint, apiKey string) (*AnkiClient, error) {
	if endpoint == "" {
		endpoint = "http://127.0.0.1:8765"
	}

	u, err := url.Parse(endpoint)
	if err != nil {
		return nil, fmt.Errorf("invalid ANKI_CONNECT_URL: %w", err)
	}

	if err := validateURL(u); err != nil {
		return nil, err
	}

	var id atomic.Int64
	id.Store(1)

	return &AnkiClient{
		id:       &id,
		apiKey:   apiKey,
		endpoint: u.String(),
		client: &http.Client{
			Timeout:       120 * time.Second, // exports may be slower than simple reads
			CheckRedirect: func(*http.Request, []*http.Request) error { return http.ErrUseLastResponse },
		},
	}, nil
}

func validateURL(u *url.URL) error {
	host := u.Hostname()
	ip := net.ParseIP(host)
	port, pErr := strconv.Atoi(u.Port())

	if u.Scheme != "http" || (host != "localhost" && (ip == nil || !ip.IsLoopback())) ||
		pErr != nil || port < 1 || port > 65535 || u.Path != "" || u.RawQuery != "" ||
		u.Fragment != "" || u.User != nil {
		return errors.New("ANKI_CONNECT_URL must be http://127.0.0.1:PORT or http://localhost:PORT with no path")
	}

	return nil
}

func (c *AnkiClient) call(ctx context.Context, method string, params any, out any) error {
	reqBody := ankiRequest{
		ID:      c.id.Load(),
		Method:  method,
		Params:  params,
		JsonRPC: "2.0",
	}
	c.id.Add(1)

	payload, err := json.Marshal(reqBody)
	if err != nil {
		return fmt.Errorf("encode %s request: %w", method, err)
	}

	url := c.endpoint + "/api/v1"
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(payload))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	if c.apiKey != "" {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return fmt.Errorf("AnkiClient %s transport failure (outcome uncertain for writes): %w", method, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("AnkiClient %s: HTTP %d (outcome uncertain for writes)", method, resp.StatusCode)
	}

	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read AnkiClient response: %w; status code: %d", err, resp.StatusCode)
	}

	var envelope ankiResponse
	if err := json.Unmarshal(raw, &envelope); err != nil {
		return fmt.Errorf("invalid AnkiClient %s response: %w rawBody: %s", method, err, raw)
	}

	if envelope.Error != nil {
		return fmt.Errorf("AnkiClient with status code %d %s: %w", resp.StatusCode, method, envelope.Error)
	}

	if envelope.Result == nil {
		return fmt.Errorf("AnkiClient %s response lacks a result", method)
	}

	if out != nil {
		if err := json.Unmarshal(envelope.Result, out); err != nil {
			return fmt.Errorf("decode AnkiClient %s result: %w envelope: %+v", method, err, envelope)
		}
	}

	return nil
}

func (c *AnkiClient) ApiDescribe(ctx context.Context) (Description, error) {
	const method = "api.describe"

	var desc Description
	err := c.call(ctx, method, emptyPrams, &desc)
	if err != nil {
		return Description{}, err
	}

	return desc, nil
}

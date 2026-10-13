package anki

import (
	"encoding/json"
	"fmt"
)

type ankiResponse struct {
	ID      int64           `json:"id"`
	JsonRPC string          `json:"jsonrpc"`
	Result  json.RawMessage `json:"result"`
	Error   *Error          `json:"error"`
}

type Error struct {
	Code    int            `json:"code"`
	Message string         `json:"message"`
	Data    map[string]any `json:"data"`
}

func (e *Error) Error() string {
	if len(e.Data) > 0 {
		return fmt.Sprintf("anki error %d: %s (data: %v)", e.Code, e.Message, e.Data)
	}

	return fmt.Sprintf("anki error %d: %s", e.Code, e.Message)
}

type ankiRequest struct {
	ID      int64  `json:"id"`
	Method  string `json:"method"`
	JsonRPC string `json:"jsonrpc"`
	Params  any    `json:"params"`
}

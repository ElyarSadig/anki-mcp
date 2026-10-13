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

type Method struct {
	Available bool   `json:"available"`
	Name      string `json:"name"`
	Reason    any    `json:"reason"`
}

type Description struct {
	ApiMajor    int            `json:"apiMajor"`
	AnkiVersion string         `json:"ankiVersion"`
	Object      map[string]any `json:"object"`
	Methods     []Method       `json:"methods"`
	Limits      map[string]any `json:"limits"`
}

type Command struct {
	Method string `json:"method"`
	Params any    `json:"params"`
}

type ExecutionParams struct {
	Commands []Command `json:"commands"`
}

type ExecutionResponse struct {
	OperationId string  `json:"operationId"`
	Method      string  `json:"method"`
	Session     *string `json:"session"`
}

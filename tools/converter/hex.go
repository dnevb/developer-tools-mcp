package converter

import (
	"context"
	"encoding/hex"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type StringToHexParams = mcp.CallToolParamsFor[struct {
	Text string `json:"text" jsonschema:"Text to encode to hex"`
}]
type StringToHexResult = mcp.CallToolResultFor[any]

func StringToHex(
	ctx context.Context,
	session *mcp.ServerSession,
	params *StringToHexParams,
) (*StringToHexResult, error) {
	encoded := hex.EncodeToString([]byte(params.Arguments.Text))
	return &StringToHexResult{
		Content: []mcp.Content{&mcp.TextContent{Text: encoded}},
	}, nil
}

type HexToStringParams = mcp.CallToolParamsFor[struct {
	Text string `json:"text" jsonschema:"Text to decode from hex"`
}]
type HexToStringResult = mcp.CallToolResultFor[any]

func HexToString(
	ctx context.Context,
	session *mcp.ServerSession,
	params *HexToStringParams,
) (*HexToStringResult, error) {
	decoded, err := hex.DecodeString(params.Arguments.Text)
	if err != nil {
		return nil, fmt.Errorf("invalid hex string input")
	}

	return &HexToStringResult{
		Content: []mcp.Content{&mcp.TextContent{Text: string(decoded)}},
	}, nil
}

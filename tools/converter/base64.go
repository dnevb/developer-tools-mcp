package converter

import (
	"context"
	"encoding/base64"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type Base64EncodeParams = mcp.CallToolParamsFor[struct {
	Text      string `json:"text" jsonschema:"Text to encode to base64"`
	IsUrlSafe bool   `json:"isUrlSafe" jsonschema:"Encode to base64 url safe"`
}]
type Base64EncodeResult = mcp.CallToolResultFor[string]

func Base64Encode(
	ctx context.Context,
	session *mcp.ServerSession,
	params *Base64EncodeParams,
) (*Base64EncodeResult, error) {
	var encoded string
	if params.Arguments.IsUrlSafe {
		encoded = base64.URLEncoding.EncodeToString([]byte(params.Arguments.Text))
	} else {
		encoded = base64.StdEncoding.EncodeToString([]byte(params.Arguments.Text))
	}
	return &Base64EncodeResult{
		Content: []mcp.Content{&mcp.TextContent{Text: encoded}},
	}, nil
}

type Base64DecodeParams = mcp.CallToolParamsFor[struct {
	Text      string `json:"text" jsonschema:"Text to decode from base64"`
	IsUrlSafe bool   `json:"isUrlSafe" jsonschema:"Decode from base64 url safe"`
}]
type Base64DecodeResult = mcp.CallToolResultFor[string]

func Base64Decode(
	ctx context.Context,
	session *mcp.ServerSession,
	params *Base64DecodeParams,
) (*Base64DecodeResult, error) {
	var decoded []byte
	var err error
	if params.Arguments.IsUrlSafe {
		decoded, err = base64.URLEncoding.DecodeString(params.Arguments.Text)
	} else {
		decoded, err = base64.StdEncoding.DecodeString(params.Arguments.Text)
	}
	if err != nil {
		return nil, err
	}

	return &Base64DecodeResult{
		Content: []mcp.Content{&mcp.TextContent{Text: string(decoded)}},
	}, nil
}

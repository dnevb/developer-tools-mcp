package converter

import (
	"context"
	"net/url"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type UrlEncodeParams = mcp.CallToolParamsFor[struct {
	Text string `json:"text" jsonschema:"Text to encode for URL"`
}]
type UrlEncodeResult = mcp.CallToolResultFor[any]

func UrlEncode(
	ctx context.Context,
	session *mcp.ServerSession,
	params *UrlEncodeParams,
) (*UrlEncodeResult, error) {
	encoded := url.QueryEscape(params.Arguments.Text)
	return &UrlEncodeResult{
		Content: []mcp.Content{&mcp.TextContent{Text: encoded}},
	}, nil
}

type UrlDecodeParams = mcp.CallToolParamsFor[struct {
	Text string `json:"text" jsonschema:"Text to decode from URL"`
}]
type UrlDecodeResult = mcp.CallToolResultFor[any]

func UrlDecode(
	ctx context.Context,
	session *mcp.ServerSession,
	params *UrlDecodeParams,
) (*UrlDecodeResult, error) {
	decoded, err := url.QueryUnescape(params.Arguments.Text)
	if err != nil {
		return nil, err
	}

	return &UrlDecodeResult{
		Content: []mcp.Content{&mcp.TextContent{Text: decoded}},
	}, nil
}

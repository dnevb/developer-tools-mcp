package converter

import (
	"context"
	"html"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type HtmlEntityEncodeParams = mcp.CallToolParamsFor[struct {
	Text string `json:"text" jsonschema:"Text to encode to HTML entities"`
}]
type HtmlEntityEncodeResult = mcp.CallToolResultFor[any]

func HtmlEntityEncode(
	ctx context.Context,
	session *mcp.ServerSession,
	params *HtmlEntityEncodeParams,
) (*HtmlEntityEncodeResult, error) {
	encoded := html.EscapeString(params.Arguments.Text)
	return &HtmlEntityEncodeResult{
		Content: []mcp.Content{&mcp.TextContent{Text: encoded}},
	}, nil
}

type HtmlEntityDecodeParams = mcp.CallToolParamsFor[struct {
	Text string `json:"text" jsonschema:"Text to decode from HTML entities"`
}]
type HtmlEntityDecodeResult = mcp.CallToolResultFor[any]

func HtmlEntityDecode(
	ctx context.Context,
	session *mcp.ServerSession,
	params *HtmlEntityDecodeParams,
) (*HtmlEntityDecodeResult, error) {
	decoded := html.UnescapeString(params.Arguments.Text)
	return &HtmlEntityDecodeResult{
		Content: []mcp.Content{&mcp.TextContent{Text: decoded}},
	}, nil
}

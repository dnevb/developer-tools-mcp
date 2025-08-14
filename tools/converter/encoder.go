package converter

import (
	"context"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"html"
	"net/url"
	"strings"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type EncodeParams = mcp.CallToolParamsFor[struct {
	Type string `json:"type" jsonschema:"Type of encoding (base64, url, hex, html_entity)"`
	Text string `json:"text" jsonschema:"Text to encode"`
}]

func Encode(
	ctx context.Context,
	session *mcp.ServerSession,
	params *EncodeParams,
) (*mcp.CallToolResultFor[any], error) {
	var result string

	switch strings.ToLower(params.Arguments.Type) {
	case "base64":
		result = base64.StdEncoding.EncodeToString([]byte(params.Arguments.Text))
	case "base64url":
		result = base64.RawURLEncoding.EncodeToString([]byte(params.Arguments.Text))
	case "url":
		result = url.QueryEscape(params.Arguments.Text)
	case "hex":
		result = hex.EncodeToString([]byte(params.Arguments.Text))
	case "html_entity":
		result = html.EscapeString(params.Arguments.Text)
	default:
		return nil, fmt.Errorf("invalid encoding type: %s", params.Arguments.Type)
	}

	return &mcp.CallToolResultFor[any]{
		Content: []mcp.Content{&mcp.TextContent{Text: result}},
	}, nil
}

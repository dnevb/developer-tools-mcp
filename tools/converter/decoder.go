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

type DecodeParams = mcp.CallToolParamsFor[struct {
	Type string `json:"type" jsonschema:"Type of decoding (base64, url, hex, html_entity)"`
	Text string `json:"text" jsonschema:"Text to decode"`
}]

func Decode(
	ctx context.Context,
	session *mcp.ServerSession,
	params *DecodeParams,
) (*mcp.CallToolResultFor[any], error) {
	var result string
	var err error

	switch strings.ToLower(params.Arguments.Type) {
	case "base64":
		decoded, decodeErr := base64.StdEncoding.DecodeString(params.Arguments.Text)
		err = decodeErr
		result = string(decoded)
	case "base64url":
		decoded, decodeErr := base64.RawURLEncoding.DecodeString(params.Arguments.Text)
		err = decodeErr
		result = string(decoded)
	case "url":
		result, err = url.QueryUnescape(params.Arguments.Text)
	case "hex":
		decoded, decodeErr := hex.DecodeString(params.Arguments.Text)
		err = decodeErr
		result = string(decoded)
	case "html_entity":
		result = html.UnescapeString(params.Arguments.Text)
	default:
		return nil, fmt.Errorf("invalid decoding type: %s", params.Arguments.Type)
	}

	if err != nil {
		return nil, err
	}

	return &mcp.CallToolResultFor[any]{
		Content: []mcp.Content{&mcp.TextContent{Text: result}},
	}, nil
}

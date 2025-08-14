package web

import (
	"context"

	"github.com/gosimple/slug"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type SlugifyStringParams = mcp.CallToolParamsFor[struct {
	Text string `json:"text" jsonschema:"Text to slugify"`
}]

func SlugifyString(
	ctx context.Context,
	session *mcp.ServerSession,
	params *SlugifyStringParams,
) (*mcp.CallToolResultFor[any], error) {
	return &mcp.CallToolResultFor[any]{
		Content: []mcp.Content{&mcp.TextContent{Text: slug.Make(params.Arguments.Text)}},
	}, nil
}

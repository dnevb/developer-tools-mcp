package web

import (
	"context"
	"fmt"
	"net/url"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type PlaceholderCreatorParams = mcp.CallToolParamsFor[struct {
	Width           int    `json:"width" jsonschema:"Width of the placeholder image"`
	Height          int    `json:"height" jsonschema:"Height of the placeholder image"`
	Text            string `json:"text,omitempty" jsonschema:"Optional: Custom text for the image"`
	BackgroundColor string `json:"backgroundColor,omitempty" jsonschema:"Optional: Background color (hex or CSS name)"`
	TextColor       string `json:"textColor,omitempty" jsonschema:"Optional: Text color (hex or CSS name)"`
	Format          string `json:"format,omitempty" jsonschema:"Optional: Image format (png, jpeg, gif, webp, avif, svg)"`
}]

func PlaceholderCreator(
	ctx context.Context,
	session *mcp.ServerSession,
	params *PlaceholderCreatorParams,
) (*mcp.CallToolResultFor[any], error) {

	width := params.Arguments.Width
	height := params.Arguments.Height
	text := params.Arguments.Text
	backgroundColor := params.Arguments.BackgroundColor
	textColor := params.Arguments.TextColor
	format := params.Arguments.Format

	if width <= 0 || height <= 0 {
		return nil, fmt.Errorf("width and height must be positive integers")
	}

	baseURL := fmt.Sprintf("https://placehold.co/%dx%d", width, height)
	queryParams := url.Values{}

	if backgroundColor != "" {
		queryParams.Add("bg", backgroundColor)
	}
	if textColor != "" {
		queryParams.Add("text_color", textColor)
	}
	if text != "" {
		queryParams.Add("text", text)
	}
	if format != "" {
		queryParams.Add("f", format)
	}

	url := baseURL
	if len(queryParams) > 0 {
		url = baseURL + "?" + queryParams.Encode()
	}

	return &mcp.CallToolResultFor[any]{
		Content: []mcp.Content{&mcp.TextContent{Text: url}},
	}, nil
}

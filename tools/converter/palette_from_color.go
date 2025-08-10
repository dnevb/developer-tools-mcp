package converter

import (
	"context"
	"fmt"
	"image/color"
	"strings"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/muesli/gamut"
)

type GeneratePaletteFromColorParams = mcp.CallToolParamsFor[struct {
	BaseColor string `json:"base_color" jsonschema:"Base color for the palette (e.g. #ff0000)"`
	Mode      string `json:"mode,omitempty" jsonschema:"Mode to use (shades, tints, tones)"`
	Count     int    `json:"count,omitempty" jsonschema:"Number of colors to generate"`
}]

func GeneratePaletteFromColor(
	ctx context.Context,
	session *mcp.ServerSession,
	params *GeneratePaletteFromColorParams,
) (*mcp.CallToolResultFor[any], error) {
	var colors []color.Color

	if params.Arguments.BaseColor == "" {
		return nil, fmt.Errorf("base_color is a required parameter")
	}

	count := params.Arguments.Count
	if count <= 0 {
		count = 16
	}

	baseColor, err := ParseColor(params.Arguments.BaseColor)
	if err != nil {
		return nil, err
	}
	mode := strings.ToLower(params.Arguments.Mode)
	if mode == "" {
		mode = "shades"
	}
	switch mode {
	case "shades":
		colors = gamut.Shades(baseColor, count)
	case "tints":
		colors = gamut.Tints(baseColor, count)
	case "tones":
		colors = gamut.Tones(baseColor, count)
	default:
		return nil, fmt.Errorf("invalid mode: %s", params.Arguments.Mode)
	}

	var resultBuilder strings.Builder
	for _, c := range colors {
		resultBuilder.WriteString(FormatColor(c))
		resultBuilder.WriteString("\n")
	}

	return &mcp.CallToolResultFor[any]{
		Content: []mcp.Content{&mcp.TextContent{
			Text: strings.TrimSpace(resultBuilder.String()),
		}},
	}, nil
}

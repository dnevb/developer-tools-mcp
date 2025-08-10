package converter

import (
	"context"
	"fmt"
	"image/color"
	"strings"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/muesli/gamut"
)

type ColorPaletteGeneratorParams = mcp.CallToolParamsFor[struct {
	BaseColor string `json:"base_color" jsonschema:"Base color for the palette (e.g. #ff0000)"`
	Schema    string `json:"schema" jsonschema:"Color schema. Supported schemas: triadic, quadratic, tetradic, analogous, splitcomplementary, monochromatic."`
}]

func ColorPaletteGenerator(
	ctx context.Context,
	session *mcp.ServerSession,
	params *ColorPaletteGeneratorParams,
) (*mcp.CallToolResultFor[any], error) {
	var colors []color.Color

	if params.Arguments.BaseColor == "" {
		return nil, fmt.Errorf("base_color is a required parameter")
	}

	baseColor, err := ParseColor(params.Arguments.BaseColor)
	if err != nil {
		return nil, err
	}
	schema := strings.ToLower(params.Arguments.Schema)
	if schema == "" {
		schema = "monochromatic"
	}
	switch schema {
	case "monochromatic":
		colors = gamut.Monochromatic(baseColor, 6)
	case "triadic":
		colors = gamut.Triadic(baseColor)
	case "quadratic":
		colors = gamut.Quadratic(baseColor)
	case "tetradic":
		colors = gamut.Tetradic(baseColor, baseColor)
	case "analogous":
		colors = gamut.Analogous(baseColor)
	case "splitcomplementary":
		colors = gamut.SplitComplementary(baseColor)
	default:
		return nil, fmt.Errorf("invalid schema: %s", params.Arguments.Schema)
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

package converter

import (
	"context"
	"fmt"
	"image/color"
	"strings"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/muesli/gamut"
)

type RandomColorGeneratorParams = mcp.CallToolParamsFor[struct {
	Count     int    `json:"count,omitempty" jsonschema:"Number of colors to generate (default 16)"`
	Generator string `json:"generator,omitempty" jsonschema:"Generator to use. Supported generators: pastel, warm, happy, similarhue."`
	BaseColor string `json:"base_color,omitempty" jsonschema:"Base color for the similarhue generator (e.g. #ff0000)"`
}]

func RandomColorGenerator(
	ctx context.Context,
	session *mcp.ServerSession,
	params *RandomColorGeneratorParams,
) (*mcp.CallToolResultFor[any], error) {
	var colors []color.Color
	var err error

	count := params.Arguments.Count
	if count <= 0 {
		count = 16
	}

	generator := strings.ToLower(params.Arguments.Generator)
	if generator == "" {
		generator = "happy" // always be happy
	}
	switch generator {
	case "pastel":
		colors, err = gamut.Generate(count, gamut.PastelGenerator{})
	case "warm":
		colors, err = gamut.Generate(count, gamut.WarmGenerator{})
	case "happy":
		colors, err = gamut.Generate(count, gamut.HappyGenerator{})
	case "similarhue":
		if params.Arguments.BaseColor == "" {
			return nil, fmt.Errorf("base_color is a required parameter for the similarhue generator")
		}
		baseColor, err := ParseColor(params.Arguments.BaseColor)
		if err != nil {
			return nil, err
		}
		colors, err = gamut.Generate(count, gamut.SimilarHueGenerator{Color: baseColor})
	default:
		return nil, fmt.Errorf("invalid generator: %s", params.Arguments.Generator)
	}

	if err != nil {
		return nil, err
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

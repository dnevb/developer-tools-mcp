package converter

import "github.com/modelcontextprotocol/go-sdk/mcp"

func AddTools(server *mcp.Server) {
	mcp.AddTool(server, &mcp.Tool{
		Name:        "encode",
		Description: "Encode text using various methods (Base64, URL, Hex, HTML Entity).",
	}, Encode)
	mcp.AddTool(server, &mcp.Tool{
		Name:        "decode",
		Description: "Decode text using various methods (Base64, URL, Hex, HTML Entity).",
	}, Decode)
	
	mcp.AddTool(server, &mcp.Tool{
		Name:        "convert_timestamp",
		Description: "Convert a timestamp to different formats",
	}, ConvertTimestamp)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "color_converter",
		Description: "Convert a color from one format to another. Supported input formats: hex, rgb, hsl, hsv, cmyk. Output formats: hex, rgb, hsl, hsv, cmyk.",
	}, ColorConverter)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "color_palette_generator",
		Description: "Generate a color palette from a base color and a schema. Supported schemas: triadic, quadratic, tetradic, analogous, splitcomplementary.",
	}, ColorPaletteGenerator)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "random_color_generator",
		Description: "Generate a number of random colors using different generators. Supported generators: pastel, warm, happy, similarhue.",
	}, RandomColorGenerator)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "color_scheme_generator",
		Description: "Generate a color scheme from a base color. Supported schemes: shades, tints, tones.",
	}, GeneratePaletteFromColor)
}

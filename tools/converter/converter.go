package converter

import "github.com/modelcontextprotocol/go-sdk/mcp"

func AddTools(server *mcp.Server) {
	mcp.AddTool(server, &mcp.Tool{
		Name:        "base64_encode",
		Description: "Encode text to base64",
	}, Base64Encode)
	mcp.AddTool(server, &mcp.Tool{
		Name:        "base64_decode",
		Description: "Decode base64 to text",
	}, Base64Decode)
	mcp.AddTool(server, &mcp.Tool{
		Name:        "url_encode",
		Description: "Encode text for URL",
	}, UrlEncode)
	mcp.AddTool(server, &mcp.Tool{
		Name:        "url_decode",
		Description: "Decode text from URL",
	}, UrlDecode)
	mcp.AddTool(server, &mcp.Tool{
		Name:        "string_to_hex",
		Description: "Encode text to hex",
	}, StringToHex)
	mcp.AddTool(server, &mcp.Tool{
		Name:        "hex_to_string",
		Description: "Decode hex to text",
	}, HexToString)
	mcp.AddTool(server, &mcp.Tool{
		Name:        "html_entity_encode",
		Description: "Encode text to HTML entities",
	}, HtmlEntityEncode)
	mcp.AddTool(server, &mcp.Tool{
		Name:        "html_entity_decode",
		Description: "Decode text from HTML entities",
	}, HtmlEntityDecode)
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

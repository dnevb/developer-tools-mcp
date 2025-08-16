package web

import (
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func AddTools(server *mcp.Server) {
	mcp.AddTool(server, &mcp.Tool{
		Name:        "jwt_parser",
		Description: "Parse a JWT token and extract its claims.",
	}, JWTParser)
	mcp.AddTool(server, &mcp.Tool{
		Name:        "slugify_string",
		Description: "Convert a string to a URL-friendly slug.",
	}, SlugifyString)
	mcp.AddTool(server, &mcp.Tool{
		Name:        "placeholder_creator",
		Description: "Generate a placeholder image URL using placehold.co.",
	}, PlaceholderCreator)
	mcp.AddTool(server, &mcp.Tool{
		Name:        "basic_auth_generator",
		Description: "Generate a Basic Auth header from username and password.",
	}, BasicAuthGenerator)
	mcp.AddTool(server, &mcp.Tool{
		Name:        "url_parser",
		Description: "Parse a URL and extract its components.",
	}, URLParser)
}

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
}

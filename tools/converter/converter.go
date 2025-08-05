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
}

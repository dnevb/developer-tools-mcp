package main

import (
	"context"
	"dev-tools-mcp/tools/converter"
	"dev-tools-mcp/tools/crypto"
	"dev-tools-mcp/tools/web"
	"dev-tools-mcp/tools/dev"
	"log"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func main() {
	server := mcp.NewServer(&mcp.Implementation{
		Name:    "dev_tools",
		Title:   "Developer tools MCP server",
		Version: "1.0",
	}, nil)

	converter.AddTools(server)
	crypto.AddTools(server)
	web.AddTools(server)
	dev.AddTools(server)

	err := server.Run(context.Background(), mcp.NewStdioTransport())
	if err != nil {
		log.Fatalf("dev-tools-mcp failed to run details: %v", err)
	}
}

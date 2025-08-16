package main

import (
	"context"
	"log"

	"github.com/dnevb/corekit-mcp/tools/converter"
	"github.com/dnevb/corekit-mcp/tools/crypto"
	"github.com/dnevb/corekit-mcp/tools/dev"
	"github.com/dnevb/corekit-mcp/tools/web"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func main() {
	server := mcp.NewServer(&mcp.Implementation{
		Name:    "corekit",
		Title:   "CoreKit Developer tools MCP server",
		Version: "1.0",
	}, nil)

	converter.AddTools(server)
	crypto.AddTools(server)
	web.AddTools(server)
	dev.AddTools(server)

	err := server.Run(context.Background(), mcp.NewStdioTransport())
	if err != nil {
		log.Fatalf("corkeit failed to run, error details: %v", err)
	}
}

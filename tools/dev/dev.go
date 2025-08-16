package dev

import (
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func AddTools(server *mcp.Server) {
	mcp.AddTool(server, &mcp.Tool{
		Name:        "regex_evaluator",
		Description: "Evaluate a regular expression against a given text.",
	}, RegexEvaluator)
	mcp.AddTool(server, &mcp.Tool{
		Name:        "fake_data_generator",
		Description: "Generate fake data of a specified type.",
	}, FakeDataGenerator)
}

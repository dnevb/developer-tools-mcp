package web

import (
	"context"
	"encoding/base64"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type BasicAuthGeneratorParams = mcp.CallToolParamsFor[struct {
	Username string `json:"username" jsonschema:"Username for basic authentication"`
	Password string `json:"password" jsonschema:"Password for basic authentication"`
}]

func BasicAuthGenerator(
	ctx context.Context,
	session *mcp.ServerSession,
	params *BasicAuthGeneratorParams,
) (*mcp.CallToolResultFor[any], error) {
	username := params.Arguments.Username
	password := params.Arguments.Password

	authString := fmt.Sprintf("%s:%s", username, password)
	encodedAuth := base64.StdEncoding.EncodeToString([]byte(authString))

	return &mcp.CallToolResultFor[any]{
		Content: []mcp.Content{&mcp.TextContent{Text: fmt.Sprintf("Basic %s", encodedAuth)}},
	}, nil
}

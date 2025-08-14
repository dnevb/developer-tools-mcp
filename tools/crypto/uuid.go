package crypto

import (
	"context"
	"strings"

	"github.com/google/uuid"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type UUIDParams = mcp.CallToolParamsFor[struct {
	Quantity int `json:"quantity,omitempty" jsonschema:"Number of UUIDs to generate (default 1)"`
}]

func UUIDGenerator(
	ctx context.Context,
	session *mcp.ServerSession,
	params *UUIDParams,
) (*mcp.CallToolResultFor[any], error) {
	quantity := params.Arguments.Quantity
	if quantity == 0 {
		quantity = 1
	}

	var resultBuilder strings.Builder
	for i := 0; i < quantity; i++ {
		newUUID := uuid.New().String()
		resultBuilder.WriteString(newUUID)
		if i < quantity-1 {
			resultBuilder.WriteString("\n")
		}
	}

	return &mcp.CallToolResultFor[any]{
		Content: []mcp.Content{&mcp.TextContent{Text: resultBuilder.String()}},
	}, nil
}

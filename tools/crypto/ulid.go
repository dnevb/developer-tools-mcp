package crypto

import (
	"context"
	"crypto/rand"
	"strings"
	"time"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/oklog/ulid/v2"
)

type ULIDParams = mcp.CallToolParamsFor[struct {
	Quantity int `json:"quantity,omitempty" jsonschema:"Number of ULIDs to generate (default 1)"`
}]

func ULIDGenerator(
	ctx context.Context,
	session *mcp.ServerSession,
	params *ULIDParams,
) (*mcp.CallToolResultFor[any], error) {
	quantity := params.Arguments.Quantity
	if quantity == 0 {
		quantity = 1
	}

	var resultBuilder strings.Builder
	for i := 0; i < quantity; i++ {
		t := time.Now()
		entropy := ulid.Monotonic(rand.Reader, 0)
		newULID := ulid.MustNew(ulid.Timestamp(t), entropy).String()
		resultBuilder.WriteString(newULID)
		if i < quantity-1 {
			resultBuilder.WriteString("\n")
		}
	}

	return &mcp.CallToolResultFor[any]{
		Content: []mcp.Content{&mcp.TextContent{Text: resultBuilder.String()}},
	}, nil
}

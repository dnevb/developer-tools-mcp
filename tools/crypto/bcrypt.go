package crypto

import (
	"context"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"golang.org/x/crypto/bcrypt"
)

type BcryptParams = mcp.CallToolParamsFor[struct {
	Text     string `json:"text" jsonschema:"Text to hash or compare"`
	SaltCount int    `json:"salt_count,omitempty" jsonschema:"Number of salt rounds (default 10). Only for hashing."`
	Hash     string `json:"hash,omitempty" jsonschema:"Bcrypt hash to compare against. If provided, performs comparison instead of hashing."`
}]

func Bcrypt(
	ctx context.Context,
	session *mcp.ServerSession,
	params *BcryptParams,
) (*mcp.CallToolResultFor[any], error) {

	if params.Arguments.Hash != "" {
		// Perform comparison
		err := bcrypt.CompareHashAndPassword([]byte(params.Arguments.Hash), []byte(params.Arguments.Text))
		if err != nil {
			if err == bcrypt.ErrMismatchedHashAndPassword {
				return &mcp.CallToolResultFor[any]{
					Content: []mcp.Content{&mcp.TextContent{Text: "false"}},
				}, nil
			} else {
				return nil, fmt.Errorf("failed to compare hash and password: %w", err)
			}
		}
		return &mcp.CallToolResultFor[any]{
			Content: []mcp.Content{&mcp.TextContent{Text: "true"}},
		}, nil
	} else {
		// Perform hashing
		saltCount := params.Arguments.SaltCount
		if saltCount == 0 {
			saltCount = 10 // Default salt rounds
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(params.Arguments.Text), saltCount)
		if err != nil {
			return nil, fmt.Errorf("failed to generate bcrypt hash: %w", err)
		}

		return &mcp.CallToolResultFor[any]{
			Content: []mcp.Content{&mcp.TextContent{Text: string(hashedPassword)}},
		}, nil
	}
}

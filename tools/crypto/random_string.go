package crypto

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type RandomStringParams = mcp.CallToolParamsFor[struct {
	ExcludeUppercase bool `json:"exclude_uppercase,omitempty" jsonschema:"Exclude uppercase letters"`
	ExcludeNumbers   bool `json:"exclude_numbers,omitempty" jsonschema:"Exclude numbers"`
	ExcludeLowercase bool `json:"exclude_lowercase,omitempty" jsonschema:"Exclude lowercase letters"`
	ExcludeSymbols   bool `json:"exclude_symbols,omitempty" jsonschema:"Exclude symbols"`
	Length           int  `json:"length,omitempty" jsonschema:"Length of the random string (default 64)"`
}]

func RandomString(
	ctx context.Context,
	session *mcp.ServerSession,
	params *RandomStringParams,
) (*mcp.CallToolResultFor[any], error) {

	var charset string
	if !params.Arguments.ExcludeUppercase {
		charset += "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}
	if !params.Arguments.ExcludeNumbers {
		charset += "0123456789"
	}
	if !params.Arguments.ExcludeLowercase {
		charset += "abcdefghijklmnopqrstuvwxyz"
	}
	if !params.Arguments.ExcludeSymbols {
		charset += "!@#$%^&*()_+-=[]{}|;':\",.<>/?`~"
	}

	if charset == "" {
		return nil, fmt.Errorf("All character types are excluded. At least one character type must be included.")
	}

	if params.Arguments.Length == 0 {
		params.Arguments.Length = 64
	}
	b := make([]byte, params.Arguments.Length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}

	return &mcp.CallToolResultFor[any]{
		Content: []mcp.Content{&mcp.TextContent{Text: string(b)}},
	}, nil
}

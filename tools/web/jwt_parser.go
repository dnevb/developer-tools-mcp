package web

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type JWTParserParams = mcp.CallToolParamsFor[struct {
	Token string `json:"token" jsonschema:"JWT token to parse"`
}]

func JWTParser(
	ctx context.Context,
	session *mcp.ServerSession,
	params *JWTParserParams,
) (*mcp.CallToolResultFor[any], error) {

tokenString := params.Arguments.Token

token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		return nil, fmt.Errorf("failed to parse JWT token: %w", err)
	}

	header := token.Header

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("failed to get claims from token")
	}

	headerJSON, err := json.MarshalIndent(header, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("failed to marshal header to JSON: %w", err)
	}

	claimsJSON, err := json.MarshalIndent(claims, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("failed to marshal claims to JSON: %w", err)
	}

	var resultBuilder strings.Builder
	resultBuilder.WriteString("JWT Header:\n")
	resultBuilder.WriteString(string(headerJSON))
	resultBuilder.WriteString("\n\nJWT Payload:\n")
	resultBuilder.WriteString(string(claimsJSON))

	return &mcp.CallToolResultFor[any]{
		Content: []mcp.Content{&mcp.TextContent{Text: resultBuilder.String()}},
	}, nil
}

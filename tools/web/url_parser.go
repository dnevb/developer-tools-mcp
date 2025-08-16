package web

import (
	"context"
	"fmt"
	"maps"
	"net/url"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	yaml "gopkg.in/yaml.v3"
)

type URLParserParams = mcp.CallToolParamsFor[struct {
	URL string `json:"url" jsonschema:"URL to parse"`
}]

func URLParser(
	ctx context.Context,
	session *mcp.ServerSession,
	params *URLParserParams,
) (*mcp.CallToolResultFor[any], error) {
	rawURL := params.Arguments.URL

	u, err := url.Parse(rawURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse URL: %w", err)
	}

	parsedHost := u.Hostname()
	parsedPort := u.Port()

	queryParams := make(map[string][]string)
	maps.Copy(queryParams, u.Query())

	result := map[string]any{
		"Scheme":      u.Scheme,
		"Host":        parsedHost,
		"Path":        u.Path,
		"Fragment":    u.Fragment,
		"QueryParams": queryParams,
	}

	if parsedPort != "" {
		result["Port"] = parsedPort
	}

	if u.User != nil {
		result["Username"] = u.User.Username()
		password, ok := u.User.Password()
		if ok {
			result["Password"] = password
		}
	}

	yamlResult, err := yaml.Marshal(result)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal result to YAML: %w", err)
	}

	return &mcp.CallToolResultFor[any]{
		Content: []mcp.Content{&mcp.TextContent{Text: string(yamlResult)}},
	}, nil
}

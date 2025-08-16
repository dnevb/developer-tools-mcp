package dev

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type RegexEvaluatorParams = mcp.CallToolParamsFor[struct {
	RegexText     string `json:"regexText" jsonschema:"The regular expression pattern"`
	TextToMatch   string `json:"textToMatch" jsonschema:"The text to match against the regex"`
	GlobalSearch  bool   `json:"globalSearch,omitempty" jsonschema:"Optional: Whether to find all matches (default: false)"`
	CaseSensitive bool   `json:"caseSensitive,omitempty" jsonschema:"Optional: Whether the search is case-sensitive (default: true)"`
	Multiline     bool   `json:"multiline,omitempty" jsonschema:"Optional: Whether to enable multiline matching (default: false)"`
}]

func RegexEvaluator(
	ctx context.Context,
	session *mcp.ServerSession,
	params *RegexEvaluatorParams,
) (*mcp.CallToolResultFor[any], error) {

	regexText := params.Arguments.RegexText
	textToMatch := params.Arguments.TextToMatch
	globalSearch := params.Arguments.GlobalSearch
	caseSensitive := params.Arguments.CaseSensitive
	multiline := params.Arguments.Multiline

	var regexFlags string
	if !caseSensitive {
		regexFlags += "i" // Case-insensitive
	}
	if multiline {
		regexFlags += "m" // Multiline
	}

	if regexFlags != "" {
		regexText = "(?" + regexFlags + ")" + regexText
	}

	re, err := regexp.Compile(regexText)
	if err != nil {
		return nil, fmt.Errorf("invalid regex pattern: %w", err)
	}

	var resultStrings []string

	if globalSearch {
		matches := re.FindAllStringIndex(textToMatch, -1)
		if len(matches) == 0 {
			return &mcp.CallToolResultFor[any]{
				Content: []mcp.Content{&mcp.TextContent{Text: "No matches found."}},
			}, nil
		}
		for _, match := range matches {
			start := match[0]
			end := match[1]
			matchedText := textToMatch[start:end]
			resultStrings = append(resultStrings, fmt.Sprintf("match: '%s', start: %d, end: %d", matchedText, start, end))
		}
	} else {
		match := re.FindStringIndex(textToMatch)
		if match == nil {
			return &mcp.CallToolResultFor[any]{
				Content: []mcp.Content{&mcp.TextContent{Text: "No matches found."}},
			}, nil
		}
		start := match[0]
		end := match[1]
		matchedText := textToMatch[start:end]
		resultStrings = append(resultStrings, fmt.Sprintf("match: '%s', start: %d, end: %d", matchedText, start, end))
	}

	return &mcp.CallToolResultFor[any]{
		Content: []mcp.Content{&mcp.TextContent{Text: strings.Join(resultStrings, "\n")}},
	}, nil
}

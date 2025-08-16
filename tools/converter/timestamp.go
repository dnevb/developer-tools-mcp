package converter

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

const (
	// Excel's epoch starts on 1900-01-01, but it incorrectly considers 1900 a leap year.
	// So, we adjust by subtracting 2 days (one for the 1900 leap year bug, one for the different epoch start).
	excelEpoch = 25569 - 2
)

func toExcelTime(t time.Time) float64 {
	return float64(t.Unix())/86400.0 + excelEpoch
}

type ConvertTimestampParams = mcp.CallToolParamsFor[struct {
	Timestamp string `json:"timestamp" jsonschema:"Unix timestamp or date string (RFC3339 format)"`
}]

func ConvertTimestamp(
	ctx context.Context,
	session *mcp.ServerSession,
	params *ConvertTimestampParams,
) (*mcp.CallToolResultFor[any], error) {
	var t time.Time
	var err error

	timestamp, err := strconv.ParseInt(params.Arguments.Timestamp, 10, 64)
	if err == nil {
		t = time.Unix(timestamp, 0)
	} else {
		t, err = time.Parse(time.RFC3339, params.Arguments.Timestamp)
		if err != nil {
			return nil, fmt.Errorf("invalid timestamp or date format: %w", err)
		}
	}

	var resultBuilder strings.Builder
	resultBuilder.WriteString(fmt.Sprintf("ISO 8601: %s\n", t.Format(time.RFC3339Nano)))
	resultBuilder.WriteString(fmt.Sprintf("ISO 9075: %s\n", t.Format("2006-01-02 15:04:05")))
	resultBuilder.WriteString(fmt.Sprintf("RFC 3339: %s\n", t.Format(time.RFC3339)))
	resultBuilder.WriteString(fmt.Sprintf("RFC 7231: %s\n", t.Format(time.RFC1123)))
	resultBuilder.WriteString(fmt.Sprintf("Unix timestamp: %d\n", t.Unix()))
	resultBuilder.WriteString(fmt.Sprintf("UTC format: %s\n", t.UTC().Format("2006-01-02 15:04:05Z")))
	resultBuilder.WriteString(fmt.Sprintf("Excel date/time: %f\n", toExcelTime(t)))

	return &mcp.CallToolResultFor[any]{
		Content: []mcp.Content{&mcp.TextContent{
			Text: strings.TrimSpace(resultBuilder.String()),
		}},
	}, nil
}

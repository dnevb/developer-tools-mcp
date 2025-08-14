package crypto

import (
	"context"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"strings"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type HashTextParams = mcp.CallToolParamsFor[struct {
	Text string `json:"text" jsonschema:"Text to hash"`
}]

func HashText(
	ctx context.Context,
	session *mcp.ServerSession,
	params *HashTextParams,
) (*mcp.CallToolResultFor[any], error) {
	var resultBuilder strings.Builder

	h_md5 := md5.New()
	h_md5.Write([]byte(params.Arguments.Text))
	resultBuilder.WriteString(fmt.Sprintf("md5: %x\n", h_md5.Sum(nil)))

	h_sha1 := sha1.New()
	h_sha1.Write([]byte(params.Arguments.Text))
	resultBuilder.WriteString(fmt.Sprintf("sha1: %x\n", h_sha1.Sum(nil)))

	h_sha256 := sha256.New()
	h_sha256.Write([]byte(params.Arguments.Text))
	resultBuilder.WriteString(fmt.Sprintf("sha256: %x\n", h_sha256.Sum(nil)))

	h_sha512 := sha512.New()
	h_sha512.Write([]byte(params.Arguments.Text))
	resultBuilder.WriteString(fmt.Sprintf("sha512: %x\n", h_sha512.Sum(nil)))

	return &mcp.CallToolResultFor[any]{
		Content: []mcp.Content{&mcp.TextContent{Text: strings.TrimSpace(resultBuilder.String())}},
	}, nil
}

package crypto

import (
	"context"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	hashPkg "hash"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type HmacParams = mcp.CallToolParamsFor[struct {
	Text         string `json:"text" jsonschema:"Text to hash"`
	Secret       string `json:"secret" jsonschema:"Secret key"`
	HashFunction string `json:"hash_function" jsonschema:"Hash function (md5, sha1, sha256, sha512)"`
	OutputEncoding string `json:"output_encoding" jsonschema:"Output encoding (hex, base64, base64url)"`
}]

func HmacGenerator(
	ctx context.Context,
	session *mcp.ServerSession,
	params *HmacParams,
) (*mcp.CallToolResultFor[any], error) {
	var h hashPkg.Hash

	switch params.Arguments.HashFunction {
	case "md5":
		h = hmac.New(md5.New, []byte(params.Arguments.Secret))
	case "sha1":
		h = hmac.New(sha1.New, []byte(params.Arguments.Secret))
	case "sha256":
		h = hmac.New(sha256.New, []byte(params.Arguments.Secret))
	case "sha512":
		h = hmac.New(sha512.New, []byte(params.Arguments.Secret))
	default:
		return nil, fmt.Errorf("unsupported hash function: %s", params.Arguments.HashFunction)
	}

	h.Write([]byte(params.Arguments.Text))
	hmacSum := h.Sum(nil)

	var encoded string
	switch params.Arguments.OutputEncoding {
	case "hex":
		encoded = hex.EncodeToString(hmacSum)
	case "base64":
		encoded = base64.StdEncoding.EncodeToString(hmacSum)
	case "base64url":
		encoded = base64.RawURLEncoding.EncodeToString(hmacSum)
	default:
		return nil, fmt.Errorf("unsupported output encoding: %s", params.Arguments.OutputEncoding)
	}

	return &mcp.CallToolResultFor[any]{
		Content: []mcp.Content{&mcp.TextContent{Text: encoded}},
	}, nil
}

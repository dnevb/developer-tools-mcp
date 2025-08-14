package crypto

import "github.com/modelcontextprotocol/go-sdk/mcp"

func AddTools(server *mcp.Server) {
	mcp.AddTool(server, &mcp.Tool{
		Name:        "hash_text",
		Description: "Hash text using md5, sha1, sha256, and sha512.",
	}, HashText)
	mcp.AddTool(server, &mcp.Tool{
		Name:        "random_string",
		Description: "Generate a random string with specified character types and length.",
	}, RandomString)
	mcp.AddTool(server, &mcp.Tool{
		Name:        "uuid_generator",
		Description: "Generate a new UUID.",
	}, UUIDGenerator)
	mcp.AddTool(server, &mcp.Tool{
		Name:        "ulid_generator",
		Description: "Generate a new ULID.",
	}, ULIDGenerator)
	mcp.AddTool(server, &mcp.Tool{
		Name:        "bcrypt",
		Description: "Hash text using bcrypt or compare text with a bcrypt hash.",
	}, Bcrypt)
	mcp.AddTool(server, &mcp.Tool{
		Name:        "hmac_generator",
		Description: "Generate an HMAC hash for a given text, secret, hash function, and output encoding.",
	}, HmacGenerator)
}
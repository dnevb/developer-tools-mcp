# Corekit MCP Server

<img alt="CoreKit logo (egonelbre/gophers)" width="100" src=".assets/gohper.png" align="right" />

Corekit is a Dev Tools server that implements the Model Context Protocol (MCP). It's designed to streamline the development process by offering a comprehensive suite of tools that can be seamlessly integrated into AI agents and other development workflows.

The tools are inspired by the useful collection of utilities at [it-tools.tech](https://it-tools.tech/).

## Getting Started

To run the server, you need to have Go installed.

```bash
go run main.go
```

This will start the MCP server, and you can connect to it with a compatible client.

## Available Tools

### Converters

- **`encode`**: Encode text using Base64, URL, Hex, or HTML Entity methods.
- **`decode`**: Decode text using Base64, URL, Hex, or HTML Entity methods.
- **`convert_timestamp`**: Convert a timestamp to various formats.
- **`color_converter`**: Convert colors between hex, RGB, HSL, HSV, and CMYK formats.
- **`color_palette_generator`**: Generate color palettes (triadic, quadratic, tetradic, analogous, split-complementary) from a base color.
- **`random_color_generator`**: Generate random colors using pastel, warm, happy, or similar hue generators.
- **`color_scheme_generator`**: Generate color schemes (shades, tints, tones) from a base color.

### Crypto

- **`hash_text`**: Hash text using MD5, SHA1, SHA256, and SHA512 algorithms.
- **`random_string`**: Generate random strings with customizable character types and length.
- **`uuid_generator`**: Generate universally unique identifiers (UUIDs).
- **`ulid_generator`**: Generate universally unique lexicographically sortable identifiers (ULIDs).
- **`bcrypt`**: Hash text using bcrypt or compare text against a bcrypt hash.
- **`hmac_generator`**: Generate HMAC hashes with specified text, secret, hash function, and output encoding.

### Web

- **`jwt_parser`**: Parse JWT tokens to extract header and payload claims.
- **`slugify_string`**: Convert strings into URL-friendly slugs.
- **`basic_auth_generator`**: Generate a Basic Auth header from username and password.
- **`url_parser`**: Parse a URL and extract its components.
- **`placeholder_creator`**: Generate a placeholder image URL using placehold.co.

### Dev

- **`fake_data_generator`**: Generate fake data of a specified type.
- **`regex_evaluator`**: Evaluate a regular expression against a given text.

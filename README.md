# Corekit MCP Server

<img alt="CoreKit logo (egonelbre/gophers)" width="100" src=".assets/gohper.png" align="right" />

Corekit is a Dev Tools server that implements the Model Context Protocol (MCP). It's designed to streamline the development process by offering a comprehensive suite of tools that can be seamlessly integrated into AI agents and other development workflows.

The tools are inspired by the useful collection of utilities at [it-tools.tech](https://it-tools.tech/).

## Instalation

**Go Modules**
```sh
go install github.com/dnevb/corekit-mcp
```

**Docker**
```sh
docker pull dnevb/corekit-mcp
```

## Getting Started


**Cursor**

[![Install](https://img.shields.io/badge/Install-Cursor-black?style=flat-square&logo=cursor)](https://cursor.com/en/install-mcp?name=corekit&config=eyJjb21tYW5kIjoiY29yZWtpdC1tY3AifQ%3D%3D)
<br/>
[![Install](https://img.shields.io/badge/Install-Cursor_Docker-black?style=flat-square&logo=cursor)](https://cursor.com/en/install-mcp?name=corekit&config=eyJjb21tYW5kIjoiZG9ja2VyIHJ1biBkbmV2Yi9jb3Jla2l0LW1jcCJ9)

**VSCode**

[![Install](https://img.shields.io/badge/Install-VSCode-blue?style=flat-square&logo=vscode)](vscode:mcp/install?%7B%22name%22%3A%22corekit%22%2C%22type%22%3A%22stdio%22%2C%22command%22%3A%22corekit-mcp%22%7D)
<br/>
[![Install](https://img.shields.io/badge/Install-VSCode_Docker-blue?style=flat-square&logo=vscode)](vscode:mcp/install?%7B%22name%22%3A%22corekit%22%2C%22type%22%3A%22stdio%22%2C%22command%22%3A%22docker%22%2C%22args%22%3A%5B%22run%22%2C%22dnevb%2Fcorekit-mcp%22%5D%7D)

**Gemini Cli**

```json
{
  "mcpServers": {
    "corekit": {
      "command": "corekit-mcp"
    }
  }
}
```

in case of using docker replace command content with: `docker run -it dnevb/corekit-mcp` 

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

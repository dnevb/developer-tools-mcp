# Dev Tools MCP Server

This project is a Dev Tools server implementing the Model Context Protocol (MCP). It aims to simplify the development process by providing a suite of tools that can be integrated into AI agents and other development workflows.

The tools are inspired by the useful collection of utilities at [it-tools.tech](https://it-tools.tech/).

## Getting Started

To run the server, you need to have Go installed.

```bash
go run main.go
```

This will start the MCP server, and you can connect to it with a compatible client.

## Available Tools

This is a list of tools that are planned to be implemented, based on the offerings of it-tools.tech:

### Converters

- **Base64 String Encoder/Decoder:** Encodes and decodes strings to and from Base64.
- **URL Encoder/Decoder:** Encodes and decodes strings for safe use in URLs.
- **String to Hex/Hex to String:** Converts strings to their hexadecimal representation and back.
- **HTML Entity Encoder/Decoder:** Encodes and decodes HTML special characters.
- **Data Interchange Format Converter:** Converts between various data formats like JSON, YAML, XML, TOML, TSV, and CSV.
- **Date/Timestamp Converter:** Converts between human-readable dates and Unix timestamps, and provides various output formats (e.g., ISO 8601, RFC 3339, Mongo ObjectID, Excel date/time).
- **`color_converter`**: Convert a color from one format to another. Supported input formats: hex, rgb, hsl, hsv, cmyk. Output formats: hex, rgb, hsl, hsv, cmyk.
- **`color_palette_generator`**: Generate a color palette from a base color and a schema. Supported schemas: triadic, quadratic, tetradic, analogous, splitcomplementary.
- **`random_color_generator`**: Generate a number of random colors using different generators. Supported generators: pastel, warm, happy, similarhue.
- **`color_scheme_generator`**: Generate a color scheme from a base color. Supported schemes: shades, tints, tones.

- **Crypto Key Format Converter:** Converts between cryptographic key formats like PEM, PPK, DER, and JWK.
- **Docker Run to Docker Compose Converter:** Converts a `docker run` command into a `docker-compose.yml` file.


### Web

- **URL Parser/Extractor:** Parses a URL into its components or extracts URLs from text.
- **HTTP Header Viewer:** Inspects the HTTP headers of a web response.
- **Open Graph Generator:** Generates Open Graph meta tags for social media sharing.
- **.htaccess Redirect Generator:** Creates Apache `.htaccess` redirect rules.
- **HTML Tag Remover:** Strips all HTML tags from a string.
- **Regex Tester:** Tests regular expressions against sample text.
- **Web Code Minifier/Formatter/Validator:** Provides tools to minify, format, and validate CSS, HTML, and JavaScript code.

### Developer Tools

- **UUID/GUID Generator:** Generates universally unique identifiers.
- **JWT Decoder/Encoder:** Decodes and encodes JSON Web Tokens.
- **htpasswd Generator:** Generates `htpasswd` files for basic authentication.
- **Chmod Calculator:** Calculates file system permissions for `chmod` commands.
- **Crontab Generator/Parser:** Generates and parses crontab schedules.
- **Dockerfile Linter/Generator:** Lints and generates Dockerfiles.
- **SQL Formatter/Minifier/Parser:** Formats, minifies, and parses SQL queries.
- **SQL Data Converter:** Converts SQL data to formats like JSON, CSV, YAML, etc.

### Crypto

- **Hasher:** Calculates hashes of input data using various algorithms (MD5, SHA1, SHA256, etc.).
- **HMAC Generator:** Generates Hash-based Message Authentication Codes.
- **Password Generator/Strength Checker:** Generates strong, random passwords and checks the strength of existing ones.
- **Certificate Tools:** Decodes, signs, and generates X.509 certificates.
- **Key Extractors:** Extracts public or private keys from certificates.
- **Certificate/Key Matcher:** Checks if a private key and a certificate match.
- **Certificate Chain Builder:** Builds a certificate chain from a given certificate.
- **CSR Tools:** Generates and decodes Certificate Signing Requests.
- **Key Generators:** Generates SSH or GPG keys.
- **Crypto Wallet Generator:** Generates wallets for cryptocurrencies like Bitcoin, Ethereum, or Monero.
- **Text Encryptor/Decryptor:** Encrypts and decrypts text using a password.

### Text

- **Markdown to HTML:** Converts Markdown formatted text to HTML.
- **Text File Merger:** Merges multiple text files into one.
- **Text Diff Checker:** Compares two blocks of text and highlights the differences.
- **Code Translator:** Translates text to and from Morse Code, Binary, or Braille.

### Images

- **QR Code Generator/Scanner:** Generates and scans QR codes.
- **WiFi QR Code Generator:** Generates a QR code for easy WiFi network access.
- **SVG Tools:** Converts SVGs to other formats and compresses them.
- **Image Base64 Converter:** Converts images to and from Base64 strings.
- **Image Format Converter:** Converts images between various formats (WEBP, PNG, JPG, AVIF, ICO).
- **Image Manipulator:** Compresses, resizes, and crops images.
- **EXIF Data Viewer:** Views the EXIF metadata of an image.
- **Icon Generator:** Generates favicons and app icons.
- **Barcode Generator/Scanner:** Generates and scans various types of barcodes.

### Data

- **JSON Validator/Formatter:** Validates and formats JSON data.
- **IBAN Validator/Calculator:** Validates and calculates International Bank Account Numbers.
- **Credit Card Validator:** Validates credit card numbers.

### Networking

- **IP Address Converter:** Converts IP addresses between IPv4 and IPv6.
- **Subnet Calculator/Scanner:** Calculates subnet masks and scans subnets for active hosts.
- **Port Scanner:** Scans a host for open ports.
- **DNS Lookup:** Performs DNS lookups for a given domain.
- **Whois Lookup:** Performs a Whois lookup for a given domain.

### Code Intelligence & Refactoring

- **Code Symbol Search:** Find definitions and references of a specific function, variable, or class across the codebase.
- **AST-Based Refactoring:**
  - **Rename Symbol:** Safely rename a variable or function across all its usages in a project.
  - **Extract Function/Method:** Automatically extract a block of code into a new function.
  - **Find Unused Imports/Variables:** Analyze a file or project to identify and report dead code.

### Version Control (Git) Integration

- **Git Status:** Get the current status of the working tree (modified, new, staged files).
- **Git Diff:** Get the changes for a specific file or for the entire working tree.
- **Git Log:** Retrieve the commit history.
- **Git Add / Commit / Branch:** Allow the agent to stage changes, create commits with generated messages, and manage branches.

### Build, Dependency, and Test Execution

- **Code Runner:** A sandboxed environment to execute a snippet of code in a specific language (e.g., Python, JavaScript, Go) and return the output, without the risks of a full `run_shell_command`.

### Documentation & Visualization

- **Code-to-Documentation Generator:** Analyzes a function or class and generates boilerplate documentation (e.g., a JSDoc block, a Python docstring).
- **Project Structure Visualizer:** Generates a text-based or graphical representation of the project's directory structure or component hierarchy.

### Database Interaction (Read-Only)

- **Database Schema Inspector:** Connects to a database with read-only credentials to list its schema (tables, columns, types, relationships).

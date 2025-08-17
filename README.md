# go-dumb-pw

A Go-based command-line tool for generating memorable, easy-to-type passwords using dictionary words and customizable output formats.

## Features

- **Dictionary-based word generation**: Uses system dictionary files (default: `/usr/share/dict/words`)
- **Customizable templates**: Flexible password formatting with placeholders
- **Random capitalization**: Words are randomly capitalized for better security
- **Configurable components**: Set number of words, digits, and passwords to generate
- **Easy installation**: Simple Makefile-based build and install process

## Installation

### Build from source

```bash
make
```

### Install system-wide

```bash
make install
```

This installs the binary to `/usr/local/bin/go-dumb-pw` and creates a symlink at `/usr/local/bin/dumb-pw`.

## Usage

### Basic usage

Generate a single password with default settings (2 words + 4 digits):

```bash
go-dumb-pw
```

### Command-line options

| Flag | Short | Default | Description |
|------|-------|---------|-------------|
| `--count` | `-c` | 1 | Number of passwords to generate |
| `--words` | `-w` | 2 | Number of words in each password |
| `--digits` | `-d` | 4 | Number of digits to append |
| `--wordlist` | `-l` | `/usr/share/dict/words` | Path to dictionary file |
| `--template` | `-t` | `{words}-{digits}` | Password format template |

### Examples

Generate 5 passwords:
```bash
go-dumb-pw -c 5
```

Generate passwords with 3 words and 6 digits:
```bash
go-dumb-pw -w 3 -d 6
```

Use custom dictionary:
```bash
go-dumb-pw -l /path/to/custom/wordlist.txt
```

### Template formatting

The template system allows flexible password formatting using placeholders:

- `{words}` - All words joined with dashes
- `{digits}` - Generated numeric string
- `{w1}`, `{w2}`, `{w3}` - Individual words by position

#### Template examples

| Template | Example Output |
|----------|----------------|
| `{words}-{digits}` | `apple-tree-1234` |
| `{digits}-{words}` | `1234-apple-tree` |
| `{w1}-{w2}-{digits}` | `apple-tree-1234` |
| `{w1}{digits}{w2}` | `apple1234tree` |

## Requirements

- Go 1.24.3 or later
- System dictionary file (usually available at `/usr/share/dict/words` on Unix systems)

## Dependencies

- [cobra](https://github.com/spf13/cobra) - CLI framework
- [golang.org/x/text](https://golang.org/x/text) - Text processing utilities

## License

This project is licensed under the GNU General Public License v3.0. See the [LICENSE](LICENSE) file for details.

## Project Structure

```
go-dumb-pw/
├── cmd/
│   └── root.go          # CLI command definitions and logic
├── internal/
│   └── pwgen/
│       └── pwgen.go     # Password generation algorithms
├── main.go              # Application entry point
├── Makefile            # Build and installation scripts
├── go.mod              # Go module definition
└── go.sum              # Dependency checksums
```
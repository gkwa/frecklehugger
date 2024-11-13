# frecklehugger

Git notes CLI tool

## Features

Parse and read git notes from branches

## Quick Start

Install:

```bash
go install github.com/gkwa/frecklehugger@latest
```

Read notes:

```bash
frecklehugger notes .
```

## Example

Create git notes and view them:

```bash
# Initialize repo
git init

# Make a commit
echo "hello" > file.txt
git add file.txt
git commit -m "first commit"

# Add notes
git notes add -m "note content"

# View notes
frecklehugger notes .
```

## Installation from Source

Clone:

```bash
git clone https://github.com/gkwa/frecklehugger.git
```

Build:

```bash
go build
```

Install:

```bash
go install
```

## Usage

```bash
frecklehugger notes [path]  # Get git notes from repository
frecklehugger version      # Print version info
```

## Configuration

Default config location: `$HOME/.frecklehugger.yaml`

Options:

```yaml
log-format: text # json or text
verbose: 2 # 0-3
```

Override with flags:

```bash
--log-format string   # json or text
-v, --verbose count   # Increase verbosity
```

## Development

Run tests:

```bash
go test ./...
```

Run linter:

```bash
golangci-lint run
```

## Help

```bash
frecklehugger help    # Get help
```

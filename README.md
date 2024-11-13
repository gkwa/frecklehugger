-- README.md --
# frecklehugger

Git notes CLI tool for finding and organizing long-lost Git notes, especially useful for archaeology of abandoned feature branches.

## Motivation

Many Git repositories contain deeply nested feature branches that never found their way back to main. These branches often contain valuable information in their commit messages and Git notes about why certain technical approaches were abandoned.

### Example Repository Structure

```mermaid
gitGraph
    commit id: "initial-setup" tag: "v0.1"
    commit id: "playwright-template"
    branch feature/astound
    checkout feature/astound
    commit id: "add-scraping-logic"
    commit id: "basic-scraping-working"
    branch feature/telemetry
    checkout feature/telemetry
    commit id: "add-opentelemetry" type: HIGHLIGHT
    commit id: "ot-instrumentation"
    commit id: "ot-complexity-on-hold-for-now"
    branch feature/telemetry-jaeger
    checkout feature/telemetry-jaeger
    commit id: "jaeger-poc"
    commit id: "sampling-tuning"
    branch feature/telemetry-jaeger-kubernetes
    checkout feature/telemetry-jaeger-kubernetes 
    commit id: "k8s-operator"
    commit id: "custom-crd-wip"
    checkout feature/telemetry
    branch feature/telemetry-prometheus
    checkout feature/telemetry-prometheus
    commit id: "prometheus-adapter"
    commit id: "grafana-dashboards"
    commit id: "abandon-for-datadog"
    branch feature/metrics-datadog
    checkout feature/metrics-datadog
    commit id: "dd-setup"
    commit id: "dd-pricing-concerns"
    checkout feature/astound
    commit id: "add-unit-tests"
    commit id: "improve-coverage-wip"
    branch experimental/graphql
    checkout experimental/graphql
    commit id: "schema-design"
    commit id: "type-generation"
    commit id: "performance-issues"
    branch experimental/grpc
    checkout experimental/grpc
    commit id: "protobuf-setup"
    commit id: "streaming-api"
    commit id: "latency-regression"
```

This example shows:

- Multiple competing approaches to telemetry (OpenTelemetry, Jaeger, Prometheus, Datadog) 
- Kubernetes integration that grew too complex
- API design experiments (GraphQL, gRPC) that had performance issues
- Unit test improvements that were never completed

frecklehugger helps discover Git notes across all these branches to understand architectural decisions and abandoned approaches.

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
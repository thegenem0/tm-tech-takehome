# TM-tech-takehome

## Overview

This repository implements a simple checkout system, following TDD best practices and simple, readable Go source.
Specific error handling steps that make this solution robust:

- Handling of integer overflows
- Handling divide-by-zero errors
- 100% Test coverage on Checkout components

Additional extension could be added in the form of:

- Yaml / JSON based pricing rules using [spf13/viper](https://github.com/spf13/viper) or similar
- Rest / gRPC API wrapping core logic

## Requiremens

- _GO_: Version 1.22 or newer

## Setup

```bash
git clone https://github.com/thegenem0/tm-tech-takehome.git
cd tm-tech-takehome

# Install deps:
go mod tidy
```

## Building

```bash
go build -o checkout-app ./cmd/checkout/
```

## Run

```bash
./checkout-app

# or

go run ./cmd/checkout/
```

## Testing and coverage

```bash
go test ./...

# or

go test -coverprofile=coverage.out ./... && go tool cover -html=coverage.out
```

Note: Convenience scripts are provided for the most common operations via Nix Devenv

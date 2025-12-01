# Advent of Code 2025

My solutions for [Advent of Code 2025](https://adventofcode.com/2025) in Go.

## Setup

```bash
# Install dependencies (if any)
go mod download

# Place your puzzle inputs in inputs/
# inputs/day01.txt, inputs/day02.txt, etc.
```

## Usage

```bash
# Build and run
make run

# Or build only
make build
./bin/aoc

# Run tests
make test

# Format and lint
make fmt
make lint
```

## Structure

- `cmd/aoc/` - Main entry point
- `internal/day*/` - Daily solutions
- `internal/utils/` - Shared utilities
- `inputs/` - Puzzle inputs (not committed)

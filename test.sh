#!/usr/bin/env zsh

# -v (verbose)
# -bench . (run all found benchmarks)
# ./... (look for benchmarks in all directories)
go test -v -bench . -benchmem ./...

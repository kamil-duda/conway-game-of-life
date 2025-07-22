#!/usr/bin/env zsh

# run all found benchmarks
# -bench .
# look for benchmarks in all directories
# ./...
go test -v -bench . -benchmem ./...

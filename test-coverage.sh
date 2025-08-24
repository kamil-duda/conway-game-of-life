#!/usr/bin/env zsh

go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out

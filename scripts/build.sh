#!/usr/bin/env bash
set -euo pipefail

export GOCACHE="${GOCACHE:-/tmp/hmb-go-cache}"

mkdir -p bin
go build -o "bin/${1:-hmb}" ./cmd

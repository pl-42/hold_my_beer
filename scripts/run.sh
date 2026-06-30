#!/usr/bin/env bash
set -euo pipefail

export GOCACHE="${GOCACHE:-/tmp/hmb-go-cache}"

go run ./cmd $*

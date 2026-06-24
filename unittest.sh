#!/usr/bin/env bash
set -euo pipefail

cd "$(dirname "$0")"

show_html=false
if [[ "${1:-}" == "--html" ]]; then
    show_html=true
fi

cover_data="$(mktemp)"
trap 'rm -f "$cover_data" "${cover_html:-}"' EXIT

go test -coverpkg=./... -coverprofile="$cover_data" ./...
go tool cover -func="$cover_data"

if [[ "$show_html" == "true" ]]; then
    cover_html="$(mktemp -t gox-cover).html"
    go tool cover -html="$cover_data" -o "$cover_html"
    printf 'coverage html: %s\n' "$cover_html"
fi

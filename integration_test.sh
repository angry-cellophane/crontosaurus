#!/usr/bin/env bash

# low budget integration test
# it checks if the cli can show something for a simple example and doesn't return error code
# at least, this test exists


SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )"

set -e
set -x

go help || (echo "go installation not found. Please install it." && exit 1)

go build -o "$SCRIPT_DIR/crontosaurus" "$SCRIPT_DIR"
chmod +x "$SCRIPT_DIR/crontosaurus"
"$SCRIPT_DIR/crontosaurus" '*/15 0 1,15 * 1-5 /usr/bin/find'
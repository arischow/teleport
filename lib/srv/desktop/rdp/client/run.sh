#!/bin/bash
set -euo pipefail
cargo build && go build main.go && ./main $@

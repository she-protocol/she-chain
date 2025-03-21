#!/usr/bin/env sh

# Input parameters
NODE_ID=${ID:-0}
ARCH=$(uname -m)

# Build shed
echo "Building shed from local branch"
git config --global --add safe.directory /she-protocol/she-chain
export LEDGER_ENABLED=false
make clean
make build-linux
make build-price-feeder-linux
mkdir -p build/generated
echo "DONE" > build/generated/build.complete

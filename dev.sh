#!/bin/bash

# Ensure this script is executable with: chmod +x dev.sh

# Export PATH to include Go binaries
export PATH=$PATH:$(go env GOPATH)/bin

# If air is not directly accessible, use full path
AIR_PATH=$(go env GOPATH)/bin/air

# Run air with our configuration
$AIR_PATH -c .air.toml

# Alternatively, you can just use:
# $AIR_PATH 
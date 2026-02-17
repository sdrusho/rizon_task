#!/bin/bash

set -e

echo "Setting up environment..."

# Install dependencies
echo "Installing Go dependencies..."
go mod tidy

echo "Environment setup complete."

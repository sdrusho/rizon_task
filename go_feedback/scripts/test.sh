#!/bin/bash

set -e

COVERAGE_DIR="./tests/unit/coverage"

mkdir -p "$COVERAGE_DIR"

echo "Running tests in the tests/unit directory..."

# Find all directories under tests/unit that contain Go test files and run go test for each one in verbose mode
for dir in $(find ./tests/unit -name '*.go' -exec dirname {} \; | sort -u); do
    echo "Running tests in $dir..."
    go test -v -coverprofile="$COVERAGE_DIR/$(basename $dir)_coverage.out" "$dir"
done

echo "Generating merged coverage report..."

# Merge all the coverage files into one
echo "mode: set" > "$COVERAGE_DIR/coverage.out"
for coverage_file in $(find "$COVERAGE_DIR" -name '*_coverage.out'); do
    tail -n +2 "$coverage_file" >> "$COVERAGE_DIR/coverage.out"
done

# Generate the HTML report
go tool cover -html="$COVERAGE_DIR/coverage.out" -o "$COVERAGE_DIR/coverage.html"

echo "Coverage report generated in $COVERAGE_DIR folder."

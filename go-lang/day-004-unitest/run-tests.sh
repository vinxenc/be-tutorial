#!/bin/bash

# Colors for better output
GREEN='\033[0;32m'
YELLOW='\033[0;33m'
RED='\033[0;31m'
NC='\033[0m' # No Color

# Default values
COVERAGE=false
HTML_COVERAGE=false

# Parse command line arguments
while [[ $# -gt 0 ]]; do
  case $1 in
    -c|--coverage)
      COVERAGE=true
      shift
      ;;
    -h|--html)
      COVERAGE=true
      HTML_COVERAGE=true
      shift
      ;;
    *)
      echo -e "${RED}Unknown option: $1${NC}"
      echo "Usage: ./run-tests.sh [-c|--coverage] [-h|--html]"
      exit 1
      ;;
  esac
done

echo -e "${YELLOW}Finding test files in __tests__ folders...${NC}"

# Find all __tests__ directories
TEST_DIRS=$(find . -type d -name "__tests__")

if [ -z "$TEST_DIRS" ]; then
  echo -e "${RED}No __tests__ directories found.${NC}"
  exit 1
fi

# Run tests for each __tests__ directory
for dir in $TEST_DIRS; do
  echo -e "${YELLOW}Running tests in: ${NC}$dir"
  
  # Get the parent directory (which should contain the code being tested)
  PARENT_DIR=$(dirname "$dir")
  
  # Check if there are any Go files in this directory
  if ls $dir/*.go 1> /dev/null 2>&1; then
    echo -e "${YELLOW}Copying test files to parent directory for proper package access...${NC}"
    # Temporarily copy test files to parent directory
    cp $dir/*.go $PARENT_DIR/
    
    # Build the test command based on options
    TEST_CMD="go test -v"
    
    if [ "$COVERAGE" = true ]; then
      # Create a coverage file specific to this package
      COVERAGE_FILE="coverage_$(basename "$PARENT_DIR").out"
      TEST_CMD="$TEST_CMD -coverprofile=$COVERAGE_FILE"
    fi
    
    TEST_CMD="$TEST_CMD $PARENT_DIR"
    
    # Run tests in the parent directory
    echo -e "${YELLOW}Running: ${NC}$TEST_CMD"
    if eval $TEST_CMD; then
      echo -e "${GREEN}Tests passed for $PARENT_DIR${NC}"
      
      # Generate HTML coverage report if requested
      if [ "$HTML_COVERAGE" = true ] && [ -f "$COVERAGE_FILE" ]; then
        echo -e "${YELLOW}Generating HTML coverage report for $PARENT_DIR${NC}"
        go tool cover -html=$COVERAGE_FILE -o coverage_$(basename "$PARENT_DIR").html
        echo -e "${GREEN}HTML coverage report generated: ${NC}coverage_$(basename "$PARENT_DIR").html"
      fi
    else
      echo -e "${RED}Tests failed for $PARENT_DIR${NC}"
      FAILED=1
    fi
    
    # Clean up - remove the copied test files
    for test_file in $dir/*.go; do
      base_name=$(basename "$test_file")
      rm -f "$PARENT_DIR/$base_name"
    done
  else
    echo -e "${YELLOW}No Go files found in $dir${NC}"
  fi
done

# If coverage was enabled, merge all coverage files
if [ "$COVERAGE" = true ]; then
  echo -e "${YELLOW}Merging coverage reports...${NC}"
  echo "mode: set" > coverage.out
  grep -h -v "mode: set" coverage_*.out >> coverage.out
  
  # Display coverage summary
  echo -e "${YELLOW}Coverage summary:${NC}"
  go tool cover -func=coverage.out
  
  # Generate combined HTML report if requested
  if [ "$HTML_COVERAGE" = true ]; then
    echo -e "${YELLOW}Generating combined HTML coverage report${NC}"
    go tool cover -html=coverage.out -o coverage.html
    echo -e "${GREEN}Combined HTML coverage report generated: ${NC}coverage.html"
  fi
fi

if [ "$FAILED" == "1" ]; then
  echo -e "${RED}Some tests failed.${NC}"
  exit 1
else
  echo -e "${GREEN}All tests passed!${NC}"
  exit 0
fi
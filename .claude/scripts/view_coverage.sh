#!/bin/bash
# Script to generate and view test coverage report

echo "Generating test coverage report..."

# Generate coverage data
go test -coverprofile=.claude/coverage/coverage.out -covermode=atomic ./...

# Generate HTML report
go tool cover -html=.claude/coverage/coverage.out -o .claude/coverage/coverage.html

# Display coverage summary
echo ""
echo "Coverage Summary:"
go tool cover -func=.claude/coverage/coverage.out | grep "total:" | awk '{print "Total Coverage: " $3}'

echo ""
echo "Coverage report generated at: .claude/coverage/coverage.html"

# Open in browser if on macOS
if [[ "$OSTYPE" == "darwin"* ]]; then
    echo "Opening coverage report in browser..."
    open .claude/coverage/coverage.html
else
    echo "To view the report, open: .claude/coverage/coverage.html"
fi
PLATFORM ?= linux

.PHONY: check dist clean cover-test cover-web bench help

check:
ifeq ($(OS),Windows_NT)
	go test ./... -short
else
	@which wget >/dev/null || (echo "Error: wget required. Install with 'brew install wget' or 'apt-get install wget'"; exit 1)
	@which bc >/dev/null || (echo "Error: bc required for the coverage gate. Install with 'brew install bc' or 'apt-get install bc'"; exit 1)
	@wget -q -O lint-project.sh https://raw.githubusercontent.com/moov-io/infra/master/go/lint-project.sh
	@chmod +x ./lint-project.sh
	COVER_THRESHOLD=50.0 COVER_EXCLUDE="cmd/" ./lint-project.sh
endif

dist: clean
ifeq ($(OS),Windows_NT)
	CGO_ENABLED=1 GOOS=windows go build -o bin/wire20022.exe github.com/moov-io/wire20022/cmd/wire20022
else
	CGO_ENABLED=1 GOOS=$(PLATFORM) go build -o bin/wire20022-$(PLATFORM)-amd64 github.com/moov-io/wire20022/cmd/wire20022
endif

clean:
	@rm -rf ./bin/ ./tmp/ coverage.txt misspell* staticcheck lint-project.sh

cover-test:
	go test -coverprofile=cover.out $(shell go list ./... | grep -v -E "cmd/")
cover-web:
	go tool cover -html=cover.out

# Help target
help:
	@echo "Available targets:"
	@echo "  check         - Run tests and linters (requires wget and bc)"
	@echo "  dist          - Build binary distribution"
	@echo "  clean         - Remove build artifacts"
	@echo "  cover-test    - Write a coverage profile"
	@echo "  cover-web     - Open the coverage profile in a browser"
	@echo "  bench         - Run XML parse/write benchmarks and print medians"

.PHONY: bench
bench:
	go test ./... -count=6 -cpu=2 -run '^$$' -bench '^(BenchmarkParseXML|BenchmarkWriteXML|BenchmarkCustomerCreditTransfer)$$' -benchmem > output.raw.txt
	python3 scripts/bench-median.py output.raw.txt | tee output.txt

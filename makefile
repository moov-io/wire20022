.PHONY: check
check:
ifeq ($(OS),Windows_NT)
	go test ./... -short
else
	@wget -O lint-project.sh https://raw.githubusercontent.com/moov-io/infra/master/go/lint-project.sh
	@chmod +x ./lint-project.sh
	COVER_THRESHOLD=50.0 DISABLE_GITLEAKS=true ./lint-project.sh
endif

dist: clean
ifeq ($(OS),Windows_NT)
	CGO_ENABLED=1 GOOS=windows go build -o bin/wire20022.exe github.com/moov-io/wire20022/cmd/wire20022
else
	CGO_ENABLED=1 GOOS=$(PLATFORM) go build -o bin/wire20022-$(PLATFORM)-amd64 github.com/moov-io/wire20022/cmd/wire20022
endif

.PHONY: setup
setup:
	docker compose up -d --force-recreate --remove-orphans

.PHONY: teardown
teardown:
	-docker compose down --remove-orphans

.PHONY: clean
clean:
	@rm -rf ./bin/ ./tmp/ coverage.txt misspell* staticcheck lint-project.sh

.PHONY: cover-test cover-web
cover-test:
	go test -coverprofile=cover.out ./...
cover-web:
	go tool cover -html=cover.out
# Define Docker image parameters
DOCKER_IMAGE_NAME ?= wire20022
DOCKER_TAG ?= latest
PLATFORM ?= linux

.PHONY: check docker build-docker push-docker
check:
ifeq ($(OS),Windows_NT)
	go test ./... -short
else
	@which wget >/dev/null || (echo "Error: wget required. Install with 'brew install wget' or 'apt-get install wget'"; exit 1)
	@wget -q -O lint-project.sh https://raw.githubusercontent.com/moov-io/infra/master/go/lint-project.sh
	@chmod +x ./lint-project.sh
	COVER_THRESHOLD=50.0 COVER_EXCLUDE="cmd/" ./lint-project.sh
endif

# Docker targets
docker: build-docker  # Alias for default docker operation

build-docker:
	docker build -t $(DOCKER_IMAGE_NAME):$(DOCKER_TAG) \
	--build-arg VERSION=$(shell git describe --tags) \
	--platform $(PLATFORM)/amd64 .

push-docker:
	docker push $(DOCKER_IMAGE_NAME):$(DOCKER_TAG)

dist: clean
ifeq ($(OS),Windows_NT)
	CGO_ENABLED=1 GOOS=windows go build -o bin/wire20022.exe github.com/moov-io/wire20022/cmd/wire20022
else
	CGO_ENABLED=1 GOOS=$(PLATFORM) go build -o bin/wire20022-$(PLATFORM)-amd64 github.com/moov-io/wire20022/cmd/wire20022
endif

# Docker compose operations
setup:
	docker compose up -d --force-recreate --remove-orphans

teardown:
	-docker compose down --remove-orphans

clean:
	@rm -rf ./bin/ ./tmp/ coverage.txt misspell* staticcheck lint-project.sh

cover-test:
	go test -coverprofile=cover.out $(shell go list ./... | grep -v -E "cmd/")
cover-web:
	go tool cover -html=cover.out

# Help target
help:
	@echo "Available targets:"
	@echo "  docker         - Build Docker image (alias for build-docker)"
	@echo "  build-docker  - Build Docker image with current source"
	@echo "  push-docker   - Push Docker image to registry"
	@echo "  setup         - Start Docker compose services"
	@echo "  teardown      - Stop Docker compose services"
	@echo "  check         - Run tests and linters"
	@echo "  dist          - Build binary distribution"
	@echo "  clean         - Remove build artifacts"

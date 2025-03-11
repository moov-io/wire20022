.PHONY: check
check:
ifeq ($(OS),Windows_NT)
	go test ./... -short
else
	@wget -O lint-project.sh https://raw.githubusercontent.com/moov-io/infra/master/go/lint-project.sh
	@chmod +x ./lint-project.sh
	COVER_THRESHOLD=50.0 DISABLE_GITLEAKS=true ./lint-project.sh
endif

GOBIN ?= $$(go env GOPATH)/bin

.PHONY: install-go-test-coverage
install-go-test-coverage:
	@go install github.com/vladopajic/go-test-coverage/v2@latest

.PHONY: check-coverage
check-coverage:
	@$(MAKE) -s clean
	@go test -timeout 1m -race ./... -coverprofile=./cover.out -covermode=atomic -coverpkg=./...
	${GOBIN}/go-test-coverage --config=./.testcoverage.yml

.PHONY: update-dependencies
update-dependencies:
	@go get -t -u ./... && go mod tidy

.PHONY: format
format:
	goimports -local github.com/josuerosadeavila/stack -w -l .

.PHONY: lint
lint:
	@$(MAKE) format
	@golangci-lint run --allow-parallel-runners ./... --max-same-issues 0

.PHONY: install-golang-ci
install-golang-ci:
	@echo "Installing golang-ci"
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1
	@echo "Golang-ci installed successfully"

.PHONY: install-goimports
install-goimports:
	@echo "Installing go imports"
	@go install golang.org/x/tools/cmd/goimports@latest
	@echo "Go imports installed successfully"

.PHONY: configure
configure:
	make install-golang-ci
	make install-goimports
	make install-go-test-coverage

.PHONY: clean
clean:
	@go clean -testcache

.PHONY: test
test:
	@$(MAKE) -s clean
	@go test -timeout 1m -race ./... -coverprofile=./cover.out.tmp -covermode=atomic -coverpkg=./... ${GOBIN}/go-test-coverage --config=./.testcoverage.yml && cat cover.out.tmp | grep -v example > ./cover.out  && go tool cover -func=cover.out
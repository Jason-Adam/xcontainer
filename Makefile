GO=go
GOTEST=$(GO) test
GOTOOL=$(GO) tool
BIN_DIR=bin

.PHONY: test
test:
	mkdir -p $(BIN_DIR) && \
		$(GOTEST) ./... -cover -coverprofile=$(BIN_DIR)/coverage.out

.PHONY: coverage
coverage: test
	$(GOTOOL) cover -html $(BIN_DIR)/coverage.out -o $(BIN_DIR)/coverage.html

.PHONY: race
race:
	$(GOTEST) ./... -race

.PHONY: benchmark
benchmark:
	$(GOTEST) ./... -bench=. -benchmem

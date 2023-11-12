INPUT_DIR := inputs

clean:
	@go clean -testcache
	@rm -rf bin/*

test:
	@go test -v ./...

build-%:
	@echo "# Building day $*"
	@go build -v -o "bin/day$*" "cmd/day$*/main.go" || exit 1

day-01: build-01
	@echo "# Running day 01"
	@echo -n "Part 1: "
	@cat $(INPUT_DIR)/day01.txt | bin/day01 1
	@echo -n "Part 2: "
	@cat $(INPUT_DIR)/day01.txt | bin/day01 3 | awk '{ sum += $$1 } END { print sum }'

.PHONY: day-% build-%

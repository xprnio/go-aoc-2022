clean:
	@go clean -testcache
	@rm -rf bin/*

test:
	@go test -v ./...

build-%:
	@echo "# Building day $*"
	@go build -v -o "bin/day$*" "cmd/day$*/main.go"

day-01: build-01
	@echo "# Running day 01"
	@echo -n "Part 1: "
	@cat inputs/day01.txt | bin/day01 1
	@echo -n "Part 2: "
	@cat inputs/day01.txt | bin/day01 3 | awk '{ sum += $$1 } END { print sum }'


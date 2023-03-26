all: build

dependency:
	@swag init
	@wire

test-dependency:
	@wire ./test

run: dependency
	@go run .

build: dependency
	@go build -o=bin/server.app .

watch: dependency
	@air --build.cmd="make build" --build.bin="./bin/server.app" --build.exclude_dir="bin,tmp,docs" --build.exclude_file="wire_gen.go" --build.kill_delay="0s"

test: test-dependency
	@go test ./test/... -v

.PHONY: dependency test-dependency

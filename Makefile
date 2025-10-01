run: build
	./bin/prxy

run-test-server: build
	./bin/test_server

run-test:
	@go test -coverpkg=./test/... -coverprofile=./test/coverage.out ./test/...

build:
	@go build -o bin/prxy cmd/main.go

clean:
	rm -rf bin/prxy

.PHONY: clean


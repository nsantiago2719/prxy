run: build
	./bin/prxy

run-test-server: build
	./bin/test_server

build:
	@go build -o bin/prxy cmd/main.go
	@go build -o bin/test_server test/main.go

clean:
	rm -rf bin/prxy

.PHONY: clean


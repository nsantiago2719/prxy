run: build
	./bin/prxy

build:
	@go build -o bin/prxy cmd/main.go

clean:
	rm -rf bin/prxy

.PHONY: clean


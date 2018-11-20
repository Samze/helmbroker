BINARY_NAME=helmbroker

all: build test

build:
	go build -v -o $(BINARY_NAME) cmd/*

clean:
	go clean
	rm -rf $(BINARY_NAME)

run: build
	./$(BINARY_NAME)

test:
	ginkgo -r

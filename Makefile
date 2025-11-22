.PHONY: build test run clean docker-build docker-run

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOCLEAN=$(GOCMD) clean
GOGET=$(GOCMD) get
BINARY_NAME=capital-gains
BINARY_UNIX=$(BINARY_NAME)_unix

all: test build

build:
	$(GOBUILD) -o $(BINARY_NAME) -v ./cmd/capital-gains

test:
	$(GOTEST) -v ./...

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)

run:
	$(GOBUILD) -o $(BINARY_NAME) -v ./cmd/capital-gains
	./$(BINARY_NAME)

docker-build:
	docker build -t capital-gains .

docker-run:
	docker run -i capital-gains

# Cross compilation
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v ./cmd/capital-gains
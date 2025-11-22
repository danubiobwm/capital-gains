APP=capital-gains
BINARY=./$(APP)
PKG=./...

.PHONY: all build test fmt vet clean run e2e docker-build

all: build

build:
	go build -o $(APP) ./cmd/capital-gains

run: build
	# Run with example input
	./$(APP) < test/examples_input.txt

test:
	go test ./...

fmt:
	go fmt ./...

vet:
	go vet ./...

clean:
	rm -f $(APP)

e2e: build
	@echo "Running end-to-end tests..."
	@bash test/e2e_test.sh test/examples_input.txt test/expected_output.txt

docker-build:
	docker build -t capital-gains:latest .

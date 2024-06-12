# Define the Go compiler and flags
GOCMD = go
GOBUILD = $(GOCMD) build
GOCLEAN = $(GOCMD) clean
GOTEST = $(GOCMD) test
GOFMT = gofmt -w
BINARY_NAME = game
SRC_DIRS = . ./src

# Build the project
build:
	$(GOBUILD) -o $(BINARY_NAME) ./components/main.go ./components/sdl_utils.go ./components/game.go ./components/renderer.go ./components/matrix.go

# Run the project
run: build
	./$(BINARY_NAME)

# Clean the project
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

# Format the code
fmt:
	$(GOFMT) $(SRC_DIRS)

# Run tests
test:
	$(GOTEST) -v ./...

# Phony targets to prevent conflicts with file names
.PHONY: build run clean fmt test

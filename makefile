BINARY_NAME=gotion.exe
# Gets the tag, or the short hash if no tag exists. 
# Adds '-dirty' if you have uncommitted changes.
VERSION=$$(git describe --tags --always --dirty)

# Local build
build:
	@go build -ldflags "-X main.Version=$(VERSION)" -o $(BINARY_NAME) .

# The "Proper" Go installation
install: build
	@go install -ldflags "-X main.Version=$(VERSION)" .
	@echo "--------------------------------------------"
	@echo "Successfully installed gotion $(VERSION)"
	@echo "Run 'gotion --version' to verify"
	@echo "--------------------------------------------"

# Run locally for testing
run: build
	@./$(BINARY_NAME)

# Clean local binary
clean:
	@if exist $(BINARY_NAME) del $(BINARY_NAME)
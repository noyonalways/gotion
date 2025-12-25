# Get the binary name from the current directory name
BINARY_NAME=gotion.exe

# Local build
build:
	@go build -o $(BINARY_NAME) .

# The "Proper" Go installation
# This will place the binary in %GOPATH%\bin (usually C:\Users\YourName\go\bin)
install:
	@go install .
	@echo "Application installed to Go bin directory."

# Run locally for testing
run: build
	@./$(BINARY_NAME)

# Clean local binary
clean:
	@if exist $(BINARY_NAME) del $(BINARY_NAME)
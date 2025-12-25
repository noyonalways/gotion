BINARY_NAME=gotion.exe
VERSION := $(shell git describe --tags --always --dirty 2>nul || echo dev)

# 1. Install to the standard Go bin directory
install:
	@go install -ldflags "-X main.Version=$(VERSION)" .
	@echo --------------------------------------------------
	@echo  Successfully installed gotion $(VERSION)
	@echo  Binary location: $(shell go env GOPATH)\bin
	@echo --------------------------------------------------

# 2. Local build for testing
build:
	@go build -ldflags "-X main.Version=$(VERSION)" -o $(BINARY_NAME) .

# 3. Cleanup
clean:
	@if exist $(BINARY_NAME) del $(BINARY_NAME)
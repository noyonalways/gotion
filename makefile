BINARY_NAME=gotion
VERSION := $(shell git describe --tags --always --dirty 2>/dev/null || echo dev)
UNAME_S := $(shell uname -s)

ifeq ($(UNAME_S),Linux)
	BINARY_NAME=gotion
	RM=rm -f
	BIN_DIR=$(shell go env GOPATH)/bin
endif
ifeq ($(UNAME_S),Darwin)
	BINARY_NAME=gotion
	RM=rm -f
	BIN_DIR=$(shell go env GOPATH)/bin
endif
ifeq ($(UNAME_S),MINGW64_NT)
	BINARY_NAME=gotion.exe
	RM=del /QBINARY_NAME=gotion
VERSION := $(shell git describe --tags --always --dirty 2>/dev/null || echo dev)
UNAME_S := $(shell uname -s)

ifeq ($(UNAME_S),Linux)
	BINARY_NAME=gotion
	RM=rm -f
	BIN_DIR=$(shell go env GOPATH)/bin
endif
ifeq ($(UNAME_S),Darwin)
	BINARY_NAME=gotion
	RM=rm -f
	BIN_DIR=$(shell go env GOPATH)/bin
endif
ifeq ($(UNAME_S),MINGW64_NT)
	BINARY_NAME=gotion.exe
	RM=del /Q
	BIN_DIR=$(shell go env GOPATH)/bin
endif

# 1. Install to the standard Go bin directory
install:
	@go install -ldflags "-X main.Version=$(VERSION)" .
	@echo --------------------------------------------------
	@echo Successfully installed $(BINARY_NAME) $(VERSION)
	@echo Binary location: $(BIN_DIR)
	@echo --------------------------------------------------

# 2. Local build for testing
build:
	@go build -ldflags "-X main.Version=$(VERSION)" -o $(BINARY_NAME) .

run: build
	@./$(BINARY_NAME)

# 3. Cleanup
clean:
	@$(RM) $(BINARY_NAME)
	BIN_DIR=$(shell go env GOPATH)/bin
endif

# 1. Install to the standard Go bin directory
install:
	@go install -ldflags "-X main.Version=$(VERSION)" .
	@echo --------------------------------------------------
	@echo Successfully installed $(BINARY_NAME) $(VERSION)
	@echo Binary location: $(BIN_DIR)
	@echo --------------------------------------------------

# 2. Local build for testing
build:
	@go build -ldflags "-X main.Version=$(VERSION)" -o $(BINARY_NAME) .

run: build
	@./$(BINARY_NAME)

# 3. Cleanup
clean:
	@$(RM) $(BINARY_NAME)
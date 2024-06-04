# while saving don't forget to indent tabs with vs code cmd + p option
ROOT_DIR := $(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))
PROTO_DIR := $(ROOT_DIR)/proto

# Determine the operating system
ifeq ($(OS), Windows_NT)
    OS = windows
    SHELL := powershell.exe
    .SHELLFLAGS := -NoProfile -Command
    PACKAGE = $(shell Get-Content go.mod -head 1 | Foreach-Object { $$data = $$_ -split " "; "{0}" -f $$data[1]})
    BIN = practice_go_grpc.exe
else
    UNAME := $(shell uname -s)
    ifeq ($(UNAME), Darwin)
        OS = macos
    else ifeq ($(UNAME), Linux)
        OS = linux
    else
        $(error OS not supported by this Makefile)  # Print detailed error message
    endif
    PACKAGE = $(shell head -1 go.mod | awk '{print $$2}')
    BIN = practice_go_grpc
endif

# Build target
build: generate
	go build -o ${BIN} .

# Generate Go code and gRPC service interface code from .proto files
generate:
	protoc -I=${PROTO_DIR} --go_out=. --go-grpc_out=. ${PROTO_DIR}/*.proto

# Update dependencies
bump: generate
	go get -u ./...

# Clean generated files and binary
clean:
	rm -f ${PROTO_DIR}/*.pb.go
	rm -f ${BIN}

# Help target to display usage instructions
help:
	@echo "Usage: make [target]"
	@echo "Targets:"
	@echo "  build     : Build the binary"
	@echo "  generate  : Generate Go code and gRPC service interface code"
	@echo "  bump      : Update dependencies"
	@echo "  clean     : Clean generated files and binary"
	@echo "  help      : Display this help message"

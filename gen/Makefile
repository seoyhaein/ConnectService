# Makefile for generating Go code from Protocol Buffers using protoc

# Path to the protoc compiler
PROTOC_PATH=/usr/local/bin/protoc

# Path to the Protocol Buffers folder
PROTO_PATH=../protos

# Path to the Protocol Buffers file
PROTO_FILE=$(PROTO_PATH)/jobs.proto

# Directory to save the generated files
OUT_DIR=$(PROTO_PATH)

# Path to the protoc-gen-go plugin
PROTOC_GEN_GO_PATH=/usr/local/bin/protoc-gen-go
PROTOC_GEN_GO_GRPC_PATH=/usr/local/bin/protoc-gen-go-grpc

# Default target
all: generate

# Generate Protocol Buffers code
generate:
	$(PROTOC_PATH) --proto_path=$(PROTO_PATH) --plugin=$(PROTOC_GEN_GO_PATH) --go_out=$(OUT_DIR) --go_opt=paths=source_relative $(PROTO_FILE)
	$(PROTOC_PATH) --proto_path=$(PROTO_PATH) --plugin=$(PROTOC_GEN_GO_GRPC_PATH) --go-grpc_out=$(OUT_DIR) --go-grpc_opt=paths=source_relative $(PROTO_FILE)

# Clean generated files (add the files to be cleaned if needed)
clean:
	rm -f $(OUT_DIR)/*.pb.go

.PHONY: all generate clean
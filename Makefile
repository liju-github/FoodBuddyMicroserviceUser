PROTO_DIR := proto

# Install necessary tools for protobuf compilation
install-tools:
	@echo "Installing necessary tools..."
	# Install protoc-gen-go and protoc-gen-go-grpc
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	@echo "Tools installed successfully."

# Compile proto files
generate-proto:
	@echo "Generating gRPC code for UserService..."
	PATH=$(HOME)/go/bin:$(PATH) protoc -I=$(PROTO_DIR) \
		--go_out=. \
		--go-grpc_out=. \
		$(PROTO_DIR)/**/*.proto

tidy:
	@echo "Downloading dependencies..."
	go mod tidy

run-app:
	go run ./cmd

# Run the entire pipeline
all: install-tools generate-proto tidy run-app
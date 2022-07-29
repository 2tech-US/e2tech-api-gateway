proto:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    internal/**/pb/*.proto
api_gateway:
	go run cmd/main.go

.PHONY: proto server
CURRENT_DIR=$(shell pwd)

proto-gen:
	@./scripts/gen-proto.sh ${CURRENT_DIR}
give-permission:
	@chmod +x scripts/gen-proto.sh
tidy:
	@go mod tidy
	@go mod vendor
run-server:
	@go run server/server.go
run-client:
	@go run client/client.go
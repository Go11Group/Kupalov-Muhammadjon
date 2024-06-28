CURRENT_DIR=$(shell pwd)

proto-gen:
	@./scripts/gen-proto.sh ${CURRENT_DIR}

give-permission:
	@chmod +x scripts/gen-proto.sh

tidy:
	@go mod tidy
	@go mod vendor

get-permission:
	@chmod +x ./scripts/gen-proto.sh

run:
	@go run cmd/main.go

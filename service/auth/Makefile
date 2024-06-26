include .env

PROTOC := protoc
PROTOC_GEN_GO := ../../bin/protoc-gen-go
PROTOC_GEN_GO_GRPC := ../../bin/protoc-gen-go-grpc
PROTOC_GEN_GRPC_GATEWAY := ../../bin/protoc-gen-grpc-gateway
PROTOC_GEN_VALIDATE := ../../bin/protoc-gen-validate
PROTOC_GEN_OPENAPIV2 := ../../bin/protoc-gen-openapiv2
GOOSE := ../../bin/goose

PROTO_SRC_DIR := api/proto/v1
PROTO_VENDOR_DIR := ../../vendor.protogen
SWAGGER_DST_DIR := ../../pkg/swagger

lint:
	golangci-lint run ./... --config ../../.golangci.pipeline.yaml

run:
	go run ./cmd/app/main.go

generate-api:
	mkdir -p $(SWAGGER_DST_DIR)
	make generate-api-access
	make generate-api-auth
	make generate-api-user

generate-api-auth:
	mkdir -p ../../pkg/grpc/auth_v1
	$(PROTOC) --proto_path=$(PROTO_SRC_DIR) --proto_path=$(PROTO_VENDOR_DIR) \
		--go_out=../../pkg/grpc/auth_v1 --go_opt=paths=source_relative \
		--plugin=protoc-gen-go=$(PROTOC_GEN_GO) \
		--go-grpc_out=../../pkg/grpc/auth_v1 --go-grpc_opt=paths=source_relative \
		--plugin=protoc-gen-go-grpc=$(PROTOC_GEN_GO_GRPC) \
		--grpc-gateway_out=../../pkg/grpc/auth_v1 --grpc-gateway_opt=paths=source_relative \
		--plugin=protoc-gen-grpc-gateway=$(PROTOC_GEN_GRPC_GATEWAY) \
		--validate_out lang=go:../../pkg/grpc/auth_v1 --validate_opt=paths=source_relative \
		--plugin=protoc-gen-validate=$(PROTOC_GEN_VALIDATE) \
		--openapiv2_out=allow_merge=true,merge_file_name=api_auth_v1:$(SWAGGER_DST_DIR) \
		--plugin=protoc-gen-openapiv2=$(PROTOC_GEN_OPENAPIV2) \
		$(PROTO_SRC_DIR)/auth.proto

generate-api-access:
	mkdir -p ../../pkg/grpc/access_v1
	$(PROTOC) --proto_path=$(PROTO_SRC_DIR) --proto_path=$(PROTO_VENDOR_DIR) \
		--go_out=../../pkg/grpc/access_v1 --go_opt=paths=source_relative \
		--plugin=protoc-gen-go=$(PROTOC_GEN_GO) \
		--go-grpc_out=../../pkg/grpc/access_v1 --go-grpc_opt=paths=source_relative \
		--plugin=protoc-gen-go-grpc=$(PROTOC_GEN_GO_GRPC) \
		--grpc-gateway_out=../../pkg/grpc/access_v1 --grpc-gateway_opt=paths=source_relative \
		--plugin=protoc-gen-grpc-gateway=$(PROTOC_GEN_GRPC_GATEWAY) \
		--openapiv2_out=allow_merge=true,merge_file_name=api_access_v1:$(SWAGGER_DST_DIR) \
		--plugin=protoc-gen-openapiv2=$(PROTOC_GEN_OPENAPIV2) \
		$(PROTO_SRC_DIR)/access.proto

generate-api-user:
	mkdir -p ../../pkg/grpc/user_v1
	$(PROTOC) --proto_path=$(PROTO_SRC_DIR) --proto_path=$(PROTO_VENDOR_DIR) \
		--go_out=../../pkg/grpc/user_v1 --go_opt=paths=source_relative \
		--plugin=protoc-gen-go=$(PROTOC_GEN_GO) \
		--go-grpc_out=../../pkg/grpc/user_v1 --go-grpc_opt=paths=source_relative \
		--plugin=protoc-gen-go-grpc=$(PROTOC_GEN_GO_GRPC) \
		--grpc-gateway_out=../../pkg/grpc/user_v1 --grpc-gateway_opt=paths=source_relative \
		--plugin=protoc-gen-grpc-gateway=$(PROTOC_GEN_GRPC_GATEWAY) \
		--validate_out lang=go:../../pkg/grpc/user_v1 --validate_opt=paths=source_relative \
		--plugin=protoc-gen-validate=$(PROTOC_GEN_VALIDATE) \
		--openapiv2_out=allow_merge=true,merge_file_name=api_user_v1:$(SWAGGER_DST_DIR) \
		--plugin=protoc-gen-openapiv2=$(PROTOC_GEN_OPENAPIV2) \
		$(PROTO_SRC_DIR)/user.proto

migrations-generate:
	$(GOOSE) -dir migrations/postgres create ${name} sql

migrations-up:
	GOOSE_DBSTRING=postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable $(GOOSE) -dir migrations/postgres postgres up

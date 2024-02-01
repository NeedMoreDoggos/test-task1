include .env

LOCAL_BIN:=$(CURDIR)/bin

LOCAL_MIGRATION_DIR=$(MIGRATION_DIR)
LOCAL_MIGRATION_DSN="host=$(HOST) port=$(PG_PORT) dbname=$(PG_DATABASE_NAME) user=$(PG_USER) password=$(PG_PASSWORD) sslmode=disable"


install-deps:
	mkdir -p $(LOCAL_BIN)
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.31.0
	GOBIN=$(LOCAL_BIN) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2.0
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@v3.14.0

get-deps:
	go mod tidy
	go mod download


docker-build-and-push:
	docker build --no-cache --platform linux/amd64 -t docker.io/needmoredoggos/testing:latest .
	docker push docker.io/needmoredoggos/testing:latest

local-migration-status:
	${LOCAL_BIN}/goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} status -v

local-migration-up:
	${LOCAL_BIN}/goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} up -v

local-migration-down:
	${LOCAL_BIN}/goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} down -v

sqlite-migration-testing-up:
	${LOCAL_BIN}/goose -dir ${LOCAL_MIGRATION_DIR} sqlite3 ./data/testing.db up -v

new-migrations:
	cd ${LOCAL_MIGRATION_DIR} && \
	${LOCAL_BIN}/goose create add_wallet_fk_to_transactions sql

test:
	go clean -testcache
	go test ./... -count=10
	
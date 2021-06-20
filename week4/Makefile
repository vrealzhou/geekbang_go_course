ENV ?= local
GOBUILD=env CGO_ENABLED=0 go build
VERSION = v1

.PHONY: all core rest gql grpc

all: core rest gql

core:
	$(info [*] Building core service)
	$(call wire,core,$(VERSION),$(ENV))
	$(call build_grpc,core,$(VERSION),$(ENV))

rest:
	$(info [*] Building rest facet service)
	# inject dependencies
	go build cmd/rest_facet/*.go

gql:
	$(info [*] Building graphql facet service)
	# inject dependencies
	go build cmd/gql_facet/*.go

grpc:
	cd pkg/todo/v1 && protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative todo.proto

start:
	$(info [*] Starting supporting services)
	docker-compose -f deploy/docker/local-compose.yaml up -d
	$(call wire,core,v1,$(ENV))
	wire gen -tags="$(ENV)" github.com/vrealzhou/geekbang_go_course/week4/cmd/core/$(VERSION)/init
	go run cmd/core/$(VERSION)/init/wire_gen.go

stop:
	$(info [*] Stopping supporting services)
	docker-compose -f deploy/docker/local-compose.yaml down

# inject dependencies
define wire
	wire gen -tags="$(3)" github.com/vrealzhou/geekbang_go_course/week4/cmd/$(1)/$(2)/initialize
endef

define build_grpc
	mkdir -p bin/$(1)/$(2) && \
	$(GOBUILD) -tags="$(3)" -o bin/$(1)/$(2)/server cmd/$(1)/$(2)/server/*.go && \
	$(GOBUILD) -tags="$(3)" -o bin/$(1)/$(2)/client cmd/$(1)/$(2)/client/*.go
endef
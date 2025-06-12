dev_config := $(wildcard env/.api.dev)

.PHONY: run
run:
	@env $(shell cat $(dev_config)) go run cmd/main.go

.PHONY: test
test:
	@env -S "`cat $(dev_config)`" go test -v -bench -coverpkg=./internal/...,./cmd/... -coverprofile=coverage.cov ./... -json -v | tparse


.PHONY: generate
generate: swagger.yaml
	rm -rf internal/api/generated/models
	rm -rf internal/api/generated/restapi/operations
	rm -rf internal/api/generated/restapi/doc.go
	rm -rf internal/api/generated/restapi/embedded_spec.go
	rm -rf internal/api/generated/restapi/server.go
	swagger generate server -f swagger.yaml -t ./internal/api/generated --main-package ../../../../cmd/ -A parcelmeadow

.PHONY: generate-mocks
generate-mocks:
	go install go.uber.org/mock/mockgen
	mockgen -destination ./internal/services/mock/generated/service.go -source ./internal/services/service.go
	mockgen -destination ./internal/api/handlers/mock/generated/handlers.go -source ./internal/api/handlers/handlers.go
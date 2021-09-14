SRC=$(shell find . -name "*.go")
.PHONY: fmt install_deps run_dev run_test run_prod

install_deps:
	$(info ******************** downloading dependencies ********************)
	go get -v ./...

run_devserver:
	$(info ******************** running development server ********************)
	NEWSMUX_ENV=dev go run ./cmd/server/main.go -dev=true -no_auth -service=api_server

run_prodserver:
	$(info ******************** running production server ********************)
	NEWSMUX_ENV=prod go run ./cmd/server/main.go -dev=false -no_auth -service=api_server

run_prodpublisher:
	$(info ******************** running publisher server ********************)
	NEWSMUX_ENV=prod go run ./cmd/publisher/main.go -dev=false -service=feed_publisher

run_devpublisher:
	$(info ******************** running publisher server ********************)
	NEWSMUX_ENV=dev go run ./cmd/publisher/main.go -dev=true -service=feed_publisher

fmt:
	$(info ******************** checking formatting ********************)
	@test -z $(shell gofmt -l $(SRC)) || (gofmt -d $(SRC); exit 1)

test:
	NEWSMUX_ENV=test go test ./...

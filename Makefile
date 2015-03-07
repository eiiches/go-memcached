.PHONY: all
all: memcached memcached-cli

flags := -race
sources := $(shell find src -type f -name \*.go) 
gopath := $(shell pwd):$(shell echo $$GOPATH)

memcached: export GOPATH = $(gopath)
memcached: $(sources) memcached.go
	go generate -tags generate -v memcached/server
	go build $(flags) -o memcached memcached.go

memcached-cli: export GOPATH = $(gopath)
memcached-cli: $(sources) memcached-cli.go
	go generate -tags generate -v memcached/server
	go build $(flags) -o memcached-cli memcached-cli.go

.PHONY: clean
clean:
	$(RM) src/memcached/server/commands.go src/memcached/server/proto_binary_commands.go memcached memcached-cli

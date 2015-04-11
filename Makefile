.PHONY: rebuild
rebuild:
	make clean
	make all

.PHONY: all
all: memcached memcached-cli

flags := -race -x
sources := $(shell find src -type f -name \*.go) 
gopath := $(shell pwd):$(shell echo $$GOPATH)

memcached: export GOPATH = $(gopath)
memcached: $(sources) memcached.go
	go generate -tags generate -v memcached
	go build $(flags) -o memcached memcached.go

memcached-cli: export GOPATH = $(gopath)
memcached-cli: $(sources) memcached-cli.go
	go generate -tags generate -v memcached
	go build $(flags) -o memcached-cli memcached-cli.go

.PHONY: clean
clean:
	$(RM) src/memcached/iface.go src/memcached/commands.go src/memcached/server_protocol_binary.go memcached memcached-cli

mix: export GOPATH = $(gopath)
mix: $(sources) mix.go
	go generate -tags generate -v memcached
	go build $(flags) -o mix mix.go

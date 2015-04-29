.PHONY: all generate install clean

null :=

build_flags := -x -gcflags '-m'
build_flags := -x

targets := \
	bin/memcached-cli \
	bin/memcached-server \
	bin/memcached-bench \
	$(null)

sources := \
	memcached/client.go \
	memcached/commands_impl_client.go \
	memcached/core/atomic.go \
	memcached/core/lrucache.go \
	memcached/errors.go \
	memcached/protocol_binary.go \
	memcached/server.go \
	memcached/server_impl.go \
	$(null)

templates := \
	memcached/iface.go.in \
	memcached/client_protocol_binary.go.in \
	memcached/server_protocol_binary.go.in \
	$(null)

generated_sources := $(templates:.go.in=.go)

all: $(targets)

generate: $(generated_sources)

install: $(generated_sources) $(sources)
	go install $(build_flags) github.com/eiiches/go-memcached/memcached

clean:
	$(RM) $(generated_sources) $(targets)

$(generated_sources): $(templates) memcached/@.go memcached/@.yml
	go generate -tags generate -v github.com/eiiches/go-memcached/memcached

bin/%: src/%.go
	go build $(build_flags) -o $@ $<

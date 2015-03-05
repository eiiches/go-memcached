memcached: export GOPATH = $(shell pwd):$(shell echo $$GOPATH)
memcached: $(shell find -type f -name *.go)
	go generate -tags generate -v memcached
	go build

.PHONY: clean
clean:
	$(RM) src/memcached/commands.go src/memcached/proto_binary_commands.go

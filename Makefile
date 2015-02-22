memcached: export GOPATH = $(shell pwd):$(shell echo $$GOPATH)
memcached: $(shell find -type f -name *.go)
	go generate -tags generate -v memcached
	go build

// +build !generate

package memcached

import "os"
import "fmt"

func (self setCommand) execute(server *MemcachedServer) {
	if self.cas != 0 {
		server.cache.CompareAndSet(self.key, self.value, self.cas, self.expire)
	} else {
		server.cache.Put(self.key, self.value, self.expire)
	}
	fmt.Fprintf(os.Stderr, "PUT %v = %v [cas: %v]\n", string(self.key), string(self.value), self.cas)
}

func (self addCommand) execute(server *MemcachedServer) {
}

func (self getCommand) execute(server *MemcachedServer) {
	value, cas := server.cache.Get(self.key)
	fmt.Fprintf(os.Stderr, "GET %v = %v [cas: %v]\n", string(self.key), string(value), cas)
}

func (self replaceCommand) execute(server *MemcachedServer) {
	if self.cas != 0 {
		server.cache.CompareAndSet(self.key, self.value, self.cas, self.expire)
	} else {
		server.cache.Replace(self.key, self.value, self.expire)
	}
}

func (self deleteCommand) execute(server *MemcachedServer) {
	if self.cas != 0 {
		server.cache.CompareAndRemove(self.key, self.cas)
	} else {
		server.cache.Remove(self.key)
	}
}

func (self incrementCommand) execute(server *MemcachedServer) {
	if self.cas != 0 {
	} else {
	}
}

func (self decrementCommand) execute(server *MemcachedServer) {
}

func (self quitCommand) execute(server *MemcachedServer) {
}

func (self flushCommand) execute(server *MemcachedServer) {
	server.cache.Clear()
}

func (self versionCommand) execute(server *MemcachedServer) {
}

func (self nopCommand) execute(server *MemcachedServer) {
}

func (self getWithKeyCommand) execute(server *MemcachedServer) {
	value, cas := server.cache.Get(self.key)
	fmt.Fprintf(os.Stderr, "GETK %v = %v [cas: %v]\n", string(self.key), string(value), cas)
}

func (self appendCommand) execute(server *MemcachedServer) {
}

func (self prependCommand) execute(server *MemcachedServer) {
}

func (self statCommand) execute(server *MemcachedServer) {
}

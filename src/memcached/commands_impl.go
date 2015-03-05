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
	fmt.Fprintf(os.Stderr, "PUT %v = %v [cas: %v]\n\r", string(self.key), string(self.value), self.cas)
	os.Stderr.Sync()
}

func (self addCommand) execute(server *MemcachedServer) {
}

func (self getCommand) execute(server *MemcachedServer) {
	value, cas := server.cache.Get(self.key)
	fmt.Print("GET", string(self.key), "=", string(value), "[cas:", cas, "]")
}

func (self replaceCommand) execute(server *MemcachedServer) {
}

func (self deleteCommand) execute(server *MemcachedServer) {
}

func (self incrementCommand) execute(server *MemcachedServer) {
}

func (self decrementCommand) execute(server *MemcachedServer) {
}

func (self quitCommand) execute(server *MemcachedServer) {
}

func (self flushCommand) execute(server *MemcachedServer) {
}

func (self versionCommand) execute(server *MemcachedServer) {
}

func (self nopCommand) execute(server *MemcachedServer) {
}

func (self getWithKeyCommand) execute(server *MemcachedServer) {
}

func (self appendCommand) execute(server *MemcachedServer) {
}

func (self prependCommand) execute(server *MemcachedServer) {
}

func (self statCommand) execute(server *MemcachedServer) {
}

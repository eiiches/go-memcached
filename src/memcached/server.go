// +build !generate

package memcached

import "os"
import "net"
import "fmt"
import "sync"
import "memcached/core"

type MemcachedServer struct {
	shutdownRequested core.AtomicBoolean
	cache             core.LruCache
	socks             []net.Listener
}

func NewMemcachedServer() *MemcachedServer {
	return &MemcachedServer{
		cache: core.NewLruCache(10 * 1024 * 1024),
		socks: make([]net.Listener, 0),
	}
}

func (self *MemcachedServer) Serve() {
	handler := newBinaryProtocolHandler()
	num_accept_threads := 4

	wg := sync.WaitGroup{}
	wg.Add(num_accept_threads * len(self.socks))
	for i := 0; i < num_accept_threads; i++ {
		for _, sock := range self.socks {
			go func() {
				defer wg.Done()
				for {
					conn, err := sock.Accept()
					if err != nil {
						if self.shutdownRequested.Get() {
							return
						}
						fmt.Fprintf(os.Stderr, "error: %+v\n", err)
						continue
					}
					go func() {
						err := handler.handleConnection(conn, self)
						conn.Close()
						if err != nil {
							fmt.Fprintf(os.Stderr, "error: %+v\n", err)
						}
					}()
				}
			}()
		}
	}
	wg.Wait()
}

func (self *MemcachedServer) Listen(proto string, laddr string) error {
	sock, err := net.Listen(proto, laddr)
	if err != nil {
		return err
	}
	self.socks = append(self.socks, sock)
	return nil
}

func (self *MemcachedServer) Shutdown() error {
	self.shutdownRequested.Set(true)
	for _, sock := range self.socks {
		sock.Close()
	}
	return nil
}

func (self *MemcachedServer) Call(c serverCommand) {
	c.execute(self)
}

func (self *MemcachedServer) Multi(cs ...serverCommand) {
}

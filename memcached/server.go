// +build !generate

package memcached

import "os"
import "net"
import "fmt"
import "sync"
import "github.com/eiiches/go-memcached/memcached/core"

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

type protocolHandler interface {
	handleConnection(conn net.Conn, server *MemcachedServer) error
}

func (self *MemcachedServer) serveLoop(sock net.Listener, handler protocolHandler) {
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
}

func (self *MemcachedServer) Serve() {
	handler := newBinaryProtocolHandler()
	acceptThreads := 1

	wg := sync.WaitGroup{}
	wg.Add(acceptThreads * len(self.socks))
	for i := 0; i < acceptThreads; i++ {
		for _, sock := range self.socks {
			go func(sock net.Listener) {
				defer wg.Done()
				self.serveLoop(sock, handler)
			}(sock)
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

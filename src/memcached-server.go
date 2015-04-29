package main

import "github.com/eiiches/go-memcached/memcached"
import "fmt"
import "os"
import "os/signal"
import "runtime"

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	server := memcached.NewMemcachedServer()

	if err := server.Listen("tcp", "0.0.0.0:11212"); err != nil {
		fmt.Fprintf(os.Stderr, "error: %+v\n", err)
		server.Shutdown()
		os.Exit(1)
	}

	if err := server.Listen("unix", "/tmp/memcached.sock"); err != nil {
		fmt.Fprintf(os.Stderr, "error: %+v\n", err)
		server.Shutdown()
		os.Exit(1)
	}

	sigch := make(chan os.Signal, 1)
	signal.Notify(sigch, os.Interrupt)
	go func() {
		for sig := range sigch {
			if sig == os.Interrupt {
				fmt.Println("Interrupted")
				server.Shutdown()
			}
		}
	}()

	server.Serve()

	// server.Call(memcached.Set([]byte("key"), []byte("value")).WithExpire(10).WithCas(1234))
	// server.Call(memcached.Add([]byte("key"), []byte("value")).WithExpire(10))
	// server.Multi(memcached.Add([]byte("key"), []byte("value")))
}

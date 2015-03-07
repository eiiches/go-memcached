package main

import "fmt"
import "os"
import "time"
import "os/signal"
import "runtime"
import "memcached"

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	server := memcached.NewMemcachedServer()

	if err := server.Listen("tcp", "0.0.0.0:11212"); err != nil {
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

	go func() {
		time.Sleep(10000)

		client, err := memcached.NewMemcachedClient("tcp", "localhost:11212")
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %+v\n", err)
			return
		}
		defer client.Close()

		client.Call(memcached.Set([]byte("key"), []byte("value")).WithExpire(10))
	}()

	server.Serve()
}

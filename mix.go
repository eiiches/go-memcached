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

		client.Set("key", "value", memcached.Set.WithExpire(10))
		client.Set("key", "value").WithExpire(10).Call()

		resp1, err := client.Set("key", "value", &memcached.SetOptions{
			Expire: 10,
		})

		resp2, err := client.Get("key", &memcached.GetOptions{
			Expire: 10,
		})

		resps, err := client.Batch([]memcached.Request{
			memcached.NewSetRequest([]byte("key"), []byte("value"), &memcached.SetOptions{
				Expire: 10,
			}),
			memcached.NewGetRequest([]byte("key"), nil),
		})

		resp, err := client.Call(memcached.NewSetRequest([]byte("key"), []byte("value"), &memcached.SetOptions{
			Expire: 10,
		}))
	}()

	server.Serve()
}

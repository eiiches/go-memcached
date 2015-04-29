package main

import "os"
import "strings"
import "fmt"
import "github.com/eiiches/go-memcached/memcached"
import "sync"
import "flag"

var (
	concurrency int
	requests    int
	address     string
)

func main() {
	flag.IntVar(&concurrency, "concurrency", 10, "the number of simultaneous connections to use")
	flag.IntVar(&requests, "requests", 10, "the number of requests per connection")
	flag.StringVar(&address, "address", "localhost:11211", "address to memcached-server, such as localhost:11211 or /tmp/memcached.sock.")
	flag.Parse()

	proto := "tcp"
	if strings.Contains(address, "/") {
		proto = "unix"
	}

	var wg sync.WaitGroup
	wg.Add(concurrency)
	for c := 0; c < concurrency; c++ {
		go func() {
			client, err := memcached.NewMemcachedClient(proto, address)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error: %+v\n", err)
				return
			}
			defer client.Close()
			for i := 0; i < requests; i++ {
				client.Set([]byte("hoge"), []byte("10"), nil)
				client.Get([]byte("hoge"), nil)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

package main

import "os"
import "fmt"
import "github.com/eiiches/go-memcached/memcached"
import "sync"
import "flag"

var (
	concurrency int
	requests    int
	host        string
	port        int
)

func main() {
	flag.IntVar(&concurrency, "concurrency", 10, "the number of simultaneous connections to use")
	flag.IntVar(&requests, "requests", 10, "the number of requests per connection")
	flag.StringVar(&host, "host", "localhost", "hostname")
	flag.IntVar(&port, "port", 11211, "port")
	flag.Parse()

	var wg sync.WaitGroup
	wg.Add(concurrency)
	for c := 0; c < concurrency; c++ {
		go func() {
			client, err := memcached.NewMemcachedClient("tcp", fmt.Sprintf("%s:%d", host, port))
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

package main

import "os"
import "fmt"
import "github.com/eiiches/go-memcached/memcached"
import "sync"

func main() {
	concurrency := 4
	requests := 1000
	var wg sync.WaitGroup
	wg.Add(concurrency)
	for c := 0; c < concurrency; c++ {
		go func() {
			client, err := memcached.NewMemcachedClient("tcp", "localhost:11212")
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

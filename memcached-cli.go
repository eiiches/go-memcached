package main

import "os"
import "fmt"
import "memcached"

func main() {
	client, err := memcached.NewMemcachedClient("tcp", "localhost:11212")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %+v\n", err)
		return
	}
	defer client.Close()

	client.Call(memcached.Set([]byte("key"), []byte("value")).WithExpire(10))
}

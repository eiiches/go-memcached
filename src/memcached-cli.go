package main

import "os"
import "fmt"
import "github.com/eiiches/go-memcached/memcached"

func main() {
	client, err := memcached.NewMemcachedClient("tcp", "localhost:11212")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %+v\n", err)
		return
	}
	defer client.Close()

	client.Set([]byte("hoge"), []byte("10"), nil)
	value, _, _, _ := client.Get([]byte("hoge"), nil)
	fmt.Println(value)
}

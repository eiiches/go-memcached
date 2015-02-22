package core

import "fmt"
import "runtime"
import "sync"
import "testing"
import "time"

func TestGet(t *testing.T) {
	runtime.GOMAXPROCS(4)

	m := NewLruCache()
	m.Get([]byte("test"))
	m.Get([]byte("test2"))
	m.Get([]byte("test9"))
	m.Get([]byte("test3"))
	m.Get([]byte("test4"))
	m.Get([]byte("hoge"))
	m.Get([]byte("fuga"))
	m.Get([]byte("piyo"))

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("m.Get ->")
			value := m.Get([]byte("key"))
			fmt.Println("<- m.Get #=>", string(value))
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("m.Put ->")
			m.Put([]byte("key"), []byte("value"))
			fmt.Println("<- m.Put")
		}
		wg.Done()
	}()
	wg.Wait()
	return
}

func TestMemStats(t *testing.T) {
	if testing.Short() {
		t.Skip("skip")
	}
	a := make([]byte, 1024*1024*100)
	a[0] = 1
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%+v\n", m)
	time.Sleep(100)
}

package main

import "github.com/eiiches/go-memcached/memcached"
import "fmt"
import "os"
import "os/signal"
import "runtime"
import "runtime/pprof"
import "flag"

var (
	pprofFile string
	tcpSocket string
	udsFile   string
)

func main() {
	flag.StringVar(&pprofFile, "pprof", "", "write pprof output to file")
	flag.StringVar(&tcpSocket, "listen-tcp", "localhost:11211", "tcp socket")
	flag.StringVar(&udsFile, "listen-unix", "/tmp/memcached.sock", "unix domain socket path")
	flag.Parse()

	if pprofFile != "" {
		f, err := os.Create(pprofFile)
		if err != nil {
			panic(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	runtime.GOMAXPROCS(runtime.NumCPU())

	server := memcached.NewMemcachedServer()

	if err := server.Listen("tcp", tcpSocket); err != nil {
		fmt.Fprintf(os.Stderr, "error: %+v\n", err)
		server.Shutdown()
		os.Exit(1)
	}

	if err := server.Listen("unix", udsFile); err != nil {
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
}

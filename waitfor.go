package main

import (
	"fmt"
	"net"
	"sync"
	"time"
	goopt "github.com/droundy/goopt"
)

var timeout = goopt.Int([]string{"-t", "--timeout"}, 60, "timeout")

func main() {
	goopt.Description = func() string {
		return "Wait for endpoints to become available."
	}
	goopt.Version = "1.0"
	goopt.Summary = "waitforit [-t N] host:port [host:port [...]]"
	goopt.Parse(nil)

	fmt.Printf("Waiting for %d seconds.\n", *timeout)

	var wg sync.WaitGroup

	for _, hostport := range goopt.Args {
		wg.Add(1)
		go func(){
			defer wg.Done()
			pingTCP(hostport, *timeout)
		}()
	}

	wg.Wait()
}

func pingTCP(hostport string, timeoutSeconds int) {
	timeout := time.Duration(timeoutSeconds) * time.Second
	start := time.Now()
	nextTry := time.Now()

	for {
		_, err := net.DialTimeout("tcp", hostport, time.Second)

		if err == nil {
			fmt.Printf("%s up!\n", hostport)
			return
		}

		fmt.Printf("%s down: %v\n", hostport, err)
		if time.Since(start) > timeout {
			panic(err)
		}

		nextTry = nextTry.Add(time.Second)
		time.Sleep(nextTry.Sub(time.Now()))
	}
}

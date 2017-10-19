package main

import "time"
import "fmt"

func main() {

	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// This `limiter` channel will receive a value
	// every 1 second. This is the regulator in
	// our rate limiting scheme.
	limiter := time.Tick(time.Second)

	// By blocking on a receive from the `limiter` channel
	// before serving each request, we limit ourselves to
	// 1 request every 1 second.
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}
}

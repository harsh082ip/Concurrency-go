package main

import (
	"fmt"
	"time"
)

/*
	Let's suppose we have a server, and in that some go-routines
	are supposed to be run indefinetly, but lets assume there are
	some other go-routines that were created in a process and are not
	cancelled till now, this can fed up our server by increasing the load
	on the ram, to cancel the child go-routines from the parent process,
	one such approch to do this is by following a DONE-CHANNEL PATTERN
*/

func doSomeWork(done <-chan bool) { // a receive-only channel

	for {
		select {
		case <-done:
			// fmt.Println(msg)
			return // stop this go-routine when any msg is received from the channel
		default:
			time.Sleep(time.Second / 2)
			fmt.Println("Doing Some Work...")
		}
	}
}

func main() {

	done := make(chan bool)

	go doSomeWork(done)

	time.Sleep(time.Second * 5)

	close(done)
}

package main

import "fmt"

/*
	Let's suppose we want, to perform different operations on the same
	set of data, so there are two ways, doing it all together, or
	doing all this thing in stages, and that where PIPELINE CONCURRENCY PATTERN comes in
*/

func sliceToChannel(nums []int) <-chan int { // return a receive only channel
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()

	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func main() {

	nums := []int{1, 2, 3, 4, 5}

	// stage 1 - load all this slice data into a channel
	dataChannel := sliceToChannel(nums)

	// stage 2 - take that channel values, are square all of them, and return a new channel
	finalChannel := sq(dataChannel)

	// stage 3 - just print the final channel recieved

	for val := range finalChannel {
		fmt.Println(val)
	}
}

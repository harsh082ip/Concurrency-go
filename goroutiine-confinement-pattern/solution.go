package goroutiineconfinementpattern

import (
	"fmt"
	"sync"
	"time"
)

/*

	The concurrency pattern used in this code is known as Goroutine Confinement.
	In this pattern, each goroutine is responsible for handling its own piece of data,
	thus avoiding shared state between goroutines and the need for explicit
	synchronization (except for waiting for all goroutines to complete).

	Here's a breakdown of how this pattern is applied in the given code:

	Goroutine for Each Task: The code spawns a new goroutine for each element
	in the input slice to process the data concurrently.

	WaitGroup for Synchronization: A sync.WaitGroup is used to wait for all goroutines
	to complete their execution before proceeding. This is necessary to ensure that the
	main function does not exit before all goroutines have finished their work.

	Confinement of Data: Each goroutine processes its own data and stores the result
	in a specific index of the result slice. This confines the data and processing to
	the individual goroutine, avoiding the need for complex synchronization mechanisms
	like mutexes for shared data.


*/

func process(data int) int {
	time.Sleep(time.Second)
	return data * 2
}

func processData(wg *sync.WaitGroup, res *int, data int) {
	defer wg.Done()

	processData := process(data)

	*res = processData

}

func Solution() {
	start := time.Now()

	var wg sync.WaitGroup

	input := []int{1, 2, 3, 4, 5}
	result := make([]int, len(input))

	for i, data := range input {
		wg.Add(1)
		go processData(&wg, &result[i], data)
	}

	wg.Wait()

	fmt.Println(time.Since(start))

	fmt.Println(result)
}

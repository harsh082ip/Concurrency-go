package goroutiineconfinementpattern

import (
	"fmt"
	"sync"
	"time"
)

var lock sync.Mutex

// PROBLEM

/* The problem here in this approach is that,
   mutex acts a bottle neck, so if the critical section,
   takes more time, then othrt go routines have to wait very long
*/

func Process(data int) int {
	time.Sleep(time.Second)
	return data * 2
}

func ProcessData(wg *sync.WaitGroup, res *[]int, data int) {
	defer wg.Done()

	processData := Process(data)

	lock.Lock()
	*res = append(*res, processData) // De-referencing a pointer
	lock.Unlock()
}

func Problem() {

	var wg sync.WaitGroup

	input := []int{1, 2, 3, 4, 5}
	result := []int{}

	for _, data := range input {
		wg.Add(1)
		go ProcessData(&wg, &result, data)
	}

	wg.Wait()

	fmt.Println(result)
}

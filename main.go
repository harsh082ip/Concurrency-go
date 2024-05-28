package main

import (
	"fmt"
	"log"
)

// "time"

// confinement "github.com/harsh082ip/Concurrency-go/goroutiine-confinement-pattern"

type Val interface {
	int | int64 | uint32 | string | bool
}

func PrintVal[T Val](values T) {
	var res T
	println(res)
	fmt.Println(res)
	log.Println(res)
}

func main() {

	// Confinement Pattern

	// for {
	// 	time.Sleep(time.Second / 2)
	// 	// confinement.Problem()
	// 	confinement.Solution()
	// }
	PrintVal(true)
	// PrintVal(12)
	println(-23)
}

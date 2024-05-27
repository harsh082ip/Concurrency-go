package main

import (
	"time"

	confinement "github.com/harsh082ip/Concurrency-go/goroutiine-confinement-pattern"
)

func main() {

	// Confinement Pattern

	for {
		time.Sleep(time.Second / 2)
		// confinement.Problem()
		confinement.Solution()
	}
}

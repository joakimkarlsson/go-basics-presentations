package main

import (
	"fmt"
	"math/rand"
	"time"
)

// START OMIT
func numbers() chan int {
	ch := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
		close(ch)
	}()

	return ch
}

func main() {
	for v := range numbers() {
		fmt.Println(v)
	}
}

// END OMIT

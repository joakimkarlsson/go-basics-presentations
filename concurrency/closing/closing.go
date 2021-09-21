package main

import (
	"fmt"
	"math/rand"
	"time"
)

// START OMIT
func numbers(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	}
	close(ch) // HL
}

func main() {
	ch := make(chan int)
	go numbers(ch)
	for v := range ch { // HL
		fmt.Println(v) // HL
	} // HL
}

// END OMIT

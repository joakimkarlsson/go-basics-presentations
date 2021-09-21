package main

import (
	"fmt"
	"math/rand"
	"time"
)

// START OMIT
func numbers(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i // HL
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	}
}

func main() {
	ch := make(chan int)
	go numbers(ch)
	for i := 0; i < 5; i++ {
		v := <-ch // HL
		fmt.Println(v)
	}
}

// END OMIT

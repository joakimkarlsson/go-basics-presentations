package main

import (
	"fmt"
	"math/rand"
	"time"
)

// START OMIT
func numbers() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	}
}

func main() {
	numbers()
}

// END OMIT

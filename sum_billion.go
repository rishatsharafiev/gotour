package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	sum := 0
	for i := 0; i < 1000000000; i++ {
		sum += i
	}
	fmt.Println("Hello, 世界")
	duration := time.Since(now)
	fmt.Println("Sum: ", sum)
	fmt.Println("Delta: ", duration)
}

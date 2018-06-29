package main

import (
	"fmt"
	"time"
	"math"
)

func testSimultaneous(k int) (int, time.Duration) {
	n := float64(k)
	startTime := time.Now()
	sum := 0
	for i := 0; i < int(n); i++ {
		sum += i
	}

	return sum, time.Since(startTime)
}

func summaryRoutine(from int, to int, c chan int) {
	sum := 0
	for i := from; i < to; i++ {
		sum += i
	}
	c <- sum
}

func testParallel(k int) (int, time.Duration)  {
	n := float64(k)
	startTime := time.Now()
	sum := 0

	for j := 0; j < 1; j++ {
		step := 25.0
		size := int(math.Ceil(n / step))
		c := make(chan int)
		var from, to int

		for i := 0; i < int(step); i++ {
			from = i * size
			to = (i + 1) * size
			if i == int(step)-1 {
				to = int(n)
			}
			go summaryRoutine(from, to, c)
		}

		for i := 0; i < int(step); i++ {
			sum += <-c
		}
	}

	return sum, time.Since(startTime)
}

func main() {
	k := 15000000000

	s1, d1 := testSimultaneous(k)
	fmt.Println("testSimultaneous(sum, duration): ", s1, d1)

	s2, d2 := testParallel(k)
	fmt.Println("testConcurency(sum, duration): ", s2, d2)

}

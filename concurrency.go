package main

import (
	"time"
	"fmt"
)

// (1/11) Go-процедуры
func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

// (2/11) Каналы
func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

// (4/11) Range и Close
func fibonacci2(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

// (5/11) Select
func fibonacci3(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	// (1/11) Go-процедуры
	fmt.Println("----------------------")

	go say("world")
	say("hello")

	// (2/11) каналы
	fmt.Println("----------------------")

	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)

	// (3/11) Буферизированные каналы
	fmt.Println("----------------------")

	ch := make(chan int, 2) // если поставить 1, то получим deadlock
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	// (4/11) Range и Close
	fmt.Println("----------------------")

	c2 := make(chan int, 10)
	go fibonacci2(cap(c2), c2)
	for i := range c2 { // range ждет пока канал закроется и выполняет блок цикла, если приходят данные
		fmt.Println(i)
	}

	// (5/11) Select
	fmt.Println("----------------------")

	c3 := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c3)
		}
		quit <- 0
	}()
	fibonacci3(c3, quit)

	// (6/11) Блок по умолчанию в Select
	fmt.Println("----------------------")

	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
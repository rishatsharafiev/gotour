package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

// (5/14) if
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

// (6-7/14) if с временной переменной в краткой инструкции
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}

// (8/14)
func Round(x float64, unit float64) float64 {
	division := math.Pow(10, unit)
	return math.Ceil(x*division) / division
}

func Sqrt(x float64) float64 {
	z := 0.5
	k := 0.0
	decimals := 2.0
	for i := 0; i < 8000; i++ {
		k = z - (z*z-x)/2*z
		if Round(k, decimals) == Round(z, decimals) {
			fmt.Printf("iteration: %d\n", i)
			break
		}
		z = k
	}
	return z
}

func main() {
	// (1/14)
	// блок инициализации: выполняется перед первой итерацией
	// условный блок: выполняется перед каждой итерацией
	// завершающий блок: выполняется в конце каждой итерации
	// (2/14)
	// блоки инициализации и завершения опциональны.
	sum := 0
	for i := 0; sum < 1000; i++ {
		sum += i
	}
	fmt.Println(sum)

	// (3/14) цикл while выглядит как оператор for с одним условием
	sum = 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	// (4/14) бесконечный цикл
	for {
		break
	}

	// (5/14)
	fmt.Println(sqrt(2), sqrt(-4))

	// (6-7/14)
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	// (8/14)
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))

	// (9/14) switch case
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
		fallthrough
	case "something1":
		fmt.Println("Something1.")
	case "linux":
		fmt.Println("Linux.")
		fallthrough
	case "something2":
		fmt.Println("Something2.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}

	// (10/14)
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	// (11/14) if-then-else
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

	// (12/14) Defer откладывает выполнение функции до возврата из окружающей функции, аргументы вычисляются сразу,
	// но вызов происходит после возврата из функции
	defer fmt.Println("world")
	fmt.Println("hello")

	// (13/14) отложенные вызовы накапливаются в стек LIFO
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

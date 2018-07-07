package main

import (
	"math"
	"fmt"
)

// (1/26) методы
type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs () float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsSecond (v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	// (1/26) методы
	v := Vertex{X: 3, Y: 4}
	fmt.Println(v.Abs())
	fmt.Println(AbsSecond(v))
}

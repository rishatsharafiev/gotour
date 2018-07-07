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

// (2/26) методы и функции
func AbsSecond (v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// (3/26) методы, продложение
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// (4/26) получатели и указатели
func (v* Vertex) AbsThird() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v* Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// (5/26) указатели и функции
func AbsForth(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func ScaleForth(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// (6/26) методы и косвенная адресация указателей
func (v *Vertex) ScaleFifth(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFuncFifth(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}


func main() {
	// (1/26) методы
	v := Vertex{X: 3, Y: 4}
	fmt.Println(v.Abs())

	// (2/26) методы и функции
	fmt.Println(AbsSecond(v))

	// (3/26) методы, продложение
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())

	// (4/26) получатели и указатели
	v.Scale(10)
	fmt.Println(v.AbsThird())

	// (5/26) указатели и функции
	v = Vertex{3, 4}
	ScaleForth(&v, 10)
	fmt.Println(AbsForth(v))

	// (6/26) методы и косвенная адресация указателей
	v = Vertex{3, 4}
	v.ScaleFifth(2)
	ScaleFuncFifth(&v, 10)

	p := &Vertex{4, 3}
	p.ScaleFifth(3)
	ScaleFuncFifth(p, 8)

	fmt.Println(v, p)
}
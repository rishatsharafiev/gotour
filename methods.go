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

// (7/26) методы и косвенная адресация указателей (2)
func (v Vertex) AbsSixth() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFuncSixth(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// (8/26) выбирая между получателем как значение или указателем
func (v *Vertex) ScaleSeventh(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex) AbsSeventh() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// (9/26) интерфейсы
type Abser interface {
	Abs2() float64
}

type MyFloat2 float64

func (f MyFloat2) Abs2() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex2 struct {
	X, Y float64
}

func (v *Vertex2) Abs2() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// (10/26) интерфейсы реализуются неявно
type I interface {
	M()
}

type T struct {
	S string
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t T) M() {
	fmt.Println(t.S)
}

// (11/26) интерфейсные значения
type I2 interface {
	M()
}

type T2 struct {
	S string
}

func (t *T2) M() {
	fmt.Println(t.S)
}

type F2 float64

func (f F2) M() {
	fmt.Println(f)
}

func describe(i I2) {
	fmt.Printf("(%v, %T)\n", i, i)
}

// (12/26) интерфейсные значения с нижележащим типом nil
type I3 interface {
	M()
}

type T3 struct {
	S string
}

func (t *T3) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

// (13/26) интерфейсное значение nil
type I4 interface {
	M()
}

// (14/26) пустой интерфейс
func describe5(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

// (16/26) switch с типами
func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

// (17/26) Stringers
type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	// (1/26) методы
	fmt.Println("-------")

	v := Vertex{X: 3, Y: 4}
	fmt.Println(v.Abs())

	// (2/26) методы и функции
	fmt.Println("-------")

	fmt.Println(AbsSecond(v))

	// (3/26) методы, продложение
	fmt.Println("-------")

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())

	// (4/26) получатели и указатели
	fmt.Println("-------")

	v.Scale(10)
	fmt.Println(v.AbsThird())

	// (5/26) указатели и функции
	fmt.Println("-------")

	v = Vertex{3, 4}
	ScaleForth(&v, 10)
	fmt.Println(AbsForth(v))

	// (6/26) методы и косвенная адресация указателей
	fmt.Println("-------")

	v = Vertex{3, 4}
	v.ScaleFifth(2)
	ScaleFuncFifth(&v, 10)

	p := &Vertex{4, 3}
	p.ScaleFifth(3)
	ScaleFuncFifth(p, 8)

	fmt.Println(v, p)

	// (7/26) методы и косвенная адресация указателей (2)
	fmt.Println("-------")

	v = Vertex{3, 4}
	fmt.Println(v.AbsSixth())
	fmt.Println(AbsFuncSixth(v))

	p = &Vertex{4, 3}
	fmt.Println(p.AbsSixth())
	fmt.Println(AbsFuncSixth(*p))

	// (8/26) выбирая между получателем как значение или указателем
	fmt.Println("-------")

	v = Vertex{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.AbsSeventh())
	v.ScaleSeventh(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.AbsSeventh())

	// (9/26) интерфейсы
	fmt.Println("-------")

	var a2 Abser
	f2 := MyFloat2(-math.Sqrt2)
	v2 := Vertex2{3, 4}

	a2 = f2  // a MyFloat implements Abser
	a2 = &v2 // a *Vertex implements Abser

	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser.
	//a2 = v2

	fmt.Println(a2.Abs2())

	// (10/26) интерфейсы реализуются неявно
	fmt.Println("-------")

	var i I = T{"hello"}
	i.M()

	// (11/26) интерфейсные значения
	fmt.Println("-------")

	i = &T2{"Hello"}
	describe(i)
	i.M()

	i = F2(math.Pi)
	describe(i)
	i.M()

	// (12/26) интерфейсные значения с нижележащим типом nil
	fmt.Println("-------")

	var i3 I3

	var t3 *T3
	i3 = t3
	describe(i3)
	i3.M()

	i3 = &T3{"hello"}
	describe(i3)
	i3.M()

	// (13/26) интерфейсное значение nil
	fmt.Println("-------")

	var i4 I4
	describe(i4)
	//defer func() {
	//	// recover from panic if one occured. Set err to nil otherwise.
	//	if (recover() != nil) {
	//		_ = errors.New("nil error")
	//	}
	//}()

	//i4.M()

	// (14/26) пустой интерфейс
	fmt.Println("-------")

	var i5 interface{}
	describe5(i5)

	i5 = 42
	describe5(i5)

	i5 = "hello"
	describe5(i5)


	// (15/26) утверждение типа
	fmt.Println("-------")

	var i6 interface{} = "hello"

	s6 := i6.(string)
	fmt.Println(s6)

	s6, ok6 := i6.(string)
	fmt.Println(s6, ok6)

	f6, ok6 := i6.(float64)
	fmt.Println(f6, ok6)

	//f6 = i6.(float64) // panic
	//fmt.Println(f6)

	// (16/26) switch с типами
	fmt.Println("-------")
	do(21)
	do("hello")
	do(true)

	// (17/26) Stringers
	fmt.Println("-------")
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
}
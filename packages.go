// (1/17) программа начинает работу в пакете main
package main

// (2/17) импорт нескольких пакетов объединяется в круглые скобки
import (
	"fmt"
	"math"
	"math/cmplx"
)

// (4/17) x, y - аргумент целового типа, возращается целый тип, (5/17) если два и более последовательных аргумента имеют
// одинаковый тип, то можно пропустить объявление типа, кроме последнего x, y int
func add(x int, y int) int {
	return x + y
}

// (6/17) функция может мовзращать множественные результаты
func swap(x, y string) (string, string) {
	return y, x
}

// (7/17) именованные возвращаемые значения ("голый возврат" ухудшают читаемость, подходит для коротких функций)
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

// (8-9/17) объявление переменных с инициализацией значений, тип указывается последним
var i, j int = 1, 2

// (11/17) базовые типы данных
/*
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // псевдоним для uint8

rune // псевдоним для int32
// представляет Unicode код

float32 float64

complex64 complex128
*/

// (16/17) нетипизированная константа принимает тип, необходимый в ее контексте.
const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	// (3/17) math.Sqrt(7) - в Go имя экспортируется, если начинается с заглавной буквы
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))

	// (4/17)
	fmt.Println(add(42, 13))

	// (6/17)
	a, b := swap("hello", "world")
	fmt.Println(a, b)

	// (7/17)
	fmt.Println(split(17))

	// (8-9/17) оператор var может быть использован на уровне пакета или функции
	// Если инициализирующее значение присутствует, то тип может быть опущен; переменная получит тип этого значения.
	var c, python, java = true, false, "no!"
	// (10/17) краткая форма объявления переменной, не импользуется вне функции
	k := 3
	fmt.Println(i, j, k, c, python, java)

	// (11/17)
	var (
		ToBe   bool       = false
		MaxInt uint64     = 1<<64 - 1
		z      complex128 = cmplx.Sqrt(-5 + 12i)
	)
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	// (12/17) переменные объявленные без указания начального значения получают нулевое значение.
	// - 0 для числовых типов,
	// - false для булевого типа, и
	// - "" (пустая строка) для строк.
	var i int
	var f float64
	var h bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, h, s)

	// (13/17) приведение типов
	var x, y int = 3, 4
	var l float64 = math.Sqrt(float64(x*x + y*y))
	var m uint = uint(f)
	fmt.Println(x, l, m)

	// (14/17) при объявлении переменной без явного указания типа (с помощью := или var =), тип переменной выводится
	// из значения на правой стороне.
	v := 42 // change me!
	fmt.Printf("v is of type %T\n", v)

	// (15/17) константы не могут быть объявленные с помощью синтаксиса :=.
	const Pi = 3.14
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)

	// (16/17)
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}

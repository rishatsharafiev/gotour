package main

import (
	"fmt"
	"golang.org/x/tour/pic"
	"golang.org/x/tour/wc"
	"math"
	"strings"
)

func main() {
	// (1/27) указатели
	fmt.Println("-------")

	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j

	// (2-3/27) структура это коллекция полей
	fmt.Println("-------")

	type Vertex struct {
		X int
		Y int
	}
	fmt.Println(Vertex{1, 2})
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)

	// (4/27) указатели на структуры
	fmt.Println("-------")

	vk := Vertex{1, 2}
	p2 := &vk
	p2.X = 1e9 // тоже самое, что v2.X = 1e9
	fmt.Println(vk)

	// (5/27) литералы структур
	fmt.Println("-------")

	var (
		v1 = Vertex{1, 2}  // has type Vertex
		v2 = Vertex{X: 1}  // Y:0 is implicit
		v3 = Vertex{}      // X:0 and Y:0
		p1 = &Vertex{1, 2} // has type *Vertex
	)
	fmt.Println(v1, p1, v2, v3)

	// (6/27) массивы [n]T, где n - неизменяемый размер массива, T - тип данных
	fmt.Println("-------")

	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1]) // доступ по индексам к значению элементов массива
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13} // литерал массива
	fmt.Println(primes)

	// (7/27) срезы []T, где T - тип среза
	fmt.Println("-------")

	primes2 := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes2[1:4] // срез на основе массива
	fmt.Println(s)

	// (8/27) срез не хранит данные, а является ссылкой на секцию нижележащего массива, модификация элементов среза
	// или массива приводит к изменениям, которые увидят другие срезы данного массива
	fmt.Println("-------")

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	g := names[0:2]
	t := names[1:3]
	fmt.Println(g, t)

	g[0] = "XXX"
	fmt.Println(g, t)
	fmt.Println(names)

	// (9/27) литералы срезов такой же как литерал массива, только без размера
	fmt.Println("-------")

	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	y := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	} // срез структур
	fmt.Println(y)

	// (10/27) значения по умолчанию в срезах
	fmt.Println("-------")

	d := []int{2, 3, 5, 7, 11, 13}

	d = d[1:4]
	fmt.Println(d)

	d = d[:2]
	fmt.Println(d)

	d = d[1:]
	fmt.Println(d)

	// (11/27) размер и вместимость среза
	// Срез имеет размер (длину) и вместимость.
	// Размер среза - это количество элементов, которые он содержит.
	// Вместимость среза - это количество элементов в его нижележащем массиве, начиная с первого элемента в срезе.
	// Размер и вместимость среза s могут быть получены с помощью len(s) и cap(s).
	fmt.Println("-------")

	k := []int{2, 3, 5, 7, 11, 13}
	printSlice(k)

	// Slice the slice to give it zero length.
	k = k[:0]
	printSlice(k)

	// Extend its length.
	k = k[:4]
	printSlice(k)

	// Drop its first two values.
	k = k[2:]
	printSlice(k)

	// (12/27) нулевые срезы
	fmt.Println("-------")

	var e []int
	fmt.Println(e, len(e), cap(e))
	if e == nil {
		fmt.Println("nil!")
	}

	// (13/27) создание среза с помощью make
	fmt.Println("-------")

	az := make([]int, 5) // создает массив с длиной 5
	printDynamicSlice("a", az)

	bz := make([]int, 0, 5) // создает массив с вместимостью 5
	printDynamicSlice("b", bz)
	bz = bz[:cap(bz)] // len(b)=5, cap(b)=5
	bz = bz[1:]       // len(b)=4, cap(b)=4

	cz := bz[:2]
	printDynamicSlice("c", cz)

	dz := cz[2:cap(cz)]
	printDynamicSlice("d", dz)

	// (14/27) срезы срезов (многомерные срезы)
	fmt.Println("-------")

	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	// (15/27)
	fmt.Println("-------")

	// Если нижележащий массив среза s слишком мал, чтобы вместить все значения, то будет создан новый массив
	// большего размера. Результирующий срез будет ссылаться на этот новый массив.
	var s2 = make([]int, 0, 2)
	printSlice(s2)

	// append works on nil slices.
	s2 = append(s2, 0)
	printSlice(s2)

	// The slice grows as needed.
	s2 = append(s2, 1)
	printSlice(s2)

	// We can add more than one element at a time.
	s2 = append(s2, 2, 3, 4) // создается новый массив, если срез превышает вместимость (cap)
	printSlice(s2)

	// (16/27) итерация по срезу с помощью for и range
	fmt.Println("-------")

	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow { // i - индекс, v - копия значения по индексу i
		fmt.Printf("2**%d = %d\n", i, v)
	}

	// (17/27) _ опустить значение, мождно опустить value в цикле с перебором массива
	fmt.Println("-------")

	pow2 := make([]int, 10)
	for i := range pow2 {
		pow2[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow2 {
		fmt.Printf("%d\n", value)
	}

	// (18/27) упражнение: срезы
	fmt.Println("-------")

	arr := Pic(4, 4)
	for _, emb := range arr {
		for _, value := range emb {
			fmt.Printf("%d ", value)
		}
		fmt.Println("")
	}
	pic.Show(Pic)

	// (19/27) карты (словари)
	fmt.Println("-------")

	type Vertex2 struct {
		Lat, Long float64
	}

	var mapping map[string]Vertex2

	mapping = make(map[string]Vertex2)
	mapping["Bell Labs"] = Vertex2{
		40.68433, -74.39967,
	}
	fmt.Println(mapping["Bell Labs"])

	// (20/27) литералы карт
	fmt.Println("-------")

	var mapping2 = map[string]Vertex2{
		"Bell Labs": Vertex2{
			40.68433, -74.39967,
		},
		"Google": Vertex2{
			37.42202, -122.08408,
		},
	}
	fmt.Println(mapping2)

	// (21/27) литералы карт, продолжение
	fmt.Println("-------")

	var mapping3 = map[string]Vertex2{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}
	fmt.Println(mapping3)

	// (22/27) изменение карт
	fmt.Println("-------")

	mapping4 := make(map[string]int)

	mapping4["Answer"] = 42 // добавление/изменение элемента
	fmt.Println("The value:", mapping4["Answer"])

	mapping4["Answer"] = 48 // добавление/изменение элемента
	fmt.Println("The value:", mapping4["Answer"])

	delete(mapping4, "Answer") // удалить элемент
	fmt.Println("The value:", mapping4["Answer"])

	val, ok := mapping4["Answer"] // проверка	 существования элемента
	fmt.Println("The value:", val, "Present?", ok)

	// (23/27) упражнение: карты
	fmt.Println("-------")

	wc.Test(WordCount)

	// (24/27) функция как значение
	fmt.Println("-------")

	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))

	// (25/27) замыкания
	fmt.Println("-------")

	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}

	// (26/27) упражнение: замыкание Фибоначчи
	fmt.Println("-------")
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

// (18/27) упражнение: срезы
func Pic(dx, dy int) [][]uint8 {
	m := make([][]uint8, dx, dx)
	for x := 0; x < len(m); x++ {
		m[x] = make([]uint8, dy, dy)
		for y := 0; y < len(m[x]); y++ {
			m[x][y] = uint8((x + y) / 2)
		}
	}
	return m
}

// (23/27) упражнение: карты
func WordCount(s string) map[string]int {
	results := make(map[string]int)
	for _, v := range strings.Fields(s) {
		if val, ok := results[v]; ok {
			results[v] = val + 1
		} else {
			results[v] = 1
		}
	}

	return results
}

// (24/27) функция как значение
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

// (25/27) замыкания
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

// (26/27) упражнение: замыкание Фибоначчи
func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func printDynamicSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

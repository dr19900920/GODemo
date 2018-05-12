package main

import "fmt"

func main() {
	var a int
	fmt.Println("a = ", a)
	//	var b,c int
	a = 10
	fmt.Println("a = ", a)

	c := 30
	fmt.Printf("c type is %T\n", c)

	d, e, f := 10, 20, 30
	d, f = f, d
	fmt.Printf("d = %d, e = %d, f = %d\n", d, e, f)

	var tmp int
	tmp, _ = d, e
	fmt.Printf("tmp = %d\n", tmp)

	var g, h, i int
	g, h, i = test()
	fmt.Printf("g = %d, h = %d, i = %d\n", g, h, i)

	const j = 10
	fmt.Printf("j type is %T\n", j)
	fmt.Printf("j = %d\n", j)

	var (
		k int
		l float64
	)

	k, l = 20, 3.14
	fmt.Printf("k = %d, l = %f\n", k, l)

	const (
		m = 10
		n = 10.22
	)
	fmt.Printf("m = %d, n = %f\n", m, n)
	test1()
	test2()
	MyFunc01(5, 6, 7)
}

func test1() {
	const (
		a = iota
		b = iota
		c = iota
	)
	fmt.Printf("a = %d, b = %d, c = %d\n", a, b, c)
	const d = iota
	fmt.Printf("d = %d\n", d)

	const (
		a1 = iota
		b1
		c1
	)
	fmt.Printf("a1 = %d, b1 = %d, c1 = %d\n", a1, b1, c1)

	const (
		a2         = iota
		b2, b3, b4 = iota, iota, iota
		c2         = iota
	)
	fmt.Printf("a2 = %d, b2 = %d, b3 = %d, b4 = %d, c2 = %d\n", a2, b2, b3, b4, c2)
}

func test2() {
	type bigint int64
	var a bigint
	fmt.Printf("a type is %T\n", a)
	if b := 10; b == 10 {
		fmt.Printf("b = %d\n", b)
	}
	score := 85
	switch {
	case score > 90:
		fmt.Print("优秀\n")
	default:
		fmt.Printf("haha\n")
	}
	sum := 0
	for i := 1; i < 3; i++ {
		sum += i
	}
	fmt.Printf("sum = %d\n", sum)

	str := "abc"
	for i, data := range str {
		fmt.Printf("str[%d]=%c\n", i, data)
	}

	for i := range str {
		fmt.Printf("str[%d]=%c\n", i, str[i])
		if i == 0 {
			goto End
		}

	}
End:
	fmt.Printf("aaa\n")
}

func test() (a, b, c int) {
	return 1, 2, 3
}

type FunType func(...int) // 别名可以实现多态
type FunType1 func(int, int) int

func MyFunc01(args ...int) {

	fmt.Println("len(args) = ", len(args))
	MyFunc02(args[:2]...)
	MyFunc02(args[2:]...)

	var fTest FunType
	fTest = MyFunc02
	fTest(1, 2, 4)
}

func MyFunc02(args ...int) {

	fmt.Println("len(args) = ", len(args))

	a := Calc(1, 1, add)
	fmt.Println("a = ", a)
}

// 回调函数
func Calc(a, b int, fTest FunType1) (result int) {

	fmt.Println("calc")
	result = fTest(a, b)
	return
}

func add(a, b int) int {
	return a + b
}

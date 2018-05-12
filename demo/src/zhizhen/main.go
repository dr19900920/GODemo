package main

import "fmt"
import "math/rand"
import "time"

func main() {
	test2()
	test3()
}

func test() {
	a := 10
	fmt.Printf("a = %d\n", a)
	fmt.Printf("address = %p\n", &a)
	var p *int
	p = &a
	fmt.Printf("p address = %p\n", p)
	*p = 666
	fmt.Printf("*p = %v, a = %v", *p, a)
}

func test1() {
	p := new(int)
	*p = 77
	fmt.Printf("p = %v\n", *p)
}

func test2() {
	var id [50]int
	for i := 0; i < len(id); i++ {
		id[i] = i + 1
	}

	// 形参数组是实参数组的复制品
	b := [5]int{0, 1, 2, 3, 4}

	c := [5]int{2: 1, 3: 4}

	// 形参切片是实参切片的引用
	e := c[2:5]

	slice := make([]int, 4)

	fmt.Println("b[1] =", b[1])
	fmt.Println("c[1] =", c[1])
	fmt.Println("c[2] =", c[2])
	fmt.Println("e[0] =", e[0])
	fmt.Println("slice[2] =", slice[2])

}

func test3() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		fmt.Println(rand.Intn(100))
	}
}

func test4() {
	// 类型 长度 容量
	//	s := make([]int, 5, 10)

}

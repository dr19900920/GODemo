package main

import "fmt"
import "os"

func main() {
	//	test1()
	//	test3()
	//	test4()
	//	test5()
}

func test5() {
	list := os.Args
	n := len(list)
	fmt.Println("n = ", n)
	for i := 0; i < n; i++ {
		fmt.Printf("list[%d] = %s\n", i, list[i])
	}
}

func test4() {
	fmt.Println("aaaaaaaaaaaa")
	defer fmt.Println("bbbbbbbbbbbbbbb")  // 函数结束前调用 延迟调用
	defer fmt.Println("dddddddddddddddd") // 先进后出
	fmt.Println("ccccccccccccccc")
}

func test3() {
	f := test2()
	//闭包不关心这些捕获了的变量是否超出作用域
	//只要有闭包在使用它,这些变量还会存在
	fmt.Println("result = ", f())
	fmt.Println("result = ", f())
	fmt.Println("result = ", f())
	fmt.Println("result = ", f())
	fmt.Println("result = ", f())
}

func test2() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func test1() {
	a := 10
	str := "mike"
	f1 := func() {
		fmt.Println("a = ", a)
		fmt.Println("str = ", str)
	}
	f1()
	type FuncType func()
	var f2 FuncType
	f2 = f1
	f2()

	func() {
		fmt.Println("str len = ", len(str))
	}()

	f3 := func(i, j int) {
		fmt.Println("i = ", i)
		fmt.Println("j = ", j)
	}
	f3(1, 2)

	func(i, j int) {
		str = "hahaha"
		fmt.Println("i = ", i)
		fmt.Println("j = ", j)
	}(10, 29)

	fmt.Println("str len = ", len(str))

	x, y := func(i, j int) (max, min int) {
		fmt.Println("i = ", i)
		fmt.Println("j = ", j)
		if i > j {
			max = i
			min = j
		} else {
			max = j
			min = i
		}
		return
	}(10, 29)
	fmt.Println("max = ", x)
	fmt.Println("min = ", y)
}

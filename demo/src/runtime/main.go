package main

import (
	"fmt"
	"runtime"
	"time"
)

// 定义一个打印机
func Printer(str string) {
	for _, data := range str {
		fmt.Printf("%c", data)
		time.Sleep(time.Second)
	}
	fmt.Printf("\n")
}

func main() {

	// 新建2个协程,代表两个人,2个人同时使用打印机 发生资源竞争
	go person1()
	go person2()

	for {

	}
}

func person1() {
	Printer("hello")
}

func person2() {
	Printer("world")
}

func test3() {
	//	n := runtime.GOMAXPROCS(1) // 指定单核运算
	n := runtime.GOMAXPROCS(4)
	fmt.Println(n)

	for {
		go fmt.Print(1)
		fmt.Print(0)
	}
}

func test1() {
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("go")
		}
	}()

	for i := 0; i < 2; i++ {
		// 让出cpu时间片
		runtime.Gosched()
		fmt.Println("hello")
	}

}

func test() {
	defer fmt.Println("ccccc")
	runtime.Goexit() //中止协程
	fmt.Println("dddddd")
}

func test2() {

	go func() {
		fmt.Println("aaaa")
		test()
		fmt.Println("bbbbb")
	}()

	for {

	}
}
